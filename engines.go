package clog

var engines = make(map[string]Engine)

// Logger Engine interface
type Engine interface {
	// Forbids another interfaces
	_dummy()

	Name() string
}

// Only EngineBase structs can implements Engine interface
type EngineBase struct{}

func (*EngineBase) _dummy() {}

func ListEngine() []string {
	var res []string

	for name := range engines {
		res = append(res, name)
	}

	return res
}

func RegisterEngine(engine Engine) {
	engines[engine.Name()] = engine
}
