package protocol

func RoomHeightmap(floorplan string, wallheight int) []byte {
	b := EmptyBuffer()
	b.WriteShort(RoomHeightMapHeader)
	b.WriteBoolean(true)
	b.WriteInt(wallheight)
	b.WriteString(floorplan)

	return b.Flush()
}

func RoomOpen() []byte {
	b := EmptyBuffer()
	b.WriteShort(RoomOpenHeader)

	return b.Flush()
}

func RoomModel(roomID int) []byte {
	b := EmptyBuffer()
	b.WriteShort(RoomModelHeader)
	b.WriteString("custom")
	b.WriteInt(roomID)

	return b.Flush()
}
