package hh

import (
	"fmt"
	"time"
	"titan/program"
	"titan/protocol"
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
	iter := r.users.iter()
	toremove := []uint64{}

	for _, u := range iter {
		if u.lockmvmnt {
			continue
		}

		// if the roomuser is in the process of being kicked, set target and inc kick cycle
		// note: kicking a player makes them walk to the door and does not remove them from the room immediately
		if u.kickcycle > -1 {
			u.targetX = r.tilemap.doorX
			u.targetY = r.tilemap.doorY
			u.kickcycle++
		}

		// check if the player wants to move
		mv := u.x != u.targetX || u.y != u.targetY

		// remove roomuser manually if being kicked but can't reach the door in given amount of ticks
		if u.kickcycle > -1 && (!mv || u.kickcycle > 20 || (u.kickcycle > 1 && !u.walking)) {
			toremove = append(toremove, u.id)

			// continue??
		}

		// clear the users movement if they don't want to move
		if !mv {
			if u.walking {
				u.actions.remove(ActionMove)
				u.setstep = true
			}
			u.walking = false

		} else { // if the user does want to move
			x, y := r.pathfind(u.x, u.y, u.z, u.targetX, u.targetY)

			if x == u.x && y == u.y {
				u.actions.remove(ActionMove)
				u.target(x, y)
				u.walking = false
				u.setstep = true
			} else {
				if u.z < 0 {
					u.z = 0
					// todo: broadcast
				}

				u.walking = true
				u.actions.remove(ActionSit)

				u.direction = direction(u.x, u.y, x, y)

				//z, _ := r.ItemSystem.HighestPointForUser(x, y)
				z := float32(r.tileheight(x, y)) // todo: replace this with above code
				r.mv(u, x, y, z)
				u.actions.add(ActionMove, Action{value: fmt.Sprintf("%v,%v,%.2f", u.x, u.y, u.z), duration: -1})
			}
		}

		if u.setstep {

			// is the user stopping on a chair?
			/* for _, val := range r.ItemSystem.itemPoints.Get(Point{X: p.x, Y: p.y}).Iter() {

				item := val.(*FloorItem)

				if item.Definition().Chair {
					p.Actions.Add("Sit", -1)
					p.z = item.z + item.Height() - 1
					p.direction = item.direction
					break
				}

			} */

			u.setstep = false

		}

		if u.actions.tick() {
			u.needsupdate = true
		}
	}

	updates := []protocol.SerializedRoomUserStatus{}
	for _, u := range iter {
		if !u.needsupdate {
			continue
		}

		u.needsupdate = false
		updates = append(updates, u.serializestatus())
	}

	if len(updates) > 0 {
		r.broadcast(protocol.RoomUserStatus(updates))
	}

	/*
		for _, usr := range toRemove {
			system.RemoveAvatar(usr)
		}
	*/

	// system.AvatarsLastTick = avatarCount
}
