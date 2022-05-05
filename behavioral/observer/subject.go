package main

/*主体*/

type subject interface {
	register(Observe observer)
	deregister(Observe observer)
	notifyAll()
}
