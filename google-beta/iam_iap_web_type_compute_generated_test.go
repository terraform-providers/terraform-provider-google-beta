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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIapWebTypeComputeIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapWebTypeComputeIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapWebTypeComputeIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor %s", fmt.Sprintf("tf-test%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor %s", fmt.Sprintf("tf-test%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", fmt.Sprintf("tf-test%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", fmt.Sprintf("tf-test%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeComputeIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeComputeIamPolicy_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_compute_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapWebTypeComputeIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"
}

resource "google_iap_web_type_compute_iam_member" "foo" {
  project = "${google_project_service.project_service.project}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapWebTypeComputeIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_iap_web_type_compute_iam_policy" "foo" {
  project = "${google_project_service.project_service.project}"
  policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}

func testAccIapWebTypeComputeIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"
}

resource "google_iap_web_type_compute_iam_binding" "foo" {
  project = "${google_project_service.project_service.project}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapWebTypeComputeIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"
}

resource "google_iap_web_type_compute_iam_binding" "foo" {
  project = "${google_project_service.project_service.project}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}

func testAccIapWebTypeComputeIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"
}

resource "google_iap_web_type_compute_iam_binding" "foo" {
  project = "${google_project_service.project_service.project}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeComputeIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"
}

resource "google_iap_web_type_compute_iam_binding" "foo" {
  project = "${google_project_service.project_service.project}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_iap_web_type_compute_iam_binding" "foo2" {
  project = "${google_project_service.project_service.project}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeComputeIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"
}

resource "google_iap_web_type_compute_iam_member" "foo" {
  project = "${google_project_service.project_service.project}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeComputeIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"
}

resource "google_iap_web_type_compute_iam_member" "foo" {
  project = "${google_project_service.project_service.project}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_iap_web_type_compute_iam_member" "foo2" {
  project = "${google_project_service.project_service.project}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeComputeIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "Expiring at midnight of 2019-12-31"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_iap_web_type_compute_iam_policy" "foo" {
  project = "${google_project_service.project_service.project}"
  policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}
