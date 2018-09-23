package joyst

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

const (
	LEFT  = 105
	RIGHT = 106
	UP    = 103
	DOWN  = 108
	ENTER = 28
)

type Event struct {
	Timeval syscall.Timeval
	Type    uint16
	Code    uint16 // The code for which button is pressed like UP, DOWN, ENTER, LEFT, or RIGHT
	Value   int32  // The preesed position returns 1 else 0
}

type Joystick struct {
	FilePath    string // name of the input file e,g, /dev/input/event1
	Timevalsize int    // Timeval size of 32 bit system is 8 byte for 64 it is 16 byte We will detect this automatically and store it here.
}

func (j *Joystick) Init() {
	//	Get input file name
	j.FilePath = inputFilePath()
}

func decodeCode(code string) int {
	codeInt, err := strconv.Atoi(code)
	if err != nil {
		glog.Error(err)
		return 0
	}
	return codeInt
}

func inputFilePath() string {
	DevInfoArray := GetInputDeviceInfo()
	var filename string
	for _, devinfo := range DevInfoArray {
		if strings.Contains(devinfo.DeviceName, "Raspberry Pi Sense HAT Joystick") {
			filename = devinfo.EventFileName
		}
	}
	return strings.Join([]string{"/dev/input/", filename}, "")
}

func (j *Joystick) Poll(echan chan<- Event) {
	// for glog flags
	flag.Parse()

	file, err := os.Open(j.FilePath)
	if err != nil {
		glog.Fatal(err)
	}
	event := Event{}
	esize := int(unsafe.Sizeof(event))
	buf := make([]byte, esize)

	for {
		_, err := file.Read(buf)
		if err != nil {
			glog.Error(err)
		}
		event.parseEvent(buf)

		if event.Value == 1 {
			fmt.Printf("%s", event)
			echan <- event
		}
		/*
			timevalsec := buf[0:4]
			//timevalusec := buf[4:8]
			//typev := buf[8:10]
			codeByte := buf[10:12]
			valueByte := buf[12:16]
			//fmt.Printf("timeval : %x , typeval: %x ,code: %x , value: %x \n", timeval, typev, code, value)
			//code_int64, _ := binary.Varint(code)
			codestr := fmt.Sprintf("%d", codeByte[0])
			valuestr := fmt.Sprintf("%d", valueByte[0])
			//fmt.Println(code_str)
			code := decodeCode(codestr)
			value := decodeCode(valuestr)
			if code != 0 && value == 1 {
				cmdcode <- code
				//fmt.Printf("code : %d value: %d", code, value)
				fmt.Printf("Timeval sec : %d", timevalsec)
			}
		*/
	}
}

func (e *Event) parseEvent(eventbyte []byte) {

	timelength := unsafe.Sizeof(e.Timeval) / 2
	epochsec := eventbyte[0:timelength]
	epochusec := eventbyte[timelength : timelength*2]
	cur := timelength * 2
	etype := eventbyte[cur : cur+2]
	ecode := eventbyte[cur+2 : cur+4]
	evalue := eventbyte[cur+4 : cur+8]
	//if timelength == 4 {
	e.Timeval.Sec = convertInt32(epochsec)
	e.Timeval.Usec = convertInt32(epochusec)
	/*
		} else {
			e.Timeval.Sec = convertInt64(epochsec)
			e.Timeval.Usec = convertInt64(epochusec)
		}
	*/
	e.Type = uint16(convertInt64(etype))
	e.Code = uint16(convertInt64(ecode))
	e.Value = convertInt32(evalue)
}
