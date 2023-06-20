#!/bin/bash

VERSION=0.0.6


make clean
git add .
git commit -m "v${VERSION} debug"
git tag "v$VERSION"
make gitpush

make docker-image
#make build