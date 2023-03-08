package grequests

import "testing"

func Test_getRequestOptions(t *testing.T) {
	for i := 0; i <= 1; i++ {
		ro := getRequestOptions()
		t.Logf("1: %p, %+v", ro, ro.Auth)
		ro.Auth = []string{"a", "b"}
		t.Logf("2: %p, %+v", ro, ro.Auth)
		putRequestOptions(ro)
	}
}
