// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package lbslns implements the JSON configuration for the LoRa Basics Station `router_config` message.
package lbslns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/band"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/frequencyplans"
	"go.thethings.network/lorawan-stack/v3/pkg/pfconfig/shared"
)

const (
	configHardwareSpecPrefix = "sx1301"
)

var errFrequencyPlan = errors.DefineInvalidArgument("frequency_plan", "invalid frequency plan `{name}`")

// DataRates encodes the available datarates of the channel plan for the Station in the format below:
// [0] -> SF (Spreading Factor; Range: 7...12 for LoRa, 0 for FSK)
// [1] -> BW (Bandwidth; 125/250/500 for LoRa, ignored for FSK)
// [2] -> DNONLY (Downlink Only; 1 = true, 0 = false)
type DataRates [16][3]int

// lbsSX1301Config contains the configuration for the SX1301 concentrator for the LoRa Basics Station `router_config` message.
// This structure incorporates modifications for the `v1.5` and `v2` concentrator reference.
type lbsSX1301Config struct {
	ClockSource         uint8
	LBTConfig           *shared.LBTConfig
	Radios              []shared.RFConfig
	Channels            []shared.IFConfig
	LoRaStandardChannel *shared.IFConfig
	FSKChannel          *shared.IFConfig
}

// RouterConfig contains the router configuration.
// This message is sent by the Gateway Server.
type RouterConfig struct {
	NetID          []int             `json:"NetID"`
	JoinEUI        [][]int           `json:"JoinEui"`
	Region         string            `json:"region"`
	HardwareSpec   string            `json:"hwspec"`
	FrequencyRange []int             `json:"freq_range"`
	DataRates      DataRates         `json:"DRs"`
	SX1301Config   []lbsSX1301Config `json:"sx1301_conf"`

	// These are debug options to be unset in production gateways.
	NoCCA       bool `json:"nocca"`
	NoDutyCycle bool `json:"nodc"`
	NoDwellTime bool `json:"nodwell"`

	MuxTime float64 `json:"MuxTime"`
}

// MarshalJSON implements json.Marshaler.
func (conf RouterConfig) MarshalJSON() ([]byte, error) {
	type Alias RouterConfig
	return json.Marshal(struct {
		Type string `json:"msgtype"`
		Alias
	}{
		Type:  "router_config",
		Alias: Alias(conf),
	})
}

type kv struct {
	key   string
	value interface{}
}

type orderedMap struct {
	kv []kv
}

func (m *orderedMap) add(k string, v interface{}) {
	m.kv = append(m.kv, kv{key: k, value: v})
}

func (m orderedMap) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")
	for i, kv := range m.kv {
		if i != 0 {
			b.WriteString(",")
		}
		key, err := json.Marshal(kv.key)
		if err != nil {
			return nil, err
		}
		b.Write(key)
		b.WriteString(":")
		val, err := json.Marshal(kv.value)
		if err != nil {
			return nil, err
		}
		b.Write(val)
	}
	b.WriteString("}")
	return b.Bytes(), nil
}

// MarshalJSON implements json.Marshaler.
func (c lbsSX1301Config) MarshalJSON() ([]byte, error) {
	var m orderedMap
	// m.add("lorawan_public", c.LoRaWANPublic)
	m.add("clksrc", c.ClockSource)
	if c.LBTConfig != nil {
		m.add("lbt_cfg", *c.LBTConfig)
	}
	for i, radio := range c.Radios {
		m.add(fmt.Sprintf("radio_%d", i), radio)
	}
	for i, channel := range c.Channels {
		m.add(fmt.Sprintf("chan_multiSF_%d", i), channel)
	}
	if c.LoRaStandardChannel != nil {
		m.add("chan_Lora_std", *c.LoRaStandardChannel)
	}
	if c.FSKChannel != nil {
		m.add("chan_FSK", *c.FSKChannel)
	}
	return json.Marshal(m)
}

