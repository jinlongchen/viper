/*
 * Copyright (c) 2019. 陈金龙.
 */

// Package remote integrates the remote features of Viper.
package viper

import (
	"bytes"
	"io"
	"os"

	crypt "github.com/xordataexchange/crypt/config"
)

type remoteConfigProvider struct{}

func (rc remoteConfigProvider) Get(rp RemoteProvider) (io.Reader, error) {
	cm, err := getConfigManager(rp)
	if err != nil {
		return nil, err
	}
	b, err := cm.Get(rp.Path())
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (rc remoteConfigProvider) Watch(rp RemoteProvider) (io.Reader, error) {
	cm, err := getConfigManager(rp)
	if err != nil {
		return nil, err
	}
	resp, err := cm.Get(rp.Path())
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(resp), nil
}

func (rc remoteConfigProvider) WatchChannel(rp RemoteProvider) (<-chan *RemoteResponse, chan bool) {
	cm, err := getConfigManager(rp)
	if err != nil {
		return nil, nil
	}
	quit := make(chan bool)
	quitwc := make(chan bool)
	viperResponsCh := make(chan *RemoteResponse)
	cryptoResponseCh := cm.Watch(rp.Path(), quit)
	// need this function to convert the Channel response form crypt.Response to Response
	go func(cr <-chan *crypt.Response, vr chan<- *RemoteResponse, quitwc <-chan bool, quit chan<- bool) {
		for {
			select {
			case <-quitwc:
				quit <- true
				return
			case resp := <-cr:
				vr <- &RemoteResponse{
					Error: resp.Error,
					Value: resp.Value,
				}

			}

		}
	}(cryptoResponseCh, viperResponsCh, quitwc, quit)

	return viperResponsCh, quitwc
}

func getConfigManager(rp RemoteProvider) (crypt.ConfigManager, error) {
	var cm crypt.ConfigManager
	var err error

	if rp.SecretKeyring() != "" {
		kr, err := os.Open(rp.SecretKeyring())
		defer kr.Close()
		if err != nil {
			return nil, err
		}
		if rp.Provider() == "etcd" {
			cm, err = crypt.NewEtcdConfigManager([]string{rp.Endpoint()}, kr)
		} else {
			cm, err = crypt.NewConsulConfigManager([]string{rp.Endpoint()}, kr)
		}
	} else {
		if rp.Provider() == "etcd" {
			cm, err = crypt.NewStandardEtcdConfigManager([]string{rp.Endpoint()})
		} else {
			cm, err = crypt.NewStandardConsulConfigManager([]string{rp.Endpoint()})
		}
	}
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func init() {
	RemoteConfig = &remoteConfigProvider{}
}
