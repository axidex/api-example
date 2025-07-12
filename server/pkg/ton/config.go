package ton

type Config struct {
	ConfigUrl     string `mapstructure:"TON_CONFIG_URL"`
	WalletAddress string `mapstructure:"TON_WALLET_ADDRESS"`
}
