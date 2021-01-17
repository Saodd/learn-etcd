package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.0.6.239:12379", "10.0.6.239:22379", "10.0.6.239:32379"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	s, _ := concurrency.NewSession(cli)
	defer s.Close()

	l := concurrency.NewMutex(s, "/locked/resource/1")
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	fmt.Println("排队取锁")
	if err := l.Lock(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("带锁工作ing……")
	time.Sleep(10 * time.Second)
	if err := l.Unlock(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("释放锁")
}
