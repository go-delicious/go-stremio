package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-delicious/go-stremio/addon"
	"github.com/go-delicious/go-stremio/catalog"
	"github.com/go-delicious/go-stremio/manifest"
)

func main() {
	addon := addon.New(&manifest.Manifest{
		ID:          "org.stremio.example",
		Name:        "Example Addon",
		Description: "An example Stremio addon",
		Version:     "1.0.0",
	})

	addon.HandleMovieCatalog("Top Movies", TopMoviesHandler)
	addon.HandleMovieCatalog("Recommended Movies", RecommendedMoviesHandler)
	addon.HandleSeriesCatalog("Top Series", TopSeriesHandler)

	addon.ListenAndServe("3000")
}

func TopMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movieCatalog := catalog.NewMovieCatalog(
		catalog.NewMovie("The Dark Knight", "tt0468569"),
		catalog.NewMovie("Inception", "tt1375666"),
	)

	json.NewEncoder(w).Encode(movieCatalog)
}

func RecommendedMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movieCatalog := catalog.NewMovieCatalog(
		catalog.NewMovie("The Shawshank Redemption", "tt0111161"),
		catalog.NewMovie("The Godfather", "tt0068646"),
		catalog.NewMovie("The Dark Knight", "tt0468569"),
	)

	json.NewEncoder(w).Encode(movieCatalog)
}

func TopSeriesHandler(w http.ResponseWriter, r *http.Request) {
	seriesCatalog := catalog.NewSeriesCatalog(
		catalog.NewSeries("Game of Thrones", "tt0944947"),
		catalog.NewSeries("The Walking Dead", "tt1520211"),
	)

	json.NewEncoder(w).Encode(seriesCatalog)
}
