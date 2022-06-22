package Log

import (
	"fmt"
	"log"
	"os"
)

func PrintLog(msg ...any) {
	fileLog, _ := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer fileLog.Close()

	mylog := log.New(fileLog, "INFO: ", log.LstdFlags|log.Lshortfile)

	mylog.Println(msg)
	fmt.Print(msg)
}

// func main() {
// 	fileLog, _ := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	defer fileLog.Close()

// 	mylog := log.New(fileLog, "INFO: ", log.LstdFlags|log.Lshortfile)
// 	//추가 실행할 문장이 있다면 여기에 추가

// 	mylog.Println("End")
// }
