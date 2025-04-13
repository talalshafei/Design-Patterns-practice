package main

type server interface {
	handleREquest(string, string) (int, string)
}
