package endpoints

import (
	"crypto/md5"
	"encoding/json"
)

func etagFor(v interface{}) string {
	hasher := md5.New()
	panicIf(json.NewEncoder(hasher).Encode(v))
	return sprintf("%x", hasher.Sum(nil))
}
