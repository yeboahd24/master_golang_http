"Accept interfaces, return structs" 
means: when your function takes input, accept an interface. When your function returns output, return a concrete struct.


// ❌ WRONG: Accepting a concrete type
func SaveLog(logger TerminalLogger, message string) {
    logger.Log(message)
}
// This ONLY works with TerminalLogger. Locked in. Can't test easily.

// ✅ RIGHT: Accepting an interface
func SaveLog(logger Logger, message string) {
    logger.Log(message)
}
// This works with ANY Logger — Terminal, File, Database, or a fake one in tests.

// ❌ WRONG: Returning an interface
func NewLogger() Logger {
    return TerminalLogger{}
}
// The caller gets a Logger interface back.
// They can't access any TerminalLogger-specific fields or methods.
// It hides what you're actually giving them.

// ✅ RIGHT: Returning a concrete struct
func NewTerminalLogger() TerminalLogger {
    return TerminalLogger{}
}
// The caller knows exactly what they got.
// They can still pass it to any function that accepts Logger,
// because TerminalLogger satisfies the Logger interface.
