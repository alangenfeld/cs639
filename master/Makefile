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


master: trie.$(su) master.$(su) runmaster.$(su)
	$(gl) -o master runmaster.$(su)
	
trie.$(su): trie.go
	$(gc) trie.go
	
master.$(su): master.go serverHeap.go
	$(gc) master.go serverHeap.go
	
runmaster.$(su): runmaster.go
	$(gc) runmaster.go
	
clean:
	-rm -f *.$(su)
	
clean-all: clean master
