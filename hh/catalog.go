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

	baseitems map[int]db.ItemBase

	items         map[int]db.CatalogItem
	computeditems map[int][]byte

	pages         map[int]db.CatalogPage
	computedpages map[int][]byte
}

func LoadCatalog() {
	catalog = &Catalog{
		baseitems: make(map[int]db.ItemBase),

		items:         make(map[int]db.CatalogItem),
		computeditems: make(map[int][]byte),

		pages:         make(map[int]db.CatalogPage),
		computedpages: make(map[int][]byte),
	}

	catalog.Lock()
	defer catalog.Unlock()

	catalog.loadbaseitems()
	catalog.loaditems()
	catalog.loadpages()
}

func (c *Catalog) loadbaseitems() {
	var items []db.ItemBase
	err := db.Conn.Find(&items).Error

	if err != nil {
		log.Error().Err(err).Msg("Failed to load item bases")
		return
	}

	for _, i := range items {
		c.baseitems[i.ID] = i
	}
}

func (c *Catalog) loaditems() {
	var items []db.CatalogItem
	err := db.Conn.Find(&items).Error

	if err != nil {
		log.Error().Err(err).Msg("Failed to load item bases")
		return
	}

	for _, i := range items {
		c.items[i.ID] = i
	}
}

func (c *Catalog) loadpages() {
	var pages []db.CatalogPage
	err := db.Conn.Order("'order'").Find(&pages).Error

	if err != nil {
		log.Error().Err(err).Msg("Failed to load catalog pages")
		return
	}

	for _, p := range pages {
		c.pages[p.ID] = p
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
