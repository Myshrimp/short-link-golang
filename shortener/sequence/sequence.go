package sequence

// Sequence is an interface that defines the method to get the next sequence number
type Sequence interface {
	Next() (seq uint64, err error)
}