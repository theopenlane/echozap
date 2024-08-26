[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=theopenlane_echozap&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=theopenlane_echozap)
[![Build status](https://badge.buildkite.com/04d19d3695dfe30fece164674e53c234b8edbfeb78857d0d10.svg)](https://buildkite.com/theopenlane/echozap)

# echo zap

Middleware for Golang [Echo](https://echo.labstack.com/) framework that provides integration with Uber´s [Zap](https://github.com/uber-go/zap)  logging library for logging HTTP requests

## Usage

```go
package main

import (
	"net/http"

	"github.com/theopenlane/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	e := echo.New()

	zapLogger, _ := zap.NewProduction()

	e.Use(echozap.ZapLogger(zapLogger))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

## Logged details

The following information is logged:

*  Status Code
*  Time
*  URI
*  Method
*  Hostname
*  Remote IP Address


