// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package mcpcli

import "io"

// Run executes the MCP server CLI command family.
func Run(args []string, stdin io.Reader, stdout io.Writer) error {
	return runWithIO(args, stdin, stdout)
}
