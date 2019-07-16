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
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/compute/v1"
)

func resourceComputeNodeGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNodeGroupCreate,
		Read:   resourceComputeNodeGroupRead,
		Update: resourceComputeNodeGroupUpdate,
		Delete: resourceComputeNodeGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNodeGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"node_template": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"size": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeNodeGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeNodeGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeNodeGroupName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	nodeTemplateProp, err := expandComputeNodeGroupNodeTemplate(d.Get("node_template"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_template"); !isEmptyValue(reflect.ValueOf(nodeTemplateProp)) && (ok || !reflect.DeepEqual(v, nodeTemplateProp)) {
		obj["nodeTemplate"] = nodeTemplateProp
	}
	sizeProp, err := expandComputeNodeGroupSize(d.Get("size"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("size"); ok || !reflect.DeepEqual(v, sizeProp) {
		obj["size"] = sizeProp
	}
	zoneProp, err := expandComputeNodeGroupZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/nodeGroups?initialNodeCount={{size}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NodeGroup: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating NodeGroup: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating NodeGroup",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NodeGroup: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating NodeGroup %q: %#v", d.Id(), res)

	return resourceComputeNodeGroupRead(d, meta)
}

func resourceComputeNodeGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeNodeGroup %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NodeGroup: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeNodeGroupCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading NodeGroup: %s", err)
	}
	if err := d.Set("description", flattenComputeNodeGroupDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading NodeGroup: %s", err)
	}
	if err := d.Set("name", flattenComputeNodeGroupName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading NodeGroup: %s", err)
	}
	if err := d.Set("node_template", flattenComputeNodeGroupNodeTemplate(res["nodeTemplate"], d)); err != nil {
		return fmt.Errorf("Error reading NodeGroup: %s", err)
	}
	if err := d.Set("size", flattenComputeNodeGroupSize(res["size"], d)); err != nil {
		return fmt.Errorf("Error reading NodeGroup: %s", err)
	}
	if err := d.Set("zone", flattenComputeNodeGroupZone(res["zone"], d)); err != nil {
		return fmt.Errorf("Error reading NodeGroup: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading NodeGroup: %s", err)
	}

	return nil
}

func resourceComputeNodeGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	d.Partial(true)

	if d.HasChange("node_template") {
		obj := make(map[string]interface{})
		nodeTemplateProp, err := expandComputeNodeGroupNodeTemplate(d.Get("node_template"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("node_template"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nodeTemplateProp)) {
			obj["nodeTemplate"] = nodeTemplateProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}/setNodeTemplate")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating NodeGroup %q: %s", d.Id(), err)
		}

		project, err := getProject(d, config)
		if err != nil {
			return err
		}
		op := &compute.Operation{}
		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating NodeGroup",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("node_template")
	}

	d.Partial(false)

	return resourceComputeNodeGroupRead(d, meta)
}

func resourceComputeNodeGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting NodeGroup %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "NodeGroup")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting NodeGroup",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NodeGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeNodeGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/nodeGroups/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeNodeGroupCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeNodeGroupDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeNodeGroupName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeNodeGroupNodeTemplate(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeNodeGroupSize(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeNodeGroupZone(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandComputeNodeGroupDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupNodeTemplate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("nodeTemplates", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for node_template: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeNodeGroupSize(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
