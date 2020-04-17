# Install go swagger
[https://goswagger.io/install.html](https://goswagger.io/install.html)

# Get go packages
Run 
1) *go get*
2)   **GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger**
3) to create spec file run
 **GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models**

# To start
run these commands parallelly
1) *go run main.go* and  
2) for swagger ui *swagger serve -F=swagger swagger.yaml*
