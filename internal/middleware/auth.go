package auth

// type AuthHeader struct {
// 	IDToken string `header:	"Authorization"`
// }

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		//
// 		//
// 		//
// 		if !isAuthorized(c) {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		c.Next()
// 	}
// }

// func isAuthorized(c *gin.Context) bool {
// 	//
// 	//
// 	//
// 	authorizationHeader := c.GetHeader("Authorization")
// 	return authorizationHeader == "my-secret-token"
// }
