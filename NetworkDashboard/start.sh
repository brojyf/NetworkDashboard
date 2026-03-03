#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

trap 'pkill -P $$ || true' EXIT INT TERM

"$SCRIPT_DIR/scripts/script.sh" &
"$SCRIPT_DIR/server" &
wait -n
