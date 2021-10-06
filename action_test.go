package document

import (
	"testing"
)

func TestAction(t *testing.T) {
	testSchema(t, ActionType, NewAction([]string{"test"}))
}
