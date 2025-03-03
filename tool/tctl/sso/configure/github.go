// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configure

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/gravitational/trace"
	"github.com/sirupsen/logrus"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/auth"
	"github.com/gravitational/teleport/lib/utils"
)

type ghExtraFlags struct {
	connectorName      string
	ignoreMissingRoles bool
}

func addGithubCommand(cmd *SSOConfigureCommand) *AuthKindCommand {
	spec := types.GithubConnectorSpecV3{}

	gh := &ghExtraFlags{}

	sub := cmd.ConfigureCmd.Command("github", "Configure GitHub auth connector.")
	// commonly used flags
	sub.Flag("name", "Connector name.").Default("github").Short('n').StringVar(&gh.connectorName)
	sub.Flag("teams-to-roles", "Sets teams-to-roles mapping using format 'organization,name,role1,role2,...'. Repeatable.").
		Short('r').
		Required().
		PlaceHolder("org,team,role1,role2,...").
		SetValue(newTeamsToRolesParser(&spec.TeamsToRoles))
	sub.Flag("display", "Sets the connector display name.").StringVar(&spec.Display)
	sub.Flag("id", "GitHub app client ID.").PlaceHolder("ID").Required().StringVar(&spec.ClientID)
	sub.Flag("secret", "GitHub app client secret.").Required().PlaceHolder("SECRET").StringVar(&spec.ClientSecret)
	sub.Flag("endpoint-url", "Endpoint URL for GitHub instance.").StringVar(&spec.EndpointURL)
	sub.Flag("api-endpoint-url", "API endpoint URL for GitHub instance.").StringVar(&spec.APIEndpointURL)

	// auto
	sub.Flag("redirect-url", "Authorization callback URL.").PlaceHolder("URL").StringVar(&spec.RedirectURL)

	// ignores
	sub.Flag("ignore-missing-roles", "Ignore missing roles referenced in --teams-to-roles.").BoolVar(&gh.ignoreMissingRoles)

	sub.Alias("gh")

	sub.Alias(`
Examples:

  > tctl sso configure gh -r octocats,admin,access,editor,auditor -r octocats,dev,access --secret GH_SECRET --id CLIENT_ID

  Generate GitHub auth connector. Two role mappings are defined:
    - members of 'admin' team in 'octocats' org will receive 'access', 'editor' and 'auditor' roles.
    - members of 'dev' team in 'octocats' org will receive 'access' role.

  The values for --secret and --id are provided by GitHub.

  > tctl sso configure gh ... | tctl sso test
  
  Generate the configuration and immediately test it using "tctl sso test" command.`)

	preset := &AuthKindCommand{
		Run: func(ctx context.Context, clt auth.ClientI) error { return ghRunFunc(ctx, cmd, &spec, gh, clt) },
	}

	sub.Action(func(ctx *kingpin.ParseContext) error {
		preset.Parsed = true
		return nil
	})

	return preset
}

func ghRunFunc(ctx context.Context, cmd *SSOConfigureCommand, spec *types.GithubConnectorSpecV3, flags *ghExtraFlags, clt auth.ClientI) error {
	if err := specCheckRoles(ctx, cmd.Logger, spec, flags.ignoreMissingRoles, clt); err != nil {
		return trace.Wrap(err)
	}

	if spec.RedirectURL == "" {
		spec.RedirectURL = ResolveCallbackURL(cmd.Logger, clt, "RedirectURL", "https://%v/v1/webapi/github/callback")
	}

	connector, err := types.NewGithubConnector(flags.connectorName, *spec)
	if err != nil {
		return trace.Wrap(err)
	}
	return trace.Wrap(utils.WriteYAML(os.Stdout, connector))
}

// ResolveCallbackURL deals with common pattern of resolving callback URL for IdP to use.
func ResolveCallbackURL(logger *logrus.Entry, clt auth.ClientI, fieldName string, callbackPattern string) string {
	var callbackURL string

	logger.Infof("%v empty, resolving automatically.", fieldName)
	proxies, err := clt.GetProxies()
	if err != nil {
		logger.WithError(err).Error("unable to get proxy list.")
	}

	// find first proxy with public addr
	for _, proxy := range proxies {
		publicAddr := proxy.GetPublicAddr()
		if publicAddr != "" {
			callbackURL = fmt.Sprintf(callbackPattern, publicAddr)
			break
		}
	}

	// check if successfully set.
	if callbackURL == "" {
		logger.Warnf("Unable to fill %v automatically, cluster's public address unknown.", fieldName)
	} else {
		logger.Infof("%v set to %q", fieldName, callbackURL)
	}
	return callbackURL
}

func specCheckRoles(ctx context.Context, logger *logrus.Entry, spec *types.GithubConnectorSpecV3, ignoreMissingRoles bool, clt auth.ClientI) error {
	allRoles, err := clt.GetRoles(ctx)
	if err != nil {
		logger.WithError(err).Warn("Unable to get roles list. Skipping teams-to-roles sanity checks.")
		return nil
	}

	roleMap := map[string]struct{}{}
	roleNames := make([]string, 0, len(allRoles))
	for _, role := range allRoles {
		roleMap[role.GetName()] = struct{}{}
		roleNames = append(roleNames, role.GetName())
	}

	for _, mapping := range spec.TeamsToRoles {
		for _, role := range mapping.Roles {
			_, found := roleMap[role]
			if !found {
				if ignoreMissingRoles {
					logger.Warnf("teams-to-roles references non-existing role: %q. Available roles: %v.", role, roleNames)
				} else {
					return trace.BadParameter("teams-to-roles references non-existing role: %v. Correct the mapping, or add --ignore-missing-roles to ignore this error. Available roles: %v.", role, roleNames)
				}
			}
		}
	}

	return nil
}
