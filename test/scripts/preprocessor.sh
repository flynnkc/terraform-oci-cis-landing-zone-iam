#!/bin/bash

declare -a CHANGE_VARS
CHANGE_VARS+="s/<TENANCY_OCID>/$TENANCY_OCID/g"

for file in $(find . -name input.auto.tfvars.template); do
    for v in $CHANGE_VARS; do
        sed -i.bac $v $file
    done
done