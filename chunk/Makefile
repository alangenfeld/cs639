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

chunkServer: chunkServer.go chunk
	$(gc) chunkServer.go
	$(gl) -o chunkServer.out chunkServer.$(su)

chunk:
	$(gc) chunk.go

fs:
	$(gc) -o chunk.$(su) chunk_fs.go

mem:
	$(gc) -o chunk.$(su) chunk_mem.go

clean:
	rm *.$(su)
