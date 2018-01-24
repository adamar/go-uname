

## Go-Uname

Get name and information about the current kernel


## GoDoc

https://godoc.org/github.com/adamar/go-uname



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
