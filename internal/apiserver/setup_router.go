package apiserver

import (
	"log"
	"logical-example/internal/apiserver/controller"
	"logical-example/internal/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	// gin.SetMode()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// fmt.Println("setup_router init")
}

// SetupRouter ...
func SetupRouter() *gin.Engine {
	// r := gin.Default()
	r := gin.New()
	r.Use(middleware.Logger())
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))

	controller.PingRouter(r)
	controller.V1Router(r)

	return r
}
