// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package hooks

import "github.com/ethosure/coding_ethos/go/internal/policy"

func newDenialTrackingID(event Event, decisions []policy.Decision) string {
	return hookTraceID(event, Result{
		Event:     event.HookEventName,
		Provider:  event.Provider(),
		Status:    statusBlocked,
		Tool:      event.ToolName,
		Decisions: decisions,
	})
}
