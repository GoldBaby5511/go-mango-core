package util

import (
	"github.com/google/uuid"
	"net"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func CurMemory() int64 {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	return int64(rtm.Alloc / 1024)
}

func ParseArgsUint32(name string, args []string) (uint32, bool) {
	for i := 0; i < len(args); i++ {
		a := strings.Split(args[i], "=")
		if len(a) != 2 {
			continue
		}
		if a[0] == name {
			v, err := strconv.Atoi(a[1])
			if err == nil {
				return uint32(v), true
			}
		}
	}
	return 0, false
}

func ParseArgsString(name string, args []string) (string, bool) {
	for i := 0; i < len(args); i++ {
		a := strings.Split(args[i], "=")
		if len(a) != 2 {
			continue
		}
		if a[0] == name {
			return a[1], true
		}
	}
	return "", false
}

func MakeUint64FromUint32(high, low uint32) uint64 {
	return uint64(high)<<32 | uint64(low)
}

func Get2Uint32FromUint64(v uint64) (uint32, uint32) {
	return GetHUint32FromUint64(v), GetLUint32FromUint64(v)
}

func GetHUint32FromUint64(v uint64) uint32 {
	return uint32(v >> 32)
}

func GetLUint32FromUint64(v uint64) uint32 {
	return uint32(v & 0xFFFFFFFF)
}

func GetIPFromIPAddress(addr string) string {
	a := strings.Split(addr, ":")
	if len(a) != 2 {
		return ""
	}
	return a[0]
}

func GetPortFromIPAddress(addr string) int {
	a := strings.Split(addr, ":")
	if len(a) != 2 {
		return 0
	}
	p, _ := strconv.Atoi(a[1])
	return p
}

func PortInUse(portNumber int) bool {
	p := strconv.Itoa(portNumber)
	addr := net.JoinHostPort("127.0.0.1", p)
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

func StrconvAsInt64(str string, defaultValue int64) int64 {
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

func GetUUID() string {
	return uuid.New().String()
}
