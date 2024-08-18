package quiz

import (
	"net/http"

	"github.com/YasserRABIE/QUIZFYv2/migrations/quiz_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	userID, _ := c.Get("user_id")

	quizzes, err := quiz_migrations.GetAll(userID.(uint))
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	r := response.NewSuccess(quizzes)
	c.JSON(http.StatusOK, r)
}
