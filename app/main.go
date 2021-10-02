package main

import (
	"log"

	_driverFactory "go-watchlist/drivers"

	_adminService "go-watchlist/business/admins"
	_adminController "go-watchlist/controllers/admins"
	_adminRepo "go-watchlist/drivers/databases/admins"

	_userService "go-watchlist/business/users"
	_userController "go-watchlist/controllers/users"
	_userRepo "go-watchlist/drivers/databases/users"

	_dbDriver "go-watchlist/drivers/mysql"

	_middleware "go-watchlist/app/middlewares"
	_routes "go-watchlist/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_adminRepo.Admins{},
		&_userRepo.Users{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.InitDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: int64(viper.GetInt(`jwt.expired`)),
	}

	e := echo.New()

	adminRepo := _driverFactory.NewAdminRepository(db)
	adminService := _adminService.NewAdminService(adminRepo, 10, &configJWT)
	adminCtrl := _adminController.NewAdminController(adminService)

	userRepo := _driverFactory.NewUserRepository(db)
	userService := _userService.NewUserService(userRepo, 10, &configJWT)
	userCtrl := _userController.NewUserController(userService)

	routesInit := _routes.ControllerList{
		AdminController: *adminCtrl,
		UserController:  *userCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
