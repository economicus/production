package conf

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	inProduction = flag.Bool("production", false, "production")
)

type App struct {
	InProduction bool
	InsecurePort string
	Domain       string
	Version      string
}

func New() *App {
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Println("[WARNING] MISSING APP ENV: empty domain")
	}
	port := os.Getenv("INSECURE_PORT")
	if port == "" {
		log.Println("[WARNING] MISSING APP ENV: empty insecure port")
	}
	version := os.Getenv("APP_VERSION")
	if version == "" {
		log.Println("[WARNING] MISSING APP ENV: empty app version")
	}
	return &App{
		InProduction: *inProduction,
		InsecurePort: port,
		Domain:       domain,
		Version:      version,
	}
}

func (a *App) Info() {
	fmt.Println("========== APP ==========")
	fmt.Println("Version: ", a.Version)
	fmt.Println("Domain: ", a.Domain)
	fmt.Println("Insecure Port: ", a.InsecurePort)
	fmt.Println("In Production: ", a.InProduction)
}
