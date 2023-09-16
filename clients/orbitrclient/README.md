# orbitrclient


`orbitrclient` is a Lantah Go SDK package that provides client access to a OrbitR server. It supports all endpoints exposed by the OrbitR API.


## Getting Started
This library is aimed at developers building Go applications that interact with the [Lantah network](https://www.lantah.org/). It allows users to query the network and submit transactions to the network. The recommended transaction builder for Go programmers is [txnbuild](https://github.com/lantah/go/tree/master/txnbuild). Together, these two libraries provide a complete Stellar SDK.

* The [orbitrclient API reference](https://godoc.org/github.com/stellar/go/clients/orbitrclient).
* The [txnbuild API reference](https://godoc.org/github.com/stellar/go/txnbuild).

### Prerequisites
* Go (this repository is officially supported on the last two releases of Go)
* [Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies

### Installing
* `go get github.com/lantah/go/clients/orbitrclient`

### Usage

``` golang
    ...
    import hClient "github.com/lantah/go/clients/orbitrclient"
    ...

    // Use the default pubnet client
    client := hClient.DefaultPublicNetClient

    // Create an account request
    accountRequest := hClient.AccountRequest{AccountID: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU"}

    // Load the account detail from the network
    account, err := client.AccountDetail(accountRequest)
    if err != nil {
        fmt.Println(err)
        return
    }
    // Account contains information about the stellar account
    fmt.Print(account)
```
For more examples, please refer to the original stellar [documentation](https://godoc.org/github.com/stellar/go/clients/orbitrclient).

## Running the tests
Run the unit tests from the package directory: `go test`

## Contributing
To submit improvements and fixes to this library, please see [CONTRIBUTING](../CONTRIBUTING.md).

## License
This project is licensed under the Apache License - see the [LICENSE](../../LICENSE) file for details.
