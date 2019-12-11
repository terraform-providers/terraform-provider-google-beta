// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Nathan is editing this to generate diffs in lots of files.
//     He won't submit this change.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudBuildTrigger_cloudbuildTriggerFilenameExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudBuildTriggerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildTrigger_cloudbuildTriggerFilenameExample(context),
			},
			{
				ResourceName:      "google_cloudbuild_trigger.filename-trigger",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudBuildTrigger_cloudbuildTriggerFilenameExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudbuild_trigger" "filename-trigger" {
  trigger_template {
    branch_name = "master"
    repo_name   = "my-repo"
  }

  substitutions = {
    _FOO = "bar"
    _BAZ = "qux"
  }

  filename = "cloudbuild.yaml"
}
`, context)
}

func testAccCheckCloudBuildTriggerDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_cloudbuild_trigger" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{CloudBuildBasePath}}projects/{{project}}/triggers/{{trigger_id}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("CloudBuildTrigger still exists at %s", url)
		}
	}

	return nil
}
