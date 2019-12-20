package svc

import (
	"fmt"
	"testing"
	"time"
)

type testService struct {
	ConstNamed
	Runnable
}

func loop(stop chan bool, quit chan error) {
	for {
		select {
		case <-stop:
			close(quit)
			return
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func Test_default_service(t *testing.T) {
	svc := &testService{
		ConstNamed: NewConstNamed(123, "name1"),
		Runnable:   NewBasicRunnable(loop),
	}

	go func() {
		m, ok := <-svc.Quit()
		fmt.Println("quited.", m, ok)
		m, ok = <-svc.Quit()
		fmt.Println("quited.", m, ok)
	}()
	fmt.Println(svc.Status())
	go svc.Run()

	fmt.Println(svc.Status())
	time.Sleep(time.Millisecond * 100)
	fmt.Println(svc.Status())
	time.Sleep(time.Millisecond * 100)

	fmt.Println("start stop")
	svc.Stop()
	fmt.Println("end stop")

	time.Sleep(time.Millisecond * 100)
	fmt.Println(svc.Status())
	time.Sleep(time.Millisecond * 100)

}

type AAAA struct {
	A int
}

type BBBB struct {
	*AAAA
	A int
	B string
}

func Test_asdf(t *testing.T) {
	b := &BBBB{
		AAAA: &AAAA{1},
		A:    2,
		B:    "abc",
	}

	fmt.Println(b)

}
