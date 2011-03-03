v0

right now mostly working:
	1 chunk per file
	1 server per chunk
	chunks phone home on launch
	client queries master, receives chunk ID and server TCPAddr
	client can write and read chunks

todo:
	switch from http to tcp RPC calls
	multiple chunks per file

![Alt text](http://lalkaka.com/rusty/StealthFS_logo.png)

stealthy...
