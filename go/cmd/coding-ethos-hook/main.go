// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package main

import (
	"os"

	"github.com/ethosure/coding_ethos/go/internal/execguard"
	"github.com/ethosure/coding_ethos/go/internal/hookcli"
)

func main() {
	execguard.Enter("coding-ethos-hook")
	os.Exit(hookcli.Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
