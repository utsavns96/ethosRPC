package main

import (
	"ethos/altEthos"
	"ethos/myRpc"
	"ethos/syscall"
	"log"
)


func init() {
	//myRpc.SetupMyRpcIncrement(increment)
	myRpc.SetupMyRpcPrintExpenseReport(printExpenseReport)
	myRpc.SetupMyRpcRemoveExpenseReport(removeExpenseReport)
	myRpc.SetupMyRpcCreateExpenseReport(CreateExpenseReport)
	myRpc.SetupMyRpcSubmitExpenseReport(SubmitExpenseReport)
	myRpc.SetupMyRpcAddItemExpenseReport(AddItemExpenseReport)
	myRpc.SetupMyRpcRemoveItemExpenseReport(RemoveItemExpensetReport)
	initializeAccounts("me")
}

func initializeAccounts(user string){
	log.Println("Initializing account for ", user)
	altEthos.SetRootFs("")
	path := "/user/" + user + "/expenses" 
	//path := "/user/me/expenses" 
	data1 := myRpc.MyType {"hello","world","foobar",1000000}
	status := altEthos.DirectoryCreate(path, &data1, "all")
	if status != syscall.StatusOk {
		log.Fatalf ("Error could not create %v  %v\n", path, status)
	}
}


func CreateExpenseReport() (myRpc.MyRpcProcedure) {
	log.Println("Creating expense report")
	return &myRpc.MyRpcCreateExpenseReportReply{syscall.StatusOk}
}

func removeExpenseReport() (myRpc.MyRpcProcedure) {
	log.Println("Removing expense report")
	return &myRpc.MyRpcRemoveItemExpenseReportReply{syscall.StatusOk}
}

func printExpenseReport() (myRpc.MyRpcProcedure) {
	log.Println("Printing expense report for user ", altEthos.GetUser())
	report := []string{"Alice", "Bob", "Cathy"}
	return &myRpc.MyRpcPrintExpenseReportReply{report, syscall.StatusOk}
}

func SubmitExpenseReport() (myRpc.MyRpcProcedure) {
	log.Println("Submitting expense report for user ", altEthos.GetUser())
	return &myRpc.MyRpcSubmitExpenseReportReply{syscall.StatusOk}
}

func AddItemExpenseReport(name string, date string, description string, amount int64) (myRpc.MyRpcProcedure) {
	log.Println("Adding item to expense report")
	log.Println("name: ", name, " | date: ", date, " | description: ", description, " | amount: $", amount)
	return &myRpc.MyRpcAddItemExpenseReportReply{23, syscall.StatusOk}
}

func RemoveItemExpensetReport(itemNumber int32) (myRpc.MyRpcProcedure) {
	log.Println("item number %v removed ", itemNumber, " for user " ,altEthos.GetUser())
	return &myRpc.MyRpcRemoveItemExpenseReportReply{syscall.StatusOk}
}

// func increment() (myRpc.MyRpcProcedure) {
// 	log.Println("called increment")
// 	myRpc_increment_counter++
// 	return &myRpc.MyRpcIncrementReply{myRpc_increment_counter}
// }

func main() {

	altEthos.LogToDirectory("application/expenseService")

	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		log.Println("Advertising service failed: ", status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error calling Import: %v\n", status)
			altEthos.Exit(status)
		}

		log.Println("new connection accepted")

		t := myRpc.MyRpc{}
		altEthos.Handle(fd, &t)
	}
}
