# Craftgate Go Client

[![Build Status](https://github.com/craftgate/craftgate-go-client/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/craftgate/craftgate-go-client/actions?query=branch%3Amaster)

This repo contains the Go library for Craftgate API.

## Requirements
- Go 1.19 or newer

## Installation

```sh
go get github.com/craftgate/craftgate-go-client
```

## Usage
You can import `Craftgate` client using:

```go
import (
    craftgate "github.com/craftgate/craftgate-go-client/adapter"
)
```

To access the Craftgate API you'll first need to obtain API credentials (e.g. an API key and a secret key). If you don't already have a Craftgate account, you can signup at [https://craftgate.io](https://craftgate.io)

Once you've obtained your API credentials, you can start using Craftgate by instantiating a `Craftgate` with your credentials.

```go
client, _ := craftgate.New("<YOUR API KEY>", "<YOUR SECRET KEY>", "https://api.craftgate.io")

request := craftgate.SearchInstallmentsRequest{
    BinNumber: "487074",
    Price:     100,
    Currency:  craftgate.Currency_TRY,
}

res, err := client.Installment.SearchInstallments(context.Background(), request)

if err != nil {
    t.Errorf("Error %s", err)
}
```

Also, you can pass the options to make localization for API responses. You can use `tr` or `en` right now.
```go
client, _ := craftgate.New("<YOUR API KEY>", "<YOUR SECRET KEY>", "https://api.craftgate.io", craftgate.WithLocalization("en"))

request := craftgate.SearchInstallmentsRequest{
    BinNumber: "487074",
    Price:     100,
    Currency:  craftgate.Currency_TRY,
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
client, _ := craftgate.New("<YOUR API KEY>", "<YOUR SECRET KEY>", "https://sandbox-api.craftgate.io");

request := craftgate.CreatePaymentRequest{
    Price:     100,
    PaidPrice: 100,
    Currency:  craftgate.Currency_TRY,
    ...
}

res, err := client.Payment.CreatePayment(context.Background(), request)

if err != nil {
    t.Errorf("Error %s", err)
}
```

## Idempotency

Mutating operations accept an optional idempotency key. Set it on the request and the client sends it as the `x-idempotency-key` header, so a request can be safely retried (e.g. after a timeout) without the operation being performed twice — the server returns the result of the first request when it sees a repeated key.

Every request embeds `BaseRequest`, which carries a `HeaderOptions` struct, so the key is available on any request:

```go
request := craftgate.CreatePaymentRequest{
    Price:     100,
    PaidPrice: 100,
    Currency:  craftgate.Currency_TRY,
}
request.HeaderOptions = craftgate.HeaderOptions{IdempotencyKey: uuid.NewString()}

res, err := client.Payment.CreatePayment(context.Background(), request)
```

It can also be set inline through the embedded struct:

```go
err := client.Payment.ExpireCheckoutPayment(context.Background(), craftgate.ExpireCheckoutPaymentRequest{
    BaseRequest: craftgate.BaseRequest{HeaderOptions: craftgate.HeaderOptions{IdempotencyKey: uuid.NewString()}},
    Token:       "456d1297-908e-4bd6-a13b-4be31a6e47d5",
})
```

> Use a fresh key per distinct operation, and reuse the same key when retrying that operation.

> The API honours the key on `POST`, `PATCH` and `DELETE` only. It is ignored on `PUT` endpoints, so retrying one of those is not de-duplicated.

`HeaderOptions` is sent as headers only — `json:"-"` keeps it out of the request body and signature, and `schema:"-"` keeps it out of the query string of read requests.

### Contributions

For all contributions to this client please see the contribution guide [here](CONTRIBUTING.md). By participating in this project, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md).

## Security
If you discover a security vulnerability, please review our [Security Policy](SECURITY.md) for how to report it responsibly.

## License
This project is licensed under the Apache License, Version 2.0 — see the [LICENSE](LICENSE) and [NOTICE](NOTICE.md) files for details.
