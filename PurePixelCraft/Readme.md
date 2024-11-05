# Pure Pixel Craft

## Project Architecture Tree

Clean Architecture에서 Infrastructure Layer를 이용한 구조

예상 Project Architecture Tree
```
📂 /cmd
 └── main.go
📂 /internal
 ├── 📂 /core
 │   └── 📂 /services
 │       ├── image_process_service.go
 │       └── data_service.go
 ├── 📂 /interfaces
 │   ├── 📂 /grpc
 │   │   └── grpc_adapter.go
 │   └── 📂 /rest
 │       └── http_adapter.go
 └── 📂 /infrastructure
     ├── 📂 /database
     │   └── maria_db.go
     ├── 📂 /storage
     │   └── s3.go
     └── 📂 /dll
         └── dll_loader.go
📂 /config
 └── config.go
📂 /pkg
 ├── errors
 └── utils
```

## Dependencies

### Gin

### grpc

### grpc-gateway

``` bash
go get -u github.com/gin-gonic/gin
go get -u google.golang.org/grpc
go get -u github.com/grpc-ecosystem/grpc-gateway/v2
```
