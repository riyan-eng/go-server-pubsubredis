package infrastructure

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func NewLocalizer() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	if _, err := bundle.LoadMessageFile("./localize/en.toml"); err != nil {
		fmt.Printf("localize: %v", err.Error())
		os.Exit(1)
	}
	if _, err := bundle.LoadMessageFile("./localize/id.toml"); err != nil {
		fmt.Printf("localize: %v", err.Error())
		os.Exit(1)
	}
	fmt.Println("localize: load successfully")
}

var conlocalizer *i18n.Localizer

func Localizer(c *fiber.Ctx) error {
	header := c.Get("Accept-Language")
	query := c.Query("lang")
	conlocalizer = i18n.NewLocalizer(bundle, header, query)
	return c.Next()
}

func Localize(params any) string {
	switch p := params.(type) {
	case string:
		return conlocalizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: p,
		})
	default:
		return conlocalizer.MustLocalize(p.(*i18n.LocalizeConfig))
	}
}
