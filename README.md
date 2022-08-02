# Craftgate Go Client

[![Build Status](https://github.com/craftgate/craftgate-go-client/workflows/Craftgate%20Go%20Client%20CI/badge.svg?branch=master)](https://github.com/craftgate/craftgate-go-client/actions)
[![Maven Central](https://maven-badges.herokuapp.com/maven-central/io.craftgate/craftgate/badge.svg)](https://maven-badges.herokuapp.com/maven-central/io.craftgate/craftgate)
[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/craftgate/craftgate-java-client)

This repo contains the Java client for Craftgate API.

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/craftgate/craftgate-java-client)

## Requirements
- Go 1.18 or newer

## Installation

```sh
go get github.com/craftgate/craftgate-go-client@v1.0.0
```

## Usage
You can import `Craftgate` client using:

```go
import (
    "github.com/craftgate/craftgate-go-client"
)
```

To access the Craftgate API you'll first need to obtain API credentials (e.g. an API key and a secret key). If you don't already have a Craftgate account, you can signup at [https://craftgate.io](https://craftgate.io)

Once you've obtained your API credentials, you can start using Craftgate by instantiating a `Craftgate` with your credentials.

```go

Craftgate := CraftgateClient("<YOUR API KEY>", "<YOUR SECRET KEY>", "<BASE URL>")
```

By default the Craftgate client connects to the production API servers at `https://api.craftgate.io`. For testing purposes, please use the sandbox URL `https://sandbox-api.craftgate.io` using the .

```go

Craftgate := CraftgateClient("<YOUR API KEY>", "<YOUR SECRET KEY>", "https://sandbox-api.craftgate.io");
```

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
