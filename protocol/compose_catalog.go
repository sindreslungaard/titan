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

func CatalogPageComposer() []byte {
	b := EmptyBuffer()
	b.WriteShort(CatalogPageHeader)

	b.WriteInt(3)                // page id
	b.WriteString("default_3x3") // cata type
	b.WriteString("default_3x3") // mode/layout

	// localization
	b.WriteInt(0) // num images
	b.WriteInt(1) // num texts
	b.WriteString("Some text")

	b.WriteInt(1) // num items

	// foreach item
	b.WriteInt(1)                    // id
	b.WriteString("plant_pineapple") // name
	b.WriteBoolean(false)            // rent
	b.WriteInt(1)                    // credits
	b.WriteInt(0)                    // points
	b.WriteInt(0)                    // pointstype
	b.WriteBoolean(true)             // can gift

	b.WriteInt(1) // num products

	b.WriteString("s")    // type
	b.WriteInt(160)       // sprite id
	b.WriteString("")     // extradata stuff for posters, bots, wallpapers etc
	b.WriteInt(1)         // quantity
	b.WriteBoolean(false) // ltd

	// continue foreach item
	b.WriteInt(1)                        // club level
	b.WriteBoolean(false)                // bundle
	b.WriteBoolean(false)                // pet
	b.WriteString("plant_pineapple.png") // preview

	// after items foreach
	b.WriteInt(0)         // offer id
	b.WriteBoolean(false) // seasonal currency

	return b.Compose()
}
