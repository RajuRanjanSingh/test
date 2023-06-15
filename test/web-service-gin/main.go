package main

import (
	"log"
	"github.com/go-pg/pg/v9"
	"github.com/gin-gonic/gin"
	"net/http"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/gin-contrib/cors"
)
// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User: "postgres",
		Password: "raju",	
		Addr: "localhost:5432",
		Database: "emp",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
	}
	log.Printf("Connected to db")
	CreateLeaveTable(db)
	InitiateDB(db)
	return db
}
type LeaveFormTable struct {
	Name     string  `json:"name"`
    Leave_Type  string  `json:"leave_type"`
    Fromdate string  `json:"fromdate"`
    Todate string `json:"todate"`
	Team_Name string `json:"team_name"`
	File_upload string `json:"file_upload"`
	Reporter string `json:"reporter"`
}
// Create User Table
func CreateLeaveTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&LeaveFormTable{}, opts)
	if createError != nil {
		log.Printf("Error while creating Leave table, Reason")
		return createError
	}
	log.Printf("Leave form table created")
	return nil
}
// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAllLeaveFormTable(c *gin.Context){
	var LeaveFormTables []LeaveFormTable
	err := dbConnect.Model(&LeaveFormTables).Select()
    if err != nil {
		log.Printf("Error while getting all leave form")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All LeaveFormTables",
		"data": LeaveFormTables,
	})
	return
}
func CreateLeave(c *gin.Context) {
	name := c.PostForm("Name")
	leave_Type := c.PostForm("Leave_Type")	
	fromdate :=c.PostForm("fromdate")
	todate:= c.PostForm("Todate")
	team_Name := c.PostForm("Team_Name")
	file_upload := c.PostForm("File_upload")
	reporter:=c.PostForm("Reporter")

	insertError := dbConnect.Insert(&LeaveFormTable{
		Name: name,
		Leave_Type : leave_Type,
		Fromdate: fromdate,
		Todate: todate,
		Team_Name: team_Name,
		File_upload: file_upload,
		Reporter:reporter,
	})
	
	if insertError != nil {
		log.Printf("Error while inserting new leave form into db")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "leaveformtable created Successfully",
		
	})
	return
}
func Routes(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500"} // Update with your front-end URL
	router.Use(cors.New(config))
	router.GET("/", welcome)
	router.GET("/LeaveFormTable", GetAllLeaveFormTable)
	router.POST("/LeaveFormTables",CreateLeave)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func main() {
	
	// Connect DB
	Connect()
	// Init Router
	router := gin.Default()
	// Route Handlers / Endpoints
	router.LoadHTMLGlob("Template/*")

	Routes(router)
	log.Fatal(router.Run(":4747"))
}


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}