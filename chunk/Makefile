LBITS := $(shell getconf LONG_BIT)
ifeq ($(LBITS),64)
gc=6g
gl=6l
else
gc=8g
gl=8l
endif

chunkServer: chunkServer.go chunk
	$(gc) chunkServer.go
	$(gl) chunkServer.6

chunk:
	$(gc) chunk.go

fs:
	$(gc) -o chunk.6 chunk_fs.go

mem:
	$(gc) -o chunk.6 chunk_mem.go

clean:
	rm *.6
