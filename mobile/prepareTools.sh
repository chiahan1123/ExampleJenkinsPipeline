#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# setup go environment variables
GOPATH=${DIR}
GOBIN=${GOPATH}/bin

# mobile development packages
GVT=github.com/FiloSottile/gvt

cleanUp() {
  echo "Cleaning up mobile development packages..."
  rm -rf "${GOPATH}/src/github.com"
  echo ""
}

cleanUp
echo "Installing GVT..."
go get -u -x -v "${GVT}"
echo ""
cleanUp

