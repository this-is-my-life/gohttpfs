package configloader

type Configuration struct {
	Listen          *string           `yaml:"listen"`           // address to listen on
	ServePrefix     *string           `yaml:"serve_prefix"`     // url path prefix to serve
	StoragePath     *string           `yaml:"storage_path"`     // path to the storage directory
	DefaultDocument *string           `yaml:"default_document"` // default document to serve
	CacheDuration   *int              `yaml:"cache_duration"`   // cache duration in seconds
	HideDotfile     *bool             `yaml:"hide_dotfile"`     // hide file/folder starting with dot
	ExtraHeaders    map[string]string `yaml:"extra_headers"`    // extra headers to add to responses
}
