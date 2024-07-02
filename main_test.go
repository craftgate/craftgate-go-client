/*
 * Craftgate API Client
 *
 * Contact: destek@craftgate.io
 */

package main

import (
	"context"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {
	t.Run("Create a new APIClient", func(t *testing.T) {

		client, _ := craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

		t.Run("Create a API request that should fail", func(t *testing.T) {
			binNumber, err := client.Installment.RetrieveBinNumber(context.Background(), "487074")
			require.NotNil(t, err)
			require.Nil(t, binNumber)
			assert.Equal(t, true, strings.Contains(err.Error(), "API credential not found"))
		})
	})
}
