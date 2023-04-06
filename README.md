# proto-sample

https://protobuf.dev/getting-started/gotutorial/

```
mkdir myproto
cd myproto
wget https://raw.githubusercontent.com/protocolbuffers/protobuf/main/examples/addressbook.proto
```

Compiling Your Protocol Buffers

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### server

```
[~/go/src/github.com/ludwig125/proto-sample/myproto/server] $ls
addressbook.proto
```

go_package が長いので編集

```
// option go_package = "github.com/protocolbuffers/protobuf/examples/go/tutorialpb";
option go_package = "examples/go/tutorialpb";
```

```
[~/go/src/github.com/ludwig125/proto-sample/myproto/server] $protoc -I=. --go_out=. addressbook.proto
[~/go/src/github.com/ludwig125/proto-sample/myproto/server] $tree
.
├── addressbook.proto
└── examples
    └── go
        └── tutorialpb
            └── addressbook.pb.go

3 directories, 2 files
[~/go/src/github.com/ludwig125/proto-sample/myproto/server] $
```

```
[~/go/src/github.com/ludwig125/proto-sample/myproto/server] $go mod init myproto
go: creating new go.mod: module myproto
go: to add module requirements and sums:
        go mod tidy
[~/go/src/github.com/ludwig125/proto-sample/myproto/server] $go mod tidy
go: finding module for package google.golang.org/protobuf/runtime/protoimpl
go: finding module for package github.com/golang/protobuf/ptypes/timestamp
go: finding module for package google.golang.org/protobuf/reflect/protoreflect
go: downloading github.com/golang/protobuf v1.5.3
go: found github.com/golang/protobuf/ptypes/timestamp in github.com/golang/protobuf v1.5.3
go: found google.golang.org/protobuf/reflect/protoreflect in google.golang.org/protobuf v1.30.0
go: found google.golang.org/protobuf/runtime/protoimpl in google.golang.org/protobuf v1.30.0
[~/go/src/github.com/ludwig125/proto-sample/myproto/server] $ls
addressbook.proto  examples/  go.mod  go.sum
[~/go/src/github.com/ludwig125/proto-sample/myproto/server] $
```

注意点

ローカルのファイルを import するときは、go.mod に書いてあるパッケージ名のあとに対象のディレクトリを指定する。

今回、

go.mod の package 名は`myproto`で、`examples/go/tutorialpb`以下の pb.go を import したいので、import 部分は以下のようになる

```
pb "myproto/examples/go/tutorialpb"
```

[~/go/src/github.com/ludwig125/proto-sample/myproto/server] $go run main.go

### client

```
[~/go/src/github.com/ludwig125/proto-sample/myproto] $cp -r server client
[~/go/src/github.com/ludwig125/proto-sample/myproto] $
```

コードを修正

```
[~/go/src/github.com/ludwig125/proto-sample/myproto/client] $go run main.go
```

server 側にログが出た

```
[~/go/src/github.com/ludwig125/proto-sample/myproto/server] $go run main.go
2023/04/06 17:05:49 Received bool value: abc
```
