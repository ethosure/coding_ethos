// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package hooks

import "github.com/ethosure/coding_ethos/go/internal/hookoutput"

const (
	outputFormatAuto  = hookoutput.FormatAuto
	outputFormatHuman = hookoutput.FormatHuman
	outputFormatJSON  = hookoutput.FormatJSON
	outputFormatTOON  = hookoutput.FormatTOON
	outputFormatEnv   = hookoutput.FormatEnv
)

func selectedOutputFormat() string {
	return hookoutput.SelectedFormat()
}

func toonCell(value string) string {
	return hookoutput.TOONCell(value)
}
