// automatically generated, do not modify

package event

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type KeyEvent struct {
	_tab flatbuffers.Table
}

func (rcv *KeyEvent) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KeyEvent) KeyEvent() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func KeyEventStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func KeyEventAddKeyEvent(builder *flatbuffers.Builder, keyEvent flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(keyEvent), 0) }
func KeyEventEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
