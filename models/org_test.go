package models

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrg(t *testing.T) {

	testSerialization := func() {
		var (
			orgA Org
			orgB Org
			doc  map[string]interface{}
		)

		orgA = Org{
			Id: NewOrgId(),
		}

		orgA.CreatedAt = time.Now()
		orgA.UpdatedAt = orgA.CreatedAt
		orgA.CreatedBy = NewUserId()
		orgA.UpdatedBy = NewUserId()

		data, err := json.Marshal(orgA)
		assert.Nil(t, err)
		fmt.Printf("%s\n", data)

		doc = make(map[string]interface{})
		assert.Nil(t, json.Unmarshal(data, &doc))
		assert.Equal(t, 5, len(doc))
		assert.NotEmpty(t, doc["id"])
		assert.NotEmpty(t, doc["createdAt"])
		assert.NotEmpty(t, doc["updatedAt"])
		assert.NotEmpty(t, doc["createdBy"])
		assert.NotEmpty(t, doc["updatedBy"])

		assert.Nil(t, json.Unmarshal(data, &orgB))
		assert.Equal(t, orgA.Id, orgB.Id)
		assert.Equal(t, orgA.CreatedAt.UnixNano(), orgB.CreatedAt.UnixNano())
		assert.Equal(t, orgA.UpdatedAt.UnixNano(), orgB.UpdatedAt.UnixNano())
		assert.Equal(t, orgA.CreatedBy, orgB.CreatedBy)
		assert.Equal(t, orgA.UpdatedBy, orgB.UpdatedBy)
	}

	testValidation := func() {
		org := Org{}

		assert.Equal(t, Errors{
			"createdBy": "is required",
			"updatedBy": "is required",
		}, org.Errors())

		org.Id = OrgId{"foo"}
		org.CreatedBy = UserId{"bin"}
		org.UpdatedBy = UserId{"baz"}
		assert.Equal(t, Errors{
			"id":        "is invalid",
			"createdBy": "is invalid",
			"updatedBy": "is invalid",
		}, org.Errors())
	}

	testSerialization()
	testValidation()
}
