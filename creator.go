package golang

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	. "github.com/davidkhala/goutils"
)

type Creator struct {
	Msp            string
	CertificatePem []byte
	Certificate    x509.Certificate `json:"-"`
}

func ParseCreator(creator []byte) (Creator) {
	var msp bytes.Buffer

	var certificateBuffer bytes.Buffer
	var mspReady = false

	for i := 0; i < len(creator); i++ {
		char := creator[i]
		if char < 127 && char > 31 {
			if !mspReady {
				msp.WriteByte(char)
			} else {
				certificateBuffer.WriteByte(char)
			}
		} else if char == 10 {
			if mspReady {
				certificateBuffer.WriteByte(char)
			}
		} else {
			if msp.Len() > 0 {
				mspReady = true
			}

		}
	}

	certBytes := certificateBuffer.Bytes()
	block, rest := pem.Decode(certBytes)
	AssertEmpty(rest, "pem decode failed:"+string(rest))
	certificate, err := x509.ParseCertificate(block.Bytes)
	PanicError(err)

	return Creator{Msp: msp.String(), CertificatePem: certBytes, Certificate: *certificate}

}
