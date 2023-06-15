package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type employee struct {
   
    Name     string  `json:"name"`
    Leave_Type  string  `json:"leave_type"`
    Fromdate string  `json:"fromdate"`
    Todate string `json:"todate"`
	Team_Name string `json:"team_name"`
	File_upload string `json:"file_upload"`
	Reporter string `json:"reporter"`
}
var employees = []employee{
    { Name: "Raju", Leave_Type: "sick leave", Fromdate: "2023-12-4",Todate: "2023-12-5",Team_Name:"secops",File_upload:"https://circleci.com/blog/gin-gonic-testing/",Reporter:"Sandeep sir"},
  
}


func HomepageHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message":"Welcome to the Tech Company listing API with Golang"})
	c.JSON(employees);
}

func main() {
    router := gin.Default()
    router.GET("/", HomepageHandler)
    router.Run()
}
