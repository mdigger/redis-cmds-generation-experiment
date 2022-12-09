package cmd

//go:generate go run generator.go -path ./spec

// Builder builds commands by reusing CommandSlice from the sync.Pool
type Builder struct {
	slot cslot
}

// NewBuilder creates a Builder and initializes the internal sync.Pool
func NewBuilder(initSlot cslot) Builder {
	return Builder{
		slot: initSlot,
	}
}

// Arbitrary allows user to build an arbitrary redis command by following Arbitrary.Keys and Arbitrary.Args
func (b Builder) Arbitrary(tokens ...string) Arbitrary {
	return Arbitrary{
		command: newCommand(tokens...),
		cslot:   b.slot,
	}
}
