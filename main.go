package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", Home)
	router.GET("/welcome", Welcome)
	router.GET("/getStudent/:rollNum", GetStudent)
	router.POST("addStudent", AddStudent)
	router.POST("addMultipleStudent", AddMultipleStudent)
	router.DELETE("/deleteStudent/:rollNum", DeleteStudent)
	router.Run(PORT)

}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to GIN API",
	})
}

func Welcome(c *gin.Context) { // /welcome?firstname=Jane&lastname=Doe
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func GetStudent(c *gin.Context) {
	rollNum, _ := strconv.Atoi(c.Param("rollNum"))
	// b, _ := json.Marshal((students[rollNum]))
	std, found := students[rollNum]
	if !found {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Student Details": std,
	})
}

func AddStudent(c *gin.Context) {
	rollNum, firstName, lastName, class, marks, contact := c.Query("rollNum"), c.Query("fName"), c.Query("lName"), c.Query("class"), c.Query("marks"), c.Query("contact")

	jsonDataBytes, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(jsonDataBytes))

	if len(rollNum) == 0 || len(firstName) == 0 || len(lastName) == 0 || len(class) == 0 || len(marks) == 0 || len(contact) == 0 {
		c.JSON(http.StatusPartialContent, "Partial Content")
		return
	}
	rollNumInt, e1 := strconv.Atoi(rollNum)
	classInt, e2 := strconv.Atoi(class)
	marksFloat, e3 := strconv.ParseFloat(marks, 64)
	contactInt, e4 := strconv.Atoi(contact)
	_, found := students[rollNumInt]
	if found {
		c.JSON(http.StatusConflict, "Roll Number already in use")
		return
	}

	if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
		fmt.Println("Error:\n", e1, "\n", e2, "\n", e3, "\n", e4)
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	newStudent := student{
		RollNum:   rollNumInt,
		FirstName: firstName,
		LastName:  lastName,
		Class:     classInt,
		Marks:     marksFloat,
		Contact:   contactInt,
	}
	students[rollNumInt] = newStudent

	c.JSON(http.StatusOK, gin.H{
		"Added": newStudent,
	})
}

func AddMultipleStudent(c *gin.Context) {
	var studentsList multipleStudent
	err := c.BindJSON(&studentsList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Student Details": studentsList,
	})
}

func DeleteStudent(c *gin.Context) {
	rollNum, _ := strconv.Atoi(c.Param("rollNum"))
	_, found := students[rollNum]
	if !found {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}
	delete(students, rollNum)
	c.JSON(http.StatusOK, gin.H{
		"Status": "Deleted",
	})

}
