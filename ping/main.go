package main 

import (
	"flag"
	"fmt"
	"os"
	"time"
	"net"
	"log"
	"bytes"
	"encoding/binary"
	"math"
)

var(
	timeout			int64 
	size			int 
	count			int 
	typ				uint8 = 8
	code			uint8 = 0
	sendCount		int
	successCount	int 
	failCount		int
	minTs			int64 = math.MaxInt32
	maxTs			int64
	totalTs			int64

)

type ICMP struct{
	Type		uint8
	Code		uint8
	CheckSum	uint16
	ID			uint16
	SequnceNum	uint16
}

func main() {
	GetComandArgs()
	desIp := os.Args[len(os.Args)-1]
	conn,err:= net.DialTimeout("ip:icmp",desIp,time.Duration(timeout)*time.Millisecond)
	if err!= nil{
		log.Fatal(err)
		return
	}
	defer conn.Close()
	fmt.Printf("Pinging %s [%s] with %d bytes of data:\n",desIp,conn.RemoteAddr(),size)
	
	for i:=0;i<count;i++{
		sendCount++
		t1 := time.Now()
		icmp  := &ICMP{
			Type :		 typ,
			Code :		 code,
			CheckSum :	 0,
			ID :		 1,
			SequnceNum : 1,
		}

		data := make([]byte,size)
		var buffer bytes.Buffer
		binary.Write(&buffer,binary.BigEndian,icmp)
		buffer.Write(data)
		data = buffer.Bytes()
		checkSum := checkSum(data)
		data[2] = byte(checkSum >> 8)
		data[3] = byte(checkSum)
		conn.SetDeadline(time.Now().Add(time.Duration(timeout)*time.Millisecond))
		n,err := conn.Write(data)
		if err != nil{
			failCount++
			log.Println(err)
			continue
		}
		buf := make([]byte,65535)
		n,err = conn.Read(buf)
		if err !=nil{
			failCount++
			log.Println(err)
			continue
		}
		successCount++
		ts:=time.Since(t1).Milliseconds()
		if minTs > ts {
			minTs = ts
		}
		if maxTs <ts {
			maxTs = ts
		}
		totalTs+=ts
		fmt.Printf("Reply from %d.%d.%d.%d: bytes=%d time=%dms TTL=%d\n",buf[12],buf[13],buf[14],buf[15],n-28,ts,buf[8])
		time.Sleep(time.Second)
	}
	fmt.Printf("Ping statistics for %s:\n\t:Packets: Sent = %d, Received = %d, Lost = %d (%.2f%% loss),\nApproximate round trip times in milli-seconds:\n\tMinimum = %dms, Maximum = %dms, Average = %dms",
	conn.RemoteAddr(),sendCount,successCount,failCount,float64(failCount)/float64(sendCount),minTs,maxTs,totalTs/int64(sendCount))
}
func GetComandArgs(){
	flag.Int64Var(&timeout,"w",1000,"请求超时时间")
	flag.IntVar(&size,"l",32,"发送字节数")
	flag.IntVar(&count,"n",4,"请求次数") 
	flag.Parse()
}

func checkSum(data []byte) uint16{
	length := len(data)
	index := 0
	var sum uint32 = 0
	for length >1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}
	if length != 0{
		sum += uint32(data[index])
	}
	hi16 := sum >>16
	for hi16 != 0{
		sum = hi16 + uint32(uint16(sum))
		hi16 = sum >>16
	}
	return uint16(^sum)
}