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

client: client.$(su)
	$(gc) test.go
	$(gl) -o client test.$(su)

client.$(su): 
	$(gc) client.go

clean:
	rm *.$(su) -f

