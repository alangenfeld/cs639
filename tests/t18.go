
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

fd1 := client.Open("hyGZ9P9pd6.txt", client.O_RDWR|client.O_CREATE)
if(fd1 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd1, []byte("AlNb1kNTgq7y2WBwzj72btP8B9vL5Yv3FBbOjEo0q3GoE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) AlNb1kNTgq7y2WBwzj72btP8B9vL5Yv3FBbOjEo0q3GoE

fd2 := client.Open("hyGZ9P9pd6.txt", client.O_RDWR|client.O_CREATE)
if(fd2 < 0) {
  panic("open failed")
}


ret = client.Close(fd1)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd2, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd2, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AlNb1kNTgq7y2WBwzj72btP8B9vL5Yv3FBbOjEo0q3GoE") {
  panic("wrong data returned")
}


ret = client.Seek(fd2, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


ret = client.Close(fd2)
if(ret != 0) {
  panic("close failed")
}


fd3 := client.Open("ufwwqNMK1w.txt", client.O_RDWR|client.O_CREATE)
if(fd3 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd3, []byte("A1YMvIxPYZnzjmToRwNWXo8VVobyzmoApaKrL8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) A1YMvIxPYZnzjmToRwNWXo8VVobyzmoApaKrL8

buf, ret = client.Read(fd3, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd3, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


buf, ret = client.Read(fd3, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "L8") {
  panic("wrong data returned")
}


ret = client.Close(fd3)
if(ret != 0) {
  panic("close failed")
}


fd4 := client.Open("RKQErSV99g.txt", client.O_RDWR|client.O_CREATE)
if(fd4 < 0) {
  panic("open failed")
}


fd5 := client.Open("jvjAJ3I1nx.txt", client.O_RDWR|client.O_CREATE)
if(fd5 < 0) {
  panic("open failed")
}


ret = client.Close(fd4)
if(ret != 0) {
  panic("close failed")
}


fd6 := client.Open("kzsba3KDvj.txt", client.O_RDWR|client.O_CREATE)
if(fd6 < 0) {
  panic("open failed")
}


ret = client.Seek(fd6, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd6)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd5, []byte("aQdGk0u8kp3EixUJg7hmKn9FI0KU5hWhRfYz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) aQdGk0u8kp3EixUJg7hmKn9FI0KU5hWhRfYz

ret = client.Close(fd5)
if(ret != 0) {
  panic("close failed")
}


fd7 := client.Open("jvjAJ3I1nx.txt", client.O_RDWR|client.O_CREATE)
if(fd7 < 0) {
  panic("open failed")
}

//fd state: (0) aQdGk0u8kp3EixUJg7hmKn9FI0KU5hWhRfYz

ret = client.Write(fd7, []byte("mBFgH9p9ZMO131xTUdEaPEMMt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) mBFgH9p9ZMO131xTUdEaPEMMt
//fd state: (25) mBFgH9p9ZMO131xTUdEaPEMMt

ret = client.Write(fd7, []byte("nDVkveTschypXTwi5V1NgQSO9J3UVFLR3Es9ekd7DB_Eu34LvuzqJngT3K5jEcsGkdpnJTk1rOnl5G_TzrxMn3ITWJ3jv1mwF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) mBFgH9p9ZMO131xTUdEaPEMMtnDVkveTschypXTwi5V1NgQSO9J3UVFLR3Es9ekd7DB_Eu34LvuzqJngT3K5jEcsGkdpnJTk1rOnl5G_TzrxMn3ITWJ3jv1mwF

ret = client.Seek(fd7, 65, client.SEEK_SET)
if(ret != 65) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 65)
  panic("seek failed")
}


buf, ret = client.Read(fd7, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "DB_Eu34LvuzqJngT") {
  panic("wrong data returned")
}


fd8 := client.Open("kzsba3KDvj.txt", client.O_RDWR|client.O_CREATE)
if(fd8 < 0) {
  panic("open failed")
}


ret = client.Seek(fd7, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Seek(fd7, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


ret = client.Close(fd8)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd7, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PEMMtnDVkveTschypXTwi5V1NgQSO9J3UVFLR3Es9ekd7DB_Eu34LvuzqJngT3K5jEcsGkdpnJTk1rO") {
  panic("wrong data returned")
}


fd9 := client.Open("06CKQJmL3s.txt", client.O_RDWR|client.O_CREATE)
if(fd9 < 0) {
  panic("open failed")
}


ret = client.Seek(fd9, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd9, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd9, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd9, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd7, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nl5G_TzrxMn3ITWJ3jv1mwF") {
  panic("wrong data returned")
}


fd10 := client.Open("vwimLePYvD.txt", client.O_RDWR|client.O_CREATE)
if(fd10 < 0) {
  panic("open failed")
}


ret = client.Close(fd9)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd7, 95, client.SEEK_SET)
if(ret != 95) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 95)
  panic("seek failed")
}


ret = client.Seek(fd7, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Close(fd7)
if(ret != 0) {
  panic("close failed")
}


fd11 := client.Open("ufwwqNMK1w.txt", client.O_RDWR|client.O_CREATE)
if(fd11 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd10, []byte("vdf4Ym5PsaNSdGM3hqqgXHvq10cajiv1vuMkTkbmbmdSu54bHU4PbkboARcDxa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) vdf4Ym5PsaNSdGM3hqqgXHvq10cajiv1vuMkTkbmbmdSu54bHU4PbkboARcDxa

fd12 := client.Open("qWFkXcSj6v.txt", client.O_RDWR|client.O_CREATE)
if(fd12 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd10, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd11)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd10, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Seek(fd12, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd10, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qqgXHvq10cajiv1vuMkTkbmbmdSu") {
  panic("wrong data returned")
}


fd13 := client.Open("cZ909fYTnT.txt", client.O_RDWR|client.O_CREATE)
if(fd13 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd13, []byte("i0Ourj1yFnr8cWMR73NLgwvIQaGgk7scZdPYyTQRWwMuHq0IG6cQiQSwwGU6hYG8nXWZF4GA0a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) i0Ourj1yFnr8cWMR73NLgwvIQaGgk7scZdPYyTQRWwMuHq0IG6cQiQSwwGU6hYG8nXWZF4GA0a

ret = client.Close(fd12)
if(ret != 0) {
  panic("close failed")
}


fd14 := client.Open("WWT0jJRj3Q.txt", client.O_RDWR|client.O_CREATE)
if(fd14 < 0) {
  panic("open failed")
}


ret = client.Seek(fd10, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd14, []byte("jyfu4x3xIe6c0LWMw7Jj99BDibu714srwWfDXYg0M9E7_iy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) jyfu4x3xIe6c0LWMw7Jj99BDibu714srwWfDXYg0M9E7_iy

fd15 := client.Open("hyGZ9P9pd6.txt", client.O_RDWR|client.O_CREATE)
if(fd15 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd15, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AlNb1kNTgq7y2WBwzj72btP8B9vL5Yv3FBbOjEo0q3GoE") {
  panic("wrong data returned")
}


fd16 := client.Open("Qsn2Nksy3K.txt", client.O_RDWR|client.O_CREATE)
if(fd16 < 0) {
  panic("open failed")
}


ret = client.Close(fd14)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd10, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


fd17 := client.Open("vwimLePYvD.txt", client.O_RDWR|client.O_CREATE)
if(fd17 < 0) {
  panic("open failed")
}


ret = client.Close(fd16)
if(ret != 0) {
  panic("close failed")
}


fd18 := client.Open("FVlYhkQ8JD.txt", client.O_RDWR|client.O_CREATE)
if(fd18 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd13, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd10, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}


fd19 := client.Open("WWT0jJRj3Q.txt", client.O_RDWR|client.O_CREATE)
if(fd19 < 0) {
  panic("open failed")
}


ret = client.Seek(fd10, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}


ret = client.Close(fd17)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd10, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd19, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Seek(fd18, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd19)
if(ret != 0) {
  panic("close failed")
}


fd20 := client.Open("CWV6d3q1rw.txt", client.O_RDWR|client.O_CREATE)
if(fd20 < 0) {
  panic("open failed")
}


ret = client.Seek(fd13, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}


buf, ret = client.Read(fd15, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd20, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd21 := client.Open("WWT0jJRj3Q.txt", client.O_RDWR|client.O_CREATE)
if(fd21 < 0) {
  panic("open failed")
}


fd22 := client.Open("QuIrnKRhbK.txt", client.O_RDWR|client.O_CREATE)
if(fd22 < 0) {
  panic("open failed")
}


ret = client.Close(fd13)
if(ret != 0) {
  panic("close failed")
}


fd23 := client.Open("yXO6X0Yzq0.txt", client.O_RDWR|client.O_CREATE)
if(fd23 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd10, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vdf4Ym5PsaNSdGM3hqqgXHvq10caj") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd18, 77)
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


ret = client.Seek(fd22, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd21, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}

//fd state: (11) jyfu4x3xIe6c0LWMw7Jj99BDibu714srwWfDXYg0M9E7_iy

ret = client.Write(fd21, []byte("beEFrW1kqEr8n2ihmQkvYhGxn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0M9E7_iy

ret = client.Close(fd15)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd23, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd23, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd21, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


buf, ret = client.Read(fd22, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (40) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0M9E7_iy

ret = client.Write(fd21, []byte("cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5CJbaphAb31dNB0BXaMk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5CJbaphAb31dNB0BXaMk
//fd state: (0) 

ret = client.Write(fd18, []byte("36nabY1ZHpbd_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) 36nabY1ZHpbd_
//fd state: (107) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5CJbaphAb31dNB0BXaMk

ret = client.Write(fd21, []byte("k3RmtuumBP68sINnjhnp4ux4klAriP1Y9ggfhlCqlSapaC96gGBFiM7ehqhUCUxsSWAd_QTQSzTQc3TtXRmq0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (192) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5CJbaphAb31dNB0BXaMkk3RmtuumBP68sINnjhnp4ux4klAriP1Y9ggfhlCqlSapaC96gGBFiM7ehqhUCUxsSWAd_QTQSzTQc3TtXRmq0

buf, ret = client.Read(fd18, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd24 := client.Open("RKQErSV99g.txt", client.O_RDWR|client.O_CREATE)
if(fd24 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd20, []byte("R5IiTiRMKq8A5RsOzuYTRb5Xbnz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) R5IiTiRMKq8A5RsOzuYTRb5Xbnz

ret = client.Seek(fd22, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd21)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd22, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd20)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd23, []byte("Li8TzRDrZWV1R2B70c_7tQM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) Li8TzRDrZWV1R2B70c_7tQM

ret = client.Seek(fd24, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd24)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd18)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd23, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd22, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd25 := client.Open("1Ub5KU0dgY.txt", client.O_RDWR|client.O_CREATE)
if(fd25 < 0) {
  panic("open failed")
}


ret = client.Seek(fd22, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd25, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd23)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd22, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd22)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd25, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd25, []byte("4icVl8n47gaMIFfjsQDyGebDw2RfN1WQTL7p7bQEddur6JAu4Swq_6ASMVdXgnNTU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) 4icVl8n47gaMIFfjsQDyGebDw2RfN1WQTL7p7bQEddur6JAu4Swq_6ASMVdXgnNTU

fd26 := client.Open("CIvzIccnON.txt", client.O_RDWR|client.O_CREATE)
if(fd26 < 0) {
  panic("open failed")
}


ret = client.Seek(fd26, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd26, []byte("_beB86tYCvFlh_IHgYlWnom3glq4MGuJ_JIyQWRalsZmp6CvZl5Q0BNyPZtyJG2GEwTBkyEPkbwLgxVMG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) _beB86tYCvFlh_IHgYlWnom3glq4MGuJ_JIyQWRalsZmp6CvZl5Q0BNyPZtyJG2GEwTBkyEPkbwLgxVMG

fd27 := client.Open("MD4hBZZLtF.txt", client.O_RDWR|client.O_CREATE)
if(fd27 < 0) {
  panic("open failed")
}


fd28 := client.Open("4OIJr9JbY6.txt", client.O_RDWR|client.O_CREATE)
if(fd28 < 0) {
  panic("open failed")
}

//fd state: (81) _beB86tYCvFlh_IHgYlWnom3glq4MGuJ_JIyQWRalsZmp6CvZl5Q0BNyPZtyJG2GEwTBkyEPkbwLgxVMG

ret = client.Write(fd26, []byte("C7lJGxJQmohOt2EKWu7u03UprHjDTCjvkMS73_EtLixZl3e32pDIifluLaZ1zAYWc4eFA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) _beB86tYCvFlh_IHgYlWnom3glq4MGuJ_JIyQWRalsZmp6CvZl5Q0BNyPZtyJG2GEwTBkyEPkbwLgxVMGC7lJGxJQmohOt2EKWu7u03UprHjDTCjvkMS73_EtLixZl3e32pDIifluLaZ1zAYWc4eFA

ret = client.Seek(fd25, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


fd29 := client.Open("WWT0jJRj3Q.txt", client.O_RDWR|client.O_CREATE)
if(fd29 < 0) {
  panic("open failed")
}

//fd state: (150) _beB86tYCvFlh_IHgYlWnom3glq4MGuJ_JIyQWRalsZmp6CvZl5Q0BNyPZtyJG2GEwTBkyEPkbwLgxVMGC7lJGxJQmohOt2EKWu7u03UprHjDTCjvkMS73_EtLixZl3e32pDIifluLaZ1zAYWc4eFA

ret = client.Write(fd26, []byte("_dzj2r5ZuIZMgkA6ps71QV3ShhG4l0ugG2xkVK0T3rxjkLya9MG4KS0_f3nztwpQnzwVAuSjr8SxOnOJ8g2H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (234) _beB86tYCvFlh_IHgYlWnom3glq4MGuJ_JIyQWRalsZmp6CvZl5Q0BNyPZtyJG2GEwTBkyEPkbwLgxVMGC7lJGxJQmohOt2EKWu7u03UprHjDTCjvkMS73_EtLixZl3e32pDIifluLaZ1zAYWc4eFA_dzj2r5ZuIZMgkA6ps71QV3ShhG4l0ugG2xkVK0T3rxjkLya9MG4KS0_f3nztwpQnzwVAuSjr8SxOnOJ8g2H

fd30 := client.Open("U9dz15Nuka.txt", client.O_RDWR|client.O_CREATE)
if(fd30 < 0) {
  panic("open failed")
}


fd31 := client.Open("fvCo57NOiR.txt", client.O_RDWR|client.O_CREATE)
if(fd31 < 0) {
  panic("open failed")
}


ret = client.Close(fd31)
if(ret != 0) {
  panic("close failed")
}

//fd state: (17) 4icVl8n47gaMIFfjsQDyGebDw2RfN1WQTL7p7bQEddur6JAu4Swq_6ASMVdXgnNTU

ret = client.Write(fd25, []byte("RhJvbW7fZFSFbUYiB31weQZw6MoMJo3rM5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) 4icVl8n47gaMIFfjsRhJvbW7fZFSFbUYiB31weQZw6MoMJo3rM5q_6ASMVdXgnNTU
//fd state: (51) 4icVl8n47gaMIFfjsRhJvbW7fZFSFbUYiB31weQZw6MoMJo3rM5q_6ASMVdXgnNTU

ret = client.Write(fd25, []byte("UkrOuLdx7oClCGss16x9p3Igzch2PQeuSBre1rDVyoUIPFX6Ib4OGWy2qJQHZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (112) 4icVl8n47gaMIFfjsRhJvbW7fZFSFbUYiB31weQZw6MoMJo3rM5UkrOuLdx7oClCGss16x9p3Igzch2PQeuSBre1rDVyoUIPFX6Ib4OGWy2qJQHZ

ret = client.Seek(fd28, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd28)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd29, 89, client.SEEK_SET)
if(ret != 89) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 89)
  panic("seek failed")
}

//fd state: (89) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5CJbaphAb31dNB0BXaMkk3RmtuumBP68sINnjhnp4ux4klAriP1Y9ggfhlCqlSapaC96gGBFiM7ehqhUCUxsSWAd_QTQSzTQc3TtXRmq0

ret = client.Write(fd29, []byte("6nrTEZy1tNYRrLTcvw0S3tGCxFBr0ihAbXFvfbQxl7PwJCv9cGszzcv9tlPczIgPIa2thr2ml5qZ_NE9MtNa09WK6tZ4Uud"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (184) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5C6nrTEZy1tNYRrLTcvw0S3tGCxFBr0ihAbXFvfbQxl7PwJCv9cGszzcv9tlPczIgPIa2thr2ml5qZ_NE9MtNa09WK6tZ4Uud3TtXRmq0

ret = client.Close(fd25)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd30, []byte("3f0K0N68iNAQvnxVokEJYVDqgrsUMzmPaF4e8ohnDygMn4RPWnZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) 3f0K0N68iNAQvnxVokEJYVDqgrsUMzmPaF4e8ohnDygMn4RPWnZ

ret = client.Seek(fd30, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


ret = client.Close(fd26)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd30)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd29)
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


fd32 := client.Open("oKf7OnRcNX.txt", client.O_RDWR|client.O_CREATE)
if(fd32 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd32, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd32, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd32, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd32, []byte("DpC4ucmvI6flse_TPlAzjgM9srl7tnIZ_dS8gqo1dPI9Uye7LoRT4f3TGFF3J_cxMdK7zEhhofB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) DpC4ucmvI6flse_TPlAzjgM9srl7tnIZ_dS8gqo1dPI9Uye7LoRT4f3TGFF3J_cxMdK7zEhhofB
//fd state: (75) DpC4ucmvI6flse_TPlAzjgM9srl7tnIZ_dS8gqo1dPI9Uye7LoRT4f3TGFF3J_cxMdK7zEhhofB

ret = client.Write(fd32, []byte("iqW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) DpC4ucmvI6flse_TPlAzjgM9srl7tnIZ_dS8gqo1dPI9Uye7LoRT4f3TGFF3J_cxMdK7zEhhofBiqW

ret = client.Seek(fd32, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


ret = client.Seek(fd32, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


fd33 := client.Open("RKQErSV99g.txt", client.O_RDWR|client.O_CREATE)
if(fd33 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd32, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "srl7tnIZ_dS8gq") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd33, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd33)
if(ret != 0) {
  panic("close failed")
}


fd34 := client.Open("06CKQJmL3s.txt", client.O_RDWR|client.O_CREATE)
if(fd34 < 0) {
  panic("open failed")
}


ret = client.Close(fd34)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd32, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "o1dPI9Uye7LoRT4f3TGFF3J_cxMdK7zEhhofBiqW") {
  panic("wrong data returned")
}


ret = client.Seek(fd32, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd32, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pC4ucmvI6flse_TPlAzjgM9srl7tnIZ_dS8gqo1dPI9Uye") {
  panic("wrong data returned")
}


ret = client.Seek(fd32, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


fd35 := client.Open("B2g001Rmnq.txt", client.O_RDWR|client.O_CREATE)
if(fd35 < 0) {
  panic("open failed")
}


ret = client.Seek(fd35, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd36 := client.Open("vwimLePYvD.txt", client.O_RDWR|client.O_CREATE)
if(fd36 < 0) {
  panic("open failed")
}

//fd state: (23) DpC4ucmvI6flse_TPlAzjgM9srl7tnIZ_dS8gqo1dPI9Uye7LoRT4f3TGFF3J_cxMdK7zEhhofBiqW

ret = client.Write(fd32, []byte("B81ryeAi25Dmi8iTfT5q9xM_2gX3FrZR_KmGwALdmXpJio3hQ28"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) DpC4ucmvI6flse_TPlAzjgMB81ryeAi25Dmi8iTfT5q9xM_2gX3FrZR_KmGwALdmXpJio3hQ28BiqW

buf, ret = client.Read(fd32, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "BiqW") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd32, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd32, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}

//fd state: (45) DpC4ucmvI6flse_TPlAzjgMB81ryeAi25Dmi8iTfT5q9xM_2gX3FrZR_KmGwALdmXpJio3hQ28BiqW

ret = client.Write(fd32, []byte("K8IXLSLIiLBcEErr4xWt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) DpC4ucmvI6flse_TPlAzjgMB81ryeAi25Dmi8iTfT5q9xK8IXLSLIiLBcEErr4xWtpJio3hQ28BiqW

fd37 := client.Open("2IzKiQkfl4.txt", client.O_RDWR|client.O_CREATE)
if(fd37 < 0) {
  panic("open failed")
}


fd38 := client.Open("B2g001Rmnq.txt", client.O_RDWR|client.O_CREATE)
if(fd38 < 0) {
  panic("open failed")
}


fd39 := client.Open("QuIrnKRhbK.txt", client.O_RDWR|client.O_CREATE)
if(fd39 < 0) {
  panic("open failed")
}


ret = client.Close(fd36)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd35, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd35, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd37, []byte("jyK3lQ309HfGDf1msE3oS4Lo44DR18Zt5U4Otcb_XCAMjkPPWoJBC6BhbLqRypInqqFmagAEIAXquzNo8EOU04Mt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) jyK3lQ309HfGDf1msE3oS4Lo44DR18Zt5U4Otcb_XCAMjkPPWoJBC6BhbLqRypInqqFmagAEIAXquzNo8EOU04Mt

fd40 := client.Open("BhfO9EYblN.txt", client.O_RDWR|client.O_CREATE)
if(fd40 < 0) {
  panic("open failed")
}


ret = client.Close(fd32)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd40, []byte("HLID8YNb9OYnBJ2V0WP0Zr8GHdg7cYvtQaNtDd71gAGSdHhuAB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) HLID8YNb9OYnBJ2V0WP0Zr8GHdg7cYvtQaNtDd71gAGSdHhuAB

fd41 := client.Open("qWFkXcSj6v.txt", client.O_RDWR|client.O_CREATE)
if(fd41 < 0) {
  panic("open failed")
}


ret = client.Seek(fd35, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd38, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd39, []byte("jddiw6ph8OKLEIeTf0SUWNcuLUWfJECPm1mbmpLRU8QLbOSyE_4bTIQBDeOAnacCAKqc6C1Imu6WfRplLx6An3FVkt2flTMy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) jddiw6ph8OKLEIeTf0SUWNcuLUWfJECPm1mbmpLRU8QLbOSyE_4bTIQBDeOAnacCAKqc6C1Imu6WfRplLx6An3FVkt2flTMy

buf, ret = client.Read(fd39, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (88) jyK3lQ309HfGDf1msE3oS4Lo44DR18Zt5U4Otcb_XCAMjkPPWoJBC6BhbLqRypInqqFmagAEIAXquzNo8EOU04Mt

ret = client.Write(fd37, []byte("OTbv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) jyK3lQ309HfGDf1msE3oS4Lo44DR18Zt5U4Otcb_XCAMjkPPWoJBC6BhbLqRypInqqFmagAEIAXquzNo8EOU04MtOTbv
//fd state: (0) 

ret = client.Write(fd38, []byte("ACNYW2GVAktiHjQVsXUhJGlqX7LCE7RvBVWaZWfQ_HqvDXIbG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) ACNYW2GVAktiHjQVsXUhJGlqX7LCE7RvBVWaZWfQ_HqvDXIbG

buf, ret = client.Read(fd39, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd40, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd35, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ACNYW2GVAktiHjQVsXUhJGlqX7LCE7R") {
  panic("wrong data returned")
}


fd42 := client.Open("fvCo57NOiR.txt", client.O_RDWR|client.O_CREATE)
if(fd42 < 0) {
  panic("open failed")
}

//fd state: (92) jyK3lQ309HfGDf1msE3oS4Lo44DR18Zt5U4Otcb_XCAMjkPPWoJBC6BhbLqRypInqqFmagAEIAXquzNo8EOU04MtOTbv

ret = client.Write(fd37, []byte("p70QLAWieWsBnk_hRuwv4fLy6W8ThcybWLqRRGGpTtaZ93KF_p2gtKkVUaXpSK7Y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (156) jyK3lQ309HfGDf1msE3oS4Lo44DR18Zt5U4Otcb_XCAMjkPPWoJBC6BhbLqRypInqqFmagAEIAXquzNo8EOU04MtOTbvp70QLAWieWsBnk_hRuwv4fLy6W8ThcybWLqRRGGpTtaZ93KF_p2gtKkVUaXpSK7Y

fd43 := client.Open("L8fQvVgv0C.txt", client.O_RDWR|client.O_CREATE)
if(fd43 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd41, []byte("8nSUR6pZLCuyD1daeWl1XLbOTDp53_i3Yo2xyzeCGsPed"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) 8nSUR6pZLCuyD1daeWl1XLbOTDp53_i3Yo2xyzeCGsPed

ret = client.Close(fd41)
if(ret != 0) {
  panic("close failed")
}


fd44 := client.Open("1Ub5KU0dgY.txt", client.O_RDWR|client.O_CREATE)
if(fd44 < 0) {
  panic("open failed")
}


fd45 := client.Open("5gE05OHOU6.txt", client.O_RDWR|client.O_CREATE)
if(fd45 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd37, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd40)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd37, 144, client.SEEK_SET)
if(ret != 144) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 144)
  panic("seek failed")
}

//fd state: (144) jyK3lQ309HfGDf1msE3oS4Lo44DR18Zt5U4Otcb_XCAMjkPPWoJBC6BhbLqRypInqqFmagAEIAXquzNo8EOU04MtOTbvp70QLAWieWsBnk_hRuwv4fLy6W8ThcybWLqRRGGpTtaZ93KF_p2gtKkVUaXpSK7Y

ret = client.Write(fd37, []byte("xHWBmvGiXP9HwPMc3Rcyk8XqmbkFaNF5D_GUUrHF4MyTJDD1FKYx37rVO3rORwWicIJOme6nXCq7H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (221) jyK3lQ309HfGDf1msE3oS4Lo44DR18Zt5U4Otcb_XCAMjkPPWoJBC6BhbLqRypInqqFmagAEIAXquzNo8EOU04MtOTbvp70QLAWieWsBnk_hRuwv4fLy6W8ThcybWLqRRGGpTtaZ93KF_p2gxHWBmvGiXP9HwPMc3Rcyk8XqmbkFaNF5D_GUUrHF4MyTJDD1FKYx37rVO3rORwWicIJOme6nXCq7H

buf, ret = client.Read(fd44, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd43, []byte("VfcVkawBBhUTaEa4Jh39YBtthNr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) VfcVkawBBhUTaEa4Jh39YBtthNr

ret = client.Close(fd45)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd39, 96, client.SEEK_SET)
if(ret != 96) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 96)
  panic("seek failed")
}


buf, ret = client.Read(fd43, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd38, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd42, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (27) VfcVkawBBhUTaEa4Jh39YBtthNr

ret = client.Write(fd43, []byte("dEvzVo1HNbp4IlfirqFqy2D2XuO3l312L8KzSnZPEF6fB8cVD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) VfcVkawBBhUTaEa4Jh39YBtthNrdEvzVo1HNbp4IlfirqFqy2D2XuO3l312L8KzSnZPEF6fB8cVD

ret = client.Seek(fd38, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd42, []byte("zIWKJzB0NvYF5LbJDn5_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) zIWKJzB0NvYF5LbJDn5_

fd46 := client.Open("4OIJr9JbY6.txt", client.O_RDWR|client.O_CREATE)
if(fd46 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd39, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd38)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd42, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd43, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd47 := client.Open("WWT0jJRj3Q.txt", client.O_RDWR|client.O_CREATE)
if(fd47 < 0) {
  panic("open failed")
}


ret = client.Close(fd42)
if(ret != 0) {
  panic("close failed")
}

//fd state: (31) ACNYW2GVAktiHjQVsXUhJGlqX7LCE7RvBVWaZWfQ_HqvDXIbG

ret = client.Write(fd35, []byte("kHl3VNyPL1RYzkeFDMikfj0DyZTiwkyK5w7_dGNTnt57Ep7sQDteJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) ACNYW2GVAktiHjQVsXUhJGlqX7LCE7RkHl3VNyPL1RYzkeFDMikfj0DyZTiwkyK5w7_dGNTnt57Ep7sQDteJ

ret = client.Seek(fd39, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}


ret = client.Seek(fd35, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Seek(fd37, 197, client.SEEK_SET)
if(ret != 197) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 197)
  panic("seek failed")
}


fd48 := client.Open("fvCo57NOiR.txt", client.O_RDWR|client.O_CREATE)
if(fd48 < 0) {
  panic("open failed")
}


ret = client.Seek(fd35, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


fd49 := client.Open("gSZMWGWvY_.txt", client.O_RDWR|client.O_CREATE)
if(fd49 < 0) {
  panic("open failed")
}


fd50 := client.Open("U9dz15Nuka.txt", client.O_RDWR|client.O_CREATE)
if(fd50 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd49, []byte("7RmhI5_zsWNrZXbPONAipLHvY7e"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) 7RmhI5_zsWNrZXbPONAipLHvY7e

buf, ret = client.Read(fd47, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJ") {
  panic("wrong data returned")
}

//fd state: (76) VfcVkawBBhUTaEa4Jh39YBtthNrdEvzVo1HNbp4IlfirqFqy2D2XuO3l312L8KzSnZPEF6fB8cVD

ret = client.Write(fd43, []byte("hFM1kir3Q7YJQoAFe5Rs28Ur6v3b2OuU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (108) VfcVkawBBhUTaEa4Jh39YBtthNrdEvzVo1HNbp4IlfirqFqy2D2XuO3l312L8KzSnZPEF6fB8cVDhFM1kir3Q7YJQoAFe5Rs28Ur6v3b2OuU

buf, ret = client.Read(fd47, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uOCou4VFCFj90TxG5C6nrTEZy1tNYRrLTcvw0S") {
  panic("wrong data returned")
}


fd51 := client.Open("QBuuZqdkWu.txt", client.O_RDWR|client.O_CREATE)
if(fd51 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd44, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4icVl8n47gaMIFfjsRhJvbW7fZFSFbUYiB31") {
  panic("wrong data returned")
}

//fd state: (0) 3f0K0N68iNAQvnxVokEJYVDqgrsUMzmPaF4e8ohnDygMn4RPWnZ

ret = client.Write(fd50, []byte("drPxMUasDcQUsXF1YnN_wsjfJC_AJEEfuKvfbYMveGrZMC2wGStl4Tlqo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) drPxMUasDcQUsXF1YnN_wsjfJC_AJEEfuKvfbYMveGrZMC2wGStl4Tlqo

ret = client.Close(fd43)
if(ret != 0) {
  panic("close failed")
}

//fd state: (57) drPxMUasDcQUsXF1YnN_wsjfJC_AJEEfuKvfbYMveGrZMC2wGStl4Tlqo

ret = client.Write(fd50, []byte("pCbrB90MwArQkLmdfkBNxkknSmqm811G8Dp1SEZWXIrfmaeY430DRpLP303rBnsJV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) drPxMUasDcQUsXF1YnN_wsjfJC_AJEEfuKvfbYMveGrZMC2wGStl4TlqopCbrB90MwArQkLmdfkBNxkknSmqm811G8Dp1SEZWXIrfmaeY430DRpLP303rBnsJV

ret = client.Seek(fd48, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


fd52 := client.Open("CIvzIccnON.txt", client.O_RDWR|client.O_CREATE)
if(fd52 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd51, []byte("HXG7945uumIC54rzzo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) HXG7945uumIC54rzzo

fd53 := client.Open("b9h5_qN6Gg.txt", client.O_RDWR|client.O_CREATE)
if(fd53 < 0) {
  panic("open failed")
}

//fd state: (109) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5C6nrTEZy1tNYRrLTcvw0S3tGCxFBr0ihAbXFvfbQxl7PwJCv9cGszzcv9tlPczIgPIa2thr2ml5qZ_NE9MtNa09WK6tZ4Uud3TtXRmq0

ret = client.Write(fd47, []byte("e"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5C6nrTEZy1tNYRrLTcvw0SetGCxFBr0ihAbXFvfbQxl7PwJCv9cGszzcv9tlPczIgPIa2thr2ml5qZ_NE9MtNa09WK6tZ4Uud3TtXRmq0
//fd state: (122) drPxMUasDcQUsXF1YnN_wsjfJC_AJEEfuKvfbYMveGrZMC2wGStl4TlqopCbrB90MwArQkLmdfkBNxkknSmqm811G8Dp1SEZWXIrfmaeY430DRpLP303rBnsJV

ret = client.Write(fd50, []byte("QX_Vsva66vCmLCrftna491JuTemQ6N_m2_2R"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (158) drPxMUasDcQUsXF1YnN_wsjfJC_AJEEfuKvfbYMveGrZMC2wGStl4TlqopCbrB90MwArQkLmdfkBNxkknSmqm811G8Dp1SEZWXIrfmaeY430DRpLP303rBnsJVQX_Vsva66vCmLCrftna491JuTemQ6N_m2_2R

fd54 := client.Open("gPKhagUzJN.txt", client.O_RDWR|client.O_CREATE)
if(fd54 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd37, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7rVO3rORwWicIJOme6nXCq7H") {
  panic("wrong data returned")
}


ret = client.Close(fd52)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd44, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


ret = client.Close(fd48)
if(ret != 0) {
  panic("close failed")
}


fd55 := client.Open("06CKQJmL3s.txt", client.O_RDWR|client.O_CREATE)
if(fd55 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd44, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "krOuLdx7oClCGs") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd54, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd54, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (110) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5C6nrTEZy1tNYRrLTcvw0SetGCxFBr0ihAbXFvfbQxl7PwJCv9cGszzcv9tlPczIgPIa2thr2ml5qZ_NE9MtNa09WK6tZ4Uud3TtXRmq0

ret = client.Write(fd47, []byte("4VHBtFsgaKiuPqmfebVyncNcKKohrE1e4L7rK4NMthrh1Kp8Rqa8EDpK95eKFF0Z69iv40ctUWlc9Tv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (189) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5C6nrTEZy1tNYRrLTcvw0Se4VHBtFsgaKiuPqmfebVyncNcKKohrE1e4L7rK4NMthrh1Kp8Rqa8EDpK95eKFF0Z69iv40ctUWlc9Tvmq0
//fd state: (0) 

ret = client.Write(fd53, []byte("qFNObSaPYwOsxRlVTQWZbkrq5AevK_k5XDiHdOWlGQ4Yt45XI_yV44ePGIU45vgapyilbr0fKlzmaGo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) qFNObSaPYwOsxRlVTQWZbkrq5AevK_k5XDiHdOWlGQ4Yt45XI_yV44ePGIU45vgapyilbr0fKlzmaGo

fd56 := client.Open("fvCo57NOiR.txt", client.O_RDWR|client.O_CREATE)
if(fd56 < 0) {
  panic("open failed")
}


ret = client.Seek(fd47, 117, client.SEEK_SET)
if(ret != 117) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 117)
  panic("seek failed")
}

//fd state: (18) HXG7945uumIC54rzzo

ret = client.Write(fd51, []byte("arLlFQO9x_abciJXU_1caO1weIv9u7WJI4G2lFiVW1jJzzjspQ5F_X6a7DWJ8nUOUe9TYdPHsdZGNATR0liIa32MRpdww"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) HXG7945uumIC54rzzoarLlFQO9x_abciJXU_1caO1weIv9u7WJI4G2lFiVW1jJzzjspQ5F_X6a7DWJ8nUOUe9TYdPHsdZGNATR0liIa32MRpdww
//fd state: (0) 

ret = client.Write(fd54, []byte("sghYOub_AD_tEFjbML5t3Ni2JsoQTnfHGkjpSToun7znog9U95UD3Hh2bnciAcMxz1pjmLkntgQT3u_IuKJUQed9MGnviX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) sghYOub_AD_tEFjbML5t3Ni2JsoQTnfHGkjpSToun7znog9U95UD3Hh2bnciAcMxz1pjmLkntgQT3u_IuKJUQed9MGnviX

ret = client.Seek(fd56, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Seek(fd44, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Seek(fd51, 60, client.SEEK_SET)
if(ret != 60) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 60)
  panic("seek failed")
}

//fd state: (117) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5C6nrTEZy1tNYRrLTcvw0Se4VHBtFsgaKiuPqmfebVyncNcKKohrE1e4L7rK4NMthrh1Kp8Rqa8EDpK95eKFF0Z69iv40ctUWlc9Tvmq0

ret = client.Write(fd47, []byte("sBQ1770iHrzT0GVN7wF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5C6nrTEZy1tNYRrLTcvw0Se4VHBtFssBQ1770iHrzT0GVN7wFohrE1e4L7rK4NMthrh1Kp8Rqa8EDpK95eKFF0Z69iv40ctUWlc9Tvmq0

fd57 := client.Open("B2g001Rmnq.txt", client.O_RDWR|client.O_CREATE)
if(fd57 < 0) {
  panic("open failed")
}

//fd state: (78) jddiw6ph8OKLEIeTf0SUWNcuLUWfJECPm1mbmpLRU8QLbOSyE_4bTIQBDeOAnacCAKqc6C1Imu6WfRplLx6An3FVkt2flTMy

ret = client.Write(fd39, []byte("m3pfshChCEa6a8JSCkhe12XVWZL2WUNa_htg0PBAjODZ_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) jddiw6ph8OKLEIeTf0SUWNcuLUWfJECPm1mbmpLRU8QLbOSyE_4bTIQBDeOAnacCAKqc6C1Imu6WfRm3pfshChCEa6a8JSCkhe12XVWZL2WUNa_htg0PBAjODZ_

fd58 := client.Open("nuJaYphaMK.txt", client.O_RDWR|client.O_CREATE)
if(fd58 < 0) {
  panic("open failed")
}

//fd state: (136) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5C6nrTEZy1tNYRrLTcvw0Se4VHBtFssBQ1770iHrzT0GVN7wFohrE1e4L7rK4NMthrh1Kp8Rqa8EDpK95eKFF0Z69iv40ctUWlc9Tvmq0

ret = client.Write(fd47, []byte("dieDTGthOGcGrJeIgrHBGLvTuThIMVNJJW28GYaaA6yg2Qjse5nF2Xg4lWZxYmh0RAfJ9X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (206) jyfu4x3xIe6beEFrW1kqEr8n2ihmQkvYhGxnXYg0cSPtMffGQNKRGJdBZi5QU2BwaR4CfnJuOCou4VFCFj90TxG5C6nrTEZy1tNYRrLTcvw0Se4VHBtFssBQ1770iHrzT0GVN7wFdieDTGthOGcGrJeIgrHBGLvTuThIMVNJJW28GYaaA6yg2Qjse5nF2Xg4lWZxYmh0RAfJ9X

ret = client.Close(fd39)
if(ret != 0) {
  panic("close failed")
}

//fd state: (2) 4icVl8n47gaMIFfjsRhJvbW7fZFSFbUYiB31weQZw6MoMJo3rM5UkrOuLdx7oClCGss16x9p3Igzch2PQeuSBre1rDVyoUIPFX6Ib4OGWy2qJQHZ

ret = client.Write(fd44, []byte("8cwaE9Ww5yED7bMP7_LbxbB6nJF2GT9hA0g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) 4i8cwaE9Ww5yED7bMP7_LbxbB6nJF2GT9hA0geQ

fd59 := client.Open("Ye5EiCI8JZ.txt", client.O_RDWR|client.O_CREATE)
if(fd59 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd49, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (37) 4i8cwaE9Ww5yED7bMP7_LbxbB6nJF2GT9hA0geQ

ret = client.Write(fd44, []byte("hjRy8M_AMTnu8Q2jtCevnUAK9NKi8iDJEoH9cF5ZdOMNyEfhKwkuYN49EhPR5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) 4i8cwaE9Ww5yED7bMP7_LbxbB6nJF2GT9hA0ghjRy8M_AMTnu8Q2jtCevnUAK9NKi8iDJEoH9cF5ZdOMNyEfhKwkuYN49EhPR5

fd60 := client.Open("k4BU33Pz5j.txt", client.O_RDWR|client.O_CREATE)
if(fd60 < 0) {
  panic("open failed")
}


ret = client.Seek(fd59, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd61 := client.Open("l512rwwcEe.txt", client.O_RDWR|client.O_CREATE)
if(fd61 < 0) {
  panic("open failed")
}

//fd state: (11) zIWKJzB0NvYF5LbJDn5_

ret = client.Write(fd56, []byte("K4iIL5rVVxWMlPTJ6Jg_XgX_oQ_cB88q2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) zIWKJzB0NvYK4iIL5rVVxWMlPTJ6Jg_XgX_oQ_cB88q2

fd62 := client.Open("W5JhWI1fwa.txt", client.O_RDWR|client.O_CREATE)
if(fd62 < 0) {
  panic("open failed")
}


ret = client.Seek(fd53, 66, client.SEEK_SET)
if(ret != 66) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 66)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd55, []byte("my6A0DpAVJ0kjSCuKQJsjkwEZk12"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) my6A0DpAVJ0kjSCuKQJsjkwEZk12

buf, ret = client.Read(fd47, 0)
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


buf, ret = client.Read(fd53, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ilbr0fKlzmaGo") {
  panic("wrong data returned")
}


ret = client.Close(fd51)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd46, []byte("4oStUEEL1ACr54To8Q7oAQkmcA3W7un2cNpT7KQITeduetNdpL4y7gAtvxhMgskB4xbTYrSQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) 4oStUEEL1ACr54To8Q7oAQkmcA3W7un2cNpT7KQITeduetNdpL4y7gAtvxhMgskB4xbTYrSQ

fd63 := client.Open("hzplpfjcW7.txt", client.O_RDWR|client.O_CREATE)
if(fd63 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd60, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd35, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GlqX7LCE7RkHl3VNyPL1RYzkeFDMikfj0DyZTiwkyK5w7_dGNTnt57Ep7sQDteJ") {
  panic("wrong data returned")
}


fd64 := client.Open("WAd44LWkY3.txt", client.O_RDWR|client.O_CREATE)
if(fd64 < 0) {
  panic("open failed")
}


ret = client.Close(fd50)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd61, []byte("P4w9eLpwLsvuheh6ilKjBh_hjofCOPNzXJDXrJdVEa40mqdDxcptPO6Lrygmpab0vtfpJjFkpv90BbyQJfkvjQzryaKPiDL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) P4w9eLpwLsvuheh6ilKjBh_hjofCOPNzXJDXrJdVEa40mqdDxcptPO6Lrygmpab0vtfpJjFkpv90BbyQJfkvjQzryaKPiDL

ret = client.Seek(fd60, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd49)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd64, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd65 := client.Open("kzsba3KDvj.txt", client.O_RDWR|client.O_CREATE)
if(fd65 < 0) {
  panic("open failed")
}


fd66 := client.Open("oKf7OnRcNX.txt", client.O_RDWR|client.O_CREATE)
if(fd66 < 0) {
  panic("open failed")
}


ret = client.Close(fd65)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd63, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd56, 7)
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


ret = client.Seek(fd54, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


fd67 := client.Open("cZ909fYTnT.txt", client.O_RDWR|client.O_CREATE)
if(fd67 < 0) {
  panic("open failed")
}


fd68 := client.Open("CWV6d3q1rw.txt", client.O_RDWR|client.O_CREATE)
if(fd68 < 0) {
  panic("open failed")
}

//fd state: (84) ACNYW2GVAktiHjQVsXUhJGlqX7LCE7RkHl3VNyPL1RYzkeFDMikfj0DyZTiwkyK5w7_dGNTnt57Ep7sQDteJ

ret = client.Write(fd35, []byte("Fk6PP6v6KB74DoFezQX1W5lzcLPwvvMYYhjN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) ACNYW2GVAktiHjQVsXUhJGlqX7LCE7RkHl3VNyPL1RYzkeFDMikfj0DyZTiwkyK5w7_dGNTnt57Ep7sQDteJFk6PP6v6KB74DoFezQX1W5lzcLPwvvMYYhjN
//fd state: (98) 4i8cwaE9Ww5yED7bMP7_LbxbB6nJF2GT9hA0ghjRy8M_AMTnu8Q2jtCevnUAK9NKi8iDJEoH9cF5ZdOMNyEfhKwkuYN49EhPR5

ret = client.Write(fd44, []byte("GmBGH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) 4i8cwaE9Ww5yED7bMP7_LbxbB6nJF2GT9hA0ghjRy8M_AMTnu8Q2jtCevnUAK9NKi8iDJEoH9cF5ZdOMNyEfhKwkuYN49EhPR5GmBGH
//fd state: (0) DpC4ucmvI6flse_TPlAzjgMB81ryeAi25Dmi8iTfT5q9xK8IXLSLIiLBcEErr4xWtpJio3hQ28BiqW

ret = client.Write(fd66, []byte("zk4VSPaMMhlU1eb3E4ZJj42v0URga6X4LXROf1XGG8DeFsRzoroKSLY4tZHNRjGk93Rcyna8eqPw3KZsCHkjmgETGu3OfJA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) zk4VSPaMMhlU1eb3E4ZJj42v0URga6X4LXROf1XGG8DeFsRzoroKSLY4tZHNRjGk93Rcyna8eqPw3KZsCHkjmgETGu3OfJA

fd69 := client.Open("JBIlfQrDNE.txt", client.O_RDWR|client.O_CREATE)
if(fd69 < 0) {
  panic("open failed")
}


fd70 := client.Open("OiejMLt5nu.txt", client.O_RDWR|client.O_CREATE)
if(fd70 < 0) {
  panic("open failed")
}

//fd state: (11) sghYOub_AD_tEFjbML5t3Ni2JsoQTnfHGkjpSToun7znog9U95UD3Hh2bnciAcMxz1pjmLkntgQT3u_IuKJUQed9MGnviX

ret = client.Write(fd54, []byte("QiVZdqNVarDzKgT40fcYoVOPBq6Bh4LY4C7FDypxjnM7ssmBtJ99SWqoAD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) sghYOub_AD_QiVZdqNVarDzKgT40fcYoVOPBq6Bh4LY4C7FDypxjnM7ssmBtJ99SWqoADLkntgQT3u_I

ret = client.Close(fd64)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd59, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd60, []byte("I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) I

ret = client.Close(fd53)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd44, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd59)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd69)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd61)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd70, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd54)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) i0Ourj1yFnr8cWMR73NLgwvIQaGgk7scZdPYyTQRWwMuHq0IG6cQiQSwwGU6hYG8nXWZF4GA0a

ret = client.Write(fd67, []byte("OiwsDbeyw4f_PsoERQIDDM4t7QyqweCKdJuNten7Xj1XNnnnM8BbTcpHTWfbqZx6wbDcKl8d3O1JK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) OiwsDbeyw4f_PsoERQIDDM4t7QyqweCKdJuNten7Xj1XNnnnM8BbTcpHTWfbqZx6wbDcKl8d3O1JK

buf, ret = client.Read(fd35, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd70, []byte("hibmICwele8dm6vP86uOdFjMABIGvCq_2KCrUneRsZ4RotMPAu7iMC0AYKIBu23oaboAJAz7G"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) hibmICwele8dm6vP86uOdFjMABIGvCq_2KCrUneRsZ4RotMPAu7iMC0AYKIBu23oaboAJAz7G

buf, ret = client.Read(fd35, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd56)
if(ret != 0) {
  panic("close failed")
}


fd71 := client.Open("nuJaYphaMK.txt", client.O_RDWR|client.O_CREATE)
if(fd71 < 0) {
  panic("open failed")
}


ret = client.Close(fd63)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd60, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 4i8cwaE9Ww5yED7bMP7_LbxbB6nJF2GT9hA0ghjRy8M_AMTnu8Q2jtCevnUAK9NKi8iDJEoH9cF5ZdOMNyEfhKwkuYN49EhPR5GmBGH

ret = client.Write(fd44, []byte("UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOC

buf, ret = client.Read(fd46, 61)
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


buf, ret = client.Read(fd71, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (38) UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOC

ret = client.Write(fd44, []byte("NmVAuq0fS0fvUq9kUy6ihpZcWh4HxuTgZBItmrCPRXy1QHYOngfheL10qJrVw6e2PAljXp_KIsHPev8nDnoNZjyMIsEd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOCNmVAuq0fS0fvUq9kUy6ihpZcWh4HxuTgZBItmrCPRXy1QHYOngfheL10qJrVw6e2PAljXp_KIsHPev8nDnoNZjyMIsEd

fd72 := client.Open("jvjAJ3I1nx.txt", client.O_RDWR|client.O_CREATE)
if(fd72 < 0) {
  panic("open failed")
}


ret = client.Close(fd70)
if(ret != 0) {
  panic("close failed")
}

//fd state: (130) UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOCNmVAuq0fS0fvUq9kUy6ihpZcWh4HxuTgZBItmrCPRXy1QHYOngfheL10qJrVw6e2PAljXp_KIsHPev8nDnoNZjyMIsEd

ret = client.Write(fd44, []byte("GpPisqDiVwboQLNpN4EG_6FWAqUNP0NBJPQnaPw_xGw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (173) UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOCNmVAuq0fS0fvUq9kUy6ihpZcWh4HxuTgZBItmrCPRXy1QHYOngfheL10qJrVw6e2PAljXp_KIsHPev8nDnoNZjyMIsEdGpPisqDiVwboQLNpN4EG_6FWAqUNP0NBJPQnaPw_xGw

buf, ret = client.Read(fd67, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd72, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


fd73 := client.Open("k4BU33Pz5j.txt", client.O_RDWR|client.O_CREATE)
if(fd73 < 0) {
  panic("open failed")
}


fd74 := client.Open("XQb0UKK6hz.txt", client.O_RDWR|client.O_CREATE)
if(fd74 < 0) {
  panic("open failed")
}


ret = client.Seek(fd67, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


buf, ret = client.Read(fd55, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd37)
if(ret != 0) {
  panic("close failed")
}

//fd state: (95) zk4VSPaMMhlU1eb3E4ZJj42v0URga6X4LXROf1XGG8DeFsRzoroKSLY4tZHNRjGk93Rcyna8eqPw3KZsCHkjmgETGu3OfJA

ret = client.Write(fd66, []byte("puMmfHpYsPf7njU9E9AHLpDQ_dy9JEZLChXc1tbwBjRchP4RRtIOgQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (149) zk4VSPaMMhlU1eb3E4ZJj42v0URga6X4LXROf1XGG8DeFsRzoroKSLY4tZHNRjGk93Rcyna8eqPw3KZsCHkjmgETGu3OfJApuMmfHpYsPf7njU9E9AHLpDQ_dy9JEZLChXc1tbwBjRchP4RRtIOgQ

ret = client.Close(fd46)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) R5IiTiRMKq8A5RsOzuYTRb5Xbnz

ret = client.Write(fd68, []byte("i70qnoU43yOqLMse0a6y8_pYMIW9WrTAxwuTWyeDFwm3d674UB9304vErCmDSwA_YRL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) i70qnoU43yOqLMse0a6y8_pYMIW9WrTAxwuTWyeDFwm3d674UB9304vErCmDSwA_YRL

buf, ret = client.Read(fd68, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd60, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd74, []byte("Wbir04f675hPrSFdNlYrPjBJt8RX0YCDbUxpBkGJazfakEXm1bdPaGOPMA8OR7pzjXUsVU60"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) Wbir04f675hPrSFdNlYrPjBJt8RX0YCDbUxpBkGJazfakEXm1bdPaGOPMA8OR7pzjXUsVU60

ret = client.Seek(fd68, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


buf, ret = client.Read(fd44, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (173) UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOCNmVAuq0fS0fvUq9kUy6ihpZcWh4HxuTgZBItmrCPRXy1QHYOngfheL10qJrVw6e2PAljXp_KIsHPev8nDnoNZjyMIsEdGpPisqDiVwboQLNpN4EG_6FWAqUNP0NBJPQnaPw_xGw

ret = client.Write(fd44, []byte("lMDaTK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (179) UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOCNmVAuq0fS0fvUq9kUy6ihpZcWh4HxuTgZBItmrCPRXy1QHYOngfheL10qJrVw6e2PAljXp_KIsHPev8nDnoNZjyMIsEdGpPisqDiVwboQLNpN4EG_6FWAqUNP0NBJPQnaPw_xGwlMDaTK

buf, ret = client.Read(fd73, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "I") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd62, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd75 := client.Open("tyPt0hL1Ex.txt", client.O_RDWR|client.O_CREATE)
if(fd75 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd71, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd60, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd72, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ZMO131xTUdEaPEMMtnDVkveTschypXTwi5V1NgQSO9J3UVFLR3Es9ekd7DB_Eu34LvuzqJngT3K5jEc") {
  panic("wrong data returned")
}


fd76 := client.Open("WclQ9pMzp5.txt", client.O_RDWR|client.O_CREATE)
if(fd76 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd67, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "beyw4f_PsoERQIDDM4t7QyqweCKdJuNten7Xj1XNnnnM8BbTcpHTWfbqZx6wbDc") {
  panic("wrong data returned")
}


ret = client.Seek(fd72, 117, client.SEEK_SET)
if(ret != 117) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 117)
  panic("seek failed")
}


buf, ret = client.Read(fd76, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd72, 91, client.SEEK_SET)
if(ret != 91) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 91)
  panic("seek failed")
}


buf, ret = client.Read(fd44, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd72)
if(ret != 0) {
  panic("close failed")
}

//fd state: (179) UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOCNmVAuq0fS0fvUq9kUy6ihpZcWh4HxuTgZBItmrCPRXy1QHYOngfheL10qJrVw6e2PAljXp_KIsHPev8nDnoNZjyMIsEdGpPisqDiVwboQLNpN4EG_6FWAqUNP0NBJPQnaPw_xGwlMDaTK

ret = client.Write(fd44, []byte("olWBXbbdx9FPj1x3Q3EuBrYxxWoXtcHpXv3WvrbgU3w2SgFYqXj_EczglFtWqkbJZk47F5xC2EbtP9ruItUlLFG5yVYllaL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (274) UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOCNmVAuq0fS0fvUq9kUy6ihpZcWh4HxuTgZBItmrCPRXy1QHYOngfheL10qJrVw6e2PAljXp_KIsHPev8nDnoNZjyMIsEdGpPisqDiVwboQLNpN4EG_6FWAqUNP0NBJPQnaPw_xGwlMDaTKolWBXbbdx9FPj1x3Q3EuBrYxxWoXtcHpXv3WvrbgU3w2SgFYqXj_EczglFtWqkbJZk47F5xC2EbtP9ruItUlLFG5yVYllaL
//fd state: (68) OiwsDbeyw4f_PsoERQIDDM4t7QyqweCKdJuNten7Xj1XNnnnM8BbTcpHTWfbqZx6wbDcKl8d3O1JK

ret = client.Write(fd67, []byte("Zab5bC5si_nHhC9ErU6ScmHgmC8NWUq5Fnf5xUUHglhPT8QvJiYMaPKJxHfqCHNpf87"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) OiwsDbeyw4f_PsoERQIDDM4t7QyqweCKdJuNten7Xj1XNnnnM8BbTcpHTWfbqZx6wbDcZab5bC5si_nHhC9ErU6ScmHgmC8NWUq5Fnf5xUUHglhPT8QvJiYMaPKJxHfqCHNpf87

fd77 := client.Open("qWFkXcSj6v.txt", client.O_RDWR|client.O_CREATE)
if(fd77 < 0) {
  panic("open failed")
}


ret = client.Seek(fd35, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}

//fd state: (72) Wbir04f675hPrSFdNlYrPjBJt8RX0YCDbUxpBkGJazfakEXm1bdPaGOPMA8OR7pzjXUsVU60

ret = client.Write(fd74, []byte("Wg71JqaCMbUrjFAtKKdb0XEVLMnYmdu7kJgMpXlljFt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) Wbir04f675hPrSFdNlYrPjBJt8RX0YCDbUxpBkGJazfakEXm1bdPaGOPMA8OR7pzjXUsVU60Wg71JqaCMbUrjFAtKKdb0XEVLMnYmdu7kJgMpXlljFt

ret = client.Close(fd67)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd66, 132, client.SEEK_SET)
if(ret != 132) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 132)
  panic("seek failed")
}


buf, ret = client.Read(fd77, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8nSUR6pZLCuyD1daeWl1XLbOTDp") {
  panic("wrong data returned")
}

//fd state: (1) I

ret = client.Write(fd73, []byte("dkuvqsi0OPWCd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) Idkuvqsi0OPWCd

ret = client.Close(fd75)
if(ret != 0) {
  panic("close failed")
}

//fd state: (132) zk4VSPaMMhlU1eb3E4ZJj42v0URga6X4LXROf1XGG8DeFsRzoroKSLY4tZHNRjGk93Rcyna8eqPw3KZsCHkjmgETGu3OfJApuMmfHpYsPf7njU9E9AHLpDQ_dy9JEZLChXc1tbwBjRchP4RRtIOgQ

ret = client.Write(fd66, []byte("y_VRg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (137) zk4VSPaMMhlU1eb3E4ZJj42v0URga6X4LXROf1XGG8DeFsRzoroKSLY4tZHNRjGk93Rcyna8eqPw3KZsCHkjmgETGu3OfJApuMmfHpYsPf7njU9E9AHLpDQ_dy9JEZLChXc1y_VRgRchP4RRtIOgQ

fd78 := client.Open("ppwA3wrpYR.txt", client.O_RDWR|client.O_CREATE)
if(fd78 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd62, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd76, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd79 := client.Open("JnokYtvQa6.txt", client.O_RDWR|client.O_CREATE)
if(fd79 < 0) {
  panic("open failed")
}


ret = client.Close(fd35)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd62)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd76, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd78)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd76)
if(ret != 0) {
  panic("close failed")
}


fd80 := client.Open("_J0i0Tbt1H.txt", client.O_RDWR|client.O_CREATE)
if(fd80 < 0) {
  panic("open failed")
}


fd81 := client.Open("UV3aPHJLgf.txt", client.O_RDWR|client.O_CREATE)
if(fd81 < 0) {
  panic("open failed")
}

//fd state: (137) zk4VSPaMMhlU1eb3E4ZJj42v0URga6X4LXROf1XGG8DeFsRzoroKSLY4tZHNRjGk93Rcyna8eqPw3KZsCHkjmgETGu3OfJApuMmfHpYsPf7njU9E9AHLpDQ_dy9JEZLChXc1y_VRgRchP4RRtIOgQ

ret = client.Write(fd66, []byte("QULljPqUDr8fTmoTvYyGZnOVpYkVN0M0RMWyPh_f_TZmSVrW37ODbcva__uUqNJzYfgPTpK2HZTBGQT1JMbnEP37LPtYZHkT4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (234) zk4VSPaMMhlU1eb3E4ZJj42v0URga6X4LXROf1XGG8DeFsRzoroKSLY4tZHNRjGk93Rcyna8eqPw3KZsCHkjmgETGu3OfJApuMmfHpYsPf7njU9E9AHLpDQ_dy9JEZLChXc1y_VRgQULljPqUDr8fTmoTvYyGZnOVpYkVN0M0RMWyPh_f_TZmSVrW37ODbcva__uUqNJzYfgPTpK2HZTBGQT1JMbnEP37LPtYZHkT4
//fd state: (27) 8nSUR6pZLCuyD1daeWl1XLbOTDp53_i3Yo2xyzeCGsPed

ret = client.Write(fd77, []byte("DtXUTHF2Hecol"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) 8nSUR6pZLCuyD1daeWl1XLbOTDpDtXUTHF2HecolGsPed

ret = client.Close(fd55)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd60)
if(ret != 0) {
  panic("close failed")
}


fd82 := client.Open("UQN06sSwTv.txt", client.O_RDWR|client.O_CREATE)
if(fd82 < 0) {
  panic("open failed")
}


fd83 := client.Open("tADgdDlJVJ.txt", client.O_RDWR|client.O_CREATE)
if(fd83 < 0) {
  panic("open failed")
}

//fd state: (40) 8nSUR6pZLCuyD1daeWl1XLbOTDpDtXUTHF2HecolGsPed

ret = client.Write(fd77, []byte("39UgYNEuepZXFuZL7GBy2I9ZXBXUBdZ6fJi2Mq7YkcL40rrhcFt4qRdr0edOcFPm3XhCOj24Vh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) 8nSUR6pZLCuyD1daeWl1XLbOTDpDtXUTHF2Hecol39UgYNEuepZXFuZL7GBy2I9ZXBXUBdZ6fJi2Mq7YkcL40rrhcFt4qRdr0edOcFPm3XhCOj24Vh

ret = client.Close(fd71)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd74, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd79, []byte("FixyoRR7f_8PY_kEAB5grR6rVyxrYE8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) FixyoRR7f_8PY_kEAB5grR6rVyxrYE8

ret = client.Seek(fd80, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd44, 171, client.SEEK_SET)
if(ret != 171) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 171)
  panic("seek failed")
}


ret = client.Seek(fd82, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd77, 86, client.SEEK_SET)
if(ret != 86) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 86)
  panic("seek failed")
}


ret = client.Close(fd66)
if(ret != 0) {
  panic("close failed")
}


fd84 := client.Open("1Ub5KU0dgY.txt", client.O_RDWR|client.O_CREATE)
if(fd84 < 0) {
  panic("open failed")
}


fd85 := client.Open("OsqrDWhO3B.txt", client.O_RDWR|client.O_CREATE)
if(fd85 < 0) {
  panic("open failed")
}


ret = client.Seek(fd77, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}


ret = client.Close(fd79)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd44, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GwlMDaTKolWBXbbdx9FPj1x3Q3EuBrYxxWoXtcHpXv3WvrbgU3w2SgFYqXj_EczglFtWqkbJZk47F5xC2E") {
  panic("wrong data returned")
}


fd86 := client.Open("hzplpfjcW7.txt", client.O_RDWR|client.O_CREATE)
if(fd86 < 0) {
  panic("open failed")
}


ret = client.Seek(fd85, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd86, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd82, []byte("s2ZQLZlOEkiomwXf2hCIDRKCpGnCEBrw46oAQGZ8jm3gKxCuOKbw5Rp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) s2ZQLZlOEkiomwXf2hCIDRKCpGnCEBrw46oAQGZ8jm3gKxCuOKbw5Rp

ret = client.Close(fd82)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd83, []byte("oOwwGYxfv96GKSwEYbFWbGbFu6IDny2wfvxNJe6TdxtvsGzieFONYXltY0BCzdwRSu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) oOwwGYxfv96GKSwEYbFWbGbFu6IDny2wfvxNJe6TdxtvsGzieFONYXltY0BCzdwRSu

ret = client.Seek(fd84, 213, client.SEEK_SET)
if(ret != 213) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 213)
  panic("seek failed")
}


fd87 := client.Open("yXO6X0Yzq0.txt", client.O_RDWR|client.O_CREATE)
if(fd87 < 0) {
  panic("open failed")
}


ret = client.Close(fd68)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd81, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd81, 79)
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

//fd state: (0) 

ret = client.Write(fd86, []byte("_W9lO2a6RaIqj6Ov_b5J4qbdcE6giizUCCLcxFB5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) _W9lO2a6RaIqj6Ov_b5J4qbdcE6giizUCCLcxFB5

ret = client.Close(fd84)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd80)
if(ret != 0) {
  panic("close failed")
}


fd88 := client.Open("ayn9sJsa92.txt", client.O_RDWR|client.O_CREATE)
if(fd88 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd88, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd89 := client.Open("jvjAJ3I1nx.txt", client.O_RDWR|client.O_CREATE)
if(fd89 < 0) {
  panic("open failed")
}


ret = client.Close(fd73)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd44)
if(ret != 0) {
  panic("close failed")
}

//fd state: (30) Wbir04f675hPrSFdNlYrPjBJt8RX0YCDbUxpBkGJazfakEXm1bdPaGOPMA8OR7pzjXUsVU60Wg71JqaCMbUrjFAtKKdb0XEVLMnYmdu7kJgMpXlljFt

ret = client.Write(fd74, []byte("7mA_39oJVCCkT7KO7DZd4376i6WBZNHFYR5JL2rqRbDIg53qLspkS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) Wbir04f675hPrSFdNlYrPjBJt8RX0Y7mA_39oJVCCkT7KO7DZd4376i6WBZNHFYR5JL2rqRbDIg53qLspkSrjFAtKKdb0XEVLMnYmdu7kJgMpXllj
//fd state: (0) 

ret = client.Write(fd88, []byte("Zund_qanIPsywJhtfUCYieEPUUzEL6HDs41SVBIgI52ir4xF_vWS7UZwwehWqXeqIpljpABiZMjj1EbGaY8i9NezWZEErVdsew"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) Zund_qanIPsywJhtfUCYieEPUUzEL6HDs41SVBIgI52ir4xF_vWS7UZwwehWqXeqIpljpABiZMjj1EbGaY8i9NezWZEErVdsew

fd90 := client.Open("RKQErSV99g.txt", client.O_RDWR|client.O_CREATE)
if(fd90 < 0) {
  panic("open failed")
}

//fd state: (0) mBFgH9p9ZMO131xTUdEaPEMMtnDVkveTschypXTwi5V1NgQSO9J3UVFLR3Es9ekd7DB_Eu34LvuzqJngT3K5jEcsGkdpnJTk1rOnl5G_TzrxMn3ITWJ3jv1mwF

ret = client.Write(fd89, []byte("1TpqgMP8fpUgOM5S6HzLR_FMQVYDfaHMokdTgjiSRJH4l4tdZDr6nItjGUKhXDrqzi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) 1TpqgMP8fpUgOM5S6HzLR_FMQVYDfaHMokdTgjiSRJH4l4tdZDr6nItjGUKhXDrqzi

fd91 := client.Open("b9h5_qN6Gg.txt", client.O_RDWR|client.O_CREATE)
if(fd91 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd89, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd87, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}

//fd state: (0) qFNObSaPYwOsxRlVTQWZbkrq5AevK_k5XDiHdOWlGQ4Yt45XI_yV44ePGIU45vgapyilbr0fKlzmaGo

ret = client.Write(fd91, []byte("E_mJxQWHIPJf4TuJpD89dE5NQ8LjEh0Z9svHErR_gjq2rnIhcDAzjunHZxNE07UC9veflOoHzNBSQh9Bj_tkb22vgt6JT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) E_mJxQWHIPJf4TuJpD89dE5NQ8LjEh0Z9svHErR_gjq2rnIhcDAzjunHZxNE07UC9veflOoHzNBSQh9Bj_tkb22vgt6JT

fd92 := client.Open("bKl3zyPTJV.txt", client.O_RDWR|client.O_CREATE)
if(fd92 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd81, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd85, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd93 := client.Open("qWezka1ce2.txt", client.O_RDWR|client.O_CREATE)
if(fd93 < 0) {
  panic("open failed")
}


ret = client.Seek(fd88, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd94 := client.Open("CIvzIccnON.txt", client.O_RDWR|client.O_CREATE)
if(fd94 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd77, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ZXFuZL7GBy2I9ZXBXUBdZ6fJi2Mq7YkcL40rrhcFt4qRdr0edOcFPm3XhCOj24Vh") {
  panic("wrong data returned")
}

//fd state: (40) _W9lO2a6RaIqj6Ov_b5J4qbdcE6giizUCCLcxFB5

ret = client.Write(fd86, []byte("BG7jTQH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) _W9lO2a6RaIqj6Ov_b5J4qbdcE6giizUCCLcxFB5BG7jTQH

ret = client.Close(fd74)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd86, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) _beB86tYCvFlh_IHgYlWnom3glq4MGuJ_JIyQWRalsZmp6CvZl5Q0BNyPZtyJG2GEwTBkyEPkbwLgxVMGC7lJGxJQmohOt2EKWu7u03UprHjDTCjvkMS73_EtLixZl3e32pDIifluLaZ1zAYWc4eFA_dzj2r5ZuIZMgkA6ps71QV3ShhG4l0ugG2xkVK0T3rxjkLya9MG4KS0_f3nztwpQnzwVAuSjr8SxOnOJ8g2H

ret = client.Write(fd94, []byte("IfRErNbk18DEP0ko6c8i0y0MUiryztO5naj62MVmL6CIjDZAhGejzG4g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) IfRErNbk18DEP0ko6c8i0y0MUiryztO5naj62MVmL6CIjDZAhGejzG4g

fd95 := client.Open("jo89Yh5jpX.txt", client.O_RDWR|client.O_CREATE)
if(fd95 < 0) {
  panic("open failed")
}

//fd state: (0) Zund_qanIPsywJhtfUCYieEPUUzEL6HDs41SVBIgI52ir4xF_vWS7UZwwehWqXeqIpljpABiZMjj1EbGaY8i9NezWZEErVdsew

ret = client.Write(fd88, []byte("BhFqbx1JYXoZ8zf90zc909BzIvoiTgBmPcG1VfodYY3N_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) BhFqbx1JYXoZ8zf90zc909BzIvoiTgBmPcG1VfodYY3N_

fd96 := client.Open("g7yDtjfm_R.txt", client.O_RDWR|client.O_CREATE)
if(fd96 < 0) {
  panic("open failed")
}


ret = client.Close(fd81)
if(ret != 0) {
  panic("close failed")
}


fd97 := client.Open("ayn9sJsa92.txt", client.O_RDWR|client.O_CREATE)
if(fd97 < 0) {
  panic("open failed")
}


fd98 := client.Open("hyGZ9P9pd6.txt", client.O_RDWR|client.O_CREATE)
if(fd98 < 0) {
  panic("open failed")
}

//fd state: (66) oOwwGYxfv96GKSwEYbFWbGbFu6IDny2wfvxNJe6TdxtvsGzieFONYXltY0BCzdwRSu

ret = client.Write(fd83, []byte("CyeSO7bi0D6SRofjP0ViRvX8S9yhibdtjpWR4nIjX1zLirkQNu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) oOwwGYxfv96GKSwEYbFWbGbFu6IDny2wfvxNJe6TdxtvsGzieFONYXltY0BCzdwRSuCyeSO7bi0D6SRofjP0ViRvX8S9yhibdtjpWR4nIjX1zLirkQNu

ret = client.Seek(fd93, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd90)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd96, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) AlNb1kNTgq7y2WBwzj72btP8B9vL5Yv3FBbOjEo0q3GoE

ret = client.Write(fd98, []byte("erLn9lkO_QanA3PUt786fNORi0DiUb3jG0TqdMhUz1dblp8bIBGqNOak3Bji4Q_zS72S_SBM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) erLn9lkO_QanA3PUt786fNORi0DiUb3jG0TqdMhUz1dblp8bIBGqNOak3Bji4Q_zS72S_SBM

buf, ret = client.Read(fd88, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd95, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd86, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Seek(fd98, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


fd99 := client.Open("UiyECf7DHD.txt", client.O_RDWR|client.O_CREATE)
if(fd99 < 0) {
  panic("open failed")
}


fd100 := client.Open("bqGWVf_eWs.txt", client.O_RDWR|client.O_CREATE)
if(fd100 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd85, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd87, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}

//fd state: (0) BhFqbx1JYXoZ8zf90zc909BzIvoiTgBmPcG1VfodYY3N_

ret = client.Write(fd97, []byte("7ShJFhMP9mLDXe1T8A4HEmVFVVL33cwd0lCzlfJ8W6V_ZkPQ6oh1xwvRRQ0Keos5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) 7ShJFhMP9mLDXe1T8A4HEmVFVVL33cwd0lCzlfJ8W6V_ZkPQ6oh1xwvRRQ0Keos5

buf, ret = client.Read(fd86, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lO2a6RaIqj6Ov_b5J4qbdcE6giizUCCLcxFB5BG7jTQH") {
  panic("wrong data returned")
}


ret = client.Seek(fd86, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


ret = client.Seek(fd86, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}

//fd state: (45) 7ShJFhMP9mLDXe1T8A4HEmVFVVL33cwd0lCzlfJ8W6V_ZkPQ6oh1xwvRRQ0Keos5

ret = client.Write(fd88, []byte("G9O6wwrvTOE9cYN0L1SbnY1AlGPUDbKEqCax16Bs3BCN3EFfAtaDHSuhLcfN1lw4Atqo249ZqRPqGev1uK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (127) 7ShJFhMP9mLDXe1T8A4HEmVFVVL33cwd0lCzlfJ8W6V_ZG9O6wwrvTOE9cYN0L1SbnY1AlGPUDbKEqCax16Bs3BCN3EFfAtaDHSuhLcfN1lw4Atqo249ZqRPqGev1uK

ret = client.Seek(fd89, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Close(fd85)
if(ret != 0) {
  panic("close failed")
}


fd101 := client.Open("nuJaYphaMK.txt", client.O_RDWR|client.O_CREATE)
if(fd101 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd86, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LcxFB5BG7jTQH") {
  panic("wrong data returned")
}

//fd state: (116) oOwwGYxfv96GKSwEYbFWbGbFu6IDny2wfvxNJe6TdxtvsGzieFONYXltY0BCzdwRSuCyeSO7bi0D6SRofjP0ViRvX8S9yhibdtjpWR4nIjX1zLirkQNu

ret = client.Write(fd83, []byte("udil6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (121) oOwwGYxfv96GKSwEYbFWbGbFu6IDny2wfvxNJe6TdxtvsGzieFONYXltY0BCzdwRSuCyeSO7bi0D6SRofjP0ViRvX8S9yhibdtjpWR4nIjX1zLirkQNuudil6

ret = client.Seek(fd88, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


fd102 := client.Open("XQb0UKK6hz.txt", client.O_RDWR|client.O_CREATE)
if(fd102 < 0) {
  panic("open failed")
}


fd103 := client.Open("1Ub5KU0dgY.txt", client.O_RDWR|client.O_CREATE)
if(fd103 < 0) {
  panic("open failed")
}


ret = client.Seek(fd83, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


ret = client.Close(fd83)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd95, []byte("8chbeI7_d6G2nSYDhaV8MgJPFdJiVhTmZz1V3A9B323XwOhEb5O50lLQHHb_1UYggpCyuAzHxig7YWY3rkh7rwYo7t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) 8chbeI7_d6G2nSYDhaV8MgJPFdJiVhTmZz1V3A9B323XwOhEb5O50lLQHHb_1UYggpCyuAzHxig7YWY3rkh7rwYo7t

ret = client.Close(fd92)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd89)
if(ret != 0) {
  panic("close failed")
}


fd104 := client.Open("QBuuZqdkWu.txt", client.O_RDWR|client.O_CREATE)
if(fd104 < 0) {
  panic("open failed")
}


fd105 := client.Open("WWT0jJRj3Q.txt", client.O_RDWR|client.O_CREATE)
if(fd105 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd93, []byte("P2iefiWcD3G6xpYO2xk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) P2iefiWcD3G6xpYO2xk

ret = client.Seek(fd97, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


fd106 := client.Open("4OIJr9JbY6.txt", client.O_RDWR|client.O_CREATE)
if(fd106 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd93, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd96, []byte("5bbOD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) 5bbOD

fd107 := client.Open("lemrozxydg.txt", client.O_RDWR|client.O_CREATE)
if(fd107 < 0) {
  panic("open failed")
}


fd108 := client.Open("fvWrPZXN19.txt", client.O_RDWR|client.O_CREATE)
if(fd108 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd86, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) UXRomjjMYe6o9sXtiUHWtH5OcXIzkE3uQvdlOCNmVAuq0fS0fvUq9kUy6ihpZcWh4HxuTgZBItmrCPRXy1QHYOngfheL10qJrVw6e2PAljXp_KIsHPev8nDnoNZjyMIsEdGpPisqDiVwboQLNpN4EG_6FWAqUNP0NBJPQnaPw_xGwlMDaTKolWBXbbdx9FPj1x3Q3EuBrYxxWoXtcHpXv3WvrbgU3w2SgFYqXj_EczglFtWqkbJZk47F5xC2EbtP9ruItUlLFG5yVYllaL

ret = client.Write(fd103, []byte("Vsy2T"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) Vsy2T

buf, ret = client.Read(fd100, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd105, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jyfu4x3xIe6b") {
  panic("wrong data returned")
}

//fd state: (0) HXG7945uumIC54rzzoarLlFQO9x_abciJXU_1caO1weIv9u7WJI4G2lFiVW1jJzzjspQ5F_X6a7DWJ8nUOUe9TYdPHsdZGNATR0liIa32MRpdww

ret = client.Write(fd104, []byte("lH76ekacLvGYIBXmny00JNF4ZBBnO8N5gKE0puQUQgPynWOXftloKbYEJkctC_IHbNoNm2CMH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) lH76ekacLvGYIBXmny00JNF4ZBBnO8N5gKE0puQUQgPynWOXftloKbYEJkctC_IHbNoNm2CMH

ret = client.Close(fd104)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd98, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


fd109 := client.Open("N7YQtgMAZW.txt", client.O_RDWR|client.O_CREATE)
if(fd109 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd91, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd110 := client.Open("nuJaYphaMK.txt", client.O_RDWR|client.O_CREATE)
if(fd110 < 0) {
  panic("open failed")
}


ret = client.Close(fd105)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd102, 113, client.SEEK_SET)
if(ret != 113) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 113)
  panic("seek failed")
}


fd111 := client.Open("1Y3DcTJm14.txt", client.O_RDWR|client.O_CREATE)
if(fd111 < 0) {
  panic("open failed")
}


	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}
