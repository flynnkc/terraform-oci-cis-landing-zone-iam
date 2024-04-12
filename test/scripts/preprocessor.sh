#!/bin/bash

# Array of sed command strings
# User authentication section
declare -a CHANGE_VARS
CHANGE_VARS+="s/<TENANCY_OCID>/$TENANCY_OCID/g"
CHANGE_VARS+="s/<USER_OCID>/$USER_OCID/g"
CHANGE_VARS+="s/<PEM_KEY_FINGERPRINT>/$FINGERPRINT/g"
CHANGE_VARS+="s/<PATH_TO_PRIVATE_KEY>/$PRIVATE_KEY_PATH/g"
CHANGE_VARS+="s/<PRIVATE_KEY_PASSWORD>/$PRIVATE_KEY_PASS/g"
CHANGE_VARS+="s/<TENANCY_REGION>/$REGION/g"

# Template variables
CHANGE_VARS+="s/<OCI-PARENT-COMPARTMENT-OCID>/$PARENT_COMPARTMENT/g"
CHANGE_VARS+="s/<OCI-COST-CENTER-TAG-OCID>/$COST_CENTER_TAG/g"
CHANGE_VARS+="s/<OCI-ENVIRONMENT-TAG-OCID>/$COST_ENVIRONMENT_TAG_OCID/g"

for file in $(find . -name input.auto.tfvars.template); do
    for v in $CHANGE_VARS; do
        sed -i $v "$file"
    done
done
