/**
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

/**
 * Enum for all route paths in the application
 * Use this for consistent route references across the app
 */
export enum RoutePath {
  // Main routes
  LANDING = '/',
  LOGOUT = '/logout',
  DASHBOARD = '/dashboard',
  DOCS = '/docs',
  SLACK = '/slack',

  // Dashboard sub-routes
  KEYS = '/dashboard/keys',
  SANDBOXES = '/dashboard/sandboxes',
  SNAPSHOTS = '/dashboard/snapshots',
  REGISTRIES = '/dashboard/registries',
  VOLUMES = '/dashboard/volumes',
  LIMITS = '/dashboard/limits',
  BILLING = '/dashboard/billing',
  MEMBERS = '/dashboard/members',
  ROLES = '/dashboard/roles',
  SETTINGS = '/dashboard/settings',
  ONBOARDING = '/dashboard/onboarding',
  AUDIT_LOGS = '/dashboard/audit-logs',

  // User routes
  USER_INVITATIONS = '/dashboard/user/invitations',
  LINKED_ACCOUNTS = '/dashboard/user/linked-accounts',
}

/**
 * Returns only the path segment for dashboard sub-routes
 * Useful for nested routes under the dashboard
 */
export const getRouteSubPath = (path: RoutePath): string => {
  return path.replace('/dashboard/', '')
}
