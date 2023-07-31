package hh

import "titan/protocol"

func e_requestcatalog(u *User, b protocol.Buffer) {
	rank := 1
	u.write(catalog.pagesfor(rank))
}

func e_requestcatalogpage(u *User, b protocol.Buffer) {
	u.write(protocol.CatalogPageComposer())
}
