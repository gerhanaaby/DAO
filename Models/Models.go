package Models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: Customer{}},
		{Model: AlamatCustomer{}},
		{Model: OfficeCustomer{}},
		{Model: DataApplicant{}},
	}
}
