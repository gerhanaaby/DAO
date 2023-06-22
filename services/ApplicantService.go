package services

import (
	"DAO/database"
	request "DAO/models/request"
	table "DAO/models/tables"
	"fmt"
	"strconv"
	"time"
)

func CreateApplicant(applicant request.SubmitApplicants, KTPImage string, SelfieImage string) error {
	var dataApplicant table.Applicants
	var dataApplicantAlamat1 table.ApplicantAlamat
	var dataApplicantAlamat2 table.ApplicantAlamat
	var errorNPWP error
	var errorIncome error
	dataApplicant.Education = applicant.RawData.DataSubmit.EducationLevel.Name
	dataApplicant.CIFNumber = applicant.CifCode
	dataApplicant.Email = applicant.Email
	dataApplicant.Name = applicant.RawData.Name
	dataApplicant.HaveNPWP, errorNPWP = strconv.ParseBool(applicant.HaveNPWP)
	if errorNPWP != nil {
		fmt.Println(errorNPWP)
		return errorNPWP
	}
	dataApplicant.Income, errorIncome = strconv.ParseUint(applicant.RawData.DataSubmit.Income, 10, 64)
	if errorIncome != nil {
		fmt.Println(errorIncome)
		return errorIncome
	}
	dataApplicant.Religion = applicant.RawData.DataSubmit.Religion
	dataApplicant.JenisKelamin = applicant.RawData.DataSubmit.JenisKelamin
	dataApplicant.KTPImage = KTPImage
	dataApplicant.SelfieImage = SelfieImage
	dataApplicant.MaritalStatus = applicant.RawData.DataSubmit.MaritalStatus
	dataApplicant.MotherName = applicant.RawData.DataSubmit.MotherName
	dataApplicant.MobileNumber = applicant.MobileNumber
	dataApplicant.TmpLahir = applicant.RawData.DataSubmit.TmpLahir
	dataApplicant.KTPNumber = applicant.RawData.DataSubmit.KtpNumber
	date, error := time.Parse("2006-01-02", applicant.RawData.DataSubmit.DOB)
	if error != nil {
		fmt.Println(error)
		return error
	}
	dataApplicant.DOB = date
	fmt.Printf(applicant.RawData.DataSubmit.Alamat)
	dataApplicantAlamat1.Alamat = applicant.RawData.DataSubmit.Alamat
	dataApplicantAlamat1.Kecamatan = applicant.RawData.DataSubmit.Kecamatan.Name
	dataApplicantAlamat1.KodePos = applicant.RawData.DataSubmit.KodePos
	dataApplicantAlamat1.Kelurahan = applicant.RawData.DataSubmit.Kelurahan.Name
	dataApplicantAlamat1.Provinsi = applicant.RawData.DataSubmit.Provinsi.Name
	dataApplicantAlamat1.Kota = applicant.RawData.DataSubmit.Kota.Name
	dataApplicantAlamat1.RT = applicant.RawData.DataSubmit.RT
	dataApplicantAlamat1.RW = applicant.RawData.DataSubmit.RW
	dataApplicantAlamat2.Alamat = applicant.RawData.DataSubmit.Alamat2
	dataApplicantAlamat2.Kecamatan = applicant.RawData.DataSubmit.Kecamatan2.Name
	dataApplicantAlamat2.KodePos = applicant.RawData.DataSubmit.KodePos2
	dataApplicantAlamat2.Kelurahan = applicant.RawData.DataSubmit.Kelurahan2.Name
	dataApplicantAlamat2.Provinsi = applicant.RawData.DataSubmit.Provinsi2.Name
	dataApplicantAlamat2.Kota = applicant.RawData.DataSubmit.Kota2.Name
	dataApplicantAlamat2.RT = applicant.RawData.DataSubmit.RT2
	dataApplicantAlamat2.RW = applicant.RawData.DataSubmit.RW2
	dataApplicant.ApplicantAlamat = append(dataApplicant.ApplicantAlamat, dataApplicantAlamat1, dataApplicantAlamat2)

	dataApplicant.ApplicantOffice.OfficeName = applicant.RawData.DataSubmit.Company_Name
	dataApplicant.ApplicantOffice.OfficeZipCode = applicant.RawData.DataSubmit.Company_Address_ZIP_Code
	dataApplicant.ApplicantOffice.OfficeAddress = applicant.RawData.DataSubmit.Company_Address
	dataApplicant.ApplicantOffice.OfficePhoneNumber = applicant.RawData.DataSubmit.Company_Phone_Number

	dataApplicant.ApplicantHistory.Notes = "-"
	dataApplicant.ApplicantHistory.Status = "Complete"
	dataApplicant.ApplicantHistory.Step = 1
	fmt.Println(dataApplicant)
	err := database.GetDB().Debug().Create(&dataApplicant).Error

	return err
}

func GetApplicant() {
	// var dataApplicant []table.Applicants
	//var applicantHistory table.ApplicantHistory
	database.GetDB().Debug().Joins("JOIN applicant_histories ON applicant_histories.applicant_id = applicants.id AND applicant_histories.status = ?", "Complete").Limit(200).Find(&table.AppQ)
}
