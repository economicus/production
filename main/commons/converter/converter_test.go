package converter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MainData struct {
	First  string   `json:"first"`
	Second int      `json:"second"`
	Third  SubData  `json:"third"`
	Forth  *SubData `json:"forth"`
}

type SubData struct {
	First  string `json:"first"`
	Second uint   `json:"second"`
}

func TestToMap(t *testing.T) {
	a := assert.New(t)
	data := MainData{
		First:  "Hello",
		Second: 14,
		Third: SubData{
			First:  "world",
			Second: 7,
		},
	}
	res := InterfaceToMap(data)
	fmt.Println(res)
	mf, ok := res["first"]
	s := mf.(string)
	a.Equal("Hello", s, "error testing main first value")
	a.Equal(true, ok, "error while testing main first ok")
	ms, ok := res["second"]
	i := ms.(int)
	a.Equal(14, i, "error testing main second value")
	a.Equal(true, ok, "error while testing main second ok")
	mt, ok := res["third"]
	sub := mt.(map[string]interface{})
	expected := map[string]interface{}{
		"first":  "world",
		"second": uint(7),
	}
	a.Equal(expected, sub, "error testing main third value")
	a.Equal(true, ok, "error while testing main third ok")
}
