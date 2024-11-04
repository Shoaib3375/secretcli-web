package cmd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mahinops/secretcli-web/internal/app/auth"
	"github.com/mahinops/secretcli-web/internal/app/secret"
	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/common"
	"gorm.io/gorm"
)

func RegisterWebRoutes(router *chi.Mux, db *gorm.DB, commonConfig *common.CommonConfig, renderer *tmplrndr.Renderer) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "index", nil) // Use the renderer instance
	})

	auth.RegisterWebRoutes(router, db, renderer)
	secret.RegisterWebRoutes(router, db, commonConfig, renderer)
}

func (a *App) StartWebServer(port string) {
	log.Println("Server is running on port " + port + "...")
	if err := http.ListenAndServe(port, a.Router); err != nil {
		log.Fatal(err)
	}
}
