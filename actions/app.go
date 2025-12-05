package actions

import (
	"github.com/gobuffalo/middleware/csrf"
	"net/http"
	"os"
	"sync"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/middleware/i18n"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

var (
	app     *buffalo.App
	appOnce sync.Once
	T       *i18n.Translator
)

func HelloHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{
		"message": "Hello from Buffalo Nano-App!",
	}))
}

func App() *buffalo.App {
	err := os.Setenv("WEBPACK_MANIFEST", "public/assets/manifest.json")
	if err != nil {
		print("Could not set manifest pah", err)
	}
	app := buffalo.New(buffalo.Options{
		Env:         ENV,
		SessionName: "_buffalo_go_web_hello_session",
	})

	// Serve static files
	//app.ServeFiles("/assets", http.FS(public.FS()))
	app.ServeFiles("/assets", http.Dir("public/assets"))

	app.GET("/.well-known/appspecific/com.chrome.devtools.json/", func(c buffalo.Context) error {
		return c.Render(200, r.String("{}"))
	})

	app.GET("/", HomeHandler)
	app.GET("/api/hello", HelloHandler)

	app.Use(csrf.New)

	return app
}
