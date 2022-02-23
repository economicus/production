package config

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	inProduction = flag.Bool("production", false, "production")
)

type AppConfig struct {
	InProduction bool
	Domain       string
	Version      string
}

func NewAppConfig() *AppConfig {
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Fatalf("error while getting domain: empty domain")
	}
	version := os.Getenv("APP_VERSION")
	if version == "" {
		log.Fatalf("error while getting version: empty version")
	}
	return &AppConfig{
		InProduction: *inProduction,
		Domain:       domain,
		Version:      version,
	}
}

func (a *AppConfig) PrintInfo() {
	fmt.Println("========== APP ==========")
	fmt.Println("Version: ", a.Version)
	fmt.Println("Domain: ", a.Domain)
	fmt.Println("In Production: ", a.InProduction)
}
