package license

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// JiraLicense represents a license object as returned from the Jira Rest API.
type JiraLicense struct {
	Valid                    bool   `json:"valid"`
	Evaluation               bool   `json:"evaluation"`
	MaximumNumberOfUsers     int    `json:"maximumNumberOfUsers"`
	LicenseType              string `json:"licenseType"`
	CreationDateString       string `json:"creationDateString"`
	ExpiryDate               int    `json:"expiryDate"`
	ExpiryDateString         string `json:"expiryDateString"`
	OrganizationName         string `json:"organizationName"`
	DataCenter               bool   `json:"dataCenter"`
	Subscription             bool   `json:"subscription"`
	RawLicense               string `json:"rawLicense"`
	Expired                  bool   `json:"expired"`
	SupportEntitlementNumber string `json:"supportEntitlementNumber"`
	Enterprise               bool   `json:"enterprise"`
	Active                   bool   `json:"active"`
	AutoRenewal              bool   `json:"autoRenewal"`
	RawJson                  string
}

// ReadJiraLicense reads the license used to allow Jira Software to work. The license
// is used for Jira itself, not for any plugin.
func ReadJiraLicense(baseurl string, user string, pass string) (JiraLicense, error) {
	endpoint := "/rest/plugins/applications/1.0/installed/jira-software/license"
	url := baseurl + endpoint
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return errorObjects(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)

	res, err := client.Do(req)
	if err != nil {
		return errorObjects(err)
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if res.StatusCode == 401 {
		return errorObjects(errors.New("must have permission to access this resource"))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errorObjects(err)
	}

	bodyString := string(body)

	result := &JiraLicense{}
	if err := json.Unmarshal(body, result); err != nil {
		return errorObjects(err)
	}
	result.RawJson = bodyString

	return *result, nil
}

func errorObjects(err error) (JiraLicense, error) {
	return JiraLicense{}, err
}
