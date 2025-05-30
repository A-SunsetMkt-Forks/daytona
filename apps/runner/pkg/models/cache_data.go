// Copyright 2025 Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package models

import (
	"time"

	"github.com/daytonaio/runner/pkg/models/enums"
)

type CacheData struct {
	SandboxState    enums.SandboxState
	SnapshotState   enums.SnapshotState
	DestructionTime *time.Time
}
