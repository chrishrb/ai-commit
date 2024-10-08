package git

type mockCommandExecutor struct {
	output string
}

func (m *mockCommandExecutor) Output() ([]byte, error) {
	return []byte(m.output), nil
}
