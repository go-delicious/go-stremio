package catalog

import "github.com/go-delicious/go-stremio/manifest"

func NewMovieCatalog(movies ...manifest.Meta) *manifest.CatalogResponse {
	return &manifest.CatalogResponse{
		Metas: movies,
	}
}

func NewMovie(name string, id string) manifest.Meta {
	return manifest.Meta{
		ID:        id,
		Name:      name,
		MediaType: manifest.Movie,
	}
}

func NewSeriesCatalog(series ...manifest.Meta) *manifest.CatalogResponse {
	return &manifest.CatalogResponse{
		Metas: series,
	}
}

func NewSeries(name string, id string) manifest.Meta {
	return manifest.Meta{
		ID:        id,
		Name:      name,
		MediaType: manifest.Series,
	}
}
