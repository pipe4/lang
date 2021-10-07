go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
rm -f pipe4/ast/*.pb.go
protoc -I=. --go_opt=module=github.com/pipe4/lang --go_out=. pipe4/ast/*.proto
