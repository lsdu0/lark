package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
)

func setGetChatMemberPushConfigResp(resp *pb_chat_member.GetChatMemberPushConfigResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatMemberService) GetChatMemberPushConfig(ctx context.Context, req *pb_chat_member.GetChatMemberPushConfigReq) (resp *pb_chat_member.GetChatMemberPushConfigResp, _ error) {
	resp = &pb_chat_member.GetChatMemberPushConfigResp{Config: &pb_chat_member.ChatMemberPushConfig{}}
	var (
		w      = entity.NewMysqlWhere()
		config *pb_chat_member.ChatMemberPushConfig
		err    error
	)
	w.SetFilter("chat_id = ?", req.ChatId)
	w.SetFilter("uid = ?", req.Uid)

	config, err = s.chatMemberRepo.ChatMemberPushConfig(w)
	if err != nil {
		setGetChatMemberPushConfigResp(resp, ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	resp.Config = config
	go s.cacheMemberPushConfig([]*pb_chat_member.ChatMemberPushConfig{config})
	return
}
