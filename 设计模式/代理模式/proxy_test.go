package proxy

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerate(t *testing.T) {
	want := `package proxy

type UserProxy struct {
	child *User
}

func NewUserProxy(child *User) *UserProxy {
	return &UserProxy{child: child}
}

func (p *UserProxy) Login(username, password string) (r0 error) {
	start := time.Now()

	r0 = p.child.Login(username, password)

	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return r0
}
`
	get, err := generate("./static_proxy.go")
	require.Nil(t, err)
	assert.Equal(t, want, get)
}
