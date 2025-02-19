package auth

import (
	"os"
	"sort"

	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func GothSetup() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http:localhost:3000/auth/google/callback"),
	)

	services := map[string]string{
		"google": "Google",
	}

	var keys []string
	for key := range services {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// providerIndex := &ProviderIndex{
	// 	Providers:    keys,
	// 	ProvidersMap: services,
	// }

	return nil
}
