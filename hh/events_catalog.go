package hh

import "titan/protocol"

func e_requestcatalog(u *User, b protocol.Buffer) {
	rank := 1
	u.write(catalog.pagesfor(rank))
}
