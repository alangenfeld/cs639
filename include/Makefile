LBITS := $(shell getconf LONG_BIT)
ifeq ($(LBITS),64)
gc=6g
gl=6l
su=6
else
gc=8g
gl=8l
su=8
endif

default: *.go
	$(gc) sfs.go
clean:
	rm *.$(su)
