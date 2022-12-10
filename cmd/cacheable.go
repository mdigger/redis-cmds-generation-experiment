package cmd

type Cacheable interface {
	// String return the command string.
	String() string

	// CacheKey returns the cache key used by the server-assisted client side caching
	CacheKey() (key, command string)

	// MGetCacheCmd returns the cache command of the MGET singular command
	MGetCacheCmd() string

	// MGetCacheKey returns the cache key of the MGET singular command
	MGetCacheKey(i int) string

	mustBeCacheable() // check interface
}

// Cacheable represents a completed Redis command which supports server-assisted client side caching,
// and it should be created by the Cache() of command builder.
type cacheable Base

func (cacheable) mustBeCacheable() {}

var _ Cacheable = (*cacheable)(nil)

func NewCacheable(b Base) Cacheable {
	return cacheable(b)
}

// String return the command string.
func (c cacheable) String() string {
	return c.command.String()
}

// CacheKey returns the cache key used by the server-assisted client side caching
func (c cacheable) CacheKey() (key, command string) {
	switch commands := c.command.list(); len(commands) {
	case 0, 1:
		return "", ""
	case 2:
		return commands[1], commands[0]
	default:
		return commands[1], joinStrings(commands[0], commands[2:]...)
	}
}

// MGetCacheCmd returns the cache command of the MGET singular command
func (c cacheable) MGetCacheCmd() string {
	if c.command.get(0)[0] == 'J' {
		return joinStrings("JSON.GET", c.command.get(-1))
	}

	return "GET"
}

// MGetCacheKey returns the cache key of the MGET singular command
func (c cacheable) MGetCacheKey(i int) string {
	return c.command.get(i + 1)
}
