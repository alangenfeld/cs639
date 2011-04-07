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

test: client.$(su) test.$(su)
	$(gl) -o test test.$(su)

client.$(su): client.go
	$(gc) client.go
	
test.$(su): test.go
	$(gc) test.go

clean:
	-rm -f *.$(su)
	
clean-all: clean test
