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

all: put sfscat get

put: put.$(su) ../client/client.go
	$(gl) -o put put.$(su)

put.$(su): put.go
	$(gc) put.go

sfscat: sfscat.$(su) ../client/client.go
	$(gl) -o sfscat sfscat.$(su)

sfscat.$(su): sfscat.go
	$(gc) sfscat.go

get: get.$(su) ../client/client.go
	$(gl) -o get get.$(su)

get.$(su): get.go
	$(gc) get.go

clean:
	-rm -f *.$(su) put sfscat get

clean-all: clean
