// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package main

import (
	"errors"
	"os"

	"github.com/ethosure/coding_ethos/go/internal/execguard"
	"github.com/ethosure/coding_ethos/go/internal/feedback"
	"github.com/ethosure/coding_ethos/go/internal/hooklogcli"
)

type exitCoder interface {
	ExitCode() int
}

func main() {
	execguard.Enter("coding-ethos-hook-log")

	status := 0

	err := hooklogcli.Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		var exitErr exitCoder
		if errors.As(err, &exitErr) {
			status = exitErr.ExitCode()
		} else {
			status = 1

			feedback.Emit(
				os.Stderr,
				feedback.Error{Message: err.Error()},
				feedback.FormatTOON,
			)
		}
	}

	os.Exit(status)
}
