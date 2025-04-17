package doubles

type SearchEngineStub struct{}

func (s *SearchEngineStub) SearchByName(name string) string {
	return "123456"
}

func (s *SearchEngineStub) SearchByPhone(phone string) string {
	return ""
}

func (s *SearchEngineStub) AddEntry(name, phone string) error {
	return nil
}
