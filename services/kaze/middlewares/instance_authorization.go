package middlewares

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/gasper/lib/mongo"
	"github.com/sdslabs/gasper/lib/utils"
	"github.com/sdslabs/gasper/types"
)

// VerifyAdmin allows the request to proceed only if the user has admin privileges
func VerifyAdmin(c *gin.Context) {
	claims := ExtractClaims(c)
	if claims == nil {
		utils.SendServerErrorResponse(c, errors.New("Failed to extract JWT claims"))
		return
	}
	if claims.IsAdmin {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(401, gin.H{
		"success": false,
		"error":   "User does not have admin privileges",
	})
}

func isInstanceOwner(c *gin.Context, instanceType string) {
	instance := c.Param(instanceType)
	claims := ExtractClaims(c)
	if claims == nil {
		utils.SendServerErrorResponse(c, errors.New("Failed to extract JWT claims"))
		return
	}
	if claims.IsAdmin {
		c.Next()
		return
	}
	count, err := mongo.CountInstances(types.M{
		"name":                instance,
		mongo.InstanceTypeKey: instanceType,
		"owner":               claims.Email,
	})
	if err != nil {
		utils.SendServerErrorResponse(c, err)
		return
	}

	if count == 0 {
		c.AbortWithStatusJSON(401, gin.H{
			"success": false,
			"error":   fmt.Sprintf("User %s is not entitled to perform operations on %s %s", claims.Email, instanceType, instance),
		})
		return
	}
	c.Next()
}

// IsAppOwner checks if a user is entitled to perform operations on an application
func IsAppOwner(c *gin.Context) {
	isInstanceOwner(c, mongo.AppInstance)
}

// IsDatabaseOwner checks if a user is entitled to perform operations on a database
func IsDatabaseOwner(c *gin.Context) {
	isInstanceOwner(c, mongo.DBInstance)
}