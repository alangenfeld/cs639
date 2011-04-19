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

cproxy: runChunkProxy.$(su) chunkproxy.$(su) chunk.$(su)
	$(gl) -o cproxy runChunkProxy.$(su)

serv: chunkServer.$(su) chunk.$(su)
	$(gl) -o serv chunkServer.$(su)

runChunkProxy.$(su): chunkproxy.$(su)
	$(gc)  runChunkProxy.go

chunkproxy.$(su): chunk.$(su)
	$(gc) chunkProxy.go

chunkServer.$(su): chunkServer.go chunk.$(su)
	$(gc) chunkServer.go
	
chunk.$(su): chunk_mem.go
	$(gc) -o chunk.$(su) chunk_mem.go

fs: chunk_fs.go
	$(gc) -o chunk.$(su) chunk_fs.go
	$(gc) chunkServer.go
	$(gl) -o chunk chunkServer.$(su)

clean:
	-rm -f chunk *.$(su) 
	
clean-all: clean serv
