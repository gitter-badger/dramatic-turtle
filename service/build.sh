#!/bin/sh
cd service
make init
make docker-build-image
cd ..