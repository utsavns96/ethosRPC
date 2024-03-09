package main

import (
	"ethos/altEthos"
	"ethos/myRpc"
	"ethos/syscall"
	"log"
)

func init() {
	myRpc.SetupMyRpcPrintExpenseReportReply(printReply)
	myRpc.SetupMyRpcAddItemExpenseReportReply(addItemReply)
	myRpc.SetupMyRpcRemoveItemExpenseReportReply(removeExpenseReply)
	myRpc.SetupMyRpcCreateExpenseReportReply(createReportReply)
	myRpc.SetupMyRpcSubmitExpenseReportReply(submitExpenseReply)
	myRpc.SetupMyRpcRemoveItemExpenseReportReply(removeExpenseReportReply)
}

func createReportReply(status syscall.Status) (myRpc.MyRpcProcedure) {
	if status != syscall.StatusOk {
		log.Printf("Expense Report creation failed %v\n", status)
	} else {
		log.Println("Expense Report created")
	}
	return nil
}

func removeExpenseReportReply(status syscall.Status) (myRpc.MyRpcProcedure){
	if status != syscall.StatusOk {
		log.Printf("Failed to remove expense report %v\n", status)
	} else {
		log.Println("Expense Report removed")
	}
	return nil
}

func printReply(list []string, status syscall.Status) (myRpc.MyRpcProcedure) {
	if status != syscall.StatusOk {
		log.Printf("Failed to fetch expense report %v\n", status)
	} else {
		log.Println("Expense report fetched")
		for i := 0; i < len(list); i++ {
			log.Println(list[i])
		}
	}
	return nil
}

func addItemReply(itemNumber int32, status syscall.Status) (myRpc.MyRpcProcedure) {
	if status != syscall.StatusOk {
		log.Printf("Adding item failed %v\n", status)
	} else {
		log.Println("Item added")
	}
	return nil
}

func removeExpenseReply(status syscall.Status) (myRpc.MyRpcProcedure) {
	if status != syscall.StatusOk {
		log.Printf("Removing expense failed %v\n", status)
	} else {
		log.Println("Expense removed")
	}
	return nil
}

func submitExpenseReply(status syscall.Status) (myRpc.MyRpcProcedure) {
	if status != syscall.StatusOk {
		log.Printf("Failed to submit expense %v\n", status)
	} else {
		log.Println("Expense submitted")
	}
	return nil
}

func main() {

	altEthos.LogToDirectory("test/expenseClient")

	log.Println("before call")

	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}
	
	log.Println("Creating Report")
	call1 := myRpc.MyRpcCreateExpenseReport{}
	status1 := altEthos.ClientCall(fd, &call1)
	if status1 != syscall.StatusOk {
		log.Printf("Creating report failed:_%v\n", status1)
		altEthos.Exit(status1)
	}
	
	//log.Println("Adding Item")
	//call2 := myRpc.MyRpcAddItemExpenseReport{"abc", "12-01-2024", "test", 37}
	//status2 := altEthos.ClientCall(fd, &call2)
	//if status2 != syscall.StatusOk {
	//	log.Printf("Adding item failed:_%v\n", status2)
	//	altEthos.Exit(status2)
	//}
	
	//log.Println("Removing Item")
	//call3 := myRpc.MyRpcRemoveItemExpenseReport{35}
	//status3 := altEthos.ClientCall(fd, &call3)
	//if status3 != syscall.StatusOk {
	//	log.Printf("Removing item failed:_%v\n", status3)
	//	altEthos.Exit(status3)
	//}
	
	log.Println("Printing Report")
	call4 := myRpc.MyRpcPrintExpenseReport{}
	status4 := altEthos.ClientCall(fd, &call4)
	if status4 != syscall.StatusOk {
		log.Printf("Printing report failed:_%v\n", status4)
		altEthos.Exit(status4)
	}
	
	log.Println("Submitting Report")
	call5 := myRpc.MyRpcSubmitExpenseReport{}
	status5 := altEthos.ClientCall(fd, &call5)
	if status5 != syscall.StatusOk {
		log.Printf("Submitting report failed:_%v\n", status5)
		altEthos.Exit(status5)
	}
	
	log.Println("Deleting Report")
	call6 := myRpc.MyRpcRemoveExpenseReport{}
	status6 := altEthos.ClientCall(fd, &call6)
	if status6 != syscall.StatusOk {
		log.Printf("Deleting report failed:_%v\n", status6)
		altEthos.Exit(status6)
	}
	
	log.Println("expenseClient: done")
}
