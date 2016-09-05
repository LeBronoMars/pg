package handlers

import (
	"net/http"
	"time"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "philgheps/survey/api/models"
)

type SurveyHandler struct {
	db *gorm.DB
}

func NewSurveyHandler(db *gorm.DB) *SurveyHandler {
	return &SurveyHandler{db}
}

//get all news
func (handler SurveyHandler) Index(c *gin.Context) {
	survey := []m.Survey{}	
	var query = handler.db

	startParam,startParamExist := c.GetQuery("start")
	limitParam,limitParamExist := c.GetQuery("limit")
	statusParam,statusParamExist := c.GetQuery("status")
	whenParam,whenParamExist := c.GetQuery("when")

	//start param exist
	if startParamExist {
		start,_ := strconv.Atoi(startParam)
		if start != 0 {
			query = query.Offset(start)				
		}
	} 

	//limit param exist
	if limitParamExist {
		limit,_ := strconv.Atoi(limitParam)
		query = query.Limit(limit)
	} else {
		query = query.Limit(10)
	}

	//status param exist
	if statusParamExist {
		query = query.Where("status = ?",statusParam)
	}

	//when param exist
	if whenParamExist {
		asia, _ := time.LoadLocation("Asia/Manila")
		now := time.Now().In(asia)				
		startOfDay := GetStartOfDay(now)
		if whenParam == "today" {
			endOfDay := GetEndOfDay(now)
			fmt.Printf("\nStart Of Day ----> %v",startOfDay)
			fmt.Printf("\nEnd of Day ----> %v",endOfDay)
			query = query.Where("created_at between ? AND ?",startOfDay,endOfDay)
		} else if whenParam == "previous" {
			query = query.Where("created_at < ?",startOfDay)
		}
	}
	query.Order("created_at desc").Find(&survey)
	c.JSON(http.StatusOK,survey)	
}

func (handler SurveyHandler) Create(c *gin.Context) {
	var survey m.Survey
	err := c.Bind(&survey)
	if err == nil {
		result := handler.db.Create(&survey)
		if result.RowsAffected > 0 {
			c.JSON(http.StatusCreated,survey)
		} else {
			respond(http.StatusBadRequest,result.Error.Error(),c,true)
		}
	} else {
		respond(http.StatusBadRequest,err.Error(),c,true)
	}
}


