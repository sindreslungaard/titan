package protocol

import "fmt"

func UserHomeRoom(home int, lobby int) []byte {
	b := EmptyBuffer()
	b.WriteShort(UserHomeRoomHeader)
	b.WriteInt(home)
	b.WriteInt(lobby)

	return b.Compose()
}

func UserEffectsList() []byte {
	b := EmptyBuffer()
	b.WriteShort(UserEffectsListHeader)
	b.WriteInt(0)

	return b.Compose()
}

func UserClothes() []byte {
	b := EmptyBuffer()
	b.WriteShort(UserClothesHeader)
	b.WriteInt(0)
	b.WriteInt(0)

	return b.Compose()
}

func NewUserIdentity() []byte {
	b := EmptyBuffer()
	b.WriteShort(NewUserIdentityHeader)
	b.WriteInt(1)

	return b.Compose()
}

func UserPermissions(rank int, clublevel int, ambassador bool) []byte {
	b := EmptyBuffer()
	b.WriteShort(UserPermissionsHeader)
	b.WriteInt(clublevel)
	b.WriteInt(rank)
	b.WriteBoolean(ambassador)

	return b.Compose()
}

func AvailabilityStatusMessage(open bool, closing bool, authentic bool) []byte {
	b := EmptyBuffer()
	b.WriteShort(AvailabilityStatusMessageHeader)
	b.WriteBoolean(open)
	b.WriteBoolean(closing)
	b.WriteBoolean(authentic)

	return b.Compose()
}

func Ping() []byte {
	b := EmptyBuffer()
	b.WriteShort(PingHeader)

	return b.Compose()
}

func EnableNotifications(enabled bool) []byte {
	b := EmptyBuffer()
	b.WriteShort(EnableNotificationsHeader)
	b.WriteBoolean(enabled)

	return b.Compose()
}

func UserAchievementScore(score int) []byte {
	b := EmptyBuffer()
	b.WriteShort(UserAchievementScoreHeader)
	b.WriteInt(score)

	return b.Compose()
}

func IsFirstLoginOfDay(value bool) []byte {
	b := EmptyBuffer()
	b.WriteShort(IsFirstLoginOfDayHeader)
	b.WriteBoolean(value)

	return b.Compose()
}

func MysteryBoxKeys() []byte {
	b := EmptyBuffer()
	b.WriteShort(MysteryBoxKeysHeader)
	b.WriteString("")
	b.WriteString("")

	return b.Compose()
}

func UserClub() []byte {
	b := EmptyBuffer()
	b.WriteShort(UserClubHeader)
	b.WriteString("habbo_club")
	b.WriteInt(365) // days 'til expiration
	b.WriteInt(0)
	b.WriteInt(0)
	b.WriteInt(1) // response type | normal=0 | login=1 | purchase=2 | discount=3 | citizendiscount=4
	b.WriteBoolean(true)
	b.WriteBoolean(true)
	b.WriteInt(0)
	b.WriteInt(30)       // past vip days
	b.WriteInt(365 * 60) // minutes 'til expiration (total)
	b.WriteInt(1)        // last modified

	return b.Compose()
}

func GenericAlert(msg string) []byte {
	b := EmptyBuffer()
	b.WriteShort(GenericAlertHeader)
	b.WriteString(msg)

	return b.Compose()
}

func UserData(id int, username string, figure string, gender string, motto string, respectReceived int, respectToGive int, petRespectToGive int) []byte {
	b := EmptyBuffer()
	b.WriteShort(UserDataHeader)

	b.WriteInt(id)
	b.WriteString(username)
	b.WriteString(figure)
	b.WriteString(gender)
	b.WriteString(motto)
	b.WriteString(username)
	b.WriteBoolean(false)
	b.WriteInt(respectReceived)
	b.WriteInt(respectToGive)
	b.WriteInt(petRespectToGive)
	b.WriteBoolean(false)
	b.WriteString("01-01-1970 00:00:00")
	b.WriteBoolean(false)
	b.WriteBoolean(false)

	return b.Compose()
}

func UserCredits(credits int) []byte {
	b := EmptyBuffer()
	b.WriteShort(UserCreditsHeader)
	b.WriteString(fmt.Sprintf("%v.0", credits))

	return b.Compose()
}

func RoomUserData(id uint64, look string, gender string, motto string, achievementscore int) []byte {
	b := EmptyBuffer()
	b.WriteShort(RoomUserDataHeader)

	b.WriteInt(int(id))
	b.WriteString(look)
	b.WriteString(gender)
	b.WriteString(motto)
	b.WriteInt(achievementscore)

	return b.Compose()
}
