package interview_accountapi

type Error struct {
	Code    string `json:"error_code"`
	Message string `json:"error_message"`
}

type Resource struct {
	Data  Account           `json:"data"`
	Links map[string]string `json:"links"`
}

type Resources struct {
	Data  []Account         `json:"data"`
	Links map[string]string `json:"links"`
}

type Account struct {
	Attributes     Attributes `json:"attributes"`
	CreatedOn      string     `json:"created_on"`
	ID             string     `json:"id"`
	ModifiedOn     string     `json:"modified_on"`
	OrganizationID string     `json:"organisation_id"`
	Type           string     `json:"type"`
	Version        int        `json:"version"`
}

type AccountParams struct {
	Attributes     Attributes `json:"attributes"`
	ID             string     `json:"id"`
	OrganizationID string     `json:"organisation_id"`
	Type           string     `json:"type"`
}

type Attributes struct {
	AccountClassification       string   `json:"account_classification"`
	AccountMatchingOptOut       bool     `json:"account_matching_opt_out"`
	AccountNumber               string   `json:"account_number"`
	AlternativeBankAccountNames []string `json:"alternative_bank_account_names"`
	BankID                      string   `json:"bank_id"`
	BankIDCode                  string   `json:"bank_id_code"`
	BaseCurrency                string   `json:"base_currency"`
	BIC                         string   `json:"bic"`
	Country                     string   `json:"country"`
	CustomerID                  string   `json:"customer_id"`
	FirstName                   string   `json:"first_name"`
	IBAN                        string   `json:"iban"`
	JointAccount                bool     `json:"joint_account"`
	Title                       string   `json:"title"`
}
