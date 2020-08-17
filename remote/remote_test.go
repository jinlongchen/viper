/*
 * Copyright (c) 2019. 陈金龙.
 */

package remote

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jinlongchen/viper"
)

func TestRemote2(t *testing.T) {
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

	allSettingsData, _ := json.Marshal(v.AllSettings())
	log.Printf("config loaded: %s\n", string(allSettingsData))

	err = v.WatchRemoteConfigOnChannel()
	if err != nil {
		log.Printf("WatchRemoteConfig err: %s\n", err.Error())
	}
	log.Println("WatchRemoteConfig -----")
	v.OnConfigChange(func(e fsnotify.Event) {
		allSettingsData, _ := json.Marshal(v.AllSettings())
		log.Printf("config reloaded: %s\n", string(allSettingsData))
	})

	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			log.Printf("%v\n", v.Get("abc.def"))
		}
	}
}
