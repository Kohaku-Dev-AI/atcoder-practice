#!/usr/bin/env bash
# Easy100/001 .. Easy100/100 を作成する（既存はそのまま）
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
for i in $(seq 1 100); do
  mkdir -p "${ROOT}/Easy100/$(printf '%03d' $i)"
done
echo "OK: ${ROOT}/Easy100/001 .. 100"
