package engine

type iEngine interface {
	Download()
	List() ([]string, error)
}

// Engines
const (
	VANILLA_ENGINE = "vanilla"
)

func EngineFactory(engine string) iEngine {
	switch engine {
	case VANILLA_ENGINE:
		return newVanillaEngine()
	default:
		panic("Invalid engine")
	}
}
