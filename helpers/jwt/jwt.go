package jwt

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v4"
)

// simple key
// var sampleSecretKey = []byte("SecretYouShouldHide")

func Sign(data TokenDetail) (Token, error) {
	sampleSecretKey := loadRsaPrivateKey()
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(sampleSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(3600 * time.Minute).Unix()
	claims["authorized"] = true
	claims["sub"] = data
	claims["data"] = data
	refresh_token, _ := createRefreshToken()
	// Generate encoded token and send it as response.
	if t, err := token.SignedString(privateKey); err != nil {
		return Token{}, err
	} else {
		return Token{
			AccessToken:      t,
			ExpiresIn:        claims["exp"],
			TokenType:        "Bearer",
			RefreshToken:     refresh_token.RefreshToken,
			RefreshExpiresIn: refresh_token.RefreshExpiresIn,
		}, nil
	}
}

func Decode(token string) (jwt.MapClaims, error) {
	sampleSecretKey := loadRsaPublicKey()
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(sampleSecretKey)
	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if err != nil {
			return publicKey, err
		}
		return publicKey, nil
	})
	if err != nil {
		return jwt.MapClaims{}, err
	}
	return claims["data"].(map[string]interface{}), nil
	// token_type := fmt.Sprintf("%s", claims["type"])
}

func createRefreshToken() (RefreshToken, error) {
	sampleSecretKey := loadRsaPrivateKey()

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	claims["authorized"] = true
	claims["typ"] = "refresh_token"
	// Generate encoded token and send it as response.
	if t, err := token.SignedString(sampleSecretKey); err != nil {
		return RefreshToken{}, err
	} else {
		return RefreshToken{
			RefreshExpiresIn: claims["exp"],
			RefreshToken:     t,
		}, nil
	}
}

func Accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}

func loadRsaPrivateKey() []byte {
	bytes, err := ioutil.ReadFile("mots.key")
	if err != nil {
		log.Println(err.Error())
	}
	// start := strings.Split(string(bytes), "-----BEGIN PRIVATE KEY-----")
	// content := strings.Split(start[1], "-----END PRIVATE KEY-----")
	// key, err := ssh.ParseRawPrivateKey(bytes)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	return bytes
}

func loadRsaPublicKey() []byte {
	bytes, err := ioutil.ReadFile("mots.pub")
	if err != nil {
		log.Println(err.Error())
	}
	// start := strings.Split(string(bytes), "-----BEGIN PUBLIC KEY-----")
	// content := strings.Split(start[1], "-----END PUBLIC KEY-----")
	return bytes
	// key, err := ssh.ParsePublicKey(bytes)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// return key.(*rsa.PublicKey)
}
