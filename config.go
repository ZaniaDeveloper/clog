package clog

type Size uint64

const (
	Octet Size = 1

	Kilo = 1000 * Octet
	Mega = 1000 * Kilo
	Giga = 1000 * Mega
	Tera = 1000 * Giga

	Kibi = 1024 * Octet
	Mebi = 1024 * Kibi
	Gibi = 1024 * Mebi
	Tebi = 1024 * Gibi
)

type Rotation struct {
	Size          Size   `json:"size" yaml:"size"`
	Timestamp     bool   `json:"append-timestamp" yaml:"append-timestamp"`
	Prefix        string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Suffix        string `json:"suffix,omitempty" yaml:"suffix,omitempty"`
	KeepExtension bool   `json:"keep-extension" yaml:"keep-extension"`
}

type Configuration struct {
	Path        string      `json:"path,omitempty" yaml:"path,omitempty"`
	Rotation    *Rotation   `json:"rotation,omitempty" yaml:"rotation,omitempty"`
	Level       string      `json:"level,omitempty" yaml:"level,omitempty"`
	TimeFormat  string      `json:"time-format,omitempty" yaml:"time-format,omitempty"`
	ThemeObject *Theme      `json:"theme-object,omitempty" yaml:"theme-object,omitempty"`
	ThemeString string      `json:"theme-string,omitempty" yaml:"theme-string,omitempty"`
	Layout      string      `json:"layout,omitempty" yaml:"layout,omitempty"`
	EngineData  interface{} `json:"-" yaml:"-"`
}

func ParseJSONConfig(src string) (*Configuration, error) {
	return nil, nil
}

func ParseYAMLConfig(src string) (*Configuration, error) {
	return nil, nil
}

func ReadJSONConfig(file string) (*Configuration, error) {
	return nil, nil
}

func ReadYAMLConfig(file string) (*Configuration, error) {
	return nil, nil
}
