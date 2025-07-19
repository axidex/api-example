package ton_test

import (
	"encoding/base64"
	"fmt"
	"github.com/axidex/api-example/transactions/pkg/ton"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xssnick/tonutils-go/tvm/cell"
	"testing"
)

func TestCellService(t *testing.T) {
	payload, err := ton.CreateCell("1230000011")
	require.NoError(t, err)
	t.Logf("payload: %s", payload)
}

func TestEncodeCell(t *testing.T) {
	before := "12345667"
	payload, err := ton.CreateCell(before)
	require.NoError(t, err)
	fmt.Println(payload)

	payloadDecoded, err := base64.StdEncoding.DecodeString(payload)
	require.NoError(t, err)

	c, err := cell.FromBOC(payloadDecoded)
	require.NoError(t, err)

	after, err := ton.DecodeStringPayload(c)
	require.NoError(t, err)

	assert.Equal(t, before, after)
}
