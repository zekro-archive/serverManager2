package main

import (
	"testing"
	"github.com/zekroTJA/serverManager2/util"
	"github.com/zekroTJA/serverManager2/core"
)

func TestConfig(t *testing.T) {
	conf := util.GetConf()
	util.CreateConf(conf)
}

func TestLogger(t *testing.T) {
	util.LogInfo("test")
	util.LogError("test")
	util.LogWarn("test")
}

func TestScreen(t *testing.T) {
	core.GetRunningScreens()
	core.GetServers("./")
	core.SliceContainsServer(&[]core.Screen {}, &core.Screen {})
	core.StartScreen(&core.Screen {}, &[]core.Screen {}, &util.Conf {}, false)
	core.StopScreen(&core.Screen {}, &[]core.Screen {}, &util.Conf {})
	core.ResumeScreen(&core.Screen {}, &[]core.Screen {}, &util.Conf {})
	core.RestartScreen(&core.Screen {}, &[]core.Screen {}, &util.Conf {}, false)
}