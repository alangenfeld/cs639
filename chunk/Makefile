
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
