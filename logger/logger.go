package logger

import(
	"exec"
	"fmt"
	"log"
	"os"
	"time"
)

type TaskId uint64
type taskInfo struct{
	TaskName string
	StartTime int64
	EndTime int64
}

var logFile *os.File
var taskMap = map[TaskId] taskInfo{}
var currTaskId TaskId
var statusDir string

const STATUS_LEN = 82
const STATUS_CMD = "./stats.sh"
const WHO = "who | awk '{print $1;}'| uniq | wc -l |awk '{print $1;}'"
const WHO_LEN = 2

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
	var info taskInfo;
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
		systemStats()
	}	
	return ""
}

func systemStats() {
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
	if info.EndTime == 0 {
		return ""
	}
	timeSpent := info.EndTime - info.StartTime
	niceTimeSpent := float64(timeSpent) / float64(1000000000)
	timeStamp := time.SecondsToLocalTime(time.Seconds())
	ret = "\n" + timeStamp.String() + ": " + info.TaskName + ": " + fmt.Sprintf("%f", niceTimeSpent) + " seconds\n"
	return ret
}

/********************************************************
 * Load score:
 * +.25 points per user up to 3 pts
 * x/5pts for memory usage, so it would be usage % * 5
 * x/2pts for cpu usage, so it would be available cpu * 2
*********************************************************/

func GetLoad() int {
	result := make([]byte, WHO_LEN)
        args := make([]string, 1)
	command, err := exec.Run(WHO, args, nil, ".", exec.PassThrough, exec.Pipe, exec.PassThrough)
        if err != nil{
                 log.Println("chunk fails in command:" + err.String())
                 log.Fatal("chunk: unable to obtain remote command")
        }
        _,err =command.Stdout.Read(result)
	return 0
}
	
