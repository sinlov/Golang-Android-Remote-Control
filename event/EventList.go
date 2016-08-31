// automatically generated, do not modify

package event

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type EventList struct {
	_tab flatbuffers.Table
}

func GetRootAsEventList(buf []byte, offset flatbuffers.UOffsetT) *EventList {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EventList{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *EventList) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EventList) Event(obj *Event, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
	if obj == nil {
		obj = new(Event)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *EventList) EventLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func EventListStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func EventListAddEvent(builder *flatbuffers.Builder, event flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(event), 0) }
func EventListStartEventVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func EventListEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
