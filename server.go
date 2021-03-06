package main

import (
	"os"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	m "philgheps/survey/api/models"
	h "philgheps/survey/api/handlers"
	"philgheps/survey/api/config"
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	db := *InitDB()
	router := gin.Default()
	LoadAPIRoutes(router, &db)	
}

func LoadAPIRoutes(r *gin.Engine, db *gorm.DB) {
	public := r.Group("/api/v1")

	//manage users
	surveyHandler := h.NewSurveyHandler(db)
	public.GET("/surveys", surveyHandler.Index)
	public.POST("/survey", surveyHandler.Create)

	r.Run(fmt.Sprintf(":%s", "7000"))
}

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		    return []byte(config.GetString("TOKEN_KEY")), nil
		})

		if err != nil || !token.Valid {
		    c.AbortWithError(401, err)
		} 
	}
}

func InitDB() *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetString("DB_USER"), config.GetString("DB_PASS"),
		config.GetString("DB_HOST"), config.GetString("DB_PORT"),
		config.GetString("DB_NAME"))
	log.Printf("\nDatabase URL: %s\n", dbURL)

	_db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database:  %s", err))
	}
	_db.DB()
	_db.LogMode(true)
	_db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&m.Survey{})
	return &_db
}

func GetPort() string {
    var port = os.Getenv("PORT")
    // Set a default port if there is nothing in the environment
    if port == "" {
        port = "9000"
        fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
    }
    fmt.Println("port -----> ", port)
    return ":" + port
}