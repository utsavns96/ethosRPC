package main

import (
	"ethos/altEthos"
	"ethos/myRpc"
	"ethos/syscall"
	"ethos/fmt"
	"log"
)


func init() {
	//myRpc.SetupMyRpcIncrement(increment)
	myRpc.SetupMyRpcPrintExpenseReport(printExpenseReport)
	myRpc.SetupMyRpcRemoveExpenseReport(removeExpenseReport)
	myRpc.SetupMyRpcCreateExpenseReport(CreateExpenseReport)
	myRpc.SetupMyRpcSubmitExpenseReport(SubmitExpenseReport)
	myRpc.SetupMyRpcAddItemExpenseReport(AddItemExpenseReport)
	myRpc.SetupMyRpcRemoveItemExpenseReport(RemoveItemExpenseReport)
	//initializeAccounts("me")
}

func initializeAccounts(user string){
	log.Println("Initializing account for ", user)
	altEthos.SetRootFs("")
	path := "/user/" + user + "/expenses" 
	//path := "/user/me/expenses" 
	data1 := myRpc.MyType{"hello","world","foobar",1000000}
	status := altEthos.DirectoryCreate(path, &data1, "all")
	if status != syscall.StatusOk {
		log.Printf ("Error could not create main dir %v  %v\n", path, status)
	}
	path2 := "/user/" + user + "/expenses/submitted"
	status2 := altEthos.DirectoryCreate(path2, &data1, "all")
	if status2 != syscall.StatusOk {
		log.Printf ("Error could not create submitted dir %v  %v\n", path2, status2)
	}
	path3 := "/user/" + user + "/expenses/open"
	status3 := altEthos.DirectoryCreate(path3, &data1, "all")
	if status3 != syscall.StatusOk {
		log.Printf ("Error could not create open dir %v  %v\n", path3, status3)
	}
	log.Println("folders initialized for " + user)
}


func CreateExpenseReport(user string) (myRpc.MyRpcProcedure) {
	log.Println("Creating expense report")
	initializeAccounts(user)
	//This function runs after creating dirs
	fp := "/user/" + user + "/expenses/open/" + fmt.Sprintf("%v", altEthos.GetTime())
	//var data1 []myRpc.MyType
	data1 := myRpc.MyType{"name,","date,","description,",-999999}
	//data1 = append(data1, &element)
	//data1 := []string("name,date,description,-9999999")
	status := altEthos.Write(fp, &data1)
	if status != syscall.StatusOk {
		log.Fatalf ("Error creating expense report %v\n", status)
	}
	return &myRpc.MyRpcCreateExpenseReportReply{syscall.StatusOk}
}

func removeExpenseReport() (myRpc.MyRpcProcedure) {
	log.Println("Removing expense report")
	//fp := "/user/" + user + "/expenses/open"
	pathopen := "/user/" + "me" + "/expenses/open"
	files, status := altEthos.SubFiles(pathopen)
	if status != syscall.StatusOk {
		log.Fatalf ("Error finding files %v\n", status)
	}
	if len(files) > 0{
		for _,f:= range files {
			p:=pathopen+"/"+f
			status = altEthos.FileRemove(p)
			if status != syscall.StatusOk {
				log.Fatalf ("Error deleting submitted file %v\n", status)
			}
		}
	} else {
		log.Fatalf ("No files to delete")
	}
	return &myRpc.MyRpcRemoveItemExpenseReportReply{syscall.StatusOk}
}

func printExpenseReport() (myRpc.MyRpcProcedure) {
	log.Println("Printing expense report for user ", altEthos.GetUser())
	path := "/user/" + "me" + "/expenses/open"
	var report1 []string
	//report := []string{"Alice", "Bob", "Cathy"}
	files, status := altEthos.SubFiles(path)
	if status != syscall.StatusOk {
		log.Fatalf ("Error finding files %v\n", status)
	}
	//f := files[len(files)-1]
	//p := path+"/"+f
	//var v myRpc.MyType
	log.Println("Found open expense report "+path)
	for _,f:= range files {
		var v myRpc.MyType
		p:=path+"/"+f
		status = altEthos.Read(p, &v)
		if status!=syscall.StatusOk {
			log.Printf("Could not read %v\n", p)
		}
		log.Println("read ", p, v)
		report1 = append(report1 , fmt.Sprintf("%v", v))
	}
	//status = altEthos.Read(p, &v)
	//if status!=syscall.StatusOk {
	//		log.Printf("Could not read %v\n", p)
	//	}
	//log.Println("read data: ", p, v)
	
	
	return &myRpc.MyRpcPrintExpenseReportReply{report1, syscall.StatusOk}
}

