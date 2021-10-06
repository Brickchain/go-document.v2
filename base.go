package document

import (
	"fmt"
	"strings"
	"time"
)

const BaseType = SchemaLocation + "/base.json"

type Base struct {
	Type        string    `json:"@type"`
	Timestamp   time.Time `json:"@timestamp"`
	ID          string    `json:"@id,omitempty"`
	Certificate string    `json:"@certificate,omitempty"`
	Realm       string    `json:"@realm,omitempty"`
	raw         []byte
}

func NewBase() *Base {
	return &Base{
		Type:      BaseType,
		Timestamp: time.Now().UTC(),
	}
}

func (b *Base) GetTimestamp() time.Time {
	return b.Timestamp
}

func (b *Base) GetCertificate() string {
	return b.Certificate
}

func (b *Base) GetType() string {
	return b.Type
}

func (b *Base) GetRaw() []byte {
	return b.raw
}

func (b *Base) SetRaw(raw []byte) {
	b.raw = raw
}

func (b *Base) SubType() string {
	return strings.Split(b.Type, "#")[1]
}

func (b *Base) SetSubType(subType string) {
	b.Type = fmt.Sprintf("%s#%s", b.Type, subType)
}
