provider "oci" {
  region               = var.region
  tenancy_ocid         = var.tenancy_ocid
  user_ocid            = var.user_ocid
  fingerprint          = var.fingerprint
  private_key_path     = var.private_key_path
  private_key_password = var.private_key_password
}

module "compartments" {
    source = "../../compartments"
    tenancy_ocid = var.tenancy_ocid
    compartments_configuration = var.compartments_configuration
}