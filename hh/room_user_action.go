package hh

import (
	"fmt"
	"strings"
	"sync"
)

type Action struct {
	value    string
	duration int
}

type ActionKey string

const (
	ActionMove ActionKey = "mv"
	ActionSit  ActionKey = "sit"
)

// duration in ticks
type Actions struct {
	mutex   sync.RWMutex
	actions map[ActionKey]Action
	changed bool
}

func actions() Actions {
	return Actions{
		actions: map[ActionKey]Action{},
	}
}

func (a *Actions) has(key ActionKey) bool {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	_, ok := a.actions[key]

	return ok
}

func (a *Actions) add(key ActionKey, value Action) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.actions[key] = value
	a.changed = true
}

func (a *Actions) remove(key ActionKey) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	_, ok := a.actions[key]

	if ok {
		delete(a.actions, key)
		a.changed = true
	}
}

func (a *Actions) tick() (needsUpdate bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	for key, action := range a.actions {
		if action.duration != -1 {
			action.duration--
			a.actions[key] = action
		}

		if action.duration == 0 {
			delete(a.actions, key)
			a.changed = true
		}
	}

	needsUpdate = a.changed
	a.changed = false

	return needsUpdate
}

func (a *Actions) serialized() string {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	result := "/"
	for key, action := range a.actions {
		result += fmt.Sprintf("%s %s/", key, action.value)
	}

	return result
}

func chataction(msg string) (string, int, bool) {
	switch strings.ToLower(msg) {
	case "o/":
		return "Wave", 6, true
	case "_b":
		return "Respect", 6, true
	case ":d", ";d":
		return "Laugh", 6, false
	case ":)", ";)":
		return "GestureSmile", 6, false
	case ":o":
		return "GestureSurprised", 6, false
	case ">:(", ":@":
		return "GestureAngry", 6, false
	case ":(", ":'(", "=[", "='[":
		return "GestureSad", 6, false
	case ":kiss":
		return "Blow", 2, true
	default:
		return "", 0, false
	}
}
