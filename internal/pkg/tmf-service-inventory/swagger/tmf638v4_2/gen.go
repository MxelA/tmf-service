package tmf638v4_2

//go:generate rm -rf server
//go:generate mkdir -p server
//go:generate swagger generate server --quiet --target server --name tmf-service-inventory-v4_2 --spec TMF638_Service_Inventory_Management_API_v4.2.0_beta_swagger.json --template-dir=./templates --exclude-main
