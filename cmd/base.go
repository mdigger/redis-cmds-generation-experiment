package cmd

// Base represents a Base Redis command from builder.
type Base struct {
	command *command
	ctags   ctags
	cslot   cslot
}

// String return the command string.
func (c Base) String() string {
	return c.command.String()
}

// func NewBase(tags ctags, ss ...string) Base {
// 	return Base{
// 		command: newCommand(ss...),
// 		ctags:   tags,
// 	}
// }

// // FIXME: maybe something should be Completed
// var (
// 	// OptInCmd is predefined CLIENT CACHING YES
// 	OptInCmd = NewBase(ctagOptIn, "CLIENT", "CACHING", "YES")
// 	// MultiCmd is predefined MULTI
// 	MultiCmd = NewBase(ctagNull, "MULTI")
// 	// ExecCmd is predefined EXEC
// 	ExecCmd = NewBase(ctagNull, "EXEC")
// 	// RoleCmd is predefined ROLE
// 	RoleCmd = NewBase(ctagNull, "ROLE")
// 	// QuitCmd is predefined QUIT
// 	QuitCmd = NewBase(ctagNull, "QUIT")
// 	// UnsubscribeCmd is predefined UNSUBSCRIBE
// 	UnsubscribeCmd = NewBase(ctagNoRet, "UNSUBSCRIBE")
// 	// PUnsubscribeCmd is predefined PUNSUBSCRIBE
// 	PUnsubscribeCmd = NewBase(ctagNoRet, "PUNSUBSCRIBE")
// 	// SUnsubscribeCmd is predefined SUNSUBSCRIBE
// 	SUnsubscribeCmd = NewBase(ctagNoRet, "SUNSUBSCRIBE")
// 	// PingCmd is predefined PING
// 	PingCmd = NewBase(ctagNull, "PING")
// 	// SlotCmd is predefined CLUSTER SLOTS
// 	SlotCmd = NewBase(ctagNull, "CLUSTER", "SLOTS")
// 	// AskingCmd is predefined CLUSTER ASKING
// 	AskingCmd = NewBase(ctagNull, "ASKING")
// 	// SentinelSubscribe is predefined SUBSCRIBE ASKING
// 	SentinelSubscribe = NewBase(ctagNoRet, "SUBSCRIBE", "+sentinel", "+switch-master", "+reboot")
// )
