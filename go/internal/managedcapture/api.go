// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package managedcapture

import (
	"io"
	"os"

	"github.com/ethosure/coding_ethos/go/diagnostics"
	"github.com/ethosure/coding_ethos/go/internal/policy"
)

const BlockedExitCode = 2

type PolicyContext struct {
	Skills       map[string]policy.Skill
	EvidenceMaps []diagnostics.EvidenceMap
	Policies     []policy.Policy
}

type CaptureOptions struct {
	PolicyContext PolicyContext
	Tool          string
	ToolPath      string
	Cwd           string
	TraceRoot     string
	OutputFormat  string
	Output        io.Writer
	Args          []string
	CodeIntel     bool
}

func Capture(options CaptureOptions) int {
	return runCapturedToolWithCodeIntel(
		options.Tool,
		options.ToolPath,
		options.Cwd,
		options.TraceRoot,
		options.Args,
		options.PolicyContext,
		options.Output,
		options.OutputFormat,
		options.CodeIntel,
	)
}

func ExecutableAvailable(path string) bool {
	return isExecutable(path)
}

func exitErr(err error) {
	emitManagedCaptureError(err)
	os.Exit(1)
}
