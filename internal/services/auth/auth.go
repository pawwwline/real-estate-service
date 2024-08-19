package auth

import(
	"github.com/dgrijalva/jwt-go"
	"real-estate-service/api/generated"
)

var TokenSign = []byte("b99f6a9b74321b8b4f4c73e3de004ad7a3bd78f3482e93c8f4a596a6b09f2208c") //TODO:move to env variables


func CreateToken(userType generated.UserType) (string, error){
	claims := jwt.MapClaims{
		"user_type": userType,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(TokenSign)
	if err != nil {
		return "", err
	}
	
	return signedToken, nil
}

