package tool 

import (
	"testing"
)

func Test_ScannerPort(t *testing.T) {
	rs := CommTool.ScannerPort("localhost:2379")
	if !rs { 
		t.Error("port not accessable")
	} else {
		t.Log("127.0.0.1:2379 ok")
	}
}