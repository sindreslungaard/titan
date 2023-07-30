package db

type RoomType int

const (
	OpenRoom RoomType = iota
	LockedRoom
	PasswordRoom
	HiddenRoom
)

type Room struct {
	ID             int
	Owner          int
	Name           string
	Description    string
	RoomType       RoomType `gorm:"default:0"`
	Password       string
	MaxUsers       int `gorm:"default:25"`
	FloorPlan      string
	DoorX          int
	DoorY          int
	WalkThrough    bool
	FloorThickness int
	WallThickness  int
	WallHeight     int
	WallHidden     bool
}
