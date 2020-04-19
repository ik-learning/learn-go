#!/bin/bash

echo "Run tests"

set -euo pipefail


usage() { echo "Usage: $0 [-p <fib | TODO >]" 1>&2; exit 1; }

: $TO_RUN

# pushd $PROJECT
go test ${TO_RUN}
# go test
# go test -bench=.
# go test -run ${TO_RUN}*
# echo "Benchmark Function"
# go test --run $TO_RUN* -bench=.
# popd || exit
