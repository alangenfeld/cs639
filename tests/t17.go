
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

fd1 := client.Open("HXhQvlhbTy.txt", client.O_RDWR|client.O_CREATE)
if(fd1 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd1, []byte("85LgfK_KVYoxID6wcX74HzCA4nmPXAmp20NN7wgBkZUuonjcGQ6gW6IBs9b1k9ctfUVZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) 85LgfK_KVYoxID6wcX74HzCA4nmPXAmp20NN7wgBkZUuonjcGQ6gW6IBs9b1k9ctfUVZ

buf, ret = client.Read(fd1, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


ret = client.Close(fd1)
if(ret != 0) {
  panic("close failed")
}


fd2 := client.Open("HXhQvlhbTy.txt", client.O_RDWR|client.O_CREATE)
if(fd2 < 0) {
  panic("open failed")
}

//fd state: (0) 85LgfK_KVYoxID6wcX74HzCA4nmPXAmp20NN7wgBkZUuonjcGQ6gW6IBs9b1k9ctfUVZ

ret = client.Write(fd2, []byte("pn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) pn

ret = client.Close(fd2)
if(ret != 0) {
  panic("close failed")
}


fd3 := client.Open("HXhQvlhbTy.txt", client.O_RDWR|client.O_CREATE)
if(fd3 < 0) {
  panic("open failed")
}


fd4 := client.Open("HXhQvlhbTy.txt", client.O_RDWR|client.O_CREATE)
if(fd4 < 0) {
  panic("open failed")
}


fd5 := client.Open("HXhQvlhbTy.txt", client.O_RDWR|client.O_CREATE)
if(fd5 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd3, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pn") {
  
}


ret = client.Seek(fd5, 0, client.SEEK_SET)
if(ret != 0) {
  panic("seek failed")
}


ret = client.Close(fd3)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd5)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd4)
if(ret != 0) {
  panic("close failed")
}


fd6 := client.Open("HXhQvlhbTy.txt", client.O_RDWR|client.O_CREATE)
if(fd6 < 0) {
  panic("open failed")
}

//fd state: (0) pn

ret = client.Write(fd6, []byte("A3DLBuFczxX88fbc3l3SjxUHveJWoEhDZGuxVe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) A3DLBuFczxX88fbc3l3SjxUHveJWoEhDZGuxVe

fd7 := client.Open("HXhQvlhbTy.txt", client.O_RDWR|client.O_CREATE)
if(fd7 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd7, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "A3DLBuFczxX88fbc3l3SjxUHveJWoEhDZGuxVe") {
  
}


ret = client.Close(fd6)
if(ret != 0) {
  panic("close failed")
}

//fd state: (38) A3DLBuFczxX88fbc3l3SjxUHveJWoEhDZGuxVe

ret = client.Write(fd7, []byte("KtuJiOCli3s6jThPoDd5cQxDX7IpNzaC0JgGPTv76RMiyaDgWrOkBs1isgfKDXfkVJOPMXWTnH1_oN3y2NaKbMKaYu4VwlIJN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) A3DLBuFczxX88fbc3l3SjxUHveJWoEhDZGuxVeKtuJiOCli3s6jThPoDd5cQxDX7IpNzaC0JgGPTv76RMiyaDgWrOkBs1isgfKDXfkVJOPMXWTnH1_oN3y2NaKbMKaYu4VwlIJN

ret = client.Seek(fd7, 47, client.SEEK_SET)
if(ret != 47) {
  panic("seek failed")
}


ret = client.Close(fd7)
if(ret != 0) {
  panic("close failed")
}


fd8 := client.Open("HXhQvlhbTy.txt", client.O_RDWR|client.O_CREATE)
if(fd8 < 0) {
  panic("open failed")
}


fd9 := client.Open("7dFdT969VZ.txt", client.O_RDWR|client.O_CREATE)
if(fd9 < 0) {
  panic("open failed")
}


fd10 := client.Open("P5QDEJ8qZO.txt", client.O_RDWR|client.O_CREATE)
if(fd10 < 0) {
  panic("open failed")
}


ret = client.Close(fd10)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd8, 95, client.SEEK_SET)
if(ret != 95) {
  panic("seek failed")
}


ret = client.Seek(fd8, 128, client.SEEK_SET)
if(ret != 128) {
  panic("seek failed")
}


ret = client.Close(fd8)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd9)
if(ret != 0) {
  panic("close failed")
}


fd11 := client.Open("544bxSGfhD.txt", client.O_RDWR|client.O_CREATE)
if(fd11 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd11, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


buf, ret = client.Read(fd11, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


fd12 := client.Open("0o_ZEbJEP8.txt", client.O_RDWR|client.O_CREATE)
if(fd12 < 0) {
  panic("open failed")
}


ret = client.Close(fd12)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd11, 0, client.SEEK_SET)
if(ret != 0) {
  panic("seek failed")
}


fd13 := client.Open("Ahwh4wqfqj.txt", client.O_RDWR|client.O_CREATE)
if(fd13 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd11, []byte("zbsO45S8pSWbAOyuy0U7B2kxOwmAvRQKotiKb4nk2kGnexn3l5yfVWl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) zbsO45S8pSWbAOyuy0U7B2kxOwmAvRQKotiKb4nk2kGnexn3l5yfVWl

buf, ret = client.Read(fd11, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


buf, ret = client.Read(fd13, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


ret = client.Seek(fd13, 0, client.SEEK_SET)
if(ret != 0) {
  panic("seek failed")
}


fd14 := client.Open("6NhptuG481.txt", client.O_RDWR|client.O_CREATE)
if(fd14 < 0) {
  panic("open failed")
}


ret = client.Seek(fd11, 28, client.SEEK_SET)
if(ret != 28) {
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd14, []byte("1XHWPMwHY3olTKzNuY3KcO4ZajxI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) 1XHWPMwHY3olTKzNuY3KcO4ZajxI

ret = client.Seek(fd14, 27, client.SEEK_SET)
if(ret != 27) {
  panic("seek failed")
}


buf, ret = client.Read(fd14, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "I") {
  
}


ret = client.Seek(fd14, 7, client.SEEK_SET)
if(ret != 7) {
  panic("seek failed")
}


fd15 := client.Open("Uory0Z7I_N.txt", client.O_RDWR|client.O_CREATE)
if(fd15 < 0) {
  panic("open failed")
}


ret = client.Close(fd11)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd14, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "HY3olTKzNuY3KcO4ZajxI") {
  
}


fd16 := client.Open("Uory0Z7I_N.txt", client.O_RDWR|client.O_CREATE)
if(fd16 < 0) {
  panic("open failed")
}


fd17 := client.Open("544bxSGfhD.txt", client.O_RDWR|client.O_CREATE)
if(fd17 < 0) {
  panic("open failed")
}


ret = client.Close(fd16)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd15, 0, client.SEEK_SET)
if(ret != 0) {
  panic("seek failed")
}


ret = client.Close(fd14)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) zbsO45S8pSWbAOyuy0U7B2kxOwmAvRQKotiKb4nk2kGnexn3l5yfVWl

ret = client.Write(fd17, []byte("AIv_B2y3_VGHSUdhmWCtfwfQZfWex3Ma37UmyyPCFYhiSJYNg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) AIv_B2y3_VGHSUdhmWCtfwfQZfWex3Ma37UmyyPCFYhiSJYNg

ret = client.Seek(fd13, 0, client.SEEK_SET)
if(ret != 0) {
  panic("seek failed")
}


ret = client.Seek(fd13, 0, client.SEEK_SET)
if(ret != 0) {
  panic("seek failed")
}


fd18 := client.Open("544bxSGfhD.txt", client.O_RDWR|client.O_CREATE)
if(fd18 < 0) {
  panic("open failed")
}


fd19 := client.Open("uowofdYYhs.txt", client.O_RDWR|client.O_CREATE)
if(fd19 < 0) {
  panic("open failed")
}


ret = client.Seek(fd13, 0, client.SEEK_SET)
if(ret != 0) {
  panic("seek failed")
}


buf, ret = client.Read(fd18, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AIv_B2y3_VGHSUdhmWCtfwfQZfWex3Ma37UmyyPCF") {
  
}

//fd state: (0) 

ret = client.Write(fd15, []byte("5_2vt6q__wFqK_dA5yZW8i7ewRhn2ilWf3xt1pbUhczaoBSW7OoY9qyBM0r15BuROlZZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) 5_2vt6q__wFqK_dA5yZW8i7ewRhn2ilWf3xt1pbUhczaoBSW7OoY9qyBM0r15BuROlZZ

buf, ret = client.Read(fd15, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


fd20 := client.Open("Rul9XOvNRO.txt", client.O_RDWR|client.O_CREATE)
if(fd20 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd19, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}

//fd state: (0) 

ret = client.Write(fd19, []byte("vobZHdWg_8mLvVPAxDzeqO2qLZXuR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) vobZHdWg_8mLvVPAxDzeqO2qLZXuR

buf, ret = client.Read(fd19, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


ret = client.Close(fd13)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd17, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


fd21 := client.Open("fM9piD8ydR.txt", client.O_RDWR|client.O_CREATE)
if(fd21 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd17, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


ret = client.Seek(fd18, 28, client.SEEK_SET)
if(ret != 28) {
  panic("seek failed")
}


buf, ret = client.Read(fd19, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


fd22 := client.Open("HXhQvlhbTy.txt", client.O_RDWR|client.O_CREATE)
if(fd22 < 0) {
  panic("open failed")
}


ret = client.Close(fd18)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd20, []byte("E4g4tcQ8Lv_7bJpxM_avdeiqrJAQf14QKoIQWyJuo60e"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) E4g4tcQ8Lv_7bJpxM_avdeiqrJAQf14QKoIQWyJuo60e

ret = client.Close(fd21)
if(ret != 0) {
  panic("close failed")
}

//fd state: (29) vobZHdWg_8mLvVPAxDzeqO2qLZXuR

ret = client.Write(fd19, []byte("FJkdAufnKa3qkWwigti_LCTK5uTE7W6Fj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) vobZHdWg_8mLvVPAxDzeqO2qLZXuRFJkdAufnKa3qkWwigti_LCTK5uTE7W6Fj

fd23 := client.Open("fM9piD8ydR.txt", client.O_RDWR|client.O_CREATE)
if(fd23 < 0) {
  panic("open failed")
}


ret = client.Seek(fd20, 0, client.SEEK_SET)
if(ret != 0) {
  panic("seek failed")
}

//fd state: (62) vobZHdWg_8mLvVPAxDzeqO2qLZXuRFJkdAufnKa3qkWwigti_LCTK5uTE7W6Fj

ret = client.Write(fd19, []byte("nCcth3o6TwwILT3SlqMKA9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) vobZHdWg_8mLvVPAxDzeqO2qLZXuRFJkdAufnKa3qkWwigti_LCTK5uTE7W6FjnCcth3o6TwwILT3SlqMKA9

ret = client.Seek(fd20, 42, client.SEEK_SET)
if(ret != 42) {
  panic("seek failed")
}


ret = client.Close(fd19)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd17, 28, client.SEEK_SET)
if(ret != 28) {
  panic("seek failed")
}


ret = client.Seek(fd15, 20, client.SEEK_SET)
if(ret != 20) {
  panic("seek failed")
}


ret = client.Close(fd20)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd22, 16, client.SEEK_SET)
if(ret != 16) {
  panic("seek failed")
}


ret = client.Close(fd22)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd17)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd23)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd15, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8i7ewRhn2ilWf3xt1pbUhczaoBSW7OoY9qyBM0r15BuROlZZ") {
  
}


buf, ret = client.Read(fd15, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  
}


ret = client.Close(fd15)
if(ret != 0) {
  panic("close failed")
}


fd24 := client.Open("6NhptuG481.txt", client.O_RDWR|client.O_CREATE)
if(fd24 < 0) {
  panic("open failed")
}


ret = client.Seek(fd24, 8, client.SEEK_SET)
if(ret != 8) {
  panic("seek failed")
}


ret = client.Close(fd24)
if(ret != 0) {
  panic("close failed")
}


fd25 := client.Open("MHp06W4cUz.txt", client.O_RDWR|client.O_CREATE)
if(fd25 < 0) {
  panic("open failed")
}


ret = client.Close(fd25)
if(ret != 0) {
  panic("close failed")
}


	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}
