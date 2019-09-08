package accountapi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
	"os"
	"strconv"
)

const defaultBaseURL = "http://localhost:8080"

// AccountClient is used to invoke Form3 Accounts API.
type AccountClient struct {
	client *resty.Client
}

// NewAccountClient returns a new instance of AccountClient.
func NewAccountClient() *AccountClient {
	client := resty.New()
	client.SetDebug(false)
	// Try getting Accounts API base URL from env var
	apiURL := os.Getenv("API_ADDR")
	if apiURL == "" {
		apiURL = defaultBaseURL
	}
	client.SetHostURL(apiURL)
	// Setting global error struct that maps to Form3's error response
	client.SetError(&Error{})

	return &AccountClient{client: client}
}

// Create registers an existing account or creates a new one.
func (a *AccountClient) Create(accountParams AccountParams) (*Resource, error) {

	resp, err := a.client.R().
		SetResult(&Resource{}).
		SetBody(map[string]AccountParams{"data": accountParams}).
		Post("/v1/organisation/accounts")

	if err != nil {
		return nil, fmt.Errorf("create account failed: %s", err)
	}

	if resp.Error() != nil {
		return nil, getAPIError(resp)
	}

	return resp.Result().(*Resource), nil
}

// List return a list of accounts.
func (a *AccountClient) List(paging PagingParams) (*Resources, error) {

	r := a.client.R().SetResult(&Resources{})

	if paging.number != "" {
		r.SetQueryParam("page[number]", paging.number)
	}

	if paging.size != "" {
		r.SetQueryParam("page[size]", paging.size)
	}
	resp, err := r.Get("/v1/organisation/accounts")

	if err != nil {
		return nil, fmt.Errorf("list accounts failed: %s", err)
	}

	if resp.Error() != nil {
		return nil, getAPIError(resp)
	}

	return resp.Result().(*Resources), nil
}

// Fetch gets a single account by ID
func (a *AccountClient) Fetch(id string) (*Resource, error) {

	// Validate the account ID
	_, err := uuid.FromString(id)
	if err != nil {
		return nil, fmt.Errorf("account ID must be a valid UUID")
	}

	resp, err := a.client.R().
		SetResult(&Resource{}).
		SetPathParams(map[string]string{"account.id": id}).
		Get("/v1/organisation/accounts/{account.id}")

	if err != nil {
		return nil, fmt.Errorf("fetch account for ID %s failed: %s", id, err)
	}

	if resp.Error() != nil {
		return nil, getAPIError(resp)
	}
	return resp.Result().(*Resource), nil
}

// Delete deletes an account by ID
func (a *AccountClient) Delete(id string, version int) error {

	// Validate the account ID
	_, err := uuid.FromString(id)
	if err != nil {
		return fmt.Errorf("account ID must be a valid UUID")
	}

	resp, err := a.client.R().
		SetPathParams(map[string]string{"account.id": id}).
		SetQueryParam("version", strconv.Itoa(version)).
		Delete("/v1/organisation/accounts/{account.id}")

	if err != nil {
		return fmt.Errorf("delete account for ID %s failed: %s", id, err)
	}

	if resp.Error() != nil {
		return getAPIError(resp)
	}
	return nil
}

// Convert error response into error message
func getAPIError(resp *resty.Response) error {
	apiError := resp.Error().(*Error)
	return fmt.Errorf("request failed [%s]: %s", apiError.Code, apiError.Message)
}
