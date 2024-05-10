#!/bin/bash

### Run in scripts directory ###

# Test
python3 preprocessor.py ../..
go test -count=1 -timeout 30m -p 1 .. | tee test_output.log
terratest_log_parser -testlog test_output.log -outputdir test_output