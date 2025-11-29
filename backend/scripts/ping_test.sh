#!/usr/bin/env bash
set -euo pipefail

DIR_PREFIX="$(cd "$(dirname "$0")/.." && pwd)"
CONFIG_DIR="$DIR_PREFIX/config"
DATA_DIR="$DIR_PREFIX/data"
OUTPUT_PING_FILE="$DATA_DIR/ping_results.csv"
OUTPUT_HOP_FILE="$DATA_DIR/hop_results.csv"
INTERVAL_SECONDS=60
PING_COUNT=4
HOP_LIMIT=15
FILE="$CONFIG_DIR/websites.txt"
rm -f "$OUTPUT_PING_FILE" "$OUTPUT_HOP_FILE"

ping_site() {
  local site="$1"
  local site_type="$2"
  local timestamp=$(date +"%H:%M")

  local ping_output
  if ! ping_output=$(ping -c "$PING_COUNT" "$site" 2>/dev/null); then
    echo "$timestamp,$site,$site_type,,," >> "$OUTPUT_PING_FILE"
    return
  fi

  local packet_times
  packet_times=$(awk -F'time=' '/time=/{sub(/ ms/, "", $2); printf "%s;", $2}' <<< "$ping_output")
  packet_times="${packet_times%;}"

  local avg_latency jitter
  read avg_latency jitter <<< "$(
    tr ';' '\n' <<< "$packet_times" |
      awk '{
            a[NR]=$1; s+=$1
          }
          END {
            avg = s / NR
            ss = 0
            for(i=1;i<=NR;i++){ d=a[i]-avg; ss+=d*d }
            jitter = sqrt(ss/NR)
            printf "%.3f %.3f", avg, jitter
          }'
  )"

  echo "$timestamp,$site,$site_type,$avg_latency,$jitter,$packet_times" >> "$OUTPUT_PING_FILE"
}

traceroute_site() {
  local site="$1"
  local line hop ip hostname latency printed

  while IFS= read -r line; do
    # detect new hop start
    if [[ "$line" =~ ^[[:space:]]*traceroute\ to ]]; then
          continue
        fi
    if [[ "$line" =~ ^[[:space:]]*([0-9]+)[[:space:]]+ ]]; then
      hop=${BASH_REMATCH[1]}
      printed=0
    fi

    # timeout (* * *)
    if [[ "$line" =~ \*[\ \t]+\*[\ \t]+\* ]]; then
      if [[ $printed -eq 0 ]]; then
        echo "$site,$hop,,," >> "$OUTPUT_HOP_FILE"
        printed=1
      fi
      continue
    fi

    # must not print twice for the same hop
    if [[ $printed -eq 1 ]]; then
      continue
    fi

    hostname=""
    ip=""
    latency=""

    # extract hostname + ip
    if [[ "$line" =~ ([a-zA-Z0-9\.\-]+)\ \(([0-9\.]+)\) ]]; then
      hostname="${BASH_REMATCH[1]}"
      ip="${BASH_REMATCH[2]}"
    fi

    # extract the first latency
    if [[ "$line" =~ ([0-9]+\.[0-9]+)\ ms ]]; then
      latency="${BASH_REMATCH[1]}"
    fi

    echo "$site,$hop,$ip,$hostname,$latency" >> "$OUTPUT_HOP_FILE"
    printed=1

  done
}

run_ping() {
  while IFS='=' read -r site site_type; do
    ping_site "$site" "$site_type"
  done < "$FILE"
}

run_traceroute() {
  while IFS='=' read -r site site_type; do
    traceroute -m 10 "$site" 2>/dev/null | traceroute_site "$site"
  done < "$FILE"
}

main() {
  [ ! -f "$FILE" ] && echo "Missing config: $FILE" && exit 1
  mkdir -p "$DATA_DIR"
  if [ ! -f "$OUTPUT_PING_FILE" ]; then
    echo "time,site,site_type,latency_ms,jitter_ms,packet_list" > "$OUTPUT_PING_FILE"
  fi
  if [ ! -f "$OUTPUT_HOP_FILE" ]; then
    echo "websites,hop,ip,hostname,latency" > "$OUTPUT_HOP_FILE"
  fi

  run_traceroute &
  while true; do
    run_ping
    sleep "$INTERVAL_SECONDS"
  done
}

main "$@"
