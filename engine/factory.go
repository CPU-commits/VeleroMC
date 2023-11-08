package engine

type iEngine interface {
	Download(version string) ([]byte, error)
	Build(javaFile []byte) error
	List() ([]string, error)
	Run(path string) error
}

// Engines
var (
	AVAILABLE_ENGINES = []string{}
)

func EngineFactory(engine string) iEngine {
	exists, err := engineRepository.ExistsEngine(engine)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("engine doesn't exists")
	}

	return newEngine(engine)
}
