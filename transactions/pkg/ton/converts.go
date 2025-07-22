package ton

const (
	Multiplier     = 1.1
	StarsPerDollar = 66.67
	TonInNanotons  = 1_000_000_000
)

func NanoTonsToStars(nanoTons uint64, TonPriceInDollars float64) int {
	tons := float64(nanoTons) / TonInNanotons
	totalUSD := tons * TonPriceInDollars
	starValueUSD := totalUSD * StarsPerDollar * Multiplier

	return int(starValueUSD)
}
