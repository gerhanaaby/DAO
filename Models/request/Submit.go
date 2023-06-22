package Models

type SubmitApplicants struct {
	Cif_Code               string  `json:"cif_code"`
	CifCode                string  `json:"cifCode"`
	Code                   string  `json:"code"`
	CreatedDate            string  `json:"createdDate"`
	CustomerConsent        string  `json:"customerConsent"`
	Email                  string  `json:"email"`
	Format                 string  `json:"format"`
	HaveNPWP               string  `json:"haveNPWP"`
	ID                     string  `json:"id"`
	KTPID                  string  `json:"ktpId"`
	LastUpdateDate         string  `json:"lastUpdatedDate"`
	MobileNumber           string  `json:"mobileNumber"`
	Mode                   string  `json:"mode"`
	OnSubmit               string  `json:"onSubmit"`
	PageName               string  `json:"pageName"`
	StatusAddress          string  `json:"statusAddress"`
	Version                string  `json:"version"`
	VersionCustomerConsent string  `json:"versiibCustomerConsent"`
	RawData                RawData `json:"rawData"`
}

type RawData struct {
	CifCode        string         `json:"cifCode"`
	FormID         string         `json:"formId"`
	ID             string         `json:"id"`
	KTPID          string         `json:"ktpId"`
	MobileNumber   string         `json:"mobileNumber"`
	Name           string         `json:"name"`
	DataAdditional DataAdditional `json:"dataAdditional"`
	DataSubmit     DataSubmit     `json:"dataSubmit"`
}

type DataAdditional struct {
	CustomerConsent        string `json:"customerConsent"`
	ProductCode            string `json:"productCode"`
	VersionCustomerConsent string `json:"versionCustomerConsent"`
}

type DataSubmit struct {
	Company_Name             string        `json:"Company_Name"`
	Company_Address_ZIP_Code string        `json:"Company_Address_ZIP_Code"`
	Company_Address          string        `json:"Company_Address"`
	Company_Phone_Number     string        `json:"Company_Phone_Number"`
	Alamat                   string        `json:"alamat"`
	Alamat2                  string        `json:"alamat2"`
	DOB                      string        `json:"dob"`
	Income                   string        `json:"income"`
	JenisKelamin             string        `json:"jenis_kelamin"`
	KodePos                  string        `json:"kode_pos"`
	KodePos2                 string        `json:"kode_pos2"`
	KtpImage                 string        `json:"ktp_image"`
	KtpNumber                string        `json:"KTP_NUMBER"`
	MaritalStatus            string        `json:"marital_status"`
	MonthlyCredAmt           string        `json:"monthlyCredAmt"`
	MonthlyCredFreq          string        `json:"monthlyCredFreq"`
	MonthlyDebtAmt           string        `json:"monthlyDebtAmt"`
	MonthlyDebtFreq          string        `json:"monthlyDebtFreq"`
	MotherName               string        `json:"mother_name"`
	PlaceBirth               string        `json:"place_birth"`
	Religion                 string        `json:"Religion"`
	RT                       string        `json:"rt"`
	RT2                      string        `json:"rt2"`
	RW                       string        `json:"rw"`
	RW2                      string        `json:"rw2"`
	SelfieImage              string        `json:"selfie_image"`
	SelfieScore              string        `json:"selfie_score"`
	Title                    string        `json:"Title"`
	TmpLahir                 string        `json:"tmp_lahir"`
	DataBank                 DataBank      `json:"dataBank"`
	EducationLevel           CodeName      `json:"Education_Level"`
	Kecamatan                CodeName      `json:"kecamatan"`
	Kecamatan2               CodeName      `json:"kecamatan2"`
	Kelurahan                CodeName      `json:"kelurahan"`
	Kelurahan2               CodeName      `json:"kelurahan2"`
	Kota                     CodeName      `json:"kota"`
	Kota2                    CodeName      `json:"kota2"`
	Pekerjaan                CodeName      `json:"pekerjaan"`
	Provinsi                 CodeName      `json:"provinci"`
	Provinsi2                CodeName      `json:"provinci2"`
	Purpose                  CodeName      `json:"purpose"`
	Source                   CodeName      `json:"source"`
	SourceAccount            SourceAccount `json:"sourceAccount"`
}

type SourceAccount struct {
}

type DataBank struct {
}

type CodeName struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
