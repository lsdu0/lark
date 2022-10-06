package service

const (
	ERROR_CODE_CHAT_PROTOCOL_UNMARSHAL_ERR int32 = 120001
	ERROR_CODE_CHAT_GRPC_SERVICE_FAILURE   int32 = 120002
	ERROR_CODE_CHAT_INSERT_VALUE_FAILED    int32 = 120003
	ERROR_CODE_CHAT_UPDATE_VALUE_FAILED    int32 = 120004
	ERROR_CODE_CHAT_QUERY_DB_FAILED        int32 = 120005
	ERROR_CODE_CHAT_INITIATOR_INVITEE_SAME int32 = 120006
	ERROR_CODE_CHAT_HAS_JOINED_CHAT        int32 = 120007
	ERROR_CODE_CHAT_PARAM_ERR              int32 = 120008
)

const (
	ERROR_CHAT_PROTOCOL_UNMARSHAL_ERR = "协议反序列化错误"
	ERROR_CHAT_GRPC_SERVICE_FAILURE   = "服务故障"
	ERROR_CHAT_INSERT_VALUE_FAILED    = "数据入库失败"
	ERROR_CHAT_UPDATE_VALUE_FAILED    = "更新Value失败"
	ERROR_CHAT_QUERY_DB_FAILED        = "查询失败"
	ERROR_CHAT_INITIATOR_INVITEE_SAME = "发起者和被邀请人为同一人"
	ERROR_CHAT_HAS_JOINED_CHAT        = "已经加入了"
	ERROR_CHAT_PARAM_ERR              = "参数错误"
)
