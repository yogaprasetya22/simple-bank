package util

import (
	"github.com/spf13/viper"
)

// Konfigurasi menyimpan semua konfigurasi aplikasi.
// Nilai dibaca oleh Viper dari file konfigurasi atau variabel lingkungan.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_ADDR"`
	ServerAddress string `mapstructure:"ADDR"`
}

// LoadConfig membaca konfigurasi dari variabel file atau lingkungan.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".envrc")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
