package main

import "fmt"

// The interface — anything that can Log is a Logger
type Logger interface {
    Log(message string) error
}

// --- Three different ways to log ---

type TerminalLogger struct{}

func (t TerminalLogger) Log(message string) error {
    fmt.Println("[TERMINAL]:", message)
    return nil
}

type FileLogger struct {
    FilePath string
}

func (f FileLogger) Log(message string) error {
    // In real code you'd write to the file here
    fmt.Printf("[FILE %s]: %s\n", f.FilePath, message)
    return nil
}

type DatabaseLogger struct {
    TableName string
}

func (d DatabaseLogger) Log(message string) error {
    // In real code you'd insert into the database here
    fmt.Printf("[DB table '%s']: %s\n", d.TableName, message)
    return nil
}

// --- Your app just depends on Logger, not the specific type ---

type App struct {
    logger Logger
}

func (a App) DoSomething() {
    a.logger.Log("DoSomething was called")
    // ... actual work here
    a.logger.Log("DoSomething finished")
}

func main() {
    // Swap the logger anytime — the App doesn't care
    app := App{logger: TerminalLogger{}}
    app.DoSomething()

    app = App{logger: FileLogger{FilePath: "/var/log/app.log"}}
    app.DoSomething()

    app = App{logger: DatabaseLogger{TableName: "logs"}}
    app.DoSomething()
}

This prints:
```
[TERMINAL]: DoSomething was called
[TERMINAL]: DoSomething finished
[FILE /var/log/app.log]: DoSomething was called
[FILE /var/log/app.log]: DoSomething finished
[DB table 'logs']: DoSomething was called
[DB table 'logs']: DoSomething finished
