cd D:\work_go\opensrc\go-zero-go-cms\app\act\cmd\api\desc
goctl api go -api act.api -dir ../  --style=goZero --home D:\work_go\opensrc\go-zero-go-cms\deploy\goctl

cd D:\work_go\opensrc\go-zero-go-cms\app\act\cmd\rpc\pb
goctl rpc protoc act.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero --home D:\work_go\opensrc\go-zero-go-cms\deploy\goctl
