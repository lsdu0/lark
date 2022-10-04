package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/proto/pb_req"
)

func setNewChatRequestResp(resp *pb_req.NewChatRequestResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatInviteService) NewChatRequest(_ context.Context, req *pb_req.NewChatRequestReq) (resp *pb_req.NewChatRequestResp, _ error) {
	resp = new(pb_req.NewChatRequestResp)
	var (
		request = new(po.ChatRequest)
		err     error
	)
	copier.Copy(request, req)
	request.RequestId = xsnowflake.NewSnowflakeID()
	err = s.chatInviteRepo.RequestCreate(request)
	if err != nil {
		setNewChatRequestResp(resp, ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED)
		xlog.Warn(resp, ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED, err)
		return
	}
	return
}
