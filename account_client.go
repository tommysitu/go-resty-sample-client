package interview_accountapi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"strconv"
)

const DefaultBaseURL = "http://localhost:8080"

type AccountClient struct {
	client *resty.Client
}

func NewAccountClient() *AccountClient {
	client := resty.New()
	client.SetDebug(false)
	apiURL := os.Getenv("API_ADDR")
	if apiURL == "" {
		apiURL = DefaultBaseURL
	}
	client.SetHostURL(apiURL)
	client.SetError(&Error{})

	return &AccountClient{client: client}
}

func (a *AccountClient) Create(accountParams AccountParams) (*Resource, error) {

	r := a.client.R().SetResult(&Resource{})

	resp, e := r.SetBody(map[string]AccountParams{"data": accountParams}).
		Post("/v1/organisation/accounts")

	if e != nil {
		return nil, fmt.Errorf("create account failed: %s", e)
	}

	return resp.Result().(*Resource), nil
}

func (a *AccountClient) List(pageNumber string, pageSize string) (*Resources, error) {

	r := a.client.R().SetResult(&Resources{})

	if pageNumber != "" {
		r.SetQueryParam("page[number]", pageNumber)
	}

	if pageSize != "" {
		r.SetQueryParam("page[size]", pageSize)
	}
	resp, e := r.Get("/v1/organisation/accounts")

	if e != nil {
		return nil, fmt.Errorf("list accounts failed: %s", e)
	}

	return resp.Result().(*Resources), nil
}

func (a *AccountClient) Fetch(id string) (*Resource, error) {
	resp, e := a.client.R().
		SetResult(&Resource{}).
		SetPathParams(map[string]string{"account.id": id}).
		Get("/v1/organisation/accounts/{account.id}")

	if e != nil {
		return nil, fmt.Errorf("fetch account for ID %s failed: %s", id, e)
	}

	if resp.Error() != nil {
		return nil, getAPIError(resp)
	}
	return resp.Result().(*Resource), nil
}


func (a *AccountClient) Delete(id string, version int) error {
	resp, e := a.client.R().
		SetPathParams(map[string]string{"account.id": id}).
		SetQueryParam("version", strconv.Itoa(version)).
		Delete("/v1/organisation/accounts/{account.id}")

	if e != nil {
		return fmt.Errorf("delete account for ID %s failed: %s", id, e)
	}

	if resp.Error() != nil {
		return getAPIError(resp)
	}
	return nil
}

func getAPIError(resp *resty.Response) error {
	apiError := resp.Error().(*Error)
	return fmt.Errorf("request failed [%s]: %s", apiError.Code, apiError.Message)
}