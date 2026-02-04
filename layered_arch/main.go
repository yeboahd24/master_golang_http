package main

func main() {
	db := connectDB()

	// Wire dependencies
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	// Setup routes
	r := gin.Default()
	r.POST("/users", userHandler.Register)
	r.Run()
}
