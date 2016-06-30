// Package static serves static files like CSS, JavaScript, and images.
package static

import (
	"net/http"
	"os"
	"path"

	"github.com/blue-jay/blueprint/controller/status"
	"github.com/blue-jay/blueprint/lib/asset"
	"github.com/blue-jay/blueprint/lib/router"
)

// Load the routes
func Load() {
	// Serve static files
	router.Get("/static/*filepath", Index)
}

// Index maps static files
func Index(w http.ResponseWriter, r *http.Request) {
	// File path
	path := path.Join(asset.Config().Folder, r.URL.Path[1:])

	// Only serve files
	if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
		http.ServeFile(w, r, path)
		return
	}

	status.Error404(w, r)
}
