package bitfield

type Bitfield interface {
    Set(index uint32, value bool)   error
    Get(index uint32)               (bool, error)
    Length()                        uint32
    AsBytes()                       []byte
}
