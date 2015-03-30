package endpoints

import "github.com/levicook/kanban-backend/repos"

type (
	transactionFactory func() repos.Transaction
	creatorFactory     func(repos.Transaction) creator
	listerFactory      func(repos.Transaction) lister
)
