#!/bin/bash -e

cd "$(dirname $0)"

echo "This one waits for the sub-process to exit before exiting."
(
  set -x
  go test -v -count=1 .
)
echo "This one returns instantly, but the sub-process is still running."
(
  set -x
  go test -v -count=1 
)
