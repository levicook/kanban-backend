package endpoints

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_boardCreator(t *testing.T) {

	testNotAcceptable := func() {
		var (
			c = boardCreator{}
			h http.Header
		)

		assert.False(t, c.notAcceptable(h))
	}

	testUnauthorized := func() {
		var (
			c = boardCreator{}
			h http.Header
		)

		assert.True(t, c.unauthorized(h))
	}

	testNotAcceptable()
	testUnauthorized()
}
