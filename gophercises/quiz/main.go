package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct{
	question  string
	answer  string
}

func parseLines(lines [][]string)[]problem{
	ret:=make([]problem,len(lines))
	for k,v :=range(lines){
		ret[k]=problem{
			question:v[0],
			answer:v[1],
		}
	}
	return ret;
}

func main(){
csvFilename:= flag.String("csv","problem.csv","a csv file in format [question,answer]")
timeLimit:=flag.Int("limit",30,"the limit for the quiz in seconds")
shuffle:=flag.Bool("shuffle",false,"shuffle the problems")
flag.Parse()

file,err:=os.Open(*csvFilename)
if err!=nil{
	fmt.Println("Error while openning file: "+*csvFilename)
	return
}

r:=csv.NewReader(file)
lines,err:=r.ReadAll()

if err!=nil{
	fmt.Println("Error while reading file: "+*csvFilename)
	return
}
problems:=parseLines(lines)
if(*shuffle){
	rand.Seed(time.Now().UnixNano())
rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
}
timer:=time.NewTimer(time.Duration(*timeLimit)*time.Second)
var counter int
for i,p:= range problems{
	fmt.Printf("Problem number %d: %s =?",i+1,p.question)
	answerCh:=make(chan string)
	go func(){
		var answer string
		fmt.Scanf("%s\n",&answer)
		answerCh<-answer
	}()

	select{
	case <-timer.C:
		fmt.Println("\nYour time is up")
		return
	case answer:=<-answerCh:
		if p.answer==strings.TrimSpace(answer){
			counter++;
			}else{
				fmt.Printf("You answered wrong!!!, Total number of correct answers is %d of %d",counter,len(problems))
				return
			}
	}	
	}
}

