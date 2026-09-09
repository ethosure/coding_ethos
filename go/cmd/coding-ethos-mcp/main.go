// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package main

import (
	"os"

	"github.com/ethosure/coding_ethos/go/internal/execguard"
	"github.com/ethosure/coding_ethos/go/internal/feedback"
	"github.com/ethosure/coding_ethos/go/internal/mcpcli"
)

func main() {
	execguard.Enter("coding-ethos-mcp")

	err := mcpcli.Run(os.Args[1:], os.Stdin, os.Stdout)
	if err != nil {
		feedback.Emit(
			os.Stderr,
			feedback.Error{Message: err.Error()},
			feedback.FormatTOON,
		)
		os.Exit(1)
	}
}