// GetRouterConfig returns the routerconfig message to be sent to the gateway.
// Currently as per the basic station docs, all frequency plans have to be from the same region (band) https://doc.sm.tc/station/tcproto.html#router-config-message.
func GetRouterConfig(bandID string, fps map[string]*frequencyplans.FrequencyPlan, isProd bool, dlTime time.Time) (RouterConfig, error) {
	for _, fp := range fps {
		if err := fp.Validate(); err != nil {
			return RouterConfig{}, errFrequencyPlan.New()
		}
	}
	conf := RouterConfig{}
	conf.JoinEUI = nil
	conf.NetID = nil

	phy, err := band.GetByID(bandID)
	if err != nil {
		return RouterConfig{}, errFrequencyPlan.New()
	}
	s := strings.Split(phy.ID, "_")
	if len(s) < 2 {
		return RouterConfig{}, errFrequencyPlan.New()
	}
	conf.Region = fmt.Sprintf("%s%s", s[0], s[1])

	min, max, err := getMinMaxFrequencies(fps)
	conf.FrequencyRange = []int{
		int(min),
		int(max),
	}

	conf.HardwareSpec = fmt.Sprintf("%s/%d", configHardwareSpecPrefix, len(fps))

	drs, err := getDataRatesFromBandID(bandID)
	if err != nil {
		return RouterConfig{}, errFrequencyPlan.New()
	}
	conf.DataRates = drs

	conf.NoCCA = !isProd
	conf.NoDutyCycle = !isProd
	conf.NoDwellTime = !isProd

	for _, fp := range fps {
		sx1301Config, err := shared.BuildSX1301Config(fp)
		if err != nil {
			return RouterConfig{}, err
		}
		lbsSX1301Config := lbsSX1301Config{
			ClockSource:         sx1301Config.ClockSource,
			LBTConfig:           sx1301Config.LBTConfig,
			LoRaStandardChannel: sx1301Config.LoRaStandardChannel,
			FSKChannel:          sx1301Config.FSKChannel,
		}
		for _, channel := range sx1301Config.Channels {
			lbsSX1301Config.Channels = append(lbsSX1301Config.Channels, channel)
		}
		for i, radio := range sx1301Config.Radios {
			if i == 0 {
				radio.TxFreqMin = 0
				radio.TxFreqMax = 0
			}
			radio.Type = ""
			lbsSX1301Config.Radios = append(lbsSX1301Config.Radios, radio)
		}
		conf.SX1301Config = append(conf.SX1301Config, lbsSX1301Config)
	}

	// Add the MuxTime for RTT measurement.
	conf.MuxTime = float64(dlTime.Unix()) + float64(dlTime.Nanosecond())/(1e9)

	return conf, nil
}

// getDataRatesFromBandID parses the available data rates from the band into DataRates.
func getDataRatesFromBandID(id string) (DataRates, error) {
	phy, err := band.GetByID(id)
	if err != nil {
		return DataRates{}, err
	}

	// Set the default values.
	drs := DataRates{}
	for _, dr := range drs {
		dr[0] = -1
		dr[1] = 0
		dr[2] = 0
	}

	for i, dr := range phy.DataRates {
		if loraDR := dr.Rate.GetLoRa(); loraDR != nil {
			loraDR.GetSpreadingFactor()
			drs[i][0] = int(loraDR.GetSpreadingFactor())
			drs[i][1] = int(loraDR.GetBandwidth() / 1000)
		} else if fskDR := dr.Rate.GetFSK(); fskDR != nil {
			drs[i][0] = 0 // must be set to 0 for FSK, the BW field is ignored.
		}
	}
	return drs, nil
}

// getMinMaxFrequencies extract the minimum and maximum frequencies between all the bands.
func getMinMaxFrequencies(fps map[string]*frequencyplans.FrequencyPlan) (uint64, uint64, error) {
	var min, max uint64
	min = math.MaxUint64
	for _, fp := range fps {
		if len(fp.Radios) == 0 {
			return 0, 0, errFrequencyPlan.New()
		}
		if fp.Radios[0].TxConfiguration.MinFrequency < min {
			min = fp.Radios[0].TxConfiguration.MinFrequency
		}
		if fp.Radios[0].TxConfiguration.MaxFrequency > max {
			max = fp.Radios[0].TxConfiguration.MaxFrequency
		}
	}
	return min, max, nil
}
