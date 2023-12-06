go build -o ./run/usercenter-rpc ./app/usercenter/cmd/rpc/usercenter.go
 
go build -o ./run/usercenter-api ./app/usercenter/cmd/api/usercenter.go
 
go build -o ./run/basic-rpc ./app/basic/cmd/rpc/basic.go
 
go build -o ./run/basic-api  ./app/basic/cmd/api/basic.go

go build -o ./run/cms-api  ./app/cms/cmd/api/cms.go
 