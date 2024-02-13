package addon

import (
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"

	"github.com/go-delicious/go-stremio/manifest"
)

type Addon struct {
	Manifest *manifest.Manifest
	Mux      *http.ServeMux
}

func New(m *manifest.Manifest) *Addon {
	return &Addon{
		Manifest: m,
		Mux:      http.NewServeMux(),
	}
}

func (a *Addon) addType(mediaType manifest.MediaType) {
	found := false
	for _, t := range a.Manifest.Types {
		if t == mediaType {
			found = true
			break
		}
	}

	if !found {
		a.Manifest.Types = append(a.Manifest.Types, mediaType)
	}
}

func (a *Addon) HandleMovieCatalog(name string, h func(w http.ResponseWriter, r *http.Request)) {
	// normalise the name
	safeName := strings.ReplaceAll(strings.ToLower(name), " ", "-")

	// add catalog to the resources array, if not already added
	if !a.resourceExists("catalog") {
		a.Manifest.Resources = append(a.Manifest.Resources, "catalog")

		// add get all catalogs handler
		a.Mux.HandleFunc("GET /catalog", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(a.Manifest.Catalogs)
		})
	}

	a.Manifest.Catalogs = append(a.Manifest.Catalogs, manifest.Catalog{
		ID:   safeName,
		Type: manifest.Movie,
		Name: name,
	})

	// add movie type if not already added
	if !a.typeExists(manifest.Movie) {
		a.addType(manifest.Movie)

		// add get all movie catalogs handler
		a.Mux.HandleFunc("GET /catalog/movie", func(w http.ResponseWriter, r *http.Request) {

			// loop through the manifest catalogs and send them as JSON
			var catalogs []manifest.Catalog
			for _, c := range a.Manifest.Catalogs {
				if c.Type == manifest.Movie {
					catalogs = append(catalogs, c)
				}
			}

			// send the list as JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(catalogs)
		})
	}

	// pattern is /catalog/movie/{name}.json
	a.Mux.HandleFunc("GET /catalog/movie/"+safeName+".json", h)
}

func (a *Addon) HandleSeriesCatalog(name string, h func(w http.ResponseWriter, r *http.Request)) {
	// make the name safe for URL
	safeName := strings.ReplaceAll(strings.ToLower(name), " ", "-")

	// add catalog to the resources array, if not already added
	if !a.resourceExists("catalog") {
		a.Manifest.Resources = append(a.Manifest.Resources, "catalog")

		// add get all catalogs handler
		a.Mux.HandleFunc("GET /catalog", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(a.Manifest.Catalogs)
		})
	}

	a.Manifest.Catalogs = append(a.Manifest.Catalogs, manifest.Catalog{
		ID:   safeName,
		Type: manifest.Series,
		Name: name,
	})

	// add series type if not already added
	if !a.typeExists(manifest.Series) {
		a.addType(manifest.Series)

		// add get all series catalogs handler
		a.Mux.HandleFunc("GET /catalog/series/", func(w http.ResponseWriter, r *http.Request) {

			// loop through the manifest catalogs and send them as JSON
			var catalogs []manifest.Catalog
			for _, c := range a.Manifest.Catalogs {
				if c.Type == manifest.Series {
					catalogs = append(catalogs, c)
				}
			}

			// send the list as JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(catalogs)
		})
	}

	// pattern is /catalog/series/{name}.json
	a.Mux.HandleFunc("GET /catalog/series/"+safeName+".json", h)
}

// helper function to check if a resource is already added
func (a *Addon) resourceExists(resource string) bool {
	for _, r := range a.Manifest.Resources {
		if r == resource {
			return true
		}
	}
	return false
}

// helper function to check if type is already added
func (a *Addon) typeExists(mediaType manifest.MediaType) bool {
	for _, t := range a.Manifest.Types {
		if t == mediaType {
			return true
		}
	}
	return false
}

// Serve the cached manifest. This method uses RLock() for concurrent reads.
func (a *Addon) HandleManifest() {
	a.Mux.HandleFunc("GET /manifest.json", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(a.Manifest); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "failed to encode manifest"})
		}
	})
}

func (a *Addon) ListenAndServe(port string) {
	a.HandleManifest()

	log.Printf("Addon running on http://localhost:%s/manifest.json", port)
	log.Printf("To add addon to streamio, go to stremio://localhost:%s/manifest.json", port)
	if err := http.ListenAndServe(":"+port, DefaultMiddleware(a.Mux)); err != nil {
		log.Fatal(err)
	}
}

func DefaultMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Cache-Control", "public, max-age=3600")
		w.Header().Set("Content-Type", "application/json")
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
