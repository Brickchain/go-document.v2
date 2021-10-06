package document

import (
	"testing"
)

func TestControllerBinding(t *testing.T) {
	realmKey := newKey(t)
	realm := NewRealmDescriptor("test", realmKey, "https://example.com/realm")
	testSchema(t, ControllerBindingType, NewControllerBinding(realm))
}
