# Craftgate Go Client

[![Build Status](https://github.com/craftgate/craftgate-go-client/workflows/Craftgate%20Go%20Client%20CI/badge.svg?branch=master)](https://github.com/craftgate/craftgate-go-client/actions)

This repo contains the Go library for Craftgate API.

## Requirements
- Go 1.18 or newer

## Installation

```sh
go get github.com/craftgate/craftgate-go-client/v1
```

## Usage
You can import `Craftgate` client using:

```go
import (
    "github.com/craftgate/craftgate-go-client/v1"
)
```

To access the Craftgate API you'll first need to obtain API credentials (e.g. an API key and a secret key). If you don't already have a Craftgate account, you can signup at [https://craftgate.io](https://craftgate.io)

Once you've obtained your API credentials, you can start using Craftgate by instantiating a `Craftgate` with your credentials.

```go
client := craftgate.New("<YOUR API KEY>", "<YOUR SECRET KEY>")

request := adapter.SearchInstallmentsRequest{
    BinNumber: "487074",
    Price:     100,
    Currency:  craftgate.Currency(craftgate.TRY),
}

res, err := client.Installment.SearchInstallments(context.Background(), request)

if err != nil {
    t.Errorf("Error %s", err)
}
```

You should use production API servers at `https://api.craftgate.io` for real world. For testing purposes, please use the sandbox URL `https://sandbox-api.craftgate.io`.

## Examples
Included in the project are a number of examples that cover almost all use-cases. Refer to [the `tests/` folder](./tests/)] for more info.

### Running the Examples
If you've cloned this repo on your development machine and wish to run the examples you can run an example with the command `go test ./...` or run single test with the command `go test tests/installment_test.go`

### Credit Card Payment Use Case
Let's quickly review an example where we implement a credit card payment scenario.

> For more examples covering almost all use-cases, check out the [examples in the `tests/` folder](./tests)

```go
Craftgate := CraftgateClient("<YOUR API KEY>", "<YOUR SECRET KEY>", "https://sandbox-api.craftgate.io");

res, err = craftgate.payment().createPayment(request);
fmt.println("Create Payment Result: %s", res);
```

### Contributions

For all contributions to this client please see the contribution guide [here](CONTRIBUTING.md).

## License
MIT
