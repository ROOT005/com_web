package routers

import (
	"com_web/controllers"
	"com_web/db"
	"com_web/models"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qor/admin"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/i18n/backends/yaml"
	"github.com/qor/media_library"
	"github.com/qor/qor"
	"net/http"
	"path/filepath"
)

func init() {
	/*链接数据库*/
	DB := db.DB
	/*本地化数据*/
	I18n := i18n.New(
		database.New(DB),
		yaml.New(filepath.Join("conf/locals")),
	)
	/*注册资源*/
	DB.AutoMigrate(&models.User{}, &models.Classfication{}, &models.Product{}, &models.Blog{}, &models.Index{}, &models.Project{}, &models.Pictures{}, &models.Seo{}, &models.MediaLibrary{}, &media_library.AssetManager{})
	Admin := admin.New(&qor.Config{DB: DB})
	media_library.RegisterCallbacks(DB)

	//为富文本工具添加Asset
	assetManager := Admin.AddResource(&media_library.AssetManager{}, &admin.Config{Invisible: true})
	/****************添加菜单**************************/
	Admin.AddResource(&models.User{})
	//产品管理
	product := Admin.AddResource(&models.Product{}, &admin.Config{PageCount: 20})
	product.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{AssetManager: assetManager, Plugins: []admin.RedactorPlugin{
		{Name: "table", Source: "/admin/asset/js/redactor_table.js"},
	}}})
	product.Meta(&admin.Meta{Name: "Classfication", Config: &admin.SelectOneConfig{AllowBlank: false}})
	product.IndexAttrs("ID", "Name", "Rate", "SourceCompany", "CreatedAt", "Classfication")
	Admin.AddResource(&models.Classfication{})
	//项目管理
	Admin.AddResource(&models.Project{}, &admin.Config{PageCount: 20})

	//博客管理
	blog := Admin.AddResource(&models.Blog{}, &admin.Config{Menu: []string{
		"Site Management"}})
	blog.Meta(&admin.Meta{Name: "Content", Config: &admin.RichEditorConfig{AssetManager: assetManager, Plugins: []admin.RedactorPlugin{
		{Name: "table", Source: "/admin/asset/js/redactor_table.js"},
	}}})
	blog.IndexAttrs("ID", "Title", "Author", "CreatedAt")

	Admin.AddResource(&models.Index{}, &admin.Config{Menu: []string{"Site Management"}})
	Admin.AddResource(&models.Pictures{}, &admin.Config{Menu: []string{"Site Management"}})
	Admin.AddResource(&models.Seo{}, &admin.Config{Menu: []string{"Site Management"}})
	Admin.AddResource(I18n)

	/*********************创建路由*********************************/
	//创建admi路由
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	beego.Handler("/admin/*", mux)
	//404
	beego.ErrorController(&controllers.ErrorController{})
	//创建主页路由
	beego.Router("/", &controllers.HomeController{})
	//产品路由
	beego.Router("/products", &controllers.ProductsController{})
	beego.Router("/products/product_info", &controllers.ProductsController{}, "get:Product_Info")
	//创建博客路由
	beego.RESTRouter("/news", &controllers.NewsController{})
}
