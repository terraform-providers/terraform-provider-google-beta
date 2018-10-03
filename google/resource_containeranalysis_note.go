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
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceContainerAnalysisNote() *schema.Resource {
	return &schema.Resource{
		Create: resourceContainerAnalysisNoteCreate,
		Read:   resourceContainerAnalysisNoteRead,
		Update: resourceContainerAnalysisNoteUpdate,
		Delete: resourceContainerAnalysisNoteDelete,

		Importer: &schema.ResourceImporter{
			State: resourceContainerAnalysisNoteImport,
		},

		Schema: map[string]*schema.Schema{
			"attestation_authority": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hint": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"human_readable_name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func resourceContainerAnalysisNoteCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandContainerAnalysisNoteName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	attestationAuthorityProp, err := expandContainerAnalysisNoteAttestationAuthority(d.Get("attestation_authority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation_authority"); !isEmptyValue(reflect.ValueOf(attestationAuthorityProp)) && (ok || !reflect.DeepEqual(v, attestationAuthorityProp)) {
		obj["attestationAuthority"] = attestationAuthorityProp
	}

	url, err := replaceVars(d, config, "https://containeranalysis.googleapis.com/v1beta1/projects/{{project}}/notes?noteId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Note: %#v", obj)
	res, err := sendRequest(config, "POST", url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Note: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Note %q: %#v", d.Id(), res)

	return resourceContainerAnalysisNoteRead(d, meta)
}

func resourceContainerAnalysisNoteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://containeranalysis.googleapis.com/v1beta1/projects/{{project}}/notes/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ContainerAnalysisNote %q", d.Id()))
	}

	if err := d.Set("name", flattenContainerAnalysisNoteName(res["name"])); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	if err := d.Set("attestation_authority", flattenContainerAnalysisNoteAttestationAuthority(res["attestationAuthority"])); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}

	return nil
}

func resourceContainerAnalysisNoteUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandContainerAnalysisNoteName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	attestationAuthorityProp, err := expandContainerAnalysisNoteAttestationAuthority(d.Get("attestation_authority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation_authority"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, attestationAuthorityProp)) {
		obj["attestationAuthority"] = attestationAuthorityProp
	}

	url, err := replaceVars(d, config, "https://containeranalysis.googleapis.com/v1beta1/projects/{{project}}/notes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Note %q: %#v", d.Id(), obj)
	updateMask := []string{}
	if d.HasChange("attestation_authority.0.hint.0.human_readable_name") {
		updateMask = append(updateMask, "attestationAuthority.hint.humanReadableName")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequest(config, "PATCH", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating Note %q: %s", d.Id(), err)
	}

	return resourceContainerAnalysisNoteRead(d, meta)
}

func resourceContainerAnalysisNoteDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://containeranalysis.googleapis.com/v1beta1/projects/{{project}}/notes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Note %q", d.Id())
	res, err := sendRequest(config, "DELETE", url, obj)
	if err != nil {
		return handleNotFoundError(err, d, "Note")
	}

	log.Printf("[DEBUG] Finished deleting Note %q: %#v", d.Id(), res)
	return nil
}

func resourceContainerAnalysisNoteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/notes/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenContainerAnalysisNoteName(v interface{}) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenContainerAnalysisNoteAttestationAuthority(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["hint"] =
		flattenContainerAnalysisNoteAttestationAuthorityHint(original["hint"])
	return []interface{}{transformed}
}
func flattenContainerAnalysisNoteAttestationAuthorityHint(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["human_readable_name"] =
		flattenContainerAnalysisNoteAttestationAuthorityHintHumanReadableName(original["humanReadableName"])
	return []interface{}{transformed}
}
func flattenContainerAnalysisNoteAttestationAuthorityHintHumanReadableName(v interface{}) interface{} {
	return v
}

func expandContainerAnalysisNoteName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisNoteAttestationAuthority(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHint, err := expandContainerAnalysisNoteAttestationAuthorityHint(original["hint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHint); val.IsValid() && !isEmptyValue(val) {
		transformed["hint"] = transformedHint
	}

	return transformed, nil
}

func expandContainerAnalysisNoteAttestationAuthorityHint(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHumanReadableName, err := expandContainerAnalysisNoteAttestationAuthorityHintHumanReadableName(original["human_readable_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHumanReadableName); val.IsValid() && !isEmptyValue(val) {
		transformed["humanReadableName"] = transformedHumanReadableName
	}

	return transformed, nil
}

func expandContainerAnalysisNoteAttestationAuthorityHintHumanReadableName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}
