package repo

import (
	"gorm.io/gorm"
	"lark/apps/request/internal/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type RequestRepository interface {
	RequestCreate(req *po.ChatRequest) (err error)
	TxRequestUpdate(tx *gorm.DB, u *entity.MysqlUpdate) (err error)
	TxRequest(tx *gorm.DB, w *entity.MysqlWhere) (request *po.ChatRequest, err error)
	RequestList(w *entity.MysqlWhere) (list []po.ChatRequest, err error)
	TxChatUsersCreate(tx *gorm.DB, users []*po.ChatUser) (err error)
}

type requestRepository struct {
}

func NewRequestRepository() RequestRepository {
	return &requestRepository{}
}

func (r *requestRepository) RequestCreate(req *po.ChatRequest) (err error) {
	db := xmysql.GetDB()
	err = db.Create(req).Error
	return
}

func (r *requestRepository) TxRequestUpdate(tx *gorm.DB, u *entity.MysqlUpdate) (err error) {
	err = tx.Model(po.ChatRequest{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *requestRepository) TxRequest(tx *gorm.DB, w *entity.MysqlWhere) (request *po.ChatRequest, err error) {
	err = tx.Where(w.Query, w.Args...).Find(request).Error
	return
}

func (r *requestRepository) RequestList(w *entity.MysqlWhere) (list []po.ChatRequest, err error) {
	list = make([]po.ChatRequest, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Limit(w.Limit).Find(&list).Error
	return
}

func (r *requestRepository) TxChatUsersCreate(tx *gorm.DB, users []*po.ChatUser) (err error) {
	err = tx.Create(users).Error
	return
}
