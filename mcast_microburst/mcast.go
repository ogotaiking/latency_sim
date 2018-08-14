package main

import "fmt"
import "net"
import "time"
import "syscall"
import "flag"

func setTTL(conn *net.UDPConn, ttl int) error {
    f, err := conn.File()
    if err != nil {
        return err
    }
    defer f.Close()
    fd := int(f.Fd())
    return syscall.SetsockoptInt(fd, syscall.SOL_IP, syscall.IP_MULTICAST_TTL, ttl)
}

func SendPacket(laddress,raddress string,pktLen,pktNumber,pktBatch int, pktInterval,tick int64) {
        raddr, err := net.ResolveUDPAddr("udp", raddress)
        if err != nil {
                panic(err)
        }
        laddr, err := net.ResolveUDPAddr("udp", laddress)

        conn, err := net.DialUDP("udp", laddr, raddr)
        if err != nil {
                panic(err)
        }
        setTTL(conn,64)
        var paddingStr string
        for i := 0 ; i < pktLen - 42; i++ {
            paddingStr = paddingStr + "A"
        }
        payload :=[]byte(paddingStr)

        for {
            tc:= time.After(time.Nanosecond * time.Duration(tick*1000*1000) )
            println(fmt.Sprintf("Sending data to address: [%s]", raddress))

            for i := 0 ; i < pktNumber ; i++ {
                conn.Write(payload)
                if pktInterval != 0 {
                   if i % pktBatch == 0 {
                       time.Sleep(time.Nanosecond * time.Duration(pktInterval*1000))
                   }
                }
            }
            <-tc
        }
}

func main() {
        var srcAddr string
        var dstAddr string

        var pktNumber int
        var pktLen  int

        var pktInterval int64
        var pktBatch int
        var tick int64

        flag.StringVar(&srcAddr,"src","1.0.0.2:8000","Source IP:port")
        flag.StringVar(&dstAddr,"dst","224.0.100.100:8000","Destination IP:port")

        flag.IntVar(&pktLen,"len",256,"Packet Length")
        if pktLen < 64 {
           pktLen = 64
        }
        flag.IntVar(&pktNumber,"num",500,"Number of packets per tick")

        flag.Int64Var(&pktInterval,"interval",0,"Interval between  multicast packet(usec)")
        flag.IntVar(&pktBatch,"batch",1,"Number of batch size ")
        flag.Int64Var(&tick,"tick",500,"Tick interval(msec)")
        flag.Parse()

        println("Multicast Micro Burst Generator......")
        time.Sleep(time.Second)
        SendPacket(srcAddr,dstAddr,pktLen,pktNumber,pktBatch,pktInterval,tick)
}


