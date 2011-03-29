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

t1: 
	$(gc) t1.go
	$(gl) -o t1 t1.$(su)

clean:
	rm -f *.$(su)
