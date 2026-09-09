// SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
// SPDX-License-Identifier: AGPL-3.0-only

package evaluators

import "github.com/ethosure/coding_ethos/go/internal/policy"

type Evaluator interface {
	Evaluate(policyDef policy.Policy, context Context) ([]policy.Decision, error)
}

type EvaluatorFunc func(
	policyDef policy.Policy,
	context Context,
) ([]policy.Decision, error)

func (fn EvaluatorFunc) Evaluate(
	policyDef policy.Policy,
	context Context,
) ([]policy.Decision, error) {
	return fn(policyDef, context)
}
