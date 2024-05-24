package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// TestCompartments plans and applies terraform for the compartments module based on the
// vision example.
func TestCompartments(t *testing.T) {

	// Init, plan, and apply
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../compartments/examples/vision",
		VarFiles:     []string{"input.auto.tfvars.template"},
		PlanFilePath: "./tf.plan",
	})

	terraform.InitAndPlan(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)
	terraform.Apply(t, terraformOptions)

	// Additional tests below

}

// TestDynamicGroups plans and applies terraform for the dynamic-groups module based
// on the vision example.
func TestDynamicGroups(t *testing.T) {

	// Init, plan, and apply
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../dynamic-groups/examples/vision",
		VarFiles:     []string{"input.auto.tfvars.template"},
		PlanFilePath: "./tf.plan",
	})

	terraform.InitAndPlan(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)
	terraform.Apply(t, terraformOptions)

	// Additional tests below

}

// TestGroups plans and applies terraform for the groups module based on the visions
// example.
func TestGroups(t *testing.T) {

	// Init, plan, and apply
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../groups/examples/vision",
		VarFiles:     []string{"input.auto.tfvars.template"},
		PlanFilePath: "./tf.plan",
	})

	terraform.InitAndPlan(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)
	terraform.Apply(t, terraformOptions)

	// Additional tests below

}
