// automatically generated, do not modify

package event

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Swipe struct {
	_tab flatbuffers.Table
}

func (rcv *Swipe) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Swipe) FromX() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Swipe) FromY() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Swipe) ToX() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Swipe) ToY() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func SwipeStart(builder *flatbuffers.Builder) { builder.StartObject(4) }
func SwipeAddFromX(builder *flatbuffers.Builder, fromX int32) { builder.PrependInt32Slot(0, fromX, 0) }
func SwipeAddFromY(builder *flatbuffers.Builder, fromY int32) { builder.PrependInt32Slot(1, fromY, 0) }
func SwipeAddToX(builder *flatbuffers.Builder, toX int32) { builder.PrependInt32Slot(2, toX, 0) }
func SwipeAddToY(builder *flatbuffers.Builder, toY int32) { builder.PrependInt32Slot(3, toY, 0) }
func SwipeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
