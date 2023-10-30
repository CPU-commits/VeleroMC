package engine

type vanillaEngine struct{}

func (*vanillaEngine) Download() {}

func (*vanillaEngine) List() ([]string, error) {
	engine, err := versionRepository.FindOneByEngine(VANILLA_ENGINE)
	if err != nil {
		return nil, err
	}

	return engine.Versions, nil
}

func newVanillaEngine() iEngine {
	return &vanillaEngine{}
}
