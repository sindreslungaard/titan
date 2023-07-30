package protocol

type CatalogPage struct {
	ID       int
	Visible  bool
	Icon     int
	Name     string
	Caption  string
	Subpages []CatalogPage
}

func CatalogPages(pages []CatalogPage) []byte {
	b := EmptyBuffer()
	b.WriteShort(CatalogPagesListHeader)

	b.WriteBoolean(true)
	b.WriteInt(0)
	b.WriteInt(-1)
	b.WriteString("root")
	b.WriteString("")
	b.WriteInt(0)

	b.WriteInt(len(pages))

	var writePage func(CatalogPage)
	writePage = func(page CatalogPage) {
		b.WriteBoolean(page.Visible)
		b.WriteInt(page.Icon)
		b.WriteInt(page.ID)
		b.WriteString(page.Name)
		b.WriteString(page.Caption)

		b.WriteInt(0) // num offers

		b.WriteInt(len(page.Subpages))

		for _, subpage := range page.Subpages {
			writePage(subpage)
		}
	}

	for _, page := range pages {
		writePage(page)
	}

	b.WriteBoolean(false)
	b.WriteString("NORMAL")

	return b.Compose()
}
