package main

type Subject interface {
	register(Observer)
	deregister(Observer)
	notifyAll()
}
