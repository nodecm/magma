/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */
'use strict';

import {omit} from 'lodash-es';
import Sequelize from 'sequelize';
import {AccessRoles} from '@fbcnms/auth/roles';

import type {DataTypes, Model} from 'sequelize';
import type {AssociateProp} from './AssociateTypes.flow';

export type UserRawType = {
  email: string,
  organization?: string,
  password: string,
  role: number,
  networkIDs?: Array<string>,
  isSuperUser?: boolean,
};

type UserModel = Model<UserRawType>;
export type StaticUserModel = Class<UserModel>;
export type UserType = UserModel & UserRawType;

export default (
  sequelize: Sequelize,
  types: DataTypes,
): StaticUserModel & AssociateProp => {
  const User = sequelize.define(
    'User',
    {
      email: types.STRING,
      organization: types.STRING,
      password: types.STRING,
      role: types.INTEGER,
      networkIDs: {
        type: types.JSON,
        allowNull: false,
        defaultValue: [],
        get() {
          return this.getDataValue('networkIDs') || [];
        },
      },
    },
    {
      getterMethods: {
        isSuperUser() {
          return this.role === AccessRoles.SUPERUSER;
        },
      },
    },
  );
  User.associate = function(_models) {
    // associations can be defined here
  };
  User.prototype.toJSON = function() {
    return omit(this.get(), 'password');
  };
  return User;
};
