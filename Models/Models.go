package models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: Applicants{}},
		{Model: ApplicantAlamat{}},
		{Model: ApplicantOffice{}},
		{Model: ApplicantDataBank{}},
		{Model: ApplicantHistory{}},
	}
}
