#Error: protoc-gen-go: program not found or is not executable Please specify a program 
#using absolute path or make sure the program is available in your PATH system variable
#Resolve the above issue make sure yo configured the GO PATH as follows
# 1. export GOROOT=/usr/local/go
# 2. export GOPATH=/Users/sunil/go
# 3. export GOBIN=$GOPATH/bin
# 4. export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

#Use this command in Mac or Linux operating systems
protoc --proto_path=src/simple --go_out=src/simple --go_opt=paths=source_relative src/simple/simple.proto
protoc --proto_path=src/enums --go_out=src/enums --go_opt=paths=source_relative src/enums/enums.proto
protoc --proto_path=src/complex --go_out=src/complex --go_opt=paths=source_relative src/complex/address.proto