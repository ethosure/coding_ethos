// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package main

import (
	"errors"
	"os"

	"github.com/ethosure/coding_ethos/go/internal/execguard"
	"github.com/ethosure/coding_ethos/go/internal/feedback"
	"github.com/ethosure/coding_ethos/go/internal/gitwrap"
	"github.com/ethosure/coding_ethos/go/internal/policygitcli"
)

func main() {
	execguard.Enter("coding-ethos-git")

	err := policygitcli.Run(os.Args[1:])
	if err == nil {
		return
	}

	var exitError gitwrap.ExitCodeError
	if errors.As(err, &exitError) {
		os.Exit(exitError.Code)
	}

	feedback.Emit(
		os.Stderr,
		feedback.Error{Message: err.Error()},
		feedback.FormatTOON,
	)
	os.Exit(1)
}
