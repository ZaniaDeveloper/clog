# clog
Customizable Logger based on logger engines:
- [Standard Go library](https://golang.org/pkg/log/)
- [Database](https://golang.org/pkg/database/sql/)
- [Log15](https://github.com/inconshreveable/log15)
- [Logrus](https://github.com/sirupsen/logrus)
- [LogXI](https://github.com/mgutz/logxi)
- [XLog](https://github.com/xfxdev/xlog)
- [Zap](https://github.com/uber-go/zap)

It formats logger lines and can be used in all loger.

```go
package main

import logger "github.com/ZaniaDeveloper/clog"

func NewLogger() Logger {
    aLogger := logger.New()
    aLogger.SetFormater(func(infos map[string]string) string {
        return plog.Format(infos)
    })

    return aLogger
}
```