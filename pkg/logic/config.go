package logic

import (
	"encoding/json"
	log "github.com/q191201771/naza/pkg/nazalog"
	"github.com/q191201771/naza/pkg/nazajson"
	"io/ioutil"
)

type Config struct {
	RTMP  RTMP       `json:"rtmp"`
	Log   log.Option `json:"log"`
	PProf PProf      `json:"pprof"`

	// v1.0.0之前不提供
	SubIdleTimeout int64   `json:"sub_idle_timeout"`
	GOPCacheNum    int     `json:"gop_cache_number"`
	HTTPFlv        HTTPFlv `json:"httpflv"`
	Pull           Pull    `json:"pull"`
}

type RTMP struct {
	Addr string `json:"addr"`
}

type PProf struct {
	Addr string `json:"addr"`
}

type HTTPFlv struct {
	SubListenAddr string `json:"sub_listen_addr"`
}

type Pull struct {
	Type                      string `json:"type"`
	Addr                      string `json:"addr"`
	ConnectTimeout            int64  `json:"connect_timeout"`
	ReadTimeout               int64  `json:"read_timeout"`
	StopPullWhileNoSubTimeout int64  `json:"stop_pull_while_no_sub_timeout"`
}

func LoadConf(confFile string) (*Config, error) {
	var config Config
	rawContent, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(rawContent, &config); err != nil {
		return nil, err
	}

	// TODO chef: check item valid.
	j, err := nazajson.New(rawContent)
	if err != nil {
		return nil, err
	}
	if !j.Exist("log.level") {
		config.Log.Level = log.LevelDebug
	}
	if !j.Exist("log.filename") {
		config.Log.Filename = "./logs/lal.log"
	}
	if !j.Exist("log.is_to_stdout") {
		config.Log.IsToStdout = true
	}
	if !j.Exist("log.is_rotate_daily") {
		config.Log.IsRotateDaily = true
	}
	if !j.Exist("log.short_file_flag") {
		config.Log.ShortFileFlag = true
	}

	return &config, nil
}