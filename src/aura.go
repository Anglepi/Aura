package main

import (
	"fmt"
	mm "github.com/Anglepi/Aura/src/memes"
	db "github.com/Anglepi/Aura/src/auramemory"
	lg "github.com/Anglepi/Aura/src/logs"
	nt "github.com/Anglepi/Aura/src/notes"
	rm "github.com/Anglepi/Aura/src/reminders"
	us "github.com/Anglepi/Aura/src/users"
)

type Aura struct {
	memeManager mm.MemeManager
	database db.AuraMemory
	logs lg.AuraLogger
	notes nt.NoteManager
	reminders rm.ReminderManager
	users us.UserManager
}

func (a Aura) listenEvent() {

}

func (a Aura) processUserMessage() {

}

func (a Aura) requestUtility() {

}

func (a Aura) answerEvent() {

}

func main(){
	fmt.Println("Hola, soy Aura")
}