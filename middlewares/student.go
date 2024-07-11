package middlewares

import (
	"net/http"

	"github.com/YasinOkat/go-school-api/repositories"
	"github.com/gin-gonic/gin"
)

func StudentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDRaw, userIDExists := c.Get("user_id")
		userTypeIDRaw, userTypeIDExists := c.Get("user_type_id")
		if !userIDExists || !userTypeIDExists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		userTypeID, _ := userTypeIDRaw.(float64)
		userID, _ := userIDRaw.(float64)

		if (int(userTypeID) == 2) || (int(userTypeID) == 3) {
			c.Next()
			return
		}

		if int(userTypeID) == 1 {
			studentIDDB, err := repositories.GetStudentIDByUserID(int(userID))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			c.Set("studentIDDB", studentIDDB)
			c.Next()
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}

		c.Abort()
	}
}
