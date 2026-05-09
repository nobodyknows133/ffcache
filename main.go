package main

import "ffcache/cache"

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	server := NewServer(opts, cache.NewCache())
	server.Start()
}
