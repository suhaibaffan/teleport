/*
Copyright 2019-2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import React from 'react';
import { Card, ButtonPrimary, Flex, Box } from 'design';
import * as Alerts from 'design/Alert';

import useAttempt from 'shared/hooks/useAttempt';
import createMfaOptions, { MfaOption } from 'shared/utils/createMfaOptions';

import { Auth2faType, PreferredMfaType } from 'shared/services';

import FieldInput from '../FieldInput';
import FieldSelect from '../FieldSelect';
import Validation, { Validator } from '../Validation';
import {
  requiredToken,
  requiredPassword,
  requiredField,
  requiredConfirmedPassword,
} from '../Validation/rules';

function FormPassword(props: Props) {
  const {
    onChangePassWithWebauthn,
    onChangePass,
    auth2faType = 'off',
    preferredMfaType,
  } = props;
  const mfaEnabled = auth2faType !== 'off';

  const [attempt, attemptActions] = useAttempt({});
  const [token, setToken] = React.useState('');
  const [oldPass, setOldPass] = React.useState('');
  const [newPass, setNewPass] = React.useState('');
  const [newPassConfirmed, setNewPassConfirmed] = React.useState('');
  const mfaOptions = React.useMemo<MfaOption[]>(
    () =>
      createMfaOptions({
        auth2faType: auth2faType,
        preferredType: preferredMfaType,
      }),
    []
  );
  const [mfaType, setMfaType] = React.useState(mfaOptions[0]);

  const { isProcessing } = attempt;

  function submit() {
    switch (mfaType?.value) {
      case 'webauthn':
        return onChangePassWithWebauthn(oldPass, newPass);
      default:
        return onChangePass(oldPass, newPass, token);
    }
  }

  function resetForm() {
    setOldPass('');
    setNewPass('');
    setNewPassConfirmed('');
    setToken('');
  }

  function onSubmit(
    e: React.MouseEvent<HTMLButtonElement>,
    validator: Validator
  ) {
    e.preventDefault();
    if (!validator.validate()) {
      return;
    }

    validator.reset();

    attemptActions.start();
    submit()
      .then(() => {
        attemptActions.stop();
        resetForm();
      })
      .catch(err => {
        attemptActions.error(err);
      });
  }

  function onSetMfaOption(option: MfaOption, validator: Validator) {
    setToken('');
    attemptActions.clear();
    validator.reset();
    setMfaType(option);
  }

  return (
    <Validation>
      {({ validator }) => (
        <Card as="form" width="456px" p="6">
          <Status attempt={attempt} />
          <FieldInput
            rule={requiredField('Current Password is required')}
            label="Current Password"
            value={oldPass}
            onChange={e => setOldPass(e.target.value)}
            type="password"
            placeholder="Password"
          />
          {mfaEnabled && (
            <Flex alignItems="flex-end" mb={4}>
              <Box width="50%" data-testid="mfa-select">
                <FieldSelect
                  label="Two-factor type"
                  value={mfaType}
                  options={mfaOptions}
                  onChange={opt => onSetMfaOption(opt as MfaOption, validator)}
                  mr={3}
                  mb={0}
                  isDisabled={isProcessing}
                />
              </Box>
              <Box width="50%">
                {mfaType.value === 'otp' && (
                  <FieldInput
                    label="Authenticator code"
                    inputMode="numeric"
                    autoComplete="one-time-code"
                    rule={requiredToken}
                    value={token}
                    onChange={e => setToken(e.target.value)}
                    placeholder="123 456"
                    mb={0}
                  />
                )}
              </Box>
            </Flex>
          )}
          <FieldInput
            rule={requiredPassword}
            label="New Password"
            value={newPass}
            onChange={e => setNewPass(e.target.value)}
            type="password"
            placeholder="New Password"
          />
          <FieldInput
            rule={requiredConfirmedPassword(newPass)}
            label="Confirm Password"
            value={newPassConfirmed}
            onChange={e => setNewPassConfirmed(e.target.value)}
            type="password"
            placeholder="Confirm Password"
          />
          <ButtonPrimary
            block
            disabled={isProcessing}
            size="large"
            onClick={e => onSubmit(e, validator)}
            mt={5}
          >
            Update Password
          </ButtonPrimary>
        </Card>
      )}
    </Validation>
  );
}

function Status({ attempt }: StatusProps) {
  if (attempt.isFailed) {
    return <Alerts.Danger>{attempt.message}</Alerts.Danger>;
  }

  if (attempt.isSuccess) {
    return <Alerts.Success>Your password has been changed!</Alerts.Success>;
  }

  return null;
}

type StatusProps = {
  attempt: ReturnType<typeof useAttempt>[0];
};

type Props = {
  auth2faType?: Auth2faType;
  preferredMfaType?: PreferredMfaType;
  onChangePass(oldPass: string, newPass: string, token: string): Promise<any>;
  onChangePassWithWebauthn(oldPass: string, newPass: string): Promise<any>;
};

export default FormPassword;
