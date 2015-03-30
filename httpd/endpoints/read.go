package endpoints

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func readJson(body io.ReadCloser, v interface{}) (err error) {
	var data []byte

	if data, err = ioutil.ReadAll(io.LimitReader(body, 10*mb)); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	if err = body.Close(); err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
