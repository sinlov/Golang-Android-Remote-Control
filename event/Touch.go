// automatically generated, do not modify

package event

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Touch struct {
	_tab flatbuffers.Table
}

func (rcv *Touch) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Touch) X() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Touch) Y() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func TouchStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func TouchAddX(builder *flatbuffers.Builder, x int32) { builder.PrependInt32Slot(0, x, 0) }
func TouchAddY(builder *flatbuffers.Builder, y int32) { builder.PrependInt32Slot(1, y, 0) }
func TouchEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
