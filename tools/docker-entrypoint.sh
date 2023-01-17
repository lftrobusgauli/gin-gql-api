#!/bin/sh

# abort on any error (including if wait-for-it fails).
set -e

# wait for the db to be up
if [ -n "$HOST" ]; then
  /app/tools/wait-for-it.sh "$HOST:${PORT:-5432}"
fi

# run the main container command.
exec "$@"
