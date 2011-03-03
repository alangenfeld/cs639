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

chunkServer: chunkServer.go
	$(gc) chunkServer.go
	$(gl) -o serv chunkServer.$(su)

fs: chunk_fs.go
	$(gc) -o chunk.$(su) chunk_fs.go
	$(gc) chunkServer.go
	$(gl) -o serv chunkServer.$(su)

mem: chunk_mem.go
	$(gc) -o chunk.$(su) chunk_mem.go
	$(gc) chunkServer.go
	$(gl) -o serv chunkServer.$(su)

clean:
	rm *.$(su)
