cd D:\work_go\opensrc\go-zero-looklook\app\payment\cmd\api\desc
goctl api go -api payment.api -dir ../  --style=goZero --home D:\work_go\opensrc\go-zero-looklook\deploy\goctl

cd D:\work_go\opensrc\go-zero-looklook\app\payment\cmd\rpc\pb
goctl rpc protoc payment.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero --home D:\work_go\opensrc\go-zero-looklook\deploy\goctl
