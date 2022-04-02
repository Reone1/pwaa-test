package router

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		// c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

		if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
		}

		c.Next()
	}
}

type bodyLogWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
    c.Writer = blw
    c.Next()
    // statusCode := c.Writer.Status()
        //ok this is an request with error, let's make a record for it
        // now print body (or log in your preferred way)
		path := c.Request.URL.Path
		if !strings.Contains(path, "swagger") {
			fmt.Println("Response body: " + blw.body.String())
		}
	}
}
func init() {
	if router == nil {
		router = gin.New()
	}
	router.Use( CORSMiddleware(), ginBodyLogMiddleware() )
}
func SetRouter() *gin.Engine{
	PwaaRouter()
	BottleRouter()
	UserRouter()
	return router
}