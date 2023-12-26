module github.com/iegad/gox/hall

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/google/uuid v1.5.0
	github.com/iegad/gox/frm v0.0.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/iegad/gox/pb v0.0.1 // indirect
	golang.org/x/net v0.17.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

go 1.21.5

replace github.com/iegad/gox/frm v0.0.1 => ../frm

replace github.com/iegad/gox/pb v0.0.1 => ../pb
