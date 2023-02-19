package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// create  a map for storing clear funcs
var clear map[string]func()

func init(){
	clear = make(map[string]func())
	clear["linux"] = func ()  {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func ()  {
		cmd  := exec.Command("cmd","/c","cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		
	}
}
// Clear Terminal function
func CallClear(){
	//runtime.GOOS -> linux,windows etc.
	val,ok := clear[runtime.GOOS] 
	if ok{
		// execute function
		val() 
	}else{
		panic("Your platform is unsupported! Couldn't clear terminal Screen :(")
	}
	// Created by: Muhammed Huseynov
}

func PlaySound(soundFile string) error{
	f, err := os.Open(soundFile)
	if err != nil{
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil{
		return err
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate,format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer,beep.Callback(func() {
		done <- true
	})))

	<- done
	return nil
	// Created by: Muhammed Huseynov
}

func Alarm(seconds int){
	timeElapsed := 0
	
	for timeElapsed < seconds {
		// sleep one second
		time.Sleep(1 * time.Second)
		CallClear()
		timeElapsed += 1
		timeLeft  := seconds - timeElapsed
		minLeft := timeLeft / 60
		secondsLeft := timeLeft % 60
		
		fmt.Printf("%12s %02d : %02d\n ------  Focus on your GOAL  ------","",minLeft,secondsLeft)
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		// Created by: Muhammed Huseynov
	}
}



func main() {
	// in seconds
	// 2 hour -> 7200 seconds
	Alarm(10)
	if err := PlaySound("alarm.mp3"); err != nil{
		log.Fatal(err)
	}
}