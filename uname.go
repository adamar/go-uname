package uname

import (
        "syscall"
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

func Uname() *UtsName {
        utsname := syscall.Utsname{}
        _ = syscall.Uname(&utsname)
	return UtsName{
		Sysname:	ctos(&utsname.Sysname),
		Nodename:	ctos(&utsname.Nodename),
		Release:	ctos(&utsname.Release),
		Version:	ctos(&utsname.Version),
		Machine:	ctos(&utsname.Machine),
		Domainname:	ctos(&utsname.Domainname),
	}
}
