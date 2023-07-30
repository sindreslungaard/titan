package hh

import (
	"strings"
	"titan/protocol"
)

func e_userdata(u *User, b protocol.Buffer) {
	u.write(protocol.UserData(
		u.data.ID,
		u.data.Username,
		u.data.Figure,
		u.data.Gender,
		u.data.Motto,
		u.data.Respect,
		u.data.RespectToGive,
		u.data.PetRespectToGive,
	))
	// todo: user perks msg
	// todo: memenu settings
}

func e_credits(u *User, b protocol.Buffer) {
	u.write(protocol.UserCredits(u.data.Credits))
}

func e_updatefigure(u *User, b protocol.Buffer) {
	gender := strings.ToUpper(b.ReadString())

	if gender != "M" && gender != "F" {
		gender = "M"
	}

	figure := b.ReadString()

	u.updatefigure(Gender(gender), figure)
}
