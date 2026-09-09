// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package policygitcli

// Run executes the policy-protected git command family.
func Run(args []string) error {
	return runWithArgs(args)
}
