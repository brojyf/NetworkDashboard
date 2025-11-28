#!/usr/bin/env bash

# Strict mode
set -euo pipefail

DIR_PREFIX="$(cd "$(dirname "$0")/.." && pwd)"
CONFIG_DIR="$DIR_PREFIX/config"
DATA_DIR="$DIR_PREFIX/data"
OUTPUT_FILE="$DATA_DIR/ping_results.csv"
INTERVAL_SECONDS=60
PING_COUNT=4

# Try write csv headers
mkdir -p "$DATA_DIR"
if [ ! -f "$OUTPUT_FILE" ]; then
  echo "time,site,site_type,latency_ms,jitter_ms,packet_list" > "$OUTPUT_FILE"
fi

ping_site() {
  local site="$1"
  local site_type="$2"
  local timestamp=$(date +"%H:%M")

  # Ping
  local ping_output
  if ! ping_output=$(ping -c "$PING_COUNT" "$site" 2>/dev/null); then
    echo "$timestamp,$site,$site_type,,," >> "$OUTPUT_FILE"
    return
  fi

 # Extract time usage using awk
  local packet_times
  packet_times=$(awk -F'time=' '/time=/{sub(/ ms/, "", $2); printf "%s;", $2}' <<< "$ping_output")
  packet_times="${packet_times%;}"

  # Calculate latency and jitter
  local avg_latency jitter
  read avg_latency jitter <<< "$(
    tr ';' '\n' <<< "$packet_times" | \
    awk '{
          a[NR]=$1; s+=$1
        }
        END {
          if(NR == 0) { print "0 0"; exit }
          avg = s / NR
          ss = 0
          for(i=1; i<=NR; i++){
            d = a[i] - avg
            ss += d * d
          }
          jitter = sqrt(ss / NR)
          printf "%.3f %.3f", avg, jitter
        }'
  )"

  echo "$timestamp,$site,$site_type,$avg_latency,$jitter,$packet_times" >> "$OUTPUT_FILE"
}


load_sites() {
  # Make sure file exists
  local file="$CONFIG_DIR/websites.txt"
  [ ! -f "$file" ] && return 0
  while IFS='=' read -r site site_type; do
    ping_site "$site" "$site_type"
  done < "$file"
}

main() {
  while true; do
    load_sites
    sleep "$INTERVAL_SECONDS"
  done
}

main "$@"
