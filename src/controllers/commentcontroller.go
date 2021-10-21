package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"app/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//  CommentController operations for Comment
type CommentController struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	driver := beego.AppConfig.String("mysql_driver")
	user := beego.AppConfig.String("mysql_user")
	pass := beego.AppConfig.String("mysql_password")
	host := beego.AppConfig.String("mysql_host")
	port := beego.AppConfig.String("mysql_port")
	db := beego.AppConfig.String("mysql_db")
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, host, port, db)
	orm.RegisterDataBase("default", driver, datasource, 30)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		msg := fmt.Sprintln("Orm.RunSyncdb エラー")
		beego.Debug(msg)
		fmt.Println(msg)
	}

	beego.AddFuncMap("dateformatJst", func(in time.Time) string {
		in = in.Add(time.Duration(0) * time.Hour)
		return in.Format("2006-01-02 15:04:05")
	})

	beego.AddFuncMap("nl2br", func(in string) string {
		return strings.Replace(in, "\n", "<br>", -1)
	})
}

// URLMapping ...
func (c *CommentController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Comment
// @Param	body		body 	models.Comment	true		"body for Comment content"
// @Success 201 {int} models.Comment
// @Failure 403 body is empty
// @router / [post]
func (c *CommentController) Post() {
	var v models.Comment

	if err := c.ParseForm(&v); err != nil {
		c.Data["Comment"] = err.Error()
		msg := fmt.Sprintln(err)
		beego.Debug(msg)
		return
	}
	if _, err := models.AddComment(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["Comment"] = v
	} else {
		c.Data["Comment"] = err.Error()
	}
	c.Ctx.Redirect(302, "/comments")
	// c.ServeJSON()
}

// func (c *CommentController) Post() {
// 	var v models.Comment
// 	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
// 	msg := fmt.Sprintln("Postの中身=>", c.Ctx.Input.RequestBody)
// 	beego.Debug(msg)
// 	if _, err := models.AddComment(&v); err == nil {
// 		c.Ctx.Output.SetStatus(201)
// 		c.Data["json"] = v
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }

// GetOne ...
// @Title Get One
// @Description get Comment by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comment
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CommentController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCommentById(id)
	if err != nil {
		c.Data["Comment"] = err.Error()
	} else {
		c.Data["Comment"] = v
	}
	c.TplName = "comment/show.html.tpl"
	//c.ServeJSON()
}

func (c *CommentController) GetOneApi() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCommentById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

type Comments []models.Comment

// GetAll ...
// @Title Get All
// @Description get Comment
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Comment
// @Failure 403
// @router / [get]
func (c *CommentController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	sortby = append(sortby, "Id") // 追加(数を合わせる)
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	order = append(order, "desc") // 追加(数を合わせる)
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["Comment"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllComment(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["Comment"] = err.Error()
	} else {
		c.Data["Comment"] = l
	}
	c.Data["Website"] = "Chat App"
	// msg := fmt.Sprintln("データベースの中身は", l)
	// beego.Debug(msg)
	c.TplName = "comment/show.html.tpl"
	// c.ServeJSON()
}

func (c *CommentController) GetAllApi() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllComment(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Comment
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Comment	true		"body for Comment content"
// @Success 200 {object} models.Comment
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CommentController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Comment{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateCommentById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Comment
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CommentController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteComment(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
