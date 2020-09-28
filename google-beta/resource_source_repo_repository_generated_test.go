// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSourceRepoRepository_sourcerepoRepositoryBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckSourceRepoRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSourceRepoRepository_sourcerepoRepositoryBasicExample(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository.my-repo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSourceRepoRepository_sourcerepoRepositoryBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}
`, context)
}

func TestAccSourceRepoRepository_sourcerepoRepositoryFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckSourceRepoRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSourceRepoRepository_sourcerepoRepositoryFullExample(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository.my-repo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSourceRepoRepository_sourcerepoRepositoryFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_service_account" "test-account" {
  account_id   = "tf-test-my-account%{random_suffix}"
  display_name = "Test Service Account"
}

resource "google_pubsub_topic" "topic" {
  name     = "tf-test-my-topic%{random_suffix}"
}

resource "google_sourcerepo_repository" "my-repo" {
  name = "tf-test-my-repository%{random_suffix}"
  pubsub_configs {
      topic = google_pubsub_topic.topic.id
      message_format = "JSON"
      service_account_email = google_service_account.test-account.email
  }
}
`, context)
}

func testAccCheckSourceRepoRepositoryDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_sourcerepo_repository" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{SourceRepoBasePath}}projects/{{project}}/repos/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("SourceRepoRepository still exists at %s", url)
			}
		}

		return nil
	}
}
