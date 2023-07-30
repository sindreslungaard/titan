package protocol

import "fmt"

type SerializedRoomUser struct {
	ID               int
	Username         string
	Motto            string
	Look             string
	VID              int
	X                int
	Y                int
	Z                float32
	Direction        int
	Gender           string
	AchievementScore int
}

type SerializedRoomUserStatus struct {
	VID          int
	PrevX        int
	PrevY        int
	PrevZ        float32
	HeadRotation int
	BodyRotation int
	Statuses     string
}

func RoomHeightmap(floorplan string, wallheight int) []byte {
	b := EmptyBuffer()
	b.WriteShort(RoomHeightMapHeader)
	b.WriteBoolean(true)
	b.WriteInt(wallheight)
	b.WriteString(floorplan)

	return b.Compose()
}

func RoomOpen() []byte {
	b := EmptyBuffer()
	b.WriteShort(RoomOpenHeader)

	return b.Compose()
}

func RoomModel(roomID int) []byte {
	b := EmptyBuffer()
	b.WriteShort(RoomModelHeader)
	b.WriteString("custom")
	b.WriteInt(roomID)

	return b.Compose()
}

func RoomUsers(users []SerializedRoomUser) []byte {
	b := EmptyBuffer()
	b.WriteShort(RoomUsersHeader)

	b.WriteInt(len(users)) // num avatars

	for _, u := range users {
		b.WriteInt(u.ID)
		b.WriteString(u.Username)
		b.WriteString(u.Motto)
		b.WriteString(u.Look)
		b.WriteInt(u.VID)
		b.WriteInt(u.X)
		b.WriteInt(u.Y)
		b.WriteString(fmt.Sprintf("%.2f", u.Z))
		b.WriteInt(u.Direction)
		b.WriteInt(1)
		b.WriteString(u.Gender)
		b.WriteInt(-1) // group
		b.WriteInt(-1)
		b.WriteString("") // group name
		b.WriteString("")
		b.WriteInt(u.AchievementScore)
		b.WriteBoolean(true)
	}

	return b.Compose()
}

func RoomUserStatus(users []SerializedRoomUserStatus) []byte {
	p := EmptyBuffer()
	p.WriteShort(RoomUserStatusHeader)

	p.WriteInt(len(users)) // num avatars

	for _, u := range users {
		p.WriteInt(u.VID)
		p.WriteInt(u.PrevX)
		p.WriteInt(u.PrevY)
		p.WriteString(fmt.Sprintf("%.2f", u.PrevZ))
		p.WriteInt(u.HeadRotation)
		p.WriteInt(u.BodyRotation)
		p.WriteString(u.Statuses)

		// "/mv x,y,z.z"

	}

	return p.Compose()
}

func RoomUserRemove(id uint64) []byte {
	b := EmptyBuffer()
	b.WriteShort(RoomUserRemoveHeader)

	b.WriteString(fmt.Sprintf("%v", id))

	return b.Compose()
}

func Talk(id uint64, emotion int, bubble int, message string) []byte {
	b := EmptyBuffer()
	b.WriteShort(RoomUserTalkHeader)

	b.WriteInt(int(id))
	b.WriteString(message)
	b.WriteInt(emotion)
	b.WriteInt(bubble)
	b.WriteInt(0)
	b.WriteInt(len([]rune(message)))

	return b.Compose()
}
