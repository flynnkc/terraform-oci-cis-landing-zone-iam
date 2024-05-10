package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

const (
	OCI_TENANT      string = "OCI_TENANCY_OCID"
	OCI_USER        string = "OCI_USER_OCID"
	OCI_FINGERPRINT string = "OCI_FINGERPRINT"
	OCI_KEY         string = "OCI_PRIVATE_KEY_PATH"
	OCI_PASS        string = "OCI_PRIVATE_KEY_PASSWORD"
	OCI_REGION      string = "OCI_REGION"
)

// TestCompartments plans and applies terraform for the compartments module based on the
// vision example.
func TestCompartments(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../compartments/examples/vision",
		VarFiles:     []string{"input.auto.tfvars.template"},
		PlanFilePath: "./tf.plan",
	})

	terraform.InitAndPlanAndShow(t, terraformOptions)

	// We defer destroy, but this will leave compartments in place
	defer terraform.Destroy(t, terraformOptions)
	terraform.Apply(t, terraformOptions)

}
