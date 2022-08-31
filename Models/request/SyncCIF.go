package models

// type Sync struct {
// 	gorm.Model
// 	SyncData SyncData
// }

// type SyncData struct {
// 	gorm.Model
// 	ID              uint   `validate:"isdefault"`
// 	BranchCode      string `gorm:"size:256;not null;" validate:"required" `
// 	ShortName       string `gorm:"size:256;not null;" validate:"required" `
// 	MotherName      string `gorm:"size:256;not null;" validate:"required" `
// 	BirthIncorpDate string `gorm:"size:256;not null;" validate:"required" `
// 	TelMobile       string `gorm:"size:256;not null;" validate:"required" `
// 	IdTypes         string `gorm:"size:256;not null;" validate:"required" `
// 	LegalIdNo       string `gorm:"size:256;not null;" validate:"required" `
// 	SmGender        string `gorm:"size:256;not null;" validate:"required" `
// 	Street          string `gorm:"size:256;not null;" validate:"required" `
// 	SuburbTown      string `gorm:"size:256;not null;" validate:"required" `
// 	CityMunicipal   string `gorm:"size:256;not null;" validate:"required" `
// 	MaritalStatus   string `gorm:"size:256;not null;" validate:"required" `
// 	SourceOfFound   string `gorm:"size:256;not null;" validate:"required" `
// 	FundPurpose     string `gorm:"size:256;not null;" validate:"required" `
// 	PlaceBirth      string `gorm:"size:256;not null;" validate:"required" `
// 	LocationCode    string `gorm:"size:256;not null;" validate:"required" `
// 	PostalCode      string `gorm:"size:256;not null;" validate:"required" `
// 	JobCode         string `gorm:"size:256;not null;" validate:"required" `
// 	UmgUniqueId     string `gorm:"size:256;not null;" validate:"required" `
// 	ChannelRefNum   string `gorm:"size:256;not null;" validate:"required" `
// 	EmailAddress    string `gorm:"size:256;not null;" validate:"required" `
// 	Nationality     string `gorm:"size:256;not null;" validate:"required" `
// 	OfficeAddress   string `gorm:"size:256;not null;" validate:"required" `
// 	SalaryRange     string `gorm:"size:256;not null;" validate:"required" `
// 	LubCusTypeBS2   string `gorm:"size:256;not null;" validate:"required" `
// 	IdExpiryDate    string `gorm:"size:256;not null;" validate:"required" `
// 	OperatorId      string `gorm:"size:256;not null;" validate:"required" `
// 	BankRel         string `gorm:"size:256;not null;" validate:"required" `
// 	ResideYN        string `gorm:"size:256;not null;" validate:"required" `
// 	TerminalId      string `gorm:"size:256;not null;" validate:"required" `
// }

type Sync struct {
	SyncData SyncData `json:"syncData"`
}

type SyncData struct {
	BranchCode      string `json:"branchCode"`
	ShortName       string `json:"shortName"`
	MotherName      string `json:"motherName"`
	BirthIncorpDate string `json:"birthIncorpDate"` //DOB
	TelMobile       string `json:"telMobile"`
	IdTypes         string `json:"idTypes"`
	LegalIdNo       string `json:"legalIdNo"` //ktp //gender
	SmGender        string `json:"smGender"`  //gender
	Street          string `json:"street"`
	SuburbTown      string `json:"suburbTown"`
	CityMunicipal   string `json:"cityMunicipal"`
	MaritalStatus   string `json:"maritalStatus"`
	SourceofFound   string `json:"sourceOfFund"`
	FundPurpose     string `json:"fundPurpose"`
	PlaceBirth      string `json:"placeBirth"`
	LocationCode    string `json:"locationCode"`
	PostalCode      string `json:"postalCode"`
	JobCode         string `json:"jobCode"`       //jobcode
	UmgUniqueId     string `json:"umgUniqueId"`   //tanya
	ChannelRefNum   string `json:"channelRefNum"` //tanya
	EmailAddress    string `json:"emailAddress"`
	Nationality     string `json:"nationality"` //ID
	OfficeAddress   string `json:"officeAddress"`
	SalaryRange     string `json:"salaryRange"`
	LubCusTypeBS2   string `json:"lubCusTypeBS2"` //hard code
	IdExpiryDate    string `json:"idExpiryDate"`  //hard code
	OperatorId      string `json:"operatorId"`    //hard code
	BankRel         string `json:"bankRel"`       //hard code
	ResideYN        string `json:"resideYN"`      //hard code
	TerminalId      string `json:"terminalId"`    //hard code
}

//citymunicipal = kecamatan
//sourceoffound = source
//
