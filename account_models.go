package accountapi

// An Error maps to Form3 API error responses
type Error struct {
	Code    string `json:"error_code,omitempty"`
	Message string `json:"error_message,omitempty"`
}

// A Resource includes account data and the HATEOAS links
type Resource struct {
	Data  Account           `json:"data,omitempty"`
	Links map[string]string `json:"links,omitempty"`
}

// A Resources includes an array of account data and the HATEOAS links
type Resources struct {
	Data  []Account         `json:"data,omitempty"`
	Links map[string]string `json:"links,omitempty"`
}

// An Account represents the account data
type Account struct {
	Attributes     Attributes `json:"attributes,omitempty"`
	CreatedOn      string     `json:"created_on,omitempty"`
	ID             string     `json:"id,omitempty"`
	ModifiedOn     string     `json:"modified_on,omitempty"`
	OrganizationID string     `json:"organisation_id,omitempty"`
	Type           string     `json:"type,omitempty"`
	Version        int        `json:"version,omitempty"`
}

// An AccountParams is the request body for creating an account
type AccountParams struct {
	Attributes     Attributes `json:"attributes,omitempty"`
	ID             string     `json:"id,omitempty"`
	OrganizationID string     `json:"organisation_id,omitempty"`
	Type           string     `json:"type,omitempty"`
}

// A PagingParams is the query parameters for pagination
type PagingParams struct {
	number string
	size   string
}

// An Attributes represents the account attributes
type Attributes struct {
	AccountClassification       string   `json:"account_classification,omitempty"`
	AccountMatchingOptOut       bool     `json:"account_matching_opt_out,omitempty"`
	AccountNumber               string   `json:"account_number,omitempty"`
	AlternativeBankAccountNames []string `json:"alternative_bank_account_names,omitempty"`
	BankID                      string   `json:"bank_id,omitempty"`
	BankIDCode                  string   `json:"bank_id_code,omitempty"`
	BankAccountName             string   `json:"bank_account_name,omitempty"`
	BaseCurrency                string   `json:"base_currency,omitempty"`
	BIC                         string   `json:"bic,omitempty"`
	Country                     string   `json:"country,omitempty"`
	CustomerID                  string   `json:"customer_id,omitempty"`
	FirstName                   string   `json:"first_name,omitempty"`
	IBAN                        string   `json:"iban,omitempty"`
	JointAccount                bool     `json:"joint_account,omitempty"`
	Title                       string   `json:"title,omitempty"`
	SecondaryIdentification     string   `json:"secondary_identification,omitempty"`
}
