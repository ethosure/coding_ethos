// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package codeintel

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/ethosure/coding_ethos/go/internal/apperror"
	"github.com/ethosure/coding_ethos/go/internal/evidence"
)

const (
	vectorBackendDuckDB    = "duckdb"
	VectorBackendDuckDBVSS = "duckdb-vss"
)

type VectorBackendConfig struct {
	Backend string
	URI     string
}

func NewVectorIndex(
	ctx context.Context,
	config VectorBackendConfig,
) (evidence.VectorIndex, error) {
	backend := strings.TrimSpace(config.Backend)
	if backend == "" {
		backend = VectorBackendDuckDBVSS
	}

	switch backend {
	case vectorBackendDuckDB, VectorBackendDuckDBVSS:
		return NewDuckDBVectorIndex(ctx, config.URI)
	default:
		return nil, apperror.Wrapf(
			apperror.StaticError("unsupported vector backend %q"),
			"unsupported vector backend %q",
			config.Backend,
		)
	}
}

func DefaultVectorPath(root string) string {
	return filepath.Join(root, ".coding-ethos", "code-intel.duckdb")
}
