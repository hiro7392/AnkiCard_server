package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHandleCardReuest(t *testing.T) {

	r := httptest.NewRequest("GET", "/card/1", nil)
	w := httptest.NewRecorder()

	handleCardRequest(w, r)

	resp := w.Result()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("cannot read test response: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("got = %d, want = 200", resp.StatusCode)
	}

	fmt.Println(body)
	// if string(body) != "hello world!" {
	// 	t.Errorf("got = %s, want = hello world!", body)
	// }

}
