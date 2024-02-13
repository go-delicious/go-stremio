package manifest

type MediaType string

const (
	Movie   MediaType = "movie"
	Series  MediaType = "series"
	Channel MediaType = "channel"
	Tv      MediaType = "tv"
)

type Manifest struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Version       string         `json:"version"`
	Resources     []interface{}  `json:"resources"`
	Types         []MediaType    `json:"types"`
	Catalogs      []Catalog      `json:"catalogs"`
	IDPrefixes    []string       `json:"idPrefixes,omitempty"`
	AddonCatalogs []Catalog      `json:"addonCatalogs,omitempty"`
	Config        []Config       `json:"config,omitempty"`
	Background    string         `json:"background,omitempty"`
	Logo          string         `json:"logo,omitempty"`
	ContactEmail  string         `json:"contactEmail,omitempty"`
	BehaviorHints *BehaviorHints `json:"behaviorHints,omitempty"`
}

type Resource struct {
	Name       string      `json:"name"`
	Types      []MediaType `json:"types,omitempty"`
	IDPrefixes []string    `json:"idPrefixes,omitempty"`
}

type Catalog struct {
	Type  MediaType `json:"type"`
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Extra []Extra   `json:"extra,omitempty"`
}

type Meta struct {
	MediaType MediaType `json:"type"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Poster    string    `json:"poster,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
}

type CatalogResponse struct {
	Metas []Meta `json:"metas"`
}

type Extra struct {
	Name         string   `json:"name"`
	IsRequired   bool     `json:"isRequired,omitempty"`
	Options      []string `json:"options,omitempty"`
	OptionsLimit int      `json:"optionsLimit,omitempty"`
}

type Config struct {
	Key      string   `json:"key"`
	Type     string   `json:"type"`
	Default  string   `json:"default,omitempty"`
	Title    string   `json:"title,omitempty"`
	Options  []string `json:"options,omitempty"`
	Required bool     `json:"required,omitempty"`
}

type BehaviorHints struct {
	Adult                 bool `json:"adult,omitempty"`
	P2P                   bool `json:"p2p,omitempty"`
	Configurable          bool `json:"configurable,omitempty"`
	ConfigurationRequired bool `json:"configurationRequired,omitempty"`
}

type ManifestOption func(*Manifest)

func WithID(id string) ManifestOption {
	return func(m *Manifest) {
		m.ID = id
	}
}

func WithName(name string) ManifestOption {
	return func(m *Manifest) {
		m.Name = name
	}
}

func WithDescription(description string) ManifestOption {
	return func(m *Manifest) {
		m.Description = description
	}
}

func WithVersion(version string) ManifestOption {
	return func(m *Manifest) {
		m.Version = version
	}
}

func WithResources(resources []interface{}) ManifestOption {
	return func(m *Manifest) {
		m.Resources = resources
	}
}

func WithBackground(background string) ManifestOption {
	return func(m *Manifest) {
		m.Background = background
	}
}

func WithLogo(logo string) ManifestOption {
	return func(m *Manifest) {
		m.Logo = logo
	}
}

func WithContactEmail(contactEmail string) ManifestOption {
	return func(m *Manifest) {
		m.ContactEmail = contactEmail
	}
}

func WithBehaviorHints(behaviorHints *BehaviorHints) ManifestOption {
	return func(m *Manifest) {
		m.BehaviorHints = behaviorHints
	}
}

func IsAdult() ManifestOption {
	return func(m *Manifest) {
		m.BehaviorHints.Adult = true
	}
}

func IsP2P() ManifestOption {
	return func(m *Manifest) {
		m.BehaviorHints.P2P = true
	}
}

func IsConfigurable() ManifestOption {
	return func(m *Manifest) {
		m.BehaviorHints.Configurable = true
	}
}

func RequiresConfiguration() ManifestOption {
	return func(m *Manifest) {
		m.BehaviorHints.ConfigurationRequired = true
	}
}

func New(options ...ManifestOption) *Manifest {
	// Initialize the Manifest with default values
	m := &Manifest{
		ID:            "org.stremio.example",
		Name:          "Example Addon",
		Description:   "An example Stremio addon",
		Version:       "1.0.0",
		Resources:     []interface{}{},
		Types:         []MediaType{},
		Catalogs:      []Catalog{},
		BehaviorHints: &BehaviorHints{},
	}

	// Apply the provided options
	for _, option := range options {
		option(m)
	}

	return m
}
