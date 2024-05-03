package cmd

import "context"

// Shared global objects to inject when building commands
// NOTE: Add fields to this struct such as clients, loggers, and io interfaces
type Context struct {
	// Parent context that could be configured to handle timeouts
	// and listen to OS Shutdown signals
	ParentCtx context.Context
}
