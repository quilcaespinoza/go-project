package commons

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang-es/prjecomments/models"
)

var (
	privateKey *rsa.PrivateKey
	// PublicKey para validar el token
	PublicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal("no se pudo leer privado")
	}
	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal("no se pudo leer publico")
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pido hacer el parse privado")
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pido hacer el parse public")
	}

}

//GenerateJWT genera el token de autenticación
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
		// Issuer: "ProgrammingJE",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo general el token")
	}
	return result
}
