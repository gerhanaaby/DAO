// package controllers

// import (
// 	models "DAO/models/request"
// 	tables "DAO/models/tables"
// 	"DAO/services"
// 	"DAO/utils"
// 	"encoding/json"
// )

// func PostStartQueue() {
// 	services.GetApplicant()
// 	//fmt.Println(len(tables.AppQ))

// 	for i := 0; i < len(tables.AppQ); i++ {
// 		// fmt.Printf(`Masuk antrian ke %d`, i)
// 		// fmt.Println()
// 		go DAO(&tables.AppQ[i], i)

// 	}
// }

// // func DAO(applicant *tables.Applicants, i int) {
// 	//SynC
// 	SyncCIFRequest := models.Sync{}
// 	SyncCIFRequest.SyncData.BranchCode     	= "0121"
// 	SyncCIFRequest.SyncData.ShortName       = applicant.Name
// 	SyncCIFRequest.SyncData.MotherName      = applicant.MotherName
// 	SyncCIFRequest.SyncData.BirthIncorpDate = applicant.DOB
// 	SyncCIFRequest.SyncData.TelMobile      	= applicant.MobileNumber
// 	SyncCIFRequest.SyncData.IdTypes         = "1"
// 	SyncCIFRequest.SyncData.LegalIdNo      	= applicant.KTPNumber
// 	SyncCIFRequest.SyncData.SmGender        = applicant.JenisKelamin
// 	SyncCIFRequest.SyncData.Street          = applicant.ApplicantAlamat[0].Alamat
// 	SyncCIFRequest.SyncData.SuburbTown      = applicant.ApplicantAlamat[0].Kelurahan
// 	SyncCIFRequest.SyncData.CityMunicipal   = applicant.ApplicantAlamat[0].Kota
// 	SyncCIFRequest.SyncData.MaritalStatus   = applicant.MaritalStatus
// 	SyncCIFRequest.SyncData.SourceofFound   = applicant.Pekerjaan
// 	SyncCIFRequest.SyncData.FundPurpose     = applicant.Income
// 	SyncCIFRequest.SyncData.PlaceBirth      = applicant.TmpLahir
// 	SyncCIFRequest.SyncData.LocationCode    = ""
// 	SyncCIFRequest.SyncData.PostalCode      = applicant.ApplicantAlamat[0].KodePos
// 	SyncCIFRequest.SyncData.JobCode         = ""
// 	SyncCIFRequest.SyncData.UmgUniqueId     = ""
// 	SyncCIFRequest.SyncData.ChannelRefNum   = ""
// 	SyncCIFRequest.SyncData.EmailAddress    = applicant.Email
// 	// SyncCIFRequest.SyncData.Nationality     =
// 	// SyncCIFRequest.SyncData.OfficeAddress   = applicant.ApplicantOffice.OfficeAddress
// 	// SyncCIFRequest.SyncData.SalaryRange     =
// 	// SyncCIFRequest.SyncData.LubCusTypeBS2  	=
// 	// SyncCIFRequest.SyncData.IdExpiryDate    =
// 	// SyncCIFRequest.SyncData.OperatorId      =
// 	// SyncCIFRequest.SyncData.BankRel         =
// 	// SyncCIFRequest.SyncData.ResideYN        =
// 	// SyncCIFRequest.SyncData.TerminalId      =
// 	JsonSyncCIFByte, err := json.Marshal(SyncCIFRequest)
// 	if err != nil {
// 		err = "Error Data Request"
// 		return1
// 	}
// 	services.ConsumeAPIService(utils.GetEndPoint("synccif", "synccif-key", "synccif-auth"), JsonSyncCIFByte)

// 	//CSR
// 	CSRRequest := models.CSR{}
// 	// CSRRequest.Service      =
// 	// CSRRequest.TemplateName =
// 	// CSRRequest.XmlTag       =
// 	// CSRRequest.WhereKondisi =

// 	// CSRRequest.TableEls.FieldParam =
// 	// CSRRequest.TableEls.FieldName  	=
// 	// CSRRequest.TableEls.FieldValue =
// 	// CSRRequest.TableEls.FlagKutip  	=
// 	// CSRRequest.TableEls.TableName  =
// 	JsonCSRByte, err := json.Marshal(CSRRequest)
// 	if err != nil {
// 		err = "Error Data Request"
// 		return
// 	}
// 	services.ConsumeAPIService(utils.GetEndPoint("csr20", "csr20-key", "csr20-auth"), JsonCSRByte)

// 	//Screening
// 	screeningRequest := models.Screening{}
// 	screeningRequest.CIFNumber    		= applicant.CIFNumber
// 	screeningRequest.Name         		= applicant.Name
// 	// screeningRequest.Nationality 		=
// 	// screeningRequest.MatchScore    		=
// 	// screeningRequest.DOB            	=applicant.DOB
// 	screeningRequest.IdentityNumber  	= applicant.KTPNumber
// 	JsonOneFCCByte, err := json.Marshal(screeningRequest)
// 	if err != nil {
// 		err = "Error Data Request"
// 		return
// 	}
// 	services.ConsumeAPIService(utils.GetEndPoint("onefcc", "onefcc-key", "onefcc-auth"), JsonOneFCCByte)

// 	//Application Complete
// 	AppCompleteRequest := models.ApplicationComplete{}
// 	AppCompleteRequest.CIF        = applicant.CIFNumber
// 	// AppCompleteRequest.RequestCif =
// 	// AppCompleteRequest.OrderId =
// 	// AppCompleteRequest.Data.AccountNumber =
// 	// AppCompleteRequest.Data.AccountName   =
// 	// AppCompleteRequest.Data.AccountType   =
// 	// AppCompleteRequest.Data.Status        =
// 	// AppCompleteRequest.Data.StatusDesc    =
// 	// AppCompleteRequest.Data.SubStatus     =
// 	// AppCompleteRequest.Data.SubStatusDesc =
// 	// AppCompleteRequest.Data.NextStep      =
// 	// AppCompleteRequest.Data.Cif    =
// 	// AppCompleteRequest.Data.OrderId  =
// 	JsonApplicationCompleteByte, err := json.Marshal(AppCompleteRequest)
// 	if err != nil {
// 		err = "Error Data Request"
// 		return
// 	}
// 	services.ConsumeAPIService(utils.GetEndPoint("eformcentral", "eformcentral-key", "eformcentral-auth"), JsonApplicationCompleteByte)

// }
