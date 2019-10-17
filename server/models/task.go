package models

var (
	Tasks [10]Task
)

type Task struct {
	id   		int64
	content int
	score_unit int
	total_score int
	start_time string
	end_time string
	date string
}

func init() {
	// for i:= 0; i < 10 ;i++{
	// 	comment := "Hello"
	// 	Tasks[i] = Task{ i, i * 10, comment}
	// }
}

func GetAllTasks() []Task {
	var slices []Task = Tasks[3:6]
	return slices
}

func AddTask(task Task) {
	
}


