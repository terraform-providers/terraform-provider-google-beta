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
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceArtifactRegistryRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceArtifactRegistryRepositoryCreate,
		Read:   resourceArtifactRegistryRepositoryRead,
		Update: resourceArtifactRegistryRepositoryUpdate,
		Delete: resourceArtifactRegistryRepositoryDelete,

		Importer: &schema.ResourceImporter{
			State: resourceArtifactRegistryRepositoryImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"format": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"DOCKER"}, false),
				Description:  `The format of packages that are stored in the repoitory. Possible values: ["DOCKER"]`,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The last part of the repository name, for example:
"repo1"`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The user-provided description of the repository.`,
			},
			"kms_key_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The Cloud KMS resource name of the customer managed encryption key that’s
used to encrypt the contents of the Repository. Has the form:
'projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key'.
This value may not be changed after the Repository has been created.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels with user-defined metadata.
This field may contain up to 64 entries. Label keys and values may be no
longer than 63 characters. Label keys must begin with a lowercase letter
and may only contain lowercase letters, numeric characters, underscores,
and dashes.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the location this repository is located in.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the repository was created.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The name of the repository, for example:
"projects/p1/locations/us-central1/repositories/repo1"`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the repository was last updated.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceArtifactRegistryRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	formatProp, err := expandArtifactRegistryRepositoryFormat(d.Get("format"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("format"); !isEmptyValue(reflect.ValueOf(formatProp)) && (ok || !reflect.DeepEqual(v, formatProp)) {
		obj["format"] = formatProp
	}
	descriptionProp, err := expandArtifactRegistryRepositoryDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandArtifactRegistryRepositoryLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	kmsKeyNameProp, err := expandArtifactRegistryRepositoryKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_key_name"); !isEmptyValue(reflect.ValueOf(kmsKeyNameProp)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}

	url, err := replaceVars(d, config, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/repositories?repository_id={{repository_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Repository: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Repository: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = artifactRegistryOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Repository", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Repository: %s", err)
	}

	if err := d.Set("name", flattenArtifactRegistryRepositoryName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Repository %q: %#v", d.Id(), res)

	return resourceArtifactRegistryRepositoryRead(d, meta)
}

func resourceArtifactRegistryRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ArtifactRegistryRepository %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}

	if err := d.Set("name", flattenArtifactRegistryRepositoryName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("format", flattenArtifactRegistryRepositoryFormat(res["format"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("description", flattenArtifactRegistryRepositoryDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("labels", flattenArtifactRegistryRepositoryLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("kms_key_name", flattenArtifactRegistryRepositoryKmsKeyName(res["kmsKeyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("create_time", flattenArtifactRegistryRepositoryCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("update_time", flattenArtifactRegistryRepositoryUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}

	return nil
}

func resourceArtifactRegistryRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	obj := make(map[string]interface{})
	formatProp, err := expandArtifactRegistryRepositoryFormat(d.Get("format"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("format"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, formatProp)) {
		obj["format"] = formatProp
	}
	descriptionProp, err := expandArtifactRegistryRepositoryDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandArtifactRegistryRepositoryLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	kmsKeyNameProp, err := expandArtifactRegistryRepositoryKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_key_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}

	url, err := replaceVars(d, config, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Repository %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PUT", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Repository %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Repository %q: %#v", d.Id(), res)
	}

	err = artifactRegistryOperationWaitTime(
		config, res, project, "Updating Repository", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceArtifactRegistryRepositoryRead(d, meta)
}

func resourceArtifactRegistryRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Repository %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Repository")
	}

	err = artifactRegistryOperationWaitTime(
		config, res, project, "Deleting Repository", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Repository %q: %#v", d.Id(), res)
	return nil
}

func resourceArtifactRegistryRepositoryImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/repositories/(?P<repository_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<repository_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<repository_id>[^/]+)",
		"(?P<repository_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenArtifactRegistryRepositoryName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenArtifactRegistryRepositoryFormat(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenArtifactRegistryRepositoryDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenArtifactRegistryRepositoryLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenArtifactRegistryRepositoryKmsKeyName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenArtifactRegistryRepositoryCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenArtifactRegistryRepositoryUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandArtifactRegistryRepositoryFormat(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandArtifactRegistryRepositoryKmsKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
