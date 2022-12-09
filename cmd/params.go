package cmd

import (
	"strings"
	"sync"
)

// command is the command container managed by the sync.Pool.
type command struct {
	params []string
}

func newCommand(strs ...string) *command {
	if len(strs) == 0 {
		return emptyCommand()
	}

	return &command{params: strs}
}

// String return the command string.
func (c command) String() string {
	return strings.Join(c.params, " ")
}

// IsEmpty checks if it is an empty command.
func (c command) IsEmpty() bool {
	return len(c.params) == 0 || c.get(0) == ""
}

// slices pool
var pool = &sync.Pool{
	New: func() any {
		return &command{
			params: make([]string, 0, 2),
		}
	},
}

func emptyCommand() *command {
	return pool.Get().(*command)
}

// free recycles the Commands.
func (c *command) free() {
	c.clear()
	pool.Put(c)
}

func (c *command) clear() {
	c.params = c.params[:0]
}

// list returns the commands as []string.
func (c command) list() []string {
	return c.params
}

func (c *command) append(keys ...string) {
	c.params = append(c.params, keys...)
}

func (c command) len() int {
	return len(c.params)
}

func (c command) get(i int) string {
	if i < 0 {
		i += len(c.params)
	}

	if i < 0 || i >= len(c.params) {
		return ""
	}

	return c.params[i]
}
