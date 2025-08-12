package tmf638v4_2

//go:generate mkdir -p server
//go:generate swagger generate server --quiet --target server --name tmf-service-order-v4_2 --spec TMF641-ServiceOrdering-v4.2.0.json --template-dir=./templates --exclude-main
