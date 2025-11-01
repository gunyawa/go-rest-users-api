# GO REST API — GET /users

This project implements a REST API endpoint `/users` that:

✅ returns users with total number of orders  
✅ supports `city`, `limit`, `offset` parameters  
✅ sorts by `total_orders DESC`, then `id DESC`  
✅ uses `LEFT JOIN + GROUP BY`  
✅ measures query time in header `X-Query-Time`  

### Run instructions

```bash
go mod init myapi
go get github.com/gin-gonic/gin
go get github.com/mattn/go-sqlite3
go run main.go
