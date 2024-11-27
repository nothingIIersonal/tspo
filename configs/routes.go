package configs

import (
	"pr8_1/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	userRepository := repositories.NewUserRepository(db)
	SetupUsersRoutes(userRepository, route)

	goodRepository := repositories.NewGoodRepository(db)
	SetupGoodsRoutes(goodRepository, route)

	roleRepository := repositories.NewRoleRepository(db)
	SetupRolesRoutes(roleRepository, route)

	featureRepository := repositories.NewFeatureRepository(db)
	SetupFeaturesRoutes(featureRepository, route)

	employeeRepository := repositories.NewEmployeeRepository(db)
	SetupEmployeesRoutes(employeeRepository, route)

	vendorRepository := repositories.NewVendorRepository(db)
	SetupVendorsRoutes(vendorRepository, route)

	orderRepository := repositories.NewOrderRepository(db)
	SetupOrdersRoutes(orderRepository, route)

	return route
}
