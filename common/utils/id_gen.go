package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/bytedance/gopkg/lang/fastrand"
)

type IDGenerator struct {
	traceIDPool <-chan string
	spanIDPool  <-chan string
	logIDPool   <-chan string
	stop        chan interface{}
}

func NewIDGenerator(maxSize int) *IDGenerator {
	stop := make(chan interface{})
	idgen := &IDGenerator{
		traceIDPool: newTraceIdPool(maxSize, stop),
		spanIDPool:  newSpanIdPool(maxSize, stop),
		logIDPool:   newLogIdPool(maxSize, stop),
		stop:        stop,
	}

	return idgen
}

func (idgen *IDGenerator) Stop() {
	select {
	case <-idgen.stop:
	default:
		close(idgen.stop)
	}
}

func (idgen *IDGenerator) NewLogID() string {
	return <-idgen.logIDPool
}

func (idgen *IDGenerator) NewTraceID() string {
	return <-idgen.traceIDPool
}

func (idgen *IDGenerator) NewSpanID(pspanID string) string {
	spanID := <-idgen.spanIDPool
	if pspanID == "" {
		return spanID
	}

	return fmt.Sprintf("%s:%s", pspanID, spanID)
}

func newTraceIdPool(size int, stop chan interface{}) <-chan string {
	pool := make(chan string, size)

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				pool <- fmt.Sprintf("%010d%d%d%d", IPv4Int(), os.Getpid(), time.Now().UnixMilli(), fastrand.Uint64())
			}
		}
	}()

	return pool
}

func newSpanIdPool(size int, stop chan interface{}) <-chan string {
	pool := make(chan string, size)

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				pool <- fmt.Sprintf("%d%d", os.Getpid(), fastrand.Uint32())
			}
		}
	}()

	return pool
}

func newLogIdPool(size int, stop chan interface{}) <-chan string {
	pool := make(chan string, size)

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				pool <- fmt.Sprintf("%010d%d%d%d", IPv4Int(), os.Getpid(), time.Now().UnixMilli(), fastrand.Uint64())
			}
		}
	}()

	return pool
}
