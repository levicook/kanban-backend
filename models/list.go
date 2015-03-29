package models

type (
	ListId struct{ identifier }

	Lists []List

	List struct {
		atStamps
		byStamps

		Id      ListId  `json:"id"`
		BoardId BoardId `json:"boardId"`
	}
)

func NewListId() ListId { return ListId{newIdentifier()} }

func (c Lists) BoardIds() []BoardId {
	ids := []BoardId{}

	for i := range c {
		if id := c[i].BoardId; id.Present() {
			ids = append(ids, id)
		}
	}

	return ids
}

func (m List) Errors() Errors {
	e := Errors{}

	if m.Id.Present() && m.Id.Invalid() {
		e["id"] = "is invalid"
	}

	switch {
	case m.BoardId.Blank():
		e["boardId"] = "is required"
	case m.BoardId.Invalid():
		e["boardId"] = "is invalid"
	}

	m.byStamps.requireValid(e)

	return e
}
