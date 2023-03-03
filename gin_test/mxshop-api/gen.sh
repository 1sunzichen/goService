function genProto {
   DOMAIN=$1
   DO=$2
    PROTO_PATH=./${DOMAIN}/proto
#    PROTO_PATH=./proto
    GO_OUT_PATH=./${DOMAIN}/proto
    protoc -I=$PROTO_PATH --go_out=$GO_OUT_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$GO_OUT_PATH --go-grpc_opt=paths=source_relative \
    $PROTO_PATH/${DO}.proto

}


genProto user-web user
#genProto  hello