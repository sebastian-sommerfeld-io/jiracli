package dummy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Echo(t *testing.T) {
	assert.Equal(t, "abc from dummy", Dummy("abc"))
}
