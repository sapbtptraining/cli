package ccv3

import (
	"bytes"
	"encoding/json"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
)

// OrganizationQuota represents a Cloud Controller organization quota.
type OrganizationQuota struct {
	Quota
}

func (client *Client) ApplyOrganizationQuota(quotaGuid, orgGuid string) (RelationshipList, Warnings, error) {

	orgs := RelationshipList{
		[]string{orgGuid},
	}

	orgBytes, err := json.Marshal(orgs)

	if err != nil {
		return RelationshipList{}, nil, err
	}
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostOrganizationQuotaApplyRequest,
		URIParams:   internal.Params{"quota_guid": quotaGuid},
		Body:        bytes.NewReader(orgBytes),
	})
	if err != nil {
		return RelationshipList{}, nil, err
	}

	var responseOrgs RelationshipList
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &responseOrgs,
	}

	err = client.connection.Make(request, &response)
	if err != nil {
		return RelationshipList{}, response.Warnings, err
	}

	return responseOrgs, response.Warnings, err
}

func (client *Client) CreateOrganizationQuota(orgQuota OrganizationQuota) (OrganizationQuota, Warnings, error) {
	var responseOrgQuota OrganizationQuota

	_, warnings, err := client.makeRequest(requestParams{
		RequestName:  internal.PostOrganizationQuotaRequest,
		RequestBody:  orgQuota,
		ResponseBody: &responseOrgQuota,
	})

	return responseOrgQuota, warnings, err
}

func (client *Client) DeleteOrganizationQuota(quotaGUID string) (JobURL, Warnings, error) {
	jobURL, warnings, err := client.makeRequest(requestParams{
		RequestName: internal.DeleteOrganizationQuotaRequest,
		URIParams:   internal.Params{"quota_guid": quotaGUID},
	})

	return jobURL, warnings, err
}

func (client *Client) GetOrganizationQuota(quotaGUID string) (OrganizationQuota, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetOrganizationQuotaRequest,
		URIParams:   internal.Params{"quota_guid": quotaGUID},
	})
	if err != nil {
		return OrganizationQuota{}, nil, err
	}
	var responseOrgQuota OrganizationQuota
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &responseOrgQuota,
	}
	err = client.connection.Make(request, &response)
	if err != nil {
		return OrganizationQuota{}, response.Warnings, err
	}

	return responseOrgQuota, response.Warnings, nil
}

func (client *Client) GetOrganizationQuotas(query ...Query) ([]OrganizationQuota, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetOrganizationQuotasRequest,
		Query:       query,
	})
	if err != nil {
		return []OrganizationQuota{}, nil, err
	}

	var orgQuotasList []OrganizationQuota
	warnings, err := client.paginate(request, OrganizationQuota{}, func(item interface{}) error {
		if orgQuota, ok := item.(OrganizationQuota); ok {
			orgQuotasList = append(orgQuotasList, orgQuota)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   OrganizationQuota{},
				Unexpected: item,
			}
		}
		return nil
	})

	return orgQuotasList, warnings, err
}

func (client *Client) UpdateOrganizationQuota(orgQuota OrganizationQuota) (OrganizationQuota, Warnings, error) {
	orgQuotaGUID := orgQuota.GUID
	orgQuota.GUID = ""

	quotaBytes, err := json.Marshal(orgQuota)
	if err != nil {
		return OrganizationQuota{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PatchOrganizationQuotaRequest,
		URIParams:   internal.Params{"quota_guid": orgQuotaGUID},
		Body:        bytes.NewReader(quotaBytes),
	})
	if err != nil {
		return OrganizationQuota{}, nil, err
	}

	var responseOrgQuota OrganizationQuota
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &responseOrgQuota,
	}

	err = client.connection.Make(request, &response)
	if err != nil {
		return OrganizationQuota{}, response.Warnings, err
	}

	return responseOrgQuota, response.Warnings, nil
}
