package document

import (
	"testing"
)

func TestControllerDescriptor(t *testing.T) {
	testSchema(t, ControllerDescriptorType, NewControllerDescriptor("test"))
}
