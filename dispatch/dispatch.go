package dispatch

import (
	"log"
	"pilosa/core"
	"pilosa/db"
	"pilosa/hold"
)

type Dispatch struct {
	service *core.Service
}

func (self *Dispatch) Init() error {
	log.Println("Starting Dispatcher")
	return nil
}

func (self *Dispatch) Close() {
	log.Println("Shutting down Dispatcher")
}

func (self *Dispatch) Run() {
	log.Println("Dispatch Run...")
	for {
		message := self.service.Transport.Receive()
		log.Println("Processing ", message)
		switch data := message.Data.(type) {
		case db.PingRequest:
			pong := db.Message{Data: db.PongRequest{Id: data.Id}}
			self.service.Transport.Send(&pong, data.Source)
		case db.PongRequest:
			hold.Hold.Set(data.Id, 1, 10)
		}

		/*
			path := message.Data.(string)

			bits := strings.Split(path, "/")

			var fragment_id util.SUUID
			var bitmaps []uint64
			var profile_id uint64
			var s uint64

			command := bits[0]
			if len(bits) > 1 {
				fragment_id = util.Hex_to_SUUID(bits[1])
			}
			if len(bits) > 2 {
				bitmap_ids := strings.Split(bits[2], ",")
				spew.Dump(bitmap_ids)
				for i := range bitmap_ids {
					spew.Dump(i, bitmap_ids[i])
					s, _ = strconv.ParseUint(bitmap_ids[i], 10, 64)
					bitmaps = append(bitmaps, s)
				}
			}
			if len(bits) > 3 {
				profile_id, _ = strconv.ParseUint(bits[3], 10, 64)
			}

			spew.Dump("COMMAND:", command)
			spew.Dump("FRAGID:", fragment_id)
			spew.Dump("BITMAPS:", bitmaps)
			spew.Dump("PROFILEID:", profile_id)

			if command == "set" {
				res, err := self.service.Process.SetBit(fragment_id, bitmaps[0], profile_id)
				spew.Dump("SET")
				spew.Dump(res)
				spew.Dump(err)
			}
			if command == "count" {
				spew.Dump("COUNT")
				bh, err := self.service.Process.Get(fragment_id, bitmaps[0])
				if err != nil {
					spew.Dump(err)
				}
				count, err := self.service.Process.Count(fragment_id, bh)
				if err != nil {
					spew.Dump(err)
				}
				spew.Dump(count)
			}
			if command == "intersect" {
				spew.Dump("INTERSECT")
				var bhs []index.BitmapHandle
				for i := range bitmaps {
					bh, _ := self.service.Process.Get(fragment_id, bitmaps[i])
					bhs = append(bhs, bh)
				}
				bhi, err := self.service.Process.Intersect(fragment_id, bhs)
				if err != nil {
					spew.Dump(err)
				}

				count, err := self.service.Process.Count(fragment_id, bhi)
				if err != nil {
					spew.Dump(err)
				}
				spew.Dump(count)
			}
			if command == "union" {
				spew.Dump("UNION")
				var bhs []index.BitmapHandle
				for i := range bitmaps {
					bh, _ := self.service.Process.Get(fragment_id, bitmaps[i])
					bhs = append(bhs, bh)
				}
				bhi, err := self.service.Process.Union(fragment_id, bhs)
				if err != nil {
					spew.Dump(err)
				}

				count, err := self.service.Process.Count(fragment_id, bhi)
				if err != nil {
					spew.Dump(err)
				}
				spew.Dump(count)
			}
		*/
	}
}

func NewDispatch(service *core.Service) *Dispatch {
	return &Dispatch{service}
}
