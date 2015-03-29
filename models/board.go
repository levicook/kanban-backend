package models

type (
	BoardId struct{ identifier }

	Boards []Board

	Board struct {
		atStamps
		byStamps

		Id    BoardId `json:"id"`
		OrgId OrgId   `json:"orgId"`
	}
)

func NewBoardId() BoardId { return BoardId{newIdentifier()} }

func (c Boards) ListIds() []OrgId {
	ids := []OrgId{}

	for i := range c {
		if id := c[i].OrgId; id.Present() {
			ids = append(ids, id)
		}
	}

	return ids
}

func (m Board) Errors() Errors {
	e := Errors{}

	if m.Id.Present() && m.Id.Invalid() {
		e["id"] = "is invalid"
	}

	switch {
	case m.OrgId.Blank():
		e["orgId"] = "is required"
	case m.OrgId.Invalid():
		e["orgId"] = "is invalid"
	}

	m.byStamps.requireValid(e)

	return e
}
