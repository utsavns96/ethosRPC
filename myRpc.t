MyRpc interface {
	CreateExpenseReport() (status Status)
	RemoveExpenseReport() (status Status)
	PrintExpenseReport() (list []string, status Status)
	SubmitExpenseReport() (status Status)
	AddItemExpenseReport(name string, date string, description string, amount int64) (itemNumber int32, status Status)
	RemoveItemExpenseReport(itemNumber int32) (status Status)
}

MyType struct {
	Field1 string
	Field2 string
	Field3 string
	Field4 int64
}