func SubmitExpenseReport() (myRpc.MyRpcProcedure) {
	//Copies files from open to submitted and deletes them from open
	log.Println("Submitting expense report for user ", altEthos.GetUser())
	pathopen := "/user/" + "me" + "/expenses/open"
	files, status := altEthos.SubFiles(pathopen)
	// submit only if there are files in /open
	if len(files) > 0{
		//create new submission dir
		data1 := myRpc.MyType{"hello","world","foobar",1000000}
		pathsubmitted := "/user/" + "me" + "/expenses/submitted/"+ fmt.Sprintf("%v", altEthos.GetTime()) + "_submission"
		status2 := altEthos.DirectoryCreate(pathsubmitted, &data1, "all")
		if status2 != syscall.StatusOk {
			log.Printf ("Error could not create expense submission %v  %v\n", pathsubmitted, status2)
		}
		//for each file in open
		for _,f:= range files {
			var v myRpc.MyType
			p:=pathopen+"/"+f
			//read file
			status = altEthos.Read(p, &v)
			if status!=syscall.StatusOk {
				log.Printf("Could not read %v\n", p)
			}
			log.Println("read ", p, v)
			//create new file in submitted/<time>_submission dir
			fp := pathsubmitted + "/" + fmt.Sprintf("%v", altEthos.GetTime())
			status3 := altEthos.Write(fp, &v)
			if status3 != syscall.StatusOk {
				log.Fatalf ("Error writing submission file %v\n", status3)
			}
			//delete file from open
			status4 := altEthos.FileRemove(p)
			if status4 != syscall.StatusOk {
				log.Fatalf ("Error deleting submitted file %v\n", status4)
			}
		}
	} else {
		log.Fatalf ("No expenses to submit")
	}
	return &myRpc.MyRpcSubmitExpenseReportReply{syscall.StatusOk}
}

func AddItemExpenseReport(name string, date string, description string, amount int64, user string) (myRpc.MyRpcProcedure) {
	log.Println("Adding item to expense report")
	log.Println("name: ", name, " | date: ", date, " | description: ", description, " | amount: $", amount)
	data := myRpc.MyType{name+",", date+",", description+",",amount}
	path := "/user/" + "me" + "/expenses/open"
	files, status := altEthos.SubFiles(path)
	if status != syscall.StatusOk {
		log.Fatalf ("Error finding files %v\n", status)
	}
	pos := int32(len(files))
	fp := path+"/"+ fmt.Sprintf("%v", altEthos.GetTime())
	status = altEthos.Write(fp, &data)
	if status != syscall.StatusOk {
		log.Fatalf ("Error creating expense report %v\n", status)
	}
	//if pos>0{
	//	f := files[pos-1]
	//	p := path+"/"+f
	//	var v []myRpc.MyType
	//	log.Println("Found open expense report "+p)
	//	status = altEthos.Read(p, &v)
	//	if status!=syscall.StatusOk {
	//			log.Printf("Could not read %v\n", p)
	//		}
	//	log.Println("read data: ", p, v)
	//	v = append(v,data)
	//	status = altEthos.Write(f, &v)
	//} else {
	//	f := path + "/" + fmt.Sprintf("%v", altEthos.GetTime())
	//	//data1 := myRpc.MyType{name+",", date+",", description+",", amount}
	//	status = altEthos.Write(f, &data)
	//}
	//
	//if status != syscall.StatusOk {
	//	log.Fatalf ("Error adding item to expense report %v\n", status)
	//}
	
	//fp := "/user/" + "me" + "/expenses/open/" + fmt.Sprintf("%v", altEthos.GetTime())
	//data1 := myRpc.MyType{name+",", date+",", description+",", amount}

	//status := altEthos.Write(fp, &data1)
	//if status != syscall.StatusOk {
	//	log.Fatalf ("Error creating expense report %v\n", status)
	//}
	log.Println("Item added to expense report \n", status)
	return &myRpc.MyRpcAddItemExpenseReportReply{pos, syscall.StatusOk}
}

func RemoveItemExpenseReport(itemNumber int32) (myRpc.MyRpcProcedure) {
	log.Println("item number %v removed ", itemNumber, " for user " ,altEthos.GetUser())
	return &myRpc.MyRpcRemoveItemExpenseReportReply{syscall.StatusOk}
}


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
