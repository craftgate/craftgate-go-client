package adapter

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Cryptogram struct {
	PublicKey []byte
}

// NewCryptogram конструктор
func NewCryptogram(path string) *Cryptogram {
	public, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &Cryptogram{
		PublicKey: public,
	}
}

// EncryptWithPublicKey encrypts data with public key
func (c *Cryptogram) EncryptWithPublicKey(msg []byte) (mes []byte, err error) {

	rawPem, err := ioutil.ReadFile("rsa_prod.rsa")

	pemBlock, _ := pem.Decode(rawPem)
	publicKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)

	/*block, _ := pem.Decode(c.PublicKey)
	  if block == nil {
	  	panic("failed to parse PEM block containing the public key")
	  }

	  pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	  if err != nil {
	  	panic("failed to parse DER encoded public key: " + err.Error())
	  }*/

	/*block, _ := pem.Decode(c.PublicKey)
	  if block == nil {
	  	fmt.Println("error 1")
	  	return nil,errors.New("private key error!")
	  }
	  pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	  if err != nil {
	  	fmt.Println("error 2")
	  	return nil, err
	  }*/
	mes, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), msg)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func adapter() {

	Cryptogram := NewCryptogram("rsa_prod.rsa")
	b, err := Cryptogram.EncryptWithPublicKey([]byte("{\"hpan\": \"ххххх\",\"cvc\":\"хххх\",\"expDate\":\"хххх\",\"terminalId\":\"хххххххх\"}"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := base64.StdEncoding.EncodeToString(b)
	fmt.Println(s)

}
