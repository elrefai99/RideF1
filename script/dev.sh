#!/usr/bin/env sh
set -eu

cd "$(dirname "$0")/../app/api"
export APP_ENV=development
exec air -c .air.toml
