// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package evaluators

import (
	"strings"
)

func currentBranch(cwd string) (string, bool) {
	cmd := gitCommand(cwd, "branch", "--show-current")

	output, err := cmd.Output()
	if err != nil {
		return "", false
	}

	return strings.TrimSpace(string(output)), true
}
