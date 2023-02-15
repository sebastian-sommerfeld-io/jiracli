package license

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// ReadJiraSoftwareLicense reads the license used to allow Jira Software to work. The license
// is used for Jira itself, not for any plugin.
func ReadJiraSoftwareLicense(baseurl string, user string, pass string) (string, error) {
	endpoint := "/rest/plugins/applications/1.0/installed/jira-software/license"
	url := baseurl + endpoint
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	result := string(body)

	unauthorizedMsg := "must have permission to access this resource"
	if strings.Contains(strings.ToLower(result), unauthorizedMsg) {
		return "", errors.New(unauthorizedMsg)
	}

	return result, nil
}
