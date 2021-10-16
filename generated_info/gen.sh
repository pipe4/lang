go run github.com/pipe4/lang/cmd/pipe4 parser bnf > pipe4.ebnf
cat pipe4.ebnf | go run github.com/alecthomas/participle/v2/cmd/railroad -w -o ./railroad.html

#go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#rm -f pipe4/ast/*.pb.go
#protoc -I=. --go_opt=module=github.com/pipe4/lang --go_out=. pipe4/ast/*.proto
