package main

import (
	"flag"
	"runtime"

	"./server"
)

func main() {
	// Optimize for current number of CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Capture command line arguments
	port := flag.Int("p", 3000, "the port on which to listen")
	caching := flag.Bool("c", false, "enables template caching for better performance")
	flag.Parse()

	// Start the server
	svr := server.New(*port)
	svr.SetTemplateCaching(*caching)
	svr.ListenAndServe()
}
