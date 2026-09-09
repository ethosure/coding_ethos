// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package policy

import "strings"

func sentence(parts ...string) string {
	return strings.Join(parts, " ")
}
