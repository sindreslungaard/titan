package hh

import "titan/protocol"

func e_navigatordata(u *User, b protocol.Buffer) {
	/*
		todo:
		NewNavigatorSettings
		NewNavigatorSavedSearches
		NewNavigatorEventCategories
	*/
	u.write(protocol.NewNavigatorMetaData())
	u.write(protocol.NewNavigatorLiftedRooms())
	u.write(protocol.NewNavigatorCollapsedCategories())
}

func e_navigatorsearch(u *User, b protocol.Buffer) {
	view := b.ReadString()
	query := b.ReadString()
	u.write(protocol.NewNavigatorSearchResults(view, query))
}
