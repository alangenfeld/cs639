
package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
)

func main(){
	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)
	var buf []byte
	buf = buf
	var ret int
	ret = ret

fd1 := client.Open("/oIsR5Rk83H.txt", client.O_RDWR|client.O_CREATE)
if(fd1 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd1, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd1, []byte("6WvRUgVZCOQmAtKHTSXr5HXSDpJdrFKfvmCwEg2uyuFs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) 6WvRUgVZCOQmAtKHTSXr5HXSDpJdrFKfvmCwEg2uyuFs

ret = client.Seek(fd1, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


ret = client.Seek(fd1, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


ret = client.Close(fd1)
if(ret != 0) {
  panic("close failed")
}


fd2 := client.Open("/iVIJjg2elX.txt", client.O_RDWR|client.O_CREATE)
if(fd2 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd2, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd2, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd3 := client.Open("/2JrRL5XTCk.txt", client.O_RDWR|client.O_CREATE)
if(fd3 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd2, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd4 := client.Open("/jaSNhtIk5k.txt", client.O_RDWR|client.O_CREATE)
if(fd4 < 0) {
  panic("open failed")
}


ret = client.Close(fd4)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd2)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd3, []byte("FpBOmAxJXZiOuRzATDSuOohUoSVFR1YM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) FpBOmAxJXZiOuRzATDSuOohUoSVFR1YM

buf, ret = client.Read(fd3, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd3, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


buf, ret = client.Read(fd3, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JXZiOuRz") {
  panic("wrong data returned")
}

//fd state: (15) FpBOmAxJXZiOuRzATDSuOohUoSVFR1YM

ret = client.Write(fd3, []byte("Fbtk2hl3geUWlgDTIo_8TAWvSmRVynNxasTH2atoRA_C2Y1ISabFhD762dVfmIE9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) FpBOmAxJXZiOuRzFbtk2hl3geUWlgDTIo_8TAWvSmRVynNxasTH2atoRA_C2Y1ISabFhD762dVfmIE9

ret = client.Close(fd3)
if(ret != 0) {
  panic("close failed")
}


fd5 := client.Open("/4cQqDopMxJ.txt", client.O_RDWR|client.O_CREATE)
if(fd5 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd5, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd5, []byte("xGsH7Re5gAlyYX2tHxhvb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) xGsH7Re5gAlyYX2tHxhvb
//fd state: (21) xGsH7Re5gAlyYX2tHxhvb

ret = client.Write(fd5, []byte("u5pEjchrmUg2jgYdZf0BE7F5XYa3CYFL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) xGsH7Re5gAlyYX2tHxhvbu5pEjchrmUg2jgYdZf0BE7F5XYa3CYFL

ret = client.Close(fd5)
if(ret != 0) {
  panic("close failed")
}


fd6 := client.Open("/4cQqDopMxJ.txt", client.O_RDWR|client.O_CREATE)
if(fd6 < 0) {
  panic("open failed")
}

//fd state: (0) xGsH7Re5gAlyYX2tHxhvbu5pEjchrmUg2jgYdZf0BE7F5XYa3CYFL

ret = client.Write(fd6, []byte("KCPdD3hBtZztTxnEDBH6JZrvcM5Hmxhw0vpXOjEmmJg6Qi6Z1dIvOnonk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) KCPdD3hBtZztTxnEDBH6JZrvcM5Hmxhw0vpXOjEmmJg6Qi6Z1dIvOnonk

fd7 := client.Open("/2VZzqpP7_6.txt", client.O_RDWR|client.O_CREATE)
if(fd7 < 0) {
  panic("open failed")
}


fd8 := client.Open("/oIsR5Rk83H.txt", client.O_RDWR|client.O_CREATE)
if(fd8 < 0) {
  panic("open failed")
}


fd9 := client.Open("/jaSNhtIk5k.txt", client.O_RDWR|client.O_CREATE)
if(fd9 < 0) {
  panic("open failed")
}


ret = client.Seek(fd7, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd10 := client.Open("/KVp2ABEVFd.txt", client.O_RDWR|client.O_CREATE)
if(fd10 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd10, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd10)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd6, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd8, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "6WvRUgVZCOQmAtKHTSXr5HXSDpJdrFKf") {
  panic("wrong data returned")
}

//fd state: (32) 6WvRUgVZCOQmAtKHTSXr5HXSDpJdrFKfvmCwEg2uyuFs

ret = client.Write(fd8, []byte("3NFAMgAIgwA0MWyAP0IHjmzKoHJJMCwmUz41Wnupj5QZKAY5PSWJ6vlmXeLIXL2OfC9k9J1SUkc2RmOrKglL3nOQTqtrwu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (126) 6WvRUgVZCOQmAtKHTSXr5HXSDpJdrFKf3NFAMgAIgwA0MWyAP0IHjmzKoHJJMCwmUz41Wnupj5QZKAY5PSWJ6vlmXeLIXL2OfC9k9J1SUkc2RmOrKglL3nOQTqtrwu

fd11 := client.Open("/jaSNhtIk5k.txt", client.O_RDWR|client.O_CREATE)
if(fd11 < 0) {
  panic("open failed")
}


ret = client.Close(fd7)
if(ret != 0) {
  panic("close failed")
}


fd12 := client.Open("/jaSNhtIk5k.txt", client.O_RDWR|client.O_CREATE)
if(fd12 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd12, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd13 := client.Open("/RUpXDwp6BU.txt", client.O_RDWR|client.O_CREATE)
if(fd13 < 0) {
  panic("open failed")
}


fd14 := client.Open("/CzVjEdkRbn.txt", client.O_RDWR|client.O_CREATE)
if(fd14 < 0) {
  panic("open failed")
}


ret = client.Seek(fd9, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (57) KCPdD3hBtZztTxnEDBH6JZrvcM5Hmxhw0vpXOjEmmJg6Qi6Z1dIvOnonk

ret = client.Write(fd6, []byte("fK8ZA_vkNxdvNY3Vw6SNrCXel79elj46m5IkvBkpUh2zOzCuVI1gv1nc3QiQyfHYMbSmlKnRiOLQxf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) KCPdD3hBtZztTxnEDBH6JZrvcM5Hmxhw0vpXOjEmmJg6Qi6Z1dIvOnonkfK8ZA_vkNxdvNY3Vw6SNrCXel79elj46m5IkvBkpUh2zOzCuVI1gv1nc3QiQyfHYMbSmlKnRiOLQxf
//fd state: (126) 6WvRUgVZCOQmAtKHTSXr5HXSDpJdrFKf3NFAMgAIgwA0MWyAP0IHjmzKoHJJMCwmUz41Wnupj5QZKAY5PSWJ6vlmXeLIXL2OfC9k9J1SUkc2RmOrKglL3nOQTqtrwu

ret = client.Write(fd8, []byte("aSq6h_FgSa3LUpdllwUCzVztiPEjalnUgc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (160) 6WvRUgVZCOQmAtKHTSXr5HXSDpJdrFKf3NFAMgAIgwA0MWyAP0IHjmzKoHJJMCwmUz41Wnupj5QZKAY5PSWJ6vlmXeLIXL2OfC9k9J1SUkc2RmOrKglL3nOQTqtrwuaSq6h_FgSa3LUpdllwUCzVztiPEjalnUgc

ret = client.Close(fd8)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd9)
if(ret != 0) {
  panic("close failed")
}


fd15 := client.Open("/2JrRL5XTCk.txt", client.O_RDWR|client.O_CREATE)
if(fd15 < 0) {
  panic("open failed")
}


ret = client.Seek(fd6, 71, client.SEEK_SET)
if(ret != 71) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 71)
  panic("seek failed")
}


ret = client.Close(fd15)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd11)
if(ret != 0) {
  panic("close failed")
}


fd16 := client.Open("/_laoVJpAG9.txt", client.O_RDWR|client.O_CREATE)
if(fd16 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd12, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd6, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


fd17 := client.Open("/3v2LLkR1Z4.txt", client.O_RDWR|client.O_CREATE)
if(fd17 < 0) {
  panic("open failed")
}


fd18 := client.Open("/rcJYn3q9Iv.txt", client.O_RDWR|client.O_CREATE)
if(fd18 < 0) {
  panic("open failed")
}


ret = client.Seek(fd16, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd14, []byte("V4p42BuHl7auv6IcbLZFdUrpUGOzpD9qmUhLGU9JdnpxzL6ZMAvLhZuZ8n4tRZHsEtLF4Lnhg4Byy2dBtx2zUUu36LupvXob"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) V4p42BuHl7auv6IcbLZFdUrpUGOzpD9qmUhLGU9JdnpxzL6ZMAvLhZuZ8n4tRZHsEtLF4Lnhg4Byy2dBtx2zUUu36LupvXob

ret = client.Close(fd14)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd17)
if(ret != 0) {
  panic("close failed")
}


fd19 := client.Open("/rcJYn3q9Iv.txt", client.O_RDWR|client.O_CREATE)
if(fd19 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd12, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd16)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd6, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "SNrC") {
  panic("wrong data returned")
}


ret = client.Seek(fd19, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd20 := client.Open("/RUpXDwp6BU.txt", client.O_RDWR|client.O_CREATE)
if(fd20 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd19, []byte("jzi0YjKyxcuJixr4GpMVaAEBi88PF4gK2_RMJAm1HmlafZUJBebkl80_MellMEgRjpy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) jzi0YjKyxcuJixr4GpMVaAEBi88PF4gK2_RMJAm1HmlafZUJBebkl80_MellMEgRjpy

buf, ret = client.Read(fd19, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd13, []byte("LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdpDuJqCfdHtmEO8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdpDuJqCfdHtmEO8

buf, ret = client.Read(fd6, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Xel79elj46m5IkvBkpUh2") {
  panic("wrong data returned")
}


ret = client.Close(fd6)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd19, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


fd21 := client.Open("/rcJYn3q9Iv.txt", client.O_RDWR|client.O_CREATE)
if(fd21 < 0) {
  panic("open failed")
}


fd22 := client.Open("/0KnXHjRLMh.txt", client.O_RDWR|client.O_CREATE)
if(fd22 < 0) {
  panic("open failed")
}


ret = client.Close(fd22)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd13, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd12)
if(ret != 0) {
  panic("close failed")
}

//fd state: (64) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdpDuJqCfdHtmEO8

ret = client.Write(fd13, []byte("o3uCIk4yiwa64J33D4DjXlFi6I4d7K8eHzx_bSPPNfoa5VobxHbyCLzAcmj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdpDuJqCfdHtmEO8o3uCIk4yiwa64J33D4DjXlFi6I4d7K8eHzx_bSPPNfoa5VobxHbyCLzAcmj

buf, ret = client.Read(fd18, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jzi0YjKyxcuJixr4GpMVaAEBi88PF4gK2") {
  panic("wrong data returned")
}

//fd state: (123) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdpDuJqCfdHtmEO8o3uCIk4yiwa64J33D4DjXlFi6I4d7K8eHzx_bSPPNfoa5VobxHbyCLzAcmj

ret = client.Write(fd13, []byte("kV_ZoWJOydFOLiNPcd0n7HXxquLfUcfadMN4tHVhAYsIaXOX7IQd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (175) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdpDuJqCfdHtmEO8o3uCIk4yiwa64J33D4DjXlFi6I4d7K8eHzx_bSPPNfoa5VobxHbyCLzAcmjkV_ZoWJOydFOLiNPcd0n7HXxquLfUcfadMN4tHVhAYsIaXOX7IQd

fd23 := client.Open("/hIqrxvPj2h.txt", client.O_RDWR|client.O_CREATE)
if(fd23 < 0) {
  panic("open failed")
}


ret = client.Seek(fd23, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd20, 128, client.SEEK_SET)
if(ret != 128) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 128)
  panic("seek failed")
}


ret = client.Close(fd13)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd23)
if(ret != 0) {
  panic("close failed")
}

//fd state: (39) jzi0YjKyxcuJixr4GpMVaAEBi88PF4gK2_RMJAm1HmlafZUJBebkl80_MellMEgRjpy

ret = client.Write(fd19, []byte("h7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) jzi0YjKyxcuJixr4GpMVaAEBi88PF4gK2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qw
//fd state: (0) jzi0YjKyxcuJixr4GpMVaAEBi88PF4gK2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qw

ret = client.Write(fd21, []byte("BcviVK8npvmqdA8hdNm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) BcviVK8npvmqdA8hdNmVaAEBi88PF4gK2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qw

buf, ret = client.Read(fd19, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd21)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd20, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}


fd24 := client.Open("/fT8o_szxtr.txt", client.O_RDWR|client.O_CREATE)
if(fd24 < 0) {
  panic("open failed")
}


ret = client.Close(fd24)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd18)
if(ret != 0) {
  panic("close failed")
}

//fd state: (50) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdpDuJqCfdHtmEO8o3uCIk4yiwa64J33D4DjXlFi6I4d7K8eHzx_bSPPNfoa5VobxHbyCLzAcmjkV_ZoWJOydFOLiNPcd0n7HXxquLfUcfadMN4tHVhAYsIaXOX7IQd

ret = client.Write(fd20, []byte("civdOVZq4rdB1sB5pX9IbDDuBFWDE8pUobt4kP1MVu1k"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdcivdOVZq4rdB1sB5pX9IbDDuBFWDE8pUobt4kP1MVu1k8eHzx_bSPPNfoa5VobxHbyCLzAcmjkV_ZoWJOydFOLiNPcd0n7HXxquLfUcfadMN4tHVhAYsIaXOX7IQd
//fd state: (95) BcviVK8npvmqdA8hdNmVaAEBi88PF4gK2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qw

ret = client.Write(fd19, []byte("d4bzTLnBwq5zVIUChZE4q8fEYfUUf6re6g5Enjkt_M4ZctMGqU2d19DtWGUVKSHhW89Gq9b5IqoAO9goSzE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (178) BcviVK8npvmqdA8hdNmVaAEBi88PF4gK2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qwd4bzTLnBwq5zVIUChZE4q8fEYfUUf6re6g5Enjkt_M4ZctMGqU2d19DtWGUVKSHhW89Gq9b5IqoAO9goSzE

ret = client.Seek(fd19, 135, client.SEEK_SET)
if(ret != 135) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 135)
  panic("seek failed")
}

//fd state: (135) BcviVK8npvmqdA8hdNmVaAEBi88PF4gK2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qwd4bzTLnBwq5zVIUChZE4q8fEYfUUf6re6g5Enjkt_M4ZctMGqU2d19DtWGUVKSHhW89Gq9b5IqoAO9goSzE

ret = client.Write(fd19, []byte("vnVtfQGreDAI1XPcrQu_7IWcxj339rRH0Why3tWFS9iCvROX7iRBJhw1CimzlvOxIGoTV20Yy8n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (210) BcviVK8npvmqdA8hdNmVaAEBi88PF4gK2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qwd4bzTLnBwq5zVIUChZE4q8fEYfUUf6re6g5EnjktvnVtfQGreDAI1XPcrQu_7IWcxj339rRH0Why3tWFS9iCvROX7iRBJhw1CimzlvOxIGoTV20Yy8n
//fd state: (94) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdcivdOVZq4rdB1sB5pX9IbDDuBFWDE8pUobt4kP1MVu1k8eHzx_bSPPNfoa5VobxHbyCLzAcmjkV_ZoWJOydFOLiNPcd0n7HXxquLfUcfadMN4tHVhAYsIaXOX7IQd

ret = client.Write(fd20, []byte("y7bh985unRNtAjmESBjboBckjpZQGjr4YVSDWNaJv9t7q9RAj4i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (145) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdcivdOVZq4rdB1sB5pX9IbDDuBFWDE8pUobt4kP1MVu1ky7bh985unRNtAjmESBjboBckjpZQGjr4YVSDWNaJv9t7q9RAj4iXxquLfUcfadMN4tHVhAYsIaXOX7IQd

fd25 := client.Open("/xlkZ2l0fT6.txt", client.O_RDWR|client.O_CREATE)
if(fd25 < 0) {
  panic("open failed")
}


ret = client.Close(fd20)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd19, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Seek(fd25, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd19, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "K2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qwd4bzTLnBwq5zVIUChZE4q") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd25, []byte("5CVzwhMK5HCqzUUeS1_DA0BItnkR6McdCVpdngvIRdocd5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) 5CVzwhMK5HCqzUUeS1_DA0BItnkR6McdCVpdngvIRdocd5

fd26 := client.Open("/2VZzqpP7_6.txt", client.O_RDWR|client.O_CREATE)
if(fd26 < 0) {
  panic("open failed")
}


fd27 := client.Open("/3v2LLkR1Z4.txt", client.O_RDWR|client.O_CREATE)
if(fd27 < 0) {
  panic("open failed")
}

//fd state: (116) BcviVK8npvmqdA8hdNmVaAEBi88PF4gK2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qwd4bzTLnBwq5zVIUChZE4q8fEYfUUf6re6g5EnjktvnVtfQGreDAI1XPcrQu_7IWcxj339rRH0Why3tWFS9iCvROX7iRBJhw1CimzlvOxIGoTV20Yy8n

ret = client.Write(fd19, []byte("gTVOCbh0rjJTPmjO8Lh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) BcviVK8npvmqdA8hdNmVaAEBi88PF4gK2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qwd4bzTLnBwq5zVIUChZE4qgTVOCbh0rjJTPmjO8LhvnVtfQGreDAI1XPcrQu_7IWcxj339rRH0Why3tWFS9iCvROX7iRBJhw1CimzlvOxIGoTV20Yy8n

ret = client.Close(fd25)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd26, []byte("8TWzKa3FMBvW6ZezH4TdgYnnDZtNjj9tdgu2C1Iw5LTA2FwZnH887"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) 8TWzKa3FMBvW6ZezH4TdgYnnDZtNjj9tdgu2C1Iw5LTA2FwZnH887

fd28 := client.Open("/jaSNhtIk5k.txt", client.O_RDWR|client.O_CREATE)
if(fd28 < 0) {
  panic("open failed")
}


ret = client.Seek(fd27, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd19)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd28, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd26)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd27, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd27)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd28, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd29 := client.Open("/4mPDBgaYe1.txt", client.O_RDWR|client.O_CREATE)
if(fd29 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd28, []byte("plUEqOv1Z8rdJXXONXbrOx_JHziDqqSXcKLHunI0B85JG1sfI2H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) plUEqOv1Z8rdJXXONXbrOx_JHziDqqSXcKLHunI0B85JG1sfI2H

buf, ret = client.Read(fd28, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd29)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd28)
if(ret != 0) {
  panic("close failed")
}


fd30 := client.Open("/hIqrxvPj2h.txt", client.O_RDWR|client.O_CREATE)
if(fd30 < 0) {
  panic("open failed")
}


fd31 := client.Open("/jEJT1wijAu.txt", client.O_RDWR|client.O_CREATE)
if(fd31 < 0) {
  panic("open failed")
}


fd32 := client.Open("/bCLvtka1QB.txt", client.O_RDWR|client.O_CREATE)
if(fd32 < 0) {
  panic("open failed")
}


fd33 := client.Open("/LZcAyv_xCm.txt", client.O_RDWR|client.O_CREATE)
if(fd33 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd31, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd30)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd32, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd34 := client.Open("/D7WnAKqAcn.txt", client.O_RDWR|client.O_CREATE)
if(fd34 < 0) {
  panic("open failed")
}


fd35 := client.Open("/LZcAyv_xCm.txt", client.O_RDWR|client.O_CREATE)
if(fd35 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd35, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd36 := client.Open("/nFxpEbR3hh.txt", client.O_RDWR|client.O_CREATE)
if(fd36 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd36, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd34, 96)
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


fd37 := client.Open("/_laoVJpAG9.txt", client.O_RDWR|client.O_CREATE)
if(fd37 < 0) {
  panic("open failed")
}


fd38 := client.Open("/k4d87VpxAe.txt", client.O_RDWR|client.O_CREATE)
if(fd38 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd33, []byte("fmUr3TA7uRrYlwqxvMkTh0Cljk2t8KAQb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) fmUr3TA7uRrYlwqxvMkTh0Cljk2t8KAQb

ret = client.Close(fd35)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd33, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


buf, ret = client.Read(fd34, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd36, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (2) fmUr3TA7uRrYlwqxvMkTh0Cljk2t8KAQb

ret = client.Write(fd33, []byte("pftJInZqnKnAARukBubzk7ZFGPtEVvYfSoEkFwGekPTWTG9K6v9yGmoarjev8LoGUfdx0jJ1XLo1xVKerzZcrSOhWKcWeKq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) fmpftJInZqnKnAARukBubzk7ZFGPtEVvYfSoEkFwGekPTWTG9K6v9yGmoarjev8LoGUfdx0jJ1XLo1xVKerzZcrSOhWKcWeKq

fd39 := client.Open("/L3a7VdOTLE.txt", client.O_RDWR|client.O_CREATE)
if(fd39 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd31, []byte("ev7_fh1MUOr5IrM8oYnIs0C56_m6cR2mZCGDWBt9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) ev7_fh1MUOr5IrM8oYnIs0C56_m6cR2mZCGDWBt9

fd40 := client.Open("/nFxpEbR3hh.txt", client.O_RDWR|client.O_CREATE)
if(fd40 < 0) {
  panic("open failed")
}


ret = client.Close(fd39)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd40, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd36)
if(ret != 0) {
  panic("close failed")
}


fd41 := client.Open("/iVIJjg2elX.txt", client.O_RDWR|client.O_CREATE)
if(fd41 < 0) {
  panic("open failed")
}


ret = client.Seek(fd31, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd38)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd33, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd31)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd37, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd41, []byte("4nPhEckHIAgdnTK7yRnZh4_2g7Pp8WMlqTBaxnUrhFElqRMlJMr4yCLti2_zg1dF3YNqjhvCuSlnGwUg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) 4nPhEckHIAgdnTK7yRnZh4_2g7Pp8WMlqTBaxnUrhFElqRMlJMr4yCLti2_zg1dF3YNqjhvCuSlnGwUg

fd42 := client.Open("/Q7Imfuqduy.txt", client.O_RDWR|client.O_CREATE)
if(fd42 < 0) {
  panic("open failed")
}


ret = client.Seek(fd33, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


fd43 := client.Open("/hIqrxvPj2h.txt", client.O_RDWR|client.O_CREATE)
if(fd43 < 0) {
  panic("open failed")
}


ret = client.Close(fd41)
if(ret != 0) {
  panic("close failed")
}

//fd state: (32) fmpftJInZqnKnAARukBubzk7ZFGPtEVvYfSoEkFwGekPTWTG9K6v9yGmoarjev8LoGUfdx0jJ1XLo1xVKerzZcrSOhWKcWeKq

ret = client.Write(fd33, []byte("flJpGQdkNc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) fmpftJInZqnKnAARukBubzk7ZFGPtEVvflJpGQdkNckPTWTG9K6v9yGmoarjev8LoGUfdx0jJ1XLo1xVKerzZcrSOhWKcWeKq

ret = client.Close(fd34)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd42, []byte("QloNu8PD5bFWZEedHwGeoA44aCr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) QloNu8PD5bFWZEedHwGeoA44aCr

ret = client.Close(fd42)
if(ret != 0) {
  panic("close failed")
}

//fd state: (42) fmpftJInZqnKnAARukBubzk7ZFGPtEVvflJpGQdkNckPTWTG9K6v9yGmoarjev8LoGUfdx0jJ1XLo1xVKerzZcrSOhWKcWeKq

ret = client.Write(fd33, []byte("lOxzccPIoNvZbiLctYYz1zjrVGOcQ6zt2pcUFT7Vv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) fmpftJInZqnKnAARukBubzk7ZFGPtEVvflJpGQdkNclOxzccPIoNvZbiLctYYz1zjrVGOcQ6zt2pcUFT7VvzZcrSOhWKcWeKq

ret = client.Close(fd40)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd43, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd37, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd43, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd33)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd43, []byte("5rsHVNr4h9COEQz9Ej8QTMJZZQh0xZH27HCj97Qblw0sv9HxH38rYxAmUa1706AMJkov3RI0WEUN5naAgQZwE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) 5rsHVNr4h9COEQz9Ej8QTMJZZQh0xZH27HCj97Qblw0sv9HxH38rYxAmUa1706AMJkov3RI0WEUN5naAgQZwE

fd44 := client.Open("/0KnXHjRLMh.txt", client.O_RDWR|client.O_CREATE)
if(fd44 < 0) {
  panic("open failed")
}


fd45 := client.Open("/nFxpEbR3hh.txt", client.O_RDWR|client.O_CREATE)
if(fd45 < 0) {
  panic("open failed")
}


ret = client.Seek(fd44, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd44, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd37, []byte("A9awL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) A9awL

fd46 := client.Open("/NDaLK8HBXg.txt", client.O_RDWR|client.O_CREATE)
if(fd46 < 0) {
  panic("open failed")
}


ret = client.Close(fd44)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd45, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd37, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd45, []byte("WWJXXW7hq63_g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) WWJXXW7hq63_g

ret = client.Seek(fd45, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


fd47 := client.Open("/RUpXDwp6BU.txt", client.O_RDWR|client.O_CREATE)
if(fd47 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd37, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd48 := client.Open("/DbkC5q9e6v.txt", client.O_RDWR|client.O_CREATE)
if(fd48 < 0) {
  panic("open failed")
}


fd49 := client.Open("/RUpXDwp6BU.txt", client.O_RDWR|client.O_CREATE)
if(fd49 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd49, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdcivdOVZ") {
  panic("wrong data returned")
}


fd50 := client.Open("/_fKtkmMG2H.txt", client.O_RDWR|client.O_CREATE)
if(fd50 < 0) {
  panic("open failed")
}


ret = client.Close(fd50)
if(ret != 0) {
  panic("close failed")
}


fd51 := client.Open("/DNkbwkXRNz.txt", client.O_RDWR|client.O_CREATE)
if(fd51 < 0) {
  panic("open failed")
}


ret = client.Close(fd45)
if(ret != 0) {
  panic("close failed")
}


fd52 := client.Open("/tikf5vEc8d.txt", client.O_RDWR|client.O_CREATE)
if(fd52 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd49, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "q4rdB1sB5pX9IbDDuBFWDE8pUobt4kP1MV") {
  panic("wrong data returned")
}


ret = client.Close(fd48)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd51, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd49, 172, client.SEEK_SET)
if(ret != 172) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 172)
  panic("seek failed")
}


ret = client.Seek(fd43, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


buf, ret = client.Read(fd49, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "IQd") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd49, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd43)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd47, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


fd53 := client.Open("/nFxpEbR3hh.txt", client.O_RDWR|client.O_CREATE)
if(fd53 < 0) {
  panic("open failed")
}


fd54 := client.Open("/hIqrxvPj2h.txt", client.O_RDWR|client.O_CREATE)
if(fd54 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd47, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd53)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd37)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd46, []byte("_Cb4_9cMg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) _Cb4_9cMg

ret = client.Seek(fd46, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


buf, ret = client.Read(fd54, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5rsHVNr4h9COEQz9Ej8QTMJZZQh0xZH27HCj97Qblw0sv9HxH38rYxAmUa1706") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd46, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "cMg") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd47, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "BFWDE8pUobt4kP1MVu1ky7bh985unRN") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd52, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd55 := client.Open("/PqQOVlfG3g.txt", client.O_RDWR|client.O_CREATE)
if(fd55 < 0) {
  panic("open failed")
}


ret = client.Close(fd46)
if(ret != 0) {
  panic("close failed")
}


fd56 := client.Open("/NDaLK8HBXg.txt", client.O_RDWR|client.O_CREATE)
if(fd56 < 0) {
  panic("open failed")
}


ret = client.Close(fd54)
if(ret != 0) {
  panic("close failed")
}

//fd state: (105) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdcivdOVZq4rdB1sB5pX9IbDDuBFWDE8pUobt4kP1MVu1ky7bh985unRNtAjmESBjboBckjpZQGjr4YVSDWNaJv9t7q9RAj4iXxquLfUcfadMN4tHVhAYsIaXOX7IQd

ret = client.Write(fd47, []byte("SDU3PBcdQ8iiLsIXacoiYh5K3D7gPH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdcivdOVZq4rdB1sB5pX9IbDDuBFWDE8pUobt4kP1MVu1ky7bh985unRNSDU3PBcdQ8iiLsIXacoiYh5K3D7gPH9t7q9RAj4iXxquLfUcfadMN4tHVhAYsIaXOX7IQd
//fd state: (0) 

ret = client.Write(fd52, []byte("Jom6S0z_sQRmVyAyzb_qnXeZIkhUYznCLVrgmNKjMxcLFCSAfmaGcMpyxSybN3eLPkJ_uvqBYVBxdvXOucLi36dDI_ZqMm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) Jom6S0z_sQRmVyAyzb_qnXeZIkhUYznCLVrgmNKjMxcLFCSAfmaGcMpyxSybN3eLPkJ_uvqBYVBxdvXOucLi36dDI_ZqMm

ret = client.Close(fd49)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd55)
if(ret != 0) {
  panic("close failed")
}


fd57 := client.Open("/_laoVJpAG9.txt", client.O_RDWR|client.O_CREATE)
if(fd57 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd56, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_Cb4_9cMg") {
  panic("wrong data returned")
}


ret = client.Seek(fd57, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd57, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd51, []byte("y6T7VqmwXHCRDXaELyOuFuhsta3os1kk2WDaq9tVotSR_6kWADAQuINqi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) y6T7VqmwXHCRDXaELyOuFuhsta3os1kk2WDaq9tVotSR_6kWADAQuINqi
//fd state: (9) _Cb4_9cMg

ret = client.Write(fd56, []byte("wbLx4w1hvzD6ri4ml0MCqWbeSH6NAPfFP6NAY3iTYmtWZKVZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) _Cb4_9cMgwbLx4w1hvzD6ri4ml0MCqWbeSH6NAPfFP6NAY3iTYmtWZKVZ

ret = client.Seek(fd57, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Close(fd52)
if(ret != 0) {
  panic("close failed")
}


fd58 := client.Open("/9RkaPG0bm2.txt", client.O_RDWR|client.O_CREATE)
if(fd58 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd58, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd57, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd58, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd57, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "A9awL") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd47, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9t7q9RAj4iXxquLfUcfadMN4tHVhAYsIaXOX7IQd") {
  panic("wrong data returned")
}


ret = client.Seek(fd47, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


buf, ret = client.Read(fd57, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd58)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd51)
if(ret != 0) {
  panic("close failed")
}

//fd state: (57) _Cb4_9cMgwbLx4w1hvzD6ri4ml0MCqWbeSH6NAPfFP6NAY3iTYmtWZKVZ

ret = client.Write(fd56, []byte("T8dnwZgJ_wn_yuyCevgBDLCFgKRyTQ1eB8bZq1Pc2t3uvsjhGqKbx0V_CnQx27xDTy98s4qzewjk5p50ZZLyj0vU8NSGc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) _Cb4_9cMgwbLx4w1hvzD6ri4ml0MCqWbeSH6NAPfFP6NAY3iTYmtWZKVZT8dnwZgJ_wn_yuyCevgBDLCFgKRyTQ1eB8bZq1Pc2t3uvsjhGqKbx0V_CnQx27xDTy98s4qzewjk5p50ZZLyj0vU8NSGc

buf, ret = client.Read(fd56, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd57)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd56)
if(ret != 0) {
  panic("close failed")
}

//fd state: (10) LIUbj_rdckWyR1_p1ZbX0tc1Xkv4RtiiHAcsCfWnJhC_nttrVdcivdOVZq4rdB1sB5pX9IbDDuBFWDE8pUobt4kP1MVu1ky7bh985unRNSDU3PBcdQ8iiLsIXacoiYh5K3D7gPH9t7q9RAj4iXxquLfUcfadMN4tHVhAYsIaXOX7IQd

ret = client.Write(fd47, []byte("qzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) LIUbj_rdckqzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAUFWDE8pUobt4kP1MVu1ky7bh985unRNSDU3PBcdQ8iiLsIXacoiYh5K3D7gPH9t7q9RAj4iXxquLfUcfadMN4tHVhAYsIaXOX7IQd

fd59 := client.Open("/9WrO_P7NpN.txt", client.O_RDWR|client.O_CREATE)
if(fd59 < 0) {
  panic("open failed")
}


fd60 := client.Open("/iVIJjg2elX.txt", client.O_RDWR|client.O_CREATE)
if(fd60 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd59, []byte("EqGbKfMvDCSK1ADCrOZlRvUOMbFrFj9inz9fakpXHtAjifdi6IfmAzxovYi_qX38Jk_gteEficrACotQS5Kpb_QQbAoqZI8jTX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) EqGbKfMvDCSK1ADCrOZlRvUOMbFrFj9inz9fakpXHtAjifdi6IfmAzxovYi_qX38Jk_gteEficrACotQS5Kpb_QQbAoqZI8jTX

buf, ret = client.Read(fd59, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd60)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd59)
if(ret != 0) {
  panic("close failed")
}

//fd state: (75) LIUbj_rdckqzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAUFWDE8pUobt4kP1MVu1ky7bh985unRNSDU3PBcdQ8iiLsIXacoiYh5K3D7gPH9t7q9RAj4iXxquLfUcfadMN4tHVhAYsIaXOX7IQd

ret = client.Write(fd47, []byte("6Sh9H5KvARhTDGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (172) LIUbj_rdckqzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAU6Sh9H5KvARhTDGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGIQd

fd61 := client.Open("/PqQOVlfG3g.txt", client.O_RDWR|client.O_CREATE)
if(fd61 < 0) {
  panic("open failed")
}

//fd state: (172) LIUbj_rdckqzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAU6Sh9H5KvARhTDGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGIQd

ret = client.Write(fd47, []byte("dxoPlKW7lbYhouaDYod8l359wnP04T1CrJ9pj2x2GFzeliIJWpQG7IBB3airdkABEq6488D"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (243) LIUbj_rdckqzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAU6Sh9H5KvARhTDGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGdxoPlKW7lbYhouaDYod8l359wnP04T1CrJ9pj2x2GFzeliIJWpQG7IBB3airdkABEq6488D

fd62 := client.Open("/nFxpEbR3hh.txt", client.O_RDWR|client.O_CREATE)
if(fd62 < 0) {
  panic("open failed")
}


ret = client.Seek(fd61, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd47)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd62, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (2) WWJXXW7hq63_g

ret = client.Write(fd62, []byte("W9jD0wWvFE8yU7Q8LSOQxkgh6m4uzNZExk5rkAGJftQG7ods2vaJ2WnZSo6mdSytwRskoTjPqE1QNDuc47HBw7RJmPMrIrio"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) WWW9jD0wWvFE8yU7Q8LSOQxkgh6m4uzNZExk5rkAGJftQG7ods2vaJ2WnZSo6mdSytwRskoTjPqE1QNDuc47HBw7RJmPMrIrio

ret = client.Seek(fd62, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


buf, ret = client.Read(fd61, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (29) WWW9jD0wWvFE8yU7Q8LSOQxkgh6m4uzNZExk5rkAGJftQG7ods2vaJ2WnZSo6mdSytwRskoTjPqE1QNDuc47HBw7RJmPMrIrio

ret = client.Write(fd62, []byte("Fz4VOVwiAkY93ClyguqNko2L3_943IqXr2xip5VUGqWzYLGzS1Mc06SQjI4h6R9epgcvuQOMLe4HZZHB0crNW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) WWW9jD0wWvFE8yU7Q8LSOQxkgh6m4Fz4VOVwiAkY93ClyguqNko2L3_943IqXr2xip5VUGqWzYLGzS1Mc06SQjI4h6R9epgcvuQOMLe4HZZHB0crNW

fd63 := client.Open("/49yoWYxM6B.txt", client.O_RDWR|client.O_CREATE)
if(fd63 < 0) {
  panic("open failed")
}


fd64 := client.Open("/fT8o_szxtr.txt", client.O_RDWR|client.O_CREATE)
if(fd64 < 0) {
  panic("open failed")
}


ret = client.Close(fd64)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd61, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd61, []byte("GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFhYf8EavuTrfTDQRB5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFhYf8EavuTrfTDQRB5
//fd state: (0) 

ret = client.Write(fd63, []byte("7wkEiOugHMQuzqR1EUlI3j2o"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) 7wkEiOugHMQuzqR1EUlI3j2o
//fd state: (114) WWW9jD0wWvFE8yU7Q8LSOQxkgh6m4Fz4VOVwiAkY93ClyguqNko2L3_943IqXr2xip5VUGqWzYLGzS1Mc06SQjI4h6R9epgcvuQOMLe4HZZHB0crNW

ret = client.Write(fd62, []byte("l0IdARg68mv2hpnc0fNJmlLFGEtlr1t7ED9RoGcVwdsykA2vG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (163) WWW9jD0wWvFE8yU7Q8LSOQxkgh6m4Fz4VOVwiAkY93ClyguqNko2L3_943IqXr2xip5VUGqWzYLGzS1Mc06SQjI4h6R9epgcvuQOMLe4HZZHB0crNWl0IdARg68mv2hpnc0fNJmlLFGEtlr1t7ED9RoGcVwdsykA2vG
//fd state: (163) WWW9jD0wWvFE8yU7Q8LSOQxkgh6m4Fz4VOVwiAkY93ClyguqNko2L3_943IqXr2xip5VUGqWzYLGzS1Mc06SQjI4h6R9epgcvuQOMLe4HZZHB0crNWl0IdARg68mv2hpnc0fNJmlLFGEtlr1t7ED9RoGcVwdsykA2vG

ret = client.Write(fd62, []byte("r1jNpAJB3rJO8Mojl9_tcIm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (186) WWW9jD0wWvFE8yU7Q8LSOQxkgh6m4Fz4VOVwiAkY93ClyguqNko2L3_943IqXr2xip5VUGqWzYLGzS1Mc06SQjI4h6R9epgcvuQOMLe4HZZHB0crNWl0IdARg68mv2hpnc0fNJmlLFGEtlr1t7ED9RoGcVwdsykA2vGr1jNpAJB3rJO8Mojl9_tcIm
//fd state: (24) 7wkEiOugHMQuzqR1EUlI3j2o

ret = client.Write(fd63, []byte("ptU2OgMVieqY7f23Qxu9VDg7og0Dx8cLmrv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23Qxu9VDg7og0Dx8cLmrv

fd65 := client.Open("/DNkbwkXRNz.txt", client.O_RDWR|client.O_CREATE)
if(fd65 < 0) {
  panic("open failed")
}


ret = client.Close(fd62)
if(ret != 0) {
  panic("close failed")
}


fd66 := client.Open("/ymeRdV7EO7.txt", client.O_RDWR|client.O_CREATE)
if(fd66 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd61, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd67 := client.Open("/0KnXHjRLMh.txt", client.O_RDWR|client.O_CREATE)
if(fd67 < 0) {
  panic("open failed")
}


ret = client.Close(fd61)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd65, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "y6T7VqmwXHCRDXaELyOuFuhsta3os1kk2WDaq9tVotS") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd66, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (59) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23Qxu9VDg7og0Dx8cLmrv

ret = client.Write(fd63, []byte("on0wfUdnDUKem5wMIzQtGxoP7jdGAb957hD0_XbPel8dpVF0sIFhFI0WoteC2YAhycEXkPF44D"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23Qxu9VDg7og0Dx8cLmrvon0wfUdnDUKem5wMIzQtGxoP7jdGAb957hD0_XbPel8dpVF0sIFhFI0WoteC2YAhycEXkPF44D

ret = client.Close(fd66)
if(ret != 0) {
  panic("close failed")
}

//fd state: (133) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23Qxu9VDg7og0Dx8cLmrvon0wfUdnDUKem5wMIzQtGxoP7jdGAb957hD0_XbPel8dpVF0sIFhFI0WoteC2YAhycEXkPF44D

ret = client.Write(fd63, []byte("__ewfLKiMnHC0KYxrCzaL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23Qxu9VDg7og0Dx8cLmrvon0wfUdnDUKem5wMIzQtGxoP7jdGAb957hD0_XbPel8dpVF0sIFhFI0WoteC2YAhycEXkPF44D__ewfLKiMnHC0KYxrCzaL

fd68 := client.Open("/CzVjEdkRbn.txt", client.O_RDWR|client.O_CREATE)
if(fd68 < 0) {
  panic("open failed")
}


fd69 := client.Open("/_fKtkmMG2H.txt", client.O_RDWR|client.O_CREATE)
if(fd69 < 0) {
  panic("open failed")
}

//fd state: (43) y6T7VqmwXHCRDXaELyOuFuhsta3os1kk2WDaq9tVotSR_6kWADAQuINqi

ret = client.Write(fd65, []byte("LxzlJQW8JzpSCUIYneiUm8H3SPq816QZ0lrcJDc1QhQ36g0N_qWttHuYik1n16pVyg3OE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (112) y6T7VqmwXHCRDXaELyOuFuhsta3os1kk2WDaq9tVotSLxzlJQW8JzpSCUIYneiUm8H3SPq816QZ0lrcJDc1QhQ36g0N_qWttHuYik1n16pVyg3OE

ret = client.Seek(fd63, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}


fd70 := client.Open("/nYDyf3ULIn.txt", client.O_RDWR|client.O_CREATE)
if(fd70 < 0) {
  panic("open failed")
}


fd71 := client.Open("/giqe_vyTh1.txt", client.O_RDWR|client.O_CREATE)
if(fd71 < 0) {
  panic("open failed")
}


ret = client.Close(fd70)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd68, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


buf, ret = client.Read(fd65, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd69, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd63)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd69, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd71, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd67, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd68, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd67, []byte("9bNoG1Ufmn2k2NqhAcctJG6WlOU_qmCtSS1wY0mND4hGriHyglXWoQ2MckQR94xStbsZm2d67eflO17KFXoDKXxby40NBL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) 9bNoG1Ufmn2k2NqhAcctJG6WlOU_qmCtSS1wY0mND4hGriHyglXWoQ2MckQR94xStbsZm2d67eflO17KFXoDKXxby40NBL

fd72 := client.Open("/giqe_vyTh1.txt", client.O_RDWR|client.O_CREATE)
if(fd72 < 0) {
  panic("open failed")
}


ret = client.Close(fd65)
if(ret != 0) {
  panic("close failed")
}


fd73 := client.Open("/KVp2ABEVFd.txt", client.O_RDWR|client.O_CREATE)
if(fd73 < 0) {
  panic("open failed")
}


ret = client.Close(fd68)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd73, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd71)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd69, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd73, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd67, 71, client.SEEK_SET)
if(ret != 71) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 71)
  panic("seek failed")
}


ret = client.Seek(fd73, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd74 := client.Open("/ojkJUB3zCr.txt", client.O_RDWR|client.O_CREATE)
if(fd74 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd73, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd69, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd75 := client.Open("/49yoWYxM6B.txt", client.O_RDWR|client.O_CREATE)
if(fd75 < 0) {
  panic("open failed")
}


fd76 := client.Open("/tikf5vEc8d.txt", client.O_RDWR|client.O_CREATE)
if(fd76 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd73, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd75, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23") {
  panic("wrong data returned")
}


ret = client.Close(fd69)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd72, []byte("TCzGHiVA_Lob_c0CcGIbvDsAQHT3RD51wkv7kleWt5cBvgvN9hlTcIVShen5q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) TCzGHiVA_Lob_c0CcGIbvDsAQHT3RD51wkv7kleWt5cBvgvN9hlTcIVShen5q

ret = client.Seek(fd76, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


ret = client.Close(fd74)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd73, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd73, []byte("_1iBTUUm7_QY_Lqz53Os_rcuBg8U6vU57WPvm_YHsV2idRgREHcgbAb8Lotx9EyF8sUHtWxXyxhaRDTXgpdt_jT9JbQjIbWK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) _1iBTUUm7_QY_Lqz53Os_rcuBg8U6vU57WPvm_YHsV2idRgREHcgbAb8Lotx9EyF8sUHtWxXyxhaRDTXgpdt_jT9JbQjIbWK
//fd state: (40) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23Qxu9VDg7og0Dx8cLmrvon0wfUdnDUKem5wMIzQtGxoP7jdGAb957hD0_XbPel8dpVF0sIFhFI0WoteC2YAhycEXkPF44D__ewfLKiMnHC0KYxrCzaL

ret = client.Write(fd75, []byte("xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJdpVF0sIFhFI0WoteC2YAhycEXkPF44D__ewfLKiMnHC0KYxrCzaL

ret = client.Close(fd76)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd72)
if(ret != 0) {
  panic("close failed")
}


fd77 := client.Open("/3v2LLkR1Z4.txt", client.O_RDWR|client.O_CREATE)
if(fd77 < 0) {
  panic("open failed")
}

//fd state: (102) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJdpVF0sIFhFI0WoteC2YAhycEXkPF44D__ewfLKiMnHC0KYxrCzaL

ret = client.Write(fd75, []byte("Z1jGm_vJ2JagjFI_5CYn4AGSAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (192) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJZ1jGm_vJ2JagjFI_5CYn4AGSAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPx

buf, ret = client.Read(fd73, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd77, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd78 := client.Open("/tikf5vEc8d.txt", client.O_RDWR|client.O_CREATE)
if(fd78 < 0) {
  panic("open failed")
}


ret = client.Seek(fd77, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (96) _1iBTUUm7_QY_Lqz53Os_rcuBg8U6vU57WPvm_YHsV2idRgREHcgbAb8Lotx9EyF8sUHtWxXyxhaRDTXgpdt_jT9JbQjIbWK

ret = client.Write(fd73, []byte("KXWBoPiuZ_si6eucC4H6bLe8zbghGY788CviYT7TtDPNdUXARtnatRrek2voasvJnsRYh2VC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (168) _1iBTUUm7_QY_Lqz53Os_rcuBg8U6vU57WPvm_YHsV2idRgREHcgbAb8Lotx9EyF8sUHtWxXyxhaRDTXgpdt_jT9JbQjIbWKKXWBoPiuZ_si6eucC4H6bLe8zbghGY788CviYT7TtDPNdUXARtnatRrek2voasvJnsRYh2VC

ret = client.Seek(fd67, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


buf, ret = client.Read(fd67, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JG6WlOU_qmCtSS1wY0mND4hGriHyglXWoQ2MckQR94xStbsZm2d67eflO17KFXoDKXxby40NBL") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd77, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd77, []byte("Lbzpk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) Lbzpk

buf, ret = client.Read(fd77, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd73, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd79 := client.Open("/oIsR5Rk83H.txt", client.O_RDWR|client.O_CREATE)
if(fd79 < 0) {
  panic("open failed")
}

//fd state: (94) 9bNoG1Ufmn2k2NqhAcctJG6WlOU_qmCtSS1wY0mND4hGriHyglXWoQ2MckQR94xStbsZm2d67eflO17KFXoDKXxby40NBL

ret = client.Write(fd67, []byte("BW3Cl2hpiEessjB3a4iLPZvYM8ImXPllfncWIcexbx9kVhEM1NMsZkZoPCLKfsIFN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (159) 9bNoG1Ufmn2k2NqhAcctJG6WlOU_qmCtSS1wY0mND4hGriHyglXWoQ2MckQR94xStbsZm2d67eflO17KFXoDKXxby40NBLBW3Cl2hpiEessjB3a4iLPZvYM8ImXPllfncWIcexbx9kVhEM1NMsZkZoPCLKfsIFN

buf, ret = client.Read(fd78, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Jom6S0z_sQRmVyAyzb_qnXe") {
  panic("wrong data returned")
}


ret = client.Close(fd79)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd75, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd67, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd67, 132, client.SEEK_SET)
if(ret != 132) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 132)
  panic("seek failed")
}


ret = client.Seek(fd78, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


ret = client.Close(fd77)
if(ret != 0) {
  panic("close failed")
}

//fd state: (192) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJZ1jGm_vJ2JagjFI_5CYn4AGSAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPx

ret = client.Write(fd75, []byte("zublhf4CaQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (202) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJZ1jGm_vJ2JagjFI_5CYn4AGSAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQ

ret = client.Close(fd67)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd78, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


ret = client.Seek(fd75, 178, client.SEEK_SET)
if(ret != 178) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 178)
  panic("seek failed")
}


ret = client.Seek(fd73, 66, client.SEEK_SET)
if(ret != 66) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 66)
  panic("seek failed")
}


ret = client.Close(fd75)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd78, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


ret = client.Seek(fd78, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


fd80 := client.Open("/3v2LLkR1Z4.txt", client.O_RDWR|client.O_CREATE)
if(fd80 < 0) {
  panic("open failed")
}

//fd state: (66) _1iBTUUm7_QY_Lqz53Os_rcuBg8U6vU57WPvm_YHsV2idRgREHcgbAb8Lotx9EyF8sUHtWxXyxhaRDTXgpdt_jT9JbQjIbWKKXWBoPiuZ_si6eucC4H6bLe8zbghGY788CviYT7TtDPNdUXARtnatRrek2voasvJnsRYh2VC

ret = client.Write(fd73, []byte("bLHxXSRxySAV84ewifbGa5_cqDeOukC8WQM5I_W5cyXDyKPpk4oVSn_vL7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) _1iBTUUm7_QY_Lqz53Os_rcuBg8U6vU57WPvm_YHsV2idRgREHcgbAb8Lotx9EyF8sbLHxXSRxySAV84ewifbGa5_cqDeOukC8WQM5I_W5cyXDyKPpk4oVSn_vL7GY788CviYT7TtDPNdUXARtnatRrek2voasvJnsRYh2VC

ret = client.Seek(fd73, 113, client.SEEK_SET)
if(ret != 113) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 113)
  panic("seek failed")
}


ret = client.Seek(fd78, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


ret = client.Close(fd78)
if(ret != 0) {
  panic("close failed")
}


fd81 := client.Open("/nYDyf3ULIn.txt", client.O_RDWR|client.O_CREATE)
if(fd81 < 0) {
  panic("open failed")
}


ret = client.Close(fd73)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd80, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Lbzpk") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd81, []byte("8ylfh0lUiOOwt1_jj3XupmWZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) 8ylfh0lUiOOwt1_jj3XupmWZ

buf, ret = client.Read(fd80, 63)
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

//fd state: (5) Lbzpk

ret = client.Write(fd80, []byte("y1vcS_ijRGx8BwPLTCKGYu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) Lbzpky1vcS_ijRGx8BwPLTCKGYu

ret = client.Close(fd80)
if(ret != 0) {
  panic("close failed")
}


fd82 := client.Open("/KqykqQdiXF.txt", client.O_RDWR|client.O_CREATE)
if(fd82 < 0) {
  panic("open failed")
}


ret = client.Seek(fd82, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd83 := client.Open("/LZcAyv_xCm.txt", client.O_RDWR|client.O_CREATE)
if(fd83 < 0) {
  panic("open failed")
}


ret = client.Seek(fd82, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd82, []byte("XqeR5BQTtne3Y1V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) XqeR5BQTtne3Y1V

fd84 := client.Open("/CzVjEdkRbn.txt", client.O_RDWR|client.O_CREATE)
if(fd84 < 0) {
  panic("open failed")
}

//fd state: (0) V4p42BuHl7auv6IcbLZFdUrpUGOzpD9qmUhLGU9JdnpxzL6ZMAvLhZuZ8n4tRZHsEtLF4Lnhg4Byy2dBtx2zUUu36LupvXob

ret = client.Write(fd84, []byte("ggY26C3_r2cnuFByadCKwjL9xl7zKo4ZJ7aKwf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) ggY26C3_r2cnuFByadCKwjL9xl7zKo4ZJ7aKwf9JdnpxzL6ZMAvLhZuZ8n4tRZHsEtLF4Lnhg4Byy2dBtx2zUUu36LupvXob
//fd state: (38) ggY26C3_r2cnuFByadCKwjL9xl7zKo4ZJ7aKwf9JdnpxzL6ZMAvLhZuZ8n4tRZHsEtLF4Lnhg4Byy2dBtx2zUUu36LupvXob

ret = client.Write(fd84, []byte("6zdm0bGj1L_KAQ7KMYvzZrtjAEDPkmX_y1E97UT9DrE1Jmw2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) ggY26C3_r2cnuFByadCKwjL9xl7zKo4ZJ7aKwf6zdm0bGj1L_KAQ7KMYvzZrtjAEDPkmX_y1E97UT9DrE1Jmw2u36LupvXob

ret = client.Close(fd82)
if(ret != 0) {
  panic("close failed")
}


fd85 := client.Open("/DbkC5q9e6v.txt", client.O_RDWR|client.O_CREATE)
if(fd85 < 0) {
  panic("open failed")
}


fd86 := client.Open("/fT8o_szxtr.txt", client.O_RDWR|client.O_CREATE)
if(fd86 < 0) {
  panic("open failed")
}


ret = client.Seek(fd84, 68, client.SEEK_SET)
if(ret != 68) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 68)
  panic("seek failed")
}


ret = client.Seek(fd86, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd87 := client.Open("/2dNN9CmhK3.txt", client.O_RDWR|client.O_CREATE)
if(fd87 < 0) {
  panic("open failed")
}


ret = client.Close(fd86)
if(ret != 0) {
  panic("close failed")
}


fd88 := client.Open("/bCLvtka1QB.txt", client.O_RDWR|client.O_CREATE)
if(fd88 < 0) {
  panic("open failed")
}


fd89 := client.Open("/ek97Se9uIB.txt", client.O_RDWR|client.O_CREATE)
if(fd89 < 0) {
  panic("open failed")
}


ret = client.Close(fd88)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd83)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd87, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd89, []byte("u2Pdd5XP9dt2z1K551eiFELJlAAB_EF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) u2Pdd5XP9dt2z1K551eiFELJlAAB_EF

ret = client.Seek(fd89, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}


fd90 := client.Open("/A2jhBeQbpL.txt", client.O_RDWR|client.O_CREATE)
if(fd90 < 0) {
  panic("open failed")
}


fd91 := client.Open("/jEJT1wijAu.txt", client.O_RDWR|client.O_CREATE)
if(fd91 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd91, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ev7_fh1MUOr5IrM8oYnIs0C56_m6cR2mZCGDWBt9") {
  panic("wrong data returned")
}

//fd state: (68) ggY26C3_r2cnuFByadCKwjL9xl7zKo4ZJ7aKwf6zdm0bGj1L_KAQ7KMYvzZrtjAEDPkmX_y1E97UT9DrE1Jmw2u36LupvXob

ret = client.Write(fd84, []byte("eTsjdHaWEtnNjNI23sSroqevIUhq7RqqfM5ZqYNTVEM4sg9w0yjotrFsPZA1InzOOuz3I4rTZ8KHkkfAd_pCj1DOI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (157) ggY26C3_r2cnuFByadCKwjL9xl7zKo4ZJ7aKwf6zdm0bGj1L_KAQ7KMYvzZrtjAEDPkmeTsjdHaWEtnNjNI23sSroqevIUhq7RqqfM5ZqYNTVEM4sg9w0yjotrFsPZA1InzOOuz3I4rTZ8KHkkfAd_pCj1DOI

ret = client.Close(fd87)
if(ret != 0) {
  panic("close failed")
}


fd92 := client.Open("/49yoWYxM6B.txt", client.O_RDWR|client.O_CREATE)
if(fd92 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd85, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd84, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd93 := client.Open("/Q7Imfuqduy.txt", client.O_RDWR|client.O_CREATE)
if(fd93 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd92, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmd") {
  panic("wrong data returned")
}


ret = client.Close(fd90)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd93, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "QloNu8PD5bFWZEedHwGeoA44aCr") {
  panic("wrong data returned")
}


ret = client.Seek(fd92, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}

//fd state: (30) u2Pdd5XP9dt2z1K551eiFELJlAAB_EF

ret = client.Write(fd89, []byte("u0bEJqEFKVUiNXejeBxxPvgLOfhvHfm5Y6UDndIKq10l2bhwIAG37ldenSy5CUiEBb8eKk3yCPGcYDaB9pSopjNlOV7EBZtwO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (127) u2Pdd5XP9dt2z1K551eiFELJlAAB_Eu0bEJqEFKVUiNXejeBxxPvgLOfhvHfm5Y6UDndIKq10l2bhwIAG37ldenSy5CUiEBb8eKk3yCPGcYDaB9pSopjNlOV7EBZtwO

fd94 := client.Open("/n16dqTGUku.txt", client.O_RDWR|client.O_CREATE)
if(fd94 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd84, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (27) QloNu8PD5bFWZEedHwGeoA44aCr

ret = client.Write(fd93, []byte("8hHY5WLj00vLaTI8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) QloNu8PD5bFWZEedHwGeoA44aCr8hHY5WLj00vLaTI8
//fd state: (0) 

ret = client.Write(fd94, []byte("vSh9BGLqQYbTCsHMCzslrne8cSw4em9wkXxiiH4Uyr3l2xdoHpTykE7busGUR5ZN8InpQKb3v2L8Hlev9iJZDaHcxD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) vSh9BGLqQYbTCsHMCzslrne8cSw4em9wkXxiiH4Uyr3l2xdoHpTykE7busGUR5ZN8InpQKb3v2L8Hlev9iJZDaHcxD

buf, ret = client.Read(fd93, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd91)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd92, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJZ1jGm_vJ2Ja") {
  panic("wrong data returned")
}


fd95 := client.Open("/uEvYTullIO.txt", client.O_RDWR|client.O_CREATE)
if(fd95 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd92, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gjFI_5CYn4AGSAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZS") {
  panic("wrong data returned")
}


fd96 := client.Open("/49yoWYxM6B.txt", client.O_RDWR|client.O_CREATE)
if(fd96 < 0) {
  panic("open failed")
}

//fd state: (0) 7wkEiOugHMQuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJZ1jGm_vJ2JagjFI_5CYn4AGSAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQ

ret = client.Write(fd96, []byte("jBS64Lgzdah"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) jBS64LgzdahuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJZ1jGm_vJ2JagjFI_5CYn4AGSAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQ

ret = client.Close(fd95)
if(ret != 0) {
  panic("close failed")
}


fd97 := client.Open("/qinn9z7zrr.txt", client.O_RDWR|client.O_CREATE)
if(fd97 < 0) {
  panic("open failed")
}


ret = client.Seek(fd93, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}


fd98 := client.Open("/qinn9z7zrr.txt", client.O_RDWR|client.O_CREATE)
if(fd98 < 0) {
  panic("open failed")
}


ret = client.Close(fd96)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd98, []byte("5HPyY8VZrAHHzXkJvUQLPN0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) 5HPyY8VZrAHHzXkJvUQLPN0

ret = client.Seek(fd94, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


buf, ret = client.Read(fd85, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd98, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd84, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


ret = client.Seek(fd89, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


buf, ret = client.Read(fd98, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd92, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQ") {
  panic("wrong data returned")
}

//fd state: (0) 5HPyY8VZrAHHzXkJvUQLPN0

ret = client.Write(fd97, []byte("AeT1OoBYcsTLvyWQenTGF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) AeT1OoBYcsTLvyWQenTGFN0

fd99 := client.Open("/Q7Imfuqduy.txt", client.O_RDWR|client.O_CREATE)
if(fd99 < 0) {
  panic("open failed")
}


ret = client.Seek(fd93, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


buf, ret = client.Read(fd84, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MYvzZrtjAEDPkmeTsjdHaWEtnNjNI23sSroqevIUhq7Rq") {
  panic("wrong data returned")
}


ret = client.Close(fd98)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd99, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


buf, ret = client.Read(fd84, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qfM5ZqYNTVEM4sg9w0yjotrFsPZA1InzOOuz3I4rTZ") {
  panic("wrong data returned")
}


fd100 := client.Open("/4cQqDopMxJ.txt", client.O_RDWR|client.O_CREATE)
if(fd100 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd100, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "KCPdD3hBtZztTxnEDBH6") {
  panic("wrong data returned")
}

//fd state: (202) jBS64LgzdahuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJZ1jGm_vJ2JagjFI_5CYn4AGSAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQ

ret = client.Write(fd92, []byte("pgygN3aOEmoYMSI9sDypzuOtD86FeN7ykkKcOm83iIXDDKOMQLAtZnnM85iOcF_3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (266) jBS64LgzdahuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJZ1jGm_vJ2JagjFI_5CYn4AGSAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQpgygN3aOEmoYMSI9sDypzuOtD86FeN7ykkKcOm83iIXDDKOMQLAtZnnM85iOcF_3

ret = client.Close(fd94)
if(ret != 0) {
  panic("close failed")
}

//fd state: (2) QloNu8PD5bFWZEedHwGeoA44aCr8hHY5WLj00vLaTI8

ret = client.Write(fd93, []byte("fFP7huvYDSYqJ8Z87Nnrkakjo374yLgXMiSv7_LoJStBdflUpJgq3hPi6Pu5bUs1p2t7XU7ETR3f2WwILslIiScO5O5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) QlfFP7huvYDSYqJ8Z87Nnrkakjo374yLgXMiSv7_LoJStBdflUpJgq3hPi6Pu5bUs1p2t7XU7ETR3f2WwILslIiScO5O5
//fd state: (0) 

ret = client.Write(fd85, []byte("y76GTNImWOrQbI8Ujy6WBFnHgX8XNp4FzL97YJ01GhBFDVxJHXkpME6zlKmwrxriiCa1uV4VwFtOn5a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) y76GTNImWOrQbI8Ujy6WBFnHgX8XNp4FzL97YJ01GhBFDVxJHXkpME6zlKmwrxriiCa1uV4VwFtOn5a
//fd state: (54) u2Pdd5XP9dt2z1K551eiFELJlAAB_Eu0bEJqEFKVUiNXejeBxxPvgLOfhvHfm5Y6UDndIKq10l2bhwIAG37ldenSy5CUiEBb8eKk3yCPGcYDaB9pSopjNlOV7EBZtwO

ret = client.Write(fd89, []byte("zD4V6oOEfVkO9MN_0XLGWuifvboAqG_O2z_PhYT06OSPURyprHWAV6VvNJOgRrwHAMuGaglLGaSnTHh3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) u2Pdd5XP9dt2z1K551eiFELJlAAB_Eu0bEJqEFKVUiNXejeBxxPvgLzD4V6oOEfVkO9MN_0XLGWuifvboAqG_O2z_PhYT06OSPURyprHWAV6VvNJOgRrwHAMuGaglLGaSnTHh3

ret = client.Close(fd93)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd97, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "N0") {
  panic("wrong data returned")
}

//fd state: (13) QlfFP7huvYDSYqJ8Z87Nnrkakjo374yLgXMiSv7_LoJStBdflUpJgq3hPi6Pu5bUs1p2t7XU7ETR3f2WwILslIiScO5O5

ret = client.Write(fd99, []byte("EcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOdQIMmIptAw0jmd9pH7tnlVuOhqbrFS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOdQIMmIptAw0jmd9pH7tnlVuOhqbrFS
//fd state: (134) u2Pdd5XP9dt2z1K551eiFELJlAAB_Eu0bEJqEFKVUiNXejeBxxPvgLzD4V6oOEfVkO9MN_0XLGWuifvboAqG_O2z_PhYT06OSPURyprHWAV6VvNJOgRrwHAMuGaglLGaSnTHh3

ret = client.Write(fd89, []byte("d011Hr0QCWeLSF1ZHxoHyaxcwATog8cXSkzlnA9AbJYuFShi21QntLwx6a4lFGswYojpxqZmynidVxpTlyQHDtMceEBLywWt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (230) u2Pdd5XP9dt2z1K551eiFELJlAAB_Eu0bEJqEFKVUiNXejeBxxPvgLzD4V6oOEfVkO9MN_0XLGWuifvboAqG_O2z_PhYT06OSPURyprHWAV6VvNJOgRrwHAMuGaglLGaSnTHh3d011Hr0QCWeLSF1ZHxoHyaxcwATog8cXSkzlnA9AbJYuFShi21QntLwx6a4lFGswYojpxqZmynidVxpTlyQHDtMceEBLywWt
//fd state: (20) KCPdD3hBtZztTxnEDBH6JZrvcM5Hmxhw0vpXOjEmmJg6Qi6Z1dIvOnonkfK8ZA_vkNxdvNY3Vw6SNrCXel79elj46m5IkvBkpUh2zOzCuVI1gv1nc3QiQyfHYMbSmlKnRiOLQxf

ret = client.Write(fd100, []byte("5kYLRELkwBL1O7324t311qNX3S0PTcNIVoD2JZdBYbWOuHIthsQV3RDO2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) KCPdD3hBtZztTxnEDBH65kYLRELkwBL1O7324t311qNX3S0PTcNIVoD2JZdBYbWOuHIthsQV3RDO2rCXel79elj46m5IkvBkpUh2zOzCuVI1gv1nc3QiQyfHYMbSmlKnRiOLQxf

ret = client.Close(fd89)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd100)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd85, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd99, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd84, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


ret = client.Seek(fd92, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


ret = client.Close(fd85)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd84, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NjNI23sSroqevIUhq7RqqfM5ZqYNTVEM4sg9w0yjotrFsPZA1InzOOuz3I4rTZ8KHkkfAd_pCj1DOI") {
  panic("wrong data returned")
}


fd101 := client.Open("/NDaLK8HBXg.txt", client.O_RDWR|client.O_CREATE)
if(fd101 < 0) {
  panic("open failed")
}


fd102 := client.Open("/AYwrkuI1uC.txt", client.O_RDWR|client.O_CREATE)
if(fd102 < 0) {
  panic("open failed")
}


fd103 := client.Open("/JmFn4HpF0R.txt", client.O_RDWR|client.O_CREATE)
if(fd103 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd103, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd101, 61, client.SEEK_SET)
if(ret != 61) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 61)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd103, []byte("aVA1fVnlWZ68e0I4dA0zoElmjVsD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) aVA1fVnlWZ68e0I4dA0zoElmjVsD

ret = client.Close(fd92)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd97, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd97, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd104 := client.Open("/gwR_qKf5l_.txt", client.O_RDWR|client.O_CREATE)
if(fd104 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd97, 72)
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


ret = client.Seek(fd102, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd99, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd104, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd99, 73, client.SEEK_SET)
if(ret != 73) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 73)
  panic("seek failed")
}

//fd state: (73) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOdQIMmIptAw0jmd9pH7tnlVuOhqbrFS

ret = client.Write(fd99, []byte("nTulIVTsfJsDkf5EfXwHt2JCV0tXQ1a3B5vdpArqfvVKGdvwm9CUcFmvHPidb7lbyhVYcV3qwK9khVcOeLcSFAWYQzxxk6b0SK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (171) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOnTulIVTsfJsDkf5EfXwHt2JCV0tXQ1a3B5vdpArqfvVKGdvwm9CUcFmvHPidb7lbyhVYcV3qwK9khVcOeLcSFAWYQzxxk6b0SK

ret = client.Seek(fd104, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd102, []byte("SHAiX2hfhaiL9Wl91x5DT5IiHMt_0WvBQ5ID"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) SHAiX2hfhaiL9Wl91x5DT5IiHMt_0WvBQ5ID

fd105 := client.Open("/DbkC5q9e6v.txt", client.O_RDWR|client.O_CREATE)
if(fd105 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd105, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "y76GTNImWOrQbI8Ujy6WBFnHgX8XNp4FzL97YJ01GhBFDVxJHXkpME6zlKmwrxriiCa1u") {
  panic("wrong data returned")
}


fd106 := client.Open("/gwR_qKf5l_.txt", client.O_RDWR|client.O_CREATE)
if(fd106 < 0) {
  panic("open failed")
}


ret = client.Close(fd103)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd106, []byte("SQNlLo8g3yfsNMAHFK30GSiJnv7sIK4JCG_uTZ5FY5CfMLJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) SQNlLo8g3yfsNMAHFK30GSiJnv7sIK4JCG_uTZ5FY5CfMLJ
//fd state: (157) ggY26C3_r2cnuFByadCKwjL9xl7zKo4ZJ7aKwf6zdm0bGj1L_KAQ7KMYvzZrtjAEDPkmeTsjdHaWEtnNjNI23sSroqevIUhq7RqqfM5ZqYNTVEM4sg9w0yjotrFsPZA1InzOOuz3I4rTZ8KHkkfAd_pCj1DOI

ret = client.Write(fd84, []byte("T9UK9UltOg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (167) ggY26C3_r2cnuFByadCKwjL9xl7zKo4ZJ7aKwf6zdm0bGj1L_KAQ7KMYvzZrtjAEDPkmeTsjdHaWEtnNjNI23sSroqevIUhq7RqqfM5ZqYNTVEM4sg9w0yjotrFsPZA1InzOOuz3I4rTZ8KHkkfAd_pCj1DOIT9UK9UltOg

ret = client.Seek(fd99, 159, client.SEEK_SET)
if(ret != 159) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 159)
  panic("seek failed")
}

//fd state: (159) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOnTulIVTsfJsDkf5EfXwHt2JCV0tXQ1a3B5vdpArqfvVKGdvwm9CUcFmvHPidb7lbyhVYcV3qwK9khVcOeLcSFAWYQzxxk6b0SK

ret = client.Write(fd99, []byte("bT1F6Xe8v8sYLjIs7nni8QzkJMP_Y681J6AsyOeB97jIxz7JJI2LnXtedmq2BQetS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (224) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOnTulIVTsfJsDkf5EfXwHt2JCV0tXQ1a3B5vdpArqfvVKGdvwm9CUcFmvHPidb7lbyhVYcV3qwK9khVcOeLcSFAbT1F6Xe8v8sYLjIs7nni8QzkJMP_Y681J6AsyOeB97jIxz7JJI2LnXtedmq2BQetS

buf, ret = client.Read(fd102, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd106, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


buf, ret = client.Read(fd102, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd106, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd104)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd84)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd105, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


fd107 := client.Open("/LZcAyv_xCm.txt", client.O_RDWR|client.O_CREATE)
if(fd107 < 0) {
  panic("open failed")
}


fd108 := client.Open("/3v2LLkR1Z4.txt", client.O_RDWR|client.O_CREATE)
if(fd108 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd105, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "I8Ujy6WBFnHgX8XNp4FzL97YJ01GhBFDVxJHXkpME6zlKmwrxri") {
  panic("wrong data returned")
}


fd109 := client.Open("/uEvYTullIO.txt", client.O_RDWR|client.O_CREATE)
if(fd109 < 0) {
  panic("open failed")
}


fd110 := client.Open("/zc2Ij4WBct.txt", client.O_RDWR|client.O_CREATE)
if(fd110 < 0) {
  panic("open failed")
}


ret = client.Seek(fd107, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


ret = client.Close(fd105)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd109, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd110, []byte("NAue8HPlFwPMz7wYpI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) NAue8HPlFwPMz7wYpI

fd111 := client.Open("/4SS3LdR3oc.txt", client.O_RDWR|client.O_CREATE)
if(fd111 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd106, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "g3yfsNM") {
  panic("wrong data returned")
}


ret = client.Close(fd110)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd111)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd102, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}

//fd state: (14) SQNlLo8g3yfsNMAHFK30GSiJnv7sIK4JCG_uTZ5FY5CfMLJ

ret = client.Write(fd106, []byte("WPbW0aMO9B1PlCHL66rpIHtPtfSdAdtCkofreovbieRSTsMX_xa1kXEkpXByQUVjD8sJtj5bw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) SQNlLo8g3yfsNMWPbW0aMO9B1PlCHL66rpIHtPtfSdAdtCkofreovbieRSTsMX_xa1kXEkpXByQUVjD8sJtj5bw

ret = client.Seek(fd102, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (224) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOnTulIVTsfJsDkf5EfXwHt2JCV0tXQ1a3B5vdpArqfvVKGdvwm9CUcFmvHPidb7lbyhVYcV3qwK9khVcOeLcSFAbT1F6Xe8v8sYLjIs7nni8QzkJMP_Y681J6AsyOeB97jIxz7JJI2LnXtedmq2BQetS

ret = client.Write(fd99, []byte("Go9g9w3jy6KybRceGNJgQhErNM88vzb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (255) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOnTulIVTsfJsDkf5EfXwHt2JCV0tXQ1a3B5vdpArqfvVKGdvwm9CUcFmvHPidb7lbyhVYcV3qwK9khVcOeLcSFAbT1F6Xe8v8sYLjIs7nni8QzkJMP_Y681J6AsyOeB97jIxz7JJI2LnXtedmq2BQetSGo9g9w3jy6KybRceGNJgQhErNM88vzb
//fd state: (0) 

ret = client.Write(fd109, []byte("LR1SxTAWFo6F0cTfqeQn1wg1VttnpwfaT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) LR1SxTAWFo6F0cTfqeQn1wg1VttnpwfaT

ret = client.Close(fd102)
if(ret != 0) {
  panic("close failed")
}


fd112 := client.Open("/N2wW4FWwc1.txt", client.O_RDWR|client.O_CREATE)
if(fd112 < 0) {
  panic("open failed")
}


fd113 := client.Open("/obmRgRRVEq.txt", client.O_RDWR|client.O_CREATE)
if(fd113 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd97, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd112, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd109)
if(ret != 0) {
  panic("close failed")
}


fd114 := client.Open("/Kavr4vF_ms.txt", client.O_RDWR|client.O_CREATE)
if(fd114 < 0) {
  panic("open failed")
}


ret = client.Close(fd97)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd112, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (255) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOnTulIVTsfJsDkf5EfXwHt2JCV0tXQ1a3B5vdpArqfvVKGdvwm9CUcFmvHPidb7lbyhVYcV3qwK9khVcOeLcSFAbT1F6Xe8v8sYLjIs7nni8QzkJMP_Y681J6AsyOeB97jIxz7JJI2LnXtedmq2BQetSGo9g9w3jy6KybRceGNJgQhErNM88vzb

ret = client.Write(fd99, []byte("9LXyxQOoBziIhetbGnOVcKMyR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (280) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOnTulIVTsfJsDkf5EfXwHt2JCV0tXQ1a3B5vdpArqfvVKGdvwm9CUcFmvHPidb7lbyhVYcV3qwK9khVcOeLcSFAbT1F6Xe8v8sYLjIs7nni8QzkJMP_Y681J6AsyOeB97jIxz7JJI2LnXtedmq2BQetSGo9g9w3jy6KybRceGNJgQhErNM88vzb9LXyxQOoBziIhetbGnOVcKMyR
//fd state: (0) 

ret = client.Write(fd114, []byte("V2nrgBzCgPIG9y_MhUy9VXUPFNw7eTTT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) V2nrgBzCgPIG9y_MhUy9VXUPFNw7eTTT

ret = client.Seek(fd114, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


ret = client.Close(fd114)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd99, 278, client.SEEK_SET)
if(ret != 278) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 278)
  panic("seek failed")
}

//fd state: (87) SQNlLo8g3yfsNMWPbW0aMO9B1PlCHL66rpIHtPtfSdAdtCkofreovbieRSTsMX_xa1kXEkpXByQUVjD8sJtj5bw

ret = client.Write(fd106, []byte("r56z91wBifhIgBaxg2kdSMyv3uroAZt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) SQNlLo8g3yfsNMWPbW0aMO9B1PlCHL66rpIHtPtfSdAdtCkofreovbieRSTsMX_xa1kXEkpXByQUVjD8sJtj5bwr56z91wBifhIgBaxg2kdSMyv3uroAZt

ret = client.Seek(fd106, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd113, []byte("vXy0QlFl0FmetWWtY9hmCCI16AFN0fmkAwkU9cXWL8OjcV5gsCD9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) vXy0QlFl0FmetWWtY9hmCCI16AFN0fmkAwkU9cXWL8OjcV5gsCD9
//fd state: (278) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOnTulIVTsfJsDkf5EfXwHt2JCV0tXQ1a3B5vdpArqfvVKGdvwm9CUcFmvHPidb7lbyhVYcV3qwK9khVcOeLcSFAbT1F6Xe8v8sYLjIs7nni8QzkJMP_Y681J6AsyOeB97jIxz7JJI2LnXtedmq2BQetSGo9g9w3jy6KybRceGNJgQhErNM88vzb9LXyxQOoBziIhetbGnOVcKMyR

ret = client.Write(fd99, []byte("VE_K0wyaCOZASU5VDhkiSGc4dI7K66NxwaNl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (314) QlfFP7huvYDSYEcW4Bh6occOSgvmWvUUeEKwpVTUJTATU3zA1K7lORLNGuLenYhmLOzIWoqyOnTulIVTsfJsDkf5EfXwHt2JCV0tXQ1a3B5vdpArqfvVKGdvwm9CUcFmvHPidb7lbyhVYcV3qwK9khVcOeLcSFAbT1F6Xe8v8sYLjIs7nni8QzkJMP_Y681J6AsyOeB97jIxz7JJI2LnXtedmq2BQetSGo9g9w3jy6KybRceGNJgQhErNM88vzb9LXyxQOoBziIhetbGnOVcKMVE_K0wyaCOZASU5VDhkiSGc4dI7K66NxwaNl

ret = client.Seek(fd106, 105, client.SEEK_SET)
if(ret != 105) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 105)
  panic("seek failed")
}


ret = client.Close(fd106)
if(ret != 0) {
  panic("close failed")
}


fd115 := client.Open("/9RkaPG0bm2.txt", client.O_RDWR|client.O_CREATE)
if(fd115 < 0) {
  panic("open failed")
}


ret = client.Close(fd108)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd99, 87, client.SEEK_SET)
if(ret != 87) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 87)
  panic("seek failed")
}

//fd state: (19) fmpftJInZqnKnAARukBubzk7ZFGPtEVvflJpGQdkNclOxzccPIoNvZbiLctYYz1zjrVGOcQ6zt2pcUFT7VvzZcrSOhWKcWeKq

ret = client.Write(fd107, []byte("xvyJsMK92fY64w3ZykwDdmS7O"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) fmpftJInZqnKnAARukBxvyJsMK92fY64w3ZykwDdmS7OxzccPIoNvZbiLctYYz1zjrVGOcQ6zt2pcUFT7VvzZcrSOhWKcWeKq

ret = client.Close(fd99)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd113, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd112, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd107)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd112, []byte("_Rjx1aHCYC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) _Rjx1aHCYC

ret = client.Close(fd113)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd115, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd116 := client.Open("/_fKtkmMG2H.txt", client.O_RDWR|client.O_CREATE)
if(fd116 < 0) {
  panic("open failed")
}

//fd state: (10) _Rjx1aHCYC

ret = client.Write(fd112, []byte("5VPTMdjbqiIGiI4FNm0mnmeZ2nOCpmb4AJzJm_syzMoxgaVttiG98Hfczn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) _Rjx1aHCYC5VPTMdjbqiIGiI4FNm0mnmeZ2nOCpmb4AJzJm_syzMoxgaVttiG98Hfczn

ret = client.Close(fd116)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd115)
if(ret != 0) {
  panic("close failed")
}


fd117 := client.Open("/8W6113_G3d.txt", client.O_RDWR|client.O_CREATE)
if(fd117 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd112, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd118 := client.Open("/finUwiQx86.txt", client.O_RDWR|client.O_CREATE)
if(fd118 < 0) {
  panic("open failed")
}


fd119 := client.Open("/JmFn4HpF0R.txt", client.O_RDWR|client.O_CREATE)
if(fd119 < 0) {
  panic("open failed")
}


ret = client.Close(fd118)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd119)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd112, 17)
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


ret = client.Close(fd112)
if(ret != 0) {
  panic("close failed")
}


fd120 := client.Open("/obmRgRRVEq.txt", client.O_RDWR|client.O_CREATE)
if(fd120 < 0) {
  panic("open failed")
}


fd121 := client.Open("/Ejqw2f8dM3.txt", client.O_RDWR|client.O_CREATE)
if(fd121 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd121, []byte("kMi9tUFXp3GFmTPHZ90biPHIkWwWOOhbPP_WglZQt1dDej1e"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) kMi9tUFXp3GFmTPHZ90biPHIkWwWOOhbPP_WglZQt1dDej1e

fd122 := client.Open("/6mwWpmkirH.txt", client.O_RDWR|client.O_CREATE)
if(fd122 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd122, []byte("6SgFpM8CYRQ1hF7ZCIF74OXogXwATCDq0YW7EY4PTyV57p9PFyBhbGEydZEcrw23eoUdxBFsbOPp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) 6SgFpM8CYRQ1hF7ZCIF74OXogXwATCDq0YW7EY4PTyV57p9PFyBhbGEydZEcrw23eoUdxBFsbOPp
//fd state: (0) vXy0QlFl0FmetWWtY9hmCCI16AFN0fmkAwkU9cXWL8OjcV5gsCD9

ret = client.Write(fd120, []byte("Ib6h1t3UzpuJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) Ib6h1t3UzpuJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848K
//fd state: (58) Ib6h1t3UzpuJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848K

ret = client.Write(fd120, []byte("Fs5HnTr67WmC33CjyrvDQdh2OtnwyjYWwl1TQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) Ib6h1t3UzpuJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848KFs5HnTr67WmC33CjyrvDQdh2OtnwyjYWwl1TQ

ret = client.Seek(fd122, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


ret = client.Close(fd122)
if(ret != 0) {
  panic("close failed")
}


fd123 := client.Open("/2MKMM17MVW.txt", client.O_RDWR|client.O_CREATE)
if(fd123 < 0) {
  panic("open failed")
}


ret = client.Seek(fd120, 64, client.SEEK_SET)
if(ret != 64) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 64)
  panic("seek failed")
}


buf, ret = client.Read(fd121, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd123, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (64) Ib6h1t3UzpuJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848KFs5HnTr67WmC33CjyrvDQdh2OtnwyjYWwl1TQ

ret = client.Write(fd120, []byte("K2gJ6hjt2YX0HFktN8fRAEuxrNmZUF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) Ib6h1t3UzpuJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848KFs5HnTK2gJ6hjt2YX0HFktN8fRAEuxrNmZUFQ

buf, ret = client.Read(fd123, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd124 := client.Open("/D7WnAKqAcn.txt", client.O_RDWR|client.O_CREATE)
if(fd124 < 0) {
  panic("open failed")
}


ret = client.Seek(fd120, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd124, []byte("Oath0iNUxu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) Oath0iNUxu

fd125 := client.Open("/3OuuWd5hLc.txt", client.O_RDWR|client.O_CREATE)
if(fd125 < 0) {
  panic("open failed")
}


fd126 := client.Open("/W4nJjAH1FT.txt", client.O_RDWR|client.O_CREATE)
if(fd126 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd123, []byte("afce4oSimmR5gAI9GN87KJlQnhmhe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) afce4oSimmR5gAI9GN87KJlQnhmhe

ret = client.Seek(fd125, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd127 := client.Open("/hofqG9YWVA.txt", client.O_RDWR|client.O_CREATE)
if(fd127 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd121, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd121, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


ret = client.Close(fd127)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd123, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd126, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd123, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (10) Oath0iNUxu

ret = client.Write(fd124, []byte("K7HJY4mVn57zabcrZ_Tf03CbHyM3rY2ICWvdMKRwZleeVn0WTamB0wgt0QplC6DMq6u_y5me0lAtuLuGV_kLhtf7yPWtHCqgj0j"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (109) Oath0iNUxuK7HJY4mVn57zabcrZ_Tf03CbHyM3rY2ICWvdMKRwZleeVn0WTamB0wgt0QplC6DMq6u_y5me0lAtuLuGV_kLhtf7yPWtHCqgj0j

buf, ret = client.Read(fd121, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PHIk") {
  panic("wrong data returned")
}


ret = client.Close(fd124)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd123)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd125)
if(ret != 0) {
  panic("close failed")
}


fd128 := client.Open("/DZlV8Pcptl.txt", client.O_RDWR|client.O_CREATE)
if(fd128 < 0) {
  panic("open failed")
}


ret = client.Seek(fd120, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}


ret = client.Close(fd121)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd128, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd129 := client.Open("/MVWR6bEOKx.txt", client.O_RDWR|client.O_CREATE)
if(fd129 < 0) {
  panic("open failed")
}


ret = client.Close(fd126)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd128, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd128, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd130 := client.Open("/BU8ZeUXfGw.txt", client.O_RDWR|client.O_CREATE)
if(fd130 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd128, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd130, []byte("dG503MzkJ1iImHSU9w8c7yp1MkbfC0SEsFaty_LKLUGNNkeCANNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV45g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) dG503MzkJ1iImHSU9w8c7yp1MkbfC0SEsFaty_LKLUGNNkeCANNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV45g

fd131 := client.Open("/nYDyf3ULIn.txt", client.O_RDWR|client.O_CREATE)
if(fd131 < 0) {
  panic("open failed")
}

//fd state: (91) dG503MzkJ1iImHSU9w8c7yp1MkbfC0SEsFaty_LKLUGNNkeCANNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV45g

ret = client.Write(fd130, []byte("KkdLdJoLo6nJZBFwSdLobuJBt3Q3KNt1mq5rdwv7XvRnEmWlafABfoihfjUdOepJeHIs_DrGoih3M"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (168) dG503MzkJ1iImHSU9w8c7yp1MkbfC0SEsFaty_LKLUGNNkeCANNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV45gKkdLdJoLo6nJZBFwSdLobuJBt3Q3KNt1mq5rdwv7XvRnEmWlafABfoihfjUdOepJeHIs_DrGoih3M

buf, ret = client.Read(fd130, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd120, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jt2YX0HFktN8fRAEuxrN") {
  panic("wrong data returned")
}


ret = client.Seek(fd129, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd131)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd130, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd130, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd128, []byte("H8y2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) H8y2

buf, ret = client.Read(fd128, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (90) Ib6h1t3UzpuJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848KFs5HnTK2gJ6hjt2YX0HFktN8fRAEuxrNmZUFQ

ret = client.Write(fd120, []byte("Eh3v4G_28ADSi5CZvYhDbqGzVgYNDA7kByhYDEjuzQa6ns9fWSeI9eKnHZDoFqpLY7FNCD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (160) Ib6h1t3UzpuJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848KFs5HnTK2gJ6hjt2YX0HFktN8fRAEuxrNEh3v4G_28ADSi5CZvYhDbqGzVgYNDA7kByhYDEjuzQa6ns9fWSeI9eKnHZDoFqpLY7FNCD

buf, ret = client.Read(fd129, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd120)
if(ret != 0) {
  panic("close failed")
}

//fd state: (4) H8y2

ret = client.Write(fd128, []byte("23do3XvYUFCgtTk8k2ByAq85K41fn5kgul7TY8wfd8NBDuPN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) H8y223do3XvYUFCgtTk8k2ByAq85K41fn5kgul7TY8wfd8NBDuPN

buf, ret = client.Read(fd129, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd130, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV45gKkdLdJoLo6nJZBFwSdLobuJBt3Q3KN") {
  panic("wrong data returned")
}


fd132 := client.Open("/TqUxY9FOiQ.txt", client.O_RDWR|client.O_CREATE)
if(fd132 < 0) {
  panic("open failed")
}


fd133 := client.Open("/ZaBfLNx_Cb.txt", client.O_RDWR|client.O_CREATE)
if(fd133 < 0) {
  panic("open failed")
}


ret = client.Close(fd128)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd133, []byte("_GlXZu3tWoM5LxHlT0rOxHSwklS5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) _GlXZu3tWoM5LxHlT0rOxHSwklS5
//fd state: (121) dG503MzkJ1iImHSU9w8c7yp1MkbfC0SEsFaty_LKLUGNNkeCANNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV45gKkdLdJoLo6nJZBFwSdLobuJBt3Q3KNt1mq5rdwv7XvRnEmWlafABfoihfjUdOepJeHIs_DrGoih3M

ret = client.Write(fd130, []byte("eawTQdk5__Ud6exoEvUEhMcwKcuz2b4cQRGZc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (158) dG503MzkJ1iImHSU9w8c7yp1MkbfC0SEsFaty_LKLUGNNkeCANNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV45gKkdLdJoLo6nJZBFwSdLobuJBt3Q3KNeawTQdk5__Ud6exoEvUEhMcwKcuz2b4cQRGZcs_DrGoih3M
//fd state: (0) 

ret = client.Write(fd132, []byte("6maJpjgkZmquT2zM2_t_adIh9sTc4RCmunB42N"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) 6maJpjgkZmquT2zM2_t_adIh9sTc4RCmunB42N

ret = client.Close(fd129)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd130)
if(ret != 0) {
  panic("close failed")
}


fd134 := client.Open("/HZVjcmUfrh.txt", client.O_RDWR|client.O_CREATE)
if(fd134 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd133, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd134, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd134, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd133, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd134, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd132, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


ret = client.Close(fd133)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd134, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd134, []byte("RALFrzldIvebb19XxX5Ee3slm0y0Wa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) RALFrzldIvebb19XxX5Ee3slm0y0Wa
//fd state: (21) 6maJpjgkZmquT2zM2_t_adIh9sTc4RCmunB42N

ret = client.Write(fd132, []byte("KJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1Pp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1Pp

ret = client.Close(fd134)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd132, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (83) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1Pp

ret = client.Write(fd132, []byte("NbIn4OilxK7ZMIYSf0rXYJSN1HdwDzk8E6tIIwxKb1ITQX2qQUX150Y61DS_sCXIom"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (149) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNbIn4OilxK7ZMIYSf0rXYJSN1HdwDzk8E6tIIwxKb1ITQX2qQUX150Y61DS_sCXIom

fd135 := client.Open("/fT8o_szxtr.txt", client.O_RDWR|client.O_CREATE)
if(fd135 < 0) {
  panic("open failed")
}


ret = client.Seek(fd135, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd135)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd132)
if(ret != 0) {
  panic("close failed")
}


fd136 := client.Open("/f_VMh3LK24.txt", client.O_RDWR|client.O_CREATE)
if(fd136 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd136, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd136, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd136)
if(ret != 0) {
  panic("close failed")
}


fd137 := client.Open("/8jCWuWK8ZO.txt", client.O_RDWR|client.O_CREATE)
if(fd137 < 0) {
  panic("open failed")
}


fd138 := client.Open("/pkf0oZfq4O.txt", client.O_RDWR|client.O_CREATE)
if(fd138 < 0) {
  panic("open failed")
}


fd139 := client.Open("/4mPDBgaYe1.txt", client.O_RDWR|client.O_CREATE)
if(fd139 < 0) {
  panic("open failed")
}


fd140 := client.Open("/oOWQmAOjGG.txt", client.O_RDWR|client.O_CREATE)
if(fd140 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd138, []byte("029OHqnyGpcoWCcQ52ihb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) 029OHqnyGpcoWCcQ52ihb

fd141 := client.Open("/9RkaPG0bm2.txt", client.O_RDWR|client.O_CREATE)
if(fd141 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd138, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd139, []byte("d532dkIAfcRzT0Sx5WZCMO6IwpxAC3vXmG57uL7VvUlfz7F4SNshpwaZ4TbunhWas8Y3pRM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (71) d532dkIAfcRzT0Sx5WZCMO6IwpxAC3vXmG57uL7VvUlfz7F4SNshpwaZ4TbunhWas8Y3pRM

buf, ret = client.Read(fd141, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd137)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd138)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd141, []byte("y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) y
//fd state: (0) 

ret = client.Write(fd140, []byte("V007DCkh4NmaQh9LFPApNyQEYLvoGuj_IuBFSvUSrQQfaUo9JValqLORE9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) V007DCkh4NmaQh9LFPApNyQEYLvoGuj_IuBFSvUSrQQfaUo9JValqLORE9

ret = client.Seek(fd141, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Close(fd140)
if(ret != 0) {
  panic("close failed")
}


fd142 := client.Open("/A2jhBeQbpL.txt", client.O_RDWR|client.O_CREATE)
if(fd142 < 0) {
  panic("open failed")
}

//fd state: (1) y

ret = client.Write(fd141, []byte("jcSRWphita7zsQjG1ARVfBkJEMlXxCJ_TReImP7rw0VWed_1pYQijWHTcpWKcQdodWNYcc9lzNF5PmC1TCwl6FEYsb8tho_xmL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) yjcSRWphita7zsQjG1ARVfBkJEMlXxCJ_TReImP7rw0VWed_1pYQijWHTcpWKcQdodWNYcc9lzNF5PmC1TCwl6FEYsb8tho_xmL

fd143 := client.Open("/qinn9z7zrr.txt", client.O_RDWR|client.O_CREATE)
if(fd143 < 0) {
  panic("open failed")
}


ret = client.Seek(fd142, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd141, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Seek(fd142, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd143)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd139, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd144 := client.Open("/nFxpEbR3hh.txt", client.O_RDWR|client.O_CREATE)
if(fd144 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd141, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "zsQjG1ARVfBkJEMlXxCJ_TReImP7rw0VWed_") {
  panic("wrong data returned")
}


ret = client.Close(fd141)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd142, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd145 := client.Open("/zZt3zhX_S8.txt", client.O_RDWR|client.O_CREATE)
if(fd145 < 0) {
  panic("open failed")
}


ret = client.Seek(fd144, 182, client.SEEK_SET)
if(ret != 182) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 182)
  panic("seek failed")
}


buf, ret = client.Read(fd142, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd142)
if(ret != 0) {
  panic("close failed")
}


fd146 := client.Open("/MaHqGlDtgN.txt", client.O_RDWR|client.O_CREATE)
if(fd146 < 0) {
  panic("open failed")
}


ret = client.Seek(fd144, 65, client.SEEK_SET)
if(ret != 65) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 65)
  panic("seek failed")
}

//fd state: (65) WWW9jD0wWvFE8yU7Q8LSOQxkgh6m4Fz4VOVwiAkY93ClyguqNko2L3_943IqXr2xip5VUGqWzYLGzS1Mc06SQjI4h6R9epgcvuQOMLe4HZZHB0crNWl0IdARg68mv2hpnc0fNJmlLFGEtlr1t7ED9RoGcVwdsykA2vGr1jNpAJB3rJO8Mojl9_tcIm

ret = client.Write(fd144, []byte("kngt2iLjWAglzQBBfBqaH0Aadwx7SRWSHOHfFY7ZxZiDF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) WWW9jD0wWvFE8yU7Q8LSOQxkgh6m4Fz4VOVwiAkY93ClyguqNko2L3_943IqXr2xikngt2iLjWAglzQBBfBqaH0Aadwx7SRWSHOHfFY7ZxZiDFcrNWl0IdARg68mv2hpnc0fNJmlLFGEtlr1t7ED9RoGcVwdsykA2vGr1jNpAJB3rJO8Mojl9_tcIm

buf, ret = client.Read(fd139, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd146)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd145, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd147 := client.Open("/7Yc4dPLSSJ.txt", client.O_RDWR|client.O_CREATE)
if(fd147 < 0) {
  panic("open failed")
}


ret = client.Seek(fd144, 92, client.SEEK_SET)
if(ret != 92) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 92)
  panic("seek failed")
}


ret = client.Close(fd139)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd147)
if(ret != 0) {
  panic("close failed")
}


fd148 := client.Open("/RUBg7QMMiO.txt", client.O_RDWR|client.O_CREATE)
if(fd148 < 0) {
  panic("open failed")
}


ret = client.Close(fd144)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd148)
if(ret != 0) {
  panic("close failed")
}


fd149 := client.Open("/D7WnAKqAcn.txt", client.O_RDWR|client.O_CREATE)
if(fd149 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd145, []byte("rxjTd1b8GxJWIATdOlyCe5zzmTcIvGi2VSh876DnU2UG7C4slgN60eVK2B7J1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) rxjTd1b8GxJWIATdOlyCe5zzmTcIvGi2VSh876DnU2UG7C4slgN60eVK2B7J1

ret = client.Close(fd149)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd145, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd145, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}

//fd state: (35) rxjTd1b8GxJWIATdOlyCe5zzmTcIvGi2VSh876DnU2UG7C4slgN60eVK2B7J1

ret = client.Write(fd145, []byte("RN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) rxjTd1b8GxJWIATdOlyCe5zzmTcIvGi2VShRN6DnU2UG7C4slgN60eVK2B7J1

ret = client.Seek(fd145, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


ret = client.Seek(fd145, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


buf, ret = client.Read(fd145, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "cIvGi2") {
  panic("wrong data returned")
}


ret = client.Seek(fd145, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


fd150 := client.Open("/HZVjcmUfrh.txt", client.O_RDWR|client.O_CREATE)
if(fd150 < 0) {
  panic("open failed")
}


fd151 := client.Open("/HRbgk5YrAz.txt", client.O_RDWR|client.O_CREATE)
if(fd151 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd151, []byte("dJ4qMyPqGAsZzT4HBKwlfG5HbDLftkMlgzavgqKhoyRGGDoi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) dJ4qMyPqGAsZzT4HBKwlfG5HbDLftkMlgzavgqKhoyRGGDoi

ret = client.Close(fd145)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd151, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd152 := client.Open("/obmRgRRVEq.txt", client.O_RDWR|client.O_CREATE)
if(fd152 < 0) {
  panic("open failed")
}


ret = client.Seek(fd150, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) Ib6h1t3UzpuJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848KFs5HnTK2gJ6hjt2YX0HFktN8fRAEuxrNEh3v4G_28ADSi5CZvYhDbqGzVgYNDA7kByhYDEjuzQa6ns9fWSeI9eKnHZDoFqpLY7FNCD

ret = client.Write(fd152, []byte("ouGfKY8mYsk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) ouGfKY8mYskJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848KFs5HnTK2gJ6hjt2YX0HFktN8fRAEuxrNEh3v4G_28ADSi5CZvYhDbqGzVgYNDA7kByhYDEjuzQa6ns9fWSeI9eKnHZDoFqpLY7FNCD

ret = client.Seek(fd152, 123, client.SEEK_SET)
if(ret != 123) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 123)
  panic("seek failed")
}


ret = client.Seek(fd151, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


fd153 := client.Open("/K1q8zuNqur.txt", client.O_RDWR|client.O_CREATE)
if(fd153 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd153, []byte("bADDD6hFBrkBtYr6ILPQr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) bADDD6hFBrkBtYr6ILPQr

fd154 := client.Open("/2JrRL5XTCk.txt", client.O_RDWR|client.O_CREATE)
if(fd154 < 0) {
  panic("open failed")
}

//fd state: (0) FpBOmAxJXZiOuRzFbtk2hl3geUWlgDTIo_8TAWvSmRVynNxasTH2atoRA_C2Y1ISabFhD762dVfmIE9

ret = client.Write(fd154, []byte("kGJEtE_LReqao0aOZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) kGJEtE_LReqao0aOZtk2hl3geUWlgDTIo_8TAWvSmRVynNxasTH2atoRA_C2Y1ISabFhD762dVfmIE9

buf, ret = client.Read(fd152, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yhYDEjuzQa6ns9fWSeI9eKnHZDoFqpLY7FNCD") {
  panic("wrong data returned")
}


fd155 := client.Open("/_CVl1cUh_V.txt", client.O_RDWR|client.O_CREATE)
if(fd155 < 0) {
  panic("open failed")
}


fd156 := client.Open("/Ex7w4vRt6O.txt", client.O_RDWR|client.O_CREATE)
if(fd156 < 0) {
  panic("open failed")
}


ret = client.Seek(fd154, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


ret = client.Close(fd151)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd156, []byte("GqLX5nqhbOhPpdK21wGcDfQ9PcSxjXbCATzEVIjON9F2UPG8yqQyIF21gt2RUsKuxox8GIan73Q7alUgZuA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) GqLX5nqhbOhPpdK21wGcDfQ9PcSxjXbCATzEVIjON9F2UPG8yqQyIF21gt2RUsKuxox8GIan73Q7alUgZuA

ret = client.Seek(fd152, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


fd157 := client.Open("/9WrO_P7NpN.txt", client.O_RDWR|client.O_CREATE)
if(fd157 < 0) {
  panic("open failed")
}


fd158 := client.Open("/QbUk9Mk7mu.txt", client.O_RDWR|client.O_CREATE)
if(fd158 < 0) {
  panic("open failed")
}


ret = client.Close(fd152)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd157)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd154)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd155)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd150)
if(ret != 0) {
  panic("close failed")
}

//fd state: (21) bADDD6hFBrkBtYr6ILPQr

ret = client.Write(fd153, []byte("FP_jkh7vpknRtqdgwJSB0An2iwXn7FZqWJy5boLze1lt3Yu8zNmgojxwBDiKQixqhuv4KSawP0NOgWaUUFrdK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (106) bADDD6hFBrkBtYr6ILPQrFP_jkh7vpknRtqdgwJSB0An2iwXn7FZqWJy5boLze1lt3Yu8zNmgojxwBDiKQixqhuv4KSawP0NOgWaUUFrdK

ret = client.Close(fd158)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd153)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd156)
if(ret != 0) {
  panic("close failed")
}


fd159 := client.Open("/DZlV8Pcptl.txt", client.O_RDWR|client.O_CREATE)
if(fd159 < 0) {
  panic("open failed")
}

//fd state: (0) H8y223do3XvYUFCgtTk8k2ByAq85K41fn5kgul7TY8wfd8NBDuPN

ret = client.Write(fd159, []byte("5_jKSax5Tg_rpZ1MwuY2fcw7AQXZ3hLIDLLYI8oVVvM3go5CB7k"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) 5_jKSax5Tg_rpZ1MwuY2fcw7AQXZ3hLIDLLYI8oVVvM3go5CB7kN

ret = client.Seek(fd159, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 5_jKSax5Tg_rpZ1MwuY2fcw7AQXZ3hLIDLLYI8oVVvM3go5CB7kN

ret = client.Write(fd159, []byte("RdJUpY78poUlbVAlyGkCZiFptMw9vMqPz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) RdJUpY78poUlbVAlyGkCZiFptMw9vMqPzLLYI8oVVvM3go5CB7kN

buf, ret = client.Read(fd159, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LLYI8oVVvM3go5CB7kN") {
  panic("wrong data returned")
}

//fd state: (52) RdJUpY78poUlbVAlyGkCZiFptMw9vMqPzLLYI8oVVvM3go5CB7kN

ret = client.Write(fd159, []byte("u5wDuSpM4fXnAmrEp30e9FN8qTy2zotFFdffRoXNlxVFXNhAMg5VAoEezsbvvefTKBRp2axxT0YYU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (129) RdJUpY78poUlbVAlyGkCZiFptMw9vMqPzLLYI8oVVvM3go5CB7kNu5wDuSpM4fXnAmrEp30e9FN8qTy2zotFFdffRoXNlxVFXNhAMg5VAoEezsbvvefTKBRp2axxT0YYU

fd160 := client.Open("/qSIF4GBaPH.txt", client.O_RDWR|client.O_CREATE)
if(fd160 < 0) {
  panic("open failed")
}


fd161 := client.Open("/W4nJjAH1FT.txt", client.O_RDWR|client.O_CREATE)
if(fd161 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd161, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd159)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd161, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd161, []byte("t2jc1Ex9YC8URWa2kUNvDRV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) t2jc1Ex9YC8URWa2kUNvDRV
//fd state: (23) t2jc1Ex9YC8URWa2kUNvDRV

ret = client.Write(fd161, []byte("n8uRB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) t2jc1Ex9YC8URWa2kUNvDRVn8uRB

ret = client.Seek(fd161, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


ret = client.Seek(fd160, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd162 := client.Open("/W4nJjAH1FT.txt", client.O_RDWR|client.O_CREATE)
if(fd162 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd160, []byte("nJDqsJJJSRkhtb6kPzqUJhaLIvANZ8cMJBGskdrtddCl_8KGie2EENJQlxAIHRhfkAl5oLPatclKh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) nJDqsJJJSRkhtb6kPzqUJhaLIvANZ8cMJBGskdrtddCl_8KGie2EENJQlxAIHRhfkAl5oLPatclKh

fd163 := client.Open("/zc2Ij4WBct.txt", client.O_RDWR|client.O_CREATE)
if(fd163 < 0) {
  panic("open failed")
}

//fd state: (27) t2jc1Ex9YC8URWa2kUNvDRVn8uRB

ret = client.Write(fd161, []byte("Nu2R7iNxjbl0UEtoCC0kqNz2VmkitU5f9TOntVXN3lq8pDz4rMcNCvwBlhnr_EjqR1enusssaVf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) t2jc1Ex9YC8URWa2kUNvDRVn8uRNu2R7iNxjbl0UEtoCC0kqNz2VmkitU5f9TOntVXN3lq8pDz4rMcNCvwBlhnr_EjqR1enusssaVf
//fd state: (0) NAue8HPlFwPMz7wYpI

ret = client.Write(fd163, []byte("pBhauuk_E1RwCof_LdU8QOyI8q2LgugjRu9w0MlnlcHcY6KzrUtfYoTzy8X2o"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) pBhauuk_E1RwCof_LdU8QOyI8q2LgugjRu9w0MlnlcHcY6KzrUtfYoTzy8X2o

fd164 := client.Open("/KqykqQdiXF.txt", client.O_RDWR|client.O_CREATE)
if(fd164 < 0) {
  panic("open failed")
}


ret = client.Seek(fd161, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Close(fd163)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd160)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd161, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


buf, ret = client.Read(fd164, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "XqeR5BQTtne3Y1V") {
  panic("wrong data returned")
}

//fd state: (15) XqeR5BQTtne3Y1V

ret = client.Write(fd164, []byte("EaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3Uo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) XqeR5BQTtne3Y1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3Uo

fd165 := client.Open("/k4d87VpxAe.txt", client.O_RDWR|client.O_CREATE)
if(fd165 < 0) {
  panic("open failed")
}


ret = client.Close(fd165)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd164)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd161, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


ret = client.Seek(fd162, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}

//fd state: (34) t2jc1Ex9YC8URWa2kUNvDRVn8uRNu2R7iNxjbl0UEtoCC0kqNz2VmkitU5f9TOntVXN3lq8pDz4rMcNCvwBlhnr_EjqR1enusssaVf

ret = client.Write(fd162, []byte("KiFbf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) t2jc1Ex9YC8URWa2kUNvDRVn8uRNu2R7iNKiFbfUEtoCC0kqNz2VmkitU5f9TOntVXN3lq8pDz4rMcNCvwBlhnr_EjqR1enusssaVf

ret = client.Seek(fd162, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}


fd166 := client.Open("/QbUk9Mk7mu.txt", client.O_RDWR|client.O_CREATE)
if(fd166 < 0) {
  panic("open failed")
}


ret = client.Seek(fd162, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


buf, ret = client.Read(fd166, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd167 := client.Open("/jSFBpvzIzk.txt", client.O_RDWR|client.O_CREATE)
if(fd167 < 0) {
  panic("open failed")
}


ret = client.Seek(fd166, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd161, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd166, []byte("CawTj8gqDjcxZqhqv8G"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) CawTj8gqDjcxZqhqv8G

fd168 := client.Open("/ZaBfLNx_Cb.txt", client.O_RDWR|client.O_CREATE)
if(fd168 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd167, []byte("2xgHENJtrX5tdaNhSUCZKDD2ySzdBjolxE3ipMsXC15W2Fmdu99zRd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) 2xgHENJtrX5tdaNhSUCZKDD2ySzdBjolxE3ipMsXC15W2Fmdu99zRd

buf, ret = client.Read(fd162, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "a2kUNvDRVn8uRNu2R7iNKiFbfUEtoCC0kqNz2VmkitU5f9TOntVXN3lq8pDz4rMcNCvw") {
  panic("wrong data returned")
}


ret = client.Close(fd167)
if(ret != 0) {
  panic("close failed")
}


fd169 := client.Open("/L3a7VdOTLE.txt", client.O_RDWR|client.O_CREATE)
if(fd169 < 0) {
  panic("open failed")
}


ret = client.Close(fd162)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd169, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd170 := client.Open("/hofqG9YWVA.txt", client.O_RDWR|client.O_CREATE)
if(fd170 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd161, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Vn8uRNu2R7") {
  panic("wrong data returned")
}


ret = client.Close(fd166)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd161, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd169, []byte("BRqIXriBgB99P5320YUjpZbfkhJUlC3QGm3AduTC8pfJ8ZArydOatPwoLqk0YOF6i6gbp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) BRqIXriBgB99P5320YUjpZbfkhJUlC3QGm3AduTC8pfJ8ZArydOatPwoLqk0YOF6i6gbp

fd171 := client.Open("/L3a7VdOTLE.txt", client.O_RDWR|client.O_CREATE)
if(fd171 < 0) {
  panic("open failed")
}


ret = client.Close(fd161)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd171, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd170, []byte("WbIe7mjcvy013VNcyhjE8HSQL82J0vsmAnzI52wH18Bv7FsBvB9IPDrNCcKEmgeG1TC4qFB_ANNhQqT87wvN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) WbIe7mjcvy013VNcyhjE8HSQL82J0vsmAnzI52wH18Bv7FsBvB9IPDrNCcKEmgeG1TC4qFB_ANNhQqT87wvN

buf, ret = client.Read(fd169, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (9) BRqIXriBgB99P5320YUjpZbfkhJUlC3QGm3AduTC8pfJ8ZArydOatPwoLqk0YOF6i6gbp

ret = client.Write(fd171, []byte("5LZvWiBjsk48riTc94JPun6lga_fhGvZzIle1C2nDX9om5eiV2yiVGXdk3m1HTNNrAgK5CHWfBsLB0Zn1rTuMT7wVOIl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) BRqIXriBg5LZvWiBjsk48riTc94JPun6lga_fhGvZzIle1C2nDX9om5eiV2yiVGXdk3m1HTNNrAgK5CHWfBsLB0Zn1rTuMT7wVOIl
//fd state: (101) BRqIXriBg5LZvWiBjsk48riTc94JPun6lga_fhGvZzIle1C2nDX9om5eiV2yiVGXdk3m1HTNNrAgK5CHWfBsLB0Zn1rTuMT7wVOIl

ret = client.Write(fd171, []byte("3huWdCCOtUAttXbGBjUEFzHSlxaaDxGhxtyOH8Ms_MduFLQofmJG0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) BRqIXriBg5LZvWiBjsk48riTc94JPun6lga_fhGvZzIle1C2nDX9om5eiV2yiVGXdk3m1HTNNrAgK5CHWfBsLB0Zn1rTuMT7wVOIl3huWdCCOtUAttXbGBjUEFzHSlxaaDxGhxtyOH8Ms_MduFLQofmJG0

fd172 := client.Open("/Qs7qyyBaX0.txt", client.O_RDWR|client.O_CREATE)
if(fd172 < 0) {
  panic("open failed")
}


ret = client.Close(fd169)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd171)
if(ret != 0) {
  panic("close failed")
}


fd173 := client.Open("/jg0Jvu7lJx.txt", client.O_RDWR|client.O_CREATE)
if(fd173 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd173, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd168, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


ret = client.Close(fd173)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd168, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


fd174 := client.Open("/KAn40utX3j.txt", client.O_RDWR|client.O_CREATE)
if(fd174 < 0) {
  panic("open failed")
}


fd175 := client.Open("/jEJT1wijAu.txt", client.O_RDWR|client.O_CREATE)
if(fd175 < 0) {
  panic("open failed")
}


ret = client.Close(fd172)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd175, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ev7_fh1MUOr5IrM8oYnIs0C56_m6cR2mZCGDWBt9") {
  panic("wrong data returned")
}


ret = client.Close(fd175)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd170)
if(ret != 0) {
  panic("close failed")
}

//fd state: (15) _GlXZu3tWoM5LxHlT0rOxHSwklS5

ret = client.Write(fd168, []byte("6IzUculJ2lbJAwDJ2Wl22btrhBpScrgR9P8F6Z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) _GlXZu3tWoM5LxH6IzUculJ2lbJAwDJ2Wl22btrhBpScrgR9P8F6Z
//fd state: (53) _GlXZu3tWoM5LxH6IzUculJ2lbJAwDJ2Wl22btrhBpScrgR9P8F6Z

ret = client.Write(fd168, []byte("eY_s78yY2InTBU8_mG7fySwLsVYAH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) _GlXZu3tWoM5LxH6IzUculJ2lbJAwDJ2Wl22btrhBpScrgR9P8F6ZeY_s78yY2InTBU8_mG7fySwLsVYAH
//fd state: (82) _GlXZu3tWoM5LxH6IzUculJ2lbJAwDJ2Wl22btrhBpScrgR9P8F6ZeY_s78yY2InTBU8_mG7fySwLsVYAH

ret = client.Write(fd168, []byte("fgUQ0RBx299sghdSy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) _GlXZu3tWoM5LxH6IzUculJ2lbJAwDJ2Wl22btrhBpScrgR9P8F6ZeY_s78yY2InTBU8_mG7fySwLsVYAHfgUQ0RBx299sghdSy

buf, ret = client.Read(fd168, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd174, []byte("GFoI48Jn_fk7HXS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) GFoI48Jn_fk7HXS

buf, ret = client.Read(fd174, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd174, 99)
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


fd176 := client.Open("/MyEfiMKv5b.txt", client.O_RDWR|client.O_CREATE)
if(fd176 < 0) {
  panic("open failed")
}


ret = client.Close(fd174)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd176, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd176, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) 

fd177 := client.Open("/6mwWpmkirH.txt", client.O_RDWR|client.O_CREATE)
if(fd177 < 0) {
  panic("open failed")
}


ret = client.Seek(fd176, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd176, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd176)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 6SgFpM8CYRQ1hF7ZCIF74OXogXwATCDq0YW7EY4PTyV57p9PFyBhbGEydZEcrw23eoUdxBFsbOPp

ret = client.Write(fd177, []byte("bWKhlV2c_K0YTuDmxI4SuAjmwAInvg8dUmVU3z3Y4p3n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) bWKhlV2c_K0YTuDmxI4SuAjmwAInvg8dUmVU3z3Y4p3n7p9PFyBhbGEydZEcrw23eoUdxBFsbOPp

fd178 := client.Open("/Cw5u_NagMb.txt", client.O_RDWR|client.O_CREATE)
if(fd178 < 0) {
  panic("open failed")
}


ret = client.Close(fd178)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd177)
if(ret != 0) {
  panic("close failed")
}


fd179 := client.Open("/jg0Jvu7lJx.txt", client.O_RDWR|client.O_CREATE)
if(fd179 < 0) {
  panic("open failed")
}


fd180 := client.Open("/Q7Imfuqduy.txt", client.O_RDWR|client.O_CREATE)
if(fd180 < 0) {
  panic("open failed")
}


ret = client.Close(fd179)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd180)
if(ret != 0) {
  panic("close failed")
}


fd181 := client.Open("/wZnvthnoX2.txt", client.O_RDWR|client.O_CREATE)
if(fd181 < 0) {
  panic("open failed")
}


ret = client.Seek(fd181, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd182 := client.Open("/obmRgRRVEq.txt", client.O_RDWR|client.O_CREATE)
if(fd182 < 0) {
  panic("open failed")
}


ret = client.Seek(fd181, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd181, []byte("ljjG6WCSkU5G"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) ljjG6WCSkU5G

ret = client.Seek(fd182, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}

//fd state: (12) ljjG6WCSkU5G

ret = client.Write(fd181, []byte("3Ddow7qXW726vUAHL5nHCieIrLwus2hDk23N5Q8yfBoRpb30tXrRo2KOjk5SXuonXj_myMAI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) ljjG6WCSkU5G3Ddow7qXW726vUAHL5nHCieIrLwus2hDk23N5Q8yfBoRpb30tXrRo2KOjk5SXuonXj_myMAI

ret = client.Seek(fd181, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


fd183 := client.Open("/dWsEcztvgd.txt", client.O_RDWR|client.O_CREATE)
if(fd183 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd183, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd184 := client.Open("/gwlhRQefvw.txt", client.O_RDWR|client.O_CREATE)
if(fd184 < 0) {
  panic("open failed")
}


ret = client.Seek(fd184, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd181, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


ret = client.Close(fd184)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd182)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd181, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


buf, ret = client.Read(fd183, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd185 := client.Open("/dkuo7HULBj.txt", client.O_RDWR|client.O_CREATE)
if(fd185 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd183, []byte("mtnBjnpjltHvUSag0u0m4b46YkcIS5xAXOopmNM32JzCGoW_TG5d5tGRRH7fqmKT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) mtnBjnpjltHvUSag0u0m4b46YkcIS5xAXOopmNM32JzCGoW_TG5d5tGRRH7fqmKT

ret = client.Close(fd181)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd185, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd183, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd183)
if(ret != 0) {
  panic("close failed")
}


fd186 := client.Open("/pLIWsilfjy.txt", client.O_RDWR|client.O_CREATE)
if(fd186 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd186, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd187 := client.Open("/dWsEcztvgd.txt", client.O_RDWR|client.O_CREATE)
if(fd187 < 0) {
  panic("open failed")
}


ret = client.Seek(fd185, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd185, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) mtnBjnpjltHvUSag0u0m4b46YkcIS5xAXOopmNM32JzCGoW_TG5d5tGRRH7fqmKT

ret = client.Write(fd187, []byte("vl6Oq4X8hLK5cnUgqlEwB6hJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) vl6Oq4X8hLK5cnUgqlEwB6hJYkcIS5xAXOopmNM32JzCGoW_TG5d5tGRRH7fqmKT

ret = client.Seek(fd185, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd185, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd186, []byte("qK4naYPq7iw_hMvnCD4pxBoiRwJmXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) qK4naYPq7iw_hMvnCD4pxBoiRwJmXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UB
//fd state: (24) vl6Oq4X8hLK5cnUgqlEwB6hJYkcIS5xAXOopmNM32JzCGoW_TG5d5tGRRH7fqmKT

ret = client.Write(fd187, []byte("5P1YAoLQQTqnK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) vl6Oq4X8hLK5cnUgqlEwB6hJ5P1YAoLQQTqnKNM32JzCGoW_TG5d5tGRRH7fqmKT

buf, ret = client.Read(fd187, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NM32JzCGoW_TG5d5tGRRH7fqmKT") {
  panic("wrong data returned")
}


ret = client.Close(fd187)
if(ret != 0) {
  panic("close failed")
}

//fd state: (91) qK4naYPq7iw_hMvnCD4pxBoiRwJmXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UB

ret = client.Write(fd186, []byte("JpYL903agqPgfa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) qK4naYPq7iw_hMvnCD4pxBoiRwJmXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UBJpYL903agqPgfa

ret = client.Close(fd185)
if(ret != 0) {
  panic("close failed")
}

//fd state: (105) qK4naYPq7iw_hMvnCD4pxBoiRwJmXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UBJpYL903agqPgfa

ret = client.Write(fd186, []byte("bhjNfq6rwRhscbXruOCeMpO7YqcF_tF5q8iSRheJPIGYZFwMCchFvcOUPu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (163) qK4naYPq7iw_hMvnCD4pxBoiRwJmXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UBJpYL903agqPgfabhjNfq6rwRhscbXruOCeMpO7YqcF_tF5q8iSRheJPIGYZFwMCchFvcOUPu

fd188 := client.Open("/Ju0L1mzSEa.txt", client.O_RDWR|client.O_CREATE)
if(fd188 < 0) {
  panic("open failed")
}


fd189 := client.Open("/3v2LLkR1Z4.txt", client.O_RDWR|client.O_CREATE)
if(fd189 < 0) {
  panic("open failed")
}


ret = client.Close(fd189)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd186, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (163) qK4naYPq7iw_hMvnCD4pxBoiRwJmXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UBJpYL903agqPgfabhjNfq6rwRhscbXruOCeMpO7YqcF_tF5q8iSRheJPIGYZFwMCchFvcOUPu

ret = client.Write(fd186, []byte("amzHVP3oCLHmKBr71EEl2Wb8jUXXif2arhF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (198) qK4naYPq7iw_hMvnCD4pxBoiRwJmXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UBJpYL903agqPgfabhjNfq6rwRhscbXruOCeMpO7YqcF_tF5q8iSRheJPIGYZFwMCchFvcOUPuamzHVP3oCLHmKBr71EEl2Wb8jUXXif2arhF

ret = client.Close(fd186)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd188)
if(ret != 0) {
  panic("close failed")
}


fd190 := client.Open("/6zCWkaOFR7.txt", client.O_RDWR|client.O_CREATE)
if(fd190 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd190, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd190, []byte("kLkORRlHF6vzmbUosRu7kPpM7dLVG7pLofh1UU8zPkreYH_kGnzi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) kLkORRlHF6vzmbUosRu7kPpM7dLVG7pLofh1UU8zPkreYH_kGnzi

fd191 := client.Open("/nUtlP4AKUi.txt", client.O_RDWR|client.O_CREATE)
if(fd191 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd190, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd192 := client.Open("/Ti2nt12Sl8.txt", client.O_RDWR|client.O_CREATE)
if(fd192 < 0) {
  panic("open failed")
}


ret = client.Seek(fd192, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (52) kLkORRlHF6vzmbUosRu7kPpM7dLVG7pLofh1UU8zPkreYH_kGnzi

ret = client.Write(fd190, []byte("QhQTdukZpuct"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) kLkORRlHF6vzmbUosRu7kPpM7dLVG7pLofh1UU8zPkreYH_kGnziQhQTdukZpuct

ret = client.Close(fd190)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd192)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd191)
if(ret != 0) {
  panic("close failed")
}


fd193 := client.Open("/49yoWYxM6B.txt", client.O_RDWR|client.O_CREATE)
if(fd193 < 0) {
  panic("open failed")
}


ret = client.Seek(fd193, 258, client.SEEK_SET)
if(ret != 258) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 258)
  panic("seek failed")
}


fd194 := client.Open("/YWJtqDTpBK.txt", client.O_RDWR|client.O_CREATE)
if(fd194 < 0) {
  panic("open failed")
}


ret = client.Close(fd194)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd193, 168, client.SEEK_SET)
if(ret != 168) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 168)
  panic("seek failed")
}


ret = client.Seek(fd193, 164, client.SEEK_SET)
if(ret != 164) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 164)
  panic("seek failed")
}


buf, ret = client.Read(fd193, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YCW5") {
  panic("wrong data returned")
}


ret = client.Close(fd193)
if(ret != 0) {
  panic("close failed")
}


fd195 := client.Open("/yGbP2_Vl8s.txt", client.O_RDWR|client.O_CREATE)
if(fd195 < 0) {
  panic("open failed")
}


ret = client.Seek(fd195, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd196 := client.Open("/9WrO_P7NpN.txt", client.O_RDWR|client.O_CREATE)
if(fd196 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd195, []byte("BcY7o5cJ15AMnrFwvThl8JB9TG1Zf056xtX3X71R3acNjwXKZWjQkJfCntXn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) BcY7o5cJ15AMnrFwvThl8JB9TG1Zf056xtX3X71R3acNjwXKZWjQkJfCntXn

ret = client.Close(fd195)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) EqGbKfMvDCSK1ADCrOZlRvUOMbFrFj9inz9fakpXHtAjifdi6IfmAzxovYi_qX38Jk_gteEficrACotQS5Kpb_QQbAoqZI8jTX

ret = client.Write(fd196, []byte("UJ56uaJ8HcLk64SvJivYqgJx8BjTQOR0BfsrjY3D_rnUkW01CeI7bTeUlmCrlcgEg2kRiL5zejtl231aSbAxYWCgnpVGG7yE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) UJ56uaJ8HcLk64SvJivYqgJx8BjTQOR0BfsrjY3D_rnUkW01CeI7bTeUlmCrlcgEg2kRiL5zejtl231aSbAxYWCgnpVGG7yETX

ret = client.Seek(fd196, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


fd197 := client.Open("/hofqG9YWVA.txt", client.O_RDWR|client.O_CREATE)
if(fd197 < 0) {
  panic("open failed")
}


ret = client.Seek(fd197, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


ret = client.Close(fd197)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd196, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "CeI7bTeUlm") {
  panic("wrong data returned")
}


ret = client.Close(fd196)
if(ret != 0) {
  panic("close failed")
}


fd198 := client.Open("/vlNW582dxT.txt", client.O_RDWR|client.O_CREATE)
if(fd198 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd198, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd198, []byte("pYuU_M_qEC2pf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) pYuU_M_qEC2pf

ret = client.Close(fd198)
if(ret != 0) {
  panic("close failed")
}


fd199 := client.Open("/95Up4yEOZw.txt", client.O_RDWR|client.O_CREATE)
if(fd199 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd199, []byte("xMn20jjWuxZalfXbyiy7QazwxIxtQX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) xMn20jjWuxZalfXbyiy7QazwxIxtQX
//fd state: (30) xMn20jjWuxZalfXbyiy7QazwxIxtQX

ret = client.Write(fd199, []byte("SfO0_ynxlnzWFz07pgPKHacDW7xZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) xMn20jjWuxZalfXbyiy7QazwxIxtQXSfO0_ynxlnzWFz07pgPKHacDW7xZ

ret = client.Seek(fd199, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


buf, ret = client.Read(fd199, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "QazwxIxtQXSfO0_ynxlnzWFz07") {
  panic("wrong data returned")
}

//fd state: (46) xMn20jjWuxZalfXbyiy7QazwxIxtQXSfO0_ynxlnzWFz07pgPKHacDW7xZ

ret = client.Write(fd199, []byte("44nlvJkPNhc1TVbuA2mhEKunxqy62chr3xsCn0UHkBlfQg0B2MJJNVypbvziH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) xMn20jjWuxZalfXbyiy7QazwxIxtQXSfO0_ynxlnzWFz0744nlvJkPNhc1TVbuA2mhEKunxqy62chr3xsCn0UHkBlfQg0B2MJJNVypbvziH

ret = client.Close(fd199)
if(ret != 0) {
  panic("close failed")
}


fd200 := client.Open("/6Okbu3K6Za.txt", client.O_RDWR|client.O_CREATE)
if(fd200 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd200, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd201 := client.Open("/b8YxyJRioX.txt", client.O_RDWR|client.O_CREATE)
if(fd201 < 0) {
  panic("open failed")
}


ret = client.Seek(fd201, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd202 := client.Open("/8W6113_G3d.txt", client.O_RDWR|client.O_CREATE)
if(fd202 < 0) {
  panic("open failed")
}


ret = client.Seek(fd202, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd202, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd200, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd201, []byte("zKTtVuTTNf84wOouvc1OeH5XENwa9QVidq4DreCwLn36qOIq7xCpWfBt5R0JYQ7TTyGFBFPvfcNc_zqBcGL5Zuk3sbG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) zKTtVuTTNf84wOouvc1OeH5XENwa9QVidq4DreCwLn36qOIq7xCpWfBt5R0JYQ7TTyGFBFPvfcNc_zqBcGL5Zuk3sbG

fd203 := client.Open("/CzVjEdkRbn.txt", client.O_RDWR|client.O_CREATE)
if(fd203 < 0) {
  panic("open failed")
}


ret = client.Close(fd202)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd201, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


fd204 := client.Open("/LZcAyv_xCm.txt", client.O_RDWR|client.O_CREATE)
if(fd204 < 0) {
  panic("open failed")
}


fd205 := client.Open("/W8T68Q5TiM.txt", client.O_RDWR|client.O_CREATE)
if(fd205 < 0) {
  panic("open failed")
}


ret = client.Close(fd204)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd205)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd203)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd200, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd201)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd200, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd200)
if(ret != 0) {
  panic("close failed")
}


fd206 := client.Open("/wZnvthnoX2.txt", client.O_RDWR|client.O_CREATE)
if(fd206 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd206, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ljjG6WCSkU5G3Ddow7qXW726vUAHL5nHCieIrLwus2hDk23N5Q8yfBoRpb30tXrRo2KOjk5SXuonXj_m") {
  panic("wrong data returned")
}


ret = client.Close(fd206)
if(ret != 0) {
  panic("close failed")
}


fd207 := client.Open("/0KnXHjRLMh.txt", client.O_RDWR|client.O_CREATE)
if(fd207 < 0) {
  panic("open failed")
}


ret = client.Close(fd207)
if(ret != 0) {
  panic("close failed")
}


fd208 := client.Open("/OHK7BLBoYt.txt", client.O_RDWR|client.O_CREATE)
if(fd208 < 0) {
  panic("open failed")
}


fd209 := client.Open("/neIEIf4G34.txt", client.O_RDWR|client.O_CREATE)
if(fd209 < 0) {
  panic("open failed")
}


ret = client.Seek(fd209, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd208, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd210 := client.Open("/FUX87rui7W.txt", client.O_RDWR|client.O_CREATE)
if(fd210 < 0) {
  panic("open failed")
}


fd211 := client.Open("/WuYpJrAIYF.txt", client.O_RDWR|client.O_CREATE)
if(fd211 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd208, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd211, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd212 := client.Open("/J1S2AYAts5.txt", client.O_RDWR|client.O_CREATE)
if(fd212 < 0) {
  panic("open failed")
}


ret = client.Close(fd211)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd209, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd208, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd213 := client.Open("/7evLhsHpf6.txt", client.O_RDWR|client.O_CREATE)
if(fd213 < 0) {
  panic("open failed")
}


ret = client.Close(fd213)
if(ret != 0) {
  panic("close failed")
}


fd214 := client.Open("/JCI3j4WuOw.txt", client.O_RDWR|client.O_CREATE)
if(fd214 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd209, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd215 := client.Open("/W2Rqw1Co6N.txt", client.O_RDWR|client.O_CREATE)
if(fd215 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd208, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd212, []byte("3yC6s99WbrFrXNlxCDMQ6Ecb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) 3yC6s99WbrFrXNlxCDMQ6Ecb

fd216 := client.Open("/4mPDBgaYe1.txt", client.O_RDWR|client.O_CREATE)
if(fd216 < 0) {
  panic("open failed")
}

//fd state: (0) d532dkIAfcRzT0Sx5WZCMO6IwpxAC3vXmG57uL7VvUlfz7F4SNshpwaZ4TbunhWas8Y3pRM

ret = client.Write(fd216, []byte("MA5Wnx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) MA5WnxIAfcRzT0Sx5WZCMO6IwpxAC3vXmG57uL7VvUlfz7F4SNshpwaZ4TbunhWas8Y3pRM

ret = client.Seek(fd215, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd208)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd209, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd212)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd215, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd216, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "IA") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd214, []byte("5WKeCJVQ9Kr1__w"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) 5WKeCJVQ9Kr1__w

ret = client.Close(fd209)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd210)
if(ret != 0) {
  panic("close failed")
}

//fd state: (8) MA5WnxIAfcRzT0Sx5WZCMO6IwpxAC3vXmG57uL7VvUlfz7F4SNshpwaZ4TbunhWas8Y3pRM

ret = client.Write(fd216, []byte("8lEbbDmlPHuNMP9YptHW17Ojw7SCK5YF2ivpfPk2U2NzCuS9FUjM84cHrEimH__"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (71) MA5WnxIA8lEbbDmlPHuNMP9YptHW17Ojw7SCK5YF2ivpfPk2U2NzCuS9FUjM84cHrEimH__

buf, ret = client.Read(fd214, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd214, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd217 := client.Open("/9yGoBNc9vR.txt", client.O_RDWR|client.O_CREATE)
if(fd217 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd216, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd215, []byte("BNYmJqqVS3nYWYtmGggMDm3gPwjnJr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) BNYmJqqVS3nYWYtmGggMDm3gPwjnJr

buf, ret = client.Read(fd217, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (30) BNYmJqqVS3nYWYtmGggMDm3gPwjnJr

ret = client.Write(fd215, []byte("1EUy7H_g7XMX1AYNu7zjda5sfLYHfN745fKzaR7Hq7vYtoDVIZX9PfHHFmEOZaNeNkPHkySkOs5T3oGD5u0YIIaaBgAULW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) BNYmJqqVS3nYWYtmGggMDm3gPwjnJr1EUy7H_g7XMX1AYNu7zjda5sfLYHfN745fKzaR7Hq7vYtoDVIZX9PfHHFmEOZaNeNkPHkySkOs5T3oGD5u0YIIaaBgAULW

buf, ret = client.Read(fd216, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd215, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd218 := client.Open("/c5pmebz5PR.txt", client.O_RDWR|client.O_CREATE)
if(fd218 < 0) {
  panic("open failed")
}


ret = client.Close(fd217)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd214, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}

//fd state: (71) MA5WnxIA8lEbbDmlPHuNMP9YptHW17Ojw7SCK5YF2ivpfPk2U2NzCuS9FUjM84cHrEimH__

ret = client.Write(fd216, []byte("_navGHXrOmJOzanpZCjZgIrxMc_zs9EaxOlZe7Eh8bfG4eyePiU9KXWaimMiN5eWnTNWU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (140) MA5WnxIA8lEbbDmlPHuNMP9YptHW17Ojw7SCK5YF2ivpfPk2U2NzCuS9FUjM84cHrEimH___navGHXrOmJOzanpZCjZgIrxMc_zs9EaxOlZe7Eh8bfG4eyePiU9KXWaimMiN5eWnTNWU

ret = client.Close(fd215)
if(ret != 0) {
  panic("close failed")
}

//fd state: (140) MA5WnxIA8lEbbDmlPHuNMP9YptHW17Ojw7SCK5YF2ivpfPk2U2NzCuS9FUjM84cHrEimH___navGHXrOmJOzanpZCjZgIrxMc_zs9EaxOlZe7Eh8bfG4eyePiU9KXWaimMiN5eWnTNWU

ret = client.Write(fd216, []byte("q9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbIi1YXJ_kG57uh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (194) MA5WnxIA8lEbbDmlPHuNMP9YptHW17Ojw7SCK5YF2ivpfPk2U2NzCuS9FUjM84cHrEimH___navGHXrOmJOzanpZCjZgIrxMc_zs9EaxOlZe7Eh8bfG4eyePiU9KXWaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbIi1YXJ_kG57uh

buf, ret = client.Read(fd214, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JVQ9Kr1__w") {
  panic("wrong data returned")
}


fd219 := client.Open("/8W6113_G3d.txt", client.O_RDWR|client.O_CREATE)
if(fd219 < 0) {
  panic("open failed")
}


fd220 := client.Open("/MVWR6bEOKx.txt", client.O_RDWR|client.O_CREATE)
if(fd220 < 0) {
  panic("open failed")
}


ret = client.Seek(fd214, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


buf, ret = client.Read(fd216, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd218)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd219, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd216, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


ret = client.Seek(fd220, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (7) 5WKeCJVQ9Kr1__w

ret = client.Write(fd214, []byte("FKHsvV3uPVT2JjpWi7dnFZEThDu1PoV62PnyeAWJh_QJe64TEjpoqAnKK6mvQNxJQaKf_E0zVxEO25bsXWi1crST"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) 5WKeCJVFKHsvV3uPVT2JjpWi7dnFZEThDu1PoV62PnyeAWJh_QJe64TEjpoqAnKK6mvQNxJQaKf_E0zVxEO25bsXWi1crST

ret = client.Seek(fd220, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (95) 5WKeCJVFKHsvV3uPVT2JjpWi7dnFZEThDu1PoV62PnyeAWJh_QJe64TEjpoqAnKK6mvQNxJQaKf_E0zVxEO25bsXWi1crST

ret = client.Write(fd214, []byte("hvfm1XqJL64nUzOiasPyppx7E_pHEUoAESRep4mr1OLRik_7s98XB4Kc4RsCJKHgTYD8K3whlksjGi1NmTX7ib3V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (183) 5WKeCJVFKHsvV3uPVT2JjpWi7dnFZEThDu1PoV62PnyeAWJh_QJe64TEjpoqAnKK6mvQNxJQaKf_E0zVxEO25bsXWi1crSThvfm1XqJL64nUzOiasPyppx7E_pHEUoAESRep4mr1OLRik_7s98XB4Kc4RsCJKHgTYD8K3whlksjGi1NmTX7ib3V

buf, ret = client.Read(fd214, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd221 := client.Open("/SbQouCL1vm.txt", client.O_RDWR|client.O_CREATE)
if(fd221 < 0) {
  panic("open failed")
}


ret = client.Seek(fd216, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


ret = client.Seek(fd220, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd221, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd219, []byte("97V_WpedoDgdE9rQlulGDV5bWWwVUkN3EYlrlZBXqW1g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) 97V_WpedoDgdE9rQlulGDV5bWWwVUkN3EYlrlZBXqW1g
//fd state: (18) MA5WnxIA8lEbbDmlPHuNMP9YptHW17Ojw7SCK5YF2ivpfPk2U2NzCuS9FUjM84cHrEimH___navGHXrOmJOzanpZCjZgIrxMc_zs9EaxOlZe7Eh8bfG4eyePiU9KXWaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbIi1YXJ_kG57uh

ret = client.Write(fd216, []byte("LiB9HbElo7j7pP8y2Uuju4Jm914QburF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) MA5WnxIA8lEbbDmlPHLiB9HbElo7j7pP8y2Uuju4Jm914QburF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1C8bfG4eyePiU9KXWaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbIi1YXJ_kG57uh

ret = client.Close(fd214)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd219, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}

//fd state: (31) 97V_WpedoDgdE9rQlulGDV5bWWwVUkN3EYlrlZBXqW1g

ret = client.Write(fd219, []byte("mZvoYmYBTJAxwaCSYXO6niMJxQXutVqi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) 97V_WpedoDgdE9rQlulGDV5bWWwVUkNmZvoYmYBTJAxwaCSYXO6niMJxQXutVqi
//fd state: (0) 

ret = client.Write(fd221, []byte("36zkNe7kIzHoUBECFdun8cUaoZo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) 36zkNe7kIzHoUBECFdun8cUaoZo

buf, ret = client.Read(fd221, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd220)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd221)
if(ret != 0) {
  panic("close failed")
}

//fd state: (111) MA5WnxIA8lEbbDmlPHLiB9HbElo7j7pP8y2Uuju4Jm914QburF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1C8bfG4eyePiU9KXWaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbIi1YXJ_kG57uh

ret = client.Write(fd216, []byte("MUFIN7wkaxagbBn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (126) MA5WnxIA8lEbbDmlPHLiB9HbElo7j7pP8y2Uuju4Jm914QburF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1CMUFIN7wkaxagbBnaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbIi1YXJ_kG57uh

ret = client.Close(fd219)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd216, 156, client.SEEK_SET)
if(ret != 156) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 156)
  panic("seek failed")
}


fd222 := client.Open("/7evLhsHpf6.txt", client.O_RDWR|client.O_CREATE)
if(fd222 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd216, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1GWaolxXhP0m2YVfYxCfXkrhbI") {
  panic("wrong data returned")
}


ret = client.Close(fd222)
if(ret != 0) {
  panic("close failed")
}

//fd state: (182) MA5WnxIA8lEbbDmlPHLiB9HbElo7j7pP8y2Uuju4Jm914QburF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1CMUFIN7wkaxagbBnaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbIi1YXJ_kG57uh

ret = client.Write(fd216, []byte("4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (269) MA5WnxIA8lEbbDmlPHLiB9HbElo7j7pP8y2Uuju4Jm914QburF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1CMUFIN7wkaxagbBnaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyD
//fd state: (269) MA5WnxIA8lEbbDmlPHLiB9HbElo7j7pP8y2Uuju4Jm914QburF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1CMUFIN7wkaxagbBnaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyD

ret = client.Write(fd216, []byte("sFP4lawaUhuZmMpiTSPNu1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (291) MA5WnxIA8lEbbDmlPHLiB9HbElo7j7pP8y2Uuju4Jm914QburF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1CMUFIN7wkaxagbBnaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1

buf, ret = client.Read(fd216, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd216, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


buf, ret = client.Read(fd216, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1CMUFIN7wkaxagbBnaimMiN5eWnTNWUq9T4q") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd216, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yHW") {
  panic("wrong data returned")
}


ret = client.Seek(fd216, 133, client.SEEK_SET)
if(ret != 133) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 133)
  panic("seek failed")
}


ret = client.Close(fd216)
if(ret != 0) {
  panic("close failed")
}


fd223 := client.Open("/4cQqDopMxJ.txt", client.O_RDWR|client.O_CREATE)
if(fd223 < 0) {
  panic("open failed")
}


ret = client.Close(fd223)
if(ret != 0) {
  panic("close failed")
}


fd224 := client.Open("/8W6113_G3d.txt", client.O_RDWR|client.O_CREATE)
if(fd224 < 0) {
  panic("open failed")
}


fd225 := client.Open("/uEvYTullIO.txt", client.O_RDWR|client.O_CREATE)
if(fd225 < 0) {
  panic("open failed")
}


ret = client.Close(fd225)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd224, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}


ret = client.Close(fd224)
if(ret != 0) {
  panic("close failed")
}


fd226 := client.Open("/K1q8zuNqur.txt", client.O_RDWR|client.O_CREATE)
if(fd226 < 0) {
  panic("open failed")
}


fd227 := client.Open("/_fKtkmMG2H.txt", client.O_RDWR|client.O_CREATE)
if(fd227 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd226, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bADDD6hFBrkBtYr6ILPQr") {
  panic("wrong data returned")
}


ret = client.Seek(fd227, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd226, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


ret = client.Seek(fd227, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd226, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "n7FZqWJy5boLze1lt3Yu8zNmgojxwBDiKQixqhuv4KSawP0NOgWaUUFrdK") {
  panic("wrong data returned")
}


ret = client.Seek(fd227, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd227, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd228 := client.Open("/XdPxjeqcPT.txt", client.O_RDWR|client.O_CREATE)
if(fd228 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd226, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd226, 61, client.SEEK_SET)
if(ret != 61) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 61)
  panic("seek failed")
}


ret = client.Seek(fd226, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Seek(fd227, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd228, []byte("KUk78xEbHsbpXuOnCESMqsiGSycSL98NkevTNBdPJ01C6xNx85TNp5ZbxEdJXYpQxm32kbqDUk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) KUk78xEbHsbpXuOnCESMqsiGSycSL98NkevTNBdPJ01C6xNx85TNp5ZbxEdJXYpQxm32kbqDUk

ret = client.Close(fd228)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd227)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd226, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}


ret = client.Close(fd226)
if(ret != 0) {
  panic("close failed")
}


fd229 := client.Open("/HRpPyOlU84.txt", client.O_RDWR|client.O_CREATE)
if(fd229 < 0) {
  panic("open failed")
}


fd230 := client.Open("/UfA4o9VB69.txt", client.O_RDWR|client.O_CREATE)
if(fd230 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd229, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd230, []byte("KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pk

fd231 := client.Open("/KVp2ABEVFd.txt", client.O_RDWR|client.O_CREATE)
if(fd231 < 0) {
  panic("open failed")
}


ret = client.Seek(fd231, 93, client.SEEK_SET)
if(ret != 93) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 93)
  panic("seek failed")
}


buf, ret = client.Read(fd229, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (77) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pk

ret = client.Write(fd230, []byte("LksfnX3kt6rGHkkjU5unyGc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGc
//fd state: (0) 

ret = client.Write(fd229, []byte("p3yczrHV40FD6zNNWkBbRlrtd3ROefhGYZ5qIt6FZw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) p3yczrHV40FD6zNNWkBbRlrtd3ROefhGYZ5qIt6FZw
//fd state: (100) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGc

ret = client.Write(fd230, []byte("oVrTHdq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdq

buf, ret = client.Read(fd231, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OukC8WQM5I_W5cyXDyKPpk4") {
  panic("wrong data returned")
}


fd232 := client.Open("/jEJT1wijAu.txt", client.O_RDWR|client.O_CREATE)
if(fd232 < 0) {
  panic("open failed")
}


ret = client.Seek(fd232, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


fd233 := client.Open("/lDebYnkzxS.txt", client.O_RDWR|client.O_CREATE)
if(fd233 < 0) {
  panic("open failed")
}


ret = client.Close(fd232)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd229, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}

//fd state: (107) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdq

ret = client.Write(fd230, []byte("HI9lanbdUW4BABvxx8mqPcJwrR61XWVBg2eUTlINt1b7YAtiDtMwnw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (161) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqPcJwrR61XWVBg2eUTlINt1b7YAtiDtMwnw

buf, ret = client.Read(fd233, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd234 := client.Open("/iBtizDgltf.txt", client.O_RDWR|client.O_CREATE)
if(fd234 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd233, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd231)
if(ret != 0) {
  panic("close failed")
}


fd235 := client.Open("/rcJYn3q9Iv.txt", client.O_RDWR|client.O_CREATE)
if(fd235 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd233, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd233, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd229)
if(ret != 0) {
  panic("close failed")
}


fd236 := client.Open("/dkuo7HULBj.txt", client.O_RDWR|client.O_CREATE)
if(fd236 < 0) {
  panic("open failed")
}


ret = client.Seek(fd230, 83, client.SEEK_SET)
if(ret != 83) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 83)
  panic("seek failed")
}

//fd state: (0) BcviVK8npvmqdA8hdNmVaAEBi88PF4gK2_RMJAmh7v53O_FG9W8kQFUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qwd4bzTLnBwq5zVIUChZE4qgTVOCbh0rjJTPmjO8LhvnVtfQGreDAI1XPcrQu_7IWcxj339rRH0Why3tWFS9iCvROX7iRBJhw1CimzlvOxIGoTV20Yy8n

ret = client.Write(fd235, []byte("_QxLDIWcnNtoyT8RscuOsZFIrXR3s57VgSaKnhrfo3PL4a54qzaI8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) _QxLDIWcnNtoyT8RscuOsZFIrXR3s57VgSaKnhrfo3PL4a54qzaI8FUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qwd4bzTLnBwq5zVIUChZE4qgTVOCbh0rjJTPmjO8LhvnVtfQGreDAI1XPcrQu_7IWcxj339rRH0Why3tWFS9iCvROX7iRBJhw1CimzlvOxIGoTV20Yy8n

ret = client.Seek(fd230, 135, client.SEEK_SET)
if(ret != 135) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 135)
  panic("seek failed")
}


buf, ret = client.Read(fd234, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd237 := client.Open("/OvTWV_UsAe.txt", client.O_RDWR|client.O_CREATE)
if(fd237 < 0) {
  panic("open failed")
}


fd238 := client.Open("/Kbe_5esbc2.txt", client.O_RDWR|client.O_CREATE)
if(fd238 < 0) {
  panic("open failed")
}


ret = client.Close(fd230)
if(ret != 0) {
  panic("close failed")
}


fd239 := client.Open("/Gg_XK1OoBp.txt", client.O_RDWR|client.O_CREATE)
if(fd239 < 0) {
  panic("open failed")
}


ret = client.Close(fd235)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd234, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd239, []byte("X1PsqzOwDTdo9Kf9NTIaev5j_X1hYhYQ1XVCifxPHfdjcMohMoF6DroadO2pD9Ict1vFb4NeTjo8aVI2Jd2BWYF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) X1PsqzOwDTdo9Kf9NTIaev5j_X1hYhYQ1XVCifxPHfdjcMohMoF6DroadO2pD9Ict1vFb4NeTjo8aVI2Jd2BWYF

ret = client.Seek(fd239, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


ret = client.Close(fd233)
if(ret != 0) {
  panic("close failed")
}


fd240 := client.Open("/XdPxjeqcPT.txt", client.O_RDWR|client.O_CREATE)
if(fd240 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd237, []byte("rC2ZsWf_0aFXO_jxOcIprtil"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) rC2ZsWf_0aFXO_jxOcIprtil

ret = client.Seek(fd237, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}

//fd state: (17) rC2ZsWf_0aFXO_jxOcIprtil

ret = client.Write(fd237, []byte("TyfRH9XU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) rC2ZsWf_0aFXO_jxOTyfRH9XU

fd241 := client.Open("/UDXib7cZUL.txt", client.O_RDWR|client.O_CREATE)
if(fd241 < 0) {
  panic("open failed")
}


ret = client.Close(fd234)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd238, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd238, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd237)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd240, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


ret = client.Seek(fd236, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd239, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PHfdjcMohMoF6DroadO2pD9Ict1vFb4NeTjo8aVI2Jd2BWYF") {
  panic("wrong data returned")
}


ret = client.Seek(fd238, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd238, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (19) KUk78xEbHsbpXuOnCESMqsiGSycSL98NkevTNBdPJ01C6xNx85TNp5ZbxEdJXYpQxm32kbqDUk

ret = client.Write(fd240, []byte("KHFxHCmRvB2DkEN06xeAhcgyLZn59htkPOou8E4HHSulBcrVkxx42FKlnhRgScRDkU7LqQFWCqizmit2g8AAB7e8ekEEJSnM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) KUk78xEbHsbpXuOnCESKHFxHCmRvB2DkEN06xeAhcgyLZn59htkPOou8E4HHSulBcrVkxx42FKlnhRgScRDkU7LqQFWCqizmit2g8AAB7e8ekEEJSnM

ret = client.Seek(fd239, 83, client.SEEK_SET)
if(ret != 83) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 83)
  panic("seek failed")
}


fd242 := client.Open("/f_VMh3LK24.txt", client.O_RDWR|client.O_CREATE)
if(fd242 < 0) {
  panic("open failed")
}


fd243 := client.Open("/OHK7BLBoYt.txt", client.O_RDWR|client.O_CREATE)
if(fd243 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd236, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd242, []byte("46CVo00yrl3np4GS1WZRdSKo60SsmTtro"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) 46CVo00yrl3np4GS1WZRdSKo60SsmTtro

buf, ret = client.Read(fd242, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (33) 46CVo00yrl3np4GS1WZRdSKo60SsmTtro

ret = client.Write(fd242, []byte("C3sp1NNbU4yOGkT34hny8LEBXu8WpXTwzCfxhyd6SCkSSfkD8qNAWPPMkhGtDH4_uVAfYRAcT5IdoKg2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) 46CVo00yrl3np4GS1WZRdSKo60SsmTtroC3sp1NNbU4yOGkT34hny8LEBXu8WpXTwzCfxhyd6SCkSSfkD8qNAWPPMkhGtDH4_uVAfYRAcT5IdoKg2
//fd state: (83) X1PsqzOwDTdo9Kf9NTIaev5j_X1hYhYQ1XVCifxPHfdjcMohMoF6DroadO2pD9Ict1vFb4NeTjo8aVI2Jd2BWYF

ret = client.Write(fd239, []byte("vg4EsuUBnZv0Dk8_miRDeX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) X1PsqzOwDTdo9Kf9NTIaev5j_X1hYhYQ1XVCifxPHfdjcMohMoF6DroadO2pD9Ict1vFb4NeTjo8aVI2Jd2vg4EsuUBnZv0Dk8_miRDeX

ret = client.Close(fd242)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd240)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd236, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd236)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd241, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd243, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd238, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd244 := client.Open("/49yoWYxM6B.txt", client.O_RDWR|client.O_CREATE)
if(fd244 < 0) {
  panic("open failed")
}


ret = client.Close(fd239)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd241, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd244, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jBS64LgzdahuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JAR") {
  panic("wrong data returned")
}


ret = client.Seek(fd241, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd241, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd241)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd243)
if(ret != 0) {
  panic("close failed")
}

//fd state: (72) jBS64LgzdahuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JARlFDyPX3qYiHqvZLLIiGijPSvekTtqJZ1jGm_vJ2JagjFI_5CYn4AGSAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQpgygN3aOEmoYMSI9sDypzuOtD86FeN7ykkKcOm83iIXDDKOMQLAtZnnM85iOcF_3

ret = client.Write(fd244, []byte("4Jii_mBSietBx4GqcOiw8NF6K6f9iFXz5AiI9kNL1IOSBVi9rYL8hp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (126) jBS64LgzdahuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JAR4Jii_mBSietBx4GqcOiw8NF6K6f9iFXz5AiI9kNL1IOSBVi9rYL8hpAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQpgygN3aOEmoYMSI9sDypzuOtD86FeN7ykkKcOm83iIXDDKOMQLAtZnnM85iOcF_3

buf, ret = client.Read(fd238, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd244, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQp") {
  panic("wrong data returned")
}


ret = client.Close(fd244)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd238, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd238, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd238, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd238)
if(ret != 0) {
  panic("close failed")
}


fd245 := client.Open("/njkggBWqLT.txt", client.O_RDWR|client.O_CREATE)
if(fd245 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd245, []byte("a8H_gXSsOfnHoKgRkn2u_mlO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) a8H_gXSsOfnHoKgRkn2u_mlO

ret = client.Close(fd245)
if(ret != 0) {
  panic("close failed")
}


fd246 := client.Open("/lK1Ra888uV.txt", client.O_RDWR|client.O_CREATE)
if(fd246 < 0) {
  panic("open failed")
}


ret = client.Seek(fd246, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd246, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd246)
if(ret != 0) {
  panic("close failed")
}


fd247 := client.Open("/4FdZoYlgde.txt", client.O_RDWR|client.O_CREATE)
if(fd247 < 0) {
  panic("open failed")
}


fd248 := client.Open("/Me76BT6BSP.txt", client.O_RDWR|client.O_CREATE)
if(fd248 < 0) {
  panic("open failed")
}


ret = client.Close(fd247)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd248, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd248, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd248, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd248)
if(ret != 0) {
  panic("close failed")
}


fd249 := client.Open("/tikf5vEc8d.txt", client.O_RDWR|client.O_CREATE)
if(fd249 < 0) {
  panic("open failed")
}

//fd state: (0) Jom6S0z_sQRmVyAyzb_qnXeZIkhUYznCLVrgmNKjMxcLFCSAfmaGcMpyxSybN3eLPkJ_uvqBYVBxdvXOucLi36dDI_ZqMm

ret = client.Write(fd249, []byte("yGSeWQm7MDfnrW9Nz1Q572eS3F1d5an9lNrBt5azIZjI34U8IUAn4ghVWBrZe4tdfG5zovDSI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) yGSeWQm7MDfnrW9Nz1Q572eS3F1d5an9lNrBt5azIZjI34U8IUAn4ghVWBrZe4tdfG5zovDSIVBxdvXOucLi36dDI_ZqMm

buf, ret = client.Read(fd249, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "VBxdvXOucLi36dDI_ZqMm") {
  panic("wrong data returned")
}


ret = client.Close(fd249)
if(ret != 0) {
  panic("close failed")
}


fd250 := client.Open("/pLIWsilfjy.txt", client.O_RDWR|client.O_CREATE)
if(fd250 < 0) {
  panic("open failed")
}

//fd state: (0) qK4naYPq7iw_hMvnCD4pxBoiRwJmXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UBJpYL903agqPgfabhjNfq6rwRhscbXruOCeMpO7YqcF_tF5q8iSRheJPIGYZFwMCchFvcOUPuamzHVP3oCLHmKBr71EEl2Wb8jUXXif2arhF

ret = client.Write(fd250, []byte("SCxPeJ7fvt2xZNehgjyHHnVMJ1s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) SCxPeJ7fvt2xZNehgjyHHnVMJ1smXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UBJpYL903agqPgfabhjNfq6rwRhscbXruOCeMpO7YqcF_tF5q8iSRheJPIGYZFwMCchFvcOUPuamzHVP3oCLHmKBr71EEl2Wb8jUXXif2arhF

ret = client.Seek(fd250, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


fd251 := client.Open("/hiNlZwaae9.txt", client.O_RDWR|client.O_CREATE)
if(fd251 < 0) {
  panic("open failed")
}


ret = client.Seek(fd251, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd252 := client.Open("/fE1zI_tMmi.txt", client.O_RDWR|client.O_CREATE)
if(fd252 < 0) {
  panic("open failed")
}


fd253 := client.Open("/KqykqQdiXF.txt", client.O_RDWR|client.O_CREATE)
if(fd253 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd250, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "mXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHy") {
  panic("wrong data returned")
}


fd254 := client.Open("/6rEUgk3ohz.txt", client.O_RDWR|client.O_CREATE)
if(fd254 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd252, []byte("7Gd2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) 7Gd2

ret = client.Close(fd253)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd250)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd254, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd251, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd255 := client.Open("/IlPekcd4f8.txt", client.O_RDWR|client.O_CREATE)
if(fd255 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd255, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd256 := client.Open("/ZaBfLNx_Cb.txt", client.O_RDWR|client.O_CREATE)
if(fd256 < 0) {
  panic("open failed")
}


ret = client.Seek(fd255, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd257 := client.Open("/oIsR5Rk83H.txt", client.O_RDWR|client.O_CREATE)
if(fd257 < 0) {
  panic("open failed")
}

//fd state: (4) 7Gd2

ret = client.Write(fd252, []byte("aR2OQ4wvoUB9JxfSL58SxeKPG9FqWCMhNLQHxLgUkZRu_iolJ6mThcz8tCCWBN4hoeEvvYu57J2EidDRRdkkb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) 7Gd2aR2OQ4wvoUB9JxfSL58SxeKPG9FqWCMhNLQHxLgUkZRu_iolJ6mThcz8tCCWBN4hoeEvvYu57J2EidDRRdkkb

buf, ret = client.Read(fd252, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd254, []byte("YqBedplBG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) YqBedplBG

fd258 := client.Open("/giqe_vyTh1.txt", client.O_RDWR|client.O_CREATE)
if(fd258 < 0) {
  panic("open failed")
}


fd259 := client.Open("/o21QrKTXjD.txt", client.O_RDWR|client.O_CREATE)
if(fd259 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd254, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd260 := client.Open("/c5R3eXHXLk.txt", client.O_RDWR|client.O_CREATE)
if(fd260 < 0) {
  panic("open failed")
}


ret = client.Close(fd259)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 6WvRUgVZCOQmAtKHTSXr5HXSDpJdrFKf3NFAMgAIgwA0MWyAP0IHjmzKoHJJMCwmUz41Wnupj5QZKAY5PSWJ6vlmXeLIXL2OfC9k9J1SUkc2RmOrKglL3nOQTqtrwuaSq6h_FgSa3LUpdllwUCzVztiPEjalnUgc

ret = client.Write(fd257, []byte("ZNjB7VzlyX7IWrlPLBaUucFAuTrFTvLCefU9oXs4joz3t0SYjUEl5GCXcvTnXPBKigGZwdXD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) ZNjB7VzlyX7IWrlPLBaUucFAuTrFTvLCefU9oXs4joz3t0SYjUEl5GCXcvTnXPBKigGZwdXDj5QZKAY5PSWJ6vlmXeLIXL2OfC9k9J1SUkc2RmOrKglL3nOQTqtrwuaSq6h_FgSa3LUpdllwUCzVztiPEjalnUgc
//fd state: (72) ZNjB7VzlyX7IWrlPLBaUucFAuTrFTvLCefU9oXs4joz3t0SYjUEl5GCXcvTnXPBKigGZwdXDj5QZKAY5PSWJ6vlmXeLIXL2OfC9k9J1SUkc2RmOrKglL3nOQTqtrwuaSq6h_FgSa3LUpdllwUCzVztiPEjalnUgc

ret = client.Write(fd257, []byte("qFfZbYPUy23eI8D0Etct"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) ZNjB7VzlyX7IWrlPLBaUucFAuTrFTvLCefU9oXs4joz3t0SYjUEl5GCXcvTnXPBKigGZwdXDqFfZbYPUy23eI8D0EtctXL2OfC9k9J1SUkc2RmOrKglL3nOQTqtrwuaSq6h_FgSa3LUpdllwUCzVztiPEjalnUgc
//fd state: (9) YqBedplBG

ret = client.Write(fd254, []byte("NDxza1D7XOGbty9mUqt1Mo_qjs6_diZFxntTZAj5L9gI2WAcq8kR1poc9HArBMZZlSPLX6_FaSRnYv7i27u6zn_jLhI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) YqBedplBGNDxza1D7XOGbty9mUqt1Mo_qjs6_diZFxntTZAj5L9gI2WAcq8kR1poc9HArBMZZlSPLX6_FaSRnYv7i27u6zn_jLhI
//fd state: (92) ZNjB7VzlyX7IWrlPLBaUucFAuTrFTvLCefU9oXs4joz3t0SYjUEl5GCXcvTnXPBKigGZwdXDqFfZbYPUy23eI8D0EtctXL2OfC9k9J1SUkc2RmOrKglL3nOQTqtrwuaSq6h_FgSa3LUpdllwUCzVztiPEjalnUgc

ret = client.Write(fd257, []byte("onR6qShWfjll2uhbCZdHAbLVlglDKF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) ZNjB7VzlyX7IWrlPLBaUucFAuTrFTvLCefU9oXs4joz3t0SYjUEl5GCXcvTnXPBKigGZwdXDqFfZbYPUy23eI8D0EtctonR6qShWfjll2uhbCZdHAbLVlglDKFtrwuaSq6h_FgSa3LUpdllwUCzVztiPEjalnUgc

ret = client.Close(fd260)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd257, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


fd261 := client.Open("/gwR_qKf5l_.txt", client.O_RDWR|client.O_CREATE)
if(fd261 < 0) {
  panic("open failed")
}

//fd state: (40) ZNjB7VzlyX7IWrlPLBaUucFAuTrFTvLCefU9oXs4joz3t0SYjUEl5GCXcvTnXPBKigGZwdXDqFfZbYPUy23eI8D0EtctonR6qShWfjll2uhbCZdHAbLVlglDKFtrwuaSq6h_FgSa3LUpdllwUCzVztiPEjalnUgc

ret = client.Write(fd257, []byte("rWWgWkrxkwWZD9imyoTdcxujkFeYUo8IBN2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) ZNjB7VzlyX7IWrlPLBaUucFAuTrFTvLCefU9oXs4rWWgWkrxkwWZD9imyoTdcxujkFeYUo8IBN2ZbYPUy23eI8D0EtctonR6qShWfjll2uhbCZdHAbLVlglDKFtrwuaSq6h_FgSa3LUpdllwUCzVztiPEjalnUgc

ret = client.Close(fd261)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) _GlXZu3tWoM5LxH6IzUculJ2lbJAwDJ2Wl22btrhBpScrgR9P8F6ZeY_s78yY2InTBU8_mG7fySwLsVYAHfgUQ0RBx299sghdSy

ret = client.Write(fd256, []byte("aB5m3mQrTStEqW0XsuCjIsAZdkKSFRgP9hcwGU2y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) aB5m3mQrTStEqW0XsuCjIsAZdkKSFRgP9hcwGU2yBpScrgR9P8F6ZeY_s78yY2InTBU8_mG7fySwLsVYAHfgUQ0RBx299sghdSy

ret = client.Close(fd255)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd251, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd254, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (89) 7Gd2aR2OQ4wvoUB9JxfSL58SxeKPG9FqWCMhNLQHxLgUkZRu_iolJ6mThcz8tCCWBN4hoeEvvYu57J2EidDRRdkkb

ret = client.Write(fd252, []byte("tc3Dpm4ypvKsavRBMW5gAXp46IZyfBQF3DMiVXEzPotxd0Ti"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (137) 7Gd2aR2OQ4wvoUB9JxfSL58SxeKPG9FqWCMhNLQHxLgUkZRu_iolJ6mThcz8tCCWBN4hoeEvvYu57J2EidDRRdkkbtc3Dpm4ypvKsavRBMW5gAXp46IZyfBQF3DMiVXEzPotxd0Ti
//fd state: (137) 7Gd2aR2OQ4wvoUB9JxfSL58SxeKPG9FqWCMhNLQHxLgUkZRu_iolJ6mThcz8tCCWBN4hoeEvvYu57J2EidDRRdkkbtc3Dpm4ypvKsavRBMW5gAXp46IZyfBQF3DMiVXEzPotxd0Ti

ret = client.Write(fd252, []byte("Fg3Cqd2J25a7c3rCth5fNgOMXfsFbW__7f262dg33zw3YHshRyOjXeHr0Vs3romsz37XpQPEQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (210) 7Gd2aR2OQ4wvoUB9JxfSL58SxeKPG9FqWCMhNLQHxLgUkZRu_iolJ6mThcz8tCCWBN4hoeEvvYu57J2EidDRRdkkbtc3Dpm4ypvKsavRBMW5gAXp46IZyfBQF3DMiVXEzPotxd0TiFg3Cqd2J25a7c3rCth5fNgOMXfsFbW__7f262dg33zw3YHshRyOjXeHr0Vs3romsz37XpQPEQ

fd262 := client.Open("/icy2T1pZHr.txt", client.O_RDWR|client.O_CREATE)
if(fd262 < 0) {
  panic("open failed")
}


fd263 := client.Open("/R6UEMyClwH.txt", client.O_RDWR|client.O_CREATE)
if(fd263 < 0) {
  panic("open failed")
}


ret = client.Close(fd252)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd256, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "BpScrgR9P8F6ZeY_s78yY2InTBU8_mG7fySwLsVYAHfgUQ0RBx299sghdSy") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd263, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd264 := client.Open("/D3Fa1l3zaJ.txt", client.O_RDWR|client.O_CREATE)
if(fd264 < 0) {
  panic("open failed")
}


fd265 := client.Open("/btiirxjYay.txt", client.O_RDWR|client.O_CREATE)
if(fd265 < 0) {
  panic("open failed")
}


ret = client.Close(fd257)
if(ret != 0) {
  panic("close failed")
}


fd266 := client.Open("/xUg22Cub2y.txt", client.O_RDWR|client.O_CREATE)
if(fd266 < 0) {
  panic("open failed")
}


ret = client.Close(fd264)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd265)
if(ret != 0) {
  panic("close failed")
}


fd267 := client.Open("/gqPLYlOPff.txt", client.O_RDWR|client.O_CREATE)
if(fd267 < 0) {
  panic("open failed")
}


ret = client.Seek(fd262, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd266, []byte("KU8TlJ6F3Dv5H7WgrptSxLKlgzXo7ylvTejOrgtPkvD7bWotYSsLBovh3O3cCsebyb1lqb6spY4pud0106J"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) KU8TlJ6F3Dv5H7WgrptSxLKlgzXo7ylvTejOrgtPkvD7bWotYSsLBovh3O3cCsebyb1lqb6spY4pud0106J
//fd state: (0) 

ret = client.Write(fd267, []byte("PWWSzAdmQ_LBMKd1n_ZhDW2P_BiO2Ilf5ovom4tiKN2PygtjSMO_mZTGlbScL330oggCdhxtKMIwweuBg8jqymZL9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) PWWSzAdmQ_LBMKd1n_ZhDW2P_BiO2Ilf5ovom4tiKN2PygtjSMO_mZTGlbScL330oggCdhxtKMIwweuBg8jqymZL9

fd268 := client.Open("/j5i_lS91rD.txt", client.O_RDWR|client.O_CREATE)
if(fd268 < 0) {
  panic("open failed")
}

//fd state: (100) YqBedplBGNDxza1D7XOGbty9mUqt1Mo_qjs6_diZFxntTZAj5L9gI2WAcq8kR1poc9HArBMZZlSPLX6_FaSRnYv7i27u6zn_jLhI

ret = client.Write(fd254, []byte("Lg2bw8r13NWWn6MASOCfzzS1FblPuzxqs2yIhcwy1dQ6KoNKs2P"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) YqBedplBGNDxza1D7XOGbty9mUqt1Mo_qjs6_diZFxntTZAj5L9gI2WAcq8kR1poc9HArBMZZlSPLX6_FaSRnYv7i27u6zn_jLhILg2bw8r13NWWn6MASOCfzzS1FblPuzxqs2yIhcwy1dQ6KoNKs2P
//fd state: (89) PWWSzAdmQ_LBMKd1n_ZhDW2P_BiO2Ilf5ovom4tiKN2PygtjSMO_mZTGlbScL330oggCdhxtKMIwweuBg8jqymZL9

ret = client.Write(fd267, []byte("WFfvR_2eBvhMis_rg6uQKNrrNSPYLlwpwmsxs9wX1_qgdAQSXMlwTLAnRP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (147) PWWSzAdmQ_LBMKd1n_ZhDW2P_BiO2Ilf5ovom4tiKN2PygtjSMO_mZTGlbScL330oggCdhxtKMIwweuBg8jqymZL9WFfvR_2eBvhMis_rg6uQKNrrNSPYLlwpwmsxs9wX1_qgdAQSXMlwTLAnRP

buf, ret = client.Read(fd258, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TCzGHiVA_Lob_c0CcGIbvDsAQHT3RD51wkv7kleWt5cBvgvN9hlTcIVShen5q") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd256, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd268)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd251, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (83) KU8TlJ6F3Dv5H7WgrptSxLKlgzXo7ylvTejOrgtPkvD7bWotYSsLBovh3O3cCsebyb1lqb6spY4pud0106J

ret = client.Write(fd266, []byte("Hy_RrLpd7pD2_rxZQ9ylfII8WCpzoCp77_YJO7sZ5bcSsgTGZFulqyO1D5z8tW_vYa8o_DyNNiQ43n0aOI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (165) KU8TlJ6F3Dv5H7WgrptSxLKlgzXo7ylvTejOrgtPkvD7bWotYSsLBovh3O3cCsebyb1lqb6spY4pud0106JHy_RrLpd7pD2_rxZQ9ylfII8WCpzoCp77_YJO7sZ5bcSsgTGZFulqyO1D5z8tW_vYa8o_DyNNiQ43n0aOI

buf, ret = client.Read(fd254, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd263, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd269 := client.Open("/W4nJjAH1FT.txt", client.O_RDWR|client.O_CREATE)
if(fd269 < 0) {
  panic("open failed")
}

//fd state: (151) YqBedplBGNDxza1D7XOGbty9mUqt1Mo_qjs6_diZFxntTZAj5L9gI2WAcq8kR1poc9HArBMZZlSPLX6_FaSRnYv7i27u6zn_jLhILg2bw8r13NWWn6MASOCfzzS1FblPuzxqs2yIhcwy1dQ6KoNKs2P

ret = client.Write(fd254, []byte("RQDTdaZPDctVjmlXGqa7bSEoop7N0O_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (182) YqBedplBGNDxza1D7XOGbty9mUqt1Mo_qjs6_diZFxntTZAj5L9gI2WAcq8kR1poc9HArBMZZlSPLX6_FaSRnYv7i27u6zn_jLhILg2bw8r13NWWn6MASOCfzzS1FblPuzxqs2yIhcwy1dQ6KoNKs2PRQDTdaZPDctVjmlXGqa7bSEoop7N0O_

ret = client.Close(fd267)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd258, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Seek(fd258, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}

//fd state: (182) YqBedplBGNDxza1D7XOGbty9mUqt1Mo_qjs6_diZFxntTZAj5L9gI2WAcq8kR1poc9HArBMZZlSPLX6_FaSRnYv7i27u6zn_jLhILg2bw8r13NWWn6MASOCfzzS1FblPuzxqs2yIhcwy1dQ6KoNKs2PRQDTdaZPDctVjmlXGqa7bSEoop7N0O_

ret = client.Write(fd254, []byte("Dc8mvJkkDXD1rlVFC027Ys7mFlALxhYKVu25xgdAR_JrVUPsgSxLAoe4RHDmGSkv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (246) YqBedplBGNDxza1D7XOGbty9mUqt1Mo_qjs6_diZFxntTZAj5L9gI2WAcq8kR1poc9HArBMZZlSPLX6_FaSRnYv7i27u6zn_jLhILg2bw8r13NWWn6MASOCfzzS1FblPuzxqs2yIhcwy1dQ6KoNKs2PRQDTdaZPDctVjmlXGqa7bSEoop7N0O_Dc8mvJkkDXD1rlVFC027Ys7mFlALxhYKVu25xgdAR_JrVUPsgSxLAoe4RHDmGSkv

fd270 := client.Open("/Ju0L1mzSEa.txt", client.O_RDWR|client.O_CREATE)
if(fd270 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd266, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd271 := client.Open("/PqQOVlfG3g.txt", client.O_RDWR|client.O_CREATE)
if(fd271 < 0) {
  panic("open failed")
}


ret = client.Seek(fd266, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd262, []byte("yIndr06_InYbZguoemjp3Rn6fbapt7c1yyKXTgLEWluD_P1Jdg4BkLjdIyCdk9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) yIndr06_InYbZguoemjp3Rn6fbapt7c1yyKXTgLEWluD_P1Jdg4BkLjdIyCdk9

ret = client.Seek(fd270, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) t2jc1Ex9YC8URWa2kUNvDRVn8uRNu2R7iNKiFbfUEtoCC0kqNz2VmkitU5f9TOntVXN3lq8pDz4rMcNCvwBlhnr_EjqR1enusssaVf

ret = client.Write(fd269, []byte("G0q4A8IJ62QGWrABuNdQaOgFXjMsy9bHArTaGUTgLpVqcXRbNzeKqlfNmVi8p89M_JYHgKw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (71) G0q4A8IJ62QGWrABuNdQaOgFXjMsy9bHArTaGUTgLpVqcXRbNzeKqlfNmVi8p89M_JYHgKwpDz4rMcNCvwBlhnr_EjqR1enusssaVf

buf, ret = client.Read(fd262, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd256)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd254)
if(ret != 0) {
  panic("close failed")
}


fd272 := client.Open("/jg0Jvu7lJx.txt", client.O_RDWR|client.O_CREATE)
if(fd272 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd263, []byte("BA22MLSZ62Fw0PQ_Ta96BELBzNoR2cnwH7fPfTnts"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) BA22MLSZ62Fw0PQ_Ta96BELBzNoR2cnwH7fPfTnts
//fd state: (0) 

ret = client.Write(fd272, []byte("fyNCTaNL_OO1r17WxQf3lUtCq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) fyNCTaNL_OO1r17WxQf3lUtCq

ret = client.Seek(fd270, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd266, 124, client.SEEK_SET)
if(ret != 124) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 124)
  panic("seek failed")
}


buf, ret = client.Read(fd263, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd263, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}

//fd state: (34) TCzGHiVA_Lob_c0CcGIbvDsAQHT3RD51wkv7kleWt5cBvgvN9hlTcIVShen5q

ret = client.Write(fd258, []byte("ho4nD1fsflImtXzfLck8XY54"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) TCzGHiVA_Lob_c0CcGIbvDsAQHT3RD51wkho4nD1fsflImtXzfLck8XY54n5q
//fd state: (124) KU8TlJ6F3Dv5H7WgrptSxLKlgzXo7ylvTejOrgtPkvD7bWotYSsLBovh3O3cCsebyb1lqb6spY4pud0106JHy_RrLpd7pD2_rxZQ9ylfII8WCpzoCp77_YJO7sZ5bcSsgTGZFulqyO1D5z8tW_vYa8o_DyNNiQ43n0aOI

ret = client.Write(fd266, []byte("eTcrayhPYwzL8YKpM1X8FRUonHRPM01kO78pMn_Yfn1f5cEcFrkf6gHyVLZ12XOrhOGQL6F9ain0DGlXBBpz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (208) KU8TlJ6F3Dv5H7WgrptSxLKlgzXo7ylvTejOrgtPkvD7bWotYSsLBovh3O3cCsebyb1lqb6spY4pud0106JHy_RrLpd7pD2_rxZQ9ylfII8WCpzoCp77_YJO7sZ5eTcrayhPYwzL8YKpM1X8FRUonHRPM01kO78pMn_Yfn1f5cEcFrkf6gHyVLZ12XOrhOGQL6F9ain0DGlXBBpz

buf, ret = client.Read(fd266, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd262)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd271, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}

//fd state: (58) TCzGHiVA_Lob_c0CcGIbvDsAQHT3RD51wkho4nD1fsflImtXzfLck8XY54n5q

ret = client.Write(fd258, []byte("bkULQ7jfAzhIXr3HrV_TzquIp6j0xbNXvpdKiP8mYKVkZGQq2okl7qVEkCbECMibogKIFKzmOogndGrlwArmjAw9fepmIMh4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) TCzGHiVA_Lob_c0CcGIbvDsAQHT3RD51wkho4nD1fsflImtXzfLck8XY54bkULQ7jfAzhIXr3HrV_TzquIp6j0xbNXvpdKiP8mYKVkZGQq2okl7qVEkCbECMibogKIFKzmOogndGrlwArmjAw9fepmIMh4

fd273 := client.Open("/1oGmSJ4tdj.txt", client.O_RDWR|client.O_CREATE)
if(fd273 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd263, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Z62Fw0PQ_Ta96BELBzNoR2cnwH") {
  panic("wrong data returned")
}


ret = client.Close(fd266)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd273)
if(ret != 0) {
  panic("close failed")
}


fd274 := client.Open("/DbkC5q9e6v.txt", client.O_RDWR|client.O_CREATE)
if(fd274 < 0) {
  panic("open failed")
}


ret = client.Close(fd274)
if(ret != 0) {
  panic("close failed")
}

//fd state: (25) fyNCTaNL_OO1r17WxQf3lUtCq

ret = client.Write(fd272, []byte("XYTr_J9MggV7__y4AaeTaaeb1iL6GCohXz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) fyNCTaNL_OO1r17WxQf3lUtCqXYTr_J9MggV7__y4AaeTaaeb1iL6GCohXz

buf, ret = client.Read(fd269, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pDz4rMcNCvwBlhnr_EjqR1enusssaVf") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd269, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd271, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "tw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkF") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd258, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd263)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd251, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd269)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd251, []byte("RTHwW7LaSlo5SQ8nEir9l4ZDsk7l"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) RTHwW7LaSlo5SQ8nEir9l4ZDsk7l

fd275 := client.Open("/T7YhA3pRq0.txt", client.O_RDWR|client.O_CREATE)
if(fd275 < 0) {
  panic("open failed")
}


ret = client.Close(fd272)
if(ret != 0) {
  panic("close failed")
}


fd276 := client.Open("/4mPDBgaYe1.txt", client.O_RDWR|client.O_CREATE)
if(fd276 < 0) {
  panic("open failed")
}


ret = client.Seek(fd275, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd277 := client.Open("/trssyo58wD.txt", client.O_RDWR|client.O_CREATE)
if(fd277 < 0) {
  panic("open failed")
}

//fd state: (62) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFhYf8EavuTrfTDQRB5

ret = client.Write(fd271, []byte("VWQ6m8Y0DgGvPpyExkwBYSeiZ82j4Yjy58Vi6vKJW0TF2mWY2RUtXPq6O5YU2qsaY4rRGzwM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFVWQ6m8Y0DgGvPpyExkwBYSeiZ82j4Yjy58Vi6vKJW0TF2mWY2RUtXPq6O5YU2qsaY4rRGzwM
//fd state: (134) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFVWQ6m8Y0DgGvPpyExkwBYSeiZ82j4Yjy58Vi6vKJW0TF2mWY2RUtXPq6O5YU2qsaY4rRGzwM

ret = client.Write(fd271, []byte("0lPqawnG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (142) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFVWQ6m8Y0DgGvPpyExkwBYSeiZ82j4Yjy58Vi6vKJW0TF2mWY2RUtXPq6O5YU2qsaY4rRGzwM0lPqawnG

buf, ret = client.Read(fd271, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd277, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd270, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd277, []byte("bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue
//fd state: (142) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFVWQ6m8Y0DgGvPpyExkwBYSeiZ82j4Yjy58Vi6vKJW0TF2mWY2RUtXPq6O5YU2qsaY4rRGzwM0lPqawnG

ret = client.Write(fd271, []byte("XmPuxMashyPeIMszs4FEtgiMMSmLgaPiVXEnPUAZGGuLTOCK9Ze3PUvhnb0mIpz29dJcETQX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (214) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFVWQ6m8Y0DgGvPpyExkwBYSeiZ82j4Yjy58Vi6vKJW0TF2mWY2RUtXPq6O5YU2qsaY4rRGzwM0lPqawnGXmPuxMashyPeIMszs4FEtgiMMSmLgaPiVXEnPUAZGGuLTOCK9Ze3PUvhnb0mIpz29dJcETQX

ret = client.Seek(fd251, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


ret = client.Close(fd251)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) MA5WnxIA8lEbbDmlPHLiB9HbElo7j7pP8y2Uuju4Jm914QburF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1CMUFIN7wkaxagbBnaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1

ret = client.Write(fd276, []byte("m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtpZFwB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtpZFwBF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1CMUFIN7wkaxagbBnaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1

ret = client.Close(fd270)
if(ret != 0) {
  panic("close failed")
}


fd278 := client.Open("/wA5Kf_Ltjb.txt", client.O_RDWR|client.O_CREATE)
if(fd278 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd258, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd258, 106, client.SEEK_SET)
if(ret != 106) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 106)
  panic("seek failed")
}


ret = client.Seek(fd277, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


ret = client.Seek(fd278, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd277, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xu93Eo3V1ue") {
  panic("wrong data returned")
}

//fd state: (106) TCzGHiVA_Lob_c0CcGIbvDsAQHT3RD51wkho4nD1fsflImtXzfLck8XY54bkULQ7jfAzhIXr3HrV_TzquIp6j0xbNXvpdKiP8mYKVkZGQq2okl7qVEkCbECMibogKIFKzmOogndGrlwArmjAw9fepmIMh4

ret = client.Write(fd258, []byte("olMkupY3sBNNUMHnRAizIE7mRHEODHkQV3eFbHYUap8ZaDLejAXnoecog3Sx8Lg5Gi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (172) TCzGHiVA_Lob_c0CcGIbvDsAQHT3RD51wkho4nD1fsflImtXzfLck8XY54bkULQ7jfAzhIXr3HrV_TzquIp6j0xbNXvpdKiP8mYKVkZGQqolMkupY3sBNNUMHnRAizIE7mRHEODHkQV3eFbHYUap8ZaDLejAXnoecog3Sx8Lg5Gi
//fd state: (49) bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue

ret = client.Write(fd277, []byte("1fkR99cxyv_7uH_HpLc50V8HufJztyunGc251vQqXkSwWxjzstgTQl3ycKGG0q1RVFN1yTpQU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue1fkR99cxyv_7uH_HpLc50V8HufJztyunGc251vQqXkSwWxjzstgTQl3ycKGG0q1RVFN1yTpQU

ret = client.Close(fd258)
if(ret != 0) {
  panic("close failed")
}

//fd state: (214) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFVWQ6m8Y0DgGvPpyExkwBYSeiZ82j4Yjy58Vi6vKJW0TF2mWY2RUtXPq6O5YU2qsaY4rRGzwM0lPqawnGXmPuxMashyPeIMszs4FEtgiMMSmLgaPiVXEnPUAZGGuLTOCK9Ze3PUvhnb0mIpz29dJcETQX

ret = client.Write(fd271, []byte("orH1KZswBHxi9I4KF2r091Q3evvZmrX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (245) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFVWQ6m8Y0DgGvPpyExkwBYSeiZ82j4Yjy58Vi6vKJW0TF2mWY2RUtXPq6O5YU2qsaY4rRGzwM0lPqawnGXmPuxMashyPeIMszs4FEtgiMMSmLgaPiVXEnPUAZGGuLTOCK9Ze3PUvhnb0mIpz29dJcETQXorH1KZswBHxi9I4KF2r091Q3evvZmrX

fd279 := client.Open("/HLZHKZfA0V.txt", client.O_RDWR|client.O_CREATE)
if(fd279 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd278, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd271, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd277, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}


fd280 := client.Open("/AYwrkuI1uC.txt", client.O_RDWR|client.O_CREATE)
if(fd280 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd275, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd277, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_HpLc50V8HufJztyunGc251vQqXkSwWxjzstg") {
  panic("wrong data returned")
}


ret = client.Close(fd280)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd278, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd281 := client.Open("/yG2d2G7zVG.txt", client.O_RDWR|client.O_CREATE)
if(fd281 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd276, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "F34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRK") {
  panic("wrong data returned")
}


fd282 := client.Open("/CY9cjfit5S.txt", client.O_RDWR|client.O_CREATE)
if(fd282 < 0) {
  panic("open failed")
}


ret = client.Close(fd275)
if(ret != 0) {
  panic("close failed")
}

//fd state: (100) bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue1fkR99cxyv_7uH_HpLc50V8HufJztyunGc251vQqXkSwWxjzstgTQl3ycKGG0q1RVFN1yTpQU

ret = client.Write(fd277, []byte("pKiD0Z1VvXkUDCTCUZrxg15u36JIFHKt7JXFbQzOKa0HNxhmeYDQe7CHUCe53mHaNzN7ZexrUIDUYm2XTKaaCOBdKa95WQv2OTQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (199) bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue1fkR99cxyv_7uH_HpLc50V8HufJztyunGc251vQqXkSwWxjzstgpKiD0Z1VvXkUDCTCUZrxg15u36JIFHKt7JXFbQzOKa0HNxhmeYDQe7CHUCe53mHaNzN7ZexrUIDUYm2XTKaaCOBdKa95WQv2OTQ

fd283 := client.Open("/4Ki0uUgl43.txt", client.O_RDWR|client.O_CREATE)
if(fd283 < 0) {
  panic("open failed")
}


fd284 := client.Open("/zGXdEW9YZx.txt", client.O_RDWR|client.O_CREATE)
if(fd284 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd282, []byte("TWPh4AOrSRcZF3jzVq362OX4eJo2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) TWPh4AOrSRcZF3jzVq362OX4eJo2

buf, ret = client.Read(fd277, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd283, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd285 := client.Open("/7evLhsHpf6.txt", client.O_RDWR|client.O_CREATE)
if(fd285 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd285, []byte("NtjInUxscPBFcvuSXDvyLBh6fUU_rLk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk

buf, ret = client.Read(fd279, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd283, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd284, []byte("FF4YZ9rQweGAGgQ71ymkZzSS96iul5asXIPDntk4aSMXTX7HlrOHdh2j11Dq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) FF4YZ9rQweGAGgQ71ymkZzSS96iul5asXIPDntk4aSMXTX7HlrOHdh2j11Dq

buf, ret = client.Read(fd284, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd271, 114, client.SEEK_SET)
if(ret != 114) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 114)
  panic("seek failed")
}


buf, ret = client.Read(fd285, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd285, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (60) FF4YZ9rQweGAGgQ71ymkZzSS96iul5asXIPDntk4aSMXTX7HlrOHdh2j11Dq

ret = client.Write(fd284, []byte("phVDNMppzKG4Mg7_45nTLkoTG9VhFLHqb2kRQ8y6vnFKKK0O4sKlxynXr8r3lxs_EuSNSc0bJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) FF4YZ9rQweGAGgQ71ymkZzSS96iul5asXIPDntk4aSMXTX7HlrOHdh2j11DqphVDNMppzKG4Mg7_45nTLkoTG9VhFLHqb2kRQ8y6vnFKKK0O4sKlxynXr8r3lxs_EuSNSc0bJ
//fd state: (105) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtpZFwBF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKMPaK1CMUFIN7wkaxagbBnaimMiN5eWnTNWUq9T4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1

ret = client.Write(fd276, []byte("eFlmnD0RVkshG9lBtyMQm06Kxa5DBJrPccapE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (142) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtpZFwBF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKeFlmnD0RVkshG9lBtyMQm06Kxa5DBJrPccapET4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1

fd286 := client.Open("/mMuSzkc3oY.txt", client.O_RDWR|client.O_CREATE)
if(fd286 < 0) {
  panic("open failed")
}


ret = client.Seek(fd271, 137, client.SEEK_SET)
if(ret != 137) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 137)
  panic("seek failed")
}


fd287 := client.Open("/bCLvtka1QB.txt", client.O_RDWR|client.O_CREATE)
if(fd287 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd286, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd288 := client.Open("/D3Fa1l3zaJ.txt", client.O_RDWR|client.O_CREATE)
if(fd288 < 0) {
  panic("open failed")
}

//fd state: (137) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFVWQ6m8Y0DgGvPpyExkwBYSeiZ82j4Yjy58Vi6vKJW0TF2mWY2RUtXPq6O5YU2qsaY4rRGzwM0lPqawnGXmPuxMashyPeIMszs4FEtgiMMSmLgaPiVXEnPUAZGGuLTOCK9Ze3PUvhnb0mIpz29dJcETQXorH1KZswBHxi9I4KF2r091Q3evvZmrX

ret = client.Write(fd271, []byte("UbMDHwnNz94NmZg4WxuAl3IPJ3XkO3AnNPFD0UOnKra3FwNRCETGljSYKSWqtBajhzNbfDb8SNJZHF5VI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (218) GhVao4WfDrSAI17MSfEjfMUx3jNtw4cQnu01TEehFECZjTMDNHbOsuc6ih2RkFVWQ6m8Y0DgGvPpyExkwBYSeiZ82j4Yjy58Vi6vKJW0TF2mWY2RUtXPq6O5YU2qsaY4rRGzwM0lPUbMDHwnNz94NmZg4WxuAl3IPJ3XkO3AnNPFD0UOnKra3FwNRCETGljSYKSWqtBajhzNbfDb8SNJZHF5VIKZswBHxi9I4KF2r091Q3evvZmrX

fd289 := client.Open("/eXNFP7XOhv.txt", client.O_RDWR|client.O_CREATE)
if(fd289 < 0) {
  panic("open failed")
}

//fd state: (31) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk

ret = client.Write(fd285, []byte("4Q02N0yeM1OPIie1f8TQY2B_TI1S4IVuibP5hswxdWqXBi7_G3sl_zn83zO0BnaSg6h9_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk4Q02N0yeM1OPIie1f8TQY2B_TI1S4IVuibP5hswxdWqXBi7_G3sl_zn83zO0BnaSg6h9_
//fd state: (0) 

ret = client.Write(fd288, []byte("NVr_siCEPrkOW41_9_XDwfeh5bf3jugCcUJkL551Rx3KJQw5XbK0J7VeHhGkBFodPeK57VYo56pv_PHwsA1I4oSE2Ftk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) NVr_siCEPrkOW41_9_XDwfeh5bf3jugCcUJkL551Rx3KJQw5XbK0J7VeHhGkBFodPeK57VYo56pv_PHwsA1I4oSE2Ftk
//fd state: (0) 

ret = client.Write(fd281, []byte("RK6qtQtzKTyx2ZVrzFblyHAajpjbO5yV9SMySUAaAsjZFUA0RBJiciHB05ijYTrqSEsgYSIYyBuJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) RK6qtQtzKTyx2ZVrzFblyHAajpjbO5yV9SMySUAaAsjZFUA0RBJiciHB05ijYTrqSEsgYSIYyBuJ

fd290 := client.Open("/40ImwTMgD0.txt", client.O_RDWR|client.O_CREATE)
if(fd290 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd290, []byte("z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeJQ_tPxFcja"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeJQ_tPxFcja
//fd state: (133) FF4YZ9rQweGAGgQ71ymkZzSS96iul5asXIPDntk4aSMXTX7HlrOHdh2j11DqphVDNMppzKG4Mg7_45nTLkoTG9VhFLHqb2kRQ8y6vnFKKK0O4sKlxynXr8r3lxs_EuSNSc0bJ

ret = client.Write(fd284, []byte("RmzuhZhgxCpcOveYTB5LVxKLD7jQ2HXKmq7VbOlYjfejgW3N"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (181) FF4YZ9rQweGAGgQ71ymkZzSS96iul5asXIPDntk4aSMXTX7HlrOHdh2j11DqphVDNMppzKG4Mg7_45nTLkoTG9VhFLHqb2kRQ8y6vnFKKK0O4sKlxynXr8r3lxs_EuSNSc0bJRmzuhZhgxCpcOveYTB5LVxKLD7jQ2HXKmq7VbOlYjfejgW3N
//fd state: (100) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk4Q02N0yeM1OPIie1f8TQY2B_TI1S4IVuibP5hswxdWqXBi7_G3sl_zn83zO0BnaSg6h9_

ret = client.Write(fd285, []byte("wc3kfTOqFt7kjILO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk4Q02N0yeM1OPIie1f8TQY2B_TI1S4IVuibP5hswxdWqXBi7_G3sl_zn83zO0BnaSg6h9_wc3kfTOqFt7kjILO
//fd state: (116) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk4Q02N0yeM1OPIie1f8TQY2B_TI1S4IVuibP5hswxdWqXBi7_G3sl_zn83zO0BnaSg6h9_wc3kfTOqFt7kjILO

ret = client.Write(fd285, []byte("dzA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (119) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk4Q02N0yeM1OPIie1f8TQY2B_TI1S4IVuibP5hswxdWqXBi7_G3sl_zn83zO0BnaSg6h9_wc3kfTOqFt7kjILOdzA
//fd state: (0) 

ret = client.Write(fd286, []byte("mBe4GzIN2ngOo2ZBDUcFHKlxgERRUUoZjxXM3CDYZS7j357qcsMHdNHHwiaBrHGx7vLJGj5CgqXzEC2vvPDkF_AMt4onGUb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) mBe4GzIN2ngOo2ZBDUcFHKlxgERRUUoZjxXM3CDYZS7j357qcsMHdNHHwiaBrHGx7vLJGj5CgqXzEC2vvPDkF_AMt4onGUb

ret = client.Close(fd287)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd276, 208, client.SEEK_SET)
if(ret != 208) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 208)
  panic("seek failed")
}


buf, ret = client.Read(fd285, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (88) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeJQ_tPxFcja

ret = client.Write(fd290, []byte("mQnZ1IY35xtwGAYiPxE81Uk5DFyx7SKsWPIarb9RZyuA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeJQ_tPxFcjamQnZ1IY35xtwGAYiPxE81Uk5DFyx7SKsWPIarb9RZyuA

ret = client.Seek(fd271, 182, client.SEEK_SET)
if(ret != 182) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 182)
  panic("seek failed")
}


buf, ret = client.Read(fd290, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (199) bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue1fkR99cxyv_7uH_HpLc50V8HufJztyunGc251vQqXkSwWxjzstgpKiD0Z1VvXkUDCTCUZrxg15u36JIFHKt7JXFbQzOKa0HNxhmeYDQe7CHUCe53mHaNzN7ZexrUIDUYm2XTKaaCOBdKa95WQv2OTQ

ret = client.Write(fd277, []byte("N0WqJyGe3pvCUyJduZ6RLyuDGB38Qeycoz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (233) bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue1fkR99cxyv_7uH_HpLc50V8HufJztyunGc251vQqXkSwWxjzstgpKiD0Z1VvXkUDCTCUZrxg15u36JIFHKt7JXFbQzOKa0HNxhmeYDQe7CHUCe53mHaNzN7ZexrUIDUYm2XTKaaCOBdKa95WQv2OTQN0WqJyGe3pvCUyJduZ6RLyuDGB38Qeycoz

ret = client.Seek(fd284, 69, client.SEEK_SET)
if(ret != 69) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 69)
  panic("seek failed")
}

//fd state: (119) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk4Q02N0yeM1OPIie1f8TQY2B_TI1S4IVuibP5hswxdWqXBi7_G3sl_zn83zO0BnaSg6h9_wc3kfTOqFt7kjILOdzA

ret = client.Write(fd285, []byte("i1GJy0hspLHnBVlmKM6O4rv8EhFXoSbqbNeM7_8unhlaiJs1Y1ncFlJRdZ_Ihyj72uENNKg4tFF5CxHGAUbvZbzP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (207) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk4Q02N0yeM1OPIie1f8TQY2B_TI1S4IVuibP5hswxdWqXBi7_G3sl_zn83zO0BnaSg6h9_wc3kfTOqFt7kjILOdzAi1GJy0hspLHnBVlmKM6O4rv8EhFXoSbqbNeM7_8unhlaiJs1Y1ncFlJRdZ_Ihyj72uENNKg4tFF5CxHGAUbvZbzP

ret = client.Close(fd282)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd278, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd285, 185, client.SEEK_SET)
if(ret != 185) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 185)
  panic("seek failed")
}

//fd state: (233) bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue1fkR99cxyv_7uH_HpLc50V8HufJztyunGc251vQqXkSwWxjzstgpKiD0Z1VvXkUDCTCUZrxg15u36JIFHKt7JXFbQzOKa0HNxhmeYDQe7CHUCe53mHaNzN7ZexrUIDUYm2XTKaaCOBdKa95WQv2OTQN0WqJyGe3pvCUyJduZ6RLyuDGB38Qeycoz

ret = client.Write(fd277, []byte("KXBHYSm5wfNxjmDoAEcmt0QOInVrvfXveG3ZM5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (271) bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue1fkR99cxyv_7uH_HpLc50V8HufJztyunGc251vQqXkSwWxjzstgpKiD0Z1VvXkUDCTCUZrxg15u36JIFHKt7JXFbQzOKa0HNxhmeYDQe7CHUCe53mHaNzN7ZexrUIDUYm2XTKaaCOBdKa95WQv2OTQN0WqJyGe3pvCUyJduZ6RLyuDGB38QeycozKXBHYSm5wfNxjmDoAEcmt0QOInVrvfXveG3ZM5

ret = client.Seek(fd286, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


buf, ret = client.Read(fd286, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "DYZS7j357qcsMHdNHHw") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd277, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd286, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "iaBrHGx7vLJGj5CgqXzEC2vvPDkF_AMt4onGUb") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd277, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (132) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeJQ_tPxFcjamQnZ1IY35xtwGAYiPxE81Uk5DFyx7SKsWPIarb9RZyuA

ret = client.Write(fd290, []byte("OLg0za72lyc75BLdZaJlNcxaFVesWssaLzQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (167) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeJQ_tPxFcjamQnZ1IY35xtwGAYiPxE81Uk5DFyx7SKsWPIarb9RZyuAOLg0za72lyc75BLdZaJlNcxaFVesWssaLzQ

ret = client.Close(fd277)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd289)
if(ret != 0) {
  panic("close failed")
}


fd291 := client.Open("/N4MHh3ZTvc.txt", client.O_RDWR|client.O_CREATE)
if(fd291 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd288, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd281)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd271, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wNRCETGljSYKSWqtBajhzNbf") {
  panic("wrong data returned")
}


ret = client.Seek(fd279, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd286)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd290, 118, client.SEEK_SET)
if(ret != 118) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 118)
  panic("seek failed")
}

//fd state: (92) NVr_siCEPrkOW41_9_XDwfeh5bf3jugCcUJkL551Rx3KJQw5XbK0J7VeHhGkBFodPeK57VYo56pv_PHwsA1I4oSE2Ftk

ret = client.Write(fd288, []byte("W_szBL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) NVr_siCEPrkOW41_9_XDwfeh5bf3jugCcUJkL551Rx3KJQw5XbK0J7VeHhGkBFodPeK57VYo56pv_PHwsA1I4oSE2FtkW_szBL
//fd state: (118) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeJQ_tPxFcjamQnZ1IY35xtwGAYiPxE81Uk5DFyx7SKsWPIarb9RZyuAOLg0za72lyc75BLdZaJlNcxaFVesWssaLzQ

ret = client.Write(fd290, []byte("VCE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (121) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeJQ_tPxFcjamQnZ1IY35xtwGAYiPxE81Uk5DFyx7SVCEPIarb9RZyuAOLg0za72lyc75BLdZaJlNcxaFVesWssaLzQ
//fd state: (0) 

ret = client.Write(fd279, []byte("6vxnhyC8raWvJo9Z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (16) 6vxnhyC8raWvJo9Z
//fd state: (0) 

ret = client.Write(fd291, []byte("rFxOLOww43w3B0MhWvOFR4AsbzEmHHPComhFjHP9bb0bHUhecNFP24IHE44"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) rFxOLOww43w3B0MhWvOFR4AsbzEmHHPComhFjHP9bb0bHUhecNFP24IHE44

ret = client.Seek(fd285, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


ret = client.Close(fd271)
if(ret != 0) {
  panic("close failed")
}


fd292 := client.Open("/ZaBfLNx_Cb.txt", client.O_RDWR|client.O_CREATE)
if(fd292 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd291, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd284, 162, client.SEEK_SET)
if(ret != 162) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 162)
  panic("seek failed")
}

//fd state: (0) aB5m3mQrTStEqW0XsuCjIsAZdkKSFRgP9hcwGU2yBpScrgR9P8F6ZeY_s78yY2InTBU8_mG7fySwLsVYAHfgUQ0RBx299sghdSy

ret = client.Write(fd292, []byte("B6xmBErBDPVgp5_V2nnQwuACisEK4hIsX3bkithlS9LMqs4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) B6xmBErBDPVgp5_V2nnQwuACisEK4hIsX3bkithlS9LMqs49P8F6ZeY_s78yY2InTBU8_mG7fySwLsVYAHfgUQ0RBx299sghdSy

ret = client.Seek(fd290, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}

//fd state: (98) NVr_siCEPrkOW41_9_XDwfeh5bf3jugCcUJkL551Rx3KJQw5XbK0J7VeHhGkBFodPeK57VYo56pv_PHwsA1I4oSE2FtkW_szBL

ret = client.Write(fd288, []byte("QflJ9akGFKOafXzawTVY2iEaB0sHiTOcjHu9OmylM7k86MH3hFWzZMbpb_A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (157) NVr_siCEPrkOW41_9_XDwfeh5bf3jugCcUJkL551Rx3KJQw5XbK0J7VeHhGkBFodPeK57VYo56pv_PHwsA1I4oSE2FtkW_szBLQflJ9akGFKOafXzawTVY2iEaB0sHiTOcjHu9OmylM7k86MH3hFWzZMbpb_A

fd293 := client.Open("/X8KDKPPquI.txt", client.O_RDWR|client.O_CREATE)
if(fd293 < 0) {
  panic("open failed")
}


ret = client.Seek(fd284, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}

//fd state: (78) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeJQ_tPxFcjamQnZ1IY35xtwGAYiPxE81Uk5DFyx7SVCEPIarb9RZyuAOLg0za72lyc75BLdZaJlNcxaFVesWssaLzQ

ret = client.Write(fd290, []byte("S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeSQ_tPxFcjamQnZ1IY35xtwGAYiPxE81Uk5DFyx7SVCEPIarb9RZyuAOLg0za72lyc75BLdZaJlNcxaFVesWssaLzQ

ret = client.Seek(fd279, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Seek(fd278, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd291, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd276, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyD") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd293, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd294 := client.Open("/W14ZY4nXnm.txt", client.O_RDWR|client.O_CREATE)
if(fd294 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd283, []byte("x1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) x1

ret = client.Close(fd294)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd293, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd276, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "sFP4lawaUhuZmMpiTSPNu1") {
  panic("wrong data returned")
}


ret = client.Close(fd283)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd279, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (33) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk4Q02N0yeM1OPIie1f8TQY2B_TI1S4IVuibP5hswxdWqXBi7_G3sl_zn83zO0BnaSg6h9_wc3kfTOqFt7kjILOdzAi1GJy0hspLHnBVlmKM6O4rv8EhFXoSbqbNeM7_8unhlaiJs1Y1ncFlJRdZ_Ihyj72uENNKg4tFF5CxHGAUbvZbzP

ret = client.Write(fd285, []byte("Wu4TJIdsYHGcQinFgU8MQJB5c7C904nFgHH2oRjZumBuIdNtOaJGGz53HN_9OvM2PkFpGz_Tic"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) NtjInUxscPBFcvuSXDvyLBh6fUU_rLk4QWu4TJIdsYHGcQinFgU8MQJB5c7C904nFgHH2oRjZumBuIdNtOaJGGz53HN_9OvM2PkFpGz_TicqFt7kjILOdzAi1GJy0hspLHnBVlmKM6O4rv8EhFXoSbqbNeM7_8unhlaiJs1Y1ncFlJRdZ_Ihyj72uENNKg4tFF5CxHGAUbvZbzP
//fd state: (0) 

ret = client.Write(fd293, []byte("lOx8cyAyByxnB759OnbcBUINxBW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) lOx8cyAyByxnB759OnbcBUINxBW

ret = client.Seek(fd291, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


ret = client.Seek(fd285, 136, client.SEEK_SET)
if(ret != 136) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 136)
  panic("seek failed")
}


buf, ret = client.Read(fd293, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd278, []byte("19Q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) 19Q

ret = client.Seek(fd290, 120, client.SEEK_SET)
if(ret != 120) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 120)
  panic("seek failed")
}


ret = client.Close(fd293)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd291)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd290, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "EPIarb9RZyuAOLg0za72lyc75BLdZaJlNcxaF") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd278, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd288, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


buf, ret = client.Read(fd292, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9P8F6ZeY_s78yY2InTBU8_mG7f") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd290, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "V") {
  panic("wrong data returned")
}


ret = client.Seek(fd292, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}

//fd state: (291) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtpZFwBF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKeFlmnD0RVkshG9lBtyMQm06Kxa5DBJrPccapET4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1

ret = client.Write(fd276, []byte("NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (334) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtpZFwBF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKeFlmnD0RVkshG9lBtyMQm06Kxa5DBJrPccapET4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

ret = client.Close(fd284)
if(ret != 0) {
  panic("close failed")
}


fd295 := client.Open("/VyUT3k7cYP.txt", client.O_RDWR|client.O_CREATE)
if(fd295 < 0) {
  panic("open failed")
}

//fd state: (12) NVr_siCEPrkOW41_9_XDwfeh5bf3jugCcUJkL551Rx3KJQw5XbK0J7VeHhGkBFodPeK57VYo56pv_PHwsA1I4oSE2FtkW_szBLQflJ9akGFKOafXzawTVY2iEaB0sHiTOcjHu9OmylM7k86MH3hFWzZMbpb_A

ret = client.Write(fd288, []byte("3QuyKKM7zVBVqBtqFEIfekTxXv4_rko3SElG76UK9mdXwJgSAdj7Jq4QLQNYj0CxM6LYfMl8f51ufGcf3NbRj_tCQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) NVr_siCEPrkO3QuyKKM7zVBVqBtqFEIfekTxXv4_rko3SElG76UK9mdXwJgSAdj7Jq4QLQNYj0CxM6LYfMl8f51ufGcf3NbRj_tCQJ9akGFKOafXzawTVY2iEaB0sHiTOcjHu9OmylM7k86MH3hFWzZMbpb_A

ret = client.Close(fd278)
if(ret != 0) {
  panic("close failed")
}

//fd state: (101) NVr_siCEPrkO3QuyKKM7zVBVqBtqFEIfekTxXv4_rko3SElG76UK9mdXwJgSAdj7Jq4QLQNYj0CxM6LYfMl8f51ufGcf3NbRj_tCQJ9akGFKOafXzawTVY2iEaB0sHiTOcjHu9OmylM7k86MH3hFWzZMbpb_A

ret = client.Write(fd288, []byte("Oy8T86inF9_D2EIbBPFASVAfp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (126) NVr_siCEPrkO3QuyKKM7zVBVqBtqFEIfekTxXv4_rko3SElG76UK9mdXwJgSAdj7Jq4QLQNYj0CxM6LYfMl8f51ufGcf3NbRj_tCQOy8T86inF9_D2EIbBPFASVAfpiTOcjHu9OmylM7k86MH3hFWzZMbpb_A

ret = client.Close(fd285)
if(ret != 0) {
  panic("close failed")
}

//fd state: (3) 6vxnhyC8raWvJo9Z

ret = client.Write(fd279, []byte("r8eBjSwOkQ4gtSBXYdHhaVos5oOAuVhx_ziqRHtzpfhzpWmXCWK0UKsrGjUscIIkeC3diJDQjwUNgX1g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) 6vxr8eBjSwOkQ4gtSBXYdHhaVos5oOAuVhx_ziqRHtzpfhzpWmXCWK0UKsrGjUscIIkeC3diJDQjwUNgX1g

fd296 := client.Open("/QbUk9Mk7mu.txt", client.O_RDWR|client.O_CREATE)
if(fd296 < 0) {
  panic("open failed")
}


ret = client.Seek(fd276, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}


buf, ret = client.Read(fd296, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "CawTj8gqDjcxZqhqv8G") {
  panic("wrong data returned")
}

//fd state: (44) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtpZFwBF34wf8SOSY3gRxVbd3AcaccKEZuiQs6zkYMsnFYw5BBaExgzFQiO7qRKeFlmnD0RVkshG9lBtyMQm06Kxa5DBJrPccapET4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

ret = client.Write(fd276, []byte("gIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtBELXt9E_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtBELXt9E_BtyMQm06Kxa5DBJrPccapET4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

buf, ret = client.Read(fd288, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "iTOcjHu9OmylM7k86MH3hFWzZMbpb_A") {
  panic("wrong data returned")
}


fd297 := client.Open("/1fAm1cEKrd.txt", client.O_RDWR|client.O_CREATE)
if(fd297 < 0) {
  panic("open failed")
}

//fd state: (19) CawTj8gqDjcxZqhqv8G

ret = client.Write(fd296, []byte("MPFXN8uMyZd1vlJOCjaA8LNlI95gnBbOSxDa7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) CawTj8gqDjcxZqhqv8GMPFXN8uMyZd1vlJOCjaA8LNlI95gnBbOSxDa7
//fd state: (0) 

ret = client.Write(fd297, []byte("BCK67o5aYAZT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) BCK67o5aYAZT

fd298 := client.Open("/HCfdrAN7le.txt", client.O_RDWR|client.O_CREATE)
if(fd298 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd295, []byte("WPPWZBwopU6eLU8CmN8ea9zXXyLbKNA03wUDdU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) WPPWZBwopU6eLU8CmN8ea9zXXyLbKNA03wUDdU

ret = client.Close(fd297)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd288, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


fd299 := client.Open("/SbQouCL1vm.txt", client.O_RDWR|client.O_CREATE)
if(fd299 < 0) {
  panic("open failed")
}

//fd state: (120) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtBELXt9E_BtyMQm06Kxa5DBJrPccapET4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

ret = client.Write(fd276, []byte("seromcoQvMvu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtBELXt9E_seromcoQvMvuDBJrPccapET4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

fd300 := client.Open("/W14ZY4nXnm.txt", client.O_RDWR|client.O_CREATE)
if(fd300 < 0) {
  panic("open failed")
}


ret = client.Seek(fd296, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


ret = client.Seek(fd279, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


ret = client.Close(fd299)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd298, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd300, []byte("QrodyHDh1mphsrl23fmpl1dcHlCPY619mbvedTZdxQKyXAH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) QrodyHDh1mphsrl23fmpl1dcHlCPY619mbvedTZdxQKyXAH

ret = client.Close(fd295)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd290, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "esWssaLzQ") {
  panic("wrong data returned")
}


ret = client.Seek(fd298, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd301 := client.Open("/rUTcfOzHpT.txt", client.O_RDWR|client.O_CREATE)
if(fd301 < 0) {
  panic("open failed")
}

//fd state: (22) CawTj8gqDjcxZqhqv8GMPFXN8uMyZd1vlJOCjaA8LNlI95gnBbOSxDa7

ret = client.Write(fd296, []byte("ytpTws7hzFPBbpYgNLSWHywS_wh1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) CawTj8gqDjcxZqhqv8GMPFytpTws7hzFPBbpYgNLSWHywS_wh1OSxDa7
//fd state: (50) CawTj8gqDjcxZqhqv8GMPFytpTws7hzFPBbpYgNLSWHywS_wh1OSxDa7

ret = client.Write(fd296, []byte("TBfWYkmTu57op_2iFoQaZr2PjXJHxBEHqbXvddHMQ58JLJFY9SEYKQd0jBx__Rwr0Dq8o28Cif1bbG4wgvM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) CawTj8gqDjcxZqhqv8GMPFytpTws7hzFPBbpYgNLSWHywS_wh1TBfWYkmTu57op_2iFoQaZr2PjXJHxBEHqbXvddHMQ58JLJFY9SEYKQd0jBx__Rwr0Dq8o28Cif1bbG4wgvM

ret = client.Seek(fd298, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd279, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


buf, ret = client.Read(fd279, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pfhzpWmXCWK0UKsrGjUsc") {
  panic("wrong data returned")
}

//fd state: (167) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeSQ_tPxFcjamQnZ1IY35xtwGAYiPxE81Uk5DFyx7SVCEPIarb9RZyuAOLg0za72lyc75BLdZaJlNcxaFVesWssaLzQ

ret = client.Write(fd290, []byte("6qwEIKYUEVkPDxmC2oYizQWcqhwpMcC6EWwdwFQvVjTpWmudrezZ0BnCM0Cg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (227) z3psbWUSgE6gApnQjQDDiOmjXkBHt2g3AdrJpTcpZnVuuVT_SplDxWqKZW3ZApl8FgPuDTifubfkHeSQ_tPxFcjamQnZ1IY35xtwGAYiPxE81Uk5DFyx7SVCEPIarb9RZyuAOLg0za72lyc75BLdZaJlNcxaFVesWssaLzQ6qwEIKYUEVkPDxmC2oYizQWcqhwpMcC6EWwdwFQvVjTpWmudrezZ0BnCM0Cg

ret = client.Seek(fd276, 260, client.SEEK_SET)
if(ret != 260) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 260)
  panic("seek failed")
}


ret = client.Seek(fd296, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}

//fd state: (12) NVr_siCEPrkO3QuyKKM7zVBVqBtqFEIfekTxXv4_rko3SElG76UK9mdXwJgSAdj7Jq4QLQNYj0CxM6LYfMl8f51ufGcf3NbRj_tCQOy8T86inF9_D2EIbBPFASVAfpiTOcjHu9OmylM7k86MH3hFWzZMbpb_A

ret = client.Write(fd288, []byte("_EhJ838UvSiByi99gDpKZyE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) NVr_siCEPrkO_EhJ838UvSiByi99gDpKZyExXv4_rko3SElG76UK9mdXwJgSAdj7Jq4QLQNYj0CxM6LYfMl8f51ufGcf3NbRj_tCQOy8T86inF9_D2EIbBPFASVAfpiTOcjHu9OmylM7k86MH3hFWzZMbpb_A

buf, ret = client.Read(fd296, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "HqbXvddH") {
  panic("wrong data returned")
}

//fd state: (35) NVr_siCEPrkO_EhJ838UvSiByi99gDpKZyExXv4_rko3SElG76UK9mdXwJgSAdj7Jq4QLQNYj0CxM6LYfMl8f51ufGcf3NbRj_tCQOy8T86inF9_D2EIbBPFASVAfpiTOcjHu9OmylM7k86MH3hFWzZMbpb_A

ret = client.Write(fd288, []byte("OT68LIHgRuUjAko"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) NVr_siCEPrkO_EhJ838UvSiByi99gDpKZyEOT68LIHgRuUjAkoUK9mdXwJgSAdj7Jq4QLQNYj0CxM6LYfMl8f51ufGcf3NbRj_tCQOy8T86inF9_D2EIbBPFASVAfpiTOcjHu9OmylM7k86MH3hFWzZMbpb_A
//fd state: (0) 

ret = client.Write(fd301, []byte("f2LJiTngqIEYQ9NfnuxpJbJq06OFvEwvK38P4opFmZtk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) f2LJiTngqIEYQ9NfnuxpJbJq06OFvEwvK38P4opFmZtk

buf, ret = client.Read(fd296, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MQ58JLJFY9SEYKQd0jBx__Rwr0Dq8o28Ci") {
  panic("wrong data returned")
}


ret = client.Close(fd288)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd279)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd296, 86, client.SEEK_SET)
if(ret != 86) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 86)
  panic("seek failed")
}


fd302 := client.Open("/NDaLK8HBXg.txt", client.O_RDWR|client.O_CREATE)
if(fd302 < 0) {
  panic("open failed")
}


fd303 := client.Open("/Q7Imfuqduy.txt", client.O_RDWR|client.O_CREATE)
if(fd303 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd290, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd303)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd300, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Close(fd296)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd276, 131, client.SEEK_SET)
if(ret != 131) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 131)
  panic("seek failed")
}

//fd state: (131) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtBELXt9E_seromcoQvMvuDBJrPccapET4qyHWCxdsHG7N1GWaolxXhP0m2YVfYxCfXkrhbI4LV39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

ret = client.Write(fd276, []byte("aEF_AnHAgCEFoLilNvkT6viy80tS62PRftK204eZuhq_ZSsYwSo9MG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (185) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtBELXt9E_seromcoQvMvaEF_AnHAgCEFoLilNvkT6viy80tS62PRftK204eZuhq_ZSsYwSo9MG39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

fd304 := client.Open("/FUX87rui7W.txt", client.O_RDWR|client.O_CREATE)
if(fd304 < 0) {
  panic("open failed")
}


ret = client.Seek(fd290, 208, client.SEEK_SET)
if(ret != 208) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 208)
  panic("seek failed")
}


fd305 := client.Open("/_fKtkmMG2H.txt", client.O_RDWR|client.O_CREATE)
if(fd305 < 0) {
  panic("open failed")
}


fd306 := client.Open("/9mO8qrylzT.txt", client.O_RDWR|client.O_CREATE)
if(fd306 < 0) {
  panic("open failed")
}


ret = client.Seek(fd305, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd298, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd305)
if(ret != 0) {
  panic("close failed")
}


fd307 := client.Open("/L8Y7EMLJqt.txt", client.O_RDWR|client.O_CREATE)
if(fd307 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd276, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGU") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd307, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd290, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jTpWmudrezZ0BnCM0Cg") {
  panic("wrong data returned")
}


ret = client.Close(fd302)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd298, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd306, []byte("JPH4gDm2hIlj9XhkaQsIs5_oZYO3U301v_K8hycgyLawcybaFbYCiwrzVXU3gJKNHe1OoIPYd0lGclOEUbOHUtqA_QzznP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) JPH4gDm2hIlj9XhkaQsIs5_oZYO3U301v_K8hycgyLawcybaFbYCiwrzVXU3gJKNHe1OoIPYd0lGclOEUbOHUtqA_QzznP

ret = client.Seek(fd300, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd300)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd307, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (44) f2LJiTngqIEYQ9NfnuxpJbJq06OFvEwvK38P4opFmZtk

ret = client.Write(fd301, []byte("MODkzG8vVqFN9hhGnVF4C_TacPmaeeBnXxW3i6nIRxqlKds1Z8jpl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) f2LJiTngqIEYQ9NfnuxpJbJq06OFvEwvK38P4opFmZtkMODkzG8vVqFN9hhGnVF4C_TacPmaeeBnXxW3i6nIRxqlKds1Z8jpl

ret = client.Close(fd306)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd304)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd290, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd308 := client.Open("/Me76BT6BSP.txt", client.O_RDWR|client.O_CREATE)
if(fd308 < 0) {
  panic("open failed")
}


ret = client.Seek(fd292, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


ret = client.Seek(fd290, 116, client.SEEK_SET)
if(ret != 116) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 116)
  panic("seek failed")
}


ret = client.Close(fd298)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd290, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7SVCEPIarb9RZyuA") {
  panic("wrong data returned")
}


ret = client.Seek(fd292, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd308, []byte("ruBwP_fiMy6ZbbJTRoXjTPlw_hiftv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) ruBwP_fiMy6ZbbJTRoXjTPlw_hiftv

ret = client.Seek(fd276, 210, client.SEEK_SET)
if(ret != 210) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 210)
  panic("seek failed")
}


fd309 := client.Open("/DbkC5q9e6v.txt", client.O_RDWR|client.O_CREATE)
if(fd309 < 0) {
  panic("open failed")
}


ret = client.Close(fd308)
if(ret != 0) {
  panic("close failed")
}


fd310 := client.Open("/4VXKthGmF8.txt", client.O_RDWR|client.O_CREATE)
if(fd310 < 0) {
  panic("open failed")
}

//fd state: (30) B6xmBErBDPVgp5_V2nnQwuACisEK4hIsX3bkithlS9LMqs49P8F6ZeY_s78yY2InTBU8_mG7fySwLsVYAHfgUQ0RBx299sghdSy

ret = client.Write(fd292, []byte("DkgjS5AtsXCLrYaYpOsB8bsX9FAb1aVfYFPBDCSp6T2o5cCvDW5ZsliQBViQW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) B6xmBErBDPVgp5_V2nnQwuACisEK4hDkgjS5AtsXCLrYaYpOsB8bsX9FAb1aVfYFPBDCSp6T2o5cCvDW5ZsliQBViQW99sghdSy

fd311 := client.Open("/Ex7w4vRt6O.txt", client.O_RDWR|client.O_CREATE)
if(fd311 < 0) {
  panic("open failed")
}


ret = client.Seek(fd301, 77, client.SEEK_SET)
if(ret != 77) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 77)
  panic("seek failed")
}


fd312 := client.Open("/jSFBpvzIzk.txt", client.O_RDWR|client.O_CREATE)
if(fd312 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd312, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2xgHENJtrX5tdaNhSUCZKDD2ySzdBjolxE3ipMsXC15W2Fmdu99zRd") {
  panic("wrong data returned")
}


fd313 := client.Open("/iJRMsrR9zU.txt", client.O_RDWR|client.O_CREATE)
if(fd313 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd313, []byte("KXJVV2kwoBVKQVKtAjo0sl6MOVebD1mQmuW8SOd1ZDlHIi0t3DPWIzzJHyeWUYpHg2L44X71THugPCqXLI1A6fhn0adPr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) KXJVV2kwoBVKQVKtAjo0sl6MOVebD1mQmuW8SOd1ZDlHIi0t3DPWIzzJHyeWUYpHg2L44X71THugPCqXLI1A6fhn0adPr

fd314 := client.Open("/ZaBfLNx_Cb.txt", client.O_RDWR|client.O_CREATE)
if(fd314 < 0) {
  panic("open failed")
}


ret = client.Close(fd313)
if(ret != 0) {
  panic("close failed")
}


fd315 := client.Open("/HCfdrAN7le.txt", client.O_RDWR|client.O_CREATE)
if(fd315 < 0) {
  panic("open failed")
}


ret = client.Close(fd290)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd312, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


buf, ret = client.Read(fd314, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "B6xmBErBDPVgp5_V2nnQwuACisEK4hDkgjS5AtsXCLrYaYpOsB") {
  panic("wrong data returned")
}


ret = client.Seek(fd310, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd310, []byte("aIhNzWLry838BnqvtgkOG_6ljpTr05fkEAHLldx6jjhCUrrqAWvR87muu2ijVS5t7FEtMI81pidAd71"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) aIhNzWLry838BnqvtgkOG_6ljpTr05fkEAHLldx6jjhCUrrqAWvR87muu2ijVS5t7FEtMI81pidAd71

buf, ret = client.Read(fd292, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "99sghdSy") {
  panic("wrong data returned")
}


ret = client.Seek(fd311, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}


fd316 := client.Open("/gwR_qKf5l_.txt", client.O_RDWR|client.O_CREATE)
if(fd316 < 0) {
  panic("open failed")
}


ret = client.Close(fd315)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd314, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


buf, ret = client.Read(fd314, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nnQwuACisEK") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd312, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NhSUCZKDD2ySzdBjolxE3ipMsXC15W2Fmdu99zRd") {
  panic("wrong data returned")
}

//fd state: (210) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtBELXt9E_seromcoQvMvaEF_AnHAgCEFoLilNvkT6viy80tS62PRftK204eZuhq_ZSsYwSo9MG39RKU8cUpvDGSm6U_MD_9qb3mC4_sDQEXAlj8l1ydQOMUZUyvqFGW7LvkX0W3EasIjg2yGUm6fV6uuhM4KyDsFP4lawaUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

ret = client.Write(fd276, []byte("0h9smremudSAocezKiFCz2jJDIHdYp2jiIpGjbk7FxiGzCOvtlC9BaiEBghp6EQ02b"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (276) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtBELXt9E_seromcoQvMvaEF_AnHAgCEFoLilNvkT6viy80tS62PRftK204eZuhq_ZSsYwSo9MG39RKU8cUpvDGSm6U_MD_9qb3m0h9smremudSAocezKiFCz2jJDIHdYp2jiIpGjbk7FxiGzCOvtlC9BaiEBghp6EQ02baUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr
//fd state: (79) aIhNzWLry838BnqvtgkOG_6ljpTr05fkEAHLldx6jjhCUrrqAWvR87muu2ijVS5t7FEtMI81pidAd71

ret = client.Write(fd310, []byte("mxtoW15Oe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) aIhNzWLry838BnqvtgkOG_6ljpTr05fkEAHLldx6jjhCUrrqAWvR87muu2ijVS5t7FEtMI81pidAd71mxtoW15Oe

ret = client.Close(fd312)
if(ret != 0) {
  panic("close failed")
}


fd317 := client.Open("/NDaLK8HBXg.txt", client.O_RDWR|client.O_CREATE)
if(fd317 < 0) {
  panic("open failed")
}


fd318 := client.Open("/_laoVJpAG9.txt", client.O_RDWR|client.O_CREATE)
if(fd318 < 0) {
  panic("open failed")
}

//fd state: (99) B6xmBErBDPVgp5_V2nnQwuACisEK4hDkgjS5AtsXCLrYaYpOsB8bsX9FAb1aVfYFPBDCSp6T2o5cCvDW5ZsliQBViQW99sghdSy

ret = client.Write(fd292, []byte("WeYpt3Lx2mkFjn70OnQJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (119) B6xmBErBDPVgp5_V2nnQwuACisEK4hDkgjS5AtsXCLrYaYpOsB8bsX9FAb1aVfYFPBDCSp6T2o5cCvDW5ZsliQBViQW99sghdSyWeYpt3Lx2mkFjn70OnQJ

buf, ret = client.Read(fd318, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "A9awL") {
  panic("wrong data returned")
}

//fd state: (88) aIhNzWLry838BnqvtgkOG_6ljpTr05fkEAHLldx6jjhCUrrqAWvR87muu2ijVS5t7FEtMI81pidAd71mxtoW15Oe

ret = client.Write(fd310, []byte("PvfZNJi_128Rtm1_lFt7TOXb_XKc7aPGXQudcD39UA2VcuDhts4IHwIslRiNgUDBps4JMBuB3w53qQRwbBO7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (172) aIhNzWLry838BnqvtgkOG_6ljpTr05fkEAHLldx6jjhCUrrqAWvR87muu2ijVS5t7FEtMI81pidAd71mxtoW15OePvfZNJi_128Rtm1_lFt7TOXb_XKc7aPGXQudcD39UA2VcuDhts4IHwIslRiNgUDBps4JMBuB3w53qQRwbBO7
//fd state: (0) SQNlLo8g3yfsNMWPbW0aMO9B1PlCHL66rpIHtPtfSdAdtCkofreovbieRSTsMX_xa1kXEkpXByQUVjD8sJtj5bwr56z91wBifhIgBaxg2kdSMyv3uroAZt

ret = client.Write(fd316, []byte("A5kGk2STUJf_URwJTffQNrd9V_e5piXpqvEqayJaF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) A5kGk2STUJf_URwJTffQNrd9V_e5piXpqvEqayJaFdAdtCkofreovbieRSTsMX_xa1kXEkpXByQUVjD8sJtj5bwr56z91wBifhIgBaxg2kdSMyv3uroAZt

fd319 := client.Open("/RUpXDwp6BU.txt", client.O_RDWR|client.O_CREATE)
if(fd319 < 0) {
  panic("open failed")
}


ret = client.Seek(fd316, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}

//fd state: (5) A9awL

ret = client.Write(fd318, []byte("Y38L"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) A9awLY38L
//fd state: (49) GqLX5nqhbOhPpdK21wGcDfQ9PcSxjXbCATzEVIjON9F2UPG8yqQyIF21gt2RUsKuxox8GIan73Q7alUgZuA

ret = client.Write(fd311, []byte("P8XglBZEFePhzyobJCR2TYJjoqWLOz1x_MO9nJl8CDv7j"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) GqLX5nqhbOhPpdK21wGcDfQ9PcSxjXbCATzEVIjON9F2UPG8yP8XglBZEFePhzyobJCR2TYJjoqWLOz1x_MO9nJl8CDv7j
//fd state: (94) GqLX5nqhbOhPpdK21wGcDfQ9PcSxjXbCATzEVIjON9F2UPG8yP8XglBZEFePhzyobJCR2TYJjoqWLOz1x_MO9nJl8CDv7j

ret = client.Write(fd311, []byte("21T"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) GqLX5nqhbOhPpdK21wGcDfQ9PcSxjXbCATzEVIjON9F2UPG8yP8XglBZEFePhzyobJCR2TYJjoqWLOz1x_MO9nJl8CDv7j21T

ret = client.Seek(fd314, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


buf, ret = client.Read(fd276, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "aUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr") {
  panic("wrong data returned")
}


fd320 := client.Open("/OwtKj7ePV0.txt", client.O_RDWR|client.O_CREATE)
if(fd320 < 0) {
  panic("open failed")
}

//fd state: (7) B6xmBErBDPVgp5_V2nnQwuACisEK4hDkgjS5AtsXCLrYaYpOsB8bsX9FAb1aVfYFPBDCSp6T2o5cCvDW5ZsliQBViQW99sghdSyWeYpt3Lx2mkFjn70OnQJ

ret = client.Write(fd314, []byte("UnLZmffy0YRwRWo2cDZXfCOV_6Th0YfSX1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) B6xmBErUnLZmffy0YRwRWo2cDZXfCOV_6Th0YfSX1LrYaYpOsB8bsX9FAb1aVfYFPBDCSp6T2o5cCvDW5ZsliQBViQW99sghdSyWeYpt3Lx2mkFjn70OnQJ

ret = client.Seek(fd314, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd317)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd301)
if(ret != 0) {
  panic("close failed")
}

//fd state: (9) A9awLY38L

ret = client.Write(fd318, []byte("iPZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) A9awLY38LiPZ
//fd state: (172) aIhNzWLry838BnqvtgkOG_6ljpTr05fkEAHLldx6jjhCUrrqAWvR87muu2ijVS5t7FEtMI81pidAd71mxtoW15OePvfZNJi_128Rtm1_lFt7TOXb_XKc7aPGXQudcD39UA2VcuDhts4IHwIslRiNgUDBps4JMBuB3w53qQRwbBO7

ret = client.Write(fd310, []byte("IZCOMCD4En59MDvY1QcZIrhp1_LZvgurKVExDC_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (211) aIhNzWLry838BnqvtgkOG_6ljpTr05fkEAHLldx6jjhCUrrqAWvR87muu2ijVS5t7FEtMI81pidAd71mxtoW15OePvfZNJi_128Rtm1_lFt7TOXb_XKc7aPGXQudcD39UA2VcuDhts4IHwIslRiNgUDBps4JMBuB3w53qQRwbBO7IZCOMCD4En59MDvY1QcZIrhp1_LZvgurKVExDC_

ret = client.Seek(fd319, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


buf, ret = client.Read(fd309, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "y76GTNImWOrQbI8Ujy6WBFnHgX8XNp4FzL97YJ01GhBFDVxJHXkpME6zlKmwrxriiCa1uV4VwFtOn5a") {
  panic("wrong data returned")
}


fd321 := client.Open("/Mnbfjg6rVs.txt", client.O_RDWR|client.O_CREATE)
if(fd321 < 0) {
  panic("open failed")
}


ret = client.Close(fd318)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd276, 113, client.SEEK_SET)
if(ret != 113) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 113)
  panic("seek failed")
}


ret = client.Close(fd310)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd309, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}


ret = client.Close(fd319)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd320)
if(ret != 0) {
  panic("close failed")
}


fd322 := client.Open("/4cQqDopMxJ.txt", client.O_RDWR|client.O_CREATE)
if(fd322 < 0) {
  panic("open failed")
}


ret = client.Close(fd311)
if(ret != 0) {
  panic("close failed")
}


fd323 := client.Open("/UTBcPmKgP5.txt", client.O_RDWR|client.O_CREATE)
if(fd323 < 0) {
  panic("open failed")
}

//fd state: (113) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtBELXt9E_seromcoQvMvaEF_AnHAgCEFoLilNvkT6viy80tS62PRftK204eZuhq_ZSsYwSo9MG39RKU8cUpvDGSm6U_MD_9qb3m0h9smremudSAocezKiFCz2jJDIHdYp2jiIpGjbk7FxiGzCOvtlC9BaiEBghp6EQ02baUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

ret = client.Write(fd276, []byte("0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtB0LXt9E_seromcoQvMvaEF_AnHAgCEFoLilNvkT6viy80tS62PRftK204eZuhq_ZSsYwSo9MG39RKU8cUpvDGSm6U_MD_9qb3m0h9smremudSAocezKiFCz2jJDIHdYp2jiIpGjbk7FxiGzCOvtlC9BaiEBghp6EQ02baUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

buf, ret = client.Read(fd322, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "KCPdD3hBtZztTxnEDBH65kYLRELkwBL1O7324t311qNX3S0PTcNIVoD2JZdB") {
  panic("wrong data returned")
}


fd324 := client.Open("/jlETz8ELRp.txt", client.O_RDWR|client.O_CREATE)
if(fd324 < 0) {
  panic("open failed")
}


ret = client.Seek(fd307, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd321, []byte("lEy4_2pn5jvI09FwUKCglK19TIsyer6ElWdH1rnXx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) lEy4_2pn5jvI09FwUKCglK19TIsyer6ElWdH1rnXx
//fd state: (114) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtB0LXt9E_seromcoQvMvaEF_AnHAgCEFoLilNvkT6viy80tS62PRftK204eZuhq_ZSsYwSo9MG39RKU8cUpvDGSm6U_MD_9qb3m0h9smremudSAocezKiFCz2jJDIHdYp2jiIpGjbk7FxiGzCOvtlC9BaiEBghp6EQ02baUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

ret = client.Write(fd276, []byte("6hfQK7HcLUlZNZvTI3zYG308vP15icG9f6C8mLcNiyk3fpXL5RuoosGbkWyI8MzZ8gUkU5mEjoD3flk466Qx17WJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (202) m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSSH5ca5hq_8gegtB06hfQK7HcLUlZNZvTI3zYG308vP15icG9f6C8mLcNiyk3fpXL5RuoosGbkWyI8MzZ8gUkU5mEjoD3flk466Qx17WJMD_9qb3m0h9smremudSAocezKiFCz2jJDIHdYp2jiIpGjbk7FxiGzCOvtlC9BaiEBghp6EQ02baUhuZmMpiTSPNu1NHk1iJcnAkXAk_Pc3reecg3y_VSg_GQJlSNuwtLPQAr

ret = client.Close(fd307)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd276, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MD_9qb3m0h9smremud") {
  panic("wrong data returned")
}

//fd state: (49) y76GTNImWOrQbI8Ujy6WBFnHgX8XNp4FzL97YJ01GhBFDVxJHXkpME6zlKmwrxriiCa1uV4VwFtOn5a

ret = client.Write(fd309, []byte("1n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) y76GTNImWOrQbI8Ujy6WBFnHgX8XNp4FzL97YJ01GhBFDVxJH1npME6zlKmwrxriiCa1uV4VwFtOn5a

ret = client.Close(fd323)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd276)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd309, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pME6zlKmwrxriiCa1uV4VwFtOn5a") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd324, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd325 := client.Open("/Vd2TJjK6jh.txt", client.O_RDWR|client.O_CREATE)
if(fd325 < 0) {
  panic("open failed")
}

//fd state: (35) A5kGk2STUJf_URwJTffQNrd9V_e5piXpqvEqayJaFdAdtCkofreovbieRSTsMX_xa1kXEkpXByQUVjD8sJtj5bwr56z91wBifhIgBaxg2kdSMyv3uroAZt

ret = client.Write(fd316, []byte("90g3MKoeHXb_fNXyuApx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) A5kGk2STUJf_URwJTffQNrd9V_e5piXpqvE90g3MKoeHXb_fNXyuApxeRSTsMX_xa1kXEkpXByQUVjD8sJtj5bwr56z91wBifhIgBaxg2kdSMyv3uroAZt
//fd state: (12) B6xmBErUnLZmffy0YRwRWo2cDZXfCOV_6Th0YfSX1LrYaYpOsB8bsX9FAb1aVfYFPBDCSp6T2o5cCvDW5ZsliQBViQW99sghdSyWeYpt3Lx2mkFjn70OnQJ

ret = client.Write(fd314, []byte("QFHK85YejVWIg8zm1uGKI3VZ1d7GF5jgzHy2GUyh8Kz1zhTALEaJlDB36S8Qf9vPC9brnrE7XuCNPqPNKfOexMbvMl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) B6xmBErUnLZmQFHK85YejVWIg8zm1uGKI3VZ1d7GF5jgzHy2GUyh8Kz1zhTALEaJlDB36S8Qf9vPC9brnrE7XuCNPqPNKfOexMbvMlpt3Lx2mkFjn70OnQJ

buf, ret = client.Read(fd321, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd326 := client.Open("/WfcD25sP7W.txt", client.O_RDWR|client.O_CREATE)
if(fd326 < 0) {
  panic("open failed")
}


ret = client.Close(fd314)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd292)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd324, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd326, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd326)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd321)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd324)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd325, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd316, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}

//fd state: (63) A5kGk2STUJf_URwJTffQNrd9V_e5piXpqvE90g3MKoeHXb_fNXyuApxeRSTsMX_xa1kXEkpXByQUVjD8sJtj5bwr56z91wBifhIgBaxg2kdSMyv3uroAZt

ret = client.Write(fd316, []byte("82pl9R4qH8_MEjGY9sXsQuxGqNQXFqgVtEqI2_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) A5kGk2STUJf_URwJTffQNrd9V_e5piXpqvE90g3MKoeHXb_fNXyuApxeRSTsMX_82pl9R4qH8_MEjGY9sXsQuxGqNQXFqgVtEqI2_axg2kdSMyv3uroAZt

fd327 := client.Open("/OlqXijaKn9.txt", client.O_RDWR|client.O_CREATE)
if(fd327 < 0) {
  panic("open failed")
}


fd328 := client.Open("/4SS3LdR3oc.txt", client.O_RDWR|client.O_CREATE)
if(fd328 < 0) {
  panic("open failed")
}


ret = client.Close(fd309)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd327, []byte("pubDyMsr09ZndTxVA8xsnfEGCxkpmXXfxJABiwBNE2XXnrkGzZ39TSE4IqvR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) pubDyMsr09ZndTxVA8xsnfEGCxkpmXXfxJABiwBNE2XXnrkGzZ39TSE4IqvR

fd329 := client.Open("/wA5Kf_Ltjb.txt", client.O_RDWR|client.O_CREATE)
if(fd329 < 0) {
  panic("open failed")
}

//fd state: (60) pubDyMsr09ZndTxVA8xsnfEGCxkpmXXfxJABiwBNE2XXnrkGzZ39TSE4IqvR

ret = client.Write(fd327, []byte("Up"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) pubDyMsr09ZndTxVA8xsnfEGCxkpmXXfxJABiwBNE2XXnrkGzZ39TSE4IqvRUp

buf, ret = client.Read(fd329, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "19Q") {
  panic("wrong data returned")
}


ret = client.Close(fd328)
if(ret != 0) {
  panic("close failed")
}

//fd state: (101) A5kGk2STUJf_URwJTffQNrd9V_e5piXpqvE90g3MKoeHXb_fNXyuApxeRSTsMX_82pl9R4qH8_MEjGY9sXsQuxGqNQXFqgVtEqI2_axg2kdSMyv3uroAZt

ret = client.Write(fd316, []byte("LJ1ShgsFm86FZWQOpWZRGBbzuH8r2IXjYg9TGgWTMNKoD63fw6aBtDAfR7HBlND7VJb23rYMWpNQhnUsAX0zLVW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (188) A5kGk2STUJf_URwJTffQNrd9V_e5piXpqvE90g3MKoeHXb_fNXyuApxeRSTsMX_82pl9R4qH8_MEjGY9sXsQuxGqNQXFqgVtEqI2_LJ1ShgsFm86FZWQOpWZRGBbzuH8r2IXjYg9TGgWTMNKoD63fw6aBtDAfR7HBlND7VJb23rYMWpNQhnUsAX0zLVW

ret = client.Seek(fd329, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


buf, ret = client.Read(fd327, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (62) pubDyMsr09ZndTxVA8xsnfEGCxkpmXXfxJABiwBNE2XXnrkGzZ39TSE4IqvRUp

ret = client.Write(fd327, []byte("NOyeurJI8yKDWRTWSbApIQ5GUm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) pubDyMsr09ZndTxVA8xsnfEGCxkpmXXfxJABiwBNE2XXnrkGzZ39TSE4IqvRUpNOyeurJI8yKDWRTWSbApIQ5GUm

buf, ret = client.Read(fd327, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (60) KCPdD3hBtZztTxnEDBH65kYLRELkwBL1O7324t311qNX3S0PTcNIVoD2JZdBYbWOuHIthsQV3RDO2rCXel79elj46m5IkvBkpUh2zOzCuVI1gv1nc3QiQyfHYMbSmlKnRiOLQxf

ret = client.Write(fd322, []byte("c4nPBhCBnsDO_4k4yE5ICAF51jQcdjlerE_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) KCPdD3hBtZztTxnEDBH65kYLRELkwBL1O7324t311qNX3S0PTcNIVoD2JZdBc4nPBhCBnsDO_4k4yE5ICAF51jQcdjlerE_kpUh2zOzCuVI1gv1nc3QiQyfHYMbSmlKnRiOLQxf

ret = client.Close(fd316)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd322, 84, client.SEEK_SET)
if(ret != 84) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 84)
  panic("seek failed")
}

//fd state: (88) pubDyMsr09ZndTxVA8xsnfEGCxkpmXXfxJABiwBNE2XXnrkGzZ39TSE4IqvRUpNOyeurJI8yKDWRTWSbApIQ5GUm

ret = client.Write(fd327, []byte("oR_3XvhL0oYQwPBbI1laS03gmQ2uNmTk28poTQ5dYnyd8cUPgzJe05P_3qUTMLyYJK_66rj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (159) pubDyMsr09ZndTxVA8xsnfEGCxkpmXXfxJABiwBNE2XXnrkGzZ39TSE4IqvRUpNOyeurJI8yKDWRTWSbApIQ5GUmoR_3XvhL0oYQwPBbI1laS03gmQ2uNmTk28poTQ5dYnyd8cUPgzJe05P_3qUTMLyYJK_66rj
//fd state: (0) 

ret = client.Write(fd325, []byte("vQhkLoMSUdBt9wZzV_ZCQTKFbaWlZQJJFpCr5E4S3Tbpt4E"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) vQhkLoMSUdBt9wZzV_ZCQTKFbaWlZQJJFpCr5E4S3Tbpt4E

ret = client.Close(fd325)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd327)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd322, 105, client.SEEK_SET)
if(ret != 105) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 105)
  panic("seek failed")
}

//fd state: (105) KCPdD3hBtZztTxnEDBH65kYLRELkwBL1O7324t311qNX3S0PTcNIVoD2JZdBc4nPBhCBnsDO_4k4yE5ICAF51jQcdjlerE_kpUh2zOzCuVI1gv1nc3QiQyfHYMbSmlKnRiOLQxf

ret = client.Write(fd322, []byte("Lz7tcr0zvFTP1B6GYr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) KCPdD3hBtZztTxnEDBH65kYLRELkwBL1O7324t311qNX3S0PTcNIVoD2JZdBc4nPBhCBnsDO_4k4yE5ICAF51jQcdjlerE_kpUh2zOzCuLz7tcr0zvFTP1B6GYrSmlKnRiOLQxf
//fd state: (3) 19Q

ret = client.Write(fd329, []byte("4lTJaMmtgjQb60lCQhx9R2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) 19Q4lTJaMmtgjQb60lCQhx9R2

fd330 := client.Open("/oy_syYGpgq.txt", client.O_RDWR|client.O_CREATE)
if(fd330 < 0) {
  panic("open failed")
}


fd331 := client.Open("/jSFBpvzIzk.txt", client.O_RDWR|client.O_CREATE)
if(fd331 < 0) {
  panic("open failed")
}


fd332 := client.Open("/S4kyEU_r6C.txt", client.O_RDWR|client.O_CREATE)
if(fd332 < 0) {
  panic("open failed")
}


ret = client.Close(fd322)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd330)
if(ret != 0) {
  panic("close failed")
}


fd333 := client.Open("/LoIyZd2AM5.txt", client.O_RDWR|client.O_CREATE)
if(fd333 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd332, []byte("pDYhk49mN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) pDYhk49mN
//fd state: (25) 19Q4lTJaMmtgjQb60lCQhx9R2

ret = client.Write(fd329, []byte("jmQXifKz5zwHsyxYBy4eszHZNe_h7vM0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) 19Q4lTJaMmtgjQb60lCQhx9R2jmQXifKz5zwHsyxYBy4eszHZNe_h7vM0

ret = client.Close(fd332)
if(ret != 0) {
  panic("close failed")
}


fd334 := client.Open("/HRpPyOlU84.txt", client.O_RDWR|client.O_CREATE)
if(fd334 < 0) {
  panic("open failed")
}

//fd state: (0) 2xgHENJtrX5tdaNhSUCZKDD2ySzdBjolxE3ipMsXC15W2Fmdu99zRd

ret = client.Write(fd331, []byte("GIp4NHMo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) GIp4NHMorX5tdaNhSUCZKDD2ySzdBjolxE3ipMsXC15W2Fmdu99zRd

ret = client.Seek(fd334, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


buf, ret = client.Read(fd333, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (57) 19Q4lTJaMmtgjQb60lCQhx9R2jmQXifKz5zwHsyxYBy4eszHZNe_h7vM0

ret = client.Write(fd329, []byte("8WaJpoYxTRzenfurIqfv4y4yw1hXKg5UTwo5GBy6sc3J6brvj0wdJmjj2F6tEDP8MjYURAFCD9W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) 19Q4lTJaMmtgjQb60lCQhx9R2jmQXifKz5zwHsyxYBy4eszHZNe_h7vM08WaJpoYxTRzenfurIqfv4y4yw1hXKg5UTwo5GBy6sc3J6brvj0wdJmjj2F6tEDP8MjYURAFCD9W

ret = client.Seek(fd331, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


ret = client.Seek(fd329, 117, client.SEEK_SET)
if(ret != 117) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 117)
  panic("seek failed")
}


fd335 := client.Open("/pqf_xZqbHE.txt", client.O_RDWR|client.O_CREATE)
if(fd335 < 0) {
  panic("open failed")
}


ret = client.Close(fd334)
if(ret != 0) {
  panic("close failed")
}


fd336 := client.Open("/NDaLK8HBXg.txt", client.O_RDWR|client.O_CREATE)
if(fd336 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd333, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd335, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd329, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Seek(fd331, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


fd337 := client.Open("/DTQl1UDSls.txt", client.O_RDWR|client.O_CREATE)
if(fd337 < 0) {
  panic("open failed")
}


ret = client.Close(fd336)
if(ret != 0) {
  panic("close failed")
}

//fd state: (40) GIp4NHMorX5tdaNhSUCZKDD2ySzdBjolxE3ipMsXC15W2Fmdu99zRd

ret = client.Write(fd331, []byte("nyTFWYZ5ZEzj4NrXIIq51TZJKbgr2FIrE_GbveLcZ8OwIF1SMxMtoM8sWO5ySxhpZGqIxRCEeYAOzDzeVQQfJv6mZjOG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) GIp4NHMorX5tdaNhSUCZKDD2ySzdBjolxE3ipMsXnyTFWYZ5ZEzj4NrXIIq51TZJKbgr2FIrE_GbveLcZ8OwIF1SMxMtoM8sWO5ySxhpZGqIxRCEeYAOzDzeVQQfJv6mZjOG

buf, ret = client.Read(fd335, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd338 := client.Open("/hIqrxvPj2h.txt", client.O_RDWR|client.O_CREATE)
if(fd338 < 0) {
  panic("open failed")
}


ret = client.Seek(fd333, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd337, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd337, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd333, []byte("i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) i

fd339 := client.Open("/R6UEMyClwH.txt", client.O_RDWR|client.O_CREATE)
if(fd339 < 0) {
  panic("open failed")
}


ret = client.Close(fd335)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd329)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd333, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


fd340 := client.Open("/NOndoHIH7p.txt", client.O_RDWR|client.O_CREATE)
if(fd340 < 0) {
  panic("open failed")
}


ret = client.Close(fd333)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd339)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd337, []byte("XdIs_MpK5Ug5tWttRcLtzVHMtzaqAjlbDH5d88r5KUCUUPq5wYk4NywW98IHPA3GWUalKGnkye418M"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) XdIs_MpK5Ug5tWttRcLtzVHMtzaqAjlbDH5d88r5KUCUUPq5wYk4NywW98IHPA3GWUalKGnkye418M

ret = client.Seek(fd331, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


ret = client.Close(fd340)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd337, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd331)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd338)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd337)
if(ret != 0) {
  panic("close failed")
}


fd341 := client.Open("/Mnbfjg6rVs.txt", client.O_RDWR|client.O_CREATE)
if(fd341 < 0) {
  panic("open failed")
}

//fd state: (0) lEy4_2pn5jvI09FwUKCglK19TIsyer6ElWdH1rnXx

ret = client.Write(fd341, []byte("jV_uhQMVyds8qlDaBycu_TDRlzCKRwv8SnW9QdKvSbTM2WhztupChwQ4S1vo4aOCUHrfyUweb6C6nnOhAe4JRgvKUfK6c1l2mS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) jV_uhQMVyds8qlDaBycu_TDRlzCKRwv8SnW9QdKvSbTM2WhztupChwQ4S1vo4aOCUHrfyUweb6C6nnOhAe4JRgvKUfK6c1l2mS

ret = client.Close(fd341)
if(ret != 0) {
  panic("close failed")
}


fd342 := client.Open("/C1ZDrk7A58.txt", client.O_RDWR|client.O_CREATE)
if(fd342 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd342, []byte("4P7C6iVFRW6_6kpqDwS3Qzdd2UuCinShZI6ZSycsVLovvXhdAjrd8ULbO71d1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) 4P7C6iVFRW6_6kpqDwS3Qzdd2UuCinShZI6ZSycsVLovvXhdAjrd8ULbO71d1

ret = client.Close(fd342)
if(ret != 0) {
  panic("close failed")
}


fd343 := client.Open("/Wc8I1sgyqN.txt", client.O_RDWR|client.O_CREATE)
if(fd343 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd343, []byte("Na3vDus5puioiFGF9DuLt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) Na3vDus5puioiFGF9DuLt

ret = client.Seek(fd343, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd344 := client.Open("/0tuonfQNeH.txt", client.O_RDWR|client.O_CREATE)
if(fd344 < 0) {
  panic("open failed")
}


ret = client.Seek(fd344, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd343)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd344, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd345 := client.Open("/7OdNvjR3av.txt", client.O_RDWR|client.O_CREATE)
if(fd345 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd345, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd345, []byte("MkcNBURxNNkLgBY1dctT6kjUqXFF0y0YVJgJ1t15bnstKHkMStM6A2MLafc5d1ogSaA55uV4MRvDoFvLapWhJBacA1DycNL7s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) MkcNBURxNNkLgBY1dctT6kjUqXFF0y0YVJgJ1t15bnstKHkMStM6A2MLafc5d1ogSaA55uV4MRvDoFvLapWhJBacA1DycNL7s

ret = client.Seek(fd344, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd344, []byte("msXnYYlR6K4E"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) msXnYYlR6K4E

fd346 := client.Open("/4mPDBgaYe1.txt", client.O_RDWR|client.O_CREATE)
if(fd346 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd345, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd347 := client.Open("/eXNFP7XOhv.txt", client.O_RDWR|client.O_CREATE)
if(fd347 < 0) {
  panic("open failed")
}


ret = client.Close(fd344)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd347, []byte("hym1SU_h"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) hym1SU_h

fd348 := client.Open("/lK1Ra888uV.txt", client.O_RDWR|client.O_CREATE)
if(fd348 < 0) {
  panic("open failed")
}

//fd state: (8) hym1SU_h

ret = client.Write(fd347, []byte("VIcDh2tq6wF7TxZCpu2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) hym1SU_hVIcDh2tq6wF7TxZCpu2

fd349 := client.Open("/RUBg7QMMiO.txt", client.O_RDWR|client.O_CREATE)
if(fd349 < 0) {
  panic("open failed")
}


ret = client.Seek(fd349, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd350 := client.Open("/_bppGnywyK.txt", client.O_RDWR|client.O_CREATE)
if(fd350 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd346, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "m7kcd4wCgLclEGMf9HMhg5Zwz75Q3KJE8YuNHZf3IejtgIrHkFzMzbsGTipZ7PAyEcQ8ZCxn3DZob1f5CZq5WxOc6KqixBlNwSS") {
  panic("wrong data returned")
}


ret = client.Close(fd348)
if(ret != 0) {
  panic("close failed")
}


fd351 := client.Open("/nHs_I8_1L0.txt", client.O_RDWR|client.O_CREATE)
if(fd351 < 0) {
  panic("open failed")
}


fd352 := client.Open("/nYSzy33lDz.txt", client.O_RDWR|client.O_CREATE)
if(fd352 < 0) {
  panic("open failed")
}


fd353 := client.Open("/qlEFbgz5IC.txt", client.O_RDWR|client.O_CREATE)
if(fd353 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd350, []byte("OoUXvf3V9L6yQ4sznxyOZTZ33t5cSGRp0xGslXc9nv2KQjIpWf3GRYkBQlpjldcW0c8a7b1LMmdXitUYzwTdz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) OoUXvf3V9L6yQ4sznxyOZTZ33t5cSGRp0xGslXc9nv2KQjIpWf3GRYkBQlpjldcW0c8a7b1LMmdXitUYzwTdz

fd354 := client.Open("/QzZYzBiP2c.txt", client.O_RDWR|client.O_CREATE)
if(fd354 < 0) {
  panic("open failed")
}


ret = client.Seek(fd350, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


ret = client.Seek(fd345, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}


buf, ret = client.Read(fd346, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "H5ca5hq_8gegtB06hfQK7HcLUlZNZvTI3zYG308vP15i") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd353, []byte("XzTCExLW1Gw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) XzTCExLW1Gw

fd355 := client.Open("/2JrRL5XTCk.txt", client.O_RDWR|client.O_CREATE)
if(fd355 < 0) {
  panic("open failed")
}


ret = client.Close(fd353)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd349)
if(ret != 0) {
  panic("close failed")
}


fd356 := client.Open("/WujY5LPXRr.txt", client.O_RDWR|client.O_CREATE)
if(fd356 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd346, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "cG9f6C8mLcNiy") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd346, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "k3fpXL5RuoosGbkWyI8MzZ8gUkU5mEjoD3flk466Qx17WJMD_9qb3m0h9smremudSAocezKiFCz2jJDIHdYp") {
  panic("wrong data returned")
}


ret = client.Seek(fd355, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd352, []byte("4tz5E04fbPh7jkJRgLQtAgDrEsu6UYWP1Pyqi0S5pQtiqQrkOBAO57zy_EBCxe1GMociBqXt2Lv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) 4tz5E04fbPh7jkJRgLQtAgDrEsu6UYWP1Pyqi0S5pQtiqQrkOBAO57zy_EBCxe1GMociBqXt2Lv
//fd state: (0) 

ret = client.Write(fd356, []byte("6gN38KhTdbkYxcR_xd0eeOfsZA9lx8byts9HBclQGtShQ0F2uw8e8DxOuTK2SVKAJxWKSfjlv45fH_2KDgCF9h4jLPD8oGRKD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) 6gN38KhTdbkYxcR_xd0eeOfsZA9lx8byts9HBclQGtShQ0F2uw8e8DxOuTK2SVKAJxWKSfjlv45fH_2KDgCF9h4jLPD8oGRKD

fd357 := client.Open("/yw9jdLKnD8.txt", client.O_RDWR|client.O_CREATE)
if(fd357 < 0) {
  panic("open failed")
}


ret = client.Close(fd355)
if(ret != 0) {
  panic("close failed")
}

//fd state: (97) 6gN38KhTdbkYxcR_xd0eeOfsZA9lx8byts9HBclQGtShQ0F2uw8e8DxOuTK2SVKAJxWKSfjlv45fH_2KDgCF9h4jLPD8oGRKD

ret = client.Write(fd356, []byte("kQP7I9xmpqdunqfuyCrtM4596bD1XoGfGrpGWc6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) 6gN38KhTdbkYxcR_xd0eeOfsZA9lx8byts9HBclQGtShQ0F2uw8e8DxOuTK2SVKAJxWKSfjlv45fH_2KDgCF9h4jLPD8oGRKDkQP7I9xmpqdunqfuyCrtM4596bD1XoGfGrpGWc6

ret = client.Close(fd346)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd352, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (75) 4tz5E04fbPh7jkJRgLQtAgDrEsu6UYWP1Pyqi0S5pQtiqQrkOBAO57zy_EBCxe1GMociBqXt2Lv

ret = client.Write(fd352, []byte("rwkETsWb4AzwpBhTMO6EW6oXL6jh4I2mprlTNJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) 4tz5E04fbPh7jkJRgLQtAgDrEsu6UYWP1Pyqi0S5pQtiqQrkOBAO57zy_EBCxe1GMociBqXt2LvrwkETsWb4AzwpBhTMO6EW6oXL6jh4I2mprlTNJ

ret = client.Close(fd345)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd354, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (113) 4tz5E04fbPh7jkJRgLQtAgDrEsu6UYWP1Pyqi0S5pQtiqQrkOBAO57zy_EBCxe1GMociBqXt2LvrwkETsWb4AzwpBhTMO6EW6oXL6jh4I2mprlTNJ

ret = client.Write(fd352, []byte("WurxnHHfXUVvEebXddWOIiNwkqW2E8LqlMlyhlW1ASH5MB6OMGnGrWXFR0cRRVlscsy47jrQPSlBSg3HQoYi0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (198) 4tz5E04fbPh7jkJRgLQtAgDrEsu6UYWP1Pyqi0S5pQtiqQrkOBAO57zy_EBCxe1GMociBqXt2LvrwkETsWb4AzwpBhTMO6EW6oXL6jh4I2mprlTNJWurxnHHfXUVvEebXddWOIiNwkqW2E8LqlMlyhlW1ASH5MB6OMGnGrWXFR0cRRVlscsy47jrQPSlBSg3HQoYi0

buf, ret = client.Read(fd357, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd358 := client.Open("/GBXTDMEc9S.txt", client.O_RDWR|client.O_CREATE)
if(fd358 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd357, []byte("g51"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) g51
//fd state: (0) 

ret = client.Write(fd358, []byte("xA5zJP1heOoOC8GYFnsgQkXQWcHWqH6ZmTMScuckqVeT8GB2CKLDEwf1cZc_jk7pAUQPfqCB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) xA5zJP1heOoOC8GYFnsgQkXQWcHWqH6ZmTMScuckqVeT8GB2CKLDEwf1cZc_jk7pAUQPfqCB

ret = client.Seek(fd347, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


ret = client.Seek(fd356, 66, client.SEEK_SET)
if(ret != 66) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 66)
  panic("seek failed")
}


ret = client.Seek(fd358, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


fd359 := client.Open("/SlueffwKXY.txt", client.O_RDWR|client.O_CREATE)
if(fd359 < 0) {
  panic("open failed")
}


fd360 := client.Open("/5L6AH59F8o.txt", client.O_RDWR|client.O_CREATE)
if(fd360 < 0) {
  panic("open failed")
}


ret = client.Close(fd360)
if(ret != 0) {
  panic("close failed")
}


fd361 := client.Open("/j5i_lS91rD.txt", client.O_RDWR|client.O_CREATE)
if(fd361 < 0) {
  panic("open failed")
}


ret = client.Close(fd359)
if(ret != 0) {
  panic("close failed")
}


fd362 := client.Open("/k4d87VpxAe.txt", client.O_RDWR|client.O_CREATE)
if(fd362 < 0) {
  panic("open failed")
}


ret = client.Seek(fd350, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


buf, ret = client.Read(fd350, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nxyOZTZ33t5cSGRp0xGslXc9nv2KQjIpWf3GRYkB") {
  panic("wrong data returned")
}


ret = client.Seek(fd361, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd357, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd357, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "51") {
  panic("wrong data returned")
}


fd363 := client.Open("/Ejqw2f8dM3.txt", client.O_RDWR|client.O_CREATE)
if(fd363 < 0) {
  panic("open failed")
}


ret = client.Close(fd357)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd363)
if(ret != 0) {
  panic("close failed")
}

//fd state: (24) hym1SU_hVIcDh2tq6wF7TxZCpu2

ret = client.Write(fd347, []byte("E8ygNH820R8K7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) hym1SU_hVIcDh2tq6wF7TxZCE8ygNH820R8K7

buf, ret = client.Read(fd352, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd352, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd356, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WKSfjlv45fH_2KDgC") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd350, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "QlpjldcW0c8a7b1LMmdXitUYzwTdz") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd361, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd364 := client.Open("/QbUk9Mk7mu.txt", client.O_RDWR|client.O_CREATE)
if(fd364 < 0) {
  panic("open failed")
}

//fd state: (62) xA5zJP1heOoOC8GYFnsgQkXQWcHWqH6ZmTMScuckqVeT8GB2CKLDEwf1cZc_jk7pAUQPfqCB

ret = client.Write(fd358, []byte("Wp8TJgqhc_gn_bHF6XIWmrYyvWn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) xA5zJP1heOoOC8GYFnsgQkXQWcHWqH6ZmTMScuckqVeT8GB2CKLDEwf1cZc_jkWp8TJgqhc_gn_bHF6XIWmrYyvWn

ret = client.Seek(fd347, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Seek(fd364, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


ret = client.Seek(fd356, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd351, []byte("L"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) L

ret = client.Close(fd354)
if(ret != 0) {
  panic("close failed")
}

//fd state: (46) CawTj8gqDjcxZqhqv8GMPFytpTws7hzFPBbpYgNLSWHywS_wh1TBfWYkmTu57op_2iFoQaZr2PjXJHxBEHqbXvddHMQ58JLJFY9SEYKQd0jBx__Rwr0Dq8o28Cif1bbG4wgvM

ret = client.Write(fd364, []byte("kVaKWScQ6vDOkFy7fqVndj83epemi0bN_QE9p5tI4ztfKboRIFGlIqUHs3HHIpviMSW2c7ToV5ysiucUiijhI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (131) CawTj8gqDjcxZqhqv8GMPFytpTws7hzFPBbpYgNLSWHywSkVaKWScQ6vDOkFy7fqVndj83epemi0bN_QE9p5tI4ztfKboRIFGlIqUHs3HHIpviMSW2c7ToV5ysiucUiijhIvM

ret = client.Close(fd358)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd352)
if(ret != 0) {
  panic("close failed")
}


fd365 := client.Open("/BU8ZeUXfGw.txt", client.O_RDWR|client.O_CREATE)
if(fd365 < 0) {
  panic("open failed")
}

//fd state: (0) dG503MzkJ1iImHSU9w8c7yp1MkbfC0SEsFaty_LKLUGNNkeCANNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV45gKkdLdJoLo6nJZBFwSdLobuJBt3Q3KNeawTQdk5__Ud6exoEvUEhMcwKcuz2b4cQRGZcs_DrGoih3M

ret = client.Write(fd365, []byte("DCeja6y6M_NMn5CmUb8rHqTXnV3c5z2u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) DCeja6y6M_NMn5CmUb8rHqTXnV3c5z2usFaty_LKLUGNNkeCANNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV45gKkdLdJoLo6nJZBFwSdLobuJBt3Q3KNeawTQdk5__Ud6exoEvUEhMcwKcuz2b4cQRGZcs_DrGoih3M

fd366 := client.Open("/WfcD25sP7W.txt", client.O_RDWR|client.O_CREATE)
if(fd366 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd347, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1SU_hVIcDh2tq6wF7TxZCE8ygNH820R8K7") {
  panic("wrong data returned")
}


fd367 := client.Open("/N2wW4FWwc1.txt", client.O_RDWR|client.O_CREATE)
if(fd367 < 0) {
  panic("open failed")
}


ret = client.Seek(fd350, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


ret = client.Close(fd367)
if(ret != 0) {
  panic("close failed")
}


fd368 := client.Open("/Wc8I1sgyqN.txt", client.O_RDWR|client.O_CREATE)
if(fd368 < 0) {
  panic("open failed")
}


fd369 := client.Open("/IIgch99vnP.txt", client.O_RDWR|client.O_CREATE)
if(fd369 < 0) {
  panic("open failed")
}


ret = client.Close(fd350)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd361)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd351)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd365, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "sFaty_LKLUGNNkeCANNN7XburV0b3oZkq") {
  panic("wrong data returned")
}


ret = client.Close(fd364)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd366, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd356, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jlv45fH_2KDgCF9h4jLPD8oGRKDkQP7I9xm") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd368, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Na3vDus5puioiFGF9DuLt") {
  panic("wrong data returned")
}


ret = client.Seek(fd368, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


fd370 := client.Open("/yunpN4Mkm8.txt", client.O_RDWR|client.O_CREATE)
if(fd370 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd366, []byte("PadKVQQvOOyhowN4VsfWa0ADbFJNbWpp1eC9mDQI6gs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) PadKVQQvOOyhowN4VsfWa0ADbFJNbWpp1eC9mDQI6gs

buf, ret = client.Read(fd368, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5puioiFGF9DuLt") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd370, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (43) PadKVQQvOOyhowN4VsfWa0ADbFJNbWpp1eC9mDQI6gs

ret = client.Write(fd366, []byte("OPlaAelhbmSOHnKl4dmQMa_sbavYvIKgCRzO26HylcUljmdlrLd0ajbzyz2c8B2j"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) PadKVQQvOOyhowN4VsfWa0ADbFJNbWpp1eC9mDQI6gsOPlaAelhbmSOHnKl4dmQMa_sbavYvIKgCRzO26HylcUljmdlrLd0ajbzyz2c8B2j
//fd state: (21) Na3vDus5puioiFGF9DuLt

ret = client.Write(fd368, []byte("Fr45g0gY6NWly7DZb34_YHc3d36E_irriZS7w8w_Tg3sk8h7sdsmh80MKIfMFRwQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) Na3vDus5puioiFGF9DuLtFr45g0gY6NWly7DZb34_YHc3d36E_irriZS7w8w_Tg3sk8h7sdsmh80MKIfMFRwQ

ret = client.Seek(fd370, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd365, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


fd371 := client.Open("/K74_aNYkXG.txt", client.O_RDWR|client.O_CREATE)
if(fd371 < 0) {
  panic("open failed")
}


fd372 := client.Open("/r_gM3BVTOr.txt", client.O_RDWR|client.O_CREATE)
if(fd372 < 0) {
  panic("open failed")
}


ret = client.Close(fd372)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd347, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


buf, ret = client.Read(fd356, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pqdunqfuyCrtM4596bD1XoGfGrpGWc6") {
  panic("wrong data returned")
}


ret = client.Close(fd371)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd356)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd368)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd365, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Seek(fd365, 89, client.SEEK_SET)
if(ret != 89) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 89)
  panic("seek failed")
}


buf, ret = client.Read(fd369, 89)
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


ret = client.Close(fd366)
if(ret != 0) {
  panic("close failed")
}

//fd state: (89) DCeja6y6M_NMn5CmUb8rHqTXnV3c5z2usFaty_LKLUGNNkeCANNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV45gKkdLdJoLo6nJZBFwSdLobuJBt3Q3KNeawTQdk5__Ud6exoEvUEhMcwKcuz2b4cQRGZcs_DrGoih3M

ret = client.Write(fd365, []byte("giwdDs6tx6gTvGZkDMiq_oNUjDZrWvH06662_g6NQ1IArdsGfsAVzovYe6VsyRdwrK9kCt50oawClwSdm6C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (172) DCeja6y6M_NMn5CmUb8rHqTXnV3c5z2usFaty_LKLUGNNkeCANNN7XburV0b3oZkqysQKqRa1nu2THqO9N62XLpV4giwdDs6tx6gTvGZkDMiq_oNUjDZrWvH06662_g6NQ1IArdsGfsAVzovYe6VsyRdwrK9kCt50oawClwSdm6C

ret = client.Close(fd362)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd369, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd370, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd370, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd369, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd365, 111, client.SEEK_SET)
if(ret != 111) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 111)
  panic("seek failed")
}


buf, ret = client.Read(fd370, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd365, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NUjDZrWvH06662_g6NQ1IArdsGfsAVzovYe6VsyRdwrK9kCt50oawClwSd") {
  panic("wrong data returned")
}


fd373 := client.Open("/3OuuWd5hLc.txt", client.O_RDWR|client.O_CREATE)
if(fd373 < 0) {
  panic("open failed")
}


fd374 := client.Open("/Kys6FX3D4n.txt", client.O_RDWR|client.O_CREATE)
if(fd374 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd373, []byte("n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) n

buf, ret = client.Read(fd373, 15)
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


fd375 := client.Open("/zj3dUeasL4.txt", client.O_RDWR|client.O_CREATE)
if(fd375 < 0) {
  panic("open failed")
}


ret = client.Seek(fd370, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd370, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd369)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd374)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd373)
if(ret != 0) {
  panic("close failed")
}


fd376 := client.Open("/YPMAsoYEb_.txt", client.O_RDWR|client.O_CREATE)
if(fd376 < 0) {
  panic("open failed")
}


ret = client.Seek(fd370, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd377 := client.Open("/QqwyHnv4d8.txt", client.O_RDWR|client.O_CREATE)
if(fd377 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd365, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "m6C") {
  panic("wrong data returned")
}


ret = client.Close(fd365)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd377, []byte("gBaxcGq7XrCKv1u7tIft_9ESj4k013ovbKhxhesZQ66d"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) gBaxcGq7XrCKv1u7tIft_9ESj4k013ovbKhxhesZQ66d

ret = client.Close(fd375)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd376)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd377, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


ret = client.Seek(fd377, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}

//fd state: (11) gBaxcGq7XrCKv1u7tIft_9ESj4k013ovbKhxhesZQ66d

ret = client.Write(fd377, []byte("pO0iHTk_gSlZCNTgpgH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) gBaxcGq7XrCpO0iHTk_gSlZCNTgpgHovbKhxhesZQ66d

fd378 := client.Open("/3wjB6vpdmW.txt", client.O_RDWR|client.O_CREATE)
if(fd378 < 0) {
  panic("open failed")
}


fd379 := client.Open("/_fKtkmMG2H.txt", client.O_RDWR|client.O_CREATE)
if(fd379 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd379, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (30) gBaxcGq7XrCpO0iHTk_gSlZCNTgpgHovbKhxhesZQ66d

ret = client.Write(fd377, []byte("Ypo6qPiR8xyf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) gBaxcGq7XrCpO0iHTk_gSlZCNTgpgHYpo6qPiR8xyf6d

fd380 := client.Open("/BFzwaNL70e.txt", client.O_RDWR|client.O_CREATE)
if(fd380 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd379, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd378, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd379)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd377, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "6d") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd370, []byte("vD6nN93r_0ViK0TYYqfU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) vD6nN93r_0ViK0TYYqfU
//fd state: (0) 

ret = client.Write(fd378, []byte("NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp11WS1g3Bg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp11WS1g3Bg

ret = client.Seek(fd380, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd381 := client.Open("/XwukkFpNt2.txt", client.O_RDWR|client.O_CREATE)
if(fd381 < 0) {
  panic("open failed")
}

//fd state: (44) gBaxcGq7XrCpO0iHTk_gSlZCNTgpgHYpo6qPiR8xyf6d

ret = client.Write(fd377, []byte("rqezc9F6h813CuRh9KuYfJCkKS6DChHZ8N58mqKwaqoHXqKuatHsHpyfZwIWqEvFXdXy3BkXQpvNW2OLVD7gIcrE21rhjqy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (139) gBaxcGq7XrCpO0iHTk_gSlZCNTgpgHYpo6qPiR8xyf6drqezc9F6h813CuRh9KuYfJCkKS6DChHZ8N58mqKwaqoHXqKuatHsHpyfZwIWqEvFXdXy3BkXQpvNW2OLVD7gIcrE21rhjqy
//fd state: (0) 

ret = client.Write(fd380, []byte("VMtiv0CSu0X7mNyDxx9yKZLgyirhnHMu4oK85hOcobv7yE2sS_jXYU9889dnySpN_zIh7X2uJ0jDwxoQ7h33nbM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) VMtiv0CSu0X7mNyDxx9yKZLgyirhnHMu4oK85hOcobv7yE2sS_jXYU9889dnySpN_zIh7X2uJ0jDwxoQ7h33nbM

ret = client.Seek(fd370, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd380)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd378, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd377, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (11) vD6nN93r_0ViK0TYYqfU

ret = client.Write(fd370, []byte("bNzUIB3gGHz0vA4pXSWAEJ6Vd7JIT9bggkQknjO0H_FViZ0R"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) vD6nN93r_0VbNzUIB3gGHz0vA4pXSWAEJ6Vd7JIT9bggkQknjO0H_FViZ0R

buf, ret = client.Read(fd377, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd370, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd378, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


ret = client.Seek(fd377, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Close(fd381)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd377, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}

//fd state: (62) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp11WS1g3Bg

ret = client.Write(fd378, []byte("lz7gcTKjXvbKKsfmnyshq8ZZ071jq9nIO4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7gcTKjXvbKKsfmnyshq8ZZ071jq9nIO4

ret = client.Close(fd377)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd378, 95, client.SEEK_SET)
if(ret != 95) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 95)
  panic("seek failed")
}


buf, ret = client.Read(fd378, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4") {
  panic("wrong data returned")
}

//fd state: (96) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7gcTKjXvbKKsfmnyshq8ZZ071jq9nIO4

ret = client.Write(fd378, []byte("5sUhZrUK_n03KYWONWJ0GSHQbmZ5ANPo5wO5TA8W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7gcTKjXvbKKsfmnyshq8ZZ071jq9nIO45sUhZrUK_n03KYWONWJ0GSHQbmZ5ANPo5wO5TA8W
//fd state: (59) vD6nN93r_0VbNzUIB3gGHz0vA4pXSWAEJ6Vd7JIT9bggkQknjO0H_FViZ0R

ret = client.Write(fd370, []byte("Gtj8WAe4m4xvdk7YS6RNIpcLqC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) vD6nN93r_0VbNzUIB3gGHz0vA4pXSWAEJ6Vd7JIT9bggkQknjO0H_FViZ0RGtj8WAe4m4xvdk7YS6RNIpcLqC

buf, ret = client.Read(fd378, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd370, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}

//fd state: (78) vD6nN93r_0VbNzUIB3gGHz0vA4pXSWAEJ6Vd7JIT9bggkQknjO0H_FViZ0RGtj8WAe4m4xvdk7YS6RNIpcLqC

ret = client.Write(fd370, []byte("WLkXkclBwoVOCQ5JJgzR4PSnFmrJrpk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (109) vD6nN93r_0VbNzUIB3gGHz0vA4pXSWAEJ6Vd7JIT9bggkQknjO0H_FViZ0RGtj8WAe4m4xvdk7YS6RWLkXkclBwoVOCQ5JJgzR4PSnFmrJrpk

ret = client.Seek(fd370, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd382 := client.Open("/KqykqQdiXF.txt", client.O_RDWR|client.O_CREATE)
if(fd382 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd378, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd370)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd382, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "XqeR5BQTtne3Y1VEaBCEHGeLyEwYX") {
  panic("wrong data returned")
}


ret = client.Close(fd382)
if(ret != 0) {
  panic("close failed")
}

//fd state: (136) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7gcTKjXvbKKsfmnyshq8ZZ071jq9nIO45sUhZrUK_n03KYWONWJ0GSHQbmZ5ANPo5wO5TA8W

ret = client.Write(fd378, []byte("uZj6dNlE8UaiiWHj2vHWs7H3q6siDYpyt_cwimLd7ZqsOiyl0TWI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (188) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7gcTKjXvbKKsfmnyshq8ZZ071jq9nIO45sUhZrUK_n03KYWONWJ0GSHQbmZ5ANPo5wO5TA8WuZj6dNlE8UaiiWHj2vHWs7H3q6siDYpyt_cwimLd7ZqsOiyl0TWI

buf, ret = client.Read(fd378, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd378, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd383 := client.Open("/na626QeuQC.txt", client.O_RDWR|client.O_CREATE)
if(fd383 < 0) {
  panic("open failed")
}


ret = client.Close(fd383)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd378, 84, client.SEEK_SET)
if(ret != 84) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 84)
  panic("seek failed")
}


buf, ret = client.Read(fd378, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ZZ071jq9nIO45sUh") {
  panic("wrong data returned")
}


ret = client.Seek(fd378, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


buf, ret = client.Read(fd378, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "UFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7gcTKjXv") {
  panic("wrong data returned")
}


ret = client.Close(fd378)
if(ret != 0) {
  panic("close failed")
}


fd384 := client.Open("/HN3w0IwSey.txt", client.O_RDWR|client.O_CREATE)
if(fd384 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd384, []byte("JeRoE9raBqOxoCpVbRlVeM8H0tXZVvx0jAL9avSB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) JeRoE9raBqOxoCpVbRlVeM8H0tXZVvx0jAL9avSB

buf, ret = client.Read(fd384, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (40) JeRoE9raBqOxoCpVbRlVeM8H0tXZVvx0jAL9avSB

ret = client.Write(fd384, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) JeRoE9raBqOxoCpVbRlVeM8H0tXZVvx0jAL9avSB
//fd state: (40) JeRoE9raBqOxoCpVbRlVeM8H0tXZVvx0jAL9avSB

ret = client.Write(fd384, []byte("Y1cbRL6CzgJ8_5TIqM8nl8d4ewbSmSihp0f_9Z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) JeRoE9raBqOxoCpVbRlVeM8H0tXZVvx0jAL9avSBY1cbRL6CzgJ8_5TIqM8nl8d4ewbSmSihp0f_9Z

ret = client.Close(fd384)
if(ret != 0) {
  panic("close failed")
}


fd385 := client.Open("/c5R3eXHXLk.txt", client.O_RDWR|client.O_CREATE)
if(fd385 < 0) {
  panic("open failed")
}


fd386 := client.Open("/3G482HgWOt.txt", client.O_RDWR|client.O_CREATE)
if(fd386 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd385, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd385, []byte("Rt0oz1DuXjpfs1L8AH5U2HfAkQ5_UwEqlqSOnheGBnOwsPzPYfuto8vsGM65OS1joguYO4dPk6A4bj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) Rt0oz1DuXjpfs1L8AH5U2HfAkQ5_UwEqlqSOnheGBnOwsPzPYfuto8vsGM65OS1joguYO4dPk6A4bj

buf, ret = client.Read(fd385, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd386, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd386)
if(ret != 0) {
  panic("close failed")
}

//fd state: (78) Rt0oz1DuXjpfs1L8AH5U2HfAkQ5_UwEqlqSOnheGBnOwsPzPYfuto8vsGM65OS1joguYO4dPk6A4bj

ret = client.Write(fd385, []byte("U2MLFl5YN9L"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) Rt0oz1DuXjpfs1L8AH5U2HfAkQ5_UwEqlqSOnheGBnOwsPzPYfuto8vsGM65OS1joguYO4dPk6A4bjU2MLFl5YN9L

ret = client.Close(fd385)
if(ret != 0) {
  panic("close failed")
}


fd387 := client.Open("/ZLj_JM8E2R.txt", client.O_RDWR|client.O_CREATE)
if(fd387 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd387, []byte("CgN9QyVu0Q48cxA8NICpDxa2ewBdiiDuIOi0kgEzYwV6QYJCYbYcbh5quKbD04zHl_OP3S3fu7a4_WAu3Di_2LzflzrROgJF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) CgN9QyVu0Q48cxA8NICpDxa2ewBdiiDuIOi0kgEzYwV6QYJCYbYcbh5quKbD04zHl_OP3S3fu7a4_WAu3Di_2LzflzrROgJF
//fd state: (96) CgN9QyVu0Q48cxA8NICpDxa2ewBdiiDuIOi0kgEzYwV6QYJCYbYcbh5quKbD04zHl_OP3S3fu7a4_WAu3Di_2LzflzrROgJF

ret = client.Write(fd387, []byte("bxyl3IhiPvedttKiLzY6fHRTQmxp2kkNaVQy3g8Q6glOw2K0sttn8Qc8KRN5Rj635UDwSgk7LVwqWDC74Ac2QzzaPQsOdm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (190) CgN9QyVu0Q48cxA8NICpDxa2ewBdiiDuIOi0kgEzYwV6QYJCYbYcbh5quKbD04zHl_OP3S3fu7a4_WAu3Di_2LzflzrROgJFbxyl3IhiPvedttKiLzY6fHRTQmxp2kkNaVQy3g8Q6glOw2K0sttn8Qc8KRN5Rj635UDwSgk7LVwqWDC74Ac2QzzaPQsOdm

ret = client.Close(fd387)
if(ret != 0) {
  panic("close failed")
}


fd388 := client.Open("/5L6AH59F8o.txt", client.O_RDWR|client.O_CREATE)
if(fd388 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd388, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd389 := client.Open("/NdZKxMANLT.txt", client.O_RDWR|client.O_CREATE)
if(fd389 < 0) {
  panic("open failed")
}


fd390 := client.Open("/IlPekcd4f8.txt", client.O_RDWR|client.O_CREATE)
if(fd390 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd390, []byte("hFUP9k0dPL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) hFUP9k0dPL

ret = client.Close(fd390)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd389, []byte("P3uX2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) P3uX2

buf, ret = client.Read(fd388, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd389, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd391 := client.Open("/gsTfmWXCkm.txt", client.O_RDWR|client.O_CREATE)
if(fd391 < 0) {
  panic("open failed")
}


ret = client.Seek(fd389, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (3) P3uX2

ret = client.Write(fd389, []byte("baJxHXF34Evj2EBqoimsaDQekRbTVpnouuQIG8jhjaOuCd97VFjeOf19uSNvxdiiKexYEs29seGsZHdtla"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) P3ubaJxHXF34Evj2EBqoimsaDQekRbTVpnouuQIG8jhjaOuCd97VFjeOf19uSNvxdiiKexYEs29seGsZHdtla

buf, ret = client.Read(fd388, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd391, []byte("fl45Wy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) fl45Wy

buf, ret = client.Read(fd388, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd388)
if(ret != 0) {
  panic("close failed")
}


fd392 := client.Open("/LxlbfoT7LK.txt", client.O_RDWR|client.O_CREATE)
if(fd392 < 0) {
  panic("open failed")
}

//fd state: (85) P3ubaJxHXF34Evj2EBqoimsaDQekRbTVpnouuQIG8jhjaOuCd97VFjeOf19uSNvxdiiKexYEs29seGsZHdtla

ret = client.Write(fd389, []byte("Ep3qZmjPVlgAsCP1XKr28jRpCtsLz0awumh_O9a0dqCNOD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (131) P3ubaJxHXF34Evj2EBqoimsaDQekRbTVpnouuQIG8jhjaOuCd97VFjeOf19uSNvxdiiKexYEs29seGsZHdtlaEp3qZmjPVlgAsCP1XKr28jRpCtsLz0awumh_O9a0dqCNOD

ret = client.Close(fd391)
if(ret != 0) {
  panic("close failed")
}


fd393 := client.Open("/jlETz8ELRp.txt", client.O_RDWR|client.O_CREATE)
if(fd393 < 0) {
  panic("open failed")
}


ret = client.Close(fd389)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd392, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd392, []byte("kLCEh4PW6edsDYtt8bNLpm7LA5PiIKd8W79DNqomBNJCs_f2B5u38di"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) kLCEh4PW6edsDYtt8bNLpm7LA5PiIKd8W79DNqomBNJCs_f2B5u38di

ret = client.Seek(fd393, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd393, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd392, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Seek(fd392, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


ret = client.Seek(fd392, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd393)
if(ret != 0) {
  panic("close failed")
}


fd394 := client.Open("/lvnB8YpciR.txt", client.O_RDWR|client.O_CREATE)
if(fd394 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd394, []byte("lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0
//fd state: (37) lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0

ret = client.Write(fd394, []byte("G842jNCzBU9Z5posvZfFFZVIlx0U9KF4SWYJBkGZZAd6f3NvqLDqxYvrGp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0G842jNCzBU9Z5posvZfFFZVIlx0U9KF4SWYJBkGZZAd6f3NvqLDqxYvrGp
//fd state: (95) lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0G842jNCzBU9Z5posvZfFFZVIlx0U9KF4SWYJBkGZZAd6f3NvqLDqxYvrGp

ret = client.Write(fd394, []byte("U_AYCLmPMQopKqZa6bPHhvYH1bC4V_mrgm4dbx4PGTGS_khCWOv8mJ281ER9bi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (157) lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0G842jNCzBU9Z5posvZfFFZVIlx0U9KF4SWYJBkGZZAd6f3NvqLDqxYvrGpU_AYCLmPMQopKqZa6bPHhvYH1bC4V_mrgm4dbx4PGTGS_khCWOv8mJ281ER9bi
//fd state: (157) lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0G842jNCzBU9Z5posvZfFFZVIlx0U9KF4SWYJBkGZZAd6f3NvqLDqxYvrGpU_AYCLmPMQopKqZa6bPHhvYH1bC4V_mrgm4dbx4PGTGS_khCWOv8mJ281ER9bi

ret = client.Write(fd394, []byte("N9bfJG4ebnmosnPRGc6Yn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (178) lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0G842jNCzBU9Z5posvZfFFZVIlx0U9KF4SWYJBkGZZAd6f3NvqLDqxYvrGpU_AYCLmPMQopKqZa6bPHhvYH1bC4V_mrgm4dbx4PGTGS_khCWOv8mJ281ER9biN9bfJG4ebnmosnPRGc6Yn
//fd state: (178) lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0G842jNCzBU9Z5posvZfFFZVIlx0U9KF4SWYJBkGZZAd6f3NvqLDqxYvrGpU_AYCLmPMQopKqZa6bPHhvYH1bC4V_mrgm4dbx4PGTGS_khCWOv8mJ281ER9biN9bfJG4ebnmosnPRGc6Yn

ret = client.Write(fd394, []byte("DTLNR9GX91ofaJKUKhfHUL49s3gsMN9qRFk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (213) lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0G842jNCzBU9Z5posvZfFFZVIlx0U9KF4SWYJBkGZZAd6f3NvqLDqxYvrGpU_AYCLmPMQopKqZa6bPHhvYH1bC4V_mrgm4dbx4PGTGS_khCWOv8mJ281ER9biN9bfJG4ebnmosnPRGc6YnDTLNR9GX91ofaJKUKhfHUL49s3gsMN9qRFk

fd395 := client.Open("/W8PDn29X7H.txt", client.O_RDWR|client.O_CREATE)
if(fd395 < 0) {
  panic("open failed")
}


ret = client.Seek(fd395, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd392, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Seek(fd395, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd396 := client.Open("/BiqkOPZuTM.txt", client.O_RDWR|client.O_CREATE)
if(fd396 < 0) {
  panic("open failed")
}


ret = client.Seek(fd394, 180, client.SEEK_SET)
if(ret != 180) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 180)
  panic("seek failed")
}


buf, ret = client.Read(fd396, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd395, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd394, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LNR9GX91ofaJKUKhfHUL49s3gsMN9qRFk") {
  panic("wrong data returned")
}


ret = client.Seek(fd392, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


buf, ret = client.Read(fd392, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd395, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd397 := client.Open("/4Ki0uUgl43.txt", client.O_RDWR|client.O_CREATE)
if(fd397 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd396, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd398 := client.Open("/L3a7VdOTLE.txt", client.O_RDWR|client.O_CREATE)
if(fd398 < 0) {
  panic("open failed")
}


ret = client.Close(fd395)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd392, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}


fd399 := client.Open("/IIgch99vnP.txt", client.O_RDWR|client.O_CREATE)
if(fd399 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd399, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd394, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd400 := client.Open("/SRZ2gm6O3V.txt", client.O_RDWR|client.O_CREATE)
if(fd400 < 0) {
  panic("open failed")
}

//fd state: (0) x1

ret = client.Write(fd397, []byte("S0zsM9IK2psm9TA5LpqBG9rqduzC2O0zCBgD4fv5RDgio9YIuxPieJS6YqX_PSydGvAKZKf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (71) S0zsM9IK2psm9TA5LpqBG9rqduzC2O0zCBgD4fv5RDgio9YIuxPieJS6YqX_PSydGvAKZKf

buf, ret = client.Read(fd394, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd399, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd394, 40)
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


fd401 := client.Open("/2CT8Qjb05c.txt", client.O_RDWR|client.O_CREATE)
if(fd401 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd401, []byte("KBZRORQSBJkqGlFs9z0uPPw6fIzDCCp2mJ8l7LNy_aWSgiBanaUnASGgOyodPz2Jo_jKjFakE5y9cV_B8C5hu6vB38iZW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) KBZRORQSBJkqGlFs9z0uPPw6fIzDCCp2mJ8l7LNy_aWSgiBanaUnASGgOyodPz2Jo_jKjFakE5y9cV_B8C5hu6vB38iZW

ret = client.Seek(fd399, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd394, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


fd402 := client.Open("/8vaK9bbYr6.txt", client.O_RDWR|client.O_CREATE)
if(fd402 < 0) {
  panic("open failed")
}


ret = client.Close(fd399)
if(ret != 0) {
  panic("close failed")
}


fd403 := client.Open("/EXgQzL0J8u.txt", client.O_RDWR|client.O_CREATE)
if(fd403 < 0) {
  panic("open failed")
}

//fd state: (25) lj9oeTlLsvxzVOLJHBttgFcW7mwVbBrpd5gG0G842jNCzBU9Z5posvZfFFZVIlx0U9KF4SWYJBkGZZAd6f3NvqLDqxYvrGpU_AYCLmPMQopKqZa6bPHhvYH1bC4V_mrgm4dbx4PGTGS_khCWOv8mJ281ER9biN9bfJG4ebnmosnPRGc6YnDTLNR9GX91ofaJKUKhfHUL49s3gsMN9qRFk

ret = client.Write(fd394, []byte("RICegANXWTZn3bNceONTGBSLXZjmHAi5FxC9Y6vZ9yRcO7Vk9_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) lj9oeTlLsvxzVOLJHBttgFcW7RICegANXWTZn3bNceONTGBSLXZjmHAi5FxC9Y6vZ9yRcO7Vk9_GZZAd6f3NvqLDqxYvrGpU_AYCLmPMQopKqZa6bPHhvYH1bC4V_mrgm4dbx4PGTGS_khCWOv8mJ281ER9biN9bfJG4ebnmosnPRGc6YnDTLNR9GX91ofaJKUKhfHUL49s3gsMN9qRFk

buf, ret = client.Read(fd392, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5u38di") {
  panic("wrong data returned")
}


fd404 := client.Open("/6zCWkaOFR7.txt", client.O_RDWR|client.O_CREATE)
if(fd404 < 0) {
  panic("open failed")
}

//fd state: (0) BRqIXriBg5LZvWiBjsk48riTc94JPun6lga_fhGvZzIle1C2nDX9om5eiV2yiVGXdk3m1HTNNrAgK5CHWfBsLB0Zn1rTuMT7wVOIl3huWdCCOtUAttXbGBjUEFzHSlxaaDxGhxtyOH8Ms_MduFLQofmJG0

ret = client.Write(fd398, []byte("x1TH2XW7xoFZB036u8z3DHxATyaL_32"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) x1TH2XW7xoFZB036u8z3DHxATyaL_326lga_fhGvZzIle1C2nDX9om5eiV2yiVGXdk3m1HTNNrAgK5CHWfBsLB0Zn1rTuMT7wVOIl3huWdCCOtUAttXbGBjUEFzHSlxaaDxGhxtyOH8Ms_MduFLQofmJG0

buf, ret = client.Read(fd392, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd403)
if(ret != 0) {
  panic("close failed")
}

//fd state: (55) kLCEh4PW6edsDYtt8bNLpm7LA5PiIKd8W79DNqomBNJCs_f2B5u38di

ret = client.Write(fd392, []byte("S9xa8utyNrqrcpPbgJQy0XpyO1RRGUX63hbWD4dsiiukwHgdwo_RbBYql_fzP3kv4YgTev7MsPz_IfE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) kLCEh4PW6edsDYtt8bNLpm7LA5PiIKd8W79DNqomBNJCs_f2B5u38diS9xa8utyNrqrcpPbgJQy0XpyO1RRGUX63hbWD4dsiiukwHgdwo_RbBYql_fzP3kv4YgTev7MsPz_IfE

ret = client.Close(fd404)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd400)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd398, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "6lga_fhGvZzIle1C2nDX9om5eiV2yiVG") {
  panic("wrong data returned")
}

//fd state: (63) x1TH2XW7xoFZB036u8z3DHxATyaL_326lga_fhGvZzIle1C2nDX9om5eiV2yiVGXdk3m1HTNNrAgK5CHWfBsLB0Zn1rTuMT7wVOIl3huWdCCOtUAttXbGBjUEFzHSlxaaDxGhxtyOH8Ms_MduFLQofmJG0

ret = client.Write(fd398, []byte("FWKkIS5JxTgF0EWQgkSBivaRckEWtJvhBrMrwEhDQDbCzw1Qq9tBCE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) x1TH2XW7xoFZB036u8z3DHxATyaL_326lga_fhGvZzIle1C2nDX9om5eiV2yiVGFWKkIS5JxTgF0EWQgkSBivaRckEWtJvhBrMrwEhDQDbCzw1Qq9tBCEBjUEFzHSlxaaDxGhxtyOH8Ms_MduFLQofmJG0

fd405 := client.Open("/giqe_vyTh1.txt", client.O_RDWR|client.O_CREATE)
if(fd405 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd398, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "BjUEFzHSlxaaD") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd397, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (130) x1TH2XW7xoFZB036u8z3DHxATyaL_326lga_fhGvZzIle1C2nDX9om5eiV2yiVGFWKkIS5JxTgF0EWQgkSBivaRckEWtJvhBrMrwEhDQDbCzw1Qq9tBCEBjUEFzHSlxaaDxGhxtyOH8Ms_MduFLQofmJG0

ret = client.Write(fd398, []byte("NUIWIP7wqp70oMV7cahoFNxZjBfPIhk4KKQWDRf6dEYuTGCv8iIaogzCORQTH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (191) x1TH2XW7xoFZB036u8z3DHxATyaL_326lga_fhGvZzIle1C2nDX9om5eiV2yiVGFWKkIS5JxTgF0EWQgkSBivaRckEWtJvhBrMrwEhDQDbCzw1Qq9tBCEBjUEFzHSlxaaDNUIWIP7wqp70oMV7cahoFNxZjBfPIhk4KKQWDRf6dEYuTGCv8iIaogzCORQTH

ret = client.Seek(fd397, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


ret = client.Close(fd401)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd394, 206, client.SEEK_SET)
if(ret != 206) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 206)
  panic("seek failed")
}


ret = client.Seek(fd392, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


buf, ret = client.Read(fd402, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd398, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


buf, ret = client.Read(fd392, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "A5PiIKd8W79DNqomBNJCs_f2B5u38diS9xa8utyNrqrcpPbgJQy0X") {
  panic("wrong data returned")
}


ret = client.Seek(fd402, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd406 := client.Open("/LtD3fxRhwd.txt", client.O_RDWR|client.O_CREATE)
if(fd406 < 0) {
  panic("open failed")
}


ret = client.Close(fd394)
if(ret != 0) {
  panic("close failed")
}


fd407 := client.Open("/etbCfgFtZR.txt", client.O_RDWR|client.O_CREATE)
if(fd407 < 0) {
  panic("open failed")
}


fd408 := client.Open("/RPLGuQQ93O.txt", client.O_RDWR|client.O_CREATE)
if(fd408 < 0) {
  panic("open failed")
}


ret = client.Close(fd392)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd406, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd398)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd402, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) TCzGHiVA_Lob_c0CcGIbvDsAQHT3RD51wkho4nD1fsflImtXzfLck8XY54bkULQ7jfAzhIXr3HrV_TzquIp6j0xbNXvpdKiP8mYKVkZGQqolMkupY3sBNNUMHnRAizIE7mRHEODHkQV3eFbHYUap8ZaDLejAXnoecog3Sx8Lg5Gi

ret = client.Write(fd405, []byte("s03"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) s03GHiVA_Lob_c0CcGIbvDsAQHT3RD51wkho4nD1fsflImtXzfLck8XY54bkULQ7jfAzhIXr3HrV_TzquIp6j0xbNXvpdKiP8mYKVkZGQqolMkupY3sBNNUMHnRAizIE7mRHEODHkQV3eFbHYUap8ZaDLejAXnoecog3Sx8Lg5Gi

fd409 := client.Open("/X8KDKPPquI.txt", client.O_RDWR|client.O_CREATE)
if(fd409 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd397, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YIuxPieJS6YqX_PSyd") {
  panic("wrong data returned")
}


ret = client.Close(fd405)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd397, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GvAKZKf") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd407, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (71) S0zsM9IK2psm9TA5LpqBG9rqduzC2O0zCBgD4fv5RDgio9YIuxPieJS6YqX_PSydGvAKZKf

ret = client.Write(fd397, []byte("yAcPt7Lb2yFrqS8Le28cHV983rLTcs4eGJ0lNQfg5NqLIIzMteNy3B"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (125) S0zsM9IK2psm9TA5LpqBG9rqduzC2O0zCBgD4fv5RDgio9YIuxPieJS6YqX_PSydGvAKZKfyAcPt7Lb2yFrqS8Le28cHV983rLTcs4eGJ0lNQfg5NqLIIzMteNy3B

ret = client.Close(fd407)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd397, 85, client.SEEK_SET)
if(ret != 85) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 85)
  panic("seek failed")
}


ret = client.Seek(fd406, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) lOx8cyAyByxnB759OnbcBUINxBW

ret = client.Write(fd409, []byte("3hMNjngbn8hZrZ3Ny2lAZUwWqSJcFrnN3D7YiWzNlFVOLWe7SiHfD2QZFGdc8a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) 3hMNjngbn8hZrZ3Ny2lAZUwWqSJcFrnN3D7YiWzNlFVOLWe7SiHfD2QZFGdc8a

ret = client.Close(fd408)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd397, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd406, []byte("sXLwmkRpNxy5kGgptuC8d2O1mzCQPf0uqT8dKGpWRLj51Yka5eyIXSx2NwMK_6LAvD7aPurIUDDd4GTN8e5WEFlvo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) sXLwmkRpNxy5kGgptuC8d2O1mzCQPf0uqT8dKGpWRLj51Yka5eyIXSx2NwMK_6LAvD7aPurIUDDd4GTN8e5WEFlvo

fd410 := client.Open("/NDaLK8HBXg.txt", client.O_RDWR|client.O_CREATE)
if(fd410 < 0) {
  panic("open failed")
}


ret = client.Seek(fd409, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


ret = client.Seek(fd410, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


buf, ret = client.Read(fd397, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "6YqX_PSydGvAKZKfyAcPt7Lb2yFrqS8Le28cHV983rLTcs4eGJ0lNQfg5NqLIIzM") {
  panic("wrong data returned")
}


fd411 := client.Open("/Cw5u_NagMb.txt", client.O_RDWR|client.O_CREATE)
if(fd411 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd411, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (21) _Cb4_9cMgwbLx4w1hvzD6ri4ml0MCqWbeSH6NAPfFP6NAY3iTYmtWZKVZT8dnwZgJ_wn_yuyCevgBDLCFgKRyTQ1eB8bZq1Pc2t3uvsjhGqKbx0V_CnQx27xDTy98s4qzewjk5p50ZZLyj0vU8NSGc

ret = client.Write(fd410, []byte("F60xU3mrcdX3_f8lTUDyDRHjb5gGcvaItubk0uhZ74HqntI5XDpCHifTJxg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) _Cb4_9cMgwbLx4w1hvzD6F60xU3mrcdX3_f8lTUDyDRHjb5gGcvaItubk0uhZ74HqntI5XDpCHifTJxgFgKRyTQ1eB8bZq1Pc2t3uvsjhGqKbx0V_CnQx27xDTy98s4qzewjk5p50ZZLyj0vU8NSGc

ret = client.Close(fd410)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd397, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd411, []byte("5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0Vc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0Vc

ret = client.Seek(fd409, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


buf, ret = client.Read(fd411, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd412 := client.Open("/zZt3zhX_S8.txt", client.O_RDWR|client.O_CREATE)
if(fd412 < 0) {
  panic("open failed")
}


fd413 := client.Open("/X8KDKPPquI.txt", client.O_RDWR|client.O_CREATE)
if(fd413 < 0) {
  panic("open failed")
}


ret = client.Seek(fd397, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


ret = client.Seek(fd402, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd409, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "FrnN3D7YiWzNlFV") {
  panic("wrong data returned")
}


ret = client.Seek(fd406, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


ret = client.Seek(fd402, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (32) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0Vc

ret = client.Write(fd411, []byte("PnUQv9hhcrw0tF_C4e2LpkBr4JMf_T0Yv4PuDGlecxbQ2IVJ4LSKQqfIwMc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0VcPnUQv9hhcrw0tF_C4e2LpkBr4JMf_T0Yv4PuDGlecxbQ2IVJ4LSKQqfIwMc

ret = client.Close(fd406)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd413)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd402, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (91) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0VcPnUQv9hhcrw0tF_C4e2LpkBr4JMf_T0Yv4PuDGlecxbQ2IVJ4LSKQqfIwMc

ret = client.Write(fd411, []byte("FYOZ_KEPz5TK9UqdA3DDUtTSSjHLIycD9vEjOmJ_I4I8S2HeOoOxkWgH5yTXgGdq5wj6YSXYNQvq9FWrNmteTwExT52Tap"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (185) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0VcPnUQv9hhcrw0tF_C4e2LpkBr4JMf_T0Yv4PuDGlecxbQ2IVJ4LSKQqfIwMcFYOZ_KEPz5TK9UqdA3DDUtTSSjHLIycD9vEjOmJ_I4I8S2HeOoOxkWgH5yTXgGdq5wj6YSXYNQvq9FWrNmteTwExT52Tap

ret = client.Seek(fd409, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd412)
if(ret != 0) {
  panic("close failed")
}

//fd state: (185) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0VcPnUQv9hhcrw0tF_C4e2LpkBr4JMf_T0Yv4PuDGlecxbQ2IVJ4LSKQqfIwMcFYOZ_KEPz5TK9UqdA3DDUtTSSjHLIycD9vEjOmJ_I4I8S2HeOoOxkWgH5yTXgGdq5wj6YSXYNQvq9FWrNmteTwExT52Tap

ret = client.Write(fd411, []byte("KQ5uOUnNhcEfOPDjYTWKsDSmx_SB4lANxsASaBPfHVZwaUa9g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (234) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0VcPnUQv9hhcrw0tF_C4e2LpkBr4JMf_T0Yv4PuDGlecxbQ2IVJ4LSKQqfIwMcFYOZ_KEPz5TK9UqdA3DDUtTSSjHLIycD9vEjOmJ_I4I8S2HeOoOxkWgH5yTXgGdq5wj6YSXYNQvq9FWrNmteTwExT52TapKQ5uOUnNhcEfOPDjYTWKsDSmx_SB4lANxsASaBPfHVZwaUa9g

ret = client.Seek(fd402, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd402)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd409, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


ret = client.Seek(fd397, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


buf, ret = client.Read(fd397, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "sm9TA5LpqBG9rqdu") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd409, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lFVOLWe7SiHfD2QZFGdc8a") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd411, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (234) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0VcPnUQv9hhcrw0tF_C4e2LpkBr4JMf_T0Yv4PuDGlecxbQ2IVJ4LSKQqfIwMcFYOZ_KEPz5TK9UqdA3DDUtTSSjHLIycD9vEjOmJ_I4I8S2HeOoOxkWgH5yTXgGdq5wj6YSXYNQvq9FWrNmteTwExT52TapKQ5uOUnNhcEfOPDjYTWKsDSmx_SB4lANxsASaBPfHVZwaUa9g

ret = client.Write(fd411, []byte("MAMgLqd4AbpI6TCzOEiKpfHqlTJ3PvACrD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (268) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0VcPnUQv9hhcrw0tF_C4e2LpkBr4JMf_T0Yv4PuDGlecxbQ2IVJ4LSKQqfIwMcFYOZ_KEPz5TK9UqdA3DDUtTSSjHLIycD9vEjOmJ_I4I8S2HeOoOxkWgH5yTXgGdq5wj6YSXYNQvq9FWrNmteTwExT52TapKQ5uOUnNhcEfOPDjYTWKsDSmx_SB4lANxsASaBPfHVZwaUa9gMAMgLqd4AbpI6TCzOEiKpfHqlTJ3PvACrD
//fd state: (62) 3hMNjngbn8hZrZ3Ny2lAZUwWqSJcFrnN3D7YiWzNlFVOLWe7SiHfD2QZFGdc8a

ret = client.Write(fd409, []byte("Rx6fXl2qRLyZKYMYz7d2k7SJvwTlZVcxJjDwiBd09LyB2_QTnE1QDsabI6xftXzyq2aWYoRuHMBHI1IkOBaA27m"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (149) 3hMNjngbn8hZrZ3Ny2lAZUwWqSJcFrnN3D7YiWzNlFVOLWe7SiHfD2QZFGdc8aRx6fXl2qRLyZKYMYz7d2k7SJvwTlZVcxJjDwiBd09LyB2_QTnE1QDsabI6xftXzyq2aWYoRuHMBHI1IkOBaA27m

ret = client.Close(fd409)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd411, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd397, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "zC2O0zCBgD4fv5RDgio9YIuxPieJS6YqX_PSydGvAKZKfyAcPt7Lb2yFrqS") {
  panic("wrong data returned")
}


fd414 := client.Open("/LoIyZd2AM5.txt", client.O_RDWR|client.O_CREATE)
if(fd414 < 0) {
  panic("open failed")
}


ret = client.Seek(fd414, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}

//fd state: (268) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0VcPnUQv9hhcrw0tF_C4e2LpkBr4JMf_T0Yv4PuDGlecxbQ2IVJ4LSKQqfIwMcFYOZ_KEPz5TK9UqdA3DDUtTSSjHLIycD9vEjOmJ_I4I8S2HeOoOxkWgH5yTXgGdq5wj6YSXYNQvq9FWrNmteTwExT52TapKQ5uOUnNhcEfOPDjYTWKsDSmx_SB4lANxsASaBPfHVZwaUa9gMAMgLqd4AbpI6TCzOEiKpfHqlTJ3PvACrD

ret = client.Write(fd411, []byte("g0ZAvWLCXyMv8LLBZZK69XCudsMRzxL51_qXOgHY5oIPeyJL4zJ50jualBgyrmSr92fsX4mXEwFx7OT2vKjhcnyoxRO93QAiA0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (366) 5yIbwCsrfJHAUU8tsU5xsc2hA53Sl0VcPnUQv9hhcrw0tF_C4e2LpkBr4JMf_T0Yv4PuDGlecxbQ2IVJ4LSKQqfIwMcFYOZ_KEPz5TK9UqdA3DDUtTSSjHLIycD9vEjOmJ_I4I8S2HeOoOxkWgH5yTXgGdq5wj6YSXYNQvq9FWrNmteTwExT52TapKQ5uOUnNhcEfOPDjYTWKsDSmx_SB4lANxsASaBPfHVZwaUa9gMAMgLqd4AbpI6TCzOEiKpfHqlTJ3PvACrDg0ZAvWLCXyMv8LLBZZK69XCudsMRzxL51_qXOgHY5oIPeyJL4zJ50jualBgyrmSr92fsX4mXEwFx7OT2vKjhcnyoxRO93QAiA0

ret = client.Seek(fd411, 64, client.SEEK_SET)
if(ret != 64) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 64)
  panic("seek failed")
}


fd415 := client.Open("/ZFdC9v3JgI.txt", client.O_RDWR|client.O_CREATE)
if(fd415 < 0) {
  panic("open failed")
}


fd416 := client.Open("/oXvVSX6Wbh.txt", client.O_RDWR|client.O_CREATE)
if(fd416 < 0) {
  panic("open failed")
}


ret = client.Seek(fd416, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd416, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd411, 139, client.SEEK_SET)
if(ret != 139) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 139)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd416, []byte("auRjHxmcWivFvUm5mRroiiGOIOhTzgPj48MuYGx35OUnlWoamYLYSr3S12T_LpOTsjjnr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) auRjHxmcWivFvUm5mRroiiGOIOhTzgPj48MuYGx35OUnlWoamYLYSr3S12T_LpOTsjjnr

ret = client.Close(fd415)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd416, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd416, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Seek(fd414, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd416, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jHxmcWivFvUm5mRroiiGOIOhTzgPj48MuYGx35OUnlWoamYLYSr3S12T_LpOTsjjnr") {
  panic("wrong data returned")
}


ret = client.Seek(fd411, 111, client.SEEK_SET)
if(ret != 111) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 111)
  panic("seek failed")
}


buf, ret = client.Read(fd414, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd417 := client.Open("/zCkZRCqM20.txt", client.O_RDWR|client.O_CREATE)
if(fd417 < 0) {
  panic("open failed")
}


ret = client.Close(fd397)
if(ret != 0) {
  panic("close failed")
}


fd418 := client.Open("/rQron0Xzpa.txt", client.O_RDWR|client.O_CREATE)
if(fd418 < 0) {
  panic("open failed")
}

//fd state: (69) auRjHxmcWivFvUm5mRroiiGOIOhTzgPj48MuYGx35OUnlWoamYLYSr3S12T_LpOTsjjnr

ret = client.Write(fd416, []byte("zaD8tZxsyEHV_ScOLI1O_ssN7cMlkln6p7DJZt6M2LRp2j9tKMyhGxn1TCwWYtBWT3oEjWt3txMKLLMgzewdSCRRmCCa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (161) auRjHxmcWivFvUm5mRroiiGOIOhTzgPj48MuYGx35OUnlWoamYLYSr3S12T_LpOTsjjnrzaD8tZxsyEHV_ScOLI1O_ssN7cMlkln6p7DJZt6M2LRp2j9tKMyhGxn1TCwWYtBWT3oEjWt3txMKLLMgzewdSCRRmCCa

ret = client.Close(fd414)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd417, []byte("Y_YGrtBXj0zFtExD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (16) Y_YGrtBXj0zFtExD

ret = client.Close(fd416)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd418, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd417, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd411, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "UtTSSjHLIycD9vEjOmJ_I4I8S2HeOoOxkWgH5yTXgGdq5wj6YSXYNQvq9FWrNmteTwExT52TapK") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd417, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd411, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Q5uOUnNhcEfOPDjYTWKsDSmx_SB4lANxsASaBPfHVZwaUa9gMAMgLqd4AbpI6T") {
  panic("wrong data returned")
}


fd419 := client.Open("/TZxRytMWBY.txt", client.O_RDWR|client.O_CREATE)
if(fd419 < 0) {
  panic("open failed")
}


fd420 := client.Open("/bQMNCmBBJY.txt", client.O_RDWR|client.O_CREATE)
if(fd420 < 0) {
  panic("open failed")
}


ret = client.Seek(fd419, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd421 := client.Open("/gwIXoPWovz.txt", client.O_RDWR|client.O_CREATE)
if(fd421 < 0) {
  panic("open failed")
}

//fd state: (16) Y_YGrtBXj0zFtExD

ret = client.Write(fd417, []byte("yexpKZtLSf5_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) Y_YGrtBXj0zFtExDyexpKZtLSf5_

ret = client.Close(fd417)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd421, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd418, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd420, []byte("ZQHRsalGFulaWxVckyeAP7GJbzQpv0I8Eml3mnFddh24jRpKORsUYGHDOG5D3U_oSqpP1rwrM1FxAxcCyDPNgnQA5K0GU6Vo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) ZQHRsalGFulaWxVckyeAP7GJbzQpv0I8Eml3mnFddh24jRpKORsUYGHDOG5D3U_oSqpP1rwrM1FxAxcCyDPNgnQA5K0GU6Vo

ret = client.Seek(fd420, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Seek(fd411, 140, client.SEEK_SET)
if(ret != 140) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 140)
  panic("seek failed")
}


ret = client.Close(fd420)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd418)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd419)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd421, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd411)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd421, []byte("2HPhr2w0p1AnIuXpdVPexi7ru_UctjWXBUAZXEfs_kwd3H5pNcFHoCf21"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) 2HPhr2w0p1AnIuXpdVPexi7ru_UctjWXBUAZXEfs_kwd3H5pNcFHoCf21
//fd state: (57) 2HPhr2w0p1AnIuXpdVPexi7ru_UctjWXBUAZXEfs_kwd3H5pNcFHoCf21

ret = client.Write(fd421, []byte("8VwgusvmNtqMmrV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) 2HPhr2w0p1AnIuXpdVPexi7ru_UctjWXBUAZXEfs_kwd3H5pNcFHoCf218VwgusvmNtqMmrV

fd422 := client.Open("/OwtKj7ePV0.txt", client.O_RDWR|client.O_CREATE)
if(fd422 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd422, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd421, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd422, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd422, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd423 := client.Open("/ZmYnI8Rd8S.txt", client.O_RDWR|client.O_CREATE)
if(fd423 < 0) {
  panic("open failed")
}


fd424 := client.Open("/lzZlQj9QSV.txt", client.O_RDWR|client.O_CREATE)
if(fd424 < 0) {
  panic("open failed")
}


ret = client.Close(fd421)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd423)
if(ret != 0) {
  panic("close failed")
}


fd425 := client.Open("/MyEfiMKv5b.txt", client.O_RDWR|client.O_CREATE)
if(fd425 < 0) {
  panic("open failed")
}


ret = client.Close(fd425)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd422, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd426 := client.Open("/QzL5Cu3b1K.txt", client.O_RDWR|client.O_CREATE)
if(fd426 < 0) {
  panic("open failed")
}


ret = client.Seek(fd426, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd426, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd426, []byte("hhlUcKwGJJ3y_bGgAdMGiwl9zoCP1NvMiRg7xVFStgdPJHr55MDVaNBac6a54Tol3d"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) hhlUcKwGJJ3y_bGgAdMGiwl9zoCP1NvMiRg7xVFStgdPJHr55MDVaNBac6a54Tol3d

ret = client.Close(fd426)
if(ret != 0) {
  panic("close failed")
}


fd427 := client.Open("/ZmYnI8Rd8S.txt", client.O_RDWR|client.O_CREATE)
if(fd427 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd427, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd427, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd424)
if(ret != 0) {
  panic("close failed")
}


fd428 := client.Open("/vlNW582dxT.txt", client.O_RDWR|client.O_CREATE)
if(fd428 < 0) {
  panic("open failed")
}


ret = client.Seek(fd422, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd427, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd427, []byte("4rS4solrePjzm5zXd1z9K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) 4rS4solrePjzm5zXd1z9K

buf, ret = client.Read(fd422, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (21) 4rS4solrePjzm5zXd1z9K

ret = client.Write(fd427, []byte("XIm1LYLLC3f8hth9mo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) 4rS4solrePjzm5zXd1z9KXIm1LYLLC3f8hth9mo
//fd state: (0) pYuU_M_qEC2pf

ret = client.Write(fd428, []byte("5xVfD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) 5xVfDM_qEC2pf
//fd state: (0) 

ret = client.Write(fd422, []byte("KQG36Lhjf6ibKiRCX7fsiSIwt25roNBhRfICTjauQsBQo9YBNnJVTbRg4P0a1UrOcbSc4FnHdFsw2b8HbMqp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) KQG36Lhjf6ibKiRCX7fsiSIwt25roNBhRfICTjauQsBQo9YBNnJVTbRg4P0a1UrOcbSc4FnHdFsw2b8HbMqp

ret = client.Seek(fd428, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


ret = client.Seek(fd427, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


fd429 := client.Open("/vlNW582dxT.txt", client.O_RDWR|client.O_CREATE)
if(fd429 < 0) {
  panic("open failed")
}


ret = client.Close(fd429)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd427, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "solrePjzm5zXd1z9KXIm1LYLLC3f") {
  panic("wrong data returned")
}

//fd state: (84) KQG36Lhjf6ibKiRCX7fsiSIwt25roNBhRfICTjauQsBQo9YBNnJVTbRg4P0a1UrOcbSc4FnHdFsw2b8HbMqp

ret = client.Write(fd422, []byte("v9kHreBA_oow90zgjW6K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) KQG36Lhjf6ibKiRCX7fsiSIwt25roNBhRfICTjauQsBQo9YBNnJVTbRg4P0a1UrOcbSc4FnHdFsw2b8HbMqpv9kHreBA_oow90zgjW6K

ret = client.Seek(fd427, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


ret = client.Close(fd428)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd422, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (29) 4rS4solrePjzm5zXd1z9KXIm1LYLLC3f8hth9mo

ret = client.Write(fd427, []byte("fNnmu1yy2fU1zht8DHrXx5kVMATZDZILxJIv8PEjiUrHXtXagSY0KpUJdv4lBYnkq0LRERmBzkF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) 4rS4solrePjzm5zXd1z9KXIm1LYLLfNnmu1yy2fU1zht8DHrXx5kVMATZDZILxJIv8PEjiUrHXtXagSY0KpUJdv4lBYnkq0LRERmBzkF

buf, ret = client.Read(fd427, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd427, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}

//fd state: (104) KQG36Lhjf6ibKiRCX7fsiSIwt25roNBhRfICTjauQsBQo9YBNnJVTbRg4P0a1UrOcbSc4FnHdFsw2b8HbMqpv9kHreBA_oow90zgjW6K

ret = client.Write(fd422, []byte("3BnxQXyijMnL3LnfYVY0rteQJeT0gxoP4UO4JGAuwnpJjWjVaJo4rSxUiWBs9vC27cGCosIVKeZTBfIX3NA7lVeruSA9E0A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (199) KQG36Lhjf6ibKiRCX7fsiSIwt25roNBhRfICTjauQsBQo9YBNnJVTbRg4P0a1UrOcbSc4FnHdFsw2b8HbMqpv9kHreBA_oow90zgjW6K3BnxQXyijMnL3LnfYVY0rteQJeT0gxoP4UO4JGAuwnpJjWjVaJo4rSxUiWBs9vC27cGCosIVKeZTBfIX3NA7lVeruSA9E0A

ret = client.Seek(fd427, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


fd430 := client.Open("/LZLmnQ7Drw.txt", client.O_RDWR|client.O_CREATE)
if(fd430 < 0) {
  panic("open failed")
}


ret = client.Seek(fd422, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Close(fd422)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd427)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd430)
if(ret != 0) {
  panic("close failed")
}


fd431 := client.Open("/bCLvtka1QB.txt", client.O_RDWR|client.O_CREATE)
if(fd431 < 0) {
  panic("open failed")
}


fd432 := client.Open("/LZcAyv_xCm.txt", client.O_RDWR|client.O_CREATE)
if(fd432 < 0) {
  panic("open failed")
}


fd433 := client.Open("/OpMYyHTi1Y.txt", client.O_RDWR|client.O_CREATE)
if(fd433 < 0) {
  panic("open failed")
}


ret = client.Seek(fd433, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd433)
if(ret != 0) {
  panic("close failed")
}


fd434 := client.Open("/QqwyHnv4d8.txt", client.O_RDWR|client.O_CREATE)
if(fd434 < 0) {
  panic("open failed")
}


ret = client.Close(fd431)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) fmpftJInZqnKnAARukBxvyJsMK92fY64w3ZykwDdmS7OxzccPIoNvZbiLctYYz1zjrVGOcQ6zt2pcUFT7VvzZcrSOhWKcWeKq

ret = client.Write(fd432, []byte("rGLuPZvpyFNB3YCoGRX_x8pFmSu_xUB09ZVKmJZctVjG8gRdw5Pev_2nLmv0ZZ3x2D5065s5a3re7xQVChPGQwmhEw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) rGLuPZvpyFNB3YCoGRX_x8pFmSu_xUB09ZVKmJZctVjG8gRdw5Pev_2nLmv0ZZ3x2D5065s5a3re7xQVChPGQwmhEwWKcWeKq

buf, ret = client.Read(fd432, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) gBaxcGq7XrCpO0iHTk_gSlZCNTgpgHYpo6qPiR8xyf6drqezc9F6h813CuRh9KuYfJCkKS6DChHZ8N58mqKwaqoHXqKuatHsHpyfZwIWqEvFXdXy3BkXQpvNW2OLVD7gIcrE21rhjqy

ret = client.Write(fd434, []byte("HjNZzWL6W4XF3X_AAPJ8HqPW1sYWoyV2b9lFA7Rjv52HbjA6LPs0deByhWcESHutHnOBFiFR6xo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) HjNZzWL6W4XF3X_AAPJ8HqPW1sYWoyV2b9lFA7Rjv52HbjA6LPs0deByhWcESHutHnOBFiFR6xoZ8N58mqKwaqoHXqKuatHsHpyfZwIWqEvFXdXy3BkXQpvNW2OLVD7gIcrE21rhjqy

ret = client.Close(fd432)
if(ret != 0) {
  panic("close failed")
}


fd435 := client.Open("/UDXib7cZUL.txt", client.O_RDWR|client.O_CREATE)
if(fd435 < 0) {
  panic("open failed")
}


ret = client.Seek(fd435, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd435, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd435, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd435)
if(ret != 0) {
  panic("close failed")
}

//fd state: (75) HjNZzWL6W4XF3X_AAPJ8HqPW1sYWoyV2b9lFA7Rjv52HbjA6LPs0deByhWcESHutHnOBFiFR6xoZ8N58mqKwaqoHXqKuatHsHpyfZwIWqEvFXdXy3BkXQpvNW2OLVD7gIcrE21rhjqy

ret = client.Write(fd434, []byte("EReItOoh7NrlW1LTpUwlj1hPDn634ZmY9eSnc5vNpOMIjFNpO6JzLqToF0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) HjNZzWL6W4XF3X_AAPJ8HqPW1sYWoyV2b9lFA7Rjv52HbjA6LPs0deByhWcESHutHnOBFiFR6xoEReItOoh7NrlW1LTpUwlj1hPDn634ZmY9eSnc5vNpOMIjFNpO6JzLqToF01rhjqy

ret = client.Seek(fd434, 112, client.SEEK_SET)
if(ret != 112) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 112)
  panic("seek failed")
}


buf, ret = client.Read(fd434, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5vNpOMIjFNpO6JzLqToF01rhjqy") {
  panic("wrong data returned")
}


ret = client.Seek(fd434, 101, client.SEEK_SET)
if(ret != 101) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 101)
  panic("seek failed")
}

//fd state: (101) HjNZzWL6W4XF3X_AAPJ8HqPW1sYWoyV2b9lFA7Rjv52HbjA6LPs0deByhWcESHutHnOBFiFR6xoEReItOoh7NrlW1LTpUwlj1hPDn634ZmY9eSnc5vNpOMIjFNpO6JzLqToF01rhjqy

ret = client.Write(fd434, []byte("0NqLyWe3DX5TSpjeApla5UFhG9Z1AMSYGkAG4LEPU9LmSMoubfPkiQXODZHCU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (162) HjNZzWL6W4XF3X_AAPJ8HqPW1sYWoyV2b9lFA7Rjv52HbjA6LPs0deByhWcESHutHnOBFiFR6xoEReItOoh7NrlW1LTpUwlj1hPDn0NqLyWe3DX5TSpjeApla5UFhG9Z1AMSYGkAG4LEPU9LmSMoubfPkiQXODZHCU

ret = client.Close(fd434)
if(ret != 0) {
  panic("close failed")
}


fd436 := client.Open("/IAWaQyZGxP.txt", client.O_RDWR|client.O_CREATE)
if(fd436 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd436, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd436, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd437 := client.Open("/nYSzy33lDz.txt", client.O_RDWR|client.O_CREATE)
if(fd437 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd436, []byte("l32PuHlCBlmKQtFL6TcpxiHru2P1gqg6l0S6pUo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) l32PuHlCBlmKQtFL6TcpxiHru2P1gqg6l0S6pUo

buf, ret = client.Read(fd437, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4t") {
  panic("wrong data returned")
}


ret = client.Close(fd437)
if(ret != 0) {
  panic("close failed")
}


fd438 := client.Open("/CeZSrGh142.txt", client.O_RDWR|client.O_CREATE)
if(fd438 < 0) {
  panic("open failed")
}


ret = client.Seek(fd438, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd438)
if(ret != 0) {
  panic("close failed")
}

//fd state: (39) l32PuHlCBlmKQtFL6TcpxiHru2P1gqg6l0S6pUo

ret = client.Write(fd436, []byte("pYFnamAFkioAVrvzcGB_HLy9y2MyFc74OdDw5rS5txHUiUeKXvYgDDcixyydAf9niuSAumdo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) l32PuHlCBlmKQtFL6TcpxiHru2P1gqg6l0S6pUopYFnamAFkioAVrvzcGB_HLy9y2MyFc74OdDw5rS5txHUiUeKXvYgDDcixyydAf9niuSAumdo

ret = client.Seek(fd436, 98, client.SEEK_SET)
if(ret != 98) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 98)
  panic("seek failed")
}


buf, ret = client.Read(fd436, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "dAf9niuSAumdo") {
  panic("wrong data returned")
}


ret = client.Seek(fd436, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


ret = client.Seek(fd436, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


ret = client.Close(fd436)
if(ret != 0) {
  panic("close failed")
}


fd439 := client.Open("/8vaK9bbYr6.txt", client.O_RDWR|client.O_CREATE)
if(fd439 < 0) {
  panic("open failed")
}


ret = client.Close(fd439)
if(ret != 0) {
  panic("close failed")
}


fd440 := client.Open("/xUg22Cub2y.txt", client.O_RDWR|client.O_CREATE)
if(fd440 < 0) {
  panic("open failed")
}


ret = client.Close(fd440)
if(ret != 0) {
  panic("close failed")
}


fd441 := client.Open("/A2jhBeQbpL.txt", client.O_RDWR|client.O_CREATE)
if(fd441 < 0) {
  panic("open failed")
}


ret = client.Close(fd441)
if(ret != 0) {
  panic("close failed")
}


fd442 := client.Open("/IvVLkoieiN.txt", client.O_RDWR|client.O_CREATE)
if(fd442 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd442, []byte("ZYQo8Ej2o_LxWM4Ydv9mBi9rKnXLe0LkITnuM5_vtw51"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) ZYQo8Ej2o_LxWM4Ydv9mBi9rKnXLe0LkITnuM5_vtw51
//fd state: (44) ZYQo8Ej2o_LxWM4Ydv9mBi9rKnXLe0LkITnuM5_vtw51

ret = client.Write(fd442, []byte("NYoPlCHyGBP3PNipnwvpPm1i_i6S_Y2dubnqx5PbIwWTgy4UWE1mDwHM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) ZYQo8Ej2o_LxWM4Ydv9mBi9rKnXLe0LkITnuM5_vtw51NYoPlCHyGBP3PNipnwvpPm1i_i6S_Y2dubnqx5PbIwWTgy4UWE1mDwHM

fd443 := client.Open("/uUuh4BivWQ.txt", client.O_RDWR|client.O_CREATE)
if(fd443 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd442, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd442)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd443, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd443)
if(ret != 0) {
  panic("close failed")
}


fd444 := client.Open("/IlPekcd4f8.txt", client.O_RDWR|client.O_CREATE)
if(fd444 < 0) {
  panic("open failed")
}


ret = client.Seek(fd444, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


buf, ret = client.Read(fd444, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9k0dPL") {
  panic("wrong data returned")
}

//fd state: (10) hFUP9k0dPL

ret = client.Write(fd444, []byte("JKZMpbu5gLOuDZra0G"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) hFUP9k0dPLJKZMpbu5gLOuDZra0G

ret = client.Close(fd444)
if(ret != 0) {
  panic("close failed")
}


fd445 := client.Open("/PYTqf9ITeb.txt", client.O_RDWR|client.O_CREATE)
if(fd445 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd445, []byte("OQZ5l"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) OQZ5l

fd446 := client.Open("/wQOfDmuJ6Y.txt", client.O_RDWR|client.O_CREATE)
if(fd446 < 0) {
  panic("open failed")
}


ret = client.Seek(fd446, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd446, []byte("yK1oNlrkJGtcSEpZJjfF2HTlhusikdDBgRj_PzofjgRBjeA6X7BnIGB5MmSrIIy4wQdnJxfaiiHD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) yK1oNlrkJGtcSEpZJjfF2HTlhusikdDBgRj_PzofjgRBjeA6X7BnIGB5MmSrIIy4wQdnJxfaiiHD
//fd state: (5) OQZ5l

ret = client.Write(fd445, []byte("ZwUNJTY7NA_hMMjnEzsPE4p"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) OQZ5lZwUNJTY7NA_hMMjnEzsPE4p
//fd state: (76) yK1oNlrkJGtcSEpZJjfF2HTlhusikdDBgRj_PzofjgRBjeA6X7BnIGB5MmSrIIy4wQdnJxfaiiHD

ret = client.Write(fd446, []byte("Hrsky7R8WX4bNFNSRg2dMHTig_iuGYEdojZQUYKmkIAvZzKjsrnHk90WuA1COl0oueswApIP6AcROPq4Z2Nitl9p"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) yK1oNlrkJGtcSEpZJjfF2HTlhusikdDBgRj_PzofjgRBjeA6X7BnIGB5MmSrIIy4wQdnJxfaiiHDHrsky7R8WX4bNFNSRg2dMHTig_iuGYEdojZQUYKmkIAvZzKjsrnHk90WuA1COl0oueswApIP6AcROPq4Z2Nitl9p

ret = client.Seek(fd446, 109, client.SEEK_SET)
if(ret != 109) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 109)
  panic("seek failed")
}


ret = client.Seek(fd446, 86, client.SEEK_SET)
if(ret != 86) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 86)
  panic("seek failed")
}


fd447 := client.Open("/Ejqw2f8dM3.txt", client.O_RDWR|client.O_CREATE)
if(fd447 < 0) {
  panic("open failed")
}


fd448 := client.Open("/Kbe_5esbc2.txt", client.O_RDWR|client.O_CREATE)
if(fd448 < 0) {
  panic("open failed")
}


fd449 := client.Open("/JPr0cEv57D.txt", client.O_RDWR|client.O_CREATE)
if(fd449 < 0) {
  panic("open failed")
}


fd450 := client.Open("/W8T68Q5TiM.txt", client.O_RDWR|client.O_CREATE)
if(fd450 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd445, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd450, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd449)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd447, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "kMi") {
  panic("wrong data returned")
}


fd451 := client.Open("/2ihr4LfQFk.txt", client.O_RDWR|client.O_CREATE)
if(fd451 < 0) {
  panic("open failed")
}


fd452 := client.Open("/NY1rdxn0ys.txt", client.O_RDWR|client.O_CREATE)
if(fd452 < 0) {
  panic("open failed")
}


ret = client.Close(fd452)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd447, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9tUFXp3GFmTPHZ90biPHIkWwWOOhbPP_WglZQt1dDej1e") {
  panic("wrong data returned")
}


ret = client.Close(fd447)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd445, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd450)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd448, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd453 := client.Open("/ojkJUB3zCr.txt", client.O_RDWR|client.O_CREATE)
if(fd453 < 0) {
  panic("open failed")
}


fd454 := client.Open("/iBtizDgltf.txt", client.O_RDWR|client.O_CREATE)
if(fd454 < 0) {
  panic("open failed")
}


ret = client.Seek(fd448, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd455 := client.Open("/S4kyEU_r6C.txt", client.O_RDWR|client.O_CREATE)
if(fd455 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd446, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4bNFNSRg2dMHTig_iuGYEdojZQUYKmkIAvZzKjsrnHk90WuA1COl0oueswApIP6AcROPq4Z2Nitl") {
  panic("wrong data returned")
}


fd456 := client.Open("/HzxAhdTV6L.txt", client.O_RDWR|client.O_CREATE)
if(fd456 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd454, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd445)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd453, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd454)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd446, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}

//fd state: (15) yK1oNlrkJGtcSEpZJjfF2HTlhusikdDBgRj_PzofjgRBjeA6X7BnIGB5MmSrIIy4wQdnJxfaiiHDHrsky7R8WX4bNFNSRg2dMHTig_iuGYEdojZQUYKmkIAvZzKjsrnHk90WuA1COl0oueswApIP6AcROPq4Z2Nitl9p

ret = client.Write(fd446, []byte("64z8PQTdaYc1VVrQjc2EMYoJIx0ENukRn9qIyURlmnaWx6gcNBeKt7GF_uH3LQBbyzpyCPiSBIr0lcOJglB3j_Zr8HnAboiF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) yK1oNlrkJGtcSEp64z8PQTdaYc1VVrQjc2EMYoJIx0ENukRn9qIyURlmnaWx6gcNBeKt7GF_uH3LQBbyzpyCPiSBIr0lcOJglB3j_Zr8HnAboiFQUYKmkIAvZzKjsrnHk90WuA1COl0oueswApIP6AcROPq4Z2Nitl9p

ret = client.Seek(fd456, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd457 := client.Open("/dWsEcztvgd.txt", client.O_RDWR|client.O_CREATE)
if(fd457 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd451, 19)
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


fd458 := client.Open("/wy3cufuk8j.txt", client.O_RDWR|client.O_CREATE)
if(fd458 < 0) {
  panic("open failed")
}


fd459 := client.Open("/UgWV3nMHHM.txt", client.O_RDWR|client.O_CREATE)
if(fd459 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd458, []byte("OsH9zwMzTaftxN0w35K11se2_c6OrGF9G25KoVgNFvF3pCj_QTmEwcxafzx3ZxT21hkxIyGxWr7yNdCpnySgurrOV_2xI6Cf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) OsH9zwMzTaftxN0w35K11se2_c6OrGF9G25KoVgNFvF3pCj_QTmEwcxafzx3ZxT21hkxIyGxWr7yNdCpnySgurrOV_2xI6Cf

ret = client.Seek(fd456, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd459)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd457, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}

//fd state: (0) pDYhk49mN

ret = client.Write(fd455, []byte("hyQ3IxuGte64OvndFOSNKV8JKpJZIjIA5GzDcAwtEdNEeu19vq89P"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) hyQ3IxuGte64OvndFOSNKV8JKpJZIjIA5GzDcAwtEdNEeu19vq89P
//fd state: (0) 

ret = client.Write(fd448, []byte("0SWYabsv5WW0JC8rUQHiar9oPF9r72UUY_K4WLy9NNSR69x7uV0i9H70"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) 0SWYabsv5WW0JC8rUQHiar9oPF9r72UUY_K4WLy9NNSR69x7uV0i9H70

buf, ret = client.Read(fd457, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7fqmKT") {
  panic("wrong data returned")
}


ret = client.Seek(fd455, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


fd460 := client.Open("/iTE9QULg8Q.txt", client.O_RDWR|client.O_CREATE)
if(fd460 < 0) {
  panic("open failed")
}

//fd state: (56) 0SWYabsv5WW0JC8rUQHiar9oPF9r72UUY_K4WLy9NNSR69x7uV0i9H70

ret = client.Write(fd448, []byte("XgydubhCXYP26Afq0FyNk8ZKOF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) 0SWYabsv5WW0JC8rUQHiar9oPF9r72UUY_K4WLy9NNSR69x7uV0i9H70XgydubhCXYP26Afq0FyNk8ZKOF

ret = client.Seek(fd448, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


ret = client.Seek(fd455, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Seek(fd460, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd458, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd458)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd457, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd453, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd455, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


ret = client.Seek(fd453, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd461 := client.Open("/ikZPJ_xCMj.txt", client.O_RDWR|client.O_CREATE)
if(fd461 < 0) {
  panic("open failed")
}


ret = client.Close(fd451)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd448)
if(ret != 0) {
  panic("close failed")
}


fd462 := client.Open("/w8tEkfoHmF.txt", client.O_RDWR|client.O_CREATE)
if(fd462 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd457, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd455, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jIA5GzDcAwtEdNEeu19vq89P") {
  panic("wrong data returned")
}


ret = client.Close(fd462)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd453, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd463 := client.Open("/2yZkRq9LHH.txt", client.O_RDWR|client.O_CREATE)
if(fd463 < 0) {
  panic("open failed")
}


ret = client.Seek(fd460, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd463, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd463, []byte("AS9HXj7dFx0mo7CBrPs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) AS9HXj7dFx0mo7CBrPs

ret = client.Seek(fd461, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd463, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd464 := client.Open("/rHbNRD2SCH.txt", client.O_RDWR|client.O_CREATE)
if(fd464 < 0) {
  panic("open failed")
}


ret = client.Close(fd463)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd461)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd460)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd453)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd456, []byte("UHdY4HQBXeeS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) UHdY4HQBXeeS

ret = client.Close(fd455)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd464)
if(ret != 0) {
  panic("close failed")
}


fd465 := client.Open("/icy2T1pZHr.txt", client.O_RDWR|client.O_CREATE)
if(fd465 < 0) {
  panic("open failed")
}


fd466 := client.Open("/UfA4o9VB69.txt", client.O_RDWR|client.O_CREATE)
if(fd466 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd466, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3i") {
  panic("wrong data returned")
}

//fd state: (12) UHdY4HQBXeeS

ret = client.Write(fd456, []byte("gdxKN1cFW3N7JYYl6bFHw8MHbmVzHwvwIEJtHTKOfpGn6sfNwDzp_4SHz6otOX7PbYwjYKW8jjc4cvjq5BYvtK4ieWg5s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) UHdY4HQBXeeSgdxKN1cFW3N7JYYl6bFHw8MHbmVzHwvwIEJtHTKOfpGn6sfNwDzp_4SHz6otOX7PbYwjYKW8jjc4cvjq5BYvtK4ieWg5s

fd467 := client.Open("/aODta09XGB.txt", client.O_RDWR|client.O_CREATE)
if(fd467 < 0) {
  panic("open failed")
}


fd468 := client.Open("/WujY5LPXRr.txt", client.O_RDWR|client.O_CREATE)
if(fd468 < 0) {
  panic("open failed")
}


ret = client.Close(fd468)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd456)
if(ret != 0) {
  panic("close failed")
}


fd469 := client.Open("/guKVxEdClf.txt", client.O_RDWR|client.O_CREATE)
if(fd469 < 0) {
  panic("open failed")
}


fd470 := client.Open("/4i_mO3QRQA.txt", client.O_RDWR|client.O_CREATE)
if(fd470 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd466, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "tdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mq") {
  panic("wrong data returned")
}


ret = client.Seek(fd469, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (127) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqPcJwrR61XWVBg2eUTlINt1b7YAtiDtMwnw

ret = client.Write(fd466, []byte("OckyeiocyiPaiKVse"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqOckyeiocyiPaiKVselINt1b7YAtiDtMwnw
//fd state: (144) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqOckyeiocyiPaiKVselINt1b7YAtiDtMwnw

ret = client.Write(fd466, []byte("UKbt4mIshGRGZUoiBRWJOTy6Tp3qKIdA4GiY_I97V4CEDFokTLPNioxo7IMa4MFTeA8u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (212) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqOckyeiocyiPaiKVseUKbt4mIshGRGZUoiBRWJOTy6Tp3qKIdA4GiY_I97V4CEDFokTLPNioxo7IMa4MFTeA8u

fd471 := client.Open("/MVWR6bEOKx.txt", client.O_RDWR|client.O_CREATE)
if(fd471 < 0) {
  panic("open failed")
}


fd472 := client.Open("/EXgQzL0J8u.txt", client.O_RDWR|client.O_CREATE)
if(fd472 < 0) {
  panic("open failed")
}


fd473 := client.Open("/dSml2M8zNP.txt", client.O_RDWR|client.O_CREATE)
if(fd473 < 0) {
  panic("open failed")
}


fd474 := client.Open("/9m1NnDd6_4.txt", client.O_RDWR|client.O_CREATE)
if(fd474 < 0) {
  panic("open failed")
}


ret = client.Close(fd470)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd473, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd475 := client.Open("/345sBFjFS5.txt", client.O_RDWR|client.O_CREATE)
if(fd475 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd474, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) yIndr06_InYbZguoemjp3Rn6fbapt7c1yyKXTgLEWluD_P1Jdg4BkLjdIyCdk9

ret = client.Write(fd465, []byte("OEZiVrsakkY5PoMs_PXK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) OEZiVrsakkY5PoMs_PXK3Rn6fbapt7c1yyKXTgLEWluD_P1Jdg4BkLjdIyCdk9

fd476 := client.Open("/HYIrFvOUl2.txt", client.O_RDWR|client.O_CREATE)
if(fd476 < 0) {
  panic("open failed")
}

//fd state: (20) OEZiVrsakkY5PoMs_PXK3Rn6fbapt7c1yyKXTgLEWluD_P1Jdg4BkLjdIyCdk9

ret = client.Write(fd465, []byte("mjmJiUjnGKMRks"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) OEZiVrsakkY5PoMs_PXKmjmJiUjnGKMRksKXTgLEWluD_P1Jdg4BkLjdIyCdk9

ret = client.Close(fd475)
if(ret != 0) {
  panic("close failed")
}

//fd state: (64) vl6Oq4X8hLK5cnUgqlEwB6hJ5P1YAoLQQTqnKNM32JzCGoW_TG5d5tGRRH7fqmKT

ret = client.Write(fd457, []byte("m98oTTLPeGcoDJEkVcnNCWhWn7b7931Cw5XYRugvv2hI6Jq0hG89bWlVSbgXbg6txE02"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) vl6Oq4X8hLK5cnUgqlEwB6hJ5P1YAoLQQTqnKNM32JzCGoW_TG5d5tGRRH7fqmKTm98oTTLPeGcoDJEkVcnNCWhWn7b7931Cw5XYRugvv2hI6Jq0hG89bWlVSbgXbg6txE02

ret = client.Close(fd469)
if(ret != 0) {
  panic("close failed")
}


fd477 := client.Open("/345sBFjFS5.txt", client.O_RDWR|client.O_CREATE)
if(fd477 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd465, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "K") {
  panic("wrong data returned")
}


fd478 := client.Open("/aGkYYZpZXG.txt", client.O_RDWR|client.O_CREATE)
if(fd478 < 0) {
  panic("open failed")
}


fd479 := client.Open("/EXgQzL0J8u.txt", client.O_RDWR|client.O_CREATE)
if(fd479 < 0) {
  panic("open failed")
}


fd480 := client.Open("/DZlV8Pcptl.txt", client.O_RDWR|client.O_CREATE)
if(fd480 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd457, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd481 := client.Open("/6Okbu3K6Za.txt", client.O_RDWR|client.O_CREATE)
if(fd481 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd478, []byte("9yO5mALiRfrI2UvYk9oFLXZ0iaOkLskbpB9fZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) 9yO5mALiRfrI2UvYk9oFLXZ0iaOkLskbpB9fZ

ret = client.Seek(fd457, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}


ret = client.Close(fd479)
if(ret != 0) {
  panic("close failed")
}


fd482 := client.Open("/tRieOJpZ5n.txt", client.O_RDWR|client.O_CREATE)
if(fd482 < 0) {
  panic("open failed")
}


ret = client.Seek(fd473, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd466, 168, client.SEEK_SET)
if(ret != 168) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 168)
  panic("seek failed")
}

//fd state: (168) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqOckyeiocyiPaiKVseUKbt4mIshGRGZUoiBRWJOTy6Tp3qKIdA4GiY_I97V4CEDFokTLPNioxo7IMa4MFTeA8u

ret = client.Write(fd466, []byte("aEGw3fSDYDcJTosOLNssw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (189) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqOckyeiocyiPaiKVseUKbt4mIshGRGZUoiBRWJOTy6aEGw3fSDYDcJTosOLNsswFokTLPNioxo7IMa4MFTeA8u
//fd state: (189) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqOckyeiocyiPaiKVseUKbt4mIshGRGZUoiBRWJOTy6aEGw3fSDYDcJTosOLNsswFokTLPNioxo7IMa4MFTeA8u

ret = client.Write(fd466, []byte("w2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (191) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqOckyeiocyiPaiKVseUKbt4mIshGRGZUoiBRWJOTy6aEGw3fSDYDcJTosOLNssww2kTLPNioxo7IMa4MFTeA8u

ret = client.Close(fd478)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd467, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd476, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd457, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


buf, ret = client.Read(fd482, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd483 := client.Open("/b8YxyJRioX.txt", client.O_RDWR|client.O_CREATE)
if(fd483 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd477, []byte("hSqf2USloqhmObReutEY2uxNjiqME5CDBsN4kyZjow254SYsrI9KvFC7il4GMX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) hSqf2USloqhmObReutEY2uxNjiqME5CDBsN4kyZjow254SYsrI9KvFC7il4GMX

fd484 := client.Open("/W14ZY4nXnm.txt", client.O_RDWR|client.O_CREATE)
if(fd484 < 0) {
  panic("open failed")
}


ret = client.Seek(fd476, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd477, 61, client.SEEK_SET)
if(ret != 61) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 61)
  panic("seek failed")
}


buf, ret = client.Read(fd476, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd482, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd485 := client.Open("/z3E8FPWTfv.txt", client.O_RDWR|client.O_CREATE)
if(fd485 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd467, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd485, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd486 := client.Open("/w8tEkfoHmF.txt", client.O_RDWR|client.O_CREATE)
if(fd486 < 0) {
  panic("open failed")
}


fd487 := client.Open("/2JrRL5XTCk.txt", client.O_RDWR|client.O_CREATE)
if(fd487 < 0) {
  panic("open failed")
}


ret = client.Seek(fd474, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd487, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd472, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd485, []byte("6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RyIIurIv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RyIIurIv

fd488 := client.Open("/D3Fa1l3zaJ.txt", client.O_RDWR|client.O_CREATE)
if(fd488 < 0) {
  panic("open failed")
}


ret = client.Seek(fd476, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd476, []byte("AF4OWrjFGnNAfKHCoIDTXnb6_IGtjrLOISH7nHziE1_1JZT2enOeynZKvzxWvyjiecNHSM7mYyo9RL46t27KnX6lMc1nD5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) AF4OWrjFGnNAfKHCoIDTXnb6_IGtjrLOISH7nHziE1_1JZT2enOeynZKvzxWvyjiecNHSM7mYyo9RL46t27KnX6lMc1nD5
//fd state: (94) AF4OWrjFGnNAfKHCoIDTXnb6_IGtjrLOISH7nHziE1_1JZT2enOeynZKvzxWvyjiecNHSM7mYyo9RL46t27KnX6lMc1nD5

ret = client.Write(fd476, []byte("J3DffhZCX8lgsH3pyrjsot3o8Z6WR2M9eSfz9lSzysfasAK4XcrO5u9L4RLqEiqX0xSfbpmiKusghHGI_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (175) AF4OWrjFGnNAfKHCoIDTXnb6_IGtjrLOISH7nHziE1_1JZT2enOeynZKvzxWvyjiecNHSM7mYyo9RL46t27KnX6lMc1nD5J3DffhZCX8lgsH3pyrjsot3o8Z6WR2M9eSfz9lSzysfasAK4XcrO5u9L4RLqEiqX0xSfbpmiKusghHGI_

fd489 := client.Open("/zZt3zhX_S8.txt", client.O_RDWR|client.O_CREATE)
if(fd489 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd467, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (1) kGJEtE_LReqao0aOZtk2hl3geUWlgDTIo_8TAWvSmRVynNxasTH2atoRA_C2Y1ISabFhD762dVfmIE9

ret = client.Write(fd487, []byte("gZHYX2wiBt2mW26NM5k1pb6iOm0p"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) kgZHYX2wiBt2mW26NM5k1pb6iOm0pDTIo_8TAWvSmRVynNxasTH2atoRA_C2Y1ISabFhD762dVfmIE9

fd490 := client.Open("/Zh1UizOUpM.txt", client.O_RDWR|client.O_CREATE)
if(fd490 < 0) {
  panic("open failed")
}


ret = client.Close(fd457)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd481)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd476, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


buf, ret = client.Read(fd487, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "DTIo_8TAWvSmRVynNxasTH2atoRA_C2Y1ISabFhD762dVfmIE9") {
  panic("wrong data returned")
}


ret = client.Seek(fd483, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Seek(fd485, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


buf, ret = client.Read(fd471, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd484)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd473, []byte("sX0ARl5lBkDPcCSSryX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) sX0ARl5lBkDPcCSSryX

ret = client.Seek(fd485, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd467, []byte("TMMEEltoHnEdOENtHTGP2KCxHViCfJDO9l6KlC8hiIweJLTyN6BJeJCFY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) TMMEEltoHnEdOENtHTGP2KCxHViCfJDO9l6KlC8hiIweJLTyN6BJeJCFY

buf, ret = client.Read(fd486, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd491 := client.Open("/ZR9PqbIJDr.txt", client.O_RDWR|client.O_CREATE)
if(fd491 < 0) {
  panic("open failed")
}


ret = client.Seek(fd480, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}

//fd state: (35) zKTtVuTTNf84wOouvc1OeH5XENwa9QVidq4DreCwLn36qOIq7xCpWfBt5R0JYQ7TTyGFBFPvfcNc_zqBcGL5Zuk3sbG

ret = client.Write(fd483, []byte("n1AnpvqBi4MdqR37Y_FmQRSvccB_HeX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) zKTtVuTTNf84wOouvc1OeH5XENwa9QVidq4n1AnpvqBi4MdqR37Y_FmQRSvccB_HeXGFBFPvfcNc_zqBcGL5Zuk3sbG

ret = client.Close(fd465)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd489, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rxjTd1b8GxJWIATdOlyCe5zzmTcIvGi2VShRN6DnU2UG7C4slgN60eVK2B7J1") {
  panic("wrong data returned")
}


ret = client.Close(fd483)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd491, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd488, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NVr_siCEPrkO_EhJ838Uv") {
  panic("wrong data returned")
}

//fd state: (79) kgZHYX2wiBt2mW26NM5k1pb6iOm0pDTIo_8TAWvSmRVynNxasTH2atoRA_C2Y1ISabFhD762dVfmIE9

ret = client.Write(fd487, []byte("nNo7Pk_BN3SJmd56jPLeR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) kgZHYX2wiBt2mW26NM5k1pb6iOm0pDTIo_8TAWvSmRVynNxasTH2atoRA_C2Y1ISabFhD762dVfmIE9nNo7Pk_BN3SJmd56jPLeR

ret = client.Close(fd482)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd466, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "kTLPNioxo7IMa4MFTeA8u") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd471, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd492 := client.Open("/V9mqXgOsu9.txt", client.O_RDWR|client.O_CREATE)
if(fd492 < 0) {
  panic("open failed")
}


fd493 := client.Open("/ZKNWUF4TFV.txt", client.O_RDWR|client.O_CREATE)
if(fd493 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd471, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd492, []byte("esHVXcuhLV9e8jyx3wcgmYRs5iE7CS1lgczm8qJufP1B0HaWzPQxXYqjyZuk0yp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) esHVXcuhLV9e8jyx3wcgmYRs5iE7CS1lgczm8qJufP1B0HaWzPQxXYqjyZuk0yp

ret = client.Seek(fd486, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd491, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd491, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd477, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


ret = client.Seek(fd467, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


fd494 := client.Open("/YrkZYzKQbd.txt", client.O_RDWR|client.O_CREATE)
if(fd494 < 0) {
  panic("open failed")
}

//fd state: (212) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqOckyeiocyiPaiKVseUKbt4mIshGRGZUoiBRWJOTy6aEGw3fSDYDcJTosOLNssww2kTLPNioxo7IMa4MFTeA8u

ret = client.Write(fd466, []byte("q8G2BM7AG8freSO0uSKnUdpxz3dUIEJGouIklHynNG_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (255) KUrB3y3xAkycY9H0EIGdvSDbjz0JD2coHtBAhKfa7ou4LIFKl9sO8BmgycaGeO3Ut3itdoATac9pkLksfnX3kt6rGHkkjU5unyGcoVrTHdqHI9lanbdUW4BABvxx8mqOckyeiocyiPaiKVseUKbt4mIshGRGZUoiBRWJOTy6aEGw3fSDYDcJTosOLNssww2kTLPNioxo7IMa4MFTeA8uq8G2BM7AG8freSO0uSKnUdpxz3dUIEJGouIklHynNG_

fd495 := client.Open("/Ex7w4vRt6O.txt", client.O_RDWR|client.O_CREATE)
if(fd495 < 0) {
  panic("open failed")
}


fd496 := client.Open("/XdPxjeqcPT.txt", client.O_RDWR|client.O_CREATE)
if(fd496 < 0) {
  panic("open failed")
}


ret = client.Close(fd476)
if(ret != 0) {
  panic("close failed")
}


fd497 := client.Open("/4VXKthGmF8.txt", client.O_RDWR|client.O_CREATE)
if(fd497 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd471, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd477, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


buf, ret = client.Read(fd467, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ViCfJDO9l6KlC8hiIweJLTyN6BJeJC") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd480, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "B7kNu5wDuSpM4fXnAmrEp30e9FN8qTy2zotFFdff") {
  panic("wrong data returned")
}


ret = client.Close(fd497)
if(ret != 0) {
  panic("close failed")
}

//fd state: (80) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RyIIurIv

ret = client.Write(fd485, []byte("ZbhSkuj8n_eHn6LL1ERsnP0X629PAyFD4T7m8TcRmTyzK2oIuXbbzF0Y7Af8pkWZqIoidMat_b1wKHcdHfuAlgM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (167) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RZbhSkuj8n_eHn6LL1ERsnP0X629PAyFD4T7m8TcRmTyzK2oIuXbbzF0Y7Af8pkWZqIoidMat_b1wKHcdHfuAlgM

buf, ret = client.Read(fd477, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "il4GMX") {
  panic("wrong data returned")
}


ret = client.Seek(fd473, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


buf, ret = client.Read(fd474, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd473)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd491, []byte("yjE5wBa74cQkosD9jOY8SATUik49w13zhTgxAl1sbPX_Denac8HRIjpB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) yjE5wBa74cQkosD9jOY8SATUik49w13zhTgxAl1sbPX_Denac8HRIjpB
//fd state: (55) TMMEEltoHnEdOENtHTGP2KCxHViCfJDO9l6KlC8hiIweJLTyN6BJeJCFY

ret = client.Write(fd467, []byte("gXfvkpTCLBurgsdz5U_QgP3I8SeiJA18vE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) TMMEEltoHnEdOENtHTGP2KCxHViCfJDO9l6KlC8hiIweJLTyN6BJeJCgXfvkpTCLBurgsdz5U_QgP3I8SeiJA18vE

buf, ret = client.Read(fd490, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd498 := client.Open("/lK1Ra888uV.txt", client.O_RDWR|client.O_CREATE)
if(fd498 < 0) {
  panic("open failed")
}


ret = client.Close(fd472)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) KUk78xEbHsbpXuOnCESKHFxHCmRvB2DkEN06xeAhcgyLZn59htkPOou8E4HHSulBcrVkxx42FKlnhRgScRDkU7LqQFWCqizmit2g8AAB7e8ekEEJSnM

ret = client.Write(fd496, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) KUk78xEbHsbpXuOnCESKHFxHCmRvB2DkEN06xeAhcgyLZn59htkPOou8E4HHSulBcrVkxx42FKlnhRgScRDkU7LqQFWCqizmit2g8AAB7e8ekEEJSnM

buf, ret = client.Read(fd491, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd498, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) KUk78xEbHsbpXuOnCESKHFxHCmRvB2DkEN06xeAhcgyLZn59htkPOou8E4HHSulBcrVkxx42FKlnhRgScRDkU7LqQFWCqizmit2g8AAB7e8ekEEJSnM

ret = client.Write(fd496, []byte("0LCoEMsusWBH3NjYh9CJGZSmser"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) 0LCoEMsusWBH3NjYh9CJGZSmservB2DkEN06xeAhcgyLZn59htkPOou8E4HHSulBcrVkxx42FKlnhRgScRDkU7LqQFWCqizmit2g8AAB7e8ekEEJSnM

fd499 := client.Open("/YcfJ8uk2Wm.txt", client.O_RDWR|client.O_CREATE)
if(fd499 < 0) {
  panic("open failed")
}


fd500 := client.Open("/4FdZoYlgde.txt", client.O_RDWR|client.O_CREATE)
if(fd500 < 0) {
  panic("open failed")
}


fd501 := client.Open("/YcfJ8uk2Wm.txt", client.O_RDWR|client.O_CREATE)
if(fd501 < 0) {
  panic("open failed")
}


fd502 := client.Open("/Ju0L1mzSEa.txt", client.O_RDWR|client.O_CREATE)
if(fd502 < 0) {
  panic("open failed")
}

//fd state: (61) rxjTd1b8GxJWIATdOlyCe5zzmTcIvGi2VShRN6DnU2UG7C4slgN60eVK2B7J1

ret = client.Write(fd489, []byte("_k9aPcS9LPj2UofR15PffbQvN8YA74NulVMGniS8IIxYD_iBdfr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (112) rxjTd1b8GxJWIATdOlyCe5zzmTcIvGi2VShRN6DnU2UG7C4slgN60eVK2B7J1_k9aPcS9LPj2UofR15PffbQvN8YA74NulVMGniS8IIxYD_iBdfr

ret = client.Seek(fd477, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd502, []byte("LHBmCmBXD1LtO9RULKBxJt9NrffqHR5MMrgcydKKJhYPUu04Nowkkkn51si"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) LHBmCmBXD1LtO9RULKBxJt9NrffqHR5MMrgcydKKJhYPUu04Nowkkkn51si

buf, ret = client.Read(fd491, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (63) esHVXcuhLV9e8jyx3wcgmYRs5iE7CS1lgczm8qJufP1B0HaWzPQxXYqjyZuk0yp

ret = client.Write(fd492, []byte("w7CTteuB9OBiB85PjDkBNELRTKV6WOnOSzAxmup1G8g7MvkQTfOVgWH2GhiIWTcmX7nZP6U_ay2smLWpaAGJi0ZmprsroNxRCZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (161) esHVXcuhLV9e8jyx3wcgmYRs5iE7CS1lgczm8qJufP1B0HaWzPQxXYqjyZuk0ypw7CTteuB9OBiB85PjDkBNELRTKV6WOnOSzAxmup1G8g7MvkQTfOVgWH2GhiIWTcmX7nZP6U_ay2smLWpaAGJi0ZmprsroNxRCZ

ret = client.Close(fd480)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd477, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


ret = client.Seek(fd467, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


fd503 := client.Open("/Zh1UizOUpM.txt", client.O_RDWR|client.O_CREATE)
if(fd503 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd503, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd504 := client.Open("/xdqG3xebi1.txt", client.O_RDWR|client.O_CREATE)
if(fd504 < 0) {
  panic("open failed")
}


ret = client.Seek(fd499, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd505 := client.Open("/7OdNvjR3av.txt", client.O_RDWR|client.O_CREATE)
if(fd505 < 0) {
  panic("open failed")
}


ret = client.Seek(fd498, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd495, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GqLX5nqhbOhPpdK21wGcDfQ9PcSxjXbCATzEVIjON9F2UPG8yP8XglBZEFePhzyobJCR2TYJjoqWLOz1x_M") {
  panic("wrong data returned")
}


ret = client.Close(fd466)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd498, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd471)
if(ret != 0) {
  panic("close failed")
}


fd506 := client.Open("/UDXib7cZUL.txt", client.O_RDWR|client.O_CREATE)
if(fd506 < 0) {
  panic("open failed")
}


ret = client.Seek(fd490, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd504, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd507 := client.Open("/o8e4tZKDW0.txt", client.O_RDWR|client.O_CREATE)
if(fd507 < 0) {
  panic("open failed")
}


ret = client.Close(fd491)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd504, []byte("GGlP5l_tCcwVsnzdvCntjP6CJJBnMPZqBbgs0JRf8gqkU440xEdEehD503UyrNT4ztSgOrX0UbJp3GlFVux"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) GGlP5l_tCcwVsnzdvCntjP6CJJBnMPZqBbgs0JRf8gqkU440xEdEehD503UyrNT4ztSgOrX0UbJp3GlFVux

ret = client.Close(fd493)
if(ret != 0) {
  panic("close failed")
}


fd508 := client.Open("/dWsEcztvgd.txt", client.O_RDWR|client.O_CREATE)
if(fd508 < 0) {
  panic("open failed")
}


ret = client.Close(fd487)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd500, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd474)
if(ret != 0) {
  panic("close failed")
}


fd509 := client.Open("/1mBp5jL7FX.txt", client.O_RDWR|client.O_CREATE)
if(fd509 < 0) {
  panic("open failed")
}


ret = client.Seek(fd502, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd500, []byte("WgSsuEBM_prV6hAuMgGeSN6rd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) WgSsuEBM_prV6hAuMgGeSN6rd

ret = client.Close(fd504)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd502, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


buf, ret = client.Read(fd508, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vl6Oq4X8hLK5cnUgqlEwB6hJ5P1YAoLQQTqnKNM32JzCGoW_TG5d5tGRRH7fqmKTm98oTTLPeGcoDJEkVcn") {
  panic("wrong data returned")
}

//fd state: (0) MkcNBURxNNkLgBY1dctT6kjUqXFF0y0YVJgJ1t15bnstKHkMStM6A2MLafc5d1ogSaA55uV4MRvDoFvLapWhJBacA1DycNL7s

ret = client.Write(fd505, []byte("8GvXi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) 8GvXiURxNNkLgBY1dctT6kjUqXFF0y0YVJgJ1t15bnstKHkMStM6A2MLafc5d1ogSaA55uV4MRvDoFvLapWhJBacA1DycNL7s

buf, ret = client.Read(fd506, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd467, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JeJCgXfvkpTCLBurgsdz5U_QgP3I8SeiJA1") {
  panic("wrong data returned")
}


ret = client.Close(fd495)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd499, []byte("YUwKleTtMubZgees0J_4HX7U1wtSjc0v_Xfynd3jMFXLgcwz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) YUwKleTtMubZgees0J_4HX7U1wtSjc0v_Xfynd3jMFXLgcwz

ret = client.Close(fd499)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd467)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd488)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd496)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd490, []byte("m0jLrbYPKisg9MdluGLr4mYaJvfJvvplJGVHpq3R7fKHdcodbKQ4S8mEHSW1Rjl4P6jlWjCT94O0HrFS8dA6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) m0jLrbYPKisg9MdluGLr4mYaJvfJvvplJGVHpq3R7fKHdcodbKQ4S8mEHSW1Rjl4P6jlWjCT94O0HrFS8dA6
//fd state: (167) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RZbhSkuj8n_eHn6LL1ERsnP0X629PAyFD4T7m8TcRmTyzK2oIuXbbzF0Y7Af8pkWZqIoidMat_b1wKHcdHfuAlgM

ret = client.Write(fd485, []byte("Z95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9my"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (228) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RZbhSkuj8n_eHn6LL1ERsnP0X629PAyFD4T7m8TcRmTyzK2oIuXbbzF0Y7Af8pkWZqIoidMat_b1wKHcdHfuAlgMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9my

buf, ret = client.Read(fd494, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd507)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd486, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd500, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


buf, ret = client.Read(fd508, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NCWhWn7b7931Cw5XYRugvv2hI6Jq0hG89bWlVSbgXbg6txE02") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd500, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AuMgGeSN6rd") {
  panic("wrong data returned")
}


ret = client.Seek(fd492, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


ret = client.Close(fd477)
if(ret != 0) {
  panic("close failed")
}


fd510 := client.Open("/1fAm1cEKrd.txt", client.O_RDWR|client.O_CREATE)
if(fd510 < 0) {
  panic("open failed")
}

//fd state: (52) LHBmCmBXD1LtO9RULKBxJt9NrffqHR5MMrgcydKKJhYPUu04Nowkkkn51si

ret = client.Write(fd502, []byte("9OnVzjNspvfIO5tJnLmdQcZLhZjasOuIlrzYPAdBsbIrnUg0RC0mIb8iK3A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) LHBmCmBXD1LtO9RULKBxJt9NrffqHR5MMrgcydKKJhYPUu04Nowk9OnVzjNspvfIO5tJnLmdQcZLhZjasOuIlrzYPAdBsbIrnUg0RC0mIb8iK3A

fd511 := client.Open("/qlEFbgz5IC.txt", client.O_RDWR|client.O_CREATE)
if(fd511 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd486, []byte("Nj7ZNp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) Nj7ZNp

ret = client.Seek(fd489, 101, client.SEEK_SET)
if(ret != 101) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 101)
  panic("seek failed")
}


ret = client.Close(fd486)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd494, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd505, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "URxNNkLgBY1dctT6kjUq") {
  panic("wrong data returned")
}


ret = client.Close(fd489)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd508, 82, client.SEEK_SET)
if(ret != 82) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 82)
  panic("seek failed")
}


fd512 := client.Open("/_yuDdDnmwK.txt", client.O_RDWR|client.O_CREATE)
if(fd512 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd509, []byte("uGKwd0isIuLIhmgnyRt78py7Q5YfJ0axKU60E25ZBn1DLoBMJgqBI062slYgYcWA57xg8sBewAQ0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) uGKwd0isIuLIhmgnyRt78py7Q5YfJ0axKU60E25ZBn1DLoBMJgqBI062slYgYcWA57xg8sBewAQ0

ret = client.Seek(fd494, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd498, []byte("QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2aHU23k1GEMDCUUCedbIkybgfX1W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2aHU23k1GEMDCUUCedbIkybgfX1W

fd513 := client.Open("/7m9F5RYwtA.txt", client.O_RDWR|client.O_CREATE)
if(fd513 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd510, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "BCK67o5aYAZT") {
  panic("wrong data returned")
}


fd514 := client.Open("/BLRdi2bjRc.txt", client.O_RDWR|client.O_CREATE)
if(fd514 < 0) {
  panic("open failed")
}


fd515 := client.Open("/IIAbSMKzYz.txt", client.O_RDWR|client.O_CREATE)
if(fd515 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd512, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (84) m0jLrbYPKisg9MdluGLr4mYaJvfJvvplJGVHpq3R7fKHdcodbKQ4S8mEHSW1Rjl4P6jlWjCT94O0HrFS8dA6

ret = client.Write(fd490, []byte("AXMBZhFMqJfpvhI3ocmJqQ5vxX7cfQ3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) m0jLrbYPKisg9MdluGLr4mYaJvfJvvplJGVHpq3R7fKHdcodbKQ4S8mEHSW1Rjl4P6jlWjCT94O0HrFS8dA6AXMBZhFMqJfpvhI3ocmJqQ5vxX7cfQ3
//fd state: (92) QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2aHU23k1GEMDCUUCedbIkybgfX1W

ret = client.Write(fd498, []byte("cEH7BTlv3cebKyDFwOwzQcxKS_L4mQr_5xFH50B8sSu_7OoqD7GGXjCOsURFs4U8JCdNyHn8I2QfVd9Nmb3Ty_rrK6Sfi6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (186) QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2aHU23k1GEMDCUUCedbIkybgfX1WcEH7BTlv3cebKyDFwOwzQcxKS_L4mQr_5xFH50B8sSu_7OoqD7GGXjCOsURFs4U8JCdNyHn8I2QfVd9Nmb3Ty_rrK6Sfi6

fd516 := client.Open("/QhEtwrbUoj.txt", client.O_RDWR|client.O_CREATE)
if(fd516 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd494, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd514)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) m0jLrbYPKisg9MdluGLr4mYaJvfJvvplJGVHpq3R7fKHdcodbKQ4S8mEHSW1Rjl4P6jlWjCT94O0HrFS8dA6AXMBZhFMqJfpvhI3ocmJqQ5vxX7cfQ3

ret = client.Write(fd503, []byte("7zv49rcMrXNLX41jNfmTTQh04IAXTBLyLypDvAGKP96AnubFaZN7qcyeyFkCviHq449h1YGNI8tyelia8Fvv8eVjzwu0nW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) 7zv49rcMrXNLX41jNfmTTQh04IAXTBLyLypDvAGKP96AnubFaZN7qcyeyFkCviHq449h1YGNI8tyelia8Fvv8eVjzwu0nWfpvhI3ocmJqQ5vxX7cfQ3
//fd state: (0) 

ret = client.Write(fd513, []byte("LXtMWUcTtU6rZDMt65FzTZDGxkRP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) LXtMWUcTtU6rZDMt65FzTZDGxkRP

ret = client.Seek(fd498, 66, client.SEEK_SET)
if(ret != 66) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 66)
  panic("seek failed")
}


buf, ret = client.Read(fd509, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd503, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "fpvhI3ocmJqQ5vxX7cfQ3") {
  panic("wrong data returned")
}


fd517 := client.Open("/UDXib7cZUL.txt", client.O_RDWR|client.O_CREATE)
if(fd517 < 0) {
  panic("open failed")
}


ret = client.Close(fd492)
if(ret != 0) {
  panic("close failed")
}


fd518 := client.Open("/goNewdKZOC.txt", client.O_RDWR|client.O_CREATE)
if(fd518 < 0) {
  panic("open failed")
}

//fd state: (228) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RZbhSkuj8n_eHn6LL1ERsnP0X629PAyFD4T7m8TcRmTyzK2oIuXbbzF0Y7Af8pkWZqIoidMat_b1wKHcdHfuAlgMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9my

ret = client.Write(fd485, []byte("YRRgEe1Iq_LoMN6buHUH7sJWq0oXAKzW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (260) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RZbhSkuj8n_eHn6LL1ERsnP0X629PAyFD4T7m8TcRmTyzK2oIuXbbzF0Y7Af8pkWZqIoidMat_b1wKHcdHfuAlgMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9myYRRgEe1Iq_LoMN6buHUH7sJWq0oXAKzW

fd519 := client.Open("/7fsL5MiIVz.txt", client.O_RDWR|client.O_CREATE)
if(fd519 < 0) {
  panic("open failed")
}


ret = client.Close(fd490)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd515, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd485, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd515)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd494, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (66) QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2aHU23k1GEMDCUUCedbIkybgfX1WcEH7BTlv3cebKyDFwOwzQcxKS_L4mQr_5xFH50B8sSu_7OoqD7GGXjCOsURFs4U8JCdNyHn8I2QfVd9Nmb3Ty_rrK6Sfi6

ret = client.Write(fd498, []byte("6qFB6PFbT36DK3NwxyEO3C1e360SA2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2a6qFB6PFbT36DK3NwxyEO3C1e360SA2BTlv3cebKyDFwOwzQcxKS_L4mQr_5xFH50B8sSu_7OoqD7GGXjCOsURFs4U8JCdNyHn8I2QfVd9Nmb3Ty_rrK6Sfi6

ret = client.Close(fd516)
if(ret != 0) {
  panic("close failed")
}

//fd state: (260) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RZbhSkuj8n_eHn6LL1ERsnP0X629PAyFD4T7m8TcRmTyzK2oIuXbbzF0Y7Af8pkWZqIoidMat_b1wKHcdHfuAlgMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9myYRRgEe1Iq_LoMN6buHUH7sJWq0oXAKzW

ret = client.Write(fd485, []byte("A22ihtGKrXZ62cEn5ckuJ_7kQ62ij4qJLSHOe7BbRGMPti0yciBC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (312) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RZbhSkuj8n_eHn6LL1ERsnP0X629PAyFD4T7m8TcRmTyzK2oIuXbbzF0Y7Af8pkWZqIoidMat_b1wKHcdHfuAlgMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9myYRRgEe1Iq_LoMN6buHUH7sJWq0oXAKzWA22ihtGKrXZ62cEn5ckuJ_7kQ62ij4qJLSHOe7BbRGMPti0yciBC

fd520 := client.Open("/AzKWKSX2Zd.txt", client.O_RDWR|client.O_CREATE)
if(fd520 < 0) {
  panic("open failed")
}


ret = client.Close(fd517)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd500)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd510, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd501)
if(ret != 0) {
  panic("close failed")
}


fd521 := client.Open("/gNxnBL3TfS.txt", client.O_RDWR|client.O_CREATE)
if(fd521 < 0) {
  panic("open failed")
}


ret = client.Close(fd503)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd485)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd506, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd494, []byte("VANAIYl9Mlf5jWo3AuSogxFHzqKpXsxDbtDsh9UwljbpX68Zvzmxg1JubzL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) VANAIYl9Mlf5jWo3AuSogxFHzqKpXsxDbtDsh9UwljbpX68Zvzmxg1JubzL

fd522 := client.Open("/14fPFDega2.txt", client.O_RDWR|client.O_CREATE)
if(fd522 < 0) {
  panic("open failed")
}


ret = client.Seek(fd512, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd494, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd513, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd523 := client.Open("/_CVl1cUh_V.txt", client.O_RDWR|client.O_CREATE)
if(fd523 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd506, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd505)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd520, []byte("an4Xn2rJNHeLZybA4nm5QA1dC55BgOoP_0eGq7SvqasmKR7pzysYBGeIUS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) an4Xn2rJNHeLZybA4nm5QA1dC55BgOoP_0eGq7SvqasmKR7pzysYBGeIUS

buf, ret = client.Read(fd521, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd524 := client.Open("/OP8ViwdWxk.txt", client.O_RDWR|client.O_CREATE)
if(fd524 < 0) {
  panic("open failed")
}


ret = client.Seek(fd506, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (28) LXtMWUcTtU6rZDMt65FzTZDGxkRP

ret = client.Write(fd513, []byte("08GRazNXbu9sPMFd82nC7IP37Bo2XJhz2mU6Oxtvi0TLySLJYP1QGHgjrofCgL4epIdIlSOh30n49a7jHO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) LXtMWUcTtU6rZDMt65FzTZDGxkRP08GRazNXbu9sPMFd82nC7IP37Bo2XJhz2mU6Oxtvi0TLySLJYP1QGHgjrofCgL4epIdIlSOh30n49a7jHO

fd525 := client.Open("/obmRgRRVEq.txt", client.O_RDWR|client.O_CREATE)
if(fd525 < 0) {
  panic("open failed")
}


ret = client.Seek(fd509, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


fd526 := client.Open("/3LEHih0czf.txt", client.O_RDWR|client.O_CREATE)
if(fd526 < 0) {
  panic("open failed")
}


ret = client.Close(fd522)
if(ret != 0) {
  panic("close failed")
}


fd527 := client.Open("/hm40EpXTIC.txt", client.O_RDWR|client.O_CREATE)
if(fd527 < 0) {
  panic("open failed")
}


ret = client.Close(fd520)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd502, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd509, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YgYcW") {
  panic("wrong data returned")
}


fd528 := client.Open("/z3E8FPWTfv.txt", client.O_RDWR|client.O_CREATE)
if(fd528 < 0) {
  panic("open failed")
}


ret = client.Seek(fd513, 94, client.SEEK_SET)
if(ret != 94) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 94)
  panic("seek failed")
}


fd529 := client.Open("/_laoVJpAG9.txt", client.O_RDWR|client.O_CREATE)
if(fd529 < 0) {
  panic("open failed")
}


ret = client.Close(fd525)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd526, []byte("I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_
//fd state: (63) uGKwd0isIuLIhmgnyRt78py7Q5YfJ0axKU60E25ZBn1DLoBMJgqBI062slYgYcWA57xg8sBewAQ0

ret = client.Write(fd509, []byte("RNfol4ydIZ94BCq4zP1vbuq78lMrrY9UUQYgyfAc8G22NS78gV6fqoE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) uGKwd0isIuLIhmgnyRt78py7Q5YfJ0axKU60E25ZBn1DLoBMJgqBI062slYgYcWRNfol4ydIZ94BCq4zP1vbuq78lMrrY9UUQYgyfAc8G22NS78gV6fqoE
//fd state: (99) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_

ret = client.Write(fd526, []byte("btwKLCx_W485Q63QqsP0Tn_Fpgbc69kea7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_btwKLCx_W485Q63QqsP0Tn_Fpgbc69kea7

buf, ret = client.Read(fd523, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd512, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd513, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


fd530 := client.Open("/h9euK4ZjF8.txt", client.O_RDWR|client.O_CREATE)
if(fd530 < 0) {
  panic("open failed")
}

//fd state: (133) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_btwKLCx_W485Q63QqsP0Tn_Fpgbc69kea7

ret = client.Write(fd526, []byte("M3EcE59KVZiOM5f8hYcyKZzfumeXCwgnTDOxQK94Qhju7G646qisU27f93zuhtU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (196) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_btwKLCx_W485Q63QqsP0Tn_Fpgbc69kea7M3EcE59KVZiOM5f8hYcyKZzfumeXCwgnTDOxQK94Qhju7G646qisU27f93zuhtU

ret = client.Seek(fd510, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}

//fd state: (118) uGKwd0isIuLIhmgnyRt78py7Q5YfJ0axKU60E25ZBn1DLoBMJgqBI062slYgYcWRNfol4ydIZ94BCq4zP1vbuq78lMrrY9UUQYgyfAc8G22NS78gV6fqoE

ret = client.Write(fd509, []byte("UbUes_DhIoM_aYgRpTaWSdMlGlcBFhKx4RLXfcLvqlEYm65OQPftAqGdzu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (176) uGKwd0isIuLIhmgnyRt78py7Q5YfJ0axKU60E25ZBn1DLoBMJgqBI062slYgYcWRNfol4ydIZ94BCq4zP1vbuq78lMrrY9UUQYgyfAc8G22NS78gV6fqoEUbUes_DhIoM_aYgRpTaWSdMlGlcBFhKx4RLXfcLvqlEYm65OQPftAqGdzu

fd531 := client.Open("/5v8sR02Va7.txt", client.O_RDWR|client.O_CREATE)
if(fd531 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd527, []byte("K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLVyvU6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLVyvU6

ret = client.Seek(fd528, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}

//fd state: (59) VANAIYl9Mlf5jWo3AuSogxFHzqKpXsxDbtDsh9UwljbpX68Zvzmxg1JubzL

ret = client.Write(fd494, []byte("qOSgzI7JylCD77SHAKd0Y1WQlAwYkBuVQqiFFZB5Vp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) VANAIYl9Mlf5jWo3AuSogxFHzqKpXsxDbtDsh9UwljbpX68Zvzmxg1JubzLqOSgzI7JylCD77SHAKd0Y1WQlAwYkBuVQqiFFZB5Vp

fd532 := client.Open("/kYPAZl3Z2l.txt", client.O_RDWR|client.O_CREATE)
if(fd532 < 0) {
  panic("open failed")
}


ret = client.Close(fd508)
if(ret != 0) {
  panic("close failed")
}

//fd state: (75) LXtMWUcTtU6rZDMt65FzTZDGxkRP08GRazNXbu9sPMFd82nC7IP37Bo2XJhz2mU6Oxtvi0TLySLJYP1QGHgjrofCgL4epIdIlSOh30n49a7jHO

ret = client.Write(fd513, []byte("1a2pY65Mq6dCB9UBhuvZEq0IC2f6_B804tc062sOA9Vgkaccoh1xSIwodQoKnbJ7Pbn1xfBS5dOb4edHFs6SdJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (161) LXtMWUcTtU6rZDMt65FzTZDGxkRP08GRazNXbu9sPMFd82nC7IP37Bo2XJhz2mU6Oxtvi0TLySL1a2pY65Mq6dCB9UBhuvZEq0IC2f6_B804tc062sOA9Vgkaccoh1xSIwodQoKnbJ7Pbn1xfBS5dOb4edHFs6SdJ

fd533 := client.Open("/FR2o5_Ez9y.txt", client.O_RDWR|client.O_CREATE)
if(fd533 < 0) {
  panic("open failed")
}


ret = client.Seek(fd526, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


fd534 := client.Open("/2VZzqpP7_6.txt", client.O_RDWR|client.O_CREATE)
if(fd534 < 0) {
  panic("open failed")
}

//fd state: (0) XzTCExLW1Gw

ret = client.Write(fd511, []byte("Lds7Q7th1LqL8YLOzOAIX13"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) Lds7Q7th1LqL8YLOzOAIX13

ret = client.Seek(fd502, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}


ret = client.Close(fd494)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd532, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd535 := client.Open("/JmFn4HpF0R.txt", client.O_RDWR|client.O_CREATE)
if(fd535 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd529, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "A9awLY38LiPZ") {
  panic("wrong data returned")
}


fd536 := client.Open("/CY9cjfit5S.txt", client.O_RDWR|client.O_CREATE)
if(fd536 < 0) {
  panic("open failed")
}


ret = client.Seek(fd532, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd533, []byte("dAef9HblEMp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) dAef9HblEMp

ret = client.Seek(fd523, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd537 := client.Open("/Asi9uxAZi0.txt", client.O_RDWR|client.O_CREATE)
if(fd537 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd498, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "BTlv3cebKyDFwOwzQcxKS_L4mQr_5xFH50B8sSu_7OoqD7GGXjCOsURFs4U8JCdNyHn8I2QfVd9Nmb3Ty_rrK6Sfi6") {
  panic("wrong data returned")
}


ret = client.Seek(fd498, 115, client.SEEK_SET)
if(ret != 115) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 115)
  panic("seek failed")
}

//fd state: (11) BCK67o5aYAZT

ret = client.Write(fd510, []byte("R2k4E0nLaWuzrjj8ImIJ6tBQPmzZ3BRIuXmLfP9R5QwUK6MfJ1m"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) BCK67o5aYAZR2k4E0nLaWuzrjj8ImIJ6tBQPmzZ3BRIuXmLfP9R5QwUK6MfJ1m

fd538 := client.Open("/finUwiQx86.txt", client.O_RDWR|client.O_CREATE)
if(fd538 < 0) {
  panic("open failed")
}


ret = client.Close(fd534)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd531, []byte("idHfzyteYDaPD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) idHfzyteYDaPD

ret = client.Close(fd535)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd509, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


buf, ret = client.Read(fd529, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd512, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd539 := client.Open("/XwukkFpNt2.txt", client.O_RDWR|client.O_CREATE)
if(fd539 < 0) {
  panic("open failed")
}


ret = client.Seek(fd527, 60, client.SEEK_SET)
if(ret != 60) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 60)
  panic("seek failed")
}


ret = client.Close(fd506)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd512, []byte("IZ9x02e4YrMreIspWfHfEZ38HKU9pgA0V9wqGEgVk_A_Z3NCscmiYKqDoFWWUvQjsMMK0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) IZ9x02e4YrMreIspWfHfEZ38HKU9pgA0V9wqGEgVk_A_Z3NCscmiYKqDoFWWUvQjsMMK0
//fd state: (0) 

ret = client.Write(fd539, []byte("gdb4dQCvYv7FuvIBoU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) gdb4dQCvYv7FuvIBoU

fd540 := client.Open("/V9WUaWK0vl.txt", client.O_RDWR|client.O_CREATE)
if(fd540 < 0) {
  panic("open failed")
}


ret = client.Seek(fd518, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd541 := client.Open("/c5pmebz5PR.txt", client.O_RDWR|client.O_CREATE)
if(fd541 < 0) {
  panic("open failed")
}


ret = client.Seek(fd532, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd523, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd509, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd509, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uGKwd0i") {
  panic("wrong data returned")
}


fd542 := client.Open("/UaSShUyc34.txt", client.O_RDWR|client.O_CREATE)
if(fd542 < 0) {
  panic("open failed")
}


ret = client.Close(fd518)
if(ret != 0) {
  panic("close failed")
}


fd543 := client.Open("/WPgv6yHnas.txt", client.O_RDWR|client.O_CREATE)
if(fd543 < 0) {
  panic("open failed")
}


ret = client.Seek(fd510, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Close(fd539)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd523, []byte("Ioy8k2Jqygrzm2KZuBfY9AX3TJwOYkdcqR3geGxPOu_frIHYX2MR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) Ioy8k2Jqygrzm2KZuBfY9AX3TJwOYkdcqR3geGxPOu_frIHYX2MR

ret = client.Seek(fd509, 163, client.SEEK_SET)
if(ret != 163) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 163)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd519, []byte("af945AQN8Ojj3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) af945AQN8Ojj3

ret = client.Seek(fd536, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


buf, ret = client.Read(fd523, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd537, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd530, []byte("PGN5MB5z9CxkxTOFR1xkXBPQ4oDob0ayfYASL2QBlmkcF_CCIHnDT9jTBDCjeJOPF8R9dXyc7Kl_Q4oTKVf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) PGN5MB5z9CxkxTOFR1xkXBPQ4oDob0ayfYASL2QBlmkcF_CCIHnDT9jTBDCjeJOPF8R9dXyc7Kl_Q4oTKVf
//fd state: (0) 

ret = client.Write(fd538, []byte("Yu3vjxKKQvMrnK9KslGCZzoz0Lbx_Ffa_MdwUxun74OTp800YFJQ6CbEsnwjCXcpJFoHGXSqFdjuLTfVLH5jPjC5zEvoIOI68KH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) Yu3vjxKKQvMrnK9KslGCZzoz0Lbx_Ffa_MdwUxun74OTp800YFJQ6CbEsnwjCXcpJFoHGXSqFdjuLTfVLH5jPjC5zEvoIOI68KH

ret = client.Seek(fd532, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd540, []byte("sj3ylVuMRfDhQg9DaX7ZTSBGwlxMVatBz5hJlAUkQib0TOImL_QmnYaj_Dyr9qC082pbLtD3WlRfph1ULWdZ9rCnmWHAjr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) sj3ylVuMRfDhQg9DaX7ZTSBGwlxMVatBz5hJlAUkQib0TOImL_QmnYaj_Dyr9qC082pbLtD3WlRfph1ULWdZ9rCnmWHAjr

fd544 := client.Open("/0r8NtQPaBF.txt", client.O_RDWR|client.O_CREATE)
if(fd544 < 0) {
  panic("open failed")
}


fd545 := client.Open("/HN3w0IwSey.txt", client.O_RDWR|client.O_CREATE)
if(fd545 < 0) {
  panic("open failed")
}

//fd state: (60) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLVyvU6

ret = client.Write(fd527, []byte("_uc4CxoSW_ZNYnswS8SNbBgt6hmTM6XipKBxyeINthBidt8jO0PjmcMyi6140SQXXzyXoWcZADZZ2dbGH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (141) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLV_uc4CxoSW_ZNYnswS8SNbBgt6hmTM6XipKBxyeINthBidt8jO0PjmcMyi6140SQXXzyXoWcZADZZ2dbGH
//fd state: (141) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLV_uc4CxoSW_ZNYnswS8SNbBgt6hmTM6XipKBxyeINthBidt8jO0PjmcMyi6140SQXXzyXoWcZADZZ2dbGH

ret = client.Write(fd527, []byte("aF7RBCs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (148) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLV_uc4CxoSW_ZNYnswS8SNbBgt6hmTM6XipKBxyeINthBidt8jO0PjmcMyi6140SQXXzyXoWcZADZZ2dbGHaF7RBCs

ret = client.Seek(fd498, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd544, []byte("ui0TklYlRhxCrelSKs2F0gWZcZW7Zgr0gZhY5BjmLUCDO5Gyuk2MFp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) ui0TklYlRhxCrelSKs2F0gWZcZW7Zgr0gZhY5BjmLUCDO5Gyuk2MFp
//fd state: (12) A9awLY38LiPZ

ret = client.Write(fd529, []byte("xxErjJwTOKWFl1W6C9qajRuMfNEyYeQNmO5Rt8RA6r71nqSNKH6csapyCILOGOjSQh1PlRW9LQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) A9awLY38LiPZxxErjJwTOKWFl1W6C9qajRuMfNEyYeQNmO5Rt8RA6r71nqSNKH6csapyCILOGOjSQh1PlRW9LQ
//fd state: (0) 

ret = client.Write(fd543, []byte("cnpHYI1cXCSPJr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) cnpHYI1cXCSPJr

fd546 := client.Open("/oIsR5Rk83H.txt", client.O_RDWR|client.O_CREATE)
if(fd546 < 0) {
  panic("open failed")
}


ret = client.Seek(fd526, 170, client.SEEK_SET)
if(ret != 170) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 170)
  panic("seek failed")
}


buf, ret = client.Read(fd512, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd523, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd547 := client.Open("/quPzIH1Oc_.txt", client.O_RDWR|client.O_CREATE)
if(fd547 < 0) {
  panic("open failed")
}


ret = client.Close(fd543)
if(ret != 0) {
  panic("close failed")
}

//fd state: (170) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_btwKLCx_W485Q63QqsP0Tn_Fpgbc69kea7M3EcE59KVZiOM5f8hYcyKZzfumeXCwgnTDOxQK94Qhju7G646qisU27f93zuhtU

ret = client.Write(fd526, []byte("Tu48U4yDEBxH_IJgJ1pdcvLas0hGMO9KypEe8NVWJ0NVBspH1igJiluuqPCROrPxHdlzK73kzvzb2ePurXmzibibDTny74"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (264) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_btwKLCx_W485Q63QqsP0Tn_Fpgbc69kea7M3EcE59KVZiOM5f8hYcyKZzfumeXCwgnTDOxQTu48U4yDEBxH_IJgJ1pdcvLas0hGMO9KypEe8NVWJ0NVBspH1igJiluuqPCROrPxHdlzK73kzvzb2ePurXmzibibDTny74
//fd state: (86) A9awLY38LiPZxxErjJwTOKWFl1W6C9qajRuMfNEyYeQNmO5Rt8RA6r71nqSNKH6csapyCILOGOjSQh1PlRW9LQ

ret = client.Write(fd529, []byte("u4hkSDszsWX5F0U"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) A9awLY38LiPZxxErjJwTOKWFl1W6C9qajRuMfNEyYeQNmO5Rt8RA6r71nqSNKH6csapyCILOGOjSQh1PlRW9LQu4hkSDszsWX5F0U
//fd state: (99) Yu3vjxKKQvMrnK9KslGCZzoz0Lbx_Ffa_MdwUxun74OTp800YFJQ6CbEsnwjCXcpJFoHGXSqFdjuLTfVLH5jPjC5zEvoIOI68KH

ret = client.Write(fd538, []byte("EMH__2X_rP81_Lx4bcm6W_Qe9Ytg05IKhiOKgUnI3WngxsqMtoO54KSCLaGPY4fnMzNmYn1zncje"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (175) Yu3vjxKKQvMrnK9KslGCZzoz0Lbx_Ffa_MdwUxun74OTp800YFJQ6CbEsnwjCXcpJFoHGXSqFdjuLTfVLH5jPjC5zEvoIOI68KHEMH__2X_rP81_Lx4bcm6W_Qe9Ytg05IKhiOKgUnI3WngxsqMtoO54KSCLaGPY4fnMzNmYn1zncje

buf, ret = client.Read(fd544, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd512, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd547, []byte("JPH9MXbXUs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) JPH9MXbXUs

buf, ret = client.Read(fd528, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3r") {
  panic("wrong data returned")
}


fd548 := client.Open("/_goIngjPLS.txt", client.O_RDWR|client.O_CREATE)
if(fd548 < 0) {
  panic("open failed")
}

//fd state: (23) Lds7Q7th1LqL8YLOzOAIX13

ret = client.Write(fd511, []byte("UgpPulvdk3dERdkx3cht9U3IdPxt19d0ELoouSxypycCbxldzIuEK32wolcayxBQ36"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) Lds7Q7th1LqL8YLOzOAIX13UgpPulvdk3dERdkx3cht9U3IdPxt19d0ELoouSxypycCbxldzIuEK32wolcayxBQ36
//fd state: (0) JeRoE9raBqOxoCpVbRlVeM8H0tXZVvx0jAL9avSBY1cbRL6CzgJ8_5TIqM8nl8d4ewbSmSihp0f_9Z

ret = client.Write(fd545, []byte("C4jGvha9OUMh3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) C4jGvha9OUMh3CpVbRlVeM8H0tXZVvx0jAL9avSBY1cbRL6CzgJ8_5TIqM8nl8d4ewbSmSihp0f_9Z
//fd state: (264) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_btwKLCx_W485Q63QqsP0Tn_Fpgbc69kea7M3EcE59KVZiOM5f8hYcyKZzfumeXCwgnTDOxQTu48U4yDEBxH_IJgJ1pdcvLas0hGMO9KypEe8NVWJ0NVBspH1igJiluuqPCROrPxHdlzK73kzvzb2ePurXmzibibDTny74

ret = client.Write(fd526, []byte("NUoGqnpkDNyS8_8RFw0Dv2XlzdYaltWe3jnZj7Sdv223jiI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (311) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_btwKLCx_W485Q63QqsP0Tn_Fpgbc69kea7M3EcE59KVZiOM5f8hYcyKZzfumeXCwgnTDOxQTu48U4yDEBxH_IJgJ1pdcvLas0hGMO9KypEe8NVWJ0NVBspH1igJiluuqPCROrPxHdlzK73kzvzb2ePurXmzibibDTny74NUoGqnpkDNyS8_8RFw0Dv2XlzdYaltWe3jnZj7Sdv223jiI

buf, ret = client.Read(fd512, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd549 := client.Open("/Buj7GpaWlK.txt", client.O_RDWR|client.O_CREATE)
if(fd549 < 0) {
  panic("open failed")
}


fd550 := client.Open("/8n_eRnEydB.txt", client.O_RDWR|client.O_CREATE)
if(fd550 < 0) {
  panic("open failed")
}


fd551 := client.Open("/5PS4vuVPFd.txt", client.O_RDWR|client.O_CREATE)
if(fd551 < 0) {
  panic("open failed")
}


fd552 := client.Open("/c5pmebz5PR.txt", client.O_RDWR|client.O_CREATE)
if(fd552 < 0) {
  panic("open failed")
}

//fd state: (148) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLV_uc4CxoSW_ZNYnswS8SNbBgt6hmTM6XipKBxyeINthBidt8jO0PjmcMyi6140SQXXzyXoWcZADZZ2dbGHaF7RBCs

ret = client.Write(fd527, []byte("b_gKochxVQgggcRVpXqzJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (169) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLV_uc4CxoSW_ZNYnswS8SNbBgt6hmTM6XipKBxyeINthBidt8jO0PjmcMyi6140SQXXzyXoWcZADZZ2dbGHaF7RBCsb_gKochxVQgggcRVpXqzJ

ret = client.Close(fd523)
if(ret != 0) {
  panic("close failed")
}

//fd state: (17) TWPh4AOrSRcZF3jzVq362OX4eJo2

ret = client.Write(fd536, []byte("SrzzUBymE2UQBT2jQcK85goHH0SBz65m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) TWPh4AOrSRcZF3jzVSrzzUBymE2UQBT2jQcK85goHH0SBz65m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeV

ret = client.Seek(fd529, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}


ret = client.Close(fd512)
if(ret != 0) {
  panic("close failed")
}

//fd state: (79) TWPh4AOrSRcZF3jzVSrzzUBymE2UQBT2jQcK85goHH0SBz65m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeV

ret = client.Write(fd536, []byte("m6nu_6tSssh_6eQp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) TWPh4AOrSRcZF3jzVSrzzUBymE2UQBT2jQcK85goHH0SBz65m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeVm6nu_6tSssh_6eQp

buf, ret = client.Read(fd549, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd553 := client.Open("/wQOfDmuJ6Y.txt", client.O_RDWR|client.O_CREATE)
if(fd553 < 0) {
  panic("open failed")
}


ret = client.Close(fd529)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd552, []byte("YT102ieSITc3XGodWnpG3xcaYsQ4hAVJoGFxiNxd_L6mC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) YT102ieSITc3XGodWnpG3xcaYsQ4hAVJoGFxiNxd_L6mC

ret = client.Close(fd510)
if(ret != 0) {
  panic("close failed")
}

//fd state: (68) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rymiePccJqO9RZbhSkuj8n_eHn6LL1ERsnP0X629PAyFD4T7m8TcRmTyzK2oIuXbbzF0Y7Af8pkWZqIoidMat_b1wKHcdHfuAlgMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9myYRRgEe1Iq_LoMN6buHUH7sJWq0oXAKzWA22ihtGKrXZ62cEn5ckuJ_7kQ62ij4qJLSHOe7BbRGMPti0yciBC

ret = client.Write(fd528, []byte("Y8T09kN5aTzqY8UacvpUp5ow_egyHmGuXqWoRySNWHoqPRZY_ZYuyKdoMp7Rc3vSmyZQJEQrFnhuNw8oLni5kczsdK5sgfKg6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (165) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rY8T09kN5aTzqY8UacvpUp5ow_egyHmGuXqWoRySNWHoqPRZY_ZYuyKdoMp7Rc3vSmyZQJEQrFnhuNw8oLni5kczsdK5sgfKg6gMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9myYRRgEe1Iq_LoMN6buHUH7sJWq0oXAKzWA22ihtGKrXZ62cEn5ckuJ_7kQ62ij4qJLSHOe7BbRGMPti0yciBC
//fd state: (81) QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2a6qFB6PFbT36DK3NwxyEO3C1e360SA2BTlv3cebKyDFwOwzQcxKS_L4mQr_5xFH50B8sSu_7OoqD7GGXjCOsURFs4U8JCdNyHn8I2QfVd9Nmb3Ty_rrK6Sfi6

ret = client.Write(fd498, []byte("fdg940IDCuGgjWefDCtRdumIDHsSFgPIoZqdtM3Nbti7o8U3UYDiSB5bWyLJipFApgFW6kPut43OR4ZSN8Sv1rk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (168) QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2a6qFB6PFbT36DK3Nfdg940IDCuGgjWefDCtRdumIDHsSFgPIoZqdtM3Nbti7o8U3UYDiSB5bWyLJipFApgFW6kPut43OR4ZSN8Sv1rkVd9Nmb3Ty_rrK6Sfi6

fd554 := client.Open("/YcfJ8uk2Wm.txt", client.O_RDWR|client.O_CREATE)
if(fd554 < 0) {
  panic("open failed")
}


fd555 := client.Open("/JGiSuqb_Uj.txt", client.O_RDWR|client.O_CREATE)
if(fd555 < 0) {
  panic("open failed")
}


ret = client.Close(fd553)
if(ret != 0) {
  panic("close failed")
}

//fd state: (13) C4jGvha9OUMh3CpVbRlVeM8H0tXZVvx0jAL9avSBY1cbRL6CzgJ8_5TIqM8nl8d4ewbSmSihp0f_9Z

ret = client.Write(fd545, []byte("QI40rW_gUYbEgLIUnssXhMsBJbmSIYgoKIRUerVh8TMPUsjBQHsyAkrHQYzrMG2mjVaSWj0kUlxWcMbDSGX1znl70J3g5Qc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (108) C4jGvha9OUMh3QI40rW_gUYbEgLIUnssXhMsBJbmSIYgoKIRUerVh8TMPUsjBQHsyAkrHQYzrMG2mjVaSWj0kUlxWcMbDSGX1znl70J3g5Qc

ret = client.Close(fd540)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd555, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd544, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd511, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd542, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd530)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd509, 125, client.SEEK_SET)
if(ret != 125) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 125)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd537, []byte("wy96mbut0TiWO9sagSXgYIhaQ5yDjLfhhJs3hWNjWEUluf1SV0Lo5IpCu3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) wy96mbut0TiWO9sagSXgYIhaQ5yDjLfhhJs3hWNjWEUluf1SV0Lo5IpCu3

ret = client.Seek(fd554, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}

//fd state: (45) YT102ieSITc3XGodWnpG3xcaYsQ4hAVJoGFxiNxd_L6mC

ret = client.Write(fd552, []byte("xv5wEJjheujt0zFW654d7AWNQPH5dHqain5dZCHiqL6Re2J8hHMLQ4JYmK8a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) YT102ieSITc3XGodWnpG3xcaYsQ4hAVJoGFxiNxd_L6mCxv5wEJjheujt0zFW654d7AWNQPH5dHqain5dZCHiqL6Re2J8hHMLQ4JYmK8a

buf, ret = client.Read(fd509, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hIoM_aYgRpTaWSdMlGlcBFhKx4RLXfcLvqlEYm65O") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd548, []byte("CjHP2I1xP_mF7FX1UjRNYf1gSNkTSDPXNl7Afqe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) CjHP2I1xP_mF7FX1UjRNYf1gSNkTSDPXNl7Afqe

ret = client.Seek(fd509, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}


ret = client.Seek(fd513, 126, client.SEEK_SET)
if(ret != 126) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 126)
  panic("seek failed")
}


ret = client.Seek(fd542, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd544)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd528, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9myYRRgEe") {
  panic("wrong data returned")
}


ret = client.Close(fd502)
if(ret != 0) {
  panic("close failed")
}

//fd state: (13) idHfzyteYDaPD

ret = client.Write(fd531, []byte("4p6t64Gk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) idHfzyteYDaPD4p6t64Gk

buf, ret = client.Read(fd526, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd526, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd556 := client.Open("/QWW32z5XWS.txt", client.O_RDWR|client.O_CREATE)
if(fd556 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd533, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd552, 99, client.SEEK_SET)
if(ret != 99) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 99)
  panic("seek failed")
}


buf, ret = client.Read(fd552, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JYmK8a") {
  panic("wrong data returned")
}

//fd state: (234) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rY8T09kN5aTzqY8UacvpUp5ow_egyHmGuXqWoRySNWHoqPRZY_ZYuyKdoMp7Rc3vSmyZQJEQrFnhuNw8oLni5kczsdK5sgfKg6gMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9myYRRgEe1Iq_LoMN6buHUH7sJWq0oXAKzWA22ihtGKrXZ62cEn5ckuJ_7kQ62ij4qJLSHOe7BbRGMPti0yciBC

ret = client.Write(fd528, []byte("ipScwePuECiTn2cHvbR0IwRoyJdkDQheHuv7dCtlOfkg00TLBSCNDr0G2To0voNhWAv1fn4SXTe6mO2c82NXFOtQiHCIiTpjlc5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (333) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rY8T09kN5aTzqY8UacvpUp5ow_egyHmGuXqWoRySNWHoqPRZY_ZYuyKdoMp7Rc3vSmyZQJEQrFnhuNw8oLni5kczsdK5sgfKg6gMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9myYRRgEeipScwePuECiTn2cHvbR0IwRoyJdkDQheHuv7dCtlOfkg00TLBSCNDr0G2To0voNhWAv1fn4SXTe6mO2c82NXFOtQiHCIiTpjlc5
//fd state: (0) 

ret = client.Write(fd532, []byte("YFrjXgdQ5Am_5DdDmccCXsqZb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) YFrjXgdQ5Am_5DdDmccCXsqZb

ret = client.Seek(fd532, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}

//fd state: (175) Yu3vjxKKQvMrnK9KslGCZzoz0Lbx_Ffa_MdwUxun74OTp800YFJQ6CbEsnwjCXcpJFoHGXSqFdjuLTfVLH5jPjC5zEvoIOI68KHEMH__2X_rP81_Lx4bcm6W_Qe9Ytg05IKhiOKgUnI3WngxsqMtoO54KSCLaGPY4fnMzNmYn1zncje

ret = client.Write(fd538, []byte("Jvr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (178) Yu3vjxKKQvMrnK9KslGCZzoz0Lbx_Ffa_MdwUxun74OTp800YFJQ6CbEsnwjCXcpJFoHGXSqFdjuLTfVLH5jPjC5zEvoIOI68KHEMH__2X_rP81_Lx4bcm6W_Qe9Ytg05IKhiOKgUnI3WngxsqMtoO54KSCLaGPY4fnMzNmYn1zncjeJvr
//fd state: (0) 

ret = client.Write(fd550, []byte("Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilgut91"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilgut91

ret = client.Seek(fd513, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


fd557 := client.Open("/_Arj1HylHr.txt", client.O_RDWR|client.O_CREATE)
if(fd557 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd555, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (311) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_btwKLCx_W485Q63QqsP0Tn_Fpgbc69kea7M3EcE59KVZiOM5f8hYcyKZzfumeXCwgnTDOxQTu48U4yDEBxH_IJgJ1pdcvLas0hGMO9KypEe8NVWJ0NVBspH1igJiluuqPCROrPxHdlzK73kzvzb2ePurXmzibibDTny74NUoGqnpkDNyS8_8RFw0Dv2XlzdYaltWe3jnZj7Sdv223jiI

ret = client.Write(fd526, []byte("YC6yqHH6NF3ERx6JAJoAeyx_HAMHvJ4mZe444QR3OuaoMh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (357) I4Qvb0eflPN_pkxcOemrIL4idxDtzGYwHUP02h7yq5NEozNK6jeoMo8TWnWnlxZeyzGz4LqLMvwUoWd4oT4i6su_LvGyqPl0OJ_btwKLCx_W485Q63QqsP0Tn_Fpgbc69kea7M3EcE59KVZiOM5f8hYcyKZzfumeXCwgnTDOxQTu48U4yDEBxH_IJgJ1pdcvLas0hGMO9KypEe8NVWJ0NVBspH1igJiluuqPCROrPxHdlzK73kzvzb2ePurXmzibibDTny74NUoGqnpkDNyS8_8RFw0Dv2XlzdYaltWe3jnZj7Sdv223jiIYC6yqHH6NF3ERx6JAJoAeyx_HAMHvJ4mZe444QR3OuaoMh

fd558 := client.Open("/j5i_lS91rD.txt", client.O_RDWR|client.O_CREATE)
if(fd558 < 0) {
  panic("open failed")
}


ret = client.Close(fd549)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd547, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd555, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd513, 136, client.SEEK_SET)
if(ret != 136) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 136)
  panic("seek failed")
}


ret = client.Close(fd546)
if(ret != 0) {
  panic("close failed")
}

//fd state: (178) Yu3vjxKKQvMrnK9KslGCZzoz0Lbx_Ffa_MdwUxun74OTp800YFJQ6CbEsnwjCXcpJFoHGXSqFdjuLTfVLH5jPjC5zEvoIOI68KHEMH__2X_rP81_Lx4bcm6W_Qe9Ytg05IKhiOKgUnI3WngxsqMtoO54KSCLaGPY4fnMzNmYn1zncjeJvr

ret = client.Write(fd538, []byte("nRkkaOz9MaXbG44Ngv5in5WzyWKX8hLRFcu0nu6fJEU5WyfqnEFQ8CV2u4xRhUIpuo3SzdOhogzA1DI19tDac7Kx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (266) Yu3vjxKKQvMrnK9KslGCZzoz0Lbx_Ffa_MdwUxun74OTp800YFJQ6CbEsnwjCXcpJFoHGXSqFdjuLTfVLH5jPjC5zEvoIOI68KHEMH__2X_rP81_Lx4bcm6W_Qe9Ytg05IKhiOKgUnI3WngxsqMtoO54KSCLaGPY4fnMzNmYn1zncjeJvrnRkkaOz9MaXbG44Ngv5in5WzyWKX8hLRFcu0nu6fJEU5WyfqnEFQ8CV2u4xRhUIpuo3SzdOhogzA1DI19tDac7Kx
//fd state: (95) TWPh4AOrSRcZF3jzVSrzzUBymE2UQBT2jQcK85goHH0SBz65m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeVm6nu_6tSssh_6eQp

ret = client.Write(fd536, []byte("XAmCjBpztoQcZ_SIDd60WEtOaVJn9yLdpxJk5HAup2INVzSFLwJxofo56"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) TWPh4AOrSRcZF3jzVSrzzUBymE2UQBT2jQcK85goHH0SBz65m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeVm6nu_6tSssh_6eQpXAmCjBpztoQcZ_SIDd60WEtOaVJn9yLdpxJk5HAup2INVzSFLwJxofo56
//fd state: (333) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rY8T09kN5aTzqY8UacvpUp5ow_egyHmGuXqWoRySNWHoqPRZY_ZYuyKdoMp7Rc3vSmyZQJEQrFnhuNw8oLni5kczsdK5sgfKg6gMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9myYRRgEeipScwePuECiTn2cHvbR0IwRoyJdkDQheHuv7dCtlOfkg00TLBSCNDr0G2To0voNhWAv1fn4SXTe6mO2c82NXFOtQiHCIiTpjlc5

ret = client.Write(fd528, []byte("LC2VqbXVAHoJM9dHRJfjQkQDAh061LQEqHCw7GDPF5YDnFcEz18z1O3U46Mqn2gagmIEbOJJOQDw3BQzyyLAxIpFG5_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (424) 6ZWnh7k9cBk9lxDuZB1ZQVFr40E3dKbmMHPAbcdooaTcjQOR9gSL_qTNg7a81l5v6x3rY8T09kN5aTzqY8UacvpUp5ow_egyHmGuXqWoRySNWHoqPRZY_ZYuyKdoMp7Rc3vSmyZQJEQrFnhuNw8oLni5kczsdK5sgfKg6gMZ95OaMcIx7sRfx519Fnqz9qWesQtyJfZOazNsivqxZJqM09UuFTeh3YBxb9myYRRgEeipScwePuECiTn2cHvbR0IwRoyJdkDQheHuv7dCtlOfkg00TLBSCNDr0G2To0voNhWAv1fn4SXTe6mO2c82NXFOtQiHCIiTpjlc5LC2VqbXVAHoJM9dHRJfjQkQDAh061LQEqHCw7GDPF5YDnFcEz18z1O3U46Mqn2gagmIEbOJJOQDw3BQzyyLAxIpFG5_

ret = client.Seek(fd557, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd511)
if(ret != 0) {
  panic("close failed")
}

//fd state: (39) CjHP2I1xP_mF7FX1UjRNYf1gSNkTSDPXNl7Afqe

ret = client.Write(fd548, []byte("iw0zCSweDGAvMrL6oLP1LxOpNLsbdCKtSgvGq_PyHpx6F7PSt5PDjKi9u8G5X7pXicWW5cmdTw1j3A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) CjHP2I1xP_mF7FX1UjRNYf1gSNkTSDPXNl7Afqeiw0zCSweDGAvMrL6oLP1LxOpNLsbdCKtSgvGq_PyHpx6F7PSt5PDjKi9u8G5X7pXicWW5cmdTw1j3A

ret = client.Close(fd538)
if(ret != 0) {
  panic("close failed")
}

//fd state: (168) QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2a6qFB6PFbT36DK3Nfdg940IDCuGgjWefDCtRdumIDHsSFgPIoZqdtM3Nbti7o8U3UYDiSB5bWyLJipFApgFW6kPut43OR4ZSN8Sv1rkVd9Nmb3Ty_rrK6Sfi6

ret = client.Write(fd498, []byte("1Y8myIY1FASQ4SlZA2_23rrJm4FPCIEhh0C2_fj5kH_j97VDB3QlOwuJDrd1TZXDN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (233) QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2a6qFB6PFbT36DK3Nfdg940IDCuGgjWefDCtRdumIDHsSFgPIoZqdtM3Nbti7o8U3UYDiSB5bWyLJipFApgFW6kPut43OR4ZSN8Sv1rk1Y8myIY1FASQ4SlZA2_23rrJm4FPCIEhh0C2_fj5kH_j97VDB3QlOwuJDrd1TZXDN

ret = client.Seek(fd550, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}


fd559 := client.Open("/9m1NnDd6_4.txt", client.O_RDWR|client.O_CREATE)
if(fd559 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd555, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (78) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilgut91

ret = client.Write(fd550, []byte("hAn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhAn91

ret = client.Seek(fd541, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


buf, ret = client.Read(fd559, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (152) TWPh4AOrSRcZF3jzVSrzzUBymE2UQBT2jQcK85goHH0SBz65m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeVm6nu_6tSssh_6eQpXAmCjBpztoQcZ_SIDd60WEtOaVJn9yLdpxJk5HAup2INVzSFLwJxofo56

ret = client.Write(fd536, []byte("Loi1VOefXNLpkB58TZu3bX0nxXKrXKQcJl9TNZ5daYop5eLd6WCVkxlfjHu950gOAS2ET"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (221) TWPh4AOrSRcZF3jzVSrzzUBymE2UQBT2jQcK85goHH0SBz65m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeVm6nu_6tSssh_6eQpXAmCjBpztoQcZ_SIDd60WEtOaVJn9yLdpxJk5HAup2INVzSFLwJxofo56Loi1VOefXNLpkB58TZu3bX0nxXKrXKQcJl9TNZ5daYop5eLd6WCVkxlfjHu950gOAS2ET

fd560 := client.Open("/RfchsYU08E.txt", client.O_RDWR|client.O_CREATE)
if(fd560 < 0) {
  panic("open failed")
}


fd561 := client.Open("/8W6113_G3d.txt", client.O_RDWR|client.O_CREATE)
if(fd561 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd550, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "91") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd557, []byte("OBREzjADYcfhI_ReHQLGaCZ4T9fwgKhxqpBWVQonBB_dilZ0NEtUf8viTVkQNhrXrHhwCURwyG_XLKEqNeDKrC8iQsb2O"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) OBREzjADYcfhI_ReHQLGaCZ4T9fwgKhxqpBWVQonBB_dilZ0NEtUf8viTVkQNhrXrHhwCURwyG_XLKEqNeDKrC8iQsb2O

buf, ret = client.Read(fd559, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd562 := client.Open("/49yoWYxM6B.txt", client.O_RDWR|client.O_CREATE)
if(fd562 < 0) {
  panic("open failed")
}


ret = client.Seek(fd547, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


buf, ret = client.Read(fd513, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bJ7Pbn1xfBS5dOb4edHFs6SdJ") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd541, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "c3XGodWnpG3xcaYsQ4hAVJoGFxiNxd_L6mCxv5wEJjheujt0zFW654d7AWNQPH5dHqain5dZCHiqL6Re2J8hHMLQ4JYmK8a") {
  panic("wrong data returned")
}

//fd state: (161) LXtMWUcTtU6rZDMt65FzTZDGxkRP08GRazNXbu9sPMFd82nC7IP37Bo2XJhz2mU6Oxtvi0TLySL1a2pY65Mq6dCB9UBhuvZEq0IC2f6_B804tc062sOA9Vgkaccoh1xSIwodQoKnbJ7Pbn1xfBS5dOb4edHFs6SdJ

ret = client.Write(fd513, []byte("Gc1xag2V07AAT1Cce895CsN6zMOR4pMxBbeXUw_fJmb1vK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (207) LXtMWUcTtU6rZDMt65FzTZDGxkRP08GRazNXbu9sPMFd82nC7IP37Bo2XJhz2mU6Oxtvi0TLySL1a2pY65Mq6dCB9UBhuvZEq0IC2f6_B804tc062sOA9Vgkaccoh1xSIwodQoKnbJ7Pbn1xfBS5dOb4edHFs6SdJGc1xag2V07AAT1Cce895CsN6zMOR4pMxBbeXUw_fJmb1vK

buf, ret = client.Read(fd562, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jBS") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd541, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd542, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd556)
if(ret != 0) {
  panic("close failed")
}

//fd state: (207) LXtMWUcTtU6rZDMt65FzTZDGxkRP08GRazNXbu9sPMFd82nC7IP37Bo2XJhz2mU6Oxtvi0TLySL1a2pY65Mq6dCB9UBhuvZEq0IC2f6_B804tc062sOA9Vgkaccoh1xSIwodQoKnbJ7Pbn1xfBS5dOb4edHFs6SdJGc1xag2V07AAT1Cce895CsN6zMOR4pMxBbeXUw_fJmb1vK

ret = client.Write(fd513, []byte("VHeTDQaSGC23h8tt764uZ6Juw8hn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (235) LXtMWUcTtU6rZDMt65FzTZDGxkRP08GRazNXbu9sPMFd82nC7IP37Bo2XJhz2mU6Oxtvi0TLySL1a2pY65Mq6dCB9UBhuvZEq0IC2f6_B804tc062sOA9Vgkaccoh1xSIwodQoKnbJ7Pbn1xfBS5dOb4edHFs6SdJGc1xag2V07AAT1Cce895CsN6zMOR4pMxBbeXUw_fJmb1vKVHeTDQaSGC23h8tt764uZ6Juw8hn

ret = client.Close(fd541)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd521, []byte("dtC3V5PxVBEIUJIoXshFP4ziJwqXORHgGNvCfH3wo_El_8HYEzqUK88A4Uj3Fb35R"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) dtC3V5PxVBEIUJIoXshFP4ziJwqXORHgGNvCfH3wo_El_8HYEzqUK88A4Uj3Fb35R

buf, ret = client.Read(fd498, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (4) YFrjXgdQ5Am_5DdDmccCXsqZb

ret = client.Write(fd532, []byte("dcOmzQ3Kl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) YFrjdcOmzQ3KlDdDmccCXsqZb
//fd state: (44) uGKwd0isIuLIhmgnyRt78py7Q5YfJ0axKU60E25ZBn1DLoBMJgqBI062slYgYcWRNfol4ydIZ94BCq4zP1vbuq78lMrrY9UUQYgyfAc8G22NS78gV6fqoEUbUes_DhIoM_aYgRpTaWSdMlGlcBFhKx4RLXfcLvqlEYm65OQPftAqGdzu

ret = client.Write(fd509, []byte("JFdYpwXbh28A21pyLhCuSOWlRPqKXHFaejS4Losl5uFNW3yLMo5fiKvtuVzWJ3VyHe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) uGKwd0isIuLIhmgnyRt78py7Q5YfJ0axKU60E25ZBn1DJFdYpwXbh28A21pyLhCuSOWlRPqKXHFaejS4Losl5uFNW3yLMo5fiKvtuVzWJ3VyHe8gV6fqoEUbUes_DhIoM_aYgRpTaWSdMlGlcBFhKx4RLXfcLvqlEYm65OQPftAqGdzu

buf, ret = client.Read(fd524, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd526)
if(ret != 0) {
  panic("close failed")
}


fd563 := client.Open("/uOmmUFMDvb.txt", client.O_RDWR|client.O_CREATE)
if(fd563 < 0) {
  panic("open failed")
}


ret = client.Seek(fd524, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd498, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (169) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLV_uc4CxoSW_ZNYnswS8SNbBgt6hmTM6XipKBxyeINthBidt8jO0PjmcMyi6140SQXXzyXoWcZADZZ2dbGHaF7RBCsb_gKochxVQgggcRVpXqzJ

ret = client.Write(fd527, []byte("KgrowNhfFlmgWptrMQ4klVenNFUMM4blo1xVuOTiy2a7Im0FnROqKPUUXDuZbjeqCm21Z9YaNR855"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (246) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLV_uc4CxoSW_ZNYnswS8SNbBgt6hmTM6XipKBxyeINthBidt8jO0PjmcMyi6140SQXXzyXoWcZADZZ2dbGHaF7RBCsb_gKochxVQgggcRVpXqzJKgrowNhfFlmgWptrMQ4klVenNFUMM4blo1xVuOTiy2a7Im0FnROqKPUUXDuZbjeqCm21Z9YaNR855

fd564 := client.Open("/Kn3GsyCkpr.txt", client.O_RDWR|client.O_CREATE)
if(fd564 < 0) {
  panic("open failed")
}


ret = client.Close(fd533)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd559, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (108) C4jGvha9OUMh3QI40rW_gUYbEgLIUnssXhMsBJbmSIYgoKIRUerVh8TMPUsjBQHsyAkrHQYzrMG2mjVaSWj0kUlxWcMbDSGX1znl70J3g5Qc

ret = client.Write(fd545, []byte("aUAs5sV1rs0YqmIV_663fQHLj83e8uD43eeTDhL6DdXxWkOqXMKMR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (161) C4jGvha9OUMh3QI40rW_gUYbEgLIUnssXhMsBJbmSIYgoKIRUerVh8TMPUsjBQHsyAkrHQYzrMG2mjVaSWj0kUlxWcMbDSGX1znl70J3g5QcaUAs5sV1rs0YqmIV_663fQHLj83e8uD43eeTDhL6DdXxWkOqXMKMR

fd565 := client.Open("/8W6113_G3d.txt", client.O_RDWR|client.O_CREATE)
if(fd565 < 0) {
  panic("open failed")
}


ret = client.Close(fd537)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd551, []byte("JcI_3CsCbU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) JcI_3CsCbU
//fd state: (0) 

ret = client.Write(fd555, []byte("JcY0VEQmWZ97gZ_lFu5ZMDTpQfDgtMf3A2SRN8jC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) JcY0VEQmWZ97gZ_lFu5ZMDTpQfDgtMf3A2SRN8jC

ret = client.Close(fd509)
if(ret != 0) {
  panic("close failed")
}


fd566 := client.Open("/rladecnREN.txt", client.O_RDWR|client.O_CREATE)
if(fd566 < 0) {
  panic("open failed")
}


ret = client.Close(fd521)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd559, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (38) YUwKleTtMubZgees0J_4HX7U1wtSjc0v_Xfynd3jMFXLgcwz

ret = client.Write(fd554, []byte("5bNhQxHlHQ9vsFLCe8Wx_WT8DzBRoEOjEgEjLSQfzI5z415d4fxaQl1eGKqLyyUczm5xOyrwEhEzo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) YUwKleTtMubZgees0J_4HX7U1wtSjc0v_Xfynd5bNhQxHlHQ9vsFLCe8Wx_WT8DzBRoEOjEgEjLSQfzI5z415d4fxaQl1eGKqLyyUczm5xOyrwEhEzo

buf, ret = client.Read(fd561, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "97V_WpedoDgdE9rQlulGDV5bWWwVUkNmZvoYmYBTJAxwaCSYXO6niMJxQXutVqi") {
  panic("wrong data returned")
}

//fd state: (161) C4jGvha9OUMh3QI40rW_gUYbEgLIUnssXhMsBJbmSIYgoKIRUerVh8TMPUsjBQHsyAkrHQYzrMG2mjVaSWj0kUlxWcMbDSGX1znl70J3g5QcaUAs5sV1rs0YqmIV_663fQHLj83e8uD43eeTDhL6DdXxWkOqXMKMR

ret = client.Write(fd545, []byte("xjgLV6P7omiQKLoIPgdgdhza9Vv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (188) C4jGvha9OUMh3QI40rW_gUYbEgLIUnssXhMsBJbmSIYgoKIRUerVh8TMPUsjBQHsyAkrHQYzrMG2mjVaSWj0kUlxWcMbDSGX1znl70J3g5QcaUAs5sV1rs0YqmIV_663fQHLj83e8uD43eeTDhL6DdXxWkOqXMKMRxjgLV6P7omiQKLoIPgdgdhza9Vv

ret = client.Close(fd566)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd542, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd551, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd548, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd547, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


ret = client.Close(fd558)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd542)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd550, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd498, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


ret = client.Close(fd528)
if(ret != 0) {
  panic("close failed")
}


fd567 := client.Open("/3wjB6vpdmW.txt", client.O_RDWR|client.O_CREATE)
if(fd567 < 0) {
  panic("open failed")
}

//fd state: (115) YUwKleTtMubZgees0J_4HX7U1wtSjc0v_Xfynd5bNhQxHlHQ9vsFLCe8Wx_WT8DzBRoEOjEgEjLSQfzI5z415d4fxaQl1eGKqLyyUczm5xOyrwEhEzo

ret = client.Write(fd554, []byte("cOm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) YUwKleTtMubZgees0J_4HX7U1wtSjc0v_Xfynd5bNhQxHlHQ9vsFLCe8Wx_WT8DzBRoEOjEgEjLSQfzI5z415d4fxaQl1eGKqLyyUczm5xOyrwEhEzocOm

buf, ret = client.Read(fd519, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (13) YFrjdcOmzQ3KlDdDmccCXsqZb

ret = client.Write(fd532, []byte("Ci0N0vHcbvtavDyYCTEq5QVEOBZL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) YFrjdcOmzQ3KlCi0N0vHcbvtavDyYCTEq5QVEOBZL
//fd state: (83) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhAn91

ret = client.Write(fd550, []byte("vUwmctdpLf2KyJbZZR1yDUIPrrY01qu9QWVt9R0QbFYdzXGP8SFzZUvJ3vV0wkCp755wYKY9HldHZW30Q5xqiVKnyQsGHeGMtQU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (182) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhAn91vUwmctdpLf2KyJbZZR1yDUIPrrY01qu9QWVt9R0QbFYdzXGP8SFzZUvJ3vV0wkCp755wYKY9HldHZW30Q5xqiVKnyQsGHeGMtQU

buf, ret = client.Read(fd519, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd568 := client.Open("/lK1Ra888uV.txt", client.O_RDWR|client.O_CREATE)
if(fd568 < 0) {
  panic("open failed")
}


ret = client.Close(fd513)
if(ret != 0) {
  panic("close failed")
}


fd569 := client.Open("/LY67aXtA3E.txt", client.O_RDWR|client.O_CREATE)
if(fd569 < 0) {
  panic("open failed")
}

//fd state: (41) YFrjdcOmzQ3KlCi0N0vHcbvtavDyYCTEq5QVEOBZL

ret = client.Write(fd532, []byte("5QFB4HhPYAm9X730ii"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) YFrjdcOmzQ3KlCi0N0vHcbvtavDyYCTEq5QVEOBZL5QFB4HhPYAm9X730ii

fd570 := client.Open("/IlPekcd4f8.txt", client.O_RDWR|client.O_CREATE)
if(fd570 < 0) {
  panic("open failed")
}


fd571 := client.Open("/OPhIpKtYFR.txt", client.O_RDWR|client.O_CREATE)
if(fd571 < 0) {
  panic("open failed")
}

//fd state: (40) JcY0VEQmWZ97gZ_lFu5ZMDTpQfDgtMf3A2SRN8jC

ret = client.Write(fd555, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) JcY0VEQmWZ97gZ_lFu5ZMDTpQfDgtMf3A2SRN8jC

ret = client.Seek(fd569, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd552, 98, client.SEEK_SET)
if(ret != 98) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 98)
  panic("seek failed")
}


buf, ret = client.Read(fd550, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd519, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


ret = client.Seek(fd564, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd572 := client.Open("/ymeRdV7EO7.txt", client.O_RDWR|client.O_CREATE)
if(fd572 < 0) {
  panic("open failed")
}

//fd state: (221) TWPh4AOrSRcZF3jzVSrzzUBymE2UQBT2jQcK85goHH0SBz65m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeVm6nu_6tSssh_6eQpXAmCjBpztoQcZ_SIDd60WEtOaVJn9yLdpxJk5HAup2INVzSFLwJxofo56Loi1VOefXNLpkB58TZu3bX0nxXKrXKQcJl9TNZ5daYop5eLd6WCVkxlfjHu950gOAS2ET

ret = client.Write(fd536, []byte("SsWTFpzGwnAa_BzSAOpwT4y11MUGiQgl7AP4vvUI_F"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (263) TWPh4AOrSRcZF3jzVSrzzUBymE2UQBT2jQcK85goHH0SBz65m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeVm6nu_6tSssh_6eQpXAmCjBpztoQcZ_SIDd60WEtOaVJn9yLdpxJk5HAup2INVzSFLwJxofo56Loi1VOefXNLpkB58TZu3bX0nxXKrXKQcJl9TNZ5daYop5eLd6WCVkxlfjHu950gOAS2ETSsWTFpzGwnAa_BzSAOpwT4y11MUGiQgl7AP4vvUI_F

ret = client.Close(fd557)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd524, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd532, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}

//fd state: (6) QaRqWzi6OwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2a6qFB6PFbT36DK3Nfdg940IDCuGgjWefDCtRdumIDHsSFgPIoZqdtM3Nbti7o8U3UYDiSB5bWyLJipFApgFW6kPut43OR4ZSN8Sv1rk1Y8myIY1FASQ4SlZA2_23rrJm4FPCIEhh0C2_fj5kH_j97VDB3QlOwuJDrd1TZXDN

ret = client.Write(fd498, []byte("8a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) QaRqWz8aOwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZNDtBFmY9Dl2a6qFB6PFbT36DK3Nfdg940IDCuGgjWefDCtRdumIDHsSFgPIoZqdtM3Nbti7o8U3UYDiSB5bWyLJipFApgFW6kPut43OR4ZSN8Sv1rk1Y8myIY1FASQ4SlZA2_23rrJm4FPCIEhh0C2_fj5kH_j97VDB3QlOwuJDrd1TZXDN
//fd state: (0) 

ret = client.Write(fd559, []byte("SqmtYHtjaCcMlTn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) SqmtYHtjaCcMlTn

fd573 := client.Open("/CKytIwcwjN.txt", client.O_RDWR|client.O_CREATE)
if(fd573 < 0) {
  panic("open failed")
}


fd574 := client.Open("/JqZAq55249.txt", client.O_RDWR|client.O_CREATE)
if(fd574 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd573, []byte("9rMJdkn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) 9rMJdkn

ret = client.Seek(fd570, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Seek(fd564, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd570)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd527)
if(ret != 0) {
  panic("close failed")
}


fd575 := client.Open("/f8ebU3uJAr.txt", client.O_RDWR|client.O_CREATE)
if(fd575 < 0) {
  panic("open failed")
}

//fd state: (15) SqmtYHtjaCcMlTn

ret = client.Write(fd559, []byte("ZSj_pWP_1m3g1hX4p"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) SqmtYHtjaCcMlTnZSj_pWP_1m3g1hX4p

buf, ret = client.Read(fd561, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd498)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd545, 167, client.SEEK_SET)
if(ret != 167) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 167)
  panic("seek failed")
}

//fd state: (98) YT102ieSITc3XGodWnpG3xcaYsQ4hAVJoGFxiNxd_L6mCxv5wEJjheujt0zFW654d7AWNQPH5dHqain5dZCHiqL6Re2J8hHMLQ4JYmK8a

ret = client.Write(fd552, []byte("AC3j8R2DEkLKtdRzf7GMdtnpXPmrj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (127) YT102ieSITc3XGodWnpG3xcaYsQ4hAVJoGFxiNxd_L6mCxv5wEJjheujt0zFW654d7AWNQPH5dHqain5dZCHiqL6Re2J8hHMLQAC3j8R2DEkLKtdRzf7GMdtnpXPmrj

fd576 := client.Open("/0PCO56vqgQ.txt", client.O_RDWR|client.O_CREATE)
if(fd576 < 0) {
  panic("open failed")
}


ret = client.Close(fd573)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd519, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "QN8Ojj3") {
  panic("wrong data returned")
}


ret = client.Seek(fd564, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd577 := client.Open("/o8e4tZKDW0.txt", client.O_RDWR|client.O_CREATE)
if(fd577 < 0) {
  panic("open failed")
}


fd578 := client.Open("/gsTfmWXCkm.txt", client.O_RDWR|client.O_CREATE)
if(fd578 < 0) {
  panic("open failed")
}


fd579 := client.Open("/59_CvOVLUp.txt", client.O_RDWR|client.O_CREATE)
if(fd579 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd554, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd554, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd580 := client.Open("/2vuPlCmIVr.txt", client.O_RDWR|client.O_CREATE)
if(fd580 < 0) {
  panic("open failed")
}


ret = client.Close(fd564)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd574)
if(ret != 0) {
  panic("close failed")
}

//fd state: (13) af945AQN8Ojj3

ret = client.Write(fd519, []byte("yXOAfd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) af945AQN8Ojj3yXOAfd

buf, ret = client.Read(fd568, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "QaRqWz8aOwwEITZpcjJrH2t2yvbBiuSo7XDMOH_IUbBchxHS5PnFAZ") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd569, []byte("0MJRJ_gsJxUoYW8oPCHLG1RjnyPWU6h_z0sC3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) 0MJRJ_gsJxUoYW8oPCHLG1RjnyPWU6h_z0sC3

fd581 := client.Open("/C0Zs6QlJna.txt", client.O_RDWR|client.O_CREATE)
if(fd581 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd568, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NDtBFmY9Dl2a6qFB6PFbT36DK3Nfdg940IDCuGgjWefDCtRdumIDHsSFg") {
  panic("wrong data returned")
}


ret = client.Seek(fd581, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (3) jBS64LgzdahuzqR1EUlI3j2optU2OgMVieqY7f23xT_p9nLbSLy2KVD66BUmdCW2aHsR3JAR4Jii_mBSietBx4GqcOiw8NF6K6f9iFXz5AiI9kNL1IOSBVi9rYL8hpAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQpgygN3aOEmoYMSI9sDypzuOtD86FeN7ykkKcOm83iIXDDKOMQLAtZnnM85iOcF_3

ret = client.Write(fd562, []byte("y60Vuotzt_hglX4goB5gScw_GfxacO8XyW4CLK2Vu2PgbKxkqvImnIouTCQhw1L_u2gKa0oAOB8MHE0iBnVP_m"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) jBSy60Vuotzt_hglX4goB5gScw_GfxacO8XyW4CLK2Vu2PgbKxkqvImnIouTCQhw1L_u2gKa0oAOB8MHE0iBnVP_mOiw8NF6K6f9iFXz5AiI9kNL1IOSBVi9rYL8hpAsztMj7XTXKCO6E2Nxi0kZSo4yPyYOZSufteArYCW5hRXkaLIf0Dl360gWTuwzttPxzublhf4CaQpgygN3aOEmoYMSI9sDypzuOtD86FeN7ykkKcOm83iIXDDKOMQLAtZnnM85iOcF_3

ret = client.Close(fd554)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd575)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd561)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd577, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd582 := client.Open("/eAmQFvoX3g.txt", client.O_RDWR|client.O_CREATE)
if(fd582 < 0) {
  panic("open failed")
}


ret = client.Seek(fd552, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}


fd583 := client.Open("/JqZAq55249.txt", client.O_RDWR|client.O_CREATE)
if(fd583 < 0) {
  panic("open failed")
}


ret = client.Seek(fd551, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


buf, ret = client.Read(fd559, 39)
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


ret = client.Seek(fd572, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd560, 72)
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


fd584 := client.Open("/Un4l5UUix3.txt", client.O_RDWR|client.O_CREATE)
if(fd584 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd577, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (40) JcY0VEQmWZ97gZ_lFu5ZMDTpQfDgtMf3A2SRN8jC

ret = client.Write(fd555, []byte("PHXE0vSUek0ztF2dQFhuuziuvMiP8NhWQIcXvuD_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) JcY0VEQmWZ97gZ_lFu5ZMDTpQfDgtMf3A2SRN8jCPHXE0vSUek0ztF2dQFhuuziuvMiP8NhWQIcXvuD_

buf, ret = client.Read(fd550, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd583, []byte("fJNDQs0eZUGYW_TwCNA3es_uhwZt8W6ZT7d2yXvQac_3PiGGTRR4tkmPUnf8SFl1a2Bm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) fJNDQs0eZUGYW_TwCNA3es_uhwZt8W6ZT7d2yXvQac_3PiGGTRR4tkmPUnf8SFl1a2Bm

fd585 := client.Open("/fs3zIIz3SH.txt", client.O_RDWR|client.O_CREATE)
if(fd585 < 0) {
  panic("open failed")
}


ret = client.Seek(fd559, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


fd586 := client.Open("/26dSM1oixo.txt", client.O_RDWR|client.O_CREATE)
if(fd586 < 0) {
  panic("open failed")
}


ret = client.Close(fd524)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd580)
if(ret != 0) {
  panic("close failed")
}


fd587 := client.Open("/3LEHih0czf.txt", client.O_RDWR|client.O_CREATE)
if(fd587 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd551, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_3CsCbU") {
  panic("wrong data returned")
}


ret = client.Close(fd583)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd531, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd588 := client.Open("/ehmPd82Tnz.txt", client.O_RDWR|client.O_CREATE)
if(fd588 < 0) {
  panic("open failed")
}


fd589 := client.Open("/VyUT3k7cYP.txt", client.O_RDWR|client.O_CREATE)
if(fd589 < 0) {
  panic("open failed")
}


ret = client.Seek(fd576, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd532, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "avDyYCTEq5QVEOBZL5QFB4H") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd571, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd545, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "P7omiQKLoIPgdgdhza9Vv") {
  panic("wrong data returned")
}


ret = client.Seek(fd589, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


ret = client.Seek(fd578, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


buf, ret = client.Read(fd565, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "97V_WpedoDgdE9rQlulGDV5bWWwVUkNmZvoYmYBTJAxwaCSYXO6niMJxQXutVqi") {
  panic("wrong data returned")
}

//fd state: (188) C4jGvha9OUMh3QI40rW_gUYbEgLIUnssXhMsBJbmSIYgoKIRUerVh8TMPUsjBQHsyAkrHQYzrMG2mjVaSWj0kUlxWcMbDSGX1znl70J3g5QcaUAs5sV1rs0YqmIV_663fQHLj83e8uD43eeTDhL6DdXxWkOqXMKMRxjgLV6P7omiQKLoIPgdgdhza9Vv

ret = client.Write(fd545, []byte("2BlMpnkyA9QnUpCL5mwFXGTpUFIeD9Yd2Al"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (223) C4jGvha9OUMh3QI40rW_gUYbEgLIUnssXhMsBJbmSIYgoKIRUerVh8TMPUsjBQHsyAkrHQYzrMG2mjVaSWj0kUlxWcMbDSGX1znl70J3g5QcaUAs5sV1rs0YqmIV_663fQHLj83e8uD43eeTDhL6DdXxWkOqXMKMRxjgLV6P7omiQKLoIPgdgdhza9Vv2BlMpnkyA9QnUpCL5mwFXGTpUFIeD9Yd2Al

fd590 := client.Open("/GAv93wc1kq.txt", client.O_RDWR|client.O_CREATE)
if(fd590 < 0) {
  panic("open failed")
}


ret = client.Seek(fd572, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd576)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd585)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd536)
if(ret != 0) {
  panic("close failed")
}


fd591 := client.Open("/rYKpiu9IG0.txt", client.O_RDWR|client.O_CREATE)
if(fd591 < 0) {
  panic("open failed")
}


ret = client.Close(fd562)
if(ret != 0) {
  panic("close failed")
}


fd592 := client.Open("/NdZKxMANLT.txt", client.O_RDWR|client.O_CREATE)
if(fd592 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd578, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Wy") {
  panic("wrong data returned")
}


ret = client.Close(fd589)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd531, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd581, []byte("47qbtVXel8wNdD3bB_FyXyNJvNSlb1sgt0W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) 47qbtVXel8wNdD3bB_FyXyNJvNSlb1sgt0W

buf, ret = client.Read(fd590, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd593 := client.Open("/Nz3Z01M3H0.txt", client.O_RDWR|client.O_CREATE)
if(fd593 < 0) {
  panic("open failed")
}


ret = client.Close(fd569)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd568)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd567)
if(ret != 0) {
  panic("close failed")
}


fd594 := client.Open("/iupwQl069R.txt", client.O_RDWR|client.O_CREATE)
if(fd594 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd588, []byte("OOUtvl5HKULKlNsVBNJLHT5tIp5GAfJ6tvABUo4LNYIfiodrX3f4g1mKY0Ib9yt4ppx0kUjRIarlXkI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) OOUtvl5HKULKlNsVBNJLHT5tIp5GAfJ6tvABUo4LNYIfiodrX3f4g1mKY0Ib9yt4ppx0kUjRIarlXkI

buf, ret = client.Read(fd550, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd595 := client.Open("/uA4d9BQc55.txt", client.O_RDWR|client.O_CREATE)
if(fd595 < 0) {
  panic("open failed")
}

//fd state: (182) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhAn91vUwmctdpLf2KyJbZZR1yDUIPrrY01qu9QWVt9R0QbFYdzXGP8SFzZUvJ3vV0wkCp755wYKY9HldHZW30Q5xqiVKnyQsGHeGMtQU

ret = client.Write(fd550, []byte("tpJxAcoBb751hnBZmszRahitRoyGtyrz5zRPEC5SEZAc1xVYbuBlYPnNP_Yld6Mv3fTp7tpKjMtofNRL7Zmu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (266) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhAn91vUwmctdpLf2KyJbZZR1yDUIPrrY01qu9QWVt9R0QbFYdzXGP8SFzZUvJ3vV0wkCp755wYKY9HldHZW30Q5xqiVKnyQsGHeGMtQUtpJxAcoBb751hnBZmszRahitRoyGtyrz5zRPEC5SEZAc1xVYbuBlYPnNP_Yld6Mv3fTp7tpKjMtofNRL7Zmu

ret = client.Seek(fd560, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd555, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


ret = client.Seek(fd532, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


fd596 := client.Open("/HRbgk5YrAz.txt", client.O_RDWR|client.O_CREATE)
if(fd596 < 0) {
  panic("open failed")
}


ret = client.Seek(fd548, 112, client.SEEK_SET)
if(ret != 112) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 112)
  panic("seek failed")
}


ret = client.Seek(fd582, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd547, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MXbXUs") {
  panic("wrong data returned")
}


fd597 := client.Open("/uA4d9BQc55.txt", client.O_RDWR|client.O_CREATE)
if(fd597 < 0) {
  panic("open failed")
}


ret = client.Seek(fd532, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


fd598 := client.Open("/HGIG0g8dvV.txt", client.O_RDWR|client.O_CREATE)
if(fd598 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd586, []byte("V87tZtyV9WfFurSBpKU3uam2K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) V87tZtyV9WfFurSBpKU3uam2K

ret = client.Seek(fd548, 61, client.SEEK_SET)
if(ret != 61) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 61)
  panic("seek failed")
}


ret = client.Seek(fd548, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}

//fd state: (266) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhAn91vUwmctdpLf2KyJbZZR1yDUIPrrY01qu9QWVt9R0QbFYdzXGP8SFzZUvJ3vV0wkCp755wYKY9HldHZW30Q5xqiVKnyQsGHeGMtQUtpJxAcoBb751hnBZmszRahitRoyGtyrz5zRPEC5SEZAc1xVYbuBlYPnNP_Yld6Mv3fTp7tpKjMtofNRL7Zmu

ret = client.Write(fd550, []byte("yy6UEpapWyh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (277) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhAn91vUwmctdpLf2KyJbZZR1yDUIPrrY01qu9QWVt9R0QbFYdzXGP8SFzZUvJ3vV0wkCp755wYKY9HldHZW30Q5xqiVKnyQsGHeGMtQUtpJxAcoBb751hnBZmszRahitRoyGtyrz5zRPEC5SEZAc1xVYbuBlYPnNP_Yld6Mv3fTp7tpKjMtofNRL7Zmuyy6UEpapWyh
//fd state: (0) 

ret = client.Write(fd560, []byte("cbEyh4wC1mLgOHDN9e8KvvO1egqa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) cbEyh4wC1mLgOHDN9e8KvvO1egqa

ret = client.Close(fd582)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd592)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd559, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (32) SqmtYHtjaCcMlTnZSj_pWP_1m3g1hX4p

ret = client.Write(fd559, []byte("0TMqXm9ee4wszWwYvjFLzqh1t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) SqmtYHtjaCcMlTnZSj_pWP_1m3g1hX4p0TMqXm9ee4wszWwYvjFLzqh1t
//fd state: (0) 

ret = client.Write(fd591, []byte("ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQD

ret = client.Close(fd551)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd531, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "64Gk") {
  panic("wrong data returned")
}


ret = client.Close(fd547)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd590, []byte("uDv_d_fSo8eh0k_bubUPkmOQNniC9qaVyVTXGt2dU2aAq1e6s_Y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) uDv_d_fSo8eh0k_bubUPkmOQNniC9qaVyVTXGt2dU2aAq1e6s_Y

ret = client.Seek(fd586, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


fd599 := client.Open("/6AgC4sSJvz.txt", client.O_RDWR|client.O_CREATE)
if(fd599 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd594, []byte("1SdSzqWWHN0fgyqdUpVH2oYNXrMR_847d5rD0CIV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) 1SdSzqWWHN0fgyqdUpVH2oYNXrMR_847d5rD0CIV

ret = client.Seek(fd550, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


ret = client.Seek(fd594, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


fd600 := client.Open("/lKIaBR3B02.txt", client.O_RDWR|client.O_CREATE)
if(fd600 < 0) {
  panic("open failed")
}


ret = client.Seek(fd560, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Seek(fd581, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Close(fd587)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd588)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd548, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}

//fd state: (19) 1SdSzqWWHN0fgyqdUpVH2oYNXrMR_847d5rD0CIV

ret = client.Write(fd594, []byte("IyApisWS6bbc97GwVuoV6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) 1SdSzqWWHN0fgyqdUpVIyApisWS6bbc97GwVuoV6

fd601 := client.Open("/QzZYzBiP2c.txt", client.O_RDWR|client.O_CREATE)
if(fd601 < 0) {
  panic("open failed")
}


ret = client.Seek(fd560, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


ret = client.Seek(fd593, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd602 := client.Open("/yXfsbC9K19.txt", client.O_RDWR|client.O_CREATE)
if(fd602 < 0) {
  panic("open failed")
}


ret = client.Close(fd571)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd579, []byte("lyhZ6MY60tHudDJpVpExG7ypQATd6697i8dIzj0WGB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) lyhZ6MY60tHudDJpVpExG7ypQATd6697i8dIzj0WGB

ret = client.Close(fd552)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd560, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "mLgOHDN9e8KvvO1e") {
  panic("wrong data returned")
}


ret = client.Close(fd599)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd597, []byte("4vAbqobHbGl7Q1RVVIsacmOLD82G7TcBdh5T9cbqXTwhlOCjrdZynJdT4B5sVcapMilIi4F5FwBX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) 4vAbqobHbGl7Q1RVVIsacmOLD82G7TcBdh5T9cbqXTwhlOCjrdZynJdT4B5sVcapMilIi4F5FwBX

ret = client.Seek(fd595, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}

//fd state: (51) uDv_d_fSo8eh0k_bubUPkmOQNniC9qaVyVTXGt2dU2aAq1e6s_Y

ret = client.Write(fd590, []byte("8yJ6dBlXVYQDwwv2HuLFgP7DED5zJNRetBjV_8_GtmPxo_jJxtx8BCa7RW_89O5_xJOuw3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (121) uDv_d_fSo8eh0k_bubUPkmOQNniC9qaVyVTXGt2dU2aAq1e6s_Y8yJ6dBlXVYQDwwv2HuLFgP7DED5zJNRetBjV_8_GtmPxo_jJxtx8BCa7RW_89O5_xJOuw3

fd603 := client.Open("/EqK6dUWAsY.txt", client.O_RDWR|client.O_CREATE)
if(fd603 < 0) {
  panic("open failed")
}


ret = client.Seek(fd591, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd604 := client.Open("/VyUT3k7cYP.txt", client.O_RDWR|client.O_CREATE)
if(fd604 < 0) {
  panic("open failed")
}


ret = client.Close(fd565)
if(ret != 0) {
  panic("close failed")
}


fd605 := client.Open("/XI3vFe5I6r.txt", client.O_RDWR|client.O_CREATE)
if(fd605 < 0) {
  panic("open failed")
}

//fd state: (79) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhAn91vUwmctdpLf2KyJbZZR1yDUIPrrY01qu9QWVt9R0QbFYdzXGP8SFzZUvJ3vV0wkCp755wYKY9HldHZW30Q5xqiVKnyQsGHeGMtQUtpJxAcoBb751hnBZmszRahitRoyGtyrz5zRPEC5SEZAc1xVYbuBlYPnNP_Yld6Mv3fTp7tpKjMtofNRL7Zmuyy6UEpapWyh

ret = client.Write(fd550, []byte("WAKtgCuG5NdBdCjAHwLWjvXuIV80ik5F"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhWAKtgCuG5NdBdCjAHwLWjvXuIV80ik5F1qu9QWVt9R0QbFYdzXGP8SFzZUvJ3vV0wkCp755wYKY9HldHZW30Q5xqiVKnyQsGHeGMtQUtpJxAcoBb751hnBZmszRahitRoyGtyrz5zRPEC5SEZAc1xVYbuBlYPnNP_Yld6Mv3fTp7tpKjMtofNRL7Zmuyy6UEpapWyh

buf, ret = client.Read(fd548, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "X1UjRNY") {
  panic("wrong data returned")
}


ret = client.Seek(fd578, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd593, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd572, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd606 := client.Open("/Ur_2N7oyWA.txt", client.O_RDWR|client.O_CREATE)
if(fd606 < 0) {
  panic("open failed")
}


ret = client.Close(fd590)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd597, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd607 := client.Open("/tr2x2m1usr.txt", client.O_RDWR|client.O_CREATE)
if(fd607 < 0) {
  panic("open failed")
}


ret = client.Close(fd601)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd531, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


buf, ret = client.Read(fd577, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd600, []byte("emmvFzI2cSWgqSgtWzNwuPQoF6pP9c6U0ME9mVFE9mGlT_8hGAJ1AZRIsrMmqDWKu5D"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) emmvFzI2cSWgqSgtWzNwuPQoF6pP9c6U0ME9mVFE9mGlT_8hGAJ1AZRIsrMmqDWKu5D

buf, ret = client.Read(fd602, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd594, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


fd608 := client.Open("/YWJtqDTpBK.txt", client.O_RDWR|client.O_CREATE)
if(fd608 < 0) {
  panic("open failed")
}


fd609 := client.Open("/njkggBWqLT.txt", client.O_RDWR|client.O_CREATE)
if(fd609 < 0) {
  panic("open failed")
}


ret = client.Close(fd602)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd593, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd608)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd603)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd591, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


buf, ret = client.Read(fd604, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WPPWZBw") {
  panic("wrong data returned")
}


ret = client.Close(fd595)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd578, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


buf, ret = client.Read(fd605, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd594, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7GwVuoV6") {
  panic("wrong data returned")
}


ret = client.Close(fd584)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd597)
if(ret != 0) {
  panic("close failed")
}


fd610 := client.Open("/oXvVSX6Wbh.txt", client.O_RDWR|client.O_CREATE)
if(fd610 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd605, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) dJ4qMyPqGAsZzT4HBKwlfG5HbDLftkMlgzavgqKhoyRGGDoi

ret = client.Write(fd596, []byte("6Qub"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) 6QubMyPqGAsZzT4HBKwlfG5HbDLftkMlgzavgqKhoyRGGDoi
//fd state: (45) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQD

ret = client.Write(fd591, []byte("vbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVx
//fd state: (6) JcY0VEQmWZ97gZ_lFu5ZMDTpQfDgtMf3A2SRN8jCPHXE0vSUek0ztF2dQFhuuziuvMiP8NhWQIcXvuD_

ret = client.Write(fd555, []byte("Dq1Yxkzp0aSunYGLmMydaYjMwUNLJ04"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) JcY0VEDq1Yxkzp0aSunYGLmMydaYjMwUNLJ048jCPHXE0vSUek0ztF2dQFhuuziuvMiP8NhWQIcXvuD_

ret = client.Close(fd563)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd605, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd598, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd611 := client.Open("/PnpsEyPcnz.txt", client.O_RDWR|client.O_CREATE)
if(fd611 < 0) {
  panic("open failed")
}


fd612 := client.Open("/hiNlZwaae9.txt", client.O_RDWR|client.O_CREATE)
if(fd612 < 0) {
  panic("open failed")
}


fd613 := client.Open("/Gr9GJJU3Ox.txt", client.O_RDWR|client.O_CREATE)
if(fd613 < 0) {
  panic("open failed")
}

//fd state: (223) C4jGvha9OUMh3QI40rW_gUYbEgLIUnssXhMsBJbmSIYgoKIRUerVh8TMPUsjBQHsyAkrHQYzrMG2mjVaSWj0kUlxWcMbDSGX1znl70J3g5QcaUAs5sV1rs0YqmIV_663fQHLj83e8uD43eeTDhL6DdXxWkOqXMKMRxjgLV6P7omiQKLoIPgdgdhza9Vv2BlMpnkyA9QnUpCL5mwFXGTpUFIeD9Yd2Al

ret = client.Write(fd545, []byte("mjV23x"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (229) C4jGvha9OUMh3QI40rW_gUYbEgLIUnssXhMsBJbmSIYgoKIRUerVh8TMPUsjBQHsyAkrHQYzrMG2mjVaSWj0kUlxWcMbDSGX1znl70J3g5QcaUAs5sV1rs0YqmIV_663fQHLj83e8uD43eeTDhL6DdXxWkOqXMKMRxjgLV6P7omiQKLoIPgdgdhza9Vv2BlMpnkyA9QnUpCL5mwFXGTpUFIeD9Yd2AlmjV23x

fd614 := client.Open("/aRXhzSthuT.txt", client.O_RDWR|client.O_CREATE)
if(fd614 < 0) {
  panic("open failed")
}

//fd state: (89) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVx

ret = client.Write(fd591, []byte("wqQiL4GYvMAfv_o55qp38IQCb5yhslNxOhywtnVEoIsK_1HTf0TCHlKf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (145) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVxwqQiL4GYvMAfv_o55qp38IQCb5yhslNxOhywtnVEoIsK_1HTf0TCHlKf
//fd state: (2) fl45Wy

ret = client.Write(fd578, []byte("KjSMalWeiSReslX63xOmSSdsKdgfflZc4sFSnbxIMCusLTw8NXOxbupd0TjUp1dU3zrdYmOltQhN5aZ3vuwAGeshdIJfJErz9zt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) flKjSMalWeiSReslX63xOmSSdsKdgfflZc4sFSnbxIMCusLTw8NXOxbupd0TjUp1dU3zrdYmOltQhN5aZ3vuwAGeshdIJfJErz9zt

ret = client.Close(fd560)
if(ret != 0) {
  panic("close failed")
}


fd615 := client.Open("/uUuh4BivWQ.txt", client.O_RDWR|client.O_CREATE)
if(fd615 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd598, []byte("f_1sI4jq2uESXFzS2OivT0QoteevtYthxDtdBO0y18e2PgkxXw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) f_1sI4jq2uESXFzS2OivT0QoteevtYthxDtdBO0y18e2PgkxXw

buf, ret = client.Read(fd532, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OmzQ3KlCi0N0vHcbvtavDyYCTEq5QVEOBZL5QFB4HhPYAm9X730ii") {
  panic("wrong data returned")
}

//fd state: (111) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhWAKtgCuG5NdBdCjAHwLWjvXuIV80ik5F1qu9QWVt9R0QbFYdzXGP8SFzZUvJ3vV0wkCp755wYKY9HldHZW30Q5xqiVKnyQsGHeGMtQUtpJxAcoBb751hnBZmszRahitRoyGtyrz5zRPEC5SEZAc1xVYbuBlYPnNP_Yld6Mv3fTp7tpKjMtofNRL7Zmuyy6UEpapWyh

ret = client.Write(fd550, []byte("9o43s2Hn20kg7ymf3pZ17JK71TzQHrMZbjMQhqSqnh7wtlmXmDCFICigt9ZzBxB5UWijhOXu530QQ4xaAPPbqh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (197) Jk9MCbWQ2lnWDRPWGFmdLJnDMYM5JjyxNgJ09KgRyEuzolDNMa0HjMVAkzNQFmGUeEncF95LzH4gilhWAKtgCuG5NdBdCjAHwLWjvXuIV80ik5F9o43s2Hn20kg7ymf3pZ17JK71TzQHrMZbjMQhqSqnh7wtlmXmDCFICigt9ZzBxB5UWijhOXu530QQ4xaAPPbqhZmszRahitRoyGtyrz5zRPEC5SEZAc1xVYbuBlYPnNP_Yld6Mv3fTp7tpKjMtofNRL7Zmuyy6UEpapWyh

ret = client.Close(fd545)
if(ret != 0) {
  panic("close failed")
}


fd616 := client.Open("/7OdNvjR3av.txt", client.O_RDWR|client.O_CREATE)
if(fd616 < 0) {
  panic("open failed")
}


ret = client.Seek(fd586, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


ret = client.Close(fd531)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd613, []byte("zv3LNJYlK3ItqY99Y8R3NRlp9joE9j3azmNVR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) zv3LNJYlK3ItqY99Y8R3NRlp9joE9j3azmNVR

fd617 := client.Open("/Z3lLNlSkX2.txt", client.O_RDWR|client.O_CREATE)
if(fd617 < 0) {
  panic("open failed")
}


ret = client.Close(fd610)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd616, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8GvXiURxNNkLgBY1dctT") {
  panic("wrong data returned")
}

//fd state: (50) f_1sI4jq2uESXFzS2OivT0QoteevtYthxDtdBO0y18e2PgkxXw

ret = client.Write(fd598, []byte("rlJblkVoQZX6F5_8XXbKWj3BOCmIdEbBGruba2yTvFBV9Ol2Bb11sWhkOKMS49j1tqSocek2fU73flEgL4QcuMSJhYC1a3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) f_1sI4jq2uESXFzS2OivT0QoteevtYthxDtdBO0y18e2PgkxXwrlJblkVoQZX6F5_8XXbKWj3BOCmIdEbBGruba2yTvFBV9Ol2Bb11sWhkOKMS49j1tqSocek2fU73flEgL4QcuMSJhYC1a3

fd618 := client.Open("/6qgmoHgYV4.txt", client.O_RDWR|client.O_CREATE)
if(fd618 < 0) {
  panic("open failed")
}


ret = client.Close(fd578)
if(ret != 0) {
  panic("close failed")
}

//fd state: (37) JcY0VEDq1Yxkzp0aSunYGLmMydaYjMwUNLJ048jCPHXE0vSUek0ztF2dQFhuuziuvMiP8NhWQIcXvuD_

ret = client.Write(fd555, []byte("x2gFFX4M7CY0399qc6Q_5f5IFoUIW95W22K0kGLBOs5qZu0IxUdV9NHN9ReriJ0rGAKGqxGGX3BNfwpz_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) JcY0VEDq1Yxkzp0aSunYGLmMydaYjMwUNLJ04x2gFFX4M7CY0399qc6Q_5f5IFoUIW95W22K0kGLBOs5qZu0IxUdV9NHN9ReriJ0rGAKGqxGGX3BNfwpz_
//fd state: (9) V87tZtyV9WfFurSBpKU3uam2K

ret = client.Write(fd586, []byte("xWofyktksrz4Zstx0bTNsk9np1IKPMJLzW_2BpCGR9q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) V87tZtyV9xWofyktksrz4Zstx0bTNsk9np1IKPMJLzW_2BpCGR9q

ret = client.Close(fd594)
if(ret != 0) {
  panic("close failed")
}


fd619 := client.Open("/GIaSRCiPEQ.txt", client.O_RDWR|client.O_CREATE)
if(fd619 < 0) {
  panic("open failed")
}


fd620 := client.Open("/Lzrygg4HmZ.txt", client.O_RDWR|client.O_CREATE)
if(fd620 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd555, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd605, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd607, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd607, []byte("whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3

ret = client.Close(fd579)
if(ret != 0) {
  panic("close failed")
}

//fd state: (67) emmvFzI2cSWgqSgtWzNwuPQoF6pP9c6U0ME9mVFE9mGlT_8hGAJ1AZRIsrMmqDWKu5D

ret = client.Write(fd600, []byte("f66YoKmMr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) emmvFzI2cSWgqSgtWzNwuPQoF6pP9c6U0ME9mVFE9mGlT_8hGAJ1AZRIsrMmqDWKu5Df66YoKmMr
//fd state: (145) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVxwqQiL4GYvMAfv_o55qp38IQCb5yhslNxOhywtnVEoIsK_1HTf0TCHlKf

ret = client.Write(fd591, []byte("qZauBlWTJe8MPGWzhqWzvhjCcOdArVaolmjQfDO3w8EH_DRJT1QRlO3WOfwYMoXcxbaQj_c"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (216) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVxwqQiL4GYvMAfv_o55qp38IQCb5yhslNxOhywtnVEoIsK_1HTf0TCHlKfqZauBlWTJe8MPGWzhqWzvhjCcOdArVaolmjQfDO3w8EH_DRJT1QRlO3WOfwYMoXcxbaQj_c

buf, ret = client.Read(fd607, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd621 := client.Open("/6mwWpmkirH.txt", client.O_RDWR|client.O_CREATE)
if(fd621 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd606, []byte("Rf81FF_mf8CAz78GK6tiTpxpTHRBmSMMU7u420nOqLZfX8Isq61dU9JjlC0s9YpvUBrvcRP7KsLvB2KxNyMqMUqjiyF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) Rf81FF_mf8CAz78GK6tiTpxpTHRBmSMMU7u420nOqLZfX8Isq61dU9JjlC0s9YpvUBrvcRP7KsLvB2KxNyMqMUqjiyF
//fd state: (0) 

ret = client.Write(fd620, []byte("rBTNVfhFuPBJUbwOJAfLA9Dktv73gRpn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) rBTNVfhFuPBJUbwOJAfLA9Dktv73gRpn

ret = client.Close(fd617)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd615, []byte("0UnqNjsZmvNooZtQ80sm1U_1RVwDW1RReVf0_WBNbQa5PWfKQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) 0UnqNjsZmvNooZtQ80sm1U_1RVwDW1RReVf0_WBNbQa5PWfKQ

fd622 := client.Open("/2VkBncInDy.txt", client.O_RDWR|client.O_CREATE)
if(fd622 < 0) {
  panic("open failed")
}


ret = client.Close(fd559)
if(ret != 0) {
  panic("close failed")
}

//fd state: (4) 6QubMyPqGAsZzT4HBKwlfG5HbDLftkMlgzavgqKhoyRGGDoi

ret = client.Write(fd596, []byte("tKapVqdHzryXTI9WnlbR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) 6QubtKapVqdHzryXTI9WnlbRbDLftkMlgzavgqKhoyRGGDoi

fd623 := client.Open("/NOndoHIH7p.txt", client.O_RDWR|client.O_CREATE)
if(fd623 < 0) {
  panic("open failed")
}

//fd state: (35) 47qbtVXel8wNdD3bB_FyXyNJvNSlb1sgt0W

ret = client.Write(fd581, []byte("7l04pRgpmIahxJhRver82m9g9cP7YWXU_bxbh7Rwwjkwm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) 47qbtVXel8wNdD3bB_FyXyNJvNSlb1sgt0W7l04pRgpmIahxJhRver82m9g9cP7YWXU_bxbh7Rwwjkwm

fd624 := client.Open("/ZKNWUF4TFV.txt", client.O_RDWR|client.O_CREATE)
if(fd624 < 0) {
  panic("open failed")
}


ret = client.Close(fd619)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd586, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd611, []byte("_Aaok7RJrWVtc04_7zjlAfwv72p62iH_OPyw5Jutzd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) _Aaok7RJrWVtc04_7zjlAfwv72p62iH_OPyw5Jutzd

ret = client.Close(fd593)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd605, []byte("dssJOEJW8cHn7ZRs1gpwUYyIjZXSIU2XVfk1vkXDy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) dssJOEJW8cHn7ZRs1gpwUYyIjZXSIU2XVfk1vkXDy

buf, ret = client.Read(fd605, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd615)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd548, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "f1gSNkTSDPXNl7Afqeiw0zCSweDGAvMrL6oLP1LxOpNLsbdCKtSgvGq_PyHpx") {
  panic("wrong data returned")
}


ret = client.Close(fd609)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd620)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd577)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd572, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd625 := client.Open("/FF3gQ3UNK5.txt", client.O_RDWR|client.O_CREATE)
if(fd625 < 0) {
  panic("open failed")
}


ret = client.Close(fd606)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd600, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Seek(fd550, 209, client.SEEK_SET)
if(ret != 209) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 209)
  panic("seek failed")
}


ret = client.Close(fd532)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd581)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd604, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd614, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd618, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd626 := client.Open("/uJt5Xeux53.txt", client.O_RDWR|client.O_CREATE)
if(fd626 < 0) {
  panic("open failed")
}


ret = client.Close(fd614)
if(ret != 0) {
  panic("close failed")
}


fd627 := client.Open("/OXdp0KPXSy.txt", client.O_RDWR|client.O_CREATE)
if(fd627 < 0) {
  panic("open failed")
}


fd628 := client.Open("/pLIWsilfjy.txt", client.O_RDWR|client.O_CREATE)
if(fd628 < 0) {
  panic("open failed")
}


ret = client.Seek(fd612, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


fd629 := client.Open("/jUx6c9_7OL.txt", client.O_RDWR|client.O_CREATE)
if(fd629 < 0) {
  panic("open failed")
}


ret = client.Close(fd605)
if(ret != 0) {
  panic("close failed")
}

//fd state: (42) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3

ret = client.Write(fd607, []byte("FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukI

buf, ret = client.Read(fd624, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd548)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd624, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd586, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd630 := client.Open("/xWWnpgOpGI.txt", client.O_RDWR|client.O_CREATE)
if(fd630 < 0) {
  panic("open failed")
}

//fd state: (216) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVxwqQiL4GYvMAfv_o55qp38IQCb5yhslNxOhywtnVEoIsK_1HTf0TCHlKfqZauBlWTJe8MPGWzhqWzvhjCcOdArVaolmjQfDO3w8EH_DRJT1QRlO3WOfwYMoXcxbaQj_c

ret = client.Write(fd591, []byte("KWanC_UIvr5yl5KMwyuZNtCSh2PWBNC4jUjXdN7mYny72mw7h42dmRfuaBtaPpeloIdqBnyRGnLz9zzpz8xq_vII"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (304) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVxwqQiL4GYvMAfv_o55qp38IQCb5yhslNxOhywtnVEoIsK_1HTf0TCHlKfqZauBlWTJe8MPGWzhqWzvhjCcOdArVaolmjQfDO3w8EH_DRJT1QRlO3WOfwYMoXcxbaQj_cKWanC_UIvr5yl5KMwyuZNtCSh2PWBNC4jUjXdN7mYny72mw7h42dmRfuaBtaPpeloIdqBnyRGnLz9zzpz8xq_vII

fd631 := client.Open("/aUFQvgjfrD.txt", client.O_RDWR|client.O_CREATE)
if(fd631 < 0) {
  panic("open failed")
}


ret = client.Close(fd613)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd622, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd591, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd555)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd629, []byte("eaCd_kAy7wiptb3T7_GOQictuRo4w45GCya_8w73F9Q2aQzh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) eaCd_kAy7wiptb3T7_GOQictuRo4w45GCya_8w73F9Q2aQzh

fd632 := client.Open("/yL4UDzcUj0.txt", client.O_RDWR|client.O_CREATE)
if(fd632 < 0) {
  panic("open failed")
}


ret = client.Seek(fd598, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Seek(fd604, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd626, []byte("3Udrlv348DOd7ipdzCKL8KUeTwzmJZsX7S144SUXGgbuznyPWH0V3pHDgBgfjolcL9EzbDurNrLhnsmnl19"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) 3Udrlv348DOd7ipdzCKL8KUeTwzmJZsX7S144SUXGgbuznyPWH0V3pHDgBgfjolcL9EzbDurNrLhnsmnl19

ret = client.Seek(fd628, 197, client.SEEK_SET)
if(ret != 197) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 197)
  panic("seek failed")
}


fd633 := client.Open("/QhEtwrbUoj.txt", client.O_RDWR|client.O_CREATE)
if(fd633 < 0) {
  panic("open failed")
}


ret = client.Close(fd633)
if(ret != 0) {
  panic("close failed")
}

//fd state: (3) WPPWZBwopU6eLU8CmN8ea9zXXyLbKNA03wUDdU

ret = client.Write(fd604, []byte("enVX0x86jytVVGV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) WPPenVX0x86jytVVGV8ea9zXXyLbKNA03wUDdU

buf, ret = client.Read(fd607, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd621, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bWKhlV") {
  panic("wrong data returned")
}


ret = client.Seek(fd625, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd634 := client.Open("/mnUQa19UWV.txt", client.O_RDWR|client.O_CREATE)
if(fd634 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd572, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd626)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd604, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Seek(fd629, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


buf, ret = client.Read(fd591, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd572, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd629, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


buf, ret = client.Read(fd625, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd635 := client.Open("/NdZKxMANLT.txt", client.O_RDWR|client.O_CREATE)
if(fd635 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd623, []byte("jD5QkUK5Fxo74PEy_sZBFizVZLMprpclALuuDyZ0VVS0TCyoJ0dC28ma"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) jD5QkUK5Fxo74PEy_sZBFizVZLMprpclALuuDyZ0VVS0TCyoJ0dC28ma

ret = client.Close(fd621)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd591, 103, client.SEEK_SET)
if(ret != 103) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 103)
  panic("seek failed")
}


ret = client.Seek(fd618, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd636 := client.Open("/uA4d9BQc55.txt", client.O_RDWR|client.O_CREATE)
if(fd636 < 0) {
  panic("open failed")
}


ret = client.Seek(fd623, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (104) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukI

ret = client.Write(fd607, []byte("QepR1DJwdn0zt8WEaC0HH8wbMDTmU0JL50fa6vMxb3vijIBkCFRz3mDVG6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (162) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HH8wbMDTmU0JL50fa6vMxb3vijIBkCFRz3mDVG6

fd637 := client.Open("/6Okbu3K6Za.txt", client.O_RDWR|client.O_CREATE)
if(fd637 < 0) {
  panic("open failed")
}


ret = client.Seek(fd618, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd586)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd596)
if(ret != 0) {
  panic("close failed")
}


fd638 := client.Open("/hjh0_HB7R_.txt", client.O_RDWR|client.O_CREATE)
if(fd638 < 0) {
  panic("open failed")
}


fd639 := client.Open("/LgfkcGUS96.txt", client.O_RDWR|client.O_CREATE)
if(fd639 < 0) {
  panic("open failed")
}


fd640 := client.Open("/k7qto9muCa.txt", client.O_RDWR|client.O_CREATE)
if(fd640 < 0) {
  panic("open failed")
}


ret = client.Seek(fd629, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


ret = client.Seek(fd591, 175, client.SEEK_SET)
if(ret != 175) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 175)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd630, []byte("QOXk1y9YCg6eAMSIkhZyy6_FiqwniTlVzcWNNKFJT6JmQ7751C32J5JiciZ50Mkftz5cTZnxjReDLv13xN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) QOXk1y9YCg6eAMSIkhZyy6_FiqwniTlVzcWNNKFJT6JmQ7751C32J5JiciZ50Mkftz5cTZnxjReDLv13xN
//fd state: (0) 

ret = client.Write(fd622, []byte("he43vcT_HjK40DrctlZ_N3uvRARCX2fOSUJVbESqAN289vFXnv8Hh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) he43vcT_HjK40DrctlZ_N3uvRARCX2fOSUJVbESqAN289vFXnv8Hh

fd641 := client.Open("/CY9cjfit5S.txt", client.O_RDWR|client.O_CREATE)
if(fd641 < 0) {
  panic("open failed")
}


fd642 := client.Open("/nrzyvFLTgm.txt", client.O_RDWR|client.O_CREATE)
if(fd642 < 0) {
  panic("open failed")
}

//fd state: (0) P3ubaJxHXF34Evj2EBqoimsaDQekRbTVpnouuQIG8jhjaOuCd97VFjeOf19uSNvxdiiKexYEs29seGsZHdtlaEp3qZmjPVlgAsCP1XKr28jRpCtsLz0awumh_O9a0dqCNOD

ret = client.Write(fd635, []byte("finui6Y4w4Pf0iYVOzxI4c8jszGZzYcyW9wpiqA7jt_cByIUbTdpJtrEj2zOFbq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) finui6Y4w4Pf0iYVOzxI4c8jszGZzYcyW9wpiqA7jt_cByIUbTdpJtrEj2zOFbqxdiiKexYEs29seGsZHdtlaEp3qZmjPVlgAsCP1XKr28jRpCtsLz0awumh_O9a0dqCNOD
//fd state: (175) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVxwqQiL4GYvMAfv_o55qp38IQCb5yhslNxOhywtnVEoIsK_1HTf0TCHlKfqZauBlWTJe8MPGWzhqWzvhjCcOdArVaolmjQfDO3w8EH_DRJT1QRlO3WOfwYMoXcxbaQj_cKWanC_UIvr5yl5KMwyuZNtCSh2PWBNC4jUjXdN7mYny72mw7h42dmRfuaBtaPpeloIdqBnyRGnLz9zzpz8xq_vII

ret = client.Write(fd591, []byte("eJuUTPJ3o3B2Y_HAQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (192) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVxwqQiL4GYvMAfv_o55qp38IQCb5yhslNxOhywtnVEoIsK_1HTf0TCHlKfqZauBlWTJe8MPGWzhqWzvhjCcOdArVeJuUTPJ3o3B2Y_HAQJT1QRlO3WOfwYMoXcxbaQj_cKWanC_UIvr5yl5KMwyuZNtCSh2PWBNC4jUjXdN7mYny72mw7h42dmRfuaBtaPpeloIdqBnyRGnLz9zzpz8xq_vII

buf, ret = client.Read(fd550, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Gtyrz5zRPEC5SEZAc1xVYbuBlYPnNP_Yld6Mv3fTp7tpKjMtofNRL7Zmuyy6UEpapWyh") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd604, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "tVVGV8ea") {
  panic("wrong data returned")
}


ret = client.Seek(fd622, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}

//fd state: (197) SCxPeJ7fvt2xZNehgjyHHnVMJ1smXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UBJpYL903agqPgfabhjNfq6rwRhscbXruOCeMpO7YqcF_tF5q8iSRheJPIGYZFwMCchFvcOUPuamzHVP3oCLHmKBr71EEl2Wb8jUXXif2arhF

ret = client.Write(fd628, []byte("RkWqRai7q8qgzuhIfhrzo03mK1gXVaMYHXrnl5YCZWPpYz4SK9Iih7kJbOwmmR5DlpYLqGAwuJk3h2JYE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (278) SCxPeJ7fvt2xZNehgjyHHnVMJ1smXxtkVXc52o4QjAHclnaBRtuqAR43GWSmQfUCrykF5S4zYHyAsOrPp6c4mffC1UBJpYL903agqPgfabhjNfq6rwRhscbXruOCeMpO7YqcF_tF5q8iSRheJPIGYZFwMCchFvcOUPuamzHVP3oCLHmKBr71EEl2Wb8jUXXif2arhRkWqRai7q8qgzuhIfhrzo03mK1gXVaMYHXrnl5YCZWPpYz4SK9Iih7kJbOwmmR5DlpYLqGAwuJk3h2JYE
//fd state: (0) 

ret = client.Write(fd632, []byte("HWlDiJBk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) HWlDiJBk
//fd state: (0) 

ret = client.Write(fd618, []byte("Qtx3MsNzTWzF8fAXKuxSOGPG6e6Vu6BCGZm2XgMC_M60TJRb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) Qtx3MsNzTWzF8fAXKuxSOGPG6e6Vu6BCGZm2XgMC_M60TJRb

ret = client.Seek(fd616, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}

//fd state: (7) f_1sI4jq2uESXFzS2OivT0QoteevtYthxDtdBO0y18e2PgkxXwrlJblkVoQZX6F5_8XXbKWj3BOCmIdEbBGruba2yTvFBV9Ol2Bb11sWhkOKMS49j1tqSocek2fU73flEgL4QcuMSJhYC1a3

ret = client.Write(fd598, []byte("viEqzDfKjBzd1eQp9Ygs2_Wxgy_TfMT4jp1y4y0LfH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) f_1sI4jviEqzDfKjBzd1eQp9Ygs2_Wxgy_TfMT4jp1y4y0LfHwrlJblkVoQZX6F5_8XXbKWj3BOCmIdEbBGruba2yTvFBV9Ol2Bb11sWhkOKMS49j1tqSocek2fU73flEgL4QcuMSJhYC1a3

fd643 := client.Open("/qnDjTQPirU.txt", client.O_RDWR|client.O_CREATE)
if(fd643 < 0) {
  panic("open failed")
}


ret = client.Seek(fd631, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd641, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TWPh4AOrSRcZF3jzVSrzzUBymE2UQBT2jQcK85goHH0SBz6") {
  panic("wrong data returned")
}


ret = client.Seek(fd634, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd644 := client.Open("/524wjnkVl2.txt", client.O_RDWR|client.O_CREATE)
if(fd644 < 0) {
  panic("open failed")
}


ret = client.Close(fd629)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd622)
if(ret != 0) {
  panic("close failed")
}


fd645 := client.Open("/w0loISNtkF.txt", client.O_RDWR|client.O_CREATE)
if(fd645 < 0) {
  panic("open failed")
}


ret = client.Seek(fd616, 83, client.SEEK_SET)
if(ret != 83) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 83)
  panic("seek failed")
}


buf, ret = client.Read(fd644, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd634, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd598, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wrlJblkVoQZX6F5_8XXbKWj3BOCmIdEbBGruba2yTvFBV9Ol2Bb11sWhkOKMS49j1tqSocek2fU") {
  panic("wrong data returned")
}

//fd state: (8) HWlDiJBk

ret = client.Write(fd632, []byte("t9JrWdNGCMerYwyErMoUbXbV0soVIW8cD4TVK2viO8uzdCBGuF8hwdO_6oYewA4wUys14GANY4f"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) HWlDiJBkt9JrWdNGCMerYwyErMoUbXbV0soVIW8cD4TVK2viO8uzdCBGuF8hwdO_6oYewA4wUys14GANY4f

buf, ret = client.Read(fd623, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jD5QkUK5Fxo74PEy_sZBFizVZLMprpclALuuDyZ0VVS0TCyoJ0dC28ma") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd616, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hJBacA1Dyc") {
  panic("wrong data returned")
}


ret = client.Close(fd616)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd607, 112, client.SEEK_SET)
if(ret != 112) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 112)
  panic("seek failed")
}

//fd state: (0) 4vAbqobHbGl7Q1RVVIsacmOLD82G7TcBdh5T9cbqXTwhlOCjrdZynJdT4B5sVcapMilIi4F5FwBX

ret = client.Write(fd636, []byte("W4oV29VOlei_6gSa8bMzKnLx8vX0k77z3xkNoTPZEex2KWqZP5JRgZbt7CcwvUGuRXVYIpRuYaKZNHiJN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) W4oV29VOlei_6gSa8bMzKnLx8vX0k77z3xkNoTPZEex2KWqZP5JRgZbt7CcwvUGuRXVYIpRuYaKZNHiJN

ret = client.Seek(fd550, 206, client.SEEK_SET)
if(ret != 206) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 206)
  panic("seek failed")
}

//fd state: (14) RTHwW7LaSlo5SQ8nEir9l4ZDsk7l

ret = client.Write(fd612, []byte("x4xJFGuK0aZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) RTHwW7LaSlo5SQx4xJFGuK0aZk7l

ret = client.Seek(fd643, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (82) QOXk1y9YCg6eAMSIkhZyy6_FiqwniTlVzcWNNKFJT6JmQ7751C32J5JiciZ50Mkftz5cTZnxjReDLv13xN

ret = client.Write(fd630, []byte("FIFOFH40oItt9Hy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) QOXk1y9YCg6eAMSIkhZyy6_FiqwniTlVzcWNNKFJT6JmQ7751C32J5JiciZ50Mkftz5cTZnxjReDLv13xNFIFOFH40oItt9Hy

buf, ret = client.Read(fd611, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (56) jD5QkUK5Fxo74PEy_sZBFizVZLMprpclALuuDyZ0VVS0TCyoJ0dC28ma

ret = client.Write(fd623, []byte("w5gOLD4hEkeXJ1aZIuWcxkc5UPBZMR6NJ1vTTCETIiSAw_uOymo2x85B4lm4tLVfNED_rOLekhc1VCVyBG5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (139) jD5QkUK5Fxo74PEy_sZBFizVZLMprpclALuuDyZ0VVS0TCyoJ0dC28maw5gOLD4hEkeXJ1aZIuWcxkc5UPBZMR6NJ1vTTCETIiSAw_uOymo2x85B4lm4tLVfNED_rOLekhc1VCVyBG5

ret = client.Seek(fd600, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Close(fd645)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd600)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd642, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd646 := client.Open("/K1q8zuNqur.txt", client.O_RDWR|client.O_CREATE)
if(fd646 < 0) {
  panic("open failed")
}


ret = client.Seek(fd625, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd647 := client.Open("/4SS3LdR3oc.txt", client.O_RDWR|client.O_CREATE)
if(fd647 < 0) {
  panic("open failed")
}


ret = client.Seek(fd591, 258, client.SEEK_SET)
if(ret != 258) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 258)
  panic("seek failed")
}


ret = client.Seek(fd624, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd648 := client.Open("/ZSbjChrY7Q.txt", client.O_RDWR|client.O_CREATE)
if(fd648 < 0) {
  panic("open failed")
}


fd649 := client.Open("/CzVjEdkRbn.txt", client.O_RDWR|client.O_CREATE)
if(fd649 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd648, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd623, 71, client.SEEK_SET)
if(ret != 71) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 71)
  panic("seek failed")
}


fd650 := client.Open("/Jo3nqa5FqY.txt", client.O_RDWR|client.O_CREATE)
if(fd650 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd598, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "73flEgL4QcuMSJhYC1") {
  panic("wrong data returned")
}


fd651 := client.Open("/hIqrxvPj2h.txt", client.O_RDWR|client.O_CREATE)
if(fd651 < 0) {
  panic("open failed")
}


ret = client.Seek(fd572, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd604, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9zXXyLbKNA03wUDdU") {
  panic("wrong data returned")
}


fd652 := client.Open("/qxUE320XyK.txt", client.O_RDWR|client.O_CREATE)
if(fd652 < 0) {
  panic("open failed")
}


ret = client.Close(fd642)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) bADDD6hFBrkBtYr6ILPQrFP_jkh7vpknRtqdgwJSB0An2iwXn7FZqWJy5boLze1lt3Yu8zNmgojxwBDiKQixqhuv4KSawP0NOgWaUUFrdK

ret = client.Write(fd646, []byte("QgolGjoGzFeUXoUPoG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) QgolGjoGzFeUXoUPoGPQrFP_jkh7vpknRtqdgwJSB0An2iwXn7FZqWJy5boLze1lt3Yu8zNmgojxwBDiKQixqhuv4KSawP0NOgWaUUFrdK

buf, ret = client.Read(fd598, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "a3") {
  panic("wrong data returned")
}


ret = client.Close(fd643)
if(ret != 0) {
  panic("close failed")
}


fd653 := client.Open("/K74_aNYkXG.txt", client.O_RDWR|client.O_CREATE)
if(fd653 < 0) {
  panic("open failed")
}


ret = client.Close(fd625)
if(ret != 0) {
  panic("close failed")
}

//fd state: (71) jD5QkUK5Fxo74PEy_sZBFizVZLMprpclALuuDyZ0VVS0TCyoJ0dC28maw5gOLD4hEkeXJ1aZIuWcxkc5UPBZMR6NJ1vTTCETIiSAw_uOymo2x85B4lm4tLVfNED_rOLekhc1VCVyBG5

ret = client.Write(fd623, []byte("LC_RJZcbBAJIqnatznkV32Dxdt_DReaOfu5SSCaNz7Anr9t5GouDh1q1ttsitnzohIf8nTiYQoNRk0er199zNPX09s6h6rNl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (167) jD5QkUK5Fxo74PEy_sZBFizVZLMprpclALuuDyZ0VVS0TCyoJ0dC28maw5gOLD4hEkeXJ1aLC_RJZcbBAJIqnatznkV32Dxdt_DReaOfu5SSCaNz7Anr9t5GouDh1q1ttsitnzohIf8nTiYQoNRk0er199zNPX09s6h6rNl

ret = client.Close(fd627)
if(ret != 0) {
  panic("close failed")
}

//fd state: (258) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVxwqQiL4GYvMAfv_o55qp38IQCb5yhslNxOhywtnVEoIsK_1HTf0TCHlKfqZauBlWTJe8MPGWzhqWzvhjCcOdArVeJuUTPJ3o3B2Y_HAQJT1QRlO3WOfwYMoXcxbaQj_cKWanC_UIvr5yl5KMwyuZNtCSh2PWBNC4jUjXdN7mYny72mw7h42dmRfuaBtaPpeloIdqBnyRGnLz9zzpz8xq_vII

ret = client.Write(fd591, []byte("FVorUwMV6BFHFQpKHwjYBK2X_cUipvLg7IjTD_zwn6yRE_1xl5A7_jpUekVkZhjH_KQdto"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (328) ARSeE37k28ICn9dXkrNPstRjCqtkbSdNNcKf7KNOHdAQDvbfmnWDOOna7UUGLXvLICkhzsb_SZ_VtmyhMk0tPVVVxwqQiL4GYvMAfv_o55qp38IQCb5yhslNxOhywtnVEoIsK_1HTf0TCHlKfqZauBlWTJe8MPGWzhqWzvhjCcOdArVeJuUTPJ3o3B2Y_HAQJT1QRlO3WOfwYMoXcxbaQj_cKWanC_UIvr5yl5KMwyuZNtCSh2PWBNC4jUjXdN7mYnFVorUwMV6BFHFQpKHwjYBK2X_cUipvLg7IjTD_zwn6yRE_1xl5A7_jpUekVkZhjH_KQdto

ret = client.Close(fd630)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd634, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd654 := client.Open("/2X6IcMWzKN.txt", client.O_RDWR|client.O_CREATE)
if(fd654 < 0) {
  panic("open failed")
}


fd655 := client.Open("/xNTmiz2SlF.txt", client.O_RDWR|client.O_CREATE)
if(fd655 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd647, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd648, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd656 := client.Open("/o8e4tZKDW0.txt", client.O_RDWR|client.O_CREATE)
if(fd656 < 0) {
  panic("open failed")
}


ret = client.Seek(fd639, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd637, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd649)
if(ret != 0) {
  panic("close failed")
}

//fd state: (81) W4oV29VOlei_6gSa8bMzKnLx8vX0k77z3xkNoTPZEex2KWqZP5JRgZbt7CcwvUGuRXVYIpRuYaKZNHiJN

ret = client.Write(fd636, []byte("WYlUe5xt0EZtIukZjUfhu8zr3d24s9qwqVKoO7NC0lVvchK2pg_K4DiyvswH505aqA3mgW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) W4oV29VOlei_6gSa8bMzKnLx8vX0k77z3xkNoTPZEex2KWqZP5JRgZbt7CcwvUGuRXVYIpRuYaKZNHiJNWYlUe5xt0EZtIukZjUfhu8zr3d24s9qwqVKoO7NC0lVvchK2pg_K4DiyvswH505aqA3mgW

buf, ret = client.Read(fd604, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd639)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd598)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd618, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (63) finui6Y4w4Pf0iYVOzxI4c8jszGZzYcyW9wpiqA7jt_cByIUbTdpJtrEj2zOFbqxdiiKexYEs29seGsZHdtlaEp3qZmjPVlgAsCP1XKr28jRpCtsLz0awumh_O9a0dqCNOD

ret = client.Write(fd635, []byte("mJLeLZljg8mQhhBJkMXh4sQoXDuysqrWUb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) finui6Y4w4Pf0iYVOzxI4c8jszGZzYcyW9wpiqA7jt_cByIUbTdpJtrEj2zOFbqmJLeLZljg8mQhhBJkMXh4sQoXDuysqrWUbsCP1XKr28jRpCtsLz0awumh_O9a0dqCNOD

buf, ret = client.Read(fd644, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd591, 215, client.SEEK_SET)
if(ret != 215) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 215)
  panic("seek failed")
}


fd657 := client.Open("/gnoGHKgG_o.txt", client.O_RDWR|client.O_CREATE)
if(fd657 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd652, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd641, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5m3LXQAVVYPIcCF3CuRZaPNnNOCPtJeVm6nu_6tSssh_6eQpXAmCjBpzt") {
  panic("wrong data returned")
}


ret = client.Close(fd638)
if(ret != 0) {
  panic("close failed")
}


fd658 := client.Open("/jEJT1wijAu.txt", client.O_RDWR|client.O_CREATE)
if(fd658 < 0) {
  panic("open failed")
}


fd659 := client.Open("/jSFBpvzIzk.txt", client.O_RDWR|client.O_CREATE)
if(fd659 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd655, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd624, []byte("QHQDczF1Od9O30iwJ6lfFQPU0S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) QHQDczF1Od9O30iwJ6lfFQPU0S

fd660 := client.Open("/Kbe_5esbc2.txt", client.O_RDWR|client.O_CREATE)
if(fd660 < 0) {
  panic("open failed")
}


ret = client.Close(fd656)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd654, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd661 := client.Open("/CKytIwcwjN.txt", client.O_RDWR|client.O_CREATE)
if(fd661 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd644, []byte("2We7zj5TawiC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) 2We7zj5TawiC
//fd state: (38) WPPenVX0x86jytVVGV8ea9zXXyLbKNA03wUDdU

ret = client.Write(fd604, []byte("NlBgW1XBiE6pOIGu2l0d84lPEhgJeB3sWcu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) WPPenVX0x86jytVVGV8ea9zXXyLbKNA03wUDdUNlBgW1XBiE6pOIGu2l0d84lPEhgJeB3sWcu

fd662 := client.Open("/bCLvtka1QB.txt", client.O_RDWR|client.O_CREATE)
if(fd662 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd611, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd607, 156, client.SEEK_SET)
if(ret != 156) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 156)
  panic("seek failed")
}


ret = client.Seek(fd655, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd631, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd591)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd650, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd607, 138, client.SEEK_SET)
if(ret != 138) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 138)
  panic("seek failed")
}


ret = client.Close(fd650)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd632)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd623)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd634, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (97) finui6Y4w4Pf0iYVOzxI4c8jszGZzYcyW9wpiqA7jt_cByIUbTdpJtrEj2zOFbqmJLeLZljg8mQhhBJkMXh4sQoXDuysqrWUbsCP1XKr28jRpCtsLz0awumh_O9a0dqCNOD

ret = client.Write(fd635, []byte("c8_iF3DZnZpxGOKqeLVcho5JzOcAQWiC7KsmCn7uJYrqUuaP4i63Co6qPCy9Iiafv3DLAWZIvsMfPlp1EKtO4KZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (184) finui6Y4w4Pf0iYVOzxI4c8jszGZzYcyW9wpiqA7jt_cByIUbTdpJtrEj2zOFbqmJLeLZljg8mQhhBJkMXh4sQoXDuysqrWUbc8_iF3DZnZpxGOKqeLVcho5JzOcAQWiC7KsmCn7uJYrqUuaP4i63Co6qPCy9Iiafv3DLAWZIvsMfPlp1EKtO4KZ

ret = client.Seek(fd611, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


fd663 := client.Open("/S4kyEU_r6C.txt", client.O_RDWR|client.O_CREATE)
if(fd663 < 0) {
  panic("open failed")
}


ret = client.Close(fd612)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd653)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd550, 265, client.SEEK_SET)
if(ret != 265) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 265)
  panic("seek failed")
}

//fd state: (0) GIp4NHMorX5tdaNhSUCZKDD2ySzdBjolxE3ipMsXnyTFWYZ5ZEzj4NrXIIq51TZJKbgr2FIrE_GbveLcZ8OwIF1SMxMtoM8sWO5ySxhpZGqIxRCEeYAOzDzeVQQfJv6mZjOG

ret = client.Write(fd659, []byte("ZWphTbVSOaeTcPBOO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) ZWphTbVSOaeTcPBOOUCZKDD2ySzdBjolxE3ipMsXnyTFWYZ5ZEzj4NrXIIq51TZJKbgr2FIrE_GbveLcZ8OwIF1SMxMtoM8sWO5ySxhpZGqIxRCEeYAOzDzeVQQfJv6mZjOG

ret = client.Close(fd659)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd641, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oQcZ_SIDd60WEtOaVJn9yLdpxJk5HAup2INVzSFLw") {
  panic("wrong data returned")
}


fd664 := client.Open("/giw08CQH59.txt", client.O_RDWR|client.O_CREATE)
if(fd664 < 0) {
  panic("open failed")
}


ret = client.Seek(fd611, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd640, []byte("ihba8LKEl1A8SuPG2cspXZUllz1hyGZqqQ_G8fF2SYiG1M7B"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) ihba8LKEl1A8SuPG2cspXZUllz1hyGZqqQ_G8fF2SYiG1M7B

ret = client.Seek(fd661, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (12) 2We7zj5TawiC

ret = client.Write(fd644, []byte("MoYPSAl0WJCfXk1B5fXgi2ayra9xiR0KH2R_wp3gqXljcUrBU8515eaSXik2J4PbUrr_1Iee8MDZdTwUCMRVxV5wBid335"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (106) 2We7zj5TawiCMoYPSAl0WJCfXk1B5fXgi2ayra9xiR0KH2R_wp3gqXljcUrBU8515eaSXik2J4PbUrr_1Iee8MDZdTwUCMRVxV5wBid335

fd665 := client.Open("/5PS4vuVPFd.txt", client.O_RDWR|client.O_CREATE)
if(fd665 < 0) {
  panic("open failed")
}


ret = client.Close(fd637)
if(ret != 0) {
  panic("close failed")
}


fd666 := client.Open("/LoIyZd2AM5.txt", client.O_RDWR|client.O_CREATE)
if(fd666 < 0) {
  panic("open failed")
}

//fd state: (106) 2We7zj5TawiCMoYPSAl0WJCfXk1B5fXgi2ayra9xiR0KH2R_wp3gqXljcUrBU8515eaSXik2J4PbUrr_1Iee8MDZdTwUCMRVxV5wBid335

ret = client.Write(fd644, []byte("pqlczkoVQV4uD6MtVkf9vidx9tyakCR897yMGBo4AadAyLeO4JBPBflYpO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) 2We7zj5TawiCMoYPSAl0WJCfXk1B5fXgi2ayra9xiR0KH2R_wp3gqXljcUrBU8515eaSXik2J4PbUrr_1Iee8MDZdTwUCMRVxV5wBid335pqlczkoVQV4uD6MtVkf9vidx9tyakCR897yMGBo4AadAyLeO4JBPBflYpO

ret = client.Close(fd660)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd662)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd640)
if(ret != 0) {
  panic("close failed")
}

//fd state: (138) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HH8wbMDTmU0JL50fa6vMxb3vijIBkCFRz3mDVG6

ret = client.Write(fd607, []byte("WgeVjzU5HJ6EgyGFSEuAHhWlQP51g503g6YI0_6YGN6V_oVblf9Uz64tW8NXzUbiEmVmQDHGVs3YHbEcRUFByNN4yYrrxzzf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (234) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HH8wbMDTmU0JL50WgeVjzU5HJ6EgyGFSEuAHhWlQP51g503g6YI0_6YGN6V_oVblf9Uz64tW8NXzUbiEmVmQDHGVs3YHbEcRUFByNN4yYrrxzzf

ret = client.Close(fd631)
if(ret != 0) {
  panic("close failed")
}


fd667 := client.Open("/JqZAq55249.txt", client.O_RDWR|client.O_CREATE)
if(fd667 < 0) {
  panic("open failed")
}


ret = client.Close(fd624)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) hyQ3IxuGte64OvndFOSNKV8JKpJZIjIA5GzDcAwtEdNEeu19vq89P

ret = client.Write(fd663, []byte("mQhc_h6tMkGQtYFotkCJOlYvIdH6xWz6_JYnjmfdCKqODXvezw8w0JZdoUOGDvd0dbDZtIMxdVZ9WYyt5j1ly_Ph"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) mQhc_h6tMkGQtYFotkCJOlYvIdH6xWz6_JYnjmfdCKqODXvezw8w0JZdoUOGDvd0dbDZtIMxdVZ9WYyt5j1ly_Ph

ret = client.Close(fd657)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd634)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd667, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


buf, ret = client.Read(fd651, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (3) 9rMJdkn

ret = client.Write(fd661, []byte("x2b8cUOlQpxvKo99TE7k0L"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) 9rMx2b8cUOlQpxvKo99TE7k0L

ret = client.Seek(fd607, 158, client.SEEK_SET)
if(ret != 158) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 158)
  panic("seek failed")
}


fd668 := client.Open("/obmRgRRVEq.txt", client.O_RDWR|client.O_CREATE)
if(fd668 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd663, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd646, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PQrFP_jkh7vpknRtqdgwJSB0An2iwXn7FZqWJy5boLze1lt3Yu8zNmgojxwBDiKQixqhuv4KSawP0NOgWaUUFrdK") {
  panic("wrong data returned")
}


fd669 := client.Open("/KqykqQdiXF.txt", client.O_RDWR|client.O_CREATE)
if(fd669 < 0) {
  panic("open failed")
}

//fd state: (151) W4oV29VOlei_6gSa8bMzKnLx8vX0k77z3xkNoTPZEex2KWqZP5JRgZbt7CcwvUGuRXVYIpRuYaKZNHiJNWYlUe5xt0EZtIukZjUfhu8zr3d24s9qwqVKoO7NC0lVvchK2pg_K4DiyvswH505aqA3mgW

ret = client.Write(fd636, []byte("B_9EzKoxOPeu2Qtv2MXNsOv3KTdRt7c99HaSf2ytwOWS8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (196) W4oV29VOlei_6gSa8bMzKnLx8vX0k77z3xkNoTPZEex2KWqZP5JRgZbt7CcwvUGuRXVYIpRuYaKZNHiJNWYlUe5xt0EZtIukZjUfhu8zr3d24s9qwqVKoO7NC0lVvchK2pg_K4DiyvswH505aqA3mgWB_9EzKoxOPeu2Qtv2MXNsOv3KTdRt7c99HaSf2ytwOWS8

ret = client.Seek(fd572, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd651)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd668, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ouGfKY8mYskJ02lCVrK2MjR_") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd635, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd644, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


fd670 := client.Open("/rF4pC11txc.txt", client.O_RDWR|client.O_CREATE)
if(fd670 < 0) {
  panic("open failed")
}


fd671 := client.Open("/K15NYmej8E.txt", client.O_RDWR|client.O_CREATE)
if(fd671 < 0) {
  panic("open failed")
}


ret = client.Close(fd644)
if(ret != 0) {
  panic("close failed")
}

//fd state: (106) QgolGjoGzFeUXoUPoGPQrFP_jkh7vpknRtqdgwJSB0An2iwXn7FZqWJy5boLze1lt3Yu8zNmgojxwBDiKQixqhuv4KSawP0NOgWaUUFrdK

ret = client.Write(fd646, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (106) QgolGjoGzFeUXoUPoGPQrFP_jkh7vpknRtqdgwJSB0An2iwXn7FZqWJy5boLze1lt3Yu8zNmgojxwBDiKQixqhuv4KSawP0NOgWaUUFrdK

buf, ret = client.Read(fd661, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd604, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd607, 233, client.SEEK_SET)
if(ret != 233) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 233)
  panic("seek failed")
}


ret = client.Seek(fd668, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


ret = client.Seek(fd655, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd664, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd672 := client.Open("/jDmBD5iFAg.txt", client.O_RDWR|client.O_CREATE)
if(fd672 < 0) {
  panic("open failed")
}

//fd state: (0) XqeR5BQTtne3Y1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3Uo

ret = client.Write(fd669, []byte("RkOPoRw9bfDF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) RkOPoRw9bfDFY1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3Uo

buf, ret = client.Read(fd618, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd618)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd667, 68, client.SEEK_SET)
if(ret != 68) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 68)
  panic("seek failed")
}


buf, ret = client.Read(fd669, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Y1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3Uo") {
  panic("wrong data returned")
}


fd673 := client.Open("/lK1Ra888uV.txt", client.O_RDWR|client.O_CREATE)
if(fd673 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd572, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd628)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd648, []byte("OCwERG9N3NstKNw3n12ntrYhXsIJsJJwr9RuD11MuKnuVUDDt_6yjdPKEiHQQsB5Lq1GXKm7GO89xtBQTixAnmKFQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) OCwERG9N3NstKNw3n12ntrYhXsIJsJJwr9RuD11MuKnuVUDDt_6yjdPKEiHQQsB5Lq1GXKm7GO89xtBQTixAnmKFQ

ret = client.Close(fd673)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd654, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd550, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uyy6UEpapWyh") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd635, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd674 := client.Open("/bppjE5LOC1.txt", client.O_RDWR|client.O_CREATE)
if(fd674 < 0) {
  panic("open failed")
}


ret = client.Close(fd670)
if(ret != 0) {
  panic("close failed")
}


fd675 := client.Open("/jg0Jvu7lJx.txt", client.O_RDWR|client.O_CREATE)
if(fd675 < 0) {
  panic("open failed")
}


ret = client.Close(fd674)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd671, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd676 := client.Open("/K15NYmej8E.txt", client.O_RDWR|client.O_CREATE)
if(fd676 < 0) {
  panic("open failed")
}

//fd state: (196) W4oV29VOlei_6gSa8bMzKnLx8vX0k77z3xkNoTPZEex2KWqZP5JRgZbt7CcwvUGuRXVYIpRuYaKZNHiJNWYlUe5xt0EZtIukZjUfhu8zr3d24s9qwqVKoO7NC0lVvchK2pg_K4DiyvswH505aqA3mgWB_9EzKoxOPeu2Qtv2MXNsOv3KTdRt7c99HaSf2ytwOWS8

ret = client.Write(fd636, []byte("OblKdE62R41TYoNuCuWciy2Fp_ncePLOygseg_TAveI0RcswqqT2lFA0sLCNNq5mgU8zKn6teqtQJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (273) W4oV29VOlei_6gSa8bMzKnLx8vX0k77z3xkNoTPZEex2KWqZP5JRgZbt7CcwvUGuRXVYIpRuYaKZNHiJNWYlUe5xt0EZtIukZjUfhu8zr3d24s9qwqVKoO7NC0lVvchK2pg_K4DiyvswH505aqA3mgWB_9EzKoxOPeu2Qtv2MXNsOv3KTdRt7c99HaSf2ytwOWS8OblKdE62R41TYoNuCuWciy2Fp_ncePLOygseg_TAveI0RcswqqT2lFA0sLCNNq5mgU8zKn6teqtQJ

ret = client.Seek(fd671, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd663, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


ret = client.Seek(fd572, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd668, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


fd677 := client.Open("/Ju0L1mzSEa.txt", client.O_RDWR|client.O_CREATE)
if(fd677 < 0) {
  panic("open failed")
}


fd678 := client.Open("/uhxp52Eb32.txt", client.O_RDWR|client.O_CREATE)
if(fd678 < 0) {
  panic("open failed")
}


ret = client.Close(fd658)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd663)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) i

ret = client.Write(fd666, []byte("AnQUCz9d53gaqtPUREz6T41Sywrq4ppC_KMog0SbJdqbBaRZciYTDg21W17DTOythd_LpcXQx0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) AnQUCz9d53gaqtPUREz6T41Sywrq4ppC_KMog0SbJdqbBaRZciYTDg21W17DTOythd_LpcXQx0
//fd state: (0) 

ret = client.Write(fd671, []byte("eUbJ8CgHPDEiuMc3nimT7SvbuXJ8RaN0lNnASKoCfxDa8NpJmtIqjI_Y67L48qpinscp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) eUbJ8CgHPDEiuMc3nimT7SvbuXJ8RaN0lNnASKoCfxDa8NpJmtIqjI_Y67L48qpinscp

buf, ret = client.Read(fd550, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (233) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HH8wbMDTmU0JL50WgeVjzU5HJ6EgyGFSEuAHhWlQP51g503g6YI0_6YGN6V_oVblf9Uz64tW8NXzUbiEmVmQDHGVs3YHbEcRUFByNN4yYrrxzzf

ret = client.Write(fd607, []byte("rBuSGjGJZQbEc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (246) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HH8wbMDTmU0JL50WgeVjzU5HJ6EgyGFSEuAHhWlQP51g503g6YI0_6YGN6V_oVblf9Uz64tW8NXzUbiEmVmQDHGVs3YHbEcRUFByNN4yYrrxzzrBuSGjGJZQbEc

ret = client.Seek(fd635, 104, client.SEEK_SET)
if(ret != 104) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 104)
  panic("seek failed")
}


fd679 := client.Open("/BoeHZjH9NR.txt", client.O_RDWR|client.O_CREATE)
if(fd679 < 0) {
  panic("open failed")
}


fd680 := client.Open("/SltVBA_XyO.txt", client.O_RDWR|client.O_CREATE)
if(fd680 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd678, []byte("_eoJHSuWFZctWXY8oT4artmWIKjnEOEihGhmnndzKoVP_ZqqdVDYLkpeaepkM4UqTN0CAESnkA1QxuOsXbpRjRYf5Ww2I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) _eoJHSuWFZctWXY8oT4artmWIKjnEOEihGhmnndzKoVP_ZqqdVDYLkpeaepkM4UqTN0CAESnkA1QxuOsXbpRjRYf5Ww2I

fd681 := client.Open("/pzfyZZwnSF.txt", client.O_RDWR|client.O_CREATE)
if(fd681 < 0) {
  panic("open failed")
}


ret = client.Close(fd647)
if(ret != 0) {
  panic("close failed")
}

//fd state: (74) AnQUCz9d53gaqtPUREz6T41Sywrq4ppC_KMog0SbJdqbBaRZciYTDg21W17DTOythd_LpcXQx0

ret = client.Write(fd666, []byte("jpuXHxY46lfj62xCh2JUd2tuS2ThjcEyQvfm0lQpVZRiTqv_eLjapM8FSOAuQ85ywh6VFAowP2yp3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) AnQUCz9d53gaqtPUREz6T41Sywrq4ppC_KMog0SbJdqbBaRZciYTDg21W17DTOythd_LpcXQx0jpuXHxY46lfj62xCh2JUd2tuS2ThjcEyQvfm0lQpVZRiTqv_eLjapM8FSOAuQ85ywh6VFAowP2yp3

fd682 := client.Open("/5QlPGMS3lm.txt", client.O_RDWR|client.O_CREATE)
if(fd682 < 0) {
  panic("open failed")
}


fd683 := client.Open("/x5DumFEOhh.txt", client.O_RDWR|client.O_CREATE)
if(fd683 < 0) {
  panic("open failed")
}

//fd state: (68) eUbJ8CgHPDEiuMc3nimT7SvbuXJ8RaN0lNnASKoCfxDa8NpJmtIqjI_Y67L48qpinscp

ret = client.Write(fd671, []byte("Q7IDqFVVcIOuy8KV4hlR_Oj6DAz5mfvVuB8_JJBoUvNr6OwtnV69iMZQ7OE3zJzfBKUO7i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (138) eUbJ8CgHPDEiuMc3nimT7SvbuXJ8RaN0lNnASKoCfxDa8NpJmtIqjI_Y67L48qpinscpQ7IDqFVVcIOuy8KV4hlR_Oj6DAz5mfvVuB8_JJBoUvNr6OwtnV69iMZQ7OE3zJzfBKUO7i

ret = client.Close(fd667)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd636, 125, client.SEEK_SET)
if(ret != 125) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 125)
  panic("seek failed")
}


fd684 := client.Open("/YWJtqDTpBK.txt", client.O_RDWR|client.O_CREATE)
if(fd684 < 0) {
  panic("open failed")
}


fd685 := client.Open("/_CVl1cUh_V.txt", client.O_RDWR|client.O_CREATE)
if(fd685 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd572, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd655, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd685, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Ioy8k2Jqygrzm2KZuBfY9AX3TJwOYkdcqR3geGxPOu_frIHYX2MR") {
  panic("wrong data returned")
}


fd686 := client.Open("/k7CcNoLmri.txt", client.O_RDWR|client.O_CREATE)
if(fd686 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd664, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (52) Ioy8k2Jqygrzm2KZuBfY9AX3TJwOYkdcqR3geGxPOu_frIHYX2MR

ret = client.Write(fd685, []byte("XNoz91CjQhUR0IzU3qHsMVuu3ZHN40E3bp0WpKh_uSHl6B17tP70mIQL_VQGYk3Huzm7tQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) Ioy8k2Jqygrzm2KZuBfY9AX3TJwOYkdcqR3geGxPOu_frIHYX2MRXNoz91CjQhUR0IzU3qHsMVuu3ZHN40E3bp0WpKh_uSHl6B17tP70mIQL_VQGYk3Huzm7tQ
//fd state: (125) W4oV29VOlei_6gSa8bMzKnLx8vX0k77z3xkNoTPZEex2KWqZP5JRgZbt7CcwvUGuRXVYIpRuYaKZNHiJNWYlUe5xt0EZtIukZjUfhu8zr3d24s9qwqVKoO7NC0lVvchK2pg_K4DiyvswH505aqA3mgWB_9EzKoxOPeu2Qtv2MXNsOv3KTdRt7c99HaSf2ytwOWS8OblKdE62R41TYoNuCuWciy2Fp_ncePLOygseg_TAveI0RcswqqT2lFA0sLCNNq5mgU8zKn6teqtQJ

ret = client.Write(fd636, []byte("ugVgxAlYNJekNE2P7kKPunrV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (149) W4oV29VOlei_6gSa8bMzKnLx8vX0k77z3xkNoTPZEex2KWqZP5JRgZbt7CcwvUGuRXVYIpRuYaKZNHiJNWYlUe5xt0EZtIukZjUfhu8zr3d24s9qwqVKoO7NC0lVvugVgxAlYNJekNE2P7kKPunrVgWB_9EzKoxOPeu2Qtv2MXNsOv3KTdRt7c99HaSf2ytwOWS8OblKdE62R41TYoNuCuWciy2Fp_ncePLOygseg_TAveI0RcswqqT2lFA0sLCNNq5mgU8zKn6teqtQJ

buf, ret = client.Read(fd550, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd675)
if(ret != 0) {
  panic("close failed")
}

//fd state: (151) AnQUCz9d53gaqtPUREz6T41Sywrq4ppC_KMog0SbJdqbBaRZciYTDg21W17DTOythd_LpcXQx0jpuXHxY46lfj62xCh2JUd2tuS2ThjcEyQvfm0lQpVZRiTqv_eLjapM8FSOAuQ85ywh6VFAowP2yp3

ret = client.Write(fd666, []byte("d6vi7rOOXTsZ8h4vT2mfvZufMFRqixkZo3Y6JFtdM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (192) AnQUCz9d53gaqtPUREz6T41Sywrq4ppC_KMog0SbJdqbBaRZciYTDg21W17DTOythd_LpcXQx0jpuXHxY46lfj62xCh2JUd2tuS2ThjcEyQvfm0lQpVZRiTqv_eLjapM8FSOAuQ85ywh6VFAowP2yp3d6vi7rOOXTsZ8h4vT2mfvZufMFRqixkZo3Y6JFtdM
//fd state: (0) 

ret = client.Write(fd682, []byte("Wlz_UQR4wubMYfWWDccgvufGj96H60DdSOg2T"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) Wlz_UQR4wubMYfWWDccgvufGj96H60DdSOg2T
//fd state: (89) OCwERG9N3NstKNw3n12ntrYhXsIJsJJwr9RuD11MuKnuVUDDt_6yjdPKEiHQQsB5Lq1GXKm7GO89xtBQTixAnmKFQ

ret = client.Write(fd648, []byte("n9HGWsM8qfTkRWAA3JFXsrDc0dxOBzhlwxhiclpI01MJPLUcvoO7eBYd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (145) OCwERG9N3NstKNw3n12ntrYhXsIJsJJwr9RuD11MuKnuVUDDt_6yjdPKEiHQQsB5Lq1GXKm7GO89xtBQTixAnmKFQn9HGWsM8qfTkRWAA3JFXsrDc0dxOBzhlwxhiclpI01MJPLUcvoO7eBYd
//fd state: (0) eUbJ8CgHPDEiuMc3nimT7SvbuXJ8RaN0lNnASKoCfxDa8NpJmtIqjI_Y67L48qpinscpQ7IDqFVVcIOuy8KV4hlR_Oj6DAz5mfvVuB8_JJBoUvNr6OwtnV69iMZQ7OE3zJzfBKUO7i

ret = client.Write(fd676, []byte("GpAaUHLWOZF3XbxkSZDdqYPsCB1i3yEfE5gF2L4pWscI8BRN7M0fiZVEii_3IQmYEkXG3nMC6UTqQy61GU_WYdEzcizcrec"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) GpAaUHLWOZF3XbxkSZDdqYPsCB1i3yEfE5gF2L4pWscI8BRN7M0fiZVEii_3IQmYEkXG3nMC6UTqQy61GU_WYdEzcizcrec5mfvVuB8_JJBoUvNr6OwtnV69iMZQ7OE3zJzfBKUO7i

ret = client.Seek(fd686, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd687 := client.Open("/_bppGnywyK.txt", client.O_RDWR|client.O_CREATE)
if(fd687 < 0) {
  panic("open failed")
}


ret = client.Seek(fd683, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd666, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd664, []byte("RkjrwizD1xB6MCo_FKjPrPb_KLECFPVNqsWFtawGdqH3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) RkjrwizD1xB6MCo_FKjPrPb_KLECFPVNqsWFtawGdqH3

buf, ret = client.Read(fd682, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd679)
if(ret != 0) {
  panic("close failed")
}


fd688 := client.Open("/gFR56ynYpp.txt", client.O_RDWR|client.O_CREATE)
if(fd688 < 0) {
  panic("open failed")
}


ret = client.Close(fd654)
if(ret != 0) {
  panic("close failed")
}


fd689 := client.Open("/6zCWkaOFR7.txt", client.O_RDWR|client.O_CREATE)
if(fd689 < 0) {
  panic("open failed")
}


ret = client.Seek(fd655, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd690 := client.Open("/MaHqGlDtgN.txt", client.O_RDWR|client.O_CREATE)
if(fd690 < 0) {
  panic("open failed")
}


ret = client.Seek(fd688, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd691 := client.Open("/TtMUgwwetl.txt", client.O_RDWR|client.O_CREATE)
if(fd691 < 0) {
  panic("open failed")
}

//fd state: (74) ouGfKY8mYskJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848KFs5HnTK2gJ6hjt2YX0HFktN8fRAEuxrNEh3v4G_28ADSi5CZvYhDbqGzVgYNDA7kByhYDEjuzQa6ns9fWSeI9eKnHZDoFqpLY7FNCD

ret = client.Write(fd668, []byte("EaY8r2XhmstUQmVGNRfOedNdDzKkkQfzE49bsBAFe61Eoxu4jYGZpIoM6QfVWSWN2aYgR4je914ZXS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) ouGfKY8mYskJ02lCVrK2MjR_ItlVqTaSs5mYvqO8Mizf1sfwjS4cwG848KFs5HnTK2gJ6hjt2YEaY8r2XhmstUQmVGNRfOedNdDzKkkQfzE49bsBAFe61Eoxu4jYGZpIoM6QfVWSWN2aYgR4je914ZXSpLY7FNCD
//fd state: (145) OCwERG9N3NstKNw3n12ntrYhXsIJsJJwr9RuD11MuKnuVUDDt_6yjdPKEiHQQsB5Lq1GXKm7GO89xtBQTixAnmKFQn9HGWsM8qfTkRWAA3JFXsrDc0dxOBzhlwxhiclpI01MJPLUcvoO7eBYd

ret = client.Write(fd648, []byte("RQe9BIYexbJCFK6juVmUas34fC2cRM4aV6M6Y5JDFQ1wb2ojpYA9qWAYnwK6UgMxUGSIfZ1efBPgzH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (223) OCwERG9N3NstKNw3n12ntrYhXsIJsJJwr9RuD11MuKnuVUDDt_6yjdPKEiHQQsB5Lq1GXKm7GO89xtBQTixAnmKFQn9HGWsM8qfTkRWAA3JFXsrDc0dxOBzhlwxhiclpI01MJPLUcvoO7eBYdRQe9BIYexbJCFK6juVmUas34fC2cRM4aV6M6Y5JDFQ1wb2ojpYA9qWAYnwK6UgMxUGSIfZ1efBPgzH

buf, ret = client.Read(fd572, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd692 := client.Open("/SbQouCL1vm.txt", client.O_RDWR|client.O_CREATE)
if(fd692 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd668, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pLY7FNCD") {
  panic("wrong data returned")
}


ret = client.Close(fd684)
if(ret != 0) {
  panic("close failed")
}

//fd state: (93) _eoJHSuWFZctWXY8oT4artmWIKjnEOEihGhmnndzKoVP_ZqqdVDYLkpeaepkM4UqTN0CAESnkA1QxuOsXbpRjRYf5Ww2I

ret = client.Write(fd678, []byte("Ds5rNtR2GtOrdzmOGnaUMGnwEOEmNgS87wmnF79rBtPlA1I96dLWLWXGjuE6jB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (155) _eoJHSuWFZctWXY8oT4artmWIKjnEOEihGhmnndzKoVP_ZqqdVDYLkpeaepkM4UqTN0CAESnkA1QxuOsXbpRjRYf5Ww2IDs5rNtR2GtOrdzmOGnaUMGnwEOEmNgS87wmnF79rBtPlA1I96dLWLWXGjuE6jB

ret = client.Seek(fd688, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd641)
if(ret != 0) {
  panic("close failed")
}

//fd state: (192) AnQUCz9d53gaqtPUREz6T41Sywrq4ppC_KMog0SbJdqbBaRZciYTDg21W17DTOythd_LpcXQx0jpuXHxY46lfj62xCh2JUd2tuS2ThjcEyQvfm0lQpVZRiTqv_eLjapM8FSOAuQ85ywh6VFAowP2yp3d6vi7rOOXTsZ8h4vT2mfvZufMFRqixkZo3Y6JFtdM

ret = client.Write(fd666, []byte("O0GWaVaiXPPYYntHJhRDJeVCyXbNPqKEXO3HrrCvswiBVg65"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (240) AnQUCz9d53gaqtPUREz6T41Sywrq4ppC_KMog0SbJdqbBaRZciYTDg21W17DTOythd_LpcXQx0jpuXHxY46lfj62xCh2JUd2tuS2ThjcEyQvfm0lQpVZRiTqv_eLjapM8FSOAuQ85ywh6VFAowP2yp3d6vi7rOOXTsZ8h4vT2mfvZufMFRqixkZo3Y6JFtdMO0GWaVaiXPPYYntHJhRDJeVCyXbNPqKEXO3HrrCvswiBVg65

ret = client.Close(fd550)
if(ret != 0) {
  panic("close failed")
}


fd693 := client.Open("/h9euK4ZjF8.txt", client.O_RDWR|client.O_CREATE)
if(fd693 < 0) {
  panic("open failed")
}


ret = client.Seek(fd676, 109, client.SEEK_SET)
if(ret != 109) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 109)
  panic("seek failed")
}


ret = client.Seek(fd688, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd694 := client.Open("/EqVLObc5V3.txt", client.O_RDWR|client.O_CREATE)
if(fd694 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd694, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd685, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


fd695 := client.Open("/uJt5Xeux53.txt", client.O_RDWR|client.O_CREATE)
if(fd695 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd694, []byte("59Mg2KkenfOfujIT11D090bjr0H2uLg_HISdlnCM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) 59Mg2KkenfOfujIT11D090bjr0H2uLg_HISdlnCM

fd696 := client.Open("/Cik6mR5AES.txt", client.O_RDWR|client.O_CREATE)
if(fd696 < 0) {
  panic("open failed")
}


ret = client.Close(fd655)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd646, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Close(fd681)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd685, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd697 := client.Open("/nzpVrSVi21.txt", client.O_RDWR|client.O_CREATE)
if(fd697 < 0) {
  panic("open failed")
}


ret = client.Close(fd664)
if(ret != 0) {
  panic("close failed")
}


fd698 := client.Open("/trssyo58wD.txt", client.O_RDWR|client.O_CREATE)
if(fd698 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd668, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd683, []byte("lMTRZO4nX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) lMTRZO4nX

ret = client.Seek(fd676, 86, client.SEEK_SET)
if(ret != 86) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 86)
  panic("seek failed")
}


fd699 := client.Open("/RTp_oPPyOZ.txt", client.O_RDWR|client.O_CREATE)
if(fd699 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd689, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "kLkORRlHF6vzmbUosRu7kPpM7dLVG7pLofh1UU8zPkreYH_kGnziQhQTdukZpuct") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd646, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "h7vpknRtqdgwJSB0An2iwXn7FZqWJy5boLze1lt3Yu8zN") {
  panic("wrong data returned")
}


ret = client.Close(fd611)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd572, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd688, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd700 := client.Open("/HN3w0IwSey.txt", client.O_RDWR|client.O_CREATE)
if(fd700 < 0) {
  panic("open failed")
}


ret = client.Seek(fd648, 213, client.SEEK_SET)
if(ret != 213) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 213)
  panic("seek failed")
}

//fd state: (0) LHBmCmBXD1LtO9RULKBxJt9NrffqHR5MMrgcydKKJhYPUu04Nowk9OnVzjNspvfIO5tJnLmdQcZLhZjasOuIlrzYPAdBsbIrnUg0RC0mIb8iK3A

ret = client.Write(fd677, []byte("XDweE5oTnZckT_DxTZik2IN36SOhl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) XDweE5oTnZckT_DxTZik2IN36SOhlR5MMrgcydKKJhYPUu04Nowk9OnVzjNspvfIO5tJnLmdQcZLhZjasOuIlrzYPAdBsbIrnUg0RC0mIb8iK3A

buf, ret = client.Read(fd686, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd701 := client.Open("/7rXbyPPk1Z.txt", client.O_RDWR|client.O_CREATE)
if(fd701 < 0) {
  panic("open failed")
}


ret = client.Seek(fd682, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


buf, ret = client.Read(fd691, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd661)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd680, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd604, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd700, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


ret = client.Seek(fd607, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) bs66QRjD1Jho6BTqkwmT8VJnK21h9AXO4ygrC_xu93Eo3V1ue1fkR99cxyv_7uH_HpLc50V8HufJztyunGc251vQqXkSwWxjzstgpKiD0Z1VvXkUDCTCUZrxg15u36JIFHKt7JXFbQzOKa0HNxhmeYDQe7CHUCe53mHaNzN7ZexrUIDUYm2XTKaaCOBdKa95WQv2OTQN0WqJyGe3pvCUyJduZ6RLyuDGB38QeycozKXBHYSm5wfNxjmDoAEcmt0QOInVrvfXveG3ZM5

ret = client.Write(fd698, []byte("NcIRjMc5eXeJjry6F11mqi_u4gPVuO3pBfv7ivRtb7fCMUe6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) NcIRjMc5eXeJjry6F11mqi_u4gPVuO3pBfv7ivRtb7fCMUe6e1fkR99cxyv_7uH_HpLc50V8HufJztyunGc251vQqXkSwWxjzstgpKiD0Z1VvXkUDCTCUZrxg15u36JIFHKt7JXFbQzOKa0HNxhmeYDQe7CHUCe53mHaNzN7ZexrUIDUYm2XTKaaCOBdKa95WQv2OTQN0WqJyGe3pvCUyJduZ6RLyuDGB38QeycozKXBHYSm5wfNxjmDoAEcmt0QOInVrvfXveG3ZM5

fd702 := client.Open("/oy_syYGpgq.txt", client.O_RDWR|client.O_CREATE)
if(fd702 < 0) {
  panic("open failed")
}


fd703 := client.Open("/TByuPJVs0i.txt", client.O_RDWR|client.O_CREATE)
if(fd703 < 0) {
  panic("open failed")
}


fd704 := client.Open("/3lWw0RID5o.txt", client.O_RDWR|client.O_CREATE)
if(fd704 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd687, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OoUXvf3V9L6yQ4sznxyOZTZ33t5cSGRp0xGslXc9nv2KQjIpWf3") {
  panic("wrong data returned")
}


ret = client.Seek(fd683, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}

//fd state: (0) Ioy8k2Jqygrzm2KZuBfY9AX3TJwOYkdcqR3geGxPOu_frIHYX2MRXNoz91CjQhUR0IzU3qHsMVuu3ZHN40E3bp0WpKh_uSHl6B17tP70mIQL_VQGYk3Huzm7tQ

ret = client.Write(fd685, []byte("Pm0RcDcMwbiQyswOVzhjNgNOpyatjXKZ4R97WGHvDLkSvjEUPKUm0pa8cYJReqLfPTCYpO3JOvFEFyE2VTIB8Gzk1mqz_VU6lPg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) Pm0RcDcMwbiQyswOVzhjNgNOpyatjXKZ4R97WGHvDLkSvjEUPKUm0pa8cYJReqLfPTCYpO3JOvFEFyE2VTIB8Gzk1mqz_VU6lPg7tP70mIQL_VQGYk3Huzm7tQ
//fd state: (0) 

ret = client.Write(fd702, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) 

fd705 := client.Open("/f_VMh3LK24.txt", client.O_RDWR|client.O_CREATE)
if(fd705 < 0) {
  panic("open failed")
}


ret = client.Close(fd702)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd704, []byte("Wl80Ay7YteKDYSQSdTXH0k_WFpu0osY_xdbkj87F4_QEvNal4d"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) Wl80Ay7YteKDYSQSdTXH0k_WFpu0osY_xdbkj87F4_QEvNal4d

ret = client.Close(fd698)
if(ret != 0) {
  panic("close failed")
}


fd706 := client.Open("/qxUE320XyK.txt", client.O_RDWR|client.O_CREATE)
if(fd706 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd572, []byte("KmjijiMRbfIhKvYBrOxFasqE4PMpfBnbRKtso4H8GIZb3W3W10o6K5IclGWUqocRd0Ek28X3Z9mqtCIlnfJgJ8WCOXuoSQFFZl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) KmjijiMRbfIhKvYBrOxFasqE4PMpfBnbRKtso4H8GIZb3W3W10o6K5IclGWUqocRd0Ek28X3Z9mqtCIlnfJgJ8WCOXuoSQFFZl

ret = client.Close(fd665)
if(ret != 0) {
  panic("close failed")
}


fd707 := client.Open("/CuUVYFzllF.txt", client.O_RDWR|client.O_CREATE)
if(fd707 < 0) {
  panic("open failed")
}


fd708 := client.Open("/IFHa_5zRVT.txt", client.O_RDWR|client.O_CREATE)
if(fd708 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd685, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7tP70mIQL_VQGYk3Huzm7tQ") {
  panic("wrong data returned")
}


ret = client.Close(fd695)
if(ret != 0) {
  panic("close failed")
}


fd709 := client.Open("/nYSzy33lDz.txt", client.O_RDWR|client.O_CREATE)
if(fd709 < 0) {
  panic("open failed")
}


ret = client.Seek(fd682, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}

//fd state: (138) GpAaUHLWOZF3XbxkSZDdqYPsCB1i3yEfE5gF2L4pWscI8BRN7M0fiZVEii_3IQmYEkXG3nMC6UTqQy61GU_WYdEzcizcrec5mfvVuB8_JJBoUvNr6OwtnV69iMZQ7OE3zJzfBKUO7i

ret = client.Write(fd671, []byte("LqO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (141) GpAaUHLWOZF3XbxkSZDdqYPsCB1i3yEfE5gF2L4pWscI8BRN7M0fiZVEii_3IQmYEkXG3nMC6UTqQy61GU_WYdEzcizcrec5mfvVuB8_JJBoUvNr6OwtnV69iMZQ7OE3zJzfBKUO7iLqO

fd710 := client.Open("/o8EgfUAG1r.txt", client.O_RDWR|client.O_CREATE)
if(fd710 < 0) {
  panic("open failed")
}


ret = client.Close(fd689)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd700)
if(ret != 0) {
  panic("close failed")
}


fd711 := client.Open("/luZC2GuVOR.txt", client.O_RDWR|client.O_CREATE)
if(fd711 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd688, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd636, 225, client.SEEK_SET)
if(ret != 225) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 225)
  panic("seek failed")
}


ret = client.Seek(fd607, 124, client.SEEK_SET)
if(ret != 124) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 124)
  panic("seek failed")
}


ret = client.Seek(fd694, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


ret = client.Seek(fd680, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd712 := client.Open("/fxr8zvWv72.txt", client.O_RDWR|client.O_CREATE)
if(fd712 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd692, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "36zkNe7kIzHoUBECFdun8cUaoZo") {
  panic("wrong data returned")
}


fd713 := client.Open("/XqiBS62H3L.txt", client.O_RDWR|client.O_CREATE)
if(fd713 < 0) {
  panic("open failed")
}


fd714 := client.Open("/ZYVT5hpnrT.txt", client.O_RDWR|client.O_CREATE)
if(fd714 < 0) {
  panic("open failed")
}

//fd state: (71) QgolGjoGzFeUXoUPoGPQrFP_jkh7vpknRtqdgwJSB0An2iwXn7FZqWJy5boLze1lt3Yu8zNmgojxwBDiKQixqhuv4KSawP0NOgWaUUFrdK

ret = client.Write(fd646, []byte("1w_AXy8Oqpt2jq7eWzw0uE50eXU_Q2MDAVKvhVh_xGBclNRwkkpz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) QgolGjoGzFeUXoUPoGPQrFP_jkh7vpknRtqdgwJSB0An2iwXn7FZqWJy5boLze1lt3Yu8zN1w_AXy8Oqpt2jq7eWzw0uE50eXU_Q2MDAVKvhVh_xGBclNRwkkpz
//fd state: (0) 

ret = client.Write(fd690, []byte("ndt8P_UIVF6a6tqeVVnk1FQ6NPoMLq0sRurPynSXCYWyEFQBpAnnxcGBcA17OvL5MYLKoUaocAEYEF2bP7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) ndt8P_UIVF6a6tqeVVnk1FQ6NPoMLq0sRurPynSXCYWyEFQBpAnnxcGBcA17OvL5MYLKoUaocAEYEF2bP7

ret = client.Close(fd604)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd711, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd685)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd703)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd691)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd686)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd707, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd707, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd715 := client.Open("/uhtzqOp7dU.txt", client.O_RDWR|client.O_CREATE)
if(fd715 < 0) {
  panic("open failed")
}


fd716 := client.Open("/osrqWqMC6C.txt", client.O_RDWR|client.O_CREATE)
if(fd716 < 0) {
  panic("open failed")
}


ret = client.Close(fd688)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd708, []byte("SL6wCRhaOaTpGp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) SL6wCRhaOaTpGp

fd717 := client.Open("/r_gM3BVTOr.txt", client.O_RDWR|client.O_CREATE)
if(fd717 < 0) {
  panic("open failed")
}


ret = client.Close(fd682)
if(ret != 0) {
  panic("close failed")
}


fd718 := client.Open("/rVqg54MwLv.txt", client.O_RDWR|client.O_CREATE)
if(fd718 < 0) {
  panic("open failed")
}


ret = client.Close(fd652)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd717, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd707, []byte("sj80xTTwM0oupgRfbswWfQV9qRS6dyBieOLSMANL_kxF2nS0_NSpnt4iy8Jvup0NsaV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) sj80xTTwM0oupgRfbswWfQV9qRS6dyBieOLSMANL_kxF2nS0_NSpnt4iy8Jvup0NsaV

buf, ret = client.Read(fd713, 6)
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


ret = client.Seek(fd678, 137, client.SEEK_SET)
if(ret != 137) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 137)
  panic("seek failed")
}


ret = client.Seek(fd706, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd711, []byte("Dtw3s7xNECDUIYobDiHJ56DAaWePBBt__mHmZl3iaWTBlW2FtN8eZZ3wuMTJLnJterWt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) Dtw3s7xNECDUIYobDiHJ56DAaWePBBt__mHmZl3iaWTBlW2FtN8eZZ3wuMTJLnJterWt

ret = client.Close(fd694)
if(ret != 0) {
  panic("close failed")
}


fd719 := client.Open("/JqZAq55249.txt", client.O_RDWR|client.O_CREATE)
if(fd719 < 0) {
  panic("open failed")
}


fd720 := client.Open("/nBwnAtneem.txt", client.O_RDWR|client.O_CREATE)
if(fd720 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd716, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd690)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd708)
if(ret != 0) {
  panic("close failed")
}


fd721 := client.Open("/TqUxY9FOiQ.txt", client.O_RDWR|client.O_CREATE)
if(fd721 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd720, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd678)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd648, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "fZ1efBPgzH") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd692, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd677, 103, client.SEEK_SET)
if(ret != 103) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 103)
  panic("seek failed")
}


fd722 := client.Open("/SlueffwKXY.txt", client.O_RDWR|client.O_CREATE)
if(fd722 < 0) {
  panic("open failed")
}


ret = client.Seek(fd711, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


ret = client.Close(fd711)
if(ret != 0) {
  panic("close failed")
}


fd723 := client.Open("/upTVLgCvxD.txt", client.O_RDWR|client.O_CREATE)
if(fd723 < 0) {
  panic("open failed")
}


ret = client.Close(fd692)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd720, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd720, []byte("c3Vh2ROhloohl_gVLECJil1Lz9RkP7nTjlt8wY8IpDZ5P9A5XkYywOQ5zE3SQjX3Y4Sfsp5Ogj1qDxWNL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) c3Vh2ROhloohl_gVLECJil1Lz9RkP7nTjlt8wY8IpDZ5P9A5XkYywOQ5zE3SQjX3Y4Sfsp5Ogj1qDxWNL
//fd state: (86) GpAaUHLWOZF3XbxkSZDdqYPsCB1i3yEfE5gF2L4pWscI8BRN7M0fiZVEii_3IQmYEkXG3nMC6UTqQy61GU_WYdEzcizcrec5mfvVuB8_JJBoUvNr6OwtnV69iMZQ7OE3zJzfBKUO7iLqO

ret = client.Write(fd676, []byte("Q95I4_jFUw3_cjHiHtnmzfSJX9upokRk1H3yN6U8gERMk3JHhbGC54lxUrXm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (146) GpAaUHLWOZF3XbxkSZDdqYPsCB1i3yEfE5gF2L4pWscI8BRN7M0fiZVEii_3IQmYEkXG3nMC6UTqQy61GU_WYdQ95I4_jFUw3_cjHiHtnmzfSJX9upokRk1H3yN6U8gERMk3JHhbGC54lxUrXm

ret = client.Seek(fd668, 96, client.SEEK_SET)
if(ret != 96) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 96)
  panic("seek failed")
}

//fd state: (0) 4tz5E04fbPh7jkJRgLQtAgDrEsu6UYWP1Pyqi0S5pQtiqQrkOBAO57zy_EBCxe1GMociBqXt2LvrwkETsWb4AzwpBhTMO6EW6oXL6jh4I2mprlTNJWurxnHHfXUVvEebXddWOIiNwkqW2E8LqlMlyhlW1ASH5MB6OMGnGrWXFR0cRRVlscsy47jrQPSlBSg3HQoYi0

ret = client.Write(fd709, []byte("_B3twDnixzSTMJh1ib3Dt7i1nTqYJOmXIQJstL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) _B3twDnixzSTMJh1ib3Dt7i1nTqYJOmXIQJstLS5pQtiqQrkOBAO57zy_EBCxe1GMociBqXt2LvrwkETsWb4AzwpBhTMO6EW6oXL6jh4I2mprlTNJWurxnHHfXUVvEebXddWOIiNwkqW2E8LqlMlyhlW1ASH5MB6OMGnGrWXFR0cRRVlscsy47jrQPSlBSg3HQoYi0

ret = client.Seek(fd677, 102, client.SEEK_SET)
if(ret != 102) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 102)
  panic("seek failed")
}


ret = client.Close(fd636)
if(ret != 0) {
  panic("close failed")
}


fd724 := client.Open("/bCLvtka1QB.txt", client.O_RDWR|client.O_CREATE)
if(fd724 < 0) {
  panic("open failed")
}


ret = client.Close(fd671)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd717, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd716, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd701)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd724, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd725 := client.Open("/XwukkFpNt2.txt", client.O_RDWR|client.O_CREATE)
if(fd725 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd680, []byte("je722qlaxOvhPJA9uuiVOONJi5ED8SfoSAM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) je722qlaxOvhPJA9uuiVOONJi5ED8SfoSAM

fd726 := client.Open("/ZYVT5hpnrT.txt", client.O_RDWR|client.O_CREATE)
if(fd726 < 0) {
  panic("open failed")
}


ret = client.Close(fd668)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd672, []byte("52ryyxV6A5qPyCglvr_Vn1kJiwgNaDUOcDpNaOCRYmCyy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) 52ryyxV6A5qPyCglvr_Vn1kJiwgNaDUOcDpNaOCRYmCyy
//fd state: (0) 

ret = client.Write(fd722, []byte("qVXC4mDnXy6Aa2uycVNhm5URrjhlX9BA4iLQrmaLI3okpqYLJq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) qVXC4mDnXy6Aa2uycVNhm5URrjhlX9BA4iLQrmaLI3okpqYLJq

ret = client.Seek(fd721, 84, client.SEEK_SET)
if(ret != 84) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 84)
  panic("seek failed")
}


buf, ret = client.Read(fd677, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "0mIb8iK3A") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd716, []byte("PIw__iil8SyiFp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) PIw__iil8SyiFp

fd727 := client.Open("/ZYVT5hpnrT.txt", client.O_RDWR|client.O_CREATE)
if(fd727 < 0) {
  panic("open failed")
}


ret = client.Seek(fd699, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd728 := client.Open("/KAn40utX3j.txt", client.O_RDWR|client.O_CREATE)
if(fd728 < 0) {
  panic("open failed")
}


ret = client.Close(fd677)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd672, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


buf, ret = client.Read(fd705, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "46CVo00yrl3np4GS1WZRdSKo60SsmTtroC3sp1NN") {
  panic("wrong data returned")
}


ret = client.Seek(fd716, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


buf, ret = client.Read(fd669, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd729 := client.Open("/BJH1VSRyX5.txt", client.O_RDWR|client.O_CREATE)
if(fd729 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd717, []byte("UkwZRGcMOgRr1y3xcjUDSWqDcjVAcB6fiqQQ2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) UkwZRGcMOgRr1y3xcjUDSWqDcjVAcB6fiqQQ2

ret = client.Seek(fd683, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd693, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


fd730 := client.Open("/3G482HgWOt.txt", client.O_RDWR|client.O_CREATE)
if(fd730 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd724, []byte("s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) s

fd731 := client.Open("/qinn9z7zrr.txt", client.O_RDWR|client.O_CREATE)
if(fd731 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd723, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (1) s

ret = client.Write(fd724, []byte("otnjUX9jgZVvt0fEajV5q_qj0aEVVUMewYydS0axEK1m4qbqq2q6rtcpHnRe47v"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) sotnjUX9jgZVvt0fEajV5q_qj0aEVVUMewYydS0axEK1m4qbqq2q6rtcpHnRe47v
//fd state: (0) 

ret = client.Write(fd710, []byte("jTrBzFdjkzCVNKiG7TFnyLh2ZCzHdHPuapGJNYHTL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) jTrBzFdjkzCVNKiG7TFnyLh2ZCzHdHPuapGJNYHTL

fd732 := client.Open("/C1ZDrk7A58.txt", client.O_RDWR|client.O_CREATE)
if(fd732 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd687, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GRYkBQlpjldcW0c8a7b1LMmdXitUYzwTdz") {
  panic("wrong data returned")
}


fd733 := client.Open("/xNTmiz2SlF.txt", client.O_RDWR|client.O_CREATE)
if(fd733 < 0) {
  panic("open failed")
}

//fd state: (98) KmjijiMRbfIhKvYBrOxFasqE4PMpfBnbRKtso4H8GIZb3W3W10o6K5IclGWUqocRd0Ek28X3Z9mqtCIlnfJgJ8WCOXuoSQFFZl

ret = client.Write(fd572, []byte("HnjUFBIIpbxjzdD4qOEQ1skmBwaqrY20S9RgzhbCoO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (140) KmjijiMRbfIhKvYBrOxFasqE4PMpfBnbRKtso4H8GIZb3W3W10o6K5IclGWUqocRd0Ek28X3Z9mqtCIlnfJgJ8WCOXuoSQFFZlHnjUFBIIpbxjzdD4qOEQ1skmBwaqrY20S9RgzhbCoO

ret = client.Close(fd680)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd676, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd707, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 4P7C6iVFRW6_6kpqDwS3Qzdd2UuCinShZI6ZSycsVLovvXhdAjrd8ULbO71d1

ret = client.Write(fd732, []byte("TDWjsklt3jebG1QNAVnvyiVKH_hjOuKB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) TDWjsklt3jebG1QNAVnvyiVKH_hjOuKBZI6ZSycsVLovvXhdAjrd8ULbO71d1

buf, ret = client.Read(fd646, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd715, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (53) RkOPoRw9bfDFY1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3Uo

ret = client.Write(fd669, []byte("dbYEm0itMYmVMNBrTmIoIBYTnJK7k8_7J14u9PTShKsXFBjRgpX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) RkOPoRw9bfDFY1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3UodbYEm0itMYmVMNBrTmIoIBYTnJK7k8_7J14u9PTShKsXFBjRgpX
//fd state: (32) TDWjsklt3jebG1QNAVnvyiVKH_hjOuKBZI6ZSycsVLovvXhdAjrd8ULbO71d1

ret = client.Write(fd732, []byte("fhNodiMqzZrqi0T_hRTW7va6K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) TDWjsklt3jebG1QNAVnvyiVKH_hjOuKBfhNodiMqzZrqi0T_hRTW7va6K71d1

buf, ret = client.Read(fd705, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bU4yOGkT34hny8LE") {
  panic("wrong data returned")
}


fd734 := client.Open("/uEvYTullIO.txt", client.O_RDWR|client.O_CREATE)
if(fd734 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd713, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd735 := client.Open("/u8ZwKWg3El.txt", client.O_RDWR|client.O_CREATE)
if(fd735 < 0) {
  panic("open failed")
}


fd736 := client.Open("/etbCfgFtZR.txt", client.O_RDWR|client.O_CREATE)
if(fd736 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd723, []byte("YcthduLLb7A6mV4nG3H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) YcthduLLb7A6mV4nG3H

buf, ret = client.Read(fd709, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "S5pQtiq") {
  panic("wrong data returned")
}


ret = client.Close(fd720)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd717, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


buf, ret = client.Read(fd710, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (124) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HH8wbMDTmU0JL50WgeVjzU5HJ6EgyGFSEuAHhWlQP51g503g6YI0_6YGN6V_oVblf9Uz64tW8NXzUbiEmVmQDHGVs3YHbEcRUFByNN4yYrrxzzrBuSGjGJZQbEc

ret = client.Write(fd607, []byte("S5RthNn0GVKF9emD4dXWbZ1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (147) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HS5RthNn0GVKF9emD4dXWbZ1J6EgyGFSEuAHhWlQP51g503g6YI0_6YGN6V_oVblf9Uz64tW8NXzUbiEmVmQDHGVs3YHbEcRUFByNN4yYrrxzzrBuSGjGJZQbEc
//fd state: (0) 

ret = client.Write(fd706, []byte("1k3sfoF8hqKkAHO2ZkqZotqHFGztzY6VbWvdwmzdocF_NuR_cQvNse08RnJqERkHI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) 1k3sfoF8hqKkAHO2ZkqZotqHFGztzY6VbWvdwmzdocF_NuR_cQvNse08RnJqERkHI

buf, ret = client.Read(fd727, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd710)
if(ret != 0) {
  panic("close failed")
}

//fd state: (19) YcthduLLb7A6mV4nG3H

ret = client.Write(fd723, []byte("YWKdLsmHh06"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) YcthduLLb7A6mV4nG3HYWKdLsmHh06

buf, ret = client.Read(fd706, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (31) 52ryyxV6A5qPyCglvr_Vn1kJiwgNaDUOcDpNaOCRYmCyy

ret = client.Write(fd672, []byte("NRERhc2ufFxIVIQXRDiDq76CuR8P3_xLzqAy1xj6_A7ZiSHmCxrrzoL_yv8n8h7ftRLFuaeqfq3kOEyoZOz035n1d"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) 52ryyxV6A5qPyCglvr_Vn1kJiwgNaDUNRERhc2ufFxIVIQXRDiDq76CuR8P3_xLzqAy1xj6_A7ZiSHmCxrrzoL_yv8n8h7ftRLFuaeqfq3kOEyoZOz035n1d

ret = client.Seek(fd719, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


fd737 := client.Open("/_laoVJpAG9.txt", client.O_RDWR|client.O_CREATE)
if(fd737 < 0) {
  panic("open failed")
}


ret = client.Close(fd729)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd730)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd713, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (45) _B3twDnixzSTMJh1ib3Dt7i1nTqYJOmXIQJstLS5pQtiqQrkOBAO57zy_EBCxe1GMociBqXt2LvrwkETsWb4AzwpBhTMO6EW6oXL6jh4I2mprlTNJWurxnHHfXUVvEebXddWOIiNwkqW2E8LqlMlyhlW1ASH5MB6OMGnGrWXFR0cRRVlscsy47jrQPSlBSg3HQoYi0

ret = client.Write(fd709, []byte("Xhk0medTHYMuLyXBBb470iUl5sgLN8tpf93LYR3sVSNLKC_Se"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) _B3twDnixzSTMJh1ib3Dt7i1nTqYJOmXIQJstLS5pQtiqXhk0medTHYMuLyXBBb470iUl5sgLN8tpf93LYR3sVSNLKC_SeEW6oXL6jh4I2mprlTNJWurxnHHfXUVvEebXddWOIiNwkqW2E8LqlMlyhlW1ASH5MB6OMGnGrWXFR0cRRVlscsy47jrQPSlBSg3HQoYi0

ret = client.Seek(fd704, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Close(fd718)
if(ret != 0) {
  panic("close failed")
}


fd738 := client.Open("/Ux0XzaWvp2.txt", client.O_RDWR|client.O_CREATE)
if(fd738 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd736, []byte("P7SUorwK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) P7SUorwK

buf, ret = client.Read(fd635, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ZnZpxGOKqeLVcho") {
  panic("wrong data returned")
}


ret = client.Close(fd716)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd727, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd739 := client.Open("/U0V545Ew8U.txt", client.O_RDWR|client.O_CREATE)
if(fd739 < 0) {
  panic("open failed")
}


ret = client.Close(fd697)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd717)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd728, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


buf, ret = client.Read(fd705, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "BXu8WpXTwzCfxhyd6SCkSSfkD8qNAWPPMkhGtDH4_uVAfYRAcT5IdoKg2") {
  panic("wrong data returned")
}


fd740 := client.Open("/ikZPJ_xCMj.txt", client.O_RDWR|client.O_CREATE)
if(fd740 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd733, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd734, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LR1SxTAWFo6F0cTfqeQn1wg1VttnpwfaT") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd736, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd722)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd713, []byte("gCoB3cq1IvGEEfzoxT7W5dtqaWlf7FhR3pIFdgn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) gCoB3cq1IvGEEfzoxT7W5dtqaWlf7FhR3pIFdgn

ret = client.Seek(fd726, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (7) GFoI48Jn_fk7HXS

ret = client.Write(fd728, []byte("8lxNSoKqYsAHFIVEqWJScfudGOXGZ4dZnMZxwPcJu_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) GFoI48J8lxNSoKqYsAHFIVEqWJScfudGOXGZ4dZnMZxwPcJu_

buf, ret = client.Read(fd733, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd635)
if(ret != 0) {
  panic("close failed")
}


fd741 := client.Open("/8W6113_G3d.txt", client.O_RDWR|client.O_CREATE)
if(fd741 < 0) {
  panic("open failed")
}


ret = client.Seek(fd733, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd672, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


ret = client.Seek(fd731, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}

//fd state: (39) gCoB3cq1IvGEEfzoxT7W5dtqaWlf7FhR3pIFdgn

ret = client.Write(fd713, []byte("epQXw_nk1Y_JEbTOp_mJJ0PGdydjQ5ZrQbtyb2AumSjjUdpKXimCKJlMWmyjHj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) gCoB3cq1IvGEEfzoxT7W5dtqaWlf7FhR3pIFdgnepQXw_nk1Y_JEbTOp_mJJ0PGdydjQ5ZrQbtyb2AumSjjUdpKXimCKJlMWmyjHj
//fd state: (0) 

ret = client.Write(fd735, []byte("gnkpqwfZDKkyc3UtoDnBPGShZyyxfhYAZXLd5enaXt_yItv10mn3FAAJYOo7Ctww"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) gnkpqwfZDKkyc3UtoDnBPGShZyyxfhYAZXLd5enaXt_yItv10mn3FAAJYOo7Ctww

ret = client.Close(fd728)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd726)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd713, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


buf, ret = client.Read(fd648, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd742 := client.Open("/zBngaNADw0.txt", client.O_RDWR|client.O_CREATE)
if(fd742 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd733, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (104) RkOPoRw9bfDFY1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3UodbYEm0itMYmVMNBrTmIoIBYTnJK7k8_7J14u9PTShKsXFBjRgpX

ret = client.Write(fd669, []byte("MCR7H4wZFSMdnWfFUojO9MP3sLp3zAx8LxqsjnjiARhyl6aUTX84AsOUN74HtbKKoaLMWejO7IYzRWRRQMzmTjvQh7OkzkrGS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (201) RkOPoRw9bfDFY1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3UodbYEm0itMYmVMNBrTmIoIBYTnJK7k8_7J14u9PTShKsXFBjRgpXMCR7H4wZFSMdnWfFUojO9MP3sLp3zAx8LxqsjnjiARhyl6aUTX84AsOUN74HtbKKoaLMWejO7IYzRWRRQMzmTjvQh7OkzkrGS

buf, ret = client.Read(fd742, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd743 := client.Open("/DqxeiYyrv3.txt", client.O_RDWR|client.O_CREATE)
if(fd743 < 0) {
  panic("open failed")
}


ret = client.Seek(fd709, 136, client.SEEK_SET)
if(ret != 136) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 136)
  panic("seek failed")
}


buf, ret = client.Read(fd712, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd744 := client.Open("/L8Y7EMLJqt.txt", client.O_RDWR|client.O_CREATE)
if(fd744 < 0) {
  panic("open failed")
}


ret = client.Seek(fd687, 85, client.SEEK_SET)
if(ret != 85) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 85)
  panic("seek failed")
}


fd745 := client.Open("/xNTmiz2SlF.txt", client.O_RDWR|client.O_CREATE)
if(fd745 < 0) {
  panic("open failed")
}


fd746 := client.Open("/Ur_2N7oyWA.txt", client.O_RDWR|client.O_CREATE)
if(fd746 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd683, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lMTRZO4nX") {
  panic("wrong data returned")
}

//fd state: (9) PGN5MB5z9CxkxTOFR1xkXBPQ4oDob0ayfYASL2QBlmkcF_CCIHnDT9jTBDCjeJOPF8R9dXyc7Kl_Q4oTKVf

ret = client.Write(fd693, []byte("HIvybpjXd58QpcMg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) PGN5MB5z9HIvybpjXd58QpcMgoDob0ayfYASL2QBlmkcF_CCIHnDT9jTBDCjeJOPF8R9dXyc7Kl_Q4oTKVf

buf, ret = client.Read(fd737, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "A9awLY38LiPZxxErjJwTOKWFl1W6") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd742, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd744)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd706, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


buf, ret = client.Read(fd731, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TGF") {
  panic("wrong data returned")
}


ret = client.Seek(fd714, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd687, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (113) 46CVo00yrl3np4GS1WZRdSKo60SsmTtroC3sp1NNbU4yOGkT34hny8LEBXu8WpXTwzCfxhyd6SCkSSfkD8qNAWPPMkhGtDH4_uVAfYRAcT5IdoKg2

ret = client.Write(fd705, []byte("hmEzfr1AlQJmX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (126) 46CVo00yrl3np4GS1WZRdSKo60SsmTtroC3sp1NNbU4yOGkT34hny8LEBXu8WpXTwzCfxhyd6SCkSSfkD8qNAWPPMkhGtDH4_uVAfYRAcT5IdoKg2hmEzfr1AlQJmX

buf, ret = client.Read(fd724, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd699, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (57) TDWjsklt3jebG1QNAVnvyiVKH_hjOuKBfhNodiMqzZrqi0T_hRTW7va6K71d1

ret = client.Write(fd732, []byte("1sNAEQRhCNkCUp7RAEaH2UkKE4Tnym_ew6mpOVmGlbTAXviTls7voa5gxOQ2OSlATtpU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (125) TDWjsklt3jebG1QNAVnvyiVKH_hjOuKBfhNodiMqzZrqi0T_hRTW7va6K1sNAEQRhCNkCUp7RAEaH2UkKE4Tnym_ew6mpOVmGlbTAXviTls7voa5gxOQ2OSlATtpU

buf, ret = client.Read(fd733, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd747 := client.Open("/kj8LVAlaZX.txt", client.O_RDWR|client.O_CREATE)
if(fd747 < 0) {
  panic("open failed")
}


ret = client.Seek(fd704, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


buf, ret = client.Read(fd669, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd713)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd733, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd742)
if(ret != 0) {
  panic("close failed")
}


fd748 := client.Open("/iugP43w9X2.txt", client.O_RDWR|client.O_CREATE)
if(fd748 < 0) {
  panic("open failed")
}


fd749 := client.Open("/tdnjjp4xp2.txt", client.O_RDWR|client.O_CREATE)
if(fd749 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd745, []byte("ZoxqSLF6dPtl1cZW4SRor9Ofk8Mg46Gs8p1w32nsMrzH_ZdduKfl7SQwu3qsOyCxLpMV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) ZoxqSLF6dPtl1cZW4SRor9Ofk8Mg46Gs8p1w32nsMrzH_ZdduKfl7SQwu3qsOyCxLpMV

ret = client.Seek(fd715, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd648, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd745)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd646)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd733, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ZoxqSLF6dPtl1cZW4SRor9Ofk8Mg46Gs8p1w32nsMrzH_ZdduKfl7SQwu3qsOyCxLpMV") {
  panic("wrong data returned")
}


fd750 := client.Open("/KLk320RUY5.txt", client.O_RDWR|client.O_CREATE)
if(fd750 < 0) {
  panic("open failed")
}

//fd state: (67) sj80xTTwM0oupgRfbswWfQV9qRS6dyBieOLSMANL_kxF2nS0_NSpnt4iy8Jvup0NsaV

ret = client.Write(fd707, []byte("KzKToMzDbXUH4aKaPO5SjBl8VotpiB8J07LkpDgIcFs6XhoVzAgvQEiNMWEl95KWPQf7GrpGbydOYyEsUSx6a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) sj80xTTwM0oupgRfbswWfQV9qRS6dyBieOLSMANL_kxF2nS0_NSpnt4iy8Jvup0NsaVKzKToMzDbXUH4aKaPO5SjBl8VotpiB8J07LkpDgIcFs6XhoVzAgvQEiNMWEl95KWPQf7GrpGbydOYyEsUSx6a

ret = client.Seek(fd737, 69, client.SEEK_SET)
if(ret != 69) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 69)
  panic("seek failed")
}


buf, ret = client.Read(fd740, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd734, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd751 := client.Open("/toGgWukH3w.txt", client.O_RDWR|client.O_CREATE)
if(fd751 < 0) {
  panic("open failed")
}


ret = client.Close(fd731)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd715, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd733, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


buf, ret = client.Read(fd736, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (64) sotnjUX9jgZVvt0fEajV5q_qj0aEVVUMewYydS0axEK1m4qbqq2q6rtcpHnRe47v

ret = client.Write(fd724, []byte("xhutyi2qYIdHKVh2GE4D3x6pNxw6nF5iqp414smJ4cbtmxQ1PKmzaog5DkxE5v8iqb2cV22YBcSTJxuwW2DgEKgX7sfi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (156) sotnjUX9jgZVvt0fEajV5q_qj0aEVVUMewYydS0axEK1m4qbqq2q6rtcpHnRe47vxhutyi2qYIdHKVh2GE4D3x6pNxw6nF5iqp414smJ4cbtmxQ1PKmzaog5DkxE5v8iqb2cV22YBcSTJxuwW2DgEKgX7sfi

ret = client.Close(fd572)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd683, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}

//fd state: (147) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HS5RthNn0GVKF9emD4dXWbZ1J6EgyGFSEuAHhWlQP51g503g6YI0_6YGN6V_oVblf9Uz64tW8NXzUbiEmVmQDHGVs3YHbEcRUFByNN4yYrrxzzrBuSGjGJZQbEc

ret = client.Write(fd607, []byte("UtlcY3X4mqHeLUetC7lBb2mfne6DfIG7CTPhRvj7vKcvVRMiuwzENHLYABqSkC7QFYiEePz9b_E3_iIua7E3rmSIbWnmbmwF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (243) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HS5RthNn0GVKF9emD4dXWbZ1UtlcY3X4mqHeLUetC7lBb2mfne6DfIG7CTPhRvj7vKcvVRMiuwzENHLYABqSkC7QFYiEePz9b_E3_iIua7E3rmSIbWnmbmwFbEc
//fd state: (0) 

ret = client.Write(fd714, []byte("AD1VDoaF9LR92pHV_YaGSXKEXddaNrlB7vVFDaf5JsgGPjO9L7JtjFAmnL3hBJgGf1s9t3KFZATyV7ymH_tooFtdG0WhKe3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) AD1VDoaF9LR92pHV_YaGSXKEXddaNrlB7vVFDaf5JsgGPjO9L7JtjFAmnL3hBJgGf1s9t3KFZATyV7ymH_tooFtdG0WhKe3
//fd state: (223) OCwERG9N3NstKNw3n12ntrYhXsIJsJJwr9RuD11MuKnuVUDDt_6yjdPKEiHQQsB5Lq1GXKm7GO89xtBQTixAnmKFQn9HGWsM8qfTkRWAA3JFXsrDc0dxOBzhlwxhiclpI01MJPLUcvoO7eBYdRQe9BIYexbJCFK6juVmUas34fC2cRM4aV6M6Y5JDFQ1wb2ojpYA9qWAYnwK6UgMxUGSIfZ1efBPgzH

ret = client.Write(fd648, []byte("zhe8Z5OHPPD1XFAAnO3mUHU3cPzjpLrOqxFGecMtBGhi8X1FPiFcTgrcEUeNquBrnvSi1M17w1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (297) OCwERG9N3NstKNw3n12ntrYhXsIJsJJwr9RuD11MuKnuVUDDt_6yjdPKEiHQQsB5Lq1GXKm7GO89xtBQTixAnmKFQn9HGWsM8qfTkRWAA3JFXsrDc0dxOBzhlwxhiclpI01MJPLUcvoO7eBYdRQe9BIYexbJCFK6juVmUas34fC2cRM4aV6M6Y5JDFQ1wb2ojpYA9qWAYnwK6UgMxUGSIfZ1efBPgzHzhe8Z5OHPPD1XFAAnO3mUHU3cPzjpLrOqxFGecMtBGhi8X1FPiFcTgrcEUeNquBrnvSi1M17w1

ret = client.Seek(fd714, 76, client.SEEK_SET)
if(ret != 76) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 76)
  panic("seek failed")
}


buf, ret = client.Read(fd747, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (85) OoUXvf3V9L6yQ4sznxyOZTZ33t5cSGRp0xGslXc9nv2KQjIpWf3GRYkBQlpjldcW0c8a7b1LMmdXitUYzwTdz

ret = client.Write(fd687, []byte("XuXx9fB4sxZjv1WzxYrxOsXj4DPojok6_rABoeYIQKiWl_j5jhWu2wKwrz2OF64j4fzvP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) OoUXvf3V9L6yQ4sznxyOZTZ33t5cSGRp0xGslXc9nv2KQjIpWf3GRYkBQlpjldcW0c8a7b1LMmdXitUYzwTdzXuXx9fB4sxZjv1WzxYrxOsXj4DPojok6_rABoeYIQKiWl_j5jhWu2wKwrz2OF64j4fzvP

fd752 := client.Open("/JPr0cEv57D.txt", client.O_RDWR|client.O_CREATE)
if(fd752 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd752, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd696)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd704, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


buf, ret = client.Read(fd669, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd719, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}

//fd state: (33) LR1SxTAWFo6F0cTfqeQn1wg1VttnpwfaT

ret = client.Write(fd734, []byte("hA9kNObJpgQmqhKpA0tW4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) LR1SxTAWFo6F0cTfqeQn1wg1VttnpwfaThA9kNObJpgQmqhKpA0tW4
//fd state: (0) 

ret = client.Write(fd712, []byte("4Me44ZFDHHDeATwq4X8AVmOLWr2D5Trg63DManK4hU0vSwn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) 4Me44ZFDHHDeATwq4X8AVmOLWr2D5Trg63DManK4hU0vSwn

buf, ret = client.Read(fd669, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd715, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd753 := client.Open("/6rEUgk3ohz.txt", client.O_RDWR|client.O_CREATE)
if(fd753 < 0) {
  panic("open failed")
}


fd754 := client.Open("/buR7_Zg0Os.txt", client.O_RDWR|client.O_CREATE)
if(fd754 < 0) {
  panic("open failed")
}


ret = client.Seek(fd705, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Close(fd753)
if(ret != 0) {
  panic("close failed")
}


fd755 := client.Open("/_fKtkmMG2H.txt", client.O_RDWR|client.O_CREATE)
if(fd755 < 0) {
  panic("open failed")
}

//fd state: (146) GpAaUHLWOZF3XbxkSZDdqYPsCB1i3yEfE5gF2L4pWscI8BRN7M0fiZVEii_3IQmYEkXG3nMC6UTqQy61GU_WYdQ95I4_jFUw3_cjHiHtnmzfSJX9upokRk1H3yN6U8gERMk3JHhbGC54lxUrXm

ret = client.Write(fd676, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (146) GpAaUHLWOZF3XbxkSZDdqYPsCB1i3yEfE5gF2L4pWscI8BRN7M0fiZVEii_3IQmYEkXG3nMC6UTqQy61GU_WYdQ95I4_jFUw3_cjHiHtnmzfSJX9upokRk1H3yN6U8gERMk3JHhbGC54lxUrXm

fd756 := client.Open("/9mO8qrylzT.txt", client.O_RDWR|client.O_CREATE)
if(fd756 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd715, []byte("oKwXawtdpw5MGzllKOP2maFGYtt12a6aMUtZV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) oKwXawtdpw5MGzllKOP2maFGYtt12a6aMUtZV

buf, ret = client.Read(fd699, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd743, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd743)
if(ret != 0) {
  panic("close failed")
}


fd757 := client.Open("/lr6giCQsli.txt", client.O_RDWR|client.O_CREATE)
if(fd757 < 0) {
  panic("open failed")
}


ret = client.Seek(fd725, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd693)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd755)
if(ret != 0) {
  panic("close failed")
}

//fd state: (156) sotnjUX9jgZVvt0fEajV5q_qj0aEVVUMewYydS0axEK1m4qbqq2q6rtcpHnRe47vxhutyi2qYIdHKVh2GE4D3x6pNxw6nF5iqp414smJ4cbtmxQ1PKmzaog5DkxE5v8iqb2cV22YBcSTJxuwW2DgEKgX7sfi

ret = client.Write(fd724, []byte("0qNxw9_Eqdwr29AAFAlkMZWeNAjpZEQ2haIv5xXzYggHeituYFG7CI7xK_1fLuG795Rb0VVBBIKvf0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (234) sotnjUX9jgZVvt0fEajV5q_qj0aEVVUMewYydS0axEK1m4qbqq2q6rtcpHnRe47vxhutyi2qYIdHKVh2GE4D3x6pNxw6nF5iqp414smJ4cbtmxQ1PKmzaog5DkxE5v8iqb2cV22YBcSTJxuwW2DgEKgX7sfi0qNxw9_Eqdwr29AAFAlkMZWeNAjpZEQ2haIv5xXzYggHeituYFG7CI7xK_1fLuG795Rb0VVBBIKvf0
//fd state: (17) 46CVo00yrl3np4GS1WZRdSKo60SsmTtroC3sp1NNbU4yOGkT34hny8LEBXu8WpXTwzCfxhyd6SCkSSfkD8qNAWPPMkhGtDH4_uVAfYRAcT5IdoKg2hmEzfr1AlQJmX

ret = client.Write(fd705, []byte("y3FKvOz4ec4fQgUZoX2EcExmxMlpdXM8d2BVTqHHP62SJQ2sr9RKEAtWkp61abYAf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) 46CVo00yrl3np4GS1y3FKvOz4ec4fQgUZoX2EcExmxMlpdXM8d2BVTqHHP62SJQ2sr9RKEAtWkp61abYAfqNAWPPMkhGtDH4_uVAfYRAcT5IdoKg2hmEzfr1AlQJmX

fd758 := client.Open("/NDaLK8HBXg.txt", client.O_RDWR|client.O_CREATE)
if(fd758 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd699, []byte("9IK2ygmQo2hsRIITowsvcPHxxJ9JCFdf09Mm8V0UJDgIfilTF5uU0zoiNvebu3wtAjLSkPhFf9pJiL_hloOmo4seJWAqcxNd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) 9IK2ygmQo2hsRIITowsvcPHxxJ9JCFdf09Mm8V0UJDgIfilTF5uU0zoiNvebu3wtAjLSkPhFf9pJiL_hloOmo4seJWAqcxNd

ret = client.Close(fd727)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd733, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


ret = client.Close(fd687)
if(ret != 0) {
  panic("close failed")
}

//fd state: (146) GpAaUHLWOZF3XbxkSZDdqYPsCB1i3yEfE5gF2L4pWscI8BRN7M0fiZVEii_3IQmYEkXG3nMC6UTqQy61GU_WYdQ95I4_jFUw3_cjHiHtnmzfSJX9upokRk1H3yN6U8gERMk3JHhbGC54lxUrXm

ret = client.Write(fd676, []byte("4oc7d1iIN6B6Fzu6HCDj6lzc6IrwgqJSU1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (180) GpAaUHLWOZF3XbxkSZDdqYPsCB1i3yEfE5gF2L4pWscI8BRN7M0fiZVEii_3IQmYEkXG3nMC6UTqQy61GU_WYdQ95I4_jFUw3_cjHiHtnmzfSJX9upokRk1H3yN6U8gERMk3JHhbGC54lxUrXm4oc7d1iIN6B6Fzu6HCDj6lzc6IrwgqJSU1

buf, ret = client.Read(fd751, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd759 := client.Open("/n16dqTGUku.txt", client.O_RDWR|client.O_CREATE)
if(fd759 < 0) {
  panic("open failed")
}


fd760 := client.Open("/8jCWuWK8ZO.txt", client.O_RDWR|client.O_CREATE)
if(fd760 < 0) {
  panic("open failed")
}


ret = client.Seek(fd706, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


buf, ret = client.Read(fd746, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Rf81FF_mf8CAz78GK6tiTpxpTHRBmSMMU7") {
  panic("wrong data returned")
}


ret = client.Close(fd759)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd704, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


ret = client.Seek(fd724, 157, client.SEEK_SET)
if(ret != 157) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 157)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd754, []byte("cZFDTi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) cZFDTi

buf, ret = client.Read(fd724, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qNxw9_Eqdwr29AAFAlkMZWeNAjpZEQ2haIv5xXzYggHeituYFG7CI7xK") {
  panic("wrong data returned")
}


fd761 := client.Open("/7evLhsHpf6.txt", client.O_RDWR|client.O_CREATE)
if(fd761 < 0) {
  panic("open failed")
}


fd762 := client.Open("/7JfTK64J_0.txt", client.O_RDWR|client.O_CREATE)
if(fd762 < 0) {
  panic("open failed")
}


ret = client.Seek(fd736, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd704, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Seek(fd757, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd715, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}

//fd state: (0) _Cb4_9cMgwbLx4w1hvzD6F60xU3mrcdX3_f8lTUDyDRHjb5gGcvaItubk0uhZ74HqntI5XDpCHifTJxgFgKRyTQ1eB8bZq1Pc2t3uvsjhGqKbx0V_CnQx27xDTy98s4qzewjk5p50ZZLyj0vU8NSGc

ret = client.Write(fd758, []byte("2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) 2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNubk0uhZ74HqntI5XDpCHifTJxgFgKRyTQ1eB8bZq1Pc2t3uvsjhGqKbx0V_CnQx27xDTy98s4qzewjk5p50ZZLyj0vU8NSGc

buf, ret = client.Read(fd724, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_1fLuG795Rb0VVBBIKvf0") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd747, []byte("695MQtUbSh7qqjw8q_SGn3aZ_aHmwxUfvbQHlgMppk0UI4ec_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) 695MQtUbSh7qqjw8q_SGn3aZ_aHmwxUfvbQHlgMppk0UI4ec_

ret = client.Close(fd757)
if(ret != 0) {
  panic("close failed")
}


fd763 := client.Open("/Wc8I1sgyqN.txt", client.O_RDWR|client.O_CREATE)
if(fd763 < 0) {
  panic("open failed")
}


ret = client.Close(fd707)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd719, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


ret = client.Seek(fd736, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Seek(fd754, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


fd764 := client.Open("/_skInyTGnN.txt", client.O_RDWR|client.O_CREATE)
if(fd764 < 0) {
  panic("open failed")
}


ret = client.Seek(fd704, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd760)
if(ret != 0) {
  panic("close failed")
}


fd765 := client.Open("/oyMaCMMJ9_.txt", client.O_RDWR|client.O_CREATE)
if(fd765 < 0) {
  panic("open failed")
}


ret = client.Seek(fd676, 137, client.SEEK_SET)
if(ret != 137) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 137)
  panic("seek failed")
}


ret = client.Close(fd748)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd706)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd607, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bEc") {
  panic("wrong data returned")
}


fd766 := client.Open("/J2zT7GDROy.txt", client.O_RDWR|client.O_CREATE)
if(fd766 < 0) {
  panic("open failed")
}


fd767 := client.Open("/Vd2TJjK6jh.txt", client.O_RDWR|client.O_CREATE)
if(fd767 < 0) {
  panic("open failed")
}


ret = client.Seek(fd762, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (3) P7SUorwK

ret = client.Write(fd736, []byte("obYghq4JMzrFWBsfMl45wygh0oGU4PsiGV9KjxQHapacNH1DpEsQOqfbsvoEuSOS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) P7SobYghq4JMzrFWBsfMl45wygh0oGU4PsiGV9KjxQHapacNH1DpEsQOqfbsvoEuSOS

ret = client.Seek(fd741, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


fd768 := client.Open("/MVWR6bEOKx.txt", client.O_RDWR|client.O_CREATE)
if(fd768 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd749, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd724, 213, client.SEEK_SET)
if(ret != 213) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 213)
  panic("seek failed")
}


buf, ret = client.Read(fd725, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uvIBoU") {
  panic("wrong data returned")
}


ret = client.Seek(fd762, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd749, []byte("L9P65aPA62EyWoyQspNcrpMUc4C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) L9P65aPA62EyWoyQspNcrpMUc4C

buf, ret = client.Read(fd738, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd648, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (84) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNbIn4OilxK7ZMIYSf0rXYJSN1HdwDzk8E6tIIwxKb1ITQX2qQUX150Y61DS_sCXIom

ret = client.Write(fd721, []byte("IK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKn4OilxK7ZMIYSf0rXYJSN1HdwDzk8E6tIIwxKb1ITQX2qQUX150Y61DS_sCXIom

ret = client.Seek(fd704, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (86) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKn4OilxK7ZMIYSf0rXYJSN1HdwDzk8E6tIIwxKb1ITQX2qQUX150Y61DS_sCXIom

ret = client.Write(fd721, []byte("iMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (159) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvk

ret = client.Seek(fd739, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd607, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd769 := client.Open("/8sNmQEjXdP.txt", client.O_RDWR|client.O_CREATE)
if(fd769 < 0) {
  panic("open failed")
}


ret = client.Close(fd767)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd719, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "tkmPUnf8SFl1a2Bm") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd714, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "V7ymH_tooFtdG0WhKe3") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd765, []byte("2r40gC23li0wDxn0z7bklbAmAUpG_onCZhZH5ifvo43yMiJVejDhPXeNF3nyJeBdP3hYagSxNbr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) 2r40gC23li0wDxn0z7bklbAmAUpG_onCZhZH5ifvo43yMiJVejDhPXeNF3nyJeBdP3hYagSxNbr

fd770 := client.Open("/tInHpwPscv.txt", client.O_RDWR|client.O_CREATE)
if(fd770 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd725, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd746)
if(ret != 0) {
  panic("close failed")
}

//fd state: (22) 52ryyxV6A5qPyCglvr_Vn1kJiwgNaDUNRERhc2ufFxIVIQXRDiDq76CuR8P3_xLzqAy1xj6_A7ZiSHmCxrrzoL_yv8n8h7ftRLFuaeqfq3kOEyoZOz035n1d

ret = client.Write(fd672, []byte("5waT_qXZh2ACBTwM7FejojB8nJ1wTQ66fbpLndjDT6JNYslxpGcb0EN2_og2HAzThwob0B39FDW6LpA6qlD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) 52ryyxV6A5qPyCglvr_Vn15waT_qXZh2ACBTwM7FejojB8nJ1wTQ66fbpLndjDT6JNYslxpGcb0EN2_og2HAzThwob0B39FDW6LpA6qlD3kOEyoZOz035n1d

buf, ret = client.Read(fd699, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (246) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HS5RthNn0GVKF9emD4dXWbZ1UtlcY3X4mqHeLUetC7lBb2mfne6DfIG7CTPhRvj7vKcvVRMiuwzENHLYABqSkC7QFYiEePz9b_E3_iIua7E3rmSIbWnmbmwFbEc

ret = client.Write(fd607, []byte("p01veSKqnIi1X5nAUfsyzyZ86B2t7r9Fv2i6EZdmlSgJUeNxej5fwcvaX1UnMNpqhOAOSihimz55A7o8eGG1BE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (332) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HS5RthNn0GVKF9emD4dXWbZ1UtlcY3X4mqHeLUetC7lBb2mfne6DfIG7CTPhRvj7vKcvVRMiuwzENHLYABqSkC7QFYiEePz9b_E3_iIua7E3rmSIbWnmbmwFbEcp01veSKqnIi1X5nAUfsyzyZ86B2t7r9Fv2i6EZdmlSgJUeNxej5fwcvaX1UnMNpqhOAOSihimz55A7o8eGG1BE
//fd state: (69) A9awLY38LiPZxxErjJwTOKWFl1W6C9qajRuMfNEyYeQNmO5Rt8RA6r71nqSNKH6csapyCILOGOjSQh1PlRW9LQu4hkSDszsWX5F0U

ret = client.Write(fd737, []byte("jwmsMSsKKbkeCKKeSMte2v7A_lCFRrtKhclFW0RJJqKrBXsoNibbsAiyH8K6GHXMu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) A9awLY38LiPZxxErjJwTOKWFl1W6C9qajRuMfNEyYeQNmO5Rt8RA6r71nqSNKH6csapyCjwmsMSsKKbkeCKKeSMte2v7A_lCFRrtKhclFW0RJJqKrBXsoNibbsAiyH8K6GHXMu

buf, ret = client.Read(fd764, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd715)
if(ret != 0) {
  panic("close failed")
}


fd771 := client.Open("/W14ZY4nXnm.txt", client.O_RDWR|client.O_CREATE)
if(fd771 < 0) {
  panic("open failed")
}


fd772 := client.Open("/ZMyenLW9Zc.txt", client.O_RDWR|client.O_CREATE)
if(fd772 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd723, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd773 := client.Open("/DsgsMQNlzS.txt", client.O_RDWR|client.O_CREATE)
if(fd773 < 0) {
  panic("open failed")
}


ret = client.Seek(fd735, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}


ret = client.Close(fd725)
if(ret != 0) {
  panic("close failed")
}


fd774 := client.Open("/L8Y7EMLJqt.txt", client.O_RDWR|client.O_CREATE)
if(fd774 < 0) {
  panic("open failed")
}


ret = client.Close(fd683)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd756, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Close(fd740)
if(ret != 0) {
  panic("close failed")
}


fd775 := client.Open("/6zCWkaOFR7.txt", client.O_RDWR|client.O_CREATE)
if(fd775 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd768, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd772, []byte("LNd2Euv9jfqXLxCG91mi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) LNd2Euv9jfqXLxCG91mi

ret = client.Close(fd762)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd766, []byte("aj_HEn3c2Q3jRGJRSPJWte3QIOus3u0GzRU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) aj_HEn3c2Q3jRGJRSPJWte3QIOus3u0GzRU

fd776 := client.Open("/odYHg5Vmi4.txt", client.O_RDWR|client.O_CREATE)
if(fd776 < 0) {
  panic("open failed")
}


ret = client.Close(fd754)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd735, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "w") {
  panic("wrong data returned")
}


ret = client.Close(fd776)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd773, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd676)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd766, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Close(fd719)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd772, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd777 := client.Open("/f_VMh3LK24.txt", client.O_RDWR|client.O_CREATE)
if(fd777 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd739, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd777, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


buf, ret = client.Read(fd765, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd758, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


fd778 := client.Open("/0UaTwBg6W9.txt", client.O_RDWR|client.O_CREATE)
if(fd778 < 0) {
  panic("open failed")
}


ret = client.Close(fd777)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd772, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


buf, ret = client.Read(fd770, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd709)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd699, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}

//fd state: (17) aj_HEn3c2Q3jRGJRSPJWte3QIOus3u0GzRU

ret = client.Write(fd766, []byte("iZSw0cy97_xZgfQgmlTUNGnKvv0DXnSRTyPImwGwglPydD8r0L6Zu5uXBCAvmP8MBTXnK9Pc1Kgx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) aj_HEn3c2Q3jRGJRSiZSw0cy97_xZgfQgmlTUNGnKvv0DXnSRTyPImwGwglPydD8r0L6Zu5uXBCAvmP8MBTXnK9Pc1Kgx

ret = client.Seek(fd724, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


fd779 := client.Open("/rcJYn3q9Iv.txt", client.O_RDWR|client.O_CREATE)
if(fd779 < 0) {
  panic("open failed")
}


ret = client.Seek(fd714, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


fd780 := client.Open("/aUFQvgjfrD.txt", client.O_RDWR|client.O_CREATE)
if(fd780 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd737, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd774, []byte("apT3Tw7dBCc61C5Lbg05ba0o0wK5H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) apT3Tw7dBCc61C5Lbg05ba0o0wK5H

ret = client.Close(fd648)
if(ret != 0) {
  panic("close failed")
}

//fd state: (49) 695MQtUbSh7qqjw8q_SGn3aZ_aHmwxUfvbQHlgMppk0UI4ec_

ret = client.Write(fd747, []byte("57A2laH3LuGl67Sgpe29vcGPLG4U7AXqutD1ifgHasYQ1I72"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) 695MQtUbSh7qqjw8q_SGn3aZ_aHmwxUfvbQHlgMppk0UI4ec_57A2laH3LuGl67Sgpe29vcGPLG4U7AXqutD1ifgHasYQ1I72
//fd state: (0) _QxLDIWcnNtoyT8RscuOsZFIrXR3s57VgSaKnhrfo3PL4a54qzaI8FUxPJyQbUqde2bdnO42vO3ArXRyZlWSMb3DUKQq_qwd4bzTLnBwq5zVIUChZE4qgTVOCbh0rjJTPmjO8LhvnVtfQGreDAI1XPcrQu_7IWcxj339rRH0Why3tWFS9iCvROX7iRBJhw1CimzlvOxIGoTV20Yy8n

ret = client.Write(fd779, []byte("o4nwK9buDZZWN6YXJ1SK8vSfajqrBiQts6eg1WXUj9s2qF6gxwh8_miQZWFNhO7eJGaovXf9U7IKaF01BUQw7i0WL0LFTu5et"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) o4nwK9buDZZWN6YXJ1SK8vSfajqrBiQts6eg1WXUj9s2qF6gxwh8_miQZWFNhO7eJGaovXf9U7IKaF01BUQw7i0WL0LFTu5etbzTLnBwq5zVIUChZE4qgTVOCbh0rjJTPmjO8LhvnVtfQGreDAI1XPcrQu_7IWcxj339rRH0Why3tWFS9iCvROX7iRBJhw1CimzlvOxIGoTV20Yy8n
//fd state: (67) P7SobYghq4JMzrFWBsfMl45wygh0oGU4PsiGV9KjxQHapacNH1DpEsQOqfbsvoEuSOS

ret = client.Write(fd736, []byte("fO0_64ckayz5sTFow8ljuZqHMMnEnwoccgVFH6O1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) P7SobYghq4JMzrFWBsfMl45wygh0oGU4PsiGV9KjxQHapacNH1DpEsQOqfbsvoEuSOSfO0_64ckayz5sTFow8ljuZqHMMnEnwoccgVFH6O1

ret = client.Close(fd769)
if(ret != 0) {
  panic("close failed")
}


fd781 := client.Open("/rDZVcO3UQt.txt", client.O_RDWR|client.O_CREATE)
if(fd781 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd774, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd750, []byte("LBJmwbSgSfg2jCol1_N5PaEQJhmQCNW9bnCzJKJ_5eDrzCgiMW7TJcUzvnqPEyb551L8bgk4F9wSe3Hzz3LhnaojqWAUL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) LBJmwbSgSfg2jCol1_N5PaEQJhmQCNW9bnCzJKJ_5eDrzCgiMW7TJcUzvnqPEyb551L8bgk4F9wSe3Hzz3LhnaojqWAUL
//fd state: (332) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HS5RthNn0GVKF9emD4dXWbZ1UtlcY3X4mqHeLUetC7lBb2mfne6DfIG7CTPhRvj7vKcvVRMiuwzENHLYABqSkC7QFYiEePz9b_E3_iIua7E3rmSIbWnmbmwFbEcp01veSKqnIi1X5nAUfsyzyZ86B2t7r9Fv2i6EZdmlSgJUeNxej5fwcvaX1UnMNpqhOAOSihimz55A7o8eGG1BE

ret = client.Write(fd607, []byte("lN3gIpZ8iF_N_OiB8CBgWl6LNjAOupO5g3sMjY0J4gzUfniSXOhGBy2WmelNL8mZmVxz5QRh82m4sm5HzSw41dZisGfmquEIl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (429) whhM5w9WPz19eSlWqFv1DUHLYTCO_2h2KouRiN1cO3FqAUfRrVJuov2tauEZn3xdkHhwaH8KeKWstxKRxcwNcQKtUJqXcLPf94l7QukIQepR1DJwdn0zt8WEaC0HS5RthNn0GVKF9emD4dXWbZ1UtlcY3X4mqHeLUetC7lBb2mfne6DfIG7CTPhRvj7vKcvVRMiuwzENHLYABqSkC7QFYiEePz9b_E3_iIua7E3rmSIbWnmbmwFbEcp01veSKqnIi1X5nAUfsyzyZ86B2t7r9Fv2i6EZdmlSgJUeNxej5fwcvaX1UnMNpqhOAOSihimz55A7o8eGG1BElN3gIpZ8iF_N_OiB8CBgWl6LNjAOupO5g3sMjY0J4gzUfniSXOhGBy2WmelNL8mZmVxz5QRh82m4sm5HzSw41dZisGfmquEIl

buf, ret = client.Read(fd764, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd782 := client.Open("/r_xlwf3nRN.txt", client.O_RDWR|client.O_CREATE)
if(fd782 < 0) {
  panic("open failed")
}


fd783 := client.Open("/PnktofjcNb.txt", client.O_RDWR|client.O_CREATE)
if(fd783 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd768, []byte("p"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) p

fd784 := client.Open("/LSKqLmfg54.txt", client.O_RDWR|client.O_CREATE)
if(fd784 < 0) {
  panic("open failed")
}


ret = client.Close(fd712)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd741, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}

//fd state: (97) 695MQtUbSh7qqjw8q_SGn3aZ_aHmwxUfvbQHlgMppk0UI4ec_57A2laH3LuGl67Sgpe29vcGPLG4U7AXqutD1ifgHasYQ1I72

ret = client.Write(fd747, []byte("uy8D8P7OvSCARlMATKMteNGiedsan80U9dqFjkRDjh6onwFwC4Tl7FW78"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) 695MQtUbSh7qqjw8q_SGn3aZ_aHmwxUfvbQHlgMppk0UI4ec_57A2laH3LuGl67Sgpe29vcGPLG4U7AXqutD1ifgHasYQ1I72uy8D8P7OvSCARlMATKMteNGiedsan80U9dqFjkRDjh6onwFwC4Tl7FW78

buf, ret = client.Read(fd775, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "kLk") {
  panic("wrong data returned")
}


ret = client.Seek(fd761, 164, client.SEEK_SET)
if(ret != 164) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 164)
  panic("seek failed")
}


ret = client.Seek(fd669, 118, client.SEEK_SET)
if(ret != 118) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 118)
  panic("seek failed")
}


ret = client.Seek(fd765, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


fd785 := client.Open("/TsBuxt90HY.txt", client.O_RDWR|client.O_CREATE)
if(fd785 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd737, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (134) A9awLY38LiPZxxErjJwTOKWFl1W6C9qajRuMfNEyYeQNmO5Rt8RA6r71nqSNKH6csapyCjwmsMSsKKbkeCKKeSMte2v7A_lCFRrtKhclFW0RJJqKrBXsoNibbsAiyH8K6GHXMu

ret = client.Write(fd737, []byte("6pNooXQcv_FQV1AnCLYd8tedUSX9VX8C6CvjLd0wpM8KigTRIkXOoVZwNj9vL6V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (197) A9awLY38LiPZxxErjJwTOKWFl1W6C9qajRuMfNEyYeQNmO5Rt8RA6r71nqSNKH6csapyCjwmsMSsKKbkeCKKeSMte2v7A_lCFRrtKhclFW0RJJqKrBXsoNibbsAiyH8K6GHXMu6pNooXQcv_FQV1AnCLYd8tedUSX9VX8C6CvjLd0wpM8KigTRIkXOoVZwNj9vL6V

fd786 := client.Open("/_q7QZEAIS7.txt", client.O_RDWR|client.O_CREATE)
if(fd786 < 0) {
  panic("open failed")
}


fd787 := client.Open("/uCE42t4glI.txt", client.O_RDWR|client.O_CREATE)
if(fd787 < 0) {
  panic("open failed")
}


fd788 := client.Open("/qIa7ZNkFqa.txt", client.O_RDWR|client.O_CREATE)
if(fd788 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd785, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd732)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd788)
if(ret != 0) {
  panic("close failed")
}


fd789 := client.Open("/9mLeHk15PZ.txt", client.O_RDWR|client.O_CREATE)
if(fd789 < 0) {
  panic("open failed")
}


ret = client.Close(fd751)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd750, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd790 := client.Open("/C7dayJrBXf.txt", client.O_RDWR|client.O_CREATE)
if(fd790 < 0) {
  panic("open failed")
}


ret = client.Close(fd774)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd785, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd791 := client.Open("/g221aMxIGc.txt", client.O_RDWR|client.O_CREATE)
if(fd791 < 0) {
  panic("open failed")
}


ret = client.Close(fd756)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd607, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Close(fd672)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd773, []byte("PIOSu4JrAr5yg40VxetXVMujkyu8ewANJfqBDXaVMkSSnQBNYJnOdu8N3wS_y8zXMQz2JTakAAQA0e08"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) PIOSu4JrAr5yg40VxetXVMujkyu8ewANJfqBDXaVMkSSnQBNYJnOdu8N3wS_y8zXMQz2JTakAAQA0e08

ret = client.Close(fd785)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd763, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Na3vDus5puioiFGF9DuLtFr45g0gY6NWly7DZb") {
  panic("wrong data returned")
}


ret = client.Seek(fd733, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


ret = client.Seek(fd699, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


ret = client.Seek(fd723, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


buf, ret = client.Read(fd791, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd704)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd783)
if(ret != 0) {
  panic("close failed")
}


fd792 := client.Open("/lhlbAGE_2Z.txt", client.O_RDWR|client.O_CREATE)
if(fd792 < 0) {
  panic("open failed")
}


fd793 := client.Open("/7kjLfusQPx.txt", client.O_RDWR|client.O_CREATE)
if(fd793 < 0) {
  panic("open failed")
}


fd794 := client.Open("/ni095iaCWC.txt", client.O_RDWR|client.O_CREATE)
if(fd794 < 0) {
  panic("open failed")
}

//fd state: (3) kLkORRlHF6vzmbUosRu7kPpM7dLVG7pLofh1UU8zPkreYH_kGnziQhQTdukZpuct

ret = client.Write(fd775, []byte("r49aYsH1yF3hsFScHbATG64GX0CgKHRU_ALIyg0jV7N7Aoe3poxWWDji"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) kLkr49aYsH1yF3hsFScHbATG64GX0CgKHRU_ALIyg0jV7N7Aoe3poxWWDjiZpuct
//fd state: (93) aj_HEn3c2Q3jRGJRSiZSw0cy97_xZgfQgmlTUNGnKvv0DXnSRTyPImwGwglPydD8r0L6Zu5uXBCAvmP8MBTXnK9Pc1Kgx

ret = client.Write(fd766, []byte("CAE_EUI1jy0lOgBv4qt6XGkB_yyNU1Yq7axS5s7R9ms6yW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (139) aj_HEn3c2Q3jRGJRSiZSw0cy97_xZgfQgmlTUNGnKvv0DXnSRTyPImwGwglPydD8r0L6Zu5uXBCAvmP8MBTXnK9Pc1KgxCAE_EUI1jy0lOgBv4qt6XGkB_yyNU1Yq7axS5s7R9ms6yW

ret = client.Seek(fd750, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


buf, ret = client.Read(fd763, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "34_YHc3d36E_irriZS7w8w_Tg3sk8h7sdsmh80MKIfMFRwQ") {
  panic("wrong data returned")
}


ret = client.Seek(fd786, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd780, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd741, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


ret = client.Seek(fd714, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


buf, ret = client.Read(fd787, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd795 := client.Open("/BFzwaNL70e.txt", client.O_RDWR|client.O_CREATE)
if(fd795 < 0) {
  panic("open failed")
}


ret = client.Close(fd779)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd737, 169, client.SEEK_SET)
if(ret != 169) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 169)
  panic("seek failed")
}


buf, ret = client.Read(fd723, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "HYWKdLsmHh06") {
  panic("wrong data returned")
}


ret = client.Close(fd763)
if(ret != 0) {
  panic("close failed")
}


fd796 := client.Open("/g_5w2E_4_7.txt", client.O_RDWR|client.O_CREATE)
if(fd796 < 0) {
  panic("open failed")
}


ret = client.Seek(fd736, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


ret = client.Seek(fd714, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


buf, ret = client.Read(fd747, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd782, []byte("8X0FUJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) 8X0FUJ

fd797 := client.Open("/2CiEz3nloe.txt", client.O_RDWR|client.O_CREATE)
if(fd797 < 0) {
  panic("open failed")
}


ret = client.Close(fd794)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd793)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd737, 195, client.SEEK_SET)
if(ret != 195) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 195)
  panic("seek failed")
}


fd798 := client.Open("/S9X5bVz0zt.txt", client.O_RDWR|client.O_CREATE)
if(fd798 < 0) {
  panic("open failed")
}


ret = client.Close(fd782)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd734)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd796, 33)
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


buf, ret = client.Read(fd786, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd764)
if(ret != 0) {
  panic("close failed")
}


fd799 := client.Open("/nBwnAtneem.txt", client.O_RDWR|client.O_CREATE)
if(fd799 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd791, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd784, []byte("Fh6bsoJq9MhzlFHoOiTOF15dtc92ejx7Yoks87JlRSKh7sK_V4nckkJQNMu7Jo3N"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) Fh6bsoJq9MhzlFHoOiTOF15dtc92ejx7Yoks87JlRSKh7sK_V4nckkJQNMu7Jo3N

ret = client.Seek(fd752, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd699)
if(ret != 0) {
  panic("close failed")
}


fd800 := client.Open("/J2zT7GDROy.txt", client.O_RDWR|client.O_CREATE)
if(fd800 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd786, []byte("ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30Gtxh8zn6J7WdiSMBQFj8SuA3HOBM6zepiYSauW1c8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30Gtxh8zn6J7WdiSMBQFj8SuA3HOBM6zepiYSauW1c8

ret = client.Seek(fd768, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd792, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (27) L9P65aPA62EyWoyQspNcrpMUc4C

ret = client.Write(fd749, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) L9P65aPA62EyWoyQspNcrpMUc4C

ret = client.Close(fd607)
if(ret != 0) {
  panic("close failed")
}

//fd state: (34) 97V_WpedoDgdE9rQlulGDV5bWWwVUkNmZvoYmYBTJAxwaCSYXO6niMJxQXutVqi

ret = client.Write(fd741, []byte("lj6G5QPPsb0Z4WpQBZWW8nyD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) 97V_WpedoDgdE9rQlulGDV5bWWwVUkNmZvlj6G5QPPsb0Z4WpQBZWW8nyDutVqi

buf, ret = client.Read(fd705, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qNAWPPMkhGtDH4_uVAfYRAcT5IdoKg2hmEzfr1AlQJmX") {
  panic("wrong data returned")
}


ret = client.Close(fd791)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd765, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "onCZhZH5ifvo43yMiJVejDhPXeNF3nyJeBdP3hYag") {
  panic("wrong data returned")
}


ret = client.Close(fd737)
if(ret != 0) {
  panic("close failed")
}


fd801 := client.Open("/Z68lMy3wwA.txt", client.O_RDWR|client.O_CREATE)
if(fd801 < 0) {
  panic("open failed")
}


ret = client.Seek(fd801, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd799)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd761, 207, client.SEEK_SET)
if(ret != 207) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 207)
  panic("seek failed")
}

//fd state: (70) 2r40gC23li0wDxn0z7bklbAmAUpG_onCZhZH5ifvo43yMiJVejDhPXeNF3nyJeBdP3hYagSxNbr

ret = client.Write(fd765, []byte("TkRG2mF2o10SKjoOXdwdmQyR7B9OjPF8WJ7423vq_UGNnXSN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) 2r40gC23li0wDxn0z7bklbAmAUpG_onCZhZH5ifvo43yMiJVejDhPXeNF3nyJeBdP3hYagTkRG2mF2o10SKjoOXdwdmQyR7B9OjPF8WJ7423vq_UGNnXSN
//fd state: (0) 

ret = client.Write(fd796, []byte("cM24y1j_uq88pmzWKopeif5aYtd8Ex1oQRpegSjsFB8tVNxJKLnV7gMzdkmZAdxffDfxnDQZjLqkaYKIDsk1d1HWtOZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) cM24y1j_uq88pmzWKopeif5aYtd8Ex1oQRpegSjsFB8tVNxJKLnV7gMzdkmZAdxffDfxnDQZjLqkaYKIDsk1d1HWtOZ
//fd state: (18) LNd2Euv9jfqXLxCG91mi

ret = client.Write(fd772, []byte("IVDMwHepJYm1wpqgt5B"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) LNd2Euv9jfqXLxCG91IVDMwHepJYm1wpqgt5B

ret = client.Seek(fd752, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd801, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd802 := client.Open("/A1A8lCyTOn.txt", client.O_RDWR|client.O_CREATE)
if(fd802 < 0) {
  panic("open failed")
}


fd803 := client.Open("/LtPMUri2sN.txt", client.O_RDWR|client.O_CREATE)
if(fd803 < 0) {
  panic("open failed")
}


ret = client.Seek(fd723, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}

//fd state: (91) cM24y1j_uq88pmzWKopeif5aYtd8Ex1oQRpegSjsFB8tVNxJKLnV7gMzdkmZAdxffDfxnDQZjLqkaYKIDsk1d1HWtOZ

ret = client.Write(fd796, []byte("4u8zpEV8kdW3mtk_Jhl9Nys2dXWCh3XxWcRuNMeZiUQe_f8Q7FIj0i6EUKpg2KiR_wtmFwWGMnBlUVJKgLLt96v8dvui"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (183) cM24y1j_uq88pmzWKopeif5aYtd8Ex1oQRpegSjsFB8tVNxJKLnV7gMzdkmZAdxffDfxnDQZjLqkaYKIDsk1d1HWtOZ4u8zpEV8kdW3mtk_Jhl9Nys2dXWCh3XxWcRuNMeZiUQe_f8Q7FIj0i6EUKpg2KiR_wtmFwWGMnBlUVJKgLLt96v8dvui

ret = client.Seek(fd787, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd736, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


fd804 := client.Open("/Barlvoj80F.txt", client.O_RDWR|client.O_CREATE)
if(fd804 < 0) {
  panic("open failed")
}

//fd state: (118) RkOPoRw9bfDFY1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3UodbYEm0itMYmVMNBrTmIoIBYTnJK7k8_7J14u9PTShKsXFBjRgpXMCR7H4wZFSMdnWfFUojO9MP3sLp3zAx8LxqsjnjiARhyl6aUTX84AsOUN74HtbKKoaLMWejO7IYzRWRRQMzmTjvQh7OkzkrGS

ret = client.Write(fd669, []byte("rb_d1f678Knj3eSMNvRQdwxCHo99lxvV3TTTLGElLq2Ayy7qB_BZT_iW0z6GwTgZPlHxmqRAtIDQFaRMYQitHI5K8l3XXGEe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (214) RkOPoRw9bfDFY1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3UodbYEm0itMYmVMNBrTmIoIBYTnJK7k8_7J14u9PTShKsXFBjRgpXMCR7H4wZFSMdnWrb_d1f678Knj3eSMNvRQdwxCHo99lxvV3TTTLGElLq2Ayy7qB_BZT_iW0z6GwTgZPlHxmqRAtIDQFaRMYQitHI5K8l3XXGEe

fd805 := client.Open("/8tf8K5Cetz.txt", client.O_RDWR|client.O_CREATE)
if(fd805 < 0) {
  panic("open failed")
}

//fd state: (214) RkOPoRw9bfDFY1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3UodbYEm0itMYmVMNBrTmIoIBYTnJK7k8_7J14u9PTShKsXFBjRgpXMCR7H4wZFSMdnWrb_d1f678Knj3eSMNvRQdwxCHo99lxvV3TTTLGElLq2Ayy7qB_BZT_iW0z6GwTgZPlHxmqRAtIDQFaRMYQitHI5K8l3XXGEe

ret = client.Write(fd669, []byte("JUDKHyviz3ZRklXoVrlj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (234) RkOPoRw9bfDFY1VEaBCEHGeLyEwYXODULp4cn0irSopUNp7t7s3UodbYEm0itMYmVMNBrTmIoIBYTnJK7k8_7J14u9PTShKsXFBjRgpXMCR7H4wZFSMdnWrb_d1f678Knj3eSMNvRQdwxCHo99lxvV3TTTLGElLq2Ayy7qB_BZT_iW0z6GwTgZPlHxmqRAtIDQFaRMYQitHI5K8l3XXGEeJUDKHyviz3ZRklXoVrlj

ret = client.Seek(fd797, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd749)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd789)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd786, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd705, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd806 := client.Open("/BiqkOPZuTM.txt", client.O_RDWR|client.O_CREATE)
if(fd806 < 0) {
  panic("open failed")
}


ret = client.Close(fd773)
if(ret != 0) {
  panic("close failed")
}

//fd state: (64) Fh6bsoJq9MhzlFHoOiTOF15dtc92ejx7Yoks87JlRSKh7sK_V4nckkJQNMu7Jo3N

ret = client.Write(fd784, []byte("gDHgKHXOEht7fRHyILt02kfe5Mb8TLVc0mMZZh4Oy1oqL4DC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (112) Fh6bsoJq9MhzlFHoOiTOF15dtc92ejx7Yoks87JlRSKh7sK_V4nckkJQNMu7Jo3NgDHgKHXOEht7fRHyILt02kfe5Mb8TLVc0mMZZh4Oy1oqL4DC

buf, ret = client.Read(fd770, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd806)
if(ret != 0) {
  panic("close failed")
}


fd807 := client.Open("/H0fKHVXLqI.txt", client.O_RDWR|client.O_CREATE)
if(fd807 < 0) {
  panic("open failed")
}


fd808 := client.Open("/eTqaZXcphg.txt", client.O_RDWR|client.O_CREATE)
if(fd808 < 0) {
  panic("open failed")
}


ret = client.Seek(fd792, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd792)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd781, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd733, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd797, []byte("aUDTsvTIMxAZX8Mr8rOZC78P04rRjwpCKSwaQLUbIgUANpUps0ZjE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) aUDTsvTIMxAZX8Mr8rOZC78P04rRjwpCKSwaQLUbIgUANpUps0ZjE

fd809 := client.Open("/x5DumFEOhh.txt", client.O_RDWR|client.O_CREATE)
if(fd809 < 0) {
  panic("open failed")
}


ret = client.Close(fd784)
if(ret != 0) {
  panic("close failed")
}


fd810 := client.Open("/MVWR6bEOKx.txt", client.O_RDWR|client.O_CREATE)
if(fd810 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd750, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "N5PaEQJhmQCNW9bnCzJKJ") {
  panic("wrong data returned")
}

//fd state: (59) kLkr49aYsH1yF3hsFScHbATG64GX0CgKHRU_ALIyg0jV7N7Aoe3poxWWDjiZpuct

ret = client.Write(fd775, []byte("k6H8bDFLTafPHEy2aflBByaIqHvczGaIjLdGlOnwF_7T8YNLcjmZDp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) kLkr49aYsH1yF3hsFScHbATG64GX0CgKHRU_ALIyg0jV7N7Aoe3poxWWDjik6H8bDFLTafPHEy2aflBByaIqHvczGaIjLdGlOnwF_7T8YNLcjmZDp
//fd state: (0) VMtiv0CSu0X7mNyDxx9yKZLgyirhnHMu4oK85hOcobv7yE2sS_jXYU9889dnySpN_zIh7X2uJ0jDwxoQ7h33nbM

ret = client.Write(fd795, []byte("sd6FNxfgwCvCVXdP8UxPkqVHOhG1dEQHGAKh3p0zQz7yWmmfkkzJhm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) sd6FNxfgwCvCVXdP8UxPkqVHOhG1dEQHGAKh3p0zQz7yWmmfkkzJhm9889dnySpN_zIh7X2uJ0jDwxoQ7h33nbM

fd811 := client.Open("/Xjz4ccT9Te.txt", client.O_RDWR|client.O_CREATE)
if(fd811 < 0) {
  panic("open failed")
}


ret = client.Close(fd790)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd771, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


fd812 := client.Open("/rDZVcO3UQt.txt", client.O_RDWR|client.O_CREATE)
if(fd812 < 0) {
  panic("open failed")
}


ret = client.Seek(fd750, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd805, []byte("NM9TDGn7wHWYevj4uJLKcW08no7uUlWIy0curRvG3Zkd4_3IepRzmcaVEtFvpJIeptDCCLhWeqr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) NM9TDGn7wHWYevj4uJLKcW08no7uUlWIy0curRvG3Zkd4_3IepRzmcaVEtFvpJIeptDCCLhWeqr
//fd state: (0) 

ret = client.Write(fd738, []byte("976FGiAP0xyX3Rm72A2vPIM0B2gUoM0GmmWFlseAtdKrp0BX5XqWVSQR4DcXYXLFYyODNVY3sBJ6T1uPRh3LzLPHQIsrfT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) 976FGiAP0xyX3Rm72A2vPIM0B2gUoM0GmmWFlseAtdKrp0BX5XqWVSQR4DcXYXLFYyODNVY3sBJ6T1uPRh3LzLPHQIsrfT
//fd state: (82) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30Gtxh8zn6J7WdiSMBQFj8SuA3HOBM6zepiYSauW1c8

ret = client.Write(fd786, []byte("pWje3i0lVeWNWxqLLV9gpn1r2dcmMJ2qwelKOcSYxbsM0WducyudwyKT5VllZV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30Gtxh8zn6J7WdiSMBQFj8SuA3HOBM6zepiYSauW1c8pWje3i0lVeWNWxqLLV9gpn1r2dcmMJ2qwelKOcSYxbsM0WducyudwyKT5VllZV

buf, ret = client.Read(fd797, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (113) kLkr49aYsH1yF3hsFScHbATG64GX0CgKHRU_ALIyg0jV7N7Aoe3poxWWDjik6H8bDFLTafPHEy2aflBByaIqHvczGaIjLdGlOnwF_7T8YNLcjmZDp

ret = client.Write(fd775, []byte("X_B9QzeVbSkBy2VVQdsC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) kLkr49aYsH1yF3hsFScHbATG64GX0CgKHRU_ALIyg0jV7N7Aoe3poxWWDjik6H8bDFLTafPHEy2aflBByaIqHvczGaIjLdGlOnwF_7T8YNLcjmZDpX_B9QzeVbSkBy2VVQdsC
//fd state: (0) 

ret = client.Write(fd802, []byte("lcZKF0mhi97Cgi1ILITVsYGK0UQKZiOnTPsT4aFuEX5jC0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) lcZKF0mhi97Cgi1ILITVsYGK0UQKZiOnTPsT4aFuEX5jC0

fd813 := client.Open("/hNMRPhVf4T.txt", client.O_RDWR|client.O_CREATE)
if(fd813 < 0) {
  panic("open failed")
}


ret = client.Seek(fd723, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Close(fd739)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd811, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd808, []byte("wV5f_ww"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) wV5f_ww
//fd state: (47) QrodyHDh1mphsrl23fmpl1dcHlCPY619mbvedTZdxQKyXAH

ret = client.Write(fd771, []byte("nPlEXJWO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) QrodyHDh1mphsrl23fmpl1dcHlCPY619mbvedTZdxQKyXAHnPlEXJWO
//fd state: (55) 2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNubk0uhZ74HqntI5XDpCHifTJxgFgKRyTQ1eB8bZq1Pc2t3uvsjhGqKbx0V_CnQx27xDTy98s4qzewjk5p50ZZLyj0vU8NSGc

ret = client.Write(fd758, []byte("8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfpr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (121) 2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNu8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTy98s4qzewjk5p50ZZLyj0vU8NSGc

buf, ret = client.Read(fd811, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd735)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd801)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd786, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd780)
if(ret != 0) {
  panic("close failed")
}


fd814 := client.Open("/OFzJgSJSh3.txt", client.O_RDWR|client.O_CREATE)
if(fd814 < 0) {
  panic("open failed")
}


ret = client.Seek(fd778, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd787, []byte("lrCsE5XMyUiVVyvxoTcMC6vm5XSOR1qxsibgD963u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) lrCsE5XMyUiVVyvxoTcMC6vm5XSOR1qxsibgD963u

fd815 := client.Open("/rQron0Xzpa.txt", client.O_RDWR|client.O_CREATE)
if(fd815 < 0) {
  panic("open failed")
}


ret = client.Seek(fd752, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd738, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Close(fd669)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd814, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd816 := client.Open("/kvtM5pSccu.txt", client.O_RDWR|client.O_CREATE)
if(fd816 < 0) {
  panic("open failed")
}


fd817 := client.Open("/1hbgHROuo3.txt", client.O_RDWR|client.O_CREATE)
if(fd817 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd813, []byte("Aey9gnL__Fr4f8dci6QzDSBmDmvF73MFf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) Aey9gnL__Fr4f8dci6QzDSBmDmvF73MFf
//fd state: (41) lrCsE5XMyUiVVyvxoTcMC6vm5XSOR1qxsibgD963u

ret = client.Write(fd787, []byte("cB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) lrCsE5XMyUiVVyvxoTcMC6vm5XSOR1qxsibgD963ucB

buf, ret = client.Read(fd705, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd741, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "utVqi") {
  panic("wrong data returned")
}


ret = client.Seek(fd761, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


buf, ret = client.Read(fd811, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd786, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd733, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}


fd818 := client.Open("/b8aFsajIyy.txt", client.O_RDWR|client.O_CREATE)
if(fd818 < 0) {
  panic("open failed")
}


ret = client.Close(fd814)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd768, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd766, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd761, 77, client.SEEK_SET)
if(ret != 77) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 77)
  panic("seek failed")
}


ret = client.Seek(fd768, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd804, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd724, 207, client.SEEK_SET)
if(ret != 207) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 207)
  panic("seek failed")
}


buf, ret = client.Read(fd724, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7CI7xK_1fLuG795Rb0VVBBIKvf0") {
  panic("wrong data returned")
}


ret = client.Close(fd781)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd752, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd819 := client.Open("/iVIJjg2elX.txt", client.O_RDWR|client.O_CREATE)
if(fd819 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd812, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd820 := client.Open("/nfLU9I_D5X.txt", client.O_RDWR|client.O_CREATE)
if(fd820 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd820, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd821 := client.Open("/3pxUopJGEQ.txt", client.O_RDWR|client.O_CREATE)
if(fd821 < 0) {
  panic("open failed")
}

//fd state: (0) aj_HEn3c2Q3jRGJRSiZSw0cy97_xZgfQgmlTUNGnKvv0DXnSRTyPImwGwglPydD8r0L6Zu5uXBCAvmP8MBTXnK9Pc1KgxCAE_EUI1jy0lOgBv4qt6XGkB_yyNU1Yq7axS5s7R9ms6yW

ret = client.Write(fd800, []byte("Gw0YB4GHWz9jSHBGDOaln7NRJlYUZHEyxHpc6gB37iKV58s5yJmMFFMaurIuFmHAwPqO0SMB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) Gw0YB4GHWz9jSHBGDOaln7NRJlYUZHEyxHpc6gB37iKV58s5yJmMFFMaurIuFmHAwPqO0SMBXBCAvmP8MBTXnK9Pc1KgxCAE_EUI1jy0lOgBv4qt6XGkB_yyNU1Yq7axS5s7R9ms6yW
//fd state: (0) 

ret = client.Write(fd804, []byte("7oDYGJexwvqwx4oJSujfsdQ6kJLvG5l2NJLcQIxv9Kel33SXJrk6elszy3DkV2WUQC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) 7oDYGJexwvqwx4oJSujfsdQ6kJLvG5l2NJLcQIxv9Kel33SXJrk6elszy3DkV2WUQC

ret = client.Close(fd800)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd818, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd772, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd786, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


ret = client.Close(fd819)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd817, []byte("LnqDLBkDevzv6XwA_XVGGdcpwUBBz0QlU_mSHbe895AbpSdQFTPNzHPsIFMBpmr36_mOVFcLM7klho3RSWcIHWLD69eNP3TFD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) LnqDLBkDevzv6XwA_XVGGdcpwUBBz0QlU_mSHbe895AbpSdQFTPNzHPsIFMBpmr36_mOVFcLM7klho3RSWcIHWLD69eNP3TFD

ret = client.Seek(fd812, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd768, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd822 := client.Open("/E64xtwodJu.txt", client.O_RDWR|client.O_CREATE)
if(fd822 < 0) {
  panic("open failed")
}


fd823 := client.Open("/rKfrTDV7yO.txt", client.O_RDWR|client.O_CREATE)
if(fd823 < 0) {
  panic("open failed")
}


ret = client.Seek(fd805, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


buf, ret = client.Read(fd772, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd824 := client.Open("/qSIF4GBaPH.txt", client.O_RDWR|client.O_CREATE)
if(fd824 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd772, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd825 := client.Open("/b_i4yt5U37.txt", client.O_RDWR|client.O_CREATE)
if(fd825 < 0) {
  panic("open failed")
}


ret = client.Seek(fd823, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd747)
if(ret != 0) {
  panic("close failed")
}


fd826 := client.Open("/HCfdrAN7le.txt", client.O_RDWR|client.O_CREATE)
if(fd826 < 0) {
  panic("open failed")
}

//fd state: (45) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30Gtxh8zn6J7WdiSMBQFj8SuA3HOBM6zepiYSauW1c8pWje3i0lVeWNWxqLLV9gpn1r2dcmMJ2qwelKOcSYxbsM0WducyudwyKT5VllZV

ret = client.Write(fd786, []byte("QtrlP142LhdMkIDrDaM5FSicY_x3TWtPrWzMmE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30GtxhQtrlP142LhdMkIDrDaM5FSicY_x3TWtPrWzMmEWje3i0lVeWNWxqLLV9gpn1r2dcmMJ2qwelKOcSYxbsM0WducyudwyKT5VllZV

ret = client.Seek(fd724, 104, client.SEEK_SET)
if(ret != 104) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 104)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd825, []byte("Oc4JpGIv1AMMreB3ZXpCIs1RA7KdfTmJEjx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) Oc4JpGIv1AMMreB3ZXpCIs1RA7KdfTmJEjx

buf, ret = client.Read(fd807, 12)
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

//fd state: (54) sd6FNxfgwCvCVXdP8UxPkqVHOhG1dEQHGAKh3p0zQz7yWmmfkkzJhm9889dnySpN_zIh7X2uJ0jDwxoQ7h33nbM

ret = client.Write(fd795, []byte("Tjg__G7QbwAQ4Rcd4WU0vNMfsW8jG3l2DOSOtmez"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) sd6FNxfgwCvCVXdP8UxPkqVHOhG1dEQHGAKh3p0zQz7yWmmfkkzJhmTjg__G7QbwAQ4Rcd4WU0vNMfsW8jG3l2DOSOtmez

ret = client.Seek(fd807, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd821, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd758, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Ty") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd824, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nJDqsJJJSRkhtb6kPzqUJhaLIvANZ8cMJBGskdrtddCl_8KG") {
  panic("wrong data returned")
}


ret = client.Seek(fd766, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}

//fd state: (81) AD1VDoaF9LR92pHV_YaGSXKEXddaNrlB7vVFDaf5JsgGPjO9L7JtjFAmnL3hBJgGf1s9t3KFZATyV7ymH_tooFtdG0WhKe3

ret = client.Write(fd714, []byte("caU8DT0S1USLH0uKmPfQnn3nmRXxD2GRMzu8CB366g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) AD1VDoaF9LR92pHV_YaGSXKEXddaNrlB7vVFDaf5JsgGPjO9L7JtjFAmnL3hBJgGf1s9t3KFZATyV7ymHcaU8DT0S1USLH0uKmPfQnn3nmRXxD2GRMzu8CB366g

fd827 := client.Open("/HGIG0g8dvV.txt", client.O_RDWR|client.O_CREATE)
if(fd827 < 0) {
  panic("open failed")
}

//fd state: (40) P7SobYghq4JMzrFWBsfMl45wygh0oGU4PsiGV9KjxQHapacNH1DpEsQOqfbsvoEuSOSfO0_64ckayz5sTFow8ljuZqHMMnEnwoccgVFH6O1

ret = client.Write(fd736, []byte("L6Prkt2p0IT4uS2V7VcU5PbTs2q6W02X0Bp3njPV_fvkB_nCiaGce8Lkeop09IYJrcpJKSrJgtIgifhpk4oGFJCgsl1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (131) P7SobYghq4JMzrFWBsfMl45wygh0oGU4PsiGV9KjL6Prkt2p0IT4uS2V7VcU5PbTs2q6W02X0Bp3njPV_fvkB_nCiaGce8Lkeop09IYJrcpJKSrJgtIgifhpk4oGFJCgsl1

ret = client.Close(fd823)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd772)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd733, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "sOyCxLpMV") {
  panic("wrong data returned")
}

//fd state: (0) f_1sI4jviEqzDfKjBzd1eQp9Ygs2_Wxgy_TfMT4jp1y4y0LfHwrlJblkVoQZX6F5_8XXbKWj3BOCmIdEbBGruba2yTvFBV9Ol2Bb11sWhkOKMS49j1tqSocek2fU73flEgL4QcuMSJhYC1a3

ret = client.Write(fd827, []byte("Ln_7gYBQoUmkZzo0HX1psPsQIYn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) Ln_7gYBQoUmkZzo0HX1psPsQIYn2_Wxgy_TfMT4jp1y4y0LfHwrlJblkVoQZX6F5_8XXbKWj3BOCmIdEbBGruba2yTvFBV9Ol2Bb11sWhkOKMS49j1tqSocek2fU73flEgL4QcuMSJhYC1a3
//fd state: (35) Oc4JpGIv1AMMreB3ZXpCIs1RA7KdfTmJEjx

ret = client.Write(fd825, []byte("Hg1p4MN5XWdyF3Rj_kTgqxzO6LL16QafFA7ErBYPA5eAkoKNesyx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) Oc4JpGIv1AMMreB3ZXpCIs1RA7KdfTmJEjxHg1p4MN5XWdyF3Rj_kTgqxzO6LL16QafFA7ErBYPA5eAkoKNesyx

buf, ret = client.Read(fd824, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ie2EENJQlxAIHRhfkAl5oLPatclKh") {
  panic("wrong data returned")
}


ret = client.Seek(fd761, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Close(fd803)
if(ret != 0) {
  panic("close failed")
}


fd828 := client.Open("/gwIXoPWovz.txt", client.O_RDWR|client.O_CREATE)
if(fd828 < 0) {
  panic("open failed")
}


ret = client.Close(fd828)
if(ret != 0) {
  panic("close failed")
}


fd829 := client.Open("/IIgch99vnP.txt", client.O_RDWR|client.O_CREATE)
if(fd829 < 0) {
  panic("open failed")
}


ret = client.Seek(fd825, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}

//fd state: (41) 976FGiAP0xyX3Rm72A2vPIM0B2gUoM0GmmWFlseAtdKrp0BX5XqWVSQR4DcXYXLFYyODNVY3sBJ6T1uPRh3LzLPHQIsrfT

ret = client.Write(fd738, []byte("s3qqts3NLqo56fy77mov8yc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) 976FGiAP0xyX3Rm72A2vPIM0B2gUoM0GmmWFlseAts3qqts3NLqo56fy77mov8ycYyODNVY3sBJ6T1uPRh3LzLPHQIsrfT
//fd state: (123) 2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNu8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTy98s4qzewjk5p50ZZLyj0vU8NSGc

ret = client.Write(fd758, []byte("oMPacd0AoaWs947QQV7TkbBdPG0bMe6bwOtURbM8UZu_cLG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (170) 2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNu8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTyoMPacd0AoaWs947QQV7TkbBdPG0bMe6bwOtURbM8UZu_cLG

buf, ret = client.Read(fd778, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd775, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Seek(fd768, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd830 := client.Open("/eXNFP7XOhv.txt", client.O_RDWR|client.O_CREATE)
if(fd830 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd809, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lMTRZO4nX") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd829, []byte("QtG88bXSQXxhp6sE75_q3vlWUXtVHD1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) QtG88bXSQXxhp6sE75_q3vlWUXtVHD1
//fd state: (64) 976FGiAP0xyX3Rm72A2vPIM0B2gUoM0GmmWFlseAts3qqts3NLqo56fy77mov8ycYyODNVY3sBJ6T1uPRh3LzLPHQIsrfT

ret = client.Write(fd738, []byte("kjh0p43H75rGEk8lwWnHGg0Ysb3RU0UJ1fC02DSspwX3npf6xXyXPcDyDfkwAn8STtf8ksG57jaxNUHxrdsn0yYRfiQdlaIM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (160) 976FGiAP0xyX3Rm72A2vPIM0B2gUoM0GmmWFlseAts3qqts3NLqo56fy77mov8yckjh0p43H75rGEk8lwWnHGg0Ysb3RU0UJ1fC02DSspwX3npf6xXyXPcDyDfkwAn8STtf8ksG57jaxNUHxrdsn0yYRfiQdlaIM

buf, ret = client.Read(fd795, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd831 := client.Open("/oW3LNI8mXJ.txt", client.O_RDWR|client.O_CREATE)
if(fd831 < 0) {
  panic("open failed")
}


fd832 := client.Open("/hiNlZwaae9.txt", client.O_RDWR|client.O_CREATE)
if(fd832 < 0) {
  panic("open failed")
}


ret = client.Seek(fd775, 102, client.SEEK_SET)
if(ret != 102) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 102)
  panic("seek failed")
}


fd833 := client.Open("/OpMYyHTi1Y.txt", client.O_RDWR|client.O_CREATE)
if(fd833 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd768, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "p") {
  panic("wrong data returned")
}


ret = client.Close(fd723)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd770, []byte("JwMUL9H42ItdkqJ1rvhNjcUnYvoZJEYIjk51Hh4aHwsXz_gR_lVjUjHttGfsnRv17Gfl7O084jbllnOa3cY7wbrm6DMJR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) JwMUL9H42ItdkqJ1rvhNjcUnYvoZJEYIjk51Hh4aHwsXz_gR_lVjUjHttGfsnRv17Gfl7O084jbllnOa3cY7wbrm6DMJR

ret = client.Close(fd714)
if(ret != 0) {
  panic("close failed")
}


fd834 := client.Open("/HPOQK_f7t0.txt", client.O_RDWR|client.O_CREATE)
if(fd834 < 0) {
  panic("open failed")
}


ret = client.Seek(fd741, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}


buf, ret = client.Read(fd778, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd797, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (104) sotnjUX9jgZVvt0fEajV5q_qj0aEVVUMewYydS0axEK1m4qbqq2q6rtcpHnRe47vxhutyi2qYIdHKVh2GE4D3x6pNxw6nF5iqp414smJ4cbtmxQ1PKmzaog5DkxE5v8iqb2cV22YBcSTJxuwW2DgEKgX7sfi0qNxw9_Eqdwr29AAFAlkMZWeNAjpZEQ2haIv5xXzYggHeituYFG7CI7xK_1fLuG795Rb0VVBBIKvf0

ret = client.Write(fd724, []byte("qDz2yTb95CZLPpm9v_SZcrG_mQDzNi6W2vGf5cfFENzDidvF4dLMG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (157) sotnjUX9jgZVvt0fEajV5q_qj0aEVVUMewYydS0axEK1m4qbqq2q6rtcpHnRe47vxhutyi2qYIdHKVh2GE4D3x6pNxw6nF5iqp414smJqDz2yTb95CZLPpm9v_SZcrG_mQDzNi6W2vGf5cfFENzDidvF4dLMGqNxw9_Eqdwr29AAFAlkMZWeNAjpZEQ2haIv5xXzYggHeituYFG7CI7xK_1fLuG795Rb0VVBBIKvf0

ret = client.Seek(fd786, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


ret = client.Close(fd765)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd807, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd825, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}


buf, ret = client.Read(fd813, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd805)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd831, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd829, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd835 := client.Open("/SuaaPq1I1G.txt", client.O_RDWR|client.O_CREATE)
if(fd835 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd770, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd836 := client.Open("/lhlbAGE_2Z.txt", client.O_RDWR|client.O_CREATE)
if(fd836 < 0) {
  panic("open failed")
}


fd837 := client.Open("/jLOS6BJRS_.txt", client.O_RDWR|client.O_CREATE)
if(fd837 < 0) {
  panic("open failed")
}


fd838 := client.Open("/3wo0tpz_HN.txt", client.O_RDWR|client.O_CREATE)
if(fd838 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd820, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd826, []byte("NJq0QDVdiMWZtm1xY_AHGe6iZOPa_Bg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) NJq0QDVdiMWZtm1xY_AHGe6iZOPa_Bg

ret = client.Seek(fd833, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd766)
if(ret != 0) {
  panic("close failed")
}


fd839 := client.Open("/NPA2Wk_fzw.txt", client.O_RDWR|client.O_CREATE)
if(fd839 < 0) {
  panic("open failed")
}


ret = client.Close(fd837)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd826, 3)
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


buf, ret = client.Read(fd812, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd808)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd811, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd804, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd810)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd787, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


fd840 := client.Open("/FMo2i7_6at.txt", client.O_RDWR|client.O_CREATE)
if(fd840 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd705, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd836, []byte("5DJs73U"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) 5DJs73U

buf, ret = client.Read(fd834, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd786, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "r5lipqCtc2H9e89X_u30GtxhQtrlP142") {
  panic("wrong data returned")
}


fd841 := client.Open("/hB_gQxlIvw.txt", client.O_RDWR|client.O_CREATE)
if(fd841 < 0) {
  panic("open failed")
}


fd842 := client.Open("/zF2ML7j1Pn.txt", client.O_RDWR|client.O_CREATE)
if(fd842 < 0) {
  panic("open failed")
}

//fd state: (159) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvk

ret = client.Write(fd721, []byte("pDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnHM73Nqy7zCD_dFp28gPbagReoJ35"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (231) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvkpDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnHM73Nqy7zCD_dFp28gPbagReoJ35

ret = client.Close(fd835)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd816, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd843 := client.Open("/7XgDJ0dpBX.txt", client.O_RDWR|client.O_CREATE)
if(fd843 < 0) {
  panic("open failed")
}


ret = client.Close(fd775)
if(ret != 0) {
  panic("close failed")
}

//fd state: (66) 7oDYGJexwvqwx4oJSujfsdQ6kJLvG5l2NJLcQIxv9Kel33SXJrk6elszy3DkV2WUQC

ret = client.Write(fd804, []byte("2flXr0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) 7oDYGJexwvqwx4oJSujfsdQ6kJLvG5l2NJLcQIxv9Kel33SXJrk6elszy3DkV2WUQC2flXr0

buf, ret = client.Read(fd812, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd750, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "J_5eDrzCgiMW7TJcUzvnqPEyb551L8bgk4F9wSe3Hzz3LhnaojqWAUL") {
  panic("wrong data returned")
}


ret = client.Seek(fd831, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd844 := client.Open("/y7zA95PazD.txt", client.O_RDWR|client.O_CREATE)
if(fd844 < 0) {
  panic("open failed")
}

//fd state: (183) cM24y1j_uq88pmzWKopeif5aYtd8Ex1oQRpegSjsFB8tVNxJKLnV7gMzdkmZAdxffDfxnDQZjLqkaYKIDsk1d1HWtOZ4u8zpEV8kdW3mtk_Jhl9Nys2dXWCh3XxWcRuNMeZiUQe_f8Q7FIj0i6EUKpg2KiR_wtmFwWGMnBlUVJKgLLt96v8dvui

ret = client.Write(fd796, []byte("sXRh5CAB1aWhpCcIVjBRktgyXXYkJITRAnWM4SJh6b2ycFxlk1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (233) cM24y1j_uq88pmzWKopeif5aYtd8Ex1oQRpegSjsFB8tVNxJKLnV7gMzdkmZAdxffDfxnDQZjLqkaYKIDsk1d1HWtOZ4u8zpEV8kdW3mtk_Jhl9Nys2dXWCh3XxWcRuNMeZiUQe_f8Q7FIj0i6EUKpg2KiR_wtmFwWGMnBlUVJKgLLt96v8dvuisXRh5CAB1aWhpCcIVjBRktgyXXYkJITRAnWM4SJh6b2ycFxlk1

buf, ret = client.Read(fd724, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qNxw9_Eqdwr29AAFAlkMZWeNAjpZEQ2haIv5xXzYggHeituYFG7CI7xK_1fLuG795Rb0VVBBIKvf0") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd768, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd816, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd839, []byte("zoscQ1PHj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) zoscQ1PHj
//fd state: (0) 

ret = client.Write(fd752, []byte("qjprDe4xMUa984v92MZWR_zPK7k8sFRbZfQkwPTDiNIE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) qjprDe4xMUa984v92MZWR_zPK7k8sFRbZfQkwPTDiNIE

buf, ret = client.Read(fd778, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd845 := client.Open("/zAAGw1724E.txt", client.O_RDWR|client.O_CREATE)
if(fd845 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd827, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2_Wxgy_TfMT4jp1y4y0LfHwrlJblkVoQZX6F5_8XXbKWj3BOCmIdEbBGruba2yTvFBV9Ol2Bb1") {
  panic("wrong data returned")
}


ret = client.Seek(fd797, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Seek(fd721, 156, client.SEEK_SET)
if(ret != 156) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 156)
  panic("seek failed")
}


ret = client.Seek(fd736, 73, client.SEEK_SET)
if(ret != 73) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 73)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd842, []byte("8IpDMCh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) 8IpDMCh
//fd state: (0) 

ret = client.Write(fd833, []byte("gOFFtD8jg3I2zsIVsneVa66kGJyFM92eqIJcw7HBfAlyQ1FSivTpWbRxMKOB1adgc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) gOFFtD8jg3I2zsIVsneVa66kGJyFM92eqIJcw7HBfAlyQ1FSivTpWbRxMKOB1adgc

fd846 := client.Open("/bSVI9qXsyi.txt", client.O_RDWR|client.O_CREATE)
if(fd846 < 0) {
  panic("open failed")
}


fd847 := client.Open("/tnQZHPR9ma.txt", client.O_RDWR|client.O_CREATE)
if(fd847 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd845, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd705)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd752, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd826)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd831, []byte("J"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) J

ret = client.Close(fd724)
if(ret != 0) {
  panic("close failed")
}

//fd state: (93) LBJmwbSgSfg2jCol1_N5PaEQJhmQCNW9bnCzJKJ_5eDrzCgiMW7TJcUzvnqPEyb551L8bgk4F9wSe3Hzz3LhnaojqWAUL

ret = client.Write(fd750, []byte("5P9N1sabnNxBfVUIUhPwZuCgX7ZLbLLbQXW4go0FsXDTbd0r6FOuA2JDrhRGN43V3OXjAY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (163) LBJmwbSgSfg2jCol1_N5PaEQJhmQCNW9bnCzJKJ_5eDrzCgiMW7TJcUzvnqPEyb551L8bgk4F9wSe3Hzz3LhnaojqWAUL5P9N1sabnNxBfVUIUhPwZuCgX7ZLbLLbQXW4go0FsXDTbd0r6FOuA2JDrhRGN43V3OXjAY

ret = client.Seek(fd736, 98, client.SEEK_SET)
if(ret != 98) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 98)
  panic("seek failed")
}


buf, ret = client.Read(fd761, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4QWu4TJIdsYHGcQinFgU8MQJB5c7C904") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd833, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd848 := client.Open("/icy2T1pZHr.txt", client.O_RDWR|client.O_CREATE)
if(fd848 < 0) {
  panic("open failed")
}


ret = client.Close(fd813)
if(ret != 0) {
  panic("close failed")
}

//fd state: (3) aUDTsvTIMxAZX8Mr8rOZC78P04rRjwpCKSwaQLUbIgUANpUps0ZjE

ret = client.Write(fd797, []byte("NQLTgh0CQqg76G49k69EtC6ZDfnTWIobWoePX07VM58EXwFrGAb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) aUDNQLTgh0CQqg76G49k69EtC6ZDfnTWIobWoePX07VM58EXwFrGAb

ret = client.Close(fd804)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd741)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd787, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}

//fd state: (53) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30GtxhQtrlP142LhdMkIDrDaM5FSicY_x3TWtPrWzMmEWje3i0lVeWNWxqLLV9gpn1r2dcmMJ2qwelKOcSYxbsM0WducyudwyKT5VllZV

ret = client.Write(fd786, []byte("w7srRNftD1NREZeWGfgwO60u6_pT29is8ytVQRhYXd8TcnMPJvW6J61"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (108) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30GtxhQtrlP142w7srRNftD1NREZeWGfgwO60u6_pT29is8ytVQRhYXd8TcnMPJvW6J61cmMJ2qwelKOcSYxbsM0WducyudwyKT5VllZV
//fd state: (0) 

ret = client.Write(fd822, []byte("mSHpTL3TgLJJKYpSNljGuD0LsOtGaFkYELrPpjoFcBvgu14trqFksvmU32azQjbTsjkj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) mSHpTL3TgLJJKYpSNljGuD0LsOtGaFkYELrPpjoFcBvgu14trqFksvmU32azQjbTsjkj
//fd state: (77) nJDqsJJJSRkhtb6kPzqUJhaLIvANZ8cMJBGskdrtddCl_8KGie2EENJQlxAIHRhfkAl5oLPatclKh

ret = client.Write(fd824, []byte("29leb_kiuOJOHC5sWupxwgUFzcVtpIdWVL4qw8VVYD2LfbfRZ6yMwoiGBl3MGs1rTGp1ESM1Zaz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) nJDqsJJJSRkhtb6kPzqUJhaLIvANZ8cMJBGskdrtddCl_8KGie2EENJQlxAIHRhfkAl5oLPatclKh29leb_kiuOJOHC5sWupxwgUFzcVtpIdWVL4qw8VVYD2LfbfRZ6yMwoiGBl3MGs1rTGp1ESM1Zaz

buf, ret = client.Read(fd811, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd849 := client.Open("/XdPxjeqcPT.txt", client.O_RDWR|client.O_CREATE)
if(fd849 < 0) {
  panic("open failed")
}

//fd state: (9) zoscQ1PHj

ret = client.Write(fd839, []byte("QrsT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) zoscQ1PHjQrsT

ret = client.Close(fd733)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd820, []byte("2_RXL6jvtGher"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) 2_RXL6jvtGher
//fd state: (0) 

ret = client.Write(fd840, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) 

ret = client.Seek(fd844, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd811)
if(ret != 0) {
  panic("close failed")
}


fd850 := client.Open("/kYPAZl3Z2l.txt", client.O_RDWR|client.O_CREATE)
if(fd850 < 0) {
  panic("open failed")
}


ret = client.Seek(fd824, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


ret = client.Seek(fd825, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


buf, ret = client.Read(fd817, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd832)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd812)
if(ret != 0) {
  panic("close failed")
}

//fd state: (108) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30GtxhQtrlP142w7srRNftD1NREZeWGfgwO60u6_pT29is8ytVQRhYXd8TcnMPJvW6J61cmMJ2qwelKOcSYxbsM0WducyudwyKT5VllZV

ret = client.Write(fd786, []byte("ER1jSz_85SRWunMrnfl5nfjWY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30GtxhQtrlP142w7srRNftD1NREZeWGfgwO60u6_pT29is8ytVQRhYXd8TcnMPJvW6J61ER1jSz_85SRWunMrnfl5nfjWYdwyKT5VllZV

buf, ret = client.Read(fd795, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd797, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Close(fd818)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd738, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd849, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}

//fd state: (0) hym1SU_hVIcDh2tq6wF7TxZCE8ygNH820R8K7

ret = client.Write(fd830, []byte("oy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) oym1SU_hVIcDh2tq6wF7TxZCE8ygNH820R8K7

fd851 := client.Open("/KCjs7VRuKx.txt", client.O_RDWR|client.O_CREATE)
if(fd851 < 0) {
  panic("open failed")
}

//fd state: (2) oym1SU_hVIcDh2tq6wF7TxZCE8ygNH820R8K7

ret = client.Write(fd830, []byte("kmPCUCvydpSD2D292vig8_od4pclDE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) oykmPCUCvydpSD2D292vig8_od4pclDE0R8K7

buf, ret = client.Read(fd822, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd778)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd836, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


buf, ret = client.Read(fd838, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd852 := client.Open("/RoesxV01Tq.txt", client.O_RDWR|client.O_CREATE)
if(fd852 < 0) {
  panic("open failed")
}


fd853 := client.Open("/Q1IxtFbzhH.txt", client.O_RDWR|client.O_CREATE)
if(fd853 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd721, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uvkpDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnH") {
  panic("wrong data returned")
}


ret = client.Close(fd848)
if(ret != 0) {
  panic("close failed")
}


fd854 := client.Open("/iQDk2CrWqb.txt", client.O_RDWR|client.O_CREATE)
if(fd854 < 0) {
  panic("open failed")
}


ret = client.Close(fd850)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd825)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd833)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd852)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd831)
if(ret != 0) {
  panic("close failed")
}

//fd state: (204) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvkpDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnHM73Nqy7zCD_dFp28gPbagReoJ35

ret = client.Write(fd721, []byte("TXrKZ36LLYSbBXx3wXRCtaBOjbCIsGOSPNShJck86nbR6FKGxjZqg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (257) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvkpDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnHTXrKZ36LLYSbBXx3wXRCtaBOjbCIsGOSPNShJck86nbR6FKGxjZqg

ret = client.Close(fd736)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd797, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


ret = client.Close(fd841)
if(ret != 0) {
  panic("close failed")
}

//fd state: (170) 2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNu8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTyoMPacd0AoaWs947QQV7TkbBdPG0bMe6bwOtURbM8UZu_cLG

ret = client.Write(fd758, []byte("5OjhIwmkICcI5CxE_xp3RGWfoK_ZQIuENtpE6xWLHb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (212) 2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNu8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTyoMPacd0AoaWs947QQV7TkbBdPG0bMe6bwOtURbM8UZu_cLG5OjhIwmkICcI5CxE_xp3RGWfoK_ZQIuENtpE6xWLHb

fd855 := client.Open("/jNEtK7FeCA.txt", client.O_RDWR|client.O_CREATE)
if(fd855 < 0) {
  panic("open failed")
}

//fd state: (13) zoscQ1PHjQrsT

ret = client.Write(fd839, []byte("gjx9LFLVNLBR4SuOQ5z0j2G3xJl_UKkOK1UubeXBDsqUJovSMhZ1gGYS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) zoscQ1PHjQrsTgjx9LFLVNLBR4SuOQ5z0j2G3xJl_UKkOK1UubeXBDsqUJovSMhZ1gGYS

ret = client.Seek(fd855, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (55) nJDqsJJJSRkhtb6kPzqUJhaLIvANZ8cMJBGskdrtddCl_8KGie2EENJQlxAIHRhfkAl5oLPatclKh29leb_kiuOJOHC5sWupxwgUFzcVtpIdWVL4qw8VVYD2LfbfRZ6yMwoiGBl3MGs1rTGp1ESM1Zaz

ret = client.Write(fd824, []byte("oAx9cDbcJn3GeGIz6ZwS_q4JCX3255259L_kqWBiq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) nJDqsJJJSRkhtb6kPzqUJhaLIvANZ8cMJBGskdrtddCl_8KGie2EENJoAx9cDbcJn3GeGIz6ZwS_q4JCX3255259L_kqWBiqxwgUFzcVtpIdWVL4qw8VVYD2LfbfRZ6yMwoiGBl3MGs1rTGp1ESM1Zaz

buf, ret = client.Read(fd827, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1sWhkOKMS49j1tqSocek2fU73flEgL4QcuMSJhYC1a3") {
  panic("wrong data returned")
}


ret = client.Seek(fd815, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd761, 120, client.SEEK_SET)
if(ret != 120) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 120)
  panic("seek failed")
}


ret = client.Close(fd750)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd822, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd770, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd768, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd758, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (31) QtG88bXSQXxhp6sE75_q3vlWUXtVHD1

ret = client.Write(fd829, []byte("D9_Z5R0okaXmi_VZYTE4z6jfx68XrYFot9qEtVYPrH69_UgbcCeE1y9zyqUnxwQ9LRf7lFDhe4d3HhLUlMImrp_cZd5io"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) QtG88bXSQXxhp6sE75_q3vlWUXtVHD1D9_Z5R0okaXmi_VZYTE4z6jfx68XrYFot9qEtVYPrH69_UgbcCeE1y9zyqUnxwQ9LRf7lFDhe4d3HhLUlMImrp_cZd5io

ret = client.Close(fd840)
if(ret != 0) {
  panic("close failed")
}

//fd state: (94) sd6FNxfgwCvCVXdP8UxPkqVHOhG1dEQHGAKh3p0zQz7yWmmfkkzJhmTjg__G7QbwAQ4Rcd4WU0vNMfsW8jG3l2DOSOtmez

ret = client.Write(fd795, []byte("FmvcwNBHkDgU0poADx56N30yRvbEU2LJrBNbz0T2NHh90EhIygZQRrtRqrGtqbqBuO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (160) sd6FNxfgwCvCVXdP8UxPkqVHOhG1dEQHGAKh3p0zQz7yWmmfkkzJhmTjg__G7QbwAQ4Rcd4WU0vNMfsW8jG3l2DOSOtmezFmvcwNBHkDgU0poADx56N30yRvbEU2LJrBNbz0T2NHh90EhIygZQRrtRqrGtqbqBuO
//fd state: (32) aUDNQLTgh0CQqg76G49k69EtC6ZDfnTWIobWoePX07VM58EXwFrGAb

ret = client.Write(fd797, []byte("z3k"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) aUDNQLTgh0CQqg76G49k69EtC6ZDfnTWz3kWoePX07VM58EXwFrGAb

ret = client.Close(fd787)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd815, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (35) aUDNQLTgh0CQqg76G49k69EtC6ZDfnTWz3kWoePX07VM58EXwFrGAb

ret = client.Write(fd797, []byte("iR3QUR2sPaNAhz2DaNLRR5OMS3O0BZpRrVh_1KvPYJEj4hl7XZUMq8tlXltgPoCwcJROTBbFkcvUpDm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) aUDNQLTgh0CQqg76G49k69EtC6ZDfnTWz3kiR3QUR2sPaNAhz2DaNLRR5OMS3O0BZpRrVh_1KvPYJEj4hl7XZUMq8tlXltgPoCwcJROTBbFkcvUpDm

ret = client.Seek(fd851, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd856 := client.Open("/rDZVcO3UQt.txt", client.O_RDWR|client.O_CREATE)
if(fd856 < 0) {
  panic("open failed")
}


fd857 := client.Open("/QUQRCvV0rG.txt", client.O_RDWR|client.O_CREATE)
if(fd857 < 0) {
  panic("open failed")
}


ret = client.Close(fd839)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd834, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd858 := client.Open("/JqZAq55249.txt", client.O_RDWR|client.O_CREATE)
if(fd858 < 0) {
  panic("open failed")
}


ret = client.Close(fd809)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd847, []byte("L5tUQuLK2HD6vJFcMCsnlp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) L5tUQuLK2HD6vJFcMCsnlp

buf, ret = client.Read(fd836, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3U") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd853, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd859 := client.Open("/eXNFP7XOhv.txt", client.O_RDWR|client.O_CREATE)
if(fd859 < 0) {
  panic("open failed")
}


ret = client.Seek(fd824, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


buf, ret = client.Read(fd758, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd830, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


ret = client.Close(fd855)
if(ret != 0) {
  panic("close failed")
}


fd860 := client.Open("/rHbNRD2SCH.txt", client.O_RDWR|client.O_CREATE)
if(fd860 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd836, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd829, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd861 := client.Open("/8Evde1kBxw.txt", client.O_RDWR|client.O_CREATE)
if(fd861 < 0) {
  panic("open failed")
}


ret = client.Close(fd856)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd838)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd761)
if(ret != 0) {
  panic("close failed")
}


fd862 := client.Open("/IPI4ET9ttc.txt", client.O_RDWR|client.O_CREATE)
if(fd862 < 0) {
  panic("open failed")
}


ret = client.Seek(fd844, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd861, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (7) 8IpDMCh

ret = client.Write(fd842, []byte("4wLVu9oGqdbxEhEH2oNVAZMH7fD_xqINY0OQ01FuEcv8vbUwbj0woVqxTByhmvpJHINNBaUpDk_jmPIXoKg_jUjHsd9IfJh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) 8IpDMCh4wLVu9oGqdbxEhEH2oNVAZMH7fD_xqINY0OQ01FuEcv8vbUwbj0woVqxTByhmvpJHINNBaUpDk_jmPIXoKg_jUjHsd9IfJh

buf, ret = client.Read(fd795, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd851, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd816, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd820, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


fd863 := client.Open("/egMmr_240Y.txt", client.O_RDWR|client.O_CREATE)
if(fd863 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd851, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd864 := client.Open("/GAv93wc1kq.txt", client.O_RDWR|client.O_CREATE)
if(fd864 < 0) {
  panic("open failed")
}


ret = client.Close(fd738)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd816)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd842, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd842)
if(ret != 0) {
  panic("close failed")
}


fd865 := client.Open("/Kbe_5esbc2.txt", client.O_RDWR|client.O_CREATE)
if(fd865 < 0) {
  panic("open failed")
}

//fd state: (124) QtG88bXSQXxhp6sE75_q3vlWUXtVHD1D9_Z5R0okaXmi_VZYTE4z6jfx68XrYFot9qEtVYPrH69_UgbcCeE1y9zyqUnxwQ9LRf7lFDhe4d3HhLUlMImrp_cZd5io

ret = client.Write(fd829, []byte("wG8NMLsMZvdVy5vTs8JHUldObd1za3x2FgewN7uygmGGkE24JUH_ezLrfu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (182) QtG88bXSQXxhp6sE75_q3vlWUXtVHD1D9_Z5R0okaXmi_VZYTE4z6jfx68XrYFot9qEtVYPrH69_UgbcCeE1y9zyqUnxwQ9LRf7lFDhe4d3HhLUlMImrp_cZd5iowG8NMLsMZvdVy5vTs8JHUldObd1za3x2FgewN7uygmGGkE24JUH_ezLrfu

ret = client.Close(fd854)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd768)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd845, []byte("qPbADj6W9q1wh5JXzMmWehU9aO6y5Vary5ojgNoDDq7S9U5MhhiGn5o6njWR95SGjC46Vzq8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) qPbADj6W9q1wh5JXzMmWehU9aO6y5Vary5ojgNoDDq7S9U5MhhiGn5o6njWR95SGjC46Vzq8

ret = client.Close(fd836)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd864)
if(ret != 0) {
  panic("close failed")
}


fd866 := client.Open("/zkSPrByLg8.txt", client.O_RDWR|client.O_CREATE)
if(fd866 < 0) {
  panic("open failed")
}

//fd state: (212) 2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNu8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTyoMPacd0AoaWs947QQV7TkbBdPG0bMe6bwOtURbM8UZu_cLG5OjhIwmkICcI5CxE_xp3RGWfoK_ZQIuENtpE6xWLHb

ret = client.Write(fd758, []byte("IHStPbR5q4SlHtYFxjK42PoiFf0qHKgixrSnWK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (250) 2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNu8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTyoMPacd0AoaWs947QQV7TkbBdPG0bMe6bwOtURbM8UZu_cLG5OjhIwmkICcI5CxE_xp3RGWfoK_ZQIuENtpE6xWLHbIHStPbR5q4SlHtYFxjK42PoiFf0qHKgixrSnWK

ret = client.Close(fd796)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd863, []byte("OeIf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) OeIf

fd867 := client.Open("/3wjB6vpdmW.txt", client.O_RDWR|client.O_CREATE)
if(fd867 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd752, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qjprDe4xMUa984v92MZWR_zPK7k8sFRbZfQkwPTDiNIE") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd865, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "0SWYabsv5WW0JC8r") {
  panic("wrong data returned")
}


fd868 := client.Open("/TZxRytMWBY.txt", client.O_RDWR|client.O_CREATE)
if(fd868 < 0) {
  panic("open failed")
}


ret = client.Close(fd795)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd820, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "her") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd849, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9CJGZSmservB2DkEN06xeAhcgyLZn59ht") {
  panic("wrong data returned")
}


ret = client.Seek(fd865, 76, client.SEEK_SET)
if(ret != 76) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 76)
  panic("seek failed")
}


buf, ret = client.Read(fd851, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd869 := client.Open("/SbJ286MbJD.txt", client.O_RDWR|client.O_CREATE)
if(fd869 < 0) {
  panic("open failed")
}


ret = client.Close(fd858)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd845, 65, client.SEEK_SET)
if(ret != 65) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 65)
  panic("seek failed")
}


fd870 := client.Open("/BU8ZeUXfGw.txt", client.O_RDWR|client.O_CREATE)
if(fd870 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd830, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lDE0R8K7") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd867, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NFUFRSpuK") {
  panic("wrong data returned")
}


ret = client.Seek(fd861, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd830, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


buf, ret = client.Read(fd847, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd821)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd870)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd867, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "CMBGLf1Bcveazsss") {
  panic("wrong data returned")
}


fd871 := client.Open("/eiV0wchkfY.txt", client.O_RDWR|client.O_CREATE)
if(fd871 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd851, []byte("l33"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) l33

buf, ret = client.Read(fd845, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "C46Vzq8") {
  panic("wrong data returned")
}

//fd state: (21) nJDqsJJJSRkhtb6kPzqUJhaLIvANZ8cMJBGskdrtddCl_8KGie2EENJoAx9cDbcJn3GeGIz6ZwS_q4JCX3255259L_kqWBiqxwgUFzcVtpIdWVL4qw8VVYD2LfbfRZ6yMwoiGBl3MGs1rTGp1ESM1Zaz

ret = client.Write(fd824, []byte("Ex5NuRWYvFwcOchs4lRvctfeJTl9OnhjFAe2MkKAZykvU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) nJDqsJJJSRkhtb6kPzqUJEx5NuRWYvFwcOchs4lRvctfeJTl9OnhjFAe2MkKAZykvUGeGIz6ZwS_q4JCX3255259L_kqWBiqxwgUFzcVtpIdWVL4qw8VVYD2LfbfRZ6yMwoiGBl3MGs1rTGp1ESM1Zaz

buf, ret = client.Read(fd822, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd822, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd822, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd851)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd752, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd857)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd866, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd872 := client.Open("/HqELFn9BNU.txt", client.O_RDWR|client.O_CREATE)
if(fd872 < 0) {
  panic("open failed")
}


ret = client.Close(fd829)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd867, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7") {
  panic("wrong data returned")
}


ret = client.Close(fd820)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd868, []byte("511A0eB4yrGZVQIogt3X_zqEtkolONd1opUSaGHLzzC8k46Hk8FbvK9OZ7hOg0967Wm_QMqyEL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) 511A0eB4yrGZVQIogt3X_zqEtkolONd1opUSaGHLzzC8k46Hk8FbvK9OZ7hOg0967Wm_QMqyEL

ret = client.Seek(fd843, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd721, 213, client.SEEK_SET)
if(ret != 213) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 213)
  panic("seek failed")
}


buf, ret = client.Read(fd817, 86)
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


ret = client.Seek(fd865, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


ret = client.Seek(fd786, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}

//fd state: (213) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvkpDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnHTXrKZ36LLYSbBXx3wXRCtaBOjbCIsGOSPNShJck86nbR6FKGxjZqg

ret = client.Write(fd721, []byte("aUGzNLKGp1U5g_uO7XZgvvmeygrOT7dd85KwS67No5xqQeuRvWR6N"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (266) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvkpDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnHTXrKZ36LLaUGzNLKGp1U5g_uO7XZgvvmeygrOT7dd85KwS67No5xqQeuRvWR6N

ret = client.Seek(fd869, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd873 := client.Open("/R6UEMyClwH.txt", client.O_RDWR|client.O_CREATE)
if(fd873 < 0) {
  panic("open failed")
}

//fd state: (74) 511A0eB4yrGZVQIogt3X_zqEtkolONd1opUSaGHLzzC8k46Hk8FbvK9OZ7hOg0967Wm_QMqyEL

ret = client.Write(fd868, []byte("f3yeJ0ShDIK5HHTn3g48KzZgUPfHlCE1dSC7SOi9hgbVnfPlx_sAXbr6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) 511A0eB4yrGZVQIogt3X_zqEtkolONd1opUSaGHLzzC8k46Hk8FbvK9OZ7hOg0967Wm_QMqyELf3yeJ0ShDIK5HHTn3g48KzZgUPfHlCE1dSC7SOi9hgbVnfPlx_sAXbr6

buf, ret = client.Read(fd863, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd867, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gcTKjXvbKKsfmnyshq8ZZ071jq9nIO45sUhZrUK_n03KYWO") {
  panic("wrong data returned")
}


fd874 := client.Open("/2baQcAu8dc.txt", client.O_RDWR|client.O_CREATE)
if(fd874 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd871, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd846, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd797, 90, client.SEEK_SET)
if(ret != 90) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 90)
  panic("seek failed")
}


ret = client.Close(fd849)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd867, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


fd875 := client.Open("/A1A8lCyTOn.txt", client.O_RDWR|client.O_CREATE)
if(fd875 < 0) {
  panic("open failed")
}


ret = client.Seek(fd860, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd822)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd830)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) lcZKF0mhi97Cgi1ILITVsYGK0UQKZiOnTPsT4aFuEX5jC0

ret = client.Write(fd875, []byte("XBxpOKq_Rbp_FkXaIGsIbBXsyz8COnrcdCQycEJWURKbThoEjgHUKaS4bzchiRcI4D"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) XBxpOKq_Rbp_FkXaIGsIbBXsyz8COnrcdCQycEJWURKbThoEjgHUKaS4bzchiRcI4D

fd876 := client.Open("/FF3gQ3UNK5.txt", client.O_RDWR|client.O_CREATE)
if(fd876 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd861, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd853, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (266) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvkpDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnHTXrKZ36LLaUGzNLKGp1U5g_uO7XZgvvmeygrOT7dd85KwS67No5xqQeuRvWR6N

ret = client.Write(fd721, []byte("_b2ddSYUuoogejXsLdKclouI4Dz2acja"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (298) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvkpDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnHTXrKZ36LLaUGzNLKGp1U5g_uO7XZgvvmeygrOT7dd85KwS67No5xqQeuRvWR6N_b2ddSYUuoogejXsLdKclouI4Dz2acja
//fd state: (44) qjprDe4xMUa984v92MZWR_zPK7k8sFRbZfQkwPTDiNIE

ret = client.Write(fd752, []byte("3g0GiuFuJTA53rXi5RvXbMlWXNWqz4JCG1LkKsOi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) qjprDe4xMUa984v92MZWR_zPK7k8sFRbZfQkwPTDiNIE3g0GiuFuJTA53rXi5RvXbMlWXNWqz4JCG1LkKsOi

buf, ret = client.Read(fd770, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd861, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd877 := client.Open("/4VXKthGmF8.txt", client.O_RDWR|client.O_CREATE)
if(fd877 < 0) {
  panic("open failed")
}

//fd state: (84) qjprDe4xMUa984v92MZWR_zPK7k8sFRbZfQkwPTDiNIE3g0GiuFuJTA53rXi5RvXbMlWXNWqz4JCG1LkKsOi

ret = client.Write(fd752, []byte("XvHMGgdxCpKSTu7vJGhkrXKRzAZaL3O5gKtZo0j_AfNtAlvLz73wPKTl3zhfiRCLWpT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) qjprDe4xMUa984v92MZWR_zPK7k8sFRbZfQkwPTDiNIE3g0GiuFuJTA53rXi5RvXbMlWXNWqz4JCG1LkKsOiXvHMGgdxCpKSTu7vJGhkrXKRzAZaL3O5gKtZo0j_AfNtAlvLz73wPKTl3zhfiRCLWpT
//fd state: (0) 

ret = client.Write(fd844, []byte("9P8kT_aEDa_eiIAyJ4_Xn_16VxVd7UPydh1bTHDk1m_WBCSRtTosLTcpCePqppBBwovkNYOS1ebhFCOdXJrt5xw3OSrCk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) 9P8kT_aEDa_eiIAyJ4_Xn_16VxVd7UPydh1bTHDk1m_WBCSRtTosLTcpCePqppBBwovkNYOS1ebhFCOdXJrt5xw3OSrCk
//fd state: (97) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30GtxhQtrlP142w7srRNftD1NREZeWGfgwO60u6_pT29is8ytVQRhYXd8TcnMPJvW6J61ER1jSz_85SRWunMrnfl5nfjWYdwyKT5VllZV

ret = client.Write(fd786, []byte("DbiBFshzTZvoqUrAhsCzIqPDgU9SlDgJMC27Fot"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) ToxY_JjlZ5sK9hv1qJq4cr5lipqCtc2H9e89X_u30GtxhQtrlP142w7srRNftD1NREZeWGfgwO60u6_pT29is8ytVQRhYXd8TDbiBFshzTZvoqUrAhsCzIqPDgU9SlDgJMC27FotKT5VllZV

ret = client.Close(fd877)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd721)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd847)
if(ret != 0) {
  panic("close failed")
}

//fd state: (151) qjprDe4xMUa984v92MZWR_zPK7k8sFRbZfQkwPTDiNIE3g0GiuFuJTA53rXi5RvXbMlWXNWqz4JCG1LkKsOiXvHMGgdxCpKSTu7vJGhkrXKRzAZaL3O5gKtZo0j_AfNtAlvLz73wPKTl3zhfiRCLWpT

ret = client.Write(fd752, []byte("QHVQVsaqySv0tP8K6xzbM8somnXSm6P1CZbLKJYrkkEaxaZ5NkG4BoTjmhTg_v"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (213) qjprDe4xMUa984v92MZWR_zPK7k8sFRbZfQkwPTDiNIE3g0GiuFuJTA53rXi5RvXbMlWXNWqz4JCG1LkKsOiXvHMGgdxCpKSTu7vJGhkrXKRzAZaL3O5gKtZo0j_AfNtAlvLz73wPKTl3zhfiRCLWpTQHVQVsaqySv0tP8K6xzbM8somnXSm6P1CZbLKJYrkkEaxaZ5NkG4BoTjmhTg_v

buf, ret = client.Read(fd869, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd853)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd876, []byte("m2q4LBpJo0JNzwV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) m2q4LBpJo0JNzwV
//fd state: (0) 

ret = client.Write(fd846, []byte("GUwUWfTrA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) GUwUWfTrA

buf, ret = client.Read(fd844, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd752, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd878 := client.Open("/HvFEDPCDzL.txt", client.O_RDWR|client.O_CREATE)
if(fd878 < 0) {
  panic("open failed")
}


fd879 := client.Open("/0tuonfQNeH.txt", client.O_RDWR|client.O_CREATE)
if(fd879 < 0) {
  panic("open failed")
}

//fd state: (15) m2q4LBpJo0JNzwV

ret = client.Write(fd876, []byte("mMRzF1QVTpxPir95twTmkXnMYThefbszczylYaSHF26uwFH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) m2q4LBpJo0JNzwVmMRzF1QVTpxPir95twTmkXnMYThefbszczylYaSHF26uwFH

buf, ret = client.Read(fd879, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "msXnYYlR6K4E") {
  panic("wrong data returned")
}

//fd state: (62) m2q4LBpJo0JNzwVmMRzF1QVTpxPir95twTmkXnMYThefbszczylYaSHF26uwFH

ret = client.Write(fd876, []byte("uJvkLOSLVBxEV8X5VfUaZdlesS0y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) m2q4LBpJo0JNzwVmMRzF1QVTpxPir95twTmkXnMYThefbszczylYaSHF26uwFHuJvkLOSLVBxEV8X5VfUaZdlesS0y

ret = client.Seek(fd863, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (66) XBxpOKq_Rbp_FkXaIGsIbBXsyz8COnrcdCQycEJWURKbThoEjgHUKaS4bzchiRcI4D

ret = client.Write(fd875, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) XBxpOKq_Rbp_FkXaIGsIbBXsyz8COnrcdCQycEJWURKbThoEjgHUKaS4bzchiRcI4D

ret = client.Close(fd871)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd752, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd880 := client.Open("/34gBdIPqJV.txt", client.O_RDWR|client.O_CREATE)
if(fd880 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd880, []byte("wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo29"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo29

ret = client.Close(fd846)
if(ret != 0) {
  panic("close failed")
}

//fd state: (72) qPbADj6W9q1wh5JXzMmWehU9aO6y5Vary5ojgNoDDq7S9U5MhhiGn5o6njWR95SGjC46Vzq8

ret = client.Write(fd845, []byte("4yKqxy1eOMhPh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) qPbADj6W9q1wh5JXzMmWehU9aO6y5Vary5ojgNoDDq7S9U5MhhiGn5o6njWR95SGjC46Vzq84yKqxy1eOMhPh

fd881 := client.Open("/Kavr4vF_ms.txt", client.O_RDWR|client.O_CREATE)
if(fd881 < 0) {
  panic("open failed")
}


ret = client.Seek(fd880, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Close(fd824)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd786)
if(ret != 0) {
  panic("close failed")
}


fd882 := client.Open("/uJt5Xeux53.txt", client.O_RDWR|client.O_CREATE)
if(fd882 < 0) {
  panic("open failed")
}

//fd state: (0) V2nrgBzCgPIG9y_MhUy9VXUPFNw7eTTT

ret = client.Write(fd881, []byte("8fQN2cNMrTU7wUkxYkVFGn4vj82oOeEUFHLV6swyxiZx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) 8fQN2cNMrTU7wUkxYkVFGn4vj82oOeEUFHLV6swyxiZx

ret = client.Seek(fd867, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


fd883 := client.Open("/oTr3IC8tAY.txt", client.O_RDWR|client.O_CREATE)
if(fd883 < 0) {
  panic("open failed")
}


fd884 := client.Open("/RfchsYU08E.txt", client.O_RDWR|client.O_CREATE)
if(fd884 < 0) {
  panic("open failed")
}


ret = client.Close(fd807)
if(ret != 0) {
  panic("close failed")
}


fd885 := client.Open("/7JfTK64J_0.txt", client.O_RDWR|client.O_CREATE)
if(fd885 < 0) {
  panic("open failed")
}


ret = client.Seek(fd868, 88, client.SEEK_SET)
if(ret != 88) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 88)
  panic("seek failed")
}


ret = client.Seek(fd866, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd859)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd881)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd817, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


buf, ret = client.Read(fd869, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd879, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Seek(fd869, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd880, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo29") {
  panic("wrong data returned")
}

//fd state: (5) msXnYYlR6K4E

ret = client.Write(fd879, []byte("W09H7U"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) msXnYW09H7UE

fd886 := client.Open("/CKytIwcwjN.txt", client.O_RDWR|client.O_CREATE)
if(fd886 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd882, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3Udrlv348DOd7ipdzCKL8KUeTwzmJZsX7S144SUXGgbuznyPWH0V3pHDgBgfjolcL9EzbDurNrLhnsmnl19") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd876, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd865, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "26Afq0FyNk8ZKOF") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd869, []byte("84uCz6da8cQPC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) 84uCz6da8cQPC

buf, ret = client.Read(fd867, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_QUPBp1lz7gc") {
  panic("wrong data returned")
}


ret = client.Seek(fd817, 86, client.SEEK_SET)
if(ret != 86) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 86)
  panic("seek failed")
}


ret = client.Close(fd861)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd860, []byte("qSBp8HSHZ2uHyMZuiir1IwFacTxdHuhNJ3X86sRBmF0Q8_aqfY5VGjhaSn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) qSBp8HSHZ2uHyMZuiir1IwFacTxdHuhNJ3X86sRBmF0Q8_aqfY5VGjhaSn

ret = client.Seek(fd860, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


buf, ret = client.Read(fd866, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd887 := client.Open("/RUpXDwp6BU.txt", client.O_RDWR|client.O_CREATE)
if(fd887 < 0) {
  panic("open failed")
}


ret = client.Close(fd797)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) BA22MLSZ62Fw0PQ_Ta96BELBzNoR2cnwH7fPfTnts

ret = client.Write(fd873, []byte("vSkl9AnquphLya0ReyryfGCGwGNZvoj8nAfT5xMIgMD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) vSkl9AnquphLya0ReyryfGCGwGNZvoj8nAfT5xMIgMD
//fd state: (93) JwMUL9H42ItdkqJ1rvhNjcUnYvoZJEYIjk51Hh4aHwsXz_gR_lVjUjHttGfsnRv17Gfl7O084jbllnOa3cY7wbrm6DMJR

ret = client.Write(fd770, []byte("OUEpXEc3quwIHghnJy3LkEDbFQr32659Sxyc39Y93_a32IlOZffh70Fvxf7zrsgZ8d6EK41dBtKCoKGPWZEEvajwc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (182) JwMUL9H42ItdkqJ1rvhNjcUnYvoZJEYIjk51Hh4aHwsXz_gR_lVjUjHttGfsnRv17Gfl7O084jbllnOa3cY7wbrm6DMJROUEpXEc3quwIHghnJy3LkEDbFQr32659Sxyc39Y93_a32IlOZffh70Fvxf7zrsgZ8d6EK41dBtKCoKGPWZEEvajwc

buf, ret = client.Read(fd883, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd845, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd863)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd758)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd882, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd888 := client.Open("/Ejqw2f8dM3.txt", client.O_RDWR|client.O_CREATE)
if(fd888 < 0) {
  panic("open failed")
}

//fd state: (86) LnqDLBkDevzv6XwA_XVGGdcpwUBBz0QlU_mSHbe895AbpSdQFTPNzHPsIFMBpmr36_mOVFcLM7klho3RSWcIHWLD69eNP3TFD

ret = client.Write(fd817, []byte("zb3odfnsNnEcvY6kgViNRRLK8L9DicjpwLo3d68M5T_w5m5mlGDA99sBesqH69o_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) LnqDLBkDevzv6XwA_XVGGdcpwUBBz0QlU_mSHbe895AbpSdQFTPNzHPsIFMBpmr36_mOVFcLM7klho3RSWcIHWzb3odfnsNnEcvY6kgViNRRLK8L9DicjpwLo3d68M5T_w5m5mlGDA99sBesqH69o_
//fd state: (150) LnqDLBkDevzv6XwA_XVGGdcpwUBBz0QlU_mSHbe895AbpSdQFTPNzHPsIFMBpmr36_mOVFcLM7klho3RSWcIHWzb3odfnsNnEcvY6kgViNRRLK8L9DicjpwLo3d68M5T_w5m5mlGDA99sBesqH69o_

ret = client.Write(fd817, []byte("dZIBbmCZ3BH1hMGayKChGCHkHhj63b4kBzzFSOB2L8ROB6S4gVxUU67cn4aSbSIYUlny"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (218) LnqDLBkDevzv6XwA_XVGGdcpwUBBz0QlU_mSHbe895AbpSdQFTPNzHPsIFMBpmr36_mOVFcLM7klho3RSWcIHWzb3odfnsNnEcvY6kgViNRRLK8L9DicjpwLo3d68M5T_w5m5mlGDA99sBesqH69o_dZIBbmCZ3BH1hMGayKChGCHkHhj63b4kBzzFSOB2L8ROB6S4gVxUU67cn4aSbSIYUlny

fd889 := client.Open("/6Bt1QPgR6L.txt", client.O_RDWR|client.O_CREATE)
if(fd889 < 0) {
  panic("open failed")
}


ret = client.Close(fd886)
if(ret != 0) {
  panic("close failed")
}

//fd state: (67) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7gcTKjXvbKKsfmnyshq8ZZ071jq9nIO45sUhZrUK_n03KYWONWJ0GSHQbmZ5ANPo5wO5TA8WuZj6dNlE8UaiiWHj2vHWs7H3q6siDYpyt_cwimLd7ZqsOiyl0TWI

ret = client.Write(fd867, []byte("evVpNocAGjTNDrYBru8kUDqdK7VhNwK7H0r8w6GnQc07o5eCcPo150emfgz6tsEawSgh5p8WWF2idcjpA8GYaXL0_dPqr7hi2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7gcevVpNocAGjTNDrYBru8kUDqdK7VhNwK7H0r8w6GnQc07o5eCcPo150emfgz6tsEawSgh5p8WWF2idcjpA8GYaXL0_dPqr7hi2DYpyt_cwimLd7ZqsOiyl0TWI
//fd state: (0) 

ret = client.Write(fd874, []byte("IUlvjAzdk4UCXqCqa8RMlkNExJ86E2iLgaZJ4EvZYadlNxRnyMuSw63UPLcR_OyGSfLE06gkLmaHJqph2Ln"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) IUlvjAzdk4UCXqCqa8RMlkNExJ86E2iLgaZJ4EvZYadlNxRnyMuSw63UPLcR_OyGSfLE06gkLmaHJqph2Ln

ret = client.Seek(fd882, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}


buf, ret = client.Read(fd884, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "cbEyh4") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd770, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd834)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd879, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "E") {
  panic("wrong data returned")
}


fd890 := client.Open("/AofvYo70oM.txt", client.O_RDWR|client.O_CREATE)
if(fd890 < 0) {
  panic("open failed")
}

//fd state: (164) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7gcevVpNocAGjTNDrYBru8kUDqdK7VhNwK7H0r8w6GnQc07o5eCcPo150emfgz6tsEawSgh5p8WWF2idcjpA8GYaXL0_dPqr7hi2DYpyt_cwimLd7ZqsOiyl0TWI

ret = client.Write(fd867, []byte("h6Y_1kfw0KyOfo09UfhyQlJqmJlXeLCGex_0eKQRdoKD18AbCj8wwifY29SV0HQmsEr_JcD49"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (237) NFUFRSpuKCMBGLf1Bcveazsssxq5Vf97QzfQI_aDG0Od4TcAGvmb5LR_QUPBp1lz7gcevVpNocAGjTNDrYBru8kUDqdK7VhNwK7H0r8w6GnQc07o5eCcPo150emfgz6tsEawSgh5p8WWF2idcjpA8GYaXL0_dPqr7hi2h6Y_1kfw0KyOfo09UfhyQlJqmJlXeLCGex_0eKQRdoKD18AbCj8wwifY29SV0HQmsEr_JcD49

buf, ret = client.Read(fd817, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd843)
if(ret != 0) {
  panic("close failed")
}


fd891 := client.Open("/GBXTDMEc9S.txt", client.O_RDWR|client.O_CREATE)
if(fd891 < 0) {
  panic("open failed")
}


ret = client.Seek(fd770, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Seek(fd815, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd888, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Close(fd888)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd869, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd892 := client.Open("/1Si5Tpuwud.txt", client.O_RDWR|client.O_CREATE)
if(fd892 < 0) {
  panic("open failed")
}


ret = client.Close(fd868)
if(ret != 0) {
  panic("close failed")
}


fd893 := client.Open("/D7PuDI5cX1.txt", client.O_RDWR|client.O_CREATE)
if(fd893 < 0) {
  panic("open failed")
}


ret = client.Seek(fd874, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


buf, ret = client.Read(fd876, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd894 := client.Open("/2ihr4LfQFk.txt", client.O_RDWR|client.O_CREATE)
if(fd894 < 0) {
  panic("open failed")
}


fd895 := client.Open("/0OgXQ0ACr1.txt", client.O_RDWR|client.O_CREATE)
if(fd895 < 0) {
  panic("open failed")
}


ret = client.Close(fd894)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd865, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd875, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd892, []byte("HdkVUC94HFp_WxUrU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) HdkVUC94HFp_WxUrU

fd896 := client.Open("/2JE4vhl6A_.txt", client.O_RDWR|client.O_CREATE)
if(fd896 < 0) {
  panic("open failed")
}


fd897 := client.Open("/EFzqZMJJ7N.txt", client.O_RDWR|client.O_CREATE)
if(fd897 < 0) {
  panic("open failed")
}


ret = client.Close(fd879)
if(ret != 0) {
  panic("close failed")
}

//fd state: (46) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo29

ret = client.Write(fd880, []byte("9Dx_LhVl8IQVIG0OQa5OoIKm99gi_nBQ4ItuDA75C9POiWydR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo299Dx_LhVl8IQVIG0OQa5OoIKm99gi_nBQ4ItuDA75C9POiWydR
//fd state: (93) 9P8kT_aEDa_eiIAyJ4_Xn_16VxVd7UPydh1bTHDk1m_WBCSRtTosLTcpCePqppBBwovkNYOS1ebhFCOdXJrt5xw3OSrCk

ret = client.Write(fd844, []byte("QWkI1xVAI_sDgO2JdAH6GKiDglS2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (121) 9P8kT_aEDa_eiIAyJ4_Xn_16VxVd7UPydh1bTHDk1m_WBCSRtTosLTcpCePqppBBwovkNYOS1ebhFCOdXJrt5xw3OSrCkQWkI1xVAI_sDgO2JdAH6GKiDglS2
//fd state: (43) vSkl9AnquphLya0ReyryfGCGwGNZvoj8nAfT5xMIgMD

ret = client.Write(fd873, []byte("AY5IJNcb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) vSkl9AnquphLya0ReyryfGCGwGNZvoj8nAfT5xMIgMDAY5IJNcb
//fd state: (51) vSkl9AnquphLya0ReyryfGCGwGNZvoj8nAfT5xMIgMDAY5IJNcb

ret = client.Write(fd873, []byte("EARTY6e1tHrngTk01HqPw_kna8shs1EQSp4oRFeNi4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) vSkl9AnquphLya0ReyryfGCGwGNZvoj8nAfT5xMIgMDAY5IJNcbEARTY6e1tHrngTk01HqPw_kna8shs1EQSp4oRFeNi4

ret = client.Seek(fd827, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


fd898 := client.Open("/qQvWz3weHf.txt", client.O_RDWR|client.O_CREATE)
if(fd898 < 0) {
  panic("open failed")
}


ret = client.Seek(fd876, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


buf, ret = client.Read(fd862, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd873)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd898, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (41) Ln_7gYBQoUmkZzo0HX1psPsQIYn2_Wxgy_TfMT4jp1y4y0LfHwrlJblkVoQZX6F5_8XXbKWj3BOCmIdEbBGruba2yTvFBV9Ol2Bb11sWhkOKMS49j1tqSocek2fU73flEgL4QcuMSJhYC1a3

ret = client.Write(fd827, []byte("bx4H_wwGJJ5cltF7CBY5gnBUO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) Ln_7gYBQoUmkZzo0HX1psPsQIYn2_Wxgy_TfMT4jpbx4H_wwGJJ5cltF7CBY5gnBUOXXbKWj3BOCmIdEbBGruba2yTvFBV9Ol2Bb11sWhkOKMS49j1tqSocek2fU73flEgL4QcuMSJhYC1a3

ret = client.Seek(fd827, 117, client.SEEK_SET)
if(ret != 117) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 117)
  panic("seek failed")
}


buf, ret = client.Read(fd898, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (121) 9P8kT_aEDa_eiIAyJ4_Xn_16VxVd7UPydh1bTHDk1m_WBCSRtTosLTcpCePqppBBwovkNYOS1ebhFCOdXJrt5xw3OSrCkQWkI1xVAI_sDgO2JdAH6GKiDglS2

ret = client.Write(fd844, []byte("a2wWGAGEDOgCHqdyYv7EtSvcCtMSVKBehSMbY8w5LFOUXQevDPaDy15Qm6RBg9cc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (185) 9P8kT_aEDa_eiIAyJ4_Xn_16VxVd7UPydh1bTHDk1m_WBCSRtTosLTcpCePqppBBwovkNYOS1ebhFCOdXJrt5xw3OSrCkQWkI1xVAI_sDgO2JdAH6GKiDglS2a2wWGAGEDOgCHqdyYv7EtSvcCtMSVKBehSMbY8w5LFOUXQevDPaDy15Qm6RBg9cc
//fd state: (16) qSBp8HSHZ2uHyMZuiir1IwFacTxdHuhNJ3X86sRBmF0Q8_aqfY5VGjhaSn

ret = client.Write(fd860, []byte("L71P9ZwT3nLMQdSS9HaTtCe5GKxQzXj1nf23QMhsPFYrO_du"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) qSBp8HSHZ2uHyMZuL71P9ZwT3nLMQdSS9HaTtCe5GKxQzXj1nf23QMhsPFYrO_du

buf, ret = client.Read(fd827, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ocek2fU73flEgL4QcuMSJhYC1a3") {
  panic("wrong data returned")
}


fd899 := client.Open("/dkuo7HULBj.txt", client.O_RDWR|client.O_CREATE)
if(fd899 < 0) {
  panic("open failed")
}


fd900 := client.Open("/2MKMM17MVW.txt", client.O_RDWR|client.O_CREATE)
if(fd900 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd889, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd901 := client.Open("/op9fw4mr10.txt", client.O_RDWR|client.O_CREATE)
if(fd901 < 0) {
  panic("open failed")
}


ret = client.Close(fd893)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd889, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd902 := client.Open("/fHgU02aR3z.txt", client.O_RDWR|client.O_CREATE)
if(fd902 < 0) {
  panic("open failed")
}


fd903 := client.Open("/5L6AH59F8o.txt", client.O_RDWR|client.O_CREATE)
if(fd903 < 0) {
  panic("open failed")
}

//fd state: (64) qSBp8HSHZ2uHyMZuL71P9ZwT3nLMQdSS9HaTtCe5GKxQzXj1nf23QMhsPFYrO_du

ret = client.Write(fd860, []byte("oNaHsttOpzd7D4Ddw_AUMS7ATHCkzReh8ThqQ7DgSPNIiILzshtuXVMjphatv870N9b5LPCT60WO9jQhPOYuAujShYQj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (156) qSBp8HSHZ2uHyMZuL71P9ZwT3nLMQdSS9HaTtCe5GKxQzXj1nf23QMhsPFYrO_duoNaHsttOpzd7D4Ddw_AUMS7ATHCkzReh8ThqQ7DgSPNIiILzshtuXVMjphatv870N9b5LPCT60WO9jQhPOYuAujShYQj

buf, ret = client.Read(fd770, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "UL9H42Itdkq") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd844, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd883, []byte("Gv4yHyHZJ_si23CZfBltDY6BEjSfDUNpSCW1MzymlgRPDij_ZgM5JOvYU9vJ4NHjNj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) Gv4yHyHZJ_si23CZfBltDY6BEjSfDUNpSCW1MzymlgRPDij_ZgM5JOvYU9vJ4NHjNj

ret = client.Seek(fd860, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd885, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd862, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd884, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wC1mLgOHDN9e8KvvO1egqa") {
  panic("wrong data returned")
}


fd904 := client.Open("/WuYpJrAIYF.txt", client.O_RDWR|client.O_CREATE)
if(fd904 < 0) {
  panic("open failed")
}


ret = client.Seek(fd869, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd876, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xPir95twT") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd901, []byte("c1TvfiU78E43VYYdQGgLYTdpKKyv5ptlrZ2XezVLjIZm6Xh4JgqMbPbVeRnknJXzPT4J7QTHQNmn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) c1TvfiU78E43VYYdQGgLYTdpKKyv5ptlrZ2XezVLjIZm6Xh4JgqMbPbVeRnknJXzPT4J7QTHQNmn

ret = client.Close(fd876)
if(ret != 0) {
  panic("close failed")
}

//fd state: (66) XBxpOKq_Rbp_FkXaIGsIbBXsyz8COnrcdCQycEJWURKbThoEjgHUKaS4bzchiRcI4D

ret = client.Write(fd875, []byte("swtx9gkzMWGm9WvUtCVoIEl9AXNlQPcAP6AASXDMgpYS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) XBxpOKq_Rbp_FkXaIGsIbBXsyz8COnrcdCQycEJWURKbThoEjgHUKaS4bzchiRcI4Dswtx9gkzMWGm9WvUtCVoIEl9AXNlQPcAP6AASXDMgpYS
//fd state: (0) 

ret = client.Write(fd904, []byte("5L6NZn3XbSE0KnD72ra0rpz5W3ZhBE2iH_8_MJy5qZhS0XmQGFORPYRJz_LE6cGP6s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) 5L6NZn3XbSE0KnD72ra0rpz5W3ZhBE2iH_8_MJy5qZhS0XmQGFORPYRJz_LE6cGP6s

buf, ret = client.Read(fd827, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd901, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


ret = client.Seek(fd878, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd895)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd752, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd878, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd891, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


buf, ret = client.Read(fd815, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd905 := client.Open("/26dSM1oixo.txt", client.O_RDWR|client.O_CREATE)
if(fd905 < 0) {
  panic("open failed")
}


ret = client.Close(fd770)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd827)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd817)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd883)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd901, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "h4JgqMbPbVeRnknJXzPT4J7QTHQNmn") {
  panic("wrong data returned")
}


fd906 := client.Open("/RUBg7QMMiO.txt", client.O_RDWR|client.O_CREATE)
if(fd906 < 0) {
  panic("open failed")
}


ret = client.Close(fd875)
if(ret != 0) {
  panic("close failed")
}

//fd state: (78) 3Udrlv348DOd7ipdzCKL8KUeTwzmJZsX7S144SUXGgbuznyPWH0V3pHDgBgfjolcL9EzbDurNrLhnsmnl19

ret = client.Write(fd882, []byte("m4GROKk0hs1rssJe7gOzo0BwP731HcEKQhBaGEXC9iUnS9dFLDMEU2nJ46A9i6K2RbFeuWEz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) 3Udrlv348DOd7ipdzCKL8KUeTwzmJZsX7S144SUXGgbuznyPWH0V3pHDgBgfjolcL9EzbDurNrLhnsm4GROKk0hs1rssJe7gOzo0BwP731HcEKQhBaGEXC9iUnS9dFLDMEU2nJ46A9i6K2RbFeuWEz

fd907 := client.Open("/Fz4V3PLcWJ.txt", client.O_RDWR|client.O_CREATE)
if(fd907 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd865, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (95) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo299Dx_LhVl8IQVIG0OQa5OoIKm99gi_nBQ4ItuDA75C9POiWydR

ret = client.Write(fd880, []byte("q30VIEp1AWRyqjnF3GoeijH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo299Dx_LhVl8IQVIG0OQa5OoIKm99gi_nBQ4ItuDA75C9POiWydRq30VIEp1AWRyqjnF3GoeijH

ret = client.Seek(fd900, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


ret = client.Seek(fd867, 170, client.SEEK_SET)
if(ret != 170) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 170)
  panic("seek failed")
}


ret = client.Seek(fd866, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd902, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd845)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd874, 60, client.SEEK_SET)
if(ret != 60) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 60)
  panic("seek failed")
}


ret = client.Close(fd907)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd898, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd902)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd899, []byte("BAKHQERBVB5tnC52eXAP8aVpohykFlMEid8sB1zYdpMSZSNew7u2TEiY9vxoMbajUFuLRQgfXCuvy9YWUfFtpIFV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) BAKHQERBVB5tnC52eXAP8aVpohykFlMEid8sB1zYdpMSZSNew7u2TEiY9vxoMbajUFuLRQgfXCuvy9YWUfFtpIFV

buf, ret = client.Read(fd899, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd903, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (28) cbEyh4wC1mLgOHDN9e8KvvO1egqa

ret = client.Write(fd884, []byte("C_piuwXm8jNDA3bSjWHz8dKsYWFNaeqQBntP7DUzW_o8QNg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) cbEyh4wC1mLgOHDN9e8KvvO1egqaC_piuwXm8jNDA3bSjWHz8dKsYWFNaeqQBntP7DUzW_o8QNg

buf, ret = client.Read(fd901, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd897)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd896)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd878, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd908 := client.Open("/RPLGuQQ93O.txt", client.O_RDWR|client.O_CREATE)
if(fd908 < 0) {
  panic("open failed")
}


ret = client.Seek(fd865, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


fd909 := client.Open("/PfWvm8Of15.txt", client.O_RDWR|client.O_CREATE)
if(fd909 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd889, []byte("pq9qnWQuuUmmLon6l3vKRcCu2Yn0kvqkRkAmMvbAmNp2ze_xd93R4IxGqv9liSO5RxJ3Bb_T69Jd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) pq9qnWQuuUmmLon6l3vKRcCu2Yn0kvqkRkAmMvbAmNp2ze_xd93R4IxGqv9liSO5RxJ3Bb_T69Jd

buf, ret = client.Read(fd874, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_OyGSfLE06gkLmaHJqph2Ln") {
  panic("wrong data returned")
}


ret = client.Seek(fd890, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd866, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd909, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (75) cbEyh4wC1mLgOHDN9e8KvvO1egqaC_piuwXm8jNDA3bSjWHz8dKsYWFNaeqQBntP7DUzW_o8QNg

ret = client.Write(fd884, []byte("4Fj6U2cpZG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) cbEyh4wC1mLgOHDN9e8KvvO1egqaC_piuwXm8jNDA3bSjWHz8dKsYWFNaeqQBntP7DUzW_o8QNg4Fj6U2cpZG

ret = client.Seek(fd860, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


fd910 := client.Open("/cnxyeAsIbW.txt", client.O_RDWR|client.O_CREATE)
if(fd910 < 0) {
  panic("open failed")
}


fd911 := client.Open("/BdWx7NOySC.txt", client.O_RDWR|client.O_CREATE)
if(fd911 < 0) {
  panic("open failed")
}


ret = client.Seek(fd911, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd899, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (62) xA5zJP1heOoOC8GYFnsgQkXQWcHWqH6ZmTMScuckqVeT8GB2CKLDEwf1cZc_jkWp8TJgqhc_gn_bHF6XIWmrYyvWn

ret = client.Write(fd891, []byte("VrBU7lCSmiuXuPX_5p7MJg0YsNDDijts5_O"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) xA5zJP1heOoOC8GYFnsgQkXQWcHWqH6ZmTMScuckqVeT8GB2CKLDEwf1cZc_jkVrBU7lCSmiuXuPX_5p7MJg0YsNDDijts5_O

ret = client.Close(fd891)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd866, []byte("TQtC7Eiy7u7KLAssO561wgm7qjgMQVHo0ZOSFToItiw9UKkjSGqKS3_GRjUj4beWZ4RHHizZL0UhhpWHkVaD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) TQtC7Eiy7u7KLAssO561wgm7qjgMQVHo0ZOSFToItiw9UKkjSGqKS3_GRjUj4beWZ4RHHizZL0UhhpWHkVaD

fd912 := client.Open("/h4M36t8DVY.txt", client.O_RDWR|client.O_CREATE)
if(fd912 < 0) {
  panic("open failed")
}


ret = client.Close(fd900)
if(ret != 0) {
  panic("close failed")
}

//fd state: (88) BAKHQERBVB5tnC52eXAP8aVpohykFlMEid8sB1zYdpMSZSNew7u2TEiY9vxoMbajUFuLRQgfXCuvy9YWUfFtpIFV

ret = client.Write(fd899, []byte("foMuuwKP159FWv5shjseVdQD2XZ36GyBKRXKgj75iCZ2FJVaBfiMOBWZnmEOenNMdqe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (155) BAKHQERBVB5tnC52eXAP8aVpohykFlMEid8sB1zYdpMSZSNew7u2TEiY9vxoMbajUFuLRQgfXCuvy9YWUfFtpIFVfoMuuwKP159FWv5shjseVdQD2XZ36GyBKRXKgj75iCZ2FJVaBfiMOBWZnmEOenNMdqe

ret = client.Seek(fd912, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (150) 3Udrlv348DOd7ipdzCKL8KUeTwzmJZsX7S144SUXGgbuznyPWH0V3pHDgBgfjolcL9EzbDurNrLhnsm4GROKk0hs1rssJe7gOzo0BwP731HcEKQhBaGEXC9iUnS9dFLDMEU2nJ46A9i6K2RbFeuWEz

ret = client.Write(fd882, []byte("hWHXNrDmXCF4pr_ECfp1tWjVqFNH47VjuIQ6GEhu5qWaM2Gcgzi56YL0a4x_vOqdD5XE3D"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (220) 3Udrlv348DOd7ipdzCKL8KUeTwzmJZsX7S144SUXGgbuznyPWH0V3pHDgBgfjolcL9EzbDurNrLhnsm4GROKk0hs1rssJe7gOzo0BwP731HcEKQhBaGEXC9iUnS9dFLDMEU2nJ46A9i6K2RbFeuWEzhWHXNrDmXCF4pr_ECfp1tWjVqFNH47VjuIQ6GEhu5qWaM2Gcgzi56YL0a4x_vOqdD5XE3D
//fd state: (118) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo299Dx_LhVl8IQVIG0OQa5OoIKm99gi_nBQ4ItuDA75C9POiWydRq30VIEp1AWRyqjnF3GoeijH

ret = client.Write(fd880, []byte("DMMQfkEnatWs390px4vMskIcGuCvqKPD3Jaa4Qh71cIqKtiB8pUAnF_psfNcJtWqTfe6qlTYEIaLtEWtci3rLH4IO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (207) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo299Dx_LhVl8IQVIG0OQa5OoIKm99gi_nBQ4ItuDA75C9POiWydRq30VIEp1AWRyqjnF3GoeijHDMMQfkEnatWs390px4vMskIcGuCvqKPD3Jaa4Qh71cIqKtiB8pUAnF_psfNcJtWqTfe6qlTYEIaLtEWtci3rLH4IO

ret = client.Seek(fd752, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd913 := client.Open("/NDaLK8HBXg.txt", client.O_RDWR|client.O_CREATE)
if(fd913 < 0) {
  panic("open failed")
}


fd914 := client.Open("/2JE4vhl6A_.txt", client.O_RDWR|client.O_CREATE)
if(fd914 < 0) {
  panic("open failed")
}


ret = client.Seek(fd860, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}


ret = client.Close(fd860)
if(ret != 0) {
  panic("close failed")
}


fd915 := client.Open("/HRbgk5YrAz.txt", client.O_RDWR|client.O_CREATE)
if(fd915 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd874, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd903)
if(ret != 0) {
  panic("close failed")
}


fd916 := client.Open("/nBwnAtneem.txt", client.O_RDWR|client.O_CREATE)
if(fd916 < 0) {
  panic("open failed")
}


ret = client.Close(fd909)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd908, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd911, []byte("kTPpQPp1k"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) kTPpQPp1k

ret = client.Close(fd910)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd867)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd885, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd752, 60, client.SEEK_SET)
if(ret != 60) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 60)
  panic("seek failed")
}


buf, ret = client.Read(fd884, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (220) 3Udrlv348DOd7ipdzCKL8KUeTwzmJZsX7S144SUXGgbuznyPWH0V3pHDgBgfjolcL9EzbDurNrLhnsm4GROKk0hs1rssJe7gOzo0BwP731HcEKQhBaGEXC9iUnS9dFLDMEU2nJ46A9i6K2RbFeuWEzhWHXNrDmXCF4pr_ECfp1tWjVqFNH47VjuIQ6GEhu5qWaM2Gcgzi56YL0a4x_vOqdD5XE3D

ret = client.Write(fd882, []byte("bzIK9_NarfFYhzhrWq_I2_5E"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (244) 3Udrlv348DOd7ipdzCKL8KUeTwzmJZsX7S144SUXGgbuznyPWH0V3pHDgBgfjolcL9EzbDurNrLhnsm4GROKk0hs1rssJe7gOzo0BwP731HcEKQhBaGEXC9iUnS9dFLDMEU2nJ46A9i6K2RbFeuWEzhWHXNrDmXCF4pr_ECfp1tWjVqFNH47VjuIQ6GEhu5qWaM2Gcgzi56YL0a4x_vOqdD5XE3DbzIK9_NarfFYhzhrWq_I2_5E

ret = client.Seek(fd874, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


ret = client.Seek(fd892, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


buf, ret = client.Read(fd898, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd892, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd882)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) LIUbj_rdckqzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAU6Sh9H5KvARhTDGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGdxoPlKW7lbYhouaDYod8l359wnP04T1CrJ9pj2x2GFzeliIJWpQG7IBB3airdkABEq6488D

ret = client.Write(fd887, []byte("RBmjO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) RBmjO_rdckqzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAU6Sh9H5KvARhTDGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGdxoPlKW7lbYhouaDYod8l359wnP04T1CrJ9pj2x2GFzeliIJWpQG7IBB3airdkABEq6488D

buf, ret = client.Read(fd912, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd862, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd912)
if(ret != 0) {
  panic("close failed")
}


fd917 := client.Open("/kYPAZl3Z2l.txt", client.O_RDWR|client.O_CREATE)
if(fd917 < 0) {
  panic("open failed")
}


fd918 := client.Open("/C0Zs6QlJna.txt", client.O_RDWR|client.O_CREATE)
if(fd918 < 0) {
  panic("open failed")
}


fd919 := client.Open("/Fcm7ZFRdzB.txt", client.O_RDWR|client.O_CREATE)
if(fd919 < 0) {
  panic("open failed")
}


ret = client.Close(fd752)
if(ret != 0) {
  panic("close failed")
}


fd920 := client.Open("/KFH4u52ZB3.txt", client.O_RDWR|client.O_CREATE)
if(fd920 < 0) {
  panic("open failed")
}

//fd state: (0) 2VWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNu8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTyoMPacd0AoaWs947QQV7TkbBdPG0bMe6bwOtURbM8UZu_cLG5OjhIwmkICcI5CxE_xp3RGWfoK_ZQIuENtpE6xWLHbIHStPbR5q4SlHtYFxjK42PoiFf0qHKgixrSnWK

ret = client.Write(fd913, []byte("d"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) dVWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNu8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTyoMPacd0AoaWs947QQV7TkbBdPG0bMe6bwOtURbM8UZu_cLG5OjhIwmkICcI5CxE_xp3RGWfoK_ZQIuENtpE6xWLHbIHStPbR5q4SlHtYFxjK42PoiFf0qHKgixrSnWK
//fd state: (0) 

ret = client.Write(fd919, []byte("WiwBybBB8ZrUK7xYkC04n9TP35SUSYpRUDEHjYcQllWLja_PYj1x6Z_dHzIShiHoyscwEze1Oz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) WiwBybBB8ZrUK7xYkC04n9TP35SUSYpRUDEHjYcQllWLja_PYj1x6Z_dHzIShiHoyscwEze1Oz

ret = client.Close(fd914)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd905, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}

//fd state: (207) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo299Dx_LhVl8IQVIG0OQa5OoIKm99gi_nBQ4ItuDA75C9POiWydRq30VIEp1AWRyqjnF3GoeijHDMMQfkEnatWs390px4vMskIcGuCvqKPD3Jaa4Qh71cIqKtiB8pUAnF_psfNcJtWqTfe6qlTYEIaLtEWtci3rLH4IO

ret = client.Write(fd880, []byte("LlWllPuMLK8k5ncqVYZ8rgAo1PJ6_R0_Ht2Em5GyvHCOcUxq5CSSHSPjOWwPOk7QhOVn9rEOYRGB0qpzEBkj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (291) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo299Dx_LhVl8IQVIG0OQa5OoIKm99gi_nBQ4ItuDA75C9POiWydRq30VIEp1AWRyqjnF3GoeijHDMMQfkEnatWs390px4vMskIcGuCvqKPD3Jaa4Qh71cIqKtiB8pUAnF_psfNcJtWqTfe6qlTYEIaLtEWtci3rLH4IOLlWllPuMLK8k5ncqVYZ8rgAo1PJ6_R0_Ht2Em5GyvHCOcUxq5CSSHSPjOWwPOk7QhOVn9rEOYRGB0qpzEBkj
//fd state: (291) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo299Dx_LhVl8IQVIG0OQa5OoIKm99gi_nBQ4ItuDA75C9POiWydRq30VIEp1AWRyqjnF3GoeijHDMMQfkEnatWs390px4vMskIcGuCvqKPD3Jaa4Qh71cIqKtiB8pUAnF_psfNcJtWqTfe6qlTYEIaLtEWtci3rLH4IOLlWllPuMLK8k5ncqVYZ8rgAo1PJ6_R0_Ht2Em5GyvHCOcUxq5CSSHSPjOWwPOk7QhOVn9rEOYRGB0qpzEBkj

ret = client.Write(fd880, []byte("HcyhbsRg0YZnAhVjBgSUyBD0e7T0G_u8zUq6iG_iSbU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (334) wI2gM9Uvhqpu_NphqkROamDkB2KS6YnmJI6Kiq9Pwvjo299Dx_LhVl8IQVIG0OQa5OoIKm99gi_nBQ4ItuDA75C9POiWydRq30VIEp1AWRyqjnF3GoeijHDMMQfkEnatWs390px4vMskIcGuCvqKPD3Jaa4Qh71cIqKtiB8pUAnF_psfNcJtWqTfe6qlTYEIaLtEWtci3rLH4IOLlWllPuMLK8k5ncqVYZ8rgAo1PJ6_R0_Ht2Em5GyvHCOcUxq5CSSHSPjOWwPOk7QhOVn9rEOYRGB0qpzEBkjHcyhbsRg0YZnAhVjBgSUyBD0e7T0G_u8zUq6iG_iSbU

ret = client.Close(fd889)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd884, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


fd921 := client.Open("/ZYVT5hpnrT.txt", client.O_RDWR|client.O_CREATE)
if(fd921 < 0) {
  panic("open failed")
}

//fd state: (16) cbEyh4wC1mLgOHDN9e8KvvO1egqaC_piuwXm8jNDA3bSjWHz8dKsYWFNaeqQBntP7DUzW_o8QNg4Fj6U2cpZG

ret = client.Write(fd884, []byte("yBHRRGqOxXNdw1EZFbscTXhXSt62tMnCzRPylFWDCwERoWdfmB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) cbEyh4wC1mLgOHDNyBHRRGqOxXNdw1EZFbscTXhXSt62tMnCzRPylFWDCwERoWdfmBUzW_o8QNg4Fj6U2cpZG

ret = client.Close(fd908)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd815, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd866, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (5) RBmjO_rdckqzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAU6Sh9H5KvARhTDGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGdxoPlKW7lbYhouaDYod8l359wnP04T1CrJ9pj2x2GFzeliIJWpQG7IBB3airdkABEq6488D

ret = client.Write(fd887, []byte("4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) RBmjO4rdckqzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAU6Sh9H5KvARhTDGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGdxoPlKW7lbYhouaDYod8l359wnP04T1CrJ9pj2x2GFzeliIJWpQG7IBB3airdkABEq6488D

ret = client.Seek(fd920, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (185) 9P8kT_aEDa_eiIAyJ4_Xn_16VxVd7UPydh1bTHDk1m_WBCSRtTosLTcpCePqppBBwovkNYOS1ebhFCOdXJrt5xw3OSrCkQWkI1xVAI_sDgO2JdAH6GKiDglS2a2wWGAGEDOgCHqdyYv7EtSvcCtMSVKBehSMbY8w5LFOUXQevDPaDy15Qm6RBg9cc

ret = client.Write(fd844, []byte("YtjYBEWdl0cD7sbifokJGFaGUEsYxIydSqs2qmzPQTGH3VLLVgdj3cZhRd75QOyQLvZ4yNq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (256) 9P8kT_aEDa_eiIAyJ4_Xn_16VxVd7UPydh1bTHDk1m_WBCSRtTosLTcpCePqppBBwovkNYOS1ebhFCOdXJrt5xw3OSrCkQWkI1xVAI_sDgO2JdAH6GKiDglS2a2wWGAGEDOgCHqdyYv7EtSvcCtMSVKBehSMbY8w5LFOUXQevDPaDy15Qm6RBg9ccYtjYBEWdl0cD7sbifokJGFaGUEsYxIydSqs2qmzPQTGH3VLLVgdj3cZhRd75QOyQLvZ4yNq

fd922 := client.Open("/obmRgRRVEq.txt", client.O_RDWR|client.O_CREATE)
if(fd922 < 0) {
  panic("open failed")
}


ret = client.Seek(fd874, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


buf, ret = client.Read(fd898, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd878, []byte("C7npuxLr6wV4Uo_IuDk2ZTTuugZiRdDIH59Sk9vmZ8TC6knTyhHcuQ6dSCvOf0BBPyzNBTBSddWn1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) C7npuxLr6wV4Uo_IuDk2ZTTuugZiRdDIH59Sk9vmZ8TC6knTyhHcuQ6dSCvOf0BBPyzNBTBSddWn1

fd923 := client.Open("/g221aMxIGc.txt", client.O_RDWR|client.O_CREATE)
if(fd923 < 0) {
  panic("open failed")
}


ret = client.Close(fd874)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 47qbtVXel8wNdD3bB_FyXyNJvNSlb1sgt0W7l04pRgpmIahxJhRver82m9g9cP7YWXU_bxbh7Rwwjkwm

ret = client.Write(fd918, []byte("ATPotmWr6QPnfon_33p"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) ATPotmWr6QPnfon_33pyXyNJvNSlb1sgt0W7l04pRgpmIahxJhRver82m9g9cP7YWXU_bxbh7Rwwjkwm

ret = client.Close(fd906)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd898, []byte("ahF9M4GLm_jofXfwUe8pZwP9HmLITWbCOv0athuJXWeW_Bb7ZKYxihzsanobUocDlY3GlsW0yyDFUojxcjKs_UXcrSuAoFpz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) ahF9M4GLm_jofXfwUe8pZwP9HmLITWbCOv0athuJXWeW_Bb7ZKYxihzsanobUocDlY3GlsW0yyDFUojxcjKs_UXcrSuAoFpz

buf, ret = client.Read(fd866, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd892, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4HFp_WxUrU") {
  panic("wrong data returned")
}

//fd state: (76) c1TvfiU78E43VYYdQGgLYTdpKKyv5ptlrZ2XezVLjIZm6Xh4JgqMbPbVeRnknJXzPT4J7QTHQNmn

ret = client.Write(fd901, []byte("GMZCyNfFP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) c1TvfiU78E43VYYdQGgLYTdpKKyv5ptlrZ2XezVLjIZm6Xh4JgqMbPbVeRnknJXzPT4J7QTHQNmnGMZCyNfFP
//fd state: (0) 

ret = client.Write(fd862, []byte("qJaTmcZGiAL4WYqkaYUX2HS5LJS7CIfU5g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) qJaTmcZGiAL4WYqkaYUX2HS5LJS7CIfU5g

buf, ret = client.Read(fd892, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd917, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


buf, ret = client.Read(fd905, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "np1IKPMJLzW_2BpCGR9q") {
  panic("wrong data returned")
}

//fd state: (6) RBmjO4rdckqzKhguHWfmy5DL9jbqUeqk22IPQFthhODpsfzMKOq0j7Lb4cRmYCrRP6GonExSJAU6Sh9H5KvARhTDGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGdxoPlKW7lbYhouaDYod8l359wnP04T1CrJ9pj2x2GFzeliIJWpQG7IBB3airdkABEq6488D

ret = client.Write(fd887, []byte("eHsy9DSrCd9Rv9_ufL5vYCrdoOlcKp0ZuW8owr89t4_5HNhUPVGqup6CWvIoYywRKVnGXPlSdIadzXwEQM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) RBmjO4eHsy9DSrCd9Rv9_ufL5vYCrdoOlcKp0ZuW8owr89t4_5HNhUPVGqup6CWvIoYywRKVnGXPlSdIadzXwEQMGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGdxoPlKW7lbYhouaDYod8l359wnP04T1CrJ9pj2x2GFzeliIJWpQG7IBB3airdkABEq6488D

ret = client.Seek(fd921, 118, client.SEEK_SET)
if(ret != 118) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 118)
  panic("seek failed")
}


buf, ret = client.Read(fd865, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "r72UUY_K4WLy9NNSR6") {
  panic("wrong data returned")
}


ret = client.Close(fd844)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd911, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd878, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd920, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd924 := client.Open("/Dg7qYVQIe7.txt", client.O_RDWR|client.O_CREATE)
if(fd924 < 0) {
  panic("open failed")
}


ret = client.Seek(fd923, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd923, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd911, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd925 := client.Open("/l0r677EYZx.txt", client.O_RDWR|client.O_CREATE)
if(fd925 < 0) {
  panic("open failed")
}


ret = client.Close(fd916)
if(ret != 0) {
  panic("close failed")
}


fd926 := client.Open("/_Arj1HylHr.txt", client.O_RDWR|client.O_CREATE)
if(fd926 < 0) {
  panic("open failed")
}


ret = client.Seek(fd884, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


fd927 := client.Open("/ZMyenLW9Zc.txt", client.O_RDWR|client.O_CREATE)
if(fd927 < 0) {
  panic("open failed")
}


fd928 := client.Open("/7Yc4dPLSSJ.txt", client.O_RDWR|client.O_CREATE)
if(fd928 < 0) {
  panic("open failed")
}


ret = client.Seek(fd917, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


fd929 := client.Open("/etbCfgFtZR.txt", client.O_RDWR|client.O_CREATE)
if(fd929 < 0) {
  panic("open failed")
}

//fd state: (81) cbEyh4wC1mLgOHDNyBHRRGqOxXNdw1EZFbscTXhXSt62tMnCzRPylFWDCwERoWdfmBUzW_o8QNg4Fj6U2cpZG

ret = client.Write(fd884, []byte("0NlCQKBp4YzYHDYTIgj8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) cbEyh4wC1mLgOHDNyBHRRGqOxXNdw1EZFbscTXhXSt62tMnCzRPylFWDCwERoWdfmBUzW_o8QNg4Fj6U20NlCQKBp4YzYHDYTIgj8

ret = client.Seek(fd866, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Close(fd922)
if(ret != 0) {
  panic("close failed")
}

//fd state: (19) ATPotmWr6QPnfon_33pyXyNJvNSlb1sgt0W7l04pRgpmIahxJhRver82m9g9cP7YWXU_bxbh7Rwwjkwm

ret = client.Write(fd918, []byte("aAYJBtvaAaa89qgZC4pL2o75ILGHUKscKyXmOl5Z44npjlEpPJNX4L9oU1m7MfxMsg3MNFb20ye13"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) ATPotmWr6QPnfon_33paAYJBtvaAaa89qgZC4pL2o75ILGHUKscKyXmOl5Z44npjlEpPJNX4L9oU1m7MfxMsg3MNFb20ye13

fd930 := client.Open("/D5iZgR6iAY.txt", client.O_RDWR|client.O_CREATE)
if(fd930 < 0) {
  panic("open failed")
}


ret = client.Close(fd880)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd815, []byte("niO4SJNC3PHAxGXxO8i3rOEcD0sBwEJ4vHy0cd9S6N0vKnzTh_EwJJkGHWEUMNTZcEh9m8TxaicmVbjVT88Tk1Mdkokm6Uy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) niO4SJNC3PHAxGXxO8i3rOEcD0sBwEJ4vHy0cd9S6N0vKnzTh_EwJJkGHWEUMNTZcEh9m8TxaicmVbjVT88Tk1Mdkokm6Uy

fd931 := client.Open("/RI0qit6pdn.txt", client.O_RDWR|client.O_CREATE)
if(fd931 < 0) {
  panic("open failed")
}


ret = client.Close(fd899)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd911, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


ret = client.Close(fd901)
if(ret != 0) {
  panic("close failed")
}


fd932 := client.Open("/2ihr4LfQFk.txt", client.O_RDWR|client.O_CREATE)
if(fd932 < 0) {
  panic("open failed")
}


ret = client.Seek(fd913, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Close(fd927)
if(ret != 0) {
  panic("close failed")
}


fd933 := client.Open("/TWDYDmh6Fj.txt", client.O_RDWR|client.O_CREATE)
if(fd933 < 0) {
  panic("open failed")
}


ret = client.Close(fd928)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd905)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd911, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "p1k") {
  panic("wrong data returned")
}


ret = client.Seek(fd892, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


buf, ret = client.Read(fd904, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd924)
if(ret != 0) {
  panic("close failed")
}


fd934 := client.Open("/_laoVJpAG9.txt", client.O_RDWR|client.O_CREATE)
if(fd934 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd865, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9x7uV0i9H70XgydubhCXYP26Afq0FyNk") {
  panic("wrong data returned")
}


ret = client.Seek(fd884, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


buf, ret = client.Read(fd904, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (13) dVWayLo4ueCP34btW8LMSoX614J8AaMWOSbW9Iuf2BYyw9ocu0A6QNu8wgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTyoMPacd0AoaWs947QQV7TkbBdPG0bMe6bwOtURbM8UZu_cLG5OjhIwmkICcI5CxE_xp3RGWfoK_ZQIuENtpE6xWLHbIHStPbR5q4SlHtYFxjK42PoiFf0qHKgixrSnWK

ret = client.Write(fd913, []byte("TztJtRGUssNVRyAsnmZypr0PBtXhYsGfX901Xovphb7c"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) dVWayLo4ueCP3TztJtRGUssNVRyAsnmZypr0PBtXhYsGfX901Xovphb7cgVShWIQHxde15AmumPd4V3xo7bICZVOq3k1HZtwBabZkHrOYP4jx8_aIJ2FRvfprTyoMPacd0AoaWs947QQV7TkbBdPG0bMe6bwOtURbM8UZu_cLG5OjhIwmkICcI5CxE_xp3RGWfoK_ZQIuENtpE6xWLHbIHStPbR5q4SlHtYFxjK42PoiFf0qHKgixrSnWK

buf, ret = client.Read(fd862, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd935 := client.Open("/TqUxY9FOiQ.txt", client.O_RDWR|client.O_CREATE)
if(fd935 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd925, []byte("OzczRH2X156dWw1RmXbYOBBXB33vIxUih8n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) OzczRH2X156dWw1RmXbYOBBXB33vIxUih8n
//fd state: (66) 5L6NZn3XbSE0KnD72ra0rpz5W3ZhBE2iH_8_MJy5qZhS0XmQGFORPYRJz_LE6cGP6s

ret = client.Write(fd904, []byte("bxbOOcHxIDLAbus_uMeVUTOWhkICegMhMV6cF6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) 5L6NZn3XbSE0KnD72ra0rpz5W3ZhBE2iH_8_MJy5qZhS0XmQGFORPYRJz_LE6cGP6sbxbOOcHxIDLAbus_uMeVUTOWhkICegMhMV6cF6

ret = client.Close(fd915)
if(ret != 0) {
  panic("close failed")
}

//fd state: (77) C7npuxLr6wV4Uo_IuDk2ZTTuugZiRdDIH59Sk9vmZ8TC6knTyhHcuQ6dSCvOf0BBPyzNBTBSddWn1

ret = client.Write(fd878, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) C7npuxLr6wV4Uo_IuDk2ZTTuugZiRdDIH59Sk9vmZ8TC6knTyhHcuQ6dSCvOf0BBPyzNBTBSddWn1
//fd state: (74) WiwBybBB8ZrUK7xYkC04n9TP35SUSYpRUDEHjYcQllWLja_PYj1x6Z_dHzIShiHoyscwEze1Oz

ret = client.Write(fd919, []byte("Srg3pXjO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) WiwBybBB8ZrUK7xYkC04n9TP35SUSYpRUDEHjYcQllWLja_PYj1x6Z_dHzIShiHoyscwEze1OzSrg3pXjO

ret = client.Seek(fd931, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (96) ahF9M4GLm_jofXfwUe8pZwP9HmLITWbCOv0athuJXWeW_Bb7ZKYxihzsanobUocDlY3GlsW0yyDFUojxcjKs_UXcrSuAoFpz

ret = client.Write(fd898, []byte("aXiqSgsu4tUlGmueskNDbm8CmVcijpA1IxfpqmyfO06mZ2hje_yJgRKKcUax56bQv1vLnCueRf7ujBNse4L"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (179) ahF9M4GLm_jofXfwUe8pZwP9HmLITWbCOv0athuJXWeW_Bb7ZKYxihzsanobUocDlY3GlsW0yyDFUojxcjKs_UXcrSuAoFpzaXiqSgsu4tUlGmueskNDbm8CmVcijpA1IxfpqmyfO06mZ2hje_yJgRKKcUax56bQv1vLnCueRf7ujBNse4L

fd936 := client.Open("/ZyM_u_ivJT.txt", client.O_RDWR|client.O_CREATE)
if(fd936 < 0) {
  panic("open failed")
}

//fd state: (35) OzczRH2X156dWw1RmXbYOBBXB33vIxUih8n

ret = client.Write(fd925, []byte("vJAc2QSo8dKarr4rfhfv86UpxPbmpuwU7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) OzczRH2X156dWw1RmXbYOBBXB33vIxUih8nvJAc2QSo8dKarr4rfhfv86UpxPbmpuwU7

ret = client.Seek(fd935, 136, client.SEEK_SET)
if(ret != 136) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 136)
  panic("seek failed")
}


ret = client.Seek(fd878, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


fd937 := client.Open("/LY67aXtA3E.txt", client.O_RDWR|client.O_CREATE)
if(fd937 < 0) {
  panic("open failed")
}


ret = client.Seek(fd815, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


ret = client.Seek(fd865, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Close(fd937)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd920, []byte("Yo22bjdvPGRI1Qdf5PW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) Yo22bjdvPGRI1Qdf5PW
//fd state: (0) 

ret = client.Write(fd923, []byte("rrUZ8NjA1ojJYnZCBSBD_YrOhw5lczgXO5NXjHR8yWaRgfNTdM4FzC_o2sQk7SGwJMz2PuF1cI6eGr1JnIP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) rrUZ8NjA1ojJYnZCBSBD_YrOhw5lczgXO5NXjHR8yWaRgfNTdM4FzC_o2sQk7SGwJMz2PuF1cI6eGr1JnIP

ret = client.Close(fd930)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd926)
if(ret != 0) {
  panic("close failed")
}


fd938 := client.Open("/g_5w2E_4_7.txt", client.O_RDWR|client.O_CREATE)
if(fd938 < 0) {
  panic("open failed")
}

//fd state: (83) rrUZ8NjA1ojJYnZCBSBD_YrOhw5lczgXO5NXjHR8yWaRgfNTdM4FzC_o2sQk7SGwJMz2PuF1cI6eGr1JnIP

ret = client.Write(fd923, []byte("W4xsfyajTIHQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) rrUZ8NjA1ojJYnZCBSBD_YrOhw5lczgXO5NXjHR8yWaRgfNTdM4FzC_o2sQk7SGwJMz2PuF1cI6eGr1JnIPW4xsfyajTIHQ

fd939 := client.Open("/elnv2rKWiI.txt", client.O_RDWR|client.O_CREATE)
if(fd939 < 0) {
  panic("open failed")
}


fd940 := client.Open("/nBwnAtneem.txt", client.O_RDWR|client.O_CREATE)
if(fd940 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd862, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd931)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd865)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd911, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Seek(fd911, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


buf, ret = client.Read(fd917, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "0vHcbvtavDyYCTEq5QVEOBZL5QFB4HhPYAm9X730ii") {
  panic("wrong data returned")
}

//fd state: (96) ATPotmWr6QPnfon_33paAYJBtvaAaa89qgZC4pL2o75ILGHUKscKyXmOl5Z44npjlEpPJNX4L9oU1m7MfxMsg3MNFb20ye13

ret = client.Write(fd918, []byte("BKVuxfOz9pA5LmFltjN8P"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) ATPotmWr6QPnfon_33paAYJBtvaAaa89qgZC4pL2o75ILGHUKscKyXmOl5Z44npjlEpPJNX4L9oU1m7MfxMsg3MNFb20ye13BKVuxfOz9pA5LmFltjN8P

fd941 := client.Open("/L9AVf81VnB.txt", client.O_RDWR|client.O_CREATE)
if(fd941 < 0) {
  panic("open failed")
}


ret = client.Close(fd939)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd917, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd941, []byte("J17WFyBam"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) J17WFyBam
//fd state: (34) qJaTmcZGiAL4WYqkaYUX2HS5LJS7CIfU5g

ret = client.Write(fd862, []byte("NaV1OFXK9oDekRnXoqIJPP66o6YnAtwIxKK9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) qJaTmcZGiAL4WYqkaYUX2HS5LJS7CIfU5gNaV1OFXK9oDekRnXoqIJPP66o6YnAtwIxKK9
//fd state: (0) cM24y1j_uq88pmzWKopeif5aYtd8Ex1oQRpegSjsFB8tVNxJKLnV7gMzdkmZAdxffDfxnDQZjLqkaYKIDsk1d1HWtOZ4u8zpEV8kdW3mtk_Jhl9Nys2dXWCh3XxWcRuNMeZiUQe_f8Q7FIj0i6EUKpg2KiR_wtmFwWGMnBlUVJKgLLt96v8dvuisXRh5CAB1aWhpCcIVjBRktgyXXYkJITRAnWM4SJh6b2ycFxlk1

ret = client.Write(fd938, []byte("wX97a51kishSSK9PRwBii1bOqb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) wX97a51kishSSK9PRwBii1bOqbd8Ex1oQRpegSjsFB8tVNxJKLnV7gMzdkmZAdxffDfxnDQZjLqkaYKIDsk1d1HWtOZ4u8zpEV8kdW3mtk_Jhl9Nys2dXWCh3XxWcRuNMeZiUQe_f8Q7FIj0i6EUKpg2KiR_wtmFwWGMnBlUVJKgLLt96v8dvuisXRh5CAB1aWhpCcIVjBRktgyXXYkJITRAnWM4SJh6b2ycFxlk1
//fd state: (70) qJaTmcZGiAL4WYqkaYUX2HS5LJS7CIfU5gNaV1OFXK9oDekRnXoqIJPP66o6YnAtwIxKK9

ret = client.Write(fd862, []byte("u0fRVSaFGRckAdMhiTULPYxoMlxglnVIJTyWF4agjXCz9UMyg1K0XhrrjMKwn_lJEaQ0Q7XAv0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) qJaTmcZGiAL4WYqkaYUX2HS5LJS7CIfU5gNaV1OFXK9oDekRnXoqIJPP66o6YnAtwIxKK9u0fRVSaFGRckAdMhiTULPYxoMlxglnVIJTyWF4agjXCz9UMyg1K0XhrrjMKwn_lJEaQ0Q7XAv0

buf, ret = client.Read(fd892, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "U") {
  panic("wrong data returned")
}


fd942 := client.Open("/EqK6dUWAsY.txt", client.O_RDWR|client.O_CREATE)
if(fd942 < 0) {
  panic("open failed")
}


fd943 := client.Open("/6Bt1QPgR6L.txt", client.O_RDWR|client.O_CREATE)
if(fd943 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd923, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (104) 5L6NZn3XbSE0KnD72ra0rpz5W3ZhBE2iH_8_MJy5qZhS0XmQGFORPYRJz_LE6cGP6sbxbOOcHxIDLAbus_uMeVUTOWhkICegMhMV6cF6

ret = client.Write(fd904, []byte("kfDlVrj7VZeLYYCku9_bBIG_aFlxNji0AgQFxdb8IUNeHplEUdfOE1VLLhTsCt0TzD2D65ngNoc7zPOTWZZ3o4yFibJUWwR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (199) 5L6NZn3XbSE0KnD72ra0rpz5W3ZhBE2iH_8_MJy5qZhS0XmQGFORPYRJz_LE6cGP6sbxbOOcHxIDLAbus_uMeVUTOWhkICegMhMV6cF6kfDlVrj7VZeLYYCku9_bBIG_aFlxNji0AgQFxdb8IUNeHplEUdfOE1VLLhTsCt0TzD2D65ngNoc7zPOTWZZ3o4yFibJUWwR

fd944 := client.Open("/aFMbLXvGS0.txt", client.O_RDWR|client.O_CREATE)
if(fd944 < 0) {
  panic("open failed")
}


fd945 := client.Open("/JrTe3FzZzG.txt", client.O_RDWR|client.O_CREATE)
if(fd945 < 0) {
  panic("open failed")
}


fd946 := client.Open("/KB64yf4rA8.txt", client.O_RDWR|client.O_CREATE)
if(fd946 < 0) {
  panic("open failed")
}


fd947 := client.Open("/ojkJUB3zCr.txt", client.O_RDWR|client.O_CREATE)
if(fd947 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd925, 46)
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

//fd state: (179) ahF9M4GLm_jofXfwUe8pZwP9HmLITWbCOv0athuJXWeW_Bb7ZKYxihzsanobUocDlY3GlsW0yyDFUojxcjKs_UXcrSuAoFpzaXiqSgsu4tUlGmueskNDbm8CmVcijpA1IxfpqmyfO06mZ2hje_yJgRKKcUax56bQv1vLnCueRf7ujBNse4L

ret = client.Write(fd898, []byte("0lc97K4YHYxqgsi5b9jpWXtF1mwqRe7cvaj3yvrKljB82rvfChIzxix_yyIh91MXZxv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (246) ahF9M4GLm_jofXfwUe8pZwP9HmLITWbCOv0athuJXWeW_Bb7ZKYxihzsanobUocDlY3GlsW0yyDFUojxcjKs_UXcrSuAoFpzaXiqSgsu4tUlGmueskNDbm8CmVcijpA1IxfpqmyfO06mZ2hje_yJgRKKcUax56bQv1vLnCueRf7ujBNse4L0lc97K4YHYxqgsi5b9jpWXtF1mwqRe7cvaj3yvrKljB82rvfChIzxix_yyIh91MXZxv

fd948 := client.Open("/3OuuWd5hLc.txt", client.O_RDWR|client.O_CREATE)
if(fd948 < 0) {
  panic("open failed")
}


fd949 := client.Open("/pJ5WYpzKxT.txt", client.O_RDWR|client.O_CREATE)
if(fd949 < 0) {
  panic("open failed")
}


ret = client.Seek(fd949, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd950 := client.Open("/h4M36t8DVY.txt", client.O_RDWR|client.O_CREATE)
if(fd950 < 0) {
  panic("open failed")
}


fd951 := client.Open("/dKU0bYhU2p.txt", client.O_RDWR|client.O_CREATE)
if(fd951 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd890, []byte("bP6ePuTYVC058MnoYboOHF9AdpCqLk8xxr5URIY6Ly27EmB5cpICXTj9X_B0VOjl048fkc6O7W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) bP6ePuTYVC058MnoYboOHF9AdpCqLk8xxr5URIY6Ly27EmB5cpICXTj9X_B0VOjl048fkc6O7W

ret = client.Close(fd929)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd920, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}

//fd state: (17) HdkVUC94HFp_WxUrU

ret = client.Write(fd892, []byte("UyVUJJsfrfLn6t3_LZlpRHCSfaQlXCkVL9P4BMYENZThAxt5o0nAgB39YX5mYw3bJj_sMyUx2mX6QTvHuqen5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) HdkVUC94HFp_WxUrUUyVUJJsfrfLn6t3_LZlpRHCSfaQlXCkVL9P4BMYENZThAxt5o0nAgB39YX5mYw3bJj_sMyUx2mX6QTvHuqen5

ret = client.Seek(fd942, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd944, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd866)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) c3Vh2ROhloohl_gVLECJil1Lz9RkP7nTjlt8wY8IpDZ5P9A5XkYywOQ5zE3SQjX3Y4Sfsp5Ogj1qDxWNL

ret = client.Write(fd940, []byte("yzRAhrfjaxhbwYiwi48NrpoCC3gDQN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) yzRAhrfjaxhbwYiwi48NrpoCC3gDQNnTjlt8wY8IpDZ5P9A5XkYywOQ5zE3SQjX3Y4Sfsp5Ogj1qDxWNL

ret = client.Close(fd936)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd947, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd932, []byte("qogceTTaIwJbY2PGee8fZbUKhvGeqkGX8VTT8OJ0Mb0VQt5c"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) qogceTTaIwJbY2PGee8fZbUKhvGeqkGX8VTT8OJ0Mb0VQt5c

fd952 := client.Open("/I9MZl6YAeT.txt", client.O_RDWR|client.O_CREATE)
if(fd952 < 0) {
  panic("open failed")
}


ret = client.Seek(fd941, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


ret = client.Seek(fd934, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


ret = client.Close(fd940)
if(ret != 0) {
  panic("close failed")
}

//fd state: (48) qogceTTaIwJbY2PGee8fZbUKhvGeqkGX8VTT8OJ0Mb0VQt5c

ret = client.Write(fd932, []byte("xcgrNKLN_K6q__dHhR3hmiM1JSIga9HLjxYTWwS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) qogceTTaIwJbY2PGee8fZbUKhvGeqkGX8VTT8OJ0Mb0VQt5cxcgrNKLN_K6q__dHhR3hmiM1JSIga9HLjxYTWwS

ret = client.Seek(fd884, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


fd953 := client.Open("/49yoWYxM6B.txt", client.O_RDWR|client.O_CREATE)
if(fd953 < 0) {
  panic("open failed")
}


fd954 := client.Open("/uITTW59r_v.txt", client.O_RDWR|client.O_CREATE)
if(fd954 < 0) {
  panic("open failed")
}


ret = client.Close(fd946)
if(ret != 0) {
  panic("close failed")
}


fd955 := client.Open("/Ra3yzForYc.txt", client.O_RDWR|client.O_CREATE)
if(fd955 < 0) {
  panic("open failed")
}

//fd state: (0) pq9qnWQuuUmmLon6l3vKRcCu2Yn0kvqkRkAmMvbAmNp2ze_xd93R4IxGqv9liSO5RxJ3Bb_T69Jd

ret = client.Write(fd943, []byte("M2dhmpvBQ4IlkLbY7FLBVOBh3rr82kAK1OfYqsKkSZeftn1KGOOjUMXWTvhTxOF43xhE2InTArqwI8oxAxWvxEpVbbjHq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) M2dhmpvBQ4IlkLbY7FLBVOBh3rr82kAK1OfYqsKkSZeftn1KGOOjUMXWTvhTxOF43xhE2InTArqwI8oxAxWvxEpVbbjHq

fd956 := client.Open("/uJt5Xeux53.txt", client.O_RDWR|client.O_CREATE)
if(fd956 < 0) {
  panic("open failed")
}


ret = client.Seek(fd945, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd957 := client.Open("/EFzqZMJJ7N.txt", client.O_RDWR|client.O_CREATE)
if(fd957 < 0) {
  panic("open failed")
}

//fd state: (0) 3Udrlv348DOd7ipdzCKL8KUeTwzmJZsX7S144SUXGgbuznyPWH0V3pHDgBgfjolcL9EzbDurNrLhnsm4GROKk0hs1rssJe7gOzo0BwP731HcEKQhBaGEXC9iUnS9dFLDMEU2nJ46A9i6K2RbFeuWEzhWHXNrDmXCF4pr_ECfp1tWjVqFNH47VjuIQ6GEhu5qWaM2Gcgzi56YL0a4x_vOqdD5XE3DbzIK9_NarfFYhzhrWq_I2_5E

ret = client.Write(fd956, []byte("z99sRxwnF7kfYHB0w1xaho7drojFikDEac_JWUYGWPYrVoVv2z2HWFE_Nnhrt62L50Cvn8JDqtY4ue4QTJTK5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) z99sRxwnF7kfYHB0w1xaho7drojFikDEac_JWUYGWPYrVoVv2z2HWFE_Nnhrt62L50Cvn8JDqtY4ue4QTJTK50hs1rssJe7gOzo0BwP731HcEKQhBaGEXC9iUnS9dFLDMEU2nJ46A9i6K2RbFeuWEzhWHXNrDmXCF4pr_ECfp1tWjVqFNH47VjuIQ6GEhu5qWaM2Gcgzi56YL0a4x_vOqdD5XE3DbzIK9_NarfFYhzhrWq_I2_5E

ret = client.Close(fd892)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd953, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jBSy60Vuotzt") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd918, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd923)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd948, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd958 := client.Open("/IAWaQyZGxP.txt", client.O_RDWR|client.O_CREATE)
if(fd958 < 0) {
  panic("open failed")
}


ret = client.Close(fd948)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd925, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd917, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}

//fd state: (87) qogceTTaIwJbY2PGee8fZbUKhvGeqkGX8VTT8OJ0Mb0VQt5cxcgrNKLN_K6q__dHhR3hmiM1JSIga9HLjxYTWwS

ret = client.Write(fd932, []byte("MjoIjVXlH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) qogceTTaIwJbY2PGee8fZbUKhvGeqkGX8VTT8OJ0Mb0VQt5cxcgrNKLN_K6q__dHhR3hmiM1JSIga9HLjxYTWwSMjoIjVXlH

buf, ret = client.Read(fd952, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd938)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd913)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd951, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd959 := client.Open("/aODta09XGB.txt", client.O_RDWR|client.O_CREATE)
if(fd959 < 0) {
  panic("open failed")
}


ret = client.Seek(fd953, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}


buf, ret = client.Read(fd943, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd956, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Seek(fd951, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd960 := client.Open("/buR7_Zg0Os.txt", client.O_RDWR|client.O_CREATE)
if(fd960 < 0) {
  panic("open failed")
}


ret = client.Close(fd898)
if(ret != 0) {
  panic("close failed")
}

//fd state: (136) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwy19eKomT5qenYg44UklVfuvkpDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnHTXrKZ36LLaUGzNLKGp1U5g_uO7XZgvvmeygrOT7dd85KwS67No5xqQeuRvWR6N_b2ddSYUuoogejXsLdKclouI4Dz2acja

ret = client.Write(fd935, []byte("N4DRatcTGzwVrf3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) 6maJpjgkZmquT2zM2_t_aKJRAo4yEIDy0EwA2Ffa47C9bAZlrW1Z3Py8wWv_CaK1lJT5zToTKbQGQJzK1PpNIKiMzMle28efBUmIJ0Q_ipl4X9lzDoLblw4pdgxk6yapJ7VWslwyN4DRatcTGzwVrf3UklVfuvkpDy9fhuAEl8uhQUKcwLpsNAjAjqeCzq1AQFK8nABK_qnHTXrKZ36LLaUGzNLKGp1U5g_uO7XZgvvmeygrOT7dd85KwS67No5xqQeuRvWR6N_b2ddSYUuoogejXsLdKclouI4Dz2acja
//fd state: (20) cbEyh4wC1mLgOHDNyBHRRGqOxXNdw1EZFbscTXhXSt62tMnCzRPylFWDCwERoWdfmBUzW_o8QNg4Fj6U20NlCQKBp4YzYHDYTIgj8

ret = client.Write(fd884, []byte("bFh7WWCPpXNAn7E1Ao6SPQ4EYksE0g1uPAMSoC79xFntwvJ_gwwKVrnuvKv8mxfkPGQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) cbEyh4wC1mLgOHDNyBHRbFh7WWCPpXNAn7E1Ao6SPQ4EYksE0g1uPAMSoC79xFntwvJ_gwwKVrnuvKv8mxfkPGQBp4YzYHDYTIgj8

buf, ret = client.Read(fd954, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd921, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "B366g") {
  panic("wrong data returned")
}


ret = client.Seek(fd959, 89, client.SEEK_SET)
if(ret != 89) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 89)
  panic("seek failed")
}


ret = client.Seek(fd885, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd957, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (68) OzczRH2X156dWw1RmXbYOBBXB33vIxUih8nvJAc2QSo8dKarr4rfhfv86UpxPbmpuwU7

ret = client.Write(fd925, []byte("apNWI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) OzczRH2X156dWw1RmXbYOBBXB33vIxUih8nvJAc2QSo8dKarr4rfhfv86UpxPbmpuwU7apNWI

ret = client.Close(fd917)
if(ret != 0) {
  panic("close failed")
}

//fd state: (199) 5L6NZn3XbSE0KnD72ra0rpz5W3ZhBE2iH_8_MJy5qZhS0XmQGFORPYRJz_LE6cGP6sbxbOOcHxIDLAbus_uMeVUTOWhkICegMhMV6cF6kfDlVrj7VZeLYYCku9_bBIG_aFlxNji0AgQFxdb8IUNeHplEUdfOE1VLLhTsCt0TzD2D65ngNoc7zPOTWZZ3o4yFibJUWwR

ret = client.Write(fd904, []byte("7w27IQvHMxVlfnT5XsqijqpnsI3yrXUroktxUrQHToaDOxBFDtYyvPqBzjprGLnpJIT5TN7usKo0PD6A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (279) 5L6NZn3XbSE0KnD72ra0rpz5W3ZhBE2iH_8_MJy5qZhS0XmQGFORPYRJz_LE6cGP6sbxbOOcHxIDLAbus_uMeVUTOWhkICegMhMV6cF6kfDlVrj7VZeLYYCku9_bBIG_aFlxNji0AgQFxdb8IUNeHplEUdfOE1VLLhTsCt0TzD2D65ngNoc7zPOTWZZ3o4yFibJUWwR7w27IQvHMxVlfnT5XsqijqpnsI3yrXUroktxUrQHToaDOxBFDtYyvPqBzjprGLnpJIT5TN7usKo0PD6A

ret = client.Seek(fd925, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}


buf, ret = client.Read(fd947, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (4) J17WFyBam

ret = client.Write(fd941, []byte("B_yZVigWdpoGQzMeXLc8lreVu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) J17WB_yZVigWdpoGQzMeXLc8lreVu

ret = client.Close(fd921)
if(ret != 0) {
  panic("close failed")
}


fd961 := client.Open("/2yZkRq9LHH.txt", client.O_RDWR|client.O_CREATE)
if(fd961 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd942, []byte("txcVka3Gmj9Mbt_cvhaZgssnCkuf1MXfy7CrTwvmegd3e6PtGcdJC1z_LsxcvF_3cQTo393xdFNM1RGxKtiZ4coW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) txcVka3Gmj9Mbt_cvhaZgssnCkuf1MXfy7CrTwvmegd3e6PtGcdJC1z_LsxcvF_3cQTo393xdFNM1RGxKtiZ4coW

buf, ret = client.Read(fd890, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd949, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd925, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


ret = client.Seek(fd959, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


ret = client.Close(fd953)
if(ret != 0) {
  panic("close failed")
}

//fd state: (96) qogceTTaIwJbY2PGee8fZbUKhvGeqkGX8VTT8OJ0Mb0VQt5cxcgrNKLN_K6q__dHhR3hmiM1JSIga9HLjxYTWwSMjoIjVXlH

ret = client.Write(fd932, []byte("XRAhyEslMHTE4Uc7cNbJZZpLEc1y5EXPVmURYr0SBcB3FHeOJpwvOJ2OxHg5d2f_uqCOafH60OZwScuTG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (177) qogceTTaIwJbY2PGee8fZbUKhvGeqkGX8VTT8OJ0Mb0VQt5cxcgrNKLN_K6q__dHhR3hmiM1JSIga9HLjxYTWwSMjoIjVXlHXRAhyEslMHTE4Uc7cNbJZZpLEc1y5EXPVmURYr0SBcB3FHeOJpwvOJ2OxHg5d2f_uqCOafH60OZwScuTG

ret = client.Close(fd862)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd918, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}


ret = client.Seek(fd887, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


buf, ret = client.Read(fd961, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AS9HXj7dFx0mo7CBrPs") {
  panic("wrong data returned")
}

//fd state: (0) cZFDTi

ret = client.Write(fd960, []byte("pNefSqjGE7e1gnSKHLNbA2dLf7rEMbN78N9vwPGcL5TxmbLiV41JKZhp87SpzrL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) pNefSqjGE7e1gnSKHLNbA2dLf7rEMbN78N9vwPGcL5TxmbLiV41JKZhp87SpzrL

buf, ret = client.Read(fd960, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd945, []byte("AR1zBDL_b6KGYmyhjZOGLUipfuAQsLzYeOiLS18mkeVaRfZBvKfnk6rT03Ktgms1aGwX7jWwXEpT_mKWlfKPXtCBGbpITZaSlVV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) AR1zBDL_b6KGYmyhjZOGLUipfuAQsLzYeOiLS18mkeVaRfZBvKfnk6rT03Ktgms1aGwX7jWwXEpT_mKWlfKPXtCBGbpITZaSlVV

fd962 := client.Open("/HCfdrAN7le.txt", client.O_RDWR|client.O_CREATE)
if(fd962 < 0) {
  panic("open failed")
}


fd963 := client.Open("/ydgmnD5zW0.txt", client.O_RDWR|client.O_CREATE)
if(fd963 < 0) {
  panic("open failed")
}


ret = client.Close(fd919)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd885)
if(ret != 0) {
  panic("close failed")
}


fd964 := client.Open("/4B_70qkxe9.txt", client.O_RDWR|client.O_CREATE)
if(fd964 < 0) {
  panic("open failed")
}


ret = client.Seek(fd887, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


ret = client.Close(fd935)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd954)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd955, []byte("wfkrnCP_8AFc2o_OC40VQplkWsTQN26Cj8uGN86XgsflPj4jhG5ufeFRon"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) wfkrnCP_8AFc2o_OC40VQplkWsTQN26Cj8uGN86XgsflPj4jhG5ufeFRon

ret = client.Seek(fd890, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}


ret = client.Seek(fd962, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


ret = client.Close(fd963)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd933)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd950, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd934, 186, client.SEEK_SET)
if(ret != 186) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 186)
  panic("seek failed")
}

//fd state: (0) 84uCz6da8cQPC

ret = client.Write(fd869, []byte("0FgTNvFVVXm7peuBP6rGEVprBpf7uLMa771CrlfDLSsAQtmH8aTF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) 0FgTNvFVVXm7peuBP6rGEVprBpf7uLMa771CrlfDLSsAQtmH8aTF

buf, ret = client.Read(fd941, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd957, []byte("PfXXiVxK0wyQGkwpgwM6nmSy2YlUroAjJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) PfXXiVxK0wyQGkwpgwM6nmSy2YlUroAjJ

ret = client.Close(fd942)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd956)
if(ret != 0) {
  panic("close failed")
}


fd965 := client.Open("/2MKMM17MVW.txt", client.O_RDWR|client.O_CREATE)
if(fd965 < 0) {
  panic("open failed")
}


ret = client.Close(fd944)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd918, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


fd966 := client.Open("/Cyzai1XC0n.txt", client.O_RDWR|client.O_CREATE)
if(fd966 < 0) {
  panic("open failed")
}

//fd state: (99) AR1zBDL_b6KGYmyhjZOGLUipfuAQsLzYeOiLS18mkeVaRfZBvKfnk6rT03Ktgms1aGwX7jWwXEpT_mKWlfKPXtCBGbpITZaSlVV

ret = client.Write(fd945, []byte("AeWNfUV8cNJ6zZON4IoZ5rqrwHa0Eu3FEUAK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) AR1zBDL_b6KGYmyhjZOGLUipfuAQsLzYeOiLS18mkeVaRfZBvKfnk6rT03Ktgms1aGwX7jWwXEpT_mKWlfKPXtCBGbpITZaSlVVAeWNfUV8cNJ6zZON4IoZ5rqrwHa0Eu3FEUAK
//fd state: (0) 

ret = client.Write(fd951, []byte("U9RcyhHbhlPaDDJB_DlzWVrvuxe72t2TlYMdDlIG647sal1jM9geks_XanQiDcUozbc2zMh_OW_JqI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) U9RcyhHbhlPaDDJB_DlzWVrvuxe72t2TlYMdDlIG647sal1jM9geks_XanQiDcUozbc2zMh_OW_JqI

ret = client.Seek(fd920, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}

//fd state: (20) NJq0QDVdiMWZtm1xY_AHGe6iZOPa_Bg

ret = client.Write(fd962, []byte("OagdaNiTIJe42"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) NJq0QDVdiMWZtm1xY_AHOagdaNiTIJe42

ret = client.Seek(fd959, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


ret = client.Seek(fd961, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Seek(fd950, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd890, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


ret = client.Seek(fd966, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd943, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd967 := client.Open("/tgCyWNIEdr.txt", client.O_RDWR|client.O_CREATE)
if(fd967 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd949, []byte("q1RrImBKv_KbDfZCwko3BZD8a28xsIJ3jzmgDPQNaC_gM0o4g6IOIaD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) q1RrImBKv_KbDfZCwko3BZD8a28xsIJ3jzmgDPQNaC_gM0o4g6IOIaD

fd968 := client.Open("/u8ZwKWg3El.txt", client.O_RDWR|client.O_CREATE)
if(fd968 < 0) {
  panic("open failed")
}


ret = client.Close(fd967)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd962, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


fd969 := client.Open("/odYHg5Vmi4.txt", client.O_RDWR|client.O_CREATE)
if(fd969 < 0) {
  panic("open failed")
}


ret = client.Close(fd964)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd951, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


fd970 := client.Open("/zmZBlQdcRy.txt", client.O_RDWR|client.O_CREATE)
if(fd970 < 0) {
  panic("open failed")
}


ret = client.Seek(fd884, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd952, []byte("EQpuZdkeGotg7zEllb28k_oBsO6kKR9_f1lFf13Sf6_HJkhs6ZLEarINtDspborL2zGz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) EQpuZdkeGotg7zEllb28k_oBsO6kKR9_f1lFf13Sf6_HJkhs6ZLEarINtDspborL2zGz

ret = client.Close(fd950)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd918, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}


buf, ret = client.Read(fd959, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lC8hiIweJLTyN6BJeJCgXfvkpTCLBurgsdz") {
  panic("wrong data returned")
}


ret = client.Close(fd890)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd965, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "afce4oSimmR5gAI9GN87KJlQnhmhe") {
  panic("wrong data returned")
}


fd971 := client.Open("/85WW49JiFd.txt", client.O_RDWR|client.O_CREATE)
if(fd971 < 0) {
  panic("open failed")
}


fd972 := client.Open("/L3a7VdOTLE.txt", client.O_RDWR|client.O_CREATE)
if(fd972 < 0) {
  panic("open failed")
}


fd973 := client.Open("/rkt95xRaA4.txt", client.O_RDWR|client.O_CREATE)
if(fd973 < 0) {
  panic("open failed")
}

//fd state: (81) niO4SJNC3PHAxGXxO8i3rOEcD0sBwEJ4vHy0cd9S6N0vKnzTh_EwJJkGHWEUMNTZcEh9m8TxaicmVbjVT88Tk1Mdkokm6Uy

ret = client.Write(fd815, []byte("1TzA5VBut47b8THRhyaKpixHBrLoHQezlQb51__wnp2HK_ZtZlzeuYTr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (137) niO4SJNC3PHAxGXxO8i3rOEcD0sBwEJ4vHy0cd9S6N0vKnzTh_EwJJkGHWEUMNTZcEh9m8TxaicmVbjVT1TzA5VBut47b8THRhyaKpixHBrLoHQezlQb51__wnp2HK_ZtZlzeuYTr

fd974 := client.Open("/Kys6FX3D4n.txt", client.O_RDWR|client.O_CREATE)
if(fd974 < 0) {
  panic("open failed")
}


ret = client.Seek(fd945, 94, client.SEEK_SET)
if(ret != 94) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 94)
  panic("seek failed")
}


ret = client.Close(fd934)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd920)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd971, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd971, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd961)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd958, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "l32PuHlCBlmKQtFL6TcpxiHru2P1gqg6l0S6pUopYFnamAF") {
  panic("wrong data returned")
}


ret = client.Close(fd918)
if(ret != 0) {
  panic("close failed")
}


fd975 := client.Open("/wQOfDmuJ6Y.txt", client.O_RDWR|client.O_CREATE)
if(fd975 < 0) {
  panic("open failed")
}


ret = client.Close(fd911)
if(ret != 0) {
  panic("close failed")
}


fd976 := client.Open("/7kjLfusQPx.txt", client.O_RDWR|client.O_CREATE)
if(fd976 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd904, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (94) AR1zBDL_b6KGYmyhjZOGLUipfuAQsLzYeOiLS18mkeVaRfZBvKfnk6rT03Ktgms1aGwX7jWwXEpT_mKWlfKPXtCBGbpITZaSlVVAeWNfUV8cNJ6zZON4IoZ5rqrwHa0Eu3FEUAK

ret = client.Write(fd945, []byte("b6F6IGV7U5jGaZ3j3Koa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) AR1zBDL_b6KGYmyhjZOGLUipfuAQsLzYeOiLS18mkeVaRfZBvKfnk6rT03Ktgms1aGwX7jWwXEpT_mKWlfKPXtCBGbpITZb6F6IGV7U5jGaZ3j3KoaN4IoZ5rqrwHa0Eu3FEUAK

buf, ret = client.Read(fd957, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd977 := client.Open("/HM33hNH0QN.txt", client.O_RDWR|client.O_CREATE)
if(fd977 < 0) {
  panic("open failed")
}


ret = client.Close(fd972)
if(ret != 0) {
  panic("close failed")
}

//fd state: (177) qogceTTaIwJbY2PGee8fZbUKhvGeqkGX8VTT8OJ0Mb0VQt5cxcgrNKLN_K6q__dHhR3hmiM1JSIga9HLjxYTWwSMjoIjVXlHXRAhyEslMHTE4Uc7cNbJZZpLEc1y5EXPVmURYr0SBcB3FHeOJpwvOJ2OxHg5d2f_uqCOafH60OZwScuTG

ret = client.Write(fd932, []byte("9doPNwFdXNg2MidHtW5VMEFugvBLj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (206) qogceTTaIwJbY2PGee8fZbUKhvGeqkGX8VTT8OJ0Mb0VQt5cxcgrNKLN_K6q__dHhR3hmiM1JSIga9HLjxYTWwSMjoIjVXlHXRAhyEslMHTE4Uc7cNbJZZpLEc1y5EXPVmURYr0SBcB3FHeOJpwvOJ2OxHg5d2f_uqCOafH60OZwScuTG9doPNwFdXNg2MidHtW5VMEFugvBLj

ret = client.Close(fd869)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd976, []byte("R0HhadJIrzdzSKkB834dVYWxFiPWWMPUMJlqXrHf9x6k9A3xvlfhgL44RdC9ayH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) R0HhadJIrzdzSKkB834dVYWxFiPWWMPUMJlqXrHf9x6k9A3xvlfhgL44RdC9ayH

buf, ret = client.Read(fd941, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd978 := client.Open("/2CT8Qjb05c.txt", client.O_RDWR|client.O_CREATE)
if(fd978 < 0) {
  panic("open failed")
}


ret = client.Close(fd975)
if(ret != 0) {
  panic("close failed")
}

//fd state: (21) NJq0QDVdiMWZtm1xY_AHOagdaNiTIJe42

ret = client.Write(fd962, []byte("CmEMIdTqFubuWzTkjEvc_VWvlGXoPO1Vx7SMHFM5qh19kawXK_24cmWxrEA4ef"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) NJq0QDVdiMWZtm1xY_AHOCmEMIdTqFubuWzTkjEvc_VWvlGXoPO1Vx7SMHFM5qh19kawXK_24cmWxrEA4ef

ret = client.Seek(fd969, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd965, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd951)
if(ret != 0) {
  panic("close failed")
}


fd979 := client.Open("/cLPoy4URMf.txt", client.O_RDWR|client.O_CREATE)
if(fd979 < 0) {
  panic("open failed")
}


ret = client.Seek(fd971, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd970, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd980 := client.Open("/1gxKAeqOET.txt", client.O_RDWR|client.O_CREATE)
if(fd980 < 0) {
  panic("open failed")
}


ret = client.Close(fd973)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd958)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd979, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd941, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


buf, ret = client.Read(fd904, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd965, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


buf, ret = client.Read(fd965, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "KJlQnhmhe") {
  panic("wrong data returned")
}

//fd state: (71) TMMEEltoHnEdOENtHTGP2KCxHViCfJDO9l6KlC8hiIweJLTyN6BJeJCgXfvkpTCLBurgsdz5U_QgP3I8SeiJA18vE

ret = client.Write(fd959, []byte("nhgFXojz2qqpk_uM3TYjsKl4LSGZjZGVU5XS0znlPPAqcDd7A7on4eYyB2GeInwgH0as9Avxw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) TMMEEltoHnEdOENtHTGP2KCxHViCfJDO9l6KlC8hiIweJLTyN6BJeJCgXfvkpTCLBurgsdznhgFXojz2qqpk_uM3TYjsKl4LSGZjZGVU5XS0znlPPAqcDd7A7on4eYyB2GeInwgH0as9Avxw

buf, ret = client.Read(fd974, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd971, []byte("S7Stp6SyeRuOM_6kfqKL0WnVdHsGE1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) S7Stp6SyeRuOM_6kfqKL0WnVdHsGE1
//fd state: (144) TMMEEltoHnEdOENtHTGP2KCxHViCfJDO9l6KlC8hiIweJLTyN6BJeJCgXfvkpTCLBurgsdznhgFXojz2qqpk_uM3TYjsKl4LSGZjZGVU5XS0znlPPAqcDd7A7on4eYyB2GeInwgH0as9Avxw

ret = client.Write(fd959, []byte("04dse55mzYiSMuq8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (160) TMMEEltoHnEdOENtHTGP2KCxHViCfJDO9l6KlC8hiIweJLTyN6BJeJCgXfvkpTCLBurgsdznhgFXojz2qqpk_uM3TYjsKl4LSGZjZGVU5XS0znlPPAqcDd7A7on4eYyB2GeInwgH0as9Avxw04dse55mzYiSMuq8

buf, ret = client.Read(fd977, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (30) S7Stp6SyeRuOM_6kfqKL0WnVdHsGE1

ret = client.Write(fd971, []byte("pUrAT4tc6SE0ffZie8c0RTVmHIgl2F3RxqLB2y1ejooM8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) S7Stp6SyeRuOM_6kfqKL0WnVdHsGE1pUrAT4tc6SE0ffZie8c0RTVmHIgl2F3RxqLB2y1ejooM8

ret = client.Seek(fd976, 61, client.SEEK_SET)
if(ret != 61) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 61)
  panic("seek failed")
}


ret = client.Seek(fd980, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (83) NJq0QDVdiMWZtm1xY_AHOCmEMIdTqFubuWzTkjEvc_VWvlGXoPO1Vx7SMHFM5qh19kawXK_24cmWxrEA4ef

ret = client.Write(fd962, []byte("AIOauax1O4rxsbygtOq5_TpzZGEGCVx3aEpi_Di8vJQS09iJ6SUW7r1yweXzRd4S5MPZvm8O8jlvO7D_9q7LRxg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (170) NJq0QDVdiMWZtm1xY_AHOCmEMIdTqFubuWzTkjEvc_VWvlGXoPO1Vx7SMHFM5qh19kawXK_24cmWxrEA4efAIOauax1O4rxsbygtOq5_TpzZGEGCVx3aEpi_Di8vJQS09iJ6SUW7r1yweXzRd4S5MPZvm8O8jlvO7D_9q7LRxg

buf, ret = client.Read(fd904, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd959, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}


ret = client.Close(fd925)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd979, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd977, []byte("PSvWe7UzrV542yUXidThQ_6GxzPqTvLL1yu3fPn1kmlC7SyH3PYo_QlEVtuZbBUlEupwlY5G_fdCIUGl5_E5v50VlfWxoyneg2V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) PSvWe7UzrV542yUXidThQ_6GxzPqTvLL1yu3fPn1kmlC7SyH3PYo_QlEVtuZbBUlEupwlY5G_fdCIUGl5_E5v50VlfWxoyneg2V

fd981 := client.Open("/ru2kz3GZxr.txt", client.O_RDWR|client.O_CREATE)
if(fd981 < 0) {
  panic("open failed")
}

//fd state: (137) niO4SJNC3PHAxGXxO8i3rOEcD0sBwEJ4vHy0cd9S6N0vKnzTh_EwJJkGHWEUMNTZcEh9m8TxaicmVbjVT1TzA5VBut47b8THRhyaKpixHBrLoHQezlQb51__wnp2HK_ZtZlzeuYTr

ret = client.Write(fd815, []byte("45Moo3jdbsa1sCC7wnepgK3X2s4IyJwcIOHL6dsCaFNpB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (182) niO4SJNC3PHAxGXxO8i3rOEcD0sBwEJ4vHy0cd9S6N0vKnzTh_EwJJkGHWEUMNTZcEh9m8TxaicmVbjVT1TzA5VBut47b8THRhyaKpixHBrLoHQezlQb51__wnp2HK_ZtZlzeuYTr45Moo3jdbsa1sCC7wnepgK3X2s4IyJwcIOHL6dsCaFNpB

buf, ret = client.Read(fd947, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd966, []byte("EMlH9l3G6MjOGIh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) EMlH9l3G6MjOGIh

ret = client.Close(fd965)
if(ret != 0) {
  panic("close failed")
}

//fd state: (55) cbEyh4wC1mLgOHDNyBHRbFh7WWCPpXNAn7E1Ao6SPQ4EYksE0g1uPAMSoC79xFntwvJ_gwwKVrnuvKv8mxfkPGQBp4YzYHDYTIgj8

ret = client.Write(fd884, []byte("wBHxJKQvYCc2YUNczCJf2vggSNGN0aBggftVAF_AUp4sbzDUKFWI4BDQQo3CBlUKH2oPtV_b1P9Mey3trrDFM5KUcuGsXvKtOqd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) cbEyh4wC1mLgOHDNyBHRbFh7WWCPpXNAn7E1Ao6SPQ4EYksE0g1uPAMwBHxJKQvYCc2YUNczCJf2vggSNGN0aBggftVAF_AUp4sbzDUKFWI4BDQQo3CBlUKH2oPtV_b1P9Mey3trrDFM5KUcuGsXvKtOqd

fd982 := client.Open("/iJRMsrR9zU.txt", client.O_RDWR|client.O_CREATE)
if(fd982 < 0) {
  panic("open failed")
}


ret = client.Seek(fd945, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


fd983 := client.Open("/0PCO56vqgQ.txt", client.O_RDWR|client.O_CREATE)
if(fd983 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd983, []byte("mTwMmae1vfEIvumPWczeAvWEEm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) mTwMmae1vfEIvumPWczeAvWEEm

ret = client.Close(fd952)
if(ret != 0) {
  panic("close failed")
}

//fd state: (63) pNefSqjGE7e1gnSKHLNbA2dLf7rEMbN78N9vwPGcL5TxmbLiV41JKZhp87SpzrL

ret = client.Write(fd960, []byte("SbMCmaszJmnBf8Tkn54Od_SfmoPzWFwOcK6AkBe9y0_7Nf8L3YSVCEwecg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (121) pNefSqjGE7e1gnSKHLNbA2dLf7rEMbN78N9vwPGcL5TxmbLiV41JKZhp87SpzrLSbMCmaszJmnBf8Tkn54Od_SfmoPzWFwOcK6AkBe9y0_7Nf8L3YSVCEwecg

fd984 := client.Open("/hB_gQxlIvw.txt", client.O_RDWR|client.O_CREATE)
if(fd984 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd947, []byte("V12ChUcC2rIzcOdN1HxKAPHDLRj1S8R4CAsRVmfZHbUzYnn5W9tvlJVDAj6yKSY6reB9h21MlIGoW9Q3VW0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) V12ChUcC2rIzcOdN1HxKAPHDLRj1S8R4CAsRVmfZHbUzYnn5W9tvlJVDAj6yKSY6reB9h21MlIGoW9Q3VW0

buf, ret = client.Read(fd959, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GZjZGVU5XS0znlPPAqcDd7A7on4eYyB2") {
  panic("wrong data returned")
}

//fd state: (93) M2dhmpvBQ4IlkLbY7FLBVOBh3rr82kAK1OfYqsKkSZeftn1KGOOjUMXWTvhTxOF43xhE2InTArqwI8oxAxWvxEpVbbjHq

ret = client.Write(fd943, []byte("e"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) M2dhmpvBQ4IlkLbY7FLBVOBh3rr82kAK1OfYqsKkSZeftn1KGOOjUMXWTvhTxOF43xhE2InTArqwI8oxAxWvxEpVbbjHqe

ret = client.Seek(fd981, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (55) q1RrImBKv_KbDfZCwko3BZD8a28xsIJ3jzmgDPQNaC_gM0o4g6IOIaD

ret = client.Write(fd949, []byte("0ZjfV8QYT7erqZg9q7zKVb8pNNKrLUGH57yeuXfvdwqyQIBAffvnUj_kAO0DNMli_LihDzPoLAQntIH4k9uDFdqMHma6hWZ5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) q1RrImBKv_KbDfZCwko3BZD8a28xsIJ3jzmgDPQNaC_gM0o4g6IOIaD0ZjfV8QYT7erqZg9q7zKVb8pNNKrLUGH57yeuXfvdwqyQIBAffvnUj_kAO0DNMli_LihDzPoLAQntIH4k9uDFdqMHma6hWZ5
//fd state: (121) pNefSqjGE7e1gnSKHLNbA2dLf7rEMbN78N9vwPGcL5TxmbLiV41JKZhp87SpzrLSbMCmaszJmnBf8Tkn54Od_SfmoPzWFwOcK6AkBe9y0_7Nf8L3YSVCEwecg

ret = client.Write(fd960, []byte("F9Hrs8tjo2wbVCVwvZhNtv2Q7hUu9nKy8iG_zjh3egVkqUMHxu4PsUKidqcNe4dIwJweO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (190) pNefSqjGE7e1gnSKHLNbA2dLf7rEMbN78N9vwPGcL5TxmbLiV41JKZhp87SpzrLSbMCmaszJmnBf8Tkn54Od_SfmoPzWFwOcK6AkBe9y0_7Nf8L3YSVCEwecgF9Hrs8tjo2wbVCVwvZhNtv2Q7hUu9nKy8iG_zjh3egVkqUMHxu4PsUKidqcNe4dIwJweO

ret = client.Close(fd932)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd970, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd976, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}

//fd state: (129) TMMEEltoHnEdOENtHTGP2KCxHViCfJDO9l6KlC8hiIweJLTyN6BJeJCgXfvkpTCLBurgsdznhgFXojz2qqpk_uM3TYjsKl4LSGZjZGVU5XS0znlPPAqcDd7A7on4eYyB2GeInwgH0as9Avxw04dse55mzYiSMuq8

ret = client.Write(fd959, []byte("JV6xDTsRRMV48JjdPiY0Nz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) TMMEEltoHnEdOENtHTGP2KCxHViCfJDO9l6KlC8hiIweJLTyN6BJeJCgXfvkpTCLBurgsdznhgFXojz2qqpk_uM3TYjsKl4LSGZjZGVU5XS0znlPPAqcDd7A7on4eYyB2JV6xDTsRRMV48JjdPiY0NzmzYiSMuq8

ret = client.Seek(fd968, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd949, 92, client.SEEK_SET)
if(ret != 92) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 92)
  panic("seek failed")
}


ret = client.Seek(fd962, 92, client.SEEK_SET)
if(ret != 92) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 92)
  panic("seek failed")
}


fd985 := client.Open("/5QlPGMS3lm.txt", client.O_RDWR|client.O_CREATE)
if(fd985 < 0) {
  panic("open failed")
}

//fd state: (43) RBmjO4eHsy9DSrCd9Rv9_ufL5vYCrdoOlcKp0ZuW8owr89t4_5HNhUPVGqup6CWvIoYywRKVnGXPlSdIadzXwEQMGZCDTKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGdxoPlKW7lbYhouaDYod8l359wnP04T1CrJ9pj2x2GFzeliIJWpQG7IBB3airdkABEq6488D

ret = client.Write(fd887, []byte("R_1tYJoamqacsOH5VFcR7jKg8011hkYjtVcEthIEtddON0cLhG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) RBmjO4eHsy9DSrCd9Rv9_ufL5vYCrdoOlcKp0ZuW8owR_1tYJoamqacsOH5VFcR7jKg8011hkYjtVcEthIEtddON0cLhGKfJj6tiIlkfetbIxiau4iJQVy0nKncwwGA9tPQDVGS3Ot7HCawWVgNzLkJhB4XQb4ppwRcDV2XDeUYGdxoPlKW7lbYhouaDYod8l359wnP04T1CrJ9pj2x2GFzeliIJWpQG7IBB3airdkABEq6488D

buf, ret = client.Read(fd982, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "KXJVV2kwoBVKQVKtAjo0sl6MOVebD1mQmuW8SOd1ZDlHIi0t3DPWIzzJHyeWUYpHg2L44X7") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd904, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (75) S7Stp6SyeRuOM_6kfqKL0WnVdHsGE1pUrAT4tc6SE0ffZie8c0RTVmHIgl2F3RxqLB2y1ejooM8

ret = client.Write(fd971, []byte("EWp__a9zx4iYpIP0oSdO1QrFfr3ybYqJx21QqcU9fjJN7YCcGp8O53WrxjZTSvTTrPjRFiy9JsLgRxXMBJ5b"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (159) S7Stp6SyeRuOM_6kfqKL0WnVdHsGE1pUrAT4tc6SE0ffZie8c0RTVmHIgl2F3RxqLB2y1ejooM8EWp__a9zx4iYpIP0oSdO1QrFfr3ybYqJx21QqcU9fjJN7YCcGp8O53WrxjZTSvTTrPjRFiy9JsLgRxXMBJ5b

ret = client.Seek(fd969, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd986 := client.Open("/D7PuDI5cX1.txt", client.O_RDWR|client.O_CREATE)
if(fd986 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd947, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd983, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Seek(fd968, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Close(fd955)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd887, 134, client.SEEK_SET)
if(ret != 134) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 134)
  panic("seek failed")
}


buf, ret = client.Read(fd959, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "mzYiSMuq8") {
  panic("wrong data returned")
}


fd987 := client.Open("/zAAGw1724E.txt", client.O_RDWR|client.O_CREATE)
if(fd987 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd978, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "KBZR") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd960, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd988 := client.Open("/ymeRdV7EO7.txt", client.O_RDWR|client.O_CREATE)
if(fd988 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd983, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "EIvumPWczeAvWEEm") {
  panic("wrong data returned")
}

//fd state: (15) EMlH9l3G6MjOGIh

ret = client.Write(fd966, []byte("0AWShfQBiTGFTqCOJckEd1cKdkq9DzD8WuA6n1h"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) EMlH9l3G6MjOGIh0AWShfQBiTGFTqCOJckEd1cKdkq9DzD8WuA6n1h
//fd state: (190) pNefSqjGE7e1gnSKHLNbA2dLf7rEMbN78N9vwPGcL5TxmbLiV41JKZhp87SpzrLSbMCmaszJmnBf8Tkn54Od_SfmoPzWFwOcK6AkBe9y0_7Nf8L3YSVCEwecgF9Hrs8tjo2wbVCVwvZhNtv2Q7hUu9nKy8iG_zjh3egVkqUMHxu4PsUKidqcNe4dIwJweO

ret = client.Write(fd960, []byte("Nqzcxz1Q9xHxsn8ITwTmA_dXxhvAMR5K0dBN9gV4hDmP7qAtCfYQZLB8Ia9J6X2rdl1NkJMazKKY__C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (269) pNefSqjGE7e1gnSKHLNbA2dLf7rEMbN78N9vwPGcL5TxmbLiV41JKZhp87SpzrLSbMCmaszJmnBf8Tkn54Od_SfmoPzWFwOcK6AkBe9y0_7Nf8L3YSVCEwecgF9Hrs8tjo2wbVCVwvZhNtv2Q7hUu9nKy8iG_zjh3egVkqUMHxu4PsUKidqcNe4dIwJweONqzcxz1Q9xHxsn8ITwTmA_dXxhvAMR5K0dBN9gV4hDmP7qAtCfYQZLB8Ia9J6X2rdl1NkJMazKKY__C

ret = client.Close(fd984)
if(ret != 0) {
  panic("close failed")
}


fd989 := client.Open("/hm40EpXTIC.txt", client.O_RDWR|client.O_CREATE)
if(fd989 < 0) {
  panic("open failed")
}


ret = client.Seek(fd979, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd990 := client.Open("/7XgDJ0dpBX.txt", client.O_RDWR|client.O_CREATE)
if(fd990 < 0) {
  panic("open failed")
}


ret = client.Close(fd968)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd904)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd981, []byte("OlEhwcrX2ADChpEnBnGq_VDtgPnI61TjbpyMd0VFz52CMlzKYeUTW9f2U6poGI_NNq8pS0rs8MfFul6FinoiLVBjaajgcB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) OlEhwcrX2ADChpEnBnGq_VDtgPnI61TjbpyMd0VFz52CMlzKYeUTW9f2U6poGI_NNq8pS0rs8MfFul6FinoiLVBjaajgcB

ret = client.Seek(fd960, 197, client.SEEK_SET)
if(ret != 197) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 197)
  panic("seek failed")
}


ret = client.Close(fd982)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd947, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd957)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd966, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd945, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vKfnk6rT03Ktgms1aGwX7jWwXEpT") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd990, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd974, []byte("cm__6cDPdGWqf6kXL_h7xgUnpSwM9C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) cm__6cDPdGWqf6kXL_h7xgUnpSwM9C

ret = client.Seek(fd949, 89, client.SEEK_SET)
if(ret != 89) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 89)
  panic("seek failed")
}


buf, ret = client.Read(fd981, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd989, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


ret = client.Seek(fd981, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


buf, ret = client.Read(fd960, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Q9xHxsn8ITwTmA_dXxhvAMR5K0dBN9gV4hDmP7qAtCfYQZLB8Ia") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd988, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "KmjijiMRbfIhKvYBrOxFasqE4PMpfBnbRKtso4H8GIZb3W3W10o6K5IclG") {
  panic("wrong data returned")
}


ret = client.Seek(fd981, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


ret = client.Seek(fd962, 110, client.SEEK_SET)
if(ret != 110) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 110)
  panic("seek failed")
}


fd991 := client.Open("/To5eyisDjM.txt", client.O_RDWR|client.O_CREATE)
if(fd991 < 0) {
  panic("open failed")
}


fd992 := client.Open("/N2wW4FWwc1.txt", client.O_RDWR|client.O_CREATE)
if(fd992 < 0) {
  panic("open failed")
}


ret = client.Close(fd988)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd815, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


buf, ret = client.Read(fd945, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_mKWlfKPXtCBGbpITZb6F6IGV7U5jGaZ3j3KoaN4IoZ5rqrwHa0Eu3FEUAK") {
  panic("wrong data returned")
}


fd993 := client.Open("/w4rp2cpYos.txt", client.O_RDWR|client.O_CREATE)
if(fd993 < 0) {
  panic("open failed")
}


ret = client.Seek(fd989, 92, client.SEEK_SET)
if(ret != 92) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 92)
  panic("seek failed")
}


fd994 := client.Open("/PHdLRfLHpq.txt", client.O_RDWR|client.O_CREATE)
if(fd994 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd985, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Wlz_UQR4wubMYfWWDccgvufGj96H60DdSOg2T") {
  panic("wrong data returned")
}


ret = client.Close(fd981)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd945, 46)
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


ret = client.Close(fd974)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd970)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd976, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


ret = client.Seek(fd815, 124, client.SEEK_SET)
if(ret != 124) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 124)
  panic("seek failed")
}


ret = client.Seek(fd959, 103, client.SEEK_SET)
if(ret != 103) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 103)
  panic("seek failed")
}


ret = client.Close(fd947)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd978)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd966, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd945, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd993, []byte("u2ynz53Xq7fiTURHsFC1sqGJn_0djMH_zWhMTIBiCnKjyyPb1quK0p3E0EnJGEe2VtahTFNrlMir9Ctx2tiHD3McTCUL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) u2ynz53Xq7fiTURHsFC1sqGJn_0djMH_zWhMTIBiCnKjyyPb1quK0p3E0EnJGEe2VtahTFNrlMir9Ctx2tiHD3McTCUL

buf, ret = client.Read(fd971, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd945, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}


buf, ret = client.Read(fd887, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "S3Ot7HCawWVgNzLkJhB4XQb4") {
  panic("wrong data returned")
}


ret = client.Close(fd983)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd966, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


ret = client.Close(fd990)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd971, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd995 := client.Open("/jv1lGjfNx_.txt", client.O_RDWR|client.O_CREATE)
if(fd995 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd971, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd887, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ppwRcDV2XDeUYGdxoPlK") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd992, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_Rjx1aHCYC5V") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd884, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd996 := client.Open("/2vuPlCmIVr.txt", client.O_RDWR|client.O_CREATE)
if(fd996 < 0) {
  panic("open failed")
}


ret = client.Close(fd995)
if(ret != 0) {
  panic("close failed")
}


fd997 := client.Open("/nRHcbaClkX.txt", client.O_RDWR|client.O_CREATE)
if(fd997 < 0) {
  panic("open failed")
}


ret = client.Seek(fd993, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


fd998 := client.Open("/413b9StdCY.txt", client.O_RDWR|client.O_CREATE)
if(fd998 < 0) {
  panic("open failed")
}


fd999 := client.Open("/DZ6ujpIdos.txt", client.O_RDWR|client.O_CREATE)
if(fd999 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd985, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd941)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd994, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd986, []byte("fp183_GhRFDrKSqDm0f7T073ytLdw018M4iDEQo5QoldTqOW7K9IiPf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) fp183_GhRFDrKSqDm0f7T073ytLdw018M4iDEQo5QoldTqOW7K9IiPf

ret = client.Seek(fd993, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd999, []byte("AwVnljDFIS24X4t2ShBtoxrDl27pp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) AwVnljDFIS24X4t2ShBtoxrDl27pp

fd1000 := client.Open("/uJt5Xeux53.txt", client.O_RDWR|client.O_CREATE)
if(fd1000 < 0) {
  panic("open failed")
}


ret = client.Seek(fd884, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


ret = client.Seek(fd999, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Close(fd994)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd993)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd999, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


fd1001 := client.Open("/csWUk3qpcU.txt", client.O_RDWR|client.O_CREATE)
if(fd1001 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd997, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd989, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pKBxyeINthBidt8jO0PjmcMyi6140") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd960, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9J6X2rdl1NkJMazKKY__C") {
  panic("wrong data returned")
}


fd1002 := client.Open("/9xkiAGJeWP.txt", client.O_RDWR|client.O_CREATE)
if(fd1002 < 0) {
  panic("open failed")
}


ret = client.Close(fd884)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd985, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


buf, ret = client.Read(fd949, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yeuXfvdwqyQIBAffvnUj_kAO0DNMli_LihDzPoLAQntIH4k9uDFdqMHma6hWZ5") {
  panic("wrong data returned")
}


fd1003 := client.Open("/N36CKWafKy.txt", client.O_RDWR|client.O_CREATE)
if(fd1003 < 0) {
  panic("open failed")
}


ret = client.Seek(fd985, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


fd1004 := client.Open("/eXNFP7XOhv.txt", client.O_RDWR|client.O_CREATE)
if(fd1004 < 0) {
  panic("open failed")
}


fd1005 := client.Open("/rYlWt17Jt6.txt", client.O_RDWR|client.O_CREATE)
if(fd1005 < 0) {
  panic("open failed")
}


fd1006 := client.Open("/zLcSjOrUCn.txt", client.O_RDWR|client.O_CREATE)
if(fd1006 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1005, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (94) M2dhmpvBQ4IlkLbY7FLBVOBh3rr82kAK1OfYqsKkSZeftn1KGOOjUMXWTvhTxOF43xhE2InTArqwI8oxAxWvxEpVbbjHqe

ret = client.Write(fd943, []byte("FKdi42yoFJZXo4El1VozYUb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) M2dhmpvBQ4IlkLbY7FLBVOBh3rr82kAK1OfYqsKkSZeftn1KGOOjUMXWTvhTxOF43xhE2InTArqwI8oxAxWvxEpVbbjHqeFKdi42yoFJZXo4El1VozYUb

ret = client.Close(fd1006)
if(ret != 0) {
  panic("close failed")
}


fd1007 := client.Open("/jg0Jvu7lJx.txt", client.O_RDWR|client.O_CREATE)
if(fd1007 < 0) {
  panic("open failed")
}


fd1008 := client.Open("/kdkOHSRy4f.txt", client.O_RDWR|client.O_CREATE)
if(fd1008 < 0) {
  panic("open failed")
}


fd1009 := client.Open("/GlIUkocnJk.txt", client.O_RDWR|client.O_CREATE)
if(fd1009 < 0) {
  panic("open failed")
}


ret = client.Close(fd1002)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd945, 120, client.SEEK_SET)
if(ret != 120) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 120)
  panic("seek failed")
}


fd1010 := client.Open("/AHQC8De8KJ.txt", client.O_RDWR|client.O_CREATE)
if(fd1010 < 0) {
  panic("open failed")
}


ret = client.Seek(fd998, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd1009, []byte("eeqsYmbYKYuUruHDNZBg3yUFGY1rGAlcJbcknCYsxBMqTPJnPkRNzIxhr3ahAnOLO3BXm3kBTTRoFsYIzAit1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) eeqsYmbYKYuUruHDNZBg3yUFGY1rGAlcJbcknCYsxBMqTPJnPkRNzIxhr3ahAnOLO3BXm3kBTTRoFsYIzAit1

ret = client.Seek(fd960, 128, client.SEEK_SET)
if(ret != 128) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 128)
  panic("seek failed")
}

//fd state: (159) S7Stp6SyeRuOM_6kfqKL0WnVdHsGE1pUrAT4tc6SE0ffZie8c0RTVmHIgl2F3RxqLB2y1ejooM8EWp__a9zx4iYpIP0oSdO1QrFfr3ybYqJx21QqcU9fjJN7YCcGp8O53WrxjZTSvTTrPjRFiy9JsLgRxXMBJ5b

ret = client.Write(fd971, []byte("8RqW_017nuZR4EEZBMkX1p8NJOZ8VY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (189) S7Stp6SyeRuOM_6kfqKL0WnVdHsGE1pUrAT4tc6SE0ffZie8c0RTVmHIgl2F3RxqLB2y1ejooM8EWp__a9zx4iYpIP0oSdO1QrFfr3ybYqJx21QqcU9fjJN7YCcGp8O53WrxjZTSvTTrPjRFiy9JsLgRxXMBJ5b8RqW_017nuZR4EEZBMkX1p8NJOZ8VY

ret = client.Seek(fd1008, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd969, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd1008, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1011 := client.Open("/HN3w0IwSey.txt", client.O_RDWR|client.O_CREATE)
if(fd1011 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1010, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd992)
if(ret != 0) {
  panic("close failed")
}

//fd state: (124) niO4SJNC3PHAxGXxO8i3rOEcD0sBwEJ4vHy0cd9S6N0vKnzTh_EwJJkGHWEUMNTZcEh9m8TxaicmVbjVT1TzA5VBut47b8THRhyaKpixHBrLoHQezlQb51__wnp2HK_ZtZlzeuYTr45Moo3jdbsa1sCC7wnepgK3X2s4IyJwcIOHL6dsCaFNpB

ret = client.Write(fd815, []byte("FxKsazYX2UjaSJFG_Riuengp6WKChqiKQqzwQLvMrDt_Fm9UDzwMAUlTgYT2ahYH59RVCRVEfjFP4PYvqsMGiBr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (211) niO4SJNC3PHAxGXxO8i3rOEcD0sBwEJ4vHy0cd9S6N0vKnzTh_EwJJkGHWEUMNTZcEh9m8TxaicmVbjVT1TzA5VBut47b8THRhyaKpixHBrLoHQezlQb51__wnp2FxKsazYX2UjaSJFG_Riuengp6WKChqiKQqzwQLvMrDt_Fm9UDzwMAUlTgYT2ahYH59RVCRVEfjFP4PYvqsMGiBr

ret = client.Seek(fd999, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


buf, ret = client.Read(fd960, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jo2wbVCVwvZhNtv2Q7hUu9nKy8iG_zjh3egVkqUMHxu4PsUKidqcNe4dIwJweONqzcxz1Q9xHxsn8ITwT") {
  panic("wrong data returned")
}


ret = client.Seek(fd987, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}


buf, ret = client.Read(fd943, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1012 := client.Open("/vseKGPl4Ii.txt", client.O_RDWR|client.O_CREATE)
if(fd1012 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd998, []byte("usqkOUdqya2LNhEXOTI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) usqkOUdqya2LNhEXOTI

buf, ret = client.Read(fd989, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "SQXXzyXoWcZADZZ2dbGHaF7RBCsb_") {
  panic("wrong data returned")
}


ret = client.Close(fd1010)
if(ret != 0) {
  panic("close failed")
}


fd1013 := client.Open("/rT5JLUyzSF.txt", client.O_RDWR|client.O_CREATE)
if(fd1013 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1009, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd1004)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd1011)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd1013)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd949, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1012, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd998)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd1003)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd991, []byte("hpFR9gMJ26ulX496X6v0C1pCGBn85pfZY97rkR4jVu3T_Hg85sL5QTN1j8NKMsNsbepUEbQJA_aUXc447Y9I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) hpFR9gMJ26ulX496X6v0C1pCGBn85pfZY97rkR4jVu3T_Hg85sL5QTN1j8NKMsNsbepUEbQJA_aUXc447Y9I

ret = client.Seek(fd997, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (150) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLV_uc4CxoSW_ZNYnswS8SNbBgt6hmTM6XipKBxyeINthBidt8jO0PjmcMyi6140SQXXzyXoWcZADZZ2dbGHaF7RBCsb_gKochxVQgggcRVpXqzJKgrowNhfFlmgWptrMQ4klVenNFUMM4blo1xVuOTiy2a7Im0FnROqKPUUXDuZbjeqCm21Z9YaNR855

ret = client.Write(fd989, []byte("gXCsmWqezvZD2ZY1D9_aVihVH0GX72NHmVQ5UYn7IlnAsdRtjwWsvWH1YrSUT_Vq84UrTWOQO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (223) K623pmcU1AiD8pUCv19A1hXoKJ9QexFLx0pnSNnDsH96LvGmed2zLlA_tbLV_uc4CxoSW_ZNYnswS8SNbBgt6hmTM6XipKBxyeINthBidt8jO0PjmcMyi6140SQXXzyXoWcZADZZ2dbGHaF7RBCsb_gXCsmWqezvZD2ZY1D9_aVihVH0GX72NHmVQ5UYn7IlnAsdRtjwWsvWH1YrSUT_Vq84UrTWOQOUUXDuZbjeqCm21Z9YaNR855

ret = client.Seek(fd1012, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (209) pNefSqjGE7e1gnSKHLNbA2dLf7rEMbN78N9vwPGcL5TxmbLiV41JKZhp87SpzrLSbMCmaszJmnBf8Tkn54Od_SfmoPzWFwOcK6AkBe9y0_7Nf8L3YSVCEwecgF9Hrs8tjo2wbVCVwvZhNtv2Q7hUu9nKy8iG_zjh3egVkqUMHxu4PsUKidqcNe4dIwJweONqzcxz1Q9xHxsn8ITwTmA_dXxhvAMR5K0dBN9gV4hDmP7qAtCfYQZLB8Ia9J6X2rdl1NkJMazKKY__C

ret = client.Write(fd960, []byte("mGFSwQQy7dj_LL2XcsAod45x7vb6m4gOIj8VQIhPm72kOK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (255) pNefSqjGE7e1gnSKHLNbA2dLf7rEMbN78N9vwPGcL5TxmbLiV41JKZhp87SpzrLSbMCmaszJmnBf8Tkn54Od_SfmoPzWFwOcK6AkBe9y0_7Nf8L3YSVCEwecgF9Hrs8tjo2wbVCVwvZhNtv2Q7hUu9nKy8iG_zjh3egVkqUMHxu4PsUKidqcNe4dIwJweONqzcxz1Q9xHxsn8ITwTmGFSwQQy7dj_LL2XcsAod45x7vb6m4gOIj8VQIhPm72kOKl1NkJMazKKY__C

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}
