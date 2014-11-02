package testutils

type MockValidator struct {
	CalledValidate bool
}

func (mv *MockValidator) Validate() {
	mv.CalledValidate = true
}
