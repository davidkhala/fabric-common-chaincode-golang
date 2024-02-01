package cid

import (
	"github.com/davidkhala/goutils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSyntax(t *testing.T) {
	t.Run("cid to json", func(t *testing.T) {
		var id_nil = ClientIdentity{
			MspID:          "msp",
			CertificatePem: "cert",
			Attrs:          nil,
		}
		assert.False(t, strings.Contains(string(goutils.ToJson(id_nil)), "Attrs"))
		id_nil.Attrs = map[string]string{}
		assert.False(t, strings.Contains(string(goutils.ToJson(id_nil)), "Attrs"))
		id_nil.Attrs["foo"] = "bar"
		assert.True(t, strings.Contains(string(goutils.ToJson(id_nil)), "Attrs"))
	})
}
