package Models

type ApplicationComplete struct {
	CIF        string `json:"service"`
	RequestCif string `json:"templateName"`
	OrderId    string `json:"xmlTag"`
	Data       Data
}

type Data struct {
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	AccountType   string `json:"accountType"`
	Status        int    `json:"status"`
	StatusDesc    string `json:"statusDesc"`
	SubStatus     string `json:"subStatus"`
	SubStatusDesc string `json:"subStatusDesc"`
	NextStep      int    `json:"nextStep"`
	Cif           string `json:"cif"`
	OrderId       string `json:"orderId"`
}
