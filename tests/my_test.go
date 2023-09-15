package test

import (
	"fmt"
	"github.com/spf13/viper"
	"pointSwitch/global"
	_ "pointSwitch/utils"
	"testing"
)

func TestM(t *testing.T) {
	fmt.Println(global.Addrs)
	viper.Set("addrs", []int{10, 4, 5})
	viper.WriteConfig()
}
