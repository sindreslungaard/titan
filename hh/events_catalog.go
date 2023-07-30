package hh

import "titan/protocol"

func e_requestcatalog(u *User, b protocol.Buffer) {
	pages := []protocol.CatalogPage{
		{
			ID:      1,
			Visible: true,
			Icon:    1,
			Name:    "Test",
			Caption: "Test",
			Subpages: []protocol.CatalogPage{
				{
					ID:      1,
					Visible: true,
					Icon:    1,
					Name:    "Test2",
					Caption: "Test2",
					Subpages: []protocol.CatalogPage{
						{
							ID:       1,
							Visible:  true,
							Icon:     1,
							Name:     "Test3",
							Caption:  "Test3",
							Subpages: []protocol.CatalogPage{},
						},
					},
				},
			},
		},
	}

	u.write(protocol.CatalogPages(pages))
}
