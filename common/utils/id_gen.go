package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
				sb := strings.Builder{}
				sb.WriteString(strconv.FormatUint(uint64(time.Now().UnixMilli()), 36))
				sb.WriteString(IPv4Hex())
				sb.WriteString(strconv.FormatUint(fastrand.Uint64(), 36))
				sb.WriteString(strconv.FormatInt(int64(os.Getpid()), 10))

				pool <- sb.String()
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
				sb := strings.Builder{}
				sb.WriteString(strconv.FormatInt(int64(os.Getpid()), 10))
				sb.WriteString(strconv.FormatUint(uint64(fastrand.Uint32()), 36))

				pool <- sb.String()
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
				sb := strings.Builder{}
				sb.WriteString(strconv.FormatUint(uint64(time.Now().UnixMilli()), 36))
				sb.WriteString(IPv4Hex())
				sb.WriteString(strconv.FormatUint(fastrand.Uint64(), 36))
				sb.WriteString(strconv.FormatUint(uint64(os.Getpid()), 10))

				pool <- sb.String()
			}
		}
	}()

	return pool
}
