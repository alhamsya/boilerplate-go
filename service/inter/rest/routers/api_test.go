package restRouters

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestRestServer_Register(t *testing.T) {
	app := fiber.New(fiber.Config{

	})
	type fields struct {
		Cfg            *config.MainConfig
		App            *fiber.App
		RestInteractor *RestInteractor
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "When_register_expectSuccess",
			fields: fields{
				Cfg: &config.MainConfig{},
				App: app,
				RestInteractor: &RestInteractor{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rest := &RestServer{
				Cfg:            tt.fields.Cfg,
				App:            tt.fields.App,
				RestInteractor: tt.fields.RestInteractor,
			}
			rest.Register()
		})
	}
}
