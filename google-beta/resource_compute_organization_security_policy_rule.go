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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceComputeOrganizationSecurityPolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeOrganizationSecurityPolicyRuleCreate,
		Read:   resourceComputeOrganizationSecurityPolicyRuleRead,
		Update: resourceComputeOrganizationSecurityPolicyRuleUpdate,
		Delete: resourceComputeOrganizationSecurityPolicyRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeOrganizationSecurityPolicyRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The Action to perform when the client connection triggers the rule. Can currently be either
"allow", "deny" or "goto_next".`,
			},
			"match": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The configuration options for matching the rule.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"layer4_config": {
										Type:        schema.TypeList,
										Required:    true,
										Description: `Pairs of IP protocols and ports that the rule should match.`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_protocol": {
													Type:     schema.TypeString,
													Required: true,
													Description: `The IP protocol to which this rule applies. The protocol
type is required when creating a firewall rule.
This value can either be one of the following well
known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp),
or the IP protocol number.`,
												},
												"ports": {
													Type:     schema.TypeList,
													Optional: true,
													Description: `An optional list of ports to which this rule applies. This field
is only applicable for UDP or TCP protocol. Each entry must be
either an integer or a range. If not specified, this rule
applies to connections through any port.

Example inputs include: ["22"], ["80","443"], and
["12345-12349"].`,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"dest_ip_ranges": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `Destination IP address range in CIDR format. Required for
EGRESS rules.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										ExactlyOneOf: []string{"match.0.config.0.src_ip_ranges", "match.0.config.0.dest_ip_ranges"},
									},
									"src_ip_ranges": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `Source IP address range in CIDR format. Required for
INGRESS rules.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										ExactlyOneOf: []string{"match.0.config.0.src_ip_ranges", "match.0.config.0.dest_ip_ranges"},
									},
								},
							},
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A description of the rule.`,
						},
						"versioned_expr": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"FIREWALL", ""}, false),
							Description: `Preconfigured versioned expression. For organization security policy rules,
the only supported type is "FIREWALL". Default value: "FIREWALL" Possible values: ["FIREWALL"]`,
							Default: "FIREWALL",
						},
					},
				},
			},
			"policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the OrganizationSecurityPolicy this rule applies to.`,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				Description: `An integer indicating the priority of a rule in the list. The priority must be a value
