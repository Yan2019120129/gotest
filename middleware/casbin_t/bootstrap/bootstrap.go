package bootstrap

import (
	"errors"
	"os"

	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/fiber/v2"

	"gotest/common/module/gorm/database"
	"gotest/middleware/casbin_t/api"
	"gotest/middleware/casbin_t/models"
	"gotest/middleware/casbin_t/router"
	"gotest/middleware/casbin_t/service"
)

const modelText = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)`

func NewApp() (*fiber.App, error) {
	if database.DB == nil {
		return nil, errors.New("GORM database is not initialized; check common/config/config.yml")
	}
	if err := database.DB.AutoMigrate(&models.User{}, &models.Test{}, &models.CasbinRule{}); err != nil {
		return nil, err
	}

	adapter, err := gormadapter.NewAdapterByDB(database.DB)
	if err != nil {
		return nil, err
	}
	model, err := casbinmodel.NewModelFromString(modelText)
	if err != nil {
		return nil, err
	}
	enforcer, err := casbin.NewEnforcer(model, adapter)
	if err != nil {
		return nil, err
	}

	secret := os.Getenv("CASBIN_T_JWT_SECRET")
	if secret == "" {
		secret = "change-this-development-secret"
	}
	auth := service.NewAuthService(database.DB, secret)
	handler := api.NewHandler(auth, service.NewTestService(database.DB))
	app := fiber.New(fiber.Config{AppName: "casbin-jwt-fiber-demo"})
	router.Register(app, handler, auth, enforcer)
	return app, nil
}
