include $(GOROOT)/src/Make.inc

TARG=bitbucket.org/taruti/libauth.go
GOFILES=\
	common.go\
	getuserpasswd.go\
	os_$(GOOS).go\


include $(GOROOT)/src/Make.pkg
