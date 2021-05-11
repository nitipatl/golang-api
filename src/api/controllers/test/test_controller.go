package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestResult struct {
	Status string `json:"status"`
}

func GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, TestResult{
		Status: "OK!",
	})
}
