package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func DllTestDef_check() (result int64, err error) {
	DllTestDef, _ := syscall.LoadLibrary("DllTestDef.dll")
	fmt.Println("+++++++syscall.LoadLibrary:", DllTestDef, "+++++++")
	defer syscall.FreeLibrary(DllTestDef)
	getlicensequantity, err := syscall.GetProcAddress(DllTestDef, "get_license_quantity")
	fmt.Println("GetProcAddress", getlicensequantity)
	ret, _, err := syscall.SyscallN(getlicensequantity,
		uintptr(6000010),
		uintptr(201001))
	fmt.Fprintln(os.Stdout, []any{"syscall return value is %d", int64(ret)}...)
	// get_license_quantity, err := syscall.GetProcAddress(DllTestDef, "get_license_quantity")
	// fmt.Println("GetProcAddress", get_license_quantity)
	// s := "130-3491583265"
	// b := append([]byte(s), 0)
	// ret, _, err := syscall.SyscallN(get_license_quantity,
	// 	uintptr(unsafe.Pointer(&b[0])),
	// 	uintptr(6000010),
	// 	uintptr(301000))

	if int(ret) == -3 {
		return 0, nil
	} else {
		return int64(ret), nil
	}
	// serialcode, err := syscall.UTF16PtrFromString(s)
	// if err == nil {

	// 	ret, _, err := syscall.SyscallN(get_license_quantity,
	// 		uintptr(unsafe.Pointer(&b[0])),
	// 		6000010,
	// 		301000)
	// 	// if err != syscall.Errno(0) {
	// 	// 	return 0, err
	// 	// }
	// 	if err == syscall.Errno(0) {
	// 		return 0, nil
	// 	}
	// 	if int(ret) == -3 {
	// 		return 0, nil
	// 	} else {
	// 		return int(ret), nil
	// 	}

	// }
	return 0, nil

}

func main1() {
	// licensecount, err := DllTestDef_check()
	// if err != nil {
	// 	fmt.Printf("<%d>\n", licensecount)

	// }
	// fmt.Printf("<%d>\n", licensecount)
	t, _ := time.ParseInLocation("20060102150405", "20240808130057", time.Local)
	fmt.Printf("<%d>\n", t.Unix())

}
