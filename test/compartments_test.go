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

const ENV_PREFIX string = "TF_VAR"

func TestCompartments(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "./compartments",
		VarFiles:     []string{"../../compartments/examples/vision/input.auto.tfvars.template"},
		PlanFilePath: "./compartments/tf.plan",
	})

	t.Logf("%+v", terraformOptions)

	//defer terraform.Destroy(t, terraformOptions)

	out := terraform.InitAndPlanAndShow(t, terraformOptions)
	fmt.Println(out)

	/* TODO Apply testing
	terraform.Apply(t, terraformOptions)
	output := terraform.OutputMapOfObjects(t, terraformOptions, "compartments")
	fmt.Println("OUTPUT:", output)
	*/

	/* Compartments do not destroy by default, this code removes deployed comaprtments
	s := extractCompartments(output)
	fmt.Println("COMPARTMENTS:", s)
	for _, cmp := range s {
		fmt.Println("Cleaning up compartment:", cmp)
		defer CleanupCompartment(t, cmp, terraformOptions)
	}
	*/
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

func CleanupCompartment(t *testing.T, c string, o *terraform.Options) error {
	provider := common.ConfigurationProviderEnvironmentVariables(
		ENV_PREFIX,
		o.Vars["private_key_password"].(string),
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

func GetVars() map[string]string {
	vars := []string{"tenancy_ocid", "user_ocid", "fingerprint", "private_key_path",
		"private_key_password", "region"}

	m := make(map[string]string)
	for _, v := range vars {
		m[v] = os.Getenv(fmt.Sprint(ENV_PREFIX, "_", v))
	}

	return m
}
