/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

import { Image } from '../entities/image.entity'

export class ImageRemovedEvent {
  constructor(public readonly image: Image) {}
}
