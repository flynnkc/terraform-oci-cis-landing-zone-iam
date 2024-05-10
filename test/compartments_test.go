package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identity"
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
		TerraformDir: "./compartments",
		VarFiles:     []string{"../../compartments/examples/vision/input.auto.tfvars.template"},
		PlanFilePath: "./compartments/tf.plan",
	})

	terraform.InitAndPlanAndShow(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)
	terraform.Apply(t, terraformOptions)
	output := terraform.OutputMapOfObjects(t, terraformOptions, "compartments")

	// Compartments do not destroy by default, this code removes deployed compartments
	s := extractCompartments(output)
	fmt.Println("COMPARTMENTS:", s)
	for _, cmp := range s {
		fmt.Println("Cleaning up compartment:", cmp)
		defer CleanupCompartment(t, cmp, terraformOptions)
	}

}

func extractCompartments(m map[string]any) []string {
	var s []string

	for _, v := range m {
		if val, ok := v.(map[string]any); ok {
			if val["id"] != "" {
				s = append(s, val["id"].(string))
			}
		}
	}

	return s
}

// Needed to delete compartments after test run
func CleanupCompartment(t *testing.T, c string, o *terraform.Options) error {
	provider := common.NewRawConfigurationProvider(
		os.Getenv(OCI_TENANT),
		os.Getenv(OCI_USER),
		os.Getenv(OCI_REGION),
		os.Getenv(OCI_FINGERPRINT),
		os.Getenv(OCI_KEY),
		common.String(os.Getenv(OCI_PASS)),
	)

	client, err := identity.NewIdentityClientWithConfigurationProvider(provider)
	if err != nil {
		return err
	}

	request := identity.DeleteCompartmentRequest{
		CompartmentId: common.String(c),
	}

	response, err := client.DeleteCompartment(context.TODO(), request)
	if err != nil {
		return err
	} else if response.RawResponse.StatusCode != 200 {
		return fmt.Errorf("unsuccessful status code returned: %v",
			response.RawResponse.StatusCode)
	}

	return nil
}
