package procstore

//go:generate protoc -I $PB_HOME/go/src/ --go_out=plugins=grpc:$PB_HOME/go/src/ $PB_HOME/go/src/dl/apis/procstore/procstore.proto



