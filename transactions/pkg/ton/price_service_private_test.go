package ton

//import (
//	"context"
//	"github.com/axidex/api-example/transactions/pkg/ton/mocks"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//	"go.uber.org/mock/gomock"
//	"testing"
//)
//
////go:generate mockgen -package=mocks -destination=./mocks/logger_mock.go github.com/axidex/api-example/server/pkg/logger Logger
//func TestGetPrice(t *testing.T) {
//	ctx := context.Background()
//	ctrl := gomock.NewController(t)
//	mockLogger := mocks.NewMockLogger(ctrl)
//	service := NewPriceService(mockLogger)
//	price, err := service.getPrice(ctx)
//	require.NoError(t, err)
//	assert.True(t, price > 0)
//	t.Log(price)
//}
