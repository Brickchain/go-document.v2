package document

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
	jose "gopkg.in/square/go-jose.v1"
)

func TestCertificateSchema(t *testing.T) {
	issuer := newKey(t)
	subject := newKey(t)

	testSchema(t, CertificateType, NewCertificate(issuer, subject, 1, 30))
}

func newKey(t *testing.T) *jose.JsonWebKey {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	jwk := jose.JsonWebKey{
		Key:       key,
		Algorithm: string(jose.ES256),
	}

	return &jwk
}

func newPublicKey(t *testing.T, private *jose.JsonWebKey) *jose.JsonWebKey {
	t.Helper()

	publicKey := &private.Key.(*ecdsa.PrivateKey).PublicKey
	return &jose.JsonWebKey{
		Key:       publicKey,
		KeyID:     private.KeyID,
		Algorithm: private.Algorithm,
	}
}
