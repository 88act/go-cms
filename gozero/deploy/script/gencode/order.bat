cd D:\work_go\opensrc\go-zero-go-cms\app\order\cmd\api\desc
goctl api go -api order.api -dir ../  --style=goZero --home D:\work_go\opensrc\go-zero-go-cms\deploy\goctl

cd D:\work_go\opensrc\go-zero-go-cms\app\order\cmd\rpc\pb
goctl rpc protoc order.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero --home D:\work_go\opensrc\go-zero-go-cms\deploy\goctl
