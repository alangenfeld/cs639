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

client: client.$(su) test.go
	$(gc) test.go
	$(gl) -o test test.$(su)

client.$(su): client.go
	$(gc) client.go

#sfs.go:
#	6g ../include/sfs.go
#	cp sfs.6 ../include/

clean:
	rm *.$(su) -f

