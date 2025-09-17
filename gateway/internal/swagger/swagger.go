package swagger

import (
	"embed"
	"io/fs"
	"net/http"
)

var swaggerUI embed.FS

func SwaggerHandler(swaggerJSON []byte) http.Handler {
	mux := http.NewServeMux()

	// Serve Swagger UI
	swaggerFS, _ := fs.Sub(swaggerUI, "swagger-ui")
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.FS(swaggerFS))))

	// Serve Swagger JSON
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(swaggerJSON)
	})

	// Redirect to Swagger UI
	mux.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger-ui/", http.StatusMovedPermanently)
	})

	return mux
}
