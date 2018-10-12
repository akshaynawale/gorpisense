# gorpisense

Library to use Raspberry Pi Sense Hat. Written in Golang. 

# Basic Usage:

Installation: 
```bash
go get github.com/akshaynawale/gorpisense/joyst
```
Import the library in your program file with:
```go
import "github.com/akshaynawale/gorpisense/joyst"
```

Create a channel of type Event which will be written by the lib function whenever any event happens.
```go
echan := make(chan joyst.Event)
```
Create a joyst struct 
```go
joy := joyst.Joystick{}
```
Call the Poll method on the joyst struct to get the events from joystick in a go routine and pass the event channel (echan) to it.
```go
go joy.Poll(echan)
```
The Poll method will write to the echan, whenever we receive a joystick event. For contineously reading this channel you can
use a for loop. Use the joyst lib constants (i.e. joyst.LEFT etc.) to compare the event.Code values to know which buttton is 
pressed. perform actions based on button pressed.
```go
for {
  event := <-echan
  Switch event.Code {
    case joyst.LEFT:
      // do somthing
      fmt.Println("LEFT Pressed")
    case joyst.RIGHT:
      // do something 
      fmt.Println("RIGHT Pressed")
    case joyst.DOWN:
      // do something
      fmt.Println("DOWN Pressed")
    case joyst.UP:
      // do something
      fmt.Println("UP Pressed")
    case joyst.Enter:
      // do something
      fmt.Println("Enter Pressed")
  }
}
```
More libs for controlling temperature, gyroscope and LED comming soon... :)

Happy Coding...