between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
highest priority and 2147483647 is the lowest prority.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A description of the rule.`,
			},
			"direction": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"INGRESS", "EGRESS", ""}, false),
				Description:  `The direction in which this rule applies. If unspecified an INGRESS rule is created. Possible values: ["INGRESS", "EGRESS"]`,
			},
			"enable_logging": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Denotes whether to enable logging for a particular rule.
If logging is enabled, logs will be exported to the
configured export destination in Stackdriver.`,
			},
			"preview": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If set to true, the specified action is not enforced.`,
			},
			"target_resources": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of network resource URLs to which this rule applies.
This field allows you to control which network's VMs get
this rule. If this field is left blank, all VMs
within the organization will receive the rule.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_service_accounts": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of service accounts indicating the sets of
instances that are applied with this rule.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceComputeOrganizationSecurityPolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeOrganizationSecurityPolicyRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	priorityProp, err := expandComputeOrganizationSecurityPolicyRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !isEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	matchProp, err := expandComputeOrganizationSecurityPolicyRuleMatch(d.Get("match"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("match"); !isEmptyValue(reflect.ValueOf(matchProp)) && (ok || !reflect.DeepEqual(v, matchProp)) {
		obj["match"] = matchProp
	}
	actionProp, err := expandComputeOrganizationSecurityPolicyRuleAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !isEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	previewProp, err := expandComputeOrganizationSecurityPolicyRulePreview(d.Get("preview"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preview"); !isEmptyValue(reflect.ValueOf(previewProp)) && (ok || !reflect.DeepEqual(v, previewProp)) {
		obj["preview"] = previewProp
	}
	directionProp, err := expandComputeOrganizationSecurityPolicyRuleDirection(d.Get("direction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("direction"); !isEmptyValue(reflect.ValueOf(directionProp)) && (ok || !reflect.DeepEqual(v, directionProp)) {
		obj["direction"] = directionProp
	}
	targetResourcesProp, err := expandComputeOrganizationSecurityPolicyRuleTargetResources(d.Get("target_resources"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_resources"); !isEmptyValue(reflect.ValueOf(targetResourcesProp)) && (ok || !reflect.DeepEqual(v, targetResourcesProp)) {
		obj["targetResources"] = targetResourcesProp
	}
	enableLoggingProp, err := expandComputeOrganizationSecurityPolicyRuleEnableLogging(d.Get("enable_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_logging"); ok || !reflect.DeepEqual(v, enableLoggingProp) {
		obj["enableLogging"] = enableLoggingProp
	}
	targetServiceAccountsProp, err := expandComputeOrganizationSecurityPolicyRuleTargetServiceAccounts(d.Get("target_service_accounts"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_service_accounts"); !isEmptyValue(reflect.ValueOf(targetServiceAccountsProp)) && (ok || !reflect.DeepEqual(v, targetServiceAccountsProp)) {
		obj["targetServiceAccounts"] = targetServiceAccountsProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}{{policy_id}}/addRule?priority={{priority}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new OrganizationSecurityPolicyRule: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating OrganizationSecurityPolicyRule: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{policy_id}}/priority/{{priority}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating OrganizationSecurityPolicyRule %q: %#v", d.Id(), res)

	// `parent` is needed to poll the asynchronous operations but its available only on the policy.

	policyUrl := fmt.Sprintf("{{ComputeBasePath}}%s", d.Get("policy_id"))
	url, err = replaceVars(d, config, policyUrl)
	if err != nil {
		return err
	}

	policyRes, err := sendRequest(config, "GET", "", url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeOrganizationSecurityPolicy %q", d.Get("policy_id")))
	}

	parent := flattenComputeOrganizationSecurityPolicyParent(policyRes["parent"], d, config)
	var opRes map[string]interface{}
	err = computeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, parent.(string), "Creating OrganizationSecurityPolicyRule",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create OrganizationSecurityPolicyRule: %s", err)
	}

	return resourceComputeOrganizationSecurityPolicyRuleRead(d, meta)
}

func resourceComputeOrganizationSecurityPolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}{{policy_id}}/getRule?priority={{priority}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeOrganizationSecurityPolicyRule %q", d.Id()))
	}

	if err := d.Set("description", flattenComputeOrganizationSecurityPolicyRuleDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicyRule: %s", err)
	}
	if err := d.Set("priority", flattenComputeOrganizationSecurityPolicyRulePriority(res["priority"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicyRule: %s", err)
	}
	if err := d.Set("match", flattenComputeOrganizationSecurityPolicyRuleMatch(res["match"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicyRule: %s", err)
	}
	if err := d.Set("action", flattenComputeOrganizationSecurityPolicyRuleAction(res["action"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicyRule: %s", err)
	}
	if err := d.Set("preview", flattenComputeOrganizationSecurityPolicyRulePreview(res["preview"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicyRule: %s", err)
	}
	if err := d.Set("direction", flattenComputeOrganizationSecurityPolicyRuleDirection(res["direction"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicyRule: %s", err)
	}
	if err := d.Set("target_resources", flattenComputeOrganizationSecurityPolicyRuleTargetResources(res["targetResources"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicyRule: %s", err)
	}
	if err := d.Set("enable_logging", flattenComputeOrganizationSecurityPolicyRuleEnableLogging(res["enableLogging"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicyRule: %s", err)
	}
	if err := d.Set("target_service_accounts", flattenComputeOrganizationSecurityPolicyRuleTargetServiceAccounts(res["targetServiceAccounts"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicyRule: %s", err)
	}

	return nil
}

func resourceComputeOrganizationSecurityPolicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeOrganizationSecurityPolicyRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	priorityProp, err := expandComputeOrganizationSecurityPolicyRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	matchProp, err := expandComputeOrganizationSecurityPolicyRuleMatch(d.Get("match"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("match"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, matchProp)) {
		obj["match"] = matchProp
	}
	actionProp, err := expandComputeOrganizationSecurityPolicyRuleAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	previewProp, err := expandComputeOrganizationSecurityPolicyRulePreview(d.Get("preview"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preview"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, previewProp)) {
		obj["preview"] = previewProp
	}
	directionProp, err := expandComputeOrganizationSecurityPolicyRuleDirection(d.Get("direction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("direction"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, directionProp)) {
		obj["direction"] = directionProp
	}
	targetResourcesProp, err := expandComputeOrganizationSecurityPolicyRuleTargetResources(d.Get("target_resources"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_resources"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetResourcesProp)) {
		obj["targetResources"] = targetResourcesProp
	}
	enableLoggingProp, err := expandComputeOrganizationSecurityPolicyRuleEnableLogging(d.Get("enable_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_logging"); ok || !reflect.DeepEqual(v, enableLoggingProp) {
		obj["enableLogging"] = enableLoggingProp
	}
	targetServiceAccountsProp, err := expandComputeOrganizationSecurityPolicyRuleTargetServiceAccounts(d.Get("target_service_accounts"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_service_accounts"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetServiceAccountsProp)) {
		obj["targetServiceAccounts"] = targetServiceAccountsProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}{{policy_id}}/patchRule?priority={{priority}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating OrganizationSecurityPolicyRule %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating OrganizationSecurityPolicyRule %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating OrganizationSecurityPolicyRule %q: %#v", d.Id(), res)
	}

	// `parent` is needed to poll the asynchronous operations but its available only on the policy.

	policyUrl := fmt.Sprintf("{{ComputeBasePath}}%s", d.Get("policy_id"))
	url, err = replaceVars(d, config, policyUrl)
	if err != nil {
		return err
	}

	policyRes, err := sendRequest(config, "GET", "", url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeOrganizationSecurityPolicy %q", d.Get("policy_id")))
	}

	parent := flattenComputeOrganizationSecurityPolicyParent(policyRes["parent"], d, config)
	var opRes map[string]interface{}
	err = computeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, parent.(string), "Creating OrganizationSecurityPolicyRule",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create OrganizationSecurityPolicyRule: %s", err)
	}
	return resourceComputeOrganizationSecurityPolicyRuleRead(d, meta)
}

func resourceComputeOrganizationSecurityPolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	url, err := replaceVars(d, config, "{{ComputeBasePath}}{{policy_id}}/removeRule?priority={{priority}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting OrganizationSecurityPolicyRule %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "OrganizationSecurityPolicyRule")
	}

	// `parent` is needed to poll the asynchronous operations but its available only on the policy.

	policyUrl := fmt.Sprintf("{{ComputeBasePath}}%s", d.Get("policy_id"))
	url, err = replaceVars(d, config, policyUrl)
	if err != nil {
		return err
	}

	policyRes, err := sendRequest(config, "GET", "", url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeOrganizationSecurityPolicy %q", d.Get("policy_id")))
	}

	parent := flattenComputeOrganizationSecurityPolicyParent(policyRes["parent"], d, config)
	var opRes map[string]interface{}
	err = computeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, parent.(string), "Creating OrganizationSecurityPolicyRule",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create OrganizationSecurityPolicyRule: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting OrganizationSecurityPolicyRule %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeOrganizationSecurityPolicyRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<policy_id>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("policy_id").(string), "/")
	if len(nameParts) != 6 {
		return nil, fmt.Errorf(
			"Saw %s when the import ID is expected to have shape %s",
			d.Get("policy_id").(string),
			"locations/global/securityPolicies/{{policy_id}}/priority/{{priority}}",
		)
	}
	if err := d.Set("policy_id", fmt.Sprintf("locations/global/securityPolicies/%s", nameParts[3])); err != nil {
		return nil, fmt.Errorf("Error setting policy_id: %s", err)
	}

	if prio, err := strconv.ParseInt(nameParts[5], 10, 64); err != nil {
		return nil, fmt.Errorf(
			"Priority %s cannot be converted to integer", nameParts[5],
		)
	} else {
		if err := d.Set("priority", prio); err != nil {
			return nil, fmt.Errorf("Error setting priority: %s", err)
		}
	}

	return []*schema.ResourceData{d}, nil
}

func flattenComputeOrganizationSecurityPolicyRuleDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRulePriority(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeOrganizationSecurityPolicyRuleMatch(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["description"] =
		flattenComputeOrganizationSecurityPolicyRuleMatchDescription(original["description"], d, config)
	transformed["versioned_expr"] =
		flattenComputeOrganizationSecurityPolicyRuleMatchVersionedExpr(original["versionedExpr"], d, config)
	transformed["config"] =
		flattenComputeOrganizationSecurityPolicyRuleMatchConfig(original["config"], d, config)
	return []interface{}{transformed}
}
func flattenComputeOrganizationSecurityPolicyRuleMatchDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRuleMatchVersionedExpr(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRuleMatchConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["src_ip_ranges"] =
		flattenComputeOrganizationSecurityPolicyRuleMatchConfigSrcIpRanges(original["srcIpRanges"], d, config)
	transformed["dest_ip_ranges"] =
		flattenComputeOrganizationSecurityPolicyRuleMatchConfigDestIpRanges(original["destIpRanges"], d, config)
	transformed["layer4_config"] =
		flattenComputeOrganizationSecurityPolicyRuleMatchConfigLayer4Config(original["layer4Configs"], d, config)
	return []interface{}{transformed}
}
func flattenComputeOrganizationSecurityPolicyRuleMatchConfigSrcIpRanges(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRuleMatchConfigDestIpRanges(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRuleMatchConfigLayer4Config(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"ip_protocol": flattenComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigIpProtocol(original["ipProtocol"], d, config),
			"ports":       flattenComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigPorts(original["ports"], d, config),
		})
	}
	return transformed
}
func flattenComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigIpProtocol(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigPorts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRuleAction(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRulePreview(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRuleDirection(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRuleTargetResources(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRuleEnableLogging(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyRuleTargetServiceAccounts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeOrganizationSecurityPolicyRuleDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRulePriority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatch(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDescription, err := expandComputeOrganizationSecurityPolicyRuleMatchDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedVersionedExpr, err := expandComputeOrganizationSecurityPolicyRuleMatchVersionedExpr(original["versioned_expr"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersionedExpr); val.IsValid() && !isEmptyValue(val) {
		transformed["versionedExpr"] = transformedVersionedExpr
	}

	transformedConfig, err := expandComputeOrganizationSecurityPolicyRuleMatchConfig(original["config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["config"] = transformedConfig
	}

	return transformed, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchVersionedExpr(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSrcIpRanges, err := expandComputeOrganizationSecurityPolicyRuleMatchConfigSrcIpRanges(original["src_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcIpRanges); val.IsValid() && !isEmptyValue(val) {
		transformed["srcIpRanges"] = transformedSrcIpRanges
	}

	transformedDestIpRanges, err := expandComputeOrganizationSecurityPolicyRuleMatchConfigDestIpRanges(original["dest_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestIpRanges); val.IsValid() && !isEmptyValue(val) {
		transformed["destIpRanges"] = transformedDestIpRanges
	}

	transformedLayer4Config, err := expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4Config(original["layer4_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLayer4Config); val.IsValid() && !isEmptyValue(val) {
		transformed["layer4Configs"] = transformedLayer4Config
	}

	return transformed, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfigSrcIpRanges(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfigDestIpRanges(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4Config(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpProtocol, err := expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigIpProtocol(original["ip_protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpProtocol); val.IsValid() && !isEmptyValue(val) {
			transformed["ipProtocol"] = transformedIpProtocol
		}

		transformedPorts, err := expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !isEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigIpProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigPorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleAction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRulePreview(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleDirection(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleTargetResources(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleEnableLogging(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleTargetServiceAccounts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
