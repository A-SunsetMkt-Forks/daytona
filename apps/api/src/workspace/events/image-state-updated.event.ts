/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

import { Image } from '../entities/image.entity'
import { ImageState } from '../enums/image-state.enum'

export class ImageStateUpdatedEvent {
  constructor(
    public readonly image: Image,
    public readonly oldState: ImageState,
    public readonly newState: ImageState,
  ) {}
}
