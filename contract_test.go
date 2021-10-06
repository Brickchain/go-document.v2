package document

import (
	"testing"
)

func TestContract(t *testing.T) {
	testSchema(t, ContractType, NewContract())
}
