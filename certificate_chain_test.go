package document

import (
	"testing"
	"time"
)

func TestCertificate(t *testing.T) {
	rootKey := newKey(t)
	rootPK := newPublicKey(t, rootKey)
	subKey := newKey(t)
	subPK := newPublicKey(t, subKey)

	b := NewCertificate(rootPK, subPK, 100, 3600)
	b.DocumentTypes = []string{"mandate"}

	if b.HasExpired() {
		t.Fatal("Certificate has expired")
	}
}

func TestCertificateHasExpired(t *testing.T) {
	rootKey := newKey(t)
	rootPK := newPublicKey(t, rootKey)
	subKey := newKey(t)
	subPK := newPublicKey(t, subKey)

	b := NewCertificate(rootPK, subPK, 100, 1)
	b.DocumentTypes = []string{"mandate"}

	time.Sleep(time.Second * 2)
	if !b.HasExpired() {
		t.Fatal("Certificate should have expired")
	}
}

func TestCertificateAllowedType(t *testing.T) {
	rootKey := newKey(t)
	rootPK := newPublicKey(t, rootKey)
	subKey := newKey(t)
	subPK := newPublicKey(t, subKey)

	b := NewCertificate(rootPK, subPK, 100, 3600)
	b.DocumentTypes = []string{"*"}

	if !b.AllowedType(NewMandate("abc@test")) {
		t.Fatal("Certificate should have allowed the mandate document type")
	}
}

func TestCertificateAllowedTypeFalse(t *testing.T) {
	rootKey := newKey(t)
	rootPK := newPublicKey(t, rootKey)
	subKey := newKey(t)
	subPK := newPublicKey(t, subKey)

	b := NewCertificate(rootPK, subPK, 100, 3600)
	b.DocumentTypes = []string{"none"}

	if b.AllowedType(NewMandate("abc@test")) {
		t.Fatal("Certificate should not have allowed the mandate document type")
	}
}

func TestCertificateAllowedSubType(t *testing.T) {
	rootKey := newKey(t)
	rootPK := newPublicKey(t, rootKey)
	subKey := newKey(t)
	subPK := newPublicKey(t, subKey)

	b := NewCertificate(rootPK, subPK, 100, 3600)
	b.DocumentTypes = []string{FactType + "#email"}

	if !b.AllowedType(NewFact("email")) {
		t.Fatal("Certificate should have allowed the fact document subtype")
	}
}

func TestCertificateAllowedSubTypeWildcard(t *testing.T) {
	rootKey := newKey(t)
	rootPK := newPublicKey(t, rootKey)
	subKey := newKey(t)
	subPK := newPublicKey(t, subKey)

	b := NewCertificate(rootPK, subPK, 100, 3600)
	b.DocumentTypes = []string{FactType + "#*"}

	if !b.AllowedType(NewFact("email")) {
		t.Fatal("Certificate should have allowed the fact document subtype")
	}
}

func TestCertificateAllowedSubTypeFalse(t *testing.T) {
	rootKey := newKey(t)
	rootPK := newPublicKey(t, rootKey)
	subKey := newKey(t)
	subPK := newPublicKey(t, subKey)

	b := NewCertificate(rootPK, subPK, 100, 3600)
	b.DocumentTypes = []string{FactType + "#email"}

	if b.AllowedType(NewFact("phone")) {
		t.Fatal("Certificate should not have allowed the fact document subtype")
	}
}
