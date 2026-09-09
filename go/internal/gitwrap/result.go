// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package gitwrap

import "github.com/ethosure/coding_ethos/go/internal/policy"

type Result struct {
	Operation string            `json:"operation,omitempty"`
	Status    string            `json:"status"`
	Decisions []policy.Decision `json:"decisions,omitempty"`
	Argv      []string          `json:"argv"`
}

func (result Result) Blocked() bool {
	return result.Status == "blocked"
}
