package logger

import(
	"exec"
	"fmt"
	"log"
	"os"
	"time"
)

type TaskId uint64
type TaskInfo struct{
	TaskName string
	StartTime int64
	EndTime int64
}

var logFile *os.File
var taskMap = map[TaskId] TaskInfo{}
var currTaskId TaskId
var statusDir string

const STATUS_LEN = 82
const STATUS_CMD = "./stats.sh"

func Init(Filename string, Directory string) os.Error{
	//open the file we've been told
	var err os.Error
	logFile, err = os.Open(Filename, os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("logger: unable to init log file: " + err.String());
		return err;
	}
	statusDir = Directory
	//initialize task id numbers
	currTaskId = 0;
	return nil;
}

func Start(TaskName string) TaskId{
	//get the start time before anything else, for consistency sake
	thisStartTime := time.Nanoseconds()
	var info TaskInfo;
	info.StartTime = thisStartTime
	info.TaskName = TaskName
	taskMap[currTaskId] = info
	//make sure you don't overwrite another task
	currTaskId ++
	return (currTaskId - 1)
}

func End(thisTask TaskId, SysStats bool) string{
	//get the end time before anything else
	thisEndTime := time.Nanoseconds()
	info, present := taskMap[thisTask]
	if !present{
		return "logger: taskId is not in use"
	}
	info.EndTime = thisEndTime
	taskMap[thisTask] = info
	_, err := logFile.WriteString(String(thisTask))
	if err != nil {
		return err.String()
	}
	if SysStats {
		SystemStats()
	}	
	return ""
}

func SystemStats() {
        args := make([]string, 1)
        var result []byte
        result = make([]byte, STATUS_LEN)
        command, err := exec.Run(STATUS_CMD, args, nil, statusDir, exec.PassThrough, exec.Pipe, exec.PassThrough)
        if err != nil{
                 log.Println("chunk fails in command:" + err.String())
                 log.Fatal("chunk: unable to obtain remote command")
        }
	err = nil
	
        time.Sleep(2100000000)
        _,err =command.Stdout.Read(result)
	if err != nil{
        	log.Println("chunk fails read from command: " + err.String())
	}
	logFile.Write(result)
}

func String(thisTask TaskId) string {
	var ret string;
	info := taskMap[thisTask]
	timeSpent := info.EndTime - info.StartTime
	niceTimeSpent := float64(timeSpent) / float64(1000000000)
	timeStamp := time.SecondsToLocalTime(time.Seconds())
	ret = timeStamp.String() + ": " + info.TaskName + ": " + fmt.Sprintf("%f", niceTimeSpent) + " seconds\n"
	return ret
}
	
