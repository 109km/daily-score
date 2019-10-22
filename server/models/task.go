package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

var (
	TaskList []*Task
)

func GetTask(id int) (task Task, err error) {

	res := Task{Id: id}
	resErr := OrmInstance.Read(&res)

	if resErr == orm.ErrNoRows {
		return res, errors.New("查询不到")
	} else if resErr == orm.ErrMissPK {
		return res, errors.New("找不到主键")
	} else {
		return res, nil
	}

	return res, errors.New("Task not exists")
}

func GetAllTasks() []*Task {
	qs := OrmInstance.QueryTable("task")
	qs.All(&TaskList)
	return TaskList
}
