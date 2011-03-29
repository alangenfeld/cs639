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
	$(gc) trie.go
	$(gc) master.go serverHeap.go
	$(gc) runmaster.go
	$(gl) -o master runmaster.$(su)
clean:
	rm -f *.$(su)
