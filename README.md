# claims-provider

Provides claims for user by email via gRPC.

Run
- Local: `go run . serve`
- Docker: `docker build -t claims-provider . && docker run --rm -p 50051:50051 -p 8081:8081 claims-provider`

Configuration (flags or env)
- GRPC_HOST / --grpc-host: gRPC listening host (default "")
- GRPC_PORT / --grpc-port: gRPC listening port (default 50051)
- STORE_CACHE_CONCURRENCY / --store-cache-concurrency: maximum concurrent cache builders (default 10)
- STORE_CACHE_EXPIRE / --store-cache-expire: cache expiration for successful entries (e.g. 60s) (default 60s)
- STORE_CACHE_ERROR_EXPIRE / --store-cache-error-expire: cache expiration for errors (default 10s)
- STORE_CACHE_CAPACITY / --store-cache-capacity: cache capacity (default 1000)
- STORE_DB_TIMEOUT / --store-db-timeout: DB query timeout (default 5s)
- Probe and PG options are provided by github.com/webtor-io/common-services (HTTP probe on 8081).

API
- gRPC service: ClaimsProvider
- Method: Get(GetRequest{email}) -> GetResponse{context, claims}

Notes
- Basic input validation returns InvalidArgument if email is empty.
- Server performs graceful shutdown for gRPC.