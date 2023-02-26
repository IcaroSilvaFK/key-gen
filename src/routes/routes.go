package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/IcaroSilvaFK/key-gen/src/internals/lib"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PasswordInput struct {
	Pass string `json:"pass"`
}

func ApplicationRoutes(router *gin.RouterGroup) {

	router.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"message": "pong"})
	})

	router.POST("/hash", func(ctx *gin.Context) {

		var input PasswordInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
			return
		}

		hash, err := lib.MakeHash(input.Pass)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"pass": hash,
		})

	})

	router.GET("/hash", func(ctx *gin.Context) {

		limitStringInput, _ := ctx.GetQuery("limit")

		fmt.Println(limitStringInput)

		limit, err := strconv.Atoi(limitStringInput)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
				"err":     err.Error(),
			})
			return
		}

		var arrHash []string

		for i := 0; i < limit; i++ {
			uu := uuid.NewString()
			hash, err := lib.MakeHash(uu)

			if err != nil {
				log.Fatal(err)
				return
			}
			arrHash = append(arrHash, hash)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"passwords": arrHash,
		})
	})

}
