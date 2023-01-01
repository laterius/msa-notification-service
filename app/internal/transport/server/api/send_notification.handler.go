package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

// SendNotificationHandler handles request to make payment
func SendNotificationHandler(service service.Service) func(c *gin.Context) {
	// Request body structure
	type Body struct {
		OrderID uuid.UUID `json:"order_id"`
		Status  bool      `json:"status"`
	}

	return func(c *gin.Context) {
		body := Body{}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		if err := service.SaveStatus(body.OrderID, body.Status); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Status sent success",
			"data":    gin.H{},
		})
	}
}
