package main

import (
	""
)

func main() {
	srv := new(todo.Server)
	srv.Run("8000")
}

