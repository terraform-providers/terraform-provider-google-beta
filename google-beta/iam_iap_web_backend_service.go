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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var IapWebBackendServiceIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"web_backend_service": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type IapWebBackendServiceIamUpdater struct {
	project           string
	webBackendService string
	d                 *schema.ResourceData
	Config            *Config
}

func IapWebBackendServiceIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	values["project"] = project
	if v, ok := d.GetOk("web_backend_service"); ok {
		values["web_backend_service"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_web/compute/services/(?P<web_backend_service>[^/]+)", "(?P<project>[^/]+)/(?P<web_backend_service>[^/]+)", "(?P<web_backend_service>[^/]+)"}, d, config, d.Get("web_backend_service").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapWebBackendServiceIamUpdater{
		project:           values["project"],
		webBackendService: values["web_backend_service"],
		d:                 d,
		Config:            config,
	}

	d.Set("project", u.project)
	d.Set("web_backend_service", u.GetResourceId())

	return u, nil
}

func IapWebBackendServiceIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	values["project"] = project

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_web/compute/services/(?P<web_backend_service>[^/]+)", "(?P<project>[^/]+)/(?P<web_backend_service>[^/]+)", "(?P<web_backend_service>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapWebBackendServiceIamUpdater{
		project:           values["project"],
		webBackendService: values["web_backend_service"],
		d:                 d,
		Config:            config,
	}
	d.Set("web_backend_service", u.GetResourceId())
	d.SetId(u.GetResourceId())
	return nil
}

func (u *IapWebBackendServiceIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyWebBackendServiceUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	obj = map[string]interface{}{
		"options": map[string]interface{}{
			"requestedPolicyVersion": iamPolicyVersion,
		},
	}

	policy, err := sendRequest(u.Config, "POST", project, url, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *IapWebBackendServiceIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyWebBackendServiceUrl("setIamPolicy")
	if err != nil {
		return err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", project, url, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *IapWebBackendServiceIamUpdater) qualifyWebBackendServiceUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{IapBasePath}}%s:%s", fmt.Sprintf("projects/%s/iap_web/compute/services/%s", u.project, u.webBackendService), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *IapWebBackendServiceIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/iap_web/compute/services/%s", u.project, u.webBackendService)
}

func (u *IapWebBackendServiceIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-iap-webbackendservice-%s", u.GetResourceId())
}

func (u *IapWebBackendServiceIamUpdater) DescribeResource() string {
	return fmt.Sprintf("iap webbackendservice %q", u.GetResourceId())
}
