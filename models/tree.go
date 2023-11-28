package models

type Key interface {
	Less(Key) bool
	Eq(Key) bool
}
