package grequests

import (
	"encoding/json"
	"sync"
)

var (
	ropool = sync.Pool{New: func() interface{} {
		return &RequestOptions{}
	}}
	ro    = &RequestOptions{}
	rb, _ = json.Marshal(&ro)
)

func getRequestOptions() *RequestOptions {
	ro := ropool.Get().(*RequestOptions)

	return ro
}

func putRequestOptions(ro *RequestOptions) {
	_ = json.Unmarshal(rb, &ro)
	ropool.Put(ro)
}
