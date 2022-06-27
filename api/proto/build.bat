echo "go proto"
@.\protoc.exe --go_out=.. gateway.proto
@.\protoc.exe --go_out=plugins=grpc:.. center.proto
@.\protoc.exe --go_out=.. login.proto
@.\protoc.exe --go_out=.. types.proto
@.\protoc.exe --go_out=.. logger.proto
@.\protoc.exe --go_out=.. config.proto
@.\protoc.exe --go_out=.. logic.proto
@.\protoc.exe --go_out=.. battle.proto
@.\protoc.exe --go_out=.. property.proto
@.\protoc.exe --go_out=.. gameddz.proto
@.\protoc.exe --go_out=.. gameAnkots.proto

pause
