#Golang service handler

#Description
Simple service handler for catching OS signal

#Installation
    go get github.com/MarinX/service
#Notes
* Your process/service must implement OnStart() / OnStop() methods
* service.Run() will call OnStart() and block until kill signal
* After kill signal, OnStop() will be called and program will exit when all service/process are exited


#Example
    // main
    package main

    import (
	    "fmt"
	    "github.com/MarinX/service"
	    "time"
    )

    type Proc1 struct{}
    type Proc2 struct{}

    func main() {

	    proc := service.New()

	    proc.Add(new(Proc1))
	    proc.Add(new(Proc2))

	    //block till kill signal
	    proc.Run()
    }

    func (t *Proc1) OnStart() {
	    fmt.Println("Proc1 started")
    }

    func (t *Proc1) OnStop() {
	    fmt.Println("Stoping proc1")
	    time.Sleep(1 * time.Second)
    }

    func (t *Proc2) OnStart() {
	    fmt.Println("Proc1 started")
    }

    func (t *Proc2) OnStop() {
	    fmt.Println("Stoping proc2")
	    time.Sleep(2 * time.Second)
    }

#License
This library is under the MIT License
#Author
Marin Basic 
