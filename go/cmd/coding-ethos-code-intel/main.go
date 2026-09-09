// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package main

import (
	"context"
	"os"

	"github.com/ethosure/coding_ethos/go/internal/codeintelcli"
	"github.com/ethosure/coding_ethos/go/internal/execguard"
	"github.com/ethosure/coding_ethos/go/internal/feedback"
)

func main() {
	execguard.Enter("coding-ethos-code-intel")

	err := codeintelcli.Run(context.Background(), os.Args[1:])
	if err != nil {
		feedback.Emit(
			os.Stderr,
			feedback.Error{Message: err.Error()},
			feedback.FormatTOON,
		)
		os.Exit(1)
	}
}
