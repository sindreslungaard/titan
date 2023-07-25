package protocol

func UserHomeRoom(home int, lobby int) []byte {
	b := EmptyBuffer()
	b.WriteShort(UserHomeRoomHeader)
	b.WriteInt(home)
	b.WriteInt(lobby)

	return b.Flush()
}

func UserEffectsList() []byte {
	b := EmptyBuffer()
	b.WriteShort(UserEffectsListHeader)
	b.WriteInt(0)

	return b.Flush()
}

func UserClothes() []byte {
	b := EmptyBuffer()
	b.WriteShort(UserClothesHeader)
	b.WriteInt(0)
	b.WriteInt(0)

	return b.Flush()
}

func NewUserIdentity() []byte {
	b := EmptyBuffer()
	b.WriteShort(NewUserIdentityHeader)
	b.WriteInt(1)

	return b.Flush()
}

func UserPermissions(rank int, clublevel int, ambassador bool) []byte {
	b := EmptyBuffer()
	b.WriteShort(UserPermissionsHeader)
	b.WriteInt(clublevel)
	b.WriteInt(rank)
	b.WriteBoolean(ambassador)

	return b.Flush()
}

func AvailabilityStatusMessage(open bool, closing bool, authentic bool) []byte {
	b := EmptyBuffer()
	b.WriteShort(AvailabilityStatusMessageHeader)
	b.WriteBoolean(open)
	b.WriteBoolean(closing)
	b.WriteBoolean(authentic)

	return b.Flush()
}

func Ping() []byte {
	b := EmptyBuffer()
	b.WriteShort(PingHeader)

	return b.Flush()
}

func EnableNotifications(enabled bool) []byte {
	b := EmptyBuffer()
	b.WriteShort(EnableNotificationsHeader)
	b.WriteBoolean(enabled)

	return b.Flush()
}

func UserAchievementScore(score int) []byte {
	b := EmptyBuffer()
	b.WriteShort(UserAchievementScoreHeader)
	b.WriteInt(score)

	return b.Flush()
}

func IsFirstLoginOfDay(value bool) []byte {
	b := EmptyBuffer()
	b.WriteShort(IsFirstLoginOfDayHeader)
	b.WriteBoolean(value)

	return b.Flush()
}

func MysteryBoxKeys() []byte {
	b := EmptyBuffer()
	b.WriteShort(MysteryBoxKeysHeader)
	b.WriteString("")
	b.WriteString("")

	return b.Flush()
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

	return b.Flush()
}

func GenericAlert(msg string) []byte {
	b := EmptyBuffer()
	b.WriteShort(GenericAlertHeader)
	b.WriteString(msg)

	return b.Flush()
}
