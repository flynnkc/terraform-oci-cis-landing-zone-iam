#!/bin/bash

# Format Files
./preprocessor.sh

# Test
go test -count=1 -timeout 30m -p 1 .