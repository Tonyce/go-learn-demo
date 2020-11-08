// +build mock

package service

import "fmt"

// PingPong ...
func PingPong(v string) string {
	return fmt.Sprintf("pong mock %s", v)
}
