package endpoints

import "github.com/dimfeld/httptreemux"

var (
	BoardCreate httptreemux.HandlerFunc
	Board       httptreemux.HandlerFunc
	BoardUpdate httptreemux.HandlerFunc
	Boards      httptreemux.HandlerFunc
	BoardCards  httptreemux.HandlerFunc
	BoardLists  httptreemux.HandlerFunc

	CardCreate httptreemux.HandlerFunc

	ListCreate httptreemux.HandlerFunc

	OrgCreate httptreemux.HandlerFunc
	Org       httptreemux.HandlerFunc
	OrgBoards httptreemux.HandlerFunc
)

func init() {

	BoardCreate = createHandler(func() creator { return &boardCreator{} })
	Board = notImplemented
	BoardUpdate = notImplemented
	Boards = listHandler(func() lister { return &boardLister{} })
	BoardCards = notImplemented
	BoardLists = notImplemented

	CardCreate = notImplemented

	ListCreate = notImplemented

	OrgCreate = notImplemented
	Org = notImplemented
	OrgBoards = notImplemented
}
