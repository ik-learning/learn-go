#!/bin/bash

set -euo pipefail

echo "Install packages"

go get github.com/stretchr/testify
go get -u github.com/gorilla/mux
go get -u github.com/aws/aws-sdk-go
go get -u github.com/urfave/cli