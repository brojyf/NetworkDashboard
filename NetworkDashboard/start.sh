#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

# Start data collection in background before launching API server
"$SCRIPT_DIR/scripts/ping_test.sh" &

exec "$SCRIPT_DIR/server"
