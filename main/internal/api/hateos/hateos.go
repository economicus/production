package hateos

import (
	"economicus/config"
	"fmt"
)

type Hateos struct {
	app *config.AppConfig
}

func NewHateos(app *config.AppConfig) *Hateos {
	return &Hateos{
		app: app,
	}
}

func (h *Hateos) GetServerDomain() string {
	if h.app.InProduction {
		return fmt.Sprintf("https://%s", h.app.Domain)
	} else {
		return "http://localhost:8080"
	}
}

func (h *Hateos) GetHref(routes string) string {
	return h.GetServerDomain() + routes
}

func (h *Hateos) LinkToLogin() map[string]string {
	return map[string]string{
		"rel":    "login",
		"href":   h.GetHref("/v1/login"),
		"method": "POST",
	}
}
