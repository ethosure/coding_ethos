// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package outputcli

import "context"

// Run executes the output-surface CLI command family.
func Run(ctx context.Context, args []string) error {
	return run(ctx, args)
}
