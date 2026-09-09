// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package main

import (
	"os"

	"github.com/ethosure/coding_ethos/go/internal/agenthookscli"
	"github.com/ethosure/coding_ethos/go/internal/execguard"
)

func main() {
	execguard.Enter("coding-ethos-agent-hooks")
	os.Exit(agenthookscli.Run(os.Args[1:]))
}
