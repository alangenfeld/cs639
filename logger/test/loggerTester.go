package loggerTester

import(
	"fmt"
	"log"
	"../logger"
//	"time"
)

func main () {
	err := logger.Init("test-log.txt", "../")
	if err != nil {
		log.Println("Error in logger.Init: " + err.String())
	}
	val := logger.GetLoad()
	log.Println("Val = " + fmt.Sprintf("%d", val))
	/*totalId := logger.Start("TotalTest");
	id := logger.Start("2 seconds")
	time.Sleep(2000000000)
	errString := logger.End(id, true)
	if errString != "" {
		log.Println("Error in logger.End: " + errString)
	}
	id = logger.Start("4 seconds")
	time.Sleep(4000000000)
	errString = logger.End(id, true)
	if errString != "" {
		log.Println("Error in logger.End: " + errString)
	}
	errString = logger.End(totalId, false)
	if errString != "" {
		log.Println("Error in logger.End: " + errString)
	}*/
}
