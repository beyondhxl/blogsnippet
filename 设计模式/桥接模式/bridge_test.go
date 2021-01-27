package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorNotification_Notify(t *testing.T) {
	sender := NewEmailMsgSender([]string{"test@test.com"})
	ntf := NewErrorNotification(sender)
	err := ntf.Notify("test")

	t.Log(ntf.sender)

	assert.Nil(t, err)
}
