package capnp

// SetData sets the underlying buffer
func (seg *Segment) SetData(data []byte) {
	seg.data = data
}

func (msg *Message) InitializeFixed() *Segment {
	msg.firstSeg = Segment{
		id:   0,
		msg:  msg,
		data: nil,
	}

	return &msg.firstSeg
}
