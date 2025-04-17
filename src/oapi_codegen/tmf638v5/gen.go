package tmf638v5

//go:generate rm -rf server
//go:generate mkdir -p server
//go:generate oapi-codegen -generate types,server -package myapi TMF638-Service_Inventory_Management-v5.0.0.oas.yaml > server/server.gen.go
