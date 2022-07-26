
# Contributing to Craftgate Go Client
As an open-source project, anyone can contribute to the development of `craftgate-go-client`. If you decide to do so, please be aware of the guidelines outlined below.

`craftgate-go-client` is written in Go, in order to contribute to the project, some familiarity with Go knowledge is required.

## Prerequisites
This project uses [gorilla](https://github.com/gorilla/schema/) as its request encoding. 
Required minimum go version is 1.18 and current build tool is gradle.

## Package Structure
The project has a straightforward package structure; the source files are located under the [adapter/](adapter), sample integrations and tests are located under [tests/](tests).

As outlined in the [README](./README.md), the bulk of the project is split into the following categories:

- Adapters: Located under the [adapter](adapter) package, these are go files that are responsible for managing a certain domain
- RestClient: Located under [adapter/rest](adapter/rest), these are utility functions to handle the traffic between backend and client.
- Enumerations and Domain Objects: Located under each adapter file and [model](model), these are enumerations, constants and domain object models that can be used by request and response classes
- Requests: Located under each adapter, these are objects to be used as requests.
- Response: Located under each adapter, these are objects to be used as responses.

## Tests and Coverage
As a payment systems client, it's important to have tests covering critical parts. In addition to tests that test crucial parts of the libraries and utilities, all samples are provided as Junit tests. 
It is strongly advised for all contributors to run all tests/ against the changes introduced in the pull request. If there are no corresponding tests against the changes introduced, owner of the pull request is responsible for adding any relevant tests or samples.
