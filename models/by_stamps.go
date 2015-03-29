package models

type byStamps struct {
	CreatedBy UserId `json:"createdBy"`
	UpdatedBy UserId `json:"updatedBy"`
}

func (m byStamps) requireValid(e Errors) {
	switch {
	case m.CreatedBy.Blank():
		e["createdBy"] = "is required"
	case m.CreatedBy.Invalid():
		e["createdBy"] = "is invalid"
	}

	switch {
	case m.UpdatedBy.Blank():
		e["updatedBy"] = "is required"
	case m.UpdatedBy.Invalid():
		e["updatedBy"] = "is invalid"
	}
}
