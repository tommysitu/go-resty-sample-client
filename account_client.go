package interview_accountapi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

const DefaultBaseURL = "http://localhost:8080"

type AccountClient struct {
	client *resty.Client
}

func NewAccountClient() *AccountClient {
	client := resty.New()
	client.SetHostURL(DefaultBaseURL)
	client.SetError(&Error{})

	return &AccountClient{client: client}
}

func (a *AccountClient) Create(accountParams *AccountParams) (*Resource, error) {

	r := a.client.R().SetResult(&Resource{})

	resp, e := r.SetBody(accountParams).
		Post("/v1/organisation/accounts")

	if e != nil {
		return nil, fmt.Errorf("create account failed: %s", e)
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
		apiError := resp.Error().(*Error)
		return fmt.Errorf("request failed [%s]: %s", apiError.Code, apiError.Message)
	}
	return nil
}