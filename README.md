# plog
Customizable Formater for Loggers

It formats logger lines and can be used in all loger.

```go
package main

import "github.com/ZaniaDeveloper/plog"

func NewLogger() Logger {
    aLogger := logger.New()
    aLogger.SetFormater(func(infos map[string]string) string {
        return plog.Format(infos)
    })

    return aLogger
}
```