// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package evaluators

import (
	"strings"

	"github.com/ethosure/coding_ethos/go/internal/celexpr"
)

func celHookCommands(
	context Context,
	expression string,
) []celexpr.HookCommandInput {
	if !strings.Contains(expression, "hook_commands") {
		return nil
	}

	if len(context.HookCommands) > 0 {
		return context.HookCommands
	}

	return celexpr.HookCommandInputs(context.Cwd, context.Files)
}
