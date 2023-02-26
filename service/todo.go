package service

import (
	"todo/dao"
	"todo/model"
)

func GetTodoList(id float64) ([]model.GetTodoListResp, error) {
	q := dao.Q.Todo
	var resp []model.GetTodoListResp
	err := q.Where(q.UserID.Eq(int32(id))).Select(q.ID, q.Todo, q.IsDone).Limit(10).Scan(&resp)
	return resp, err
}

func AddTodoItem(id int32, todo string) error {
	q := dao.Q.Todo
	item := model.Todo{
		UserID: id,
		Todo:   todo,
		IsDone: false,
	}
	return q.Create(&item)
}
