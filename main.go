package main

import (
	"book-rental-app/config"
	"book-rental-app/handler"
	"book-rental-app/middleware"
	"book-rental-app/model"
	"book-rental-app/repository"
	"book-rental-app/service"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No env file found\n")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Auto migrate (DEV ONLY)
	// db.Migrator().DropTable(
	// 	&model.Loan{},
	// 	&model.Book{},
	// 	&model.Author{},
	// 	&model.Genre{},
	// 	&model.User{},
	// )

	db.AutoMigrate(
		&model.Author{},
		&model.Genre{},
		&model.Book{},
		&model.Loan{},
		&model.User{},
	)

	//repository
	userRepository := repository.NewUserRepository(db)
	authorRepository := repository.NewAuthorRepository(db)
	genreRepository := repository.NewGenreRepository(db)
	bookRepository := repository.NewBookRepository(db)
	loanRepository := repository.NewLoanRepository(db)

	//service
	userService := service.NewUserService(userRepository)
	authorService := service.NewAuthorService(authorRepository)
	genreService := service.NewGenreService(genreRepository)
	bookService := service.NewBookService(bookRepository)
	loanService := service.NewLoanService(loanRepository)

	//handler
	userHandler := handler.NewUserHandler(userService)
	authorHandler := handler.NewAuthorHandler(authorService)
	genreHandler := handler.NewGenreHandler(genreService)
	bookHandler := handler.NewBookHandler(bookService)
	loanHandler := handler.NewLoanHandler(loanService)

	e := echo.New()

	admin := e.Group("/admin")
	admin.GET("/authors", authorHandler.GetAllAuthor)
	admin.POST("/authors", authorHandler.CreateAuthor)

	admin.GET("/genres", genreHandler.GetAllGenre)
	admin.POST("/genres", genreHandler.CreateGenre)

	admin.GET("/users", userHandler.GetAllUser)

	e.POST("/users/register", userHandler.CreateUser)
	e.POST("/users/login", userHandler.Login)

	// JWT protected routes
	jwtSecret := os.Getenv("JWT_SECRET")
	auth := e.Group("")
	auth.Use(middleware.JWTMiddleware(jwtSecret))

	auth.GET("/users/me", userHandler.GetUserByEmail)
	auth.GET("/books", bookHandler.GetBook)
	auth.POST("/loans", loanHandler.CreateLoan)

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
