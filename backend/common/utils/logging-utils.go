package utils

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		if c.Request.Body != nil {
			body, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		log.Printf("Request Body: %s", body)
		c.Next()
	}
}
