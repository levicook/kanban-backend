package endpoints

import (
	"fmt"

	"github.com/dimfeld/httptreemux"
	"github.com/levicook/kanban-backend/repos"
	"github.com/levicook/slog"
)

var (
	BoardCreate httptreemux.HandlerFunc
	Board       httptreemux.HandlerFunc
	BoardUpdate httptreemux.HandlerFunc
	Boards      httptreemux.HandlerFunc

	BoardCards httptreemux.HandlerFunc
	BoardLists httptreemux.HandlerFunc

	CardCreate httptreemux.HandlerFunc

	ListCreate httptreemux.HandlerFunc

	OrgCreate httptreemux.HandlerFunc
	Org       httptreemux.HandlerFunc
	OrgBoards httptreemux.HandlerFunc
)

func init() {

	BoardCreate = createHandler(
		repos.NewTransaction,
		func(t repos.Transaction) creator {
			return &boardCreator{
				boardRepo: repos.NewBoardRepo(t),
			}
		},
	)

	Board = notImplemented

	BoardUpdate = notImplemented

	Boards = notImplemented

	BoardCards = notImplemented

	BoardLists = notImplemented

	CardCreate = notImplemented

	ListCreate = notImplemented

	OrgCreate = notImplemented

	Org = notImplemented

	OrgBoards = notImplemented
}

var (
	panicIf = slog.PanicIf
	sprintf = fmt.Sprintf
)
