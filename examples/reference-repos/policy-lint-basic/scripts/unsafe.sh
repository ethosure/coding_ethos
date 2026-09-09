#!/usr/bin/env bash
# SPDX-FileCopyrightText: 2026 Ethosure Governance Inc. <oss@ethosure.com>
# SPDX-License-Identifier: AGPL-3.0-only

set -euo pipefail

name=${1:-world}
printf '%s\n' "$name"
