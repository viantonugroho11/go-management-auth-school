package interfacepkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInArray(t *testing.T) {
	res, index := InArray("abc", []string{"abc", "def"})
	assert.Equal(t, res, true)
	assert.Equal(t, index, 0)

	res, index = InArray("abc", []string{"def"})
	assert.Equal(t, res, false)
	assert.Equal(t, index, -1)
}

func TestInArrayNoErr(t *testing.T) {
	res := InArrayNoErr("abc", []string{"abc", "def"})
	assert.Equal(t, res, true)

	res = InArrayNoErr("abc", []string{"def"})
	assert.Equal(t, res, false)
}

func TestInterfaceArrayToString(t *testing.T) {
	res := InterfaceArrayToString([]interface{}{"abc", "def"})
	assert.Equal(t, res, "abc, def")
}

func TestMarshal(t *testing.T) {
	type User struct {
		Name string `json:"Name"`
	}
	object := User{"John Doe"}
	res := Marshal(object)
	assert.Equal(t, res, `{"Name":"John Doe"}`)

	res = Marshal(make(chan int))
	assert.Equal(t, res, "")
}

func TestUnmarshall(t *testing.T) {
	res := Unmarshall(`{"Name": "John Doe"}`)
	assert.Equal(t, res, map[string]interface{}{"Name": "John Doe"})
}

func TestUnmarshallCb(t *testing.T) {
	type User struct {
		Name string `json:"Name"`
	}
	var res User
	UnmarshallCb(`{"Name": "John Doe"}`, &res)
	assert.Equal(t, res, User{"John Doe"})
}

func TestUnmarshallCbInterface(t *testing.T) {
	type User struct {
		Name string `json:"Name"`
	}
	object := User{"John Doe"}
	var res User
	UnmarshallCbInterface(object, &res)
	assert.Equal(t, object, res)
}

func TestMarshallMap(t *testing.T) {
	object := map[string]interface{}{"Name": "John Doe"}
	res := MarshallMap(object)
	assert.Equal(t, res, "{\"Name\":\"John Doe\"}")

	object = map[string]interface{}{"Name": make(chan int)}
	res = MarshallMap(object)
	assert.Equal(t, res, "")
}

func TestInterfaceStringToString(t *testing.T) {
	res := InterfaceStringToString(nil, "")
	assert.Equal(t, res, "")

	res = InterfaceStringToString(map[string]interface{}{}, "Name")
	assert.Equal(t, res, "")

	res = InterfaceStringToString(map[string]interface{}{"Name": "John Doe"}, "Name")
	assert.Equal(t, res, "John Doe")
}

func TestConvert(t *testing.T) {
	type User struct {
		Name string `json:"Name"`
	}
	var res User
	err := Convert("{\"Name\":\"John Doe\"}", res)
	assert.NoError(t, err)
}

func TestConvertStrToInterface(t *testing.T) {
	type User struct {
		Name string `json:"Name"`
	}
	var res User
	err := ConvertStrToInterface("{\"Name\":\"John Doe\"}", res)
	assert.NoError(t, err)
}
