package main

import (
	"github.com/jaloldinov/books-catalog/api/docs"

	"fmt"

	"github.com/jaloldinov/books-catalog/config"
	"github.com/jaloldinov/books-catalog/handler"
	"github.com/jaloldinov/books-catalog/storage/postgres"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/gin-gonic/gin"
)

// @title           Book Store
// @version         0.0.1
// @description     Book Store beta application

// @contact.name   API Support
// @contact.url    http://t.me/jaloldinovs
// @contact.email  jaloldinovuz@gmail.com
// @license.name   MIT
// @BasePath  /api/v1
func main() {
	cfg := config.Load()

	docs.SwaggerInfo.Host = fmt.Sprintf("%v%v", cfg.ServiceHost, cfg.HTTPPort)

	str := fmt.Sprintf("port=%d host=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PostgresPort, cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresDatabase, cfg.PostgresPassword, cfg.PostgresSSLMode,
	)

	strg := postgres.NewPostgres(str)
	defer strg.CloseDB()

	handler := handler.NewHandler(strg)

	switch cfg.Environment {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		authors := v1.Group("/authors")
		{
			authors.POST("/", handler.CreateAuthor)
			authors.GET("/", handler.GetAllAuthors)
			authors.GET("/:id", handler.GetAuthor)
			authors.PUT("/:id", handler.UpdateAuthor)
			authors.DELETE("/:id", handler.DeleteAuthor)
		}

		book_category := v1.Group("/book_category")
		{
			book_category.POST("/", handler.CreateBookCategory)
			book_category.GET("/", handler.GetAllBookCategories)
			book_category.GET("/:id", handler.GetBookCategory)
			book_category.PUT("/:id", handler.UpdateBookCategory)
			book_category.DELETE("/:id", handler.DeleteBookCategory)
		}

		books := v1.Group("/books")
		{
			books.POST("/", handler.CreateBook)
			books.GET("/", handler.GetAllBooks)
			books.GET("/:id", handler.GetBook)
			books.PUT("/:id", handler.UpdateBook)
			books.DELETE("/:id", handler.DeleteBook)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
