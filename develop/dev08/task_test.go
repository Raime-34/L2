package main

import (
	"testing"
)

func Test_cd(t *testing.T) {
	type args struct {
		path string
	}

	cd("B:\\Go")

	if pwd() != "B:\\Go" {
		t.Errorf("shite")
	}

}

func Test_pwd(t *testing.T) {
	if pwd() != "b:\\Go\\L2\\develop\\dev08" {
		t.Errorf("shite")
	}
}

func Test_ps(t *testing.T) {

	if len(ps()) == 0 {
		t.Errorf("shite")
	}

}

func Test_kill(t *testing.T) {
	type args struct {
		pid int
	}

	err := kill(ps()["steam.exe"])

	if err != nil {
		t.Errorf("shite")
	}

}
