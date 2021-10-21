package routers

import (
	"app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// Comment
	beego.Router("/comments", &controllers.CommentController{}, "get:GetAll")
	beego.Router("/comment/:id", &controllers.CommentController{}, "get:GetOne")
	beego.Router("/post", &controllers.CommentController{}, "post:Post")
	// Comment API (Return は JSON)
	beego.Router("/api/comment", &controllers.CommentController{}, "get:GetAllApi")
	beego.Router("/api/comment/:id", &controllers.CommentController{}, "get:GetOneApi")
}
