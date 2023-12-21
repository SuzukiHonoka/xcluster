package hive

import (
	"log"
	"sync"
)

// Minor: Capacity limits
//const MaxCapacity = 65535

type Hive struct {
	//Bees      map[server.ID]*Bee
	Bees      *sync.Map
	Engage    chan *Bee
	Refrain   chan *Bee
	BroadCast chan Message
	working   bool
}

func NewHive() *Hive {
	return &Hive{
		Bees:    &sync.Map{},
		Engage:  make(chan *Bee),
		Refrain: make(chan *Bee),
		working: true,
	}
}

func (h *Hive) Open() {
	go h.open()
}

func (h *Hive) Close() {
	var err error
	h.Bees.Range(func(key, value any) bool {
		bee := value.(*Bee)
		if err = bee.GoAway(); err != nil {
			log.Printf("hive: bee disconnect failed, err=%s", err)
			// will not stop iteration
		}
		h.Bees.Delete(key)
		return true
	})
}

func (h *Hive) open() {
	for h.working {
		select {
		case bee := <-h.Engage:
			h.Bees.Store(bee.ServerID, bee)
			log.Printf("hive: [%s] engaged", bee)
		case bee := <-h.Refrain:
			if err := bee.GoAway(); err != nil {
				log.Printf("hive: [%s] gone with error, err=%s", bee, err)
				// continue removes it from the hive
			}
			if _, ok := h.Bees.LoadAndDelete(bee.ServerID); ok {
				log.Printf("hive: [%s] refrained", bee)
			} else {
				log.Printf("hive: [%s] refrain failed, not found", bee)
			}
		case msg := <-h.BroadCast:
			h.Bees.Range(func(_, value any) bool {
				bee := value.(*Bee)
				select {
				case bee.MessageReceiver <- msg:
				default:
					close(bee.MessageReceiver)
					h.Bees.Delete(bee.ServerID)
				}
				return true
			})
		}
	}
}
