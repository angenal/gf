package gf

import (
	"sync"

	"github.com/valyala/fasthttp"
)

// Version of current fiber package
const Version = "1.0.0"

// App denotes the Fiber application.
type App struct {
	mutex sync.Mutex
	// Route stack divided by HTTP methods
	stack [][]*Route
	// Route stack divided by HTTP methods and route prefixes
	treeStack []map[string][]*Route
	// contains the information if the route stack has been changed to build the optimized tree
	routesRefreshed bool
	// Amount of registered routes
	routesCount uint32
	// Amount of registered handlers
	handlersCount uint32
	// Ctx pool
	pool sync.Pool
	// Fast http server instance
	server *fasthttp.Server
	// App config instance
	config Config
	// Converts string to a byte slice
	getBytes func(s string) (b []byte)
	// Converts byte slice to a string
	getString func(b []byte) string
	// mount prefix -> error handler
	errorHandlers map[string]ErrorHandler
}
