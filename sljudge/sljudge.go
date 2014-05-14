package sljudge

//
//this package supports the judge of solution
//
import (
	"GoOnlineJudge/config"
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func compare(std_out, user_out string) int {
	if std_out == user_out {
		return config.JudgeAC
	} else {
		return config.JudgeWA
	}
}

func SJudge(lang, timelimt, memroylimt, pid int, code string) (res, s_time, s_mem int) { //多人同时就会奔溃，建议按solution id创建运行目录
	log.Println("run solution")
	res = config.JudgePD

	pwd, err := os.Getwd()
	os.Chdir("H:\\ACM\\" + strconv.Itoa(pid))
	defer os.Chdir(pwd)
	std_in, std_out_f := os.Stdin, os.Stdout
	std_in, err = os.Open("test.in")
	if err != nil {
		log.Println("Open std_in file failed")
		return
	}
	defer std_in.Close()

	//begin read the std_out text
	var std_out string
	std_out_f, err = os.Open("test.out")
	if err != nil {
		log.Println("open std_out file failed")
		return
	} else {
		reader := bufio.NewReader(std_out_f)
		for true {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				std_out = std_out + string(line) //+ "\n"
				break
			} else if err != nil {
				break
			}
			std_out = std_out + string(line) + "\r\n"
		}
	}
	defer std_out_f.Close()
	//end read the std_out text

	//begin write code to file
	fl, err := os.Open("test.cpp")
	for err == nil {
		fl.Close()
		fl, err = os.Open("test.cpp")
	}
	codefile, err := os.Create("test.cpp")
	defer func() {
		codefile.Close()
		os.Remove("test.cpp")
	}()
	_, err = codefile.WriteString(code)
	if err != nil {
		log.Println("source code writing to file failed")
	}
	//end write code to file

	//begin compile source code
	var out bytes.Buffer
	cmd := exec.Command("g++", "-o", "Main.exe", "-g", "-Wall", "test.cpp") //compile
	err = cmd.Run()
	if err != nil {
		log.Println(err, "1111")
		res = config.JudgeCE //compiler error
		return
	}
	defer os.Remove("Main.exe")
	//end compile source code

	target, err := os.Open("Main.exe") //Memroy of the target
	defer target.Close()
	t_info, err := target.Stat()
	if s_mem = int(t_info.Size()) / (1024 * 1024); s_mem > memroylimt {
		res = config.JudgeMLE
		return
	}

	//begin run
	channel := make(chan int)
	//process_id := make(chan int)
	defer close(channel)
	log.Println("Here")
	bcmd := exec.Command("Main.exe")
	bcmd.Stdin = std_in
	bcmd.Stdout = &out
	bcmd.Start()
	go func() {
		t_time := time.Now()
		err := bcmd.Wait()
		s_time = int(time.Since(t_time).Seconds() * 1000)
		if err != nil {
			channel <- config.JudgeRE //runtime error
			return
		}
		channel <- config.JudgeCP
	}()
	//end run

	select {
	case res = <-channel:
		if res == config.JudgeCP {
			user_out := out.String()
			res = compare(std_out, user_out)
		}
	case <-time.After(time.Second * 1):
		bcmd.Process.Kill()
		<-channel
		res = config.JudgeTLE
	}
	return
}
