/*
 * Copyright (c) 2019. 陈金龙.
 */

package remote

import (
	"testing"

	"github.com/jinlongchen/viper"
)

func TestRemotePrecedence2(t *testing.T) {
	v := viper.New()
	v.SetConfigType("toml")
	err := v.AddRemoteProvider("etcd", "http://192.168.2.42:2379", "/configs/a.toml")
	if err != nil {
		panic(err)
	}
	err = v.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
}
