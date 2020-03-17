set PROTOC_HOME=D:\Tools\protoc-3.10.1-win64
set PROTOC=%PROTOC_HOME%\bin\protoc.exe
set PROTOC_OPTS=-ID:\Go\src -I%PROTOC_HOME%\include -I%GOPATH%\src -I.\third_party\googleapis

%PROTOC% %PROTOC_OPTS% github.com/zxfishhack/grpc-gateway-examples/pkgs/proto/examples.proto ^
  --gogo_out=plugins=grpc:%GOPATH%\src --swagger_out=logtostderr=true,grpc_api_configuration=.\pkgs\proto\examples.yaml:pkgs/assets/web/swagger ^
  --grpc-gateway_out=logtostderr=true,grpc_api_configuration=.\pkgs\proto\examples.yaml:%GOPATH%\src

