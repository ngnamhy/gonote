#!/bin/sh
set -e

echo "APP_MODE=$APP_MODE"

if [ "$APP_MODE" = "debug" ]; then
  echo ">>> Running with Delve"
  exec dlv exec /app/app \
    --listen=:40000 \
    --headless=true \
    --api-version=2 \
    --accept-multiclient \
    --continue
else
  echo ">>> Running normally"
  exec /app/app
fi
