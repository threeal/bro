package systemctl

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"
)

type StatusInfo struct {
	Active  string
	MainPid string
}

func Status(name string) (StatusInfo, error) {
	var info StatusInfo
	cmd := exec.Command("systemctl", "status", name)
	out, _ := cmd.CombinedOutput()
	activeRegex, err := regexp.Compile("Active:.*")
	if err != nil {
		return info, err
	}
	activeString := activeRegex.FindString(string(out))
	if activeString == "" {
		return info, errors.New("Could not find active info from the output")
	}
	info.Active = strings.Split(activeString, ": ")[1]
	PIDRegex, err := regexp.Compile("Main PID:.*")
	if err != nil {
		return info, err
	}
	mainPIDString := PIDRegex.FindString(string(out))
	if mainPIDString != "" {
		info.MainPid = mainPIDString
	}
	return info, nil
}
