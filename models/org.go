package models

func NewOrgId() OrgId { return OrgId{newIdentifier()} }

type OrgId struct{ identifier }

type Orgs []Org

type Org struct {
	atStamps
	byStamps

	Id OrgId `json:"id"`
}

func (m Org) Errors() Errors {
	e := Errors{}

	if m.Id.Present() && m.Id.Invalid() {
		e["id"] = "is invalid"
	}

	m.byStamps.requireValid(e)

	return e
}
