package pkg

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKEY = []byte(os.Getenv("secret_key"))

// Gerador de JWT
func GenerateJWT() (string, error) {

	// formato do JWT : Header.Payload.Signature
	var tokenString, err = jwt.NewWithClaims(
		jwt.SigningMethodHS256, // = Header

		jwt.MapClaims{ // = Payload
			"user_id": 3200,
			"iss":     "Daniel",
			"aud":     "trueAPI",
			"exp":     time.Now().Add(time.Hour * 2).Unix(), // expirar: tempo atual + 2 horas
		},
	).SignedString(secretKEY) // = Signature

	return tokenString, err
}

func ValidateJWT(tokenString string) bool {
	tokenString = _RemoveBearerPrefix(tokenString)

	// Analisa o Token
	var token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar se o método é do tipo criptográfico HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "Não é do tipo HMAC", nil
		}
		// utiliza a `secretKEY` para validação
		return secretKEY, nil
	})
	if err != nil {
		fmt.Println(err)
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("claims: ", claims)
		return true
	} else {
		fmt.Println("condição inválida")
		return false
	}
}

// Remover o prefixo `Bearer`
func _RemoveBearerPrefix(TokenString string) string {
	if strings.HasPrefix(TokenString, "Bearer ") {
		TokenString = strings.TrimPrefix(TokenString, "Bearer ")
	}

	return TokenString
}
