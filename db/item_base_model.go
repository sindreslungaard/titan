package db

type ItemType string

const (
	FloorItem ItemType = "s"
	WallItem  ItemType = "w"
)

type ItemBase struct {
	ID               int
	Name             string
	Sprite           string
	SpriteID         int
	Type             ItemType
	Width            int
	Length           int
	Height           float32
	Behaviour        string
	InteractionCount int
	CanStack         bool
	CanSit           bool
	CanWalk          bool
	CanGift          bool
	CanTrade         bool
}

func (ItemBase) TableName() string {
	return "items_base"
}
