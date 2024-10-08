#!/bin/bash

set -euo pipefail

usage() { echo "Usage: $0 [-p <cloudtrail>]" 1>&2; exit 0; }

while getopts "p:f:h" opt; do
  case $opt in
    p) export PROJECT="$OPTARG";;
    f) export TO_RUN="$OPTARG";;
    c) export COMMAND="$OPTARG";;
    h) usage;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
    :)
      echo "Option -$OPTARG requires an argument." >&2
      exit 1
      ;;
  esac
done


: $PROJECT
: $TO_RUN

pushd $PROJECT
go test "${TO_RUN}"
popd || exit
