package hh

import (
	"sync"
	"titan/db"
	"titan/protocol"

	"github.com/rs/zerolog/log"
)

var catalog = &Catalog{}

type Catalog struct {
	sync.RWMutex
	pages         map[int]db.CatalogPage
	computedpages map[int][]byte
}

func LoadCatalog() {
	catalog.Lock()
	defer catalog.Unlock()

	catalog = &Catalog{
		pages:         make(map[int]db.CatalogPage),
		computedpages: make(map[int][]byte),
	}

	var pages []db.CatalogPage
	err := db.Conn.Order("'order'").Find(&pages).Error

	if err != nil {
		log.Error().Err(err).Msg("Failed to load catalog pages")
		return
	}

	for _, p := range pages {
		catalog.pages[p.ID] = p
	}

	log.Info().Int("pages", len(pages)).Msg("Loaded catalog")
}

func (c *Catalog) pagesfor(rank int) []byte {
	c.RLock()
	defer c.RUnlock()

	cache, ok := c.computedpages[rank]

	if ok {
		return cache
	}

	pages := []protocol.CatalogPage{}

	var crawl func(db.CatalogPage) protocol.CatalogPage
	crawl = func(page db.CatalogPage) protocol.CatalogPage {
		p := protocol.CatalogPage{
			ID:       page.ID,
			Visible:  true,
			Icon:     page.Icon,
			Name:     page.Name,
			Caption:  page.Name,
			Subpages: []protocol.CatalogPage{},
		}

		for _, subpage := range catalog.pages {
			if subpage.ParentID != page.ID || subpage.MinRank > rank {
				continue
			}
			p.Subpages = append(p.Subpages, crawl(subpage))
		}

		return p
	}

	for _, p := range catalog.pages {
		if p.ParentID != -1 || p.MinRank > rank {
			continue
		}

		pages = append(pages, crawl(p))
	}

	log.Debug().Int("rank", rank).Msg("Created catalog pages cache")

	data := protocol.CatalogPages(pages)
	catalog.computedpages[rank] = data

	return data
}
