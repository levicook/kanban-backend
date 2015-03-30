package routes

import "github.com/levicook/kanban-backend/httpd/endpoints"

var Routes = routeSet{

	// ====================================================================
	// api routes
	// ====================================================================

	{
		"board_create",
		"POST", "/api/boards", endpoints.BoardCreate,
	}, {
		"board",
		"GET", "/api/boards/:boardId", endpoints.Board,
	}, {
		"board_update",
		"PUT", "/api/boards/:boardId", endpoints.BoardUpdate,
	}, {
		"boards",
		"GET", "/api/boards", endpoints.Boards,
	}, {
		"board_cards",
		"GET", "/api/boards/:boardId/cards", endpoints.BoardCards,
	}, {
		"board_lists",
		"GET", "/api/boards/:boardId/lists", endpoints.BoardLists,
	},

	// --------------------------------------------------------------------

	{
		"card_create",
		"POST", "/api/cards", endpoints.CardCreate,
	},

	// --------------------------------------------------------------------

	{
		"list_create",
		"POST", "/api/lists", endpoints.ListCreate,
	},

	// --------------------------------------------------------------------

	{
		"org_create",
		"POST", "/api/orgs", endpoints.OrgCreate,
	}, {
		"orgs",
		"GET", "/api/orgs", endpoints.Org,
	}, {
		"org_boards",
		"GET", "/api/orgs/:orgId/boards", endpoints.OrgBoards,
	},

	// ====================================================================
	// diagnostic routes
	// ====================================================================
	{
		"panic",
		"GET", "/api/panic", endpoints.Panic,
	}, {
		"ping",
		"GET", "/api/ping", endpoints.Ping,
	},
}
