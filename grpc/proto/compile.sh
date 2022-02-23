protoc -I./ --go_out=/Users/jiheo/Desktop/workspace/Golang/economicus/main/grpc/proto \
 --go_opt=paths=source_relative --go-grpc_out=/Users/jiheo/Desktop/workspace/Golang/economicus/main/grpc/proto \
 --go-grpc_opt=paths=source_relative quant.proto

 python -m grpc_tools.protoc -I. \
 --python_out=/Users/jiheo/Desktop/workspace/Golang/economicus/quant/grpc-server/proto \
 --grpc_python_out=/Users/jiheo/Desktop/workspace/Golang/economicus/quant/grpc-server/proto quant.proto
