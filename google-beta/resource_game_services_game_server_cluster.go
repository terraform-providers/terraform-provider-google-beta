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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func suppressSuffixDiff(_, old, new string, _ *schema.ResourceData) bool {
	if strings.HasSuffix(old, new) {
		log.Printf("[INFO] suppressing diff as %s is the same as the full path of %s", new, old)
		return true
	}

	return false
}

func resourceGameServicesGameServerCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceGameServicesGameServerClusterCreate,
		Read:   resourceGameServicesGameServerClusterRead,
		Update: resourceGameServicesGameServerClusterUpdate,
		Delete: resourceGameServicesGameServerClusterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGameServicesGameServerClusterImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. The resource name of the game server cluster`,
			},
			"connection_info": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Description: `Game server cluster connection information. This information is used to
manage game server clusters.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gke_cluster_reference": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `Reference of the GKE cluster where the game servers are installed.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: suppressSuffixDiff,
										Description: `The full or partial name of a GKE cluster, using one of the following
forms:

* 'projects/{project_id}/locations/{location}/clusters/{cluster_id}'
* 'locations/{location}/clusters/{cluster_id}'
* '{cluster_id}'

If project and location are not specified, the project and location of the
GameServerCluster resource are used to generate the full name of the
GKE cluster.`,
									},
								},
							},
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Namespace designated on the game server cluster where the game server
instances will be created. The namespace existence will be validated
during creation.`,
						},
					},
				},
			},
			"realm_id": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The realm id of the game server realm.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Human readable description of the cluster.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels associated with this game server cluster. Each label is a
key-value pair.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Location of the Cluster.`,
				Default:     "global",
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource id of the game server cluster, eg:

'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'.
For example,

'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.`,
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

func resourceGameServicesGameServerClusterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	labelsProp, err := expandGameServicesGameServerClusterLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	connectionInfoProp, err := expandGameServicesGameServerClusterConnectionInfo(d.Get("connection_info"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connection_info"); !isEmptyValue(reflect.ValueOf(connectionInfoProp)) && (ok || !reflect.DeepEqual(v, connectionInfoProp)) {
		obj["connectionInfo"] = connectionInfoProp
	}
	descriptionProp, err := expandGameServicesGameServerClusterDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := replaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters?gameServerClusterId={{cluster_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GameServerCluster: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GameServerCluster: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = gameServicesOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating GameServerCluster",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GameServerCluster: %s", err)
	}

	if err := d.Set("name", flattenGameServicesGameServerClusterName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating GameServerCluster %q: %#v", d.Id(), res)

	return resourceGameServicesGameServerClusterRead(d, meta)
}

func resourceGameServicesGameServerClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("GameServicesGameServerCluster %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GameServerCluster: %s", err)
	}

	if err := d.Set("name", flattenGameServicesGameServerClusterName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerCluster: %s", err)
	}
	if err := d.Set("labels", flattenGameServicesGameServerClusterLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerCluster: %s", err)
	}
	if err := d.Set("connection_info", flattenGameServicesGameServerClusterConnectionInfo(res["connectionInfo"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerCluster: %s", err)
	}
	if err := d.Set("description", flattenGameServicesGameServerClusterDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerCluster: %s", err)
	}

	return nil
}

func resourceGameServicesGameServerClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandGameServicesGameServerClusterLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandGameServicesGameServerClusterDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := replaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating GameServerCluster %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating GameServerCluster %q: %s", d.Id(), err)
	}

	err = gameServicesOperationWaitTime(
		config, res, project, "Updating GameServerCluster",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceGameServicesGameServerClusterRead(d, meta)
}

func resourceGameServicesGameServerClusterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GameServerCluster %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GameServerCluster")
	}

	err = gameServicesOperationWaitTime(
		config, res, project, "Deleting GameServerCluster",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GameServerCluster %q: %#v", d.Id(), res)
	return nil
}

func resourceGameServicesGameServerClusterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/realms/(?P<realm_id>[^/]+)/gameServerClusters/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<realm_id>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<realm_id>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGameServicesGameServerClusterName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerClusterLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerClusterConnectionInfo(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["gke_cluster_reference"] =
		flattenGameServicesGameServerClusterConnectionInfoGkeClusterReference(original["gkeClusterReference"], d, config)
	transformed["namespace"] =
		flattenGameServicesGameServerClusterConnectionInfoNamespace(original["namespace"], d, config)
	return []interface{}{transformed}
}
func flattenGameServicesGameServerClusterConnectionInfoGkeClusterReference(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["cluster"] =
		flattenGameServicesGameServerClusterConnectionInfoGkeClusterReferenceCluster(original["cluster"], d, config)
	return []interface{}{transformed}
}
func flattenGameServicesGameServerClusterConnectionInfoGkeClusterReferenceCluster(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerClusterConnectionInfoNamespace(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerClusterDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandGameServicesGameServerClusterLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGameServicesGameServerClusterConnectionInfo(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGkeClusterReference, err := expandGameServicesGameServerClusterConnectionInfoGkeClusterReference(original["gke_cluster_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGkeClusterReference); val.IsValid() && !isEmptyValue(val) {
		transformed["gkeClusterReference"] = transformedGkeClusterReference
	}

	transformedNamespace, err := expandGameServicesGameServerClusterConnectionInfoNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandGameServicesGameServerClusterConnectionInfoGkeClusterReference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCluster, err := expandGameServicesGameServerClusterConnectionInfoGkeClusterReferenceCluster(original["cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCluster); val.IsValid() && !isEmptyValue(val) {
		transformed["cluster"] = transformedCluster
	}

	return transformed, nil
}

func expandGameServicesGameServerClusterConnectionInfoGkeClusterReferenceCluster(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerClusterConnectionInfoNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerClusterDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
