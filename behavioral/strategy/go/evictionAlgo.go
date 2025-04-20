package main

type EvictionAlgo interface {
	evict(*Cache)
}
