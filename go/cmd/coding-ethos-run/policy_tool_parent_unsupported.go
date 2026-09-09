// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

//go:build !linux

package main

import (
	"fmt"
	"runtime"

	"github.com/ethosure/coding_ethos/go/internal/apperror"
)

func policyToolParentExecutable() (string, error) {
	return "", fmt.Errorf(
		"%w on %s",
		apperror.StaticError(
			"actionlint ShellCheck protocol is unsupported because parent executable identity cannot be authenticated",
		),
		runtime.GOOS,
	)
}
