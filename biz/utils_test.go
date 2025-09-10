package biz

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetHello(t *testing.T) {
	assert.Equal(t, "hello", GetHello())
}
