variable "tenancy_ocid" {
  default = ""
  type = string
  }
variable "region" {
  default = ""
  type = string
  }
variable "user_ocid" {
  default = ""
  type = string
  }
variable "fingerprint" {
  default = ""
  type = string
  }
variable "private_key_path" {
  default = ""
  type = string
  }
variable "private_key_password" {
  default = ""
  type = string
  }

variable "compartments_configuration" {
  description = "The compartments configuration. Use the compartments attribute to define your topology. OCI supports compartment hierarchies up to six levels."
  type = object({
    default_parent_id = optional(string) # the default parent for all top (first level) compartments. Use parent_id attribute within each compartment to specify different parents.
    default_defined_tags = optional(map(string)) # applies to all compartments, unless overriden by defined_tags in a compartment object
    default_freeform_tags = optional(map(string)) # applies to all compartments, unless overriden by freeform_tags in a compartment object
    enable_delete = optional(bool) # whether or not compartments are physically deleted when destroyed. Default is false.
    compartments = map(object({
      name          = string
      description   = string
      parent_id   = optional(string)
      defined_tags  = optional(map(string))
      freeform_tags = optional(map(string))
      tag_defaults     = optional(map(object({
        tag_id = string,
        default_value = string,
        is_user_required = optional(bool)
      })))
      children      = optional(map(object({
        name          = string
        description   = string
        defined_tags  = optional(map(string))
        freeform_tags = optional(map(string))
        tag_defaults     = optional(map(object({
            tag_id = string,
            default_value = string,
            is_user_required = optional(bool)
          })))
        children      = optional(map(object({
          name          = string
          description   = string
          defined_tags  = optional(map(string))
          freeform_tags = optional(map(string))
          tag_defaults     = optional(map(object({
            tag_id = string,
            default_value = string,
            is_user_required = optional(bool)
          })))
          children      = optional(map(object({
            name          = string
            description   = string
            defined_tags  = optional(map(string))
            freeform_tags = optional(map(string))
            tag_defaults     = optional(map(object({
              tag_id = string,
              default_value = string,
              is_user_required = optional(bool)
            })))
            children      = optional(map(object({
              name          = string
              description   = string
              defined_tags  = optional(map(string))
              freeform_tags = optional(map(string))
              tag_defaults     = optional(map(object({
                tag_id = string,
                default_value = string,
                is_user_required = optional(bool)
              })))
              children      = optional(map(object({
                name          = string
                description   = string
                defined_tags  = optional(map(string))
                freeform_tags = optional(map(string))
                tag_defaults     = optional(map(object({
                  tag_id = string,
                  default_value = string,
                  is_user_required = optional(bool)
                })))
              })))  
            })))
          })))
        })))
      })))  
    }))
  })
  default = null
}