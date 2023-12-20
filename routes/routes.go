package routes

import (
	"log"
	middlewares "match_card/middleware"
	auth_response "match_card/pkg/auth/http"
	filter_response "match_card/pkg/filter/http"
	upload_utils "match_card/utils/upload"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	// Note: Rate Limiter
	// * 5 reqs/second: "5-S"
	// * 10 reqs/minute: "10-M"
	// * 1000 reqs/hour: "1000-H"
	// * 2000 reqs/day: "2000-D"

	// limit to 1000 requests per second. if exceed, will return http 429 (too many req)
	rate, err := limiter.NewRateFromFormatted("1000-S")
	if err != nil {
		log.Fatal(err)
		return route
	}

	store := memory.NewStore()

	// Create a new middleware with the limiter instance using in memory golang.
	middleware := mgin.NewMiddleware(limiter.New(store, rate))

	// Forward / Save Client ip to go memory.
	route.ForwardedByClientIP = true

	// Use Middleware rate limiter
	route.Use(middleware)

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE", "GET", "OPTIONS", "TRACE", "CONNECT"},
		AllowHeaders:     []string{"Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Origin", "Content-Type", "Content-Length", "Date", "origin", "Origins", "x-requested-with", "access-control-allow-methods", "access-control-allow-credentials", "apikey"},
		ExposeHeaders:    []string{"Content-Length"},
	}))

	//swagger only in dev
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := route.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	score := protected.Group("score")
	{
		score.POST("post", auth_response.CreateHandler)
		route.GET("score/get", auth_response.GetAllUser)
	}

	filter := protected.Group("filter")
	{
		filter.POST("post", filter_response.CreateHandler)
		route.GET("filter/get", filter_response.GetAllUser)
	}

	route.Static("/files", "./public")

	upload := protected.Group("/upload")
	{
		upload.POST("/file", upload_utils.UploadFieldHandler)
	}
	return route
}
