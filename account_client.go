package interview_accountapi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
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

	resp, err := r.SetBody(map[string]AccountParams{"data": accountParams}).
		Post("/v1/organisation/accounts")

	if err != nil {
		return nil, fmt.Errorf("create account failed: %s", err)
	}

	if resp.Error() != nil {
		return nil, getAPIError(resp)
	}

	return resp.Result().(*Resource), nil
}

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

func (a *AccountClient) Fetch(id string) (*Resource, error) {

	_, err := uuid.FromString(id)
	if err != nil {
		return nil, fmt.Errorf("account ID must be a valid v4 UUID")
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

func (a *AccountClient) Delete(id string, version int) error {

	_, err := uuid.FromString(id)
	if err != nil {
		return fmt.Errorf("account ID must be a valid v4 UUID")
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

func getAPIError(resp *resty.Response) error {
	apiError := resp.Error().(*Error)
	return fmt.Errorf("request failed [%s]: %s", apiError.Code, apiError.Message)
}
