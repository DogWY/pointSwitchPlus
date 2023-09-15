package services

import (
	"github.com/spf13/viper"
	"pointSwitch/global"
	"sort"
	"strconv"
)

func AddAddrs(addrs []string) {
	for _, addr := range addrs {
		addrInt, _ := strconv.Atoi(addr)
		if _, ok := global.States[addrInt]; ok {
			continue
		}
		global.Addrs = append(global.Addrs, addrInt)
		global.States[addrInt] = 0
	}
	sort.Ints(global.Addrs)
	viper.Set("addrs", global.Addrs)
	viper.WriteConfig()
}

func RemoveAddrs(addrs []string) {
	for _, addr := range addrs {
		addrInt, _ := strconv.Atoi(addr)
		if _, ok := global.States[addrInt]; ok {
			idx := sort.SearchInts(global.Addrs, addrInt)
			global.Addrs = append(global.Addrs[:idx], global.Addrs[idx+1:]...)
			delete(global.States, addrInt)
			delete(global.MDis, addrInt)
			delete(global.MTimes, addrInt)
		}

	}
	viper.Set("addrs", global.Addrs)
	viper.WriteConfig()
}
