package router

import (
	swaggerFiles "github.com/swaggo/files"
	"vue3_admin/controller"
	"vue3_admin/logger"
	"vue3_admin/middlewares"
	"vue3_admin/settings"

	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "vue3_admin/docs"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()

	// 图片路由设置
	r.MaxMultipartMemory = 4 << 20 // 4 MiB
	r.Static("/static", settings.Conf.Static.Path)

	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.Cors())

	// swag route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DocExpansion("none")))

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "I'm OK!")
	})
	adminAclGroup := r.Group("/admin/acl")
	adminAclGroup.POST("/index/login", controller.UserController.Login)
	adminAclGroup.POST("/index/logout", controller.UserController.Logout)
	adminAclGroup.Use(middlewares.JWTAuthMiddleware()) // 应用 JWT认证中间件
	{
		adminAclGroup.GET("/index/info", controller.UserController.GetInfo)

		// 用户管理
		adminAclGroup.POST("/user/save", controller.UserController.SignUp)
		adminAclGroup.GET("/user/:page/:limit", controller.UserController.GetUser)
		adminAclGroup.GET("/user/toAssign/:adminId", controller.UserController.ToAssign)
		adminAclGroup.POST("/user/doAssignRole", controller.UserController.DoAssignRole)

		// 角色管理
		adminAclGroup.GET("/role/:page/:limit", controller.RoleController.GetRole)
		adminAclGroup.POST("/role/save", controller.RoleController.SaveRole)
		adminAclGroup.PUT("/role/update", controller.RoleController.UpdateRole)
		adminAclGroup.DELETE("/role/remove/:id", controller.RoleController.DeleteRole)

		// 菜单管理
		adminAclGroup.GET("/permission", controller.MenuController.GetMenu)
		adminAclGroup.POST("/permission/save", controller.MenuController.SaveMenu)
		adminAclGroup.PUT("/permission/update", controller.MenuController.UpdateMenu)
		adminAclGroup.DELETE("/permission/remove/:id", controller.MenuController.DeleteMenu)
		adminAclGroup.GET("/permission/toAssign/:roleId", controller.MenuController.ToAssign)
		adminAclGroup.POST("/permission/doAssign", controller.MenuController.DoAssign)
	}

	adminProductGroup := r.Group("/admin/product")
	adminProductGroup.Use(middlewares.JWTAuthMiddleware())
	{
		// 文件上传
		adminProductGroup.POST("/fileUpload", controller.FileController.FileUpload)

		// 品牌管理
		adminProductGroup.POST("/baseTrademark/save", controller.TrademarkController.CreateTrademark)
		adminProductGroup.GET("/baseTrademark/:page/:limit", controller.TrademarkController.GetTrademark)
		adminProductGroup.PUT("/baseTrademark/update", controller.TrademarkController.UpdateTrademark)
		adminProductGroup.DELETE("/baseTrademark/remove/:id", controller.TrademarkController.DeleteTrademark)
		adminProductGroup.GET("/baseTrademark/getTrademarkList", controller.TrademarkController.GetAllTrademarkList)

		// 分类管理
		adminProductGroup.GET("/getCategory1", controller.CategoryController.GetCategory1)
		adminProductGroup.GET("/getCategory2/:id", controller.CategoryController.GetCategory2)
		adminProductGroup.GET("/getCategory3/:id", controller.CategoryController.GetCategory3)

		// 如下两个接口没有使用到
		adminProductGroup.POST("/saveCategory2", controller.CategoryController.CreateCategory2)
		adminProductGroup.POST("/saveCategory3", controller.CategoryController.CreateCategory3)

		// 属性管理
		adminProductGroup.POST("/saveAttrInfo", controller.AttrController.SaveAttrInfo)
		adminProductGroup.GET("/attrInfoList/:c1Id/:c2Id/:c3Id", controller.AttrController.GetAttr)
		adminProductGroup.DELETE("/deleteAttr/:attrId", controller.AttrController.DeleteAttr)

		// 商品 SPU 接口
		adminProductGroup.GET("/baseSaleAttrList", controller.SpuController.GetSaleAttrList)
		adminProductGroup.POST("/saveSpuInfo", controller.SpuController.SaveSpuInfo)
		adminProductGroup.GET("/:page/:limit", controller.SpuController.GetSpuList)
		adminProductGroup.POST("/updateSpuInfo", controller.SpuController.UpdateSpuInfo)
		adminProductGroup.DELETE("/deleteSpu/:id", controller.SpuController.DeleteSpu)

		// 商品 SKU 接口
		adminProductGroup.GET("/spuImageList/:id", controller.SpuController.GetSpuImageList)
		adminProductGroup.GET("/spuSaleAttrList/:id", controller.SpuController.GetSpuSaleAttrList)
		adminProductGroup.POST("/saveSkuInfo", controller.SkuController.SaveSkuInfo)
		adminProductGroup.GET("/findBySpuId/:id", controller.SkuController.FindBySpuId)
		adminProductGroup.GET("/list/:page/:limit", controller.SkuController.GetSkuList)
		adminProductGroup.GET("/onSale/:id", controller.SkuController.OnSaleSku)
		adminProductGroup.GET("/cancelSale/:id", controller.SkuController.CancelSaleSku)
		adminProductGroup.DELETE("/deleteSku/:id", controller.SkuController.DeleteSku)
		adminProductGroup.GET("/getSkuInfo/:id", controller.SkuController.GetSkuInfo)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
