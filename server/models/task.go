package models

var (
	TaskList []*Task
)

func GetTask(id int64) (task Task, err error) {

	res := Task{Id: id}
	resErr := OrmInstance.Read(&res)

	errorMsg := ProceedSearchError(resErr)
	return res, errorMsg
}

func GetAllTasks() []*Task {
	qs := OrmInstance.QueryTable("task")
	qs.All(&TaskList)
	return TaskList
}
