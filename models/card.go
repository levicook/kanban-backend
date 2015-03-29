package models

type (
	CardId struct{ identifier }

	Cards []Card

	Card struct {
		atStamps
		byStamps

		Id      CardId  `json:"id"`
		BoardId BoardId `json:"boardId"`
	}
)

func NewCardId() CardId { return CardId{newIdentifier()} }

func (c Cards) BoardIds() []BoardId {
	ids := []BoardId{}

	for i := range c {
		if id := c[i].BoardId; id.Present() {
			ids = append(ids, id)
		}
	}

	return ids
}

func (m Card) Errors() Errors {
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
