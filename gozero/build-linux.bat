SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
 
 
go build -o ./run/usercenter-api ./app/usercenter/cmd/api/usercenter.go
  
go build -o ./run/basic-api  ./app/basic/cmd/api/basic.go

go build -o ./run/cms-api  ./app/cms/cmd/api/cms.go
 
