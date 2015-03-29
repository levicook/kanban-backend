package models

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {

	testSerialization := func() {
		var (
			boardA Board
			boardB Board
			doc    map[string]interface{}
		)

		boardA = Board{
			Id:    NewBoardId(),
			OrgId: NewOrgId(),
		}

		boardA.CreatedAt = time.Now()
		boardA.UpdatedAt = boardA.CreatedAt
		boardA.CreatedBy = NewUserId()
		boardA.UpdatedBy = NewUserId()

		data, err := json.Marshal(boardA)
		assert.Nil(t, err)
		fmt.Printf("%s\n", data)

		doc = make(map[string]interface{})
		assert.Nil(t, json.Unmarshal(data, &doc))
		assert.Equal(t, 6, len(doc))
		assert.NotEmpty(t, doc["id"])
		assert.NotEmpty(t, doc["orgId"])
		assert.NotEmpty(t, doc["createdAt"])
		assert.NotEmpty(t, doc["updatedAt"])
		assert.NotEmpty(t, doc["createdBy"])
		assert.NotEmpty(t, doc["updatedBy"])

		assert.Nil(t, json.Unmarshal(data, &boardB))
		assert.Equal(t, boardA.Id, boardB.Id)
		assert.Equal(t, boardA.OrgId, boardB.OrgId)
		assert.Equal(t, boardA.CreatedAt.UnixNano(), boardB.CreatedAt.UnixNano())
		assert.Equal(t, boardA.UpdatedAt.UnixNano(), boardB.UpdatedAt.UnixNano())
		assert.Equal(t, boardA.CreatedBy, boardB.CreatedBy)
		assert.Equal(t, boardA.UpdatedBy, boardB.UpdatedBy)
	}

	testValidation := func() {
		b := Board{}

		assert.Equal(t, Errors{
			"orgId":     "is required",
			"createdBy": "is required",
			"updatedBy": "is required",
		}, b.Errors())

		b.Id = BoardId{"foo"}
		b.OrgId = OrgId{"bar"}
		b.CreatedBy = UserId{"bin"}
		b.UpdatedBy = UserId{"baz"}
		assert.Equal(t, Errors{
			"id":        "is invalid",
			"orgId":     "is invalid",
			"createdBy": "is invalid",
			"updatedBy": "is invalid",
		}, b.Errors())
	}

	testSerialization()
	testValidation()
}
