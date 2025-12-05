package grifts

import (
	"buffalo-go-web-hello/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
