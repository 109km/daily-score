package models

var (
	EventsList []*Event
)

func GetEvent(id int64) (event Event, err error) {

	res := Event{Id: id}
	resErr := OrmInstance.Read(&res)

	errorMsg := ProceedSearchError(resErr)
	return res, errorMsg
}

func GetAllEvents() []*Event {
	qs := OrmInstance.QueryTable("event")
	qs.All(&EventsList)
	return EventsList
}

func AddOneEvent(title string, level int, total_score int, start_time string, end_time string) (eid int64, errorMsg error) {
	var event Event
	event.Title = title
	event.Level = level
	event.TotalScore = total_score
	event.StartTime = start_time
	event.EndTime = end_time

	eid, responseError = OrmInstance.Insert(&event)
	errorMsg = ProceedSearchError(responseError)
	return eid, errorMsg
}
