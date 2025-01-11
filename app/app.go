package app

import (
	"github.com/abdlmunemdhaw/customs/config"
	"github.com/abdlmunemdhaw/customs/controller"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

func App(router *gin.Engine) {
	municipalityService := Services.NewMunicipalityService(config.Db);
	municipalityController := controller.NewMunicipalityController(municipalityService);
	municipality := router.Group("/municipality")
	{
		municipality.POST("/",municipalityController.CreateMunicipality);
		municipality.GET("/municipalities",municipalityController.GetMunicipalities);
		municipality.PATCH("/:id",municipalityController.UpdateMunicipality);
		municipality.DELETE("/:id",municipalityController.DeleteMunicipality);
	}
	typeService := Services.NewTypeService(config.Db);
	typeController := controller.NewTypeController(typeService);
	type_table  := router.Group("/type")
	{
		type_table.POST("/",typeController.CreateType);
		type_table.GET("/types",typeController.GetTypes);
		type_table.PATCH("/:id",typeController.UpdateType);
		type_table.DELETE("/:id",typeController.DeleteType);
	}
	sectorService := Services.NewSectorService(config.Db);
	sectorController := controller.NewSectorController(sectorService);
	sector := router.Group("/sector")
	{
		sector.GET("sectors",sectorController.GetAllSector);
		sector.POST("/",sectorController.CreateSector);
		sector.PATCH("/:id",sectorController.UpdateSector);
		sector.DELETE("/:id",sectorController.DeleteSector);
	}
	fieldService := Services.NewFieldService(config.Db);
	fieldController := controller.NewFieldController(fieldService);
	field := router.Group("/field")
	{
		field.POST("/",fieldController.CreateField);
		field.GET("/fields",fieldController.GetFields);
		field.PATCH("/:id",fieldController.UpdateField);
		field.DELETE("/:id",fieldController.DeleteField);
	}
	businessService := Services.NewBusinessService(config.Db);
	businessController := controller.NewBusinessController(businessService);
	business := router.Group("/business");
	{
		business.GET("/businesses",businessController.GetBusiness);
		business.POST("/",businessController.CreateBusiness);
		business.DELETE("/:id",businessController.DeleteBusines);
		business.PATCH("/:id",businessController.UpdateBusiness);
	}
	hsCodeService := Services.NewHsCodeService(config.Db);
	hsCodeController := controller.NewHsCodeController(hsCodeService);
	hs_code := router.Group("/hscode");
	{
		hs_code.GET("/hscodes",hsCodeController.GetHsCodes);
		hs_code.POST("/",hsCodeController.CreateHsCode);
		hs_code.PATCH("/:id",hsCodeController.UpdateHsCode);
		hs_code.DELETE("/:id",hsCodeController.DeleteHsCode);
	}
	documentServic := Services.NewDocumentService(config.Db);
	documentController := controller.NewDocumentController(documentServic);
	document := router.Group("/document");
	{
		document.GET("/documents",documentController.GetDocument);
		document.POST("/",documentController.CreateDocument);
		document.DELETE("/:id",documentController.DeleteDocument);
		document.PATCH("/:id",documentController.UpdateDocument);
	}
	bulletinIssureService := Services.NewBulletinIssuerService(config.Db);
	bulletinIssureController := controller.NewBulletinIssuerController(bulletinIssureService);
	bulletinIssure := router.Group("/bulletinissure");
	{
		bulletinIssure.GET("/bulletinsIssures",bulletinIssureController.GetBulletinIssueres);
		bulletinIssure.POST("/",bulletinIssureController.CreateBusiness);
		bulletinIssure.PATCH("/:id",bulletinIssureController.UpdateBulletinIssuer);
		bulletinIssure.DELETE("/:id",bulletinIssureController.DeleteBulletinIssuer);
	}
	locationService := Services.NewLocationService(config.Db);
	locationConroller := controller.NewLocationConroller(locationService);
	location := router.Group("/location");
	{
		location.GET("/locations",locationConroller.GetLocations);
		location.POST("/",locationConroller.CreateLocation);
		location.DELETE("/:id",locationConroller.DeleteLocation);
		location.PATCH("/:id",locationConroller.UpdateLocation);
	}
	categoryService := Services.NewCategoryService(config.Db);
	categoryController := controller.NewCategoryController(categoryService);
	category := router.Group("/category");
	{
		category.GET("/categories",categoryController.GetCategory);
		category.POST("/",categoryController.CreateCategory);
		category.PATCH("/:id",categoryController.UpdateCategory);
		category.DELETE("/:id",categoryController.DeleteCategory);
	}
	economicBulltenService := Services.NewEconomicBulletinService(config.Db);
	economicBulltenController := controller.NewEconomicBulletinController(economicBulltenService);
	economicBullten := router.Group("/economicbulletin");
	{
		economicBullten.GET("/economicbulletins",economicBulltenController.GetEconomicBulletins);
		economicBullten.POST("/",economicBulltenController.CreateEconomicBulletin);
		economicBullten.PATCH("/:id",economicBulltenController.UpdateEconomicBulletin);
		economicBullten.DELETE("/:id",economicBulltenController.DeleteEconomicBulletin);
	} 
}