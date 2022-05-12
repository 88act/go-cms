cd D:\work_go\opensrc\go-zero-go-cms\app\usercenter\cmd\api\desc
goctl api go -api usercenter.api -dir ../  --style=goZero --home D:\work_go\opensrc\go-zero-go-cms\deploy\goctl

cd D:\work_go\opensrc\go-zero-go-cms\app\usercenter\cmd\rpc\pb
goctl rpc protoc usercenter.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero --home D:\work_go\opensrc\go-zero-go-cms\deploy\goctl
