package config

func InitConfig() *Config {
	// read config from file
	return &Config{}
}

type Config struct {
	Database DatabaseConfig
	Product  ProductConfig
	Promo    PromoConfig
	User     UserConfig
}

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
}

type ProductConfig struct {
	Host           string
	GetProductPath string
}
type PromoConfig struct {
	Host         string
	GetPromoPath string
}

type UserConfig struct {
	Host            string
	GetUserInfoPath string
}
