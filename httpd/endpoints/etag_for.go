package endpoints

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/levicook/slog"
)

func etagFor(v interface{}) string {
	hasher := md5.New()
	slog.PanicIf(json.NewEncoder(hasher).Encode(v))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
