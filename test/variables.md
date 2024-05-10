# Landing Zone Modules Template Variables

## Tenancy Connectivity Variables

| Template Variable Name | Description |
| --- | --- |
| <TENANCY_OCID> | Oracle Cloud ID of the OCI tenancy |
| <USER_OCID> | *(Required for user principal)* Oracle Cloud ID of the user running the script |
| <PEM_KEY_FINGERPRINT> | *(Required for user principal)* Fingerprint of the PEM key associated with the user running the script |
| <PATH_TO_PRIVATE_KEY> | *(Required for user principal)* File path to the private key associated with the user |
| <PRIVATE_KEY_PASSWORD> | *(Required for user principal)* Password to private key; may be empty string |
| <TENANCY_REGION> | OCI Region to deploy script assets |

## Compartments Module Variables

| Template Variable Name | Description |
| --- | --- |
| <ENCLOSING_COMPARTMENT_OCID> | OCID of the compartment that will be the parent of deployed compartments |
| <OCI_COST_CENTER_TAG_OCID> | OCID of the tag to act as cost tracking tag for compartments |
| <OCI_ENVIRONMENT_TAG_OCID> | OCID of the tag to act as environment label |

## Dynamic Groups Module Variables

| Template Variable Name | Description |
| --- | --- |
| <SECURITY_COMPARTMENT_OCID> | Oracle Cloud ID of the security compartment |
| <APPLICATION_COMPARTMENT_OCID> | Oracle Cloud ID of the applicaiton compartment |
| <DATABASE_COMPARTMENT_OCID> | Oracle Cloud ID of the database compartment |

## Identity Domains Module Variables

| Template Variable Name | Description |
| --- | --- |
| <ENCLOSING_COMPARTMENT_OCID> | OCID of compartment to deploy domain(s) |
| <HOME_REGION> | Home region of Identity Domain; may be empty string |
| <ADMIN_EMAIL> | Email address of domain administrator |
| <ADMIN_FIRST_NAME> | First name of domain administrator |
| <ADMIN_LAST_NAME> | Last name of the domain administrator |
| <ADMIN_USERNAME> | Username for domain administrator |
| <SECURITY_COMPARTMENT_OCID> | Oracle Cloud ID of the security compartment |
| <APPLICATION_COMPARTMENT_OCID> | Oracle Cloud ID of the application compartment |
| <DATABASE_COMPARTMENT_OCID> | Oracle Cloud ID of the database compartment |

## Policies Module Variables

| Template Variable Name | Description |
| --- | --- |
| <ENCLOSING_COMPARTMENT_OCID> | OCID of compartment to deploy domain(s) |
