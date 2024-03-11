export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=ethos
export GOLINUXINCLUDE=linux
export BUILD=ethos

export ETHOSROOT=client/rootfs
export MINIMALTDROOT=client/minimaltdfs


.PHONY: all install clean
all: expenseService expenseClient expenseClient2

ethos:
	mkdir ethos
	cp -pr /usr/lib64/go/pkg/ethos_$(GOARCH)/* ethos

myRpc.go: myRpc.t
	$(ETN2GO) . myRpc $^

myRpc.goo.ethos : myRpc.go ethos
	ethosGoPackage  myRpc ethos myRpc.go

#myType.go: myType.t ethos
#	$(ETN2GO) . myType myType.t

#myType.goo.ethos: myType.go
#	ethosGoPackage myType ethos myType.go
#myType.goo.ethos
expenseService: expenseService.go myRpc.goo.ethos 
	ethosGo expenseService.go
#myType.goo.ethos
expenseClient: expenseClient.go myRpc.goo.ethos 
	ethosGo expenseClient.go
	
expenseClient2: expenseClient2.go myRpc.goo.ethos 
	ethosGo expenseClient2.go

# install types, service,
install: all
	sudo rm -rf client
	(ethosParams client && cd client && ethosMinimaltdBuilder)
	ethosTypeInstall myRpc
	ethosDirCreate $(ETHOSROOT)/services/myRpc   $(ETHOSROOT)/types/spec/myRpc/MyRpc all
	install -D  expenseClient expenseClient2 expenseService                   $(ETHOSROOT)/programs
	ethosStringEncode /programs/expenseService    > $(ETHOSROOT)/etc/init/services/expenseService
	ethosStringEncode /programs/expenseClient       > $(ETHOSROOT)/etc/init/services/expenseClient
	ethosStringEncode /programs/expenseClient2       > $(ETHOSROOT)/etc/init/services/expenseClient2

# remove build artifacts
clean:
	rm -rf myRpc/ myRpcIndex/ ethos clent
	rm -f myRpc.go
	rm -f expenseClient
	rm -f expenseService
	rm -f expenseClient2
	rm -f expenseService.goo.ethos 
	rm -f expenseClient.goo.ethos 
	rm -f expenseClient2.goo.ethos 
	rm -f myRpc.goo.ethos
	rm -rf myType/ myTypeIndex/ client ethos
	rm -f myType.go
	rm -f myType.goo.ethos