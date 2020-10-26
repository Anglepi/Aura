package reminders

import "google.golang.org/genproto/googleapis/type/datetime"

type Reminder struct {
	id int
	reminder string
	time_notify string
}

func (r Reminder) load() {

}

func (r Reminder) save() {

}

func (r Reminder) edit() {
	
}