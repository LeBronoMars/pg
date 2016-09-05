package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func respond(statusCode int, responseMessage string, c *gin.Context, isError bool) {
	response := &Response{Message: responseMessage}
	c.JSON(statusCode,response)
}

type Response struct {
	Message string `json:"message"`
}

//change timezone of date
func changeTimeZone(t time.Time) time.Time {
	loc,_ := time.LoadLocation("Asia/Manila")
	newTime,_ := time.ParseInLocation(time.RFC3339,t.Format(time.RFC3339),loc)
	return newTime
}

func GetStartOfDay(t time.Time) time.Time {
    year, month, day := t.Date()
    return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func GetEndOfDay(t time.Time) time.Time {
    year, month, day := t.Date()
    return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}



