//go:build windows

// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package codeintel

import "github.com/ethosure/coding_ethos/go/internal/apperror"

func duckDBRebuildLockPIDStale(_ int) (bool, error) {
	return false, apperror.StaticError("pid liveness is unavailable on Windows")
}
