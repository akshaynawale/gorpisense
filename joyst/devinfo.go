package joyst

import (
	"github.com/golang/glog"
	"io/ioutil"
	"strings"
)

var FileName = "/proc/bus/input/devices"

type DevInfo struct {
	DeviceName    string // Device name
	EventFileName string // Corresponding event file path
}

func GetInputDeviceInfo() []DevInfo {
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		glog.Errorf("Unable to read file %s. Error: %s", FileName, err)
	}
	DevInfoArray := []DevInfo{}
	info := []string{}
	dataS := string(data)
	dataList := strings.Split(dataS, "\n")
	var devdata string
	for _, line := range dataList {
		line = strings.TrimSpace(line)
		// Ignore all the line empty lines in the file, before first entry
		if line == "" && devdata == "" {
			continue
		}
		if line == "" {
			info = append(info, devdata)
			devdata = ""
		}
		devdata = strings.Join([]string{devdata, line}, "\n")
	}

	for _, d := range info {
		nameline := []string{}
		handline := []string{}
		for _, line := range strings.Split(d, "\n") {
			if strings.Contains(line, "Name=") {
				nameline = strings.Split(line, "Name=")
			}
			if strings.Contains(line, "Handlers=") {
				handline = strings.Split(line, " ")
			}
		}
		if len(nameline) != 0 && len(handline) != 0 {
			DevInfoArray = append(
				DevInfoArray,
				DevInfo{
					DeviceName:    nameline[len(nameline)-1],
					EventFileName: handline[len(handline)-1],
				},
			)
		}
	}
	return DevInfoArray
}
