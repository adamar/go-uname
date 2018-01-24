package uname

import (
        "syscall"
	"fmt"
)

type UtsName struct {
	Sysname		string
 	Nodename	string
	Release		string
	Version		string
	Machine		string
	Domainname	string
}

func ctos(input *[65]int8) string {
	length := 0
        str := make([]byte, len(input))
        for ; length < len(input); length++ {
                if input[length] == 0 {
                        break
                }
                str[length] = uint8(input[length])
        }
        return string(str[0:length])
}

func Uname() {
        utsname := syscall.Utsname{}
        _ = syscall.Uname(&utsname)
	return UtsName{
		Sysname:	Ctos(&utsname.Sysname),
		Nodename:	Ctos(&utsname.Nodename),
		Release:	Ctos(&utsname.Release),
		Version:	Ctos(&utsname.Version),
		Machine:	Ctos(&utsname.Machine),
		Domainname:	Ctos(&utsname.Domainname),
	}
}
