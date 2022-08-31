package models

import table "DAO/models/tables"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: table.Applicants{}},
		{Model: table.ApplicantHistory{}},
		{Model: table.ApplicantDataBank{}},
		{Model: table.ApplicantOffice{}},
		{Model: table.ApplicantAlamat{}},
	}
}
