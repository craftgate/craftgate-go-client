#!/usr/bin/env bash
# Simulates a release tagged $1 without pushing a tag: serves HEAD as that version from a local module proxy,
# adds it to a throwaway module with "go get" like a user would, and checks the x-client-version header.
set -euo pipefail
tag="${1:?usage: release-dry-run.sh <tag>}"
if [[ ! "$tag" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "::error::Go module versions need a vX.Y.Z tag, got '$tag'"
  exit 1
fi
module="github.com/craftgate/craftgate-go-client"
repo_dir="$(cd "$(dirname "$0")/../.." && pwd)"
work_dir="$(mktemp -d)"
trap 'chmod -R u+w "$work_dir"; rm -rf "$work_dir"' EXIT

proxy_dir="$work_dir/proxy/$module/@v"
mkdir -p "$proxy_dir"
git -C "$repo_dir" archive --format=zip --prefix="$module@$tag/" HEAD > "$proxy_dir/$tag.zip"
cp "$repo_dir/go.mod" "$proxy_dir/$tag.mod"
echo "{\"Version\":\"$tag\",\"Time\":\"$(date -u +%Y-%m-%dT%H:%M:%SZ)\"}" > "$proxy_dir/$tag.info"
echo "$tag" > "$proxy_dir/list"

mkdir "$work_dir/app"
cat > "$work_dir/app/main.go" <<'GO'
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/craftgate/craftgate-go-client/adapter"
)

func main() {
	client, _ := adapter.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")
	req, err := client.NewRequest(context.Background(), http.MethodGet, "/installment/v1/bins/525864", nil)
	if err != nil {
		panic(err)
	}
	actual := req.Header.Get("x-client-version")
	expected := "craftgate-go-client:" + os.Args[1]
	if actual != expected {
		fmt.Fprintf(os.Stderr, "Expected header %q, got %q\n", expected, actual)
		os.Exit(1)
	}
	fmt.Println("OK x-client-version: " + actual)
}
GO

cd "$work_dir/app"
export GOPROXY="file://$work_dir/proxy,https://proxy.golang.org" GONOSUMDB="$module" GOFLAGS=-mod=mod GOPATH="$work_dir/gopath"
go mod init example.com/release-dry-run >/dev/null 2>&1
go get "$module@$tag"
go run . "${tag#v}"
