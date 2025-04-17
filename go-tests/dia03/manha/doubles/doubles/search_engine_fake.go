package doubles

type FakeSearchEngine struct {
	// key: Nome, value: Phone
	DB map[string]string
}

func (m *FakeSearchEngine) SearchByName(name string) string {
	return m.DB[name]
}

func (m *FakeSearchEngine) SearchByPhone(phone string) string {
	for key, value := range m.DB {
		if value == phone {
			return key
		}
	}
	return ""
}

func (m *FakeSearchEngine) AddEntry(name, phone string) error {
	if m.DB == nil {
		m.DB = map[string]string{}
	}
	m.DB[name] = phone
	return nil
}
