package document

import (
	"testing"
)

func TestActionDescriptor(t *testing.T) {
	testSchema(t, ActionDescriptorType, NewActionDescriptor("test", []string{"test"}, 1, "https://example.com/action"))
}
