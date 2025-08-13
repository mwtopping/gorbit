package gorbit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	myname := "Michael"
	err := HelloWorld(myname)
	assert.NoError(t, err)

	myemptyname := ""
	err = HelloWorld(myemptyname)
	assert.Error(t, err)
}
