/**
 * Copyright 2023 Gravitational, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { EmployeeOptions, TeamOptions, TitleOptions } from './constants';

export type QuestionnaireProps = {
  username: string;
};

export type QuestionProps = {
  updateFields: (fields: Partial<QuestionnaireFormFields>) => void;
};

export type CompanyProps = QuestionProps & {
  companyName: string;
  numberOfEmployees: EmployeeOptionsStrings;
};

export type RoleProps = QuestionProps & {
  role: TitleOptionsStrings;
  team: TeamOptionsStrings;
};

export type ResourceType = {
  label: Resource;
  image: string;
};

export type ResourcesProps = QuestionProps & {
  resources: ResourceType[];
  checked: Resource[];
};

export enum Resource {
  RESOURCE_WINDOWS_DESKTOPS = 'Windows Desktops',
  RESOURCE_SERVER_SSH = 'Server/SSH',
  RESOURCE_DATABASES = 'Databases',
  RESOURCE_KUBERNETES = 'Kubernetes',
  RESOURCE_WEB_APPLICATIONS = 'Web Applications',
}

export type QuestionnaireFormFields = {
  companyName: string;
  employeeCount: EmployeeOptionsStrings;
  role: TitleOptionsStrings;
  team: TeamOptionsStrings;
  resources: Resource[];
};

export type EmployeeOptionsStrings = keyof typeof EmployeeOptions;
export type TeamOptionsStrings = keyof typeof TeamOptions;
export type TitleOptionsStrings = keyof typeof TitleOptions;
