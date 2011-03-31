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

all:t1 t2 t3

t1:t1.go
	$(gc) t1.go
	$(gl) -o t1 t1.$(su)

t2:t2.go
	$(gc) t2.go
	$(gl) -o t2 t2.$(su)

t3:t3.go
	$(gc) t3.go
	$(gl) -o t3 t3.$(su)

clean:
	rm -f *.$(su) t1 t2 t3
