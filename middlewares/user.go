package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
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

		requestedUserID, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}

		if (int(userTypeID) == 2) || (int(userTypeID) == 3) {
			c.Next()
			return
		}

		if int(userTypeID) == 1 && int(userID) == requestedUserID {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()
	}
}
