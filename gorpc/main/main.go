package main

import (
	"gorpc"
	"log"
	"net"
	"sync"
	"time"
)

type Foo int

type Args struct {
	Num1, Num2 int
}

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addr chan string) {
	var foo Foo
	if err := gorpc.Register(&foo); err != nil {
		log.Fatal("register error:", err)
	}

	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	gorpc.Accept(l)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	// in fact, following code is like a simple geerpc client
	//conn, _ := net.Dial("tcp", <-addr)
	//defer func() { _ = conn.Close() }()
	//
	//time.Sleep(time.Second)
	//// send options
	//_ = json.NewEncoder(conn).Encode(gorpc.DefaultOption)
	//cc := codec.NewGobCodec(conn)
	//// send request & receive response
	//for i := 0; i < 5; i++ {
	//	h := &codec.Header{
	//		ServiceMethod: "Foo.Sum",
	//		Seq:           uint64(i),
	//	}
	//	_ = cc.Write(h, fmt.Sprintf("gorpc req %d", h.Seq))
	//	_ = cc.ReadHeader(h)
	//	var reply string
	//	_ = cc.ReadBody(&reply)
	//	log.Println("reply:", reply)
	//}

	client, _ := gorpc.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			//args := fmt.Sprintf("gorpc req %d", i)
			//
			//var reply string
			//if err := client.Call("Foo.Sum", args, &reply); err != nil {
			//	log.Fatal("call Foo.Sum error:", err)
			//}
			//
			//log.Println("reply:", reply)

			args := &Args{Num1: i, Num2: i * i}
			var reply int
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}

	wg.Wait()
}
