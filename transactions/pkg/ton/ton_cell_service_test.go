package ton_test

import (
	"github.com/axidex/api-example/transactions/pkg/ton"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCellService(t *testing.T) {
	payload, err := ton.CreateCell("1230000011")
	require.NoError(t, err)
	t.Logf("payload: %s", payload)
}
