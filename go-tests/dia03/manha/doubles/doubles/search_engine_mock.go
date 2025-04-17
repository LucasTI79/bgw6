package doubles

type MockSearchEngine struct {
	SearchByNameWasCalled bool
}

func (m *MockSearchEngine) SearchByName(nombre string) string {
	m.SearchByNameWasCalled = true
	return "12345678"
}

func (m *MockSearchEngine) SearchByPhone(phone string) string {
	return ""
}

func (m *MockSearchEngine) AddEntry(name, phone string) error {
	return nil
}
