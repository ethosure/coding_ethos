// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package main

import (
	"os"

	"github.com/ethosure/coding_ethos/go/internal/execguard"
	"github.com/ethosure/coding_ethos/go/internal/toolchaincli"
)

func main() {
	execguard.Enter("coding-ethos-toolchain")
	os.Exit(toolchaincli.Run(os.Args[1:]))
}
