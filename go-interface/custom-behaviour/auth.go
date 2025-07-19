package auth

import "fmt"

type Permission byte

const (
	Read Permission = iota + 1
	Write
	Admin
)

// String implements fmt.Stringer
func (p Permission) String() string {
	switch p {
	case Read:
		return "Read"
	case Write:
		return "Write"
	case Admin:
		return "Admin"
	}
	return fmt.Sprintf("<Permission: %d>", p)
}
