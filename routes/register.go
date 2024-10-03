package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest-api.com/models"
)

func registerForEvents(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})

		return
	}

	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this event."})

		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event."})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered for event."})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
	}

	var event models.Event

	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel register for event."})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled !!!"})

}
