package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

type SYSTEMTIME struct {
	wYear         uint16
	wMonth        uint16
	wDayOfWeek    uint16
	wDay          uint16
	wHour         uint16
	wMinutes      uint16
	wSecond       uint16
	wMilliseconds uint16
}

func setSystemTime(year, month, day, hour, min, sec uint16) error {
	s := SYSTEMTIME{wYear: year, wMonth: month, wDay: day, wHour: hour, wMinutes: min, wSecond: sec}
	DllTestDef, _ := syscall.LoadLibrary("kernel32.dll")
	defer syscall.FreeLibrary(DllTestDef)
	SetSystemTime, err := syscall.GetProcAddress(DllTestDef, "SetSystemTime")
	ret, _, err := syscall.SyscallN(SetSystemTime, uintptr(unsafe.Pointer(&s)))
	if ret <= 0 {
		return err
	} else {
		return nil
	}
}

func main() {
	// if len(os.Args) < 3 {
	// 	fmt.Println("Usage: go run main.go <parameter>")
	// 	return
	// }
	fmt.Printf("Begin parse parameter\n")
	month := os.Args[1]
	fmt.Printf("The first argument is: %s\n", month)
	nmonth, _ := strconv.Atoi(month)

	day := os.Args[2]
	fmt.Printf("The first argument is: %s\n", day)
	nday, _ := strconv.Atoi(day)

	hour := os.Args[3]
	fmt.Printf("The first argument is: %s\n", hour)
	nhour, _ := strconv.Atoi(hour)
	minute := os.Args[4]
	fmt.Printf("The first argument is: %s\n", minute)
	nminute, _ := strconv.Atoi(minute)
	second := os.Args[5]
	fmt.Printf("The first argument is: %s\n", second)
	nsecond, _ := strconv.Atoi(second)

	// 设置系统时间为2023年4月1日 12:00:00
	err := setSystemTime(2024, uint16(nmonth), uint16(nday), uint16(nhour), uint16(nminute), uint16(nsecond))
	if err != nil {
		fmt.Printf("Error setting system time: %v\n", err)
	} else {
		fmt.Println("System time set successfully.")
	}

	// 等待几秒钟以便改变生效
	time.Sleep(5 * time.Second)
}
