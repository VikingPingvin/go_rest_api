package config

import (
	"time"
)

type Config struct {
	Server serverConf
}

type serverConf struct {
	PORT         int
	TimeoutRead  time.Duration
	TimeoutWrite time.Duration
	TimeoutIdle  time.Duration
}
