package repos

type Transaction interface {
	Commit() error
	Rollback() error
}

func NewTransaction() Transaction {
	return nil
}
