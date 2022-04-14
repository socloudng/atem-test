@REM %USERPROFILE%\.nuget\packages\grpc.tools\2.40.0\tools\windows_x64\protoc.exe ^
md pbs
protoc --proto_path=./protos --go_out=plugins=grpc:./pbs hello.proto
protoc --proto_path=./protos --go_out=plugins=grpc:./pbs student.proto