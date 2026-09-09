// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package managedcapture

import (
	"context"
	"fmt"

	"github.com/ethosure/coding_ethos/go/internal/codeintel"
	"github.com/ethosure/coding_ethos/go/internal/hookoutput"
	"github.com/ethosure/coding_ethos/go/internal/lint"
	"github.com/ethosure/coding_ethos/go/internal/outputsurface"
)

func logCapturedToolResult(
	cwd string,
	result lint.Result,
) string {
	tracePath, err := lint.LogResult(cwd, result)
	if err != nil {
		emitManagedCaptureText("warning: lint trace not written: " + err.Error())

		return ""
	}

	err = writeCapturedToolSARIFSidecar(tracePath, result)
	if err != nil {
		emitManagedCaptureText("warning: lint SARIF sidecar not written: " + err.Error())
	}

	err = outputsurface.AutoPruneSurface(
		context.Background(),
		cwd,
		"lint_traces",
		false,
	)
	if err != nil {
		emitManagedCaptureText("warning: lint trace auto-prune failed: " + err.Error())
	}

	return tracePath
}

func writeCapturedToolSARIFSidecar(tracePath string, result lint.Result) error {
	err := hookoutput.WriteLintSARIFSidecar(tracePath, result)
	if err != nil {
		return fmt.Errorf("write captured tool SARIF sidecar: %w", err)
	}

	return nil
}

func refreshCapturedToolCodeIntel(root, tracePath string, changedFiles []string) {
	ctx := context.Background()

	err := codeintel.IngestLintTraceFile(ctx, root, tracePath)
	if err != nil {
		emitManagedCaptureText(
			"warning: captured lint trace not ingested into code-intel: " + err.Error(),
		)
	}

	if len(changedFiles) == 0 {
		return
	}

	_, err = codeintel.RefreshLintFiles(ctx, root, changedFiles)
	if err != nil {
		emitManagedCaptureText(
			"warning: captured lint code-intel refresh failed: " + err.Error(),
		)
	}
}
