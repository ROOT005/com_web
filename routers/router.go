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
	"github.com/qor/media_library"
	//"github.com/qor/notification"
	//notice_data "github.com/qor/notification/channels/database"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/validations"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/confirm"
	_ "gopkg.in/authboss.v0/recover"
	_ "gopkg.in/authboss.v0/register"
	"net/http"
)

type AdminAuth struct{}

var (
	Auth = authboss.New()
)

func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/logout"
}
func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	email, err := c.Request.Cookie("id")
	value, _ := c.Request.Cookie("see")
	if err == nil && value.Value == "BgQDwQ3THJn9F7NPLBi6hTI3Fwz55h47jQUVCOL6iq" {
		var user models.AdminUser
		if !db.DB.First(&user, "email = ?", email).RecordNotFound() {
			return &user
		}
	}
	return nil
}
func init() {
	/*链接数据库*/
	DB := db.DB
	/*本地化数据*/
	I18n := i18n.New(
		database.New(DB),
	)
	/*注册资源*/
	DB.AutoMigrate(&models.User{}, &models.Classfication{}, &models.Product{}, &models.Blog{}, &models.Index{}, &models.Project{}, &models.Pictures{}, &models.Seo{}, &models.MediaLibrary{}, &media_library.AssetManager{}, &models.CountRange{}, &models.Category{}, &models.ProjectRange{}, &models.BlogCategory{}, &models.AdminUser{})
	Admin := admin.New(&qor.Config{DB: DB})
	media_library.RegisterCallbacks(DB)

	//为富文本工具添加Asset
	assetManager := Admin.AddResource(&media_library.AssetManager{}, &admin.Config{Invisible: true})

	/****************添加菜单**************************/
	//控制面板
	Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin"})
	//管理员管理
	admin_auth := AdminAuth{}
	Admin.SetAuth(&admin_auth)
	adminuser := Admin.AddResource(&models.AdminUser{})
	adminuser.Meta(&admin.Meta{Name: "Password",
		Type:            "password",
		FormattedValuer: func(interface{}, *qor.Context) interface{} { return "" },
		Setter: func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			values := metaValue.Value.([]string)
			if len(values) > 0 {
				if newPassword := values[0]; newPassword != "" {
					bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
					if err != nil {
						context.DB.AddError(validations.NewError(adminuser, "Password", "无法加密密码"))
						return
					}
					u := resource.(*models.AdminUser)
					u.Password = string(bcryptPassword)
				}
			}
		},
	})

	//用户管理
	user := Admin.AddResource(&models.User{})
	user.Filter(&admin.Filter{
		Name: "State",
		Config: &admin.SelectOneConfig{
			Collection: []string{"未查看", "已查看"},
		},
	})
	user.Action(&admin.Action{
		Name: "Check It",
		Handle: func(actionArgument *admin.ActionArgument) error {
			for _, record := range actionArgument.FindSelectedRecords() {
				actionArgument.Context.DB.Model(record.(*models.User)).Update("state", "已查看")
			}
			return nil
		},
		Modes: []string{"index", "edit", "show", "menu_item"},
	})
	//消息管理
	/*var (
		users   interface{}
		context *qor.Context
	)
	Notification := notification.New(&notification.Config{})
	Notification.RegisterChannel(notice_data.New(&notice_data.Config{DB: db.DB}))
	Notification.GetNotifications(users, context)
	Admin.NewResource(Notification)*/
	//产品管理
	product := Admin.AddResource(&models.Product{}, &admin.Config{PageCount: 20, Menu: []string{"Product Management"}})
	product.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{AssetManager: assetManager, Plugins: []admin.RedactorPlugin{
		{Name: "table", Source: "/admin/asset/js/redactor_table.js"},
	}}})
	product.Meta(&admin.Meta{Name: "Classfication", Config: &admin.SelectOneConfig{AllowBlank: false}})
	product.Meta(&admin.Meta{Name: "CountRange", Config: &admin.SelectOneConfig{AllowBlank: false}})
	product.IndexAttrs("ID", "Name", "Rate", "SourceCompany", "CreatedAt")
	Admin.AddResource(&models.Classfication{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.CountRange{}, &admin.Config{Menu: []string{"Product Management"}})

	//项目管理
	project := Admin.AddResource(&models.Project{}, &admin.Config{PageCount: 20, Menu: []string{"Order Management"}})
	project.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{AssetManager: assetManager, Plugins: []admin.RedactorPlugin{
		{Name: "table", Source: "/admin/asset/js/redactor_table.js"},
	}}})
	project.Meta(&admin.Meta{Name: "Category", Config: &admin.SelectOneConfig{AllowBlank: false}})
	project.Meta(&admin.Meta{Name: "ProjectRange", Config: &admin.SelectOneConfig{AllowBlank: false}})
	project.IndexAttrs("ID", "Name", "Count", "CreatedAt")
	Admin.AddResource(&models.Category{}, &admin.Config{Menu: []string{"Order Management"}})
	Admin.AddResource(&models.ProjectRange{}, &admin.Config{Menu: []string{"Order Management"}})

	//博客管理
	blog := Admin.AddResource(&models.Blog{}, &admin.Config{Menu: []string{
		"Site Management"}})
	blog.Meta(&admin.Meta{Name: "Content", Config: &admin.RichEditorConfig{AssetManager: assetManager, Plugins: []admin.RedactorPlugin{
		{Name: "table", Source: "/admin/asset/js/redactor_table.js"},
	}}})
	blog.Meta(&admin.Meta{Name: "Abstract", Config: &admin.RichEditorConfig{AssetManager: assetManager, Plugins: []admin.RedactorPlugin{
		{Name: "table", Source: "/admin/asset/js/redactor_table.js"},
	}}})
	blog.IndexAttrs("ID", "Title", "Author", "CreatedAt")
	Admin.AddResource(&models.BlogCategory{}, &admin.Config{Menu: []string{"Site Management"}})
	index := Admin.AddResource(&models.Index{}, &admin.Config{Menu: []string{"Site Management"}})
	index.IndexAttrs("ID", "Title", "CreatedAt")
	Admin.AddResource(&models.Seo{}, &admin.Config{Menu: []string{"Site Management"}})
	Admin.AddResource(I18n)

	// Register Auth for QOR Admin
	/*********************创建路由*********************************/
	//创建admin路由
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
	//创建项目路由
	beego.Router("/projects", &controllers.ProjectController{})
	beego.Router("/projects/project_info", &controllers.ProjectController{}, "get:Project_Info")
	//创建新闻页路由
	beego.Router("/news", &controllers.NewsController{})
	beego.Router("/news/newinfo", &controllers.NewsController{}, "get:News_Info")
	//创建关于我们页
	beego.Router("/about", &controllers.AboutController{})
	//咨询页
	beego.Router("/consult", &controllers.ConsultController{})
	beego.Router("/verify", &controllers.VerifyController{})
	beego.Router("/submit", &controllers.SubmitController{})
	beego.Router("/submit/phone", &controllers.SubmitController{}, "get:Phone")
	//登录页
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	//后台用户数据获取
	beego.Router("/users", &controllers.NewUsersController{})
	beego.Router("/dowload", &controllers.DowloadController{})
}
