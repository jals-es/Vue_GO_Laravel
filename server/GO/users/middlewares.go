package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/satori/go.uuid"
	"appbar/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"fmt"
)

func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, my_user_id uuid.UUID) {
	var myUserModel UserModel
	// if my_user_id != uuid.NewV4 {
		
	// }

	db := common.GetDB()
	/*err :=*/ db.First(&myUserModel, my_user_id)
	
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
	// }else{
		
	// }
	c.Set("my_user_id", my_user_id)
	c.Set("my_user_model", myUserModel)
}

// You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//  r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// UpdateContextUserModel(c, uuid.UUID)
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.NBSecretPassword))
			return b, nil
		})
		if err != nil {
			if auto401 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			get_my_id := uuid.Must(uuid.FromString(fmt.Sprintf("%v", claims["id"])))

			UpdateContextUserModel(c, get_my_id)
		}
	}
}