package auto

import (
	"encoding/base64"
)

const (
	AssetsBuildDate = "Thu, 14 Apr 2016 21:00:33 GMT"
)

func Assets() map[string][]byte {
	var assets = make(map[string][]byte, 1)

	assets["index.html"], _ = base64.StdEncoding.DecodeString("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")
	return assets
}
