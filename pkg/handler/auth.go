package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/model"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var body model.SignUpRequest

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	h.service.Auth.SignUp(&body)

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "ok",
	})
}

func (h *Handler) signIn(ctx *gin.Context) {
	var body model.SignInRequest

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	user, err := h.service.Auth.SignIn(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Username or password is wrong",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("token", tokenString, 3600*24*7, "", "", false, true)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func (h *Handler) signOut(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "", "", false, true)
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (h *Handler) authMiddleware(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := h.service.User.FindById(int(claims["sub"].(float64)))

		if user.ID == 0 || err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
