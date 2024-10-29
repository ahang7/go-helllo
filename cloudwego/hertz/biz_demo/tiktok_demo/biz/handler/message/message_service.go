// Code generated by hertz generator.

package message

import (
	"context"

	message "github.com/ahang7/go-hello/cloudwego/hertz/biz_demo/tiktok_demo/biz/model/social/message"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// MessageChat .
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req message.DouyinMessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(message.DouyinMessageChatResponse)

	c.JSON(consts.StatusOK, resp)
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req message.DouyinMessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(message.DouyinMessageActionResponse)

	c.JSON(consts.StatusOK, resp)
}
