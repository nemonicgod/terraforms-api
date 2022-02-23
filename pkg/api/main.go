package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/kubastick/ginception"
	"github.com/nemonicgod/terraforms-api/backend"
	"github.com/nemonicgod/terraforms-api/config"

	ginlogrus "github.com/toorop/gin-logrus"

	v1Parcels "github.com/nemonicgod/terraforms-api/controllers/v1/parcels"
)

func BackendMiddleware(b *backend.Backend) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("backend", b)
		c.Next()
	}
}

// CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ignoreLogger(p string) gin.HandlerFunc {
	return logger.SetLogger(
		logger.WithSkipPath([]string{p}),
	)
}

func main() {
	role := os.Getenv("ROLE")

	if role != config.ROLE_API {
		panic(fmt.Errorf("JOB: exited, incorrect role passed %s", role))
	}

	// Initalize a new client, the base entrpy point to the application code
	// the true, true is a poor way of turning DB/Redis on/off, cmon bro
	b, e := backend.NewBackend(true, true)
	if e != nil {
		panic(e)
	}

	// Database connect, defer close
	pqDB, err := b.R.D.DB()
	if err != nil {
		panic(err)
	}
	defer pqDB.Close()

	b.L.Println("API: main() starting...")

	// Enable and/or set cors
	cf := cors.DefaultConfig()
	cf.AllowAllOrigins = true
	cf.AllowCredentials = true
	cf.AddAllowHeaders("authorization")

	// Set the initial API instance
	r := gin.New()

	r.Use(cors.New(cf))
	r.Use(cors.Default())
	r.Use(CORSMiddleware())
	r.Use(ginception.Middleware())

	// Set the backend
	r.Use(BackendMiddleware(b))

	r.GET("/", ignoreLogger("/"), func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("OK"))
	})

	r.GET("/health", ignoreLogger("/health"), func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("OK"))
	})

	// parcels
	r.GET("/v1/parcels/:id", ginlogrus.Logger(b.L), v1Parcels.GetParcel)

	//curve.GetDailyTimestampsBlockNums()
	r.Run((":" + b.C.GetString("port")))
}
