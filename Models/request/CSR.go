package Models

type CSR struct {
	Service      string     `json:"service"`
	TemplateName string     `json:"templateName"`
	XmlTag       string     `json:"xmlTag"`
	WhereKondisi string     `json:"whereKondisi"`
	TableEls     []TableEls `json:"tableEls"`
}

type TableEls struct {
	FieldParam string `json:"fieldParam"`
	FieldName  string `json:"fieldName"`
	FieldValue string `json:"fieldValue"`
	FlagKutip  string `json:"flagKutip"`
	TableName  string `json:"tableName"`
}
