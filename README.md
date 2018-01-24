

## Go-Uname

Get name and information about the current kernel


## GoDoc

https://godoc.org/github.com/adamar/go-uname


## Usage

Example
```
package main

import (
        gouname "github.com/adamar/go-uname"
        "fmt"
        "reflect"
)

func main() {

        uts := gouname.Uname()

        s := reflect.ValueOf(&uts).Elem()
        typeOfT := s.Type()

        for i := 0; i < s.NumField(); i++ {
                f := s.Field(i)
                fmt.Printf("%s:\t%v\n", typeOfT.Field(i).Name, f.Interface())
        }

}
```

Returns
```
Sysname:	Linux
Nodename:	testbox
Release:	3.10.0-327.22.2.el7.x86_64
Version:	#1 SMP Thu Jun 23 17:05:11 UTC 2016
Machine:	x86_64
Domainname:	(none)
```


## Underlying Linux Lib

The utsname struct is defined in <sys/utsname.h>

    struct utsname {
               char sysname[];    /* Operating system name (e.g., "Linux") */
               char nodename[];   /* Name within "some implementation-defined network" */
               char release[];    /* Operating system release (e.g., "2.6.28") */
               char version[];    /* Operating system version */
               char machine[];    /* Hardware identifier */
               char domainname[]; /* NIS or YP domain name */
           };
