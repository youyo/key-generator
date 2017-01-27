package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// request & response
type (
	Request struct {
		CommonName             string `json:"common_name"`
		Country                string `json:"country"`
		State                  string `json:"state"`
		Locality               string `json:"locality"`
		OrganizationName       string `json:"organization_name"`
		OrganizationalUnitName string `json:"organizational_unit_name"`
	}

	Response struct {
		Response []Key `json:"response"`
		Status   bool  `json:"status"`
		Error    error `json:"error"`
	}
	Key struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}
)

func main() {
	// initialilze
	e := echo.New()

	// cors
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error { return index(c) })

	// create csr & privatekey
	e.POST("/generate", func(c echo.Context) error { return generateKeys(c) })

	// server start
	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Status: true, Error: nil})
}

func NewRequest() (r *Request) {
	return &Request{}
}

func generateKeys(c echo.Context) (err error) {
	// bind json
	r := NewRequest()
	if err = c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Response: nil, Error: err})
	}

	// generate private key
	p, err := generatePrivateKeyBytes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Response: nil, Error: err})
	}
	privateKey, err := exportPrivateKey(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Response: nil, Error: err})
	}

	// generate csr
	csrBytes, err := r.generateCsrBytes(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Response: nil, Error: err})
	}
	csr, err := exportCsr(csrBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Response: nil, Error: err})
	}

	// return success response
	response := Response{
		Response: []Key{
			Key{
				Name: r.CommonName + ".key",
				Data: privateKey,
			},
			Key{
				Name: r.CommonName + ".csr",
				Data: csr,
			},
		},
		Status: true,
	}
	return c.JSON(http.StatusCreated, response)
}

func generatePrivateKeyBytes() (p *rsa.PrivateKey, err error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func exportPrivateKey(p *rsa.PrivateKey) (privateKey string, err error) {
	privateKey = string(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(p),
			},
		),
	)
	return privateKey, nil
}

func decodePrivateKey(p string) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(p))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func (r Request) generateCsrBytes(privateKey *rsa.PrivateKey) (csr []byte, err error) {
	subj := pkix.Name{
		CommonName:         r.CommonName,
		Country:            []string{r.Country},
		Province:           []string{r.State},
		Locality:           []string{r.Locality},
		Organization:       []string{r.OrganizationName},
		OrganizationalUnit: []string{r.OrganizationalUnitName},
	}
	asn1Subj, err := asn1.Marshal(subj.ToRDNSequence())
	if err != nil {
		return nil, err
	}
	template := x509.CertificateRequest{
		RawSubject:         asn1Subj,
		SignatureAlgorithm: x509.SHA256WithRSA,
	}
	return x509.CreateCertificateRequest(rand.Reader, &template, privateKey)
}

func exportCsr(csrBytes []byte) (csr string, err error) {
	csr = string(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "CERTIFICATE REQUEST",
				Bytes: csrBytes,
			},
		),
	)
	return csr, nil
}
