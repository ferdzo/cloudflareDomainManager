package secrets

import (
	"github.com/spf13/viper"
	"log"
)

func LoadSecrets() *Secret {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	secrets := &Secret{
		X_Auth_Key:   viper.GetString("X_AUTH_KEY"),
		X_Auth_Email: viper.GetString("X_AUTH_EMAIL"),
		Zone_ID:      viper.GetString("ZONE_ID"),
	}
	return secrets
}
