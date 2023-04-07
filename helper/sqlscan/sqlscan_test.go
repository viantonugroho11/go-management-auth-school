package sqlscan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	unmarshalData UnmarshalData = UnmarshalData{"stringval"}
)

func TestScan(t *testing.T) {
	err := unmarshalData.Scan("stringval")
	assert.NoError(t, err)

	err = unmarshalData.Scan([]byte(`{
        "Name": "John Doe"
    }`))
    assert.NoError(t, err)
}
