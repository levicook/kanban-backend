package endpoints

type mockTransaction struct {
	commitCalls   int
	rollbackCalls int
}

func (m *mockTransaction) Commit() error {
	m.commitCalls++
	return nil
}

func (m *mockTransaction) Rollback() error {
	m.rollbackCalls++
	return nil
}
