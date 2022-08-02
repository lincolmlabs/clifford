# To Compile

Run<br />
`go build -o clifford`


# To validate Tests BDD with Ginkgo
1. Install Ginkgo:<br />
   `go install github.com/onsi/ginkgo/v2/ginkgo`

2. Execute Ginkgo<br />
   `export PATH=$PATH:$(go env GOPATH)/bin`<br />
   ginkgo build

# Running Tests
Running tests will require Ginkgo configured on path. Once fine, run this command:<br />
<br />
`go test`