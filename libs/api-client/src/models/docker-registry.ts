/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: Apache-2.0
 */

/* tslint:disable */

/**
 * Daytona
 * Daytona AI platform API Docs
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@daytona.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 *
 * @export
 * @interface DockerRegistry
 */
export interface DockerRegistry {
  /**
   * Registry ID
   * @type {string}
   * @memberof DockerRegistry
   */
  id: string
  /**
   * Registry name
   * @type {string}
   * @memberof DockerRegistry
   */
  name: string
  /**
   * Registry URL
   * @type {string}
   * @memberof DockerRegistry
   */
  url: string
  /**
   * Registry username
   * @type {string}
   * @memberof DockerRegistry
   */
  username: string
  /**
   * Registry project
   * @type {string}
   * @memberof DockerRegistry
   */
  project: string
  /**
   * Registry type
   * @type {string}
   * @memberof DockerRegistry
   */
  registryType: DockerRegistryRegistryTypeEnum
  /**
   * Creation timestamp
   * @type {Date}
   * @memberof DockerRegistry
   */
  createdAt: Date
  /**
   * Last update timestamp
   * @type {Date}
   * @memberof DockerRegistry
   */
  updatedAt: Date
}

export const DockerRegistryRegistryTypeEnum = {
  INTERNAL: 'internal',
  ORGANIZATION: 'organization',
  PUBLIC: 'public',
  TRANSIENT: 'transient',
} as const

export type DockerRegistryRegistryTypeEnum =
  (typeof DockerRegistryRegistryTypeEnum)[keyof typeof DockerRegistryRegistryTypeEnum]
