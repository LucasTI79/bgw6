package doubles

type SearchEngineDummy struct{}

func (s *SearchEngineDummy) SearchByName(name string) string {
	return ""
}

func (s *SearchEngineDummy) SearchByPhone(phone string) string {
	return ""
}

func (s *SearchEngineDummy) AddEntry(name, phone string) error {
	return nil
}
