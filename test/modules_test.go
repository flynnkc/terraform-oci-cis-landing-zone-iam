package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// TestCompartments plans and applies terraform for the compartments module based on the
// vision example.
func TestCompartments(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../compartments/examples/vision",
		VarFiles:     []string{"input.auto.tfvars.template"},
		PlanFilePath: "./tf.plan",
	})

	terraform.InitAndPlan(t, terraformOptions)

	// We defer destroy, but this will leave compartments in place
	defer terraform.Destroy(t, terraformOptions)
	terraform.Apply(t, terraformOptions)

}

func TestDynamicGroups(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../dynamic-groups/examples/vision",
		VarFiles:     []string{"input.auto.tfvars.template"},
		PlanFilePath: "./tf.plan",
	})

	terraform.InitAndPlan(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)
	terraform.Apply(t, terraformOptions)

}
