package ton

import (
	"context"
	"github.com/JulianToledano/goingecko/v3/api"
	"github.com/axidex/api-example/server/pkg/logger"
	"sync"
	"time"
)

type PriceService interface {
	Start()
	GetPrice() float64
}

type PriceServiceCoinGecko struct {
	client    *api.Client
	cacheMu   sync.RWMutex
	usdPrice  float64
	updatedAt time.Time
	logger    logger.Logger
}

func NewPriceService(ctx context.Context, logger logger.Logger) (*PriceServiceCoinGecko, error) {
	client := api.NewDefaultClient()

	service := &PriceServiceCoinGecko{
		client:  client,
		logger:  logger,
		cacheMu: sync.RWMutex{},
	}

	basePrice, err := service.getPrice(ctx)
	if err != nil {
		return nil, err
	}
	
	service.usdPrice = basePrice

	return service, nil
}

func (s *PriceServiceCoinGecko) Start() {
	go s.runBackgroundUpdater()
}

func (s *PriceServiceCoinGecko) runBackgroundUpdater() {
	ctx := context.Background()

	s.updateCache(ctx)

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		s.updateCache(ctx)
	}
}

func (s *PriceServiceCoinGecko) updateCache(ctx context.Context) {
	price, err := s.getPrice(ctx)
	if err != nil {
		s.logger.Error(ctx, "failed to get price", logger.NewAttribute("error", err.Error()))
		return
	}

	s.cacheMu.Lock()
	s.usdPrice = price
	s.updatedAt = time.Now().UTC()
	s.cacheMu.Unlock()

	s.logger.Info(ctx, "updated price", logger.NewAttribute("price", price))
}

func (s *PriceServiceCoinGecko) getPrice(ctx context.Context) (float64, error) {
	data, err := s.client.CoinsId(ctx, "the-open-network")
	if err != nil {
		return 0, err
	}

	return data.MarketData.CurrentPrice.Usd, nil
}

func (s *PriceServiceCoinGecko) GetPrice() float64 {
	s.cacheMu.RLock()
	defer s.cacheMu.RUnlock()
	return s.usdPrice
}
