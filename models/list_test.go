package models

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {

	testSerialization := func() {
		var (
			listA List
			listB List
			doc   map[string]interface{}
		)

		listA = List{
			Id:      NewListId(),
			BoardId: NewBoardId(),
		}

		listA.CreatedAt = time.Now()
		listA.UpdatedAt = listA.CreatedAt
		listA.CreatedBy = NewUserId()
		listA.UpdatedBy = NewUserId()

		data, err := json.Marshal(listA)
		assert.Nil(t, err)
		fmt.Printf("%s\n", data)

		doc = make(map[string]interface{})
		assert.Nil(t, json.Unmarshal(data, &doc))
		assert.Equal(t, 6, len(doc))
		assert.NotEmpty(t, doc["id"])
		assert.NotEmpty(t, doc["boardId"])
		assert.NotEmpty(t, doc["createdAt"])
		assert.NotEmpty(t, doc["updatedAt"])
		assert.NotEmpty(t, doc["createdBy"])
		assert.NotEmpty(t, doc["updatedBy"])

		assert.Nil(t, json.Unmarshal(data, &listB))
		assert.Equal(t, listA.Id, listB.Id)
		assert.Equal(t, listA.BoardId, listB.BoardId)
		assert.Equal(t, listA.CreatedAt.UnixNano(), listB.CreatedAt.UnixNano())
		assert.Equal(t, listA.UpdatedAt.UnixNano(), listB.UpdatedAt.UnixNano())
		assert.Equal(t, listA.CreatedBy, listB.CreatedBy)
		assert.Equal(t, listA.UpdatedBy, listB.UpdatedBy)
	}

	testValidation := func() {
		b := List{}

		assert.Equal(t, Errors{
			"boardId":   "is required",
			"createdBy": "is required",
			"updatedBy": "is required",
		}, b.Errors())

		b.Id = ListId{"foo"}
		b.BoardId = BoardId{"bar"}
		b.CreatedBy = UserId{"bin"}
		b.UpdatedBy = UserId{"baz"}
		assert.Equal(t, Errors{
			"id":        "is invalid",
			"boardId":   "is invalid",
			"createdBy": "is invalid",
			"updatedBy": "is invalid",
		}, b.Errors())
	}

	testSerialization()
	testValidation()
}
