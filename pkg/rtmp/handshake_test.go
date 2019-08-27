package rtmp

import (
	"bytes"
	"github.com/q191201771/nezha/pkg/assert"
	"testing"
)

func TestAll(t *testing.T) {
	var err error
	var hc HandshakeClient
	var hs HandshakeServer
	b := &bytes.Buffer{}
	err = hc.WriteC0C1(b)
	assert.Equal(t, nil, err)
	err = hs.ReadC0C1(b)
	assert.Equal(t, nil, err)
	err = hs.WriteS0S1S2(b)
	assert.Equal(t, nil, err)
	err = hc.ReadS0S1S2(b)
	assert.Equal(t, nil, err)
	err = hc.WriteC2(b)
	assert.Equal(t, nil, err)
	err = hs.ReadC2(b)
	assert.Equal(t, nil, err)
}
