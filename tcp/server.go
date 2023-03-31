// Package tcp
//
// ----------------develop info----------------
// @Author zhihaohe@rastar.com
// @DateTime 2023-3-31 17:17
// --------------------------------------------
//
package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ListenAndServe(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err: %v", err))
	}
	defer func() {
		listener.Close()
		log.Println("listener close success")
	}()

	log.Println(fmt.Printf("bind: %s, start listening...", address))
	for {
		// Accept 会一直阻塞直到有新的链接建立或者listen中断才会返回
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("accept err:%v", err))
		}
		//开启新的 goroutine 处理链接
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		// ReadString 会一直阻塞直到遇到分隔符 '\n'
		// 遇到分隔符后 ReadString 会返回上次遇到分隔符到现在收到所有数据
		// 若在遇到分隔符之前发生异常，ReadString会返回已收到的数据和错误信息
		msg, err := reader.ReadString('\n')
		if err != nil {
			//连接中断
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
		}
		b := []byte(msg)
		conn.Write(b)
	}
}

func main() {
	ListenAndServe(":8080")
}
