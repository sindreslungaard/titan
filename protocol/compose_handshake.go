package protocol

func SecureLoginOK() []byte {
	b := EmptyBuffer()
	b.WriteShort(SecureLoginOKHeader)
	return b.Compose()
}
