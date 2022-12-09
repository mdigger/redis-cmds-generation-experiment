package cmd

type ctags uint16

const (
	ctagNull     = ctags(0)                    // default tags
	ctagOptIn    = ctags(1 << 15)              // client side caching opt-int command
	ctagBlock    = ctags(1 << 14)              // blocking command which needs to be process by dedicated connection
	ctagReadOnly = ctags(1 << 13)              // readonly command and can be retried when network error
	ctagNoRet    = ctags(1<<12) | ctagReadOnly // one of the SUBSCRIBE, PSUBSCRIBE, UNSUBSCRIBE or PUNSUBSCRIBE commands
	ctagMtGet    = ctags(1<<11) | ctagReadOnly // the command is MGET
)

// IsOptIn checks if it is client side caching opt-int command.
func (c ctags) IsOptIn() bool {
	return c&ctagOptIn == ctagOptIn
}

// IsBlock checks if it is blocking command which needs to be process by dedicated connection.
func (c ctags) IsBlock() bool {
	return c&ctagBlock == ctagBlock
}

// NoReply checks if it is one of the SUBSCRIBE, PSUBSCRIBE, UNSUBSCRIBE or PUNSUBSCRIBE commands.
func (c ctags) NoReply() bool {
	return c&ctagNoRet == ctagNoRet
}

// IsReadOnly checks if it is readonly command and can be retried when network error.
func (c ctags) IsReadOnly() bool {
	return c&ctagReadOnly == ctagReadOnly
}

// IsMGet returns if the command is MGET
func (c ctags) IsMGet() bool {
	return c == ctagMtGet
}

// toBlock marks the command with blockTag
func (c *ctags) toBlock() {
	*c |= ctagBlock
}
