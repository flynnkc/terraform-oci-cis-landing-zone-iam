## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | < 1.3.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_oci"></a> [oci](#provider\_oci) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [oci_identity_domain.these](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/identity_domain) | resource |
| [oci_identity_domains_dynamic_resource_group.these](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/identity_domains_dynamic_resource_group) | resource |
| [oci_identity_domains_group.these](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/identity_domains_group) | resource |
| [oci_identity_domain.dyngrp_domain](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/identity_domain) | data source |
| [oci_identity_domain.grp_domain](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/identity_domain) | data source |
| [oci_identity_domains_users.these](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/identity_domains_users) | data source |
| [oci_identity_regions.these](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/identity_regions) | data source |
| [oci_identity_tenancy.this](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/identity_tenancy) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_compartments_dependency"></a> [compartments\_dependency](#input\_compartments\_dependency) | A map of objects containing the externally managed compartments this module may depend on. All map objects must have the same type and must contain at least an 'id' attribute (representing the compartment OCID) of string type. | `map(any)` | `null` | no |
| <a name="input_identity_domain_dynamic_groups_configuration"></a> [identity\_domain\_dynamic\_groups\_configuration](#input\_identity\_domain\_dynamic\_groups\_configuration) | The identity domain dynamic groups configuration. | <pre>object({<br>    default_identity_domain_id  = optional(string)<br>    default_defined_tags        = optional(map(string))<br>    default_freeform_tags       = optional(map(string))<br>    dynamic_groups = map(object({<br>      identity_domain_id        = optional(string),<br>      name                      = string,<br>      description               = optional(string),<br>      matching_rule             = string,<br>      defined_tags              = optional(map(string)),<br>      freeform_tags             = optional(map(string))<br>    }))<br>  })</pre> | `null` | no |
| <a name="input_identity_domain_groups_configuration"></a> [identity\_domain\_groups\_configuration](#input\_identity\_domain\_groups\_configuration) | The identity domain groups configuration. | <pre>object({<br>    default_identity_domain_id  = optional(string)<br>    default_defined_tags        = optional(map(string))<br>    default_freeform_tags       = optional(map(string))<br>    groups = map(object({<br>      identity_domain_id        = optional(string),<br>      name                      = string,<br>      description               = optional(string),<br>      requestable               = optional(bool),<br>      members                   = optional(list(string)),<br>      defined_tags              = optional(map(string)),<br>      freeform_tags             = optional(map(string))<br>    }))<br>  })</pre> | `null` | no |
| <a name="input_identity_domains_configuration"></a> [identity\_domains\_configuration](#input\_identity\_domains\_configuration) | The identity domains configuration. | <pre>object({<br>    default_compartment_id = optional(string)<br>    default_defined_tags   = optional(map(string))<br>    default_freeform_tags  = optional(map(string))<br>    identity_domains = map(object({<br>      compartment_id            = optional(string),<br>      display_name              = string,<br>      description               = string,<br>      home_region               = optional(string),<br>      license_type              = string,<br>      admin_email               = optional(string),<br>      admin_first_name          = optional(string),<br>      admin_last_name           = optional(string),<br>      admin_user_name           = optional(string),<br>      is_hidden_on_login        = optional(bool),<br>      is_notification_bypassed  = optional(bool),<br>      is_primary_email_required = optional(bool),<br>      defined_tags              = optional(map(string)),<br>      freeform_tags             = optional(map(string))<br>    }))<br>  })</pre> | `null` | no |
| <a name="input_module_name"></a> [module\_name](#input\_module\_name) | The module name. | `string` | `"iam-identity-domains"` | no |
| <a name="input_tenancy_ocid"></a> [tenancy\_ocid](#input\_tenancy\_ocid) | The OCID of the tenancy. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_identity_domain_dynamic_groups"></a> [identity\_domain\_dynamic\_groups](#output\_identity\_domain\_dynamic\_groups) | The identity domain groups |
| <a name="output_identity_domain_groups"></a> [identity\_domain\_groups](#output\_identity\_domain\_groups) | The identity domain groups |
| <a name="output_identity_domains"></a> [identity\_domains](#output\_identity\_domains) | The identity domains. |
