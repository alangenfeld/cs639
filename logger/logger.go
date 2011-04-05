package logger

import(
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

func Init(Filename string){
	//open the file we've been told
	var err os.Error
	logFile, err = os.Open(Filename, os.O_CREAT | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("logger: unable to init log file: " + err.String());
	}
	//initialize task id numbers
	currTaskId = 0;
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

func End(thisTask TaskId) {
	//get the end time before anything else
	thisEndTime := time.Nanoseconds()
	info, present := taskMap[thisTask]
	if !present{
		log.Fatal("logger: taskId is not in use");
	}
	info.EndTime = thisEndTime
	taskMap[thisTask] = info
	logFile.WriteString(String(thisTask))
}

func String(thisTask TaskId) string {
	var ret string;
	info := taskMap[thisTask]
	timeSpent := info.EndTime - info.StartTime
	ret = info.TaskName + ": " + fmt.Sprintf("%u", timeSpent) + "\n"
	return ret
}
	
