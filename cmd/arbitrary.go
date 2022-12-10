package cmd

import "strings"

// arbitrary allows user to build an arbitrary redis command with Builder.Arbitrary
type Arbitrary Base

// Keys calculate which key slot the command belongs to.
// Users must use Keys to construct the key part of the command, otherwise
// the command will not be sent to correct redis node.
func (c Arbitrary) Keys(keys ...string) Arbitrary {
	switch c.cslot & cslotNo {
	case cslotNo:
		if len(keys) > 0 {
			c.cslot = cslotNo | getSlot(keys[0])
		}

	default:
		for _, k := range keys {
			c.cslot.set(getSlot(k))
		}
	}

	c.command.append(keys...)

	return c
}

// Args is used to construct non-key parts of the command.
func (c Arbitrary) Args(args ...string) Arbitrary {
	c.command.append(args...)

	return c
}

// Build is used to complete constructing a command
func (c Arbitrary) Build() Completed {
	if c.command.IsEmpty() {
		panic(arbitraryNoCommand)
	}

	if strings.HasSuffix(strings.ToUpper(c.command.get(0)), "SUBSCRIBE") {
		panic(arbitrarySubscribe)
	}

	return completed(c)
}

// Blocking is used to complete constructing a command and mark it as blocking command.
// Blocking command will occupy a connection from a separated connection pool.
func (c Arbitrary) Blocking() Completed {
	c.ctags = ctagBlock

	return c.Build()
}

// ReadOnly is used to complete constructing a command and mark it as readonly command.
// ReadOnly will be retried under network issues.
func (c Arbitrary) ReadOnly() Completed {
	c.ctags = ctagReadOnly

	return c.Build()
}

// MultiGet is used to complete constructing a command and mark it as mtGetTag command.
func (c Arbitrary) MultiGet() Completed {
	if c.command.IsEmpty() {
		panic(arbitraryNoCommand)
	}

	if cmd := c.command.get(0); cmd != "MGET" && cmd != "JSON.MGET" {
		panic(arbitraryMultiGet)
	}

	c.ctags = ctagMtGet

	return c.Build()
}

// // IsZero is used to test if Arbitrary is initialized
// func (c Arbitrary) IsZero() bool {
// 	return c.command == nil
// }

var (
	arbitraryNoCommand = "arbitrary: should be provided with redis command"
	arbitrarySubscribe = "arbitrary: does not support SUBSCRIBE/UNSUBSCRIBE"
	arbitraryMultiGet  = "arbitrary: MultiGet is only valid for MGET and JSON.MGET"
)
