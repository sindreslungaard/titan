package hh

import (
	"time"
	"titan/program"
)

func (r *Room) process() {
	defer program.Recover()
	defer r.close()

	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()

	lowPrioTicker := time.NewTicker(time.Second * 60)
	defer lowPrioTicker.Stop()

	for {

		select {
		case <-r.sigclose:
			{
				return
			}
		case <-lowPrioTicker.C:
			{
				/* if r.IsInactive() {
					return
				} */
			}
		case <-ticker.C:
			{
				r.tick()

				// do this after an update to lock up resources only when there's
				// some time until the next update

				/* if time.Now().Sub(r.ItemSystem.lastItemSave) >= ItemSaveDelay {
					r.ItemSystem.SaveQueuedItems()
				} */
			}
		}

	}
}

func (r *Room) tick() {
	r.tickroomusers()
}

func (r *Room) tickroomusers() {

}
