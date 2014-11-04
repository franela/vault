package testutils

type MockValidator struct {
	ValidateMock func()
}

func (m *MockValidator) Validate() { m.ValidateMock() }
