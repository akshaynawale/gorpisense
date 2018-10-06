package joyst

import (
	"fmt"
	"github.com/golang/glog"
	"strconv"
)

func reverse(input []byte) []byte {
	// Retruns reverse of the byte array
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func convertInt32(input []byte) int32 {
	input = reverse(input)
	inputs := fmt.Sprintf("%x", input)
	out, err := strconv.ParseInt(inputs, 16, 32)
	if err != nil {
		glog.Error(err)
	}
	return int32(out)
}

func convertInt64(input []byte) int64 {
	input = reverse(input)
	inputs := fmt.Sprintf("%x", input)
	out, err := strconv.ParseInt(inputs, 16, 64)
	if err != nil {
		glog.Error(err)
	}
	return out
}
