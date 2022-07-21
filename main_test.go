package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func setup() {
	fmt.Println("Start test..")
}

func Test_main(t *testing.T){
	assert.Equal(t,"jj", test(),"字元錯誤")
}