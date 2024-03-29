package log

import (
	"fmt"
	nlog "log"

	"github.com/revx-official/output/console"
)

const (
	// The constant prefix for trace messages.
	UniformConstantTrace = "TRACE"

	// The constant prefix for debug messages.
	UniformConstantDebug = "DEBUG"

	// The constant prefix for info messages.
	UniformConstantInfo = "INFO "

	// The constant prefix for warn messages.
	UniformConstantWarn = "WARN "

	// The constant prefix for error messages.
	UniformConstantError = "ERROR"

	// The constant prefix for fatal messages.
	UniformConstantFatal = "FATAL"
)

// Description:
//
//	Definition for the trace color.
//	Can be overriden from the outside.
var ConsoleColorTrace = console.NewConsoleColor(console.FgColorMagenta)

// Description:
//
//	Definition for the debug color.
//	Can be overriden from the outside.
var ConsoleColorDebug = console.NewConsoleColor(console.FgColorCyan)

// Description:
//
//	Definition for the info color.
//	Can be overriden from the outside.
var ConsoleColorInfo = console.NewConsoleColor()

// Description:
//
//	Definition for the warn color.
//	Can be overriden from the outside.
var ConsoleColorWarn = console.NewConsoleColor(console.FgColorYellow)

// Description:
//
//	Definition for the error color.
//	Can be overriden from the outside.
var ConsoleColorError = console.NewConsoleColor(console.FgColorRed)

// Description:
//
//	Definition for the fatal color.
//	Can be overriden from the outside.
var ConsoleColorFatal = console.NewConsoleColor(console.FgColorIntenseRed, console.EmphasisReverseVideo)

// Description:
//
//	The global log level.
var Level = LevelTrace

// Description:
//
//	Traces a formatted message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Tracef(format string, args ...interface{}) {
	if Level > LevelTrace {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log := ConsoleColorTrace.Wrap("[%s] %s", UniformConstantTrace, msg)

	nlog.Println(log)
}

// Description:
//
//	Logs a formatted debug message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Debugf(format string, args ...interface{}) {
	if Level > LevelDebug {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log := ConsoleColorDebug.Wrap("[%s] %s", UniformConstantDebug, msg)

	nlog.Println(log)
}

// Description:
//
//	Logs a formatted info message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Infof(format string, args ...interface{}) {
	if Level > LevelInfo {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log := ConsoleColorInfo.Wrap("[%s] %s", UniformConstantInfo, msg)

	nlog.Println(log)
}

// Description:
//
//	Logs a formatted warning message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Warnf(format string, args ...interface{}) {
	if Level > LevelWarn {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log := ConsoleColorWarn.Wrap("[%s] %s", UniformConstantWarn, msg)

	nlog.Println(log)
}

// Description:
//
//	Logs a formatted error message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Errorf(format string, args ...interface{}) {
	if Level > LevelError {
		return
	}

	msg := fmt.Sprintf(format, args...)
	log := ConsoleColorError.Wrap("[%s] %s", UniformConstantError, msg)

	nlog.Println(log)
}

// Description:
//
//	Logs a formatted fatal message.
//	Automatically calls `os.Exit(1)` afterwards.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log := ConsoleColorFatal.Wrap("[%s] %s", UniformConstantFatal, msg)

	nlog.Fatalf("%s\n", log)
}
