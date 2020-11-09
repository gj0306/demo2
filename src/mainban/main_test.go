package mainban

import "testing"

func TestMainBan(t *testing.T)  {
	// 方法1
	defer func() { select {} }()
	//方法2
	defer func() { <-make(chan bool) }()
}
