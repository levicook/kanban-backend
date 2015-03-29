package models

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {

	testSerialization := func() {
		var (
			cardA Card
			cardB Card
			doc   map[string]interface{}
		)

		cardA = Card{
			Id:      NewCardId(),
			BoardId: NewBoardId(),
		}

		cardA.CreatedAt = time.Now()
		cardA.UpdatedAt = cardA.CreatedAt
		cardA.CreatedBy = NewUserId()
		cardA.UpdatedBy = NewUserId()

		data, err := json.Marshal(cardA)
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

		assert.Nil(t, json.Unmarshal(data, &cardB))
		assert.Equal(t, cardA.Id, cardB.Id)
		assert.Equal(t, cardA.BoardId, cardB.BoardId)
		assert.Equal(t, cardA.CreatedAt.UnixNano(), cardB.CreatedAt.UnixNano())
		assert.Equal(t, cardA.UpdatedAt.UnixNano(), cardB.UpdatedAt.UnixNano())
		assert.Equal(t, cardA.CreatedBy, cardB.CreatedBy)
		assert.Equal(t, cardA.UpdatedBy, cardB.UpdatedBy)
	}

	testValidation := func() {
		b := Card{}

		assert.Equal(t, Errors{
			"boardId":   "is required",
			"createdBy": "is required",
			"updatedBy": "is required",
		}, b.Errors())

		b.Id = CardId{"foo"}
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
