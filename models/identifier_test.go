package models

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_identifier(t *testing.T) {
	var (
		idA identifier
		idB identifier
	)

	idA = newIdentifier()
	assert.False(t, idA.Blank())
	assert.False(t, idA.Invalid())
	assert.True(t, idA.Present())
	assert.True(t, idA.Valid())

	data, err := json.Marshal(idA)
	require.Nil(t, err)
	fmt.Printf("%s\n", data)

	require.Nil(t, json.Unmarshal(data, &idB))
	assert.Equal(t, idA, idB)
}
