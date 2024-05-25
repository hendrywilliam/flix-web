package routes

import (
	"errors"
	schema "flix/db"
	"flix/domain"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesCore(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// v1
	v1 := r.Group("/api/v1")
	NewMovieRoute(v1)
	NewSignupRoute(v1, db)
	v1.POST("/match", func(c *gin.Context) {
		var request domain.SignupRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user schema.User
		db.Where("email = ?", request.Email).Find(&user)

		match, err := argon2id.ComparePasswordAndHash(request.Password, user.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if match {
			c.JSON(http.StatusOK, gin.H{"status": "password does match."})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("password does not match")})
	})
	return r
}
