package main 

import (
	"ethos/syscall"
	"ethos/ethos"
        "ethos/efmt"
	"log"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	fd, status := ethos.OpenDirectoryPath(path)
	if status != syscall.StatusOk {
		log.Fatalf ("Error opening %v: %v\n", path, status)
	}


	data    := MyType {5, 6, 7, 8 }

	data.Write(fd)
	data.WriteVar(path +"foobar")
        efmt.Println("val:",data.x1,data.y1,data.x2,data.y2)
        data.x1 = 2*data.x1+1
        data.y1 = 2*data.y1+1
        data.x2 = 2*data.x2+1
        data.y2 = 2*data.y2+1
        efmt.Println("val:",data.x1,data.y1,data.x2,data.y2)

}
