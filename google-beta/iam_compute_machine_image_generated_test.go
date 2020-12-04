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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeMachineImageIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccComputeMachineImageIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeMachineImageIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccComputeMachineImageIamPolicy_emptyBinding(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamBinding_withConditionGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamBinding_withAndWithoutConditionGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamMember_withConditionGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamMember_withAndWithoutConditionGenerated(context),
			},
		},
	})
}

func TestAccComputeMachineImageIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.admin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImageIamPolicy_withConditionGenerated(context),
			},
		},
	})
}

func testAccComputeMachineImageIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_member" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeMachineImageIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_machine_image_iam_policy" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeMachineImageIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_compute_machine_image_iam_policy" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeMachineImageIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_binding" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeMachineImageIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_binding" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}

func testAccComputeMachineImageIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_binding" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
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

func testAccComputeMachineImageIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_binding" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_compute_machine_image_iam_binding" "foo2" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
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

func testAccComputeMachineImageIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_member" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
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

func testAccComputeMachineImageIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_machine_image_iam_member" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_compute_machine_image_iam_member" "foo2" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
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

func testAccComputeMachineImageIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}

data "google_iam_policy" "foo" {
  provider = google-beta
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

resource "google_compute_machine_image_iam_policy" "foo" {
  provider = google-beta
  project = google_compute_machine_image.image.project
  machine_image = google_compute_machine_image.image.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
