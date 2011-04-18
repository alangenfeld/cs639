
package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
	"time"
)

func main(){
	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)
	var buf []byte
	buf = buf
	var ret int
	ret = ret
last := time.Nanoseconds()
now := time.Nanoseconds()

fd1 := client.Open("/x4uQZLfO13.txt", client.O_RDWR|client.O_CREATE)
if(fd1 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd1, []byte("uR98v7OLlaOcfQDM6dK5WxXeD1ua0H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) uR98v7OLlaOcfQDM6dK5WxXeD1ua0H

ret = client.Seek(fd1, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


ret = client.Seek(fd1, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


ret = client.Close(fd1)
if(ret != 0) {
  panic("close failed")
}


fd2 := client.Open("/_BLukPDMrv.txt", client.O_RDWR|client.O_CREATE)
if(fd2 < 0) {
  panic("open failed")
}


ret = client.Seek(fd2, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd2, []byte("1MrsxizlJz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) 1MrsxizlJz

fd3 := client.Open("/LK829EmRZQ.txt", client.O_RDWR|client.O_CREATE)
if(fd3 < 0) {
  panic("open failed")
}


ret = client.Close(fd2)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd3, []byte("IZwvr3F2uAIMc2ggv4Zf9NuVyHrZd2FrT2yX3fw2WDPj8z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) IZwvr3F2uAIMc2ggv4Zf9NuVyHrZd2FrT2yX3fw2WDPj8z
//fd state: (46) IZwvr3F2uAIMc2ggv4Zf9NuVyHrZd2FrT2yX3fw2WDPj8z

ret = client.Write(fd3, []byte("CILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (109) IZwvr3F2uAIMc2ggv4Zf9NuVyHrZd2FrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLT
//fd state: (109) IZwvr3F2uAIMc2ggv4Zf9NuVyHrZd2FrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLT

ret = client.Write(fd3, []byte("x1drJMj_jJtbWl2C9JtH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (129) IZwvr3F2uAIMc2ggv4Zf9NuVyHrZd2FrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLTx1drJMj_jJtbWl2C9JtH

buf, ret = client.Read(fd3, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd4 := client.Open("/1ME66mmhaH.txt", client.O_RDWR|client.O_CREATE)
if(fd4 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd3, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd4, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd5 := client.Open("/YfmUYg5AyQ.txt", client.O_RDWR|client.O_CREATE)
if(fd5 < 0) {
  panic("open failed")
}


ret = client.Close(fd3)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd5)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd4, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd4, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd6 := client.Open("/Z2evREiOSO.txt", client.O_RDWR|client.O_CREATE)
if(fd6 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd4, []byte("TUOGs3FBdlpHEJ6nZBrH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) TUOGs3FBdlpHEJ6nZBrH

fd7 := client.Open("/CdZyI4bMVY.txt", client.O_RDWR|client.O_CREATE)
if(fd7 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd7, []byte("ylwLPUz4IfCymaNGuSE2Hgp8JocFAGYXNCviposSb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) ylwLPUz4IfCymaNGuSE2Hgp8JocFAGYXNCviposSb

buf, ret = client.Read(fd4, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (41) ylwLPUz4IfCymaNGuSE2Hgp8JocFAGYXNCviposSb

ret = client.Write(fd7, []byte("6IViaWLJKOxaqgxTfVWouoMrnBlfnhJcVSmgxET6fScLdyqX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) ylwLPUz4IfCymaNGuSE2Hgp8JocFAGYXNCviposSb6IViaWLJKOxaqgxTfVWouoMrnBlfnhJcVSmgxET6fScLdyqX
//fd state: (0) 

ret = client.Write(fd6, []byte("9POj0HsGwb9FJK9vzSKIpcK_HHn8ORe_nlLKL_YKpmFw9oH8gA93deNbtC4jUx5Kijv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) 9POj0HsGwb9FJK9vzSKIpcK_HHn8ORe_nlLKL_YKpmFw9oH8gA93deNbtC4jUx5Kijv

buf, ret = client.Read(fd7, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd8 := client.Open("/1ME66mmhaH.txt", client.O_RDWR|client.O_CREATE)
if(fd8 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd8, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TUOGs3FBdlpHEJ6nZBrH") {
  panic("wrong data returned")
}


ret = client.Close(fd6)
if(ret != 0) {
  panic("close failed")
}


fd9 := client.Open("/rWIIG4WJd_.txt", client.O_RDWR|client.O_CREATE)
if(fd9 < 0) {
  panic("open failed")
}


fd10 := client.Open("/rWIIG4WJd_.txt", client.O_RDWR|client.O_CREATE)
if(fd10 < 0) {
  panic("open failed")
}


fd11 := client.Open("/4300GnPlIR.txt", client.O_RDWR|client.O_CREATE)
if(fd11 < 0) {
  panic("open failed")
}


fd12 := client.Open("/otUQRxsRyj.txt", client.O_RDWR|client.O_CREATE)
if(fd12 < 0) {
  panic("open failed")
}


ret = client.Seek(fd11, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd4, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


ret = client.Seek(fd8, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


fd13 := client.Open("/YVbJEFJUJd.txt", client.O_RDWR|client.O_CREATE)
if(fd13 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd7, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd7)
if(ret != 0) {
  panic("close failed")
}


fd14 := client.Open("/8lrlg9FbK5.txt", client.O_RDWR|client.O_CREATE)
if(fd14 < 0) {
  panic("open failed")
}


fd15 := client.Open("/4300GnPlIR.txt", client.O_RDWR|client.O_CREATE)
if(fd15 < 0) {
  panic("open failed")
}


ret = client.Close(fd12)
if(ret != 0) {
  panic("close failed")
}


fd16 := client.Open("/S3clCGhVDt.txt", client.O_RDWR|client.O_CREATE)
if(fd16 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd10, []byte("p0OQ_0akqNiLtvDISsFwtqxYDrFNNHPnZRstd5H6Ejflv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) p0OQ_0akqNiLtvDISsFwtqxYDrFNNHPnZRstd5H6Ejflv

ret = client.Seek(fd13, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd15, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd11, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd17 := client.Open("/Ac1BgFRgda.txt", client.O_RDWR|client.O_CREATE)
if(fd17 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd11, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd16, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd15)
if(ret != 0) {
  panic("close failed")
}


fd18 := client.Open("/1ME66mmhaH.txt", client.O_RDWR|client.O_CREATE)
if(fd18 < 0) {
  panic("open failed")
}


fd19 := client.Open("/wCijCfQcBC.txt", client.O_RDWR|client.O_CREATE)
if(fd19 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd14, []byte("fMEeIrgo6Qj0rv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) fMEeIrgo6Qj0rv

ret = client.Seek(fd4, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd19, []byte("tkOPj86z4ehj7kBG461RpktI1BFiV_odgpFwS3XbU4A_WB52pqvj7zEe3QXiGkNaLn_1DEkudo6NO0PhSkmr4mAnP4i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) tkOPj86z4ehj7kBG461RpktI1BFiV_odgpFwS3XbU4A_WB52pqvj7zEe3QXiGkNaLn_1DEkudo6NO0PhSkmr4mAnP4i
//fd state: (0) p0OQ_0akqNiLtvDISsFwtqxYDrFNNHPnZRstd5H6Ejflv

ret = client.Write(fd9, []byte("WgptUaqjIA7147_4YEzZc0U5jnV4avAVljR7_JfnG5vXCPwJilWlcMYv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) WgptUaqjIA7147_4YEzZc0U5jnV4avAVljR7_JfnG5vXCPwJilWlcMYv

ret = client.Close(fd16)
if(ret != 0) {
  panic("close failed")
}

//fd state: (91) tkOPj86z4ehj7kBG461RpktI1BFiV_odgpFwS3XbU4A_WB52pqvj7zEe3QXiGkNaLn_1DEkudo6NO0PhSkmr4mAnP4i

ret = client.Write(fd19, []byte("ZEMF7YCtorYRiw1vCtuhzrMgExf7nEaJzo4_GvfEs6uLJOccAZS4np8xhJ2TSJaDZ9qFpigSitrkd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (168) tkOPj86z4ehj7kBG461RpktI1BFiV_odgpFwS3XbU4A_WB52pqvj7zEe3QXiGkNaLn_1DEkudo6NO0PhSkmr4mAnP4iZEMF7YCtorYRiw1vCtuhzrMgExf7nEaJzo4_GvfEs6uLJOccAZS4np8xhJ2TSJaDZ9qFpigSitrkd

fd20 := client.Open("/s4SMoZEfz7.txt", client.O_RDWR|client.O_CREATE)
if(fd20 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd20, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd4)
if(ret != 0) {
  panic("close failed")
}


fd21 := client.Open("/JPFo4bf_mL.txt", client.O_RDWR|client.O_CREATE)
if(fd21 < 0) {
  panic("open failed")
}


ret = client.Seek(fd17, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd8, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


buf, ret = client.Read(fd18, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TUOGs3FBdlpHEJ6nZBrH") {
  panic("wrong data returned")
}

//fd state: (20) TUOGs3FBdlpHEJ6nZBrH

ret = client.Write(fd18, []byte("XRQjOspCjq3YmhDlnZVOF2uRt9ej0h8Z76h4bHfHK0zPMb1r8fudHfzW0Uygdj8iyALtG8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) TUOGs3FBdlpHEJ6nZBrHXRQjOspCjq3YmhDlnZVOF2uRt9ej0h8Z76h4bHfHK0zPMb1r8fudHfzW0Uygdj8iyALtG8

ret = client.Close(fd10)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd11)
if(ret != 0) {
  panic("close failed")
}


fd22 := client.Open("/wCijCfQcBC.txt", client.O_RDWR|client.O_CREATE)
if(fd22 < 0) {
  panic("open failed")
}

//fd state: (0) tkOPj86z4ehj7kBG461RpktI1BFiV_odgpFwS3XbU4A_WB52pqvj7zEe3QXiGkNaLn_1DEkudo6NO0PhSkmr4mAnP4iZEMF7YCtorYRiw1vCtuhzrMgExf7nEaJzo4_GvfEs6uLJOccAZS4np8xhJ2TSJaDZ9qFpigSitrkd

ret = client.Write(fd22, []byte("NuXqDiL0Mvby_iSpWzJRZdOlvbwUzifvr_g5x97FApg79sKdoG7oX4lNHcLrklkCRyZmkVxy8n5993tqecFEENUtM0LKL9daCwe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) NuXqDiL0Mvby_iSpWzJRZdOlvbwUzifvr_g5x97FApg79sKdoG7oX4lNHcLrklkCRyZmkVxy8n5993tqecFEENUtM0LKL9daCweorYRiw1vCtuhzrMgExf7nEaJzo4_GvfEs6uLJOccAZS4np8xhJ2TSJaDZ9qFpigSitrkd

ret = client.Seek(fd18, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


fd23 := client.Open("/4300GnPlIR.txt", client.O_RDWR|client.O_CREATE)
if(fd23 < 0) {
  panic("open failed")
}


fd24 := client.Open("/LW1OckBfGg.txt", client.O_RDWR|client.O_CREATE)
if(fd24 < 0) {
  panic("open failed")
}


fd25 := client.Open("/YfmUYg5AyQ.txt", client.O_RDWR|client.O_CREATE)
if(fd25 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd18, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OspCjq3YmhDlnZVOF2uRt9ej0h8Z76h4bHfHK0zPMb1r8fudHfzW0Uy") {
  panic("wrong data returned")
}


ret = client.Close(fd22)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd18, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


buf, ret = client.Read(fd24, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd18)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd24)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd23, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd14, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd8, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rHXRQjOspCjq3YmhDlnZVOF2uRt9ej0h8Z7") {
  panic("wrong data returned")
}

//fd state: (14) fMEeIrgo6Qj0rv

ret = client.Write(fd14, []byte("pC9YwSMd26S7Rpz0HE8yJF8kG3UIqnpCxxwEHh0z1WwLlRLIkkYjd22evF3w4lD2GEKvPkuUwdVV3Kpqapf_PiZP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) fMEeIrgo6Qj0rvpC9YwSMd26S7Rpz0HE8yJF8kG3UIqnpCxxwEHh0z1WwLlRLIkkYjd22evF3w4lD2GEKvPkuUwdVV3Kpqapf_PiZP

ret = client.Close(fd21)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd13, []byte("S_KajJMZfBB0gddHDL55vW9C7TqbfYOiTJA3LC8QiRuuNEUOICa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) S_KajJMZfBB0gddHDL55vW9C7TqbfYOiTJA3LC8QiRuuNEUOICa

buf, ret = client.Read(fd9, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (102) fMEeIrgo6Qj0rvpC9YwSMd26S7Rpz0HE8yJF8kG3UIqnpCxxwEHh0z1WwLlRLIkkYjd22evF3w4lD2GEKvPkuUwdVV3Kpqapf_PiZP

ret = client.Write(fd14, []byte("NNRn03vKBxylSuMxUj0Gf94OeFNAPEODG8Qxs4G2Su2jaOJ9y7_K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) fMEeIrgo6Qj0rvpC9YwSMd26S7Rpz0HE8yJF8kG3UIqnpCxxwEHh0z1WwLlRLIkkYjd22evF3w4lD2GEKvPkuUwdVV3Kpqapf_PiZPNNRn03vKBxylSuMxUj0Gf94OeFNAPEODG8Qxs4G2Su2jaOJ9y7_K

fd26 := client.Open("/JR6nH_boRB.txt", client.O_RDWR|client.O_CREATE)
if(fd26 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd25, []byte("m1CC_0SgjxqfeXLB9oXh4moURNXAFkkfXa7n25m7zp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) m1CC_0SgjxqfeXLB9oXh4moURNXAFkkfXa7n25m7zp

fd27 := client.Open("/CdZyI4bMVY.txt", client.O_RDWR|client.O_CREATE)
if(fd27 < 0) {
  panic("open failed")
}


ret = client.Seek(fd9, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Seek(fd27, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


ret = client.Close(fd20)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd14)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (42) m1CC_0SgjxqfeXLB9oXh4moURNXAFkkfXa7n25m7zp

ret = client.Write(fd25, []byte("beyaz9JirAlnjKUZXZ0bk5J"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) m1CC_0SgjxqfeXLB9oXh4moURNXAFkkfXa7n25m7zpbeyaz9JirAlnjKUZXZ0bk5J

ret = client.Close(fd19)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd17, []byte("TDXJOfypwsevp66He4fi0Mw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) TDXJOfypwsevp66He4fi0Mw

buf, ret = client.Read(fd13, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (28) ylwLPUz4IfCymaNGuSE2Hgp8JocFAGYXNCviposSb6IViaWLJKOxaqgxTfVWouoMrnBlfnhJcVSmgxET6fScLdyqX

ret = client.Write(fd27, []byte("IzNlkBBv6gSuclvLNFQaK7NBZDrCw8omFf4eVqSVo38mx7jM7604vtuID1g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) ylwLPUz4IfCymaNGuSE2Hgp8JocFIzNlkBBv6gSuclvLNFQaK7NBZDrCw8omFf4eVqSVo38mx7jM7604vtuID1gqX

ret = client.Seek(fd25, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


buf, ret = client.Read(fd13, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd25, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "feXLB9oXh4moURNXAFkkfXa7n25m7zpbeyaz9JirAlnjKUZXZ0bk5J") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd26, []byte("ThkPCe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) ThkPCe

ret = client.Close(fd13)
if(ret != 0) {
  panic("close failed")
}

//fd state: (6) ThkPCe

ret = client.Write(fd26, []byte("fQXKEHif4Eoh_PbVibIFFfXABMCiG4CwWJHZaaIibagmepOCuynPZ8l_d1C4lcqe_dZil_CrpkdoJa4De"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) ThkPCefQXKEHif4Eoh_PbVibIFFfXABMCiG4CwWJHZaaIibagmepOCuynPZ8l_d1C4lcqe_dZil_CrpkdoJa4De

buf, ret = client.Read(fd25, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd9, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Seek(fd27, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


fd28 := client.Open("/_mbJ__Dujj.txt", client.O_RDWR|client.O_CREATE)
if(fd28 < 0) {
  panic("open failed")
}


fd29 := client.Open("/fncK2K1IOC.txt", client.O_RDWR|client.O_CREATE)
if(fd29 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd29, []byte("zT_Y526EQfhTCxT4DgOu7y3WIW9leNJMeixI0yB5iFmiMmo081HEG_0mXMQJlC5dC__zEajYEG0ZITFY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) zT_Y526EQfhTCxT4DgOu7y3WIW9leNJMeixI0yB5iFmiMmo081HEG_0mXMQJlC5dC__zEajYEG0ZITFY

ret = client.Seek(fd28, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd23, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (23) TDXJOfypwsevp66He4fi0Mw

ret = client.Write(fd17, []byte("Yq0hIu7G9yDlMWcA87t8wF8AqXEokg9kbyKAkirA38eqEPzSxDRTueYzx1VwEx34pk5KCkxvfb48vPEWeE6M_3l"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) TDXJOfypwsevp66He4fi0MwYq0hIu7G9yDlMWcA87t8wF8AqXEokg9kbyKAkirA38eqEPzSxDRTueYzx1VwEx34pk5KCkxvfb48vPEWeE6M_3l

ret = client.Seek(fd25, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Close(fd8)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd26, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


fd30 := client.Open("/0prf8WI95c.txt", client.O_RDWR|client.O_CREATE)
if(fd30 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd23, []byte("_ockFJ0zIB4X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) _ockFJ0zIB4X

fd31 := client.Open("/S3clCGhVDt.txt", client.O_RDWR|client.O_CREATE)
if(fd31 < 0) {
  panic("open failed")
}


ret = client.Close(fd30)
if(ret != 0) {
  panic("close failed")
}


fd32 := client.Open("/w2FXc3l6HU.txt", client.O_RDWR|client.O_CREATE)
if(fd32 < 0) {
  panic("open failed")
}


fd33 := client.Open("/x4uQZLfO13.txt", client.O_RDWR|client.O_CREATE)
if(fd33 < 0) {
  panic("open failed")
}


ret = client.Seek(fd31, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd17, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


ret = client.Seek(fd32, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd34 := client.Open("/ezy6lpPkBl.txt", client.O_RDWR|client.O_CREATE)
if(fd34 < 0) {
  panic("open failed")
}


fd35 := client.Open("/wCijCfQcBC.txt", client.O_RDWR|client.O_CREATE)
if(fd35 < 0) {
  panic("open failed")
}


ret = client.Seek(fd32, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd36 := client.Open("/JR6nH_boRB.txt", client.O_RDWR|client.O_CREATE)
if(fd36 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd9, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "47_4YEzZc0U5jnV4avAVljR7_JfnG5vXCPwJilWlcMYv") {
  panic("wrong data returned")
}


fd37 := client.Open("/YfmUYg5AyQ.txt", client.O_RDWR|client.O_CREATE)
if(fd37 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd17, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "A87t8wF8AqXEokg9kbyKAki") {
  panic("wrong data returned")
}


ret = client.Seek(fd37, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


buf, ret = client.Read(fd23, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd25, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}


fd38 := client.Open("/4300GnPlIR.txt", client.O_RDWR|client.O_CREATE)
if(fd38 < 0) {
  panic("open failed")
}

//fd state: (12) _ockFJ0zIB4X

ret = client.Write(fd23, []byte("Uf7XOaDaKSsTth5LKGvIaIQeIq6xMJBZ3EGdHx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) _ockFJ0zIB4XUf7XOaDaKSsTth5LKGvIaIQeIq6xMJBZ3EGdHx

buf, ret = client.Read(fd36, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ThkPCefQXKEHif4Eoh_PbVibIFFfXABMCiG4CwWJHZaaIibagmepOCuynPZ8l_d1C4lcqe_dZil_CrpkdoJa4De") {
  panic("wrong data returned")
}


ret = client.Close(fd28)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd9)
if(ret != 0) {
  panic("close failed")
}

//fd state: (56) ThkPCefQXKEHif4Eoh_PbVibIFFfXABMCiG4CwWJHZaaIibagmepOCuynPZ8l_d1C4lcqe_dZil_CrpkdoJa4De

ret = client.Write(fd26, []byte("tm573NVFSJNjSrkdponjYwiwRs6BYS5bHc8th7hI7YPajMfjlVYKMepjhzXwkbyFu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (121) ThkPCefQXKEHif4Eoh_PbVibIFFfXABMCiG4CwWJHZaaIibagmepOCuytm573NVFSJNjSrkdponjYwiwRs6BYS5bHc8th7hI7YPajMfjlVYKMepjhzXwkbyFu

ret = client.Seek(fd34, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd39 := client.Open("/_BLukPDMrv.txt", client.O_RDWR|client.O_CREATE)
if(fd39 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd17, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rA38eqEPzSxDRTueYzx1VwEx34pk5KCkxvfb48vPEWeE6M_3l") {
  panic("wrong data returned")
}


ret = client.Close(fd23)
if(ret != 0) {
  panic("close failed")
}

//fd state: (43) m1CC_0SgjxqfeXLB9oXh4moURNXAFkkfXa7n25m7zpbeyaz9JirAlnjKUZXZ0bk5J

ret = client.Write(fd37, []byte("TpdydASOize3Wb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) m1CC_0SgjxqfeXLB9oXh4moURNXAFkkfXa7n25m7zpbTpdydASOize3WbZXZ0bk5J

ret = client.Seek(fd29, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


buf, ret = client.Read(fd39, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1MrsxizlJz") {
  panic("wrong data returned")
}


ret = client.Close(fd27)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd39)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd32, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd40 := client.Open("/_mbJ__Dujj.txt", client.O_RDWR|client.O_CREATE)
if(fd40 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd35, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NuXqDiL0Mvby_iSpWzJRZdOlvbwUzifvr_g5x97FApg79sKdoG7oX4lNHcLr") {
  panic("wrong data returned")
}


fd41 := client.Open("/hCnVwbgHwm.txt", client.O_RDWR|client.O_CREATE)
if(fd41 < 0) {
  panic("open failed")
}


fd42 := client.Open("/CdZyI4bMVY.txt", client.O_RDWR|client.O_CREATE)
if(fd42 < 0) {
  panic("open failed")
}


fd43 := client.Open("/OqvlaCPI05.txt", client.O_RDWR|client.O_CREATE)
if(fd43 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd43, []byte("3xm6oLn1kYf12wTlosat9l80Z_k76VheqslktiFBFV57y0F49QH9Mr_yVQCEyEpKJsYqaqkj_4psDdn2auoQz06Yu795i6rfYnI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) 3xm6oLn1kYf12wTlosat9l80Z_k76VheqslktiFBFV57y0F49QH9Mr_yVQCEyEpKJsYqaqkj_4psDdn2auoQz06Yu795i6rfYnI

ret = client.Close(fd29)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd26, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


fd44 := client.Open("/Xa2of_COvq.txt", client.O_RDWR|client.O_CREATE)
if(fd44 < 0) {
  panic("open failed")
}


ret = client.Close(fd42)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd36, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


fd45 := client.Open("/q6rX79dbpd.txt", client.O_RDWR|client.O_CREATE)
if(fd45 < 0) {
  panic("open failed")
}


ret = client.Close(fd26)
if(ret != 0) {
  panic("close failed")
}

//fd state: (48) ThkPCefQXKEHif4Eoh_PbVibIFFfXABMCiG4CwWJHZaaIibagmepOCuytm573NVFSJNjSrkdponjYwiwRs6BYS5bHc8th7hI7YPajMfjlVYKMepjhzXwkbyFu

ret = client.Write(fd36, []byte("tWVFHCvaRrnJLhxJ79YQ_eTR1WfjPnlQgdmhb7ocD2J9zezvriPdEIgCuOMMrahuOso5Mv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) ThkPCefQXKEHif4Eoh_PbVibIFFfXABMCiG4CwWJHZaaIibatWVFHCvaRrnJLhxJ79YQ_eTR1WfjPnlQgdmhb7ocD2J9zezvriPdEIgCuOMMrahuOso5MvyFu

ret = client.Seek(fd33, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


fd46 := client.Open("/mFvxS3bd1N.txt", client.O_RDWR|client.O_CREATE)
if(fd46 < 0) {
  panic("open failed")
}


ret = client.Seek(fd34, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd36, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


buf, ret = client.Read(fd45, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd41)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd32, []byte("GBV3FQYXb28zaXgqNy0k_l2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) GBV3FQYXb28zaXgqNy0k_l2

ret = client.Close(fd38)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd45, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd37)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd43)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd45, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd47 := client.Open("/9q9ndfi0_e.txt", client.O_RDWR|client.O_CREATE)
if(fd47 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd44, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd46)
if(ret != 0) {
  panic("close failed")
}


fd48 := client.Open("/mFvxS3bd1N.txt", client.O_RDWR|client.O_CREATE)
if(fd48 < 0) {
  panic("open failed")
}

//fd state: (24) ThkPCefQXKEHif4Eoh_PbVibIFFfXABMCiG4CwWJHZaaIibatWVFHCvaRrnJLhxJ79YQ_eTR1WfjPnlQgdmhb7ocD2J9zezvriPdEIgCuOMMrahuOso5MvyFu

ret = client.Write(fd36, []byte("9QAHWEwc4g8IZAaxxyvcKvj1CXx1xXfwwvhs1QxmTjtwPZEKiWx_9AMIlq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) ThkPCefQXKEHif4Eoh_PbVib9QAHWEwc4g8IZAaxxyvcKvj1CXx1xXfwwvhs1QxmTjtwPZEKiWx_9AMIlqmhb7ocD2J9zezvriPdEIgCuOMMrahuOso5MvyFu

fd49 := client.Open("/ZkuXs_QYnQ.txt", client.O_RDWR|client.O_CREATE)
if(fd49 < 0) {
  panic("open failed")
}

//fd state: (2) uR98v7OLlaOcfQDM6dK5WxXeD1ua0H

ret = client.Write(fd33, []byte("gQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLW

ret = client.Close(fd25)
if(ret != 0) {
  panic("close failed")
}


fd50 := client.Open("/PIYzUD6TOm.txt", client.O_RDWR|client.O_CREATE)
if(fd50 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd45, []byte("xvqRQRFhNMmTHjIolcUSJYf7X8V362wipzyN7RukBD3nYQfzFrqLi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) xvqRQRFhNMmTHjIolcUSJYf7X8V362wipzyN7RukBD3nYQfzFrqLi

ret = client.Seek(fd49, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (23) GBV3FQYXb28zaXgqNy0k_l2

ret = client.Write(fd32, []byte("3ZFeitsqslQ26bsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqtbTKkbCPOgSlGLDAlyqKlCdZPjfs1dYcV7_0b0hq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) GBV3FQYXb28zaXgqNy0k_l23ZFeitsqslQ26bsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqtbTKkbCPOgSlGLDAlyqKlCdZPjfs1dYcV7_0b0hq
//fd state: (116) GBV3FQYXb28zaXgqNy0k_l23ZFeitsqslQ26bsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqtbTKkbCPOgSlGLDAlyqKlCdZPjfs1dYcV7_0b0hq

ret = client.Write(fd32, []byte("cLumGDJf4lD1WsmIQLzBcKh8DFCc8DYT8FUI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) GBV3FQYXb28zaXgqNy0k_l23ZFeitsqslQ26bsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqtbTKkbCPOgSlGLDAlyqKlCdZPjfs1dYcV7_0b0hqcLumGDJf4lD1WsmIQLzBcKh8DFCc8DYT8FUI

fd51 := client.Open("/PIpopWJsly.txt", client.O_RDWR|client.O_CREATE)
if(fd51 < 0) {
  panic("open failed")
}


fd52 := client.Open("/_mbJ__Dujj.txt", client.O_RDWR|client.O_CREATE)
if(fd52 < 0) {
  panic("open failed")
}


ret = client.Seek(fd49, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd53 := client.Open("/JR6nH_boRB.txt", client.O_RDWR|client.O_CREATE)
if(fd53 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd34, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd54 := client.Open("/RRGCNYlXez.txt", client.O_RDWR|client.O_CREATE)
if(fd54 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd51, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (53) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLW

ret = client.Write(fd33, []byte("QYG7m4krfHoN62QdXJEniebBS6SK_6XVazPQLnak7ETb1VCNC_PMcYZEkuNI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLWQYG7m4krfHoN62QdXJEniebBS6SK_6XVazPQLnak7ETb1VCNC_PMcYZEkuNI

ret = client.Close(fd40)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd36, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "mhb7ocD2J9zezvriPdEIgCuOMMrahuOso5MvyFu") {
  panic("wrong data returned")
}


fd55 := client.Open("/Da3xOtaW62.txt", client.O_RDWR|client.O_CREATE)
if(fd55 < 0) {
  panic("open failed")
}


ret = client.Seek(fd54, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (152) GBV3FQYXb28zaXgqNy0k_l23ZFeitsqslQ26bsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqtbTKkbCPOgSlGLDAlyqKlCdZPjfs1dYcV7_0b0hqcLumGDJf4lD1WsmIQLzBcKh8DFCc8DYT8FUI

ret = client.Write(fd32, []byte("pW5jPHIqG2AUdk5_wmfq7su93VoGr7Da5Ia"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (187) GBV3FQYXb28zaXgqNy0k_l23ZFeitsqslQ26bsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqtbTKkbCPOgSlGLDAlyqKlCdZPjfs1dYcV7_0b0hqcLumGDJf4lD1WsmIQLzBcKh8DFCc8DYT8FUIpW5jPHIqG2AUdk5_wmfq7su93VoGr7Da5Ia

fd56 := client.Open("/iQtDHYElyv.txt", client.O_RDWR|client.O_CREATE)
if(fd56 < 0) {
  panic("open failed")
}

//fd state: (113) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLWQYG7m4krfHoN62QdXJEniebBS6SK_6XVazPQLnak7ETb1VCNC_PMcYZEkuNI

ret = client.Write(fd33, []byte("9BzjvoBevWIezCCowA1QQRh9eSSMZNgSMo4N5CYVNRPw9riuaDqKW6InyOe8KrNdsm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (179) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLWQYG7m4krfHoN62QdXJEniebBS6SK_6XVazPQLnak7ETb1VCNC_PMcYZEkuNI9BzjvoBevWIezCCowA1QQRh9eSSMZNgSMo4N5CYVNRPw9riuaDqKW6InyOe8KrNdsm

buf, ret = client.Read(fd52, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd51, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd57 := client.Open("/ydk9BRFkKJ.txt", client.O_RDWR|client.O_CREATE)
if(fd57 < 0) {
  panic("open failed")
}


fd58 := client.Open("/kqZrpEP1nq.txt", client.O_RDWR|client.O_CREATE)
if(fd58 < 0) {
  panic("open failed")
}


ret = client.Close(fd54)
if(ret != 0) {
  panic("close failed")
}


fd59 := client.Open("/74mEwYxfha.txt", client.O_RDWR|client.O_CREATE)
if(fd59 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd50, []byte("HDA2GSp9eJQyurMHUHoArw6jmf5wXgq5MGitfvSpmgmxyo8g_QIZjp3P9JQIHRJJS6eMs8ZDc4kt4E9y6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) HDA2GSp9eJQyurMHUHoArw6jmf5wXgq5MGitfvSpmgmxyo8g_QIZjp3P9JQIHRJJS6eMs8ZDc4kt4E9y6

buf, ret = client.Read(fd56, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd52, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd17)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd55, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd60 := client.Open("/DMPyQLTMXr.txt", client.O_RDWR|client.O_CREATE)
if(fd60 < 0) {
  panic("open failed")
}


fd61 := client.Open("/Klubr3k6Z5.txt", client.O_RDWR|client.O_CREATE)
if(fd61 < 0) {
  panic("open failed")
}


ret = client.Seek(fd55, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd31)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd34)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd48, []byte("JPCWa6G6gXvzyjEbekdPmZ9i_l3b1f4Kgo8r_Mp6if"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) JPCWa6G6gXvzyjEbekdPmZ9i_l3b1f4Kgo8r_Mp6if

fd62 := client.Open("/H8q3MipPbg.txt", client.O_RDWR|client.O_CREATE)
if(fd62 < 0) {
  panic("open failed")
}


ret = client.Close(fd59)
if(ret != 0) {
  panic("close failed")
}


fd63 := client.Open("/_mbJ__Dujj.txt", client.O_RDWR|client.O_CREATE)
if(fd63 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd49, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd32)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd63, []byte("mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0Mr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0Mr

ret = client.Seek(fd51, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd56)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd57, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd64 := client.Open("/Da3xOtaW62.txt", client.O_RDWR|client.O_CREATE)
if(fd64 < 0) {
  panic("open failed")
}

//fd state: (53) xvqRQRFhNMmTHjIolcUSJYf7X8V362wipzyN7RukBD3nYQfzFrqLi

ret = client.Write(fd45, []byte("A2DbYRMsWtYApSn8kfslEOiKt5k4KqXFrCDeEImK3tCHjO45NeoWTkhjiMybKLy0Df_904"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) xvqRQRFhNMmTHjIolcUSJYf7X8V362wipzyN7RukBD3nYQfzFrqLiA2DbYRMsWtYApSn8kfslEOiKt5k4KqXFrCDeEImK3tCHjO45NeoWTkhjiMybKLy0Df_904

buf, ret = client.Read(fd35, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "kl") {
  panic("wrong data returned")
}


ret = client.Close(fd52)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd44)
if(ret != 0) {
  panic("close failed")
}


fd65 := client.Open("/gg2gBLYHOn.txt", client.O_RDWR|client.O_CREATE)
if(fd65 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd55, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd64, []byte("eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI8

fd66 := client.Open("/JR6nH_boRB.txt", client.O_RDWR|client.O_CREATE)
if(fd66 < 0) {
  panic("open failed")
}

//fd state: (94) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0Mr

ret = client.Write(fd63, []byte("EvETHm8Iw3pdoZPd4Llk9foVqcHqOoVjrKUphgiKtWhFeiiBxUEzaXnz9K_GXqME4EHS2r2UNIFW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (170) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0MrEvETHm8Iw3pdoZPd4Llk9foVqcHqOoVjrKUphgiKtWhFeiiBxUEzaXnz9K_GXqME4EHS2r2UNIFW

buf, ret = client.Read(fd50, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd67 := client.Open("/qgChWbDgrs.txt", client.O_RDWR|client.O_CREATE)
if(fd67 < 0) {
  panic("open failed")
}

//fd state: (0) ThkPCefQXKEHif4Eoh_PbVib9QAHWEwc4g8IZAaxxyvcKvj1CXx1xXfwwvhs1QxmTjtwPZEKiWx_9AMIlqmhb7ocD2J9zezvriPdEIgCuOMMrahuOso5MvyFu

ret = client.Write(fd66, []byte("Tpte3x0Ff3B87In5r91QWwL5OjDdwC8e_mqMqHAa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) Tpte3x0Ff3B87In5r91QWwL5OjDdwC8e_mqMqHAaxyvcKvj1CXx1xXfwwvhs1QxmTjtwPZEKiWx_9AMIlqmhb7ocD2J9zezvriPdEIgCuOMMrahuOso5MvyFu

ret = client.Seek(fd47, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd58)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd48, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


buf, ret = client.Read(fd36, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd68 := client.Open("/rWIIG4WJd_.txt", client.O_RDWR|client.O_CREATE)
if(fd68 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd51, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd45, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd69 := client.Open("/DMPyQLTMXr.txt", client.O_RDWR|client.O_CREATE)
if(fd69 < 0) {
  panic("open failed")
}


fd70 := client.Open("/otUQRxsRyj.txt", client.O_RDWR|client.O_CREATE)
if(fd70 < 0) {
  panic("open failed")
}


ret = client.Close(fd36)
if(ret != 0) {
  panic("close failed")
}


fd71 := client.Open("/OmH6ue7qfz.txt", client.O_RDWR|client.O_CREATE)
if(fd71 < 0) {
  panic("open failed")
}


fd72 := client.Open("/OlMsanXw7a.txt", client.O_RDWR|client.O_CREATE)
if(fd72 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd55, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI8") {
  panic("wrong data returned")
}


ret = client.Close(fd70)
if(ret != 0) {
  panic("close failed")
}

//fd state: (46) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI8

ret = client.Write(fd64, []byte("6h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNE

buf, ret = client.Read(fd35, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "kCRyZmkVxy8n5993tqecFEENUtM0LKL9daCweorYRiw1vCtuhzr") {
  panic("wrong data returned")
}


ret = client.Close(fd51)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd68)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd47, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (179) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLWQYG7m4krfHoN62QdXJEniebBS6SK_6XVazPQLnak7ETb1VCNC_PMcYZEkuNI9BzjvoBevWIezCCowA1QQRh9eSSMZNgSMo4N5CYVNRPw9riuaDqKW6InyOe8KrNdsm

ret = client.Write(fd33, []byte("xxyxIiE4jJ8vYZCBoMrMpYqgGoPgQHkJRw5KBdd4f9VFu4rjdKL23dlU44xzX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (240) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLWQYG7m4krfHoN62QdXJEniebBS6SK_6XVazPQLnak7ETb1VCNC_PMcYZEkuNI9BzjvoBevWIezCCowA1QQRh9eSSMZNgSMo4N5CYVNRPw9riuaDqKW6InyOe8KrNdsmxxyxIiE4jJ8vYZCBoMrMpYqgGoPgQHkJRw5KBdd4f9VFu4rjdKL23dlU44xzX
//fd state: (0) Tpte3x0Ff3B87In5r91QWwL5OjDdwC8e_mqMqHAaxyvcKvj1CXx1xXfwwvhs1QxmTjtwPZEKiWx_9AMIlqmhb7ocD2J9zezvriPdEIgCuOMMrahuOso5MvyFu

ret = client.Write(fd53, []byte("X9foy10Xnd93eeA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) X9foy10Xnd93eeA5r91QWwL5OjDdwC8e_mqMqHAaxyvcKvj1CXx1xXfwwvhs1QxmTjtwPZEKiWx_9AMIlqmhb7ocD2J9zezvriPdEIgCuOMMrahuOso5MvyFu

ret = client.Seek(fd47, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd73 := client.Open("/gA71a_GvUv.txt", client.O_RDWR|client.O_CREATE)
if(fd73 < 0) {
  panic("open failed")
}

//fd state: (83) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNE

ret = client.Write(fd64, []byte("G_p18LgggmGjr0G9xzknUOlo4MYax2kRVAUGlnxiG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNEG_p18LgggmGjr0G9xzknUOlo4MYax2kRVAUGlnxiG

buf, ret = client.Read(fd35, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MgExf7nEaJzo4_GvfEs6uLJOccAZS4np8xhJ2TSJaDZ9qFpigSitrkd") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd48, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "r_") {
  panic("wrong data returned")
}


ret = client.Seek(fd73, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd66, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xyvcKvj1CXx1xXfwwvhs1QxmTjtwPZEKiWx_9AMIlqm") {
  panic("wrong data returned")
}


fd74 := client.Open("/OmH6ue7qfz.txt", client.O_RDWR|client.O_CREATE)
if(fd74 < 0) {
  panic("open failed")
}


ret = client.Seek(fd64, 91, client.SEEK_SET)
if(ret != 91) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 91)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd67, []byte("Ktyb2PoDE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) Ktyb2PoDE

ret = client.Seek(fd73, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd57, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd74)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd72)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd71, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd67)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd69, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd75 := client.Open("/oT8HBbhOs3.txt", client.O_RDWR|client.O_CREATE)
if(fd75 < 0) {
  panic("open failed")
}


fd76 := client.Open("/OmH6ue7qfz.txt", client.O_RDWR|client.O_CREATE)
if(fd76 < 0) {
  panic("open failed")
}


fd77 := client.Open("/di6atYCfA6.txt", client.O_RDWR|client.O_CREATE)
if(fd77 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd49, []byte("RjvXrtHhNrBlRD6F7gM0K5ZzjCvSEy9B1F9rWfkxfHIqW0bziEWlWXHlAo4ntxFCV3aG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) RjvXrtHhNrBlRD6F7gM0K5ZzjCvSEy9B1F9rWfkxfHIqW0bziEWlWXHlAo4ntxFCV3aG

ret = client.Close(fd53)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd45)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd73, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd78 := client.Open("/SgtN9yKcDb.txt", client.O_RDWR|client.O_CREATE)
if(fd78 < 0) {
  panic("open failed")
}

//fd state: (240) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLWQYG7m4krfHoN62QdXJEniebBS6SK_6XVazPQLnak7ETb1VCNC_PMcYZEkuNI9BzjvoBevWIezCCowA1QQRh9eSSMZNgSMo4N5CYVNRPw9riuaDqKW6InyOe8KrNdsmxxyxIiE4jJ8vYZCBoMrMpYqgGoPgQHkJRw5KBdd4f9VFu4rjdKL23dlU44xzX

ret = client.Write(fd33, []byte("_M0EImPcfPiBdVCvGnIOa7Q7L4Tg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (268) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLWQYG7m4krfHoN62QdXJEniebBS6SK_6XVazPQLnak7ETb1VCNC_PMcYZEkuNI9BzjvoBevWIezCCowA1QQRh9eSSMZNgSMo4N5CYVNRPw9riuaDqKW6InyOe8KrNdsmxxyxIiE4jJ8vYZCBoMrMpYqgGoPgQHkJRw5KBdd4f9VFu4rjdKL23dlU44xzX_M0EImPcfPiBdVCvGnIOa7Q7L4Tg

fd79 := client.Open("/Z2evREiOSO.txt", client.O_RDWR|client.O_CREATE)
if(fd79 < 0) {
  panic("open failed")
}


ret = client.Close(fd50)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd66, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


fd80 := client.Open("/gg2gBLYHOn.txt", client.O_RDWR|client.O_CREATE)
if(fd80 < 0) {
  panic("open failed")
}

//fd state: (36) X9foy10Xnd93eeA5r91QWwL5OjDdwC8e_mqMqHAaxyvcKvj1CXx1xXfwwvhs1QxmTjtwPZEKiWx_9AMIlqmhb7ocD2J9zezvriPdEIgCuOMMrahuOso5MvyFu

ret = client.Write(fd66, []byte("depWmfIw5ko5QmxeG_cqjjItKsAy4bQK7bLTyhqFXMNQROb8HpaJB32e"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) X9foy10Xnd93eeA5r91QWwL5OjDdwC8e_mqMdepWmfIw5ko5QmxeG_cqjjItKsAy4bQK7bLTyhqFXMNQROb8HpaJB32ezezvriPdEIgCuOMMrahuOso5MvyFu
//fd state: (0) 9POj0HsGwb9FJK9vzSKIpcK_HHn8ORe_nlLKL_YKpmFw9oH8gA93deNbtC4jUx5Kijv

ret = client.Write(fd79, []byte("GKveKr5wb3_boCmkLVlf577PKuNy5qge6jr6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) GKveKr5wb3_boCmkLVlf577PKuNy5qge6jr6L_YKpmFw9oH8gA93deNbtC4jUx5Kijv

fd81 := client.Open("/DWICQRTyV5.txt", client.O_RDWR|client.O_CREATE)
if(fd81 < 0) {
  panic("open failed")
}

//fd state: (36) GKveKr5wb3_boCmkLVlf577PKuNy5qge6jr6L_YKpmFw9oH8gA93deNbtC4jUx5Kijv

ret = client.Write(fd79, []byte("IxG4XK9GsA7KuCgoE0a9uL526zyN9NchY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) GKveKr5wb3_boCmkLVlf577PKuNy5qge6jr6IxG4XK9GsA7KuCgoE0a9uL526zyN9NchY

buf, ret = client.Read(fd81, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd75)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd63, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (68) RjvXrtHhNrBlRD6F7gM0K5ZzjCvSEy9B1F9rWfkxfHIqW0bziEWlWXHlAo4ntxFCV3aG

ret = client.Write(fd49, []byte("ug_BJTsmLrKN5OMrVLETKQCHtiykk7pkiyklff5NE00yl40kbjAxSK4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) RjvXrtHhNrBlRD6F7gM0K5ZzjCvSEy9B1F9rWfkxfHIqW0bziEWlWXHlAo4ntxFCV3aGug_BJTsmLrKN5OMrVLETKQCHtiykk7pkiyklff5NE00yl40kbjAxSK4

ret = client.Close(fd69)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd61)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd66)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd35)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd71, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd55, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "6h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNEG_p18LgggmGjr0G9xzknUOlo4") {
  panic("wrong data returned")
}


fd82 := client.Open("/rWIIG4WJd_.txt", client.O_RDWR|client.O_CREATE)
if(fd82 < 0) {
  panic("open failed")
}


fd83 := client.Open("/pB99UkEeea.txt", client.O_RDWR|client.O_CREATE)
if(fd83 < 0) {
  panic("open failed")
}


fd84 := client.Open("/B_hZtR1Qiv.txt", client.O_RDWR|client.O_CREATE)
if(fd84 < 0) {
  panic("open failed")
}


fd85 := client.Open("/YfmUYg5AyQ.txt", client.O_RDWR|client.O_CREATE)
if(fd85 < 0) {
  panic("open failed")
}


ret = client.Close(fd77)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd63, 166, client.SEEK_SET)
if(ret != 166) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 166)
  panic("seek failed")
}


fd86 := client.Open("/Z2evREiOSO.txt", client.O_RDWR|client.O_CREATE)
if(fd86 < 0) {
  panic("open failed")
}


ret = client.Seek(fd47, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd48, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Mp6if") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd81, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd87 := client.Open("/4ldmhfQNz1.txt", client.O_RDWR|client.O_CREATE)
if(fd87 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd87, []byte("BOx3SouSO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) BOx3SouSO

ret = client.Close(fd55)
if(ret != 0) {
  panic("close failed")
}


fd88 := client.Open("/SgtN9yKcDb.txt", client.O_RDWR|client.O_CREATE)
if(fd88 < 0) {
  panic("open failed")
}


ret = client.Seek(fd88, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd73, []byte("FHTzJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) FHTzJ

buf, ret = client.Read(fd48, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd48, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Close(fd88)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd87)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd73, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd81)
if(ret != 0) {
  panic("close failed")
}

//fd state: (166) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0MrEvETHm8Iw3pdoZPd4Llk9foVqcHqOoVjrKUphgiKtWhFeiiBxUEzaXnz9K_GXqME4EHS2r2UNIFW

ret = client.Write(fd63, []byte("wjZ2uBZWjVi7VERL1C2sju745RTQMwV_PDwwJv7c2UmvwII1PBNF4AIXDcCLlkGlH1IIf_whpaEXSYHOgRP8Cvq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (253) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0MrEvETHm8Iw3pdoZPd4Llk9foVqcHqOoVjrKUphgiKtWhFeiiBxUEzaXnz9K_GXqME4EHS2r2UwjZ2uBZWjVi7VERL1C2sju745RTQMwV_PDwwJv7c2UmvwII1PBNF4AIXDcCLlkGlH1IIf_whpaEXSYHOgRP8Cvq

ret = client.Seek(fd85, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


ret = client.Close(fd49)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd33, 153, client.SEEK_SET)
if(ret != 153) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 153)
  panic("seek failed")
}


buf, ret = client.Read(fd60, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd89 := client.Open("/jyOocvYJ4z.txt", client.O_RDWR|client.O_CREATE)
if(fd89 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd62, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd76, []byte("0NUSBWlxgvZCB_gsv7K2iMnSw1tkb05OxzISIxbdt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) 0NUSBWlxgvZCB_gsv7K2iMnSw1tkb05OxzISIxbdt

buf, ret = client.Read(fd47, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd80, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd57, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (5) JPCWa6G6gXvzyjEbekdPmZ9i_l3b1f4Kgo8r_Mp6if

ret = client.Write(fd48, []byte("DecnaoiDYuS0gKsqy8Zw63R"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) JPCWaDecnaoiDYuS0gKsqy8Zw63R1f4Kgo8r_Mp6if

fd90 := client.Open("/fncK2K1IOC.txt", client.O_RDWR|client.O_CREATE)
if(fd90 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd79, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd91 := client.Open("/LK829EmRZQ.txt", client.O_RDWR|client.O_CREATE)
if(fd91 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd47, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd92 := client.Open("/XyzWAQDWkU.txt", client.O_RDWR|client.O_CREATE)
if(fd92 < 0) {
  panic("open failed")
}


ret = client.Close(fd85)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd86, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


ret = client.Seek(fd86, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


buf, ret = client.Read(fd63, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (91) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNEG_p18LgggmGjr0G9xzknUOlo4MYax2kRVAUGlnxiG

ret = client.Write(fd64, []byte("860ZaWJPvsjEV9Eq4WpzM2BPtv0mGjHd6KUNDlW2Yl1c6ytehoM4wi5tYeF5cDMvBAsTw7iJqRN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (166) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNEG_p18Lgg860ZaWJPvsjEV9Eq4WpzM2BPtv0mGjHd6KUNDlW2Yl1c6ytehoM4wi5tYeF5cDMvBAsTw7iJqRN

ret = client.Seek(fd71, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


buf, ret = client.Read(fd62, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd47)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd33, 102, client.SEEK_SET)
if(ret != 102) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 102)
  panic("seek failed")
}


ret = client.Seek(fd57, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd92, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd93 := client.Open("/Da3xOtaW62.txt", client.O_RDWR|client.O_CREATE)
if(fd93 < 0) {
  panic("open failed")
}


ret = client.Close(fd90)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd83, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd64, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (102) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLWQYG7m4krfHoN62QdXJEniebBS6SK_6XVazPQLnak7ETb1VCNC_PMcYZEkuNI9BzjvoBevWIezCCowA1QQRh9eSSMZNgSMo4N5CYVNRPw9riuaDqKW6InyOe8KrNdsmxxyxIiE4jJ8vYZCBoMrMpYqgGoPgQHkJRw5KBdd4f9VFu4rjdKL23dlU44xzX_M0EImPcfPiBdVCvGnIOa7Q7L4Tg

ret = client.Write(fd33, []byte("blRwPEW7UGzGwwjWmYJwkhCs4KA7nTzbYYsoW3wt4aSP6vETciC38kU7YWq9aZ0_epp30lJ4IPCBUUSf04RP_9xOKJ7XF8RQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (198) uRgQ_OiBdUL4aS_U6MIAssSM_Zt936a8JVfr_79qJZRqBN5ql6WLWQYG7m4krfHoN62QdXJEniebBS6SK_6XVazPQLnak7ETb1VCNCblRwPEW7UGzGwwjWmYJwkhCs4KA7nTzbYYsoW3wt4aSP6vETciC38kU7YWq9aZ0_epp30lJ4IPCBUUSf04RP_9xOKJ7XF8RQMpYqgGoPgQHkJRw5KBdd4f9VFu4rjdKL23dlU44xzX_M0EImPcfPiBdVCvGnIOa7Q7L4Tg

ret = client.Seek(fd63, 112, client.SEEK_SET)
if(ret != 112) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 112)
  panic("seek failed")
}


buf, ret = client.Read(fd60, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd60, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd65, []byte("OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgU

buf, ret = client.Read(fd78, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd94 := client.Open("/OqvlaCPI05.txt", client.O_RDWR|client.O_CREATE)
if(fd94 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd57, []byte("nnBaL1SCiqtTmA2BvDsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRNY4ILShBprPOHDd4Pz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) nnBaL1SCiqtTmA2BvDsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRNY4ILShBprPOHDd4Pz
//fd state: (0) 

ret = client.Write(fd92, []byte("9abPdTLRkIulNX0oX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) 9abPdTLRkIulNX0oX

ret = client.Close(fd57)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd64, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd95 := client.Open("/seP9Yl_hKG.txt", client.O_RDWR|client.O_CREATE)
if(fd95 < 0) {
  panic("open failed")
}


ret = client.Close(fd94)
if(ret != 0) {
  panic("close failed")
}


fd96 := client.Open("/jAMfPtHSb4.txt", client.O_RDWR|client.O_CREATE)
if(fd96 < 0) {
  panic("open failed")
}


fd97 := client.Open("/Avi_dS2O4_.txt", client.O_RDWR|client.O_CREATE)
if(fd97 < 0) {
  panic("open failed")
}


ret = client.Seek(fd96, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd98 := client.Open("/OqvlaCPI05.txt", client.O_RDWR|client.O_CREATE)
if(fd98 < 0) {
  panic("open failed")
}


ret = client.Seek(fd79, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Close(fd71)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd83, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd84, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd60, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (1) GKveKr5wb3_boCmkLVlf577PKuNy5qge6jr6IxG4XK9GsA7KuCgoE0a9uL526zyN9NchY

ret = client.Write(fd79, []byte("vgWI2qEEjzNR9GP5zZl2o9cZAsPDidWBMiIX9_gPhOLpRpjz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) GvgWI2qEEjzNR9GP5zZl2o9cZAsPDidWBMiIX9_gPhOLpRpjzCgoE0a9uL526zyN9NchY

buf, ret = client.Read(fd65, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (40) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgU

ret = client.Write(fd65, []byte("Erha1T_MHD1W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1W
//fd state: (41) 0NUSBWlxgvZCB_gsv7K2iMnSw1tkb05OxzISIxbdt

ret = client.Write(fd76, []byte("rhpQlfiGVPO618DIYFIIetS23QVihBL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) 0NUSBWlxgvZCB_gsv7K2iMnSw1tkb05OxzISIxbdtrhpQlfiGVPO618DIYFIIetS23QVihBL

ret = client.Seek(fd86, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}


fd99 := client.Open("/i7rZ5jaEkR.txt", client.O_RDWR|client.O_CREATE)
if(fd99 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd79, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "CgoE0a9uL526zyN9NchY") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd92, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (0) 3xm6oLn1kYf12wTlosat9l80Z_k76VheqslktiFBFV57y0F49QH9Mr_yVQCEyEpKJsYqaqkj_4psDdn2auoQz06Yu795i6rfYnI

ret = client.Write(fd98, []byte("44jkIiTN6S_iy9I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) 44jkIiTN6S_iy9Ilosat9l80Z_k76VheqslktiFBFV57y0F49QH9Mr_yVQCEyEpKJsYqaqkj_4psDdn2auoQz06Yu795i6rfYnI

fd100 := client.Open("/ydk9BRFkKJ.txt", client.O_RDWR|client.O_CREATE)
if(fd100 < 0) {
  panic("open failed")
}


ret = client.Close(fd95)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd83, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (28) JPCWaDecnaoiDYuS0gKsqy8Zw63R1f4Kgo8r_Mp6if

ret = client.Write(fd48, []byte("zwGvlmDSwVf1lr9YLhYi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) JPCWaDecnaoiDYuS0gKsqy8Zw63RzwGvlmDSwVf1lr9YLhYi
//fd state: (0) 

ret = client.Write(fd84, []byte("v715Dz37A5f21oV_KmIy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) v715Dz37A5f21oV_KmIy

buf, ret = client.Read(fd100, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nnBaL1SCiqtTmA2BvDsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRNY4ILShBp") {
  panic("wrong data returned")
}


ret = client.Seek(fd99, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd96, []byte("utxBQzB5n_lRdnUOo3nT6lWYcW82XY2CK7nllcuPt9thV9oLPQ9JGZ6WIF_RjukJVrAMX9D9wqe3LqZO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) utxBQzB5n_lRdnUOo3nT6lWYcW82XY2CK7nllcuPt9thV9oLPQ9JGZ6WIF_RjukJVrAMX9D9wqe3LqZO
//fd state: (112) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0MrEvETHm8Iw3pdoZPd4Llk9foVqcHqOoVjrKUphgiKtWhFeiiBxUEzaXnz9K_GXqME4EHS2r2UwjZ2uBZWjVi7VERL1C2sju745RTQMwV_PDwwJv7c2UmvwII1PBNF4AIXDcCLlkGlH1IIf_whpaEXSYHOgRP8Cvq

ret = client.Write(fd63, []byte("HVYelZGFrc7UaFjVH1YlRrO5uWLJKxmpu0RSwna33ZcOXbby3kgjoVI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (167) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0MrEvETHm8Iw3pdoZPd4LHVYelZGFrc7UaFjVH1YlRrO5uWLJKxmpu0RSwna33ZcOXbby3kgjoVIjZ2uBZWjVi7VERL1C2sju745RTQMwV_PDwwJv7c2UmvwII1PBNF4AIXDcCLlkGlH1IIf_whpaEXSYHOgRP8Cvq

ret = client.Close(fd82)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd92)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd80, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1W") {
  panic("wrong data returned")
}


ret = client.Seek(fd60, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd91, 114, client.SEEK_SET)
if(ret != 114) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 114)
  panic("seek failed")
}


buf, ret = client.Read(fd60, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd101 := client.Open("/74mEwYxfha.txt", client.O_RDWR|client.O_CREATE)
if(fd101 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd91, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Mj_jJtbWl2C9JtH") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd97, []byte("4LgZZ8tIb6AJvWLMklVa6A1Euio40rOy7j2_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) 4LgZZ8tIb6AJvWLMklVa6A1Euio40rOy7j2_

ret = client.Seek(fd76, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}

//fd state: (52) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1W

ret = client.Write(fd65, []byte("_VTr6PcbChxTGJKChscPhPuxIZz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1W_VTr6PcbChxTGJKChscPhPuxIZz

fd102 := client.Open("/XyzWAQDWkU.txt", client.O_RDWR|client.O_CREATE)
if(fd102 < 0) {
  panic("open failed")
}


fd103 := client.Open("/DWICQRTyV5.txt", client.O_RDWR|client.O_CREATE)
if(fd103 < 0) {
  panic("open failed")
}


ret = client.Close(fd79)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd78, []byte("PMo1oF2M9EEWm6plLqk1v0EoQrli"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) PMo1oF2M9EEWm6plLqk1v0EoQrli

fd104 := client.Open("/Q23IzSXTDq.txt", client.O_RDWR|client.O_CREATE)
if(fd104 < 0) {
  panic("open failed")
}

//fd state: (48) JPCWaDecnaoiDYuS0gKsqy8Zw63RzwGvlmDSwVf1lr9YLhYi

ret = client.Write(fd48, []byte("kA5484"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) JPCWaDecnaoiDYuS0gKsqy8Zw63RzwGvlmDSwVf1lr9YLhYikA5484

buf, ret = client.Read(fd65, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd105 := client.Open("/w2FXc3l6HU.txt", client.O_RDWR|client.O_CREATE)
if(fd105 < 0) {
  panic("open failed")
}


ret = client.Seek(fd93, 71, client.SEEK_SET)
if(ret != 71) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 71)
  panic("seek failed")
}


ret = client.Seek(fd80, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


fd106 := client.Open("/cFf3kxV9cS.txt", client.O_RDWR|client.O_CREATE)
if(fd106 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd78, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd107 := client.Open("/fncK2K1IOC.txt", client.O_RDWR|client.O_CREATE)
if(fd107 < 0) {
  panic("open failed")
}


ret = client.Close(fd89)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd106, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd108 := client.Open("/_BLukPDMrv.txt", client.O_RDWR|client.O_CREATE)
if(fd108 < 0) {
  panic("open failed")
}


fd109 := client.Open("/LK829EmRZQ.txt", client.O_RDWR|client.O_CREATE)
if(fd109 < 0) {
  panic("open failed")
}


ret = client.Seek(fd107, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


ret = client.Close(fd60)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd48)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd78, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd65, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}


ret = client.Seek(fd84, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd107)
if(ret != 0) {
  panic("close failed")
}


fd110 := client.Open("/jAMfPtHSb4.txt", client.O_RDWR|client.O_CREATE)
if(fd110 < 0) {
  panic("open failed")
}


fd111 := client.Open("/mFvxS3bd1N.txt", client.O_RDWR|client.O_CREATE)
if(fd111 < 0) {
  panic("open failed")
}

//fd state: (0) IZwvr3F2uAIMc2ggv4Zf9NuVyHrZd2FrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLTx1drJMj_jJtbWl2C9JtH

ret = client.Write(fd109, []byte("luB78SS7YLQzQltMjgyaZ1_CIk1mxJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) luB78SS7YLQzQltMjgyaZ1_CIk1mxJFrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLTx1drJMj_jJtbWl2C9JtH

buf, ret = client.Read(fd109, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "FrT2y") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd93, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lulJocL_wrNEG_p18Lgg860ZaWJPvsjEV9Eq4WpzM2BPtv0mGjHd6KUNDlW2Yl1c6ytehoM4wi5tYeF5cDMvBAsTw7i") {
  panic("wrong data returned")
}


ret = client.Seek(fd104, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd109, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


fd112 := client.Open("/PIpopWJsly.txt", client.O_RDWR|client.O_CREATE)
if(fd112 < 0) {
  panic("open failed")
}


ret = client.Seek(fd78, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd101, []byte("eoeeB8qZkgJtN8Lf2LOqCAeVx_0VkFGbZ98wXJJu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) eoeeB8qZkgJtN8Lf2LOqCAeVx_0VkFGbZ98wXJJu

buf, ret = client.Read(fd108, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1MrsxizlJz") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd104, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (167) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0MrEvETHm8Iw3pdoZPd4LHVYelZGFrc7UaFjVH1YlRrO5uWLJKxmpu0RSwna33ZcOXbby3kgjoVIjZ2uBZWjVi7VERL1C2sju745RTQMwV_PDwwJv7c2UmvwII1PBNF4AIXDcCLlkGlH1IIf_whpaEXSYHOgRP8Cvq

ret = client.Write(fd63, []byte("UlXRRGh0BBITxi4RuJZe9DTCw3f6B0Z5lc7s0rHq2xdIUqsVRgIWFUGy__ovJuq6K6_1TG8XTPdmQZeskloatexQibZzLDXOue"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (265) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0MrEvETHm8Iw3pdoZPd4LHVYelZGFrc7UaFjVH1YlRrO5uWLJKxmpu0RSwna33ZcOXbby3kgjoVIUlXRRGh0BBITxi4RuJZe9DTCw3f6B0Z5lc7s0rHq2xdIUqsVRgIWFUGy__ovJuq6K6_1TG8XTPdmQZeskloatexQibZzLDXOue
//fd state: (0) GBV3FQYXb28zaXgqNy0k_l23ZFeitsqslQ26bsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqtbTKkbCPOgSlGLDAlyqKlCdZPjfs1dYcV7_0b0hqcLumGDJf4lD1WsmIQLzBcKh8DFCc8DYT8FUIpW5jPHIqG2AUdk5_wmfq7su93VoGr7Da5Ia

ret = client.Write(fd105, []byte("HjCrlg8CbeArGDTwFObXD6JAt8bsCztnzRF3t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) HjCrlg8CbeArGDTwFObXD6JAt8bsCztnzRF3tsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqtbTKkbCPOgSlGLDAlyqKlCdZPjfs1dYcV7_0b0hqcLumGDJf4lD1WsmIQLzBcKh8DFCc8DYT8FUIpW5jPHIqG2AUdk5_wmfq7su93VoGr7Da5Ia

fd113 := client.Open("/x7r7aEPrq5.txt", client.O_RDWR|client.O_CREATE)
if(fd113 < 0) {
  panic("open failed")
}


ret = client.Close(fd99)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd101, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd106)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd96, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd114 := client.Open("/ZWSlf0INI_.txt", client.O_RDWR|client.O_CREATE)
if(fd114 < 0) {
  panic("open failed")
}


ret = client.Seek(fd102, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


fd115 := client.Open("/PIpopWJsly.txt", client.O_RDWR|client.O_CREATE)
if(fd115 < 0) {
  panic("open failed")
}

//fd state: (36) 4LgZZ8tIb6AJvWLMklVa6A1Euio40rOy7j2_

ret = client.Write(fd97, []byte("C0WDcu1vy5yRqNX17UtBNSHw2tmSOHwudow96e7weF7Xpk21F6kVIJJDRO70LyQjUAbZnudqYaiBC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) 4LgZZ8tIb6AJvWLMklVa6A1Euio40rOy7j2_C0WDcu1vy5yRqNX17UtBNSHw2tmSOHwudow96e7weF7Xpk21F6kVIJJDRO70LyQjUAbZnudqYaiBC

ret = client.Seek(fd62, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd103)
if(ret != 0) {
  panic("close failed")
}

//fd state: (7) v715Dz37A5f21oV_KmIy

ret = client.Write(fd84, []byte("4IWoGQVx63AjHueKVXVi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) v715Dz34IWoGQVx63AjHueKVXVi
//fd state: (7) 9abPdTLRkIulNX0oX

ret = client.Write(fd102, []byte("hlhAUmBz5WkfxBAjGoRb3E0YuxjK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) 9abPdTLhlhAUmBz5WkfxBAjGoRb3E0YuxjK
//fd state: (129) luB78SS7YLQzQltMjgyaZ1_CIk1mxJFrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLTx1drJMj_jJtbWl2C9JtH

ret = client.Write(fd91, []byte("xa2gb5XhxxvnvbgUNmAMu7U5p5VDiv4ir5tBqIRRyXGoW6EBOOqQNcL2GsqhOGoA9FpnjtC8dU1cof2MA2sUEfG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (216) luB78SS7YLQzQltMjgyaZ1_CIk1mxJFrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLTx1drJMj_jJtbWl2C9JtHxa2gb5XhxxvnvbgUNmAMu7U5p5VDiv4ir5tBqIRRyXGoW6EBOOqQNcL2GsqhOGoA9FpnjtC8dU1cof2MA2sUEfG

fd116 := client.Open("/jRPMyJY9NL.txt", client.O_RDWR|client.O_CREATE)
if(fd116 < 0) {
  panic("open failed")
}


ret = client.Seek(fd91, 123, client.SEEK_SET)
if(ret != 123) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 123)
  panic("seek failed")
}


ret = client.Close(fd104)
if(ret != 0) {
  panic("close failed")
}


fd117 := client.Open("/PIpopWJsly.txt", client.O_RDWR|client.O_CREATE)
if(fd117 < 0) {
  panic("open failed")
}


ret = client.Close(fd86)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd91, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2C9JtHxa2gb5XhxxvnvbgUNmAMu7U5p5VDiv4ir5tBqIRRyXGoW6EBOOqQNcL2Gsqh") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd96, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd101, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd101)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) JPCWaDecnaoiDYuS0gKsqy8Zw63RzwGvlmDSwVf1lr9YLhYikA5484

ret = client.Write(fd111, []byte("WZqHUkaK4jCQKj8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) WZqHUkaK4jCQKj8S0gKsqy8Zw63RzwGvlmDSwVf1lr9YLhYikA5484

fd118 := client.Open("/kqZrpEP1nq.txt", client.O_RDWR|client.O_CREATE)
if(fd118 < 0) {
  panic("open failed")
}


ret = client.Seek(fd76, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}

//fd state: (189) luB78SS7YLQzQltMjgyaZ1_CIk1mxJFrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLTx1drJMj_jJtbWl2C9JtHxa2gb5XhxxvnvbgUNmAMu7U5p5VDiv4ir5tBqIRRyXGoW6EBOOqQNcL2GsqhOGoA9FpnjtC8dU1cof2MA2sUEfG

ret = client.Write(fd91, []byte("q4mA2BwwEVjoK_CG20tFVifAfohqCXcuBeZa25R6oDijnusWaPpACJ59xko2RkY9YzjnK2tru0CgxIHqrA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (271) luB78SS7YLQzQltMjgyaZ1_CIk1mxJFrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLTx1drJMj_jJtbWl2C9JtHxa2gb5XhxxvnvbgUNmAMu7U5p5VDiv4ir5tBqIRRyXGoW6EBOOqQNcL2Gsqhq4mA2BwwEVjoK_CG20tFVifAfohqCXcuBeZa25R6oDijnusWaPpACJ59xko2RkY9YzjnK2tru0CgxIHqrA

buf, ret = client.Read(fd109, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rT2yX3fw") {
  panic("wrong data returned")
}

//fd state: (265) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0MrEvETHm8Iw3pdoZPd4LHVYelZGFrc7UaFjVH1YlRrO5uWLJKxmpu0RSwna33ZcOXbby3kgjoVIUlXRRGh0BBITxi4RuJZe9DTCw3f6B0Z5lc7s0rHq2xdIUqsVRgIWFUGy__ovJuq6K6_1TG8XTPdmQZeskloatexQibZzLDXOue

ret = client.Write(fd63, []byte("uqOZb_YoB80rM8jLo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (282) mBgHPD9a60lNwoP7ZrogHCCt6dd1tlCllTY0nIB6rSORdhfysfH0SsZHidzXpt89mugnQiJDKZ49kTsdosZVud3YwLz0MrEvETHm8Iw3pdoZPd4LHVYelZGFrc7UaFjVH1YlRrO5uWLJKxmpu0RSwna33ZcOXbby3kgjoVIUlXRRGh0BBITxi4RuJZe9DTCw3f6B0Z5lc7s0rHq2xdIUqsVRgIWFUGy__ovJuq6K6_1TG8XTPdmQZeskloatexQibZzLDXOueuqOZb_YoB80rM8jLo

buf, ret = client.Read(fd63, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd119 := client.Open("/jwmw0Yfr5L.txt", client.O_RDWR|client.O_CREATE)
if(fd119 < 0) {
  panic("open failed")
}


fd120 := client.Open("/jAMfPtHSb4.txt", client.O_RDWR|client.O_CREATE)
if(fd120 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd78, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oF2M9EEWm6plLqk1v0EoQrli") {
  panic("wrong data returned")
}


ret = client.Seek(fd116, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd121 := client.Open("/LsRInew_XM.txt", client.O_RDWR|client.O_CREATE)
if(fd121 < 0) {
  panic("open failed")
}


ret = client.Seek(fd63, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


ret = client.Seek(fd109, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


buf, ret = client.Read(fd114, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd98, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "losat9l80Z_k76VheqslktiFBFV57y0F49QH9Mr_yVQCEyEpKJsYqaqkj_") {
  panic("wrong data returned")
}


ret = client.Seek(fd80, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd121)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd73, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd115, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd122 := client.Open("/8lrlg9FbK5.txt", client.O_RDWR|client.O_CREATE)
if(fd122 < 0) {
  panic("open failed")
}


fd123 := client.Open("/XyzWAQDWkU.txt", client.O_RDWR|client.O_CREATE)
if(fd123 < 0) {
  panic("open failed")
}


fd124 := client.Open("/fncK2K1IOC.txt", client.O_RDWR|client.O_CREATE)
if(fd124 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd119, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd78)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd112, []byte("DuTgaI19rxYDs8bDc0lxXKGPFyiQJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) DuTgaI19rxYDs8bDc0lxXKGPFyiQJ

ret = client.Seek(fd98, 64, client.SEEK_SET)
if(ret != 64) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 64)
  panic("seek failed")
}


buf, ret = client.Read(fd76, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lfiGVPO618DIYFIIetS23QVi") {
  panic("wrong data returned")
}


ret = client.Seek(fd108, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


fd125 := client.Open("/FGjIWf9TaF.txt", client.O_RDWR|client.O_CREATE)
if(fd125 < 0) {
  panic("open failed")
}


ret = client.Seek(fd113, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd126 := client.Open("/AWEExtr57V.txt", client.O_RDWR|client.O_CREATE)
if(fd126 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd114, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd127 := client.Open("/bFM001bg35.txt", client.O_RDWR|client.O_CREATE)
if(fd127 < 0) {
  panic("open failed")
}

//fd state: (166) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNEG_p18Lgg860ZaWJPvsjEV9Eq4WpzM2BPtv0mGjHd6KUNDlW2Yl1c6ytehoM4wi5tYeF5cDMvBAsTw7iJqRN

ret = client.Write(fd64, []byte("wMMEv8Zj72FRbpKD1ooPgN1ohLxqUsIE9EoovNjDytYZUXs5FB4zJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (219) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNEG_p18Lgg860ZaWJPvsjEV9Eq4WpzM2BPtv0mGjHd6KUNDlW2Yl1c6ytehoM4wi5tYeF5cDMvBAsTw7iJqRNwMMEv8Zj72FRbpKD1ooPgN1ohLxqUsIE9EoovNjDytYZUXs5FB4zJ

ret = client.Close(fd84)
if(ret != 0) {
  panic("close failed")
}


fd128 := client.Open("/V_46nNBIIS.txt", client.O_RDWR|client.O_CREATE)
if(fd128 < 0) {
  panic("open failed")
}


ret = client.Seek(fd97, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


buf, ret = client.Read(fd98, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JsYqaqkj_4psDdn2auoQz06Yu795i6rfYnI") {
  panic("wrong data returned")
}


ret = client.Close(fd118)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd127)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd100, 66, client.SEEK_SET)
if(ret != 66) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 66)
  panic("seek failed")
}


ret = client.Seek(fd126, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd73, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd76)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd120)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd125)
if(ret != 0) {
  panic("close failed")
}


fd129 := client.Open("/bQtVNGcMxc.txt", client.O_RDWR|client.O_CREATE)
if(fd129 < 0) {
  panic("open failed")
}

//fd state: (271) luB78SS7YLQzQltMjgyaZ1_CIk1mxJFrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLTx1drJMj_jJtbWl2C9JtHxa2gb5XhxxvnvbgUNmAMu7U5p5VDiv4ir5tBqIRRyXGoW6EBOOqQNcL2Gsqhq4mA2BwwEVjoK_CG20tFVifAfohqCXcuBeZa25R6oDijnusWaPpACJ59xko2RkY9YzjnK2tru0CgxIHqrA

ret = client.Write(fd91, []byte("NwbAAb6XC8LVshKfSZ3HD4eTzNcdB1ood2wEmhtxPFda_NP52_9Q_r0JtTvpq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (332) luB78SS7YLQzQltMjgyaZ1_CIk1mxJFrT2yX3fw2WDPj8zCILUjsCvdCxByqkrcHcLiuMHh1671YyF7ihiLQFWEJBPej9LxTGWDNoLdw6BkLTx1drJMj_jJtbWl2C9JtHxa2gb5XhxxvnvbgUNmAMu7U5p5VDiv4ir5tBqIRRyXGoW6EBOOqQNcL2Gsqhq4mA2BwwEVjoK_CG20tFVifAfohqCXcuBeZa25R6oDijnusWaPpACJ59xko2RkY9YzjnK2tru0CgxIHqrANwbAAb6XC8LVshKfSZ3HD4eTzNcdB1ood2wEmhtxPFda_NP52_9Q_r0JtTvpq

ret = client.Seek(fd123, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


fd130 := client.Open("/YbYhTNaGQb.txt", client.O_RDWR|client.O_CREATE)
if(fd130 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd97, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "klVa6A1") {
  panic("wrong data returned")
}


fd131 := client.Open("/FS0Irk9hhz.txt", client.O_RDWR|client.O_CREATE)
if(fd131 < 0) {
  panic("open failed")
}


ret = client.Seek(fd117, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Close(fd122)
if(ret != 0) {
  panic("close failed")
}

//fd state: (15) WZqHUkaK4jCQKj8S0gKsqy8Zw63RzwGvlmDSwVf1lr9YLhYikA5484

ret = client.Write(fd111, []byte("3j2zrYKcLlCVp_fS12TmakDiP2kJ3aoN77NU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) WZqHUkaK4jCQKj83j2zrYKcLlCVp_fS12TmakDiP2kJ3aoN77NU484

buf, ret = client.Read(fd116, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd132 := client.Open("/drnz2FxSfe.txt", client.O_RDWR|client.O_CREATE)
if(fd132 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd65, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1W_VTr6PcbChxTGJKChscPh") {
  panic("wrong data returned")
}


ret = client.Close(fd93)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd91, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd133 := client.Open("/EzNlejskyS.txt", client.O_RDWR|client.O_CREATE)
if(fd133 < 0) {
  panic("open failed")
}


ret = client.Close(fd128)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd98)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd64, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd134 := client.Open("/jRPMyJY9NL.txt", client.O_RDWR|client.O_CREATE)
if(fd134 < 0) {
  panic("open failed")
}

//fd state: (219) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNEG_p18Lgg860ZaWJPvsjEV9Eq4WpzM2BPtv0mGjHd6KUNDlW2Yl1c6ytehoM4wi5tYeF5cDMvBAsTw7iJqRNwMMEv8Zj72FRbpKD1ooPgN1ohLxqUsIE9EoovNjDytYZUXs5FB4zJ

ret = client.Write(fd64, []byte("X5E5DzDFaLxWUhe_tFV7TMlNNNgTeAb7CibOQLTXN42ce3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (265) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNEG_p18Lgg860ZaWJPvsjEV9Eq4WpzM2BPtv0mGjHd6KUNDlW2Yl1c6ytehoM4wi5tYeF5cDMvBAsTw7iJqRNwMMEv8Zj72FRbpKD1ooPgN1ohLxqUsIE9EoovNjDytYZUXs5FB4zJX5E5DzDFaLxWUhe_tFV7TMlNNNgTeAb7CibOQLTXN42ce3
//fd state: (0) 

ret = client.Write(fd130, []byte("1LC7RoKzOuHe3xV3zJI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) 1LC7RoKzOuHe3xV3zJI

ret = client.Close(fd91)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd112)
if(ret != 0) {
  panic("close failed")
}


fd135 := client.Open("/gg2gBLYHOn.txt", client.O_RDWR|client.O_CREATE)
if(fd135 < 0) {
  panic("open failed")
}


fd136 := client.Open("/Da3xOtaW62.txt", client.O_RDWR|client.O_CREATE)
if(fd136 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd62, []byte("gi_anMd4cVNPBFGZ7A3z9d5of_JhyNJBYyB5RlTrDwPFYducNWwdTknZLR13dFFA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) gi_anMd4cVNPBFGZ7A3z9d5of_JhyNJBYyB5RlTrDwPFYducNWwdTknZLR13dFFA
//fd state: (5) FHTzJ

ret = client.Write(fd73, []byte("M2ues"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) FHTzJM2ues

ret = client.Close(fd97)
if(ret != 0) {
  panic("close failed")
}


fd137 := client.Open("/FSTYY4xE5x.txt", client.O_RDWR|client.O_CREATE)
if(fd137 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd137, []byte("bxdNzy37Aa0DP5mmPY7J8CK7C5oS2SMmmE8jQxyRA6ILmg9vDTBLDrUmC3b57ETMatdSHZX2f1hQ7XrWLvRacLFQz_N4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) bxdNzy37Aa0DP5mmPY7J8CK7C5oS2SMmmE8jQxyRA6ILmg9vDTBLDrUmC3b57ETMatdSHZX2f1hQ7XrWLvRacLFQz_N4
//fd state: (51) WZqHUkaK4jCQKj83j2zrYKcLlCVp_fS12TmakDiP2kJ3aoN77NU484

ret = client.Write(fd111, []byte("tQTpva44CxAci_NHE4jfVcl3Teda33hXOkGJRzQiA9tHY_V5fcV8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) WZqHUkaK4jCQKj83j2zrYKcLlCVp_fS12TmakDiP2kJ3aoN77NUtQTpva44CxAci_NHE4jfVcl3Teda33hXOkGJRzQiA9tHY_V5fcV8

fd138 := client.Open("/SZ3w7aLh_u.txt", client.O_RDWR|client.O_CREATE)
if(fd138 < 0) {
  panic("open failed")
}


ret = client.Close(fd63)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd133, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd80)
if(ret != 0) {
  panic("close failed")
}


fd139 := client.Open("/YYMVsNBZSh.txt", client.O_RDWR|client.O_CREATE)
if(fd139 < 0) {
  panic("open failed")
}


ret = client.Close(fd105)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd83, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd73)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd115)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) zT_Y526EQfhTCxT4DgOu7y3WIW9leNJMeixI0yB5iFmiMmo081HEG_0mXMQJlC5dC__zEajYEG0ZITFY

ret = client.Write(fd124, []byte("s8WBl3LzAiq2bkhL32EELxo_gAVjtFO0vmOEhp_H3o_o"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) s8WBl3LzAiq2bkhL32EELxo_gAVjtFO0vmOEhp_H3o_oMmo081HEG_0mXMQJlC5dC__zEajYEG0ZITFY

ret = client.Seek(fd135, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


buf, ret = client.Read(fd137, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd135, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


ret = client.Seek(fd102, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd131, []byte("8tvZirCd71oK5O9A33mWu1IyugSANxpK0IIVeWAPxnsZnDzYTlfTA6bBx4YYK51"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) 8tvZirCd71oK5O9A33mWu1IyugSANxpK0IIVeWAPxnsZnDzYTlfTA6bBx4YYK51

ret = client.Close(fd129)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd33, 93, client.SEEK_SET)
if(ret != 93) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 93)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd113, []byte("FVufQKNseI93OoaNQW61pdh5Sb9b3PiDmzBMgWwIklMDOPOElzeuCNYBB08fiUxSsm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) FVufQKNseI93OoaNQW61pdh5Sb9b3PiDmzBMgWwIklMDOPOElzeuCNYBB08fiUxSsm

ret = client.Close(fd116)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd126, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd140 := client.Open("/YbYhTNaGQb.txt", client.O_RDWR|client.O_CREATE)
if(fd140 < 0) {
  panic("open failed")
}

//fd state: (19) 1LC7RoKzOuHe3xV3zJI

ret = client.Write(fd130, []byte("t_lL45wMqR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) 1LC7RoKzOuHe3xV3zJIt_lL45wMqR
//fd state: (7) 1MrsxizlJz

ret = client.Write(fd108, []byte("DXboeLrW7WR57R3UqVwaXEK7AsMGzcXdipgMQGgdj2lXLPwgKZUZhKUohriB2AIyp5NYnUpQa5NjXePvK4meHqPEED7L"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) 1MrsxizDXboeLrW7WR57R3UqVwaXEK7AsMGzcXdipgMQGgdj2lXLPwgKZUZhKUohriB2AIyp5NYnUpQa5NjXePvK4meHqPEED7L
//fd state: (265) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNEG_p18Lgg860ZaWJPvsjEV9Eq4WpzM2BPtv0mGjHd6KUNDlW2Yl1c6ytehoM4wi5tYeF5cDMvBAsTw7iJqRNwMMEv8Zj72FRbpKD1ooPgN1ohLxqUsIE9EoovNjDytYZUXs5FB4zJX5E5DzDFaLxWUhe_tFV7TMlNNNgTeAb7CibOQLTXN42ce3

ret = client.Write(fd64, []byte("Th4Ge1ohWCBGYj3SNV_oiTHwfTki3NcMOWEVKInxPYvEUv_wNdZCcKtZ9THRyDA0Qlr0z_9n4kgD46SGR9xB34y9azy1hv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (359) eP9nHy4I1dlxAcSSj_e5VrmiRdyIR98JoG6oXdkMjc5NI86h68jl06TubZ0i_qC2ae5QDVJlulJocL_wrNEG_p18Lgg860ZaWJPvsjEV9Eq4WpzM2BPtv0mGjHd6KUNDlW2Yl1c6ytehoM4wi5tYeF5cDMvBAsTw7iJqRNwMMEv8Zj72FRbpKD1ooPgN1ohLxqUsIE9EoovNjDytYZUXs5FB4zJX5E5DzDFaLxWUhe_tFV7TMlNNNgTeAb7CibOQLTXN42ce3Th4Ge1ohWCBGYj3SNV_oiTHwfTki3NcMOWEVKInxPYvEUv_wNdZCcKtZ9THRyDA0Qlr0z_9n4kgD46SGR9xB34y9azy1hv

ret = client.Close(fd114)
if(ret != 0) {
  panic("close failed")
}

//fd state: (103) WZqHUkaK4jCQKj83j2zrYKcLlCVp_fS12TmakDiP2kJ3aoN77NUtQTpva44CxAci_NHE4jfVcl3Teda33hXOkGJRzQiA9tHY_V5fcV8

ret = client.Write(fd111, []byte("IoZCVmJr50W_I5dkjcyQ2GwBVWRxhYXGXF8ZDWATL0LhJVviXpyLgHUPbsb1u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) WZqHUkaK4jCQKj83j2zrYKcLlCVp_fS12TmakDiP2kJ3aoN77NUtQTpva44CxAci_NHE4jfVcl3Teda33hXOkGJRzQiA9tHY_V5fcV8IoZCVmJr50W_I5dkjcyQ2GwBVWRxhYXGXF8ZDWATL0LhJVviXpyLgHUPbsb1u

fd141 := client.Open("/SgtN9yKcDb.txt", client.O_RDWR|client.O_CREATE)
if(fd141 < 0) {
  panic("open failed")
}


ret = client.Close(fd110)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd119, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd138, []byte("F8iipPi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) F8iipPi

fd142 := client.Open("/LW1OckBfGg.txt", client.O_RDWR|client.O_CREATE)
if(fd142 < 0) {
  panic("open failed")
}


ret = client.Seek(fd96, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}


fd143 := client.Open("/fOI1tyDNiH.txt", client.O_RDWR|client.O_CREATE)
if(fd143 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd126, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd140)
if(ret != 0) {
  panic("close failed")
}


fd144 := client.Open("/YfmUYg5AyQ.txt", client.O_RDWR|client.O_CREATE)
if(fd144 < 0) {
  panic("open failed")
}

//fd state: (66) FVufQKNseI93OoaNQW61pdh5Sb9b3PiDmzBMgWwIklMDOPOElzeuCNYBB08fiUxSsm

ret = client.Write(fd113, []byte("Ai6s3NXFylcYMX6TAFILcL3CYz_DWUUOd7COX7WQpXRqakAAyf3Sw790VWXRshColR6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) FVufQKNseI93OoaNQW61pdh5Sb9b3PiDmzBMgWwIklMDOPOElzeuCNYBB08fiUxSsmAi6s3NXFylcYMX6TAFILcL3CYz_DWUUOd7COX7WQpXRqakAAyf3Sw790VWXRshColR6

ret = client.Close(fd108)
if(ret != 0) {
  panic("close failed")
}

//fd state: (64) gi_anMd4cVNPBFGZ7A3z9d5of_JhyNJBYyB5RlTrDwPFYducNWwdTknZLR13dFFA

ret = client.Write(fd62, []byte("3Qq0Zgz65lReDTdZfdppOX0XqzILB_Fx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) gi_anMd4cVNPBFGZ7A3z9d5of_JhyNJBYyB5RlTrDwPFYducNWwdTknZLR13dFFA3Qq0Zgz65lReDTdZfdppOX0XqzILB_Fx

ret = client.Close(fd119)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd145 := client.Open("/kTj6SkX0xi.txt", client.O_RDWR|client.O_CREATE)
if(fd145 < 0) {
  panic("open failed")
}


fd146 := client.Open("/dfRXaFO9tW.txt", client.O_RDWR|client.O_CREATE)
if(fd146 < 0) {
  panic("open failed")
}


ret = client.Close(fd64)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd134)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd130, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd145, []byte("vb73i6FuLJQIxbE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) vb73i6FuLJQIxbE

ret = client.Close(fd137)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd130)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd142)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd100, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "prPOHDd4Pz") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd124, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Mmo081HEG_0mXMQJlC5dC__zEajYEG0ZITFY") {
  panic("wrong data returned")
}


fd147 := client.Open("/QlSfrOOHqy.txt", client.O_RDWR|client.O_CREATE)
if(fd147 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd133, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd96)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd143)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd33)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd141, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PMo1oF2M9EEWm6plLqk1v0EoQrli") {
  panic("wrong data returned")
}


ret = client.Seek(fd146, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd146, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd148 := client.Open("/ksGXj4_LYC.txt", client.O_RDWR|client.O_CREATE)
if(fd148 < 0) {
  panic("open failed")
}


ret = client.Close(fd111)
if(ret != 0) {
  panic("close failed")
}

//fd state: (7) F8iipPi

ret = client.Write(fd138, []byte("HaSATt6OZaNFa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) F8iipPiHaSATt6OZaNFa

ret = client.Close(fd113)
if(ret != 0) {
  panic("close failed")
}


fd149 := client.Open("/ti7WCCrD9X.txt", client.O_RDWR|client.O_CREATE)
if(fd149 < 0) {
  panic("open failed")
}

//fd state: (73) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1W_VTr6PcbChxTGJKChscPhPuxIZz

ret = client.Write(fd65, []byte("r1WAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7o"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (126) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1W_VTr6PcbChxTGJKChscPhr1WAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7o
//fd state: (51) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1W_VTr6PcbChxTGJKChscPhr1WAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7o

ret = client.Write(fd135, []byte("gzmbGANpt2sz3fGu6gekHaqKn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1gzmbGANpt2sz3fGu6gekHaqKnAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7o

buf, ret = client.Read(fd135, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7o") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd109, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yaZ1_CIk1mxJFrT2yX3fw") {
  panic("wrong data returned")
}

//fd state: (96) gi_anMd4cVNPBFGZ7A3z9d5of_JhyNJBYyB5RlTrDwPFYducNWwdTknZLR13dFFA3Qq0Zgz65lReDTdZfdppOX0XqzILB_Fx

ret = client.Write(fd62, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) gi_anMd4cVNPBFGZ7A3z9d5of_JhyNJBYyB5RlTrDwPFYducNWwdTknZLR13dFFA3Qq0Zgz65lReDTdZfdppOX0XqzILB_Fx
//fd state: (0) m1CC_0SgjxqfeXLB9oXh4moURNXAFkkfXa7n25m7zpbTpdydASOize3WbZXZ0bk5J

ret = client.Write(fd144, []byte("s3CiGwjJKyc0G1XeCc3nwDfEosFLGcgxsfBmu2VKlc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) s3CiGwjJKyc0G1XeCc3nwDfEosFLGcgxsfBmu2VKlcbTpdydASOize3WbZXZ0bk5J

buf, ret = client.Read(fd133, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (33) 9abPdTLhlhAUmBz5WkfxBAjGoRb3E0YuxjK

ret = client.Write(fd123, []byte("bCxFfvqdMy6kgRXPKoH9Gb_vAQSkCtE2Gy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) 9abPdTLhlhAUmBz5WkfxBAjGoRb3E0YuxbCxFfvqdMy6kgRXPKoH9Gb_vAQSkCtE2Gy

fd150 := client.Open("/EzNlejskyS.txt", client.O_RDWR|client.O_CREATE)
if(fd150 < 0) {
  panic("open failed")
}


ret = client.Close(fd124)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd141, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd151 := client.Open("/0prf8WI95c.txt", client.O_RDWR|client.O_CREATE)
if(fd151 < 0) {
  panic("open failed")
}

//fd state: (126) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1gzmbGANpt2sz3fGu6gekHaqKnAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7o

ret = client.Write(fd135, []byte("4pI_vFhv0jr4Hn35x8QLVyXL58NE0CoWTkEAhL4OofBSk0JOvhm5l"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (179) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1gzmbGANpt2sz3fGu6gekHaqKnAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7o4pI_vFhv0jr4Hn35x8QLVyXL58NE0CoWTkEAhL4OofBSk0JOvhm5l

ret = client.Seek(fd148, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd146, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd152 := client.Open("/7HWK34MAn3.txt", client.O_RDWR|client.O_CREATE)
if(fd152 < 0) {
  panic("open failed")
}


ret = client.Close(fd83)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd148, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd117)
if(ret != 0) {
  panic("close failed")
}


fd153 := client.Open("/kTj6SkX0xi.txt", client.O_RDWR|client.O_CREATE)
if(fd153 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd147, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd100)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd133, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd133)
if(ret != 0) {
  panic("close failed")
}


fd154 := client.Open("/EzNlejskyS.txt", client.O_RDWR|client.O_CREATE)
if(fd154 < 0) {
  panic("open failed")
}

//fd state: (96) gi_anMd4cVNPBFGZ7A3z9d5of_JhyNJBYyB5RlTrDwPFYducNWwdTknZLR13dFFA3Qq0Zgz65lReDTdZfdppOX0XqzILB_Fx

ret = client.Write(fd62, []byte("nHeqNxPfLSQ5oFeFFfFpbyb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (119) gi_anMd4cVNPBFGZ7A3z9d5of_JhyNJBYyB5RlTrDwPFYducNWwdTknZLR13dFFA3Qq0Zgz65lReDTdZfdppOX0XqzILB_FxnHeqNxPfLSQ5oFeFFfFpbyb

buf, ret = client.Read(fd139, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd155 := client.Open("/SIYaGzyyUN.txt", client.O_RDWR|client.O_CREATE)
if(fd155 < 0) {
  panic("open failed")
}


ret = client.Seek(fd102, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Seek(fd151, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd151)
if(ret != 0) {
  panic("close failed")
}


fd156 := client.Open("/rWIIG4WJd_.txt", client.O_RDWR|client.O_CREATE)
if(fd156 < 0) {
  panic("open failed")
}


ret = client.Close(fd146)
if(ret != 0) {
  panic("close failed")
}


fd157 := client.Open("/Q23IzSXTDq.txt", client.O_RDWR|client.O_CREATE)
if(fd157 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd148, []byte("t4fwIO4JoQMzVIM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) t4fwIO4JoQMzVIM
//fd state: (67) 9abPdTLhlhAUmBz5WkfxBAjGoRb3E0YuxbCxFfvqdMy6kgRXPKoH9Gb_vAQSkCtE2Gy

ret = client.Write(fd123, []byte("DxTUU5QYBIdOPU_VNP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) 9abPdTLhlhAUmBz5WkfxBAjGoRb3E0YuxbCxFfvqdMy6kgRXPKoH9Gb_vAQSkCtE2GyDxTUU5QYBIdOPU_VNP
//fd state: (85) 9abPdTLhlhAUmBz5WkfxBAjGoRb3E0YuxbCxFfvqdMy6kgRXPKoH9Gb_vAQSkCtE2GyDxTUU5QYBIdOPU_VNP

ret = client.Write(fd123, []byte("cD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) 9abPdTLhlhAUmBz5WkfxBAjGoRb3E0YuxbCxFfvqdMy6kgRXPKoH9Gb_vAQSkCtE2GyDxTUU5QYBIdOPU_VNPcD

ret = client.Close(fd135)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd138, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (15) vb73i6FuLJQIxbE

ret = client.Write(fd145, []byte("P9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1A

fd158 := client.Open("/0i3NHfYg9J.txt", client.O_RDWR|client.O_CREATE)
if(fd158 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd144, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bTpdydASOize3WbZXZ0bk5J") {
  panic("wrong data returned")
}


ret = client.Seek(fd145, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}

//fd state: (87) 9abPdTLhlhAUmBz5WkfxBAjGoRb3E0YuxbCxFfvqdMy6kgRXPKoH9Gb_vAQSkCtE2GyDxTUU5QYBIdOPU_VNPcD

ret = client.Write(fd123, []byte("mZWK60qAibs7t_xD38Z8DhgUsrK4nDiAKtCWHyqn6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (128) 9abPdTLhlhAUmBz5WkfxBAjGoRb3E0YuxbCxFfvqdMy6kgRXPKoH9Gb_vAQSkCtE2GyDxTUU5QYBIdOPU_VNPcDmZWK60qAibs7t_xD38Z8DhgUsrK4nDiAKtCWHyqn6
//fd state: (0) 

ret = client.Write(fd154, []byte("o1wsjwi2KPmPGgr1XZzZ0KMIoAL_JEaKZOB4dfUrPx36mEIZhZWszPQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) o1wsjwi2KPmPGgr1XZzZ0KMIoAL_JEaKZOB4dfUrPx36mEIZhZWszPQ

ret = client.Close(fd136)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd155, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd159 := client.Open("/VghMsYjnXc.txt", client.O_RDWR|client.O_CREATE)
if(fd159 < 0) {
  panic("open failed")
}


ret = client.Close(fd154)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd126, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd126, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd126, []byte("YELE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) YELE
//fd state: (0) 

ret = client.Write(fd159, []byte("kKWGrBhx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) kKWGrBhx

ret = client.Seek(fd159, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


fd160 := client.Open("/CWlKR1JLzO.txt", client.O_RDWR|client.O_CREATE)
if(fd160 < 0) {
  panic("open failed")
}


ret = client.Seek(fd132, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd126, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd161 := client.Open("/V0gWwEDbM4.txt", client.O_RDWR|client.O_CREATE)
if(fd161 < 0) {
  panic("open failed")
}


ret = client.Seek(fd157, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd132, []byte("5L2IXW3vtB79j_IH2OXo4SjY0HbfOMCh1Tu4WhOja4R7gOrAO3aGVCBJBnDiYI4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) 5L2IXW3vtB79j_IH2OXo4SjY0HbfOMCh1Tu4WhOja4R7gOrAO3aGVCBJBnDiYI4

ret = client.Seek(fd156, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd152, []byte("2DzlNJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) 2DzlNJ

fd162 := client.Open("/V0gWwEDbM4.txt", client.O_RDWR|client.O_CREATE)
if(fd162 < 0) {
  panic("open failed")
}


ret = client.Close(fd62)
if(ret != 0) {
  panic("close failed")
}


fd163 := client.Open("/4300GnPlIR.txt", client.O_RDWR|client.O_CREATE)
if(fd163 < 0) {
  panic("open failed")
}


ret = client.Seek(fd153, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


ret = client.Seek(fd109, 192, client.SEEK_SET)
if(ret != 192) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 192)
  panic("seek failed")
}

//fd state: (0) o1wsjwi2KPmPGgr1XZzZ0KMIoAL_JEaKZOB4dfUrPx36mEIZhZWszPQ

ret = client.Write(fd150, []byte("woTwOTj4doF7Z6xMoDY1IWr0z39POY0d8StjfsrfEsNUhKWfci"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) woTwOTj4doF7Z6xMoDY1IWr0z39POY0d8StjfsrfEsNUhKWfciWszPQ

buf, ret = client.Read(fd159, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rBhx") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd159, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd164 := client.Open("/UHzput6Vyv.txt", client.O_RDWR|client.O_CREATE)
if(fd164 < 0) {
  panic("open failed")
}


ret = client.Close(fd160)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd139, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd155, []byte("a_fUMcwextYd8iOlTrvqpYLGz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) a_fUMcwextYd8iOlTrvqpYLGz

ret = client.Seek(fd126, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd161, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd155, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Seek(fd141, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) _ockFJ0zIB4XUf7XOaDaKSsTth5LKGvIaIQeIq6xMJBZ3EGdHx

ret = client.Write(fd163, []byte("Zl_7nfBqenOlSiRgrFHeVLRIpCgot1lz5prcLXqXCgM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) Zl_7nfBqenOlSiRgrFHeVLRIpCgot1lz5prcLXqXCgMZ3EGdHx

ret = client.Seek(fd123, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


buf, ret = client.Read(fd138, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd126, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ELE") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd158, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd147, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd157, []byte("jOdX6B0jC5Fm5vz2fNdzP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) jOdX6B0jC5Fm5vz2fNdzP

ret = client.Close(fd138)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd149, []byte("gphVZsJMAS3ncKU0btXjxTI8mnv4jCGCc3o5e9IlkAy0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) gphVZsJMAS3ncKU0btXjxTI8mnv4jCGCc3o5e9IlkAy0

buf, ret = client.Read(fd145, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Yv2uVZQuMwRoAJSw7lErR1A") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd148, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd102)
if(ret != 0) {
  panic("close failed")
}

//fd state: (8) kKWGrBhx

ret = client.Write(fd159, []byte("5ONQXQ0hfTmF5xjIZ8lYtWeTNRwgKqXvE_iG88UEORJu7HL3Wcc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) kKWGrBhx5ONQXQ0hfTmF5xjIZ8lYtWeTNRwgKqXvE_iG88UEORJu7HL3Wcc

buf, ret = client.Read(fd132, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd141)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd161)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd145, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


ret = client.Seek(fd158, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd165 := client.Open("/ep7dA5t5Cw.txt", client.O_RDWR|client.O_CREATE)
if(fd165 < 0) {
  panic("open failed")
}


ret = client.Seek(fd155, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}

//fd state: (4) YELE

ret = client.Write(fd126, []byte("acaUs4JFnWjD7IiqKa9ARD7bX94IVIXevc7noZpodRhyhWf59NAZ_4c"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) YELEacaUs4JFnWjD7IiqKa9ARD7bX94IVIXevc7noZpodRhyhWf59NAZ_4c

fd166 := client.Open("/CcuIbz5Tou.txt", client.O_RDWR|client.O_CREATE)
if(fd166 < 0) {
  panic("open failed")
}


ret = client.Close(fd126)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd132)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd109)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd145, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AJSw7lErR") {
  panic("wrong data returned")
}


ret = client.Seek(fd144, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}

//fd state: (20) WgptUaqjIA7147_4YEzZc0U5jnV4avAVljR7_JfnG5vXCPwJilWlcMYv

ret = client.Write(fd156, []byte("bO7UNwkHC3Ai_HP_AgDtuPMPNv4QVvolZzJnIdFKIHTLvY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) WgptUaqjIA7147_4YEzZbO7UNwkHC3Ai_HP_AgDtuPMPNv4QVvolZzJnIdFKIHTLvY

ret = client.Seek(fd156, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}

//fd state: (59) kKWGrBhx5ONQXQ0hfTmF5xjIZ8lYtWeTNRwgKqXvE_iG88UEORJu7HL3Wcc

ret = client.Write(fd159, []byte("w8STn3MR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) kKWGrBhx5ONQXQ0hfTmF5xjIZ8lYtWeTNRwgKqXvE_iG88UEORJu7HL3Wccw8STn3MR

buf, ret = client.Read(fd150, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WszP") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd139, []byte("_lIMrdvTX1NCHpH3DLLpDAshKP3PmhXw41Wc_ZmUjjvi3qo4JUNwmU7n2gLkZkkMtVgkXKvpP9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) _lIMrdvTX1NCHpH3DLLpDAshKP3PmhXw41Wc_ZmUjjvi3qo4JUNwmU7n2gLkZkkMtVgkXKvpP9

ret = client.Close(fd153)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd139, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (126) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1gzmbGANpt2sz3fGu6gekHaqKnAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7o4pI_vFhv0jr4Hn35x8QLVyXL58NE0CoWTkEAhL4OofBSk0JOvhm5l

ret = client.Write(fd65, []byte("VcqeEjIstIWzYfro2Gyix_jlGr3iCDabBrdCaj2kgsR8TTWtAlS3Ph7Pi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (183) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1gzmbGANpt2sz3fGu6gekHaqKnAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7oVcqeEjIstIWzYfro2Gyix_jlGr3iCDabBrdCaj2kgsR8TTWtAlS3Ph7Pi
//fd state: (47) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1A

ret = client.Write(fd145, []byte("1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXj

buf, ret = client.Read(fd149, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (136) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXj

ret = client.Write(fd145, []byte("LLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiag"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (224) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiag

ret = client.Close(fd159)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd145, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd150)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd152, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd158, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd167 := client.Open("/Mi8VZiqzTq.txt", client.O_RDWR|client.O_CREATE)
if(fd167 < 0) {
  panic("open failed")
}

//fd state: (25) a_fUMcwextYd8iOlTrvqpYLGz

ret = client.Write(fd155, []byte("ryZyvrS8cWxEmBoRJMmxyF_YP3v7GPNItvYKtOWVQVbee26fXUhTkRs14dN13KsDRhfJA3AN0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) a_fUMcwextYd8iOlTrvqpYLGzryZyvrS8cWxEmBoRJMmxyF_YP3v7GPNItvYKtOWVQVbee26fXUhTkRs14dN13KsDRhfJA3AN0

ret = client.Seek(fd163, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


ret = client.Close(fd156)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd157, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


buf, ret = client.Read(fd148, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (15) t4fwIO4JoQMzVIM

ret = client.Write(fd148, []byte("A1xjVEvL0IhG00WFqGZNxn8CDkrwLKOw4HG_2hBEFa8RbJMznzS2lTLkzfALBJ7ttFx1FdAwWfFex_PuarDpZLSepplGX6A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) t4fwIO4JoQMzVIMA1xjVEvL0IhG00WFqGZNxn8CDkrwLKOw4HG_2hBEFa8RbJMznzS2lTLkzfALBJ7ttFx1FdAwWfFex_PuarDpZLSepplGX6A
//fd state: (4) Zl_7nfBqenOlSiRgrFHeVLRIpCgot1lz5prcLXqXCgMZ3EGdHx

ret = client.Write(fd163, []byte("dRc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) Zl_7dRcqenOlSiRgrFHeVLRIpCgot1lz5prcLXqXCgMZ3EGdHx
//fd state: (0) 

ret = client.Write(fd147, []byte("nMfUfmOTLfsKpfUip7dQNNPL2BkQMe0pWgUpGHfv5CZx2tbK_Pj6MwXZeZMsNoy0PAMaw0TCQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) nMfUfmOTLfsKpfUip7dQNNPL2BkQMe0pWgUpGHfv5CZx2tbK_Pj6MwXZeZMsNoy0PAMaw0TCQ

ret = client.Close(fd147)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd165, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd149, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd163, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd164, []byte("cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1t
//fd state: (0) 

ret = client.Write(fd167, []byte("crMqUzA79wb7zbBk3DqTw4tMYHHLJylkJZgimUahsXalff_I0lCnMAS5GB059S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) crMqUzA79wb7zbBk3DqTw4tMYHHLJylkJZgimUahsXalff_I0lCnMAS5GB059S

ret = client.Seek(fd155, 85, client.SEEK_SET)
if(ret != 85) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 85)
  panic("seek failed")
}


ret = client.Close(fd165)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd148)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd163, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LRIpCgot1lz5prcLXqXCgMZ") {
  panic("wrong data returned")
}


ret = client.Seek(fd158, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd168 := client.Open("/fOI1tyDNiH.txt", client.O_RDWR|client.O_CREATE)
if(fd168 < 0) {
  panic("open failed")
}


ret = client.Seek(fd123, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}

//fd state: (224) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiag

ret = client.Write(fd145, []byte("vSJLxhAzvUrSop8hr__YmWyQK19kga_h7SKcfR2U5IwspgpYWDLV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (276) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiagvSJLxhAzvUrSop8hr__YmWyQK19kga_h7SKcfR2U5IwspgpYWDLV

ret = client.Seek(fd155, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}

//fd state: (44) Zl_7dRcqenOlSiRgrFHeVLRIpCgot1lz5prcLXqXCgMZ3EGdHx

ret = client.Write(fd163, []byte("sFrz7T83PEFKaISdQDSmMkqnsRbe9k0hF9TJ1IimQsIZZqNK8YiQ8w7rpMGCpZctmmgrjBHs_Rs1zoi7czoiLxjt0xuFuD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (138) Zl_7dRcqenOlSiRgrFHeVLRIpCgot1lz5prcLXqXCgMZsFrz7T83PEFKaISdQDSmMkqnsRbe9k0hF9TJ1IimQsIZZqNK8YiQ8w7rpMGCpZctmmgrjBHs_Rs1zoi7czoiLxjt0xuFuD

ret = client.Seek(fd131, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


buf, ret = client.Read(fd163, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (138) Zl_7dRcqenOlSiRgrFHeVLRIpCgot1lz5prcLXqXCgMZsFrz7T83PEFKaISdQDSmMkqnsRbe9k0hF9TJ1IimQsIZZqNK8YiQ8w7rpMGCpZctmmgrjBHs_Rs1zoi7czoiLxjt0xuFuD

ret = client.Write(fd163, []byte("iFLcbt6S2BPUp3SaMKLzYLh_JeL97yBn3SnFT_FYacswxoXHqsyevbbm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (194) Zl_7dRcqenOlSiRgrFHeVLRIpCgot1lz5prcLXqXCgMZsFrz7T83PEFKaISdQDSmMkqnsRbe9k0hF9TJ1IimQsIZZqNK8YiQ8w7rpMGCpZctmmgrjBHs_Rs1zoi7czoiLxjt0xuFuDiFLcbt6S2BPUp3SaMKLzYLh_JeL97yBn3SnFT_FYacswxoXHqsyevbbm

ret = client.Close(fd158)
if(ret != 0) {
  panic("close failed")
}


fd169 := client.Open("/K1YhjaImgW.txt", client.O_RDWR|client.O_CREATE)
if(fd169 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd167, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd144, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3CiGwjJ") {
  panic("wrong data returned")
}


ret = client.Seek(fd152, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd163, 181, client.SEEK_SET)
if(ret != 181) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 181)
  panic("seek failed")
}

//fd state: (276) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiagvSJLxhAzvUrSop8hr__YmWyQK19kga_h7SKcfR2U5IwspgpYWDLV

ret = client.Write(fd145, []byte("Og1OiCMb8J8InUXJRJEYXvFJH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (301) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiagvSJLxhAzvUrSop8hr__YmWyQK19kga_h7SKcfR2U5IwspgpYWDLVOg1OiCMb8J8InUXJRJEYXvFJH
//fd state: (8) s3CiGwjJKyc0G1XeCc3nwDfEosFLGcgxsfBmu2VKlcbTpdydASOize3WbZXZ0bk5J

ret = client.Write(fd144, []byte("LhEtf9VYgVX4jQDPmrQYaS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaSgxsfBmu2VKlcbTpdydASOize3WbZXZ0bk5J

ret = client.Seek(fd166, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd157, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


ret = client.Seek(fd149, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


fd170 := client.Open("/YYMVsNBZSh.txt", client.O_RDWR|client.O_CREATE)
if(fd170 < 0) {
  panic("open failed")
}


ret = client.Seek(fd152, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


buf, ret = client.Read(fd157, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "z2fNdzP") {
  panic("wrong data returned")
}


ret = client.Seek(fd155, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


buf, ret = client.Read(fd162, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd166)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd170)
if(ret != 0) {
  panic("close failed")
}

//fd state: (42) a_fUMcwextYd8iOlTrvqpYLGzryZyvrS8cWxEmBoRJMmxyF_YP3v7GPNItvYKtOWVQVbee26fXUhTkRs14dN13KsDRhfJA3AN0

ret = client.Write(fd155, []byte("jRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) a_fUMcwextYd8iOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOn

buf, ret = client.Read(fd169, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd168, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd163, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


buf, ret = client.Read(fd169, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (30) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaSgxsfBmu2VKlcbTpdydASOize3WbZXZ0bk5J

ret = client.Write(fd144, []byte("2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1t
//fd state: (84) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1t

ret = client.Write(fd144, []byte("XhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9

fd171 := client.Open("/k06ZPn3SfU.txt", client.O_RDWR|client.O_CREATE)
if(fd171 < 0) {
  panic("open failed")
}


ret = client.Seek(fd168, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (97) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1t

ret = client.Write(fd164, []byte("Yj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (181) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719g

ret = client.Close(fd149)
if(ret != 0) {
  panic("close failed")
}


fd172 := client.Open("/Ng2aSPb5ni.txt", client.O_RDWR|client.O_CREATE)
if(fd172 < 0) {
  panic("open failed")
}


ret = client.Close(fd65)
if(ret != 0) {
  panic("close failed")
}

//fd state: (99) a_fUMcwextYd8iOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOn

ret = client.Write(fd155, []byte("lkdnk_IsnSZuVyQtxd2WHXKEKQYrqWC39NT8Uwnwj7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (141) a_fUMcwextYd8iOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOnlkdnk_IsnSZuVyQtxd2WHXKEKQYrqWC39NT8Uwnwj7
//fd state: (181) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719g

ret = client.Write(fd164, []byte("UPCTWLyXdV0xNBNYebaHHWzLOZ4Sig"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (211) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4Sig

ret = client.Close(fd139)
if(ret != 0) {
  panic("close failed")
}

//fd state: (6) 2DzlNJ

ret = client.Write(fd152, []byte("yjWmtP5jB1LbHUF09cihyz2rojouzaEcwGCeHBiYxS8Zj1jDsyOEPO1fY9KVLUhD7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (71) 2DzlNJyjWmtP5jB1LbHUF09cihyz2rojouzaEcwGCeHBiYxS8Zj1jDsyOEPO1fY9KVLUhD7
//fd state: (141) a_fUMcwextYd8iOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOnlkdnk_IsnSZuVyQtxd2WHXKEKQYrqWC39NT8Uwnwj7

ret = client.Write(fd155, []byte("4oruCYzPSUwb97TLcpjeMqGKhbeKJTzn6C1F86Ge20SZFRh11ZKC9WRUBDRvLtEUYH2dA0WDyF0vk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (218) a_fUMcwextYd8iOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOnlkdnk_IsnSZuVyQtxd2WHXKEKQYrqWC39NT8Uwnwj74oruCYzPSUwb97TLcpjeMqGKhbeKJTzn6C1F86Ge20SZFRh11ZKC9WRUBDRvLtEUYH2dA0WDyF0vk

ret = client.Seek(fd164, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


buf, ret = client.Read(fd123, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "b3E0YuxbCxFfvq") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd145, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd152, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}


fd173 := client.Open("/LsRInew_XM.txt", client.O_RDWR|client.O_CREATE)
if(fd173 < 0) {
  panic("open failed")
}


ret = client.Close(fd131)
if(ret != 0) {
  panic("close failed")
}

//fd state: (3) Zl_7dRcqenOlSiRgrFHeVLRIpCgot1lz5prcLXqXCgMZsFrz7T83PEFKaISdQDSmMkqnsRbe9k0hF9TJ1IimQsIZZqNK8YiQ8w7rpMGCpZctmmgrjBHs_Rs1zoi7czoiLxjt0xuFuDiFLcbt6S2BPUp3SaMKLzYLh_JeL97yBn3SnFT_FYacswxoXHqsyevbbm

ret = client.Write(fd163, []byte("Kn0bAf7gtnUaiFeV7Rr672ew1ePNfxiB7cv25krenR91tj2m_KHbebwBNieojuGdiJZEpkVijfECgnto2t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) Zl_Kn0bAf7gtnUaiFeV7Rr672ew1ePNfxiB7cv25krenR91tj2m_KHbebwBNieojuGdiJZEpkVijfECgnto2tsIZZqNK8YiQ8w7rpMGCpZctmmgrjBHs_Rs1zoi7czoiLxjt0xuFuDiFLcbt6S2BPUp3SaMKLzYLh_JeL97yBn3SnFT_FYacswxoXHqsyevbbm

ret = client.Close(fd163)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd123, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd123, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "dMy6kgRXPKoH9Gb_vAQSkCtE2GyDxTUU5QY") {
  panic("wrong data returned")
}


fd174 := client.Open("/LW1OckBfGg.txt", client.O_RDWR|client.O_CREATE)
if(fd174 < 0) {
  panic("open failed")
}


fd175 := client.Open("/Iw63byftWQ.txt", client.O_RDWR|client.O_CREATE)
if(fd175 < 0) {
  panic("open failed")
}

//fd state: (62) crMqUzA79wb7zbBk3DqTw4tMYHHLJylkJZgimUahsXalff_I0lCnMAS5GB059S

ret = client.Write(fd167, []byte("Jboo7m1NLdIVxkMmqx0CNEC2iJAA9phw3FlQ262g4zvSTiH_N9c2Iit3Oxt0YIDbxwJ9Cqyqlvbj5R2RnFHiuga_Cs0k11eKgk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (160) crMqUzA79wb7zbBk3DqTw4tMYHHLJylkJZgimUahsXalff_I0lCnMAS5GB059SJboo7m1NLdIVxkMmqx0CNEC2iJAA9phw3FlQ262g4zvSTiH_N9c2Iit3Oxt0YIDbxwJ9Cqyqlvbj5R2RnFHiuga_Cs0k11eKgk

ret = client.Seek(fd123, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd175, []byte("SYkNMe2LUfDd3TxKrU1_n2I2sW1FUypSBbqREwG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) SYkNMe2LUfDd3TxKrU1_n2I2sW1FUypSBbqREwG

ret = client.Close(fd157)
if(ret != 0) {
  panic("close failed")
}

//fd state: (39) SYkNMe2LUfDd3TxKrU1_n2I2sW1FUypSBbqREwG

ret = client.Write(fd175, []byte("8SybvUk_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) SYkNMe2LUfDd3TxKrU1_n2I2sW1FUypSBbqREwG8SybvUk_

fd176 := client.Open("/dfRXaFO9tW.txt", client.O_RDWR|client.O_CREATE)
if(fd176 < 0) {
  panic("open failed")
}


ret = client.Close(fd155)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd172, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd177 := client.Open("/FS0Irk9hhz.txt", client.O_RDWR|client.O_CREATE)
if(fd177 < 0) {
  panic("open failed")
}


fd178 := client.Open("/YbYhTNaGQb.txt", client.O_RDWR|client.O_CREATE)
if(fd178 < 0) {
  panic("open failed")
}


ret = client.Close(fd177)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd176, []byte("YWFH3qBKfvN7Rvm3mBul2AMRxbzAY6zOvQPu6ZO4NhjU1gHSaX50BDx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) YWFH3qBKfvN7Rvm3mBul2AMRxbzAY6zOvQPu6ZO4NhjU1gHSaX50BDx
//fd state: (47) SYkNMe2LUfDd3TxKrU1_n2I2sW1FUypSBbqREwG8SybvUk_

ret = client.Write(fd175, []byte("eJ5eIdPGTqJmob84p5UCq67FIzVRMCDv1Wl7pF8Kva7aAlKes2qmKOJa2_s8xNBo8wDEDM80gIHs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) SYkNMe2LUfDd3TxKrU1_n2I2sW1FUypSBbqREwG8SybvUk_eJ5eIdPGTqJmob84p5UCq67FIzVRMCDv1Wl7pF8Kva7aAlKes2qmKOJa2_s8xNBo8wDEDM80gIHs

fd179 := client.Open("/Xa2of_COvq.txt", client.O_RDWR|client.O_CREATE)
if(fd179 < 0) {
  panic("open failed")
}


ret = client.Close(fd164)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd171, []byte("W1jjQ9K4vCDt5v9PIx7Op2Bec"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) W1jjQ9K4vCDt5v9PIx7Op2Bec
//fd state: (123) SYkNMe2LUfDd3TxKrU1_n2I2sW1FUypSBbqREwG8SybvUk_eJ5eIdPGTqJmob84p5UCq67FIzVRMCDv1Wl7pF8Kva7aAlKes2qmKOJa2_s8xNBo8wDEDM80gIHs

ret = client.Write(fd175, []byte("AlXQngR1Ard3f7wjYvkwnu9o"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (147) SYkNMe2LUfDd3TxKrU1_n2I2sW1FUypSBbqREwG8SybvUk_eJ5eIdPGTqJmob84p5UCq67FIzVRMCDv1Wl7pF8Kva7aAlKes2qmKOJa2_s8xNBo8wDEDM80gIHsAlXQngR1Ard3f7wjYvkwnu9o
//fd state: (0) 

ret = client.Write(fd173, []byte("5HLC9P5JC5m6KeeKjYOXLT0cuwVSHK9m"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) 5HLC9P5JC5m6KeeKjYOXLT0cuwVSHK9m

ret = client.Seek(fd174, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (25) 9abPdTLhlhAUmBz5WkfxBAjGoRb3E0YuxbCxFfvqdMy6kgRXPKoH9Gb_vAQSkCtE2GyDxTUU5QYBIdOPU_VNPcDmZWK60qAibs7t_xD38Z8DhgUsrK4nDiAKtCWHyqn6

ret = client.Write(fd123, []byte("aN3jeJC7wkxL2m2RMeaqDyLe5SeWX0q3deX4RTv99zygkQ3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) 9abPdTLhlhAUmBz5WkfxBAjGoaN3jeJC7wkxL2m2RMeaqDyLe5SeWX0q3deX4RTv99zygkQ35QYBIdOPU_VNPcDmZWK60qAibs7t_xD38Z8DhgUsrK4nDiAKtCWHyqn6

ret = client.Seek(fd123, 92, client.SEEK_SET)
if(ret != 92) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 92)
  panic("seek failed")
}


buf, ret = client.Read(fd167, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (301) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiagvSJLxhAzvUrSop8hr__YmWyQK19kga_h7SKcfR2U5IwspgpYWDLVOg1OiCMb8J8InUXJRJEYXvFJH

ret = client.Write(fd145, []byte("avQvrvbDlTpidwvZIjD0brf3CFH4_LlOZjKvROBPWpn6x8DEVZMUUDH6XPMV8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (362) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiagvSJLxhAzvUrSop8hr__YmWyQK19kga_h7SKcfR2U5IwspgpYWDLVOg1OiCMb8J8InUXJRJEYXvFJHavQvrvbDlTpidwvZIjD0brf3CFH4_LlOZjKvROBPWpn6x8DEVZMUUDH6XPMV8
//fd state: (0) 

ret = client.Write(fd174, []byte("vwDBrUWp6YLKpvN5M_CGPbIA7DjIwaa6yDGnFJW0nVtu9PWX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) vwDBrUWp6YLKpvN5M_CGPbIA7DjIwaa6yDGnFJW0nVtu9PWX

fd180 := client.Open("/plv8EuK8Z3.txt", client.O_RDWR|client.O_CREATE)
if(fd180 < 0) {
  panic("open failed")
}


ret = client.Close(fd172)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd144)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 1LC7RoKzOuHe3xV3zJIt_lL45wMqR

ret = client.Write(fd178, []byte("Oq_Pzj8StdxDIMRztxchZTE0AK8psOT0VlziWgob6GMCLqdb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) Oq_Pzj8StdxDIMRztxchZTE0AK8psOT0VlziWgob6GMCLqdb

ret = client.Seek(fd171, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


buf, ret = client.Read(fd171, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Q9K4vCDt5v9PIx7Op2Bec") {
  panic("wrong data returned")
}

//fd state: (147) SYkNMe2LUfDd3TxKrU1_n2I2sW1FUypSBbqREwG8SybvUk_eJ5eIdPGTqJmob84p5UCq67FIzVRMCDv1Wl7pF8Kva7aAlKes2qmKOJa2_s8xNBo8wDEDM80gIHsAlXQngR1Ard3f7wjYvkwnu9o

ret = client.Write(fd175, []byte("8fv6kO7U5OKB1qkJXCYIwkhonB6e_LooFGeOSCu8iTubAQKSiI745_ZgbkHnb3Lq2HvLyL9H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (219) SYkNMe2LUfDd3TxKrU1_n2I2sW1FUypSBbqREwG8SybvUk_eJ5eIdPGTqJmob84p5UCq67FIzVRMCDv1Wl7pF8Kva7aAlKes2qmKOJa2_s8xNBo8wDEDM80gIHsAlXQngR1Ard3f7wjYvkwnu9o8fv6kO7U5OKB1qkJXCYIwkhonB6e_LooFGeOSCu8iTubAQKSiI745_ZgbkHnb3Lq2HvLyL9H

ret = client.Close(fd176)
if(ret != 0) {
  panic("close failed")
}


fd181 := client.Open("/JR6nH_boRB.txt", client.O_RDWR|client.O_CREATE)
if(fd181 < 0) {
  panic("open failed")
}


ret = client.Seek(fd178, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


ret = client.Seek(fd174, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}

//fd state: (92) 9abPdTLhlhAUmBz5WkfxBAjGoaN3jeJC7wkxL2m2RMeaqDyLe5SeWX0q3deX4RTv99zygkQ35QYBIdOPU_VNPcDmZWK60qAibs7t_xD38Z8DhgUsrK4nDiAKtCWHyqn6

ret = client.Write(fd123, []byte("lw4w791JXvwZRJZbV44SNFh_RRpU75o8bUzu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (128) 9abPdTLhlhAUmBz5WkfxBAjGoaN3jeJC7wkxL2m2RMeaqDyLe5SeWX0q3deX4RTv99zygkQ35QYBIdOPU_VNPcDmZWK6lw4w791JXvwZRJZbV44SNFh_RRpU75o8bUzu
//fd state: (160) crMqUzA79wb7zbBk3DqTw4tMYHHLJylkJZgimUahsXalff_I0lCnMAS5GB059SJboo7m1NLdIVxkMmqx0CNEC2iJAA9phw3FlQ262g4zvSTiH_N9c2Iit3Oxt0YIDbxwJ9Cqyqlvbj5R2RnFHiuga_Cs0k11eKgk

ret = client.Write(fd167, []byte("Uxwl_gKFG10gzm01LbPoBLSNKvWqVQGYFiiKAJVNwrFwsJUwWbf1S1xh82twW72yyvWYDZYog"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (233) crMqUzA79wb7zbBk3DqTw4tMYHHLJylkJZgimUahsXalff_I0lCnMAS5GB059SJboo7m1NLdIVxkMmqx0CNEC2iJAA9phw3FlQ262g4zvSTiH_N9c2Iit3Oxt0YIDbxwJ9Cqyqlvbj5R2RnFHiuga_Cs0k11eKgkUxwl_gKFG10gzm01LbPoBLSNKvWqVQGYFiiKAJVNwrFwsJUwWbf1S1xh82twW72yyvWYDZYog
//fd state: (0) 

ret = client.Write(fd169, []byte("E8kFPyWxSRzQdcRlIkUEEQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) E8kFPyWxSRzQdcRlIkUEEQ
//fd state: (32) 5HLC9P5JC5m6KeeKjYOXLT0cuwVSHK9m

ret = client.Write(fd173, []byte("9RZWiH_JbDT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) 5HLC9P5JC5m6KeeKjYOXLT0cuwVSHK9m9RZWiH_JbDT
//fd state: (43) 5HLC9P5JC5m6KeeKjYOXLT0cuwVSHK9m9RZWiH_JbDT

ret = client.Write(fd173, []byte("owo6mVpgu_ptMcqtpataA6EgZPfD9aDZLYTCW6y4B2qOZTGZx1YG5qGZFuEE4KcRSiW2mrT2e4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) 5HLC9P5JC5m6KeeKjYOXLT0cuwVSHK9m9RZWiH_JbDTowo6mVpgu_ptMcqtpataA6EgZPfD9aDZLYTCW6y4B2qOZTGZx1YG5qGZFuEE4KcRSiW2mrT2e4

ret = client.Close(fd152)
if(ret != 0) {
  panic("close failed")
}


fd182 := client.Open("/SIYaGzyyUN.txt", client.O_RDWR|client.O_CREATE)
if(fd182 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd180, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd168)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd162)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd167, 164, client.SEEK_SET)
if(ret != 164) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 164)
  panic("seek failed")
}


ret = client.Seek(fd180, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd180, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd174, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5M_CGPbIA7DjIwaa6yDGn") {
  panic("wrong data returned")
}


ret = client.Close(fd175)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) a_fUMcwextYd8iOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOnlkdnk_IsnSZuVyQtxd2WHXKEKQYrqWC39NT8Uwnwj74oruCYzPSUwb97TLcpjeMqGKhbeKJTzn6C1F86Ge20SZFRh11ZKC9WRUBDRvLtEUYH2dA0WDyF0vk

ret = client.Write(fd182, []byte("DfKeJmPN1B7Goo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) DfKeJmPN1B7GooOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOnlkdnk_IsnSZuVyQtxd2WHXKEKQYrqWC39NT8Uwnwj74oruCYzPSUwb97TLcpjeMqGKhbeKJTzn6C1F86Ge20SZFRh11ZKC9WRUBDRvLtEUYH2dA0WDyF0vk

ret = client.Seek(fd173, 88, client.SEEK_SET)
if(ret != 88) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 88)
  panic("seek failed")
}


ret = client.Seek(fd171, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Seek(fd181, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


buf, ret = client.Read(fd179, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd145, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}


ret = client.Seek(fd167, 146, client.SEEK_SET)
if(ret != 146) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 146)
  panic("seek failed")
}


fd183 := client.Open("/CWlKR1JLzO.txt", client.O_RDWR|client.O_CREATE)
if(fd183 < 0) {
  panic("open failed")
}


ret = client.Seek(fd179, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd181, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}

//fd state: (14) Oq_Pzj8StdxDIMRztxchZTE0AK8psOT0VlziWgob6GMCLqdb

ret = client.Write(fd178, []byte("7c06RbG_PpO_k60WoOoiO6nTKbPjlmHDfyfVbhTqhiVqFPxfNJsQaclvbczMpq1FtGXy8y1AoMEfF3rH7Pn2cp4ysIMe0XwCp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) Oq_Pzj8StdxDIM7c06RbG_PpO_k60WoOoiO6nTKbPjlmHDfyfVbhTqhiVqFPxfNJsQaclvbczMpq1FtGXy8y1AoMEfF3rH7Pn2cp4ysIMe0XwCp

ret = client.Seek(fd181, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


ret = client.Seek(fd178, 90, client.SEEK_SET)
if(ret != 90) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 90)
  panic("seek failed")
}


fd184 := client.Open("/CcuIbz5Tou.txt", client.O_RDWR|client.O_CREATE)
if(fd184 < 0) {
  panic("open failed")
}


ret = client.Close(fd171)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd181)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd167, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}


ret = client.Close(fd123)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd178, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Close(fd169)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd179)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd173, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TGZx1YG5qGZFu") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd145, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "A9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd167, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Q262g4zvSTiH_N9c2Iit3Oxt0YIDbxwJ9Cqyqlvbj5R2RnFHiuga_Cs") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd180, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd167, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "0k11eKgkUxwl_gKFG10gz") {
  panic("wrong data returned")
}


ret = client.Seek(fd173, 64, client.SEEK_SET)
if(ret != 64) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 64)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd183, []byte("79Gkuc5UD11chdFPzQa8RfCoceAy28Kz_UPxn4ZTDwdTAel1ZZ38zIgr7G1c"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) 79Gkuc5UD11chdFPzQa8RfCoceAy28Kz_UPxn4ZTDwdTAel1ZZ38zIgr7G1c

buf, ret = client.Read(fd174, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "FJW0nVtu9PWX") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd185 := client.Open("/seP9Yl_hKG.txt", client.O_RDWR|client.O_CREATE)
if(fd185 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd183, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (173) crMqUzA79wb7zbBk3DqTw4tMYHHLJylkJZgimUahsXalff_I0lCnMAS5GB059SJboo7m1NLdIVxkMmqx0CNEC2iJAA9phw3FlQ262g4zvSTiH_N9c2Iit3Oxt0YIDbxwJ9Cqyqlvbj5R2RnFHiuga_Cs0k11eKgkUxwl_gKFG10gzm01LbPoBLSNKvWqVQGYFiiKAJVNwrFwsJUwWbf1S1xh82twW72yyvWYDZYog

ret = client.Write(fd167, []byte("ikjv2wBzDw4_B4RNwAU1VkxTS1GCGkr0R"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (206) crMqUzA79wb7zbBk3DqTw4tMYHHLJylkJZgimUahsXalff_I0lCnMAS5GB059SJboo7m1NLdIVxkMmqx0CNEC2iJAA9phw3FlQ262g4zvSTiH_N9c2Iit3Oxt0YIDbxwJ9Cqyqlvbj5R2RnFHiuga_Cs0k11eKgkUxwl_gKFG10gzikjv2wBzDw4_B4RNwAU1VkxTS1GCGkr0RUwWbf1S1xh82twW72yyvWYDZYog
//fd state: (64) 5HLC9P5JC5m6KeeKjYOXLT0cuwVSHK9m9RZWiH_JbDTowo6mVpgu_ptMcqtpataA6EgZPfD9aDZLYTCW6y4B2qOZTGZx1YG5qGZFuEE4KcRSiW2mrT2e4

ret = client.Write(fd173, []byte("cS5PTFAoA2D3HLowYA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) 5HLC9P5JC5m6KeeKjYOXLT0cuwVSHK9m9RZWiH_JbDTowo6mVpgu_ptMcqtpataAcS5PTFAoA2D3HLowYA4B2qOZTGZx1YG5qGZFuEE4KcRSiW2mrT2e4

ret = client.Seek(fd185, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd182, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn") {
  panic("wrong data returned")
}


ret = client.Close(fd185)
if(ret != 0) {
  panic("close failed")
}


fd186 := client.Open("/cDsrk7sXrQ.txt", client.O_RDWR|client.O_CREATE)
if(fd186 < 0) {
  panic("open failed")
}


ret = client.Close(fd167)
if(ret != 0) {
  panic("close failed")
}


fd187 := client.Open("/Ac1BgFRgda.txt", client.O_RDWR|client.O_CREATE)
if(fd187 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd174, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd145, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DW") {
  panic("wrong data returned")
}

//fd state: (48) vwDBrUWp6YLKpvN5M_CGPbIA7DjIwaa6yDGnFJW0nVtu9PWX

ret = client.Write(fd174, []byte("uq3vroj3TCCcyj_G_1poA1eioGLYkIPnmucYOIYdvcLXTnw5FcWdGI0SLXo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) vwDBrUWp6YLKpvN5M_CGPbIA7DjIwaa6yDGnFJW0nVtu9PWXuq3vroj3TCCcyj_G_1poA1eioGLYkIPnmucYOIYdvcLXTnw5FcWdGI0SLXo

ret = client.Close(fd184)
if(ret != 0) {
  panic("close failed")
}


fd188 := client.Open("/DBm2gV4Brr.txt", client.O_RDWR|client.O_CREATE)
if(fd188 < 0) {
  panic("open failed")
}


fd189 := client.Open("/B_hZtR1Qiv.txt", client.O_RDWR|client.O_CREATE)
if(fd189 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd182, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_") {
  panic("wrong data returned")
}


ret = client.Close(fd180)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd188, []byte("otY1YddnAdlP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) otY1YddnAdlP

ret = client.Seek(fd183, 60, client.SEEK_SET)
if(ret != 60) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 60)
  panic("seek failed")
}

//fd state: (12) otY1YddnAdlP

ret = client.Write(fd188, []byte("LH5m24_P1XI0ceEAlaN_9tfRLMvxS2K6HscEZS0p8MFH6cD2EVulMo4sIXT71TmYfz2KNjw2A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) otY1YddnAdlPLH5m24_P1XI0ceEAlaN_9tfRLMvxS2K6HscEZS0p8MFH6cD2EVulMo4sIXT71TmYfz2KNjw2A
//fd state: (0) v715Dz34IWoGQVx63AjHueKVXVi

ret = client.Write(fd189, []byte("wcCG57d_axXbeMp3AHirRnn7D3W84PqYak8d4j3WsNzPu263Yla1RcqMT7D7d"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) wcCG57d_axXbeMp3AHirRnn7D3W84PqYak8d4j3WsNzPu263Yla1RcqMT7D7d

ret = client.Seek(fd188, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


fd190 := client.Open("/K1YhjaImgW.txt", client.O_RDWR|client.O_CREATE)
if(fd190 < 0) {
  panic("open failed")
}

//fd state: (17) Oq_Pzj8StdxDIM7c06RbG_PpO_k60WoOoiO6nTKbPjlmHDfyfVbhTqhiVqFPxfNJsQaclvbczMpq1FtGXy8y1AoMEfF3rH7Pn2cp4ysIMe0XwCp

ret = client.Write(fd178, []byte("HWFZl8ZpFHO03TRVEahm2r5s7wBfjV7FHFpM6miij1Dc6WC4jtcA3b9vUK5AlyPOqWRJzpzApmWGlMNUd4nZQhqAdftp_UeK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) Oq_Pzj8StdxDIM7c0HWFZl8ZpFHO03TRVEahm2r5s7wBfjV7FHFpM6miij1Dc6WC4jtcA3b9vUK5AlyPOqWRJzpzApmWGlMNUd4nZQhqAdftp_UeK

buf, ret = client.Read(fd188, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vxS2K6HscEZS0p8MFH6cD2EVu") {
  panic("wrong data returned")
}


ret = client.Seek(fd182, 173, client.SEEK_SET)
if(ret != 173) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 173)
  panic("seek failed")
}


buf, ret = client.Read(fd178, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd191 := client.Open("/OmH6ue7qfz.txt", client.O_RDWR|client.O_CREATE)
if(fd191 < 0) {
  panic("open failed")
}


ret = client.Seek(fd187, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd174)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd178, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd186)
if(ret != 0) {
  panic("close failed")
}


fd192 := client.Open("/8BJSeBJnmX.txt", client.O_RDWR|client.O_CREATE)
if(fd192 < 0) {
  panic("open failed")
}


ret = client.Seek(fd188, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


fd193 := client.Open("/DWICQRTyV5.txt", client.O_RDWR|client.O_CREATE)
if(fd193 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd178, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd190)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd193)
if(ret != 0) {
  panic("close failed")
}


fd194 := client.Open("/gg2gBLYHOn.txt", client.O_RDWR|client.O_CREATE)
if(fd194 < 0) {
  panic("open failed")
}


ret = client.Seek(fd178, 71, client.SEEK_SET)
if(ret != 71) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 71)
  panic("seek failed")
}


ret = client.Close(fd189)
if(ret != 0) {
  panic("close failed")
}


fd195 := client.Open("/mHY5KcNlPE.txt", client.O_RDWR|client.O_CREATE)
if(fd195 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd191, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "0NUSBWlxgvZCB_gsv7K2iMnSw1t") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd192, []byte("AihOwtba5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) AihOwtba5

ret = client.Seek(fd187, 109, client.SEEK_SET)
if(ret != 109) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 109)
  panic("seek failed")
}


buf, ret = client.Read(fd192, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (109) TDXJOfypwsevp66He4fi0MwYq0hIu7G9yDlMWcA87t8wF8AqXEokg9kbyKAkirA38eqEPzSxDRTueYzx1VwEx34pk5KCkxvfb48vPEWeE6M_3l

ret = client.Write(fd187, []byte("su88HRcW7DreO2BMnuI25kAoYowPC35cPkS010IyWHfg0HE7ILVm6s96GL6Y0AJxD36kiqzyP1WHfPnGAt61dCSmRNrmTR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (203) TDXJOfypwsevp66He4fi0MwYq0hIu7G9yDlMWcA87t8wF8AqXEokg9kbyKAkirA38eqEPzSxDRTueYzx1VwEx34pk5KCkxvfb48vPEWeE6M_3su88HRcW7DreO2BMnuI25kAoYowPC35cPkS010IyWHfg0HE7ILVm6s96GL6Y0AJxD36kiqzyP1WHfPnGAt61dCSmRNrmTR

ret = client.Seek(fd194, 136, client.SEEK_SET)
if(ret != 136) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 136)
  panic("seek failed")
}


buf, ret = client.Read(fd173, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4B2qOZTGZx1YG5qGZFuEE4KcRSiW2mrT2e4") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd191, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "kb05OxzISIxbdt") {
  panic("wrong data returned")
}


ret = client.Close(fd178)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd195, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd196 := client.Open("/jAMfPtHSb4.txt", client.O_RDWR|client.O_CREATE)
if(fd196 < 0) {
  panic("open failed")
}


ret = client.Seek(fd173, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}

//fd state: (173) DfKeJmPN1B7GooOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOnlkdnk_IsnSZuVyQtxd2WHXKEKQYrqWC39NT8Uwnwj74oruCYzPSUwb97TLcpjeMqGKhbeKJTzn6C1F86Ge20SZFRh11ZKC9WRUBDRvLtEUYH2dA0WDyF0vk

ret = client.Write(fd182, []byte("AqbEthzIK4nRA9vPNgQUsRL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (196) DfKeJmPN1B7GooOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOnlkdnk_IsnSZuVyQtxd2WHXKEKQYrqWC39NT8Uwnwj74oruCYzPSUwb97TLcpjeMqGKhbeKJTznAqbEthzIK4nRA9vPNgQUsRLUBDRvLtEUYH2dA0WDyF0vk
//fd state: (60) 79Gkuc5UD11chdFPzQa8RfCoceAy28Kz_UPxn4ZTDwdTAel1ZZ38zIgr7G1c

ret = client.Write(fd183, []byte("VKsVPRHRt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) 79Gkuc5UD11chdFPzQa8RfCoceAy28Kz_UPxn4ZTDwdTAel1ZZ38zIgr7G1cVKsVPRHRt
//fd state: (136) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1gzmbGANpt2sz3fGu6gekHaqKnAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7oVcqeEjIstIWzYfro2Gyix_jlGr3iCDabBrdCaj2kgsR8TTWtAlS3Ph7Pi

ret = client.Write(fd194, []byte("fQkcLIhMHAAktpO24zOxhME8LR3UZvI0CsALPSnLhybAybMQ_9JDhVPdCu3w439J1df3X7AsQgaahAdKjgCSbK_jTZtOdAQZd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (233) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1gzmbGANpt2sz3fGu6gekHaqKnAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7oVcqeEjIstIfQkcLIhMHAAktpO24zOxhME8LR3UZvI0CsALPSnLhybAybMQ_9JDhVPdCu3w439J1df3X7AsQgaahAdKjgCSbK_jTZtOdAQZd

ret = client.Seek(fd145, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Seek(fd188, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Seek(fd145, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


ret = client.Close(fd196)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd194, 120, client.SEEK_SET)
if(ret != 120) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 120)
  panic("seek failed")
}


ret = client.Close(fd191)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd194, 225, client.SEEK_SET)
if(ret != 225) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 225)
  panic("seek failed")
}


buf, ret = client.Read(fd183, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (9) AihOwtba5

ret = client.Write(fd192, []byte("H3qSpTCMtJRpJgXz1G3KL2WWGWcjjALNthedw8a5Hz4H9yRUnrL_eW4V8LpWbcft0WAkQSpbnp2z3ziR_58qOyGd8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) AihOwtba5H3qSpTCMtJRpJgXz1G3KL2WWGWcjjALNthedw8a5Hz4H9yRUnrL_eW4V8LpWbcft0WAkQSpbnp2z3ziR_58qOyGd8

buf, ret = client.Read(fd182, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "UBDRvLtEUYH2dA0WDyF0vk") {
  panic("wrong data returned")
}


ret = client.Close(fd194)
if(ret != 0) {
  panic("close failed")
}


fd197 := client.Open("/t2Dyv3YWoA.txt", client.O_RDWR|client.O_CREATE)
if(fd197 < 0) {
  panic("open failed")
}


fd198 := client.Open("/UHzput6Vyv.txt", client.O_RDWR|client.O_CREATE)
if(fd198 < 0) {
  panic("open failed")
}


ret = client.Seek(fd183, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


buf, ret = client.Read(fd183, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "DwdTAel1ZZ38zIgr7G1cVKsVPRHRt") {
  panic("wrong data returned")
}

//fd state: (81) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9SiY2P8OF0chwpSsmyqL22bswMBaH0h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiagvSJLxhAzvUrSop8hr__YmWyQK19kga_h7SKcfR2U5IwspgpYWDLVOg1OiCMb8J8InUXJRJEYXvFJHavQvrvbDlTpidwvZIjD0brf3CFH4_LlOZjKvROBPWpn6x8DEVZMUUDH6XPMV8

ret = client.Write(fd145, []byte("xg6wnoxE6bRo3UWAjuujvTFChHe6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (109) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9Sxg6wnoxE6bRo3UWAjuujvTFChHe6h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiagvSJLxhAzvUrSop8hr__YmWyQK19kga_h7SKcfR2U5IwspgpYWDLVOg1OiCMb8J8InUXJRJEYXvFJHavQvrvbDlTpidwvZIjD0brf3CFH4_LlOZjKvROBPWpn6x8DEVZMUUDH6XPMV8

fd199 := client.Open("/eXTFVE24CN.txt", client.O_RDWR|client.O_CREATE)
if(fd199 < 0) {
  panic("open failed")
}


ret = client.Close(fd173)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd182)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd145, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "h44ng3FeeQKYZz") {
  panic("wrong data returned")
}


ret = client.Seek(fd199, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd188, 77, client.SEEK_SET)
if(ret != 77) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 77)
  panic("seek failed")
}

//fd state: (77) otY1YddnAdlPLH5m24_P1XI0ceEAlaN_9tfRLMvxS2K6HscEZS0p8MFH6cD2EVulMo4sIXT71TmYfz2KNjw2A

ret = client.Write(fd188, []byte("eeUkEdSQLS_ZoSHWXfZY7Q7ceDobBBykh_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) otY1YddnAdlPLH5m24_P1XI0ceEAlaN_9tfRLMvxS2K6HscEZS0p8MFH6cD2EVulMo4sIXT71TmYfeeUkEdSQLS_ZoSHWXfZY7Q7ceDobBBykh_

ret = client.Seek(fd192, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}

//fd state: (37) AihOwtba5H3qSpTCMtJRpJgXz1G3KL2WWGWcjjALNthedw8a5Hz4H9yRUnrL_eW4V8LpWbcft0WAkQSpbnp2z3ziR_58qOyGd8

ret = client.Write(fd192, []byte("vAStfTd_P2JEebvrqPZvVyIF7fGtg9oZmKrmm2NmNoDo95UKJmTc_9MqxcgJJnUfdcPnE5Q4ZhSiT85vsQiJ3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) AihOwtba5H3qSpTCMtJRpJgXz1G3KL2WWGWcjvAStfTd_P2JEebvrqPZvVyIF7fGtg9oZmKrmm2NmNoDo95UKJmTc_9MqxcgJJnUfdcPnE5Q4ZhSiT85vsQiJ3

buf, ret = client.Read(fd187, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd200 := client.Open("/CdZyI4bMVY.txt", client.O_RDWR|client.O_CREATE)
if(fd200 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd188, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd192)
if(ret != 0) {
  panic("close failed")
}


fd201 := client.Open("/QIPEg8dwnM.txt", client.O_RDWR|client.O_CREATE)
if(fd201 < 0) {
  panic("open failed")
}


fd202 := client.Open("/iQtDHYElyv.txt", client.O_RDWR|client.O_CREATE)
if(fd202 < 0) {
  panic("open failed")
}


ret = client.Close(fd183)
if(ret != 0) {
  panic("close failed")
}


fd203 := client.Open("/K1YhjaImgW.txt", client.O_RDWR|client.O_CREATE)
if(fd203 < 0) {
  panic("open failed")
}


ret = client.Close(fd188)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd197)
if(ret != 0) {
  panic("close failed")
}


fd204 := client.Open("/i7rZ5jaEkR.txt", client.O_RDWR|client.O_CREATE)
if(fd204 < 0) {
  panic("open failed")
}


ret = client.Seek(fd204, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd205 := client.Open("/s4SMoZEfz7.txt", client.O_RDWR|client.O_CREATE)
if(fd205 < 0) {
  panic("open failed")
}


ret = client.Close(fd199)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd205, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd202, []byte("q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) q
//fd state: (0) E8kFPyWxSRzQdcRlIkUEEQ

ret = client.Write(fd203, []byte("hXnSISfeUo4y_s3jbsXgtxo5NI5ulgVyxOjLbpyBDBfuGWSUqFG3O8XL_WgsMDMO01N50XWBv8arXrUgNWNTCRNAiMTwpb0wtG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) hXnSISfeUo4y_s3jbsXgtxo5NI5ulgVyxOjLbpyBDBfuGWSUqFG3O8XL_WgsMDMO01N50XWBv8arXrUgNWNTCRNAiMTwpb0wtG
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd198, 100, client.SEEK_SET)
if(ret != 100) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 100)
  panic("seek failed")
}


ret = client.Seek(fd202, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}

//fd state: (98) hXnSISfeUo4y_s3jbsXgtxo5NI5ulgVyxOjLbpyBDBfuGWSUqFG3O8XL_WgsMDMO01N50XWBv8arXrUgNWNTCRNAiMTwpb0wtG

ret = client.Write(fd203, []byte("6ULRt2wBdLHDLq2HuQGFRcaVJLUiID6i0zDchzG2vupiyxrfnasqeC2dqSRo6ETem4FYSxZZTqRoPfRv7anR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (182) hXnSISfeUo4y_s3jbsXgtxo5NI5ulgVyxOjLbpyBDBfuGWSUqFG3O8XL_WgsMDMO01N50XWBv8arXrUgNWNTCRNAiMTwpb0wtG6ULRt2wBdLHDLq2HuQGFRcaVJLUiID6i0zDchzG2vupiyxrfnasqeC2dqSRo6ETem4FYSxZZTqRoPfRv7anR

fd206 := client.Open("/MLw9D0Mx68.txt", client.O_RDWR|client.O_CREATE)
if(fd206 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd200, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ylwLPUz4IfCymaNGuSE2") {
  panic("wrong data returned")
}


fd207 := client.Open("/AHpFNlzRlq.txt", client.O_RDWR|client.O_CREATE)
if(fd207 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd207, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd205, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd208 := client.Open("/zSQ7ASIRaX.txt", client.O_RDWR|client.O_CREATE)
if(fd208 < 0) {
  panic("open failed")
}


fd209 := client.Open("/PIYzUD6TOm.txt", client.O_RDWR|client.O_CREATE)
if(fd209 < 0) {
  panic("open failed")
}


ret = client.Seek(fd145, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


buf, ret = client.Read(fd209, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "HDA2GSp9eJQyurMHUHoArw6jmf5wXgq5MGitfvSpmgmxyo8g_QIZjp") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd201, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd200, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Hgp8JocFIzNlkBBv6gSuclvLNFQaK7NBZDrCw8omFf4eVqSVo38mx7jM7604vtuID1gqX") {
  panic("wrong data returned")
}


ret = client.Close(fd198)
if(ret != 0) {
  panic("close failed")
}


fd210 := client.Open("/_gbGbZZtm6.txt", client.O_RDWR|client.O_CREATE)
if(fd210 < 0) {
  panic("open failed")
}


ret = client.Close(fd200)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd202)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd203, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd211 := client.Open("/Avi_dS2O4_.txt", client.O_RDWR|client.O_CREATE)
if(fd211 < 0) {
  panic("open failed")
}


ret = client.Seek(fd203, 150, client.SEEK_SET)
if(ret != 150) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 150)
  panic("seek failed")
}


ret = client.Seek(fd195, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd201)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd206, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd206)
if(ret != 0) {
  panic("close failed")
}


fd212 := client.Open("/9hxNYjJQOk.txt", client.O_RDWR|client.O_CREATE)
if(fd212 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd145, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9Sxg6wnoxE6bRo3UWAjuujvTFChHe6h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2v") {
  panic("wrong data returned")
}


ret = client.Close(fd187)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd205)
if(ret != 0) {
  panic("close failed")
}


fd213 := client.Open("/SgtN9yKcDb.txt", client.O_RDWR|client.O_CREATE)
if(fd213 < 0) {
  panic("open failed")
}


fd214 := client.Open("/VghMsYjnXc.txt", client.O_RDWR|client.O_CREATE)
if(fd214 < 0) {
  panic("open failed")
}

//fd state: (0) kKWGrBhx5ONQXQ0hfTmF5xjIZ8lYtWeTNRwgKqXvE_iG88UEORJu7HL3Wccw8STn3MR

ret = client.Write(fd214, []byte("rKMS0KhTQDP6431qc4jEVLgWGE5cbYYdCIH0FM4EChefYSKxzJkgfTye_ucH_UsVl_efbPsEbY4JqoOET6mFW2X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) rKMS0KhTQDP6431qc4jEVLgWGE5cbYYdCIH0FM4EChefYSKxzJkgfTye_ucH_UsVl_efbPsEbY4JqoOET6mFW2X

ret = client.Seek(fd209, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


fd215 := client.Open("/YYMVsNBZSh.txt", client.O_RDWR|client.O_CREATE)
if(fd215 < 0) {
  panic("open failed")
}

//fd state: (142) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9Sxg6wnoxE6bRo3UWAjuujvTFChHe6h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2voJXmhJ4PAcK_eKCToK4pP7uzYMHw5ViPwegifzke0kE2iVs4DWxKIEYxZT0Z65yQ3rv0CoRL4B9GhKyiagvSJLxhAzvUrSop8hr__YmWyQK19kga_h7SKcfR2U5IwspgpYWDLVOg1OiCMb8J8InUXJRJEYXvFJHavQvrvbDlTpidwvZIjD0brf3CFH4_LlOZjKvROBPWpn6x8DEVZMUUDH6XPMV8

ret = client.Write(fd145, []byte("wBUR5242gnLu5pi67NI6gBdPtwh0sm_iNXaGCjKfUvqx_6ZHMcwn0GgQM9d"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (201) vb73i6FuLJQIxbEP9IMIxHokCgYv2uVZQuMwRoAJSw7lErR1Js2M_0yc_VPA9Pgj8IieJIdRO5IzM3B9Sxg6wnoxE6bRo3UWAjuujvTFChHe6h44ng3FeeQKYZzXGb0mHyXXDeXjLLoR2vwBUR5242gnLu5pi67NI6gBdPtwh0sm_iNXaGCjKfUvqx_6ZHMcwn0GgQM9dZ65yQ3rv0CoRL4B9GhKyiagvSJLxhAzvUrSop8hr__YmWyQK19kga_h7SKcfR2U5IwspgpYWDLVOg1OiCMb8J8InUXJRJEYXvFJHavQvrvbDlTpidwvZIjD0brf3CFH4_LlOZjKvROBPWpn6x8DEVZMUUDH6XPMV8

buf, ret = client.Read(fd212, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) _lIMrdvTX1NCHpH3DLLpDAshKP3PmhXw41Wc_ZmUjjvi3qo4JUNwmU7n2gLkZkkMtVgkXKvpP9

ret = client.Write(fd215, []byte("xfXyPorcGAnwUVemL26sl_Ig2uapFifjUvLt8TfXl5O2tiEidIPXOMmODS0vmvBFzZn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) xfXyPorcGAnwUVemL26sl_Ig2uapFifjUvLt8TfXl5O2tiEidIPXOMmODS0vmvBFzZnkXKvpP9
//fd state: (0) 

ret = client.Write(fd208, []byte("27bIlYzGmi8Inz4zCtLrrRcCnzrhkobHgkV2GQ9covn09cqVUHXT_XigEkwLiz_7vaiHZcrhKO1Zj3BC88mGL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) 27bIlYzGmi8Inz4zCtLrrRcCnzrhkobHgkV2GQ9covn09cqVUHXT_XigEkwLiz_7vaiHZcrhKO1Zj3BC88mGL

ret = client.Close(fd145)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd209, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}


ret = client.Close(fd211)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd203, 65, client.SEEK_SET)
if(ret != 65) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 65)
  panic("seek failed")
}


ret = client.Close(fd208)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd195, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd216 := client.Open("/8mk1VIqezX.txt", client.O_RDWR|client.O_CREATE)
if(fd216 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd216, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd209, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


buf, ret = client.Read(fd216, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (65) hXnSISfeUo4y_s3jbsXgtxo5NI5ulgVyxOjLbpyBDBfuGWSUqFG3O8XL_WgsMDMO01N50XWBv8arXrUgNWNTCRNAiMTwpb0wtG6ULRt2wBdLHDLq2HuQGFRcaVJLUiID6i0zDchzG2vupiyxrfnasqeC2dqSRo6ETem4FYSxZZTqRoPfRv7anR

ret = client.Write(fd203, []byte("rJmHPWTn97D6rdXBw82WXuyevEUqP9O3wdEx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) hXnSISfeUo4y_s3jbsXgtxo5NI5ulgVyxOjLbpyBDBfuGWSUqFG3O8XL_WgsMDMO0rJmHPWTn97D6rdXBw82WXuyevEUqP9O3wdExRt2wBdLHDLq2HuQGFRcaVJLUiID6i0zDchzG2vupiyxrfnasqeC2dqSRo6ETem4FYSxZZTqRoPfRv7anR
//fd state: (0) 

ret = client.Write(fd195, []byte("eDqhrw3RKPvL4x"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) eDqhrw3RKPvL4x

ret = client.Close(fd213)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd209)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd216, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd215, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}

//fd state: (87) rKMS0KhTQDP6431qc4jEVLgWGE5cbYYdCIH0FM4EChefYSKxzJkgfTye_ucH_UsVl_efbPsEbY4JqoOET6mFW2X

ret = client.Write(fd214, []byte("9H5qJTo1Uv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) rKMS0KhTQDP6431qc4jEVLgWGE5cbYYdCIH0FM4EChefYSKxzJkgfTye_ucH_UsVl_efbPsEbY4JqoOET6mFW2X9H5qJTo1Uv

buf, ret = client.Read(fd216, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd216, []byte("hkpyyuu5qZsHWxRlcqKakp3WZ1XRZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) hkpyyuu5qZsHWxRlcqKakp3WZ1XRZ

ret = client.Close(fd210)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd214, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd204, []byte("RpewFtsQ_UKw4Gp82K9rAGTrcIEUUazShcCqtGQsgVOP8GoOe6Y0QdlkL6OorxJMfZuU_9QrEmFyjQ4RJxwV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) RpewFtsQ_UKw4Gp82K9rAGTrcIEUUazShcCqtGQsgVOP8GoOe6Y0QdlkL6OorxJMfZuU_9QrEmFyjQ4RJxwV

ret = client.Seek(fd212, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd217 := client.Open("/dg_fg5cemT.txt", client.O_RDWR|client.O_CREATE)
if(fd217 < 0) {
  panic("open failed")
}


fd218 := client.Open("/zf1Y4vkWDq.txt", client.O_RDWR|client.O_CREATE)
if(fd218 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd218, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd212, []byte("eWwIPkXy_wlKVezkGPj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) eWwIPkXy_wlKVezkGPj

ret = client.Seek(fd203, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}

//fd state: (14) eDqhrw3RKPvL4x

ret = client.Write(fd195, []byte("NeTEXLO867sF6k0gWDmy8tL4StOg71tQYWUoEqrqELgmqqqI05_aOUP0UGcywSrwhEkYjGnz6Sh4ujCNPbYnPx7JT1b2TneK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) eDqhrw3RKPvL4xNeTEXLO867sF6k0gWDmy8tL4StOg71tQYWUoEqrqELgmqqqI05_aOUP0UGcywSrwhEkYjGnz6Sh4ujCNPbYnPx7JT1b2TneK
//fd state: (84) RpewFtsQ_UKw4Gp82K9rAGTrcIEUUazShcCqtGQsgVOP8GoOe6Y0QdlkL6OorxJMfZuU_9QrEmFyjQ4RJxwV

ret = client.Write(fd204, []byte("OsekYvMD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) RpewFtsQ_UKw4Gp82K9rAGTrcIEUUazShcCqtGQsgVOP8GoOe6Y0QdlkL6OorxJMfZuU_9QrEmFyjQ4RJxwVOsekYvMD

fd219 := client.Open("/di6atYCfA6.txt", client.O_RDWR|client.O_CREATE)
if(fd219 < 0) {
  panic("open failed")
}


ret = client.Seek(fd203, 162, client.SEEK_SET)
if(ret != 162) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 162)
  panic("seek failed")
}


ret = client.Close(fd204)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd207, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd214, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


fd220 := client.Open("/B_hZtR1Qiv.txt", client.O_RDWR|client.O_CREATE)
if(fd220 < 0) {
  panic("open failed")
}

//fd state: (110) eDqhrw3RKPvL4xNeTEXLO867sF6k0gWDmy8tL4StOg71tQYWUoEqrqELgmqqqI05_aOUP0UGcywSrwhEkYjGnz6Sh4ujCNPbYnPx7JT1b2TneK

ret = client.Write(fd195, []byte("sV3d19vxWeuDE_rlgQsDAxA7J_jHkDVpcy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) eDqhrw3RKPvL4xNeTEXLO867sF6k0gWDmy8tL4StOg71tQYWUoEqrqELgmqqqI05_aOUP0UGcywSrwhEkYjGnz6Sh4ujCNPbYnPx7JT1b2TneKsV3d19vxWeuDE_rlgQsDAxA7J_jHkDVpcy

buf, ret = client.Read(fd203, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "m4FYSxZZTqRoPfRv7anR") {
  panic("wrong data returned")
}


fd221 := client.Open("/Wod_MWssCh.txt", client.O_RDWR|client.O_CREATE)
if(fd221 < 0) {
  panic("open failed")
}


fd222 := client.Open("/gMb5eMbeqb.txt", client.O_RDWR|client.O_CREATE)
if(fd222 < 0) {
  panic("open failed")
}


ret = client.Seek(fd207, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd220, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wcCG57d_axXbeMp3AHirRnn7D3W84PqYak8d4j3WsNzPu263Yla1RcqMT7D7d") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd221, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd207)
if(ret != 0) {
  panic("close failed")
}

//fd state: (19) rKMS0KhTQDP6431qc4jEVLgWGE5cbYYdCIH0FM4EChefYSKxzJkgfTye_ucH_UsVl_efbPsEbY4JqoOET6mFW2X9H5qJTo1Uv

ret = client.Write(fd214, []byte("7I2xATpAGC1z79E4pY0m17gmHcIdisWUnuRuzvoDeHTCC4jsCkKv6h8GxCnZtYi0ouzdiRutHmzBLu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) rKMS0KhTQDP6431qc4j7I2xATpAGC1z79E4pY0m17gmHcIdisWUnuRuzvoDeHTCC4jsCkKv6h8GxCnZtYi0ouzdiRutHmzBLu

fd223 := client.Open("/l0WXb6xP4b.txt", client.O_RDWR|client.O_CREATE)
if(fd223 < 0) {
  panic("open failed")
}


fd224 := client.Open("/gg2gBLYHOn.txt", client.O_RDWR|client.O_CREATE)
if(fd224 < 0) {
  panic("open failed")
}


ret = client.Close(fd215)
if(ret != 0) {
  panic("close failed")
}


fd225 := client.Open("/jwmw0Yfr5L.txt", client.O_RDWR|client.O_CREATE)
if(fd225 < 0) {
  panic("open failed")
}


fd226 := client.Open("/ywoKRDO28J.txt", client.O_RDWR|client.O_CREATE)
if(fd226 < 0) {
  panic("open failed")
}


ret = client.Close(fd214)
if(ret != 0) {
  panic("close failed")
}


fd227 := client.Open("/epwKfrBhnk.txt", client.O_RDWR|client.O_CREATE)
if(fd227 < 0) {
  panic("open failed")
}


ret = client.Close(fd227)
if(ret != 0) {
  panic("close failed")
}

//fd state: (144) eDqhrw3RKPvL4xNeTEXLO867sF6k0gWDmy8tL4StOg71tQYWUoEqrqELgmqqqI05_aOUP0UGcywSrwhEkYjGnz6Sh4ujCNPbYnPx7JT1b2TneKsV3d19vxWeuDE_rlgQsDAxA7J_jHkDVpcy

ret = client.Write(fd195, []byte("s7vTSTgpbFw6sbmgFx_MHmSS7FaJ8wK29DoYOaU8ICH9OzuBIEUJ35jI0RgPhJG5HCF8SZ4P7NO3VzgEgZvbxUIdL01q9As"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (239) eDqhrw3RKPvL4xNeTEXLO867sF6k0gWDmy8tL4StOg71tQYWUoEqrqELgmqqqI05_aOUP0UGcywSrwhEkYjGnz6Sh4ujCNPbYnPx7JT1b2TneKsV3d19vxWeuDE_rlgQsDAxA7J_jHkDVpcys7vTSTgpbFw6sbmgFx_MHmSS7FaJ8wK29DoYOaU8ICH9OzuBIEUJ35jI0RgPhJG5HCF8SZ4P7NO3VzgEgZvbxUIdL01q9As

buf, ret = client.Read(fd195, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd218, []byte("2e4dUch"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) 2e4dUch

buf, ret = client.Read(fd220, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd203, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd221, []byte("7ZkJmDDHSpIJCat0ARtlwE5dI1XAQLrYHOG3xKS0vD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) 7ZkJmDDHSpIJCat0ARtlwE5dI1XAQLrYHOG3xKS0vD

buf, ret = client.Read(fd220, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd222)
if(ret != 0) {
  panic("close failed")
}


fd228 := client.Open("/nHgSrOUu58.txt", client.O_RDWR|client.O_CREATE)
if(fd228 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd219, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd225, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd228, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) OzbMSpH7AdFBKchOOFJVtSyTDioQhNo57ynR4dgUErha1T_MHD1gzmbGANpt2sz3fGu6gekHaqKnAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7oVcqeEjIstIfQkcLIhMHAAktpO24zOxhME8LR3UZvI0CsALPSnLhybAybMQ_9JDhVPdCu3w439J1df3X7AsQgaahAdKjgCSbK_jTZtOdAQZd

ret = client.Write(fd224, []byte("Lj1Koq_DR8J87Xi6Hyu9izfxP71Do5GOCMsYGaJejdMbCma72gItlhEH0xPdFT56MTa_7zVcY4mw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) Lj1Koq_DR8J87Xi6Hyu9izfxP71Do5GOCMsYGaJejdMbCma72gItlhEH0xPdFT56MTa_7zVcY4mwAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7oVcqeEjIstIfQkcLIhMHAAktpO24zOxhME8LR3UZvI0CsALPSnLhybAybMQ_9JDhVPdCu3w439J1df3X7AsQgaahAdKjgCSbK_jTZtOdAQZd

buf, ret = client.Read(fd228, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (76) Lj1Koq_DR8J87Xi6Hyu9izfxP71Do5GOCMsYGaJejdMbCma72gItlhEH0xPdFT56MTa_7zVcY4mwAEGzryPL1SmMAaf_2Oih_bRAY6WuzqcB9wZNehYi8y5mx_Yf7oVcqeEjIstIfQkcLIhMHAAktpO24zOxhME8LR3UZvI0CsALPSnLhybAybMQ_9JDhVPdCu3w439J1df3X7AsQgaahAdKjgCSbK_jTZtOdAQZd

ret = client.Write(fd224, []byte("LpQhz2r7cYh16mgS89qz9QdPhsP3GviR87pXneEhb11cF1t8FomuKRytKCOrepgkld"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (142) Lj1Koq_DR8J87Xi6Hyu9izfxP71Do5GOCMsYGaJejdMbCma72gItlhEH0xPdFT56MTa_7zVcY4mwLpQhz2r7cYh16mgS89qz9QdPhsP3GviR87pXneEhb11cF1t8FomuKRytKCOrepgkldhMHAAktpO24zOxhME8LR3UZvI0CsALPSnLhybAybMQ_9JDhVPdCu3w439J1df3X7AsQgaahAdKjgCSbK_jTZtOdAQZd

buf, ret = client.Read(fd219, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd216, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


ret = client.Close(fd226)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd224)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd228)
if(ret != 0) {
  panic("close failed")
}


fd229 := client.Open("/90hEKufnYM.txt", client.O_RDWR|client.O_CREATE)
if(fd229 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd229, []byte("mpJLASTrVJSdahxAZHmQKouYtt15M8K2H6vLugFHj9hNCVH4XOnR1sU8RbOeMscfhz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) mpJLASTrVJSdahxAZHmQKouYtt15M8K2H6vLugFHj9hNCVH4XOnR1sU8RbOeMscfhz
//fd state: (0) 

ret = client.Write(fd223, []byte("5e3qoV6cxD0Mnph6_YlKu64tTF7zkeacda1Gy32HrJA_QtidLYzPEibJAUNM8BMm3kOBL0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) 5e3qoV6cxD0Mnph6_YlKu64tTF7zkeacda1Gy32HrJA_QtidLYzPEibJAUNM8BMm3kOBL0

fd230 := client.Open("/FSTYY4xE5x.txt", client.O_RDWR|client.O_CREATE)
if(fd230 < 0) {
  panic("open failed")
}


fd231 := client.Open("/LK829EmRZQ.txt", client.O_RDWR|client.O_CREATE)
if(fd231 < 0) {
  panic("open failed")
}

//fd state: (239) eDqhrw3RKPvL4xNeTEXLO867sF6k0gWDmy8tL4StOg71tQYWUoEqrqELgmqqqI05_aOUP0UGcywSrwhEkYjGnz6Sh4ujCNPbYnPx7JT1b2TneKsV3d19vxWeuDE_rlgQsDAxA7J_jHkDVpcys7vTSTgpbFw6sbmgFx_MHmSS7FaJ8wK29DoYOaU8ICH9OzuBIEUJ35jI0RgPhJG5HCF8SZ4P7NO3VzgEgZvbxUIdL01q9As

ret = client.Write(fd195, []byte("SBPySJai72K6VJXs5qCbtdHMLToy4GjfVQHFXJj8WYujZtvcml2uF0T"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (294) eDqhrw3RKPvL4xNeTEXLO867sF6k0gWDmy8tL4StOg71tQYWUoEqrqELgmqqqI05_aOUP0UGcywSrwhEkYjGnz6Sh4ujCNPbYnPx7JT1b2TneKsV3d19vxWeuDE_rlgQsDAxA7J_jHkDVpcys7vTSTgpbFw6sbmgFx_MHmSS7FaJ8wK29DoYOaU8ICH9OzuBIEUJ35jI0RgPhJG5HCF8SZ4P7NO3VzgEgZvbxUIdL01q9AsSBPySJai72K6VJXs5qCbtdHMLToy4GjfVQHFXJj8WYujZtvcml2uF0T

ret = client.Seek(fd212, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


fd232 := client.Open("/2Kwg_3a7Gr.txt", client.O_RDWR|client.O_CREATE)
if(fd232 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd229, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd233 := client.Open("/JR6nH_boRB.txt", client.O_RDWR|client.O_CREATE)
if(fd233 < 0) {
  panic("open failed")
}


ret = client.Seek(fd218, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


buf, ret = client.Read(fd223, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) X9foy10Xnd93eeA5r91QWwL5OjDdwC8e_mqMdepWmfIw5ko5QmxeG_cqjjItKsAy4bQK7bLTyhqFXMNQROb8HpaJB32ezezvriPdEIgCuOMMrahuOso5MvyFu

ret = client.Write(fd233, []byte("EqoRamUWoBsrVug9PVsyL9BKb4rBnp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) EqoRamUWoBsrVug9PVsyL9BKb4rBnp8e_mqMdepWmfIw5ko5QmxeG_cqjjItKsAy4bQK7bLTyhqFXMNQROb8HpaJB32ezezvriPdEIgCuOMMrahuOso5MvyFu

ret = client.Close(fd217)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd230, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


ret = client.Close(fd232)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd203)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd219)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd212)
if(ret != 0) {
  panic("close failed")
}


fd234 := client.Open("/kiQZc5jb4w.txt", client.O_RDWR|client.O_CREATE)
if(fd234 < 0) {
  panic("open failed")
}

//fd state: (66) mpJLASTrVJSdahxAZHmQKouYtt15M8K2H6vLugFHj9hNCVH4XOnR1sU8RbOeMscfhz

ret = client.Write(fd229, []byte("dfD3nKpuFPWEMyIxk_vPLqNfLrVyCVpyyd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) mpJLASTrVJSdahxAZHmQKouYtt15M8K2H6vLugFHj9hNCVH4XOnR1sU8RbOeMscfhzdfD3nKpuFPWEMyIxk_vPLqNfLrVyCVpyyd
//fd state: (70) 5e3qoV6cxD0Mnph6_YlKu64tTF7zkeacda1Gy32HrJA_QtidLYzPEibJAUNM8BMm3kOBL0

ret = client.Write(fd223, []byte("m5kIG8EOpIMAh4WFk2S9sPn26cLj8j7BZXM0AKgWW6RBVB9rdgk9u4Q1VxiKSsZzMOfzIJ0JBVLp9avq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) 5e3qoV6cxD0Mnph6_YlKu64tTF7zkeacda1Gy32HrJA_QtidLYzPEibJAUNM8BMm3kOBL0m5kIG8EOpIMAh4WFk2S9sPn26cLj8j7BZXM0AKgWW6RBVB9rdgk9u4Q1VxiKSsZzMOfzIJ0JBVLp9avq

fd235 := client.Open("/B7_Q2wkRCf.txt", client.O_RDWR|client.O_CREATE)
if(fd235 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd229, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd195)
if(ret != 0) {
  panic("close failed")
}

//fd state: (61) wcCG57d_axXbeMp3AHirRnn7D3W84PqYak8d4j3WsNzPu263Yla1RcqMT7D7d

ret = client.Write(fd220, []byte("SK6tj1ivuYWbGi6M2BTMw3mt3VvVPjbvwbMn0TpxI52o_ehUWarkSsGX08Ql21Ib6u7IpC_0uJYvlbENC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (142) wcCG57d_axXbeMp3AHirRnn7D3W84PqYak8d4j3WsNzPu263Yla1RcqMT7D7dSK6tj1ivuYWbGi6M2BTMw3mt3VvVPjbvwbMn0TpxI52o_ehUWarkSsGX08Ql21Ib6u7IpC_0uJYvlbENC

fd236 := client.Open("/iWExPMWmQy.txt", client.O_RDWR|client.O_CREATE)
if(fd236 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd229, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (42) 7ZkJmDDHSpIJCat0ARtlwE5dI1XAQLrYHOG3xKS0vD

ret = client.Write(fd221, []byte("c_Q5ek2Z2v2DozvY_oLEhMtD3SrA7yVOjUublmgXD1EPkX_lx3Ko"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) 7ZkJmDDHSpIJCat0ARtlwE5dI1XAQLrYHOG3xKS0vDc_Q5ek2Z2v2DozvY_oLEhMtD3SrA7yVOjUublmgXD1EPkX_lx3Ko

ret = client.Close(fd220)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd235, []byte("UX9JbkDJ2KrtoMwK5SgzJXaBFbaK2E2E9QMxQLNMaKJYLgm2X6WU5Ezw1ZSI2yIZE_CHQPX5oskRqOWshnc6N5vzHBRrKT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) UX9JbkDJ2KrtoMwK5SgzJXaBFbaK2E2E9QMxQLNMaKJYLgm2X6WU5Ezw1ZSI2yIZE_CHQPX5oskRqOWshnc6N5vzHBRrKT

buf, ret = client.Read(fd221, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd225, []byte("r1O4gn3n37s5biNW0wp67bSjjl50Ckt9cU2LrfDI9KisfbpUXa9hekDnK2HWannEIPNucPZEboDYhf2FcRuPG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) r1O4gn3n37s5biNW0wp67bSjjl50Ckt9cU2LrfDI9KisfbpUXa9hekDnK2HWannEIPNucPZEboDYhf2FcRuPG

buf, ret = client.Read(fd233, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8e_mqMdepWmfIw5ko5QmxeG_cqjjItKsAy4bQK7bLTyhqFXMNQROb8HpaJB32ezezvriPdEIgCuOMMrahuOso5MvyFu") {
  panic("wrong data returned")
}

//fd state: (121) EqoRamUWoBsrVug9PVsyL9BKb4rBnp8e_mqMdepWmfIw5ko5QmxeG_cqjjItKsAy4bQK7bLTyhqFXMNQROb8HpaJB32ezezvriPdEIgCuOMMrahuOso5MvyFu

ret = client.Write(fd233, []byte("UbxghNU3JAsBEZ9ut8WOhAL4hlrfVTU1YO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (155) EqoRamUWoBsrVug9PVsyL9BKb4rBnp8e_mqMdepWmfIw5ko5QmxeG_cqjjItKsAy4bQK7bLTyhqFXMNQROb8HpaJB32ezezvriPdEIgCuOMMrahuOso5MvyFuUbxghNU3JAsBEZ9ut8WOhAL4hlrfVTU1YO

ret = client.Seek(fd223, 138, client.SEEK_SET)
if(ret != 138) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 138)
  panic("seek failed")
}


ret = client.Close(fd231)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd221, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd235, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


ret = client.Seek(fd235, 83, client.SEEK_SET)
if(ret != 83) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 83)
  panic("seek failed")
}


ret = client.Close(fd216)
if(ret != 0) {
  panic("close failed")
}


fd237 := client.Open("/EzNlejskyS.txt", client.O_RDWR|client.O_CREATE)
if(fd237 < 0) {
  panic("open failed")
}


ret = client.Seek(fd230, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


fd238 := client.Open("/k06ZPn3SfU.txt", client.O_RDWR|client.O_CREATE)
if(fd238 < 0) {
  panic("open failed")
}


ret = client.Seek(fd218, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


buf, ret = client.Read(fd234, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd236, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) W1jjQ9K4vCDt5v9PIx7Op2Bec

ret = client.Write(fd238, []byte("0CDpjaLp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) 0CDpjaLpvCDt5v9PIx7Op2Bec

fd239 := client.Open("/ApkZMhqPog.txt", client.O_RDWR|client.O_CREATE)
if(fd239 < 0) {
  panic("open failed")
}


fd240 := client.Open("/w2FXc3l6HU.txt", client.O_RDWR|client.O_CREATE)
if(fd240 < 0) {
  panic("open failed")
}


ret = client.Seek(fd239, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd229)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd239, []byte("ZC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) ZC

fd241 := client.Open("/ezy6lpPkBl.txt", client.O_RDWR|client.O_CREATE)
if(fd241 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd234, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd240, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "HjCrlg8CbeArGDTwFObXD6JAt8bsCztnzRF3tsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqt") {
  panic("wrong data returned")
}


ret = client.Close(fd241)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd234, []byte("5rHjocc8RSzv47ZMP3naKuAD69BBaoKiabvd9ARG_uTsIz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) 5rHjocc8RSzv47ZMP3naKuAD69BBaoKiabvd9ARG_uTsIz

ret = client.Seek(fd238, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}

//fd state: (46) 5rHjocc8RSzv47ZMP3naKuAD69BBaoKiabvd9ARG_uTsIz

ret = client.Write(fd234, []byte("fvqWrNKuA9ERSraGzYJKLAI9iXKDJ3Y6Jofgdfnsaw0z5W2zoCLaR4XkBt8sx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) 5rHjocc8RSzv47ZMP3naKuAD69BBaoKiabvd9ARG_uTsIzfvqWrNKuA9ERSraGzYJKLAI9iXKDJ3Y6Jofgdfnsaw0z5W2zoCLaR4XkBt8sx

ret = client.Close(fd225)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd230, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}

//fd state: (83) UX9JbkDJ2KrtoMwK5SgzJXaBFbaK2E2E9QMxQLNMaKJYLgm2X6WU5Ezw1ZSI2yIZE_CHQPX5oskRqOWshnc6N5vzHBRrKT

ret = client.Write(fd235, []byte("EwbCSoxsCdWf1otQ1jL0dcscjry49Xg_8o4PhL7l2pJI73KdQrKY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) UX9JbkDJ2KrtoMwK5SgzJXaBFbaK2E2E9QMxQLNMaKJYLgm2X6WU5Ezw1ZSI2yIZE_CHQPX5oskRqOWshncEwbCSoxsCdWf1otQ1jL0dcscjry49Xg_8o4PhL7l2pJI73KdQrKY
//fd state: (2) ZC

ret = client.Write(fd239, []byte("_FHsd6d_FEIu2itMbAKA2vHuiuiTO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) ZC_FHsd6d_FEIu2itMbAKA2vHuiuiTO

ret = client.Seek(fd238, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Seek(fd223, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Close(fd234)
if(ret != 0) {
  panic("close failed")
}


fd242 := client.Open("/jRPMyJY9NL.txt", client.O_RDWR|client.O_CREATE)
if(fd242 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd230, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MmmE8jQxyR") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd242, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd221)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd235, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


buf, ret = client.Read(fd230, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "A6ILmg") {
  panic("wrong data returned")
}


ret = client.Close(fd242)
if(ret != 0) {
  panic("close failed")
}

//fd state: (5) 2e4dUch

ret = client.Write(fd218, []byte("VU9ICU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) 2e4dUVU9ICU

fd243 := client.Open("/PIYzUD6TOm.txt", client.O_RDWR|client.O_CREATE)
if(fd243 < 0) {
  panic("open failed")
}

//fd state: (11) 2e4dUVU9ICU

ret = client.Write(fd218, []byte("74Rxg8HqM7NtMXq8yzGaz5rseNUNF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) 2e4dUVU9ICU74Rxg8HqM7NtMXq8yzGaz5rseNUNF

buf, ret = client.Read(fd235, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oMwK5SgzJXaBFbaK2E2E9QMxQLNMaKJYLgm2X6WU5Ezw1ZSI2yIZE_CHQPX5oskRqOWshncEwbCSoxsCdWf1otQ1jL0dcscj") {
  panic("wrong data returned")
}


fd244 := client.Open("/_BLukPDMrv.txt", client.O_RDWR|client.O_CREATE)
if(fd244 < 0) {
  panic("open failed")
}

//fd state: (77) HjCrlg8CbeArGDTwFObXD6JAt8bsCztnzRF3tsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqtbTKkbCPOgSlGLDAlyqKlCdZPjfs1dYcV7_0b0hqcLumGDJf4lD1WsmIQLzBcKh8DFCc8DYT8FUIpW5jPHIqG2AUdk5_wmfq7su93VoGr7Da5Ia

ret = client.Write(fd240, []byte("5R7KHop0EOZE5BKxcjKy9E"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) HjCrlg8CbeArGDTwFObXD6JAt8bsCztnzRF3tsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqt5R7KHop0EOZE5BKxcjKy9EZPjfs1dYcV7_0b0hqcLumGDJf4lD1WsmIQLzBcKh8DFCc8DYT8FUIpW5jPHIqG2AUdk5_wmfq7su93VoGr7Da5Ia

fd245 := client.Open("/bFM001bg35.txt", client.O_RDWR|client.O_CREATE)
if(fd245 < 0) {
  panic("open failed")
}


fd246 := client.Open("/DBm2gV4Brr.txt", client.O_RDWR|client.O_CREATE)
if(fd246 < 0) {
  panic("open failed")
}


ret = client.Seek(fd218, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


buf, ret = client.Read(fd235, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ry49Xg_8o4PhL7l2p") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd235, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JI73KdQrKY") {
  panic("wrong data returned")
}


ret = client.Seek(fd238, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


fd247 := client.Open("/kiQZc5jb4w.txt", client.O_RDWR|client.O_CREATE)
if(fd247 < 0) {
  panic("open failed")
}


ret = client.Close(fd238)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd236, []byte("NbLXU71JgOEE2sFAXExb_KZGNZHgA8Hdx2MxCBMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) NbLXU71JgOEE2sFAXExb_KZGNZHgA8Hdx2MxCBMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lG
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (46) bxdNzy37Aa0DP5mmPY7J8CK7C5oS2SMmmE8jQxyRA6ILmg9vDTBLDrUmC3b57ETMatdSHZX2f1hQ7XrWLvRacLFQz_N4

ret = client.Write(fd230, []byte("mFnC8oFN0IOPdmZZ0GsVhYAWu9dfnf9Y_vQ6w"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) bxdNzy37Aa0DP5mmPY7J8CK7C5oS2SMmmE8jQxyRA6ILmgmFnC8oFN0IOPdmZZ0GsVhYAWu9dfnf9Y_vQ6wacLFQz_N4

ret = client.Close(fd246)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd239, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd239, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


buf, ret = client.Read(fd244, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1MrsxizDXboeLrW7WR57R3UqVwaXEK7AsMGzcXdipgMQGgdj2lXLPwgKZUZhKUohriB2AI") {
  panic("wrong data returned")
}


ret = client.Seek(fd244, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd245, []byte("lO2uy1HZ9VsGUlen1r04sbnVnh46GbxKYq4N7JcygGFV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) lO2uy1HZ9VsGUlen1r04sbnVnh46GbxKYq4N7JcygGFV

buf, ret = client.Read(fd235, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd243)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd233, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Close(fd239)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd235)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd218, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


ret = client.Seek(fd230, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}

//fd state: (52) bxdNzy37Aa0DP5mmPY7J8CK7C5oS2SMmmE8jQxyRA6ILmgmFnC8oFN0IOPdmZZ0GsVhYAWu9dfnf9Y_vQ6wacLFQz_N4

ret = client.Write(fd230, []byte("1xrKMiIoQEe4_3Rb_UdklfaX3seE5J5FDlCCyBzNf1EBr5yywixgdCakEeMLzcceE7KbeoMP0kVzgN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) bxdNzy37Aa0DP5mmPY7J8CK7C5oS2SMmmE8jQxyRA6ILmgmFnC8o1xrKMiIoQEe4_3Rb_UdklfaX3seE5J5FDlCCyBzNf1EBr5yywixgdCakEeMLzcceE7KbeoMP0kVzgN

buf, ret = client.Read(fd230, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd230, 105, client.SEEK_SET)
if(ret != 105) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 105)
  panic("seek failed")
}


fd248 := client.Open("/gA71a_GvUv.txt", client.O_RDWR|client.O_CREATE)
if(fd248 < 0) {
  panic("open failed")
}


ret = client.Seek(fd223, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


fd249 := client.Open("/ezy6lpPkBl.txt", client.O_RDWR|client.O_CREATE)
if(fd249 < 0) {
  panic("open failed")
}


ret = client.Seek(fd237, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Close(fd233)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd248)
if(ret != 0) {
  panic("close failed")
}


fd250 := client.Open("/nHgSrOUu58.txt", client.O_RDWR|client.O_CREATE)
if(fd250 < 0) {
  panic("open failed")
}

//fd state: (105) bxdNzy37Aa0DP5mmPY7J8CK7C5oS2SMmmE8jQxyRA6ILmgmFnC8o1xrKMiIoQEe4_3Rb_UdklfaX3seE5J5FDlCCyBzNf1EBr5yywixgdCakEeMLzcceE7KbeoMP0kVzgN

ret = client.Write(fd230, []byte("dn4BqI8WGX7ihj5EwMiTF3XdEeKglfUCSeW7KKP3ireNmkJ0uVBoAx5lHzR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) bxdNzy37Aa0DP5mmPY7J8CK7C5oS2SMmmE8jQxyRA6ILmgmFnC8o1xrKMiIoQEe4_3Rb_UdklfaX3seE5J5FDlCCyBzNf1EBr5yywixgddn4BqI8WGX7ihj5EwMiTF3XdEeKglfUCSeW7KKP3ireNmkJ0uVBoAx5lHzR
//fd state: (99) HjCrlg8CbeArGDTwFObXD6JAt8bsCztnzRF3tsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqt5R7KHop0EOZE5BKxcjKy9EZPjfs1dYcV7_0b0hqcLumGDJf4lD1WsmIQLzBcKh8DFCc8DYT8FUIpW5jPHIqG2AUdk5_wmfq7su93VoGr7Da5Ia

ret = client.Write(fd240, []byte("KiN8ixcp1zPN12L5q7OHcP9OnofhUG23y8dikApbTExBATSHC2_bU4sitD5Kgcl43F6TDkIP5Vx3wP6JOKpHb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (184) HjCrlg8CbeArGDTwFObXD6JAt8bsCztnzRF3tsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqt5R7KHop0EOZE5BKxcjKy9EKiN8ixcp1zPN12L5q7OHcP9OnofhUG23y8dikApbTExBATSHC2_bU4sitD5Kgcl43F6TDkIP5Vx3wP6JOKpHb5Ia

buf, ret = client.Read(fd230, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd240, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5Ia") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd240, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd251 := client.Open("/fncK2K1IOC.txt", client.O_RDWR|client.O_CREATE)
if(fd251 < 0) {
  panic("open failed")
}


ret = client.Seek(fd230, 144, client.SEEK_SET)
if(ret != 144) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 144)
  panic("seek failed")
}

//fd state: (80) NbLXU71JgOEE2sFAXExb_KZGNZHgA8Hdx2MxCBMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lG

ret = client.Write(fd236, []byte("kueqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (153) NbLXU71JgOEE2sFAXExb_KZGNZHgA8Hdx2MxCBMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lGkueqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYS

buf, ret = client.Read(fd230, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3ireNmkJ0uVBoAx5lHzR") {
  panic("wrong data returned")
}


ret = client.Close(fd223)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 5rHjocc8RSzv47ZMP3naKuAD69BBaoKiabvd9ARG_uTsIzfvqWrNKuA9ERSraGzYJKLAI9iXKDJ3Y6Jofgdfnsaw0z5W2zoCLaR4XkBt8sx

ret = client.Write(fd247, []byte("2Nyse3gMaFwP2KOnOr4P6vVz5O8L2TSXpQZ4Hc7XVFDhvXbiUpcJaL63NUZzzUojgBln43vOHhupYiI_TJ_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) 2Nyse3gMaFwP2KOnOr4P6vVz5O8L2TSXpQZ4Hc7XVFDhvXbiUpcJaL63NUZzzUojgBln43vOHhupYiI_TJ_fnsaw0z5W2zoCLaR4XkBt8sx

ret = client.Close(fd240)
if(ret != 0) {
  panic("close failed")
}

//fd state: (83) 2Nyse3gMaFwP2KOnOr4P6vVz5O8L2TSXpQZ4Hc7XVFDhvXbiUpcJaL63NUZzzUojgBln43vOHhupYiI_TJ_fnsaw0z5W2zoCLaR4XkBt8sx

ret = client.Write(fd247, []byte("VgNLyM71ZpdGtNFf9Gjs02JPpc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (109) 2Nyse3gMaFwP2KOnOr4P6vVz5O8L2TSXpQZ4Hc7XVFDhvXbiUpcJaL63NUZzzUojgBln43vOHhupYiI_TJ_VgNLyM71ZpdGtNFf9Gjs02JPpc

ret = client.Close(fd249)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd245)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd247, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


buf, ret = client.Read(fd237, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Tj4doF7Z6xMoDY1IWr0z") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd250, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd252 := client.Open("/DWICQRTyV5.txt", client.O_RDWR|client.O_CREATE)
if(fd252 < 0) {
  panic("open failed")
}

//fd state: (24) 1MrsxizDXboeLrW7WR57R3UqVwaXEK7AsMGzcXdipgMQGgdj2lXLPwgKZUZhKUohriB2AIyp5NYnUpQa5NjXePvK4meHqPEED7L

ret = client.Write(fd244, []byte("49D8bABDEB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) 1MrsxizDXboeLrW7WR57R3Uq49D8bABDEBGzcXdipgMQGgdj2lXLPwgKZUZhKUohriB2AIyp5NYnUpQa5NjXePvK4meHqPEED7L

buf, ret = client.Read(fd251, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "s8WBl3LzAiq2bkhL32EELxo_gAVjtFO0vmOEhp_H3o_oMmo081HEG_0mXMQJlC5dC__zEajYEG0ZITFY") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd247, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3gMaFw") {
  panic("wrong data returned")
}


ret = client.Seek(fd252, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (153) NbLXU71JgOEE2sFAXExb_KZGNZHgA8Hdx2MxCBMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lGkueqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYS

ret = client.Write(fd236, []byte("tXtZ2C_qG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (162) NbLXU71JgOEE2sFAXExb_KZGNZHgA8Hdx2MxCBMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lGkueqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYStXtZ2C_qG

ret = client.Seek(fd230, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


ret = client.Seek(fd247, 66, client.SEEK_SET)
if(ret != 66) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 66)
  panic("seek failed")
}


fd253 := client.Open("/mFvxS3bd1N.txt", client.O_RDWR|client.O_CREATE)
if(fd253 < 0) {
  panic("open failed")
}


ret = client.Close(fd252)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd247)
if(ret != 0) {
  panic("close failed")
}

//fd state: (4) bxdNzy37Aa0DP5mmPY7J8CK7C5oS2SMmmE8jQxyRA6ILmgmFnC8o1xrKMiIoQEe4_3Rb_UdklfaX3seE5J5FDlCCyBzNf1EBr5yywixgddn4BqI8WGX7ihj5EwMiTF3XdEeKglfUCSeW7KKP3ireNmkJ0uVBoAx5lHzR

ret = client.Write(fd230, []byte("aW07RBZoeLphJZVapZk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) bxdNaW07RBZoeLphJZVapZk7C5oS2SMmmE8jQxyRA6ILmgmFnC8o1xrKMiIoQEe4_3Rb_UdklfaX3seE5J5FDlCCyBzNf1EBr5yywixgddn4BqI8WGX7ihj5EwMiTF3XdEeKglfUCSeW7KKP3ireNmkJ0uVBoAx5lHzR

ret = client.Seek(fd253, 111, client.SEEK_SET)
if(ret != 111) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 111)
  panic("seek failed")
}


ret = client.Close(fd218)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd244)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd251, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


ret = client.Seek(fd230, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


fd254 := client.Open("/oT8HBbhOs3.txt", client.O_RDWR|client.O_CREATE)
if(fd254 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd250, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd255 := client.Open("/q8LarXbP95.txt", client.O_RDWR|client.O_CREATE)
if(fd255 < 0) {
  panic("open failed")
}


ret = client.Seek(fd251, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


ret = client.Close(fd251)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd255)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd253)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd230, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


ret = client.Close(fd236)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd254, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd237, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}


buf, ret = client.Read(fd230, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MiIoQEe4_3Rb_UdklfaX3seE5J5FDlCCyBzNf1EBr") {
  panic("wrong data returned")
}


ret = client.Seek(fd237, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


buf, ret = client.Read(fd237, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "doF7Z6xMoDY1IWr0z39POY0d8StjfsrfEsNUhKWfci") {
  panic("wrong data returned")
}


ret = client.Seek(fd230, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}


ret = client.Close(fd254)
if(ret != 0) {
  panic("close failed")
}

//fd state: (78) bxdNaW07RBZoeLphJZVapZk7C5oS2SMmmE8jQxyRA6ILmgmFnC8o1xrKMiIoQEe4_3Rb_UdklfaX3seE5J5FDlCCyBzNf1EBr5yywixgddn4BqI8WGX7ihj5EwMiTF3XdEeKglfUCSeW7KKP3ireNmkJ0uVBoAx5lHzR

ret = client.Write(fd230, []byte("wHKPh7ufloTkFNkYHesyTFxkdlUvEVI15YNelqn66P9PBeb1I122HKTVUYSK3LDnY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (143) bxdNaW07RBZoeLphJZVapZk7C5oS2SMmmE8jQxyRA6ILmgmFnC8o1xrKMiIoQEe4_3Rb_UdklfaX3swHKPh7ufloTkFNkYHesyTFxkdlUvEVI15YNelqn66P9PBeb1I122HKTVUYSK3LDnYP3ireNmkJ0uVBoAx5lHzR

fd256 := client.Open("/oO3jPd5LqR.txt", client.O_RDWR|client.O_CREATE)
if(fd256 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd230, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "P3ireNmkJ0uVBoAx5lHzR") {
  panic("wrong data returned")
}


ret = client.Close(fd230)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd250, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd257 := client.Open("/2Kwg_3a7Gr.txt", client.O_RDWR|client.O_CREATE)
if(fd257 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd256, []byte("fiIX8W13tiwEN_S02pUGzCt3ffWsHMGTWTEnC4X51TMkJQORxv_ji0YEN0IV6a7pv1DJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) fiIX8W13tiwEN_S02pUGzCt3ffWsHMGTWTEnC4X51TMkJQORxv_ji0YEN0IV6a7pv1DJ

fd258 := client.Open("/Uq03S3GN57.txt", client.O_RDWR|client.O_CREATE)
if(fd258 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd258, []byte("7z0iqAkzNjxwUoP63900yAZEcgi9CePJauaVb6pC8PQrKhi5HLcTpHw7I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) 7z0iqAkzNjxwUoP63900yAZEcgi9CePJauaVb6pC8PQrKhi5HLcTpHw7I

ret = client.Seek(fd250, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd256)
if(ret != 0) {
  panic("close failed")
}


fd259 := client.Open("/CHOx395o99.txt", client.O_RDWR|client.O_CREATE)
if(fd259 < 0) {
  panic("open failed")
}


ret = client.Seek(fd257, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd260 := client.Open("/JPG9dFmZGp.txt", client.O_RDWR|client.O_CREATE)
if(fd260 < 0) {
  panic("open failed")
}


fd261 := client.Open("/6xOy8mDYIZ.txt", client.O_RDWR|client.O_CREATE)
if(fd261 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd260, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd258)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd250, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd262 := client.Open("/VKesPBLV5v.txt", client.O_RDWR|client.O_CREATE)
if(fd262 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd260, []byte("LFI5_kulXo2oa2YXf5a7IVl5L9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) LFI5_kulXo2oa2YXf5a7IVl5L9

ret = client.Seek(fd259, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd263 := client.Open("/1ME66mmhaH.txt", client.O_RDWR|client.O_CREATE)
if(fd263 < 0) {
  panic("open failed")
}


ret = client.Seek(fd260, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Seek(fd259, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd264 := client.Open("/DWICQRTyV5.txt", client.O_RDWR|client.O_CREATE)
if(fd264 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd265 := client.Open("/PIpopWJsly.txt", client.O_RDWR|client.O_CREATE)
if(fd265 < 0) {
  panic("open failed")
}


ret = client.Seek(fd250, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd262)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd257, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (2) LFI5_kulXo2oa2YXf5a7IVl5L9

ret = client.Write(fd260, []byte("7fpmfvbxwYldJCgpPriT82dC30zTWroobWEdQa08UqMFtC5oc3tDHocI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) LF7fpmfvbxwYldJCgpPriT82dC30zTWroobWEdQa08UqMFtC5oc3tDHocI

ret = client.Close(fd265)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd257, []byte("s8MQ46VzRfrBBbZXwBxXnIykCkgLxHgyxJJeXjantBjWWhiZqDrnYR1_RhM2ubOe5bZg7GQpHAKQoo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) s8MQ46VzRfrBBbZXwBxXnIykCkgLxHgyxJJeXjantBjWWhiZqDrnYR1_RhM2ubOe5bZg7GQpHAKQoo

ret = client.Seek(fd257, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}


ret = client.Seek(fd264, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd257, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd264, []byte("fOZldEMctXk8TODHCChQnhFVFTBA4Gj4AUrMxzozsrc64Syq0bnwtiQCQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) fOZldEMctXk8TODHCChQnhFVFTBA4Gj4AUrMxzozsrc64Syq0bnwtiQCQ

ret = client.Close(fd263)
if(ret != 0) {
  panic("close failed")
}

//fd state: (50) woTwOTj4doF7Z6xMoDY1IWr0z39POY0d8StjfsrfEsNUhKWfciWszPQ

ret = client.Write(fd237, []byte("JaZNk5qcJlyqLQEy8tM5k32z2xHzChXZgTPAgk2uI9o_m2WeMWMME2w"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) woTwOTj4doF7Z6xMoDY1IWr0z39POY0d8StjfsrfEsNUhKWfciJaZNk5qcJlyqLQEy8tM5k32z2xHzChXZgTPAgk2uI9o_m2WeMWMME2w

ret = client.Seek(fd257, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


fd266 := client.Open("/kWlBQtlrUQ.txt", client.O_RDWR|client.O_CREATE)
if(fd266 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd266, []byte("5iIMrEIpKj0J5BhwYZI7ua1YD0feWFqR5E7z_dIn0eluSelb4JG1U5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) 5iIMrEIpKj0J5BhwYZI7ua1YD0feWFqR5E7z_dIn0eluSelb4JG1U5

fd267 := client.Open("/0prf8WI95c.txt", client.O_RDWR|client.O_CREATE)
if(fd267 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd257, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YR1_RhM2ubOe5bZg7GQpHAKQoo") {
  panic("wrong data returned")
}


ret = client.Close(fd266)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd257, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Close(fd250)
if(ret != 0) {
  panic("close failed")
}

//fd state: (47) s8MQ46VzRfrBBbZXwBxXnIykCkgLxHgyxJJeXjantBjWWhiZqDrnYR1_RhM2ubOe5bZg7GQpHAKQoo

ret = client.Write(fd257, []byte("qPthJ_WnB75FE8OP7Tngv8yUKnEDoNgkx6c4eAw_4FuEliZ981vsv0CzVqTpsfzkvKdgPJ8sQGYBFh52uqrsB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) s8MQ46VzRfrBBbZXwBxXnIykCkgLxHgyxJJeXjantBjWWhiqPthJ_WnB75FE8OP7Tngv8yUKnEDoNgkx6c4eAw_4FuEliZ981vsv0CzVqTpsfzkvKdgPJ8sQGYBFh52uqrsB

fd268 := client.Open("/ja3oo3k0R9.txt", client.O_RDWR|client.O_CREATE)
if(fd268 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd237, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd268, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd269 := client.Open("/6pv1r2ZR8B.txt", client.O_RDWR|client.O_CREATE)
if(fd269 < 0) {
  panic("open failed")
}


ret = client.Close(fd268)
if(ret != 0) {
  panic("close failed")
}


fd270 := client.Open("/ZOIAW57u8n.txt", client.O_RDWR|client.O_CREATE)
if(fd270 < 0) {
  panic("open failed")
}


ret = client.Close(fd237)
if(ret != 0) {
  panic("close failed")
}

//fd state: (132) s8MQ46VzRfrBBbZXwBxXnIykCkgLxHgyxJJeXjantBjWWhiqPthJ_WnB75FE8OP7Tngv8yUKnEDoNgkx6c4eAw_4FuEliZ981vsv0CzVqTpsfzkvKdgPJ8sQGYBFh52uqrsB

ret = client.Write(fd257, []byte("sThb1F4MUyTTsrWzYJjGnX7atgDVY4E5kU4WOYyBkITXHi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (178) s8MQ46VzRfrBBbZXwBxXnIykCkgLxHgyxJJeXjantBjWWhiqPthJ_WnB75FE8OP7Tngv8yUKnEDoNgkx6c4eAw_4FuEliZ981vsv0CzVqTpsfzkvKdgPJ8sQGYBFh52uqrsBsThb1F4MUyTTsrWzYJjGnX7atgDVY4E5kU4WOYyBkITXHi

buf, ret = client.Read(fd259, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd267, []byte("O_BkGN7xo7x885c95fxznNwGWIWrELz2WUKVvTF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) O_BkGN7xo7x885c95fxznNwGWIWrELz2WUKVvTF
//fd state: (58) LF7fpmfvbxwYldJCgpPriT82dC30zTWroobWEdQa08UqMFtC5oc3tDHocI

ret = client.Write(fd260, []byte("rdeYeFlCFb26smvfxaeeqIq3JUpHRiKpHC5qchb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) LF7fpmfvbxwYldJCgpPriT82dC30zTWroobWEdQa08UqMFtC5oc3tDHocIrdeYeFlCFb26smvfxaeeqIq3JUpHRiKpHC5qchb

ret = client.Close(fd261)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd270, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd271 := client.Open("/T7X58qc9mk.txt", client.O_RDWR|client.O_CREATE)
if(fd271 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd269, []byte("ekuKyzv8Z6JzR3lHQfq64Me_a27K1BH2yup6nWcAiFtbewKktdFN2BHvfHNbQItEo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) ekuKyzv8Z6JzR3lHQfq64Me_a27K1BH2yup6nWcAiFtbewKktdFN2BHvfHNbQItEo

buf, ret = client.Read(fd269, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd272 := client.Open("/Avi_dS2O4_.txt", client.O_RDWR|client.O_CREATE)
if(fd272 < 0) {
  panic("open failed")
}


fd273 := client.Open("/x7r7aEPrq5.txt", client.O_RDWR|client.O_CREATE)
if(fd273 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd264, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd257, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


fd274 := client.Open("/kiQZc5jb4w.txt", client.O_RDWR|client.O_CREATE)
if(fd274 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd273, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "FVufQKNseI93OoaNQW61pdh5Sb9b3PiDmzBMgWwIklMD") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd271, []byte("Vxid_T73syQ9fc63CYUwceJE9NRDfFUikfoW4AyQcXtSqmBtJIW0WjdUsYfDD7Of3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) Vxid_T73syQ9fc63CYUwceJE9NRDfFUikfoW4AyQcXtSqmBtJIW0WjdUsYfDD7Of3

buf, ret = client.Read(fd267, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (65) Vxid_T73syQ9fc63CYUwceJE9NRDfFUikfoW4AyQcXtSqmBtJIW0WjdUsYfDD7Of3

ret = client.Write(fd271, []byte("HVGaHObjxGMmTQqPFDinvR3XN4i7B6sppzqEnxCbinQixL8HQAh807st09n5jxT6Y9QtD89zJt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (139) Vxid_T73syQ9fc63CYUwceJE9NRDfFUikfoW4AyQcXtSqmBtJIW0WjdUsYfDD7Of3HVGaHObjxGMmTQqPFDinvR3XN4i7B6sppzqEnxCbinQixL8HQAh807st09n5jxT6Y9QtD89zJt

buf, ret = client.Read(fd272, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4LgZZ8tIb6AJv") {
  panic("wrong data returned")
}


fd275 := client.Open("/9HehunRkpL.txt", client.O_RDWR|client.O_CREATE)
if(fd275 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd259, []byte("d0DRyGvTe7Jr648NvHumNLQK0y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) d0DRyGvTe7Jr648NvHumNLQK0y
//fd state: (13) 4LgZZ8tIb6AJvWLMklVa6A1Euio40rOy7j2_C0WDcu1vy5yRqNX17UtBNSHw2tmSOHwudow96e7weF7Xpk21F6kVIJJDRO70LyQjUAbZnudqYaiBC

ret = client.Write(fd272, []byte("tHhsoDzD7qTazPky2MgJtxtoSv7jzn4ZLrxBm4FLRI8F1YduEtoS1YfH_fxVG3y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) 4LgZZ8tIb6AJvtHhsoDzD7qTazPky2MgJtxtoSv7jzn4ZLrxBm4FLRI8F1YduEtoS1YfH_fxVG3yeF7Xpk21F6kVIJJDRO70LyQjUAbZnudqYaiBC

ret = client.Seek(fd275, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd273, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


fd276 := client.Open("/pnwfk4ua8G.txt", client.O_RDWR|client.O_CREATE)
if(fd276 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd260, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd272, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


buf, ret = client.Read(fd259, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd275)
if(ret != 0) {
  panic("close failed")
}

//fd state: (139) Vxid_T73syQ9fc63CYUwceJE9NRDfFUikfoW4AyQcXtSqmBtJIW0WjdUsYfDD7Of3HVGaHObjxGMmTQqPFDinvR3XN4i7B6sppzqEnxCbinQixL8HQAh807st09n5jxT6Y9QtD89zJt

ret = client.Write(fd271, []byte("c663naDLrdkVYGNfWTR9_HENCB5vnwpwJK7VcrS2OgyfB0OL90QxNKYCRfgU1TUtFHaBzWT6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (211) Vxid_T73syQ9fc63CYUwceJE9NRDfFUikfoW4AyQcXtSqmBtJIW0WjdUsYfDD7Of3HVGaHObjxGMmTQqPFDinvR3XN4i7B6sppzqEnxCbinQixL8HQAh807st09n5jxT6Y9QtD89zJtc663naDLrdkVYGNfWTR9_HENCB5vnwpwJK7VcrS2OgyfB0OL90QxNKYCRfgU1TUtFHaBzWT6

ret = client.Seek(fd264, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Close(fd270)
if(ret != 0) {
  panic("close failed")
}

//fd state: (65) ekuKyzv8Z6JzR3lHQfq64Me_a27K1BH2yup6nWcAiFtbewKktdFN2BHvfHNbQItEo

ret = client.Write(fd269, []byte("d3Q7_uueI7ee3RZVHRjmALZjieeMvs3csbm4qTxRwfLRlarEJhIMt7353vJ6SWskiRrxK1RixLygAP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (143) ekuKyzv8Z6JzR3lHQfq64Me_a27K1BH2yup6nWcAiFtbewKktdFN2BHvfHNbQItEod3Q7_uueI7ee3RZVHRjmALZjieeMvs3csbm4qTxRwfLRlarEJhIMt7353vJ6SWskiRrxK1RixLygAP
//fd state: (143) ekuKyzv8Z6JzR3lHQfq64Me_a27K1BH2yup6nWcAiFtbewKktdFN2BHvfHNbQItEod3Q7_uueI7ee3RZVHRjmALZjieeMvs3csbm4qTxRwfLRlarEJhIMt7353vJ6SWskiRrxK1RixLygAP

ret = client.Write(fd269, []byte("PIr09M6ApJyMLrH3vYbV4ZcesBxNREnSDu8LOk66u8Hr1JedB7h9nnYFHEHSOYTAalngH2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (213) ekuKyzv8Z6JzR3lHQfq64Me_a27K1BH2yup6nWcAiFtbewKktdFN2BHvfHNbQItEod3Q7_uueI7ee3RZVHRjmALZjieeMvs3csbm4qTxRwfLRlarEJhIMt7353vJ6SWskiRrxK1RixLygAPPIr09M6ApJyMLrH3vYbV4ZcesBxNREnSDu8LOk66u8Hr1JedB7h9nnYFHEHSOYTAalngH2
//fd state: (0) 

ret = client.Write(fd276, []byte("04w0Xl_aqJNXI157zNyaTXV6NzPKqh6y20ocNMtD7tgU0I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) 04w0Xl_aqJNXI157zNyaTXV6NzPKqh6y20ocNMtD7tgU0I

buf, ret = client.Read(fd272, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hsoDzD7qTazPky2MgJtxtoSv7jzn4ZLrxBm4FLRI8F1YduEtoS1YfH_fxVG3yeF7Xpk21F6kVIJJDRO70") {
  panic("wrong data returned")
}


ret = client.Seek(fd264, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


buf, ret = client.Read(fd267, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd257, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "CkgLxHgyxJJeXjantBjWWhiqPthJ") {
  panic("wrong data returned")
}


fd277 := client.Open("/oT8HBbhOs3.txt", client.O_RDWR|client.O_CREATE)
if(fd277 < 0) {
  panic("open failed")
}


fd278 := client.Open("/pvFMa6eFhT.txt", client.O_RDWR|client.O_CREATE)
if(fd278 < 0) {
  panic("open failed")
}


ret = client.Close(fd257)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd278, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd278, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd278)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd260)
if(ret != 0) {
  panic("close failed")
}


fd279 := client.Open("/_gbGbZZtm6.txt", client.O_RDWR|client.O_CREATE)
if(fd279 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd269, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd280 := client.Open("/9HehunRkpL.txt", client.O_RDWR|client.O_CREATE)
if(fd280 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd271, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd267)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd259, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


buf, ret = client.Read(fd271, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd272, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LyQjUAbZnudqYaiBC") {
  panic("wrong data returned")
}


fd281 := client.Open("/6JfRSSuVna.txt", client.O_RDWR|client.O_CREATE)
if(fd281 < 0) {
  panic("open failed")
}


fd282 := client.Open("/q757Jr9Xdi.txt", client.O_RDWR|client.O_CREATE)
if(fd282 < 0) {
  panic("open failed")
}


fd283 := client.Open("/KcqL0rAr8o.txt", client.O_RDWR|client.O_CREATE)
if(fd283 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd280, []byte("kwTE5Xbd_0pgRLA1SAdt5PPhDiNVKX3uEE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) kwTE5Xbd_0pgRLA1SAdt5PPhDiNVKX3uEE

ret = client.Close(fd272)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 2Nyse3gMaFwP2KOnOr4P6vVz5O8L2TSXpQZ4Hc7XVFDhvXbiUpcJaL63NUZzzUojgBln43vOHhupYiI_TJ_VgNLyM71ZpdGtNFf9Gjs02JPpc

ret = client.Write(fd274, []byte("6mAjiaksdRIdO14dmpbSU1LsqvFd2e_JCGomwY1Tewttinm6S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) 6mAjiaksdRIdO14dmpbSU1LsqvFd2e_JCGomwY1Tewttinm6SpcJaL63NUZzzUojgBln43vOHhupYiI_TJ_VgNLyM71ZpdGtNFf9Gjs02JPpc

buf, ret = client.Read(fd271, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd277, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd271, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd284 := client.Open("/zf1Y4vkWDq.txt", client.O_RDWR|client.O_CREATE)
if(fd284 < 0) {
  panic("open failed")
}

//fd state: (46) 04w0Xl_aqJNXI157zNyaTXV6NzPKqh6y20ocNMtD7tgU0I

ret = client.Write(fd276, []byte("nSsN5Gsu6z5KAnVQtOA_4byKwQwmSFUzh_bY5BwGs6TNTkhaa1T9N"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) 04w0Xl_aqJNXI157zNyaTXV6NzPKqh6y20ocNMtD7tgU0InSsN5Gsu6z5KAnVQtOA_4byKwQwmSFUzh_bY5BwGs6TNTkhaa1T9N

fd285 := client.Open("/6hTT93U1mS.txt", client.O_RDWR|client.O_CREATE)
if(fd285 < 0) {
  panic("open failed")
}


ret = client.Close(fd276)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd264, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


buf, ret = client.Read(fd264, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "zozsrc64Syq0bnwtiQCQ") {
  panic("wrong data returned")
}


fd286 := client.Open("/Wod_MWssCh.txt", client.O_RDWR|client.O_CREATE)
if(fd286 < 0) {
  panic("open failed")
}

//fd state: (24) d0DRyGvTe7Jr648NvHumNLQK0y

ret = client.Write(fd259, []byte("gQhPZo2jquO_ays0zgJYawIxegMUsIcgjGfb4we9l24aYosHqe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) d0DRyGvTe7Jr648NvHumNLQKgQhPZo2jquO_ays0zgJYawIxegMUsIcgjGfb4we9l24aYosHqe
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd287 := client.Open("/dfRXaFO9tW.txt", client.O_RDWR|client.O_CREATE)
if(fd287 < 0) {
  panic("open failed")
}


ret = client.Close(fd273)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd283)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 7ZkJmDDHSpIJCat0ARtlwE5dI1XAQLrYHOG3xKS0vDc_Q5ek2Z2v2DozvY_oLEhMtD3SrA7yVOjUublmgXD1EPkX_lx3Ko

ret = client.Write(fd286, []byte("VEueXWkpsqroazahWlbxiFY8SwMaj0ty5y3Xjo8j98Ar1Iazx1tSPfOcJyOpjsF0C2PJUga3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) VEueXWkpsqroazahWlbxiFY8SwMaj0ty5y3Xjo8j98Ar1Iazx1tSPfOcJyOpjsF0C2PJUga3VOjUublmgXD1EPkX_lx3Ko

buf, ret = client.Read(fd282, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd280)
if(ret != 0) {
  panic("close failed")
}

//fd state: (57) fOZldEMctXk8TODHCChQnhFVFTBA4Gj4AUrMxzozsrc64Syq0bnwtiQCQ

ret = client.Write(fd264, []byte("nh_odKvaoc7CnJQgRzKWXTTLg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) fOZldEMctXk8TODHCChQnhFVFTBA4Gj4AUrMxzozsrc64Syq0bnwtiQCQnh_odKvaoc7CnJQgRzKWXTTLg

ret = client.Close(fd285)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd284)
if(ret != 0) {
  panic("close failed")
}

//fd state: (213) ekuKyzv8Z6JzR3lHQfq64Me_a27K1BH2yup6nWcAiFtbewKktdFN2BHvfHNbQItEod3Q7_uueI7ee3RZVHRjmALZjieeMvs3csbm4qTxRwfLRlarEJhIMt7353vJ6SWskiRrxK1RixLygAPPIr09M6ApJyMLrH3vYbV4ZcesBxNREnSDu8LOk66u8Hr1JedB7h9nnYFHEHSOYTAalngH2

ret = client.Write(fd269, []byte("IVIUl4Ytz5UKnEWkVjbzNDOzQOg5pwBYJt7WYkUdkYlKaLGLYtw9BOV3LRFQzJUeqoPUGBvYTt8cWPfr5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (294) ekuKyzv8Z6JzR3lHQfq64Me_a27K1BH2yup6nWcAiFtbewKktdFN2BHvfHNbQItEod3Q7_uueI7ee3RZVHRjmALZjieeMvs3csbm4qTxRwfLRlarEJhIMt7353vJ6SWskiRrxK1RixLygAPPIr09M6ApJyMLrH3vYbV4ZcesBxNREnSDu8LOk66u8Hr1JedB7h9nnYFHEHSOYTAalngH2IVIUl4Ytz5UKnEWkVjbzNDOzQOg5pwBYJt7WYkUdkYlKaLGLYtw9BOV3LRFQzJUeqoPUGBvYTt8cWPfr5

ret = client.Close(fd259)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd274)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd277, []byte("HI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) HI

buf, ret = client.Read(fd282, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (2) HI

ret = client.Write(fd277, []byte("9u133pXLBtbPsdqZjbxkcuGVIyzxjc8GJuSQyi7lguTvEmp3o97UbG9USvXe9s00"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) HI9u133pXLBtbPsdqZjbxkcuGVIyzxjc8GJuSQyi7lguTvEmp3o97UbG9USvXe9s00

buf, ret = client.Read(fd279, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd271, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd282)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd281, []byte("y0X7vxuPZnVW4CFvyHzD43W1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) y0X7vxuPZnVW4CFvyHzD43W1

ret = client.Seek(fd286, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd277, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd286)
if(ret != 0) {
  panic("close failed")
}


fd288 := client.Open("/SZ3w7aLh_u.txt", client.O_RDWR|client.O_CREATE)
if(fd288 < 0) {
  panic("open failed")
}


fd289 := client.Open("/1ME66mmhaH.txt", client.O_RDWR|client.O_CREATE)
if(fd289 < 0) {
  panic("open failed")
}


ret = client.Close(fd277)
if(ret != 0) {
  panic("close failed")
}


fd290 := client.Open("/zvfpB2u_jD.txt", client.O_RDWR|client.O_CREATE)
if(fd290 < 0) {
  panic("open failed")
}

//fd state: (0) YWFH3qBKfvN7Rvm3mBul2AMRxbzAY6zOvQPu6ZO4NhjU1gHSaX50BDx

ret = client.Write(fd287, []byte("BQzXB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) BQzXBqBKfvN7Rvm3mBul2AMRxbzAY6zOvQPu6ZO4NhjU1gHSaX50BDx

buf, ret = client.Read(fd264, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd264)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd287, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd279)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd271)
if(ret != 0) {
  panic("close failed")
}


fd291 := client.Open("/PiOdP6CVcO.txt", client.O_RDWR|client.O_CREATE)
if(fd291 < 0) {
  panic("open failed")
}

//fd state: (24) y0X7vxuPZnVW4CFvyHzD43W1

ret = client.Write(fd281, []byte("t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) y0X7vxuPZnVW4CFvyHzD43W1t

fd292 := client.Open("/ydk9BRFkKJ.txt", client.O_RDWR|client.O_CREATE)
if(fd292 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd291, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd281, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


fd293 := client.Open("/UcVD44wg3a.txt", client.O_RDWR|client.O_CREATE)
if(fd293 < 0) {
  panic("open failed")
}


ret = client.Close(fd287)
if(ret != 0) {
  panic("close failed")
}


fd294 := client.Open("/l0WXb6xP4b.txt", client.O_RDWR|client.O_CREATE)
if(fd294 < 0) {
  panic("open failed")
}


ret = client.Seek(fd288, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


ret = client.Seek(fd289, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


buf, ret = client.Read(fd291, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd291, []byte("AkaXbzc28ZwkaWjIHD213MrNJOBnweNfv_JgFK3AQQ3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) AkaXbzc28ZwkaWjIHD213MrNJOBnweNfv_JgFK3AQQ3
//fd state: (37) TUOGs3FBdlpHEJ6nZBrHXRQjOspCjq3YmhDlnZVOF2uRt9ej0h8Z76h4bHfHK0zPMb1r8fudHfzW0Uygdj8iyALtG8

ret = client.Write(fd289, []byte("l5gXTVyIaY6vXRza"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) TUOGs3FBdlpHEJ6nZBrHXRQjOspCjq3YmhDlnl5gXTVyIaY6vXRza6h4bHfHK0zPMb1r8fudHfzW0Uygdj8iyALtG8
//fd state: (0) 

ret = client.Write(fd290, []byte("xyHoXDc70wYMuR1TKTL9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) xyHoXDc70wYMuR1TKTL9
//fd state: (0) 

ret = client.Write(fd293, []byte("NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7n

ret = client.Close(fd292)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd290, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd295 := client.Open("/IwCBPIfR8t.txt", client.O_RDWR|client.O_CREATE)
if(fd295 < 0) {
  panic("open failed")
}


ret = client.Seek(fd290, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd296 := client.Open("/pZv_3JdBVv.txt", client.O_RDWR|client.O_CREATE)
if(fd296 < 0) {
  panic("open failed")
}


fd297 := client.Open("/R8FiGZHyRq.txt", client.O_RDWR|client.O_CREATE)
if(fd297 < 0) {
  panic("open failed")
}


ret = client.Close(fd291)
if(ret != 0) {
  panic("close failed")
}


fd298 := client.Open("/JgvYBa4b20.txt", client.O_RDWR|client.O_CREATE)
if(fd298 < 0) {
  panic("open failed")
}


ret = client.Close(fd269)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd289, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


fd299 := client.Open("/Wuhr1LEWJy.txt", client.O_RDWR|client.O_CREATE)
if(fd299 < 0) {
  panic("open failed")
}


ret = client.Close(fd296)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd294)
if(ret != 0) {
  panic("close failed")
}


fd300 := client.Open("/GrMYXD3DIK.txt", client.O_RDWR|client.O_CREATE)
if(fd300 < 0) {
  panic("open failed")
}


ret = client.Seek(fd298, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd289, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


fd301 := client.Open("/jMLCbPratd.txt", client.O_RDWR|client.O_CREATE)
if(fd301 < 0) {
  panic("open failed")
}

//fd state: (81) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7n

ret = client.Write(fd293, []byte("T9qmn6nJWWOFonsu5Qsah3sYRNfUtGPqyjZlSpLMNZwX202pamULAI1fXpWH1XEKxhzO4xSEYBnWQFPkXYchgcqtd7O"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (172) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7nT9qmn6nJWWOFonsu5Qsah3sYRNfUtGPqyjZlSpLMNZwX202pamULAI1fXpWH1XEKxhzO4xSEYBnWQFPkXYchgcqtd7O
//fd state: (172) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7nT9qmn6nJWWOFonsu5Qsah3sYRNfUtGPqyjZlSpLMNZwX202pamULAI1fXpWH1XEKxhzO4xSEYBnWQFPkXYchgcqtd7O

ret = client.Write(fd293, []byte("znfyVDk2vZAuS_FZr4ubPl5I0rqK1jvThMKhcAnv6OrItYXdJKNHJosp7OZVU0mqQ4iuhaqehu1LBMYDDkL9J32yXCvUKauzEEq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (271) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7nT9qmn6nJWWOFonsu5Qsah3sYRNfUtGPqyjZlSpLMNZwX202pamULAI1fXpWH1XEKxhzO4xSEYBnWQFPkXYchgcqtd7OznfyVDk2vZAuS_FZr4ubPl5I0rqK1jvThMKhcAnv6OrItYXdJKNHJosp7OZVU0mqQ4iuhaqehu1LBMYDDkL9J32yXCvUKauzEEq
//fd state: (0) 

ret = client.Write(fd295, []byte("JKvPnuLf0QPpRBKzGGz2Qp1rjJbmPe7KOO1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) JKvPnuLf0QPpRBKzGGz2Qp1rjJbmPe7KOO1
//fd state: (0) 

ret = client.Write(fd301, []byte("wqYNkMTfDmBxC3KBQ7wVlzECLvRsXYFf1RJS1_yKSELmkB4VT5P5rLSV22Yw5ML53z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) wqYNkMTfDmBxC3KBQ7wVlzECLvRsXYFf1RJS1_yKSELmkB4VT5P5rLSV22Yw5ML53z

ret = client.Seek(fd298, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd293, 126, client.SEEK_SET)
if(ret != 126) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 126)
  panic("seek failed")
}

//fd state: (3) xyHoXDc70wYMuR1TKTL9

ret = client.Write(fd290, []byte("Fttu_tPoeRaZg9dy6ghopA_AnUBK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) xyHFttu_tPoeRaZg9dy6ghopA_AnUBK

buf, ret = client.Read(fd289, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vXRza6h4bHfHK0zPMb1r8fudHfzW0Uygdj8iyALtG8") {
  panic("wrong data returned")
}


fd302 := client.Open("/kWlBQtlrUQ.txt", client.O_RDWR|client.O_CREATE)
if(fd302 < 0) {
  panic("open failed")
}

//fd state: (23) y0X7vxuPZnVW4CFvyHzD43W1t

ret = client.Write(fd281, []byte("_Mj2kWxY9oylLJHagx6V0S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) y0X7vxuPZnVW4CFvyHzD43W_Mj2kWxY9oylLJHagx6V0S

ret = client.Close(fd300)
if(ret != 0) {
  panic("close failed")
}


fd303 := client.Open("/XgqlK6vDL2.txt", client.O_RDWR|client.O_CREATE)
if(fd303 < 0) {
  panic("open failed")
}


ret = client.Seek(fd289, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Close(fd295)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd281)
if(ret != 0) {
  panic("close failed")
}

//fd state: (9) F8iipPiHaSATt6OZaNFa

ret = client.Write(fd288, []byte("wxWX_Ch2fqid_H04r9m99xtMgy558xL1LtfCRLyGFsfFDPxug0cN26ZKRYMkVHZLMQzBUC5AFp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) F8iipPiHawxWX_Ch2fqid_H04r9m99xtMgy558xL1LtfCRLyGFsfFDPxug0cN26ZKRYMkVHZLMQzBUC5AFp

ret = client.Seek(fd290, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Seek(fd293, 152, client.SEEK_SET)
if(ret != 152) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 152)
  panic("seek failed")
}


buf, ret = client.Read(fd301, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd302, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


buf, ret = client.Read(fd303, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd297, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd289)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd288, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


buf, ret = client.Read(fd303, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd288)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd297, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd293, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "EYBnWQFPkXYchgcqtd7OznfyVDk2vZAu") {
  panic("wrong data returned")
}


ret = client.Seek(fd293, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


ret = client.Seek(fd290, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Seek(fd290, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Seek(fd293, 82, client.SEEK_SET)
if(ret != 82) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 82)
  panic("seek failed")
}


fd304 := client.Open("/OqvlaCPI05.txt", client.O_RDWR|client.O_CREATE)
if(fd304 < 0) {
  panic("open failed")
}


ret = client.Seek(fd298, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd298, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd297, []byte("9bjFpe9PcAhYWrCu2q5CnNcIFhuZ4vOrllH3jnpkjgW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) 9bjFpe9PcAhYWrCu2q5CnNcIFhuZ4vOrllH3jnpkjgW
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd290, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "eRaZg9dy6ghopA_AnUBK") {
  panic("wrong data returned")
}


fd305 := client.Open("/ja3oo3k0R9.txt", client.O_RDWR|client.O_CREATE)
if(fd305 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd290, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd305, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd297, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


ret = client.Close(fd304)
if(ret != 0) {
  panic("close failed")
}


fd306 := client.Open("/4k_5BxY9TN.txt", client.O_RDWR|client.O_CREATE)
if(fd306 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd303, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd298, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd307 := client.Open("/eCvA_JYFK6.txt", client.O_RDWR|client.O_CREATE)
if(fd307 < 0) {
  panic("open failed")
}


ret = client.Close(fd305)
if(ret != 0) {
  panic("close failed")
}

//fd state: (11) 5iIMrEIpKj0J5BhwYZI7ua1YD0feWFqR5E7z_dIn0eluSelb4JG1U5

ret = client.Write(fd302, []byte("czDjTGp9weRYWPxqoxlt1nZMaCdcRSmE3MkzfqPAXF1J5Sbv2jHCPexDR97PNaKH5hqMRzd8RUedVksHf23UNuD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) 5iIMrEIpKj0czDjTGp9weRYWPxqoxlt1nZMaCdcRSmE3MkzfqPAXF1J5Sbv2jHCPexDR97PNaKH5hqMRzd8RUedVksHf23UNuD
//fd state: (0) 

ret = client.Write(fd303, []byte("WZ_fJw_Z4SOlCL5rgTDoGtIXFrrSyZVJ7lnMZfRnNoIwAHXRQO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) WZ_fJw_Z4SOlCL5rgTDoGtIXFrrSyZVJ7lnMZfRnNoIwAHXRQO

ret = client.Seek(fd297, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


fd308 := client.Open("/CHOx395o99.txt", client.O_RDWR|client.O_CREATE)
if(fd308 < 0) {
  panic("open failed")
}


fd309 := client.Open("/879TeeQEuG.txt", client.O_RDWR|client.O_CREATE)
if(fd309 < 0) {
  panic("open failed")
}


ret = client.Close(fd298)
if(ret != 0) {
  panic("close failed")
}

//fd state: (31) xyHFttu_tPoeRaZg9dy6ghopA_AnUBK

ret = client.Write(fd290, []byte("8xGWi5_m2Nl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) xyHFttu_tPoeRaZg9dy6ghopA_AnUBK8xGWi5_m2Nl

fd310 := client.Open("/vA8vL6aicz.txt", client.O_RDWR|client.O_CREATE)
if(fd310 < 0) {
  panic("open failed")
}


fd311 := client.Open("/JPG9dFmZGp.txt", client.O_RDWR|client.O_CREATE)
if(fd311 < 0) {
  panic("open failed")
}


fd312 := client.Open("/UkyvZgKXCW.txt", client.O_RDWR|client.O_CREATE)
if(fd312 < 0) {
  panic("open failed")
}


ret = client.Close(fd303)
if(ret != 0) {
  panic("close failed")
}

//fd state: (82) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7nT9qmn6nJWWOFonsu5Qsah3sYRNfUtGPqyjZlSpLMNZwX202pamULAI1fXpWH1XEKxhzO4xSEYBnWQFPkXYchgcqtd7OznfyVDk2vZAuS_FZr4ubPl5I0rqK1jvThMKhcAnv6OrItYXdJKNHJosp7OZVU0mqQ4iuhaqehu1LBMYDDkL9J32yXCvUKauzEEq

ret = client.Write(fd293, []byte("xz1cfPxp9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7nTxz1cfPxp9OFonsu5Qsah3sYRNfUtGPqyjZlSpLMNZwX202pamULAI1fXpWH1XEKxhzO4xSEYBnWQFPkXYchgcqtd7OznfyVDk2vZAuS_FZr4ubPl5I0rqK1jvThMKhcAnv6OrItYXdJKNHJosp7OZVU0mqQ4iuhaqehu1LBMYDDkL9J32yXCvUKauzEEq

buf, ret = client.Read(fd297, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lH3jnpkjgW") {
  panic("wrong data returned")
}


ret = client.Close(fd302)
if(ret != 0) {
  panic("close failed")
}


fd313 := client.Open("/XYg89mK8WS.txt", client.O_RDWR|client.O_CREATE)
if(fd313 < 0) {
  panic("open failed")
}


ret = client.Seek(fd301, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


fd314 := client.Open("/tL_sr4v2I3.txt", client.O_RDWR|client.O_CREATE)
if(fd314 < 0) {
  panic("open failed")
}


ret = client.Close(fd313)
if(ret != 0) {
  panic("close failed")
}


fd315 := client.Open("/eCvA_JYFK6.txt", client.O_RDWR|client.O_CREATE)
if(fd315 < 0) {
  panic("open failed")
}


fd316 := client.Open("/HG_fqUj4cf.txt", client.O_RDWR|client.O_CREATE)
if(fd316 < 0) {
  panic("open failed")
}


fd317 := client.Open("/YSlKQ6IsZU.txt", client.O_RDWR|client.O_CREATE)
if(fd317 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd308, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "d0DRyGvTe7Jr648NvHumNLQKgQhPZo") {
  panic("wrong data returned")
}


fd318 := client.Open("/LsRInew_XM.txt", client.O_RDWR|client.O_CREATE)
if(fd318 < 0) {
  panic("open failed")
}


ret = client.Seek(fd293, 162, client.SEEK_SET)
if(ret != 162) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 162)
  panic("seek failed")
}


ret = client.Close(fd311)
if(ret != 0) {
  panic("close failed")
}

//fd state: (30) d0DRyGvTe7Jr648NvHumNLQKgQhPZo2jquO_ays0zgJYawIxegMUsIcgjGfb4we9l24aYosHqe

ret = client.Write(fd308, []byte("jzFFuQYSYYYh93TPQbqGDoz61W6XgF4D0dv4KWlpQZNEct87UKobSBzXi4eOdgsbHxujmfLkD_nR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (106) d0DRyGvTe7Jr648NvHumNLQKgQhPZojzFFuQYSYYYh93TPQbqGDoz61W6XgF4D0dv4KWlpQZNEct87UKobSBzXi4eOdgsbHxujmfLkD_nR

buf, ret = client.Read(fd312, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd310, []byte("bKWTHYlPttkq801dLsMI4HyInrR2HWVB0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0
//fd state: (0) 5HLC9P5JC5m6KeeKjYOXLT0cuwVSHK9m9RZWiH_JbDTowo6mVpgu_ptMcqtpataAcS5PTFAoA2D3HLowYA4B2qOZTGZx1YG5qGZFuEE4KcRSiW2mrT2e4

ret = client.Write(fd318, []byte("FFbEigE1tfjVWcC7Z_AmJGUXv9S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) FFbEigE1tfjVWcC7Z_AmJGUXv9SSHK9m9RZWiH_JbDTowo6mVpgu_ptMcqtpataAcS5PTFAoA2D3HLowYA4B2qOZTGZx1YG5qGZFuEE4KcRSiW2mrT2e4

ret = client.Close(fd306)
if(ret != 0) {
  panic("close failed")
}


fd319 := client.Open("/rWIIG4WJd_.txt", client.O_RDWR|client.O_CREATE)
if(fd319 < 0) {
  panic("open failed")
}


ret = client.Close(fd319)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd317)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd297)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd290)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd315, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd314)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd301, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd320 := client.Open("/eXTFVE24CN.txt", client.O_RDWR|client.O_CREATE)
if(fd320 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd316, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd321 := client.Open("/raavVeJ6eA.txt", client.O_RDWR|client.O_CREATE)
if(fd321 < 0) {
  panic("open failed")
}


fd322 := client.Open("/Z2evREiOSO.txt", client.O_RDWR|client.O_CREATE)
if(fd322 < 0) {
  panic("open failed")
}


ret = client.Seek(fd312, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd293, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Ychgcqtd7OznfyVDk2vZAuS_FZr4ubP") {
  panic("wrong data returned")
}

//fd state: (106) d0DRyGvTe7Jr648NvHumNLQKgQhPZojzFFuQYSYYYh93TPQbqGDoz61W6XgF4D0dv4KWlpQZNEct87UKobSBzXi4eOdgsbHxujmfLkD_nR

ret = client.Write(fd308, []byte("mX9Po67SG7eW4xvcx_kszYFFCre3dQpv3QqhKl5YPrX_MywXc8Sm2cFeCd27fuDzD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (171) d0DRyGvTe7Jr648NvHumNLQKgQhPZojzFFuQYSYYYh93TPQbqGDoz61W6XgF4D0dv4KWlpQZNEct87UKobSBzXi4eOdgsbHxujmfLkD_nRmX9Po67SG7eW4xvcx_kszYFFCre3dQpv3QqhKl5YPrX_MywXc8Sm2cFeCd27fuDzD
//fd state: (0) 

ret = client.Write(fd315, []byte("0WfFhR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) 0WfFhR

buf, ret = client.Read(fd315, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd299)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 0WfFhR

ret = client.Write(fd307, []byte("ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_a5QuAFpyNJ8e4IsZKs3aEOtP1ZDRmX4L5veAPY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_a5QuAFpyNJ8e4IsZKs3aEOtP1ZDRmX4L5veAPY
//fd state: (0) 

ret = client.Write(fd309, []byte("56oa5_PWNhGj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) 56oa5_PWNhGj
//fd state: (79) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_a5QuAFpyNJ8e4IsZKs3aEOtP1ZDRmX4L5veAPY

ret = client.Write(fd307, []byte("tdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzdZKU_ygvfxbXMt7QubaeUtF2_V2pnR50DwZ5wQVJfVm59bMqR_wcSAH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (171) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_a5QuAFpyNJ8e4IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzdZKU_ygvfxbXMt7QubaeUtF2_V2pnR50DwZ5wQVJfVm59bMqR_wcSAH

ret = client.Close(fd315)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd293, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "l5I0rqK1jvThMKhcAnv6OrItYXdJKNHJosp7OZVU0mqQ4iuhaqehu1LB") {
  panic("wrong data returned")
}

//fd state: (12) 56oa5_PWNhGj

ret = client.Write(fd309, []byte("Pmk1FzPt3OUyW1AXllWUpg_D5ziodqkum8dpcBJJa9PJLlwV033SaNkYgAYpErfj74YfGDHcUni2ivZl5S8jKTGVhf44pN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (106) 56oa5_PWNhGjPmk1FzPt3OUyW1AXllWUpg_D5ziodqkum8dpcBJJa9PJLlwV033SaNkYgAYpErfj74YfGDHcUni2ivZl5S8jKTGVhf44pN

ret = client.Seek(fd309, 64, client.SEEK_SET)
if(ret != 64) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 64)
  panic("seek failed")
}


ret = client.Close(fd307)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd309)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd321, []byte("LsX2upRewHuGPKOI6otReaIgSHqD10MRV8hKSrxLqSQdNbnRugD0gFnKmqlDGiQqEIM1LvGx1lrD1W0YiEX6Haf0O"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) LsX2upRewHuGPKOI6otReaIgSHqD10MRV8hKSrxLqSQdNbnRugD0gFnKmqlDGiQqEIM1LvGx1lrD1W0YiEX6Haf0O

ret = client.Close(fd308)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd293, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MYDDkL9J32yXCvUKauzEEq") {
  panic("wrong data returned")
}


fd323 := client.Open("/62kKbTRIzd.txt", client.O_RDWR|client.O_CREATE)
if(fd323 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd316, []byte("Eyp9BWYKPl0D9bGPcgEEUDlvbkoF3j19lmL90s1rK5PSs5SxF4RtR44Nf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) Eyp9BWYKPl0D9bGPcgEEUDlvbkoF3j19lmL90s1rK5PSs5SxF4RtR44Nf

ret = client.Seek(fd293, 255, client.SEEK_SET)
if(ret != 255) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 255)
  panic("seek failed")
}


fd324 := client.Open("/R8FiGZHyRq.txt", client.O_RDWR|client.O_CREATE)
if(fd324 < 0) {
  panic("open failed")
}


ret = client.Close(fd316)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd312, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd312, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd312, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd321)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd318, 87, client.SEEK_SET)
if(ret != 87) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 87)
  panic("seek failed")
}


buf, ret = client.Read(fd310, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (87) FFbEigE1tfjVWcC7Z_AmJGUXv9SSHK9m9RZWiH_JbDTowo6mVpgu_ptMcqtpataAcS5PTFAoA2D3HLowYA4B2qOZTGZx1YG5qGZFuEE4KcRSiW2mrT2e4

ret = client.Write(fd318, []byte("uCZ0bph6UPqAP6GpcVPkuQsYLobQKZe4C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) FFbEigE1tfjVWcC7Z_AmJGUXv9SSHK9m9RZWiH_JbDTowo6mVpgu_ptMcqtpataAcS5PTFAoA2D3HLowYA4B2qOuCZ0bph6UPqAP6GpcVPkuQsYLobQKZe4C
//fd state: (0) 

ret = client.Write(fd323, []byte("pIdaBHc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) pIdaBHc
//fd state: (7) pIdaBHc

ret = client.Write(fd323, []byte("H35nXzQpPgQg6hOT3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) pIdaBHcH35nXzQpPgQg6hOT3
//fd state: (255) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7nTxz1cfPxp9OFonsu5Qsah3sYRNfUtGPqyjZlSpLMNZwX202pamULAI1fXpWH1XEKxhzO4xSEYBnWQFPkXYchgcqtd7OznfyVDk2vZAuS_FZr4ubPl5I0rqK1jvThMKhcAnv6OrItYXdJKNHJosp7OZVU0mqQ4iuhaqehu1LBMYDDkL9J32yXCvUKauzEEq

ret = client.Write(fd293, []byte("Svf51Zz4ofN6vTXP5_HE1Y0TxnfOdhlg8P6_05f2gA4lSaSumtNfgn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (309) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7nTxz1cfPxp9OFonsu5Qsah3sYRNfUtGPqyjZlSpLMNZwX202pamULAI1fXpWH1XEKxhzO4xSEYBnWQFPkXYchgcqtd7OznfyVDk2vZAuS_FZr4ubPl5I0rqK1jvThMKhcAnv6OrItYXdJKNHJosp7OZVU0mqQ4iuhaqehu1LBMYDDkLSvf51Zz4ofN6vTXP5_HE1Y0TxnfOdhlg8P6_05f2gA4lSaSumtNfgn
//fd state: (33) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0

ret = client.Write(fd310, []byte("Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILot"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILot
//fd state: (3) wqYNkMTfDmBxC3KBQ7wVlzECLvRsXYFf1RJS1_yKSELmkB4VT5P5rLSV22Yw5ML53z

ret = client.Write(fd301, []byte("VnAjP6G9KDsf9Q_dYKmZuYddqYXqkjJVPVyC8TpnCz8htDuMrodZSjqYf7AR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) wqYVnAjP6G9KDsf9Q_dYKmZuYddqYXqkjJVPVyC8TpnCz8htDuMrodZSjqYf7AR53z
//fd state: (73) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILot

ret = client.Write(fd310, []byte("ZOQ9DRW6aiwvb6PqZtXlvB9Qs5T67ZM2EtdTqPfz87iwb5nQ3alYUhb5H8FlClzkVKzDGANhi9LMfJJDJeZyPuej6ePC1dY7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (169) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILotZOQ9DRW6aiwvb6PqZtXlvB9Qs5T67ZM2EtdTqPfz87iwb5nQ3alYUhb5H8FlClzkVKzDGANhi9LMfJJDJeZyPuej6ePC1dY7

fd325 := client.Open("/VKesPBLV5v.txt", client.O_RDWR|client.O_CREATE)
if(fd325 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd312, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd324)
if(ret != 0) {
  panic("close failed")
}

//fd state: (309) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7nTxz1cfPxp9OFonsu5Qsah3sYRNfUtGPqyjZlSpLMNZwX202pamULAI1fXpWH1XEKxhzO4xSEYBnWQFPkXYchgcqtd7OznfyVDk2vZAuS_FZr4ubPl5I0rqK1jvThMKhcAnv6OrItYXdJKNHJosp7OZVU0mqQ4iuhaqehu1LBMYDDkLSvf51Zz4ofN6vTXP5_HE1Y0TxnfOdhlg8P6_05f2gA4lSaSumtNfgn

ret = client.Write(fd293, []byte("kh6gQdmGzWBE_hrtTi7Bj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (330) NkE1P8Jt_ZDnvx1Q6aR52IvYHbqUlAtyG4Mpcduq_Wt6MTbJj9EZfultn9M9Lht33Sg1Eua7iN9EDsZ7nTxz1cfPxp9OFonsu5Qsah3sYRNfUtGPqyjZlSpLMNZwX202pamULAI1fXpWH1XEKxhzO4xSEYBnWQFPkXYchgcqtd7OznfyVDk2vZAuS_FZr4ubPl5I0rqK1jvThMKhcAnv6OrItYXdJKNHJosp7OZVU0mqQ4iuhaqehu1LBMYDDkLSvf51Zz4ofN6vTXP5_HE1Y0TxnfOdhlg8P6_05f2gA4lSaSumtNfgnkh6gQdmGzWBE_hrtTi7Bj

ret = client.Close(fd310)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd325)
if(ret != 0) {
  panic("close failed")
}


fd326 := client.Open("/JR6nH_boRB.txt", client.O_RDWR|client.O_CREATE)
if(fd326 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd312, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd322)
if(ret != 0) {
  panic("close failed")
}


fd327 := client.Open("/PS_GRjWAR1.txt", client.O_RDWR|client.O_CREATE)
if(fd327 < 0) {
  panic("open failed")
}


ret = client.Close(fd320)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (0) 

ret = client.Write(fd327, []byte("mr6XPPo67LLCf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) mr6XPPo67LLCf

ret = client.Close(fd293)
if(ret != 0) {
  panic("close failed")
}


fd328 := client.Open("/VKesPBLV5v.txt", client.O_RDWR|client.O_CREATE)
if(fd328 < 0) {
  panic("open failed")
}


ret = client.Close(fd323)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd312, []byte("44vMiU1nR1byF3C2QujZaL80W65WvzxRZjJ0O2uNo7OYvW5LsWt0H_ws4jzqIgH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) 44vMiU1nR1byF3C2QujZaL80W65WvzxRZjJ0O2uNo7OYvW5LsWt0H_ws4jzqIgH

ret = client.Close(fd327)
if(ret != 0) {
  panic("close failed")
}

//fd state: (63) 44vMiU1nR1byF3C2QujZaL80W65WvzxRZjJ0O2uNo7OYvW5LsWt0H_ws4jzqIgH

ret = client.Write(fd312, []byte("ezZso4UNrpPFEVGS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) 44vMiU1nR1byF3C2QujZaL80W65WvzxRZjJ0O2uNo7OYvW5LsWt0H_ws4jzqIgHezZso4UNrpPFEVGS
//fd state: (0) 

ret = client.Write(fd328, []byte("gbspdWG5U_A3z549jZpgLXApGgQPyZh1TwkTThw7jPnrMtZnFvF8XJHwzN_i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) gbspdWG5U_A3z549jZpgLXApGgQPyZh1TwkTThw7jPnrMtZnFvF8XJHwzN_i

ret = client.Close(fd301)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd326, 77, client.SEEK_SET)
if(ret != 77) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 77)
  panic("seek failed")
}


ret = client.Close(fd312)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd326)
if(ret != 0) {
  panic("close failed")
}


fd329 := client.Open("/5C4wbbcHKO.txt", client.O_RDWR|client.O_CREATE)
if(fd329 < 0) {
  panic("open failed")
}


fd330 := client.Open("/rWIIG4WJd_.txt", client.O_RDWR|client.O_CREATE)
if(fd330 < 0) {
  panic("open failed")
}


ret = client.Close(fd328)
if(ret != 0) {
  panic("close failed")
}


fd331 := client.Open("/tbLiMuf1nh.txt", client.O_RDWR|client.O_CREATE)
if(fd331 < 0) {
  panic("open failed")
}


ret = client.Close(fd318)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd331)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd329)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd330)
if(ret != 0) {
  panic("close failed")
}


fd332 := client.Open("/6AwM3ujffY.txt", client.O_RDWR|client.O_CREATE)
if(fd332 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd332, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd332, []byte("uXMQlSo7IpgfF8lAwqt2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQbr3UwUGeUijW7DRhrsDrFZMxwW5tDv1uK4MT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) uXMQlSo7IpgfF8lAwqt2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQbr3UwUGeUijW7DRhrsDrFZMxwW5tDv1uK4MT

ret = client.Seek(fd332, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


fd333 := client.Open("/Phw5p58_St.txt", client.O_RDWR|client.O_CREATE)
if(fd333 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd333, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd332, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


buf, ret = client.Read(fd333, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd333, []byte("AfqVQfctptYND4yJoiJHGUMjP6H9KqQH_yS9oB_PlE3lxl3FT6qxeYgNI_dBYftHwONB9upkjc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) AfqVQfctptYND4yJoiJHGUMjP6H9KqQH_yS9oB_PlE3lxl3FT6qxeYgNI_dBYftHwONB9upkjc

buf, ret = client.Read(fd332, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQb") {
  panic("wrong data returned")
}


ret = client.Close(fd333)
if(ret != 0) {
  panic("close failed")
}

//fd state: (59) uXMQlSo7IpgfF8lAwqt2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQbr3UwUGeUijW7DRhrsDrFZMxwW5tDv1uK4MT

ret = client.Write(fd332, []byte("LssZIuboWvWIGzoYTfUkc82WQcmtPhdLnZAAEJDBg9vHvpSgM77g97drh27D0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) uXMQlSo7IpgfF8lAwqt2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQbLssZIuboWvWIGzoYTfUkc82WQcmtPhdLnZAAEJDBg9vHvpSgM77g97drh27D0
//fd state: (120) uXMQlSo7IpgfF8lAwqt2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQbLssZIuboWvWIGzoYTfUkc82WQcmtPhdLnZAAEJDBg9vHvpSgM77g97drh27D0

ret = client.Write(fd332, []byte("z7MkRGwz5Xd07fLXi7B3Anx_WUnJUlMfdfLL2ijvCzJ3JfQL_7kzOaoFIo5G7HA2jPM3GLs1b"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (193) uXMQlSo7IpgfF8lAwqt2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQbLssZIuboWvWIGzoYTfUkc82WQcmtPhdLnZAAEJDBg9vHvpSgM77g97drh27D0z7MkRGwz5Xd07fLXi7B3Anx_WUnJUlMfdfLL2ijvCzJ3JfQL_7kzOaoFIo5G7HA2jPM3GLs1b

ret = client.Seek(fd332, 168, client.SEEK_SET)
if(ret != 168) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 168)
  panic("seek failed")
}


ret = client.Seek(fd332, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


ret = client.Seek(fd332, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


buf, ret = client.Read(fd332, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "dd0mXGYTaTbaGVwQbLssZIuboWvWIGzoYTfUkc82W") {
  panic("wrong data returned")
}


ret = client.Seek(fd332, 149, client.SEEK_SET)
if(ret != 149) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 149)
  panic("seek failed")
}


ret = client.Seek(fd332, 131, client.SEEK_SET)
if(ret != 131) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 131)
  panic("seek failed")
}


buf, ret = client.Read(fd332, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "07fLXi7B3Anx_WUnJUlMfdfLL2ijvCzJ3JfQL_7kzOaoFIo5G7HA2jPM3GLs1b") {
  panic("wrong data returned")
}


fd334 := client.Open("/oo6gpjhf0k.txt", client.O_RDWR|client.O_CREATE)
if(fd334 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd332, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (193) uXMQlSo7IpgfF8lAwqt2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQbLssZIuboWvWIGzoYTfUkc82WQcmtPhdLnZAAEJDBg9vHvpSgM77g97drh27D0z7MkRGwz5Xd07fLXi7B3Anx_WUnJUlMfdfLL2ijvCzJ3JfQL_7kzOaoFIo5G7HA2jPM3GLs1b

ret = client.Write(fd332, []byte("MfFYxCU_n7GFSUJDAnOlyy8THPzLFMuoWjm0VrBgz52NUAeLGpeucT80X_nR4eeorGswdNJ4QESP1N61I6Cmq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (278) uXMQlSo7IpgfF8lAwqt2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQbLssZIuboWvWIGzoYTfUkc82WQcmtPhdLnZAAEJDBg9vHvpSgM77g97drh27D0z7MkRGwz5Xd07fLXi7B3Anx_WUnJUlMfdfLL2ijvCzJ3JfQL_7kzOaoFIo5G7HA2jPM3GLs1bMfFYxCU_n7GFSUJDAnOlyy8THPzLFMuoWjm0VrBgz52NUAeLGpeucT80X_nR4eeorGswdNJ4QESP1N61I6Cmq
//fd state: (0) 

ret = client.Write(fd334, []byte("N69mY4R8NWbX3qrbShAtsxrLCEqJ4NQ46F8ybPRb4XXZW67ynVGLkxtBqRUYEWhtHsHFpjXGzwF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) N69mY4R8NWbX3qrbShAtsxrLCEqJ4NQ46F8ybPRb4XXZW67ynVGLkxtBqRUYEWhtHsHFpjXGzwF

buf, ret = client.Read(fd334, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd334, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (278) uXMQlSo7IpgfF8lAwqt2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQbLssZIuboWvWIGzoYTfUkc82WQcmtPhdLnZAAEJDBg9vHvpSgM77g97drh27D0z7MkRGwz5Xd07fLXi7B3Anx_WUnJUlMfdfLL2ijvCzJ3JfQL_7kzOaoFIo5G7HA2jPM3GLs1bMfFYxCU_n7GFSUJDAnOlyy8THPzLFMuoWjm0VrBgz52NUAeLGpeucT80X_nR4eeorGswdNJ4QESP1N61I6Cmq

ret = client.Write(fd332, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (278) uXMQlSo7IpgfF8lAwqt2DJwMIdCQ58WEbuN1U31Vhfdd0mXGYTaTbaGVwQbLssZIuboWvWIGzoYTfUkc82WQcmtPhdLnZAAEJDBg9vHvpSgM77g97drh27D0z7MkRGwz5Xd07fLXi7B3Anx_WUnJUlMfdfLL2ijvCzJ3JfQL_7kzOaoFIo5G7HA2jPM3GLs1bMfFYxCU_n7GFSUJDAnOlyy8THPzLFMuoWjm0VrBgz52NUAeLGpeucT80X_nR4eeorGswdNJ4QESP1N61I6Cmq

ret = client.Seek(fd334, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


ret = client.Seek(fd334, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


fd335 := client.Open("/UkkCIjbjVT.txt", client.O_RDWR|client.O_CREATE)
if(fd335 < 0) {
  panic("open failed")
}


fd336 := client.Open("/ZkuXs_QYnQ.txt", client.O_RDWR|client.O_CREATE)
if(fd336 < 0) {
  panic("open failed")
}


ret = client.Close(fd335)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd334)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd332)
if(ret != 0) {
  panic("close failed")
}


fd337 := client.Open("/9q9ndfi0_e.txt", client.O_RDWR|client.O_CREATE)
if(fd337 < 0) {
  panic("open failed")
}


ret = client.Close(fd336)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd337, []byte("cPC46h5xLt23HA74ghk2iVoArCXcMm8fvNiqQe_bACxNyra6N6xYenomcVLFfJsaHu1b9qUzrUYiFIyH38r9zCrSRuk68lwJJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) cPC46h5xLt23HA74ghk2iVoArCXcMm8fvNiqQe_bACxNyra6N6xYenomcVLFfJsaHu1b9qUzrUYiFIyH38r9zCrSRuk68lwJJ

buf, ret = client.Read(fd337, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd337, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (97) cPC46h5xLt23HA74ghk2iVoArCXcMm8fvNiqQe_bACxNyra6N6xYenomcVLFfJsaHu1b9qUzrUYiFIyH38r9zCrSRuk68lwJJ

ret = client.Write(fd337, []byte("cZrhvqnTF4vLYspQpte58bTTbkHUWFoFu7UCLhRn6QQa0XX6pu0EVsPLUCVL9drmg9Fk0JmeACzW6WtcEOMcd6D402ufq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (190) cPC46h5xLt23HA74ghk2iVoArCXcMm8fvNiqQe_bACxNyra6N6xYenomcVLFfJsaHu1b9qUzrUYiFIyH38r9zCrSRuk68lwJJcZrhvqnTF4vLYspQpte58bTTbkHUWFoFu7UCLhRn6QQa0XX6pu0EVsPLUCVL9drmg9Fk0JmeACzW6WtcEOMcd6D402ufq

fd338 := client.Open("/PecpTE5aC2.txt", client.O_RDWR|client.O_CREATE)
if(fd338 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd338, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd339 := client.Open("/CGxX_rlbuS.txt", client.O_RDWR|client.O_CREATE)
if(fd339 < 0) {
  panic("open failed")
}


ret = client.Close(fd338)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd337, 148, client.SEEK_SET)
if(ret != 148) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 148)
  panic("seek failed")
}


buf, ret = client.Read(fd339, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd339, []byte("dmRkrLSkSS3R8OlgEVggvo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) dmRkrLSkSS3R8OlgEVggvo

ret = client.Close(fd339)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd337, 92, client.SEEK_SET)
if(ret != 92) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 92)
  panic("seek failed")
}


ret = client.Close(fd337)
if(ret != 0) {
  panic("close failed")
}


fd340 := client.Open("/llTndggwBW.txt", client.O_RDWR|client.O_CREATE)
if(fd340 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd340, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd341 := client.Open("/eCvA_JYFK6.txt", client.O_RDWR|client.O_CREATE)
if(fd341 < 0) {
  panic("open failed")
}


ret = client.Seek(fd341, 106, client.SEEK_SET)
if(ret != 106) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 106)
  panic("seek failed")
}


ret = client.Seek(fd340, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd341, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "58knH6d4CzdZKU_ygvfxbX") {
  panic("wrong data returned")
}


fd342 := client.Open("/xpQvmuTyo2.txt", client.O_RDWR|client.O_CREATE)
if(fd342 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd342, []byte("majyE2L13bVVck0nT8gY_InwWKbWtbZeyyW1Nrs3J6LJvR2CjYtPS1xM5qEbnM9RZ20QyeDeD6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) majyE2L13bVVck0nT8gY_InwWKbWtbZeyyW1Nrs3J6LJvR2CjYtPS1xM5qEbnM9RZ20QyeDeD6
//fd state: (128) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_a5QuAFpyNJ8e4IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzdZKU_ygvfxbXMt7QubaeUtF2_V2pnR50DwZ5wQVJfVm59bMqR_wcSAH

ret = client.Write(fd341, []byte("F3wLNwbzM7ZzJLDoNLle2tL2tM86e7OkaOYrsHizgJrb6V9MuCjeLHc7SPm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (187) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_a5QuAFpyNJ8e4IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzdZKU_ygvfxbXF3wLNwbzM7ZzJLDoNLle2tL2tM86e7OkaOYrsHizgJrb6V9MuCjeLHc7SPm

buf, ret = client.Read(fd340, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (74) majyE2L13bVVck0nT8gY_InwWKbWtbZeyyW1Nrs3J6LJvR2CjYtPS1xM5qEbnM9RZ20QyeDeD6

ret = client.Write(fd342, []byte("y07aYELz72d55qe_3W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) majyE2L13bVVck0nT8gY_InwWKbWtbZeyyW1Nrs3J6LJvR2CjYtPS1xM5qEbnM9RZ20QyeDeD6y07aYELz72d55qe_3W

ret = client.Seek(fd340, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd343 := client.Open("/Z96QJ6_Xyc.txt", client.O_RDWR|client.O_CREATE)
if(fd343 < 0) {
  panic("open failed")
}


fd344 := client.Open("/IwCBPIfR8t.txt", client.O_RDWR|client.O_CREATE)
if(fd344 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd340, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd343, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd341, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


fd345 := client.Open("/CR4qfzt_Wc.txt", client.O_RDWR|client.O_CREATE)
if(fd345 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd340, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd344, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JKvPnuLf0QPpRBKzGGz2Qp1rjJbmPe7KOO1") {
  panic("wrong data returned")
}


ret = client.Seek(fd342, 85, client.SEEK_SET)
if(ret != 85) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 85)
  panic("seek failed")
}


buf, ret = client.Read(fd341, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "00MVkcDMpmJ2_a5QuAFpyNJ8e4IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd343, []byte("6FY5Txn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) 6FY5Txn

buf, ret = client.Read(fd341, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4Cz") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd343, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd344, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}

//fd state: (116) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_a5QuAFpyNJ8e4IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzdZKU_ygvfxbXF3wLNwbzM7ZzJLDoNLle2tL2tM86e7OkaOYrsHizgJrb6V9MuCjeLHc7SPm

ret = client.Write(fd341, []byte("xFwV3XfG3iZvHvuaiXbwMl1GZgTw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_a5QuAFpyNJ8e4IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzxFwV3XfG3iZvHvuaiXbwMl1GZgTwNLle2tL2tM86e7OkaOYrsHizgJrb6V9MuCjeLHc7SPm

ret = client.Seek(fd342, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


buf, ret = client.Read(fd340, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd342)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd340)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd343, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd345, []byte("cbrWhsPG1xyWR1F7JVvML2HVNjnDiF1oKSB3H3cx8zmSYE46FtpOIGl0MBh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) cbrWhsPG1xyWR1F7JVvML2HVNjnDiF1oKSB3H3cx8zmSYE46FtpOIGl0MBh

ret = client.Seek(fd341, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}

//fd state: (42) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_a5QuAFpyNJ8e4IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzxFwV3XfG3iZvHvuaiXbwMl1GZgTwNLle2tL2tM86e7OkaOYrsHizgJrb6V9MuCjeLHc7SPm

ret = client.Write(fd341, []byte("ZPos1xdEoVh7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_aZPos1xdEoVh7IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzxFwV3XfG3iZvHvuaiXbwMl1GZgTwNLle2tL2tM86e7OkaOYrsHizgJrb6V9MuCjeLHc7SPm

ret = client.Seek(fd344, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Seek(fd341, 166, client.SEEK_SET)
if(ret != 166) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 166)
  panic("seek failed")
}


ret = client.Close(fd345)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd343, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "6FY5") {
  panic("wrong data returned")
}

//fd state: (166) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_aZPos1xdEoVh7IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzxFwV3XfG3iZvHvuaiXbwMl1GZgTwNLle2tL2tM86e7OkaOYrsHizgJrb6V9MuCjeLHc7SPm

ret = client.Write(fd341, []byte("PWq_rZ1ZJQ6CjxL8JenrWB5S13GRsLn8Qb8qQdLE5N"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (208) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_aZPos1xdEoVh7IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzxFwV3XfG3iZvHvuaiXbwMl1GZgTwNLle2tL2tM86e7OkaOYrsHPWq_rZ1ZJQ6CjxL8JenrWB5S13GRsLn8Qb8qQdLE5N
//fd state: (4) 6FY5Txn

ret = client.Write(fd343, []byte("wFg0S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) 6FY5wFg0S

ret = client.Seek(fd341, 89, client.SEEK_SET)
if(ret != 89) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 89)
  panic("seek failed")
}


fd346 := client.Open("/GI8sQv42jD.txt", client.O_RDWR|client.O_CREATE)
if(fd346 < 0) {
  panic("open failed")
}


ret = client.Seek(fd344, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd341, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "RJGb9jOtesMDBMAGo58knH6d4CzxFwV3XfG3iZvHvuaiXbwMl1GZgTwNLle2tL2tM86e7OkaOYrsHPWq_rZ1") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd341, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ZJQ6CjxL8JenrWB5S13GRsLn8Qb8qQdLE5N") {
  panic("wrong data returned")
}


ret = client.Close(fd343)
if(ret != 0) {
  panic("close failed")
}


fd347 := client.Open("/ZLjr5ROzuJ.txt", client.O_RDWR|client.O_CREATE)
if(fd347 < 0) {
  panic("open failed")
}


fd348 := client.Open("/DRtZWfLYIa.txt", client.O_RDWR|client.O_CREATE)
if(fd348 < 0) {
  panic("open failed")
}


ret = client.Close(fd346)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd344)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd348, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd348, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd347, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd347)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd348, []byte("RetaPJTRxfhVRgnOq3Kurgt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) RetaPJTRxfhVRgnOq3Kurgt

ret = client.Seek(fd348, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}

//fd state: (22) RetaPJTRxfhVRgnOq3Kurgt

ret = client.Write(fd348, []byte("kWQa_gREnI_BUvUrvh5Q1zFLqAFYwlvQ9vnghmN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) RetaPJTRxfhVRgnOq3KurgkWQa_gREnI_BUvUrvh5Q1zFLqAFYwlvQ9vnghmN

buf, ret = client.Read(fd341, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (61) RetaPJTRxfhVRgnOq3KurgkWQa_gREnI_BUvUrvh5Q1zFLqAFYwlvQ9vnghmN

ret = client.Write(fd348, []byte("UJibCbN7NdmJHccFnqlbTyLD0x5UQjpT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) RetaPJTRxfhVRgnOq3KurgkWQa_gREnI_BUvUrvh5Q1zFLqAFYwlvQ9vnghmNUJibCbN7NdmJHccFnqlbTyLD0x5UQjpT

ret = client.Close(fd348)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd341, 166, client.SEEK_SET)
if(ret != 166) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 166)
  panic("seek failed")
}


fd349 := client.Open("/H8q3MipPbg.txt", client.O_RDWR|client.O_CREATE)
if(fd349 < 0) {
  panic("open failed")
}


ret = client.Close(fd349)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd341, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PWq_rZ1ZJQ6CjxL8JenrWB5S13GRsLn8Qb8qQdLE5N") {
  panic("wrong data returned")
}


ret = client.Close(fd341)
if(ret != 0) {
  panic("close failed")
}


fd350 := client.Open("/Z34Wgzj0C3.txt", client.O_RDWR|client.O_CREATE)
if(fd350 < 0) {
  panic("open failed")
}


ret = client.Close(fd350)
if(ret != 0) {
  panic("close failed")
}


fd351 := client.Open("/VTasqgKgMi.txt", client.O_RDWR|client.O_CREATE)
if(fd351 < 0) {
  panic("open failed")
}


ret = client.Close(fd351)
if(ret != 0) {
  panic("close failed")
}


fd352 := client.Open("/FIXnoumSKW.txt", client.O_RDWR|client.O_CREATE)
if(fd352 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd352, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd352, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd352, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd352, []byte("85_CH07eEYLGrLTfMdeNwFk_ziMARs8aNR4vAo2Ys3GbdLDx5chKgaPGbod4dG_QZ1i8Tz9522wV8UmPK6LkvJCwyQoNWuMbac"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) 85_CH07eEYLGrLTfMdeNwFk_ziMARs8aNR4vAo2Ys3GbdLDx5chKgaPGbod4dG_QZ1i8Tz9522wV8UmPK6LkvJCwyQoNWuMbac

fd353 := client.Open("/rkU9KrdLrD.txt", client.O_RDWR|client.O_CREATE)
if(fd353 < 0) {
  panic("open failed")
}

//fd state: (98) 85_CH07eEYLGrLTfMdeNwFk_ziMARs8aNR4vAo2Ys3GbdLDx5chKgaPGbod4dG_QZ1i8Tz9522wV8UmPK6LkvJCwyQoNWuMbac

ret = client.Write(fd352, []byte("Yo0rjTZRxAZSwj259I0Rf17iape1agxvha_s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) 85_CH07eEYLGrLTfMdeNwFk_ziMARs8aNR4vAo2Ys3GbdLDx5chKgaPGbod4dG_QZ1i8Tz9522wV8UmPK6LkvJCwyQoNWuMbacYo0rjTZRxAZSwj259I0Rf17iape1agxvha_s

buf, ret = client.Read(fd353, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd354 := client.Open("/dg_fg5cemT.txt", client.O_RDWR|client.O_CREATE)
if(fd354 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd353, []byte("djyGYSLV3sDgD8t3ga95c7U0b5mwKDHY3W1Ct0ntvw80ezIMgI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) djyGYSLV3sDgD8t3ga95c7U0b5mwKDHY3W1Ct0ntvw80ezIMgI

ret = client.Seek(fd352, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Close(fd353)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd354, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd354, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd354, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd355 := client.Open("/iwLLSjsyMF.txt", client.O_RDWR|client.O_CREATE)
if(fd355 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd354, []byte("6thLrFYNOP67TYEZjX22pZU17rdPOBAriwG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) 6thLrFYNOP67TYEZjX22pZU17rdPOBAriwG

buf, ret = client.Read(fd352, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "aNR4vAo2Ys3GbdLDx5chKgaPGbod4dG_QZ1i8Tz9522wV8UmPK6LkvJCwyQoNWuMbacYo0rjTZR") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd354, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd355, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd356 := client.Open("/xtYmEEvAxH.txt", client.O_RDWR|client.O_CREATE)
if(fd356 < 0) {
  panic("open failed")
}


ret = client.Close(fd356)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd355, []byte("0QNjxbktDhb3bLTMxTnxflCEHaM1aTIMrpVFenba5MHz0aR080uo7oxXE5EVI82mtXwk6fHLhV6K7marRaWJd9p48HgZe7qo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) 0QNjxbktDhb3bLTMxTnxflCEHaM1aTIMrpVFenba5MHz0aR080uo7oxXE5EVI82mtXwk6fHLhV6K7marRaWJd9p48HgZe7qo
//fd state: (106) 85_CH07eEYLGrLTfMdeNwFk_ziMARs8aNR4vAo2Ys3GbdLDx5chKgaPGbod4dG_QZ1i8Tz9522wV8UmPK6LkvJCwyQoNWuMbacYo0rjTZRxAZSwj259I0Rf17iape1agxvha_s

ret = client.Write(fd352, []byte("qKgAsyknDWxJI_eSsBsCJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (127) 85_CH07eEYLGrLTfMdeNwFk_ziMARs8aNR4vAo2Ys3GbdLDx5chKgaPGbod4dG_QZ1i8Tz9522wV8UmPK6LkvJCwyQoNWuMbacYo0rjTZRqKgAsyknDWxJI_eSsBsCJgxvha_s

buf, ret = client.Read(fd352, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gxvha_s") {
  panic("wrong data returned")
}


fd357 := client.Open("/rkkje4LeOI.txt", client.O_RDWR|client.O_CREATE)
if(fd357 < 0) {
  panic("open failed")
}


ret = client.Close(fd355)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd352, 113, client.SEEK_SET)
if(ret != 113) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 113)
  panic("seek failed")
}


fd358 := client.Open("/O9U7jQNIkQ.txt", client.O_RDWR|client.O_CREATE)
if(fd358 < 0) {
  panic("open failed")
}


ret = client.Seek(fd352, 104, client.SEEK_SET)
if(ret != 104) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 104)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd358, []byte("BajfYiIvHz5VB5vieLHoEzm5tPjXqDQgBA3srmRXwGRGkuHCNnOekCxp4DnCCwtRRytxeUGjGTvU8Z283kX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) BajfYiIvHz5VB5vieLHoEzm5tPjXqDQgBA3srmRXwGRGkuHCNnOekCxp4DnCCwtRRytxeUGjGTvU8Z283kX

ret = client.Close(fd358)
if(ret != 0) {
  panic("close failed")
}

//fd state: (35) 6thLrFYNOP67TYEZjX22pZU17rdPOBAriwG

ret = client.Write(fd354, []byte("dUpplR9OtLPDYofsezxEqi2tu7PqKPO7AP6yjmAcXaZo1QGHv83OZWYngCvdVP59edaXcsGsWx8jfl0dGtUk_YeiXvuxaHaH99"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) 6thLrFYNOP67TYEZjX22pZU17rdPOBAriwGdUpplR9OtLPDYofsezxEqi2tu7PqKPO7AP6yjmAcXaZo1QGHv83OZWYngCvdVP59edaXcsGsWx8jfl0dGtUk_YeiXvuxaHaH99

ret = client.Seek(fd354, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


fd359 := client.Open("/1db9qrUlfV.txt", client.O_RDWR|client.O_CREATE)
if(fd359 < 0) {
  panic("open failed")
}

//fd state: (104) 85_CH07eEYLGrLTfMdeNwFk_ziMARs8aNR4vAo2Ys3GbdLDx5chKgaPGbod4dG_QZ1i8Tz9522wV8UmPK6LkvJCwyQoNWuMbacYo0rjTZRqKgAsyknDWxJI_eSsBsCJgxvha_s

ret = client.Write(fd352, []byte("DJdC10GQMRCj_R04pzjFPzsekH5ExD4PpLdRuO0hrC8Yomg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) 85_CH07eEYLGrLTfMdeNwFk_ziMARs8aNR4vAo2Ys3GbdLDx5chKgaPGbod4dG_QZ1i8Tz9522wV8UmPK6LkvJCwyQoNWuMbacYo0rjTDJdC10GQMRCj_R04pzjFPzsekH5ExD4PpLdRuO0hrC8Yomg

ret = client.Close(fd357)
if(ret != 0) {
  panic("close failed")
}


fd360 := client.Open("/fe6x_kC5oC.txt", client.O_RDWR|client.O_CREATE)
if(fd360 < 0) {
  panic("open failed")
}


fd361 := client.Open("/3djEq6AYXg.txt", client.O_RDWR|client.O_CREATE)
if(fd361 < 0) {
  panic("open failed")
}


ret = client.Close(fd354)
if(ret != 0) {
  panic("close failed")
}


fd362 := client.Open("/6xOy8mDYIZ.txt", client.O_RDWR|client.O_CREATE)
if(fd362 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd361, []byte("xe3HRD4_dauG7Tz5irlDpRk49VGPlcyKoaxLzqrK2yKMtp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) xe3HRD4_dauG7Tz5irlDpRk49VGPlcyKoaxLzqrK2yKMtp

ret = client.Close(fd359)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd352, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


ret = client.Close(fd352)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd360, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd363 := client.Open("/ySY1MqWeNO.txt", client.O_RDWR|client.O_CREATE)
if(fd363 < 0) {
  panic("open failed")
}


fd364 := client.Open("/xiflUJdfBL.txt", client.O_RDWR|client.O_CREATE)
if(fd364 < 0) {
  panic("open failed")
}


fd365 := client.Open("/0chQNY9PXi.txt", client.O_RDWR|client.O_CREATE)
if(fd365 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd363, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd362, []byte("fUI4bIIpLGCKTznRVb1YiZnvFant3dOhdvELISoUDUjf6PyJ8Fz7KzjvM6Nx0scVFBwxkt2lGMbt23aaHAE6cjqRl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) fUI4bIIpLGCKTznRVb1YiZnvFant3dOhdvELISoUDUjf6PyJ8Fz7KzjvM6Nx0scVFBwxkt2lGMbt23aaHAE6cjqRl

fd366 := client.Open("/nOo_rJHfBF.txt", client.O_RDWR|client.O_CREATE)
if(fd366 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd366, []byte("F8G2ZxsGcNvAcXPIUYq_DIM4xd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) F8G2ZxsGcNvAcXPIUYq_DIM4xd

buf, ret = client.Read(fd363, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd363, []byte("CN07Gp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) CN07Gp

ret = client.Seek(fd361, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


fd367 := client.Open("/YJ0bjUslc0.txt", client.O_RDWR|client.O_CREATE)
if(fd367 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd360, []byte("RIJYaXkdTlPr4CjImsxIdsmbXz3vhMXMPyDePRnF2fRdfASQpqsULnyU4G87K7M2G"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) RIJYaXkdTlPr4CjImsxIdsmbXz3vhMXMPyDePRnF2fRdfASQpqsULnyU4G87K7M2G

fd368 := client.Open("/eCvA_JYFK6.txt", client.O_RDWR|client.O_CREATE)
if(fd368 < 0) {
  panic("open failed")
}


ret = client.Seek(fd368, 159, client.SEEK_SET)
if(ret != 159) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 159)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd365, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd366, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Seek(fd365, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd364)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd365, []byte("odSH7nuqvS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) odSH7nuqvS

ret = client.Seek(fd361, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}

//fd state: (159) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_aZPos1xdEoVh7IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzxFwV3XfG3iZvHvuaiXbwMl1GZgTwNLle2tL2tM86e7OkaOYrsHPWq_rZ1ZJQ6CjxL8JenrWB5S13GRsLn8Qb8qQdLE5N

ret = client.Write(fd368, []byte("B4HkdILhC5yKOzWMcpOPkwMAnJEsbJ2AYh6MqtVvqmK0L7sU2UhIfMP2jo_wUkR5eZNNa0vzJiu5Yk09RKQMmBpVvf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (249) ztcZa_ZQtiGFyF_bDUD5APuEzf4t00MVkcDMpmJ2_aZPos1xdEoVh7IsZKs3aEOtP1ZDRmX4L5veAPYtdP9seHqbKRJGb9jOtesMDBMAGo58knH6d4CzxFwV3XfG3iZvHvuaiXbwMl1GZgTwNLle2tL2tM86e7OB4HkdILhC5yKOzWMcpOPkwMAnJEsbJ2AYh6MqtVvqmK0L7sU2UhIfMP2jo_wUkR5eZNNa0vzJiu5Yk09RKQMmBpVvf

ret = client.Close(fd368)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd366)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd367, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd367, []byte("QteKtL9lfbDCUGdadX7xngI2olRLoi6IiaQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) QteKtL9lfbDCUGdadX7xngI2olRLoi6IiaQ

buf, ret = client.Read(fd360, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (38) xe3HRD4_dauG7Tz5irlDpRk49VGPlcyKoaxLzqrK2yKMtp

ret = client.Write(fd361, []byte("JvxFUetsXcy6dV58i7QWT2YLqeDemq8usx2iFFZh1pbjtSka6lYrPTPTFMrBAMefb8PygNeziHOUDziqhOg7jee4TNssB4WU2ue"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (137) xe3HRD4_dauG7Tz5irlDpRk49VGPlcyKoaxLzqJvxFUetsXcy6dV58i7QWT2YLqeDemq8usx2iFFZh1pbjtSka6lYrPTPTFMrBAMefb8PygNeziHOUDziqhOg7jee4TNssB4WU2ue

ret = client.Close(fd367)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd362)
if(ret != 0) {
  panic("close failed")
}


fd369 := client.Open("/NzPRIhLV5e.txt", client.O_RDWR|client.O_CREATE)
if(fd369 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd369, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd369, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd365)
if(ret != 0) {
  panic("close failed")
}


fd370 := client.Open("/yY3JR0W531.txt", client.O_RDWR|client.O_CREATE)
if(fd370 < 0) {
  panic("open failed")
}

//fd state: (137) xe3HRD4_dauG7Tz5irlDpRk49VGPlcyKoaxLzqJvxFUetsXcy6dV58i7QWT2YLqeDemq8usx2iFFZh1pbjtSka6lYrPTPTFMrBAMefb8PygNeziHOUDziqhOg7jee4TNssB4WU2ue

ret = client.Write(fd361, []byte("uTUAcKdFUtSxhVd4nosbRuhrqm_l2LcpoQZOGVzhEZCzQSG5n0iKsH3drRpU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (197) xe3HRD4_dauG7Tz5irlDpRk49VGPlcyKoaxLzqJvxFUetsXcy6dV58i7QWT2YLqeDemq8usx2iFFZh1pbjtSka6lYrPTPTFMrBAMefb8PygNeziHOUDziqhOg7jee4TNssB4WU2ueuTUAcKdFUtSxhVd4nosbRuhrqm_l2LcpoQZOGVzhEZCzQSG5n0iKsH3drRpU
//fd state: (0) 

ret = client.Write(fd369, []byte("Bshh7U"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) Bshh7U

ret = client.Seek(fd370, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd360)
if(ret != 0) {
  panic("close failed")
}

//fd state: (6) Bshh7U

ret = client.Write(fd369, []byte("Tzbh5Q5HwIpPsZkULozyNf8JkUJ7RkEseL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) Bshh7UTzbh5Q5HwIpPsZkULozyNf8JkUJ7RkEseL

ret = client.Seek(fd370, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd363, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


buf, ret = client.Read(fd369, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd370, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd361, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd369, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd361)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd363)
if(ret != 0) {
  panic("close failed")
}


fd371 := client.Open("/0prf8WI95c.txt", client.O_RDWR|client.O_CREATE)
if(fd371 < 0) {
  panic("open failed")
}


fd372 := client.Open("/FMtf38GO1K.txt", client.O_RDWR|client.O_CREATE)
if(fd372 < 0) {
  panic("open failed")
}


fd373 := client.Open("/ti7WCCrD9X.txt", client.O_RDWR|client.O_CREATE)
if(fd373 < 0) {
  panic("open failed")
}


ret = client.Close(fd369)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) O_BkGN7xo7x885c95fxznNwGWIWrELz2WUKVvTF

ret = client.Write(fd371, []byte("jypep8ehQwLUOlsnfDI8dE4qfT8m76YRXfO3SBqhcq29Ga7nYBdNZOwOqsBYetnLP5GIk6l0x3gvQ9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) jypep8ehQwLUOlsnfDI8dE4qfT8m76YRXfO3SBqhcq29Ga7nYBdNZOwOqsBYetnLP5GIk6l0x3gvQ9

buf, ret = client.Read(fd373, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gphVZsJMAS3ncKU0btXjxTI8mnv4jCGCc3o5e9IlkAy0") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd372, []byte("zRSZWCY92heUeswKgy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) zRSZWCY92heUeswKgy

ret = client.Close(fd370)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd371, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd371, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd372)
if(ret != 0) {
  panic("close failed")
}


fd374 := client.Open("/MLw9D0Mx68.txt", client.O_RDWR|client.O_CREATE)
if(fd374 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd374, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd373)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd374, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd371)
if(ret != 0) {
  panic("close failed")
}


fd375 := client.Open("/NzPRIhLV5e.txt", client.O_RDWR|client.O_CREATE)
if(fd375 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd374, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) Bshh7UTzbh5Q5HwIpPsZkULozyNf8JkUJ7RkEseL

ret = client.Write(fd375, []byte("GKOVa_cmtn6lveIsfkYk4RUM3l5u0eq3TOo6rG4C6xznN_wCkYIrVDP0cKEDOQYw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) GKOVa_cmtn6lveIsfkYk4RUM3l5u0eq3TOo6rG4C6xznN_wCkYIrVDP0cKEDOQYw

ret = client.Close(fd374)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd375, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd376 := client.Open("/1wZUGH1ts_.txt", client.O_RDWR|client.O_CREATE)
if(fd376 < 0) {
  panic("open failed")
}

//fd state: (64) GKOVa_cmtn6lveIsfkYk4RUM3l5u0eq3TOo6rG4C6xznN_wCkYIrVDP0cKEDOQYw

ret = client.Write(fd375, []byte("xURr4MS2r"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) GKOVa_cmtn6lveIsfkYk4RUM3l5u0eq3TOo6rG4C6xznN_wCkYIrVDP0cKEDOQYwxURr4MS2r

ret = client.Close(fd375)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd376)
if(ret != 0) {
  panic("close failed")
}


fd377 := client.Open("/4nOzh3S2jn.txt", client.O_RDWR|client.O_CREATE)
if(fd377 < 0) {
  panic("open failed")
}


ret = client.Close(fd377)
if(ret != 0) {
  panic("close failed")
}


fd378 := client.Open("/vA8vL6aicz.txt", client.O_RDWR|client.O_CREATE)
if(fd378 < 0) {
  panic("open failed")
}


ret = client.Seek(fd378, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


ret = client.Seek(fd378, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


buf, ret = client.Read(fd378, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WTHYlPttkq801dLsMI4HyInrR2HWVB0Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILotZOQ9DRW6aiwvb6PqZtXlvB9Qs5T6") {
  panic("wrong data returned")
}

//fd state: (101) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILotZOQ9DRW6aiwvb6PqZtXlvB9Qs5T67ZM2EtdTqPfz87iwb5nQ3alYUhb5H8FlClzkVKzDGANhi9LMfJJDJeZyPuej6ePC1dY7

ret = client.Write(fd378, []byte("TcB_dsXI0nxDYKXvKQE1Gf4bgSC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (128) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILotZOQ9DRW6aiwvb6PqZtXlvB9Qs5T6TcB_dsXI0nxDYKXvKQE1Gf4bgSC5H8FlClzkVKzDGANhi9LMfJJDJeZyPuej6ePC1dY7
//fd state: (128) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILotZOQ9DRW6aiwvb6PqZtXlvB9Qs5T6TcB_dsXI0nxDYKXvKQE1Gf4bgSC5H8FlClzkVKzDGANhi9LMfJJDJeZyPuej6ePC1dY7

ret = client.Write(fd378, []byte("uIgkJ4JRm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (137) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILotZOQ9DRW6aiwvb6PqZtXlvB9Qs5T6TcB_dsXI0nxDYKXvKQE1Gf4bgSCuIgkJ4JRmVKzDGANhi9LMfJJDJeZyPuej6ePC1dY7
//fd state: (137) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILotZOQ9DRW6aiwvb6PqZtXlvB9Qs5T6TcB_dsXI0nxDYKXvKQE1Gf4bgSCuIgkJ4JRmVKzDGANhi9LMfJJDJeZyPuej6ePC1dY7

ret = client.Write(fd378, []byte("REx8K0oiiUVj1UU4CeCxpT9g1y6QaqZZl4d6pnip"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (177) bKWTHYlPttkq801dLsMI4HyInrR2HWVB0Wsfuq_eHtTxfEJkHiFuptC7oSSlHV4ocwwDtILotZOQ9DRW6aiwvb6PqZtXlvB9Qs5T6TcB_dsXI0nxDYKXvKQE1Gf4bgSCuIgkJ4JRmREx8K0oiiUVj1UU4CeCxpT9g1y6QaqZZl4d6pnip

fd379 := client.Open("/cDsrk7sXrQ.txt", client.O_RDWR|client.O_CREATE)
if(fd379 < 0) {
  panic("open failed")
}


ret = client.Close(fd378)
if(ret != 0) {
  panic("close failed")
}


fd380 := client.Open("/9yJ4B8r9Jr.txt", client.O_RDWR|client.O_CREATE)
if(fd380 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd380, []byte("W8H61C5wBk4OZBhwIjpsotZnsc5DKCeC509AQLRPmHUX4x_p0bl6fiCNQmKYRItTr4QjwBhfci8IndXWvja0s_qiQ_A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) W8H61C5wBk4OZBhwIjpsotZnsc5DKCeC509AQLRPmHUX4x_p0bl6fiCNQmKYRItTr4QjwBhfci8IndXWvja0s_qiQ_A

ret = client.Close(fd379)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd380)
if(ret != 0) {
  panic("close failed")
}


fd381 := client.Open("/pZv_3JdBVv.txt", client.O_RDWR|client.O_CREATE)
if(fd381 < 0) {
  panic("open failed")
}


ret = client.Close(fd381)
if(ret != 0) {
  panic("close failed")
}


fd382 := client.Open("/74mEwYxfha.txt", client.O_RDWR|client.O_CREATE)
if(fd382 < 0) {
  panic("open failed")
}

//fd state: (0) eoeeB8qZkgJtN8Lf2LOqCAeVx_0VkFGbZ98wXJJu

ret = client.Write(fd382, []byte("HvRa2ENPhB3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) HvRa2ENPhB3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaA

buf, ret = client.Read(fd382, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (44) HvRa2ENPhB3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaA

ret = client.Write(fd382, []byte("YTtNueQjzs9GvnM146GOEVteJffSDvKyMaSfioXjUPfo4iz7aakQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) HvRa2ENPhB3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaAYTtNueQjzs9GvnM146GOEVteJffSDvKyMaSfioXjUPfo4iz7aakQ

buf, ret = client.Read(fd382, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (96) HvRa2ENPhB3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaAYTtNueQjzs9GvnM146GOEVteJffSDvKyMaSfioXjUPfo4iz7aakQ

ret = client.Write(fd382, []byte("4IenrGLNZlPnIodCODHE4bLw7m3qDbrtGk9aZlLBM2SE0MHFG1RKPMI0KSJyFlJ0HFdT7tWSS_dvYgD4qf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (178) HvRa2ENPhB3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaAYTtNueQjzs9GvnM146GOEVteJffSDvKyMaSfioXjUPfo4iz7aakQ4IenrGLNZlPnIodCODHE4bLw7m3qDbrtGk9aZlLBM2SE0MHFG1RKPMI0KSJyFlJ0HFdT7tWSS_dvYgD4qf

ret = client.Seek(fd382, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


buf, ret = client.Read(fd382, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "aAYTtNueQjzs9GvnM146GOEVteJffS") {
  panic("wrong data returned")
}


ret = client.Close(fd382)
if(ret != 0) {
  panic("close failed")
}


fd383 := client.Open("/ylHMXBDQc2.txt", client.O_RDWR|client.O_CREATE)
if(fd383 < 0) {
  panic("open failed")
}


ret = client.Close(fd383)
if(ret != 0) {
  panic("close failed")
}


fd384 := client.Open("/YbYhTNaGQb.txt", client.O_RDWR|client.O_CREATE)
if(fd384 < 0) {
  panic("open failed")
}

//fd state: (0) Oq_Pzj8StdxDIM7c0HWFZl8ZpFHO03TRVEahm2r5s7wBfjV7FHFpM6miij1Dc6WC4jtcA3b9vUK5AlyPOqWRJzpzApmWGlMNUd4nZQhqAdftp_UeK

ret = client.Write(fd384, []byte("aITibAqnyiZ7bT79sgJBKFKVPFc3A8MQb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) aITibAqnyiZ7bT79sgJBKFKVPFc3A8MQbEahm2r5s7wBfjV7FHFpM6miij1Dc6WC4jtcA3b9vUK5AlyPOqWRJzpzApmWGlMNUd4nZQhqAdftp_UeK
//fd state: (33) aITibAqnyiZ7bT79sgJBKFKVPFc3A8MQbEahm2r5s7wBfjV7FHFpM6miij1Dc6WC4jtcA3b9vUK5AlyPOqWRJzpzApmWGlMNUd4nZQhqAdftp_UeK

ret = client.Write(fd384, []byte("YgWUdlx_tsBKN_5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) aITibAqnyiZ7bT79sgJBKFKVPFc3A8MQbYgWUdlx_tsBKN_5FHFpM6miij1Dc6WC4jtcA3b9vUK5AlyPOqWRJzpzApmWGlMNUd4nZQhqAdftp_UeK

ret = client.Seek(fd384, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


ret = client.Close(fd384)
if(ret != 0) {
  panic("close failed")
}


fd385 := client.Open("/RPoih7SFkr.txt", client.O_RDWR|client.O_CREATE)
if(fd385 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd385, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd386 := client.Open("/Uq03S3GN57.txt", client.O_RDWR|client.O_CREATE)
if(fd386 < 0) {
  panic("open failed")
}


fd387 := client.Open("/y0dHmoJ4hA.txt", client.O_RDWR|client.O_CREATE)
if(fd387 < 0) {
  panic("open failed")
}


fd388 := client.Open("/LcGiEGoTe0.txt", client.O_RDWR|client.O_CREATE)
if(fd388 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd388, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd385, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd389 := client.Open("/UHzput6Vyv.txt", client.O_RDWR|client.O_CREATE)
if(fd389 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd389, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHU") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd388, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd389, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd386)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd388, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd390 := client.Open("/otUQRxsRyj.txt", client.O_RDWR|client.O_CREATE)
if(fd390 < 0) {
  panic("open failed")
}


ret = client.Seek(fd385, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd387, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd387, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd385)
if(ret != 0) {
  panic("close failed")
}


fd391 := client.Open("/UkkCIjbjVT.txt", client.O_RDWR|client.O_CREATE)
if(fd391 < 0) {
  panic("open failed")
}


ret = client.Close(fd390)
if(ret != 0) {
  panic("close failed")
}


fd392 := client.Open("/GI8sQv42jD.txt", client.O_RDWR|client.O_CREATE)
if(fd392 < 0) {
  panic("open failed")
}


fd393 := client.Open("/8LABDJ_PCD.txt", client.O_RDWR|client.O_CREATE)
if(fd393 < 0) {
  panic("open failed")
}


ret = client.Close(fd391)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd393, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd387)
if(ret != 0) {
  panic("close failed")
}


fd394 := client.Open("/879TeeQEuG.txt", client.O_RDWR|client.O_CREATE)
if(fd394 < 0) {
  panic("open failed")
}


ret = client.Close(fd392)
if(ret != 0) {
  panic("close failed")
}


fd395 := client.Open("/JgvYBa4b20.txt", client.O_RDWR|client.O_CREATE)
if(fd395 < 0) {
  panic("open failed")
}


ret = client.Seek(fd388, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd396 := client.Open("/S5KOScMkUi.txt", client.O_RDWR|client.O_CREATE)
if(fd396 < 0) {
  panic("open failed")
}


ret = client.Close(fd388)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd394)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd389, 183, client.SEEK_SET)
if(ret != 183) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 183)
  panic("seek failed")
}


fd397 := client.Open("/gfY_Y9jHYW.txt", client.O_RDWR|client.O_CREATE)
if(fd397 < 0) {
  panic("open failed")
}


ret = client.Seek(fd393, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd395, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd397)
if(ret != 0) {
  panic("close failed")
}


fd398 := client.Open("/R7CKgaxysw.txt", client.O_RDWR|client.O_CREATE)
if(fd398 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd389, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "CTWLyXdV0xNBNYebaHHWzLOZ4Sig") {
  panic("wrong data returned")
}


fd399 := client.Open("/2KEc6YkBum.txt", client.O_RDWR|client.O_CREATE)
if(fd399 < 0) {
  panic("open failed")
}


ret = client.Close(fd393)
if(ret != 0) {
  panic("close failed")
}

//fd state: (211) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4Sig

ret = client.Write(fd389, []byte("ffyZTuOsmEQcwdfboQy0YzpcA7vfCOvSlnkrp6aR3wN3XoaAHtxKDdMyE3T8JD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (273) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4SigffyZTuOsmEQcwdfboQy0YzpcA7vfCOvSlnkrp6aR3wN3XoaAHtxKDdMyE3T8JD
//fd state: (0) 

ret = client.Write(fd395, []byte("JgfqoH6GMqDU9OZ6hP1tAmNr_CQPYqgHb2st75HgRmPpIe2J"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) JgfqoH6GMqDU9OZ6hP1tAmNr_CQPYqgHb2st75HgRmPpIe2J
//fd state: (273) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4SigffyZTuOsmEQcwdfboQy0YzpcA7vfCOvSlnkrp6aR3wN3XoaAHtxKDdMyE3T8JD

ret = client.Write(fd389, []byte("0NAUm3l0XhF_KaUJvAZDnUzuup"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (299) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4SigffyZTuOsmEQcwdfboQy0YzpcA7vfCOvSlnkrp6aR3wN3XoaAHtxKDdMyE3T8JD0NAUm3l0XhF_KaUJvAZDnUzuup

ret = client.Seek(fd399, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd399, []byte("4NqJsTuzXef2ZvwVKyOl3lJEyKWgLVK7R4g2Y3lqoLOtGymqlSkNrVHt8FuWkGFY3eTtTkC32"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) 4NqJsTuzXef2ZvwVKyOl3lJEyKWgLVK7R4g2Y3lqoLOtGymqlSkNrVHt8FuWkGFY3eTtTkC32

fd400 := client.Open("/_gbGbZZtm6.txt", client.O_RDWR|client.O_CREATE)
if(fd400 < 0) {
  panic("open failed")
}


fd401 := client.Open("/T4EfyJXeyW.txt", client.O_RDWR|client.O_CREATE)
if(fd401 < 0) {
  panic("open failed")
}


ret = client.Seek(fd399, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


buf, ret = client.Read(fd396, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (48) JgfqoH6GMqDU9OZ6hP1tAmNr_CQPYqgHb2st75HgRmPpIe2J

ret = client.Write(fd395, []byte("fQn2jBXRTsyw0qsmQlWZTK4oqUtLHjkapXAc9PiVv9O5dMY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) JgfqoH6GMqDU9OZ6hP1tAmNr_CQPYqgHb2st75HgRmPpIe2JfQn2jBXRTsyw0qsmQlWZTK4oqUtLHjkapXAc9PiVv9O5dMY

fd402 := client.Open("/nQ29N2qWtv.txt", client.O_RDWR|client.O_CREATE)
if(fd402 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd400, []byte("zy1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) zy1

fd403 := client.Open("/Xnvew95MH0.txt", client.O_RDWR|client.O_CREATE)
if(fd403 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd402, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd400, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd396, []byte("SJwQ0p_rp_S4BeohliazdTvD0BUZFDk_Oan7ypSz8KiqUOOUPBC2rD3791vI98korSDavulMEs_I6wp1UQq3SYfn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) SJwQ0p_rp_S4BeohliazdTvD0BUZFDk_Oan7ypSz8KiqUOOUPBC2rD3791vI98korSDavulMEs_I6wp1UQq3SYfn
//fd state: (299) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4SigffyZTuOsmEQcwdfboQy0YzpcA7vfCOvSlnkrp6aR3wN3XoaAHtxKDdMyE3T8JD0NAUm3l0XhF_KaUJvAZDnUzuup

ret = client.Write(fd389, []byte("EgM_Prnh0GT4hDKgCBgc5U6oyqFqWy72pJ1Ik8rcnom"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (342) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4SigffyZTuOsmEQcwdfboQy0YzpcA7vfCOvSlnkrp6aR3wN3XoaAHtxKDdMyE3T8JD0NAUm3l0XhF_KaUJvAZDnUzuupEgM_Prnh0GT4hDKgCBgc5U6oyqFqWy72pJ1Ik8rcnom

ret = client.Close(fd401)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd389, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd404 := client.Open("/VVprEbK7Bu.txt", client.O_RDWR|client.O_CREATE)
if(fd404 < 0) {
  panic("open failed")
}


ret = client.Close(fd403)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd398, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd405 := client.Open("/tbLiMuf1nh.txt", client.O_RDWR|client.O_CREATE)
if(fd405 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd400, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd406 := client.Open("/i7rZ5jaEkR.txt", client.O_RDWR|client.O_CREATE)
if(fd406 < 0) {
  panic("open failed")
}


fd407 := client.Open("/XYg89mK8WS.txt", client.O_RDWR|client.O_CREATE)
if(fd407 < 0) {
  panic("open failed")
}


ret = client.Seek(fd405, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd389, 244, client.SEEK_SET)
if(ret != 244) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 244)
  panic("seek failed")
}


buf, ret = client.Read(fd396, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd396, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}


fd408 := client.Open("/fe6x_kC5oC.txt", client.O_RDWR|client.O_CREATE)
if(fd408 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd406, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "RpewFtsQ_UKw4Gp82K9rAGTrcIEUUazShcCqtGQsgVOP8GoOe6Y0QdlkL6OorxJMfZuU_9QrEmFyjQ4RJxwVOsekY") {
  panic("wrong data returned")
}


fd409 := client.Open("/iSfKjSx1vt.txt", client.O_RDWR|client.O_CREATE)
if(fd409 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd398, []byte("ebjMzvbq25vAW10MK43PdaYk9FQ6ejYsA26"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) ebjMzvbq25vAW10MK43PdaYk9FQ6ejYsA26

buf, ret = client.Read(fd395, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd396)
if(ret != 0) {
  panic("close failed")
}


fd410 := client.Open("/zvfpB2u_jD.txt", client.O_RDWR|client.O_CREATE)
if(fd410 < 0) {
  panic("open failed")
}

//fd state: (244) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4SigffyZTuOsmEQcwdfboQy0YzpcA7vfCOvSlnkrp6aR3wN3XoaAHtxKDdMyE3T8JD0NAUm3l0XhF_KaUJvAZDnUzuupEgM_Prnh0GT4hDKgCBgc5U6oyqFqWy72pJ1Ik8rcnom

ret = client.Write(fd389, []byte("AT9OUCu8JSyavuFWqbalQ6WewNmaUB4KZsL9SOyhSuO5m7mDhGK6G5w9h8Tmfx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (306) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4SigffyZTuOsmEQcwdfboQy0YzpcA7vfCOvSlAT9OUCu8JSyavuFWqbalQ6WewNmaUB4KZsL9SOyhSuO5m7mDhGK6G5w9h8Tmfxh0GT4hDKgCBgc5U6oyqFqWy72pJ1Ik8rcnom

ret = client.Seek(fd405, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd408, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}

//fd state: (95) JgfqoH6GMqDU9OZ6hP1tAmNr_CQPYqgHb2st75HgRmPpIe2JfQn2jBXRTsyw0qsmQlWZTK4oqUtLHjkapXAc9PiVv9O5dMY

ret = client.Write(fd395, []byte("WAPeq8qCdffJF9_czrKpiBPgwUrEjFMBz4H14n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) JgfqoH6GMqDU9OZ6hP1tAmNr_CQPYqgHb2st75HgRmPpIe2JfQn2jBXRTsyw0qsmQlWZTK4oqUtLHjkapXAc9PiVv9O5dMYWAPeq8qCdffJF9_czrKpiBPgwUrEjFMBz4H14n

ret = client.Seek(fd404, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd405)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd409)
if(ret != 0) {
  panic("close failed")
}


fd411 := client.Open("/T6Y4E8A7YK.txt", client.O_RDWR|client.O_CREATE)
if(fd411 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd402, []byte("V01lnejFlnzulKWVCHqZwqNHqZ7AUxn2au1NCePZ_UkV5KUXFRQjK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) V01lnejFlnzulKWVCHqZwqNHqZ7AUxn2au1NCePZ_UkV5KUXFRQjK
//fd state: (0) 

ret = client.Write(fd407, []byte("zZW9TQmBWHlV2IOqMox26y4Fehs55kVegUKxuv5YXbgAf0gx5XkRJdYoE_HlKp6SEa7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) zZW9TQmBWHlV2IOqMox26y4Fehs55kVegUKxuv5YXbgAf0gx5XkRJdYoE_HlKp6SEa7

fd412 := client.Open("/Avi_dS2O4_.txt", client.O_RDWR|client.O_CREATE)
if(fd412 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd404, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd398, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd408, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4G87K7M2G") {
  panic("wrong data returned")
}


fd413 := client.Open("/PAI60AcLwF.txt", client.O_RDWR|client.O_CREATE)
if(fd413 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd413, []byte("_vlTJfgIXwIgIkDivOfNktpjGvw9rc7dWQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) _vlTJfgIXwIgIkDivOfNktpjGvw9rc7dWQ

buf, ret = client.Read(fd410, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xyHFttu_tPoeRaZg9dy6ghopA_AnUBK8xGWi5_m2Nl") {
  panic("wrong data returned")
}


ret = client.Seek(fd410, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


ret = client.Close(fd406)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd410, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "o") {
  panic("wrong data returned")
}


fd414 := client.Open("/SZ3w7aLh_u.txt", client.O_RDWR|client.O_CREATE)
if(fd414 < 0) {
  panic("open failed")
}


ret = client.Close(fd411)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd404, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd415 := client.Open("/fi0xquNQcl.txt", client.O_RDWR|client.O_CREATE)
if(fd415 < 0) {
  panic("open failed")
}

//fd state: (34) _vlTJfgIXwIgIkDivOfNktpjGvw9rc7dWQ

ret = client.Write(fd413, []byte("xElJSaOK6tObVduLvT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) _vlTJfgIXwIgIkDivOfNktpjGvw9rc7dWQxElJSaOK6tObVduLvT
//fd state: (306) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4SigffyZTuOsmEQcwdfboQy0YzpcA7vfCOvSlAT9OUCu8JSyavuFWqbalQ6WewNmaUB4KZsL9SOyhSuO5m7mDhGK6G5w9h8Tmfxh0GT4hDKgCBgc5U6oyqFqWy72pJ1Ik8rcnom

ret = client.Write(fd389, []byte("f6ELL7mxh8WV4srPaF32Q6YIfEi1cDOQG_tNIZb9WWiY6THbPBv_4mFOU20B"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (366) cY53Kbu8BURD4NgIWyUxM2i3mmxqUWxUfEM90oRPRHUvtQUkCK22Z_43JheVikrZbMCGpfTplS4nz0iR2hf5yr8TvnTo6nW1tYj4fSD7PDj56k8OfeFKvZw5s3mOh1xTVtiGtgIGzvVxm7Ax0XlySxrsevBnZ9lNuEMh7KQWRlBdFkWRT719gUPCTWLyXdV0xNBNYebaHHWzLOZ4SigffyZTuOsmEQcwdfboQy0YzpcA7vfCOvSlAT9OUCu8JSyavuFWqbalQ6WewNmaUB4KZsL9SOyhSuO5m7mDhGK6G5w9h8Tmfxf6ELL7mxh8WV4srPaF32Q6YIfEi1cDOQG_tNIZb9WWiY6THbPBv_4mFOU20B
//fd state: (133) JgfqoH6GMqDU9OZ6hP1tAmNr_CQPYqgHb2st75HgRmPpIe2JfQn2jBXRTsyw0qsmQlWZTK4oqUtLHjkapXAc9PiVv9O5dMYWAPeq8qCdffJF9_czrKpiBPgwUrEjFMBz4H14n

ret = client.Write(fd395, []byte("j5QOyNF2nm5cpZFd5NFerENnXPeuHb6iQuQJo57EdsR9h5gl0SrNpkqXP4KLbMhH_Om"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (200) JgfqoH6GMqDU9OZ6hP1tAmNr_CQPYqgHb2st75HgRmPpIe2JfQn2jBXRTsyw0qsmQlWZTK4oqUtLHjkapXAc9PiVv9O5dMYWAPeq8qCdffJF9_czrKpiBPgwUrEjFMBz4H14nj5QOyNF2nm5cpZFd5NFerENnXPeuHb6iQuQJo57EdsR9h5gl0SrNpkqXP4KLbMhH_Om

ret = client.Seek(fd400, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd414, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd408)
if(ret != 0) {
  panic("close failed")
}


fd416 := client.Open("/fncK2K1IOC.txt", client.O_RDWR|client.O_CREATE)
if(fd416 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd413)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd416)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd389)
if(ret != 0) {
  panic("close failed")
}

//fd state: (16) 4NqJsTuzXef2ZvwVKyOl3lJEyKWgLVK7R4g2Y3lqoLOtGymqlSkNrVHt8FuWkGFY3eTtTkC32

ret = client.Write(fd399, []byte("Kx4PIwT8nvF3e2USjQ5mHBjzgKyMRWKXXrQDI6MV3vTewyltVR3XocNM4UfFOIZF5YJmpPCtZforMnXxSCG9ECZbuVy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) 4NqJsTuzXef2ZvwVKx4PIwT8nvF3e2USjQ5mHBjzgKyMRWKXXrQDI6MV3vTewyltVR3XocNM4UfFOIZF5YJmpPCtZforMnXxSCG9ECZbuVy
//fd state: (107) 4NqJsTuzXef2ZvwVKx4PIwT8nvF3e2USjQ5mHBjzgKyMRWKXXrQDI6MV3vTewyltVR3XocNM4UfFOIZF5YJmpPCtZforMnXxSCG9ECZbuVy

ret = client.Write(fd399, []byte("82URixB8w_rMJ0BfSBpF08jic0n2t1uuh4BmMloS0iQLNYM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) 4NqJsTuzXef2ZvwVKx4PIwT8nvF3e2USjQ5mHBjzgKyMRWKXXrQDI6MV3vTewyltVR3XocNM4UfFOIZF5YJmpPCtZforMnXxSCG9ECZbuVy82URixB8w_rMJ0BfSBpF08jic0n2t1uuh4BmMloS0iQLNYM

buf, ret = client.Read(fd410, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pA_An") {
  panic("wrong data returned")
}

//fd state: (28) xyHFttu_tPoeRaZg9dy6ghopA_AnUBK8xGWi5_m2Nl

ret = client.Write(fd410, []byte("q7ojTfcWEaG7MmXIad0fyYGrlib6kc8cRf9p4NOBk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) xyHFttu_tPoeRaZg9dy6ghopA_Anq7ojTfcWEaG7MmXIad0fyYGrlib6kc8cRf9p4NOBk

ret = client.Seek(fd400, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd398, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd412)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd415, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (35) ebjMzvbq25vAW10MK43PdaYk9FQ6ejYsA26

ret = client.Write(fd398, []byte("biuzSe9XgApClS3EaiB9UcnK3ivDbgvCJNBbZye0EJ16hGjIKYELEFxw5yXC7EIrHopk6MTFl9cUZiTkyMrX6tZSOhUvs6oe6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) ebjMzvbq25vAW10MK43PdaYk9FQ6ejYsA26biuzSe9XgApClS3EaiB9UcnK3ivDbgvCJNBbZye0EJ16hGjIKYELEFxw5yXC7EIrHopk6MTFl9cUZiTkyMrX6tZSOhUvs6oe6

fd417 := client.Open("/S5KOScMkUi.txt", client.O_RDWR|client.O_CREATE)
if(fd417 < 0) {
  panic("open failed")
}


fd418 := client.Open("/5ZBJLLgBa4.txt", client.O_RDWR|client.O_CREATE)
if(fd418 < 0) {
  panic("open failed")
}


ret = client.Seek(fd395, 118, client.SEEK_SET)
if(ret != 118) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 118)
  panic("seek failed")
}


buf, ret = client.Read(fd418, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (154) 4NqJsTuzXef2ZvwVKx4PIwT8nvF3e2USjQ5mHBjzgKyMRWKXXrQDI6MV3vTewyltVR3XocNM4UfFOIZF5YJmpPCtZforMnXxSCG9ECZbuVy82URixB8w_rMJ0BfSBpF08jic0n2t1uuh4BmMloS0iQLNYM

ret = client.Write(fd399, []byte("TavQ7Tp4v0IQO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (167) 4NqJsTuzXef2ZvwVKx4PIwT8nvF3e2USjQ5mHBjzgKyMRWKXXrQDI6MV3vTewyltVR3XocNM4UfFOIZF5YJmpPCtZforMnXxSCG9ECZbuVy82URixB8w_rMJ0BfSBpF08jic0n2t1uuh4BmMloS0iQLNYMTavQ7Tp4v0IQO

ret = client.Close(fd417)
if(ret != 0) {
  panic("close failed")
}

//fd state: (7) F8iipPiHawxWX_Ch2fqid_H04r9m99xtMgy558xL1LtfCRLyGFsfFDPxug0cN26ZKRYMkVHZLMQzBUC5AFp

ret = client.Write(fd414, []byte("xcUW2Ock4YjQsXL_t9Kijfn0UdWXBwxVKllTVvrgrNxODQofVpmknDczTdhPHbh1gK2dn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) F8iipPixcUW2Ock4YjQsXL_t9Kijfn0UdWXBwxVKllTVvrgrNxODQofVpmknDczTdhPHbh1gK2dnBUC5AFp

ret = client.Seek(fd407, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


ret = client.Seek(fd414, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


fd419 := client.Open("/gF3CBhfROh.txt", client.O_RDWR|client.O_CREATE)
if(fd419 < 0) {
  panic("open failed")
}


fd420 := client.Open("/wCijCfQcBC.txt", client.O_RDWR|client.O_CREATE)
if(fd420 < 0) {
  panic("open failed")
}


ret = client.Close(fd420)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd419, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd414, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "knDczTdhPHbh1gK2dnBUC5AFp") {
  panic("wrong data returned")
}


fd421 := client.Open("/eaI_co6DSS.txt", client.O_RDWR|client.O_CREATE)
if(fd421 < 0) {
  panic("open failed")
}


ret = client.Close(fd415)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd395)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd418, []byte("JSx2FqpSSepAJKI3nli5ZvgsQsFA_X5p_E2PbgIet"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) JSx2FqpSSepAJKI3nli5ZvgsQsFA_X5p_E2PbgIet

buf, ret = client.Read(fd418, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd422 := client.Open("/qkb5m1FlgT.txt", client.O_RDWR|client.O_CREATE)
if(fd422 < 0) {
  panic("open failed")
}


ret = client.Seek(fd399, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


fd423 := client.Open("/o_zaeN62w2.txt", client.O_RDWR|client.O_CREATE)
if(fd423 < 0) {
  panic("open failed")
}


ret = client.Seek(fd410, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}


ret = client.Close(fd419)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd407, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


fd424 := client.Open("/_BLukPDMrv.txt", client.O_RDWR|client.O_CREATE)
if(fd424 < 0) {
  panic("open failed")
}


ret = client.Seek(fd398, 132, client.SEEK_SET)
if(ret != 132) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 132)
  panic("seek failed")
}


ret = client.Close(fd407)
if(ret != 0) {
  panic("close failed")
}


fd425 := client.Open("/ALVR8XJH6r.txt", client.O_RDWR|client.O_CREATE)
if(fd425 < 0) {
  panic("open failed")
}


ret = client.Seek(fd399, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


ret = client.Seek(fd399, 161, client.SEEK_SET)
if(ret != 161) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 161)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd423, []byte("g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJ
//fd state: (93) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJ

ret = client.Write(fd423, []byte("e4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (142) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO6

ret = client.Seek(fd414, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}

//fd state: (142) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO6

ret = client.Write(fd423, []byte("2C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (211) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsV

buf, ret = client.Read(fd399, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4v0IQO") {
  panic("wrong data returned")
}


ret = client.Seek(fd398, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


ret = client.Close(fd421)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd404, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd398)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd418, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd410)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd399)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd418, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (1) F8iipPixcUW2Ock4YjQsXL_t9Kijfn0UdWXBwxVKllTVvrgrNxODQofVpmknDczTdhPHbh1gK2dnBUC5AFp

ret = client.Write(fd414, []byte("3CyOeYJn7Jhk6sCp8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) F3CyOeYJn7Jhk6sCp8QsXL_t9Kijfn0UdWXBwxVKllTVvrgrNxODQofVpmknDczTdhPHbh1gK2dnBUC5AFp

ret = client.Close(fd422)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd425, []byte("SUbZTwa5r_XprdQjLe7wEjPjmqF5M4DH7M98idROQJyTauPsSmN7WVU0RAvBsnrhuKLKkNfccksDkHSLyFzOYmMpNbi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) SUbZTwa5r_XprdQjLe7wEjPjmqF5M4DH7M98idROQJyTauPsSmN7WVU0RAvBsnrhuKLKkNfccksDkHSLyFzOYmMpNbi

ret = client.Close(fd425)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd404, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (53) V01lnejFlnzulKWVCHqZwqNHqZ7AUxn2au1NCePZ_UkV5KUXFRQjK

ret = client.Write(fd402, []byte("FDoifU3WzdFFIQlEZMj7U"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) V01lnejFlnzulKWVCHqZwqNHqZ7AUxn2au1NCePZ_UkV5KUXFRQjKFDoifU3WzdFFIQlEZMj7U

ret = client.Close(fd402)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd404, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd400, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd404, []byte("HRFrlIiuWPP59snHSeo1QyzgwJsxIMLzLBfYK7VZ4nXcgmtu2RRV2A0yBiv_DlXbYNJQ1D_C6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) HRFrlIiuWPP59snHSeo1QyzgwJsxIMLzLBfYK7VZ4nXcgmtu2RRV2A0yBiv_DlXbYNJQ1D_C6

ret = client.Seek(fd400, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Close(fd418)
if(ret != 0) {
  panic("close failed")
}


fd426 := client.Open("/kiQZc5jb4w.txt", client.O_RDWR|client.O_CREATE)
if(fd426 < 0) {
  panic("open failed")
}


fd427 := client.Open("/9n5lJZ73Ah.txt", client.O_RDWR|client.O_CREATE)
if(fd427 < 0) {
  panic("open failed")
}


ret = client.Close(fd424)
if(ret != 0) {
  panic("close failed")
}


fd428 := client.Open("/XyzWAQDWkU.txt", client.O_RDWR|client.O_CREATE)
if(fd428 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd423, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd426)
if(ret != 0) {
  panic("close failed")
}

//fd state: (2) zy1

ret = client.Write(fd400, []byte("mbI2sxWd7cv8rGAegfVguM8IzyCip1NB4lqUFBdeHRx7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) zymbI2sxWd7cv8rGAegfVguM8IzyCip1NB4lqUFBdeHRx7

fd429 := client.Open("/ZmjsjhxUuS.txt", client.O_RDWR|client.O_CREATE)
if(fd429 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd429, []byte("AMng8qGz1LsfSIILvvMxLy47HpC5498qRcsU9K5WoOt6v1NCwlS7tsOvADVXRbPqriBZGoipg_vmSrtxptEnuTK4gQ69cpmPuzx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) AMng8qGz1LsfSIILvvMxLy47HpC5498qRcsU9K5WoOt6v1NCwlS7tsOvADVXRbPqriBZGoipg_vmSrtxptEnuTK4gQ69cpmPuzx

fd430 := client.Open("/PIpopWJsly.txt", client.O_RDWR|client.O_CREATE)
if(fd430 < 0) {
  panic("open failed")
}


ret = client.Seek(fd414, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}


buf, ret = client.Read(fd430, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "DuTgaI19rxYDs8bDc0lxXKGPFyiQJ") {
  panic("wrong data returned")
}


ret = client.Close(fd430)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd404, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 9abPdTLhlhAUmBz5WkfxBAjGoaN3jeJC7wkxL2m2RMeaqDyLe5SeWX0q3deX4RTv99zygkQ35QYBIdOPU_VNPcDmZWK6lw4w791JXvwZRJZbV44SNFh_RRpU75o8bUzu

ret = client.Write(fd428, []byte("075KjHENMdWWbgCH2PJjMB4OQwus6X4eG5j69FJkDaeK1b2lkWmvAnArnTYDRs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) 075KjHENMdWWbgCH2PJjMB4OQwus6X4eG5j69FJkDaeK1b2lkWmvAnArnTYDRsTv99zygkQ35QYBIdOPU_VNPcDmZWK6lw4w791JXvwZRJZbV44SNFh_RRpU75o8bUzu

buf, ret = client.Read(fd429, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (211) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsV

ret = client.Write(fd423, []byte("B7m701lwVXCKJjxrbajNwaByif9zOxAIoec3ng8rEdCzkTvAwlV0sk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (265) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsVB7m701lwVXCKJjxrbajNwaByif9zOxAIoec3ng8rEdCzkTvAwlV0sk
//fd state: (73) HRFrlIiuWPP59snHSeo1QyzgwJsxIMLzLBfYK7VZ4nXcgmtu2RRV2A0yBiv_DlXbYNJQ1D_C6

ret = client.Write(fd404, []byte("GPjqgsGon1LtE1PNr1T5y6VCuaWIqHiBW1rqSzvH_GOOV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) HRFrlIiuWPP59snHSeo1QyzgwJsxIMLzLBfYK7VZ4nXcgmtu2RRV2A0yBiv_DlXbYNJQ1D_C6GPjqgsGon1LtE1PNr1T5y6VCuaWIqHiBW1rqSzvH_GOOV

fd431 := client.Open("/jMLCbPratd.txt", client.O_RDWR|client.O_CREATE)
if(fd431 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd404, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd429, 83, client.SEEK_SET)
if(ret != 83) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 83)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd427, []byte("Z7HI2E0SlH9teELZ6f0C1kBRYYsCXYK9Chwc8FQEF5_L3o_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) Z7HI2E0SlH9teELZ6f0C1kBRYYsCXYK9Chwc8FQEF5_L3o_

buf, ret = client.Read(fd427, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (47) Z7HI2E0SlH9teELZ6f0C1kBRYYsCXYK9Chwc8FQEF5_L3o_

ret = client.Write(fd427, []byte("MQN9D2AhgzEFisCNnnvfK4gqzpMQYRxSKTSdvMwem9z2HCEnDPg2G40m4vL2dSu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) Z7HI2E0SlH9teELZ6f0C1kBRYYsCXYK9Chwc8FQEF5_L3o_MQN9D2AhgzEFisCNnnvfK4gqzpMQYRxSKTSdvMwem9z2HCEnDPg2G40m4vL2dSu

ret = client.Close(fd400)
if(ret != 0) {
  panic("close failed")
}

//fd state: (265) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsVB7m701lwVXCKJjxrbajNwaByif9zOxAIoec3ng8rEdCzkTvAwlV0sk

ret = client.Write(fd423, []byte("5egFeeRyycRNUL8NUPMIZsjCFHeczWybPeVvWR5Xb3EY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (309) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsVB7m701lwVXCKJjxrbajNwaByif9zOxAIoec3ng8rEdCzkTvAwlV0sk5egFeeRyycRNUL8NUPMIZsjCFHeczWybPeVvWR5Xb3EY

ret = client.Seek(fd427, 104, client.SEEK_SET)
if(ret != 104) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 104)
  panic("seek failed")
}


buf, ret = client.Read(fd414, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "K2d") {
  panic("wrong data returned")
}


ret = client.Seek(fd414, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


buf, ret = client.Read(fd427, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vL2dSu") {
  panic("wrong data returned")
}


ret = client.Seek(fd404, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


buf, ret = client.Read(fd428, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Tv99zygkQ35QYBIdOPU_VNPcDmZWK6lw4w") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd427)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd429, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}

//fd state: (309) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsVB7m701lwVXCKJjxrbajNwaByif9zOxAIoec3ng8rEdCzkTvAwlV0sk5egFeeRyycRNUL8NUPMIZsjCFHeczWybPeVvWR5Xb3EY

ret = client.Write(fd423, []byte("XbYK5k5nMcM8DZQpHxaCUy1ihcalcIze98T2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (345) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsVB7m701lwVXCKJjxrbajNwaByif9zOxAIoec3ng8rEdCzkTvAwlV0sk5egFeeRyycRNUL8NUPMIZsjCFHeczWybPeVvWR5Xb3EYXbYK5k5nMcM8DZQpHxaCUy1ihcalcIze98T2
//fd state: (345) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsVB7m701lwVXCKJjxrbajNwaByif9zOxAIoec3ng8rEdCzkTvAwlV0sk5egFeeRyycRNUL8NUPMIZsjCFHeczWybPeVvWR5Xb3EYXbYK5k5nMcM8DZQpHxaCUy1ihcalcIze98T2

ret = client.Write(fd423, []byte("CB6YsSvztnyIf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (358) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsVB7m701lwVXCKJjxrbajNwaByif9zOxAIoec3ng8rEdCzkTvAwlV0sk5egFeeRyycRNUL8NUPMIZsjCFHeczWybPeVvWR5Xb3EYXbYK5k5nMcM8DZQpHxaCUy1ihcalcIze98T2CB6YsSvztnyIf

ret = client.Close(fd404)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd414, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xVKllTVvrgrNxODQofVpmknDczTdhPHbh1gK2dnBUC5AFp") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd431, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wqYVnAjP6G9KDsf9Q_dYKmZuYddqYXqkjJVPVyC8TpnCz8htDuMrodZSjqYf7AR53z") {
  panic("wrong data returned")
}


fd432 := client.Open("/62lJCMOXTv.txt", client.O_RDWR|client.O_CREATE)
if(fd432 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd432, []byte("3HewuwrjrFv3s26iaPtj4WzWbJN1g6aFz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) 3HewuwrjrFv3s26iaPtj4WzWbJN1g6aFz

fd433 := client.Open("/3oDyCHmOH9.txt", client.O_RDWR|client.O_CREATE)
if(fd433 < 0) {
  panic("open failed")
}


ret = client.Close(fd414)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd431, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


ret = client.Close(fd432)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd423, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd423, 142, client.SEEK_SET)
if(ret != 142) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 142)
  panic("seek failed")
}


buf, ret = client.Read(fd423, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4X") {
  panic("wrong data returned")
}


fd434 := client.Open("/T4EfyJXeyW.txt", client.O_RDWR|client.O_CREATE)
if(fd434 < 0) {
  panic("open failed")
}


ret = client.Seek(fd431, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


fd435 := client.Open("/5Z_ri0JLY3.txt", client.O_RDWR|client.O_CREATE)
if(fd435 < 0) {
  panic("open failed")
}


fd436 := client.Open("/ICHMbbOWY4.txt", client.O_RDWR|client.O_CREATE)
if(fd436 < 0) {
  panic("open failed")
}


ret = client.Seek(fd423, 265, client.SEEK_SET)
if(ret != 265) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 265)
  panic("seek failed")
}

//fd state: (265) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsVB7m701lwVXCKJjxrbajNwaByif9zOxAIoec3ng8rEdCzkTvAwlV0sk5egFeeRyycRNUL8NUPMIZsjCFHeczWybPeVvWR5Xb3EYXbYK5k5nMcM8DZQpHxaCUy1ihcalcIze98T2CB6YsSvztnyIf

ret = client.Write(fd423, []byte("dtbTogCBQir84ANjEQtM2mhAiwLLxgqJ1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (298) g23u9jhoQ8d7KNzTIsE03nq01KO4gRfJnJN4d2pbbGAXlhpSCzjBhnT4XrDiooCP8Z8U42ieTQPf_pUy0sjyPNBVIlAyJe4Sk0hVEfLoQ_rWjAWVVNAfYLca0HXijItaQOtHuyjKuGwpO62C3LU7zjr_aZ9OtoxwE2_F00ZFyb6cIK_9F0aSByL7Wh4XAKTKSPTkF_dR8QDHcqfCIsVB7m701lwVXCKJjxrbajNwaByif9zOxAIoec3ng8rEdCzkTvAwlV0skdtbTogCBQir84ANjEQtM2mhAiwLLxgqJ1eVvWR5Xb3EYXbYK5k5nMcM8DZQpHxaCUy1ihcalcIze98T2CB6YsSvztnyIf
//fd state: (0) 

ret = client.Write(fd433, []byte("V2u6g3Uw7AZj8Rtb9mO9VoKmG_RXEMTEjSPLRoDT7cn9tk42Gv_z_A58T"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) V2u6g3Uw7AZj8Rtb9mO9VoKmG_RXEMTEjSPLRoDT7cn9tk42Gv_z_A58T
//fd state: (96) 075KjHENMdWWbgCH2PJjMB4OQwus6X4eG5j69FJkDaeK1b2lkWmvAnArnTYDRsTv99zygkQ35QYBIdOPU_VNPcDmZWK6lw4w791JXvwZRJZbV44SNFh_RRpU75o8bUzu

ret = client.Write(fd428, []byte("6k5KumtcVrWBqVjlIWndV2W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (119) 075KjHENMdWWbgCH2PJjMB4OQwus6X4eG5j69FJkDaeK1b2lkWmvAnArnTYDRsTv99zygkQ35QYBIdOPU_VNPcDmZWK6lw4w6k5KumtcVrWBqVjlIWndV2WU75o8bUzu

ret = client.Close(fd433)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd431)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd436, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd428)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd436)
if(ret != 0) {
  panic("close failed")
}

//fd state: (5) AMng8qGz1LsfSIILvvMxLy47HpC5498qRcsU9K5WoOt6v1NCwlS7tsOvADVXRbPqriBZGoipg_vmSrtxptEnuTK4gQ69cpmPuzx

ret = client.Write(fd429, []byte("ee5_U"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) AMng8ee5_UsfSIILvvMxLy47HpC5498qRcsU9K5WoOt6v1NCwlS7tsOvADVXRbPqriBZGoipg_vmSrtxptEnuTK4gQ69cpmPuzx

fd437 := client.Open("/ebgGYWPeVA.txt", client.O_RDWR|client.O_CREATE)
if(fd437 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd435, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd438 := client.Open("/ydk9BRFkKJ.txt", client.O_RDWR|client.O_CREATE)
if(fd438 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd438, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nnBaL1SCiqtTmA2BvDsRyRiQV6tG7yVJi5UoCIdr5f4B0") {
  panic("wrong data returned")
}


ret = client.Close(fd435)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd434, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd439 := client.Open("/y0dHmoJ4hA.txt", client.O_RDWR|client.O_CREATE)
if(fd439 < 0) {
  panic("open failed")
}


fd440 := client.Open("/0WU4yNMJLm.txt", client.O_RDWR|client.O_CREATE)
if(fd440 < 0) {
  panic("open failed")
}


ret = client.Close(fd439)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd438, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Seek(fd429, 94, client.SEEK_SET)
if(ret != 94) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 94)
  panic("seek failed")
}


ret = client.Seek(fd429, 88, client.SEEK_SET)
if(ret != 88) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 88)
  panic("seek failed")
}


ret = client.Close(fd440)
if(ret != 0) {
  panic("close failed")
}


fd441 := client.Open("/4ldmhfQNz1.txt", client.O_RDWR|client.O_CREATE)
if(fd441 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd437, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd441, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd442 := client.Open("/bFM001bg35.txt", client.O_RDWR|client.O_CREATE)
if(fd442 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd434, []byte("GTN2OvkNGa5BycCXrOIvtKbe3xQZqR0F3A1YijsPMdeDG1RKXdTXgTwEor0zEcumpM1REFP1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) GTN2OvkNGa5BycCXrOIvtKbe3xQZqR0F3A1YijsPMdeDG1RKXdTXgTwEor0zEcumpM1REFP1

ret = client.Close(fd442)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd429, 71, client.SEEK_SET)
if(ret != 71) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 71)
  panic("seek failed")
}


fd443 := client.Open("/gF3CBhfROh.txt", client.O_RDWR|client.O_CREATE)
if(fd443 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd443, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (13) nnBaL1SCiqtTmA2BvDsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRNY4ILShBprPOHDd4Pz

ret = client.Write(fd438, []byte("4jsyY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) nnBaL1SCiqtTm4jsyYsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRNY4ILShBprPOHDd4Pz

ret = client.Seek(fd437, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd438, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}

//fd state: (58) nnBaL1SCiqtTm4jsyYsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRNY4ILShBprPOHDd4Pz

ret = client.Write(fd438, []byte("UFK8CFlcP4MqLMoUNiBWIIS3KGVw9T53feTxxdrpVYvEVeISgj6TBBeMLEijvnVXEEdRFz0gQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (131) nnBaL1SCiqtTm4jsyYsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRUFK8CFlcP4MqLMoUNiBWIIS3KGVw9T53feTxxdrpVYvEVeISgj6TBBeMLEijvnVXEEdRFz0gQ

ret = client.Close(fd443)
if(ret != 0) {
  panic("close failed")
}


fd444 := client.Open("/ylHMXBDQc2.txt", client.O_RDWR|client.O_CREATE)
if(fd444 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd434, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd429, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pg_vmSrtxptEnuTK4gQ69cpmPuz") {
  panic("wrong data returned")
}


ret = client.Close(fd423)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd429, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "x") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd438, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd437, []byte("Zh6CW7PV1FV8y8ajXx25Aa7l"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) Zh6CW7PV1FV8y8ajXx25Aa7l
//fd state: (72) GTN2OvkNGa5BycCXrOIvtKbe3xQZqR0F3A1YijsPMdeDG1RKXdTXgTwEor0zEcumpM1REFP1

ret = client.Write(fd434, []byte("hgvhG8rUwunqJlWIs6j_IxNIbJminsvsyMbckUcRj8YbiId9rBnEmTqYQ15ef5GYGcJQR4x0Q0jbJlGUGk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) GTN2OvkNGa5BycCXrOIvtKbe3xQZqR0F3A1YijsPMdeDG1RKXdTXgTwEor0zEcumpM1REFP1hgvhG8rUwunqJlWIs6j_IxNIbJminsvsyMbckUcRj8YbiId9rBnEmTqYQ15ef5GYGcJQR4x0Q0jbJlGUGk

fd445 := client.Open("/UHbNJwnB4n.txt", client.O_RDWR|client.O_CREATE)
if(fd445 < 0) {
  panic("open failed")
}


fd446 := client.Open("/c3zYasdKwN.txt", client.O_RDWR|client.O_CREATE)
if(fd446 < 0) {
  panic("open failed")
}


fd447 := client.Open("/uE9P5oukcF.txt", client.O_RDWR|client.O_CREATE)
if(fd447 < 0) {
  panic("open failed")
}


ret = client.Seek(fd434, 85, client.SEEK_SET)
if(ret != 85) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 85)
  panic("seek failed")
}


ret = client.Close(fd445)
if(ret != 0) {
  panic("close failed")
}

//fd state: (131) nnBaL1SCiqtTm4jsyYsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRUFK8CFlcP4MqLMoUNiBWIIS3KGVw9T53feTxxdrpVYvEVeISgj6TBBeMLEijvnVXEEdRFz0gQ

ret = client.Write(fd438, []byte("GqbJQ3cB_nEpKOLl4KnUqaI6DLagY99O3OGkxA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (169) nnBaL1SCiqtTm4jsyYsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRUFK8CFlcP4MqLMoUNiBWIIS3KGVw9T53feTxxdrpVYvEVeISgj6TBBeMLEijvnVXEEdRFz0gQGqbJQ3cB_nEpKOLl4KnUqaI6DLagY99O3OGkxA
//fd state: (0) 

ret = client.Write(fd446, []byte("OkZljU26x5o7ziZhg96vhnKtgFSlvAHoOBzC6j"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) OkZljU26x5o7ziZhg96vhnKtgFSlvAHoOBzC6j

buf, ret = client.Read(fd441, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3SouSO") {
  panic("wrong data returned")
}


ret = client.Close(fd437)
if(ret != 0) {
  panic("close failed")
}


fd448 := client.Open("/PiOdP6CVcO.txt", client.O_RDWR|client.O_CREATE)
if(fd448 < 0) {
  panic("open failed")
}


ret = client.Seek(fd434, 82, client.SEEK_SET)
if(ret != 82) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 82)
  panic("seek failed")
}


ret = client.Seek(fd429, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}


fd449 := client.Open("/JfPWUs3vJS.txt", client.O_RDWR|client.O_CREATE)
if(fd449 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd449, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd434, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nqJlWIs6") {
  panic("wrong data returned")
}


ret = client.Close(fd434)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd447)
if(ret != 0) {
  panic("close failed")
}


fd450 := client.Open("/56n5wy1x4r.txt", client.O_RDWR|client.O_CREATE)
if(fd450 < 0) {
  panic("open failed")
}


ret = client.Seek(fd441, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}

//fd state: (169) nnBaL1SCiqtTm4jsyYsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRUFK8CFlcP4MqLMoUNiBWIIS3KGVw9T53feTxxdrpVYvEVeISgj6TBBeMLEijvnVXEEdRFz0gQGqbJQ3cB_nEpKOLl4KnUqaI6DLagY99O3OGkxA

ret = client.Write(fd438, []byte("J7B9LRcHK5SZ_MHuFJyEOLbIoEdLJTL9LqKXA31EcywulwHhQdzT3CkPlt5lPngOEN9Anp6YGwx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (244) nnBaL1SCiqtTm4jsyYsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRUFK8CFlcP4MqLMoUNiBWIIS3KGVw9T53feTxxdrpVYvEVeISgj6TBBeMLEijvnVXEEdRFz0gQGqbJQ3cB_nEpKOLl4KnUqaI6DLagY99O3OGkxAJ7B9LRcHK5SZ_MHuFJyEOLbIoEdLJTL9LqKXA31EcywulwHhQdzT3CkPlt5lPngOEN9Anp6YGwx
//fd state: (0) AkaXbzc28ZwkaWjIHD213MrNJOBnweNfv_JgFK3AQQ3

ret = client.Write(fd448, []byte("BYoRAOQLAbIs0AZRvZ9_JJvFleqwQFxV4mu3U3PudQCJRIfrfDpdsGoWKgAVC1FQ_Tuo2LDbl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) BYoRAOQLAbIs0AZRvZ9_JJvFleqwQFxV4mu3U3PudQCJRIfrfDpdsGoWKgAVC1FQ_Tuo2LDbl
//fd state: (5) BOx3SouSO

ret = client.Write(fd441, []byte("_JXfrmzTDq_PjKS0hhEjV8AejA3Fk4N832gtvrnAP3uG0hJqqmy1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) BOx3S_JXfrmzTDq_PjKS0hhEjV8AejA3Fk4N832gtvrnAP3uG0hJqqmy1

ret = client.Seek(fd446, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


fd451 := client.Open("/yY3JR0W531.txt", client.O_RDWR|client.O_CREATE)
if(fd451 < 0) {
  panic("open failed")
}


ret = client.Close(fd429)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd450)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd444, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd452 := client.Open("/vYftS72Zrw.txt", client.O_RDWR|client.O_CREATE)
if(fd452 < 0) {
  panic("open failed")
}


ret = client.Close(fd444)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd438, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}


ret = client.Close(fd438)
if(ret != 0) {
  panic("close failed")
}


fd453 := client.Open("/tOmCdGHlGk.txt", client.O_RDWR|client.O_CREATE)
if(fd453 < 0) {
  panic("open failed")
}


ret = client.Seek(fd446, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


buf, ret = client.Read(fd452, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd448, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd441)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd453, []byte("5yFGmBF_D7_i1IIq8WNisib8SnJ2zk5qLNx1Yga5egtKkp0v"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) 5yFGmBF_D7_i1IIq8WNisib8SnJ2zk5qLNx1Yga5egtKkp0v

ret = client.Seek(fd449, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd449, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (33) OkZljU26x5o7ziZhg96vhnKtgFSlvAHoOBzC6j

ret = client.Write(fd446, []byte("MFRNjslvNP_4JDLmbFPXmh1g6dcjTO3SJykse7fo8XzEVUqP9C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) OkZljU26x5o7ziZhg96vhnKtgFSlvAHoOMFRNjslvNP_4JDLmbFPXmh1g6dcjTO3SJykse7fo8XzEVUqP9C

ret = client.Seek(fd451, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd449)
if(ret != 0) {
  panic("close failed")
}


fd454 := client.Open("/li9xFJJEVH.txt", client.O_RDWR|client.O_CREATE)
if(fd454 < 0) {
  panic("open failed")
}


fd455 := client.Open("/ZoszLk9N4b.txt", client.O_RDWR|client.O_CREATE)
if(fd455 < 0) {
  panic("open failed")
}


fd456 := client.Open("/nQ29N2qWtv.txt", client.O_RDWR|client.O_CREATE)
if(fd456 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd453, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd453)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd455, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) V01lnejFlnzulKWVCHqZwqNHqZ7AUxn2au1NCePZ_UkV5KUXFRQjKFDoifU3WzdFFIQlEZMj7U

ret = client.Write(fd456, []byte("a5GOQvtggCLN1rQsUVDy4mVBeOpCSk_puqjmi0rw8LHsLck7lc6ElsrAzmxB8s8fIUMB4eYqLH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) a5GOQvtggCLN1rQsUVDy4mVBeOpCSk_puqjmi0rw8LHsLck7lc6ElsrAzmxB8s8fIUMB4eYqLH

ret = client.Seek(fd454, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd456, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


fd457 := client.Open("/Da3xOtaW62.txt", client.O_RDWR|client.O_CREATE)
if(fd457 < 0) {
  panic("open failed")
}


fd458 := client.Open("/OlMsanXw7a.txt", client.O_RDWR|client.O_CREATE)
if(fd458 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd448, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd459 := client.Open("/SgtN9yKcDb.txt", client.O_RDWR|client.O_CREATE)
if(fd459 < 0) {
  panic("open failed")
}


ret = client.Close(fd454)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd459, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PMo1oF2M9EEWm6plLqk1v0EoQrli") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd455, []byte("SM5PdxpsyQhNmilnWrPeKCHUAVGezJb6cZDrpyfs_gg0IX1UfNetOeu0EZpgmC5WoXCOWhwOQ1LRg1GpMwgb9xPPrfL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) SM5PdxpsyQhNmilnWrPeKCHUAVGezJb6cZDrpyfs_gg0IX1UfNetOeu0EZpgmC5WoXCOWhwOQ1LRg1GpMwgb9xPPrfL

ret = client.Close(fd457)
if(ret != 0) {
  panic("close failed")
}


fd460 := client.Open("/T4EfyJXeyW.txt", client.O_RDWR|client.O_CREATE)
if(fd460 < 0) {
  panic("open failed")
}


fd461 := client.Open("/oT8dR_w3Q8.txt", client.O_RDWR|client.O_CREATE)
if(fd461 < 0) {
  panic("open failed")
}


ret = client.Seek(fd458, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd462 := client.Open("/D24_Z9DpA2.txt", client.O_RDWR|client.O_CREATE)
if(fd462 < 0) {
  panic("open failed")
}


fd463 := client.Open("/kGLXVETOMb.txt", client.O_RDWR|client.O_CREATE)
if(fd463 < 0) {
  panic("open failed")
}


ret = client.Seek(fd461, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd451, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd451, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (91) SM5PdxpsyQhNmilnWrPeKCHUAVGezJb6cZDrpyfs_gg0IX1UfNetOeu0EZpgmC5WoXCOWhwOQ1LRg1GpMwgb9xPPrfL

ret = client.Write(fd455, []byte("LBWSwNXoxsCGwOAvxsVbnsNg7b83fziIG1HRPrR9kR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) SM5PdxpsyQhNmilnWrPeKCHUAVGezJb6cZDrpyfs_gg0IX1UfNetOeu0EZpgmC5WoXCOWhwOQ1LRg1GpMwgb9xPPrfLLBWSwNXoxsCGwOAvxsVbnsNg7b83fziIG1HRPrR9kR

ret = client.Close(fd462)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd461)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd448, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd456, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Close(fd458)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd456)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd451)
if(ret != 0) {
  panic("close failed")
}


fd464 := client.Open("/oJ_uY3Fb1U.txt", client.O_RDWR|client.O_CREATE)
if(fd464 < 0) {
  panic("open failed")
}


ret = client.Seek(fd446, 64, client.SEEK_SET)
if(ret != 64) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 64)
  panic("seek failed")
}


fd465 := client.Open("/g6sj5UBE8k.txt", client.O_RDWR|client.O_CREATE)
if(fd465 < 0) {
  panic("open failed")
}


fd466 := client.Open("/eAQn6wQOjv.txt", client.O_RDWR|client.O_CREATE)
if(fd466 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd452, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd448)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd463, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd466, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd446, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "SJykse7fo8XzEVUqP9C") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd452, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd463, []byte("8tJsFwlR24FbAD_XmTrcxrKaE4G7uX1yIutSEhRYgOfdy2euScHoGYvKjNscZRB6jA8se2E5FM7MR14xFFgr4F"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) 8tJsFwlR24FbAD_XmTrcxrKaE4G7uX1yIutSEhRYgOfdy2euScHoGYvKjNscZRB6jA8se2E5FM7MR14xFFgr4F

fd467 := client.Open("/oJi1_Yfhik.txt", client.O_RDWR|client.O_CREATE)
if(fd467 < 0) {
  panic("open failed")
}


ret = client.Close(fd452)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd465)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd467, []byte("uOrEmwea3_zJqJk3C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) uOrEmwea3_zJqJk3C

ret = client.Close(fd467)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd466, []byte("hnOceT6CX5NNeBJfRTh0xUi_i0xvsuGPUjPAfUa5htvQlh3lzWKXdoRsG68hJy94YhX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) hnOceT6CX5NNeBJfRTh0xUi_i0xvsuGPUjPAfUa5htvQlh3lzWKXdoRsG68hJy94YhX

ret = client.Seek(fd446, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}


fd468 := client.Open("/CdZyI4bMVY.txt", client.O_RDWR|client.O_CREATE)
if(fd468 < 0) {
  panic("open failed")
}


ret = client.Close(fd468)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd460, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GTN2OvkNGa5BycCXrOIvtKbe3xQZqR0F3A1Yi") {
  panic("wrong data returned")
}


fd469 := client.Open("/otUQRxsRyj.txt", client.O_RDWR|client.O_CREATE)
if(fd469 < 0) {
  panic("open failed")
}

//fd state: (30) OkZljU26x5o7ziZhg96vhnKtgFSlvAHoOMFRNjslvNP_4JDLmbFPXmh1g6dcjTO3SJykse7fo8XzEVUqP9C

ret = client.Write(fd446, []byte("9fHsa4vRZ3wlRIenjASoL1hNAtwIjqnDAYqTQWLW5_c1nPLr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) OkZljU26x5o7ziZhg96vhnKtgFSlvA9fHsa4vRZ3wlRIenjASoL1hNAtwIjqnDAYqTQWLW5_c1nPLrUqP9C

ret = client.Seek(fd459, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


buf, ret = client.Read(fd463, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd446, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "UqP9C") {
  panic("wrong data returned")
}


fd470 := client.Open("/mbhXIMTJYM.txt", client.O_RDWR|client.O_CREATE)
if(fd470 < 0) {
  panic("open failed")
}

//fd state: (25) PMo1oF2M9EEWm6plLqk1v0EoQrli

ret = client.Write(fd459, []byte("xc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4Cg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) PMo1oF2M9EEWm6plLqk1v0EoQxc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4Cg

fd471 := client.Open("/Y6QS20zabJ.txt", client.O_RDWR|client.O_CREATE)
if(fd471 < 0) {
  panic("open failed")
}


ret = client.Close(fd466)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd464)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd459, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd469)
if(ret != 0) {
  panic("close failed")
}


fd472 := client.Open("/FGjIWf9TaF.txt", client.O_RDWR|client.O_CREATE)
if(fd472 < 0) {
  panic("open failed")
}

//fd state: (86) 8tJsFwlR24FbAD_XmTrcxrKaE4G7uX1yIutSEhRYgOfdy2euScHoGYvKjNscZRB6jA8se2E5FM7MR14xFFgr4F

ret = client.Write(fd463, []byte("bWcGpW5PKpJ0VcEsbqsx_VqS2DewUfToHRXAwF1qj2h9D52Udv9zaSKVzvwreQR4HxIn0cjdUxzygTbDwwAHXiHSdRgSO95O"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (182) 8tJsFwlR24FbAD_XmTrcxrKaE4G7uX1yIutSEhRYgOfdy2euScHoGYvKjNscZRB6jA8se2E5FM7MR14xFFgr4FbWcGpW5PKpJ0VcEsbqsx_VqS2DewUfToHRXAwF1qj2h9D52Udv9zaSKVzvwreQR4HxIn0cjdUxzygTbDwwAHXiHSdRgSO95O

buf, ret = client.Read(fd472, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd472, []byte("rSdmkPqGLLX1DzseJYW4ypzYbgaLfODGPGno2oF55oGtybO3lskgeI3RSr_yT16KAKTSOlSZ784sTltywV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) rSdmkPqGLLX1DzseJYW4ypzYbgaLfODGPGno2oF55oGtybO3lskgeI3RSr_yT16KAKTSOlSZ784sTltywV
//fd state: (0) 

ret = client.Write(fd471, []byte("BblVW7G"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) BblVW7G
//fd state: (83) OkZljU26x5o7ziZhg96vhnKtgFSlvA9fHsa4vRZ3wlRIenjASoL1hNAtwIjqnDAYqTQWLW5_c1nPLrUqP9C

ret = client.Write(fd446, []byte("1IfvY5z07So85ARWWImzeiz4ra8A8Ufuk7TpmNhDT0mD9qUcu4iPq8dvFEbqnmqwc0b2qlOVqzC2uTg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (162) OkZljU26x5o7ziZhg96vhnKtgFSlvA9fHsa4vRZ3wlRIenjASoL1hNAtwIjqnDAYqTQWLW5_c1nPLrUqP9C1IfvY5z07So85ARWWImzeiz4ra8A8Ufuk7TpmNhDT0mD9qUcu4iPq8dvFEbqnmqwc0b2qlOVqzC2uTg

ret = client.Close(fd472)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd470)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd460, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jsPMdeDG1RKXdTXgTwEor0zEcumpM1REFP1hgvhG8rUwunqJlWIs6j_IxNIbJminsvsyMbckUcRj8YbiId9rBnEmTqY") {
  panic("wrong data returned")
}


ret = client.Seek(fd460, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


buf, ret = client.Read(fd455, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd463, 129, client.SEEK_SET)
if(ret != 129) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 129)
  panic("seek failed")
}


ret = client.Seek(fd460, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Close(fd455)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd463, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9D52Udv9zaSKVzvwreQR4HxIn0cjdUxzygTbDwwAHXiHSdRgSO95O") {
  panic("wrong data returned")
}


fd473 := client.Open("/V0gWwEDbM4.txt", client.O_RDWR|client.O_CREATE)
if(fd473 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd446, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd446)
if(ret != 0) {
  panic("close failed")
}

//fd state: (65) PMo1oF2M9EEWm6plLqk1v0EoQxc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4Cg

ret = client.Write(fd459, []byte("HeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (131) PMo1oF2M9EEWm6plLqk1v0EoQxc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4CgHeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpn

fd474 := client.Open("/DBm2gV4Brr.txt", client.O_RDWR|client.O_CREATE)
if(fd474 < 0) {
  panic("open failed")
}


fd475 := client.Open("/mFvxS3bd1N.txt", client.O_RDWR|client.O_CREATE)
if(fd475 < 0) {
  panic("open failed")
}


ret = client.Close(fd474)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd463)
if(ret != 0) {
  panic("close failed")
}

//fd state: (131) PMo1oF2M9EEWm6plLqk1v0EoQxc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4CgHeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpn

ret = client.Write(fd459, []byte("fIz0nNvJzCb0LD7rU20C0brf8x3ulLgKU54"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (166) PMo1oF2M9EEWm6plLqk1v0EoQxc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4CgHeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpnfIz0nNvJzCb0LD7rU20C0brf8x3ulLgKU54

ret = client.Close(fd473)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) WZqHUkaK4jCQKj83j2zrYKcLlCVp_fS12TmakDiP2kJ3aoN77NUtQTpva44CxAci_NHE4jfVcl3Teda33hXOkGJRzQiA9tHY_V5fcV8IoZCVmJr50W_I5dkjcyQ2GwBVWRxhYXGXF8ZDWATL0LhJVviXpyLgHUPbsb1u

ret = client.Write(fd475, []byte("GKNj4OFcylSx_GpEneynClXV3YNOcB3vXvLrKc4kloY4ODCkU1lCK5goyFyr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) GKNj4OFcylSx_GpEneynClXV3YNOcB3vXvLrKc4kloY4ODCkU1lCK5goyFyrxAci_NHE4jfVcl3Teda33hXOkGJRzQiA9tHY_V5fcV8IoZCVmJr50W_I5dkjcyQ2GwBVWRxhYXGXF8ZDWATL0LhJVviXpyLgHUPbsb1u
//fd state: (166) PMo1oF2M9EEWm6plLqk1v0EoQxc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4CgHeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpnfIz0nNvJzCb0LD7rU20C0brf8x3ulLgKU54

ret = client.Write(fd459, []byte("ZikIC8eCN049O6g8uXuQB4Va0zv1ChPHK1jcoRM0mC7Q7GBN5gMGfIFIJQx62sJliIVfmpr2_Gfonf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (244) PMo1oF2M9EEWm6plLqk1v0EoQxc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4CgHeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpnfIz0nNvJzCb0LD7rU20C0brf8x3ulLgKU54ZikIC8eCN049O6g8uXuQB4Va0zv1ChPHK1jcoRM0mC7Q7GBN5gMGfIFIJQx62sJliIVfmpr2_Gfonf
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd460)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd471, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd476 := client.Open("/PLzJpEhWKF.txt", client.O_RDWR|client.O_CREATE)
if(fd476 < 0) {
  panic("open failed")
}


ret = client.Seek(fd471, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


buf, ret = client.Read(fd459, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd476)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd475)
if(ret != 0) {
  panic("close failed")
}

//fd state: (244) PMo1oF2M9EEWm6plLqk1v0EoQxc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4CgHeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpnfIz0nNvJzCb0LD7rU20C0brf8x3ulLgKU54ZikIC8eCN049O6g8uXuQB4Va0zv1ChPHK1jcoRM0mC7Q7GBN5gMGfIFIJQx62sJliIVfmpr2_Gfonf

ret = client.Write(fd459, []byte("pHUqEx0xjBbdYcj8MudzGgYlc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (269) PMo1oF2M9EEWm6plLqk1v0EoQxc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4CgHeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpnfIz0nNvJzCb0LD7rU20C0brf8x3ulLgKU54ZikIC8eCN049O6g8uXuQB4Va0zv1ChPHK1jcoRM0mC7Q7GBN5gMGfIFIJQx62sJliIVfmpr2_GfonfpHUqEx0xjBbdYcj8MudzGgYlc

buf, ret = client.Read(fd459, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (5) BblVW7G

ret = client.Write(fd471, []byte("DSqNPY3P0YDOqe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) BblVWDSqNPY3P0YDOqe

ret = client.Close(fd459)
if(ret != 0) {
  panic("close failed")
}

//fd state: (19) BblVWDSqNPY3P0YDOqe

ret = client.Write(fd471, []byte("SKi58ltCmr_yRONs4Ab2Cz56FrZf9M8EG0QPvX55vdr9ebcB8a0LcwQcNkggg4TdaVdVOrXn0tCqUYNrEMglTV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) BblVWDSqNPY3P0YDOqeSKi58ltCmr_yRONs4Ab2Cz56FrZf9M8EG0QPvX55vdr9ebcB8a0LcwQcNkggg4TdaVdVOrXn0tCqUYNrEMglTV

ret = client.Seek(fd471, 99, client.SEEK_SET)
if(ret != 99) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 99)
  panic("seek failed")
}


fd477 := client.Open("/YYMVsNBZSh.txt", client.O_RDWR|client.O_CREATE)
if(fd477 < 0) {
  panic("open failed")
}


ret = client.Close(fd477)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd471, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


ret = client.Close(fd471)
if(ret != 0) {
  panic("close failed")
}


fd478 := client.Open("/YVbJEFJUJd.txt", client.O_RDWR|client.O_CREATE)
if(fd478 < 0) {
  panic("open failed")
}


fd479 := client.Open("/hCnVwbgHwm.txt", client.O_RDWR|client.O_CREATE)
if(fd479 < 0) {
  panic("open failed")
}


ret = client.Close(fd479)
if(ret != 0) {
  panic("close failed")
}


fd480 := client.Open("/iHqPqm1btq.txt", client.O_RDWR|client.O_CREATE)
if(fd480 < 0) {
  panic("open failed")
}


ret = client.Seek(fd480, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd480, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd480, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd480)
if(ret != 0) {
  panic("close failed")
}


fd481 := client.Open("/b2Gr0PrgBp.txt", client.O_RDWR|client.O_CREATE)
if(fd481 < 0) {
  panic("open failed")
}


fd482 := client.Open("/GI8sQv42jD.txt", client.O_RDWR|client.O_CREATE)
if(fd482 < 0) {
  panic("open failed")
}


ret = client.Close(fd482)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd478)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd481, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd481, []byte("cMZoKQUNI2mneFONtH6tXoVVJB8o0FRNG6Ryy9JIldO3veNcgEmXulYiY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) cMZoKQUNI2mneFONtH6tXoVVJB8o0FRNG6Ryy9JIldO3veNcgEmXulYiY
//fd state: (57) cMZoKQUNI2mneFONtH6tXoVVJB8o0FRNG6Ryy9JIldO3veNcgEmXulYiY

ret = client.Write(fd481, []byte("6mGhkUwHR7fV0vunIUzPZDE7HqJESel8JipUIRouqFCzsuYD1Ow0hXVJRjPo1Om4Nz2KoF7dy4yf8bLIB833GLtuRkmqEG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) cMZoKQUNI2mneFONtH6tXoVVJB8o0FRNG6Ryy9JIldO3veNcgEmXulYiY6mGhkUwHR7fV0vunIUzPZDE7HqJESel8JipUIRouqFCzsuYD1Ow0hXVJRjPo1Om4Nz2KoF7dy4yf8bLIB833GLtuRkmqEG

ret = client.Seek(fd481, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}


fd483 := client.Open("/hCnVwbgHwm.txt", client.O_RDWR|client.O_CREATE)
if(fd483 < 0) {
  panic("open failed")
}


fd484 := client.Open("/9g7Nr3wNRE.txt", client.O_RDWR|client.O_CREATE)
if(fd484 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd483, []byte("sBgF7CQKZHSibsJ7wzTTEzOeQVBvEOzreL3a0P"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) sBgF7CQKZHSibsJ7wzTTEzOeQVBvEOzreL3a0P

fd485 := client.Open("/aWiyrgOINA.txt", client.O_RDWR|client.O_CREATE)
if(fd485 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd484, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd485, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd484, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd481)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd484, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd485, []byte("owau6G368X7UsbPB5QyprSyZ8oQ9m8ZwgPYkf9QCkWDipERDel2cABQdEuPLetAscml"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) owau6G368X7UsbPB5QyprSyZ8oQ9m8ZwgPYkf9QCkWDipERDel2cABQdEuPLetAscml

ret = client.Close(fd483)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd484, []byte("w6dQu_3UuiXEOZn2qzOoUnU2tID8jyGGXno5napHXwuWbmlXAg4kVb4H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) w6dQu_3UuiXEOZn2qzOoUnU2tID8jyGGXno5napHXwuWbmlXAg4kVb4H
//fd state: (56) w6dQu_3UuiXEOZn2qzOoUnU2tID8jyGGXno5napHXwuWbmlXAg4kVb4H

ret = client.Write(fd484, []byte("QqsPaS9XvlzI0pHdLAL3Qh0GP0QD3ZJllUo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) w6dQu_3UuiXEOZn2qzOoUnU2tID8jyGGXno5napHXwuWbmlXAg4kVb4HQqsPaS9XvlzI0pHdLAL3Qh0GP0QD3ZJllUo

ret = client.Close(fd485)
if(ret != 0) {
  panic("close failed")
}


fd486 := client.Open("/zvfpB2u_jD.txt", client.O_RDWR|client.O_CREATE)
if(fd486 < 0) {
  panic("open failed")
}


ret = client.Close(fd484)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd486)
if(ret != 0) {
  panic("close failed")
}


fd487 := client.Open("/tbLiMuf1nh.txt", client.O_RDWR|client.O_CREATE)
if(fd487 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd487, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd487)
if(ret != 0) {
  panic("close failed")
}


fd488 := client.Open("/0WU4yNMJLm.txt", client.O_RDWR|client.O_CREATE)
if(fd488 < 0) {
  panic("open failed")
}


ret = client.Seek(fd488, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd489 := client.Open("/sG15GUiEku.txt", client.O_RDWR|client.O_CREATE)
if(fd489 < 0) {
  panic("open failed")
}


ret = client.Close(fd489)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd488, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd488, []byte("EiH1yLauEKscNzOetrMCzuGn1k8Emq7ZbHYhnBL4uK0shXGFnVPGYIikVD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) EiH1yLauEKscNzOetrMCzuGn1k8Emq7ZbHYhnBL4uK0shXGFnVPGYIikVD
//fd state: (58) EiH1yLauEKscNzOetrMCzuGn1k8Emq7ZbHYhnBL4uK0shXGFnVPGYIikVD

ret = client.Write(fd488, []byte("c5gTgsYrZQTRkErSEGKtIwJ9BhZ01Vbse_gW8IbaXdpe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) EiH1yLauEKscNzOetrMCzuGn1k8Emq7ZbHYhnBL4uK0shXGFnVPGYIikVDc5gTgsYrZQTRkErSEGKtIwJ9BhZ01Vbse_gW8IbaXdpe

fd490 := client.Open("/JzZk_6uCpc.txt", client.O_RDWR|client.O_CREATE)
if(fd490 < 0) {
  panic("open failed")
}


ret = client.Seek(fd488, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd490, []byte("5lh_gujNIB_yLPM8cNBRhMqqG3tPCRXwwkEduQA_kt4qrj865nUJRBMCtOtKfUSDDca5pzY2_ITOrmGfYFD8j"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) 5lh_gujNIB_yLPM8cNBRhMqqG3tPCRXwwkEduQA_kt4qrj865nUJRBMCtOtKfUSDDca5pzY2_ITOrmGfYFD8j

ret = client.Seek(fd490, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}

//fd state: (58) 5lh_gujNIB_yLPM8cNBRhMqqG3tPCRXwwkEduQA_kt4qrj865nUJRBMCtOtKfUSDDca5pzY2_ITOrmGfYFD8j

ret = client.Write(fd490, []byte("mdgnFI5us8lV2eFkwdmvLw05Z8HNUL7sOyTLBMC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) 5lh_gujNIB_yLPM8cNBRhMqqG3tPCRXwwkEduQA_kt4qrj865nUJRBMCtOmdgnFI5us8lV2eFkwdmvLw05Z8HNUL7sOyTLBMC

fd491 := client.Open("/q6rX79dbpd.txt", client.O_RDWR|client.O_CREATE)
if(fd491 < 0) {
  panic("open failed")
}


ret = client.Seek(fd488, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd491)
if(ret != 0) {
  panic("close failed")
}


fd492 := client.Open("/bRoENni7CO.txt", client.O_RDWR|client.O_CREATE)
if(fd492 < 0) {
  panic("open failed")
}


ret = client.Seek(fd490, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}


ret = client.Seek(fd488, 87, client.SEEK_SET)
if(ret != 87) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 87)
  panic("seek failed")
}


buf, ret = client.Read(fd488, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Vbse_gW8IbaXdpe") {
  panic("wrong data returned")
}


ret = client.Close(fd488)
if(ret != 0) {
  panic("close failed")
}

//fd state: (72) 5lh_gujNIB_yLPM8cNBRhMqqG3tPCRXwwkEduQA_kt4qrj865nUJRBMCtOmdgnFI5us8lV2eFkwdmvLw05Z8HNUL7sOyTLBMC

ret = client.Write(fd490, []byte("ZwrFoNmw15qnnBcOSh8HHJ3Vno5qigZkilnNPA9slsFXfjyA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) 5lh_gujNIB_yLPM8cNBRhMqqG3tPCRXwwkEduQA_kt4qrj865nUJRBMCtOmdgnFI5us8lV2eZwrFoNmw15qnnBcOSh8HHJ3Vno5qigZkilnNPA9slsFXfjyA

fd493 := client.Open("/kWlBQtlrUQ.txt", client.O_RDWR|client.O_CREATE)
if(fd493 < 0) {
  panic("open failed")
}


ret = client.Seek(fd493, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


ret = client.Close(fd490)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd493, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oxlt1nZMaCdcRSmE3MkzfqPAXF1J5Sbv2jHCPexDR97PNaKH5hqMRzd8RUedVksHf23UNuD") {
  panic("wrong data returned")
}


ret = client.Close(fd492)
if(ret != 0) {
  panic("close failed")
}

//fd state: (98) 5iIMrEIpKj0czDjTGp9weRYWPxqoxlt1nZMaCdcRSmE3MkzfqPAXF1J5Sbv2jHCPexDR97PNaKH5hqMRzd8RUedVksHf23UNuD

ret = client.Write(fd493, []byte("HOWUXgGox1fpfIEr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) 5iIMrEIpKj0czDjTGp9weRYWPxqoxlt1nZMaCdcRSmE3MkzfqPAXF1J5Sbv2jHCPexDR97PNaKH5hqMRzd8RUedVksHf23UNuDHOWUXgGox1fpfIEr

buf, ret = client.Read(fd493, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd494 := client.Open("/w2FXc3l6HU.txt", client.O_RDWR|client.O_CREATE)
if(fd494 < 0) {
  panic("open failed")
}


ret = client.Seek(fd493, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}


ret = client.Seek(fd493, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


buf, ret = client.Read(fd494, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "HjCrlg8CbeArGDTwFO") {
  panic("wrong data returned")
}


ret = client.Seek(fd494, 102, client.SEEK_SET)
if(ret != 102) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 102)
  panic("seek failed")
}


ret = client.Seek(fd494, 183, client.SEEK_SET)
if(ret != 183) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 183)
  panic("seek failed")
}


ret = client.Seek(fd493, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 5iIMrEIpKj0czDjTGp9weRYWPxqoxlt1nZMaCdcRSmE3MkzfqPAXF1J5Sbv2jHCPexDR97PNaKH5hqMRzd8RUedVksHf23UNuDHOWUXgGox1fpfIEr

ret = client.Write(fd493, []byte("Svdc8oNdhEHpEN1syQgngq0AtLNGruKP8fFO0coZB1GixvRFj3Fiv6I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) Svdc8oNdhEHpEN1syQgngq0AtLNGruKP8fFO0coZB1GixvRFj3Fiv6I5Sbv2jHCPexDR97PNaKH5hqMRzd8RUedVksHf23UNuDHOWUXgGox1fpfIEr
//fd state: (183) HjCrlg8CbeArGDTwFObXD6JAt8bsCztnzRF3tsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqt5R7KHop0EOZE5BKxcjKy9EKiN8ixcp1zPN12L5q7OHcP9OnofhUG23y8dikApbTExBATSHC2_bU4sitD5Kgcl43F6TDkIP5Vx3wP6JOKpHb5Ia

ret = client.Write(fd494, []byte("9PFX0qwMtu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (193) HjCrlg8CbeArGDTwFObXD6JAt8bsCztnzRF3tsQvEihANcMWOBqrE5St56Q8V4_fSu5NkKak5MFqt5R7KHop0EOZE5BKxcjKy9EKiN8ixcp1zPN12L5q7OHcP9OnofhUG23y8dikApbTExBATSHC2_bU4sitD5Kgcl43F6TDkIP5Vx3wP6JOKpH9PFX0qwMtu

ret = client.Close(fd494)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd493, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5Sbv2jHCPexDR97PNaKH5hqMRzd8RUedVksHf23UNuDHOWUXgGox1fpfIEr") {
  panic("wrong data returned")
}


ret = client.Seek(fd493, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


fd495 := client.Open("/YJ0bjUslc0.txt", client.O_RDWR|client.O_CREATE)
if(fd495 < 0) {
  panic("open failed")
}

//fd state: (81) Svdc8oNdhEHpEN1syQgngq0AtLNGruKP8fFO0coZB1GixvRFj3Fiv6I5Sbv2jHCPexDR97PNaKH5hqMRzd8RUedVksHf23UNuDHOWUXgGox1fpfIEr

ret = client.Write(fd493, []byte("E26OnmrztUQiaVzCrXVBC3vXqV5zEAMXuBKJDq2jAIQYVNTYsVb3wUu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) Svdc8oNdhEHpEN1syQgngq0AtLNGruKP8fFO0coZB1GixvRFj3Fiv6I5Sbv2jHCPexDR97PNaKH5hqMRzE26OnmrztUQiaVzCrXVBC3vXqV5zEAMXuBKJDq2jAIQYVNTYsVb3wUu
//fd state: (136) Svdc8oNdhEHpEN1syQgngq0AtLNGruKP8fFO0coZB1GixvRFj3Fiv6I5Sbv2jHCPexDR97PNaKH5hqMRzE26OnmrztUQiaVzCrXVBC3vXqV5zEAMXuBKJDq2jAIQYVNTYsVb3wUu

ret = client.Write(fd493, []byte("PthCeJEQkdeC0vgpx5aJImL8In1yssxSOwejz5_REGou2Uvusx9HLuhq27q_iMyaqJizKBnRweuQwRyAkVtMRmmw1n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (226) Svdc8oNdhEHpEN1syQgngq0AtLNGruKP8fFO0coZB1GixvRFj3Fiv6I5Sbv2jHCPexDR97PNaKH5hqMRzE26OnmrztUQiaVzCrXVBC3vXqV5zEAMXuBKJDq2jAIQYVNTYsVb3wUuPthCeJEQkdeC0vgpx5aJImL8In1yssxSOwejz5_REGou2Uvusx9HLuhq27q_iMyaqJizKBnRweuQwRyAkVtMRmmw1n

fd496 := client.Open("/6JfRSSuVna.txt", client.O_RDWR|client.O_CREATE)
if(fd496 < 0) {
  panic("open failed")
}


fd497 := client.Open("/OmH6ue7qfz.txt", client.O_RDWR|client.O_CREATE)
if(fd497 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd493, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd498 := client.Open("/xxTCIY1dXp.txt", client.O_RDWR|client.O_CREATE)
if(fd498 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd498)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd493, 156, client.SEEK_SET)
if(ret != 156) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 156)
  panic("seek failed")
}

//fd state: (0) y0X7vxuPZnVW4CFvyHzD43W_Mj2kWxY9oylLJHagx6V0S

ret = client.Write(fd496, []byte("MxdC21RfQtyxCF7t1HkKMM0_1dpF2lCxPNUGxMU_Vxnqg5DcdtJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) MxdC21RfQtyxCF7t1HkKMM0_1dpF2lCxPNUGxMU_Vxnqg5DcdtJ

ret = client.Seek(fd493, 203, client.SEEK_SET)
if(ret != 203) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 203)
  panic("seek failed")
}


ret = client.Seek(fd497, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


fd499 := client.Open("/6JfRSSuVna.txt", client.O_RDWR|client.O_CREATE)
if(fd499 < 0) {
  panic("open failed")
}


ret = client.Close(fd495)
if(ret != 0) {
  panic("close failed")
}


fd500 := client.Open("/PecpTE5aC2.txt", client.O_RDWR|client.O_CREATE)
if(fd500 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd499, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MxdC21RfQtyxCF7t1HkKMM0_1dpF2lCxPNUGxMU_Vxnqg5DcdtJ") {
  panic("wrong data returned")
}


ret = client.Close(fd500)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd496)
if(ret != 0) {
  panic("close failed")
}


fd501 := client.Open("/d5bTRIS3Qm.txt", client.O_RDWR|client.O_CREATE)
if(fd501 < 0) {
  panic("open failed")
}

//fd state: (54) 0NUSBWlxgvZCB_gsv7K2iMnSw1tkb05OxzISIxbdtrhpQlfiGVPO618DIYFIIetS23QVihBL

ret = client.Write(fd497, []byte("oL2jGetMZ5ScwoKdiq_BbepVxUNRTCYQd2nfsCIENMtS32An12"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) 0NUSBWlxgvZCB_gsv7K2iMnSw1tkb05OxzISIxbdtrhpQlfiGVPO61oL2jGetMZ5ScwoKdiq_BbepVxUNRTCYQd2nfsCIENMtS32An12

buf, ret = client.Read(fd501, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd501, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (104) 0NUSBWlxgvZCB_gsv7K2iMnSw1tkb05OxzISIxbdtrhpQlfiGVPO61oL2jGetMZ5ScwoKdiq_BbepVxUNRTCYQd2nfsCIENMtS32An12

ret = client.Write(fd497, []byte("d295rxNcFiuFQSBHkTtYFdSidEA3t1BqJvhFuyTeP3_Pl9DgZrW4wVPxTpzZLCJSuopJrrGY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (176) 0NUSBWlxgvZCB_gsv7K2iMnSw1tkb05OxzISIxbdtrhpQlfiGVPO61oL2jGetMZ5ScwoKdiq_BbepVxUNRTCYQd2nfsCIENMtS32An12d295rxNcFiuFQSBHkTtYFdSidEA3t1BqJvhFuyTeP3_Pl9DgZrW4wVPxTpzZLCJSuopJrrGY

ret = client.Seek(fd493, 221, client.SEEK_SET)
if(ret != 221) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 221)
  panic("seek failed")
}


ret = client.Seek(fd501, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd502 := client.Open("/CR4qfzt_Wc.txt", client.O_RDWR|client.O_CREATE)
if(fd502 < 0) {
  panic("open failed")
}


fd503 := client.Open("/f4k1qvMKPz.txt", client.O_RDWR|client.O_CREATE)
if(fd503 < 0) {
  panic("open failed")
}


ret = client.Seek(fd499, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


fd504 := client.Open("/cPJ5D3QIxK.txt", client.O_RDWR|client.O_CREATE)
if(fd504 < 0) {
  panic("open failed")
}

//fd state: (11) MxdC21RfQtyxCF7t1HkKMM0_1dpF2lCxPNUGxMU_Vxnqg5DcdtJ

ret = client.Write(fd499, []byte("whsNlD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) MxdC21RfQtywhsNlDHkKMM0_1dpF2lCxPNUGxMU_Vxnqg5DcdtJ
//fd state: (0) 

ret = client.Write(fd503, []byte("FsPRS6rmY1YBosCBzjSt9VZGjpT8X_fEmVrBpjCQYpP0QF2Waq6t9Yuy013FjjIQU5M5N5qbbzzpLBz9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) FsPRS6rmY1YBosCBzjSt9VZGjpT8X_fEmVrBpjCQYpP0QF2Waq6t9Yuy013FjjIQU5M5N5qbbzzpLBz9
//fd state: (0) cbrWhsPG1xyWR1F7JVvML2HVNjnDiF1oKSB3H3cx8zmSYE46FtpOIGl0MBh

ret = client.Write(fd502, []byte("HQGr4wnIatTQeNB42hPav0X_JUiSeA48l4rDA5FZBw5uevZqHPS_ydlV_jUM7ozljBW5TALIAq2scikovPl5_BrwY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) HQGr4wnIatTQeNB42hPav0X_JUiSeA48l4rDA5FZBw5uevZqHPS_ydlV_jUM7ozljBW5TALIAq2scikovPl5_BrwY

fd505 := client.Open("/QIPEg8dwnM.txt", client.O_RDWR|client.O_CREATE)
if(fd505 < 0) {
  panic("open failed")
}


ret = client.Close(fd497)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd502, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd505, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd493)
if(ret != 0) {
  panic("close failed")
}


fd506 := client.Open("/xxTCIY1dXp.txt", client.O_RDWR|client.O_CREATE)
if(fd506 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd505, []byte("ACQ2h855f2Sr8Q8sHm3yPIKoKL002tq0dQTsrGxus6tsHHdMBridBaDn8fYln6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) ACQ2h855f2Sr8Q8sHm3yPIKoKL002tq0dQTsrGxus6tsHHdMBridBaDn8fYln6

fd507 := client.Open("/yY3JR0W531.txt", client.O_RDWR|client.O_CREATE)
if(fd507 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd507, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd501, []byte("17hWpe6DK73XI_Gcmz3xtFMrwOMBJLfIM3GM1R31WXzctPfLVVb_QGJb6NGR4mdqMEgIa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) 17hWpe6DK73XI_Gcmz3xtFMrwOMBJLfIM3GM1R31WXzctPfLVVb_QGJb6NGR4mdqMEgIa

buf, ret = client.Read(fd507, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd507, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd508 := client.Open("/90hEKufnYM.txt", client.O_RDWR|client.O_CREATE)
if(fd508 < 0) {
  panic("open failed")
}


ret = client.Seek(fd506, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd509 := client.Open("/VTasqgKgMi.txt", client.O_RDWR|client.O_CREATE)
if(fd509 < 0) {
  panic("open failed")
}


ret = client.Close(fd504)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd501, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}

//fd state: (80) FsPRS6rmY1YBosCBzjSt9VZGjpT8X_fEmVrBpjCQYpP0QF2Waq6t9Yuy013FjjIQU5M5N5qbbzzpLBz9

ret = client.Write(fd503, []byte("vzOVdOz84t1c9rzifLDOCidAWb1rE7zy3noNc28gPKA9VW4JF4OWvbKWCdctULNlCgQldvdk53zcYCQaAsZNFmsiIK52PCY6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (176) FsPRS6rmY1YBosCBzjSt9VZGjpT8X_fEmVrBpjCQYpP0QF2Waq6t9Yuy013FjjIQU5M5N5qbbzzpLBz9vzOVdOz84t1c9rzifLDOCidAWb1rE7zy3noNc28gPKA9VW4JF4OWvbKWCdctULNlCgQldvdk53zcYCQaAsZNFmsiIK52PCY6

buf, ret = client.Read(fd501, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "M3GM1R31WXzc") {
  panic("wrong data returned")
}

//fd state: (176) FsPRS6rmY1YBosCBzjSt9VZGjpT8X_fEmVrBpjCQYpP0QF2Waq6t9Yuy013FjjIQU5M5N5qbbzzpLBz9vzOVdOz84t1c9rzifLDOCidAWb1rE7zy3noNc28gPKA9VW4JF4OWvbKWCdctULNlCgQldvdk53zcYCQaAsZNFmsiIK52PCY6

ret = client.Write(fd503, []byte("kfYOIxViwDdi8e"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (190) FsPRS6rmY1YBosCBzjSt9VZGjpT8X_fEmVrBpjCQYpP0QF2Waq6t9Yuy013FjjIQU5M5N5qbbzzpLBz9vzOVdOz84t1c9rzifLDOCidAWb1rE7zy3noNc28gPKA9VW4JF4OWvbKWCdctULNlCgQldvdk53zcYCQaAsZNFmsiIK52PCY6kfYOIxViwDdi8e

ret = client.Close(fd505)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) mpJLASTrVJSdahxAZHmQKouYtt15M8K2H6vLugFHj9hNCVH4XOnR1sU8RbOeMscfhzdfD3nKpuFPWEMyIxk_vPLqNfLrVyCVpyyd

ret = client.Write(fd508, []byte("rpiIiZukGBYsz1A5cQwgNenwVW8Y8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) rpiIiZukGBYsz1A5cQwgNenwVW8Y88K2H6vLugFHj9hNCVH4XOnR1sU8RbOeMscfhzdfD3nKpuFPWEMyIxk_vPLqNfLrVyCVpyyd

ret = client.Seek(fd507, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd507, []byte("9l5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) 9l5

ret = client.Seek(fd501, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd507)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd508, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8K2H6vLugFHj9hNCVH4XOnR1") {
  panic("wrong data returned")
}


ret = client.Seek(fd508, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


ret = client.Close(fd501)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd509, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd499)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd508, 100, client.SEEK_SET)
if(ret != 100) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 100)
  panic("seek failed")
}


ret = client.Close(fd502)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd509, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd509, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd510 := client.Open("/FBwyYmw985.txt", client.O_RDWR|client.O_CREATE)
if(fd510 < 0) {
  panic("open failed")
}


ret = client.Seek(fd508, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


fd511 := client.Open("/CR4qfzt_Wc.txt", client.O_RDWR|client.O_CREATE)
if(fd511 < 0) {
  panic("open failed")
}

//fd state: (190) FsPRS6rmY1YBosCBzjSt9VZGjpT8X_fEmVrBpjCQYpP0QF2Waq6t9Yuy013FjjIQU5M5N5qbbzzpLBz9vzOVdOz84t1c9rzifLDOCidAWb1rE7zy3noNc28gPKA9VW4JF4OWvbKWCdctULNlCgQldvdk53zcYCQaAsZNFmsiIK52PCY6kfYOIxViwDdi8e

ret = client.Write(fd503, []byte("jjJHjOwNac64BHNGvThKOTWwP4VtP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (219) FsPRS6rmY1YBosCBzjSt9VZGjpT8X_fEmVrBpjCQYpP0QF2Waq6t9Yuy013FjjIQU5M5N5qbbzzpLBz9vzOVdOz84t1c9rzifLDOCidAWb1rE7zy3noNc28gPKA9VW4JF4OWvbKWCdctULNlCgQldvdk53zcYCQaAsZNFmsiIK52PCY6kfYOIxViwDdi8ejjJHjOwNac64BHNGvThKOTWwP4VtP

buf, ret = client.Read(fd509, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) HQGr4wnIatTQeNB42hPav0X_JUiSeA48l4rDA5FZBw5uevZqHPS_ydlV_jUM7ozljBW5TALIAq2scikovPl5_BrwY

ret = client.Write(fd511, []byte("ZE68XXff"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) ZE68XXffatTQeNB42hPav0X_JUiSeA48l4rDA5FZBw5uevZqHPS_ydlV_jUM7ozljBW5TALIAq2scikovPl5_BrwY

ret = client.Seek(fd506, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd512 := client.Open("/otTHquN_0Q.txt", client.O_RDWR|client.O_CREATE)
if(fd512 < 0) {
  panic("open failed")
}


fd513 := client.Open("/AVIwDYxhCx.txt", client.O_RDWR|client.O_CREATE)
if(fd513 < 0) {
  panic("open failed")
}


ret = client.Seek(fd512, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd509)
if(ret != 0) {
  panic("close failed")
}


fd514 := client.Open("/Crr1MV2zwA.txt", client.O_RDWR|client.O_CREATE)
if(fd514 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd512, []byte("zYyJ6AFWIkkae6fRrdH2_nkGZ3dkzWSTf7U1Nnmw7rTdEIF6uzMmMJooWLpOMcglqxY9OmLtIrvXp1zvdONuXK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) zYyJ6AFWIkkae6fRrdH2_nkGZ3dkzWSTf7U1Nnmw7rTdEIF6uzMmMJooWLpOMcglqxY9OmLtIrvXp1zvdONuXK

buf, ret = client.Read(fd514, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd514, []byte("sIKO8uPjOaEzDyY4hITnesh9R8mEO8gDfrqferM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) sIKO8uPjOaEzDyY4hITnesh9R8mEO8gDfrqferM

ret = client.Seek(fd513, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd506, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd512)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd508)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd506, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd503)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd513, []byte("ZchtiRMg6mcTqjjgJu55joGSsCIIzPVT95A3PiiQnz4JM1wnWsCDrf0g0FfLSTNl_emg50ioaDgAYKpO7q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) ZchtiRMg6mcTqjjgJu55joGSsCIIzPVT95A3PiiQnz4JM1wnWsCDrf0g0FfLSTNl_emg50ioaDgAYKpO7q

ret = client.Close(fd514)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd513, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd511, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "atTQeNB42") {
  panic("wrong data returned")
}


ret = client.Seek(fd506, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd510, []byte("NpET0RHOjmK2VZDzYfQ5ySY3xAC8gdDFakTYE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) NpET0RHOjmK2VZDzYfQ5ySY3xAC8gdDFakTYE

fd515 := client.Open("/YJ0bjUslc0.txt", client.O_RDWR|client.O_CREATE)
if(fd515 < 0) {
  panic("open failed")
}


ret = client.Close(fd506)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd511, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hPav0X_JUiS") {
  panic("wrong data returned")
}


fd516 := client.Open("/rfd_kqAe9A.txt", client.O_RDWR|client.O_CREATE)
if(fd516 < 0) {
  panic("open failed")
}


ret = client.Seek(fd515, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd516, []byte("slA9mkGPCDvzhghqqdbnv8RPNocZ00xshEBc2pwnXkRWxvX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) slA9mkGPCDvzhghqqdbnv8RPNocZ00xshEBc2pwnXkRWxvX

buf, ret = client.Read(fd510, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd511, 89, client.SEEK_SET)
if(ret != 89) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 89)
  panic("seek failed")
}

//fd state: (82) ZchtiRMg6mcTqjjgJu55joGSsCIIzPVT95A3PiiQnz4JM1wnWsCDrf0g0FfLSTNl_emg50ioaDgAYKpO7q

ret = client.Write(fd513, []byte("dKwd_AYZ4l1RyBmsST18Cfd4oGivoJIYSMCSTB04K94Rt8tv6nLzyPYlSklc1oJ1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (146) ZchtiRMg6mcTqjjgJu55joGSsCIIzPVT95A3PiiQnz4JM1wnWsCDrf0g0FfLSTNl_emg50ioaDgAYKpO7qdKwd_AYZ4l1RyBmsST18Cfd4oGivoJIYSMCSTB04K94Rt8tv6nLzyPYlSklc1oJ1
//fd state: (89) ZE68XXffatTQeNB42hPav0X_JUiSeA48l4rDA5FZBw5uevZqHPS_ydlV_jUM7ozljBW5TALIAq2scikovPl5_BrwY

ret = client.Write(fd511, []byte("ea"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) ZE68XXffatTQeNB42hPav0X_JUiSeA48l4rDA5FZBw5uevZqHPS_ydlV_jUM7ozljBW5TALIAq2scikovPl5_BrwYea

ret = client.Close(fd515)
if(ret != 0) {
  panic("close failed")
}


fd517 := client.Open("/4k_5BxY9TN.txt", client.O_RDWR|client.O_CREATE)
if(fd517 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd510, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd518 := client.Open("/ct9OAepmrS.txt", client.O_RDWR|client.O_CREATE)
if(fd518 < 0) {
  panic("open failed")
}

//fd state: (91) ZE68XXffatTQeNB42hPav0X_JUiSeA48l4rDA5FZBw5uevZqHPS_ydlV_jUM7ozljBW5TALIAq2scikovPl5_BrwYea

ret = client.Write(fd511, []byte("fCqTPaiarXqflyjL6H32qChljQzV5NhLkKRy5uwbBJDiwtp3to8A7uo9RZodIqjJMF9X_VF96GZraLB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (170) ZE68XXffatTQeNB42hPav0X_JUiSeA48l4rDA5FZBw5uevZqHPS_ydlV_jUM7ozljBW5TALIAq2scikovPl5_BrwYeafCqTPaiarXqflyjL6H32qChljQzV5NhLkKRy5uwbBJDiwtp3to8A7uo9RZodIqjJMF9X_VF96GZraLB

buf, ret = client.Read(fd513, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd510, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd511, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}


buf, ret = client.Read(fd510, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd519 := client.Open("/gA71a_GvUv.txt", client.O_RDWR|client.O_CREATE)
if(fd519 < 0) {
  panic("open failed")
}

//fd state: (47) slA9mkGPCDvzhghqqdbnv8RPNocZ00xshEBc2pwnXkRWxvX

ret = client.Write(fd516, []byte("_SgSW_TBmq2O5CI8W0DXUVMmzd2cYb06Jgtn4XJ107scy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) slA9mkGPCDvzhghqqdbnv8RPNocZ00xshEBc2pwnXkRWxvX_SgSW_TBmq2O5CI8W0DXUVMmzd2cYb06Jgtn4XJ107scy

ret = client.Seek(fd511, 141, client.SEEK_SET)
if(ret != 141) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 141)
  panic("seek failed")
}


fd520 := client.Open("/ebgGYWPeVA.txt", client.O_RDWR|client.O_CREATE)
if(fd520 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd520, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Zh6CW7PV1FV8y8ajXx25Aa7l") {
  panic("wrong data returned")
}


fd521 := client.Open("/x7r7aEPrq5.txt", client.O_RDWR|client.O_CREATE)
if(fd521 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd513, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd519)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd517, []byte("nf519OXTMDZY0dv5VyBkscG4l04ejwmZF0bBb9oMDyJNFe5AYqbIFZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) nf519OXTMDZY0dv5VyBkscG4l04ejwmZF0bBb9oMDyJNFe5AYqbIFZ

ret = client.Seek(fd516, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


fd522 := client.Open("/V0gWwEDbM4.txt", client.O_RDWR|client.O_CREATE)
if(fd522 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd522, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (146) ZchtiRMg6mcTqjjgJu55joGSsCIIzPVT95A3PiiQnz4JM1wnWsCDrf0g0FfLSTNl_emg50ioaDgAYKpO7qdKwd_AYZ4l1RyBmsST18Cfd4oGivoJIYSMCSTB04K94Rt8tv6nLzyPYlSklc1oJ1

ret = client.Write(fd513, []byte("FREO4nGA4xOqcW0Hi0x"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (165) ZchtiRMg6mcTqjjgJu55joGSsCIIzPVT95A3PiiQnz4JM1wnWsCDrf0g0FfLSTNl_emg50ioaDgAYKpO7qdKwd_AYZ4l1RyBmsST18Cfd4oGivoJIYSMCSTB04K94Rt8tv6nLzyPYlSklc1oJ1FREO4nGA4xOqcW0Hi0x

fd523 := client.Open("/WqvRmW3xO_.txt", client.O_RDWR|client.O_CREATE)
if(fd523 < 0) {
  panic("open failed")
}


ret = client.Seek(fd518, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (54) nf519OXTMDZY0dv5VyBkscG4l04ejwmZF0bBb9oMDyJNFe5AYqbIFZ

ret = client.Write(fd517, []byte("HfP4Re8TEVjtTsnpJcqZZ7DfyGLxdXZYxnaQ2Ucxs6OlXSYYE7_ClG9sHiIYHqpaS7Eof"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) nf519OXTMDZY0dv5VyBkscG4l04ejwmZF0bBb9oMDyJNFe5AYqbIFZHfP4Re8TEVjtTsnpJcqZZ7DfyGLxdXZYxnaQ2Ucxs6OlXSYYE7_ClG9sHiIYHqpaS7Eof

ret = client.Close(fd516)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd518, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd524 := client.Open("/0i3NHfYg9J.txt", client.O_RDWR|client.O_CREATE)
if(fd524 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd523, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd517, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd520)
if(ret != 0) {
  panic("close failed")
}


fd525 := client.Open("/PS_GRjWAR1.txt", client.O_RDWR|client.O_CREATE)
if(fd525 < 0) {
  panic("open failed")
}


ret = client.Seek(fd525, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Close(fd510)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd511, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8A7uo9RZodIqjJMF9X_VF96GZraLB") {
  panic("wrong data returned")
}


fd526 := client.Open("/_vdpUCIYrP.txt", client.O_RDWR|client.O_CREATE)
if(fd526 < 0) {
  panic("open failed")
}


ret = client.Seek(fd521, 103, client.SEEK_SET)
if(ret != 103) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 103)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd523, []byte("6FS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) 6FS

ret = client.Close(fd518)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd526)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd513, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd527 := client.Open("/fhDLD2y7yr.txt", client.O_RDWR|client.O_CREATE)
if(fd527 < 0) {
  panic("open failed")
}


fd528 := client.Open("/Crr1MV2zwA.txt", client.O_RDWR|client.O_CREATE)
if(fd528 < 0) {
  panic("open failed")
}


ret = client.Close(fd523)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd525)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd527, []byte("qrfa5gWgDlSO4RWqHjwD4u_LKymwLsRXoHpKxr2I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) qrfa5gWgDlSO4RWqHjwD4u_LKymwLsRXoHpKxr2I

ret = client.Close(fd511)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd528, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "sIKO8uPjOaEzDyY4hITnesh9R8mEO8gDfrqferM") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd521, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7WQpXRqakAAyf3S") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd513, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (40) qrfa5gWgDlSO4RWqHjwD4u_LKymwLsRXoHpKxr2I

ret = client.Write(fd527, []byte("Q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) qrfa5gWgDlSO4RWqHjwD4u_LKymwLsRXoHpKxr2IQ

ret = client.Seek(fd527, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


ret = client.Seek(fd524, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd527, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


ret = client.Seek(fd522, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (118) FVufQKNseI93OoaNQW61pdh5Sb9b3PiDmzBMgWwIklMDOPOElzeuCNYBB08fiUxSsmAi6s3NXFylcYMX6TAFILcL3CYz_DWUUOd7COX7WQpXRqakAAyf3Sw790VWXRshColR6

ret = client.Write(fd521, []byte("_cEKsrotJnxKg3lIuHmgPKxHVFXcCilDA1kshSeFgTLyP13GXTjbccNK6lYXJw_g5S9oTCm5S8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (192) FVufQKNseI93OoaNQW61pdh5Sb9b3PiDmzBMgWwIklMDOPOElzeuCNYBB08fiUxSsmAi6s3NXFylcYMX6TAFILcL3CYz_DWUUOd7COX7WQpXRqakAAyf3S_cEKsrotJnxKg3lIuHmgPKxHVFXcCilDA1kshSeFgTLyP13GXTjbccNK6lYXJw_g5S9oTCm5S8

buf, ret = client.Read(fd517, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd517, 114, client.SEEK_SET)
if(ret != 114) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 114)
  panic("seek failed")
}


buf, ret = client.Read(fd527, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oHpKxr2IQ") {
  panic("wrong data returned")
}


ret = client.Close(fd521)
if(ret != 0) {
  panic("close failed")
}


fd529 := client.Open("/otUQRxsRyj.txt", client.O_RDWR|client.O_CREATE)
if(fd529 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd529, []byte("2RxZH_kAhThQC7bFt3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) 2RxZH_kAhThQC7bFt3
//fd state: (114) nf519OXTMDZY0dv5VyBkscG4l04ejwmZF0bBb9oMDyJNFe5AYqbIFZHfP4Re8TEVjtTsnpJcqZZ7DfyGLxdXZYxnaQ2Ucxs6OlXSYYE7_ClG9sHiIYHqpaS7Eof

ret = client.Write(fd517, []byte("f_PMW6bcZsldDxGDq6EU2_45Ru4fsc1E0XpgOdB_10ce8mdxAWKWaSIlHtDOhLW0q2Q7kGhxDGGQ1AaSvEiXC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (199) nf519OXTMDZY0dv5VyBkscG4l04ejwmZF0bBb9oMDyJNFe5AYqbIFZHfP4Re8TEVjtTsnpJcqZZ7DfyGLxdXZYxnaQ2Ucxs6OlXSYYE7_ClG9sHiIYf_PMW6bcZsldDxGDq6EU2_45Ru4fsc1E0XpgOdB_10ce8mdxAWKWaSIlHtDOhLW0q2Q7kGhxDGGQ1AaSvEiXC

ret = client.Seek(fd529, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


buf, ret = client.Read(fd527, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd524)
if(ret != 0) {
  panic("close failed")
}

//fd state: (199) nf519OXTMDZY0dv5VyBkscG4l04ejwmZF0bBb9oMDyJNFe5AYqbIFZHfP4Re8TEVjtTsnpJcqZZ7DfyGLxdXZYxnaQ2Ucxs6OlXSYYE7_ClG9sHiIYf_PMW6bcZsldDxGDq6EU2_45Ru4fsc1E0XpgOdB_10ce8mdxAWKWaSIlHtDOhLW0q2Q7kGhxDGGQ1AaSvEiXC

ret = client.Write(fd517, []byte("UK405FTYjqCe6lwx2Afs0AFHBJ0yW5dXDoVyS8aHBcQ92mAANba8uA3TPCKNqXQmfG9JckiwTuY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (274) nf519OXTMDZY0dv5VyBkscG4l04ejwmZF0bBb9oMDyJNFe5AYqbIFZHfP4Re8TEVjtTsnpJcqZZ7DfyGLxdXZYxnaQ2Ucxs6OlXSYYE7_ClG9sHiIYf_PMW6bcZsldDxGDq6EU2_45Ru4fsc1E0XpgOdB_10ce8mdxAWKWaSIlHtDOhLW0q2Q7kGhxDGGQ1AaSvEiXCUK405FTYjqCe6lwx2Afs0AFHBJ0yW5dXDoVyS8aHBcQ92mAANba8uA3TPCKNqXQmfG9JckiwTuY

ret = client.Close(fd529)
if(ret != 0) {
  panic("close failed")
}

//fd state: (274) nf519OXTMDZY0dv5VyBkscG4l04ejwmZF0bBb9oMDyJNFe5AYqbIFZHfP4Re8TEVjtTsnpJcqZZ7DfyGLxdXZYxnaQ2Ucxs6OlXSYYE7_ClG9sHiIYf_PMW6bcZsldDxGDq6EU2_45Ru4fsc1E0XpgOdB_10ce8mdxAWKWaSIlHtDOhLW0q2Q7kGhxDGGQ1AaSvEiXCUK405FTYjqCe6lwx2Afs0AFHBJ0yW5dXDoVyS8aHBcQ92mAANba8uA3TPCKNqXQmfG9JckiwTuY

ret = client.Write(fd517, []byte("L6tOzhWAq3l4e6jYm5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (292) nf519OXTMDZY0dv5VyBkscG4l04ejwmZF0bBb9oMDyJNFe5AYqbIFZHfP4Re8TEVjtTsnpJcqZZ7DfyGLxdXZYxnaQ2Ucxs6OlXSYYE7_ClG9sHiIYf_PMW6bcZsldDxGDq6EU2_45Ru4fsc1E0XpgOdB_10ce8mdxAWKWaSIlHtDOhLW0q2Q7kGhxDGGQ1AaSvEiXCUK405FTYjqCe6lwx2Afs0AFHBJ0yW5dXDoVyS8aHBcQ92mAANba8uA3TPCKNqXQmfG9JckiwTuYL6tOzhWAq3l4e6jYm5

ret = client.Seek(fd522, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd522)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd513)
if(ret != 0) {
  panic("close failed")
}


fd530 := client.Open("/nHS8yQkY68.txt", client.O_RDWR|client.O_CREATE)
if(fd530 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd530, []byte("VZcc26tdqqfEBmq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) VZcc26tdqqfEBmq

ret = client.Close(fd530)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd528, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd517, 105, client.SEEK_SET)
if(ret != 105) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 105)
  panic("seek failed")
}


ret = client.Close(fd528)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd517)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd527, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd531 := client.Open("/ogbVTic8um.txt", client.O_RDWR|client.O_CREATE)
if(fd531 < 0) {
  panic("open failed")
}


ret = client.Seek(fd531, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd527, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd532 := client.Open("/yY3JR0W531.txt", client.O_RDWR|client.O_CREATE)
if(fd532 < 0) {
  panic("open failed")
}


ret = client.Seek(fd531, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd532)
if(ret != 0) {
  panic("close failed")
}


fd533 := client.Open("/R2ldcPasbO.txt", client.O_RDWR|client.O_CREATE)
if(fd533 < 0) {
  panic("open failed")
}


ret = client.Seek(fd531, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd533, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd533, []byte("cUfvOp9XKegLEIGv3Ro7YE_YpxhyOqUMxu8e7EcCzevwnK9rk1VlOUw1EkQOGJp03vgSTKDo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) cUfvOp9XKegLEIGv3Ro7YE_YpxhyOqUMxu8e7EcCzevwnK9rk1VlOUw1EkQOGJp03vgSTKDo
//fd state: (72) cUfvOp9XKegLEIGv3Ro7YE_YpxhyOqUMxu8e7EcCzevwnK9rk1VlOUw1EkQOGJp03vgSTKDo

ret = client.Write(fd533, []byte("bnRoWBJNzxAH8PvNijh6O3VyrLHHOHPjP8CFhPLhKVIQg2TTfPsvubhLl9J7Hg0ar4M8IxDtrEWVZmQIxs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) cUfvOp9XKegLEIGv3Ro7YE_YpxhyOqUMxu8e7EcCzevwnK9rk1VlOUw1EkQOGJp03vgSTKDobnRoWBJNzxAH8PvNijh6O3VyrLHHOHPjP8CFhPLhKVIQg2TTfPsvubhLl9J7Hg0ar4M8IxDtrEWVZmQIxs

ret = client.Close(fd533)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd531)
if(ret != 0) {
  panic("close failed")
}

//fd state: (41) qrfa5gWgDlSO4RWqHjwD4u_LKymwLsRXoHpKxr2IQ

ret = client.Write(fd527, []byte("uNFafxBG9_ZVFRPxTt6yNA8BH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) qrfa5gWgDlSO4RWqHjwD4u_LKymwLsRXoHpKxr2IQuNFafxBG9_ZVFRPxTt6yNA8BH

ret = client.Close(fd527)
if(ret != 0) {
  panic("close failed")
}


fd534 := client.Open("/9nmD1_VinR.txt", client.O_RDWR|client.O_CREATE)
if(fd534 < 0) {
  panic("open failed")
}


ret = client.Close(fd534)
if(ret != 0) {
  panic("close failed")
}


fd535 := client.Open("/Klubr3k6Z5.txt", client.O_RDWR|client.O_CREATE)
if(fd535 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd535, []byte("zTcGNhc_a3CniE20bpwVLJgfGHLWl8K2rhkXh9Snxb1FJyvlagevAaMwEl6aVg98OFmsmwLWuWp4mA5YU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) zTcGNhc_a3CniE20bpwVLJgfGHLWl8K2rhkXh9Snxb1FJyvlagevAaMwEl6aVg98OFmsmwLWuWp4mA5YU

fd536 := client.Open("/hythcT1zzr.txt", client.O_RDWR|client.O_CREATE)
if(fd536 < 0) {
  panic("open failed")
}


fd537 := client.Open("/B_hZtR1Qiv.txt", client.O_RDWR|client.O_CREATE)
if(fd537 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd536, []byte("grjiTRo9t9wkqH4x1IirAGpqwXhF_YnPJc8ZlK6q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) grjiTRo9t9wkqH4x1IirAGpqwXhF_YnPJc8ZlK6q

buf, ret = client.Read(fd535, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (40) grjiTRo9t9wkqH4x1IirAGpqwXhF_YnPJc8ZlK6q

ret = client.Write(fd536, []byte("hgiJ0BAzfZWjsA72nQ6XExQdeaPH8k3OOl7c_PfY_gF7vtb_SPHY5KrVVJIrb2iPvB95uDbVAggJGBqfDAZ7YwaZuGx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (131) grjiTRo9t9wkqH4x1IirAGpqwXhF_YnPJc8ZlK6qhgiJ0BAzfZWjsA72nQ6XExQdeaPH8k3OOl7c_PfY_gF7vtb_SPHY5KrVVJIrb2iPvB95uDbVAggJGBqfDAZ7YwaZuGx
//fd state: (81) zTcGNhc_a3CniE20bpwVLJgfGHLWl8K2rhkXh9Snxb1FJyvlagevAaMwEl6aVg98OFmsmwLWuWp4mA5YU

ret = client.Write(fd535, []byte("GLSjktnuzgrew_OHZf2X98bFfRQJF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) zTcGNhc_a3CniE20bpwVLJgfGHLWl8K2rhkXh9Snxb1FJyvlagevAaMwEl6aVg98OFmsmwLWuWp4mA5YUGLSjktnuzgrew_OHZf2X98bFfRQJF

fd538 := client.Open("/K1YhjaImgW.txt", client.O_RDWR|client.O_CREATE)
if(fd538 < 0) {
  panic("open failed")
}


ret = client.Close(fd537)
if(ret != 0) {
  panic("close failed")
}


fd539 := client.Open("/0chQNY9PXi.txt", client.O_RDWR|client.O_CREATE)
if(fd539 < 0) {
  panic("open failed")
}


ret = client.Seek(fd535, 99, client.SEEK_SET)
if(ret != 99) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 99)
  panic("seek failed")
}


ret = client.Close(fd538)
if(ret != 0) {
  panic("close failed")
}

//fd state: (131) grjiTRo9t9wkqH4x1IirAGpqwXhF_YnPJc8ZlK6qhgiJ0BAzfZWjsA72nQ6XExQdeaPH8k3OOl7c_PfY_gF7vtb_SPHY5KrVVJIrb2iPvB95uDbVAggJGBqfDAZ7YwaZuGx

ret = client.Write(fd536, []byte("WoBKPJLxtbVuvVzm6WHzU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) grjiTRo9t9wkqH4x1IirAGpqwXhF_YnPJc8ZlK6qhgiJ0BAzfZWjsA72nQ6XExQdeaPH8k3OOl7c_PfY_gF7vtb_SPHY5KrVVJIrb2iPvB95uDbVAggJGBqfDAZ7YwaZuGxWoBKPJLxtbVuvVzm6WHzU

fd540 := client.Open("/9yJ4B8r9Jr.txt", client.O_RDWR|client.O_CREATE)
if(fd540 < 0) {
  panic("open failed")
}


fd541 := client.Open("/q8LarXbP95.txt", client.O_RDWR|client.O_CREATE)
if(fd541 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd541, []byte("3HOXzTKG0GOrEwu5NS8NwYlwRiqR3MuNthCcwiICIbRbRksJQ31GXxFkvyErywJpOOPK50FRX64nM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) 3HOXzTKG0GOrEwu5NS8NwYlwRiqR3MuNthCcwiICIbRbRksJQ31GXxFkvyErywJpOOPK50FRX64nM

buf, ret = client.Read(fd535, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2X98bFfRQJF") {
  panic("wrong data returned")
}


ret = client.Close(fd536)
if(ret != 0) {
  panic("close failed")
}

//fd state: (110) zTcGNhc_a3CniE20bpwVLJgfGHLWl8K2rhkXh9Snxb1FJyvlagevAaMwEl6aVg98OFmsmwLWuWp4mA5YUGLSjktnuzgrew_OHZf2X98bFfRQJF

ret = client.Write(fd535, []byte("G2HlM6_vYuNI4dtzeBHrfBUQHAgwoAUldLzTlFHMsZ9sZomFKW5n4cWh6NasP_pCb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (175) zTcGNhc_a3CniE20bpwVLJgfGHLWl8K2rhkXh9Snxb1FJyvlagevAaMwEl6aVg98OFmsmwLWuWp4mA5YUGLSjktnuzgrew_OHZf2X98bFfRQJFG2HlM6_vYuNI4dtzeBHrfBUQHAgwoAUldLzTlFHMsZ9sZomFKW5n4cWh6NasP_pCb

buf, ret = client.Read(fd539, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "odSH7nuqvS") {
  panic("wrong data returned")
}


fd542 := client.Open("/8Tqb7q4mha.txt", client.O_RDWR|client.O_CREATE)
if(fd542 < 0) {
  panic("open failed")
}

//fd state: (175) zTcGNhc_a3CniE20bpwVLJgfGHLWl8K2rhkXh9Snxb1FJyvlagevAaMwEl6aVg98OFmsmwLWuWp4mA5YUGLSjktnuzgrew_OHZf2X98bFfRQJFG2HlM6_vYuNI4dtzeBHrfBUQHAgwoAUldLzTlFHMsZ9sZomFKW5n4cWh6NasP_pCb

ret = client.Write(fd535, []byte("YfcJ0NCZWSJfHqAluRDibr_OEAPUjuT9Y3xm2gcsi8DHmLMtjA1zWryJScIJcViF5o"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (241) zTcGNhc_a3CniE20bpwVLJgfGHLWl8K2rhkXh9Snxb1FJyvlagevAaMwEl6aVg98OFmsmwLWuWp4mA5YUGLSjktnuzgrew_OHZf2X98bFfRQJFG2HlM6_vYuNI4dtzeBHrfBUQHAgwoAUldLzTlFHMsZ9sZomFKW5n4cWh6NasP_pCbYfcJ0NCZWSJfHqAluRDibr_OEAPUjuT9Y3xm2gcsi8DHmLMtjA1zWryJScIJcViF5o

ret = client.Close(fd539)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd540, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "W8H61C5wBk4OZBhwIjpsotZnsc") {
  panic("wrong data returned")
}

//fd state: (26) W8H61C5wBk4OZBhwIjpsotZnsc5DKCeC509AQLRPmHUX4x_p0bl6fiCNQmKYRItTr4QjwBhfci8IndXWvja0s_qiQ_A

ret = client.Write(fd540, []byte("ZthO3t5b9iVlLlIhmfRyYTjtbsWE8gB7vb6nQk9XZvXO4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) W8H61C5wBk4OZBhwIjpsotZnscZthO3t5b9iVlLlIhmfRyYTjtbsWE8gB7vb6nQk9XZvXO4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpK
//fd state: (124) W8H61C5wBk4OZBhwIjpsotZnscZthO3t5b9iVlLlIhmfRyYTjtbsWE8gB7vb6nQk9XZvXO4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpK

ret = client.Write(fd540, []byte("eyLzy41eRxSNsL1Q9IAhKULwwemEgktFrpqbjQhnnmQywB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (170) W8H61C5wBk4OZBhwIjpsotZnscZthO3t5b9iVlLlIhmfRyYTjtbsWE8gB7vb6nQk9XZvXO4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpKeyLzy41eRxSNsL1Q9IAhKULwwemEgktFrpqbjQhnnmQywB

fd543 := client.Open("/R1VJ6U8gR7.txt", client.O_RDWR|client.O_CREATE)
if(fd543 < 0) {
  panic("open failed")
}


ret = client.Seek(fd541, 77, client.SEEK_SET)
if(ret != 77) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 77)
  panic("seek failed")
}


ret = client.Seek(fd540, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


buf, ret = client.Read(fd540, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lIhmfRyYTjtbsWE8gB7vb6nQk9XZvXO4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpK") {
  panic("wrong data returned")
}


ret = client.Close(fd535)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd542)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd543, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd544 := client.Open("/24JwXc5xc3.txt", client.O_RDWR|client.O_CREATE)
if(fd544 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd543, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (124) W8H61C5wBk4OZBhwIjpsotZnscZthO3t5b9iVlLlIhmfRyYTjtbsWE8gB7vb6nQk9XZvXO4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpKeyLzy41eRxSNsL1Q9IAhKULwwemEgktFrpqbjQhnnmQywB

ret = client.Write(fd540, []byte("KxDKhvrh19IT1W1Z4PlF93B1ObQVMj2gwufj4dyDSOcyDSnXTtyS2jqHszsGiuJCKFJssqXw4Jb2zd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (202) W8H61C5wBk4OZBhwIjpsotZnscZthO3t5b9iVlLlIhmfRyYTjtbsWE8gB7vb6nQk9XZvXO4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpKKxDKhvrh19IT1W1Z4PlF93B1ObQVMj2gwufj4dyDSOcyDSnXTtyS2jqHszsGiuJCKFJssqXw4Jb2zd
//fd state: (0) 

ret = client.Write(fd543, []byte("iCyQAa0N1TiYXbdRFTi_en3qcKdb2A5JdfFQBm78vxguj_dMR8nYxldTXnssjXhCfj6SR9H7ofr3LYm3k27qkwd3Ppxv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) iCyQAa0N1TiYXbdRFTi_en3qcKdb2A5JdfFQBm78vxguj_dMR8nYxldTXnssjXhCfj6SR9H7ofr3LYm3k27qkwd3Ppxv
//fd state: (77) 3HOXzTKG0GOrEwu5NS8NwYlwRiqR3MuNthCcwiICIbRbRksJQ31GXxFkvyErywJpOOPK50FRX64nM

ret = client.Write(fd541, []byte("HcfyCLR0RTcRQqh4Sld1gWilx8ViMthNDMH0Z7fG4bwipaNtYqgFuHZYtREtaBXxlZDuh2p0CZ5um_rRkYOk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (161) 3HOXzTKG0GOrEwu5NS8NwYlwRiqR3MuNthCcwiICIbRbRksJQ31GXxFkvyErywJpOOPK50FRX64nMHcfyCLR0RTcRQqh4Sld1gWilx8ViMthNDMH0Z7fG4bwipaNtYqgFuHZYtREtaBXxlZDuh2p0CZ5um_rRkYOk

ret = client.Close(fd543)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd540, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd545 := client.Open("/K1YhjaImgW.txt", client.O_RDWR|client.O_CREATE)
if(fd545 < 0) {
  panic("open failed")
}


ret = client.Seek(fd541, 87, client.SEEK_SET)
if(ret != 87) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 87)
  panic("seek failed")
}


ret = client.Seek(fd541, 119, client.SEEK_SET)
if(ret != 119) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 119)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd544, []byte("fnQcriQLBZ5QzMCrGWL53X0c9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) fnQcriQLBZ5QzMCrGWL53X0c9

ret = client.Close(fd544)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd540, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}


buf, ret = client.Read(fd540, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpKKxDKhvrh19IT1W1Z4PlF") {
  panic("wrong data returned")
}

//fd state: (144) W8H61C5wBk4OZBhwIjpsotZnscZthO3t5b9iVlLlIhmfRyYTjtbsWE8gB7vb6nQk9XZvXO4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpKKxDKhvrh19IT1W1Z4PlF93B1ObQVMj2gwufj4dyDSOcyDSnXTtyS2jqHszsGiuJCKFJssqXw4Jb2zd

ret = client.Write(fd540, []byte("JVQVnOqLQKsw52B3w9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (162) W8H61C5wBk4OZBhwIjpsotZnscZthO3t5b9iVlLlIhmfRyYTjtbsWE8gB7vb6nQk9XZvXO4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpKKxDKhvrh19IT1W1Z4PlFJVQVnOqLQKsw52B3w9yDSOcyDSnXTtyS2jqHszsGiuJCKFJssqXw4Jb2zd

ret = client.Close(fd541)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) hXnSISfeUo4y_s3jbsXgtxo5NI5ulgVyxOjLbpyBDBfuGWSUqFG3O8XL_WgsMDMO0rJmHPWTn97D6rdXBw82WXuyevEUqP9O3wdExRt2wBdLHDLq2HuQGFRcaVJLUiID6i0zDchzG2vupiyxrfnasqeC2dqSRo6ETem4FYSxZZTqRoPfRv7anR

ret = client.Write(fd545, []byte("_R_DPCJezXOeVZ5nZTzgv8eEUIOtHH9UnnDXY5NZQuTAy9mRlnSzo_tm7vQErhNxiEr7r1fib5PCUKnpA_4y1Y9r0fR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) _R_DPCJezXOeVZ5nZTzgv8eEUIOtHH9UnnDXY5NZQuTAy9mRlnSzo_tm7vQErhNxiEr7r1fib5PCUKnpA_4y1Y9r0fRUqP9O3wdExRt2wBdLHDLq2HuQGFRcaVJLUiID6i0zDchzG2vupiyxrfnasqeC2dqSRo6ETem4FYSxZZTqRoPfRv7anR

fd546 := client.Open("/sKX3ZdVyBq.txt", client.O_RDWR|client.O_CREATE)
if(fd546 < 0) {
  panic("open failed")
}

//fd state: (91) _R_DPCJezXOeVZ5nZTzgv8eEUIOtHH9UnnDXY5NZQuTAy9mRlnSzo_tm7vQErhNxiEr7r1fib5PCUKnpA_4y1Y9r0fRUqP9O3wdExRt2wBdLHDLq2HuQGFRcaVJLUiID6i0zDchzG2vupiyxrfnasqeC2dqSRo6ETem4FYSxZZTqRoPfRv7anR

ret = client.Write(fd545, []byte("58jFcMlx8dTodfYOfbMPBEFQuMu8fdOVGFeJMwU8RsEajjXEkNy1DLyT1rZlZFbodx9APXThZIjNDkze"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (171) _R_DPCJezXOeVZ5nZTzgv8eEUIOtHH9UnnDXY5NZQuTAy9mRlnSzo_tm7vQErhNxiEr7r1fib5PCUKnpA_4y1Y9r0fR58jFcMlx8dTodfYOfbMPBEFQuMu8fdOVGFeJMwU8RsEajjXEkNy1DLyT1rZlZFbodx9APXThZIjNDkzeqRoPfRv7anR

ret = client.Seek(fd540, 193, client.SEEK_SET)
if(ret != 193) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 193)
  panic("seek failed")
}


ret = client.Close(fd540)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd545)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd546, []byte("PXoDTuKPxKT5I_HyVM1oEwOdLfm1DiHDkHOA22IF2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) PXoDTuKPxKT5I_HyVM1oEwOdLfm1DiHDkHOA22IF2

buf, ret = client.Read(fd546, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd546)
if(ret != 0) {
  panic("close failed")
}


fd547 := client.Open("/gQMBiUhjLf.txt", client.O_RDWR|client.O_CREATE)
if(fd547 < 0) {
  panic("open failed")
}


ret = client.Close(fd547)
if(ret != 0) {
  panic("close failed")
}


fd548 := client.Open("/PmQIOvOqZw.txt", client.O_RDWR|client.O_CREATE)
if(fd548 < 0) {
  panic("open failed")
}


fd549 := client.Open("/hqQ4RRTmCR.txt", client.O_RDWR|client.O_CREATE)
if(fd549 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd548, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd550 := client.Open("/uV1nPoDUtI.txt", client.O_RDWR|client.O_CREATE)
if(fd550 < 0) {
  panic("open failed")
}


ret = client.Seek(fd550, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd548, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd551 := client.Open("/xtYmEEvAxH.txt", client.O_RDWR|client.O_CREATE)
if(fd551 < 0) {
  panic("open failed")
}


ret = client.Seek(fd550, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd550, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd550, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd551, []byte("_FRYnOm1HTymqLTM4S809ezU5C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) _FRYnOm1HTymqLTM4S809ezU5C

fd552 := client.Open("/dQ5BvVzU2a.txt", client.O_RDWR|client.O_CREATE)
if(fd552 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd548, []byte("888MnhCjWzpf6aJYKSwmcPJqSlTPzSDCvgBJ0xWwo1RF1S8Occ4kK2KGf1eT2df2_lNccRsnGaGw7IsE1Gf8t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) 888MnhCjWzpf6aJYKSwmcPJqSlTPzSDCvgBJ0xWwo1RF1S8Occ4kK2KGf1eT2df2_lNccRsnGaGw7IsE1Gf8t

ret = client.Seek(fd550, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (26) _FRYnOm1HTymqLTM4S809ezU5C

ret = client.Write(fd551, []byte("FtL9GK3iTs0Qj_ZnWeAmAXgo1zod6CNEvfUnET8K46y_YhNonfHWKZmMu8TZ21hbugQZBJGXvSX1H6T8JbDOQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) _FRYnOm1HTymqLTM4S809ezU5CFtL9GK3iTs0Qj_ZnWeAmAXgo1zod6CNEvfUnET8K46y_YhNonfHWKZmMu8TZ21hbugQZBJGXvSX1H6T8JbDOQ

ret = client.Close(fd550)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd552, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd548)
if(ret != 0) {
  panic("close failed")
}


fd553 := client.Open("/QWehDPrfor.txt", client.O_RDWR|client.O_CREATE)
if(fd553 < 0) {
  panic("open failed")
}


ret = client.Seek(fd551, 61, client.SEEK_SET)
if(ret != 61) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 61)
  panic("seek failed")
}


fd554 := client.Open("/pj4lLaSsh8.txt", client.O_RDWR|client.O_CREATE)
if(fd554 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd549, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd555 := client.Open("/Ru9u2Mvc2A.txt", client.O_RDWR|client.O_CREATE)
if(fd555 < 0) {
  panic("open failed")
}


ret = client.Close(fd551)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd555)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd549, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd556 := client.Open("/9TFyXXeDd1.txt", client.O_RDWR|client.O_CREATE)
if(fd556 < 0) {
  panic("open failed")
}


ret = client.Close(fd556)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd554, []byte("whPSDFF8me"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) whPSDFF8me

ret = client.Close(fd549)
if(ret != 0) {
  panic("close failed")
}

//fd state: (10) whPSDFF8me

ret = client.Write(fd554, []byte("Ghk8iAX7SdfzGRSiCjByRgGO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) whPSDFF8meGhk8iAX7SdfzGRSiCjByRgGO
//fd state: (0) 

ret = client.Write(fd552, []byte("AxxznkIjtif4oETVQKkWYABPqUmvfGNPZz12MSixvq7nYXCHYanH4KEsqhpJLya4TUEXq4O9cf7FfXrtB1WBozMhPGWd77z6j6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) AxxznkIjtif4oETVQKkWYABPqUmvfGNPZz12MSixvq7nYXCHYanH4KEsqhpJLya4TUEXq4O9cf7FfXrtB1WBozMhPGWd77z6j6

buf, ret = client.Read(fd552, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd552, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd557 := client.Open("/bRoENni7CO.txt", client.O_RDWR|client.O_CREATE)
if(fd557 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd554, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd554, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


ret = client.Close(fd554)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd557)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd553)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd552)
if(ret != 0) {
  panic("close failed")
}


fd558 := client.Open("/Ky3khvieYX.txt", client.O_RDWR|client.O_CREATE)
if(fd558 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd558, []byte("5Z9tYcPeAOSwhQe7SfEOHyqb6UtnfZTMBN9R4D7yfktW4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) 5Z9tYcPeAOSwhQe7SfEOHyqb6UtnfZTMBN9R4D7yfktW4

ret = client.Close(fd558)
if(ret != 0) {
  panic("close failed")
}


fd559 := client.Open("/YfmUYg5AyQ.txt", client.O_RDWR|client.O_CREATE)
if(fd559 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd559)
if(ret != 0) {
  panic("close failed")
}


fd560 := client.Open("/FopR58AGxt.txt", client.O_RDWR|client.O_CREATE)
if(fd560 < 0) {
  panic("open failed")
}


fd561 := client.Open("/kqZrpEP1nq.txt", client.O_RDWR|client.O_CREATE)
if(fd561 < 0) {
  panic("open failed")
}


fd562 := client.Open("/jMLCbPratd.txt", client.O_RDWR|client.O_CREATE)
if(fd562 < 0) {
  panic("open failed")
}


fd563 := client.Open("/OmH6ue7qfz.txt", client.O_RDWR|client.O_CREATE)
if(fd563 < 0) {
  panic("open failed")
}


fd564 := client.Open("/yX6wGc_UzJ.txt", client.O_RDWR|client.O_CREATE)
if(fd564 < 0) {
  panic("open failed")
}


fd565 := client.Open("/JfhBLqKHCq.txt", client.O_RDWR|client.O_CREATE)
if(fd565 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd563, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "0NUSBWlxgvZCB") {
  panic("wrong data returned")
}


ret = client.Seek(fd564, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd564)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd565, []byte("ywz9AK8BIlvdCMyPMheA9qQOnov9qt9oXRLO_gbrjdqpVZO2vgFTktLxgA_hQKft84gK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) ywz9AK8BIlvdCMyPMheA9qQOnov9qt9oXRLO_gbrjdqpVZO2vgFTktLxgA_hQKft84gK

ret = client.Close(fd565)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd562, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


buf, ret = client.Read(fd561, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd566 := client.Open("/0hmueH147q.txt", client.O_RDWR|client.O_CREATE)
if(fd566 < 0) {
  panic("open failed")
}

//fd state: (13) 0NUSBWlxgvZCB_gsv7K2iMnSw1tkb05OxzISIxbdtrhpQlfiGVPO61oL2jGetMZ5ScwoKdiq_BbepVxUNRTCYQd2nfsCIENMtS32An12d295rxNcFiuFQSBHkTtYFdSidEA3t1BqJvhFuyTeP3_Pl9DgZrW4wVPxTpzZLCJSuopJrrGY

ret = client.Write(fd563, []byte("hP_bFkgZt45wL_ocKYI36scnLwLvP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) 0NUSBWlxgvZCBhP_bFkgZt45wL_ocKYI36scnLwLvPhpQlfiGVPO61oL2jGetMZ5ScwoKdiq_BbepVxUNRTCYQd2nfsCIENMtS32An12d295rxNcFiuFQSBHkTtYFdSidEA3t1BqJvhFuyTeP3_Pl9DgZrW4wVPxTpzZLCJSuopJrrGY

ret = client.Seek(fd563, 165, client.SEEK_SET)
if(ret != 165) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 165)
  panic("seek failed")
}


buf, ret = client.Read(fd560, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd567 := client.Open("/sKX3ZdVyBq.txt", client.O_RDWR|client.O_CREATE)
if(fd567 < 0) {
  panic("open failed")
}

//fd state: (165) 0NUSBWlxgvZCBhP_bFkgZt45wL_ocKYI36scnLwLvPhpQlfiGVPO61oL2jGetMZ5ScwoKdiq_BbepVxUNRTCYQd2nfsCIENMtS32An12d295rxNcFiuFQSBHkTtYFdSidEA3t1BqJvhFuyTeP3_Pl9DgZrW4wVPxTpzZLCJSuopJrrGY

ret = client.Write(fd563, []byte("0kLYOymOErnUnh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (179) 0NUSBWlxgvZCBhP_bFkgZt45wL_ocKYI36scnLwLvPhpQlfiGVPO61oL2jGetMZ5ScwoKdiq_BbepVxUNRTCYQd2nfsCIENMtS32An12d295rxNcFiuFQSBHkTtYFdSidEA3t1BqJvhFuyTeP3_Pl9DgZrW4wVPxTpzZL0kLYOymOErnUnh

buf, ret = client.Read(fd560, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd561, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd563, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


fd568 := client.Open("/mHY5KcNlPE.txt", client.O_RDWR|client.O_CREATE)
if(fd568 < 0) {
  panic("open failed")
}

//fd state: (0) PXoDTuKPxKT5I_HyVM1oEwOdLfm1DiHDkHOA22IF2

ret = client.Write(fd567, []byte("mMVB0Y38aGaZSF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) mMVB0Y38aGaZSFHyVM1oEwOdLfm1DiHDkHOA22IF2
//fd state: (9) wqYVnAjP6G9KDsf9Q_dYKmZuYddqYXqkjJVPVyC8TpnCz8htDuMrodZSjqYf7AR53z

ret = client.Write(fd562, []byte("pK6uhTNAkGY5kkCKvDQGdsuTcJ9q3Q9EbcZC5ETp1OZZL2bo_QC8wd3hJhiqSrtpB3A8O9wOH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) wqYVnAjP6pK6uhTNAkGY5kkCKvDQGdsuTcJ9q3Q9EbcZC5ETp1OZZL2bo_QC8wd3hJhiqSrtpB3A8O9wOH

ret = client.Close(fd560)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd562)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd561)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd566, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) 

ret = client.Seek(fd566, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (14) mMVB0Y38aGaZSFHyVM1oEwOdLfm1DiHDkHOA22IF2

ret = client.Write(fd567, []byte("ees_ZMzwV_61l0grdHK0ExoxKRZ9cZKf9mobu9eHu_HRMeJ8BrSZhF877Jr6HC48cqV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) mMVB0Y38aGaZSFees_ZMzwV_61l0grdHK0ExoxKRZ9cZKf9mobu9eHu_HRMeJ8BrSZhF877Jr6HC48cqV

buf, ret = client.Read(fd568, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "eDqhrw3RKPvL4xNeTEXLO867sF6k0gWDmy8tL4StOg71tQYWUoEqrqELgmqqqI05_aOUP0UGcywSrwhEkYjG") {
  panic("wrong data returned")
}

//fd state: (81) mMVB0Y38aGaZSFees_ZMzwV_61l0grdHK0ExoxKRZ9cZKf9mobu9eHu_HRMeJ8BrSZhF877Jr6HC48cqV

ret = client.Write(fd567, []byte("qF2yXZ8vh4KsjXK9pQpmtgaadsd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (108) mMVB0Y38aGaZSFees_ZMzwV_61l0grdHK0ExoxKRZ9cZKf9mobu9eHu_HRMeJ8BrSZhF877Jr6HC48cqVqF2yXZ8vh4KsjXK9pQpmtgaadsd

ret = client.Seek(fd566, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd568)
if(ret != 0) {
  panic("close failed")
}

//fd state: (108) mMVB0Y38aGaZSFees_ZMzwV_61l0grdHK0ExoxKRZ9cZKf9mobu9eHu_HRMeJ8BrSZhF877Jr6HC48cqVqF2yXZ8vh4KsjXK9pQpmtgaadsd

ret = client.Write(fd567, []byte("vELARwnVEWAxbi8R42fF5kDH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) mMVB0Y38aGaZSFees_ZMzwV_61l0grdHK0ExoxKRZ9cZKf9mobu9eHu_HRMeJ8BrSZhF877Jr6HC48cqVqF2yXZ8vh4KsjXK9pQpmtgaadsdvELARwnVEWAxbi8R42fF5kDH

ret = client.Seek(fd566, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd566)
if(ret != 0) {
  panic("close failed")
}


fd569 := client.Open("/Rlp_5jzvVy.txt", client.O_RDWR|client.O_CREATE)
if(fd569 < 0) {
  panic("open failed")
}


ret = client.Close(fd569)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd567, 60, client.SEEK_SET)
if(ret != 60) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 60)
  panic("seek failed")
}


ret = client.Close(fd567)
if(ret != 0) {
  panic("close failed")
}


fd570 := client.Open("/xxjdlztmab.txt", client.O_RDWR|client.O_CREATE)
if(fd570 < 0) {
  panic("open failed")
}


fd571 := client.Open("/JfhBLqKHCq.txt", client.O_RDWR|client.O_CREATE)
if(fd571 < 0) {
  panic("open failed")
}


ret = client.Close(fd563)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) ywz9AK8BIlvdCMyPMheA9qQOnov9qt9oXRLO_gbrjdqpVZO2vgFTktLxgA_hQKft84gK

ret = client.Write(fd571, []byte("MvzAvBFqEer5NNi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) MvzAvBFqEer5NNiPMheA9qQOnov9qt9oXRLO_gbrjdqpVZO2vgFTktLxgA_hQKft84gK
//fd state: (15) MvzAvBFqEer5NNiPMheA9qQOnov9qt9oXRLO_gbrjdqpVZO2vgFTktLxgA_hQKft84gK

ret = client.Write(fd571, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) MvzAvBFqEer5NNiPMheA9qQOnov9qt9oXRLO_gbrjdqpVZO2vgFTktLxgA_hQKft84gK

ret = client.Seek(fd571, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


buf, ret = client.Read(fd571, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "FqEer5NNiPMheA9qQOno") {
  panic("wrong data returned")
}


ret = client.Close(fd570)
if(ret != 0) {
  panic("close failed")
}


fd572 := client.Open("/Qc15HraxHA.txt", client.O_RDWR|client.O_CREATE)
if(fd572 < 0) {
  panic("open failed")
}


ret = client.Close(fd571)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd572, []byte("V3X2ykwBdJoSof_YFrSwdyL4a6xwdqBQcFrpNXJ_gJlrTfqgTnp4V97cBm_dJI6DLBIRrBITfp19CCFvuOZE0Shv2_YM97OoJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) V3X2ykwBdJoSof_YFrSwdyL4a6xwdqBQcFrpNXJ_gJlrTfqgTnp4V97cBm_dJI6DLBIRrBITfp19CCFvuOZE0Shv2_YM97OoJ

buf, ret = client.Read(fd572, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd572)
if(ret != 0) {
  panic("close failed")
}


fd573 := client.Open("/BJqa1WNiUz.txt", client.O_RDWR|client.O_CREATE)
if(fd573 < 0) {
  panic("open failed")
}


ret = client.Seek(fd573, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd573, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd573)
if(ret != 0) {
  panic("close failed")
}


fd574 := client.Open("/U_WXa7M3fC.txt", client.O_RDWR|client.O_CREATE)
if(fd574 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd574, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd575 := client.Open("/q8LarXbP95.txt", client.O_RDWR|client.O_CREATE)
if(fd575 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd574, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd574, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd576 := client.Open("/nHgSrOUu58.txt", client.O_RDWR|client.O_CREATE)
if(fd576 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd576, []byte("EbAy7WzBttksknh55gAHR9k1PwrxlCZxQMUVpZmvDxW5JLMwzGLA17Y9jcVGKwW7137H2_UsfxWyuQWFTKSXfqHrFtg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) EbAy7WzBttksknh55gAHR9k1PwrxlCZxQMUVpZmvDxW5JLMwzGLA17Y9jcVGKwW7137H2_UsfxWyuQWFTKSXfqHrFtg

ret = client.Close(fd574)
if(ret != 0) {
  panic("close failed")
}

//fd state: (91) EbAy7WzBttksknh55gAHR9k1PwrxlCZxQMUVpZmvDxW5JLMwzGLA17Y9jcVGKwW7137H2_UsfxWyuQWFTKSXfqHrFtg

ret = client.Write(fd576, []byte("hW009K8tEtwGur1p4oobRhbUE4P0Ayn33K3krqCPwIqCMOebQ_h0MKAc4sSlZlAXO6ry9E0NHiURxgET"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (171) EbAy7WzBttksknh55gAHR9k1PwrxlCZxQMUVpZmvDxW5JLMwzGLA17Y9jcVGKwW7137H2_UsfxWyuQWFTKSXfqHrFtghW009K8tEtwGur1p4oobRhbUE4P0Ayn33K3krqCPwIqCMOebQ_h0MKAc4sSlZlAXO6ry9E0NHiURxgET

buf, ret = client.Read(fd576, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd575, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


ret = client.Seek(fd576, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


ret = client.Close(fd575)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd576, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "FTKSXfqHrFtghW009K8tEtwGur1") {
  panic("wrong data returned")
}

//fd state: (106) EbAy7WzBttksknh55gAHR9k1PwrxlCZxQMUVpZmvDxW5JLMwzGLA17Y9jcVGKwW7137H2_UsfxWyuQWFTKSXfqHrFtghW009K8tEtwGur1p4oobRhbUE4P0Ayn33K3krqCPwIqCMOebQ_h0MKAc4sSlZlAXO6ry9E0NHiURxgET

ret = client.Write(fd576, []byte("wDy9FWYxoujljPRAqUB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (125) EbAy7WzBttksknh55gAHR9k1PwrxlCZxQMUVpZmvDxW5JLMwzGLA17Y9jcVGKwW7137H2_UsfxWyuQWFTKSXfqHrFtghW009K8tEtwGur1wDy9FWYxoujljPRAqUB3krqCPwIqCMOebQ_h0MKAc4sSlZlAXO6ry9E0NHiURxgET

fd577 := client.Open("/llTndggwBW.txt", client.O_RDWR|client.O_CREATE)
if(fd577 < 0) {
  panic("open failed")
}


ret = client.Seek(fd577, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd576, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3krqCPwIqCMOebQ_h0MKAc4sSlZlAXO6ry9E0NHiURxgET") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd576, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd577, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (171) EbAy7WzBttksknh55gAHR9k1PwrxlCZxQMUVpZmvDxW5JLMwzGLA17Y9jcVGKwW7137H2_UsfxWyuQWFTKSXfqHrFtghW009K8tEtwGur1wDy9FWYxoujljPRAqUB3krqCPwIqCMOebQ_h0MKAc4sSlZlAXO6ry9E0NHiURxgET

ret = client.Write(fd576, []byte("zhno1_SHXgwQesqmyeyYfWZ528m98tj0krB_lJ2VCAi061W3bIepIFL4jnL6qZCXfdPSmdsK4peIRGMF78_xH2sK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (259) EbAy7WzBttksknh55gAHR9k1PwrxlCZxQMUVpZmvDxW5JLMwzGLA17Y9jcVGKwW7137H2_UsfxWyuQWFTKSXfqHrFtghW009K8tEtwGur1wDy9FWYxoujljPRAqUB3krqCPwIqCMOebQ_h0MKAc4sSlZlAXO6ry9E0NHiURxgETzhno1_SHXgwQesqmyeyYfWZ528m98tj0krB_lJ2VCAi061W3bIepIFL4jnL6qZCXfdPSmdsK4peIRGMF78_xH2sK

ret = client.Close(fd576)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd577, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd577, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd577, []byte("4Mwt93rqLa91USjecWRbTQ7zz8JoZxZri7Cq7gOrZWglQ6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) 4Mwt93rqLa91USjecWRbTQ7zz8JoZxZri7Cq7gOrZWglQ6
//fd state: (46) 4Mwt93rqLa91USjecWRbTQ7zz8JoZxZri7Cq7gOrZWglQ6

ret = client.Write(fd577, []byte("l7CqvEw8XGF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) 4Mwt93rqLa91USjecWRbTQ7zz8JoZxZri7Cq7gOrZWglQ6l7CqvEw8XGF

buf, ret = client.Read(fd577, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd577, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


fd578 := client.Open("/mbhXIMTJYM.txt", client.O_RDWR|client.O_CREATE)
if(fd578 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd578, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd578)
if(ret != 0) {
  panic("close failed")
}

//fd state: (15) 4Mwt93rqLa91USjecWRbTQ7zz8JoZxZri7Cq7gOrZWglQ6l7CqvEw8XGF

ret = client.Write(fd577, []byte("H3K0oaUWTPmXYI8usYITtMunAXC2qHgMJXN7TWwqkiPWrLPIzMizjAdAztYrRW3mUgJwTjtLwqdHthfY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) 4Mwt93rqLa91USjH3K0oaUWTPmXYI8usYITtMunAXC2qHgMJXN7TWwqkiPWrLPIzMizjAdAztYrRW3mUgJwTjtLwqdHthfY

buf, ret = client.Read(fd577, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd577, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}

//fd state: (20) 4Mwt93rqLa91USjH3K0oaUWTPmXYI8usYITtMunAXC2qHgMJXN7TWwqkiPWrLPIzMizjAdAztYrRW3mUgJwTjtLwqdHthfY

ret = client.Write(fd577, []byte("YuiLpbFK3yQYOmf0QROoUgWCreQz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) 4Mwt93rqLa91USjH3K0oYuiLpbFK3yQYOmf0QROoUgWCreQzXN7TWwqkiPWrLPIzMizjAdAztYrRW3mUgJwTjtLwqdHthfY

ret = client.Close(fd577)
if(ret != 0) {
  panic("close failed")
}


fd579 := client.Open("/ALA4lYOD2g.txt", client.O_RDWR|client.O_CREATE)
if(fd579 < 0) {
  panic("open failed")
}


fd580 := client.Open("/aWiyrgOINA.txt", client.O_RDWR|client.O_CREATE)
if(fd580 < 0) {
  panic("open failed")
}


ret = client.Close(fd579)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) owau6G368X7UsbPB5QyprSyZ8oQ9m8ZwgPYkf9QCkWDipERDel2cABQdEuPLetAscml

ret = client.Write(fd580, []byte("9k2BEf9pBBO9iTsJA7f2Xjc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) 9k2BEf9pBBO9iTsJA7f2XjcZ8oQ9m8ZwgPYkf9QCkWDipERDel2cABQdEuPLetAscml
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd580)
if(ret != 0) {
  panic("close failed")
}


fd581 := client.Open("/PIpkJxo4gp.txt", client.O_RDWR|client.O_CREATE)
if(fd581 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd581, []byte("JuomKku8QhOHffZySaiLl4MJI1FIqdDqqh1yYIER"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) JuomKku8QhOHffZySaiLl4MJI1FIqdDqqh1yYIER

buf, ret = client.Read(fd581, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd581, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


buf, ret = client.Read(fd581, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ySaiLl4MJI1FIqdDqqh1yYIER") {
  panic("wrong data returned")
}


ret = client.Seek(fd581, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


fd582 := client.Open("/CHOx395o99.txt", client.O_RDWR|client.O_CREATE)
if(fd582 < 0) {
  panic("open failed")
}

//fd state: (0) d0DRyGvTe7Jr648NvHumNLQKgQhPZojzFFuQYSYYYh93TPQbqGDoz61W6XgF4D0dv4KWlpQZNEct87UKobSBzXi4eOdgsbHxujmfLkD_nRmX9Po67SG7eW4xvcx_kszYFFCre3dQpv3QqhKl5YPrX_MywXc8Sm2cFeCd27fuDzD

ret = client.Write(fd582, []byte("vRXJnbrSnjmVde0pRwBFtki2a2UurdkJ_CZPBSI3unZ5Ft04am1PoTfhH5wzaCSFv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) vRXJnbrSnjmVde0pRwBFtki2a2UurdkJ_CZPBSI3unZ5Ft04am1PoTfhH5wzaCSFv4KWlpQZNEct87UKobSBzXi4eOdgsbHxujmfLkD_nRmX9Po67SG7eW4xvcx_kszYFFCre3dQpv3QqhKl5YPrX_MywXc8Sm2cFeCd27fuDzD
//fd state: (32) JuomKku8QhOHffZySaiLl4MJI1FIqdDqqh1yYIER

ret = client.Write(fd581, []byte("Xb7oyeVpdLOPfCT40tq0De2JbX8o9BaqVCNezCLhhI0xvSWU4Nn3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) JuomKku8QhOHffZySaiLl4MJI1FIqdDqXb7oyeVpdLOPfCT40tq0De2JbX8o9BaqVCNezCLhhI0xvSWU4Nn3

ret = client.Seek(fd581, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}

//fd state: (17) JuomKku8QhOHffZySaiLl4MJI1FIqdDqXb7oyeVpdLOPfCT40tq0De2JbX8o9BaqVCNezCLhhI0xvSWU4Nn3

ret = client.Write(fd581, []byte("IVq26a93PyaZmbwtuJPEwUOOA97lyT7fk9ttEgvsjW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) JuomKku8QhOHffZySIVq26a93PyaZmbwtuJPEwUOOA97lyT7fk9ttEgvsjWo9BaqVCNezCLhhI0xvSWU4Nn3

ret = client.Close(fd581)
if(ret != 0) {
  panic("close failed")
}

//fd state: (65) vRXJnbrSnjmVde0pRwBFtki2a2UurdkJ_CZPBSI3unZ5Ft04am1PoTfhH5wzaCSFv4KWlpQZNEct87UKobSBzXi4eOdgsbHxujmfLkD_nRmX9Po67SG7eW4xvcx_kszYFFCre3dQpv3QqhKl5YPrX_MywXc8Sm2cFeCd27fuDzD

ret = client.Write(fd582, []byte("ZVm6v7JlDyIsSOtlDH0S4gcBgPyoI1qd7HDvCW2Ok0NJuxAD4v"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) vRXJnbrSnjmVde0pRwBFtki2a2UurdkJ_CZPBSI3unZ5Ft04am1PoTfhH5wzaCSFvZVm6v7JlDyIsSOtlDH0S4gcBgPyoI1qd7HDvCW2Ok0NJuxAD4v7eW4xvcx_kszYFFCre3dQpv3QqhKl5YPrX_MywXc8Sm2cFeCd27fuDzD

ret = client.Seek(fd582, 129, client.SEEK_SET)
if(ret != 129) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 129)
  panic("seek failed")
}


ret = client.Seek(fd582, 154, client.SEEK_SET)
if(ret != 154) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 154)
  panic("seek failed")
}


ret = client.Seek(fd582, 158, client.SEEK_SET)
if(ret != 158) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 158)
  panic("seek failed")
}

//fd state: (158) vRXJnbrSnjmVde0pRwBFtki2a2UurdkJ_CZPBSI3unZ5Ft04am1PoTfhH5wzaCSFvZVm6v7JlDyIsSOtlDH0S4gcBgPyoI1qd7HDvCW2Ok0NJuxAD4v7eW4xvcx_kszYFFCre3dQpv3QqhKl5YPrX_MywXc8Sm2cFeCd27fuDzD

ret = client.Write(fd582, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (158) vRXJnbrSnjmVde0pRwBFtki2a2UurdkJ_CZPBSI3unZ5Ft04am1PoTfhH5wzaCSFvZVm6v7JlDyIsSOtlDH0S4gcBgPyoI1qd7HDvCW2Ok0NJuxAD4v7eW4xvcx_kszYFFCre3dQpv3QqhKl5YPrX_MywXc8Sm2cFeCd27fuDzD
//fd state: (158) vRXJnbrSnjmVde0pRwBFtki2a2UurdkJ_CZPBSI3unZ5Ft04am1PoTfhH5wzaCSFvZVm6v7JlDyIsSOtlDH0S4gcBgPyoI1qd7HDvCW2Ok0NJuxAD4v7eW4xvcx_kszYFFCre3dQpv3QqhKl5YPrX_MywXc8Sm2cFeCd27fuDzD

ret = client.Write(fd582, []byte("z0dXbVv_bp_fraLtTtiqnDi70FD_awhoBdHWldKsxZtNhq8AX5AaNO1tOUHpSUFiODuH0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (227) vRXJnbrSnjmVde0pRwBFtki2a2UurdkJ_CZPBSI3unZ5Ft04am1PoTfhH5wzaCSFvZVm6v7JlDyIsSOtlDH0S4gcBgPyoI1qd7HDvCW2Ok0NJuxAD4v7eW4xvcx_kszYFFCre3dQpv3QqhKl5YPrX_MywXc8Smz0dXbVv_bp_fraLtTtiqnDi70FD_awhoBdHWldKsxZtNhq8AX5AaNO1tOUHpSUFiODuH0

fd583 := client.Open("/7bbJM2y7gL.txt", client.O_RDWR|client.O_CREATE)
if(fd583 < 0) {
  panic("open failed")
}


ret = client.Close(fd583)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd582, 155, client.SEEK_SET)
if(ret != 155) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 155)
  panic("seek failed")
}


ret = client.Close(fd582)
if(ret != 0) {
  panic("close failed")
}


fd584 := client.Open("/ge1KupepMM.txt", client.O_RDWR|client.O_CREATE)
if(fd584 < 0) {
  panic("open failed")
}


ret = client.Seek(fd584, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd584, []byte("JAvJCao5BW5ZJdVwHbga88xb4ugEOZIGi5owtBv5HDzBRj1UwPN_F_azXGOvYrxXYFStVtVjvvX6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) JAvJCao5BW5ZJdVwHbga88xb4ugEOZIGi5owtBv5HDzBRj1UwPN_F_azXGOvYrxXYFStVtVjvvX6

ret = client.Seek(fd584, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


ret = client.Seek(fd584, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


buf, ret = client.Read(fd584, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "tVtVjvvX6") {
  panic("wrong data returned")
}

//fd state: (76) JAvJCao5BW5ZJdVwHbga88xb4ugEOZIGi5owtBv5HDzBRj1UwPN_F_azXGOvYrxXYFStVtVjvvX6

ret = client.Write(fd584, []byte("Z8xUeVX9EkudXLGsBcHbUQoIERfNbomdlpg0zdJvH3JBSPbXGbrHwCBXXDoD4242YmqnS5ozR5Nrx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (153) JAvJCao5BW5ZJdVwHbga88xb4ugEOZIGi5owtBv5HDzBRj1UwPN_F_azXGOvYrxXYFStVtVjvvX6Z8xUeVX9EkudXLGsBcHbUQoIERfNbomdlpg0zdJvH3JBSPbXGbrHwCBXXDoD4242YmqnS5ozR5Nrx

fd585 := client.Open("/0_jDN0GKli.txt", client.O_RDWR|client.O_CREATE)
if(fd585 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd585, []byte("KHgGg00XJURyKK4nWTOLQ0two6TZYZE2CAmxk84kaZai87QbmHU49uA6zOM7fgZjhNegOanXanE_HXIfZw6ilw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) KHgGg00XJURyKK4nWTOLQ0two6TZYZE2CAmxk84kaZai87QbmHU49uA6zOM7fgZjhNegOanXanE_HXIfZw6ilw

fd586 := client.Open("/SZ3w7aLh_u.txt", client.O_RDWR|client.O_CREATE)
if(fd586 < 0) {
  panic("open failed")
}


fd587 := client.Open("/FGjIWf9TaF.txt", client.O_RDWR|client.O_CREATE)
if(fd587 < 0) {
  panic("open failed")
}


ret = client.Seek(fd585, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


fd588 := client.Open("/cWx90GSelK.txt", client.O_RDWR|client.O_CREATE)
if(fd588 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd584, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd586)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd587)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd588)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd584, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd589 := client.Open("/iWExPMWmQy.txt", client.O_RDWR|client.O_CREATE)
if(fd589 < 0) {
  panic("open failed")
}


ret = client.Seek(fd585, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Seek(fd585, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


fd590 := client.Open("/VVprEbK7Bu.txt", client.O_RDWR|client.O_CREATE)
if(fd590 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd590, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "HRFrlIiuWPP59snHSeo1QyzgwJsxIMLzLBfYK7VZ4nXcgmtu2RRV2A0yBiv_DlXbYN") {
  panic("wrong data returned")
}


ret = client.Seek(fd585, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


buf, ret = client.Read(fd590, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JQ1D_C6GP") {
  panic("wrong data returned")
}

//fd state: (32) KHgGg00XJURyKK4nWTOLQ0two6TZYZE2CAmxk84kaZai87QbmHU49uA6zOM7fgZjhNegOanXanE_HXIfZw6ilw

ret = client.Write(fd585, []byte("Jl0WaV4UrtST8yhFLbwACI0MXc3LvYHlMryDegTKX0fm5gTEHaV3E71nwgPSP3VVF7g1Ko6k"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) KHgGg00XJURyKK4nWTOLQ0two6TZYZE2Jl0WaV4UrtST8yhFLbwACI0MXc3LvYHlMryDegTKX0fm5gTEHaV3E71nwgPSP3VVF7g1Ko6k

ret = client.Close(fd584)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd590)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) NbLXU71JgOEE2sFAXExb_KZGNZHgA8Hdx2MxCBMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lGkueqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYStXtZ2C_qG

ret = client.Write(fd589, []byte("SGj0tyfq6IFkdxFZqEvnNLFXOOh1PXL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) SGj0tyfq6IFkdxFZqEvnNLFXOOh1PXLdx2MxCBMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lGkueqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYStXtZ2C_qG

fd591 := client.Open("/kcHPGxVqjP.txt", client.O_RDWR|client.O_CREATE)
if(fd591 < 0) {
  panic("open failed")
}

//fd state: (104) KHgGg00XJURyKK4nWTOLQ0two6TZYZE2Jl0WaV4UrtST8yhFLbwACI0MXc3LvYHlMryDegTKX0fm5gTEHaV3E71nwgPSP3VVF7g1Ko6k

ret = client.Write(fd585, []byte("a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) KHgGg00XJURyKK4nWTOLQ0two6TZYZE2Jl0WaV4UrtST8yhFLbwACI0MXc3LvYHlMryDegTKX0fm5gTEHaV3E71nwgPSP3VVF7g1Ko6ka

buf, ret = client.Read(fd585, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (31) SGj0tyfq6IFkdxFZqEvnNLFXOOh1PXLdx2MxCBMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lGkueqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYStXtZ2C_qG

ret = client.Write(fd589, []byte("XXUpvRx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) SGj0tyfq6IFkdxFZqEvnNLFXOOh1PXLXXUpvRxMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lGkueqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYStXtZ2C_qG

buf, ret = client.Read(fd591, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (38) SGj0tyfq6IFkdxFZqEvnNLFXOOh1PXLXXUpvRxMT6yvqHmnAnRW_OEVZAD6Lr3L0k4goV1ISTtf8i5lGkueqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYStXtZ2C_qG

ret = client.Write(fd589, []byte("g_JaPdc7tkYx6Mr8x_Coi7V1QLY9sKmbka1_NhRFZ1Nx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) SGj0tyfq6IFkdxFZqEvnNLFXOOh1PXLXXUpvRxg_JaPdc7tkYx6Mr8x_Coi7V1QLY9sKmbka1_NhRFZ1NxeqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYStXtZ2C_qG

buf, ret = client.Read(fd591, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd589, 155, client.SEEK_SET)
if(ret != 155) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 155)
  panic("seek failed")
}


buf, ret = client.Read(fd585, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd585, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (155) SGj0tyfq6IFkdxFZqEvnNLFXOOh1PXLXXUpvRxg_JaPdc7tkYx6Mr8x_Coi7V1QLY9sKmbka1_NhRFZ1NxeqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYStXtZ2C_qG

ret = client.Write(fd589, []byte("x0PYbAYwAyi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (166) SGj0tyfq6IFkdxFZqEvnNLFXOOh1PXLXXUpvRxg_JaPdc7tkYx6Mr8x_Coi7V1QLY9sKmbka1_NhRFZ1NxeqONqm_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiytkZbSXwJeELYZfZryNDQYStXx0PYbAYwAyi

fd592 := client.Open("/8Tqb7q4mha.txt", client.O_RDWR|client.O_CREATE)
if(fd592 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd591, []byte("vek5yBnmqCyeowwng1hCtBFC7WpEALfVVyV9qrdXdDdVv9TsXsgfbDhE4nzm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) vek5yBnmqCyeowwng1hCtBFC7WpEALfVVyV9qrdXdDdVv9TsXsgfbDhE4nzm

fd593 := client.Open("/nMC82xZiaO.txt", client.O_RDWR|client.O_CREATE)
if(fd593 < 0) {
  panic("open failed")
}


fd594 := client.Open("/ja3oo3k0R9.txt", client.O_RDWR|client.O_CREATE)
if(fd594 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd593, []byte("sZlGie3zd69HcAkHRaUsQH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) sZlGie3zd69HcAkHRaUsQH

buf, ret = client.Read(fd589, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd589, 64, client.SEEK_SET)
if(ret != 64) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 64)
  panic("seek failed")
}


ret = client.Close(fd589)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd585, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd592, []byte("l5prdy7UuwUWD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) l5prdy7UuwUWD

buf, ret = client.Read(fd593, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd592, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Close(fd594)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd591, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd595 := client.Open("/Xa2of_COvq.txt", client.O_RDWR|client.O_CREATE)
if(fd595 < 0) {
  panic("open failed")
}


fd596 := client.Open("/raavVeJ6eA.txt", client.O_RDWR|client.O_CREATE)
if(fd596 < 0) {
  panic("open failed")
}

//fd state: (22) sZlGie3zd69HcAkHRaUsQH

ret = client.Write(fd593, []byte("nC_AXipwz2XsmuIBdP7QbahWIlTx250gXl54RDebL4kQ63FWJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (71) sZlGie3zd69HcAkHRaUsQHnC_AXipwz2XsmuIBdP7QbahWIlTx250gXl54RDebL4kQ63FWJ
//fd state: (71) sZlGie3zd69HcAkHRaUsQHnC_AXipwz2XsmuIBdP7QbahWIlTx250gXl54RDebL4kQ63FWJ

ret = client.Write(fd593, []byte("qxdULafFclh5J0fxWjn05Uurhqjqgwy6zBK8EXSh7xKTXw02AScJ9myp1LGrlkHH1u51uEcyX75Pen5lu2SRZg52"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (159) sZlGie3zd69HcAkHRaUsQHnC_AXipwz2XsmuIBdP7QbahWIlTx250gXl54RDebL4kQ63FWJqxdULafFclh5J0fxWjn05Uurhqjqgwy6zBK8EXSh7xKTXw02AScJ9myp1LGrlkHH1u51uEcyX75Pen5lu2SRZg52

buf, ret = client.Read(fd596, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LsX2upRewHuGPKOI6otReaIgSHqD10MRV8hKSrxLqSQdNbnRugD0gFnKmqlDG") {
  panic("wrong data returned")
}

//fd state: (159) sZlGie3zd69HcAkHRaUsQHnC_AXipwz2XsmuIBdP7QbahWIlTx250gXl54RDebL4kQ63FWJqxdULafFclh5J0fxWjn05Uurhqjqgwy6zBK8EXSh7xKTXw02AScJ9myp1LGrlkHH1u51uEcyX75Pen5lu2SRZg52

ret = client.Write(fd593, []byte("bqJCP23QF9W2dMWKa8GOKT9q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (183) sZlGie3zd69HcAkHRaUsQHnC_AXipwz2XsmuIBdP7QbahWIlTx250gXl54RDebL4kQ63FWJqxdULafFclh5J0fxWjn05Uurhqjqgwy6zBK8EXSh7xKTXw02AScJ9myp1LGrlkHH1u51uEcyX75Pen5lu2SRZg52bqJCP23QF9W2dMWKa8GOKT9q

ret = client.Close(fd593)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd595, []byte("j9AakDYS1IM1f7WVdiV5AhuZuLslMPFQU1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) j9AakDYS1IM1f7WVdiV5AhuZuLslMPFQU1

buf, ret = client.Read(fd591, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd585)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd591, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


fd597 := client.Open("/_H9N9W_3zL.txt", client.O_RDWR|client.O_CREATE)
if(fd597 < 0) {
  panic("open failed")
}


ret = client.Seek(fd597, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd596)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd597, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd591)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd597, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd598 := client.Open("/LWBzlj5ZUH.txt", client.O_RDWR|client.O_CREATE)
if(fd598 < 0) {
  panic("open failed")
}


ret = client.Seek(fd598, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd597, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd598)
if(ret != 0) {
  panic("close failed")
}

//fd state: (34) j9AakDYS1IM1f7WVdiV5AhuZuLslMPFQU1

ret = client.Write(fd595, []byte("V9rKWecAG2ynUcctKw7gdWdBHIDUV6opqsBmxV1aUTk120vZhnIuPNuBj7wv3ySoE9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) j9AakDYS1IM1f7WVdiV5AhuZuLslMPFQU1V9rKWecAG2ynUcctKw7gdWdBHIDUV6opqsBmxV1aUTk120vZhnIuPNuBj7wv3ySoE9

ret = client.Seek(fd597, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd597, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd597)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd595, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd599 := client.Open("/Wuhr1LEWJy.txt", client.O_RDWR|client.O_CREATE)
if(fd599 < 0) {
  panic("open failed")
}

//fd state: (5) l5prdy7UuwUWD

ret = client.Write(fd592, []byte("OIXN2EmxXns80fLInuXSAkgmgwyHHNPcN89lOX0xYmz4hPLpJXhE3OlZx0mPfFgfHEPVPw4OTIJP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) l5prdOIXN2EmxXns80fLInuXSAkgmgwyHHNPcN89lOX0xYmz4hPLpJXhE3OlZx0mPfFgfHEPVPw4OTIJP

ret = client.Seek(fd592, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Seek(fd599, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd592, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "s80fLInuXSAkgmgwyHHNPcN89lOX0xYmz") {
  panic("wrong data returned")
}


ret = client.Close(fd599)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd595, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd595, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd592)
if(ret != 0) {
  panic("close failed")
}


fd600 := client.Open("/JPFo4bf_mL.txt", client.O_RDWR|client.O_CREATE)
if(fd600 < 0) {
  panic("open failed")
}


ret = client.Close(fd600)
if(ret != 0) {
  panic("close failed")
}

//fd state: (11) j9AakDYS1IM1f7WVdiV5AhuZuLslMPFQU1V9rKWecAG2ynUcctKw7gdWdBHIDUV6opqsBmxV1aUTk120vZhnIuPNuBj7wv3ySoE9

ret = client.Write(fd595, []byte("6zKKObFbnS13jjqXT8EFMGiweHtx4kaAFy552hwthOF6wz82p3KWDv1HM53WXLpemh1OQ6DON4iTTYdRyTJjhzN0Jg3qlf_KPwk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) j9AakDYS1IM6zKKObFbnS13jjqXT8EFMGiweHtx4kaAFy552hwthOF6wz82p3KWDv1HM53WXLpemh1OQ6DON4iTTYdRyTJjhzN0Jg3qlf_KPwk

buf, ret = client.Read(fd595, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd601 := client.Open("/4nOzh3S2jn.txt", client.O_RDWR|client.O_CREATE)
if(fd601 < 0) {
  panic("open failed")
}


ret = client.Close(fd595)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd601, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd601)
if(ret != 0) {
  panic("close failed")
}


fd602 := client.Open("/epwKfrBhnk.txt", client.O_RDWR|client.O_CREATE)
if(fd602 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd602, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd603 := client.Open("/9TFyXXeDd1.txt", client.O_RDWR|client.O_CREATE)
if(fd603 < 0) {
  panic("open failed")
}


ret = client.Close(fd602)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd603, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd603)
if(ret != 0) {
  panic("close failed")
}


fd604 := client.Open("/PHNZPMwu3G.txt", client.O_RDWR|client.O_CREATE)
if(fd604 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd604, []byte("H694BHWWEwlXR8mlKvj_McVeU3uWmP37kYiu5slXQiJi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) H694BHWWEwlXR8mlKvj_McVeU3uWmP37kYiu5slXQiJi

ret = client.Seek(fd604, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


ret = client.Close(fd604)
if(ret != 0) {
  panic("close failed")
}


fd605 := client.Open("/R2ldcPasbO.txt", client.O_RDWR|client.O_CREATE)
if(fd605 < 0) {
  panic("open failed")
}


ret = client.Close(fd605)
if(ret != 0) {
  panic("close failed")
}


fd606 := client.Open("/Z34Wgzj0C3.txt", client.O_RDWR|client.O_CREATE)
if(fd606 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd606, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd606, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd607 := client.Open("/8fT6sTxzZ9.txt", client.O_RDWR|client.O_CREATE)
if(fd607 < 0) {
  panic("open failed")
}


ret = client.Close(fd607)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd606, []byte("YGRy_KeYPvbdfeKmaK4fzGKx6qaF8xEOpVRa9yEjE1jUXWEGe3bI14QaAWia3uh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) YGRy_KeYPvbdfeKmaK4fzGKx6qaF8xEOpVRa9yEjE1jUXWEGe3bI14QaAWia3uh

ret = client.Seek(fd606, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


ret = client.Close(fd606)
if(ret != 0) {
  panic("close failed")
}


fd608 := client.Open("/ZJKKvj2wHY.txt", client.O_RDWR|client.O_CREATE)
if(fd608 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd608, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd609 := client.Open("/iwLLSjsyMF.txt", client.O_RDWR|client.O_CREATE)
if(fd609 < 0) {
  panic("open failed")
}


ret = client.Seek(fd609, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


fd610 := client.Open("/VupjQasWWX.txt", client.O_RDWR|client.O_CREATE)
if(fd610 < 0) {
  panic("open failed")
}


fd611 := client.Open("/UM7bUFua9_.txt", client.O_RDWR|client.O_CREATE)
if(fd611 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd609, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "k6fHLhV6K7marRaWJd9p48HgZe") {
  panic("wrong data returned")
}


ret = client.Seek(fd608, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd609)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd611, []byte("AbPC79WZrF0qG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) AbPC79WZrF0qG
//fd state: (0) 

ret = client.Write(fd608, []byte("2HTaxAYs4W7_G3vsNQFqJl3xtsgNCYaLQ3f6ZDTHuIKHLgnKhphiiKxZJKQx6bqMGnYNAd3nFUH0kdFoqtj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) 2HTaxAYs4W7_G3vsNQFqJl3xtsgNCYaLQ3f6ZDTHuIKHLgnKhphiiKxZJKQx6bqMGnYNAd3nFUH0kdFoqtj
//fd state: (83) 2HTaxAYs4W7_G3vsNQFqJl3xtsgNCYaLQ3f6ZDTHuIKHLgnKhphiiKxZJKQx6bqMGnYNAd3nFUH0kdFoqtj

ret = client.Write(fd608, []byte("EckaaQ_skfP6FXSVGzZdaqVnMJWsAmlEC729SsKQZ9YViRScIzw0sqt7WJivPQlvCDFWt62J0zdosm8tnphXU4gQw8I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (174) 2HTaxAYs4W7_G3vsNQFqJl3xtsgNCYaLQ3f6ZDTHuIKHLgnKhphiiKxZJKQx6bqMGnYNAd3nFUH0kdFoqtjEckaaQ_skfP6FXSVGzZdaqVnMJWsAmlEC729SsKQZ9YViRScIzw0sqt7WJivPQlvCDFWt62J0zdosm8tnphXU4gQw8I

fd612 := client.Open("/vCY8hmzxY1.txt", client.O_RDWR|client.O_CREATE)
if(fd612 < 0) {
  panic("open failed")
}


fd613 := client.Open("/FS0Irk9hhz.txt", client.O_RDWR|client.O_CREATE)
if(fd613 < 0) {
  panic("open failed")
}


fd614 := client.Open("/ogbVTic8um.txt", client.O_RDWR|client.O_CREATE)
if(fd614 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd614, []byte("jYAFkcma4Mel1mbS2m9GE0ZR6xCQDVIU1ZPuNldRzzG3o25chya0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) jYAFkcma4Mel1mbS2m9GE0ZR6xCQDVIU1ZPuNldRzzG3o25chya0

ret = client.Close(fd608)
if(ret != 0) {
  panic("close failed")
}


fd615 := client.Open("/ZOIAW57u8n.txt", client.O_RDWR|client.O_CREATE)
if(fd615 < 0) {
  panic("open failed")
}


ret = client.Seek(fd614, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


fd616 := client.Open("/ApkZMhqPog.txt", client.O_RDWR|client.O_CREATE)
if(fd616 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd613, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8tvZirCd71oK5O9A33mWu1IyugSANxpK0IIVeWAPxnsZnDzYTlfTA6bBx4Y") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd615, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd610, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd612, []byte("wMFLhfhoRKOlVEaQOGMeUbQkXRiintGt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) wMFLhfhoRKOlVEaQOGMeUbQkXRiintGt

fd617 := client.Open("/ZkuXs_QYnQ.txt", client.O_RDWR|client.O_CREATE)
if(fd617 < 0) {
  panic("open failed")
}


ret = client.Seek(fd614, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


buf, ret = client.Read(fd617, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "RjvXrtHhNrBlRD6F7gM0K5Zz") {
  panic("wrong data returned")
}


ret = client.Seek(fd617, 106, client.SEEK_SET)
if(ret != 106) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 106)
  panic("seek failed")
}

//fd state: (32) wMFLhfhoRKOlVEaQOGMeUbQkXRiintGt

ret = client.Write(fd612, []byte("5sA_f3bhjqOB3YyQySUl2dcAxKj94iNheyHQ5cuX6HgHsDiCeKUxvhG9xKBX6jk8heN4QFJnlqQ8b8x_fXeSpvWD4ZD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) wMFLhfhoRKOlVEaQOGMeUbQkXRiintGt5sA_f3bhjqOB3YyQySUl2dcAxKj94iNheyHQ5cuX6HgHsDiCeKUxvhG9xKBX6jk8heN4QFJnlqQ8b8x_fXeSpvWD4ZD

ret = client.Close(fd613)
if(ret != 0) {
  panic("close failed")
}


fd618 := client.Open("/x4uQZLfO13.txt", client.O_RDWR|client.O_CREATE)
if(fd618 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd610, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd612)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd610, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd619 := client.Open("/t2Dyv3YWoA.txt", client.O_RDWR|client.O_CREATE)
if(fd619 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd610, []byte("JTO0A2OMdDP682E25I1_7tvHXNaEke6ZZh3I86QnrsTr8PCnPCRKpIybfxE8tMt4vj92en_KE_g99pi32W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) JTO0A2OMdDP682E25I1_7tvHXNaEke6ZZh3I86QnrsTr8PCnPCRKpIybfxE8tMt4vj92en_KE_g99pi32W

fd620 := client.Open("/s5ORu3HSjk.txt", client.O_RDWR|client.O_CREATE)
if(fd620 < 0) {
  panic("open failed")
}


ret = client.Seek(fd614, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


ret = client.Seek(fd614, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd618)
if(ret != 0) {
  panic("close failed")
}


fd621 := client.Open("/HeMi5T8fXN.txt", client.O_RDWR|client.O_CREATE)
if(fd621 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd619, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd616, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


fd622 := client.Open("/gMb5eMbeqb.txt", client.O_RDWR|client.O_CREATE)
if(fd622 < 0) {
  panic("open failed")
}


ret = client.Seek(fd615, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd615, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd615, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd621, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd610)
if(ret != 0) {
  panic("close failed")
}

//fd state: (12) jYAFkcma4Mel1mbS2m9GE0ZR6xCQDVIU1ZPuNldRzzG3o25chya0

ret = client.Write(fd614, []byte("BfWJBkgP_zdwdk39taj5lPtLJTFrZ6Ply0L8txBUvrlDeFDWe34bBXZrfesK0p2sLJFr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) jYAFkcma4MelBfWJBkgP_zdwdk39taj5lPtLJTFrZ6Ply0L8txBUvrlDeFDWe34bBXZrfesK0p2sLJFr
//fd state: (0) 

ret = client.Write(fd620, []byte("u8guiERY5i2ocaOvqRQWPrgqwFoHINzMT9Bcxu_5_YlpVLvxvDgjw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) u8guiERY5i2ocaOvqRQWPrgqwFoHINzMT9Bcxu_5_YlpVLvxvDgjw
//fd state: (0) 

ret = client.Write(fd619, []byte("DBC3IlGDA976MMhhmr38uwnMzYuScQhgRUN4XB0CYcrVF0ptAzLzExOx2PHbB8r3Zo1DDfywhwJ4DZwHwnKz5OdjTQ37vG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) DBC3IlGDA976MMhhmr38uwnMzYuScQhgRUN4XB0CYcrVF0ptAzLzExOx2PHbB8r3Zo1DDfywhwJ4DZwHwnKz5OdjTQ37vG
//fd state: (80) jYAFkcma4MelBfWJBkgP_zdwdk39taj5lPtLJTFrZ6Ply0L8txBUvrlDeFDWe34bBXZrfesK0p2sLJFr

ret = client.Write(fd614, []byte("XUu0qWVh5f9_b15XFRAgINWDK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) jYAFkcma4MelBfWJBkgP_zdwdk39taj5lPtLJTFrZ6Ply0L8txBUvrlDeFDWe34bBXZrfesK0p2sLJFrXUu0qWVh5f9_b15XFRAgINWDK

ret = client.Close(fd615)
if(ret != 0) {
  panic("close failed")
}

//fd state: (105) jYAFkcma4MelBfWJBkgP_zdwdk39taj5lPtLJTFrZ6Ply0L8txBUvrlDeFDWe34bBXZrfesK0p2sLJFrXUu0qWVh5f9_b15XFRAgINWDK

ret = client.Write(fd614, []byte("CEqmGvson6V1WWlIQVkS9I7dYHxPy6uEpIx14wv55KunpEp1Vi2CmEIiII6R4NZF4maL6MPpk0sE6G5vv2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (187) jYAFkcma4MelBfWJBkgP_zdwdk39taj5lPtLJTFrZ6Ply0L8txBUvrlDeFDWe34bBXZrfesK0p2sLJFrXUu0qWVh5f9_b15XFRAgINWDKCEqmGvson6V1WWlIQVkS9I7dYHxPy6uEpIx14wv55KunpEp1Vi2CmEIiII6R4NZF4maL6MPpk0sE6G5vv2

fd623 := client.Open("/2KEc6YkBum.txt", client.O_RDWR|client.O_CREATE)
if(fd623 < 0) {
  panic("open failed")
}


ret = client.Seek(fd621, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd616, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "6d_FEIu2itMbAKA2vHuiuiTO") {
  panic("wrong data returned")
}


ret = client.Seek(fd621, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd624 := client.Open("/gfY_Y9jHYW.txt", client.O_RDWR|client.O_CREATE)
if(fd624 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd625 := client.Open("/UmfTLCMLDx.txt", client.O_RDWR|client.O_CREATE)
if(fd625 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd622, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd626 := client.Open("/08lzRI0lEO.txt", client.O_RDWR|client.O_CREATE)
if(fd626 < 0) {
  panic("open failed")
}


ret = client.Close(fd611)
if(ret != 0) {
  panic("close failed")
}


fd627 := client.Open("/ezy6lpPkBl.txt", client.O_RDWR|client.O_CREATE)
if(fd627 < 0) {
  panic("open failed")
}

//fd state: (0) 4NqJsTuzXef2ZvwVKx4PIwT8nvF3e2USjQ5mHBjzgKyMRWKXXrQDI6MV3vTewyltVR3XocNM4UfFOIZF5YJmpPCtZforMnXxSCG9ECZbuVy82URixB8w_rMJ0BfSBpF08jic0n2t1uuh4BmMloS0iQLNYMTavQ7Tp4v0IQO

ret = client.Write(fd623, []byte("H797mR4acFjxfWvyDt8_kaVva6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) H797mR4acFjxfWvyDt8_kaVva6F3e2USjQ5mHBjzgKyMRWKXXrQDI6MV3vTewyltVR3XocNM4UfFOIZF5YJmpPCtZforMnXxSCG9ECZbuVy82URixB8w_rMJ0BfSBpF08jic0n2t1uuh4BmMloS0iQLNYMTavQ7Tp4v0IQO
//fd state: (0) 

ret = client.Write(fd626, []byte("jaWRIH8cglLg6w84lUtpKF_BJO3gKIxu5FyLTzagU513q6d873Tl97jiMBDckB189azuhP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) jaWRIH8cglLg6w84lUtpKF_BJO3gKIxu5FyLTzagU513q6d873Tl97jiMBDckB189azuhP
//fd state: (26) H797mR4acFjxfWvyDt8_kaVva6F3e2USjQ5mHBjzgKyMRWKXXrQDI6MV3vTewyltVR3XocNM4UfFOIZF5YJmpPCtZforMnXxSCG9ECZbuVy82URixB8w_rMJ0BfSBpF08jic0n2t1uuh4BmMloS0iQLNYMTavQ7Tp4v0IQO

ret = client.Write(fd623, []byte("shaLjOd1N32dBBFylPUgRYWjecDQmz0J84tHAGHqLE9BUDjRCF3yxrz_dSxtDXnd7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) H797mR4acFjxfWvyDt8_kaVva6shaLjOd1N32dBBFylPUgRYWjecDQmz0J84tHAGHqLE9BUDjRCF3yxrz_dSxtDXnd7rMnXxSCG9ECZbuVy82URixB8w_rMJ0BfSBpF08jic0n2t1uuh4BmMloS0iQLNYMTavQ7Tp4v0IQO
//fd state: (0) 

ret = client.Write(fd627, []byte("kYnAxWLSSyz0xoYPvGlL8HPSRS5fbdMRA8o"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) kYnAxWLSSyz0xoYPvGlL8HPSRS5fbdMRA8o

buf, ret = client.Read(fd620, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd627, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd628 := client.Open("/jH8lev2VAs.txt", client.O_RDWR|client.O_CREATE)
if(fd628 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd617, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5NE00yl40kbjAxSK4") {
  panic("wrong data returned")
}


fd629 := client.Open("/08lzRI0lEO.txt", client.O_RDWR|client.O_CREATE)
if(fd629 < 0) {
  panic("open failed")
}


fd630 := client.Open("/ZOIAW57u8n.txt", client.O_RDWR|client.O_CREATE)
if(fd630 < 0) {
  panic("open failed")
}


ret = client.Seek(fd616, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


ret = client.Close(fd627)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd629)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd628, []byte("HNaAl36BDTyUn5bjDSne"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) HNaAl36BDTyUn5bjDSne

ret = client.Close(fd622)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd628)
if(ret != 0) {
  panic("close failed")
}


fd631 := client.Open("/MLw9D0Mx68.txt", client.O_RDWR|client.O_CREATE)
if(fd631 < 0) {
  panic("open failed")
}


ret = client.Seek(fd617, 98, client.SEEK_SET)
if(ret != 98) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 98)
  panic("seek failed")
}


ret = client.Seek(fd631, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd626, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


buf, ret = client.Read(fd614, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd632 := client.Open("/OqvlaCPI05.txt", client.O_RDWR|client.O_CREATE)
if(fd632 < 0) {
  panic("open failed")
}


ret = client.Seek(fd620, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}


ret = client.Seek(fd625, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd630, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd617, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pkiyklff5NE00yl40kbjAxSK4") {
  panic("wrong data returned")
}


fd633 := client.Open("/FSTYY4xE5x.txt", client.O_RDWR|client.O_CREATE)
if(fd633 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd617, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd630, []byte("azjcA2yPem5kHRUpg3q9JvK_ShXl8H60XimmidVsdWPJKFoArcUvtZYYlRCPQ8BmT2O0xUJ1K4dXk7eMjQxc21TvsGyB3Wdz2S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) azjcA2yPem5kHRUpg3q9JvK_ShXl8H60XimmidVsdWPJKFoArcUvtZYYlRCPQ8BmT2O0xUJ1K4dXk7eMjQxc21TvsGyB3Wdz2S

buf, ret = client.Read(fd616, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2itMbAKA2vHuiuiTO") {
  panic("wrong data returned")
}


ret = client.Seek(fd625, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd620)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd626)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd630, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd614, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd634 := client.Open("/U27uigcAVw.txt", client.O_RDWR|client.O_CREATE)
if(fd634 < 0) {
  panic("open failed")
}


ret = client.Seek(fd633, 122, client.SEEK_SET)
if(ret != 122) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 122)
  panic("seek failed")
}


buf, ret = client.Read(fd630, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd633)
if(ret != 0) {
  panic("close failed")
}


fd635 := client.Open("/Z34Wgzj0C3.txt", client.O_RDWR|client.O_CREATE)
if(fd635 < 0) {
  panic("open failed")
}


ret = client.Seek(fd621, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd614)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd624, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd636 := client.Open("/wNZPweOOek.txt", client.O_RDWR|client.O_CREATE)
if(fd636 < 0) {
  panic("open failed")
}


ret = client.Seek(fd619, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd624, []byte("1qU_wlfuJSgy1tUgBnhr6RSWnRn8_cnAX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) 1qU_wlfuJSgy1tUgBnhr6RSWnRn8_cnAX

ret = client.Seek(fd616, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


buf, ret = client.Read(fd623, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rMnXxSC") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd634, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd624)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd617)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd625, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd637 := client.Open("/Crr1MV2zwA.txt", client.O_RDWR|client.O_CREATE)
if(fd637 < 0) {
  panic("open failed")
}


ret = client.Close(fd621)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd634, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd623, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "G9ECZbuVy82URixB8w_rMJ0BfSBpF08jic0n2t1uuh4BmMloS0iQLNY") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd632, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "44jkIiTN6S_iy9Ilosat9l80Z_k76VheqslktiFBFV57y0F49QH9Mr_yVQCEyEpKJsYqaqkj_4psDdn2a") {
  panic("wrong data returned")
}


fd638 := client.Open("/XYg89mK8WS.txt", client.O_RDWR|client.O_CREATE)
if(fd638 < 0) {
  panic("open failed")
}


ret = client.Close(fd619)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd637)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd632, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uoQz06Yu795i6rfYnI") {
  panic("wrong data returned")
}


fd639 := client.Open("/PIpkJxo4gp.txt", client.O_RDWR|client.O_CREATE)
if(fd639 < 0) {
  panic("open failed")
}


ret = client.Seek(fd636, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd634, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd635)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd630)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd631)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd625, []byte("lbqC8mRwmwg0bVMRbEKY4Xc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) lbqC8mRwmwg0bVMRbEKY4Xc

ret = client.Close(fd636)
if(ret != 0) {
  panic("close failed")
}


fd640 := client.Open("/oJi1_Yfhik.txt", client.O_RDWR|client.O_CREATE)
if(fd640 < 0) {
  panic("open failed")
}


fd641 := client.Open("/nMC82xZiaO.txt", client.O_RDWR|client.O_CREATE)
if(fd641 < 0) {
  panic("open failed")
}


fd642 := client.Open("/C2X2B8VghN.txt", client.O_RDWR|client.O_CREATE)
if(fd642 < 0) {
  panic("open failed")
}


fd643 := client.Open("/3NpJGV4eDd.txt", client.O_RDWR|client.O_CREATE)
if(fd643 < 0) {
  panic("open failed")
}


ret = client.Close(fd641)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd623, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MTavQ7Tp4v0IQO") {
  panic("wrong data returned")
}


ret = client.Seek(fd634, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd616)
if(ret != 0) {
  panic("close failed")
}


fd644 := client.Open("/LWBzlj5ZUH.txt", client.O_RDWR|client.O_CREATE)
if(fd644 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd642, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd643, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd640)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd644)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd638)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd632, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd623, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


fd645 := client.Open("/ji0cGwVP1l.txt", client.O_RDWR|client.O_CREATE)
if(fd645 < 0) {
  panic("open failed")
}


ret = client.Seek(fd643, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd646 := client.Open("/DGye9zYlsd.txt", client.O_RDWR|client.O_CREATE)
if(fd646 < 0) {
  panic("open failed")
}

//fd state: (0) JuomKku8QhOHffZySIVq26a93PyaZmbwtuJPEwUOOA97lyT7fk9ttEgvsjWo9BaqVCNezCLhhI0xvSWU4Nn3

ret = client.Write(fd639, []byte("5jR1kCPm4aRGJg05eFQbL7vRZw_bv8JcOWhId9JmYWUEVw7kdeZjwHLzgSc5E5jZ8ihbAlUcEdfIxxpoz84f71f5UP4e_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) 5jR1kCPm4aRGJg05eFQbL7vRZw_bv8JcOWhId9JmYWUEVw7kdeZjwHLzgSc5E5jZ8ihbAlUcEdfIxxpoz84f71f5UP4e_

fd647 := client.Open("/enZ6N2BD1f.txt", client.O_RDWR|client.O_CREATE)
if(fd647 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd625, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd632, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd625, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd648 := client.Open("/zvfpB2u_jD.txt", client.O_RDWR|client.O_CREATE)
if(fd648 < 0) {
  panic("open failed")
}


fd649 := client.Open("/nqFTzh7bC4.txt", client.O_RDWR|client.O_CREATE)
if(fd649 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd632, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd623, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "84tHAGHqLE9BUDjRCF3yxrz_dSxtDXnd7rMnXxSCG9ECZbuVy82URixB") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd645, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd643)
if(ret != 0) {
  panic("close failed")
}


fd650 := client.Open("/74mEwYxfha.txt", client.O_RDWR|client.O_CREATE)
if(fd650 < 0) {
  panic("open failed")
}


ret = client.Seek(fd645, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd634, []byte("n8vuiZuGQNS1JqoBjHpTvQ9kMfrnpfaKxo3CqbvJ_ib5xfVxWJbCtQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) n8vuiZuGQNS1JqoBjHpTvQ9kMfrnpfaKxo3CqbvJ_ib5xfVxWJbCtQ
//fd state: (99) 44jkIiTN6S_iy9Ilosat9l80Z_k76VheqslktiFBFV57y0F49QH9Mr_yVQCEyEpKJsYqaqkj_4psDdn2auoQz06Yu795i6rfYnI

ret = client.Write(fd632, []byte("haJLzc03phoHXh3n13CUcRJIhpvqX9hYJ6F_HXtC6Yf2wHrOjWAl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) 44jkIiTN6S_iy9Ilosat9l80Z_k76VheqslktiFBFV57y0F49QH9Mr_yVQCEyEpKJsYqaqkj_4psDdn2auoQz06Yu795i6rfYnIhaJLzc03phoHXh3n13CUcRJIhpvqX9hYJ6F_HXtC6Yf2wHrOjWAl
//fd state: (3) lbqC8mRwmwg0bVMRbEKY4Xc

ret = client.Write(fd625, []byte("KLCCYhA_XDpbuI5fz0HLKGdcGJPnltzZkBUlBf46HLM327YZPCZB0wnKVsibAA_YxnN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) lbqKLCCYhA_XDpbuI5fz0HLKGdcGJPnltzZkBUlBf46HLM327YZPCZB0wnKVsibAA_YxnN

ret = client.Seek(fd648, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}

//fd state: (0) HvRa2ENPhB3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaAYTtNueQjzs9GvnM146GOEVteJffSDvKyMaSfioXjUPfo4iz7aakQ4IenrGLNZlPnIodCODHE4bLw7m3qDbrtGk9aZlLBM2SE0MHFG1RKPMI0KSJyFlJ0HFdT7tWSS_dvYgD4qf

ret = client.Write(fd650, []byte("MINo6zG7ZS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) MINo6zG7ZS3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaAYTtNueQjzs9GvnM146GOEVteJffSDvKyMaSfioXjUPfo4iz7aakQ4IenrGLNZlPnIodCODHE4bLw7m3qDbrtGk9aZlLBM2SE0MHFG1RKPMI0KSJyFlJ0HFdT7tWSS_dvYgD4qf

ret = client.Seek(fd648, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}


buf, ret = client.Read(fd646, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd646, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (70) lbqKLCCYhA_XDpbuI5fz0HLKGdcGJPnltzZkBUlBf46HLM327YZPCZB0wnKVsibAA_YxnN

ret = client.Write(fd625, []byte("sfAAJAwQqJzW0GvZEV2ebxODR0cd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) lbqKLCCYhA_XDpbuI5fz0HLKGdcGJPnltzZkBUlBf46HLM327YZPCZB0wnKVsibAA_YxnNsfAAJAwQqJzW0GvZEV2ebxODR0cd

ret = client.Seek(fd642, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd645, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd651 := client.Open("/iEctF_9T1b.txt", client.O_RDWR|client.O_CREATE)
if(fd651 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd650, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaAYTtNueQjzs9GvnM146GOEVteJffSDvK") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd647, []byte("6cEzTGsvmH5p9XXUNW4mvkOW1saO6s8ai099pPjLE2ZJwjzw2JkmG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) 6cEzTGsvmH5p9XXUNW4mvkOW1saO6s8ai099pPjLE2ZJwjzw2JkmG

buf, ret = client.Read(fd651, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd625, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd652 := client.Open("/rG17EcN2Dj.txt", client.O_RDWR|client.O_CREATE)
if(fd652 < 0) {
  panic("open failed")
}


fd653 := client.Open("/QVh9ftowiN.txt", client.O_RDWR|client.O_CREATE)
if(fd653 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd652, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd651, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd654 := client.Open("/mmm4TPjG41.txt", client.O_RDWR|client.O_CREATE)
if(fd654 < 0) {
  panic("open failed")
}


fd655 := client.Open("/TXZ9KqgAYy.txt", client.O_RDWR|client.O_CREATE)
if(fd655 < 0) {
  panic("open failed")
}


ret = client.Seek(fd634, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


fd656 := client.Open("/rp8DgXYLb4.txt", client.O_RDWR|client.O_CREATE)
if(fd656 < 0) {
  panic("open failed")
}


fd657 := client.Open("/MFvyYPutcX.txt", client.O_RDWR|client.O_CREATE)
if(fd657 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd646, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd657)
if(ret != 0) {
  panic("close failed")
}

//fd state: (114) H797mR4acFjxfWvyDt8_kaVva6shaLjOd1N32dBBFylPUgRYWjecDQmz0J84tHAGHqLE9BUDjRCF3yxrz_dSxtDXnd7rMnXxSCG9ECZbuVy82URixB8w_rMJ0BfSBpF08jic0n2t1uuh4BmMloS0iQLNYMTavQ7Tp4v0IQO

ret = client.Write(fd623, []byte("4YoqExRJtHeDiZyh9r8BWQHEtFdVRjCfXDzGVl5LqNn5pXU5R1K58ifBaF9iQKmKB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (179) H797mR4acFjxfWvyDt8_kaVva6shaLjOd1N32dBBFylPUgRYWjecDQmz0J84tHAGHqLE9BUDjRCF3yxrz_dSxtDXnd7rMnXxSCG9ECZbuVy82URixB4YoqExRJtHeDiZyh9r8BWQHEtFdVRjCfXDzGVl5LqNn5pXU5R1K58ifBaF9iQKmKB

ret = client.Close(fd642)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd655, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd634)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd653)
if(ret != 0) {
  panic("close failed")
}


fd658 := client.Open("/h6MQl3Ciwy.txt", client.O_RDWR|client.O_CREATE)
if(fd658 < 0) {
  panic("open failed")
}


fd659 := client.Open("/FvKjuV322b.txt", client.O_RDWR|client.O_CREATE)
if(fd659 < 0) {
  panic("open failed")
}


fd660 := client.Open("/R8FiGZHyRq.txt", client.O_RDWR|client.O_CREATE)
if(fd660 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd654, []byte("0SamE4NSuceSkjrqXYwz9_Gq2CcoXELXHJPmPHbBBWSiWj3tjUrqWReMFJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) 0SamE4NSuceSkjrqXYwz9_Gq2CcoXELXHJPmPHbBBWSiWj3tjUrqWReMFJ

ret = client.Seek(fd646, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd652, []byte("1XVlIE5JBGVVbc4p_YTZg4RYX3XKYuuZtrrlomWgX4kHnTYE9fydiRdCrml9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) 1XVlIE5JBGVVbc4p_YTZg4RYX3XKYuuZtrrlomWgX4kHnTYE9fydiRdCrml9
//fd state: (75) MINo6zG7ZS3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaAYTtNueQjzs9GvnM146GOEVteJffSDvKyMaSfioXjUPfo4iz7aakQ4IenrGLNZlPnIodCODHE4bLw7m3qDbrtGk9aZlLBM2SE0MHFG1RKPMI0KSJyFlJ0HFdT7tWSS_dvYgD4qf

ret = client.Write(fd650, []byte("nBfx5iIH8zh6IPKrv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) MINo6zG7ZS3i561JZMl7UGrg97WUbtAWGp2PoSK1FJaAYTtNueQjzs9GvnM146GOEVteJffSDvKnBfx5iIH8zh6IPKrvaakQ4IenrGLNZlPnIodCODHE4bLw7m3qDbrtGk9aZlLBM2SE0MHFG1RKPMI0KSJyFlJ0HFdT7tWSS_dvYgD4qf

fd661 := client.Open("/cPJ5D3QIxK.txt", client.O_RDWR|client.O_CREATE)
if(fd661 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd656, []byte("NzgFAuHI9e97ccS4Svrw9t6fJtuEp37CxoYVu6as"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) NzgFAuHI9e97ccS4Svrw9t6fJtuEp37CxoYVu6as
//fd state: (40) NzgFAuHI9e97ccS4Svrw9t6fJtuEp37CxoYVu6as

ret = client.Write(fd656, []byte("g7SOYdNmu0P5AjStB6f3lfOBIMdWuh0xVodTgyqcYYPuu_DTwVruqp7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) NzgFAuHI9e97ccS4Svrw9t6fJtuEp37CxoYVu6asg7SOYdNmu0P5AjStB6f3lfOBIMdWuh0xVodTgyqcYYPuu_DTwVruqp7

ret = client.Seek(fd632, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Seek(fd652, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


buf, ret = client.Read(fd654, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd650, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


fd662 := client.Open("/__oKbcjfmB.txt", client.O_RDWR|client.O_CREATE)
if(fd662 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd661, []byte("tOiImKhlAGdg3Eq3ungbgVa0IFTj2Wufgx64Oe74ZFnRycbz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) tOiImKhlAGdg3Eq3ungbgVa0IFTj2Wufgx64Oe74ZFnRycbz

ret = client.Close(fd647)
if(ret != 0) {
  panic("close failed")
}

//fd state: (13) 44jkIiTN6S_iy9Ilosat9l80Z_k76VheqslktiFBFV57y0F49QH9Mr_yVQCEyEpKJsYqaqkj_4psDdn2auoQz06Yu795i6rfYnIhaJLzc03phoHXh3n13CUcRJIhpvqX9hYJ6F_HXtC6Yf2wHrOjWAl

ret = client.Write(fd632, []byte("wji4xoat"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) 44jkIiTN6S_iywji4xoatl80Z_k76VheqslktiFBFV57y0F49QH9Mr_yVQCEyEpKJsYqaqkj_4psDdn2auoQz06Yu795i6rfYnIhaJLzc03phoHXh3n13CUcRJIhpvqX9hYJ6F_HXtC6Yf2wHrOjWAl

ret = client.Close(fd662)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd632)
if(ret != 0) {
  panic("close failed")
}


fd663 := client.Open("/kRBIeb8ihu.txt", client.O_RDWR|client.O_CREATE)
if(fd663 < 0) {
  panic("open failed")
}


ret = client.Seek(fd660, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Close(fd650)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd652, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


buf, ret = client.Read(fd625, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "KLCCYh") {
  panic("wrong data returned")
}


ret = client.Seek(fd648, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd651)
if(ret != 0) {
  panic("close failed")
}


fd664 := client.Open("/qgChWbDgrs.txt", client.O_RDWR|client.O_CREATE)
if(fd664 < 0) {
  panic("open failed")
}


ret = client.Close(fd649)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd654)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd655, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd663, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd665 := client.Open("/Cnq8lxUTLg.txt", client.O_RDWR|client.O_CREATE)
if(fd665 < 0) {
  panic("open failed")
}


fd666 := client.Open("/IKie96SW_E.txt", client.O_RDWR|client.O_CREATE)
if(fd666 < 0) {
  panic("open failed")
}


fd667 := client.Open("/qkb5m1FlgT.txt", client.O_RDWR|client.O_CREATE)
if(fd667 < 0) {
  panic("open failed")
}


fd668 := client.Open("/KyMhRXeJjY.txt", client.O_RDWR|client.O_CREATE)
if(fd668 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd639, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd669 := client.Open("/3djEq6AYXg.txt", client.O_RDWR|client.O_CREATE)
if(fd669 < 0) {
  panic("open failed")
}


ret = client.Close(fd645)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd668, []byte("QSp0JT5YkVpJ9o8ieHLG5zPkjoDlQvCnoKUUPmc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) QSp0JT5YkVpJ9o8ieHLG5zPkjoDlQvCnoKUUPmc

ret = client.Close(fd625)
if(ret != 0) {
  panic("close failed")
}


fd670 := client.Open("/VmU5r1Iala.txt", client.O_RDWR|client.O_CREATE)
if(fd670 < 0) {
  panic("open failed")
}


ret = client.Close(fd623)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd664)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd670, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd652, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd655, []byte("dAi9fTmKCqmpJ9jV142uF88HXUYQmlRbbvbubGM1pn4_DqQF58t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) dAi9fTmKCqmpJ9jV142uF88HXUYQmlRbbvbubGM1pn4_DqQF58t

fd671 := client.Open("/Xa2of_COvq.txt", client.O_RDWR|client.O_CREATE)
if(fd671 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd663, []byte("SfnASbknr6tCq63MfMoTFLDfswmpkJWZG6geFJuiB8s68YgGbBwF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) SfnASbknr6tCq63MfMoTFLDfswmpkJWZG6geFJuiB8s68YgGbBwF

fd672 := client.Open("/jKyoNeaMC8.txt", client.O_RDWR|client.O_CREATE)
if(fd672 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd658, []byte("8IQgHX8hVAvQgW3Qu3A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) 8IQgHX8hVAvQgW3Qu3A

ret = client.Close(fd655)
if(ret != 0) {
  panic("close failed")
}

//fd state: (39) QSp0JT5YkVpJ9o8ieHLG5zPkjoDlQvCnoKUUPmc

ret = client.Write(fd668, []byte("tZ7_HRMd0I2f_TqkZ2Mre6prJLbc7FNyVJGOsLz2HJiXQ5Iq9ZpWrTsYCC8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) QSp0JT5YkVpJ9o8ieHLG5zPkjoDlQvCnoKUUPmctZ7_HRMd0I2f_TqkZ2Mre6prJLbc7FNyVJGOsLz2HJiXQ5Iq9ZpWrTsYCC8

fd673 := client.Open("/eaI_co6DSS.txt", client.O_RDWR|client.O_CREATE)
if(fd673 < 0) {
  panic("open failed")
}


fd674 := client.Open("/epwKfrBhnk.txt", client.O_RDWR|client.O_CREATE)
if(fd674 < 0) {
  panic("open failed")
}


fd675 := client.Open("/gwNdAj_Oux.txt", client.O_RDWR|client.O_CREATE)
if(fd675 < 0) {
  panic("open failed")
}

//fd state: (0) xe3HRD4_dauG7Tz5irlDpRk49VGPlcyKoaxLzqJvxFUetsXcy6dV58i7QWT2YLqeDemq8usx2iFFZh1pbjtSka6lYrPTPTFMrBAMefb8PygNeziHOUDziqhOg7jee4TNssB4WU2ueuTUAcKdFUtSxhVd4nosbRuhrqm_l2LcpoQZOGVzhEZCzQSG5n0iKsH3drRpU

ret = client.Write(fd669, []byte("E33lO3Lfz5yWPQHy69"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) E33lO3Lfz5yWPQHy69lDpRk49VGPlcyKoaxLzqJvxFUetsXcy6dV58i7QWT2YLqeDemq8usx2iFFZh1pbjtSka6lYrPTPTFMrBAMefb8PygNeziHOUDziqhOg7jee4TNssB4WU2ueuTUAcKdFUtSxhVd4nosbRuhrqm_l2LcpoQZOGVzhEZCzQSG5n0iKsH3drRpU

ret = client.Close(fd663)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd666, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd665, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd661, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Seek(fd660, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd652)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd672, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd659, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd658)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd667, []byte("GmwEQtHPtRALrQSQcPEqAMIpRcQFip_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) GmwEQtHPtRALrQSQcPEqAMIpRcQFip_

buf, ret = client.Read(fd648, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "eRaZg9dy6ghopA_Anq7") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd666, []byte("aI6PxoC3yui0yPp3YSscLLVn1wOMKcOVKtC4WT1Mq55qaifUWYEaqw3NEQDQyFrhcPWymsHY1V0VSK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) aI6PxoC3yui0yPp3YSscLLVn1wOMKcOVKtC4WT1Mq55qaifUWYEaqw3NEQDQyFrhcPWymsHY1V0VSK
//fd state: (31) GmwEQtHPtRALrQSQcPEqAMIpRcQFip_

ret = client.Write(fd667, []byte("TTpemESX2hogFCNdxAUATo9zxd7auSgIpfIdgkx3FMw8IFt1JKggGkzOANORSCf_Jh47W_17h"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) GmwEQtHPtRALrQSQcPEqAMIpRcQFip_TTpemESX2hogFCNdxAUATo9zxd7auSgIpfIdgkx3FMw8IFt1JKggGkzOANORSCf_Jh47W_17h

ret = client.Seek(fd665, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd660, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YWrCu2q5CnN") {
  panic("wrong data returned")
}


ret = client.Seek(fd656, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd675, []byte("BHdkAmv_tyMrSIL45UySg1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) BHdkAmv_tyMrSIL45UySg1

ret = client.Seek(fd660, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Seek(fd665, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd660, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YWrC") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd668, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd646, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd676 := client.Open("/ydk9BRFkKJ.txt", client.O_RDWR|client.O_CREATE)
if(fd676 < 0) {
  panic("open failed")
}


fd677 := client.Open("/1g_SO9gjIH.txt", client.O_RDWR|client.O_CREATE)
if(fd677 < 0) {
  panic("open failed")
}


ret = client.Close(fd659)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd639, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}


ret = client.Close(fd672)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd671, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


ret = client.Seek(fd674, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd678 := client.Open("/mJJZDlZYe9.txt", client.O_RDWR|client.O_CREATE)
if(fd678 < 0) {
  panic("open failed")
}

//fd state: (30) xyHFttu_tPoeRaZg9dy6ghopA_Anq7ojTfcWEaG7MmXIad0fyYGrlib6kc8cRf9p4NOBk

ret = client.Write(fd648, []byte("lq8ceMerS1EBpiwNtkFkuuZ5mQmEjyyosYlo43iuwt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) xyHFttu_tPoeRaZg9dy6ghopA_Anq7lq8ceMerS1EBpiwNtkFkuuZ5mQmEjyyosYlo43iuwt

ret = client.Close(fd673)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd660, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


fd679 := client.Open("/Iw63byftWQ.txt", client.O_RDWR|client.O_CREATE)
if(fd679 < 0) {
  panic("open failed")
}


ret = client.Seek(fd669, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


buf, ret = client.Read(fd665, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd666)
if(ret != 0) {
  panic("close failed")
}


fd680 := client.Open("/iWExPMWmQy.txt", client.O_RDWR|client.O_CREATE)
if(fd680 < 0) {
  panic("open failed")
}


fd681 := client.Open("/mbhXIMTJYM.txt", client.O_RDWR|client.O_CREATE)
if(fd681 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd648, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd660)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd674, []byte("aJNo8FHH494_r6GuJmFIlzeoIA06Fhm1FmDmLKA3MSqGO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) aJNo8FHH494_r6GuJmFIlzeoIA06Fhm1FmDmLKA3MSqGO

buf, ret = client.Read(fd648, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (98) QSp0JT5YkVpJ9o8ieHLG5zPkjoDlQvCnoKUUPmctZ7_HRMd0I2f_TqkZ2Mre6prJLbc7FNyVJGOsLz2HJiXQ5Iq9ZpWrTsYCC8

ret = client.Write(fd668, []byte("EZvoI9Opp8rXPrWqfsVutm5I39n_kH6eHjPzSNlia"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (139) QSp0JT5YkVpJ9o8ieHLG5zPkjoDlQvCnoKUUPmctZ7_HRMd0I2f_TqkZ2Mre6prJLbc7FNyVJGOsLz2HJiXQ5Iq9ZpWrTsYCC8EZvoI9Opp8rXPrWqfsVutm5I39n_kH6eHjPzSNlia

fd682 := client.Open("/YbYhTNaGQb.txt", client.O_RDWR|client.O_CREATE)
if(fd682 < 0) {
  panic("open failed")
}


ret = client.Close(fd661)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd665, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd682)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd639)
if(ret != 0) {
  panic("close failed")
}


fd683 := client.Open("/Xa2of_COvq.txt", client.O_RDWR|client.O_CREATE)
if(fd683 < 0) {
  panic("open failed")
}


ret = client.Close(fd656)
if(ret != 0) {
  panic("close failed")
}


fd684 := client.Open("/tZiK0BxhAL.txt", client.O_RDWR|client.O_CREATE)
if(fd684 < 0) {
  panic("open failed")
}

//fd state: (139) QSp0JT5YkVpJ9o8ieHLG5zPkjoDlQvCnoKUUPmctZ7_HRMd0I2f_TqkZ2Mre6prJLbc7FNyVJGOsLz2HJiXQ5Iq9ZpWrTsYCC8EZvoI9Opp8rXPrWqfsVutm5I39n_kH6eHjPzSNlia

ret = client.Write(fd668, []byte("2vfTLln2twgf3WVZrjjOHnJdzOsHBc67R91OCjUW681jLioHr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (188) QSp0JT5YkVpJ9o8ieHLG5zPkjoDlQvCnoKUUPmctZ7_HRMd0I2f_TqkZ2Mre6prJLbc7FNyVJGOsLz2HJiXQ5Iq9ZpWrTsYCC8EZvoI9Opp8rXPrWqfsVutm5I39n_kH6eHjPzSNlia2vfTLln2twgf3WVZrjjOHnJdzOsHBc67R91OCjUW681jLioHr

ret = client.Close(fd670)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd681, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd676, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nnBaL1SCiqtTm4jsyYsRyRiQV6tG7yVJi5UoCIdr5f4B0e_VURZvaKvFWRUFK") {
  panic("wrong data returned")
}


ret = client.Seek(fd683, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


ret = client.Seek(fd669, 103, client.SEEK_SET)
if(ret != 103) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 103)
  panic("seek failed")
}


buf, ret = client.Read(fd683, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "M53WXLpemh1OQ6DON4iTTYdRyTJjhzN0Jg3qlf_KPwk") {
  panic("wrong data returned")
}


fd685 := client.Open("/AAMt9Vzr1x.txt", client.O_RDWR|client.O_CREATE)
if(fd685 < 0) {
  panic("open failed")
}


ret = client.Seek(fd665, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd679)
if(ret != 0) {
  panic("close failed")
}


fd686 := client.Open("/74mEwYxfha.txt", client.O_RDWR|client.O_CREATE)
if(fd686 < 0) {
  panic("open failed")
}


ret = client.Seek(fd676, 217, client.SEEK_SET)
if(ret != 217) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 217)
  panic("seek failed")
}


ret = client.Close(fd665)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd681)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd680, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "SGj0tyfq6IFkdxFZqEvnNLFXOOh1PXLXXUpvRxg_JaPdc7tkYx6Mr8x_Coi7V1QLY9sKmbka1_NhRFZ1NxeqONqm") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd685, []byte("sbYPJzK8QUbafkbJprVZsuPr7q4xXA80DgvWNaQYrP6KcMKOmRJKGDj0fnwLF7SBv_TkjU8XE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) sbYPJzK8QUbafkbJprVZsuPr7q4xXA80DgvWNaQYrP6KcMKOmRJKGDj0fnwLF7SBv_TkjU8XE

ret = client.Close(fd677)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd680, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_WYgiToNnJeJaseT8LLI_rUIdb1imB0aTqEQXoEtBiyt") {
  panic("wrong data returned")
}


fd687 := client.Open("/VXJbIbbGgX.txt", client.O_RDWR|client.O_CREATE)
if(fd687 < 0) {
  panic("open failed")
}


ret = client.Close(fd668)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd676, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Q") {
  panic("wrong data returned")
}


ret = client.Seek(fd676, 206, client.SEEK_SET)
if(ret != 206) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 206)
  panic("seek failed")
}


ret = client.Seek(fd678, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd646)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd687)
if(ret != 0) {
  panic("close failed")
}


fd688 := client.Open("/H8H4MRWssQ.txt", client.O_RDWR|client.O_CREATE)
if(fd688 < 0) {
  panic("open failed")
}


ret = client.Seek(fd686, 99, client.SEEK_SET)
if(ret != 99) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 99)
  panic("seek failed")
}


ret = client.Close(fd678)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd675)
if(ret != 0) {
  panic("close failed")
}


fd689 := client.Open("/U1r0MhQgKJ.txt", client.O_RDWR|client.O_CREATE)
if(fd689 < 0) {
  panic("open failed")
}


ret = client.Close(fd680)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd686)
if(ret != 0) {
  panic("close failed")
}


fd690 := client.Open("/DGye9zYlsd.txt", client.O_RDWR|client.O_CREATE)
if(fd690 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd676, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "31EcywulwHhQdzT3CkPlt5lPngOEN9Anp6YGwx") {
  panic("wrong data returned")
}


fd691 := client.Open("/RDgMjMG9OR.txt", client.O_RDWR|client.O_CREATE)
if(fd691 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd683, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd669, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


ret = client.Close(fd689)
if(ret != 0) {
  panic("close failed")
}

//fd state: (73) sbYPJzK8QUbafkbJprVZsuPr7q4xXA80DgvWNaQYrP6KcMKOmRJKGDj0fnwLF7SBv_TkjU8XE

ret = client.Write(fd685, []byte("3Rzf4i_PW6fAEcJXp9w"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) sbYPJzK8QUbafkbJprVZsuPr7q4xXA80DgvWNaQYrP6KcMKOmRJKGDj0fnwLF7SBv_TkjU8XE3Rzf4i_PW6fAEcJXp9w

fd692 := client.Open("/knLVMcM41S.txt", client.O_RDWR|client.O_CREATE)
if(fd692 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd671, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "tx4") {
  panic("wrong data returned")
}


ret = client.Close(fd674)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd691, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd693 := client.Open("/8Tqb7q4mha.txt", client.O_RDWR|client.O_CREATE)
if(fd693 < 0) {
  panic("open failed")
}


fd694 := client.Open("/Xd5e7z7Cv4.txt", client.O_RDWR|client.O_CREATE)
if(fd694 < 0) {
  panic("open failed")
}


ret = client.Seek(fd693, 68, client.SEEK_SET)
if(ret != 68) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 68)
  panic("seek failed")
}


ret = client.Close(fd690)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd693)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd676, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd692, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd691, []byte("8vx1W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) 8vx1W

buf, ret = client.Read(fd684, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd683, 108, client.SEEK_SET)
if(ret != 108) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 108)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd683, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


fd695 := client.Open("/R1VJ6U8gR7.txt", client.O_RDWR|client.O_CREATE)
if(fd695 < 0) {
  panic("open failed")
}


ret = client.Close(fd676)
if(ret != 0) {
  panic("close failed")
}


fd696 := client.Open("/WGTJdUdKLz.txt", client.O_RDWR|client.O_CREATE)
if(fd696 < 0) {
  panic("open failed")
}

//fd state: (104) GmwEQtHPtRALrQSQcPEqAMIpRcQFip_TTpemESX2hogFCNdxAUATo9zxd7auSgIpfIdgkx3FMw8IFt1JKggGkzOANORSCf_Jh47W_17h

ret = client.Write(fd667, []byte("kig_W_OW14jRlCQ1C_rvhpA298"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) GmwEQtHPtRALrQSQcPEqAMIpRcQFip_TTpemESX2hogFCNdxAUATo9zxd7auSgIpfIdgkx3FMw8IFt1JKggGkzOANORSCf_Jh47W_17hkig_W_OW14jRlCQ1C_rvhpA298

ret = client.Seek(fd669, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}


ret = client.Seek(fd671, 76, client.SEEK_SET)
if(ret != 76) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 76)
  panic("seek failed")
}

//fd state: (0) iCyQAa0N1TiYXbdRFTi_en3qcKdb2A5JdfFQBm78vxguj_dMR8nYxldTXnssjXhCfj6SR9H7ofr3LYm3k27qkwd3Ppxv

ret = client.Write(fd695, []byte("BuNEBWgVJYhQTONLKuQzia50E7cMh2hn9xl6nnGrI8L6QprNkpwWm6PQ8g9ZneWQO_tU4f4DLhCTwyeuM8X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) BuNEBWgVJYhQTONLKuQzia50E7cMh2hn9xl6nnGrI8L6QprNkpwWm6PQ8g9ZneWQO_tU4f4DLhCTwyeuM8Xqkwd3Ppxv
//fd state: (0) 

ret = client.Write(fd692, []byte("Gj8dzMbIwB5LHgYjP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) Gj8dzMbIwB5LHgYjP

ret = client.Close(fd696)
if(ret != 0) {
  panic("close failed")
}

//fd state: (76) j9AakDYS1IM6zKKObFbnS13jjqXT8EFMGiweHtx4kaAFy552hwthOF6wz82p3KWDv1HM53WXLpemh1OQ6DON4iTTYdRyTJjhzN0Jg3qlf_KPwk

ret = client.Write(fd671, []byte("vUY2hVTEtnLaR6uHh_bheoDG4ccTMkdqU8jluPj16buzMvvuwLOgpCnn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) j9AakDYS1IM6zKKObFbnS13jjqXT8EFMGiweHtx4kaAFy552hwthOF6wz82p3KWDv1HM53WXLpemvUY2hVTEtnLaR6uHh_bheoDG4ccTMkdqU8jluPj16buzMvvuwLOgpCnn

fd697 := client.Open("/dQ5BvVzU2a.txt", client.O_RDWR|client.O_CREATE)
if(fd697 < 0) {
  panic("open failed")
}

//fd state: (57) E33lO3Lfz5yWPQHy69lDpRk49VGPlcyKoaxLzqJvxFUetsXcy6dV58i7QWT2YLqeDemq8usx2iFFZh1pbjtSka6lYrPTPTFMrBAMefb8PygNeziHOUDziqhOg7jee4TNssB4WU2ueuTUAcKdFUtSxhVd4nosbRuhrqm_l2LcpoQZOGVzhEZCzQSG5n0iKsH3drRpU

ret = client.Write(fd669, []byte("W1SIiv6ZPWTiME8ejHSSOkazs_fin3oaB1iCeAhlMFyEnM9qHPHCVPFtWWONfQoQkFu8N"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (126) E33lO3Lfz5yWPQHy69lDpRk49VGPlcyKoaxLzqJvxFUetsXcy6dV58i7QW1SIiv6ZPWTiME8ejHSSOkazs_fin3oaB1iCeAhlMFyEnM9qHPHCVPFtWWONfQoQkFu8NTNssB4WU2ueuTUAcKdFUtSxhVd4nosbRuhrqm_l2LcpoQZOGVzhEZCzQSG5n0iKsH3drRpU

buf, ret = client.Read(fd648, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd692)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd694, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (83) BuNEBWgVJYhQTONLKuQzia50E7cMh2hn9xl6nnGrI8L6QprNkpwWm6PQ8g9ZneWQO_tU4f4DLhCTwyeuM8Xqkwd3Ppxv

ret = client.Write(fd695, []byte("zCbddbBqktrnHSKlioRvFQKPI4ejJhdEp1j8V2uBnYpN30IaqlK1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) BuNEBWgVJYhQTONLKuQzia50E7cMh2hn9xl6nnGrI8L6QprNkpwWm6PQ8g9ZneWQO_tU4f4DLhCTwyeuM8XzCbddbBqktrnHSKlioRvFQKPI4ejJhdEp1j8V2uBnYpN30IaqlK1

fd698 := client.Open("/kFxvy7JBKC.txt", client.O_RDWR|client.O_CREATE)
if(fd698 < 0) {
  panic("open failed")
}


ret = client.Close(fd685)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd688, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd691, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd688, []byte("KQLvUo8hzWpb4h3iaM5zwenL0Di"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) KQLvUo8hzWpb4h3iaM5zwenL0Di

ret = client.Close(fd691)
if(ret != 0) {
  panic("close failed")
}

//fd state: (132) j9AakDYS1IM6zKKObFbnS13jjqXT8EFMGiweHtx4kaAFy552hwthOF6wz82p3KWDv1HM53WXLpemvUY2hVTEtnLaR6uHh_bheoDG4ccTMkdqU8jluPj16buzMvvuwLOgpCnn

ret = client.Write(fd671, []byte("AyJpkpvbQCykvIlRs3bYZvk88AUU4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (161) j9AakDYS1IM6zKKObFbnS13jjqXT8EFMGiweHtx4kaAFy552hwthOF6wz82p3KWDv1HM53WXLpemvUY2hVTEtnLaR6uHh_bheoDG4ccTMkdqU8jluPj16buzMvvuwLOgpCnnAyJpkpvbQCykvIlRs3bYZvk88AUU4

buf, ret = client.Read(fd698, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (161) j9AakDYS1IM6zKKObFbnS13jjqXT8EFMGiweHtx4kaAFy552hwthOF6wz82p3KWDv1HM53WXLpemvUY2hVTEtnLaR6uHh_bheoDG4ccTMkdqU8jluPj16buzMvvuwLOgpCnnAyJpkpvbQCykvIlRs3bYZvk88AUU4

ret = client.Write(fd671, []byte("fxfyP4vQo0LkJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (174) j9AakDYS1IM6zKKObFbnS13jjqXT8EFMGiweHtx4kaAFy552hwthOF6wz82p3KWDv1HM53WXLpemvUY2hVTEtnLaR6uHh_bheoDG4ccTMkdqU8jluPj16buzMvvuwLOgpCnnAyJpkpvbQCykvIlRs3bYZvk88AUU4fxfyP4vQo0LkJ
//fd state: (130) GmwEQtHPtRALrQSQcPEqAMIpRcQFip_TTpemESX2hogFCNdxAUATo9zxd7auSgIpfIdgkx3FMw8IFt1JKggGkzOANORSCf_Jh47W_17hkig_W_OW14jRlCQ1C_rvhpA298

ret = client.Write(fd667, []byte("mO9dWEror9D6AH_8TS4r46"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) GmwEQtHPtRALrQSQcPEqAMIpRcQFip_TTpemESX2hogFCNdxAUATo9zxd7auSgIpfIdgkx3FMw8IFt1JKggGkzOANORSCf_Jh47W_17hkig_W_OW14jRlCQ1C_rvhpA298mO9dWEror9D6AH_8TS4r46

fd699 := client.Open("/Ri3JhqskKq.txt", client.O_RDWR|client.O_CREATE)
if(fd699 < 0) {
  panic("open failed")
}


ret = client.Close(fd684)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd697, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AxxznkIjtif4oETVQKkWYABPqUmvfGNPZz12MSix") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd694, []byte("z8G214a4uXqj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) z8G214a4uXqj
//fd state: (47) j9AakDYS1IM6zKKObFbnS13jjqXT8EFMGiweHtx4kaAFy552hwthOF6wz82p3KWDv1HM53WXLpemvUY2hVTEtnLaR6uHh_bheoDG4ccTMkdqU8jluPj16buzMvvuwLOgpCnnAyJpkpvbQCykvIlRs3bYZvk88AUU4fxfyP4vQo0LkJ

ret = client.Write(fd683, []byte("qfhoJaxRNTaedfeS0uwCSqM14ry6P6lj12Nrb76AnhgegmYfG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) j9AakDYS1IM6zKKObFbnS13jjqXT8EFMGiweHtx4kaAFy55qfhoJaxRNTaedfeS0uwCSqM14ry6P6lj12Nrb76AnhgegmYfGeoDG4ccTMkdqU8jluPj16buzMvvuwLOgpCnnAyJpkpvbQCykvIlRs3bYZvk88AUU4fxfyP4vQo0LkJ

buf, ret = client.Read(fd688, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd700 := client.Open("/HN_h2_khZ5.txt", client.O_RDWR|client.O_CREATE)
if(fd700 < 0) {
  panic("open failed")
}


ret = client.Close(fd669)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd671, 87, client.SEEK_SET)
if(ret != 87) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 87)
  panic("seek failed")
}


ret = client.Close(fd648)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd671)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd695, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd694, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


fd701 := client.Open("/wVLioSOIGX.txt", client.O_RDWR|client.O_CREATE)
if(fd701 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd700, []byte("mQlOk69D"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) mQlOk69D

ret = client.Seek(fd667, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


ret = client.Close(fd698)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd700)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd701, []byte("MnK7Gr6dgNvZNuI2BmJHw91L4Ywc7hRi_i0VRlJ2dae13VXg1V8mQ8GrnlQqg4Ab_Dn2Aa4qQZ7CiBdXsxWYk8GrSDh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) MnK7Gr6dgNvZNuI2BmJHw91L4Ywc7hRi_i0VRlJ2dae13VXg1V8mQ8GrnlQqg4Ab_Dn2Aa4qQZ7CiBdXsxWYk8GrSDh

ret = client.Seek(fd688, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Seek(fd694, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd699, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (91) MnK7Gr6dgNvZNuI2BmJHw91L4Ywc7hRi_i0VRlJ2dae13VXg1V8mQ8GrnlQqg4Ab_Dn2Aa4qQZ7CiBdXsxWYk8GrSDh

ret = client.Write(fd701, []byte("fWWyTmmjDqjwc3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) MnK7Gr6dgNvZNuI2BmJHw91L4Ywc7hRi_i0VRlJ2dae13VXg1V8mQ8GrnlQqg4Ab_Dn2Aa4qQZ7CiBdXsxWYk8GrSDhfWWyTmmjDqjwc3

fd702 := client.Open("/HG_fqUj4cf.txt", client.O_RDWR|client.O_CREATE)
if(fd702 < 0) {
  panic("open failed")
}


fd703 := client.Open("/b9IOiiNWX_.txt", client.O_RDWR|client.O_CREATE)
if(fd703 < 0) {
  panic("open failed")
}


fd704 := client.Open("/2SLyLKR9En.txt", client.O_RDWR|client.O_CREATE)
if(fd704 < 0) {
  panic("open failed")
}


fd705 := client.Open("/nDeXh0ab79.txt", client.O_RDWR|client.O_CREATE)
if(fd705 < 0) {
  panic("open failed")
}


ret = client.Close(fd703)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd697, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd699, []byte("S3g5IaXiKi2OvIi9c4rqbAfy07VtjWGkXy8q5dYuwE4d4fYj5ScgAlr1ScpjnPSKRR3DKfPDMqflQZJeYbE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) S3g5IaXiKi2OvIi9c4rqbAfy07VtjWGkXy8q5dYuwE4d4fYj5ScgAlr1ScpjnPSKRR3DKfPDMqflQZJeYbE

fd706 := client.Open("/CfKi92RL9_.txt", client.O_RDWR|client.O_CREATE)
if(fd706 < 0) {
  panic("open failed")
}

//fd state: (0) Eyp9BWYKPl0D9bGPcgEEUDlvbkoF3j19lmL90s1rK5PSs5SxF4RtR44Nf

ret = client.Write(fd702, []byte("spVew7rXMeh5Hi4VfXxEJ1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) spVew7rXMeh5Hi4VfXxEJ1lvbkoF3j19lmL90s1rK5PSs5SxF4RtR44Nf

ret = client.Seek(fd699, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Seek(fd702, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


ret = client.Close(fd701)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd694, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8G214a4uXqj") {
  panic("wrong data returned")
}


fd707 := client.Open("/JRYsh5HY6W.txt", client.O_RDWR|client.O_CREATE)
if(fd707 < 0) {
  panic("open failed")
}


fd708 := client.Open("/jmVadlwBeg.txt", client.O_RDWR|client.O_CREATE)
if(fd708 < 0) {
  panic("open failed")
}

//fd state: (80) AxxznkIjtif4oETVQKkWYABPqUmvfGNPZz12MSixvq7nYXCHYanH4KEsqhpJLya4TUEXq4O9cf7FfXrtB1WBozMhPGWd77z6j6

ret = client.Write(fd697, []byte("FtXdD0J6sR3zWStpayQplQ2ypiMfIq9L9xc9vGTKWBvkow51rnHbN_JeNPxb2alZPUzuBOJv_gwsid"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (158) AxxznkIjtif4oETVQKkWYABPqUmvfGNPZz12MSixvq7nYXCHYanH4KEsqhpJLya4TUEXq4O9cf7FfXrtFtXdD0J6sR3zWStpayQplQ2ypiMfIq9L9xc9vGTKWBvkow51rnHbN_JeNPxb2alZPUzuBOJv_gwsid

ret = client.Seek(fd667, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}

//fd state: (19) spVew7rXMeh5Hi4VfXxEJ1lvbkoF3j19lmL90s1rK5PSs5SxF4RtR44Nf

ret = client.Write(fd702, []byte("heFH9xSBIxyZ9hYd6T6IAU0eEpQ33to79xtKhqAUvC5Jx6haDNBKTAwVokGFjc68UBWtifm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) spVew7rXMeh5Hi4VfXxheFH9xSBIxyZ9hYd6T6IAU0eEpQ33to79xtKhqAUvC5Jx6haDNBKTAwVokGFjc68UBWtifm

fd709 := client.Open("/YBc2F9uh1_.txt", client.O_RDWR|client.O_CREATE)
if(fd709 < 0) {
  panic("open failed")
}


ret = client.Close(fd699)
if(ret != 0) {
  panic("close failed")
}


fd710 := client.Open("/eaI_co6DSS.txt", client.O_RDWR|client.O_CREATE)
if(fd710 < 0) {
  panic("open failed")
}


fd711 := client.Open("/rmUv9JJeUn.txt", client.O_RDWR|client.O_CREATE)
if(fd711 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd707, []byte("mYs2m889NDnq2oo_PJHJATypzrtIQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) mYs2m889NDnq2oo_PJHJATypzrtIQ
//fd state: (0) 

ret = client.Write(fd709, []byte("deBKtsqqWNRdB9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) deBKtsqqWNRdB9

ret = client.Seek(fd695, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (12) z8G214a4uXqj

ret = client.Write(fd694, []byte("JQzZXyJtJpQMGez4OH4SIIKXmQDc3r4zCDpVX4I0pPol37QfHa1jZjVXudAk_EprBZmfdZAvXu9XutoG7QfCKWvXs4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) z8G214a4uXqjJQzZXyJtJpQMGez4OH4SIIKXmQDc3r4zCDpVX4I0pPol37QfHa1jZjVXudAk_EprBZmfdZAvXu9XutoG7QfCKWvXs4
//fd state: (158) AxxznkIjtif4oETVQKkWYABPqUmvfGNPZz12MSixvq7nYXCHYanH4KEsqhpJLya4TUEXq4O9cf7FfXrtFtXdD0J6sR3zWStpayQplQ2ypiMfIq9L9xc9vGTKWBvkow51rnHbN_JeNPxb2alZPUzuBOJv_gwsid

ret = client.Write(fd697, []byte("7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (159) AxxznkIjtif4oETVQKkWYABPqUmvfGNPZz12MSixvq7nYXCHYanH4KEsqhpJLya4TUEXq4O9cf7FfXrtFtXdD0J6sR3zWStpayQplQ2ypiMfIq9L9xc9vGTKWBvkow51rnHbN_JeNPxb2alZPUzuBOJv_gwsid7

buf, ret = client.Read(fd694, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd705, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd683)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd688, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


buf, ret = client.Read(fd695, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NEBWgVJYhQTONLKuQzia50E7cMh2hn9xl6nnGrI8L6QprNkpwWm6PQ8") {
  panic("wrong data returned")
}


fd712 := client.Open("/wvtPwOJxDc.txt", client.O_RDWR|client.O_CREATE)
if(fd712 < 0) {
  panic("open failed")
}


fd713 := client.Open("/jmVadlwBeg.txt", client.O_RDWR|client.O_CREATE)
if(fd713 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd713, []byte("1GmwXTYHumQZMmD9nw1tMQRlz03V8XtNisFsYZQAOhH4jy02VynVM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) 1GmwXTYHumQZMmD9nw1tMQRlz03V8XtNisFsYZQAOhH4jy02VynVM
//fd state: (53) 1GmwXTYHumQZMmD9nw1tMQRlz03V8XtNisFsYZQAOhH4jy02VynVM

ret = client.Write(fd713, []byte("UMtW9iZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) 1GmwXTYHumQZMmD9nw1tMQRlz03V8XtNisFsYZQAOhH4jy02VynVMUMtW9iZ

buf, ret = client.Read(fd708, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1GmwXTYHumQZMmD9nw1tMQRlz0") {
  panic("wrong data returned")
}


ret = client.Seek(fd712, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd667)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd695)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd712, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd694, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd714 := client.Open("/0N4kwgmdsj.txt", client.O_RDWR|client.O_CREATE)
if(fd714 < 0) {
  panic("open failed")
}


ret = client.Close(fd712)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd711)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd708, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3V8XtNisFsYZQAOhH4jy02VynVMUMtW9iZ") {
  panic("wrong data returned")
}


fd715 := client.Open("/QVh9ftowiN.txt", client.O_RDWR|client.O_CREATE)
if(fd715 < 0) {
  panic("open failed")
}


fd716 := client.Open("/HPtMBtpc99.txt", client.O_RDWR|client.O_CREATE)
if(fd716 < 0) {
  panic("open failed")
}


ret = client.Seek(fd710, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd709, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (0) 

ret = client.Write(fd715, []byte("TLFv_fzdb2zfleXDGZ31NdlbnWcBI9vxHwxlm_C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) TLFv_fzdb2zfleXDGZ31NdlbnWcBI9vxHwxlm_C
//fd state: (0) 

ret = client.Write(fd704, []byte("BAbeARDfR2H1GvO0ub1CoNS9krcv16H1tWIxyDSJ4KDQ05G1AE2zRVSCE5QBFceheQh5JqBgJo3RedLgtB1usIYhqen1ApkR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) BAbeARDfR2H1GvO0ub1CoNS9krcv16H1tWIxyDSJ4KDQ05G1AE2zRVSCE5QBFceheQh5JqBgJo3RedLgtB1usIYhqen1ApkR

ret = client.Close(fd710)
if(ret != 0) {
  panic("close failed")
}

//fd state: (29) mYs2m889NDnq2oo_PJHJATypzrtIQ

ret = client.Write(fd707, []byte("EKSeNu7RX9F6QV7XlYDQBSwGQaH6IaVteiBf0NlPaVmjv8iqkMTW0hVg04CV7VNw2DXPts391DQA6ZgND7VGu4hDP0y5Uj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) mYs2m889NDnq2oo_PJHJATypzrtIQEKSeNu7RX9F6QV7XlYDQBSwGQaH6IaVteiBf0NlPaVmjv8iqkMTW0hVg04CV7VNw2DXPts391DQA6ZgND7VGu4hDP0y5Uj
//fd state: (159) AxxznkIjtif4oETVQKkWYABPqUmvfGNPZz12MSixvq7nYXCHYanH4KEsqhpJLya4TUEXq4O9cf7FfXrtFtXdD0J6sR3zWStpayQplQ2ypiMfIq9L9xc9vGTKWBvkow51rnHbN_JeNPxb2alZPUzuBOJv_gwsid7

ret = client.Write(fd697, []byte("nAUm416FFrSsPGqy77MQOdaxKMdMLZocgnZeRk6pvv9y2Z8Q8EEJzsQRFJn3c3k0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (223) AxxznkIjtif4oETVQKkWYABPqUmvfGNPZz12MSixvq7nYXCHYanH4KEsqhpJLya4TUEXq4O9cf7FfXrtFtXdD0J6sR3zWStpayQplQ2ypiMfIq9L9xc9vGTKWBvkow51rnHbN_JeNPxb2alZPUzuBOJv_gwsid7nAUm416FFrSsPGqy77MQOdaxKMdMLZocgnZeRk6pvv9y2Z8Q8EEJzsQRFJn3c3k0
//fd state: (0) 

ret = client.Write(fd716, []byte("TfRXbWX4kpC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) TfRXbWX4kpC

fd717 := client.Open("/RPoih7SFkr.txt", client.O_RDWR|client.O_CREATE)
if(fd717 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd706, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd706)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd705, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (123) mYs2m889NDnq2oo_PJHJATypzrtIQEKSeNu7RX9F6QV7XlYDQBSwGQaH6IaVteiBf0NlPaVmjv8iqkMTW0hVg04CV7VNw2DXPts391DQA6ZgND7VGu4hDP0y5Uj

ret = client.Write(fd707, []byte("eRPDuA1DE1a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) mYs2m889NDnq2oo_PJHJATypzrtIQEKSeNu7RX9F6QV7XlYDQBSwGQaH6IaVteiBf0NlPaVmjv8iqkMTW0hVg04CV7VNw2DXPts391DQA6ZgND7VGu4hDP0y5UjeRPDuA1DE1a

fd718 := client.Open("/jpgEv1Jz8d.txt", client.O_RDWR|client.O_CREATE)
if(fd718 < 0) {
  panic("open failed")
}


fd719 := client.Open("/B4EwowUnon.txt", client.O_RDWR|client.O_CREATE)
if(fd719 < 0) {
  panic("open failed")
}


ret = client.Seek(fd688, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Close(fd715)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd713)
if(ret != 0) {
  panic("close failed")
}


fd720 := client.Open("/6qD6aooX0t.txt", client.O_RDWR|client.O_CREATE)
if(fd720 < 0) {
  panic("open failed")
}


ret = client.Seek(fd702, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}


ret = client.Seek(fd705, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd702)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd705, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd714)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd720, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd717)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd697)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd709, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "sqqWNRdB9") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd705, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd709, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


buf, ret = client.Read(fd716, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd719, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd721 := client.Open("/295LYlO2fY.txt", client.O_RDWR|client.O_CREATE)
if(fd721 < 0) {
  panic("open failed")
}


ret = client.Seek(fd719, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd688, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd722 := client.Open("/tL_sr4v2I3.txt", client.O_RDWR|client.O_CREATE)
if(fd722 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd694, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (96) BAbeARDfR2H1GvO0ub1CoNS9krcv16H1tWIxyDSJ4KDQ05G1AE2zRVSCE5QBFceheQh5JqBgJo3RedLgtB1usIYhqen1ApkR

ret = client.Write(fd704, []byte("IyfuBXKjcRtfUPJgom4XGbPvxJN1DYOH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (128) BAbeARDfR2H1GvO0ub1CoNS9krcv16H1tWIxyDSJ4KDQ05G1AE2zRVSCE5QBFceheQh5JqBgJo3RedLgtB1usIYhqen1ApkRIyfuBXKjcRtfUPJgom4XGbPvxJN1DYOH

fd723 := client.Open("/U1r0MhQgKJ.txt", client.O_RDWR|client.O_CREATE)
if(fd723 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd721, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (3) KQLvUo8hzWpb4h3iaM5zwenL0Di

ret = client.Write(fd688, []byte("QI87vjPceeTO7nOA6R86_MfhvKLPPEiztZwr0CZW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) KQLQI87vjPceeTO7nOA6R86_MfhvKLPPEiztZwr0CZW

fd724 := client.Open("/V56SxLcZii.txt", client.O_RDWR|client.O_CREATE)
if(fd724 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd719, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (8) deBKtsqqWNRdB9

ret = client.Write(fd709, []byte("jnIcLEI5HFqbdxN5EiBbXhagxQP9s3gH1TbpZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) deBKtsqqjnIcLEI5HFqbdxN5EiBbXhagxQP9s3gH1TbpZ
//fd state: (134) mYs2m889NDnq2oo_PJHJATypzrtIQEKSeNu7RX9F6QV7XlYDQBSwGQaH6IaVteiBf0NlPaVmjv8iqkMTW0hVg04CV7VNw2DXPts391DQA6ZgND7VGu4hDP0y5UjeRPDuA1DE1a

ret = client.Write(fd707, []byte("bx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) mYs2m889NDnq2oo_PJHJATypzrtIQEKSeNu7RX9F6QV7XlYDQBSwGQaH6IaVteiBf0NlPaVmjv8iqkMTW0hVg04CV7VNw2DXPts391DQA6ZgND7VGu4hDP0y5UjeRPDuA1DE1abx

fd725 := client.Open("/JRYsh5HY6W.txt", client.O_RDWR|client.O_CREATE)
if(fd725 < 0) {
  panic("open failed")
}


ret = client.Seek(fd725, 91, client.SEEK_SET)
if(ret != 91) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 91)
  panic("seek failed")
}


fd726 := client.Open("/ycFkAjvXbq.txt", client.O_RDWR|client.O_CREATE)
if(fd726 < 0) {
  panic("open failed")
}


ret = client.Close(fd707)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd722)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd704)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd720)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd725, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Nw2DXPts391DQA6ZgND7VGu4hD") {
  panic("wrong data returned")
}


ret = client.Seek(fd726, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd716)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd694, 76, client.SEEK_SET)
if(ret != 76) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 76)
  panic("seek failed")
}


ret = client.Close(fd726)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd708, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (60) 1GmwXTYHumQZMmD9nw1tMQRlz03V8XtNisFsYZQAOhH4jy02VynVMUMtW9iZ

ret = client.Write(fd708, []byte("cwYepUvcEv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) 1GmwXTYHumQZMmD9nw1tMQRlz03V8XtNisFsYZQAOhH4jy02VynVMUMtW9iZcwYepUvcEv

buf, ret = client.Read(fd721, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd724, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd705, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (117) mYs2m889NDnq2oo_PJHJATypzrtIQEKSeNu7RX9F6QV7XlYDQBSwGQaH6IaVteiBf0NlPaVmjv8iqkMTW0hVg04CV7VNw2DXPts391DQA6ZgND7VGu4hDP0y5UjeRPDuA1DE1abx

ret = client.Write(fd725, []byte("AElVtJlIugaeQva56LyFosw5voP_G"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (146) mYs2m889NDnq2oo_PJHJATypzrtIQEKSeNu7RX9F6QV7XlYDQBSwGQaH6IaVteiBf0NlPaVmjv8iqkMTW0hVg04CV7VNw2DXPts391DQA6ZgND7VGu4hDAElVtJlIugaeQva56LyFosw5voP_G

buf, ret = client.Read(fd725, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd718, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd723, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd718)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd721, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd724, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd694, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}

//fd state: (70) 1GmwXTYHumQZMmD9nw1tMQRlz03V8XtNisFsYZQAOhH4jy02VynVMUMtW9iZcwYepUvcEv

ret = client.Write(fd708, []byte("WCNt9e3t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) 1GmwXTYHumQZMmD9nw1tMQRlz03V8XtNisFsYZQAOhH4jy02VynVMUMtW9iZcwYepUvcEvWCNt9e3t
//fd state: (0) 

ret = client.Write(fd705, []byte("8SOI8g1vTErtnsY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) 8SOI8g1vTErtnsY
//fd state: (45) deBKtsqqjnIcLEI5HFqbdxN5EiBbXhagxQP9s3gH1TbpZ

ret = client.Write(fd709, []byte("JKt5FQmgVDkgQWFxoNcDW__Ev_JPR7RX3Y1X_v1KDQxC8R"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) deBKtsqqjnIcLEI5HFqbdxN5EiBbXhagxQP9s3gH1TbpZJKt5FQmgVDkgQWFxoNcDW__Ev_JPR7RX3Y1X_v1KDQxC8R

ret = client.Close(fd719)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd708, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd709, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (91) deBKtsqqjnIcLEI5HFqbdxN5EiBbXhagxQP9s3gH1TbpZJKt5FQmgVDkgQWFxoNcDW__Ev_JPR7RX3Y1X_v1KDQxC8R

ret = client.Write(fd709, []byte("ExTYFlQQhIceqacBgaN1JCJYrQULqRm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) deBKtsqqjnIcLEI5HFqbdxN5EiBbXhagxQP9s3gH1TbpZJKt5FQmgVDkgQWFxoNcDW__Ev_JPR7RX3Y1X_v1KDQxC8RExTYFlQQhIceqacBgaN1JCJYrQULqRm
//fd state: (15) 8SOI8g1vTErtnsY

ret = client.Write(fd705, []byte("j4BFWkWINkNGGZkSDLjKFIPLcDzUVDYJivQNWn4i2h8p78MUhc0G1Fk6FYaqEQ6EPAn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) 8SOI8g1vTErtnsYj4BFWkWINkNGGZkSDLjKFIPLcDzUVDYJivQNWn4i2h8p78MUhc0G1Fk6FYaqEQ6EPAn

fd727 := client.Open("/B7_Q2wkRCf.txt", client.O_RDWR|client.O_CREATE)
if(fd727 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd724, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd728 := client.Open("/bb3zIBjBgh.txt", client.O_RDWR|client.O_CREATE)
if(fd728 < 0) {
  panic("open failed")
}


ret = client.Close(fd728)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd688)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd708, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1GmwXTYHumQZMmD9nw1tMQRlz03") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd694, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "a4uXqjJQzZXyJtJpQMGez4OH4SIIKXmQDc3r4zCDpVX4I0pPol37QfHa1jZjVXudAk_EprBZmfdZAvXu9XutoG7Qf") {
  panic("wrong data returned")
}


fd729 := client.Open("/e_y2AYyBzC.txt", client.O_RDWR|client.O_CREATE)
if(fd729 < 0) {
  panic("open failed")
}


ret = client.Seek(fd729, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd723, []byte("dNGrZCCcgyXVAeLpBMIsMU66wVwpfxwWhTdOP9cuVCzy7J"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) dNGrZCCcgyXVAeLpBMIsMU66wVwpfxwWhTdOP9cuVCzy7J

ret = client.Close(fd729)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd725, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}

//fd state: (46) dNGrZCCcgyXVAeLpBMIsMU66wVwpfxwWhTdOP9cuVCzy7J

ret = client.Write(fd723, []byte("86QYCDFNX1PgFk_n44iTJpDsHCv7I6EESZ9QmcRkXtP9wUKmJ0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) dNGrZCCcgyXVAeLpBMIsMU66wVwpfxwWhTdOP9cuVCzy7J86QYCDFNX1PgFk_n44iTJpDsHCv7I6EESZ9QmcRkXtP9wUKmJ0

ret = client.Seek(fd727, 123, client.SEEK_SET)
if(ret != 123) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 123)
  panic("seek failed")
}


buf, ret = client.Read(fd708, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "V8XtNisFsYZQAOhH4jy02VynVMUMtW9iZcwYepUvcEvWCNt9e3t") {
  panic("wrong data returned")
}


ret = client.Close(fd724)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd708)
if(ret != 0) {
  panic("close failed")
}

//fd state: (48) mYs2m889NDnq2oo_PJHJATypzrtIQEKSeNu7RX9F6QV7XlYDQBSwGQaH6IaVteiBf0NlPaVmjv8iqkMTW0hVg04CV7VNw2DXPts391DQA6ZgND7VGu4hDAElVtJlIugaeQva56LyFosw5voP_G

ret = client.Write(fd725, []byte("sdajpgbaBktgVHijpJFmbk3V9zhvIpFL9dqH2z8IDVJwZj1mHJs7QmW_9M7LhS5Ox5JDhXGJRYQe0NDB8Jr94"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) mYs2m889NDnq2oo_PJHJATypzrtIQEKSeNu7RX9F6QV7XlYDsdajpgbaBktgVHijpJFmbk3V9zhvIpFL9dqH2z8IDVJwZj1mHJs7QmW_9M7LhS5Ox5JDhXGJRYQe0NDB8Jr946LyFosw5voP_G

ret = client.Seek(fd727, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Seek(fd725, 94, client.SEEK_SET)
if(ret != 94) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 94)
  panic("seek failed")
}

//fd state: (82) 8SOI8g1vTErtnsYj4BFWkWINkNGGZkSDLjKFIPLcDzUVDYJivQNWn4i2h8p78MUhc0G1Fk6FYaqEQ6EPAn

ret = client.Write(fd705, []byte("tdtFd8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) 8SOI8g1vTErtnsYj4BFWkWINkNGGZkSDLjKFIPLcDzUVDYJivQNWn4i2h8p78MUhc0G1Fk6FYaqEQ6EPAntdtFd8

fd730 := client.Open("/6AwM3ujffY.txt", client.O_RDWR|client.O_CREATE)
if(fd730 < 0) {
  panic("open failed")
}


ret = client.Close(fd730)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd727)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd725, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1mHJs7QmW_9M7LhS5Ox5JDhXGJRYQe0NDB8Jr946LyFosw5voP_G") {
  panic("wrong data returned")
}


fd731 := client.Open("/uV1nPoDUtI.txt", client.O_RDWR|client.O_CREATE)
if(fd731 < 0) {
  panic("open failed")
}


ret = client.Seek(fd723, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd731, []byte("6Pa8c574ZlFOPq8oFDtIQ9gMMS921HAfj4euCCFQIq9JVVsaf91F1u739GOds_0ggdefL7Q2JpIQEzVpU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) 6Pa8c574ZlFOPq8oFDtIQ9gMMS921HAfj4euCCFQIq9JVVsaf91F1u739GOds_0ggdefL7Q2JpIQEzVpU

ret = client.Close(fd709)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd705, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


fd732 := client.Open("/l0WXb6xP4b.txt", client.O_RDWR|client.O_CREATE)
if(fd732 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd725, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd733 := client.Open("/jAMfPtHSb4.txt", client.O_RDWR|client.O_CREATE)
if(fd733 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd731, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd723, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "v7I6EESZ9QmcRkXtP9wUKmJ0") {
  panic("wrong data returned")
}


ret = client.Seek(fd732, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Close(fd694)
if(ret != 0) {
  panic("close failed")
}


fd734 := client.Open("/xkdPTt0bab.txt", client.O_RDWR|client.O_CREATE)
if(fd734 < 0) {
  panic("open failed")
}


ret = client.Close(fd721)
if(ret != 0) {
  panic("close failed")
}

//fd state: (96) dNGrZCCcgyXVAeLpBMIsMU66wVwpfxwWhTdOP9cuVCzy7J86QYCDFNX1PgFk_n44iTJpDsHCv7I6EESZ9QmcRkXtP9wUKmJ0

ret = client.Write(fd723, []byte("tNmot2s7DP_ZWY9Q3eew1qWd_l4gjaNiFWyavIekDzdoEGq_Ja0vrDjdgoSEEoCc5uKQwOIPRH6qr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (173) dNGrZCCcgyXVAeLpBMIsMU66wVwpfxwWhTdOP9cuVCzy7J86QYCDFNX1PgFk_n44iTJpDsHCv7I6EESZ9QmcRkXtP9wUKmJ0tNmot2s7DP_ZWY9Q3eew1qWd_l4gjaNiFWyavIekDzdoEGq_Ja0vrDjdgoSEEoCc5uKQwOIPRH6qr

buf, ret = client.Read(fd705, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "sYj4BFWkWINkNGGZkSDLjKFIPLcDzUVDYJivQNWn4i2h8p78MUhc0G1Fk6FYaqEQ6EPAntdtFd8") {
  panic("wrong data returned")
}


fd735 := client.Open("/D2eWjCFlOC.txt", client.O_RDWR|client.O_CREATE)
if(fd735 < 0) {
  panic("open failed")
}


ret = client.Close(fd735)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd732)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd723, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}


buf, ret = client.Read(fd725, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd733)
if(ret != 0) {
  panic("close failed")
}

//fd state: (81) 6Pa8c574ZlFOPq8oFDtIQ9gMMS921HAfj4euCCFQIq9JVVsaf91F1u739GOds_0ggdefL7Q2JpIQEzVpU

ret = client.Write(fd731, []byte("VbiXXIEMqoqaBR40o"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) 6Pa8c574ZlFOPq8oFDtIQ9gMMS921HAfj4euCCFQIq9JVVsaf91F1u739GOds_0ggdefL7Q2JpIQEzVpUVbiXXIEMqoqaBR40o
//fd state: (0) 

ret = client.Write(fd734, []byte("eBHOhiXShIGRN7OnHxP1kUQjXxcEkf1YaTO6Xg8XYI3kY03XOjBWdd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) eBHOhiXShIGRN7OnHxP1kUQjXxcEkf1YaTO6Xg8XYI3kY03XOjBWdd

buf, ret = client.Read(fd734, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd736 := client.Open("/SZ3w7aLh_u.txt", client.O_RDWR|client.O_CREATE)
if(fd736 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd723, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9QmcRkXtP9wUKmJ0tNmot2s7DP_ZWY9Q3eew1qWd") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd736, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "F3CyOeYJn7Jhk6sCp8QsXL_t9Kijfn0UdWXBwxVKllTVvrgrNxODQofVpmknDczTdhPHbh1gK2dnBUC") {
  panic("wrong data returned")
}


fd737 := client.Open("/AjUgpmjviU.txt", client.O_RDWR|client.O_CREATE)
if(fd737 < 0) {
  panic("open failed")
}


ret = client.Close(fd723)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd736, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


fd738 := client.Open("/8fT6sTxzZ9.txt", client.O_RDWR|client.O_CREATE)
if(fd738 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd736, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "n7Jhk6sCp8QsXL_t9Kijfn0UdWXBwxVKllTVvrgrNxODQofVpmknDczTdhPHbh1gK2dnBUC5AFp") {
  panic("wrong data returned")
}


ret = client.Close(fd731)
if(ret != 0) {
  panic("close failed")
}


fd739 := client.Open("/m8VIThj_Gv.txt", client.O_RDWR|client.O_CREATE)
if(fd739 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd739, []byte("O8ul0Kkl0eHkyRZUTu_hPZMzEDsZCECvXC00BSdmddJxAf6S_OvKnehtC5C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) O8ul0Kkl0eHkyRZUTu_hPZMzEDsZCECvXC00BSdmddJxAf6S_OvKnehtC5C

fd740 := client.Open("/JPG9dFmZGp.txt", client.O_RDWR|client.O_CREATE)
if(fd740 < 0) {
  panic("open failed")
}


ret = client.Close(fd738)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd739)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd705, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd736)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd725, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


ret = client.Seek(fd705, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


fd741 := client.Open("/U_EGQ6_dMX.txt", client.O_RDWR|client.O_CREATE)
if(fd741 < 0) {
  panic("open failed")
}


ret = client.Close(fd741)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd705, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


buf, ret = client.Read(fd705, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "cDzUVDYJivQNWn4i2h8p78MUhc0G1Fk6FYaq") {
  panic("wrong data returned")
}


ret = client.Seek(fd734, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


ret = client.Close(fd740)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd725)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd737, []byte("l"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) l

ret = client.Seek(fd737, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd705, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}


ret = client.Close(fd737)
if(ret != 0) {
  panic("close failed")
}


fd742 := client.Open("/8d9mIP7BXQ.txt", client.O_RDWR|client.O_CREATE)
if(fd742 < 0) {
  panic("open failed")
}


ret = client.Seek(fd742, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd734, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


buf, ret = client.Read(fd705, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YaqEQ6EPAntdtFd8") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd734, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "aTO6Xg") {
  panic("wrong data returned")
}


ret = client.Close(fd734)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd742, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd742, []byte("NbpqDsk5EZKyCYpFmLvvHaPRrd4Rv3CGFNJ0SrkyCVPod4OQU8O7ZqG1WPXJzEvW9XOQrRNEpHjYHCM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) NbpqDsk5EZKyCYpFmLvvHaPRrd4Rv3CGFNJ0SrkyCVPod4OQU8O7ZqG1WPXJzEvW9XOQrRNEpHjYHCM

buf, ret = client.Read(fd705, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (79) NbpqDsk5EZKyCYpFmLvvHaPRrd4Rv3CGFNJ0SrkyCVPod4OQU8O7ZqG1WPXJzEvW9XOQrRNEpHjYHCM

ret = client.Write(fd742, []byte("geMcrbou71TSlryUeEIS0f_eGYVKZiFYm3oNz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) NbpqDsk5EZKyCYpFmLvvHaPRrd4Rv3CGFNJ0SrkyCVPod4OQU8O7ZqG1WPXJzEvW9XOQrRNEpHjYHCMgeMcrbou71TSlryUeEIS0f_eGYVKZiFYm3oNz

fd743 := client.Open("/jcVS5hPSWx.txt", client.O_RDWR|client.O_CREATE)
if(fd743 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd705, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd705, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (2) 8SOI8g1vTErtnsYj4BFWkWINkNGGZkSDLjKFIPLcDzUVDYJivQNWn4i2h8p78MUhc0G1Fk6FYaqEQ6EPAntdtFd8

ret = client.Write(fd705, []byte("DRKIgsQ2d7HqhpaWyjbbd0SddNQGT3V7DWhUEyB9RrvPysZOppmGlj1CQETN4SBCO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) 8SDRKIgsQ2d7HqhpaWyjbbd0SddNQGT3V7DWhUEyB9RrvPysZOppmGlj1CQETN4SBCO1Fk6FYaqEQ6EPAntdtFd8
//fd state: (67) 8SDRKIgsQ2d7HqhpaWyjbbd0SddNQGT3V7DWhUEyB9RrvPysZOppmGlj1CQETN4SBCO1Fk6FYaqEQ6EPAntdtFd8

ret = client.Write(fd705, []byte("imQwlrZZY0CJAY_yXuvaEZfal6IBCs2LEM4uIwWk31lXH6hCDnoGaxdiXEswToaaCYm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) 8SDRKIgsQ2d7HqhpaWyjbbd0SddNQGT3V7DWhUEyB9RrvPysZOppmGlj1CQETN4SBCOimQwlrZZY0CJAY_yXuvaEZfal6IBCs2LEM4uIwWk31lXH6hCDnoGaxdiXEswToaaCYm
//fd state: (0) 

ret = client.Write(fd743, []byte("rMuaTGlIB1zttStHXsQd7ls0fad5cHFDGN6WL4uXZfDNg5FC1zW8Vd4qvEH5gJfSR_QslAodPvi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) rMuaTGlIB1zttStHXsQd7ls0fad5cHFDGN6WL4uXZfDNg5FC1zW8Vd4qvEH5gJfSR_QslAodPvi

ret = client.Close(fd742)
if(ret != 0) {
  panic("close failed")
}

//fd state: (75) rMuaTGlIB1zttStHXsQd7ls0fad5cHFDGN6WL4uXZfDNg5FC1zW8Vd4qvEH5gJfSR_QslAodPvi

ret = client.Write(fd743, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) rMuaTGlIB1zttStHXsQd7ls0fad5cHFDGN6WL4uXZfDNg5FC1zW8Vd4qvEH5gJfSR_QslAodPvi

ret = client.Close(fd743)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd705, 109, client.SEEK_SET)
if(ret != 109) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 109)
  panic("seek failed")
}


ret = client.Seek(fd705, 60, client.SEEK_SET)
if(ret != 60) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 60)
  panic("seek failed")
}


fd744 := client.Open("/hqQ4RRTmCR.txt", client.O_RDWR|client.O_CREATE)
if(fd744 < 0) {
  panic("open failed")
}


ret = client.Seek(fd705, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


buf, ret = client.Read(fd744, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd705, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_yXuvaEZfal6IBCs2LEM4uIwWk") {
  panic("wrong data returned")
}


fd745 := client.Open("/b9IOiiNWX_.txt", client.O_RDWR|client.O_CREATE)
if(fd745 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd744, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd745, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd744, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd746 := client.Open("/AHpFNlzRlq.txt", client.O_RDWR|client.O_CREATE)
if(fd746 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd745, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd744, []byte("lF9PRZHyVEvw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) lF9PRZHyVEvw

buf, ret = client.Read(fd744, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd746, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd747 := client.Open("/6pv1r2ZR8B.txt", client.O_RDWR|client.O_CREATE)
if(fd747 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd705, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "31lXH6hCDnoGaxdiXEswToaaCYm") {
  panic("wrong data returned")
}


ret = client.Seek(fd747, 221, client.SEEK_SET)
if(ret != 221) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 221)
  panic("seek failed")
}


ret = client.Close(fd746)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd705)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd745, []byte("I5KoskDjPxTLUGmL4w5CIF6CNsnckHO7aOKkAxyY3uFcVFyHr4oov"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) I5KoskDjPxTLUGmL4w5CIF6CNsnckHO7aOKkAxyY3uFcVFyHr4oov

buf, ret = client.Read(fd745, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd748 := client.Open("/UukjxufYwH.txt", client.O_RDWR|client.O_CREATE)
if(fd748 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd747, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "z5UKnEWkVjbzNDOzQOg5pwBYJt7WYkUdkYlKaLGLYtw9BOV") {
  panic("wrong data returned")
}


ret = client.Close(fd747)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd748, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (53) I5KoskDjPxTLUGmL4w5CIF6CNsnckHO7aOKkAxyY3uFcVFyHr4oov

ret = client.Write(fd745, []byte("usQJCa2ov9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) I5KoskDjPxTLUGmL4w5CIF6CNsnckHO7aOKkAxyY3uFcVFyHr4oovusQJCa2ov9

ret = client.Seek(fd748, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd744)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd745, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


ret = client.Seek(fd745, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


buf, ret = client.Read(fd745, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PxTLUGmL4w5CIF6CNsnckHO7aOKkAxyY3uFcVFyHr4oovusQJCa2ov9") {
  panic("wrong data returned")
}


fd749 := client.Open("/BYVaUwRVUC.txt", client.O_RDWR|client.O_CREATE)
if(fd749 < 0) {
  panic("open failed")
}


fd750 := client.Open("/9g7Nr3wNRE.txt", client.O_RDWR|client.O_CREATE)
if(fd750 < 0) {
  panic("open failed")
}


fd751 := client.Open("/AHpFNlzRlq.txt", client.O_RDWR|client.O_CREATE)
if(fd751 < 0) {
  panic("open failed")
}


ret = client.Close(fd748)
if(ret != 0) {
  panic("close failed")
}


fd752 := client.Open("/ALVR8XJH6r.txt", client.O_RDWR|client.O_CREATE)
if(fd752 < 0) {
  panic("open failed")
}


fd753 := client.Open("/YfmUYg5AyQ.txt", client.O_RDWR|client.O_CREATE)
if(fd753 < 0) {
  panic("open failed")
}


fd754 := client.Open("/TkdRT14v9h.txt", client.O_RDWR|client.O_CREATE)
if(fd754 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd749, []byte("zCUgKeCLMCDjgIKO_yc1hFbigmibCY9U5XSY6y40uPcLFshUxvvdl4TZO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) zCUgKeCLMCDjgIKO_yc1hFbigmibCY9U5XSY6y40uPcLFshUxvvdl4TZO

buf, ret = client.Read(fd745, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd753, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8") {
  panic("wrong data returned")
}


ret = client.Seek(fd745, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


buf, ret = client.Read(fd751, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (35) I5KoskDjPxTLUGmL4w5CIF6CNsnckHO7aOKkAxyY3uFcVFyHr4oovusQJCa2ov9

ret = client.Write(fd745, []byte("NDE1CAO7p3ZhjisODBkwnrlo45_E0QaaofVUINbhvWcIS9EwGUhNN1T6jWS_GadfPWVAXMB_Fa5b47lfgXCJ2bYHGN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (125) I5KoskDjPxTLUGmL4w5CIF6CNsnckHO7aOKNDE1CAO7p3ZhjisODBkwnrlo45_E0QaaofVUINbhvWcIS9EwGUhNN1T6jWS_GadfPWVAXMB_Fa5b47lfgXCJ2bYHGN

buf, ret = client.Read(fd750, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "w6dQu_3UuiXEOZn2qzOoUnU2tID8jyGGXno5napHXwuWbmlXAg4kVb4HQqsPaS9XvlzI0pHdLAL3Qh0GP0") {
  panic("wrong data returned")
}

//fd state: (82) w6dQu_3UuiXEOZn2qzOoUnU2tID8jyGGXno5napHXwuWbmlXAg4kVb4HQqsPaS9XvlzI0pHdLAL3Qh0GP0QD3ZJllUo

ret = client.Write(fd750, []byte("ZDtBpwvtaSarOTITW6dwjsRFllzEhc466mFTgPbSogLX4tnswtMw6zX6wst3H1BAilFMzAFQmjo7H0Xry5iZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (166) w6dQu_3UuiXEOZn2qzOoUnU2tID8jyGGXno5napHXwuWbmlXAg4kVb4HQqsPaS9XvlzI0pHdLAL3Qh0GP0ZDtBpwvtaSarOTITW6dwjsRFllzEhc466mFTgPbSogLX4tnswtMw6zX6wst3H1BAilFMzAFQmjo7H0Xry5iZ

buf, ret = client.Read(fd745, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd755 := client.Open("/7KLt2RqLK2.txt", client.O_RDWR|client.O_CREATE)
if(fd755 < 0) {
  panic("open failed")
}


ret = client.Close(fd755)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd749, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd753, 112, client.SEEK_SET)
if(ret != 112) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 112)
  panic("seek failed")
}


fd756 := client.Open("/wvtPwOJxDc.txt", client.O_RDWR|client.O_CREATE)
if(fd756 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd745, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (125) I5KoskDjPxTLUGmL4w5CIF6CNsnckHO7aOKNDE1CAO7p3ZhjisODBkwnrlo45_E0QaaofVUINbhvWcIS9EwGUhNN1T6jWS_GadfPWVAXMB_Fa5b47lfgXCJ2bYHGN

ret = client.Write(fd745, []byte("K10mky3dSscst"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (138) I5KoskDjPxTLUGmL4w5CIF6CNsnckHO7aOKNDE1CAO7p3ZhjisODBkwnrlo45_E0QaaofVUINbhvWcIS9EwGUhNN1T6jWS_GadfPWVAXMB_Fa5b47lfgXCJ2bYHGNK10mky3dSscst

buf, ret = client.Read(fd753, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GlOanht9") {
  panic("wrong data returned")
}


ret = client.Seek(fd754, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd757 := client.Open("/56n5wy1x4r.txt", client.O_RDWR|client.O_CREATE)
if(fd757 < 0) {
  panic("open failed")
}


ret = client.Close(fd757)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd756)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd753, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd754, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd749, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Close(fd745)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd750, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd751, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd758 := client.Open("/_mbJ__Dujj.txt", client.O_RDWR|client.O_CREATE)
if(fd758 < 0) {
  panic("open failed")
}


fd759 := client.Open("/JawSnCt88v.txt", client.O_RDWR|client.O_CREATE)
if(fd759 < 0) {
  panic("open failed")
}


ret = client.Close(fd758)
if(ret != 0) {
  panic("close failed")
}

//fd state: (120) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9

ret = client.Write(fd753, []byte("bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (214) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hg

fd760 := client.Open("/90hEKufnYM.txt", client.O_RDWR|client.O_CREATE)
if(fd760 < 0) {
  panic("open failed")
}


fd761 := client.Open("/1_pU8JJR0U.txt", client.O_RDWR|client.O_CREATE)
if(fd761 < 0) {
  panic("open failed")
}


fd762 := client.Open("/b2Gr0PrgBp.txt", client.O_RDWR|client.O_CREATE)
if(fd762 < 0) {
  panic("open failed")
}


ret = client.Close(fd754)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd752, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "SUbZTwa5r_XprdQjLe7wEjPjmqF5M4DH7M98idROQJyTauPsSmN7WVU0RAvBsn") {
  panic("wrong data returned")
}


ret = client.Close(fd760)
if(ret != 0) {
  panic("close failed")
}


fd763 := client.Open("/0_jDN0GKli.txt", client.O_RDWR|client.O_CREATE)
if(fd763 < 0) {
  panic("open failed")
}

//fd state: (214) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hg

ret = client.Write(fd753, []byte("jEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (292) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfX

fd764 := client.Open("/ewflQBQZRS.txt", client.O_RDWR|client.O_CREATE)
if(fd764 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd761, []byte("iJpQ1HjPqsrTCIFX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (16) iJpQ1HjPqsrTCIFX

ret = client.Seek(fd750, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


fd765 := client.Open("/SZD6wU8PL1.txt", client.O_RDWR|client.O_CREATE)
if(fd765 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd751, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd752, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}

//fd state: (28) w6dQu_3UuiXEOZn2qzOoUnU2tID8jyGGXno5napHXwuWbmlXAg4kVb4HQqsPaS9XvlzI0pHdLAL3Qh0GP0ZDtBpwvtaSarOTITW6dwjsRFllzEhc466mFTgPbSogLX4tnswtMw6zX6wst3H1BAilFMzAFQmjo7H0Xry5iZ

ret = client.Write(fd750, []byte("kw5M7mjiDuczuBfTb8uztx9aT1LNxPqRsjxgfo8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) w6dQu_3UuiXEOZn2qzOoUnU2tID8kw5M7mjiDuczuBfTb8uztx9aT1LNxPqRsjxgfo8I0pHdLAL3Qh0GP0ZDtBpwvtaSarOTITW6dwjsRFllzEhc466mFTgPbSogLX4tnswtMw6zX6wst3H1BAilFMzAFQmjo7H0Xry5iZ

fd766 := client.Open("/9jmiz_VRmL.txt", client.O_RDWR|client.O_CREATE)
if(fd766 < 0) {
  panic("open failed")
}


ret = client.Seek(fd762, 124, client.SEEK_SET)
if(ret != 124) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 124)
  panic("seek failed")
}


ret = client.Seek(fd765, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd766, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd749)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd762, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "KoF7dy4yf8bLIB833GLtuRkmqEG") {
  panic("wrong data returned")
}


ret = client.Close(fd762)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd759, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd752, 68, client.SEEK_SET)
if(ret != 68) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 68)
  panic("seek failed")
}


buf, ret = client.Read(fd752, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "kNfccksDkHSLyFzOYmMpNbi") {
  panic("wrong data returned")
}


ret = client.Close(fd761)
if(ret != 0) {
  panic("close failed")
}

//fd state: (292) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfX

ret = client.Write(fd753, []byte("aU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (294) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfXaU

ret = client.Seek(fd759, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd767 := client.Open("/D24_Z9DpA2.txt", client.O_RDWR|client.O_CREATE)
if(fd767 < 0) {
  panic("open failed")
}


ret = client.Close(fd764)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd767, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd759)
if(ret != 0) {
  panic("close failed")
}


fd768 := client.Open("/y0dHmoJ4hA.txt", client.O_RDWR|client.O_CREATE)
if(fd768 < 0) {
  panic("open failed")
}


ret = client.Seek(fd750, 140, client.SEEK_SET)
if(ret != 140) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 140)
  panic("seek failed")
}


buf, ret = client.Read(fd753, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd767, []byte("EKuAKSo0u1MDgE3G6mosTGkI3DNglaesXiLXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Ayw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) EKuAKSo0u1MDgE3G6mosTGkI3DNglaesXiLXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Ayw
//fd state: (0) 

ret = client.Write(fd751, []byte("LWNgmCe4DaFAnSdZ66rgcDAWGAdwGEyPttAcFZb2fvPyZuItv110WzRHbAX_xAweLyZA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) LWNgmCe4DaFAnSdZ66rgcDAWGAdwGEyPttAcFZb2fvPyZuItv110WzRHbAX_xAweLyZA

ret = client.Seek(fd763, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}

//fd state: (77) EKuAKSo0u1MDgE3G6mosTGkI3DNglaesXiLXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Ayw

ret = client.Write(fd767, []byte("x96fGgHygyhE5RBRS4lwzeX4EGxRsvomj6Ex96G_Dd255OLVd7Nd3Nj2LaPwghcAWj1mPIWhPKkBjd7OQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (158) EKuAKSo0u1MDgE3G6mosTGkI3DNglaesXiLXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Aywx96fGgHygyhE5RBRS4lwzeX4EGxRsvomj6Ex96G_Dd255OLVd7Nd3Nj2LaPwghcAWj1mPIWhPKkBjd7OQ

fd769 := client.Open("/2Kwg_3a7Gr.txt", client.O_RDWR|client.O_CREATE)
if(fd769 < 0) {
  panic("open failed")
}


ret = client.Seek(fd750, 141, client.SEEK_SET)
if(ret != 141) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 141)
  panic("seek failed")
}


ret = client.Close(fd769)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd751)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd768)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd766)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd765, []byte("JPgko2upu2zguld_mC3hGZ23NDQ8Ba6alKfTmrAZCVWTpecfSSNCna9gpV7eidOd4WhKZ128l9cowlsstKWA9mznTcX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) JPgko2upu2zguld_mC3hGZ23NDQ8Ba6alKfTmrAZCVWTpecfSSNCna9gpV7eidOd4WhKZ128l9cowlsstKWA9mznTcX

ret = client.Seek(fd767, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


buf, ret = client.Read(fd753, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (91) SUbZTwa5r_XprdQjLe7wEjPjmqF5M4DH7M98idROQJyTauPsSmN7WVU0RAvBsnrhuKLKkNfccksDkHSLyFzOYmMpNbi

ret = client.Write(fd752, []byte("FV7uDEq3L6hh76XIVeWV6c4y_e1ptIFlYveuH747pYvVkAjg8__mzkck5j3o1HcxAwlBIJ_PY2AmSCpxvNNb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (175) SUbZTwa5r_XprdQjLe7wEjPjmqF5M4DH7M98idROQJyTauPsSmN7WVU0RAvBsnrhuKLKkNfccksDkHSLyFzOYmMpNbiFV7uDEq3L6hh76XIVeWV6c4y_e1ptIFlYveuH747pYvVkAjg8__mzkck5j3o1HcxAwlBIJ_PY2AmSCpxvNNb

buf, ret = client.Read(fd767, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Aywx96fGgHygyhE5RBRS4lwzeX4EGxRsvomj") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd765, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (110) EKuAKSo0u1MDgE3G6mosTGkI3DNglaesXiLXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Aywx96fGgHygyhE5RBRS4lwzeX4EGxRsvomj6Ex96G_Dd255OLVd7Nd3Nj2LaPwghcAWj1mPIWhPKkBjd7OQ

ret = client.Write(fd767, []byte("XMkFsInZ2_M3HG3gkQuGtOAEt5HBTaWYCSaeJk5JELhVkazJUdz1ipkCTBWxaYDUmB_60RPkbnc4mWN5Sn68OomwOglnlUj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (205) EKuAKSo0u1MDgE3G6mosTGkI3DNglaesXiLXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Aywx96fGgHygyhE5RBRS4lwzeX4EGxRsvomjXMkFsInZ2_M3HG3gkQuGtOAEt5HBTaWYCSaeJk5JELhVkazJUdz1ipkCTBWxaYDUmB_60RPkbnc4mWN5Sn68OomwOglnlUj
//fd state: (44) KHgGg00XJURyKK4nWTOLQ0two6TZYZE2Jl0WaV4UrtST8yhFLbwACI0MXc3LvYHlMryDegTKX0fm5gTEHaV3E71nwgPSP3VVF7g1Ko6ka

ret = client.Write(fd763, []byte("XAlZCgcFCxY_34J8BHBxg80Htl2mqlERO3roHl0G4fnZOAAiW0uqDIH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) KHgGg00XJURyKK4nWTOLQ0two6TZYZE2Jl0WaV4UrtSTXAlZCgcFCxY_34J8BHBxg80Htl2mqlERO3roHl0G4fnZOAAiW0uqDIH1Ko6ka

ret = client.Close(fd765)
if(ret != 0) {
  panic("close failed")
}


fd770 := client.Open("/T5Ff3G8IgM.txt", client.O_RDWR|client.O_CREATE)
if(fd770 < 0) {
  panic("open failed")
}


ret = client.Seek(fd752, 145, client.SEEK_SET)
if(ret != 145) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 145)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (294) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfXaU

ret = client.Write(fd753, []byte("ulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (369) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfXaUulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOy

ret = client.Close(fd752)
if(ret != 0) {
  panic("close failed")
}


fd771 := client.Open("/685Nb0NdyX.txt", client.O_RDWR|client.O_CREATE)
if(fd771 < 0) {
  panic("open failed")
}


fd772 := client.Open("/VTasqgKgMi.txt", client.O_RDWR|client.O_CREATE)
if(fd772 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd770, []byte("qWRFEkKsEBioMONs5nTzIgxREhqaj8I2avld8w3q1xzeUiwSwYOpuGUyM0MWH3EBV773lzfYmLwu9USbttGYZhq6wKChCfKzyJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) qWRFEkKsEBioMONs5nTzIgxREhqaj8I2avld8w3q1xzeUiwSwYOpuGUyM0MWH3EBV773lzfYmLwu9USbttGYZhq6wKChCfKzyJ
//fd state: (205) EKuAKSo0u1MDgE3G6mosTGkI3DNglaesXiLXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Aywx96fGgHygyhE5RBRS4lwzeX4EGxRsvomjXMkFsInZ2_M3HG3gkQuGtOAEt5HBTaWYCSaeJk5JELhVkazJUdz1ipkCTBWxaYDUmB_60RPkbnc4mWN5Sn68OomwOglnlUj

ret = client.Write(fd767, []byte("GgnohotQbydllUgWfNCBh01GNevYDiJcSwmMJ5Qe9QdJGCklaqOEh9JeQMHWEwMFaEK5stL2ZOsmq4ocheEgl7JI4ragcvx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (300) EKuAKSo0u1MDgE3G6mosTGkI3DNglaesXiLXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Aywx96fGgHygyhE5RBRS4lwzeX4EGxRsvomjXMkFsInZ2_M3HG3gkQuGtOAEt5HBTaWYCSaeJk5JELhVkazJUdz1ipkCTBWxaYDUmB_60RPkbnc4mWN5Sn68OomwOglnlUjGgnohotQbydllUgWfNCBh01GNevYDiJcSwmMJ5Qe9QdJGCklaqOEh9JeQMHWEwMFaEK5stL2ZOsmq4ocheEgl7JI4ragcvx

ret = client.Seek(fd770, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


ret = client.Seek(fd750, 157, client.SEEK_SET)
if(ret != 157) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 157)
  panic("seek failed")
}

//fd state: (369) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfXaUulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOy

ret = client.Write(fd753, []byte("OnlCZNge1gb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (380) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfXaUulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOyOnlCZNge1gb
//fd state: (300) EKuAKSo0u1MDgE3G6mosTGkI3DNglaesXiLXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Aywx96fGgHygyhE5RBRS4lwzeX4EGxRsvomjXMkFsInZ2_M3HG3gkQuGtOAEt5HBTaWYCSaeJk5JELhVkazJUdz1ipkCTBWxaYDUmB_60RPkbnc4mWN5Sn68OomwOglnlUjGgnohotQbydllUgWfNCBh01GNevYDiJcSwmMJ5Qe9QdJGCklaqOEh9JeQMHWEwMFaEK5stL2ZOsmq4ocheEgl7JI4ragcvx

ret = client.Write(fd767, []byte("vVhRhOPcHAf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (311) EKuAKSo0u1MDgE3G6mosTGkI3DNglaesXiLXGAfo6Vgl3ALHxJMhGlfJEp4uOTRNq53NizOCf4Aywx96fGgHygyhE5RBRS4lwzeX4EGxRsvomjXMkFsInZ2_M3HG3gkQuGtOAEt5HBTaWYCSaeJk5JELhVkazJUdz1ipkCTBWxaYDUmB_60RPkbnc4mWN5Sn68OomwOglnlUjGgnohotQbydllUgWfNCBh01GNevYDiJcSwmMJ5Qe9QdJGCklaqOEh9JeQMHWEwMFaEK5stL2ZOsmq4ocheEgl7JI4ragcvxvVhRhOPcHAf
//fd state: (157) w6dQu_3UuiXEOZn2qzOoUnU2tID8kw5M7mjiDuczuBfTb8uztx9aT1LNxPqRsjxgfo8I0pHdLAL3Qh0GP0ZDtBpwvtaSarOTITW6dwjsRFllzEhc466mFTgPbSogLX4tnswtMw6zX6wst3H1BAilFMzAFQmjo7H0Xry5iZ

ret = client.Write(fd750, []byte("b"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (158) w6dQu_3UuiXEOZn2qzOoUnU2tID8kw5M7mjiDuczuBfTb8uztx9aT1LNxPqRsjxgfo8I0pHdLAL3Qh0GP0ZDtBpwvtaSarOTITW6dwjsRFllzEhc466mFTgPbSogLX4tnswtMw6zX6wst3H1BAilFMzAFQmjobH0Xry5iZ
//fd state: (158) w6dQu_3UuiXEOZn2qzOoUnU2tID8kw5M7mjiDuczuBfTb8uztx9aT1LNxPqRsjxgfo8I0pHdLAL3Qh0GP0ZDtBpwvtaSarOTITW6dwjsRFllzEhc466mFTgPbSogLX4tnswtMw6zX6wst3H1BAilFMzAFQmjobH0Xry5iZ

ret = client.Write(fd750, []byte("Eu69f1rrEAE6gmIlrztNu4by8PDTeTdbd075"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (194) w6dQu_3UuiXEOZn2qzOoUnU2tID8kw5M7mjiDuczuBfTb8uztx9aT1LNxPqRsjxgfo8I0pHdLAL3Qh0GP0ZDtBpwvtaSarOTITW6dwjsRFllzEhc466mFTgPbSogLX4tnswtMw6zX6wst3H1BAilFMzAFQmjobEu69f1rrEAE6gmIlrztNu4by8PDTeTdbd075

ret = client.Seek(fd753, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}

//fd state: (40) qWRFEkKsEBioMONs5nTzIgxREhqaj8I2avld8w3q1xzeUiwSwYOpuGUyM0MWH3EBV773lzfYmLwu9USbttGYZhq6wKChCfKzyJ

ret = client.Write(fd770, []byte("IHF9NFTwy7tMWm3QyTXDtoaIYopm7b9J_fZFW23WaMYtDJcdE6ODnOZdRNQx3BvXzsUiXDk1Sp_0KR_zOSMU48nMXHQQ1xEp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) qWRFEkKsEBioMONs5nTzIgxREhqaj8I2avld8w3qIHF9NFTwy7tMWm3QyTXDtoaIYopm7b9J_fZFW23WaMYtDJcdE6ODnOZdRNQx3BvXzsUiXDk1Sp_0KR_zOSMU48nMXHQQ1xEp

buf, ret = client.Read(fd753, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1") {
  panic("wrong data returned")
}


fd773 := client.Open("/cDsrk7sXrQ.txt", client.O_RDWR|client.O_CREATE)
if(fd773 < 0) {
  panic("open failed")
}

//fd state: (99) KHgGg00XJURyKK4nWTOLQ0two6TZYZE2Jl0WaV4UrtSTXAlZCgcFCxY_34J8BHBxg80Htl2mqlERO3roHl0G4fnZOAAiW0uqDIH1Ko6ka

ret = client.Write(fd763, []byte("VaCHC4wpIyYQsr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) KHgGg00XJURyKK4nWTOLQ0two6TZYZE2Jl0WaV4UrtSTXAlZCgcFCxY_34J8BHBxg80Htl2mqlERO3roHl0G4fnZOAAiW0uqDIHVaCHC4wpIyYQsr

ret = client.Seek(fd773, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd774 := client.Open("/CFuIs75vUa.txt", client.O_RDWR|client.O_CREATE)
if(fd774 < 0) {
  panic("open failed")
}

//fd state: (101) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1zCP8WnKRYAWGlOanht9bNHiyEeBwXXECvPBZQl4Xi8DTa2tKgIQ1XGKzUFmlOInBSR_6GgEpz8myYUsfYVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfXaUulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOyOnlCZNge1gb

ret = client.Write(fd753, []byte("atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (181) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfXaUulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOyOnlCZNge1gb
//fd state: (0) 

ret = client.Write(fd772, []byte("VjxjucZNESYlW4ALpVH2O_0BluUvHpOfpf6RC_tV6O46FH6HRIuAUplhSFvl8d1e83QrQatl2U0QyuLXM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) VjxjucZNESYlW4ALpVH2O_0BluUvHpOfpf6RC_tV6O46FH6HRIuAUplhSFvl8d1e83QrQatl2U0QyuLXM

fd775 := client.Open("/GqdBSdeMYc.txt", client.O_RDWR|client.O_CREATE)
if(fd775 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd773, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd772, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd772)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd753, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKgu") {
  panic("wrong data returned")
}


ret = client.Close(fd750)
if(ret != 0) {
  panic("close failed")
}


fd776 := client.Open("/bTnE2AVRQj.txt", client.O_RDWR|client.O_CREATE)
if(fd776 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd763, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd775, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd777 := client.Open("/QWehDPrfor.txt", client.O_RDWR|client.O_CREATE)
if(fd777 < 0) {
  panic("open failed")
}


ret = client.Seek(fd771, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd770, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd778 := client.Open("/jmVadlwBeg.txt", client.O_RDWR|client.O_CREATE)
if(fd778 < 0) {
  panic("open failed")
}


fd779 := client.Open("/xdFaEq7KDB.txt", client.O_RDWR|client.O_CREATE)
if(fd779 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd777, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd776, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd779, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd770, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "FEkKsEBioMONs5nTzIgxREhqaj8I2avld8w3qIHF9NFTwy7tMWm3QyTXDtoaIYopm7b9J_") {
  panic("wrong data returned")
}


ret = client.Seek(fd753, 197, client.SEEK_SET)
if(ret != 197) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 197)
  panic("seek failed")
}


ret = client.Close(fd778)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd777, []byte("Z3TbnTz73vbYO_byxHDQSxUWfd2bOrGDtotHIv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) Z3TbnTz73vbYO_byxHDQSxUWfd2bOrGDtotHIv

ret = client.Seek(fd773, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd773, []byte("7TobMM27p3DK8X5JHVwAGsWMz4XgnoB6rFKfLP7P9ClGaNAYCBIcDZR2R2hsl2mCa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) 7TobMM27p3DK8X5JHVwAGsWMz4XgnoB6rFKfLP7P9ClGaNAYCBIcDZR2R2hsl2mCa

ret = client.Close(fd770)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd773, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd771)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd763)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd774, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd767, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd780 := client.Open("/JKZTf7z8OC.txt", client.O_RDWR|client.O_CREATE)
if(fd780 < 0) {
  panic("open failed")
}


ret = client.Close(fd777)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd774)
if(ret != 0) {
  panic("close failed")
}


fd781 := client.Open("/fcHb173aGv.txt", client.O_RDWR|client.O_CREATE)
if(fd781 < 0) {
  panic("open failed")
}


ret = client.Close(fd767)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd776, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd776)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd773)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd779)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd780, []byte("Jl8hPyrSHnj5zG4frm9YEJYZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) Jl8hPyrSHnj5zG4frm9YEJYZ

ret = client.Seek(fd781, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd775, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd780, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd782 := client.Open("/9q9ndfi0_e.txt", client.O_RDWR|client.O_CREATE)
if(fd782 < 0) {
  panic("open failed")
}


ret = client.Close(fd782)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd781, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd781, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd783 := client.Open("/GrMYXD3DIK.txt", client.O_RDWR|client.O_CREATE)
if(fd783 < 0) {
  panic("open failed")
}

//fd state: (197) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALB3fD0rWrcPvOCwW8hgjEfHuKguEVfDUQiXF0hWltlCxSukfUIOCKU8kbmzYD6WTCyq79tdXMbjPiR8mNhSzSVnyHgYDA8nfXaUulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOyOnlCZNge1gb

ret = client.Write(fd753, []byte("Ik1danbkEhwTvHkoAuITMapEu1uUP5Cf1RFbJIRQpUx54upJbmdExk__jvIEuMiAyTjA15CB2ziDFLSPhcf7ezY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (284) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALBIk1danbkEhwTvHkoAuITMapEu1uUP5Cf1RFbJIRQpUx54upJbmdExk__jvIEuMiAyTjA15CB2ziDFLSPhcf7ezYgYDA8nfXaUulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOyOnlCZNge1gb
//fd state: (24) Jl8hPyrSHnj5zG4frm9YEJYZ

ret = client.Write(fd780, []byte("NP6I1HMoqlb9yJiLDbJgAqU4UViwUp9F1vLoyMfgH2bxvrAzA906ocnEhSxAQvEM1dQqiUaP74xTnYH_CECjrCp2zhxUYX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) Jl8hPyrSHnj5zG4frm9YEJYZNP6I1HMoqlb9yJiLDbJgAqU4UViwUp9F1vLoyMfgH2bxvrAzA906ocnEhSxAQvEM1dQqiUaP74xTnYH_CECjrCp2zhxUYX

ret = client.Seek(fd775, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (284) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALBIk1danbkEhwTvHkoAuITMapEu1uUP5Cf1RFbJIRQpUx54upJbmdExk__jvIEuMiAyTjA15CB2ziDFLSPhcf7ezYgYDA8nfXaUulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOyOnlCZNge1gb

ret = client.Write(fd753, []byte("wsxUYsP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (291) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALBIk1danbkEhwTvHkoAuITMapEu1uUP5Cf1RFbJIRQpUx54upJbmdExk__jvIEuMiAyTjA15CB2ziDFLSPhcf7ezYwsxUYsPXaUulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOyOnlCZNge1gb

ret = client.Close(fd783)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd781, []byte("i3lKaeXkv2G03lkeD0igYophfi45LwKA8ifN31wNtGzksA8AErhoYOZy4yByCaR5VKHxO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) i3lKaeXkv2G03lkeD0igYophfi45LwKA8ifN31wNtGzksA8AErhoYOZy4yByCaR5VKHxO

ret = client.Close(fd775)
if(ret != 0) {
  panic("close failed")
}

//fd state: (291) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALBIk1danbkEhwTvHkoAuITMapEu1uUP5Cf1RFbJIRQpUx54upJbmdExk__jvIEuMiAyTjA15CB2ziDFLSPhcf7ezYwsxUYsPXaUulm5wcFwqo03MS2NhpC7fnws9GXdynmiJVxj7uby6RAxzV2D1MxVjGVYySGfwcPVY0FcLtJLzOyOnlCZNge1gb

ret = client.Write(fd753, []byte("XgWk46AvlSP_GgsEWn5LunlYKEpU55HX84cwwMwqVjIO3Lgr92svd6XWSEFZdpHvZkewW82f1uNgK5lW_m3sj2g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (378) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALBIk1danbkEhwTvHkoAuITMapEu1uUP5Cf1RFbJIRQpUx54upJbmdExk__jvIEuMiAyTjA15CB2ziDFLSPhcf7ezYwsxUYsPXgWk46AvlSP_GgsEWn5LunlYKEpU55HX84cwwMwqVjIO3Lgr92svd6XWSEFZdpHvZkewW82f1uNgK5lW_m3sj2ggb

fd784 := client.Open("/_pPqWlscvp.txt", client.O_RDWR|client.O_CREATE)
if(fd784 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd780, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd780, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}


ret = client.Seek(fd781, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


fd785 := client.Open("/_gbGbZZtm6.txt", client.O_RDWR|client.O_CREATE)
if(fd785 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd753, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gb") {
  panic("wrong data returned")
}

//fd state: (50) Jl8hPyrSHnj5zG4frm9YEJYZNP6I1HMoqlb9yJiLDbJgAqU4UViwUp9F1vLoyMfgH2bxvrAzA906ocnEhSxAQvEM1dQqiUaP74xTnYH_CECjrCp2zhxUYX

ret = client.Write(fd780, []byte("tDqB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) Jl8hPyrSHnj5zG4frm9YEJYZNP6I1HMoqlb9yJiLDbJgAqU4UVtDqB9F1vLoyMfgH2bxvrAzA906ocnEhSxAQvEM1dQqiUaP74xTnYH_CECjrCp2zhxUYX

fd786 := client.Open("/wVLioSOIGX.txt", client.O_RDWR|client.O_CREATE)
if(fd786 < 0) {
  panic("open failed")
}


fd787 := client.Open("/t5jEm7yeiJ.txt", client.O_RDWR|client.O_CREATE)
if(fd787 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd780, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9F1vLoyMfgH2bxvrAzA906ocnEhSxAQvEM1dQqiUaP74xTnYH_CECjrCp2z") {
  panic("wrong data returned")
}

//fd state: (380) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALBIk1danbkEhwTvHkoAuITMapEu1uUP5Cf1RFbJIRQpUx54upJbmdExk__jvIEuMiAyTjA15CB2ziDFLSPhcf7ezYwsxUYsPXgWk46AvlSP_GgsEWn5LunlYKEpU55HX84cwwMwqVjIO3Lgr92svd6XWSEFZdpHvZkewW82f1uNgK5lW_m3sj2ggb

ret = client.Write(fd753, []byte("H8N6pdRSU6bSaob93a1rDY7navR8SwWmBRGmLmheThOvpTb07z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (430) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALBIk1danbkEhwTvHkoAuITMapEu1uUP5Cf1RFbJIRQpUx54upJbmdExk__jvIEuMiAyTjA15CB2ziDFLSPhcf7ezYwsxUYsPXgWk46AvlSP_GgsEWn5LunlYKEpU55HX84cwwMwqVjIO3Lgr92svd6XWSEFZdpHvZkewW82f1uNgK5lW_m3sj2ggbH8N6pdRSU6bSaob93a1rDY7navR8SwWmBRGmLmheThOvpTb07z

ret = client.Close(fd780)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd784)
if(ret != 0) {
  panic("close failed")
}


fd788 := client.Open("/fyMbR1YJOy.txt", client.O_RDWR|client.O_CREATE)
if(fd788 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd753, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd781, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


ret = client.Seek(fd786, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd787, []byte("bMGNonU9ethbElp5KbiR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) bMGNonU9ethbElp5KbiR

buf, ret = client.Read(fd785, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "zymbI2sxWd7cv8rGAegfVguM8IzyCip1NB4lqUFBdeHRx7") {
  panic("wrong data returned")
}

//fd state: (430) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALBIk1danbkEhwTvHkoAuITMapEu1uUP5Cf1RFbJIRQpUx54upJbmdExk__jvIEuMiAyTjA15CB2ziDFLSPhcf7ezYwsxUYsPXgWk46AvlSP_GgsEWn5LunlYKEpU55HX84cwwMwqVjIO3Lgr92svd6XWSEFZdpHvZkewW82f1uNgK5lW_m3sj2ggbH8N6pdRSU6bSaob93a1rDY7navR8SwWmBRGmLmheThOvpTb07z

ret = client.Write(fd753, []byte("v35RVhsnCew80ilmLSIWUi5YhehArtpEcgJaQZGamkheeou843eDruW6EoV_n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (491) s3CiGwjJLhEtf9VYgVX4jQDPmrQYaS2OCBm4FSVyASCIzb29r9tWgcoqSb9fgHgtOfXaa8fI9F1McY0A8d1tXhFCbQ_8QYkYLcI_1atLc3M3XTIj1k9AXnUCIDpPUBzhy30f7T4gRT8Xhxtv3JA8JgNKhhFcv_ePEcQHhs7DLkPMcLtj1BKv6YVXn0ByZKedKQALBIk1danbkEhwTvHkoAuITMapEu1uUP5Cf1RFbJIRQpUx54upJbmdExk__jvIEuMiAyTjA15CB2ziDFLSPhcf7ezYwsxUYsPXgWk46AvlSP_GgsEWn5LunlYKEpU55HX84cwwMwqVjIO3Lgr92svd6XWSEFZdpHvZkewW82f1uNgK5lW_m3sj2ggbH8N6pdRSU6bSaob93a1rDY7navR8SwWmBRGmLmheThOvpTb07zv35RVhsnCew80ilmLSIWUi5YhehArtpEcgJaQZGamkheeou843eDruW6EoV_n
//fd state: (20) bMGNonU9ethbElp5KbiR

ret = client.Write(fd787, []byte("Cnptr06zLrZjaNxXoYWW_0WHSDnyw9F1z34gZxP0GJzafD8h5orG6fuRXlHI8yWtnQn0WgpNOwxzSIdAAb08e6XEEwBzLwtmncK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (119) bMGNonU9ethbElp5KbiRCnptr06zLrZjaNxXoYWW_0WHSDnyw9F1z34gZxP0GJzafD8h5orG6fuRXlHI8yWtnQn0WgpNOwxzSIdAAb08e6XEEwBzLwtmncK

ret = client.Close(fd787)
if(ret != 0) {
  panic("close failed")
}


fd789 := client.Open("/fyMbR1YJOy.txt", client.O_RDWR|client.O_CREATE)
if(fd789 < 0) {
  panic("open failed")
}


ret = client.Close(fd753)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd789)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd788)
if(ret != 0) {
  panic("close failed")
}

//fd state: (97) MnK7Gr6dgNvZNuI2BmJHw91L4Ywc7hRi_i0VRlJ2dae13VXg1V8mQ8GrnlQqg4Ab_Dn2Aa4qQZ7CiBdXsxWYk8GrSDhfWWyTmmjDqjwc3

ret = client.Write(fd786, []byte("65hhkKTUqJP9KcHU0cwJjquIP7Orp8lJ3nDm2k0raOwOjDFv011UoVIY45rQxAz5SE7C_9vDxpP0f1vxJqOLtw8q7cP3a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (190) MnK7Gr6dgNvZNuI2BmJHw91L4Ywc7hRi_i0VRlJ2dae13VXg1V8mQ8GrnlQqg4Ab_Dn2Aa4qQZ7CiBdXsxWYk8GrSDhfWWyTm65hhkKTUqJP9KcHU0cwJjquIP7Orp8lJ3nDm2k0raOwOjDFv011UoVIY45rQxAz5SE7C_9vDxpP0f1vxJqOLtw8q7cP3a
//fd state: (48) i3lKaeXkv2G03lkeD0igYophfi45LwKA8ifN31wNtGzksA8AErhoYOZy4yByCaR5VKHxO

ret = client.Write(fd781, []byte("zSyOUvLvJBti3fr4iU6UtvfZUq0sXKOdNlsvdhg0g4BIp_DbLCQXEb7v8eYF7Tu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) i3lKaeXkv2G03lkeD0igYophfi45LwKA8ifN31wNtGzksA8AzSyOUvLvJBti3fr4iU6UtvfZUq0sXKOdNlsvdhg0g4BIp_DbLCQXEb7v8eYF7Tu

ret = client.Close(fd781)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd786)
if(ret != 0) {
  panic("close failed")
}

//fd state: (46) zymbI2sxWd7cv8rGAegfVguM8IzyCip1NB4lqUFBdeHRx7

ret = client.Write(fd785, []byte("eqVl4b8pdvdwBYZMvYEODZSwin0ezowaJ8Hq8O3Vh9yqtP6ba0uPhwr_wbJsr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) zymbI2sxWd7cv8rGAegfVguM8IzyCip1NB4lqUFBdeHRx7eqVl4b8pdvdwBYZMvYEODZSwin0ezowaJ8Hq8O3Vh9yqtP6ba0uPhwr_wbJsr
//fd state: (107) zymbI2sxWd7cv8rGAegfVguM8IzyCip1NB4lqUFBdeHRx7eqVl4b8pdvdwBYZMvYEODZSwin0ezowaJ8Hq8O3Vh9yqtP6ba0uPhwr_wbJsr

ret = client.Write(fd785, []byte("XjWpYnaW2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) zymbI2sxWd7cv8rGAegfVguM8IzyCip1NB4lqUFBdeHRx7eqVl4b8pdvdwBYZMvYEODZSwin0ezowaJ8Hq8O3Vh9yqtP6ba0uPhwr_wbJsrXjWpYnaW2

ret = client.Seek(fd785, 76, client.SEEK_SET)
if(ret != 76) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 76)
  panic("seek failed")
}


ret = client.Close(fd785)
if(ret != 0) {
  panic("close failed")
}


fd790 := client.Open("/7KLt2RqLK2.txt", client.O_RDWR|client.O_CREATE)
if(fd790 < 0) {
  panic("open failed")
}


ret = client.Close(fd790)
if(ret != 0) {
  panic("close failed")
}


fd791 := client.Open("/d6tVKLHZtE.txt", client.O_RDWR|client.O_CREATE)
if(fd791 < 0) {
  panic("open failed")
}


ret = client.Seek(fd791, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd791, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd791, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd791, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd791, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd792 := client.Open("/KxLxGxYJjk.txt", client.O_RDWR|client.O_CREATE)
if(fd792 < 0) {
  panic("open failed")
}


fd793 := client.Open("/_xt2CaFERA.txt", client.O_RDWR|client.O_CREATE)
if(fd793 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd792, []byte("BwoVg4CrlHTGXSkJE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) BwoVg4CrlHTGXSkJE
//fd state: (0) 

ret = client.Write(fd793, []byte("qOUGsZ9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) qOUGsZ9

ret = client.Seek(fd793, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


buf, ret = client.Read(fd792, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd793, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GsZ9") {
  panic("wrong data returned")
}

//fd state: (7) qOUGsZ9

ret = client.Write(fd793, []byte("LGQOsbR3pNwZ8zOuvE0_ptjW5XbICFZA1hoKGkL2EMk4LCe1QE4nNQTq5dPZHnWYHxt88yzCmK3j42Q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) qOUGsZ9LGQOsbR3pNwZ8zOuvE0_ptjW5XbICFZA1hoKGkL2EMk4LCe1QE4nNQTq5dPZHnWYHxt88yzCmK3j42Q

ret = client.Close(fd791)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd793, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (17) BwoVg4CrlHTGXSkJE

ret = client.Write(fd792, []byte("J0papzdWH6q1isIJ6oxL24jWKboE9sfB6vdZkk5SIKBfooEbXNceuDlQikFEyycRYjX52gczT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) BwoVg4CrlHTGXSkJEJ0papzdWH6q1isIJ6oxL24jWKboE9sfB6vdZkk5SIKBfooEbXNceuDlQikFEyycRYjX52gczT

fd794 := client.Open("/Wuhr1LEWJy.txt", client.O_RDWR|client.O_CREATE)
if(fd794 < 0) {
  panic("open failed")
}


ret = client.Close(fd792)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd794, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (2) qOUGsZ9LGQOsbR3pNwZ8zOuvE0_ptjW5XbICFZA1hoKGkL2EMk4LCe1QE4nNQTq5dPZHnWYHxt88yzCmK3j42Q

ret = client.Write(fd793, []byte("mbotjc3VuBkBHnRWwN5iymTzJoXRRGnJw9gJwRWmTMJ8huwaTorcuDyjAi0pE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) qOmbotjc3VuBkBHnRWwN5iymTzJoXRRGnJw9gJwRWmTMJ8huwaTorcuDyjAi0pE5dPZHnWYHxt88yzCmK3j42Q

fd795 := client.Open("/nqFTzh7bC4.txt", client.O_RDWR|client.O_CREATE)
if(fd795 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd795, []byte("_lLLsWXoFdA3AWQsgR2g4eidtST5RtbkkN5pYBTfZfmXUbc6drfVa9RHKd_nJ6NcuTwKXIqiirK0icju50b5clYAO7XyaQN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) _lLLsWXoFdA3AWQsgR2g4eidtST5RtbkkN5pYBTfZfmXUbc6drfVa9RHKd_nJ6NcuTwKXIqiirK0icju50b5clYAO7XyaQN
//fd state: (63) qOmbotjc3VuBkBHnRWwN5iymTzJoXRRGnJw9gJwRWmTMJ8huwaTorcuDyjAi0pE5dPZHnWYHxt88yzCmK3j42Q

ret = client.Write(fd793, []byte("ssGKzL2WuWwq3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) qOmbotjc3VuBkBHnRWwN5iymTzJoXRRGnJw9gJwRWmTMJ8huwaTorcuDyjAi0pEssGKzL2WuWwq3yzCmK3j42Q
//fd state: (0) 

ret = client.Write(fd794, []byte("KwuqE5pPlr02_fKICSQtEj1_JQJ_Qj5oB1772MEe6snxHv4h4kvgiavGHv46huGsn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) KwuqE5pPlr02_fKICSQtEj1_JQJ_Qj5oB1772MEe6snxHv4h4kvgiavGHv46huGsn

ret = client.Seek(fd793, 84, client.SEEK_SET)
if(ret != 84) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 84)
  panic("seek failed")
}


fd796 := client.Open("/1yGSi4FeUR.txt", client.O_RDWR|client.O_CREATE)
if(fd796 < 0) {
  panic("open failed")
}

//fd state: (95) _lLLsWXoFdA3AWQsgR2g4eidtST5RtbkkN5pYBTfZfmXUbc6drfVa9RHKd_nJ6NcuTwKXIqiirK0icju50b5clYAO7XyaQN

ret = client.Write(fd795, []byte("8BrGxkxP7PRzRJ9GVcNjGw_sUTisg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) _lLLsWXoFdA3AWQsgR2g4eidtST5RtbkkN5pYBTfZfmXUbc6drfVa9RHKd_nJ6NcuTwKXIqiirK0icju50b5clYAO7XyaQN8BrGxkxP7PRzRJ9GVcNjGw_sUTisg
//fd state: (65) KwuqE5pPlr02_fKICSQtEj1_JQJ_Qj5oB1772MEe6snxHv4h4kvgiavGHv46huGsn

ret = client.Write(fd794, []byte("y4bSc2rvuqPx2N2TnrMy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) KwuqE5pPlr02_fKICSQtEj1_JQJ_Qj5oB1772MEe6snxHv4h4kvgiavGHv46huGsny4bSc2rvuqPx2N2TnrMy
//fd state: (85) KwuqE5pPlr02_fKICSQtEj1_JQJ_Qj5oB1772MEe6snxHv4h4kvgiavGHv46huGsny4bSc2rvuqPx2N2TnrMy

ret = client.Write(fd794, []byte("wq0qOoK0eWCRhWjdepFT1cAnYvGEsbWe3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) KwuqE5pPlr02_fKICSQtEj1_JQJ_Qj5oB1772MEe6snxHv4h4kvgiavGHv46huGsny4bSc2rvuqPx2N2TnrMywq0qOoK0eWCRhWjdepFT1cAnYvGEsbWe3

ret = client.Close(fd795)
if(ret != 0) {
  panic("close failed")
}


fd797 := client.Open("/y0dHmoJ4hA.txt", client.O_RDWR|client.O_CREATE)
if(fd797 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd796, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd796)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd793, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2Q") {
  panic("wrong data returned")
}

//fd state: (118) KwuqE5pPlr02_fKICSQtEj1_JQJ_Qj5oB1772MEe6snxHv4h4kvgiavGHv46huGsny4bSc2rvuqPx2N2TnrMywq0qOoK0eWCRhWjdepFT1cAnYvGEsbWe3

ret = client.Write(fd794, []byte("Wq50bQBbekCg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) KwuqE5pPlr02_fKICSQtEj1_JQJ_Qj5oB1772MEe6snxHv4h4kvgiavGHv46huGsny4bSc2rvuqPx2N2TnrMywq0qOoK0eWCRhWjdepFT1cAnYvGEsbWe3Wq50bQBbekCg

buf, ret = client.Read(fd794, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd797, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd793, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd794)
if(ret != 0) {
  panic("close failed")
}


fd798 := client.Open("/wKMHwsSahz.txt", client.O_RDWR|client.O_CREATE)
if(fd798 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd797, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd798)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd797)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd793, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd793, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd793, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd793)
if(ret != 0) {
  panic("close failed")
}


fd799 := client.Open("/HtoO9sC2Y7.txt", client.O_RDWR|client.O_CREATE)
if(fd799 < 0) {
  panic("open failed")
}


fd800 := client.Open("/gQMBiUhjLf.txt", client.O_RDWR|client.O_CREATE)
if(fd800 < 0) {
  panic("open failed")
}


ret = client.Close(fd799)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd800)
if(ret != 0) {
  panic("close failed")
}


fd801 := client.Open("/iL9P8HAhw7.txt", client.O_RDWR|client.O_CREATE)
if(fd801 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd801, []byte("mkbajuq8ClP1D3gvkn0AcG9wEPdGVbVDSMPbszL77K_UdcsfrQc_n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) mkbajuq8ClP1D3gvkn0AcG9wEPdGVbVDSMPbszL77K_UdcsfrQc_n

fd802 := client.Open("/fe6x_kC5oC.txt", client.O_RDWR|client.O_CREATE)
if(fd802 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd801, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd803 := client.Open("/otTHquN_0Q.txt", client.O_RDWR|client.O_CREATE)
if(fd803 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd801, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd802, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}

//fd state: (0) zYyJ6AFWIkkae6fRrdH2_nkGZ3dkzWSTf7U1Nnmw7rTdEIF6uzMmMJooWLpOMcglqxY9OmLtIrvXp1zvdONuXK

ret = client.Write(fd803, []byte("dAVFd34TQNh_zFId0b3zaDF9OFOp_n8bjBvWvCZgXN0eaP9K8L6QyMyseu7vyXROb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) dAVFd34TQNh_zFId0b3zaDF9OFOp_n8bjBvWvCZgXN0eaP9K8L6QyMyseu7vyXRObxY9OmLtIrvXp1zvdONuXK

fd804 := client.Open("/b9IOiiNWX_.txt", client.O_RDWR|client.O_CREATE)
if(fd804 < 0) {
  panic("open failed")
}


fd805 := client.Open("/T6AcK3BvDF.txt", client.O_RDWR|client.O_CREATE)
if(fd805 < 0) {
  panic("open failed")
}


ret = client.Close(fd803)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) I5KoskDjPxTLUGmL4w5CIF6CNsnckHO7aOKNDE1CAO7p3ZhjisODBkwnrlo45_E0QaaofVUINbhvWcIS9EwGUhNN1T6jWS_GadfPWVAXMB_Fa5b47lfgXCJ2bYHGNK10mky3dSscst

ret = client.Write(fd804, []byte("oqlj44ehpr70188mfqJW6n8v2R8Kvil7IeJFNTnVoc1fb1uo43rFgt6FG2uQXw9s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) oqlj44ehpr70188mfqJW6n8v2R8Kvil7IeJFNTnVoc1fb1uo43rFgt6FG2uQXw9sQaaofVUINbhvWcIS9EwGUhNN1T6jWS_GadfPWVAXMB_Fa5b47lfgXCJ2bYHGNK10mky3dSscst

ret = client.Seek(fd801, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


ret = client.Close(fd805)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd802, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xIdsmbXz3vhMXMPyDePR") {
  panic("wrong data returned")
}


ret = client.Seek(fd802, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd804, 87, client.SEEK_SET)
if(ret != 87) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 87)
  panic("seek failed")
}


fd806 := client.Open("/apJFPTK9eS.txt", client.O_RDWR|client.O_CREATE)
if(fd806 < 0) {
  panic("open failed")
}


fd807 := client.Open("/x_st9d12KV.txt", client.O_RDWR|client.O_CREATE)
if(fd807 < 0) {
  panic("open failed")
}


ret = client.Close(fd804)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd806, []byte("bKSFmYUKbdiDW65ELgit8sEdxK9p9mv6VJQF2znUEsTHLr0qQ2Edh4S4vUgXzyB9gd2hCr1CyQMceaJnbLUeg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) bKSFmYUKbdiDW65ELgit8sEdxK9p9mv6VJQF2znUEsTHLr0qQ2Edh4S4vUgXzyB9gd2hCr1CyQMceaJnbLUeg

buf, ret = client.Read(fd807, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd808 := client.Open("/3q5Lyv0cAM.txt", client.O_RDWR|client.O_CREATE)
if(fd808 < 0) {
  panic("open failed")
}


fd809 := client.Open("/offJDM4Vh_.txt", client.O_RDWR|client.O_CREATE)
if(fd809 < 0) {
  panic("open failed")
}


ret = client.Close(fd801)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd806)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd809)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd802, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


fd810 := client.Open("/k06ZPn3SfU.txt", client.O_RDWR|client.O_CREATE)
if(fd810 < 0) {
  panic("open failed")
}


ret = client.Close(fd810)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd808, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd802)
if(ret != 0) {
  panic("close failed")
}


fd811 := client.Open("/1I9MVaXUzk.txt", client.O_RDWR|client.O_CREATE)
if(fd811 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd811, []byte("DAcnkVDmuXBHf3yzwiC3hYz5Hv5_3Dc6jEOTroE4CzIctvJS1ahy0Z9KBvS_jTcxQ9PDeLQALimcKHjlGYAgcGKoR6T0ihBPm_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) DAcnkVDmuXBHf3yzwiC3hYz5Hv5_3Dc6jEOTroE4CzIctvJS1ahy0Z9KBvS_jTcxQ9PDeLQALimcKHjlGYAgcGKoR6T0ihBPm_

fd812 := client.Open("/cK6VMUXJqT.txt", client.O_RDWR|client.O_CREATE)
if(fd812 < 0) {
  panic("open failed")
}


ret = client.Close(fd812)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd808)
if(ret != 0) {
  panic("close failed")
}


fd813 := client.Open("/zv13uXSfcj.txt", client.O_RDWR|client.O_CREATE)
if(fd813 < 0) {
  panic("open failed")
}


fd814 := client.Open("/V56SxLcZii.txt", client.O_RDWR|client.O_CREATE)
if(fd814 < 0) {
  panic("open failed")
}


fd815 := client.Open("/T4EfyJXeyW.txt", client.O_RDWR|client.O_CREATE)
if(fd815 < 0) {
  panic("open failed")
}


ret = client.Close(fd813)
if(ret != 0) {
  panic("close failed")
}

//fd state: (98) DAcnkVDmuXBHf3yzwiC3hYz5Hv5_3Dc6jEOTroE4CzIctvJS1ahy0Z9KBvS_jTcxQ9PDeLQALimcKHjlGYAgcGKoR6T0ihBPm_

ret = client.Write(fd811, []byte("2W9InDGhiRHgIr07zATnZr6v"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) DAcnkVDmuXBHf3yzwiC3hYz5Hv5_3Dc6jEOTroE4CzIctvJS1ahy0Z9KBvS_jTcxQ9PDeLQALimcKHjlGYAgcGKoR6T0ihBPm_2W9InDGhiRHgIr07zATnZr6v
//fd state: (0) GTN2OvkNGa5BycCXrOIvtKbe3xQZqR0F3A1YijsPMdeDG1RKXdTXgTwEor0zEcumpM1REFP1hgvhG8rUwunqJlWIs6j_IxNIbJminsvsyMbckUcRj8YbiId9rBnEmTqYQ15ef5GYGcJQR4x0Q0jbJlGUGk

ret = client.Write(fd815, []byte("9AuAWIAF6Ndedpvacn02e79_RR53qQPCSzITMECKzO4BHMgyzJwQX_qDRrebDtnM__sPFF1We"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) 9AuAWIAF6Ndedpvacn02e79_RR53qQPCSzITMECKzO4BHMgyzJwQX_qDRrebDtnM__sPFF1WegvhG8rUwunqJlWIs6j_IxNIbJminsvsyMbckUcRj8YbiId9rBnEmTqYQ15ef5GYGcJQR4x0Q0jbJlGUGk
//fd state: (73) 9AuAWIAF6Ndedpvacn02e79_RR53qQPCSzITMECKzO4BHMgyzJwQX_qDRrebDtnM__sPFF1WegvhG8rUwunqJlWIs6j_IxNIbJminsvsyMbckUcRj8YbiId9rBnEmTqYQ15ef5GYGcJQR4x0Q0jbJlGUGk

ret = client.Write(fd815, []byte("ylBj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) 9AuAWIAF6Ndedpvacn02e79_RR53qQPCSzITMECKzO4BHMgyzJwQX_qDRrebDtnM__sPFF1WeylBj8rUwunqJlWIs6j_IxNIbJminsvsyMbckUcRj8YbiId9rBnEmTqYQ15ef5GYGcJQR4x0Q0jbJlGUGk

buf, ret = client.Read(fd814, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd807, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd815)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd807)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd814, []byte("AWQI4XHJRDqFAgkr_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) AWQI4XHJRDqFAgkr_

buf, ret = client.Read(fd814, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (122) DAcnkVDmuXBHf3yzwiC3hYz5Hv5_3Dc6jEOTroE4CzIctvJS1ahy0Z9KBvS_jTcxQ9PDeLQALimcKHjlGYAgcGKoR6T0ihBPm_2W9InDGhiRHgIr07zATnZr6v

ret = client.Write(fd811, []byte("0XjbAT7XjgNIbsq78Z7bvYUyHT4c913SvNt54Wt8Kyohr3PRBHsPZx19K_JcRUAvJLpN9W8wYok_s18jEM6VBi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (208) DAcnkVDmuXBHf3yzwiC3hYz5Hv5_3Dc6jEOTroE4CzIctvJS1ahy0Z9KBvS_jTcxQ9PDeLQALimcKHjlGYAgcGKoR6T0ihBPm_2W9InDGhiRHgIr07zATnZr6v0XjbAT7XjgNIbsq78Z7bvYUyHT4c913SvNt54Wt8Kyohr3PRBHsPZx19K_JcRUAvJLpN9W8wYok_s18jEM6VBi

ret = client.Seek(fd811, 139, client.SEEK_SET)
if(ret != 139) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 139)
  panic("seek failed")
}


buf, ret = client.Read(fd814, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd814)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd811, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


buf, ret = client.Read(fd811, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wiC3hYz5Hv5_3Dc6jEOTroE4C") {
  panic("wrong data returned")
}


ret = client.Seek(fd811, 149, client.SEEK_SET)
if(ret != 149) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 149)
  panic("seek failed")
}


buf, ret = client.Read(fd811, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "c913SvNt5") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd811, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4Wt8Kyohr3PRBHsPZx19K_JcRUAvJLpN9W8wY") {
  panic("wrong data returned")
}


ret = client.Close(fd811)
if(ret != 0) {
  panic("close failed")
}


fd816 := client.Open("/cK6VMUXJqT.txt", client.O_RDWR|client.O_CREATE)
if(fd816 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd816, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd817 := client.Open("/SZD6wU8PL1.txt", client.O_RDWR|client.O_CREATE)
if(fd817 < 0) {
  panic("open failed")
}


ret = client.Close(fd817)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd816, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd816, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd816, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd818 := client.Open("/0N4kwgmdsj.txt", client.O_RDWR|client.O_CREATE)
if(fd818 < 0) {
  panic("open failed")
}


ret = client.Close(fd818)
if(ret != 0) {
  panic("close failed")
}


fd819 := client.Open("/3E2MnTHNVS.txt", client.O_RDWR|client.O_CREATE)
if(fd819 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd819, []byte("zYEjrOoYOnw5XvJXYAcURk9qNIvGBVd0G8FIP8A8A7rhXGHQyfxLWAlcXUpCNgHEosoqu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) zYEjrOoYOnw5XvJXYAcURk9qNIvGBVd0G8FIP8A8A7rhXGHQyfxLWAlcXUpCNgHEosoqu

buf, ret = client.Read(fd816, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd816)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd819)
if(ret != 0) {
  panic("close failed")
}


fd820 := client.Open("/1c3zKAUoNJ.txt", client.O_RDWR|client.O_CREATE)
if(fd820 < 0) {
  panic("open failed")
}


fd821 := client.Open("/YYMVsNBZSh.txt", client.O_RDWR|client.O_CREATE)
if(fd821 < 0) {
  panic("open failed")
}


ret = client.Close(fd821)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd820, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd820, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd822 := client.Open("/69ZpRby0jr.txt", client.O_RDWR|client.O_CREATE)
if(fd822 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd822, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd820, []byte("sl43rMbiO1wMJu6pQtB5HVRFJJlu520"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) sl43rMbiO1wMJu6pQtB5HVRFJJlu520
//fd state: (31) sl43rMbiO1wMJu6pQtB5HVRFJJlu520

ret = client.Write(fd820, []byte("qNfB4KrThKRVbWs4MdM_91yoRfRGEXWdoMGjv8csifoSQ4S0aIpsA6ePdr8rRZRJskgrx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) sl43rMbiO1wMJu6pQtB5HVRFJJlu520qNfB4KrThKRVbWs4MdM_91yoRfRGEXWdoMGjv8csifoSQ4S0aIpsA6ePdr8rRZRJskgrx

ret = client.Seek(fd820, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


buf, ret = client.Read(fd822, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd823 := client.Open("/bFM001bg35.txt", client.O_RDWR|client.O_CREATE)
if(fd823 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd823, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lO2uy1HZ9VsGUlen1r04sbnVnh46Gb") {
  panic("wrong data returned")
}


ret = client.Seek(fd820, 84, client.SEEK_SET)
if(ret != 84) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 84)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd822, []byte("Eiz_q4gM2uU1oBs7X7Fspz7EzlSKVhSpm_tm3nlKeTTImpcWd2XuacHVv8QLGGZvmU6khlF2KoEZi9JFcKpNf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) Eiz_q4gM2uU1oBs7X7Fspz7EzlSKVhSpm_tm3nlKeTTImpcWd2XuacHVv8QLGGZvmU6khlF2KoEZi9JFcKpNf
//fd state: (30) lO2uy1HZ9VsGUlen1r04sbnVnh46GbxKYq4N7JcygGFV

ret = client.Write(fd823, []byte("6EavUro24q2moLap7DA9FloOxX_6IgBXASlw9nUevmlGO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) lO2uy1HZ9VsGUlen1r04sbnVnh46Gb6EavUro24q2moLap7DA9FloOxX_6IgBXASlw9nUevmlGO

fd824 := client.Open("/1HmdFpbrnZ.txt", client.O_RDWR|client.O_CREATE)
if(fd824 < 0) {
  panic("open failed")
}


ret = client.Close(fd823)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd824)
if(ret != 0) {
  panic("close failed")
}


fd825 := client.Open("/BuCA1_kEx5.txt", client.O_RDWR|client.O_CREATE)
if(fd825 < 0) {
  panic("open failed")
}


ret = client.Close(fd820)
if(ret != 0) {
  panic("close failed")
}


fd826 := client.Open("/QlSfrOOHqy.txt", client.O_RDWR|client.O_CREATE)
if(fd826 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd825, []byte("mnxHj0bIEjDPQMU76rdc2K8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) mnxHj0bIEjDPQMU76rdc2K8
//fd state: (23) mnxHj0bIEjDPQMU76rdc2K8

ret = client.Write(fd825, []byte("32q1F_fydJcpox"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) mnxHj0bIEjDPQMU76rdc2K832q1F_fydJcpox

buf, ret = client.Read(fd825, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd825, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd826, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Close(fd825)
if(ret != 0) {
  panic("close failed")
}

//fd state: (85) Eiz_q4gM2uU1oBs7X7Fspz7EzlSKVhSpm_tm3nlKeTTImpcWd2XuacHVv8QLGGZvmU6khlF2KoEZi9JFcKpNf

ret = client.Write(fd822, []byte("42EznF7eznE2vt6vSlH268jzmD2LAHyUwmFi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (121) Eiz_q4gM2uU1oBs7X7Fspz7EzlSKVhSpm_tm3nlKeTTImpcWd2XuacHVv8QLGGZvmU6khlF2KoEZi9JFcKpNf42EznF7eznE2vt6vSlH268jzmD2LAHyUwmFi

fd827 := client.Open("/iY0nrsRuhc.txt", client.O_RDWR|client.O_CREATE)
if(fd827 < 0) {
  panic("open failed")
}


ret = client.Seek(fd822, 117, client.SEEK_SET)
if(ret != 117) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 117)
  panic("seek failed")
}


buf, ret = client.Read(fd827, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (35) nMfUfmOTLfsKpfUip7dQNNPL2BkQMe0pWgUpGHfv5CZx2tbK_Pj6MwXZeZMsNoy0PAMaw0TCQ

ret = client.Write(fd826, []byte("v"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) nMfUfmOTLfsKpfUip7dQNNPL2BkQMe0pWgUvGHfv5CZx2tbK_Pj6MwXZeZMsNoy0PAMaw0TCQ

ret = client.Close(fd826)
if(ret != 0) {
  panic("close failed")
}

//fd state: (117) Eiz_q4gM2uU1oBs7X7Fspz7EzlSKVhSpm_tm3nlKeTTImpcWd2XuacHVv8QLGGZvmU6khlF2KoEZi9JFcKpNf42EznF7eznE2vt6vSlH268jzmD2LAHyUwmFi

ret = client.Write(fd822, []byte("jAi2xKX0zkAnaOtcZCyfvJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (139) Eiz_q4gM2uU1oBs7X7Fspz7EzlSKVhSpm_tm3nlKeTTImpcWd2XuacHVv8QLGGZvmU6khlF2KoEZi9JFcKpNf42EznF7eznE2vt6vSlH268jzmD2LAHyUjAi2xKX0zkAnaOtcZCyfvJ

ret = client.Seek(fd827, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd827)
if(ret != 0) {
  panic("close failed")
}


fd828 := client.Open("/0RJX1d5gRl.txt", client.O_RDWR|client.O_CREATE)
if(fd828 < 0) {
  panic("open failed")
}

//fd state: (139) Eiz_q4gM2uU1oBs7X7Fspz7EzlSKVhSpm_tm3nlKeTTImpcWd2XuacHVv8QLGGZvmU6khlF2KoEZi9JFcKpNf42EznF7eznE2vt6vSlH268jzmD2LAHyUjAi2xKX0zkAnaOtcZCyfvJ

ret = client.Write(fd822, []byte("KMOQ9FMN6kGVKIs6ESv9u1ozznsY1aQ25uN7I0MsSY9kNW8grdA4ZHzxRsm0CJixcY8WDqrHRqYKuUj6qv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (221) Eiz_q4gM2uU1oBs7X7Fspz7EzlSKVhSpm_tm3nlKeTTImpcWd2XuacHVv8QLGGZvmU6khlF2KoEZi9JFcKpNf42EznF7eznE2vt6vSlH268jzmD2LAHyUjAi2xKX0zkAnaOtcZCyfvJKMOQ9FMN6kGVKIs6ESv9u1ozznsY1aQ25uN7I0MsSY9kNW8grdA4ZHzxRsm0CJixcY8WDqrHRqYKuUj6qv

fd829 := client.Open("/_9fBycs2Go.txt", client.O_RDWR|client.O_CREATE)
if(fd829 < 0) {
  panic("open failed")
}

//fd state: (221) Eiz_q4gM2uU1oBs7X7Fspz7EzlSKVhSpm_tm3nlKeTTImpcWd2XuacHVv8QLGGZvmU6khlF2KoEZi9JFcKpNf42EznF7eznE2vt6vSlH268jzmD2LAHyUjAi2xKX0zkAnaOtcZCyfvJKMOQ9FMN6kGVKIs6ESv9u1ozznsY1aQ25uN7I0MsSY9kNW8grdA4ZHzxRsm0CJixcY8WDqrHRqYKuUj6qv

ret = client.Write(fd822, []byte("NVW67FmM08ARDnRoCZkgAFSEMDJYL2M01nYJA24BUXozNR4GzFyoAMt6g19krmTVbtTZX9WanrDXHIqcB4s1uG_XI3pQo4Bw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (317) Eiz_q4gM2uU1oBs7X7Fspz7EzlSKVhSpm_tm3nlKeTTImpcWd2XuacHVv8QLGGZvmU6khlF2KoEZi9JFcKpNf42EznF7eznE2vt6vSlH268jzmD2LAHyUjAi2xKX0zkAnaOtcZCyfvJKMOQ9FMN6kGVKIs6ESv9u1ozznsY1aQ25uN7I0MsSY9kNW8grdA4ZHzxRsm0CJixcY8WDqrHRqYKuUj6qvNVW67FmM08ARDnRoCZkgAFSEMDJYL2M01nYJA24BUXozNR4GzFyoAMt6g19krmTVbtTZX9WanrDXHIqcB4s1uG_XI3pQo4Bw

ret = client.Close(fd822)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd829, []byte("iGSdaPc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) iGSdaPc

buf, ret = client.Read(fd829, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd828, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (7) iGSdaPc

ret = client.Write(fd829, []byte("EpoU_LohfjC_td2GXxgy18A_w7ugbr0C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) iGSdaPcEpoU_LohfjC_td2GXxgy18A_w7ugbr0C

fd830 := client.Open("/sAKrpzKuzA.txt", client.O_RDWR|client.O_CREATE)
if(fd830 < 0) {
  panic("open failed")
}


fd831 := client.Open("/Q8GlO2l7Nx.txt", client.O_RDWR|client.O_CREATE)
if(fd831 < 0) {
  panic("open failed")
}


fd832 := client.Open("/wCijCfQcBC.txt", client.O_RDWR|client.O_CREATE)
if(fd832 < 0) {
  panic("open failed")
}


ret = client.Seek(fd829, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


ret = client.Seek(fd828, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (22) iGSdaPcEpoU_LohfjC_td2GXxgy18A_w7ugbr0C

ret = client.Write(fd829, []byte("MWcDEJMgquejVCEfdbkj1Qaxtb6J"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) iGSdaPcEpoU_LohfjC_td2MWcDEJMgquejVCEfdbkj1Qaxtb6J

ret = client.Close(fd829)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd831, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd833 := client.Open("/jAMfPtHSb4.txt", client.O_RDWR|client.O_CREATE)
if(fd833 < 0) {
  panic("open failed")
}


fd834 := client.Open("/_xt2CaFERA.txt", client.O_RDWR|client.O_CREATE)
if(fd834 < 0) {
  panic("open failed")
}


fd835 := client.Open("/di6atYCfA6.txt", client.O_RDWR|client.O_CREATE)
if(fd835 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd828, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd828, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd832, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd830)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd835)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd836 := client.Open("/hUsfMv9kV4.txt", client.O_RDWR|client.O_CREATE)
if(fd836 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd834, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qOmbotjc3VuBkBHnRWwN5iymTzJoXRRGnJw9gJwRWmTMJ8huwaTorcuDyjAi0pEssGKzL2WuW") {
  panic("wrong data returned")
}


ret = client.Seek(fd833, 68, client.SEEK_SET)
if(ret != 68) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 68)
  panic("seek failed")
}

//fd state: (68) utxBQzB5n_lRdnUOo3nT6lWYcW82XY2CK7nllcuPt9thV9oLPQ9JGZ6WIF_RjukJVrAMX9D9wqe3LqZO

ret = client.Write(fd833, []byte("XfHuw0sLMLAyvicT7f2KPbxHFul7pUAoqSECRCQpOlV9BhJiBBHZIXtxLOM213gWr07ZrgKCE0K4n_kiGKvqeIw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (155) utxBQzB5n_lRdnUOo3nT6lWYcW82XY2CK7nllcuPt9thV9oLPQ9JGZ6WIF_RjukJVrAMXfHuw0sLMLAyvicT7f2KPbxHFul7pUAoqSECRCQpOlV9BhJiBBHZIXtxLOM213gWr07ZrgKCE0K4n_kiGKvqeIw

ret = client.Close(fd832)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd831, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd831, []byte("jp6TWNAa6GsKsrBH5fxI7mpJEsQCjri9IJVJELfC9NE7_Ta4xqpCjij7yAbpBCqilRaaZhAm_Gj9zO7h2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) jp6TWNAa6GsKsrBH5fxI7mpJEsQCjri9IJVJELfC9NE7_Ta4xqpCjij7yAbpBCqilRaaZhAm_Gj9zO7h2

fd837 := client.Open("/1g_SO9gjIH.txt", client.O_RDWR|client.O_CREATE)
if(fd837 < 0) {
  panic("open failed")
}


fd838 := client.Open("/DRtZWfLYIa.txt", client.O_RDWR|client.O_CREATE)
if(fd838 < 0) {
  panic("open failed")
}

//fd state: (73) qOmbotjc3VuBkBHnRWwN5iymTzJoXRRGnJw9gJwRWmTMJ8huwaTorcuDyjAi0pEssGKzL2WuWwq3yzCmK3j42Q

ret = client.Write(fd834, []byte("cg6XJdZYUIuZRwq1Yng0p80X1M2ITRM9awx_nfwi6M5oFzWgKtD1KKlWKZwwMVtAzmFAp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (142) qOmbotjc3VuBkBHnRWwN5iymTzJoXRRGnJw9gJwRWmTMJ8huwaTorcuDyjAi0pEssGKzL2WuWcg6XJdZYUIuZRwq1Yng0p80X1M2ITRM9awx_nfwi6M5oFzWgKtD1KKlWKZwwMVtAzmFAp

fd839 := client.Open("/fhDLD2y7yr.txt", client.O_RDWR|client.O_CREATE)
if(fd839 < 0) {
  panic("open failed")
}


fd840 := client.Open("/CX0xtnQDqz.txt", client.O_RDWR|client.O_CREATE)
if(fd840 < 0) {
  panic("open failed")
}


fd841 := client.Open("/Uq03S3GN57.txt", client.O_RDWR|client.O_CREATE)
if(fd841 < 0) {
  panic("open failed")
}


ret = client.Close(fd840)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd828, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd836, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd837, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd833, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


ret = client.Close(fd838)
if(ret != 0) {
  panic("close failed")
}


fd842 := client.Open("/kqZrpEP1nq.txt", client.O_RDWR|client.O_CREATE)
if(fd842 < 0) {
  panic("open failed")
}


fd843 := client.Open("/xpQvmuTyo2.txt", client.O_RDWR|client.O_CREATE)
if(fd843 < 0) {
  panic("open failed")
}

//fd state: (0) 7z0iqAkzNjxwUoP63900yAZEcgi9CePJauaVb6pC8PQrKhi5HLcTpHw7I

ret = client.Write(fd841, []byte("FTpeZdzAAsoo8NJRZWKyueSly9CVj_s2yKInTLwXOgcdFmAKrdEvEh2DKZZzNtmLpQrCyZ5yUx1fKdjw4AUwCOpuB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) FTpeZdzAAsoo8NJRZWKyueSly9CVj_s2yKInTLwXOgcdFmAKrdEvEh2DKZZzNtmLpQrCyZ5yUx1fKdjw4AUwCOpuB
//fd state: (0) 

ret = client.Write(fd837, []byte("3xfgHWYhKZN8kCNWm9s2ik9IYr0qssI2ywWmqPpCgWVBwN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) 3xfgHWYhKZN8kCNWm9s2ik9IYr0qssI2ywWmqPpCgWVBwN
//fd state: (46) 3xfgHWYhKZN8kCNWm9s2ik9IYr0qssI2ywWmqPpCgWVBwN

ret = client.Write(fd837, []byte("ofilIiYGPcURttQ6KdVtHQMXUjZu2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) 3xfgHWYhKZN8kCNWm9s2ik9IYr0qssI2ywWmqPpCgWVBwNofilIiYGPcURttQ6KdVtHQMXUjZu2

buf, ret = client.Read(fd843, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "majyE2L13bVVck0nT8gY_InwWKbWtbZeyyW1Nrs3J6LJvR2CjYtPS1xM") {
  panic("wrong data returned")
}

//fd state: (81) jp6TWNAa6GsKsrBH5fxI7mpJEsQCjri9IJVJELfC9NE7_Ta4xqpCjij7yAbpBCqilRaaZhAm_Gj9zO7h2

ret = client.Write(fd831, []byte("LxHf3BBCDdZN7YIF1mHEjnQ41Nz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (108) jp6TWNAa6GsKsrBH5fxI7mpJEsQCjri9IJVJELfC9NE7_Ta4xqpCjij7yAbpBCqilRaaZhAm_Gj9zO7h2LxHf3BBCDdZN7YIF1mHEjnQ41Nz

buf, ret = client.Read(fd842, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd828, []byte("9sRPnzwVrcYypeV6VVBDwM5Tj2rmgoR6UWSFC8ocomP1aWUtjgRz5H4G4TxJAR7OVAxGvB9qkX6gzX4xhlR89vq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) 9sRPnzwVrcYypeV6VVBDwM5Tj2rmgoR6UWSFC8ocomP1aWUtjgRz5H4G4TxJAR7OVAxGvB9qkX6gzX4xhlR89vq
//fd state: (108) jp6TWNAa6GsKsrBH5fxI7mpJEsQCjri9IJVJELfC9NE7_Ta4xqpCjij7yAbpBCqilRaaZhAm_Gj9zO7h2LxHf3BBCDdZN7YIF1mHEjnQ41Nz

ret = client.Write(fd831, []byte("cbIZSfqJQqTtzeWgi7FONsiPiltRYOqXVMZj55jTZRY7Z4T_5JNlfxPp9cIGa0I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (171) jp6TWNAa6GsKsrBH5fxI7mpJEsQCjri9IJVJELfC9NE7_Ta4xqpCjij7yAbpBCqilRaaZhAm_Gj9zO7h2LxHf3BBCDdZN7YIF1mHEjnQ41NzcbIZSfqJQqTtzeWgi7FONsiPiltRYOqXVMZj55jTZRY7Z4T_5JNlfxPp9cIGa0I

ret = client.Close(fd841)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd833)
if(ret != 0) {
  panic("close failed")
}

//fd state: (56) majyE2L13bVVck0nT8gY_InwWKbWtbZeyyW1Nrs3J6LJvR2CjYtPS1xM5qEbnM9RZ20QyeDeD6y07aYELz72d55qe_3W

ret = client.Write(fd843, []byte("h48QhDeApcW8owGAv2fjb_9cvPnpgqlidj4DkGxST15hzvGK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) majyE2L13bVVck0nT8gY_InwWKbWtbZeyyW1Nrs3J6LJvR2CjYtPS1xMh48QhDeApcW8owGAv2fjb_9cvPnpgqlidj4DkGxST15hzvGK

ret = client.Seek(fd831, 152, client.SEEK_SET)
if(ret != 152) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 152)
  panic("seek failed")
}

//fd state: (0) qrfa5gWgDlSO4RWqHjwD4u_LKymwLsRXoHpKxr2IQuNFafxBG9_ZVFRPxTt6yNA8BH

ret = client.Write(fd839, []byte("sYZjP_4swgdv1LqnLUndvkeM2cGiOFRTZTt2fxG13MHon2j6zStncrSBuj6YVk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) sYZjP_4swgdv1LqnLUndvkeM2cGiOFRTZTt2fxG13MHon2j6zStncrSBuj6YVkA8BH

ret = client.Close(fd836)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd828, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd839, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Close(fd837)
if(ret != 0) {
  panic("close failed")
}


fd844 := client.Open("/hFOV2oTMMC.txt", client.O_RDWR|client.O_CREATE)
if(fd844 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd844, []byte("l9DktMw668a8BrAOo3oKt20mVoPOdwB0fDu1dLx3mVYTmkxFKvAHY0QHn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) l9DktMw668a8BrAOo3oKt20mVoPOdwB0fDu1dLx3mVYTmkxFKvAHY0QHn

fd845 := client.Open("/V0gWwEDbM4.txt", client.O_RDWR|client.O_CREATE)
if(fd845 < 0) {
  panic("open failed")
}


ret = client.Seek(fd828, 69, client.SEEK_SET)
if(ret != 69) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 69)
  panic("seek failed")
}


ret = client.Close(fd843)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd839)
if(ret != 0) {
  panic("close failed")
}


fd846 := client.Open("/JFMbWQchvB.txt", client.O_RDWR|client.O_CREATE)
if(fd846 < 0) {
  panic("open failed")
}


ret = client.Close(fd845)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd834)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd842, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd831, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


buf, ret = client.Read(fd828, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "B9qkX6gzX4xhlR89vq") {
  panic("wrong data returned")
}


ret = client.Close(fd842)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd844, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


ret = client.Seek(fd844, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


ret = client.Seek(fd831, 159, client.SEEK_SET)
if(ret != 159) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 159)
  panic("seek failed")
}


ret = client.Close(fd831)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd844, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Kt20mVoPOdwB0fDu1dLx3mVYTmkxFKvAHY0QHn") {
  panic("wrong data returned")
}


ret = client.Close(fd846)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd844, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


fd847 := client.Open("/fhDLD2y7yr.txt", client.O_RDWR|client.O_CREATE)
if(fd847 < 0) {
  panic("open failed")
}

//fd state: (87) 9sRPnzwVrcYypeV6VVBDwM5Tj2rmgoR6UWSFC8ocomP1aWUtjgRz5H4G4TxJAR7OVAxGvB9qkX6gzX4xhlR89vq

ret = client.Write(fd828, []byte("RCRa9TZp_N1JvvnpxGhKP72wtEWoZMXifzwdr1NlN0iFH5HqgQ1zxBkwi_ow3yUSULnRGlTtv8pOQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) 9sRPnzwVrcYypeV6VVBDwM5Tj2rmgoR6UWSFC8ocomP1aWUtjgRz5H4G4TxJAR7OVAxGvB9qkX6gzX4xhlR89vqRCRa9TZp_N1JvvnpxGhKP72wtEWoZMXifzwdr1NlN0iFH5HqgQ1zxBkwi_ow3yUSULnRGlTtv8pOQ

ret = client.Seek(fd828, 100, client.SEEK_SET)
if(ret != 100) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 100)
  panic("seek failed")
}

//fd state: (100) 9sRPnzwVrcYypeV6VVBDwM5Tj2rmgoR6UWSFC8ocomP1aWUtjgRz5H4G4TxJAR7OVAxGvB9qkX6gzX4xhlR89vqRCRa9TZp_N1JvvnpxGhKP72wtEWoZMXifzwdr1NlN0iFH5HqgQ1zxBkwi_ow3yUSULnRGlTtv8pOQ

ret = client.Write(fd828, []byte("tj9iPSE4HRF1SiELSKApg9jA2_IGzBe_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) 9sRPnzwVrcYypeV6VVBDwM5Tj2rmgoR6UWSFC8ocomP1aWUtjgRz5H4G4TxJAR7OVAxGvB9qkX6gzX4xhlR89vqRCRa9TZp_N1Jvtj9iPSE4HRF1SiELSKApg9jA2_IGzBe_5HqgQ1zxBkwi_ow3yUSULnRGlTtv8pOQ

buf, ret = client.Read(fd844, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1dLx3mVYTmkxFKvAHY0QHn") {
  panic("wrong data returned")
}


ret = client.Close(fd844)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) sYZjP_4swgdv1LqnLUndvkeM2cGiOFRTZTt2fxG13MHon2j6zStncrSBuj6YVkA8BH

ret = client.Write(fd847, []byte("3vCXcKK2wuwsLwO2ozBgZFrCniqaj2ofyY0CJKB4kJe9WcUcVb684L5_xMqFo5zT6ZHUas33ag91nMGKsLKk5lhWTTn_f"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) 3vCXcKK2wuwsLwO2ozBgZFrCniqaj2ofyY0CJKB4kJe9WcUcVb684L5_xMqFo5zT6ZHUas33ag91nMGKsLKk5lhWTTn_f

fd848 := client.Open("/LG8n_Q468c.txt", client.O_RDWR|client.O_CREATE)
if(fd848 < 0) {
  panic("open failed")
}


ret = client.Close(fd848)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd828, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5HqgQ1zxBkwi_ow3yUSULnRGlTtv8pOQ") {
  panic("wrong data returned")
}


fd849 := client.Open("/kcHPGxVqjP.txt", client.O_RDWR|client.O_CREATE)
if(fd849 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd849, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vek5yBnmqCyeowwng1hCtBFC7WpEALfVVyV9qrdXdDdVv9TsXsgfbDhE4nzm") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd828, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd850 := client.Open("/ad457d9kC_.txt", client.O_RDWR|client.O_CREATE)
if(fd850 < 0) {
  panic("open failed")
}


ret = client.Close(fd847)
if(ret != 0) {
  panic("close failed")
}


fd851 := client.Open("/1_pU8JJR0U.txt", client.O_RDWR|client.O_CREATE)
if(fd851 < 0) {
  panic("open failed")
}


ret = client.Close(fd828)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) iJpQ1HjPqsrTCIFX

ret = client.Write(fd851, []byte("7ULH0cbQy6EVebUDg8wbbtnOdy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) 7ULH0cbQy6EVebUDg8wbbtnOdy

fd852 := client.Open("/frZMlkWF_t.txt", client.O_RDWR|client.O_CREATE)
if(fd852 < 0) {
  panic("open failed")
}


fd853 := client.Open("/SIYaGzyyUN.txt", client.O_RDWR|client.O_CREATE)
if(fd853 < 0) {
  panic("open failed")
}


fd854 := client.Open("/AatP41mIdG.txt", client.O_RDWR|client.O_CREATE)
if(fd854 < 0) {
  panic("open failed")
}


fd855 := client.Open("/62kKbTRIzd.txt", client.O_RDWR|client.O_CREATE)
if(fd855 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd853, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "DfKeJmPN1B7GooOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8K") {
  panic("wrong data returned")
}


ret = client.Close(fd850)
if(ret != 0) {
  panic("close failed")
}

//fd state: (60) vek5yBnmqCyeowwng1hCtBFC7WpEALfVVyV9qrdXdDdVv9TsXsgfbDhE4nzm

ret = client.Write(fd849, []byte("bi7mspi1j2XHnCJO67FbITXRVyGPTBosVe2BsxXtvBEVAjBvMnuhamDWKKgRawz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) vek5yBnmqCyeowwng1hCtBFC7WpEALfVVyV9qrdXdDdVv9TsXsgfbDhE4nzmbi7mspi1j2XHnCJO67FbITXRVyGPTBosVe2BsxXtvBEVAjBvMnuhamDWKKgRawz

ret = client.Close(fd849)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd851)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd852, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) pIdaBHcH35nXzQpPgQg6hOT3

ret = client.Write(fd855, []byte("wbLc1zhuukPHor5esoUincVEQQnE__QjnKyXU9PyhmMSw9cEs1P6vU_QpuQfAP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) wbLc1zhuukPHor5esoUincVEQQnE__QjnKyXU9PyhmMSw9cEs1P6vU_QpuQfAP

ret = client.Close(fd853)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd854, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (62) wbLc1zhuukPHor5esoUincVEQQnE__QjnKyXU9PyhmMSw9cEs1P6vU_QpuQfAP

ret = client.Write(fd855, []byte("RttFjGmuRqhaXs_tw58cwgerCDdExXmDtoBjgqICgsG7xFNeazjidcTk2VD82KW7T5a8r4k_H3MalUlqGr1RGcaG0xRVv4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (156) wbLc1zhuukPHor5esoUincVEQQnE__QjnKyXU9PyhmMSw9cEs1P6vU_QpuQfAPRttFjGmuRqhaXs_tw58cwgerCDdExXmDtoBjgqICgsG7xFNeazjidcTk2VD82KW7T5a8r4k_H3MalUlqGr1RGcaG0xRVv4

ret = client.Seek(fd852, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd854)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd852, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd852, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd855)
if(ret != 0) {
  panic("close failed")
}


fd856 := client.Open("/knLVMcM41S.txt", client.O_RDWR|client.O_CREATE)
if(fd856 < 0) {
  panic("open failed")
}


ret = client.Close(fd852)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd856, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (2) Gj8dzMbIwB5LHgYjP

ret = client.Write(fd856, []byte("wi5aEcipI7XqdW8E3FaaSZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) Gjwi5aEcipI7XqdW8E3FaaSZ
//fd state: (24) Gjwi5aEcipI7XqdW8E3FaaSZ

ret = client.Write(fd856, []byte("6hjo18H2E_tI8YtcjKpHFnfwFSxmhUs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) Gjwi5aEcipI7XqdW8E3FaaSZ6hjo18H2E_tI8YtcjKpHFnfwFSxmhUs
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd856)
if(ret != 0) {
  panic("close failed")
}


fd857 := client.Open("/miLvheV4je.txt", client.O_RDWR|client.O_CREATE)
if(fd857 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd857, []byte("hA1SXtuyf_7fYRPBijsoeuN68frOaYd4EfuZvVBm3hz_uKAFxN3TS_7l_DoI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) hA1SXtuyf_7fYRPBijsoeuN68frOaYd4EfuZvVBm3hz_uKAFxN3TS_7l_DoI

ret = client.Seek(fd857, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


fd858 := client.Open("/eaI_co6DSS.txt", client.O_RDWR|client.O_CREATE)
if(fd858 < 0) {
  panic("open failed")
}


fd859 := client.Open("/DOVDokaeNv.txt", client.O_RDWR|client.O_CREATE)
if(fd859 < 0) {
  panic("open failed")
}


fd860 := client.Open("/OKVWtGYx1I.txt", client.O_RDWR|client.O_CREATE)
if(fd860 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd857, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_7fYRPBijsoeuN68frOaYd4EfuZvVBm3hz_uKAFxN3TS_7l_DoI") {
  panic("wrong data returned")
}


ret = client.Close(fd859)
if(ret != 0) {
  panic("close failed")
}


fd861 := client.Open("/YVbJEFJUJd.txt", client.O_RDWR|client.O_CREATE)
if(fd861 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd861, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "S_KajJMZfBB0gddHDL5") {
  panic("wrong data returned")
}

//fd state: (19) S_KajJMZfBB0gddHDL55vW9C7TqbfYOiTJA3LC8QiRuuNEUOICa

ret = client.Write(fd861, []byte("jeMRnQT64AXKZWgD1tohSWSw3E"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) S_KajJMZfBB0gddHDL5jeMRnQT64AXKZWgD1tohSWSw3EEUOICa

buf, ret = client.Read(fd861, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "EUOICa") {
  panic("wrong data returned")
}


ret = client.Close(fd857)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd860, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (51) S_KajJMZfBB0gddHDL5jeMRnQT64AXKZWgD1tohSWSw3EEUOICa

ret = client.Write(fd861, []byte("2CNMHprpbWrS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) S_KajJMZfBB0gddHDL5jeMRnQT64AXKZWgD1tohSWSw3EEUOICa2CNMHprpbWrS

ret = client.Seek(fd858, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd861, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd860, []byte("DnNQBqJREfy8AdbheINxEr92AREPYczembJPkBXDuoqLdExOTW6vR2BIGup0cYuLaQG__Ke39SZxQrPAO8FUKfhEH1U"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) DnNQBqJREfy8AdbheINxEr92AREPYczembJPkBXDuoqLdExOTW6vR2BIGup0cYuLaQG__Ke39SZxQrPAO8FUKfhEH1U

ret = client.Seek(fd858, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (63) S_KajJMZfBB0gddHDL5jeMRnQT64AXKZWgD1tohSWSw3EEUOICa2CNMHprpbWrS

ret = client.Write(fd861, []byte("kQsw3j2jKHTDooWtjHQE5vUDIRqvyVxZW3b5mAikjjvqu3sZO9NkUya2CUIoVOciPM5vJdQLvDtNYYYtQG8FMn9dHYejbi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (157) S_KajJMZfBB0gddHDL5jeMRnQT64AXKZWgD1tohSWSw3EEUOICa2CNMHprpbWrSkQsw3j2jKHTDooWtjHQE5vUDIRqvyVxZW3b5mAikjjvqu3sZO9NkUya2CUIoVOciPM5vJdQLvDtNYYYtQG8FMn9dHYejbi

buf, ret = client.Read(fd860, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd862 := client.Open("/gfY_Y9jHYW.txt", client.O_RDWR|client.O_CREATE)
if(fd862 < 0) {
  panic("open failed")
}


ret = client.Seek(fd858, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd863 := client.Open("/DDLIkxI9an.txt", client.O_RDWR|client.O_CREATE)
if(fd863 < 0) {
  panic("open failed")
}


fd864 := client.Open("/SNTEj6ymnz.txt", client.O_RDWR|client.O_CREATE)
if(fd864 < 0) {
  panic("open failed")
}


ret = client.Close(fd863)
if(ret != 0) {
  panic("close failed")
}


fd865 := client.Open("/VmU5r1Iala.txt", client.O_RDWR|client.O_CREATE)
if(fd865 < 0) {
  panic("open failed")
}


fd866 := client.Open("/VghMsYjnXc.txt", client.O_RDWR|client.O_CREATE)
if(fd866 < 0) {
  panic("open failed")
}


fd867 := client.Open("/zR6Zy4PV6d.txt", client.O_RDWR|client.O_CREATE)
if(fd867 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd864, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd867, []byte("HW9Jcs40flXFs1ykmA3wc34By5wXyhodaChiM9MuH4HoXNIFP9GOOXGb_pjz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) HW9Jcs40flXFs1ykmA3wc34By5wXyhodaChiM9MuH4HoXNIFP9GOOXGb_pjz

buf, ret = client.Read(fd860, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd861, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd864, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd860, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd864, []byte("H3MaYrnM11HXhCHi17owW8Jr94zSgJtq6vErHR56QMFdI7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) H3MaYrnM11HXhCHi17owW8Jr94zSgJtq6vErHR56QMFdI7

ret = client.Close(fd865)
if(ret != 0) {
  panic("close failed")
}

//fd state: (60) HW9Jcs40flXFs1ykmA3wc34By5wXyhodaChiM9MuH4HoXNIFP9GOOXGb_pjz

ret = client.Write(fd867, []byte("5g3fSsFUYESBQ2sHMJNIs3018SiKHfrxhnVXXM8dVJx1U467E"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (109) HW9Jcs40flXFs1ykmA3wc34By5wXyhodaChiM9MuH4HoXNIFP9GOOXGb_pjz5g3fSsFUYESBQ2sHMJNIs3018SiKHfrxhnVXXM8dVJx1U467E

ret = client.Seek(fd861, 137, client.SEEK_SET)
if(ret != 137) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 137)
  panic("seek failed")
}


ret = client.Close(fd866)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd858)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd864, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


ret = client.Seek(fd860, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


buf, ret = client.Read(fd867, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd861)
if(ret != 0) {
  panic("close failed")
}


fd868 := client.Open("/K1YhjaImgW.txt", client.O_RDWR|client.O_CREATE)
if(fd868 < 0) {
  panic("open failed")
}


fd869 := client.Open("/U27uigcAVw.txt", client.O_RDWR|client.O_CREATE)
if(fd869 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd860, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xQrPAO8FUKfhEH1U") {
  panic("wrong data returned")
}


ret = client.Seek(fd867, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


ret = client.Close(fd869)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd864, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}

//fd state: (91) DnNQBqJREfy8AdbheINxEr92AREPYczembJPkBXDuoqLdExOTW6vR2BIGup0cYuLaQG__Ke39SZxQrPAO8FUKfhEH1U

ret = client.Write(fd860, []byte("SrMb1kp7Hm38uj66xDMpW_PKieLPrD_4ArEuhhRxPFy4jHr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (138) DnNQBqJREfy8AdbheINxEr92AREPYczembJPkBXDuoqLdExOTW6vR2BIGup0cYuLaQG__Ke39SZxQrPAO8FUKfhEH1USrMb1kp7Hm38uj66xDMpW_PKieLPrD_4ArEuhhRxPFy4jHr
//fd state: (29) H3MaYrnM11HXhCHi17owW8Jr94zSgJtq6vErHR56QMFdI7

ret = client.Write(fd864, []byte("KAq9Jp2se37SXddjz9vssYELgbRnNVHT0h6VE4P6X4gw09r7zvgKVSE5yX9NbgR4jhGT5_GKEWNE0E9pAT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) H3MaYrnM11HXhCHi17owW8Jr94zSgKAq9Jp2se37SXddjz9vssYELgbRnNVHT0h6VE4P6X4gw09r7zvgKVSE5yX9NbgR4jhGT5_GKEWNE0E9pAT

buf, ret = client.Read(fd867, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "HoXNIFP9GOOXGb_pjz5g3fSs") {
  panic("wrong data returned")
}


ret = client.Close(fd860)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd868, 82, client.SEEK_SET)
if(ret != 82) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 82)
  panic("seek failed")
}


ret = client.Seek(fd862, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


fd870 := client.Open("/JKZTf7z8OC.txt", client.O_RDWR|client.O_CREATE)
if(fd870 < 0) {
  panic("open failed")
}

//fd state: (5) 1qU_wlfuJSgy1tUgBnhr6RSWnRn8_cnAX

ret = client.Write(fd862, []byte("nzcVKQZRFA6uV61UrbMzJoWrubrW4M5u361pAOEYGAu4JP4DCtyASmvkfgRBrGEujXZ2ans07e"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) 1qU_wnzcVKQZRFA6uV61UrbMzJoWrubrW4M5u361pAOEYGAu4JP4DCtyASmvkfgRBrGEujXZ2ans07e
//fd state: (79) 1qU_wnzcVKQZRFA6uV61UrbMzJoWrubrW4M5u361pAOEYGAu4JP4DCtyASmvkfgRBrGEujXZ2ans07e

ret = client.Write(fd862, []byte("zDQxURfkTbi5tNadJc_wDwk7W7q51lnLalasno2UsgF6EpKdDc5k4muUfFi6AwDdJhsuqDhnMCoap0XkNPYzRbwhxZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (169) 1qU_wnzcVKQZRFA6uV61UrbMzJoWrubrW4M5u361pAOEYGAu4JP4DCtyASmvkfgRBrGEujXZ2ans07ezDQxURfkTbi5tNadJc_wDwk7W7q51lnLalasno2UsgF6EpKdDc5k4muUfFi6AwDdJhsuqDhnMCoap0XkNPYzRbwhxZ
//fd state: (169) 1qU_wnzcVKQZRFA6uV61UrbMzJoWrubrW4M5u361pAOEYGAu4JP4DCtyASmvkfgRBrGEujXZ2ans07ezDQxURfkTbi5tNadJc_wDwk7W7q51lnLalasno2UsgF6EpKdDc5k4muUfFi6AwDdJhsuqDhnMCoap0XkNPYzRbwhxZ

ret = client.Write(fd862, []byte("p80K39_Y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (177) 1qU_wnzcVKQZRFA6uV61UrbMzJoWrubrW4M5u361pAOEYGAu4JP4DCtyASmvkfgRBrGEujXZ2ans07ezDQxURfkTbi5tNadJc_wDwk7W7q51lnLalasno2UsgF6EpKdDc5k4muUfFi6AwDdJhsuqDhnMCoap0XkNPYzRbwhxZp80K39_Y
//fd state: (0) Jl8hPyrSHnj5zG4frm9YEJYZNP6I1HMoqlb9yJiLDbJgAqU4UVtDqB9F1vLoyMfgH2bxvrAzA906ocnEhSxAQvEM1dQqiUaP74xTnYH_CECjrCp2zhxUYX

ret = client.Write(fd870, []byte("pBu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) pBuhPyrSHnj5zG4frm9YEJYZNP6I1HMoqlb9yJiLDbJgAqU4UVtDqB9F1vLoyMfgH2bxvrAzA906ocnEhSxAQvEM1dQqiUaP74xTnYH_CECjrCp2zhxUYX

ret = client.Close(fd864)
if(ret != 0) {
  panic("close failed")
}

//fd state: (3) pBuhPyrSHnj5zG4frm9YEJYZNP6I1HMoqlb9yJiLDbJgAqU4UVtDqB9F1vLoyMfgH2bxvrAzA906ocnEhSxAQvEM1dQqiUaP74xTnYH_CECjrCp2zhxUYX

ret = client.Write(fd870, []byte("TLDu23yhg7CnAPP8JKl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) pBuTLDu23yhg7CnAPP8JKlYZNP6I1HMoqlb9yJiLDbJgAqU4UVtDqB9F1vLoyMfgH2bxvrAzA906ocnEhSxAQvEM1dQqiUaP74xTnYH_CECjrCp2zhxUYX

fd871 := client.Open("/apJFPTK9eS.txt", client.O_RDWR|client.O_CREATE)
if(fd871 < 0) {
  panic("open failed")
}


ret = client.Close(fd867)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd868, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4y1Y9r0fR58jFcMlx8dTodfYOfb") {
  panic("wrong data returned")
}

//fd state: (109) _R_DPCJezXOeVZ5nZTzgv8eEUIOtHH9UnnDXY5NZQuTAy9mRlnSzo_tm7vQErhNxiEr7r1fib5PCUKnpA_4y1Y9r0fR58jFcMlx8dTodfYOfbMPBEFQuMu8fdOVGFeJMwU8RsEajjXEkNy1DLyT1rZlZFbodx9APXThZIjNDkzeqRoPfRv7anR

ret = client.Write(fd868, []byte("9WZEBuG9GvkYVMyh4IyDLx7lcTIPg7Rcv3WST6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (147) _R_DPCJezXOeVZ5nZTzgv8eEUIOtHH9UnnDXY5NZQuTAy9mRlnSzo_tm7vQErhNxiEr7r1fib5PCUKnpA_4y1Y9r0fR58jFcMlx8dTodfYOfb9WZEBuG9GvkYVMyh4IyDLx7lcTIPg7Rcv3WST61rZlZFbodx9APXThZIjNDkzeqRoPfRv7anR

ret = client.Close(fd870)
if(ret != 0) {
  panic("close failed")
}

//fd state: (147) _R_DPCJezXOeVZ5nZTzgv8eEUIOtHH9UnnDXY5NZQuTAy9mRlnSzo_tm7vQErhNxiEr7r1fib5PCUKnpA_4y1Y9r0fR58jFcMlx8dTodfYOfb9WZEBuG9GvkYVMyh4IyDLx7lcTIPg7Rcv3WST61rZlZFbodx9APXThZIjNDkzeqRoPfRv7anR

ret = client.Write(fd868, []byte("Z3RQirwpw0k0aYdTBGEwLDGgTyzBj0ey7Pjy2TRqTOhKR1mjdBONmw6q7kJEGEyfEPndRabRU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (220) _R_DPCJezXOeVZ5nZTzgv8eEUIOtHH9UnnDXY5NZQuTAy9mRlnSzo_tm7vQErhNxiEr7r1fib5PCUKnpA_4y1Y9r0fR58jFcMlx8dTodfYOfb9WZEBuG9GvkYVMyh4IyDLx7lcTIPg7Rcv3WST6Z3RQirwpw0k0aYdTBGEwLDGgTyzBj0ey7Pjy2TRqTOhKR1mjdBONmw6q7kJEGEyfEPndRabRU

ret = client.Seek(fd868, 190, client.SEEK_SET)
if(ret != 190) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 190)
  panic("seek failed")
}


ret = client.Close(fd871)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd868)
if(ret != 0) {
  panic("close failed")
}


fd872 := client.Open("/ez0X4ohZiJ.txt", client.O_RDWR|client.O_CREATE)
if(fd872 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd872, []byte("ryVoiuNk7Y8EfVHJxXKncSCRKkLCBn4W7wNTFWbh9ikDxGVKXJAPZbhf3FIN8y9gXuFdQVgwLXv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) ryVoiuNk7Y8EfVHJxXKncSCRKkLCBn4W7wNTFWbh9ikDxGVKXJAPZbhf3FIN8y9gXuFdQVgwLXv

buf, ret = client.Read(fd862, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd872, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd872, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


buf, ret = client.Read(fd872, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Wbh9ikDx") {
  panic("wrong data returned")
}


fd873 := client.Open("/U27uigcAVw.txt", client.O_RDWR|client.O_CREATE)
if(fd873 < 0) {
  panic("open failed")
}


ret = client.Seek(fd872, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Seek(fd862, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


buf, ret = client.Read(fd873, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "n8vuiZuGQNS1JqoBjHpTvQ9kMfrnpfaKxo3CqbvJ_ib5xfVxWJbCtQ") {
  panic("wrong data returned")
}

//fd state: (35) ryVoiuNk7Y8EfVHJxXKncSCRKkLCBn4W7wNTFWbh9ikDxGVKXJAPZbhf3FIN8y9gXuFdQVgwLXv

ret = client.Write(fd872, []byte("v5mTC2fIBP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) ryVoiuNk7Y8EfVHJxXKncSCRKkLCBn4W7wNv5mTC2fIBPGVKXJAPZbhf3FIN8y9gXuFdQVgwLXv

fd874 := client.Open("/3emFEq5txq.txt", client.O_RDWR|client.O_CREATE)
if(fd874 < 0) {
  panic("open failed")
}


fd875 := client.Open("/ZGMeCpQK4Z.txt", client.O_RDWR|client.O_CREATE)
if(fd875 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd872, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GVKXJAPZbhf3FIN8y9gXuFdQVgwLXv") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd862, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pAOEYGAu4JP4DCtyAS") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd874, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd872)
if(ret != 0) {
  panic("close failed")
}


fd876 := client.Open("/i8RkTeMGDt.txt", client.O_RDWR|client.O_CREATE)
if(fd876 < 0) {
  panic("open failed")
}


fd877 := client.Open("/cDsrk7sXrQ.txt", client.O_RDWR|client.O_CREATE)
if(fd877 < 0) {
  panic("open failed")
}


ret = client.Seek(fd877, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd876, []byte("Wn6QeXnXgYmKG1o0cJNTS6vWCOUKADrprZcyjwkbrFPlK29X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) Wn6QeXnXgYmKG1o0cJNTS6vWCOUKADrprZcyjwkbrFPlK29X

buf, ret = client.Read(fd877, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JHVwAGsWMz4XgnoB") {
  panic("wrong data returned")
}


ret = client.Seek(fd877, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


fd878 := client.Open("/0tyzNGz888.txt", client.O_RDWR|client.O_CREATE)
if(fd878 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd874, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd879 := client.Open("/3uj6y_uQsP.txt", client.O_RDWR|client.O_CREATE)
if(fd879 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd873)
if(ret != 0) {
  panic("close failed")
}


fd880 := client.Open("/SZ8EuT89RO.txt", client.O_RDWR|client.O_CREATE)
if(fd880 < 0) {
  panic("open failed")
}


ret = client.Seek(fd876, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd862)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd875, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd876)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd875, []byte("biPKU55nmZ3aoFaXPcL1Mr1cjydSWbrtpqkb4vcsmsCWgN7LPRlyMPgz2bp9nZWzJKQN_muYU9ZUK0GLSYFJohRWUG72H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) biPKU55nmZ3aoFaXPcL1Mr1cjydSWbrtpqkb4vcsmsCWgN7LPRlyMPgz2bp9nZWzJKQN_muYU9ZUK0GLSYFJohRWUG72H

ret = client.Seek(fd880, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd874, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd874, []byte("P5D2617EtIU1_TwLYaqlNIV4h_hV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) P5D2617EtIU1_TwLYaqlNIV4h_hV

ret = client.Seek(fd874, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


buf, ret = client.Read(fd880, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd874, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "U1_TwLYaqlNIV4h_hV") {
  panic("wrong data returned")
}


fd881 := client.Open("/VLlA7iApWX.txt", client.O_RDWR|client.O_CREATE)
if(fd881 < 0) {
  panic("open failed")
}


ret = client.Seek(fd875, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (28) P5D2617EtIU1_TwLYaqlNIV4h_hV

ret = client.Write(fd874, []byte("VxicuZXWbz8xNyKfSLwwMq_Fhs1B3YBnomSwiDYfLeoZB9xmsSJqubOIdT5o9TNWEy2HXKcVL4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) P5D2617EtIU1_TwLYaqlNIV4h_hVVxicuZXWbz8xNyKfSLwwMq_Fhs1B3YBnomSwiDYfLeoZB9xmsSJqubOIdT5o9TNWEy2HXKcVL4

fd882 := client.Open("/kA9LLagYX3.txt", client.O_RDWR|client.O_CREATE)
if(fd882 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd874, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd881, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd882, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (38) 7TobMM27p3DK8X5JHVwAGsWMz4XgnoB6rFKfLP7P9ClGaNAYCBIcDZR2R2hsl2mCa

ret = client.Write(fd877, []byte("KhKlf1ii1qqWyHczOz2tVoLaQVMt8P_MAPkqN_oTsd6G5s7qH1ig7sPeJAh3rGmKJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) 7TobMM27p3DK8X5JHVwAGsWMz4XgnoB6rFKfLPKhKlf1ii1qqWyHczOz2tVoLaQVMt8P_MAPkqN_oTsd6G5s7qH1ig7sPeJAh3rGmKJ

buf, ret = client.Read(fd874, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd880, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd877, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd883 := client.Open("/JDaZBTumnQ.txt", client.O_RDWR|client.O_CREATE)
if(fd883 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd883, []byte("F8CJXije0VglEcReGcONOjTNKkRFqWduFE4BhRGFP__cOxZDEc_hMhMWkoFoV9m1w2ykY_XY4qybX3C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) F8CJXije0VglEcReGcONOjTNKkRFqWduFE4BhRGFP__cOxZDEc_hMhMWkoFoV9m1w2ykY_XY4qybX3C

ret = client.Close(fd880)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd877, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd884 := client.Open("/YponeBVrOZ.txt", client.O_RDWR|client.O_CREATE)
if(fd884 < 0) {
  panic("open failed")
}


ret = client.Close(fd874)
if(ret != 0) {
  panic("close failed")
}


fd885 := client.Open("/_3UzxCa0Yu.txt", client.O_RDWR|client.O_CREATE)
if(fd885 < 0) {
  panic("open failed")
}


ret = client.Seek(fd885, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd878, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd886 := client.Open("/6cvgQ4BXlZ.txt", client.O_RDWR|client.O_CREATE)
if(fd886 < 0) {
  panic("open failed")
}


fd887 := client.Open("/hqQ4RRTmCR.txt", client.O_RDWR|client.O_CREATE)
if(fd887 < 0) {
  panic("open failed")
}


ret = client.Close(fd881)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd878, []byte("efc52EnkFVRmOZUhCY6H_iK2o8LbuvQdqCpP0oWW12Dd64XxtY1LHnYdssF46cxeeXN2k4YJ8L_OwCB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) efc52EnkFVRmOZUhCY6H_iK2o8LbuvQdqCpP0oWW12Dd64XxtY1LHnYdssF46cxeeXN2k4YJ8L_OwCB

fd888 := client.Open("/i9w3nnu9X1.txt", client.O_RDWR|client.O_CREATE)
if(fd888 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd886, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd888, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd889 := client.Open("/NntGsxrgb1.txt", client.O_RDWR|client.O_CREATE)
if(fd889 < 0) {
  panic("open failed")
}


ret = client.Seek(fd888, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd885, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd879, []byte("ltX13sRhCxD1IB7i75iMcPJwXHj_5_QbOBryk9DTBdVWfX4eCSIr1IDqaG7Z6KXA_k3Qynpl1uxWFb1dYcW7m"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) ltX13sRhCxD1IB7i75iMcPJwXHj_5_QbOBryk9DTBdVWfX4eCSIr1IDqaG7Z6KXA_k3Qynpl1uxWFb1dYcW7m

ret = client.Seek(fd877, 64, client.SEEK_SET)
if(ret != 64) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 64)
  panic("seek failed")
}


ret = client.Close(fd887)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd878, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd889)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd877, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


fd890 := client.Open("/3BfdqXctXO.txt", client.O_RDWR|client.O_CREATE)
if(fd890 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd890, []byte("kdjtU548C_CGHqXtMqcGuCI64unoR0GU62wohIrp3nSwJBrpGJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) kdjtU548C_CGHqXtMqcGuCI64unoR0GU62wohIrp3nSwJBrpGJ

fd891 := client.Open("/gF3CBhfROh.txt", client.O_RDWR|client.O_CREATE)
if(fd891 < 0) {
  panic("open failed")
}


fd892 := client.Open("/OZOyt6VkFz.txt", client.O_RDWR|client.O_CREATE)
if(fd892 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd892, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd886, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd893 := client.Open("/bsCsfx792m.txt", client.O_RDWR|client.O_CREATE)
if(fd893 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd893, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd894 := client.Open("/7wIy5jUAqf.txt", client.O_RDWR|client.O_CREATE)
if(fd894 < 0) {
  panic("open failed")
}


ret = client.Close(fd888)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd885, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd877)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd894, []byte("rd7PAH_dlF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) rd7PAH_dlF

buf, ret = client.Read(fd885, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd883, 65, client.SEEK_SET)
if(ret != 65) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 65)
  panic("seek failed")
}


ret = client.Seek(fd891, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd885, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd883, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd891, []byte("UJRzNQS3cm9neJMwqDP16VcbapuGu6Ouow9dJ88qfGoqoyr1pLT9WJD7fnw7SUVDoZx14qq9yX22u11SMn3f_4Q0yi35198"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) UJRzNQS3cm9neJMwqDP16VcbapuGu6Ouow9dJ88qfGoqoyr1pLT9WJD7fnw7SUVDoZx14qq9yX22u11SMn3f_4Q0yi35198

ret = client.Seek(fd875, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


fd895 := client.Open("/Rlp_5jzvVy.txt", client.O_RDWR|client.O_CREATE)
if(fd895 < 0) {
  panic("open failed")
}


ret = client.Seek(fd883, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Close(fd894)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd885)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd893)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd892, []byte("VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjG

ret = client.Seek(fd890, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


buf, ret = client.Read(fd884, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd884)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd879)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd882, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd891, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


buf, ret = client.Read(fd891, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Gu6Ouow9dJ88qfGoqoyr") {
  panic("wrong data returned")
}


ret = client.Close(fd883)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd891, 84, client.SEEK_SET)
if(ret != 84) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 84)
  panic("seek failed")
}

//fd state: (74) biPKU55nmZ3aoFaXPcL1Mr1cjydSWbrtpqkb4vcsmsCWgN7LPRlyMPgz2bp9nZWzJKQN_muYU9ZUK0GLSYFJohRWUG72H

ret = client.Write(fd875, []byte("9_vCQtFRtwYJ7wsuIpvZDRBwHP5DKod4_RfUQYSBXSnNpALs6eYOKDZa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) biPKU55nmZ3aoFaXPcL1Mr1cjydSWbrtpqkb4vcsmsCWgN7LPRlyMPgz2bp9nZWzJKQN_muYU99_vCQtFRtwYJ7wsuIpvZDRBwHP5DKod4_RfUQYSBXSnNpALs6eYOKDZa

fd896 := client.Open("/t5jEm7yeiJ.txt", client.O_RDWR|client.O_CREATE)
if(fd896 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd882, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd882, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd897 := client.Open("/KyMhRXeJjY.txt", client.O_RDWR|client.O_CREATE)
if(fd897 < 0) {
  panic("open failed")
}

//fd state: (55) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjG

ret = client.Write(fd892, []byte("bvyumrw1qG0ujLv0vUHuT_Lbc8ltRazarHZMA0MhtzkRRSCbMLFfCbTy_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (112) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjGbvyumrw1qG0ujLv0vUHuT_Lbc8ltRazarHZMA0MhtzkRRSCbMLFfCbTy_

ret = client.Seek(fd875, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}


buf, ret = client.Read(fd892, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (112) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjGbvyumrw1qG0ujLv0vUHuT_Lbc8ltRazarHZMA0MhtzkRRSCbMLFfCbTy_

ret = client.Write(fd892, []byte("GjyiNEGD1ebrSnvbpJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjGbvyumrw1qG0ujLv0vUHuT_Lbc8ltRazarHZMA0MhtzkRRSCbMLFfCbTy_GjyiNEGD1ebrSnvbpJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu

buf, ret = client.Read(fd878, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd878)
if(ret != 0) {
  panic("close failed")
}

//fd state: (84) UJRzNQS3cm9neJMwqDP16VcbapuGu6Ouow9dJ88qfGoqoyr1pLT9WJD7fnw7SUVDoZx14qq9yX22u11SMn3f_4Q0yi35198

ret = client.Write(fd891, []byte("PZP6ny4il5UPl0P_lY4Mh2GuuICjly_dd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) UJRzNQS3cm9neJMwqDP16VcbapuGu6Ouow9dJ88qfGoqoyr1pLT9WJD7fnw7SUVDoZx14qq9yX22u11SMn3fPZP6ny4il5UPl0P_lY4Mh2GuuICjly_dd
//fd state: (0) bMGNonU9ethbElp5KbiRCnptr06zLrZjaNxXoYWW_0WHSDnyw9F1z34gZxP0GJzafD8h5orG6fuRXlHI8yWtnQn0WgpNOwxzSIdAAb08e6XEEwBzLwtmncK

ret = client.Write(fd896, []byte("I6wJBlFSQ7bpSK4HGMj6XWUY_Q5UarWJkYhMpIy_K44VpYeALde7P7Mw3TTSO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) I6wJBlFSQ7bpSK4HGMj6XWUY_Q5UarWJkYhMpIy_K44VpYeALde7P7Mw3TTSOJzafD8h5orG6fuRXlHI8yWtnQn0WgpNOwxzSIdAAb08e6XEEwBzLwtmncK

ret = client.Seek(fd895, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd886)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd897, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "QSp0JT5YkVpJ9o8ieHLG5") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd898 := client.Open("/i5pxa6lkdZ.txt", client.O_RDWR|client.O_CREATE)
if(fd898 < 0) {
  panic("open failed")
}


ret = client.Seek(fd892, 76, client.SEEK_SET)
if(ret != 76) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 76)
  panic("seek failed")
}


fd899 := client.Open("/kcHPGxVqjP.txt", client.O_RDWR|client.O_CREATE)
if(fd899 < 0) {
  panic("open failed")
}


ret = client.Seek(fd892, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}


buf, ret = client.Read(fd895, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd890)
if(ret != 0) {
  panic("close failed")
}


fd900 := client.Open("/ad457d9kC_.txt", client.O_RDWR|client.O_CREATE)
if(fd900 < 0) {
  panic("open failed")
}


fd901 := client.Open("/zv13uXSfcj.txt", client.O_RDWR|client.O_CREATE)
if(fd901 < 0) {
  panic("open failed")
}


ret = client.Seek(fd895, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd897)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd898, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd875)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd901, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (70) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjGbvyumrw1qG0ujLv0vUHuT_Lbc8ltRazarHZMA0MhtzkRRSCbMLFfCbTy_GjyiNEGD1ebrSnvbpJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu

ret = client.Write(fd892, []byte("6r5zNYybw2b7g9GSdX3Y374li_bGkJg41dlHb7hq2YsIHgMF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjGbvyumrw1qG0ujLv6r5zNYybw2b7g9GSdX3Y374li_bGkJg41dlHb7hq2YsIHgMFGD1ebrSnvbpJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu

ret = client.Seek(fd895, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd901, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd898, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd900, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd902 := client.Open("/YJ0bjUslc0.txt", client.O_RDWR|client.O_CREATE)
if(fd902 < 0) {
  panic("open failed")
}


fd903 := client.Open("/08lzRI0lEO.txt", client.O_RDWR|client.O_CREATE)
if(fd903 < 0) {
  panic("open failed")
}

//fd state: (118) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjGbvyumrw1qG0ujLv6r5zNYybw2b7g9GSdX3Y374li_bGkJg41dlHb7hq2YsIHgMFGD1ebrSnvbpJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu

ret = client.Write(fd892, []byte("VzEz4L2Fwws"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (129) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjGbvyumrw1qG0ujLv6r5zNYybw2b7g9GSdX3Y374li_bGkJg41dlHb7hq2YsIHgMFVzEz4L2FwwsJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu

fd904 := client.Open("/XiBvGyIwPY.txt", client.O_RDWR|client.O_CREATE)
if(fd904 < 0) {
  panic("open failed")
}


ret = client.Seek(fd902, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


buf, ret = client.Read(fd898, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd900, []byte("L17IYF214mNZ7CGZQFPPrMvAzZ54X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) L17IYF214mNZ7CGZQFPPrMvAzZ54X

buf, ret = client.Read(fd882, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd896)
if(ret != 0) {
  panic("close failed")
}

//fd state: (21) QteKtL9lfbDCUGdadX7xngI2olRLoi6IiaQ

ret = client.Write(fd902, []byte("Zi6CcM_zubTLWx2RU4YkPXzT61uvCgjfdr56ykNzG2Ox2ITfuzze6vAYgDqquuoSRZSxesDtf8DSECVALBCuC4WOA2LZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) QteKtL9lfbDCUGdadX7xnZi6CcM_zubTLWx2RU4YkPXzT61uvCgjfdr56ykNzG2Ox2ITfuzze6vAYgDqquuoSRZSxesDtf8DSECVALBCuC4WOA2LZ

ret = client.Close(fd900)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd892, 99, client.SEEK_SET)
if(ret != 99) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 99)
  panic("seek failed")
}


fd905 := client.Open("/t6aEQSm_lw.txt", client.O_RDWR|client.O_CREATE)
if(fd905 < 0) {
  panic("open failed")
}


ret = client.Seek(fd895, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd904, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd906 := client.Open("/OYbOtLvtcd.txt", client.O_RDWR|client.O_CREATE)
if(fd906 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd898, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd901)
if(ret != 0) {
  panic("close failed")
}


fd907 := client.Open("/SNTEj6ymnz.txt", client.O_RDWR|client.O_CREATE)
if(fd907 < 0) {
  panic("open failed")
}


fd908 := client.Open("/jxEtjy8zaN.txt", client.O_RDWR|client.O_CREATE)
if(fd908 < 0) {
  panic("open failed")
}


ret = client.Seek(fd906, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd907)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd882)
if(ret != 0) {
  panic("close failed")
}


fd909 := client.Open("/d5bTRIS3Qm.txt", client.O_RDWR|client.O_CREATE)
if(fd909 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd898, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd908, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd898, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) vek5yBnmqCyeowwng1hCtBFC7WpEALfVVyV9qrdXdDdVv9TsXsgfbDhE4nzmbi7mspi1j2XHnCJO67FbITXRVyGPTBosVe2BsxXtvBEVAjBvMnuhamDWKKgRawz

ret = client.Write(fd899, []byte("9F2bD7crnT9w6WyJ9EbntsoDzms85GgX8oC6zT2IaUMNjiL9NckojxwkfjEVRGIcaFMW3ENsHgTad_7Kypu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) 9F2bD7crnT9w6WyJ9EbntsoDzms85GgX8oC6zT2IaUMNjiL9NckojxwkfjEVRGIcaFMW3ENsHgTad_7KypuRVyGPTBosVe2BsxXtvBEVAjBvMnuhamDWKKgRawz

fd910 := client.Open("/apJFPTK9eS.txt", client.O_RDWR|client.O_CREATE)
if(fd910 < 0) {
  panic("open failed")
}


ret = client.Close(fd899)
if(ret != 0) {
  panic("close failed")
}


fd911 := client.Open("/kqZrpEP1nq.txt", client.O_RDWR|client.O_CREATE)
if(fd911 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd905, []byte("ZF1AqzVqnESVLAwev6qEAMVj4zygDkq6BVwhw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) ZF1AqzVqnESVLAwev6qEAMVj4zygDkq6BVwhw

buf, ret = client.Read(fd909, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "17hWpe6DK73XI_Gcmz3xtFMrwOMBJLfIM3GM1R31WXzctPfLVVb_QGJ") {
  panic("wrong data returned")
}


ret = client.Seek(fd891, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd898, []byte("4DxuzE5rAWSlCfCVhMm3uPh486uPKQXR0_ZkGBfuej2DZioPjA_5IOGvHJDhF48XWrZvKR2d_A4L0AFWnTp1q7JIZrKX4hG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) 4DxuzE5rAWSlCfCVhMm3uPh486uPKQXR0_ZkGBfuej2DZioPjA_5IOGvHJDhF48XWrZvKR2d_A4L0AFWnTp1q7JIZrKX4hG

buf, ret = client.Read(fd892, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Jg41dlHb7hq2YsIHgMFVzEz4L2FwwsJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu") {
  panic("wrong data returned")
}


ret = client.Close(fd909)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd892, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd908)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd905, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd910, 68, client.SEEK_SET)
if(ret != 68) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 68)
  panic("seek failed")
}


fd912 := client.Open("/xJA7eZFNHf.txt", client.O_RDWR|client.O_CREATE)
if(fd912 < 0) {
  panic("open failed")
}


ret = client.Seek(fd892, 123, client.SEEK_SET)
if(ret != 123) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 123)
  panic("seek failed")
}


ret = client.Seek(fd902, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}

//fd state: (0) jaWRIH8cglLg6w84lUtpKF_BJO3gKIxu5FyLTzagU513q6d873Tl97jiMBDckB189azuhP

ret = client.Write(fd903, []byte("lnNBxrQAw_jjnUknDUJMm2UuHj0BPOUseIdCxrFiBG3h"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) lnNBxrQAw_jjnUknDUJMm2UuHj0BPOUseIdCxrFiBG3hq6d873Tl97jiMBDckB189azuhP

ret = client.Close(fd904)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd902, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Zi6CcM_zubTLWx2RU4YkPXzT61uvCgjfdr56ykNzG2Ox2ITfu") {
  panic("wrong data returned")
}


ret = client.Close(fd898)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd892, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd911, []byte("jqDXuBgBukHZzdc1J_4WXYcBFT8z9mlcDvB3W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) jqDXuBgBukHZzdc1J_4WXYcBFT8z9mlcDvB3W

buf, ret = client.Read(fd903, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "q6d873Tl97jiMBD") {
  panic("wrong data returned")
}


fd913 := client.Open("/JlmdaEMRbu.txt", client.O_RDWR|client.O_CREATE)
if(fd913 < 0) {
  panic("open failed")
}


ret = client.Seek(fd903, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd895, []byte("pABrwes79MUkv6UGBUru1T2Nf1ZnlDtYORBUWdUl5GJKVxbydtmfO3zexQhj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) pABrwes79MUkv6UGBUru1T2Nf1ZnlDtYORBUWdUl5GJKVxbydtmfO3zexQhj

ret = client.Close(fd892)
if(ret != 0) {
  panic("close failed")
}


fd914 := client.Open("/BLyjaSEDXE.txt", client.O_RDWR|client.O_CREATE)
if(fd914 < 0) {
  panic("open failed")
}


ret = client.Close(fd914)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd913, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (72) UJRzNQS3cm9neJMwqDP16VcbapuGu6Ouow9dJ88qfGoqoyr1pLT9WJD7fnw7SUVDoZx14qq9yX22u11SMn3fPZP6ny4il5UPl0P_lY4Mh2GuuICjly_dd

ret = client.Write(fd891, []byte("N8Uda83QOLm48I1g8DlUozon3aO6_LOPl49f31GmYZ7rEuSlUsh5AiAZQZy5pw87FSUY7PXgkzvRcRdEw3UUREQzZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (161) UJRzNQS3cm9neJMwqDP16VcbapuGu6Ouow9dJ88qfGoqoyr1pLT9WJD7fnw7SUVDoZx14qq9N8Uda83QOLm48I1g8DlUozon3aO6_LOPl49f31GmYZ7rEuSlUsh5AiAZQZy5pw87FSUY7PXgkzvRcRdEw3UUREQzZ
//fd state: (60) pABrwes79MUkv6UGBUru1T2Nf1ZnlDtYORBUWdUl5GJKVxbydtmfO3zexQhj

ret = client.Write(fd895, []byte("tcZJ9RhAkDFoxuUY0QOEDCXJPIcL1IjLUQ6hT28sw7jd7T_mElC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) pABrwes79MUkv6UGBUru1T2Nf1ZnlDtYORBUWdUl5GJKVxbydtmfO3zexQhjtcZJ9RhAkDFoxuUY0QOEDCXJPIcL1IjLUQ6hT28sw7jd7T_mElC

ret = client.Close(fd891)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd902, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "zze6vAYgDqquuoSRZSx") {
  panic("wrong data returned")
}

//fd state: (37) jqDXuBgBukHZzdc1J_4WXYcBFT8z9mlcDvB3W

ret = client.Write(fd911, []byte("uHZLNt2lPpg8R_4mXt5KkbI85sLO5zIzCIeOO2i1swcUdPF8A37ipqYJjF4ifXx6Ews1oHv2hT16juOO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) jqDXuBgBukHZzdc1J_4WXYcBFT8z9mlcDvB3WuHZLNt2lPpg8R_4mXt5KkbI85sLO5zIzCIeOO2i1swcUdPF8A37ipqYJjF4ifXx6Ews1oHv2hT16juOO
//fd state: (0) 

ret = client.Write(fd912, []byte("2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5

ret = client.Seek(fd902, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}

//fd state: (54) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5

ret = client.Write(fd912, []byte("VbV3SR4KfjTr_PGs_b5LA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5VbV3SR4KfjTr_PGs_b5LA

buf, ret = client.Read(fd906, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd910, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Cr1CyQMceaJnbLUeg") {
  panic("wrong data returned")
}


fd915 := client.Open("/RFIaVpGCI_.txt", client.O_RDWR|client.O_CREATE)
if(fd915 < 0) {
  panic("open failed")
}


fd916 := client.Open("/XAbT_qsPjH.txt", client.O_RDWR|client.O_CREATE)
if(fd916 < 0) {
  panic("open failed")
}


ret = client.Close(fd910)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd895, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd895)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd905)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd915, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd916, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd917 := client.Open("/EZgeVdm5j2.txt", client.O_RDWR|client.O_CREATE)
if(fd917 < 0) {
  panic("open failed")
}

//fd state: (8) QteKtL9lfbDCUGdadX7xnZi6CcM_zubTLWx2RU4YkPXzT61uvCgjfdr56ykNzG2Ox2ITfuzze6vAYgDqquuoSRZSxesDtf8DSECVALBCuC4WOA2LZ

ret = client.Write(fd902, []byte("u8o7ypaF6r7G1TCDFZmRQaDf48MlMPGFnCDr_Xk4MkkiCfViNKI4qetw9ZCsI2blf829BndTtkcslame"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) QteKtL9lu8o7ypaF6r7G1TCDFZmRQaDf48MlMPGFnCDr_Xk4MkkiCfViNKI4qetw9ZCsI2blf829BndTtkcslamexesDtf8DSECVALBCuC4WOA2LZ
//fd state: (0) 

ret = client.Write(fd917, []byte("z3S69gRqQuRqVbEy3wv3sJ5fj9fxFMMpiuvLFDWG4TLX0Xow6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) z3S69gRqQuRqVbEy3wv3sJ5fj9fxFMMpiuvLFDWG4TLX0Xow6

ret = client.Close(fd903)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd912, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd906, []byte("rnS2OqyzVTXz7MqI90O__K92HP0SAJJyEPRTcNivUU4yh8yyuz2ak8DlsnlCt4KRZJLsLBq21uE9xrQvE8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) rnS2OqyzVTXz7MqI90O__K92HP0SAJJyEPRTcNivUU4yh8yyuz2ak8DlsnlCt4KRZJLsLBq21uE9xrQvE8
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd917, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd916, []byte("70ZBWqoBrQuaE3puwp7OW_rqdW4gvEmKv4qvYWC68A8ODQjSHW2t9BXy6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) 70ZBWqoBrQuaE3puwp7OW_rqdW4gvEmKv4qvYWC68A8ODQjSHW2t9BXy6
//fd state: (88) QteKtL9lu8o7ypaF6r7G1TCDFZmRQaDf48MlMPGFnCDr_Xk4MkkiCfViNKI4qetw9ZCsI2blf829BndTtkcslamexesDtf8DSECVALBCuC4WOA2LZ

ret = client.Write(fd902, []byte("MUvYWt9MG2_fwdwCGHLGdiOr6G4EQVgT2Q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) QteKtL9lu8o7ypaF6r7G1TCDFZmRQaDf48MlMPGFnCDr_Xk4MkkiCfViNKI4qetw9ZCsI2blf829BndTtkcslameMUvYWt9MG2_fwdwCGHLGdiOr6G4EQVgT2Q

buf, ret = client.Read(fd917, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "z3S69gRqQuRqVbEy3wv3sJ5fj9fxFMMpiuvLFDWG4TLX0Xow6") {
  panic("wrong data returned")
}


ret = client.Seek(fd906, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


buf, ret = client.Read(fd913, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd917, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd902, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (75) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5VbV3SR4KfjTr_PGs_b5LA

ret = client.Write(fd912, []byte("ah0NKaQ41g3BrDzkE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5VbV3SR4KfjTr_PGs_b5LAah0NKaQ41g3BrDzkE

ret = client.Close(fd911)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd916, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


ret = client.Seek(fd916, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}

//fd state: (92) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5VbV3SR4KfjTr_PGs_b5LAah0NKaQ41g3BrDzkE

ret = client.Write(fd912, []byte("Ii8l6pV_AUodycPXqPrzZP7QuRhCxPgPiR5QKj8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (131) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5VbV3SR4KfjTr_PGs_b5LAah0NKaQ41g3BrDzkEIi8l6pV_AUodycPXqPrzZP7QuRhCxPgPiR5QKj8

ret = client.Seek(fd916, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


buf, ret = client.Read(fd916, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "aE3puwp7OW_rqdW4gvEmKv4qvYWC68A8ODQjSHW2t9BXy6") {
  panic("wrong data returned")
}

//fd state: (57) 70ZBWqoBrQuaE3puwp7OW_rqdW4gvEmKv4qvYWC68A8ODQjSHW2t9BXy6

ret = client.Write(fd916, []byte("c_tZWmOlKy0odi5bfj23QLVVzoL3Oh__YVffOlKxp3Oh0BJArmA1sKOpt9o4xmz_wp6qM0EN14DKfscr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (137) 70ZBWqoBrQuaE3puwp7OW_rqdW4gvEmKv4qvYWC68A8ODQjSHW2t9BXy6c_tZWmOlKy0odi5bfj23QLVVzoL3Oh__YVffOlKxp3Oh0BJArmA1sKOpt9o4xmz_wp6qM0EN14DKfscr

fd918 := client.Open("/4jun8IF924.txt", client.O_RDWR|client.O_CREATE)
if(fd918 < 0) {
  panic("open failed")
}


ret = client.Close(fd902)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd913, []byte("T1lf9JBolgFh6jf28W3NJsWZO1y3lGeiMCLOFJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) T1lf9JBolgFh6jf28W3NJsWZO1y3lGeiMCLOFJ

ret = client.Seek(fd906, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


ret = client.Close(fd916)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd913, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd915)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd917)
if(ret != 0) {
  panic("close failed")
}

//fd state: (131) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5VbV3SR4KfjTr_PGs_b5LAah0NKaQ41g3BrDzkEIi8l6pV_AUodycPXqPrzZP7QuRhCxPgPiR5QKj8

ret = client.Write(fd912, []byte("cxpPjEspG8WipDEQjwQZOULpmWLdh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (160) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5VbV3SR4KfjTr_PGs_b5LAah0NKaQ41g3BrDzkEIi8l6pV_AUodycPXqPrzZP7QuRhCxPgPiR5QKj8cxpPjEspG8WipDEQjwQZOULpmWLdh

fd919 := client.Open("/w5X53pcgws.txt", client.O_RDWR|client.O_CREATE)
if(fd919 < 0) {
  panic("open failed")
}


fd920 := client.Open("/hythcT1zzr.txt", client.O_RDWR|client.O_CREATE)
if(fd920 < 0) {
  panic("open failed")
}


ret = client.Close(fd919)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd906, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


fd921 := client.Open("/cK6VMUXJqT.txt", client.O_RDWR|client.O_CREATE)
if(fd921 < 0) {
  panic("open failed")
}


fd922 := client.Open("/mTxm62dMSA.txt", client.O_RDWR|client.O_CREATE)
if(fd922 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd918, []byte("my9rqCIiksJdwStw7npFYkQ_K28ju1DW0G0NlDEJDFz__8UyBiigTB6gCqVHrZO5JLMhqnFPCJm5FUE6VP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) my9rqCIiksJdwStw7npFYkQ_K28ju1DW0G0NlDEJDFz__8UyBiigTB6gCqVHrZO5JLMhqnFPCJm5FUE6VP
//fd state: (160) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5VbV3SR4KfjTr_PGs_b5LAah0NKaQ41g3BrDzkEIi8l6pV_AUodycPXqPrzZP7QuRhCxPgPiR5QKj8cxpPjEspG8WipDEQjwQZOULpmWLdh

ret = client.Write(fd912, []byte("oDT8YWAjY3BHZWTJgc4dlEMIa0cAJzAVW2gBdIKzqq8Y_MFD72NxuWtgbTQ4eP9w8UcCHv2qjvgL0OD4yp9AORG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (247) 2NQpUXeJCdDAhuE3QWaItnQZiYuYW9uLqzxZoDoWFYM0aheREK70n5VbV3SR4KfjTr_PGs_b5LAah0NKaQ41g3BrDzkEIi8l6pV_AUodycPXqPrzZP7QuRhCxPgPiR5QKj8cxpPjEspG8WipDEQjwQZOULpmWLdhoDT8YWAjY3BHZWTJgc4dlEMIa0cAJzAVW2gBdIKzqq8Y_MFD72NxuWtgbTQ4eP9w8UcCHv2qjvgL0OD4yp9AORG

ret = client.Close(fd912)
if(ret != 0) {
  panic("close failed")
}

//fd state: (38) T1lf9JBolgFh6jf28W3NJsWZO1y3lGeiMCLOFJ

ret = client.Write(fd913, []byte("isuThQcJW3nuLjKuMIRO5tHCsdkifsHaApi2tWO3eGhLyiSiC9bKQxvBYRWfcQAMrmSXUpmaJly4L8pDh9ihQAOdsqSj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) T1lf9JBolgFh6jf28W3NJsWZO1y3lGeiMCLOFJisuThQcJW3nuLjKuMIRO5tHCsdkifsHaApi2tWO3eGhLyiSiC9bKQxvBYRWfcQAMrmSXUpmaJly4L8pDh9ihQAOdsqSj

ret = client.Close(fd920)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd906, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yEPRTcNivUU4yh8yyuz2ak8DlsnlCt4KRZJ") {
  panic("wrong data returned")
}


fd923 := client.Open("/b3GRSU2uJb.txt", client.O_RDWR|client.O_CREATE)
if(fd923 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd906, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LsLBq21uE9xrQvE8") {
  panic("wrong data returned")
}

//fd state: (130) T1lf9JBolgFh6jf28W3NJsWZO1y3lGeiMCLOFJisuThQcJW3nuLjKuMIRO5tHCsdkifsHaApi2tWO3eGhLyiSiC9bKQxvBYRWfcQAMrmSXUpmaJly4L8pDh9ihQAOdsqSj

ret = client.Write(fd913, []byte("D1B4TUZH5cGzyk5lmSyjcAWg0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (155) T1lf9JBolgFh6jf28W3NJsWZO1y3lGeiMCLOFJisuThQcJW3nuLjKuMIRO5tHCsdkifsHaApi2tWO3eGhLyiSiC9bKQxvBYRWfcQAMrmSXUpmaJly4L8pDh9ihQAOdsqSjD1B4TUZH5cGzyk5lmSyjcAWg0

ret = client.Close(fd906)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd913, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd921, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd921, []byte("xDuGoi2nmwqCcB0coq6ToDlv1ALG40jOh57iSGSCl6bOJNTn3CQaRxm0OZUG9pnRM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) xDuGoi2nmwqCcB0coq6ToDlv1ALG40jOh57iSGSCl6bOJNTn3CQaRxm0OZUG9pnRM
//fd state: (155) T1lf9JBolgFh6jf28W3NJsWZO1y3lGeiMCLOFJisuThQcJW3nuLjKuMIRO5tHCsdkifsHaApi2tWO3eGhLyiSiC9bKQxvBYRWfcQAMrmSXUpmaJly4L8pDh9ihQAOdsqSjD1B4TUZH5cGzyk5lmSyjcAWg0

ret = client.Write(fd913, []byte("fR8KNe7oHMg2hmCxYziGC0hWRzu2FNPD8Jd85gwWozUd4vGyWi7Sv4bCw4RGzlEVE2_3Q4MaH52upM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (233) T1lf9JBolgFh6jf28W3NJsWZO1y3lGeiMCLOFJisuThQcJW3nuLjKuMIRO5tHCsdkifsHaApi2tWO3eGhLyiSiC9bKQxvBYRWfcQAMrmSXUpmaJly4L8pDh9ihQAOdsqSjD1B4TUZH5cGzyk5lmSyjcAWg0fR8KNe7oHMg2hmCxYziGC0hWRzu2FNPD8Jd85gwWozUd4vGyWi7Sv4bCw4RGzlEVE2_3Q4MaH52upM

ret = client.Close(fd913)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd918, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd921)
if(ret != 0) {
  panic("close failed")
}


fd924 := client.Open("/XgqlK6vDL2.txt", client.O_RDWR|client.O_CREATE)
if(fd924 < 0) {
  panic("open failed")
}

//fd state: (82) my9rqCIiksJdwStw7npFYkQ_K28ju1DW0G0NlDEJDFz__8UyBiigTB6gCqVHrZO5JLMhqnFPCJm5FUE6VP

ret = client.Write(fd918, []byte("C9TOJxT10SA0_bDrP2kVrLLDQ28Qs2gFaKw2B5Ri982"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (125) my9rqCIiksJdwStw7npFYkQ_K28ju1DW0G0NlDEJDFz__8UyBiigTB6gCqVHrZO5JLMhqnFPCJm5FUE6VPC9TOJxT10SA0_bDrP2kVrLLDQ28Qs2gFaKw2B5Ri982
//fd state: (0) 

ret = client.Write(fd923, []byte("80SIZWQThbcPzbr5szAeDtDmXPXyUJMicdzXwi7FhBMeodLyBr3d25REbtP7S5NI_pOBPSF3AAAU0iNsnS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) 80SIZWQThbcPzbr5szAeDtDmXPXyUJMicdzXwi7FhBMeodLyBr3d25REbtP7S5NI_pOBPSF3AAAU0iNsnS

ret = client.Close(fd918)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd922)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd923, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd924, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WZ_fJw_Z4SOlCL5rgTDoGtI") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd924, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "XFrrSyZVJ7lnMZfRnN") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd923, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (82) 80SIZWQThbcPzbr5szAeDtDmXPXyUJMicdzXwi7FhBMeodLyBr3d25REbtP7S5NI_pOBPSF3AAAU0iNsnS

ret = client.Write(fd923, []byte("yGwpyOoHzaOj4Nc41veQ2tAr_U6MpQoZ1uoMAZH9id8rOhv8PZ7aRl4iZkpSG1znl7Yh6HxYpB6dFXGJDVS7B9bcbSAUVa9D"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (178) 80SIZWQThbcPzbr5szAeDtDmXPXyUJMicdzXwi7FhBMeodLyBr3d25REbtP7S5NI_pOBPSF3AAAU0iNsnSyGwpyOoHzaOj4Nc41veQ2tAr_U6MpQoZ1uoMAZH9id8rOhv8PZ7aRl4iZkpSG1znl7Yh6HxYpB6dFXGJDVS7B9bcbSAUVa9D

ret = client.Close(fd924)
if(ret != 0) {
  panic("close failed")
}

//fd state: (178) 80SIZWQThbcPzbr5szAeDtDmXPXyUJMicdzXwi7FhBMeodLyBr3d25REbtP7S5NI_pOBPSF3AAAU0iNsnSyGwpyOoHzaOj4Nc41veQ2tAr_U6MpQoZ1uoMAZH9id8rOhv8PZ7aRl4iZkpSG1znl7Yh6HxYpB6dFXGJDVS7B9bcbSAUVa9D

ret = client.Write(fd923, []byte("xkbmBaw6WBOnP2lksS5l1BqEE5nxhDe4DW5_LeRMw8XkGqs9LgdrT55xHwjY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (238) 80SIZWQThbcPzbr5szAeDtDmXPXyUJMicdzXwi7FhBMeodLyBr3d25REbtP7S5NI_pOBPSF3AAAU0iNsnSyGwpyOoHzaOj4Nc41veQ2tAr_U6MpQoZ1uoMAZH9id8rOhv8PZ7aRl4iZkpSG1znl7Yh6HxYpB6dFXGJDVS7B9bcbSAUVa9DxkbmBaw6WBOnP2lksS5l1BqEE5nxhDe4DW5_LeRMw8XkGqs9LgdrT55xHwjY

buf, ret = client.Read(fd923, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd923, 186, client.SEEK_SET)
if(ret != 186) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 186)
  panic("seek failed")
}

//fd state: (186) 80SIZWQThbcPzbr5szAeDtDmXPXyUJMicdzXwi7FhBMeodLyBr3d25REbtP7S5NI_pOBPSF3AAAU0iNsnSyGwpyOoHzaOj4Nc41veQ2tAr_U6MpQoZ1uoMAZH9id8rOhv8PZ7aRl4iZkpSG1znl7Yh6HxYpB6dFXGJDVS7B9bcbSAUVa9DxkbmBaw6WBOnP2lksS5l1BqEE5nxhDe4DW5_LeRMw8XkGqs9LgdrT55xHwjY

ret = client.Write(fd923, []byte("cuILvKlJFSvsfZtz5Qqk79BwQI85WRVKP0MLmqlDvumG27OZzVRke6Cb9W2Y7Os3eFUtLxhsHeIMwZIyINCMzPiamsz00"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (279) 80SIZWQThbcPzbr5szAeDtDmXPXyUJMicdzXwi7FhBMeodLyBr3d25REbtP7S5NI_pOBPSF3AAAU0iNsnSyGwpyOoHzaOj4Nc41veQ2tAr_U6MpQoZ1uoMAZH9id8rOhv8PZ7aRl4iZkpSG1znl7Yh6HxYpB6dFXGJDVS7B9bcbSAUVa9DxkbmBaw6cuILvKlJFSvsfZtz5Qqk79BwQI85WRVKP0MLmqlDvumG27OZzVRke6Cb9W2Y7Os3eFUtLxhsHeIMwZIyINCMzPiamsz00

fd925 := client.Open("/7PgE4R0k3A.txt", client.O_RDWR|client.O_CREATE)
if(fd925 < 0) {
  panic("open failed")
}


ret = client.Close(fd923)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd925, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd925, []byte("RYwsSAfQZTUvFBYNHyKCaCbft5cNi4szeJkAjHCvSij0sQZ_4Pe9f5_rbTuZLchO_aZB9VJa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) RYwsSAfQZTUvFBYNHyKCaCbft5cNi4szeJkAjHCvSij0sQZ_4Pe9f5_rbTuZLchO_aZB9VJa
//fd state: (72) RYwsSAfQZTUvFBYNHyKCaCbft5cNi4szeJkAjHCvSij0sQZ_4Pe9f5_rbTuZLchO_aZB9VJa

ret = client.Write(fd925, []byte("O2KYgzKC3K09VI02h95EGxESSSmnyE9zJzgzQNsWthyVLSJ3apxKj_3NiL2enYFKUg1izWw58p2rZ6YpL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (153) RYwsSAfQZTUvFBYNHyKCaCbft5cNi4szeJkAjHCvSij0sQZ_4Pe9f5_rbTuZLchO_aZB9VJaO2KYgzKC3K09VI02h95EGxESSSmnyE9zJzgzQNsWthyVLSJ3apxKj_3NiL2enYFKUg1izWw58p2rZ6YpL
//fd state: (153) RYwsSAfQZTUvFBYNHyKCaCbft5cNi4szeJkAjHCvSij0sQZ_4Pe9f5_rbTuZLchO_aZB9VJaO2KYgzKC3K09VI02h95EGxESSSmnyE9zJzgzQNsWthyVLSJ3apxKj_3NiL2enYFKUg1izWw58p2rZ6YpL

ret = client.Write(fd925, []byte("VSa0kgOydLoPKMQR_hNsTP_S6XMX5IK8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (185) RYwsSAfQZTUvFBYNHyKCaCbft5cNi4szeJkAjHCvSij0sQZ_4Pe9f5_rbTuZLchO_aZB9VJaO2KYgzKC3K09VI02h95EGxESSSmnyE9zJzgzQNsWthyVLSJ3apxKj_3NiL2enYFKUg1izWw58p2rZ6YpLVSa0kgOydLoPKMQR_hNsTP_S6XMX5IK8

fd926 := client.Open("/FDM_Iwp3SC.txt", client.O_RDWR|client.O_CREATE)
if(fd926 < 0) {
  panic("open failed")
}


fd927 := client.Open("/ct9OAepmrS.txt", client.O_RDWR|client.O_CREATE)
if(fd927 < 0) {
  panic("open failed")
}


ret = client.Close(fd925)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd927, []byte("zx5d2NpYdbfMaM1o4UayBsCXo18waoWAeviy0ScNpMAuU25ndDukuPYfGPoKhIPOybtWFo0J0MK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) zx5d2NpYdbfMaM1o4UayBsCXo18waoWAeviy0ScNpMAuU25ndDukuPYfGPoKhIPOybtWFo0J0MK

fd928 := client.Open("/5MukqgTcRB.txt", client.O_RDWR|client.O_CREATE)
if(fd928 < 0) {
  panic("open failed")
}


ret = client.Close(fd927)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd928, []byte("t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) t

ret = client.Close(fd928)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd926, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd926)
if(ret != 0) {
  panic("close failed")
}


fd929 := client.Open("/cK6VMUXJqT.txt", client.O_RDWR|client.O_CREATE)
if(fd929 < 0) {
  panic("open failed")
}


ret = client.Seek(fd929, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (2) xDuGoi2nmwqCcB0coq6ToDlv1ALG40jOh57iSGSCl6bOJNTn3CQaRxm0OZUG9pnRM

ret = client.Write(fd929, []byte("A8ND8bjNRb06obH91mCSOT5z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) xDA8ND8bjNRb06obH91mCSOT5zLG40jOh57iSGSCl6bOJNTn3CQaRxm0OZUG9pnRM

buf, ret = client.Read(fd929, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LG") {
  panic("wrong data returned")
}


ret = client.Seek(fd929, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}

//fd state: (46) xDA8ND8bjNRb06obH91mCSOT5zLG40jOh57iSGSCl6bOJNTn3CQaRxm0OZUG9pnRM

ret = client.Write(fd929, []byte("vq8tFvkpKORuIQbr3oUfnGfvo9FW8WMnoCQWOvl2cNOYJAIWU77_YXMjqJD6r1Lka8ECg5DWzb_iP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) xDA8ND8bjNRb06obH91mCSOT5zLG40jOh57iSGSCl6bOJNvq8tFvkpKORuIQbr3oUfnGfvo9FW8WMnoCQWOvl2cNOYJAIWU77_YXMjqJD6r1Lka8ECg5DWzb_iP

buf, ret = client.Read(fd929, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd929)
if(ret != 0) {
  panic("close failed")
}


fd930 := client.Open("/7hzYrijoQ0.txt", client.O_RDWR|client.O_CREATE)
if(fd930 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd930, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd930, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd930, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd930, []byte("W_jJuO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) W_jJuO

buf, ret = client.Read(fd930, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd930, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd931 := client.Open("/nQ29N2qWtv.txt", client.O_RDWR|client.O_CREATE)
if(fd931 < 0) {
  panic("open failed")
}


fd932 := client.Open("/YTjyADNkIX.txt", client.O_RDWR|client.O_CREATE)
if(fd932 < 0) {
  panic("open failed")
}

//fd state: (6) W_jJuO

ret = client.Write(fd930, []byte("rSVCetJz1IgiI4m0pYiNMJFBxgIa_iErKO3YfkuGxq0ei8lSeZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) W_jJuOrSVCetJz1IgiI4m0pYiNMJFBxgIa_iErKO3YfkuGxq0ei8lSeZ

fd933 := client.Open("/I3wdKtSiv1.txt", client.O_RDWR|client.O_CREATE)
if(fd933 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd931, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "a5GOQvtggCLN1rQsUVDy4mVBeOpCSk_puqjmi0rw8LHsLck7lc") {
  panic("wrong data returned")
}


ret = client.Close(fd931)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd930)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd932, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd934 := client.Open("/q757Jr9Xdi.txt", client.O_RDWR|client.O_CREATE)
if(fd934 < 0) {
  panic("open failed")
}


ret = client.Close(fd932)
if(ret != 0) {
  panic("close failed")
}


fd935 := client.Open("/y0dHmoJ4hA.txt", client.O_RDWR|client.O_CREATE)
if(fd935 < 0) {
  panic("open failed")
}


fd936 := client.Open("/1K0omrYdDW.txt", client.O_RDWR|client.O_CREATE)
if(fd936 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd934, []byte("Asx4jUMA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) Asx4jUMA

buf, ret = client.Read(fd936, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd935, []byte("twojFB7nyHbI4FuVYNPOWN7f1WnAQrQZYReSGm77wGxsDEeImy0XNq1Vc8Z6s0SJ4ST_WWEhlr7NmsRMD9X82Weh8jcZqF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) twojFB7nyHbI4FuVYNPOWN7f1WnAQrQZYReSGm77wGxsDEeImy0XNq1Vc8Z6s0SJ4ST_WWEhlr7NmsRMD9X82Weh8jcZqF

ret = client.Seek(fd934, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Close(fd936)
if(ret != 0) {
  panic("close failed")
}

//fd state: (94) twojFB7nyHbI4FuVYNPOWN7f1WnAQrQZYReSGm77wGxsDEeImy0XNq1Vc8Z6s0SJ4ST_WWEhlr7NmsRMD9X82Weh8jcZqF

ret = client.Write(fd935, []byte("_UcT0fjwCuxqBsgAYImN2hcoH_387dW0z8uEX5ycMYFuIpVr4o7ej_akQjIVx8_0Pkzl8C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) twojFB7nyHbI4FuVYNPOWN7f1WnAQrQZYReSGm77wGxsDEeImy0XNq1Vc8Z6s0SJ4ST_WWEhlr7NmsRMD9X82Weh8jcZqF_UcT0fjwCuxqBsgAYImN2hcoH_387dW0z8uEX5ycMYFuIpVr4o7ej_akQjIVx8_0Pkzl8C

buf, ret = client.Read(fd933, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd937 := client.Open("/FyoozucGND.txt", client.O_RDWR|client.O_CREATE)
if(fd937 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd935, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd938 := client.Open("/_gbGbZZtm6.txt", client.O_RDWR|client.O_CREATE)
if(fd938 < 0) {
  panic("open failed")
}


fd939 := client.Open("/X9xnpa9VHS.txt", client.O_RDWR|client.O_CREATE)
if(fd939 < 0) {
  panic("open failed")
}


fd940 := client.Open("/Ik6J5ZX4Xj.txt", client.O_RDWR|client.O_CREATE)
if(fd940 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd939, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd934, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


fd941 := client.Open("/mHY5KcNlPE.txt", client.O_RDWR|client.O_CREATE)
if(fd941 < 0) {
  panic("open failed")
}


ret = client.Close(fd939)
if(ret != 0) {
  panic("close failed")
}


fd942 := client.Open("/nGCywDCqun.txt", client.O_RDWR|client.O_CREATE)
if(fd942 < 0) {
  panic("open failed")
}


ret = client.Seek(fd934, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Seek(fd941, 247, client.SEEK_SET)
if(ret != 247) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 247)
  panic("seek failed")
}


ret = client.Close(fd941)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd940, []byte("oqJ9yAnLjzDOgFqwbquTEAF9mV3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) oqJ9yAnLjzDOgFqwbquTEAF9mV3

fd943 := client.Open("/7_nT83eVgg.txt", client.O_RDWR|client.O_CREATE)
if(fd943 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd943, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd942, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd944 := client.Open("/nQ29N2qWtv.txt", client.O_RDWR|client.O_CREATE)
if(fd944 < 0) {
  panic("open failed")
}


fd945 := client.Open("/7hzYrijoQ0.txt", client.O_RDWR|client.O_CREATE)
if(fd945 < 0) {
  panic("open failed")
}


ret = client.Seek(fd935, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}


ret = client.Seek(fd938, 101, client.SEEK_SET)
if(ret != 101) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 101)
  panic("seek failed")
}


buf, ret = client.Read(fd942, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (27) oqJ9yAnLjzDOgFqwbquTEAF9mV3

ret = client.Write(fd940, []byte("UINVLVmuLt8cMw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) oqJ9yAnLjzDOgFqwbquTEAF9mV3UINVLVmuLt8cMw

buf, ret = client.Read(fd937, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd944)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd933, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (41) oqJ9yAnLjzDOgFqwbquTEAF9mV3UINVLVmuLt8cMw

ret = client.Write(fd940, []byte("UM879BmTdwIUfqP9Fy0cCdXgcEI40QMjs0zJbvbSLqGHuWqwJaI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) oqJ9yAnLjzDOgFqwbquTEAF9mV3UINVLVmuLt8cMwUM879BmTdwIUfqP9Fy0cCdXgcEI40QMjs0zJbvbSLqGHuWqwJaI
//fd state: (92) oqJ9yAnLjzDOgFqwbquTEAF9mV3UINVLVmuLt8cMwUM879BmTdwIUfqP9Fy0cCdXgcEI40QMjs0zJbvbSLqGHuWqwJaI

ret = client.Write(fd940, []byte("fPRE3MhhUHv_ndJyiVWbfhNH7uBxjxqkRpbVbvN_Z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) oqJ9yAnLjzDOgFqwbquTEAF9mV3UINVLVmuLt8cMwUM879BmTdwIUfqP9Fy0cCdXgcEI40QMjs0zJbvbSLqGHuWqwJaIfPRE3MhhUHv_ndJyiVWbfhNH7uBxjxqkRpbVbvN_Z

buf, ret = client.Read(fd943, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd935, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8Z6s0SJ4ST_WWEhlr7NmsRMD9X82Weh8jcZqF_UcT0fjwCuxqBs") {
  panic("wrong data returned")
}


ret = client.Seek(fd940, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


buf, ret = client.Read(fd940, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "VmuLt8cMwUM879Bm") {
  panic("wrong data returned")
}


ret = client.Close(fd943)
if(ret != 0) {
  panic("close failed")
}


fd946 := client.Open("/BQuaEYmALY.txt", client.O_RDWR|client.O_CREATE)
if(fd946 < 0) {
  panic("open failed")
}


ret = client.Seek(fd945, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


ret = client.Close(fd946)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd938)
if(ret != 0) {
  panic("close failed")
}


fd947 := client.Open("/LsRInew_XM.txt", client.O_RDWR|client.O_CREATE)
if(fd947 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd945, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pYiNMJFBxgIa_iErKO3YfkuGxq0ei8lSeZ") {
  panic("wrong data returned")
}


fd948 := client.Open("/offJDM4Vh_.txt", client.O_RDWR|client.O_CREATE)
if(fd948 < 0) {
  panic("open failed")
}


fd949 := client.Open("/SgtN9yKcDb.txt", client.O_RDWR|client.O_CREATE)
if(fd949 < 0) {
  panic("open failed")
}

//fd state: (5) Asx4jUMA

ret = client.Write(fd934, []byte("ZbP_PJvN6X6qc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) Asx4jZbP_PJvN6X6qc

buf, ret = client.Read(fd935, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gAYImN2hcoH_387dW0z8uEX5ycMYFuIpVr4o7ej_akQjIVx8_0Pkzl8C") {
  panic("wrong data returned")
}


fd950 := client.Open("/1_pU8JJR0U.txt", client.O_RDWR|client.O_CREATE)
if(fd950 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd933, []byte("T3M4_76WaZ7qmoD9fpOj4GiBFGwy7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) T3M4_76WaZ7qmoD9fpOj4GiBFGwy7

fd951 := client.Open("/iGeqw2DPgJ.txt", client.O_RDWR|client.O_CREATE)
if(fd951 < 0) {
  panic("open failed")
}


ret = client.Close(fd940)
if(ret != 0) {
  panic("close failed")
}


fd952 := client.Open("/kIN_CKC2TK.txt", client.O_RDWR|client.O_CREATE)
if(fd952 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd937, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd933)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd942)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd947)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd951, []byte("NQApf5euFXp4yXpDQj41L2Lb7B57UPUuecIy7E0QVp42vnScWILpteskmra7vxbv7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) NQApf5euFXp4yXpDQj41L2Lb7B57UPUuecIy7E0QVp42vnScWILpteskmra7vxbv7

buf, ret = client.Read(fd945, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd953 := client.Open("/yX6wGc_UzJ.txt", client.O_RDWR|client.O_CREATE)
if(fd953 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd952, []byte("feXOfeqoVxwwd6gvDQ70nk_ogxMgGRXgo_ktHxFPuuUWaU5_NDxcEHWGR96Yue8mTJiw1KZjo7Na"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) feXOfeqoVxwwd6gvDQ70nk_ogxMgGRXgo_ktHxFPuuUWaU5_NDxcEHWGR96Yue8mTJiw1KZjo7Na

ret = client.Seek(fd950, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


buf, ret = client.Read(fd948, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd952)
if(ret != 0) {
  panic("close failed")
}


fd954 := client.Open("/zwxEZHcWOj.txt", client.O_RDWR|client.O_CREATE)
if(fd954 < 0) {
  panic("open failed")
}


ret = client.Seek(fd953, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (65) NQApf5euFXp4yXpDQj41L2Lb7B57UPUuecIy7E0QVp42vnScWILpteskmra7vxbv7

ret = client.Write(fd951, []byte("tut3Sm9DWb557Xndb3AxMnglLNxaBWXwRjs2C3t2AiK0SNXJMh3fuZdgHg_MhkZaPh9HYAl4ra_hSg1cF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (146) NQApf5euFXp4yXpDQj41L2Lb7B57UPUuecIy7E0QVp42vnScWILpteskmra7vxbv7tut3Sm9DWb557Xndb3AxMnglLNxaBWXwRjs2C3t2AiK0SNXJMh3fuZdgHg_MhkZaPh9HYAl4ra_hSg1cF

buf, ret = client.Read(fd937, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd945, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd955 := client.Open("/I4d0giEtWy.txt", client.O_RDWR|client.O_CREATE)
if(fd955 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd953, []byte("3Gefn9Bo4EyQwbr0QcRdumXRfJiAbKmfnvdoCBDLzB5Ip4HJOq34KIZrzqbtu4ZmY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) 3Gefn9Bo4EyQwbr0QcRdumXRfJiAbKmfnvdoCBDLzB5Ip4HJOq34KIZrzqbtu4ZmY

ret = client.Seek(fd951, 95, client.SEEK_SET)
if(ret != 95) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 95)
  panic("seek failed")
}

//fd state: (65) 3Gefn9Bo4EyQwbr0QcRdumXRfJiAbKmfnvdoCBDLzB5Ip4HJOq34KIZrzqbtu4ZmY

ret = client.Write(fd953, []byte("FOcSQ6cTLF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) 3Gefn9Bo4EyQwbr0QcRdumXRfJiAbKmfnvdoCBDLzB5Ip4HJOq34KIZrzqbtu4ZmYFOcSQ6cTLF
//fd state: (0) 

ret = client.Write(fd955, []byte("Lmqx8S328D6rPhv6en35r0lj_5_yc8w"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) Lmqx8S328D6rPhv6en35r0lj_5_yc8w

fd956 := client.Open("/C2X2B8VghN.txt", client.O_RDWR|client.O_CREATE)
if(fd956 < 0) {
  panic("open failed")
}


ret = client.Close(fd951)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd937, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (75) 3Gefn9Bo4EyQwbr0QcRdumXRfJiAbKmfnvdoCBDLzB5Ip4HJOq34KIZrzqbtu4ZmYFOcSQ6cTLF

ret = client.Write(fd953, []byte("JFqNnWthNAfaNALssTk5cvdRBcnuDoXwGD3hM2RwFmRFe761rPjKd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (128) 3Gefn9Bo4EyQwbr0QcRdumXRfJiAbKmfnvdoCBDLzB5Ip4HJOq34KIZrzqbtu4ZmYFOcSQ6cTLFJFqNnWthNAfaNALssTk5cvdRBcnuDoXwGD3hM2RwFmRFe761rPjKd

ret = client.Close(fd956)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd955, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


ret = client.Seek(fd935, 149, client.SEEK_SET)
if(ret != 149) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 149)
  panic("seek failed")
}


buf, ret = client.Read(fd955, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "r0lj_5_yc8w") {
  panic("wrong data returned")
}

//fd state: (0) PMo1oF2M9EEWm6plLqk1v0EoQxc4gBvJWV4KRKBwYqZOSbjmRocZ8RLw0VyiEg4CgHeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpnfIz0nNvJzCb0LD7rU20C0brf8x3ulLgKU54ZikIC8eCN049O6g8uXuQB4Va0zv1ChPHK1jcoRM0mC7Q7GBN5gMGfIFIJQx62sJliIVfmpr2_GfonfpHUqEx0xjBbdYcj8MudzGgYlc

ret = client.Write(fd949, []byte("HLFnmE6rEHw7lwi33W_jNLbtpWuhcX5aGDjww8MuT7gJXvezKce64"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) HLFnmE6rEHw7lwi33W_jNLbtpWuhcX5aGDjww8MuT7gJXvezKce64RLw0VyiEg4CgHeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpnfIz0nNvJzCb0LD7rU20C0brf8x3ulLgKU54ZikIC8eCN049O6g8uXuQB4Va0zv1ChPHK1jcoRM0mC7Q7GBN5gMGfIFIJQx62sJliIVfmpr2_GfonfpHUqEx0xjBbdYcj8MudzGgYlc
//fd state: (0) 

ret = client.Write(fd954, []byte("u6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) u6

buf, ret = client.Read(fd950, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "H0cbQy6EVebUDg8wbbtnOdy") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd948, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd955, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd935)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd953, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}


buf, ret = client.Read(fd948, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd957 := client.Open("/IKie96SW_E.txt", client.O_RDWR|client.O_CREATE)
if(fd957 < 0) {
  panic("open failed")
}


fd958 := client.Open("/ZLjr5ROzuJ.txt", client.O_RDWR|client.O_CREATE)
if(fd958 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd955, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd959 := client.Open("/SIYaGzyyUN.txt", client.O_RDWR|client.O_CREATE)
if(fd959 < 0) {
  panic("open failed")
}


ret = client.Close(fd948)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd934, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd960 := client.Open("/id33r0vGmj.txt", client.O_RDWR|client.O_CREATE)
if(fd960 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd949, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "RLw0VyiEg4CgHeE8uXpYheucydGV_CPUryjlAxbB37V5S82v5TsFSpYiNmr1F5fC50V2nYsoHwqjpnfIz0") {
  panic("wrong data returned")
}


ret = client.Close(fd953)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd945, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (26) 7ULH0cbQy6EVebUDg8wbbtnOdy

ret = client.Write(fd950, []byte("4R"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) 7ULH0cbQy6EVebUDg8wbbtnOdy4R

buf, ret = client.Read(fd934, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) aI6PxoC3yui0yPp3YSscLLVn1wOMKcOVKtC4WT1Mq55qaifUWYEaqw3NEQDQyFrhcPWymsHY1V0VSK

ret = client.Write(fd957, []byte("4uAK7gTQCjFRNWm2SCDAkBhgfnbb11SQnc6XDvEDlKU6KF1Ed7fVEKjwRvrs277TiXOhP9nRO8JomE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) 4uAK7gTQCjFRNWm2SCDAkBhgfnbb11SQnc6XDvEDlKU6KF1Ed7fVEKjwRvrs277TiXOhP9nRO8JomE
//fd state: (28) 7ULH0cbQy6EVebUDg8wbbtnOdy4R

ret = client.Write(fd950, []byte("Wn7OMylfEajZuD8LD9M0pumTgHbZ3K8Vd9AnZvqB7EpvzBCmPVxWsuKKAVJdN5lycNteFLzJ69Nd4_4BBEXZZWd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) 7ULH0cbQy6EVebUDg8wbbtnOdy4RWn7OMylfEajZuD8LD9M0pumTgHbZ3K8Vd9AnZvqB7EpvzBCmPVxWsuKKAVJdN5lycNteFLzJ69Nd4_4BBEXZZWd

fd961 := client.Open("/wJEXdnqHbD.txt", client.O_RDWR|client.O_CREATE)
if(fd961 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd957, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd961)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd958, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd957, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


ret = client.Close(fd957)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd950, 64, client.SEEK_SET)
if(ret != 64) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 64)
  panic("seek failed")
}


ret = client.Close(fd937)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd955)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd960)
if(ret != 0) {
  panic("close failed")
}

//fd state: (56) W_jJuOrSVCetJz1IgiI4m0pYiNMJFBxgIa_iErKO3YfkuGxq0ei8lSeZ

ret = client.Write(fd945, []byte("vah37z3v5irLKlKUpXidYaGxnNWkWWvG0VlghL84epPKOzF9zsfx8VE7_PcUqBRnzbRbPxceT3E4Ue_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) W_jJuOrSVCetJz1IgiI4m0pYiNMJFBxgIa_iErKO3YfkuGxq0ei8lSeZvah37z3v5irLKlKUpXidYaGxnNWkWWvG0VlghL84epPKOzF9zsfx8VE7_PcUqBRnzbRbPxceT3E4Ue_

ret = client.Close(fd950)
if(ret != 0) {
  panic("close failed")
}


fd962 := client.Open("/POUUp2IIUt.txt", client.O_RDWR|client.O_CREATE)
if(fd962 < 0) {
  panic("open failed")
}


fd963 := client.Open("/qS706qI7aE.txt", client.O_RDWR|client.O_CREATE)
if(fd963 < 0) {
  panic("open failed")
}


fd964 := client.Open("/boIbts5KMP.txt", client.O_RDWR|client.O_CREATE)
if(fd964 < 0) {
  panic("open failed")
}

//fd state: (2) u6

ret = client.Write(fd954, []byte("u3zg_Jicm3H6ueOhAb2xGAKbjtMiY7oP2WNycsS0p"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) u6u3zg_Jicm3H6ueOhAb2xGAKbjtMiY7oP2WNycsS0p

buf, ret = client.Read(fd954, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd964, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd949, 249, client.SEEK_SET)
if(ret != 249) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 249)
  panic("seek failed")
}


ret = client.Seek(fd964, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd965 := client.Open("/P6zKAKzRI9.txt", client.O_RDWR|client.O_CREATE)
if(fd965 < 0) {
  panic("open failed")
}


fd966 := client.Open("/_vdpUCIYrP.txt", client.O_RDWR|client.O_CREATE)
if(fd966 < 0) {
  panic("open failed")
}

//fd state: (0) DfKeJmPN1B7GooOlTrvqpYLGzryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOnlkdnk_IsnSZuVyQtxd2WHXKEKQYrqWC39NT8Uwnwj74oruCYzPSUwb97TLcpjeMqGKhbeKJTznAqbEthzIK4nRA9vPNgQUsRLUBDRvLtEUYH2dA0WDyF0vk

ret = client.Write(fd959, []byte("9Zsg356VmXjU80XQuwJq9Z46"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) 9Zsg356VmXjU80XQuwJq9Z46zryZyvrS8cWxEmBoRJjRdL_2ttuj8PdobNxcn_3wzjpEJ6y0F0UM6cTSq8S4_8KPgoZb_7biLOnlkdnk_IsnSZuVyQtxd2WHXKEKQYrqWC39NT8Uwnwj74oruCYzPSUwb97TLcpjeMqGKhbeKJTznAqbEthzIK4nRA9vPNgQUsRLUBDRvLtEUYH2dA0WDyF0vk

buf, ret = client.Read(fd949, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "x0xjBbdYcj8MudzGgYlc") {
  panic("wrong data returned")
}


fd967 := client.Open("/DWICQRTyV5.txt", client.O_RDWR|client.O_CREATE)
if(fd967 < 0) {
  panic("open failed")
}


ret = client.Seek(fd965, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd963, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd968 := client.Open("/zoKlXZrNEU.txt", client.O_RDWR|client.O_CREATE)
if(fd968 < 0) {
  panic("open failed")
}


ret = client.Seek(fd968, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd969 := client.Open("/r3k_UAG5DF.txt", client.O_RDWR|client.O_CREATE)
if(fd969 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd949, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd954, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd934, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd949, 158, client.SEEK_SET)
if(ret != 158) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 158)
  panic("seek failed")
}


buf, ret = client.Read(fd969, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd958, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd962, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd959, 195, client.SEEK_SET)
if(ret != 195) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 195)
  panic("seek failed")
}


buf, ret = client.Read(fd954, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd969, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd964, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd934)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd945, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd962)
if(ret != 0) {
  panic("close failed")
}


fd970 := client.Open("/vhi67tCvDp.txt", client.O_RDWR|client.O_CREATE)
if(fd970 < 0) {
  panic("open failed")
}


fd971 := client.Open("/uV1nPoDUtI.txt", client.O_RDWR|client.O_CREATE)
if(fd971 < 0) {
  panic("open failed")
}


fd972 := client.Open("/QHvsDv_2uV.txt", client.O_RDWR|client.O_CREATE)
if(fd972 < 0) {
  panic("open failed")
}


ret = client.Close(fd958)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd954)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd967)
if(ret != 0) {
  panic("close failed")
}


fd973 := client.Open("/62lJCMOXTv.txt", client.O_RDWR|client.O_CREATE)
if(fd973 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd964, []byte("fkdALUeiNQ0PckPkuMMCGcVZsm5xQ9MpDLATvuiLgJZ1Zvl4C3tolMktJHXlCZcszvsR19rAXwHVncAxRjSSX8iYsyaebol"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) fkdALUeiNQ0PckPkuMMCGcVZsm5xQ9MpDLATvuiLgJZ1Zvl4C3tolMktJHXlCZcszvsR19rAXwHVncAxRjSSX8iYsyaebol

ret = client.Seek(fd972, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd974 := client.Open("/JRYsh5HY6W.txt", client.O_RDWR|client.O_CREATE)
if(fd974 < 0) {
  panic("open failed")
}

//fd state: (0) mYs2m889NDnq2oo_PJHJATypzrtIQEKSeNu7RX9F6QV7XlYDsdajpgbaBktgVHijpJFmbk3V9zhvIpFL9dqH2z8IDVJwZj1mHJs7QmW_9M7LhS5Ox5JDhXGJRYQe0NDB8Jr946LyFosw5voP_G

ret = client.Write(fd974, []byte("4SwxFlpJporc_IeuLAK_a28PiBKXY0JeYZX5ub2BwwRvL8VTHeg1w"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) 4SwxFlpJporc_IeuLAK_a28PiBKXY0JeYZX5ub2BwwRvL8VTHeg1wgbaBktgVHijpJFmbk3V9zhvIpFL9dqH2z8IDVJwZj1mHJs7QmW_9M7LhS5Ox5JDhXGJRYQe0NDB8Jr946LyFosw5voP_G

ret = client.Seek(fd949, 222, client.SEEK_SET)
if(ret != 222) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 222)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd968, []byte("XqipMfkE1uoPrdCuIaj8MJO2uKJmd3m5cnqTedZWVOxvQ3qpkI57tMceXxM9E9EqH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) XqipMfkE1uoPrdCuIaj8MJO2uKJmd3m5cnqTedZWVOxvQ3qpkI57tMceXxM9E9EqH

buf, ret = client.Read(fd972, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd959)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd971)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd973, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3HewuwrjrFv3s26iaPtj4WzWbJN1g6aFz") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd949, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JQ") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd969, []byte("fkZrt_jSrdcL3DHzIeGQhrmsL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) fkZrt_jSrdcL3DHzIeGQhrmsL

buf, ret = client.Read(fd965, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (33) 3HewuwrjrFv3s26iaPtj4WzWbJN1g6aFz

ret = client.Write(fd973, []byte("IEFTEX9w_KW7JId5kP40zQA3CBekIdGXQrIUdNNt8uO4c"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) 3HewuwrjrFv3s26iaPtj4WzWbJN1g6aFzIEFTEX9w_KW7JId5kP40zQA3CBekIdGXQrIUdNNt8uO4c

fd975 := client.Open("/avTOA9vu6F.txt", client.O_RDWR|client.O_CREATE)
if(fd975 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd972, []byte("Nu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) Nu

ret = client.Seek(fd966, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd976 := client.Open("/1jA88lQxIi.txt", client.O_RDWR|client.O_CREATE)
if(fd976 < 0) {
  panic("open failed")
}

//fd state: (2) Nu

ret = client.Write(fd972, []byte("xkhlfVWRll1dRDsIFa4Faw0u6jh_VWXEgf9U1ULfdluk3fMMSiPlgKAWHwLvtPDIxWXcIsnfQ4yl326vgHmQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) NuxkhlfVWRll1dRDsIFa4Faw0u6jh_VWXEgf9U1ULfdluk3fMMSiPlgKAWHwLvtPDIxWXcIsnfQ4yl326vgHmQ

ret = client.Seek(fd972, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


fd977 := client.Open("/OlMsanXw7a.txt", client.O_RDWR|client.O_CREATE)
if(fd977 < 0) {
  panic("open failed")
}


fd978 := client.Open("/S5XuqQbVeb.txt", client.O_RDWR|client.O_CREATE)
if(fd978 < 0) {
  panic("open failed")
}


ret = client.Seek(fd965, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd977, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd963, []byte("sLG6qn_mgSsgBHXT5DDTeHtqnAKjuqntjiOYRyfMXmfiH1x91wSisQ9yw2i_nVelfPPme5g5yrpaD7gRU_AYTtU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) sLG6qn_mgSsgBHXT5DDTeHtqnAKjuqntjiOYRyfMXmfiH1x91wSisQ9yw2i_nVelfPPme5g5yrpaD7gRU_AYTtU
//fd state: (78) 3HewuwrjrFv3s26iaPtj4WzWbJN1g6aFzIEFTEX9w_KW7JId5kP40zQA3CBekIdGXQrIUdNNt8uO4c

ret = client.Write(fd973, []byte("auh_sAwHIJz1wu6dBkqVa0C_y3FLa6TGRTWNBYOUGtm1IItsP5YpkzuwvZLCOX6b6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (143) 3HewuwrjrFv3s26iaPtj4WzWbJN1g6aFzIEFTEX9w_KW7JId5kP40zQA3CBekIdGXQrIUdNNt8uO4cauh_sAwHIJz1wu6dBkqVa0C_y3FLa6TGRTWNBYOUGtm1IItsP5YpkzuwvZLCOX6b6
//fd state: (81) NuxkhlfVWRll1dRDsIFa4Faw0u6jh_VWXEgf9U1ULfdluk3fMMSiPlgKAWHwLvtPDIxWXcIsnfQ4yl326vgHmQ

ret = client.Write(fd972, []byte("kWVN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) NuxkhlfVWRll1dRDsIFa4Faw0u6jh_VWXEgf9U1ULfdluk3fMMSiPlgKAWHwLvtPDIxWXcIsnfQ4yl326kWVNQ

fd979 := client.Open("/OlMsanXw7a.txt", client.O_RDWR|client.O_CREATE)
if(fd979 < 0) {
  panic("open failed")
}


ret = client.Close(fd970)
if(ret != 0) {
  panic("close failed")
}


fd980 := client.Open("/QlSfrOOHqy.txt", client.O_RDWR|client.O_CREATE)
if(fd980 < 0) {
  panic("open failed")
}


ret = client.Close(fd966)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd978, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd949, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "x6") {
  panic("wrong data returned")
}


ret = client.Close(fd965)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd968)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd977)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd972, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd949)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd980, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nMfUfmOTLfsKpfUip7") {
  panic("wrong data returned")
}


ret = client.Close(fd980)
if(ret != 0) {
  panic("close failed")
}

//fd state: (25) fkZrt_jSrdcL3DHzIeGQhrmsL

ret = client.Write(fd969, []byte("1Wy9bvxdJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) fkZrt_jSrdcL3DHzIeGQhrmsL1Wy9bvxdJ

ret = client.Close(fd972)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd978, []byte("Tc5TTTBRQZuljW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) Tc5TTTBRQZuljW

ret = client.Seek(fd978, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (34) fkZrt_jSrdcL3DHzIeGQhrmsL1Wy9bvxdJ

ret = client.Write(fd969, []byte("B0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) fkZrt_jSrdcL3DHzIeGQhrmsL1Wy9bvxdJB0

ret = client.Seek(fd973, 123, client.SEEK_SET)
if(ret != 123) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 123)
  panic("seek failed")
}


buf, ret = client.Read(fd964, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd969, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (123) 3HewuwrjrFv3s26iaPtj4WzWbJN1g6aFzIEFTEX9w_KW7JId5kP40zQA3CBekIdGXQrIUdNNt8uO4cauh_sAwHIJz1wu6dBkqVa0C_y3FLa6TGRTWNBYOUGtm1IItsP5YpkzuwvZLCOX6b6

ret = client.Write(fd973, []byte("wscBQntwtGQt7NuQg3cJnACZUL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (149) 3HewuwrjrFv3s26iaPtj4WzWbJN1g6aFzIEFTEX9w_KW7JId5kP40zQA3CBekIdGXQrIUdNNt8uO4cauh_sAwHIJz1wu6dBkqVa0C_y3FLa6TGRTWNBYOUGtm1IwscBQntwtGQt7NuQg3cJnACZUL

buf, ret = client.Read(fd964, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd945, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


fd981 := client.Open("/mhq5vt2bFc.txt", client.O_RDWR|client.O_CREATE)
if(fd981 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd963, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd969)
if(ret != 0) {
  panic("close failed")
}

//fd state: (95) fkdALUeiNQ0PckPkuMMCGcVZsm5xQ9MpDLATvuiLgJZ1Zvl4C3tolMktJHXlCZcszvsR19rAXwHVncAxRjSSX8iYsyaebol

ret = client.Write(fd964, []byte("I8NNAK1snyR91r8BaVrHwRPtrQNjP41_oLOLYi9m7ymR5LBESP85CfndX2K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) fkdALUeiNQ0PckPkuMMCGcVZsm5xQ9MpDLATvuiLgJZ1Zvl4C3tolMktJHXlCZcszvsR19rAXwHVncAxRjSSX8iYsyaebolI8NNAK1snyR91r8BaVrHwRPtrQNjP41_oLOLYi9m7ymR5LBESP85CfndX2K

ret = client.Seek(fd945, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


buf, ret = client.Read(fd981, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd974, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


fd982 := client.Open("/x3feW4_Qsd.txt", client.O_RDWR|client.O_CREATE)
if(fd982 < 0) {
  panic("open failed")
}


ret = client.Close(fd978)
if(ret != 0) {
  panic("close failed")
}


fd983 := client.Open("/H71v73_fdb.txt", client.O_RDWR|client.O_CREATE)
if(fd983 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd982, []byte("K9iQieTtRiZIyiCCLUMueHnb6Cgiw2mBiIne30P_MP1JKEKWBipDr7j3TrTsyYEY1xWjuSj6MWBz_0HLM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) K9iQieTtRiZIyiCCLUMueHnb6Cgiw2mBiIne30P_MP1JKEKWBipDr7j3TrTsyYEY1xWjuSj6MWBz_0HLM
//fd state: (87) sLG6qn_mgSsgBHXT5DDTeHtqnAKjuqntjiOYRyfMXmfiH1x91wSisQ9yw2i_nVelfPPme5g5yrpaD7gRU_AYTtU

ret = client.Write(fd963, []byte("3od7lMLdr2waanf95_GHUTZrzxBN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) sLG6qn_mgSsgBHXT5DDTeHtqnAKjuqntjiOYRyfMXmfiH1x91wSisQ9yw2i_nVelfPPme5g5yrpaD7gRU_AYTtU3od7lMLdr2waanf95_GHUTZrzxBN

fd984 := client.Open("/f4k1qvMKPz.txt", client.O_RDWR|client.O_CREATE)
if(fd984 < 0) {
  panic("open failed")
}

//fd state: (81) K9iQieTtRiZIyiCCLUMueHnb6Cgiw2mBiIne30P_MP1JKEKWBipDr7j3TrTsyYEY1xWjuSj6MWBz_0HLM

ret = client.Write(fd982, []byte("T8_SMc392tNuqjVkBxjRj7SP5UJJB9SPXp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) K9iQieTtRiZIyiCCLUMueHnb6Cgiw2mBiIne30P_MP1JKEKWBipDr7j3TrTsyYEY1xWjuSj6MWBz_0HLMT8_SMc392tNuqjVkBxjRj7SP5UJJB9SPXp
//fd state: (115) K9iQieTtRiZIyiCCLUMueHnb6Cgiw2mBiIne30P_MP1JKEKWBipDr7j3TrTsyYEY1xWjuSj6MWBz_0HLMT8_SMc392tNuqjVkBxjRj7SP5UJJB9SPXp

ret = client.Write(fd982, []byte("1S8CMnyqGMw0nFa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) K9iQieTtRiZIyiCCLUMueHnb6Cgiw2mBiIne30P_MP1JKEKWBipDr7j3TrTsyYEY1xWjuSj6MWBz_0HLMT8_SMc392tNuqjVkBxjRj7SP5UJJB9SPXp1S8CMnyqGMw0nFa

buf, ret = client.Read(fd979, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd985 := client.Open("/6XoL2LuxTf.txt", client.O_RDWR|client.O_CREATE)
if(fd985 < 0) {
  panic("open failed")
}


ret = client.Close(fd973)
if(ret != 0) {
  panic("close failed")
}


fd986 := client.Open("/xtYmEEvAxH.txt", client.O_RDWR|client.O_CREATE)
if(fd986 < 0) {
  panic("open failed")
}


fd987 := client.Open("/eUVkxN2kF2.txt", client.O_RDWR|client.O_CREATE)
if(fd987 < 0) {
  panic("open failed")
}

//fd state: (0) _FRYnOm1HTymqLTM4S809ezU5CFtL9GK3iTs0Qj_ZnWeAmAXgo1zod6CNEvfUnET8K46y_YhNonfHWKZmMu8TZ21hbugQZBJGXvSX1H6T8JbDOQ

ret = client.Write(fd986, []byte("nJDuakxB0_vLkXMrI2GwQeU_8CI3eHxW9A0A1cXCsmDL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) nJDuakxB0_vLkXMrI2GwQeU_8CI3eHxW9A0A1cXCsmDLAmAXgo1zod6CNEvfUnET8K46y_YhNonfHWKZmMu8TZ21hbugQZBJGXvSX1H6T8JbDOQ

buf, ret = client.Read(fd983, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd987)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd985, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd988 := client.Open("/b28U0Ezj36.txt", client.O_RDWR|client.O_CREATE)
if(fd988 < 0) {
  panic("open failed")
}


ret = client.Seek(fd985, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd989 := client.Open("/MFvyYPutcX.txt", client.O_RDWR|client.O_CREATE)
if(fd989 < 0) {
  panic("open failed")
}


ret = client.Seek(fd963, 111, client.SEEK_SET)
if(ret != 111) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 111)
  panic("seek failed")
}


ret = client.Seek(fd974, 95, client.SEEK_SET)
if(ret != 95) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 95)
  panic("seek failed")
}


fd990 := client.Open("/OZOyt6VkFz.txt", client.O_RDWR|client.O_CREATE)
if(fd990 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd981, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd991 := client.Open("/RH71tgp37R.txt", client.O_RDWR|client.O_CREATE)
if(fd991 < 0) {
  panic("open failed")
}


ret = client.Seek(fd981, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (10) W_jJuOrSVCetJz1IgiI4m0pYiNMJFBxgIa_iErKO3YfkuGxq0ei8lSeZvah37z3v5irLKlKUpXidYaGxnNWkWWvG0VlghL84epPKOzF9zsfx8VE7_PcUqBRnzbRbPxceT3E4Ue_

ret = client.Write(fd945, []byte("cB6zBLJhM3gZkE5CKVYM5szf20NhG4T5IpltjqusyAfrWfAQCB1GU3TVD7k3vRE3ql3aGo9u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) W_jJuOrSVCcB6zBLJhM3gZkE5CKVYM5szf20NhG4T5IpltjqusyAfrWfAQCB1GU3TVD7k3vRE3ql3aGo9uWkWWvG0VlghL84epPKOzF9zsfx8VE7_PcUqBRnzbRbPxceT3E4Ue_

ret = client.Close(fd986)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd975, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd992 := client.Open("/aqTjeF62CH.txt", client.O_RDWR|client.O_CREATE)
if(fd992 < 0) {
  panic("open failed")
}


ret = client.Close(fd982)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd983)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd963, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


ret = client.Close(fd974)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd988, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd981, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd976, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd991)
if(ret != 0) {
  panic("close failed")
}

//fd state: (43) sLG6qn_mgSsgBHXT5DDTeHtqnAKjuqntjiOYRyfMXmfiH1x91wSisQ9yw2i_nVelfPPme5g5yrpaD7gRU_AYTtU3od7lMLdr2waanf95_GHUTZrzxBN

ret = client.Write(fd963, []byte("mW2U2eZF0L9Q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) sLG6qn_mgSsgBHXT5DDTeHtqnAKjuqntjiOYRyfMXmfmW2U2eZF0L9Qyw2i_nVelfPPme5g5yrpaD7gRU_AYTtU3od7lMLdr2waanf95_GHUTZrzxBN

ret = client.Close(fd963)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd976, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd993 := client.Open("/XgqlK6vDL2.txt", client.O_RDWR|client.O_CREATE)
if(fd993 < 0) {
  panic("open failed")
}


ret = client.Close(fd984)
if(ret != 0) {
  panic("close failed")
}


fd994 := client.Open("/QWehDPrfor.txt", client.O_RDWR|client.O_CREATE)
if(fd994 < 0) {
  panic("open failed")
}


fd995 := client.Open("/hqQ4RRTmCR.txt", client.O_RDWR|client.O_CREATE)
if(fd995 < 0) {
  panic("open failed")
}


fd996 := client.Open("/LcGiEGoTe0.txt", client.O_RDWR|client.O_CREATE)
if(fd996 < 0) {
  panic("open failed")
}


ret = client.Seek(fd990, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}


ret = client.Seek(fd981, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd994, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


fd997 := client.Open("/1y8xHZOuxw.txt", client.O_RDWR|client.O_CREATE)
if(fd997 < 0) {
  panic("open failed")
}

//fd state: (50) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMzAjGbvyumrw1qG0ujLv6r5zNYybw2b7g9GSdX3Y374li_bGkJg41dlHb7hq2YsIHgMFVzEz4L2FwwsJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu

ret = client.Write(fd990, []byte("MMJbRAoOHx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMMJbRAoOHxrw1qG0ujLv6r5zNYybw2b7g9GSdX3Y374li_bGkJg41dlHb7hq2YsIHgMFVzEz4L2FwwsJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu

buf, ret = client.Read(fd997, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd981, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd993, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd988, []byte("Q5RfYkne20dGDauWjSw4_UW9MpzZ9Rw6Q0GgZN3baitxeG4Idp52kAh8HWsw446tzIpfiP13PXz85sGRI9WG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) Q5RfYkne20dGDauWjSw4_UW9MpzZ9Rw6Q0GgZN3baitxeG4Idp52kAh8HWsw446tzIpfiP13PXz85sGRI9WG

ret = client.Seek(fd993, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


buf, ret = client.Read(fd981, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd998 := client.Open("/bb3zIBjBgh.txt", client.O_RDWR|client.O_CREATE)
if(fd998 < 0) {
  panic("open failed")
}


fd999 := client.Open("/CvNkTigzOc.txt", client.O_RDWR|client.O_CREATE)
if(fd999 < 0) {
  panic("open failed")
}

//fd state: (35) Z3TbnTz73vbYO_byxHDQSxUWfd2bOrGDtotHIv

ret = client.Write(fd994, []byte("KsUFB7zGNhSYkTlXmBsBIYGTsw8TT79qTrpA2K8LYH1sAs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) Z3TbnTz73vbYO_byxHDQSxUWfd2bOrGDtotKsUFB7zGNhSYkTlXmBsBIYGTsw8TT79qTrpA2K8LYH1sAs
//fd state: (0) 

ret = client.Write(fd985, []byte("3YtzpZpUGPiUOOQB4Ddvo13VPq1wYAA6Xl8JLCIiflzPCL66QdfkXX3LNhhQ14gcDxkChWkM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) 3YtzpZpUGPiUOOQB4Ddvo13VPq1wYAA6Xl8JLCIiflzPCL66QdfkXX3LNhhQ14gcDxkChWkM

buf, ret = client.Read(fd975, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd975, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd989)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd998, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd995, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


buf, ret = client.Read(fd981, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd997, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd999)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd997)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd993, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd996)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd981, []byte("VkHWNYauG57nVq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) VkHWNYauG57nVq

ret = client.Seek(fd979, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd990, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rw1qG0ujLv6r") {
  panic("wrong data returned")
}

//fd state: (72) 3YtzpZpUGPiUOOQB4Ddvo13VPq1wYAA6Xl8JLCIiflzPCL66QdfkXX3LNhhQ14gcDxkChWkM

ret = client.Write(fd985, []byte("ZY4S5E6y0nUmZgRzqdaOlveSY0BPDtGrYOLOs9VrRo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) 3YtzpZpUGPiUOOQB4Ddvo13VPq1wYAA6Xl8JLCIiflzPCL66QdfkXX3LNhhQ14gcDxkChWkMZY4S5E6y0nUmZgRzqdaOlveSY0BPDtGrYOLOs9VrRo

ret = client.Seek(fd992, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd995, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Evw") {
  panic("wrong data returned")
}


ret = client.Seek(fd994, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd995, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (84) Q5RfYkne20dGDauWjSw4_UW9MpzZ9Rw6Q0GgZN3baitxeG4Idp52kAh8HWsw446tzIpfiP13PXz85sGRI9WG

ret = client.Write(fd988, []byte("YOqtUhPXr9fPpvmwYPnSUHa2Rcq1veefaA0pP_miGmbnBOtuHEyxv4kYSBCNclG3C3Hwr89Jb0J1j7DdnzQbWvLqRo1ICCjn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (180) Q5RfYkne20dGDauWjSw4_UW9MpzZ9Rw6Q0GgZN3baitxeG4Idp52kAh8HWsw446tzIpfiP13PXz85sGRI9WGYOqtUhPXr9fPpvmwYPnSUHa2Rcq1veefaA0pP_miGmbnBOtuHEyxv4kYSBCNclG3C3Hwr89Jb0J1j7DdnzQbWvLqRo1ICCjn

ret = client.Seek(fd995, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Seek(fd998, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (82) W_jJuOrSVCcB6zBLJhM3gZkE5CKVYM5szf20NhG4T5IpltjqusyAfrWfAQCB1GU3TVD7k3vRE3ql3aGo9uWkWWvG0VlghL84epPKOzF9zsfx8VE7_PcUqBRnzbRbPxceT3E4Ue_

ret = client.Write(fd945, []byte("8KiCYlJlyZ1HiNDYKzGQID76WIFJ_uTyqGcjOdbaq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) W_jJuOrSVCcB6zBLJhM3gZkE5CKVYM5szf20NhG4T5IpltjqusyAfrWfAQCB1GU3TVD7k3vRE3ql3aGo9u8KiCYlJlyZ1HiNDYKzGQID76WIFJ_uTyqGcjOdbaqbPxceT3E4Ue_

fd1000 := client.Open("/FopR58AGxt.txt", client.O_RDWR|client.O_CREATE)
if(fd1000 < 0) {
  panic("open failed")
}

//fd state: (72) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMMJbRAoOHxrw1qG0ujLv6r5zNYybw2b7g9GSdX3Y374li_bGkJg41dlHb7hq2YsIHgMFVzEz4L2FwwsJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu

ret = client.Write(fd990, []byte("Y95m5AcZVLJGwXJN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) VobYcNxJs3xr7FxVnLapsI8b4nRJNBlnNu7l5_8fGZEpe6kS2iMMJbRAoOHxrw1qG0ujLv6rY95m5AcZVLJGwXJN3Y374li_bGkJg41dlHb7hq2YsIHgMFVzEz4L2FwwsJNxQFAO0kT2uHBADsOZP8W70lI1k5heA6pu

ret = client.Seek(fd981, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd1000)
if(ret != 0) {
  panic("close failed")
}

//fd state: (114) 3YtzpZpUGPiUOOQB4Ddvo13VPq1wYAA6Xl8JLCIiflzPCL66QdfkXX3LNhhQ14gcDxkChWkMZY4S5E6y0nUmZgRzqdaOlveSY0BPDtGrYOLOs9VrRo

ret = client.Write(fd985, []byte("94FdGEVwdvg7Lm9nJ0BOcTluV5CFiKDf4QG01ui40gPF_Xt9Fm5MAreFtD7CTQihXiVkmyl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (185) 3YtzpZpUGPiUOOQB4Ddvo13VPq1wYAA6Xl8JLCIiflzPCL66QdfkXX3LNhhQ14gcDxkChWkMZY4S5E6y0nUmZgRzqdaOlveSY0BPDtGrYOLOs9VrRo94FdGEVwdvg7Lm9nJ0BOcTluV5CFiKDf4QG01ui40gPF_Xt9Fm5MAreFtD7CTQihXiVkmyl

ret = client.Close(fd985)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd993)
if(ret != 0) {
  panic("close failed")
}


fd1001 := client.Open("/hVemHBzD6Z.txt", client.O_RDWR|client.O_CREATE)
if(fd1001 < 0) {
  panic("open failed")
}

//fd state: (12) VkHWNYauG57nVq

ret = client.Write(fd981, []byte("MYSQ4RVdnkJOxpXqcVdmCvaCBxPDCs1lwUI5vFZz5JqSCnxxLuRLdcBBnpLJRamxGQB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) VkHWNYauG57nMYSQ4RVdnkJOxpXqcVdmCvaCBxPDCs1lwUI5vFZz5JqSCnxxLuRLdcBBnpLJRamxGQB

buf, ret = client.Read(fd1001, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd945)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd994, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


buf, ret = client.Read(fd964, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1002 := client.Open("/raavVeJ6eA.txt", client.O_RDWR|client.O_CREATE)
if(fd1002 < 0) {
  panic("open failed")
}


fd1003 := client.Open("/4SF7BMPmZu.txt", client.O_RDWR|client.O_CREATE)
if(fd1003 < 0) {
  panic("open failed")
}

//fd state: (0) LsX2upRewHuGPKOI6otReaIgSHqD10MRV8hKSrxLqSQdNbnRugD0gFnKmqlDGiQqEIM1LvGx1lrD1W0YiEX6Haf0O

ret = client.Write(fd1002, []byte("BBubaR93z596axIy2N6EIPg4G3_1chsZ8g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) BBubaR93z596axIy2N6EIPg4G3_1chsZ8ghKSrxLqSQdNbnRugD0gFnKmqlDGiQqEIM1LvGx1lrD1W0YiEX6Haf0O

buf, ret = client.Read(fd964, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1003, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1004 := client.Open("/SJ6nAqZ9nA.txt", client.O_RDWR|client.O_CREATE)
if(fd1004 < 0) {
  panic("open failed")
}


ret = client.Close(fd981)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd994)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1003, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd998)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd1003, []byte("PdpPUJGQhMo2LcpcWdFLF19VgdkQ2tT5gN6X7LI3HkbC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) PdpPUJGQhMo2LcpcWdFLF19VgdkQ2tT5gN6X7LI3HkbC
//fd state: (0) 

ret = client.Write(fd1004, []byte("giEx0ofWx0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) giEx0ofWx0

ret = client.Close(fd1001)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1002, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hKSrxLqSQdNbnRugD0gFnKmqlDGiQqEIM1LvGx1lrD1W0YiEX6Haf0O") {
  panic("wrong data returned")
}


fd1005 := client.Open("/plv8EuK8Z3.txt", client.O_RDWR|client.O_CREATE)
if(fd1005 < 0) {
  panic("open failed")
}


ret = client.Close(fd1005)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1003, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1006 := client.Open("/6v5z4HEYRW.txt", client.O_RDWR|client.O_CREATE)
if(fd1006 < 0) {
  panic("open failed")
}


ret = client.Close(fd995)
if(ret != 0) {
  panic("close failed")
}


fd1007 := client.Open("/PiOdP6CVcO.txt", client.O_RDWR|client.O_CREATE)
if(fd1007 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd976, []byte("8pck8kb9sXylhjT5idd6JlQOt8Wc4F9r0mJq60I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) 8pck8kb9sXylhjT5idd6JlQOt8Wc4F9r0mJq60I

ret = client.Close(fd990)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd975)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd992)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd988, 169, client.SEEK_SET)
if(ret != 169) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 169)
  panic("seek failed")
}


buf, ret = client.Read(fd1006, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) BYoRAOQLAbIs0AZRvZ9_JJvFleqwQFxV4mu3U3PudQCJRIfrfDpdsGoWKgAVC1FQ_Tuo2LDbl

ret = client.Write(fd1007, []byte("aM0DeexqsLi5Vsqyo5hwy3RMUy7XMdC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) aM0DeexqsLi5Vsqyo5hwy3RMUy7XMdCV4mu3U3PudQCJRIfrfDpdsGoWKgAVC1FQ_Tuo2LDbl

ret = client.Seek(fd1006, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd1002, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


buf, ret = client.Read(fd1007, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "V4mu3U3PudQCJRIfrfDpdsGoWKgAVC") {
  panic("wrong data returned")
}


ret = client.Close(fd964)
if(ret != 0) {
  panic("close failed")
}

//fd state: (10) giEx0ofWx0

ret = client.Write(fd1004, []byte("Za9sfjGp5Z_MJp34vGJG2nQMas646owhxoPUrKsJhd651kxezuM14EKxlhT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) giEx0ofWx0Za9sfjGp5Z_MJp34vGJG2nQMas646owhxoPUrKsJhd651kxezuM14EKxlhT
//fd state: (39) 8pck8kb9sXylhjT5idd6JlQOt8Wc4F9r0mJq60I

ret = client.Write(fd976, []byte("OKUbSlZ8CmiABceN2mjYr5ycTBDyBuZT8jPGGhy75UAE1xwDggfvukAiBEw9I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) 8pck8kb9sXylhjT5idd6JlQOt8Wc4F9r0mJq60IOKUbSlZ8CmiABceN2mjYr5ycTBDyBuZT8jPGGhy75UAE1xwDggfvukAiBEw9I

fd1008 := client.Open("/685Nb0NdyX.txt", client.O_RDWR|client.O_CREATE)
if(fd1008 < 0) {
  panic("open failed")
}


ret = client.Seek(fd988, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}


ret = client.Seek(fd1004, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Close(fd1007)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1006, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1009 := client.Open("/9yJ4B8r9Jr.txt", client.O_RDWR|client.O_CREATE)
if(fd1009 < 0) {
  panic("open failed")
}

//fd state: (0) W8H61C5wBk4OZBhwIjpsotZnscZthO3t5b9iVlLlIhmfRyYTjtbsWE8gB7vb6nQk9XZvXO4EjHTfkmQ6doahe3utAbzMIJrJSx7iZrAj0hk9upSv29n9VksY3XpKKxDKhvrh19IT1W1Z4PlFJVQVnOqLQKsw52B3w9yDSOcyDSnXTtyS2jqHszsGiuJCKFJssqXw4Jb2zd

ret = client.Write(fd1009, []byte("tjpJZTNuph4cC3zy1F1AbJFO3kUeOf_rYqCt0aakZ3NcVfxeXKsUMRry5PJFDw2waVrXVVNrrHLJTd6GNqrCVbLRpVx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) tjpJZTNuph4cC3zy1F1AbJFO3kUeOf_rYqCt0aakZ3NcVfxeXKsUMRry5PJFDw2waVrXVVNrrHLJTd6GNqrCVbLRpVxMIJrJSx7iZrAj0hk9upSv29n9VksY3XpKKxDKhvrh19IT1W1Z4PlFJVQVnOqLQKsw52B3w9yDSOcyDSnXTtyS2jqHszsGiuJCKFJssqXw4Jb2zd

buf, ret = client.Read(fd1006, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd1003)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1008, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (100) 8pck8kb9sXylhjT5idd6JlQOt8Wc4F9r0mJq60IOKUbSlZ8CmiABceN2mjYr5ycTBDyBuZT8jPGGhy75UAE1xwDggfvukAiBEw9I

ret = client.Write(fd976, []byte("wPuMeaFXVdHecUdYz3gXhA_bOVTnhWICsTmnt9WbUJcISut9in2WrRkvGrNfIU8x44QXjYsKc4NuVLqN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (180) 8pck8kb9sXylhjT5idd6JlQOt8Wc4F9r0mJq60IOKUbSlZ8CmiABceN2mjYr5ycTBDyBuZT8jPGGhy75UAE1xwDggfvukAiBEw9IwPuMeaFXVdHecUdYz3gXhA_bOVTnhWICsTmnt9WbUJcISut9in2WrRkvGrNfIU8x44QXjYsKc4NuVLqN

ret = client.Seek(fd1006, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd1002, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "EIPg4G3_1chsZ8ghKSrxLqSQdNbnRugD0gFnKmqlDGiQqEIM1LvGx1lrD1W0YiEX6Haf0O") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1006, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (89) BBubaR93z596axIy2N6EIPg4G3_1chsZ8ghKSrxLqSQdNbnRugD0gFnKmqlDGiQqEIM1LvGx1lrD1W0YiEX6Haf0O

ret = client.Write(fd1002, []byte("Fi94lLr2PnA6lBKBOGA3nEpNEyProc6YqLBuAkJ3D5d_2dhHo6IaCv9YgIgmfFbbO8LG1uRWblTs0x"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (167) BBubaR93z596axIy2N6EIPg4G3_1chsZ8ghKSrxLqSQdNbnRugD0gFnKmqlDGiQqEIM1LvGx1lrD1W0YiEX6Haf0OFi94lLr2PnA6lBKBOGA3nEpNEyProc6YqLBuAkJ3D5d_2dhHo6IaCv9YgIgmfFbbO8LG1uRWblTs0x

ret = client.Close(fd1004)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1006, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1010 := client.Open("/GpshBQBhVD.txt", client.O_RDWR|client.O_CREATE)
if(fd1010 < 0) {
  panic("open failed")
}


ret = client.Close(fd1008)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd976)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1006, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1011 := client.Open("/EE2wAtM0EI.txt", client.O_RDWR|client.O_CREATE)
if(fd1011 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd979, []byte("V8P8jZ7MGAuxR0Se0mOSeypl2mBzXm8TYtx03wZN86cxpynuBK2XzENlv15fXTuh9Y_wsr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) V8P8jZ7MGAuxR0Se0mOSeypl2mBzXm8TYtx03wZN86cxpynuBK2XzENlv15fXTuh9Y_wsr

ret = client.Close(fd1009)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1011, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd979, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}

//fd state: (167) BBubaR93z596axIy2N6EIPg4G3_1chsZ8ghKSrxLqSQdNbnRugD0gFnKmqlDGiQqEIM1LvGx1lrD1W0YiEX6Haf0OFi94lLr2PnA6lBKBOGA3nEpNEyProc6YqLBuAkJ3D5d_2dhHo6IaCv9YgIgmfFbbO8LG1uRWblTs0x

ret = client.Write(fd1002, []byte("aGI5D1BS4IKnBpHx5zpi0lzh6riHzq2LQiYvjE4AAR4hftpo2z8an3e1rvsHZTqb2aFAj6QAI0pZFj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (245) BBubaR93z596axIy2N6EIPg4G3_1chsZ8ghKSrxLqSQdNbnRugD0gFnKmqlDGiQqEIM1LvGx1lrD1W0YiEX6Haf0OFi94lLr2PnA6lBKBOGA3nEpNEyProc6YqLBuAkJ3D5d_2dhHo6IaCv9YgIgmfFbbO8LG1uRWblTs0xaGI5D1BS4IKnBpHx5zpi0lzh6riHzq2LQiYvjE4AAR4hftpo2z8an3e1rvsHZTqb2aFAj6QAI0pZFj

ret = client.Seek(fd1002, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}


buf, ret = client.Read(fd1010, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd1002, 132, client.SEEK_SET)
if(ret != 132) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 132)
  panic("seek failed")
}


ret = client.Seek(fd988, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


ret = client.Seek(fd1002, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


fd1012 := client.Open("/dyKKqLJYGZ.txt", client.O_RDWR|client.O_CREATE)
if(fd1012 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1010, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd988, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


ret = client.Close(fd1012)
if(ret != 0) {
  panic("close failed")
}


fd1013 := client.Open("/PiOdP6CVcO.txt", client.O_RDWR|client.O_CREATE)
if(fd1013 < 0) {
  panic("open failed")
}


ret = client.Close(fd1011)
if(ret != 0) {
  panic("close failed")
}


fd1014 := client.Open("/R1VJ6U8gR7.txt", client.O_RDWR|client.O_CREATE)
if(fd1014 < 0) {
  panic("open failed")
}


ret = client.Close(fd1010)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd1013)
if(ret != 0) {
  panic("close failed")
}


fd1015 := client.Open("/rGYU992m68.txt", client.O_RDWR|client.O_CREATE)
if(fd1015 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1002, 229, client.SEEK_SET)
if(ret != 229) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 229)
  panic("seek failed")
}

//fd state: (52) Q5RfYkne20dGDauWjSw4_UW9MpzZ9Rw6Q0GgZN3baitxeG4Idp52kAh8HWsw446tzIpfiP13PXz85sGRI9WGYOqtUhPXr9fPpvmwYPnSUHa2Rcq1veefaA0pP_miGmbnBOtuHEyxv4kYSBCNclG3C3Hwr89Jb0J1j7DdnzQbWvLqRo1ICCjn

ret = client.Write(fd988, []byte("d_DFCru_9jyTLvnmJWE6dremAXLe14zlKDXugLc_gHgCm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) Q5RfYkne20dGDauWjSw4_UW9MpzZ9Rw6Q0GgZN3baitxeG4Idp52d_DFCru_9jyTLvnmJWE6dremAXLe14zlKDXugLc_gHgCmvmwYPnSUHa2Rcq1veefaA0pP_miGmbnBOtuHEyxv4kYSBCNclG3C3Hwr89Jb0J1j7DdnzQbWvLqRo1ICCjn

buf, ret = client.Read(fd1002, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qb2aFAj") {
  panic("wrong data returned")
}


fd1016 := client.Open("/YBc2F9uh1_.txt", client.O_RDWR|client.O_CREATE)
if(fd1016 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1015, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd979)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1016, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "deBKtsqqjnIcLEI5HFqbdxN5EiBbXhagxQP9s3gH1TbpZJKt5FQmgVDkgQWFxoNcDW__Ev_JPR7RX3Y1X_v1KDQxC8") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd1017 := client.Open("/tZiK0BxhAL.txt", client.O_RDWR|client.O_CREATE)
if(fd1017 < 0) {
  panic("open failed")
}


fd1018 := client.Open("/Q6r48G5r_c.txt", client.O_RDWR|client.O_CREATE)
if(fd1018 < 0) {
  panic("open failed")
}


ret = client.Close(fd1002)
if(ret != 0) {
  panic("close failed")
}


fd1019 := client.Open("/6XoL2LuxTf.txt", client.O_RDWR|client.O_CREATE)
if(fd1019 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1017, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd1006, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd1016)
if(ret != 0) {
  panic("close failed")
}


fd1020 := client.Open("/xdFaEq7KDB.txt", client.O_RDWR|client.O_CREATE)
if(fd1020 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd988, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vmwYPnSUHa2Rcq1veefaA0pP_miGmbnBOtuHEyxv4kYSBCNclG3C") {
  panic("wrong data returned")
}

//fd state: (0) 3YtzpZpUGPiUOOQB4Ddvo13VPq1wYAA6Xl8JLCIiflzPCL66QdfkXX3LNhhQ14gcDxkChWkMZY4S5E6y0nUmZgRzqdaOlveSY0BPDtGrYOLOs9VrRo94FdGEVwdvg7Lm9nJ0BOcTluV5CFiKDf4QG01ui40gPF_Xt9Fm5MAreFtD7CTQihXiVkmyl

ret = client.Write(fd1019, []byte("1OBgmzMgUJfr0NdLGUfcuHBImzlkNg3aC483KlzwpbUdhVbb9sLLytkpVSY8YqcVHkdpcXlkpeJinTi9KwhGwW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) 1OBgmzMgUJfr0NdLGUfcuHBImzlkNg3aC483KlzwpbUdhVbb9sLLytkpVSY8YqcVHkdpcXlkpeJinTi9KwhGwWRzqdaOlveSY0BPDtGrYOLOs9VrRo94FdGEVwdvg7Lm9nJ0BOcTluV5CFiKDf4QG01ui40gPF_Xt9Fm5MAreFtD7CTQihXiVkmyl

ret = client.Close(fd1017)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1014, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


fd1021 := client.Open("/_xt2CaFERA.txt", client.O_RDWR|client.O_CREATE)
if(fd1021 < 0) {
  panic("open failed")
}


fd1022 := client.Open("/rkU9KrdLrD.txt", client.O_RDWR|client.O_CREATE)
if(fd1022 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1022, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}


ret = client.Seek(fd1015, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd988)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1014, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Seek(fd1020, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (2) BuNEBWgVJYhQTONLKuQzia50E7cMh2hn9xl6nnGrI8L6QprNkpwWm6PQ8g9ZneWQO_tU4f4DLhCTwyeuM8XzCbddbBqktrnHSKlioRvFQKPI4ejJhdEp1j8V2uBnYpN30IaqlK1

ret = client.Write(fd1014, []byte("pPboFBdgT4pg_KZ7EwGL8YeWi8xr2FawqXVqKwB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) BupPboFBdgT4pg_KZ7EwGL8YeWi8xr2FawqXVqKwB8L6QprNkpwWm6PQ8g9ZneWQO_tU4f4DLhCTwyeuM8XzCbddbBqktrnHSKlioRvFQKPI4ejJhdEp1j8V2uBnYpN30IaqlK1

ret = client.Close(fd1019)
if(ret != 0) {
  panic("close failed")
}


fd1023 := client.Open("/zvfpB2u_jD.txt", client.O_RDWR|client.O_CREATE)
if(fd1023 < 0) {
  panic("open failed")
}


ret = client.Close(fd1018)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1014, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8L6QprNkpwWm6PQ8g9ZneWQO_tU4f4DLhCTwyeuM8XzCbddbBqktrnHSKlioRvFQKPI4ejJhdEp1j8V2uBnYpN3") {
  panic("wrong data returned")
}


ret = client.Seek(fd1015, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd1022, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


fd1024 := client.Open("/vphjpP0lT0.txt", client.O_RDWR|client.O_CREATE)
if(fd1024 < 0) {
  panic("open failed")
}


ret = client.Close(fd1006)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd1021)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1024, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1022, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "c7U0b5mwKDHY3W1Ct0ntvw80ezIMgI") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1024, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd1014)
if(ret != 0) {
  panic("close failed")
}


fd1025 := client.Open("/p90S07207w.txt", client.O_RDWR|client.O_CREATE)
if(fd1025 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1025, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd1023, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xyHFttu") {
  panic("wrong data returned")
}


ret = client.Seek(fd1022, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}

//fd state: (17) djyGYSLV3sDgD8t3ga95c7U0b5mwKDHY3W1Ct0ntvw80ezIMgI

ret = client.Write(fd1022, []byte("tjsy5uHq5cnZ57h0TpHpQXzMwaVfi8gKOMBVV4i8gAN4Tld63g_6hogDITeWO3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) djyGYSLV3sDgD8t3gtjsy5uHq5cnZ57h0TpHpQXzMwaVfi8gKOMBVV4i8gAN4Tld63g_6hogDITeWO3

fd1026 := client.Open("/JzZk_6uCpc.txt", client.O_RDWR|client.O_CREATE)
if(fd1026 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd1024, []byte("8njtAVk5HM1tZ5KrJRgVuxiRp8g7OWsDf7AhEosbepFiUEJvNqOltc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) 8njtAVk5HM1tZ5KrJRgVuxiRp8g7OWsDf7AhEosbepFiUEJvNqOltc
//fd state: (79) djyGYSLV3sDgD8t3gtjsy5uHq5cnZ57h0TpHpQXzMwaVfi8gKOMBVV4i8gAN4Tld63g_6hogDITeWO3

ret = client.Write(fd1022, []byte("fVSPs5raHb_S3H3vRjxq28irBGbmeK3YqpCxv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) djyGYSLV3sDgD8t3gtjsy5uHq5cnZ57h0TpHpQXzMwaVfi8gKOMBVV4i8gAN4Tld63g_6hogDITeWO3fVSPs5raHb_S3H3vRjxq28irBGbmeK3YqpCxv

fd1027 := client.Open("/q8LarXbP95.txt", client.O_RDWR|client.O_CREATE)
if(fd1027 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd1025, []byte("v5RyfoNgAdFyuD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) v5RyfoNgAdFyuD

ret = client.Seek(fd1026, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


buf, ret = client.Read(fd1027, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd1023)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1027, 98, client.SEEK_SET)
if(ret != 98) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 98)
  panic("seek failed")
}


fd1028 := client.Open("/0tyzNGz888.txt", client.O_RDWR|client.O_CREATE)
if(fd1028 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1015, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd1015, []byte("neHRCIcmT0c"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) neHRCIcmT0c

ret = client.Close(fd1015)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd1022)
if(ret != 0) {
  panic("close failed")
}


fd1029 := client.Open("/U3v_bL_lmK.txt", client.O_RDWR|client.O_CREATE)
if(fd1029 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd1020, []byte("UiyhbJg2GHuHaoAYhYILf2w2h22USPRM71S29_QH4QyP_SmQkbZddO6wdH8Dl3sKxKlqC5VCyXistzSgMc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) UiyhbJg2GHuHaoAYhYILf2w2h22USPRM71S29_QH4QyP_SmQkbZddO6wdH8Dl3sKxKlqC5VCyXistzSgMc

buf, ret = client.Read(fd1028, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "efc52EnkFVRmOZUhCY6H_iK2o8LbuvQdqCpP0oWW12Dd64XxtY1LHnYdssF46cxeeXN2k4YJ8L_OwCB") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1028, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd1020, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}


ret = client.Seek(fd1026, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}

//fd state: (14) v5RyfoNgAdFyuD

ret = client.Write(fd1025, []byte("xVGrh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) v5RyfoNgAdFyuDxVGrh

ret = client.Seek(fd1025, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


fd1030 := client.Open("/offJDM4Vh_.txt", client.O_RDWR|client.O_CREATE)
if(fd1030 < 0) {
  panic("open failed")
}


ret = client.Close(fd1028)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1020, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "SgMc") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1025, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "RyfoNgAdFyuDxVGrh") {
  panic("wrong data returned")
}


fd1031 := client.Open("/0_jDN0GKli.txt", client.O_RDWR|client.O_CREATE)
if(fd1031 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1026, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rj865nUJRBMCtOmdgnFI5us8lV2eZwrFoNmw15qnnBcOSh8HHJ3Vno5qigZkilnNPA9slsFXfjyA") {
  panic("wrong data returned")
}


ret = client.Close(fd1031)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1029, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1032 := client.Open("/nG1YGGFGtK.txt", client.O_RDWR|client.O_CREATE)
if(fd1032 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1020, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


buf, ret = client.Read(fd1025, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1033 := client.Open("/FBwyYmw985.txt", client.O_RDWR|client.O_CREATE)
if(fd1033 < 0) {
  panic("open failed")
}


fd1034 := client.Open("/skSFr__AzL.txt", client.O_RDWR|client.O_CREATE)
if(fd1034 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1020, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "29_QH4QyP_SmQkbZddO6wdH8Dl3sKxKlqC") {
  panic("wrong data returned")
}

//fd state: (0) NpET0RHOjmK2VZDzYfQ5ySY3xAC8gdDFakTYE

ret = client.Write(fd1033, []byte("qzLkNPN01cz0Vc2Mtu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) qzLkNPN01cz0Vc2MtuQ5ySY3xAC8gdDFakTYE

ret = client.Close(fd1029)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1034, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd1027)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1034, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1035 := client.Open("/WGTJdUdKLz.txt", client.O_RDWR|client.O_CREATE)
if(fd1035 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1025, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


fd1036 := client.Open("/ja3oo3k0R9.txt", client.O_RDWR|client.O_CREATE)
if(fd1036 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1026, 83, client.SEEK_SET)
if(ret != 83) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 83)
  panic("seek failed")
}


fd1037 := client.Open("/DMPyQLTMXr.txt", client.O_RDWR|client.O_CREATE)
if(fd1037 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1034, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd1037, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd1035, []byte("bfeFRfApaDYKBi9wIL0t9MRXk8YzYTvO05NNoznFGINe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) bfeFRfApaDYKBi9wIL0t9MRXk8YzYTvO05NNoznFGINe

buf, ret = client.Read(fd1037, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1038 := client.Open("/d9joSjnqtY.txt", client.O_RDWR|client.O_CREATE)
if(fd1038 < 0) {
  panic("open failed")
}


ret = client.Close(fd1034)
if(ret != 0) {
  panic("close failed")
}

//fd state: (44) bfeFRfApaDYKBi9wIL0t9MRXk8YzYTvO05NNoznFGINe

ret = client.Write(fd1035, []byte("ItY65Kd5tWP_uStojsnW7LowlkoEu1_Zro07qT2q8Nb7uE23mbwq_2uUkdWklnrlF4Lar_9jGMwLpAW2AU_Hf9nEL0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) bfeFRfApaDYKBi9wIL0t9MRXk8YzYTvO05NNoznFGINeItY65Kd5tWP_uStojsnW7LowlkoEu1_Zro07qT2q8Nb7uE23mbwq_2uUkdWklnrlF4Lar_9jGMwLpAW2AU_Hf9nEL0

buf, ret = client.Read(fd1033, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Q5ySY3xAC8gdDFakTYE") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1035, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1039 := client.Open("/DBm2gV4Brr.txt", client.O_RDWR|client.O_CREATE)
if(fd1039 < 0) {
  panic("open failed")
}


fd1040 := client.Open("/wA277eurDT.txt", client.O_RDWR|client.O_CREATE)
if(fd1040 < 0) {
  panic("open failed")
}


ret = client.Close(fd1037)
if(ret != 0) {
  panic("close failed")
}


fd1041 := client.Open("/bWUvYr_axB.txt", client.O_RDWR|client.O_CREATE)
if(fd1041 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1041, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1042 := client.Open("/oXkOXESJ1G.txt", client.O_RDWR|client.O_CREATE)
if(fd1042 < 0) {
  panic("open failed")
}

//fd state: (37) qzLkNPN01cz0Vc2MtuQ5ySY3xAC8gdDFakTYE

ret = client.Write(fd1033, []byte("OcSkyKd21iJDDciXK1kzycDcPtd4n250ZtfrBZbZJdmWmM0bU_Cz95xXW68QJwyUlo2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) qzLkNPN01cz0Vc2MtuQ5ySY3xAC8gdDFakTYEOcSkyKd21iJDDciXK1kzycDcPtd4n250ZtfrBZbZJdmWmM0bU_Cz95xXW68QJwyUlo2
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}
