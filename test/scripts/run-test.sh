#!/bin/bash


# Test
go test -count=1 -timeout 30m -p 1 | tee test_output.log
terratest_log_parser -testlog test_output.log -outputdir test_output