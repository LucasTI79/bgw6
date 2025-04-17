package service

type SearchEngine struct {
	Engine Engine
}

type Engine interface {
	SearchByName(name string) string
	SearchByPhone(phone string) string
	AddEntry(name, phone string) error
}

func (s *SearchEngine) GetVersion() string {
	return "1.0.0"
}

func NewSearchEngine(engine Engine) *SearchEngine {
	return &SearchEngine{
		Engine: engine,
	}
}
