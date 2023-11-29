package imp

import "fmt"

type Redis struct{}

func (r *Redis) Use() {
	fmt.Println("Redis")
}

func (r *Redis) NewDatabase() {
	fmt.Println("Redis")
}
