// automatically generated, do not modify

package event

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Text struct {
	_tab flatbuffers.Table
}

func (rcv *Text) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Text) Text() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TextStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func TextAddText(builder *flatbuffers.Builder, text flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(text), 0) }
func TextEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
