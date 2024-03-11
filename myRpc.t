MyRpc interface {
	CreateExpenseReport(user string) (status Status)
	RemoveExpenseReport(user string) (status Status)
	PrintExpenseReport(user string) (list []string, status Status)
	SubmitExpenseReport(user string) (status Status)
	AddItemExpenseReport(name string, date string, description string, amount int64, user string) (itemNumber int32, status Status)
	RemoveItemExpenseReport(itemNumber int32, user string) (status Status)
}

MyType struct {
	Field1 string
	Field2 string
	Field3 string
	Field4 int64
}
