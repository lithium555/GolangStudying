package Models

import "time"

type OwnAlarm struct{
	ID int
	WelcomeMessage string
	playDate time.Time
}

func (a OwnAlarm) SetTime(t time.Time) {
	panic("implement me")
}

func (a OwnAlarm) SetAlarmSound(s string) {
	panic("implement me")
}

func (a OwnAlarm) PlayAlarm() {
	panic("implement me")
}

