package license

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func jsonLicense() string {
	return `{"valid":true,"evaluation":true,"maximumNumberOfUsers":-1,"licenseType":"Commercial","creationDateString":"14/Feb/23","expiryDate":1678971600000,"expiryDateString":"16/Mar/23","organizationName":"sebastian@sommerfeld.io","dataCenter":true,"subscription":true,"rawLicense":"THE_ACTUAL_LICENSE","expired":false,"supportEntitlementNumber":"SEN-L19188898","enterprise":false,"active":true,"autoRenewal":false}`
}

func Test_ShouldGetLicense(t *testing.T) {
	testTable := []struct {
		name             string
		server           *httptest.Server
		expectedResponse *JiraLicense
		expectedErr      error
	}{
		{
			name: "ShouldGetLicense",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r.SetBasicAuth("admin", "admin")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(jsonLicense()))
			})),
			expectedResponse: &JiraLicense{
				Valid:                    true,
				Evaluation:               true,
				MaximumNumberOfUsers:     -1,
				LicenseType:              "Commercial",
				CreationDateString:       "14/Feb/23",
				ExpiryDate:               1678971600000,
				ExpiryDateString:         "16/Mar/23",
				OrganizationName:         "sebastian@sommerfeld.io",
				DataCenter:               true,
				Subscription:             true,
				RawLicense:               "THE_ACTUAL_LICENSE",
				Expired:                  false,
				SupportEntitlementNumber: "SEN-L19188898",
				Enterprise:               false,
				Active:                   true,
				AutoRenewal:              false,
				RawJson:                  jsonLicense(),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			result, err := ReadJiraLicense(tc.server.URL, "", "")

			assert.Nil(t, err)
			assert.True(t, json.Valid([]byte(result.RawJson)), "JSON should be valid")
		})
	}
}

func Test_ShouldGetError(t *testing.T) {
	testTable := []struct {
		name             string
		server           *httptest.Server
		expectedResponse *JiraLicense
		expectedErr      error
	}{
		{
			name: "ShouldGetUnauthorizedError",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r.SetBasicAuth("cloud", "cloud")
				w.WriteHeader(http.StatusUnauthorized)
			})),
			expectedResponse: &JiraLicense{},
			expectedErr:      errors.New("must have permission to access this resource"),
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			result, err := ReadJiraLicense(tc.server.URL, "", "")

			assert.NotNil(t, err)
			assert.EqualErrorf(t, err, tc.expectedErr.Error(), "Error should be: %v, got: %v", tc.expectedErr.Error(), err)
			assert.Equal(t, JiraLicense{}, result)
		})
	}
}
