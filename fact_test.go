package document

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFactHash(t *testing.T) {
	data := `{"data":{"b":"1", "a":"2", "sub":{"c": "3", "array":["3","1","2", {"1":"2"}]}}}`

	f := Fact{}
	if err := json.Unmarshal([]byte(data), &f); err != nil {
		t.Fatal(err)
	}

	r, err := f.serialize(f.Data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)

	h, err := f.HashData()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(h)
}
