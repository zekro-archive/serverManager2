package main

import (
	"testing"
	"fmt"
	t "time"
	u "github.com/zekroTJA/serverManager2/util"
	c "github.com/zekroTJA/serverManager2/core"
)

var conf *u.Conf

func TestConfig(t *testing.T) {
	conf = u.GetConf("./testconf.json")
	u.CreateConf(conf)
	fmt.Println()
	u.LogInfo(fmt.Sprintf("Config successfully created: %+v", conf))
	conf.ServerLocation = "./testservers"
	conf.Logging = 1
	u.LogInfo(fmt.Sprintf("Changed config for test:     %+v", conf))
}

// func TestLogger(t *testing.T) {
// 	u.LogInfo("test")
// 	u.LogError("test")
// 	u.LogWarn("test")
// }


func TestScreen(t *testing.T) {
	c.LogLocation = "./logs/"
	u.LogInfo(fmt.Sprintf("Changed log location to:     %+v", c.LogLocation))

	u.LogInfo("screen#GetServers:")
	fmt.Println(c.GetServers("./testservers"))

	testscreen := &c.Screen {0, "", "test1", ""}

	c.StartScreen(
		testscreen,
		&[]c.Screen {},
		conf, 
		true)
	u.LogInfo("Started testserver")

	c.RestartScreen(
		testscreen,
		c.GetRunningScreens(),
		conf,
		true)
	u.LogInfo("Restarted testserver")

	c.ResumeScreen(
		testscreen,
		c.GetRunningScreens(),
		conf)
	u.LogInfo("Resumed testserver")
	u.LogWarn("While testing screen can not be attached to current terminal.")

	c.StopScreen(
		testscreen,
		c.GetRunningScreens(),
		conf)
	u.LogInfo("Stopped testserver")
}

func TestGetSrceenBenchmark(test *testing.T) {
	runs := 10
	times := make(chan t.Duration, runs)

	go func() {
		for i := 0; i < runs; i++ {
			start := t.Now()
			c.GetServers("./testservers")
			times <- t.Since(start)
		}
		close(times)
	}()

	var max, min, sum int64
	for time := range times {
		itime := int64(time)
		if max == 0 || itime > max {
			max = itime
		}
		if min == 0 || itime < min {
			min = itime
		}
		sum += itime
	}
	u.LogInfo(fmt.Sprintf(
`screen#GetServers() benchmark results:
  | Runs:                %d
  | Servers:             100
  | Result:   | Min:     %d microseconds
              | Max:     %d microseconds
              | Average: %d microseconds`,
	runs, 
	min / 1000, 
	max / 1000, 
	sum / int64(runs) / 1000))
}