package endpoints

type mockTransaction struct {
	commitCalls   int
	rollbackCalls int
}

func (m *mockTransaction) Commit() error {
	m.commitCalls += 1
	return nil
}

func (m *mockTransaction) Rollback() error {
	m.rollbackCalls += 1
	return nil
}
