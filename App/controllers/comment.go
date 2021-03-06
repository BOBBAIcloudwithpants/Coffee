package controllers

import (
	"github.com/XMatrixStudio/Coffee/App/services"
	"github.com/globalsign/mgo/bson"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"html/template"
)

// CommentController 评论
type CommentController struct {
	Ctx     iris.Context
	Service services.CommentService
	Session *sessions.Session
}

// CommentRes 评论回复
type CommentRes struct {
	State string
	Data  []services.CommentForContent
}

// GetBy GET /comment/{contentID} 获取指定内容的评论
func (c *CommentController) GetBy(id string) (res CommentRes) {
	if !bson.IsObjectIdHex(id) {
		res.State = StatusBadReq
		return
	}
	res.State = StatusSuccess
	res.Data = c.Service.GetComment(id)
	return
}

// postCommentReq 评论数据请求
type postCommentReq struct {
	ContentID string `json:"contentId"`
	FatherID  string `json:"fatherId"`
	Content   string `json:"content"`
	IsReply   bool   `json:"isReply"`
}

// Post POST /comment 增加评论
func (c *CommentController) Post() (res CommonRes) {
	if c.Session.Get("id") == nil {
		res.State = StatusNotLogin
		return
	}
	req := postCommentReq{}
	if err := c.Ctx.ReadJSON(&req); err != nil {
		res.State = StatusBadReq
		return
	}
	if !(bson.IsObjectIdHex(req.ContentID) && (bson.IsObjectIdHex(req.FatherID) || req.FatherID == "")) {
		res.State = StatusBadReq
		return
	}
	if req.Content == "" {
		res.State = StatusBadReq
		return
	}
	// 过滤字符
	req.Content = template.HTMLEscapeString(req.Content)
	err := c.Service.AddComment(c.Session.GetString("id"), req.ContentID, req.FatherID, req.Content, req.IsReply)
	if err != nil {
		res.State = err.Error()
		return
	}
	res.State = StatusSuccess
	return
}

// DeleteBy DELETE /comment/{commentID} 删除指定评论
func (c *CommentController) DeleteBy(id string) (res CommonRes) {
	if c.Session.Get("id") == nil {
		res.State = StatusNotLogin
		return
	}
	if !bson.IsObjectIdHex(id) {
		res.State = StatusBadReq
		return
	}
	err := c.Service.DeleteComment(id, c.Session.GetString("id"))
	if err != nil {
		res.State = err.Error()
		return
	}
	res.State = StatusSuccess
	return
}
