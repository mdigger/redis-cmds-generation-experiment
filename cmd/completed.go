package cmd

type Completed interface {
	// String return the command string.
	String() string

	// IsEmpty checks if it is an empty command.
	IsEmpty() bool

	// IsOptIn checks if it is client side caching opt-int command.
	IsOptIn() bool

	// IsBlock checks if it is blocking command which needs to be process by dedicated connection.
	IsBlock() bool

	// NoReply checks if it is one of the SUBSCRIBE, PSUBSCRIBE, UNSUBSCRIBE or PUNSUBSCRIBE commands.
	NoReply() bool

	// IsReadOnly checks if it is readonly command and can be retried when network error.
	IsReadOnly() bool

	// IsWrite checks if it is not readonly command.
	IsWrite() bool

	// IsMGet returns if the command is MGET
	IsMGet() bool

	mustBeCompleted() // check interface
}

// completed represents a internal completed Redis command.
type completed Base

func (completed) mustBeCompleted()

var _ Completed = (*completed)(nil)

func NewCompleted(tags ctags, cmds ...string) Completed {
	return completed{
		ctags:   tags,
		command: newCommand(cmds...),
	}
}

// String return the command string.
func (c completed) String() string {
	return c.command.String()
}

// IsEmpty checks if it is an empty command.
func (c completed) IsEmpty() bool {
	return c.command.IsEmpty()
}

// IsOptIn checks if it is client side caching opt-int command.
func (c completed) IsOptIn() bool {
	return c.ctags.IsOptIn()
}

// IsBlock checks if it is blocking command which needs to be process by dedicated connection.
func (c completed) IsBlock() bool {
	return c.ctags.IsBlock()
}

// NoReply checks if it is one of the SUBSCRIBE, PSUBSCRIBE, UNSUBSCRIBE or PUNSUBSCRIBE commands.
func (c completed) NoReply() bool {
	return c.ctags.NoReply()
}

// IsReadOnly checks if it is readonly command and can be retried when network error.
func (c completed) IsReadOnly() bool {
	return c.ctags.IsReadOnly()
}

// IsWrite checks if it is not readonly command.
func (c completed) IsWrite() bool {
	return !c.IsReadOnly()
}

// IsMGet returns if the command is MGET
func (c completed) IsMGet() bool {
	return c.ctags.IsMGet()
}

// NewBlockingCompleted creates an arbitrary blocking Completed command.
func NewBlockingCompleted(cmds ...string) Completed {
	return NewCompleted(ctagBlock, cmds...)
}

// NewReadOnlyCompleted creates an arbitrary readonly Completed command.
func NewReadOnlyCompleted(cmds ...string) Completed {
	return NewCompleted(ctagReadOnly, cmds...)
}

// NewMGetCompleted creates an arbitrary readonly Completed command.
func NewMGetCompleted(cmds ...string) Completed {
	return NewCompleted(ctagMtGet, cmds...)
}

// MGets groups keys by their slot and returns multi MGET commands
func MGets(keys []string) map[cslot]Completed {
	return slotMGets("MGET", keys)
}

// JsonMGets groups keys by their slot and returns multi JSON.MGET commands
func JsonMGets(keys []string, path string) map[cslot]Completed {
	ret := slotMGets("JSON.MGET", keys)
	for _, jsonmget := range ret {
		jsonmget.(*completed).command.append(path)
	}

	return ret
}

// NewMultiCompleted creates multiple arbitrary Completed commands.
func NewMultiCompleted(cmds [][]string) []Completed {
	ret := make([]Completed, len(cmds))
	for i, c := range cmds {
		ret[i] = NewCompleted(ctagNull, c...)
	}

	return ret
}

func slotMGets(cmd string, keys []string) map[cslot]Completed {
	ret := make(map[cslot]Completed, 16)
	for _, key := range keys {
		slot := getSlot(key)
		if cp, ok := ret[slot]; ok {
			cp.(*completed).command.append(key)
			continue
		}

		cmd := NewCompleted(ctagMtGet, cmd, key)
		cmd.(*completed).cslot.set(slot)
		ret[slot] = cmd
	}

	return ret
}
