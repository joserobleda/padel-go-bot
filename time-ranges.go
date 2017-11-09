package main

import (
	"time"
)

var exact, _ = time.ParseDuration("0m")
var halfHourBefore, _ = time.ParseDuration("-30m")
var halfHourAfter, _ = time.ParseDuration("30m")

var TimesRanges = []time.Duration{ exact, halfHourBefore, halfHourAfter }
