package protocol

func NewNavigatorMetaData() []byte {
	b := EmptyBuffer()
	b.WriteShort(NewNavigatorMetaDataHeader)

	b.WriteInt(3)

	b.WriteString("official_view")
	b.WriteInt(0)

	b.WriteString("hotel_view")
	b.WriteInt(0)

	/* b.WriteString("roomads_view")
	b.WriteInt(0) */

	b.WriteString("myworld_view")
	b.WriteInt(0)

	return b.Compose()
}

func NewNavigatorLiftedRooms() []byte {
	b := EmptyBuffer()
	b.WriteShort(NewNavigatorLiftedRoomsHeader)
	b.WriteInt(0)

	return b.Compose()
}

func NewNavigatorCollapsedCategories() []byte {
	b := EmptyBuffer()
	b.WriteShort(NewNavigatorCollapsedCategoriesHeader)
	b.WriteInt(46)
	b.WriteString("new_ads")
	b.WriteString("friend_finding")
	b.WriteString("staffpicks")
	b.WriteString("with_friends")
	b.WriteString("with_rights")
	b.WriteString("query")
	b.WriteString("recommended")
	b.WriteString("my_groups")
	b.WriteString("favorites")
	b.WriteString("history")
	b.WriteString("top_promotions")
	b.WriteString("campaign_target")
	b.WriteString("friends_rooms")
	b.WriteString("groups")
	b.WriteString("metadata")
	b.WriteString("history_freq")
	b.WriteString("highest_score")
	b.WriteString("competition")
	b.WriteString("category__Agencies")
	b.WriteString("category__Role Playing")
	b.WriteString("category__Global Chat & Discussi")
	b.WriteString("category__GLOBAL BUILDING AND DE")
	b.WriteString("category__global party")
	b.WriteString("category__global games")
	b.WriteString("category__global fansite")
	b.WriteString("category__global help")
	b.WriteString("category__Trading")
	b.WriteString("category__global personal space")
	b.WriteString("category__Habbo Life")
	b.WriteString("category__TRADING")
	b.WriteString("category__global official")
	b.WriteString("category__global trade")
	b.WriteString("category__global reviews")
	b.WriteString("category__global bc")
	b.WriteString("category__global personal space")
	b.WriteString("eventcategory__Hottest Events")
	b.WriteString("eventcategory__Parties & Music")
	b.WriteString("eventcategory__Role Play")
	b.WriteString("eventcategory__Help Desk")
	b.WriteString("eventcategory__Trading")
	b.WriteString("eventcategory__Games")
	b.WriteString("eventcategory__Debates & Discuss")
	b.WriteString("eventcategory__Grand Openings")
	b.WriteString("eventcategory__Friending")
	b.WriteString("eventcategory__Jobs")
	b.WriteString("eventcategory__Group Events")

	return b.Compose()
}

type NavigatorRoom struct {
	ID          int
	Name        string
	Description string
	Owner       int
	OwnerName   string
	RoomType    int
	MaxUsers    int
	UserCount   int
}

func NewNavigatorSearchResults(view string, query string, rooms []NavigatorRoom) []byte {
	b := EmptyBuffer()
	b.WriteShort(NewNavigatorSearchResultsHeader)

	b.WriteString(view)
	b.WriteString(query)

	// todo: handle different views, all rooms vs my rooms etc

	b.WriteInt(1) // num categories

	b.WriteString("")
	b.WriteString("Rooms")
	b.WriteInt(0)
	b.WriteBoolean(false)
	b.WriteInt(0)

	b.WriteInt(len(rooms)) // foreach category, num rooms

	for _, r := range rooms {
		b.WriteInt(r.ID)             // room id
		b.WriteString(r.Name)        // room Room
		b.WriteInt(r.Owner)          // owner id
		b.WriteString(r.OwnerName)   // owner name
		b.WriteInt(int(r.RoomType))  // 0=open | 1=locked | 2=pw | 3=invis
		b.WriteInt(r.UserCount)      // users in room
		b.WriteInt(r.MaxUsers)       // max users
		b.WriteString(r.Description) // room desc
		b.WriteInt(0)
		b.WriteInt(0) // upvotes?
		b.WriteInt(0)
		b.WriteInt(1) // category, prob referring to collapsed categories index
		b.WriteInt(0) // num tags
		b.WriteInt(0) // bitmask, base=shift0, group=shift2, promoted=shift4, public=shift8
	}

	return b.Compose()
}
