package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/Shahboz4131/api-gateway-bot/api/handlers/models"
	pb "github.com/Shahboz4131/api-gateway-bot/genproto"
	l "github.com/Shahboz4131/api-gateway-bot/pkg/logger"
)

// GetMessage ...
// @Router /v1/ [get]
// @Summary GetMessage
// @Description This API for getting message
// @Tags message
// @Accept  json
// @Produce  json
// @Param Task request body models.Message true "messageGetRequest"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) GetMessage(c *gin.Context) {
	var (
		body         pb.Message
		jsonbMarshal protojson.MarshalOptions
	)
	jsonbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	response, err := h.serviceManager.BotService().GetMessage(context.Background(), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to send message", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}
