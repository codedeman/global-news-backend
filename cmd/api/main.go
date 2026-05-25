package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/codedeman/neura-backend/internal/agentpermission"
	"github.com/codedeman/neura-backend/internal/agentprofile"
	"github.com/codedeman/neura-backend/internal/article"
	"github.com/codedeman/neura-backend/internal/db"
	"github.com/codedeman/neura-backend/internal/googlenews"
	"github.com/codedeman/neura-backend/internal/interestcategory"
	"github.com/codedeman/neura-backend/internal/johari"
	"github.com/codedeman/neura-backend/internal/rawevent"
	"github.com/codedeman/neura-backend/internal/rss"
	"github.com/codedeman/neura-backend/internal/userbehavior"
	"github.com/codedeman/neura-backend/internal/user"
	"github.com/codedeman/neura-backend/internal/userinsight"
	"github.com/codedeman/neura-backend/internal/userinterest"
	"github.com/codedeman/neura-backend/internal/usermemory"
	"github.com/codedeman/neura-backend/internal/userprofile"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "root:02111997@tcp(localhost:3306)/ai-agent-postgres?parseTime=true&charset=utf8mb4"
	}

	mysql, err := db.NewMySQLDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer mysql.Close()

	go rss.StartWorker()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Global News Backend Running 🚀"})
	})
	router.HEAD("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// RSS news routes (in-memory)
	router.GET("/articles", article.GetArticlesHandler)
	router.GET("/google-news", googlenews.Handler)

	// DB-backed routes
	user.NewHandler(user.NewStore(mysql)).RegisterRoutes(router.Group("/users"))

	userprofile.NewHandler(userprofile.NewStore(mysql)).RegisterRoutes(router.Group("/user-profiles"))

	usermemory.NewHandler(usermemory.NewStore(mysql)).RegisterRoutes(router.Group("/user-memories"))

	userinterest.NewHandler(userinterest.NewStore(mysql)).RegisterRoutes(router.Group("/user-interests"))

	userinsight.NewHandler(userinsight.NewStore(mysql)).RegisterRoutes(router.Group("/user-insights"))

	userbehavior.NewHandler(userbehavior.NewStore(mysql)).RegisterRoutes(router.Group("/user-signals"))

	rawevent.NewHandler(rawevent.NewStore(mysql)).RegisterRoutes(router.Group("/raw-events"))

	interestcategory.NewHandler(interestcategory.NewStore(mysql)).RegisterRoutes(router.Group("/interest-categories"))

	agentprofile.NewHandler(agentprofile.NewStore(mysql)).RegisterRoutes(router.Group("/agent-profiles"))

	agentpermission.NewHandler(agentpermission.NewStore(mysql)).RegisterRoutes(router.Group("/agent-permissions"))

	johari.NewHandler(johari.NewStore(mysql)).RegisterRoutes(router.Group("/johari"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on :%s", port)
	router.Run(":" + port)
}
