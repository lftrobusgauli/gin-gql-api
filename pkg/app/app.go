package app

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"thegame/controller"
	"thegame/graph"
	"thegame/middleware"
	"thegame/model"
)

type App struct {
	Storage *gorm.DB
}

// Run is the core entry function for this service which blocks forever and
// is meant to be called by the main function.
func (app *App) Run() error {
	// Setting up Gin
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/query", graphqlHandler(app.Storage))
	r.GET("/", playgroundHandler())
	r.Run()
	return nil
}

// Defining the Graphql handler
func graphqlHandler(database *gorm.DB) gin.HandlerFunc {
	c := graph.Config{
		Resolvers: &controller.Resolver{
			UserService: &model.User{},
		},
	}
	h := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	return middleware.CreateDbContext(database, h)
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
