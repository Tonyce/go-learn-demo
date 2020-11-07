package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger ...
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c := cc.Copy()
		// go func() {
		// simulate a long task with time.Sleep(). 5 seconds
		// time.Sleep(5 * time.Second)

		// // note that you are using the copied context "cCp", IMPORTANT
		// log.Println("Done! in path " + cCp.Request.URL.Path)

		t := time.Now()

		clientIP := c.ClientIP()
		if clientIP == "" {
			clientIP = "-"
		}
		method := c.Request.Method
		path := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery
		if rawQuery == "" {
			rawQuery = "-"
		}
		proto := c.Request.Proto
		ua := c.Request.UserAgent()

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// reqBodySize := c.Reader.Size()

		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		// log.Print(latency)

		// access the status we are sending
		statusCode := c.Writer.Status()
		// bodySize := c.Writer.Size()
		// c.Errors.ByType(ErrorTypePrivate).String()
		// log.Println(status)
		log.Println(fmt.Sprintf("%s %d >> %s %s %s %s %s %s << %s",
			clientIP,
			statusCode,
			// t.Format(time.RFC3339),
			ua,
			method,
			proto,
			path,
			rawQuery,
			string(bodyBytes),
			latency,

			// bodySize,
		))
		// }()
	}
}
