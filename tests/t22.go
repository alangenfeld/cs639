
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

fd1 := client.Open("/DqNYJpnyj0.txt", client.O_RDWR|client.O_CREATE)
if(fd1 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd1, []byte("iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8isS7VaBxXZon8Fdr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8isS7VaBxXZon8Fdr

buf, ret = client.Read(fd1, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd1)
if(ret != 0) {
  panic("close failed")
}


fd2 := client.Open("/IfH1CmBLH1.txt", client.O_RDWR|client.O_CREATE)
if(fd2 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd2, []byte("ND3tAEKv1AHhS7m7qPNmjst4RJ8eD853T2o4rdtOzGiOcakpOGyYN5K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) ND3tAEKv1AHhS7m7qPNmjst4RJ8eD853T2o4rdtOzGiOcakpOGyYN5K
//fd state: (55) ND3tAEKv1AHhS7m7qPNmjst4RJ8eD853T2o4rdtOzGiOcakpOGyYN5K

ret = client.Write(fd2, []byte("RKDGTovxKHHL6S7S866ruV24a7otrOXn6ihJ3K9BSSdKcck8513ywa2k4o4pBZaIUFbqm4IiuHtwB4RqkeHB3Zm_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (143) ND3tAEKv1AHhS7m7qPNmjst4RJ8eD853T2o4rdtOzGiOcakpOGyYN5KRKDGTovxKHHL6S7S866ruV24a7otrOXn6ihJ3K9BSSdKcck8513ywa2k4o4pBZaIUFbqm4IiuHtwB4RqkeHB3Zm_

ret = client.Seek(fd2, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


ret = client.Seek(fd2, 98, client.SEEK_SET)
if(ret != 98) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 98)
  panic("seek failed")
}


buf, ret = client.Read(fd2, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Kcck8513ywa2k4o4pBZaIUFbqm4IiuHtwB4RqkeHB3Zm_") {
  panic("wrong data returned")
}


ret = client.Close(fd2)
if(ret != 0) {
  panic("close failed")
}


fd3 := client.Open("/IfH1CmBLH1.txt", client.O_RDWR|client.O_CREATE)
if(fd3 < 0) {
  panic("open failed")
}


fd4 := client.Open("/IfH1CmBLH1.txt", client.O_RDWR|client.O_CREATE)
if(fd4 < 0) {
  panic("open failed")
}


fd5 := client.Open("/DqNYJpnyj0.txt", client.O_RDWR|client.O_CREATE)
if(fd5 < 0) {
  panic("open failed")
}


fd6 := client.Open("/IfH1CmBLH1.txt", client.O_RDWR|client.O_CREATE)
if(fd6 < 0) {
  panic("open failed")
}


ret = client.Close(fd5)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd3)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) ND3tAEKv1AHhS7m7qPNmjst4RJ8eD853T2o4rdtOzGiOcakpOGyYN5KRKDGTovxKHHL6S7S866ruV24a7otrOXn6ihJ3K9BSSdKcck8513ywa2k4o4pBZaIUFbqm4IiuHtwB4RqkeHB3Zm_

ret = client.Write(fd6, []byte("vChC9cO32j7QX5FgEFhbArjg9yaRDexF2GvvoWk4sF0x3GABIA3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) vChC9cO32j7QX5FgEFhbArjg9yaRDexF2GvvoWk4sF0x3GABIA3YN5KRKDGTovxKHHL6S7S866ruV24a7otrOXn6ihJ3K9BSSdKcck8513ywa2k4o4pBZaIUFbqm4IiuHtwB4RqkeHB3Zm_

ret = client.Close(fd6)
if(ret != 0) {
  panic("close failed")
}


fd7 := client.Open("/JhxzFpd2v3.txt", client.O_RDWR|client.O_CREATE)
if(fd7 < 0) {
  panic("open failed")
}


fd8 := client.Open("/IXVYf7E6rp.txt", client.O_RDWR|client.O_CREATE)
if(fd8 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd7, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd8)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd7, []byte("hPFqesnWZT_iom"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) hPFqesnWZT_iom

fd9 := client.Open("/JhxzFpd2v3.txt", client.O_RDWR|client.O_CREATE)
if(fd9 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd7, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd9, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


fd10 := client.Open("/XFfR2z8Yx8.txt", client.O_RDWR|client.O_CREATE)
if(fd10 < 0) {
  panic("open failed")
}


fd11 := client.Open("/cej2zXEjS3.txt", client.O_RDWR|client.O_CREATE)
if(fd11 < 0) {
  panic("open failed")
}


fd12 := client.Open("/L3M3DlnPNb.txt", client.O_RDWR|client.O_CREATE)
if(fd12 < 0) {
  panic("open failed")
}


ret = client.Close(fd11)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd7, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Seek(fd7, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


fd13 := client.Open("/L3M3DlnPNb.txt", client.O_RDWR|client.O_CREATE)
if(fd13 < 0) {
  panic("open failed")
}

//fd state: (5) hPFqesnWZT_iom

ret = client.Write(fd9, []byte("ZC_oM7N_tU9SjJAPRB53FraVn2jYe3twtZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) hPFqeZC_oM7N_tU9SjJAPRB53FraVn2jYe3twtZ
//fd state: (0) 

ret = client.Write(fd12, []byte("fFfVoY2FmhwBOGc66UxGwzT40xPuMBvnQl57pLDYdZ76HA2BmxowX9zgHtibc0LWARKsB7iuFNCXjDXn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) fFfVoY2FmhwBOGc66UxGwzT40xPuMBvnQl57pLDYdZ76HA2BmxowX9zgHtibc0LWARKsB7iuFNCXjDXn

ret = client.Close(fd10)
if(ret != 0) {
  panic("close failed")
}


fd14 := client.Open("/DqNYJpnyj0.txt", client.O_RDWR|client.O_CREATE)
if(fd14 < 0) {
  panic("open failed")
}

//fd state: (0) fFfVoY2FmhwBOGc66UxGwzT40xPuMBvnQl57pLDYdZ76HA2BmxowX9zgHtibc0LWARKsB7iuFNCXjDXn

ret = client.Write(fd13, []byte("TWVSw81B0Lkooi5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) TWVSw81B0Lkooi566UxGwzT40xPuMBvnQl57pLDYdZ76HA2BmxowX9zgHtibc0LWARKsB7iuFNCXjDXn

fd15 := client.Open("/8ldA_h4XX0.txt", client.O_RDWR|client.O_CREATE)
if(fd15 < 0) {
  panic("open failed")
}


fd16 := client.Open("/JhxzFpd2v3.txt", client.O_RDWR|client.O_CREATE)
if(fd16 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd7, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "tU9SjJAPRB53FraVn2jYe3twtZ") {
  panic("wrong data returned")
}

//fd state: (80) TWVSw81B0Lkooi566UxGwzT40xPuMBvnQl57pLDYdZ76HA2BmxowX9zgHtibc0LWARKsB7iuFNCXjDXn

ret = client.Write(fd12, []byte("ibnfrHQ4cfRbxz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) TWVSw81B0Lkooi566UxGwzT40xPuMBvnQl57pLDYdZ76HA2BmxowX9zgHtibc0LWARKsB7iuFNCXjDXnibnfrHQ4cfRbxz

buf, ret = client.Read(fd9, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) vChC9cO32j7QX5FgEFhbArjg9yaRDexF2GvvoWk4sF0x3GABIA3YN5KRKDGTovxKHHL6S7S866ruV24a7otrOXn6ihJ3K9BSSdKcck8513ywa2k4o4pBZaIUFbqm4IiuHtwB4RqkeHB3Zm_

ret = client.Write(fd4, []byte("xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0oWk4sF0x3GABIA3YN5KRKDGTovxKHHL6S7S866ruV24a7otrOXn6ihJ3K9BSSdKcck8513ywa2k4o4pBZaIUFbqm4IiuHtwB4RqkeHB3Zm_

buf, ret = client.Read(fd14, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8isS7VaBxXZon8Fdr") {
  panic("wrong data returned")
}


fd17 := client.Open("/JW2zRIl6yS.txt", client.O_RDWR|client.O_CREATE)
if(fd17 < 0) {
  panic("open failed")
}


ret = client.Seek(fd9, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd17, []byte("mdlsvS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) mdlsvS

ret = client.Seek(fd14, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}

//fd state: (36) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0oWk4sF0x3GABIA3YN5KRKDGTovxKHHL6S7S866ruV24a7otrOXn6ihJ3K9BSSdKcck8513ywa2k4o4pBZaIUFbqm4IiuHtwB4RqkeHB3Zm_

ret = client.Write(fd4, []byte("IcY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcY4sF0x3GABIA3YN5KRKDGTovxKHHL6S7S866ruV24a7otrOXn6ihJ3K9BSSdKcck8513ywa2k4o4pBZaIUFbqm4IiuHtwB4RqkeHB3Zm_

fd18 := client.Open("/qksOWSAKgO.txt", client.O_RDWR|client.O_CREATE)
if(fd18 < 0) {
  panic("open failed")
}


ret = client.Seek(fd14, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}

//fd state: (15) hPFqeZC_oM7N_tU9SjJAPRB53FraVn2jYe3twtZ

ret = client.Write(fd9, []byte("4TMQWHS3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) hPFqeZC_oM7N_tU4TMQWHS353FraVn2jYe3twtZ

buf, ret = client.Read(fd17, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (39) hPFqeZC_oM7N_tU4TMQWHS353FraVn2jYe3twtZ

ret = client.Write(fd7, []byte("IOpSL5pITC4h8VJqho0w2uH4AMaVX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) hPFqeZC_oM7N_tU4TMQWHS353FraVn2jYe3twtZIOpSL5pITC4h8VJqho0w2uH4AMaVX

ret = client.Seek(fd16, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}

//fd state: (51) hPFqeZC_oM7N_tU4TMQWHS353FraVn2jYe3twtZIOpSL5pITC4h8VJqho0w2uH4AMaVX

ret = client.Write(fd16, []byte("mtKlxrh2bPGEdR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) hPFqeZC_oM7N_tU4TMQWHS353FraVn2jYe3twtZIOpSL5pITC4hmtKlxrh2bPGEdRaVX

ret = client.Close(fd7)
if(ret != 0) {
  panic("close failed")
}

//fd state: (39) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcY4sF0x3GABIA3YN5KRKDGTovxKHHL6S7S866ruV24a7otrOXn6ihJ3K9BSSdKcck8513ywa2k4o4pBZaIUFbqm4IiuHtwB4RqkeHB3Zm_

ret = client.Write(fd4, []byte("WO8fYcxG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcYWO8fYcxG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_

ret = client.Close(fd13)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd12)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd17, 51)
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


ret = client.Seek(fd18, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (48) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8isS7VaBxXZon8Fdr

ret = client.Write(fd14, []byte("iq2wxFhrV1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1on8Fdr

fd19 := client.Open("/JW2zRIl6yS.txt", client.O_RDWR|client.O_CREATE)
if(fd19 < 0) {
  panic("open failed")
}


ret = client.Close(fd17)
if(ret != 0) {
  panic("close failed")
}


fd20 := client.Open("/b09IPCVYla.txt", client.O_RDWR|client.O_CREATE)
if(fd20 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd20, []byte("vZ_fdlZxnhaOVoPtgFVfwePtyfb5VHzpXZBRdDTa0jp4i1xjl24KWLDLloaOerXs_s5zyeq_BmLkWnOhn4mq7zVDl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) vZ_fdlZxnhaOVoPtgFVfwePtyfb5VHzpXZBRdDTa0jp4i1xjl24KWLDLloaOerXs_s5zyeq_BmLkWnOhn4mq7zVDl

ret = client.Close(fd20)
if(ret != 0) {
  panic("close failed")
}


fd21 := client.Open("/pgOFTKSQaT.txt", client.O_RDWR|client.O_CREATE)
if(fd21 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd15, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd15, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (58) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1on8Fdr

ret = client.Write(fd14, []byte("RuKxeSFOppB1E4nqgWOGB84Q6GWgNPMRjEDnlxdNvlYhRVshD6jceDTNbuY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1RuKxeSFOppB1E4nqgWOGB84Q6GWgNPMRjEDnlxdNvlYhRVshD6jceDTNbuY

ret = client.Close(fd19)
if(ret != 0) {
  panic("close failed")
}


fd22 := client.Open("/LQ9tyeq1jT.txt", client.O_RDWR|client.O_CREATE)
if(fd22 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd15, []byte("VO8qjnBUAIIyU8RkhyuN9S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) VO8qjnBUAIIyU8RkhyuN9S

ret = client.Close(fd22)
if(ret != 0) {
  panic("close failed")
}


fd23 := client.Open("/3HJvawgO69.txt", client.O_RDWR|client.O_CREATE)
if(fd23 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd18, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd14)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd4)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd9, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd23, []byte("VoLI6ui5PWAam6QQbZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) VoLI6ui5PWAam6QQbZ
//fd state: (0) 

ret = client.Write(fd21, []byte("ju6cvEqolfKV24yQF1KHRBOzP7c7yM_1s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) ju6cvEqolfKV24yQF1KHRBOzP7c7yM_1s
//fd state: (27) hPFqeZC_oM7N_tU4TMQWHS353FraVn2jYe3twtZIOpSL5pITC4hmtKlxrh2bPGEdRaVX

ret = client.Write(fd9, []byte("q_GniK2TtaZhJZi4NtJ6ycLr6j705qVpB5jf6NtsvP5E5FE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) hPFqeZC_oM7N_tU4TMQWHS353Frq_GniK2TtaZhJZi4NtJ6ycLr6j705qVpB5jf6NtsvP5E5FE

buf, ret = client.Read(fd18, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd21, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Close(fd21)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd9)
if(ret != 0) {
  panic("close failed")
}


fd24 := client.Open("/O5jLTz9ceU.txt", client.O_RDWR|client.O_CREATE)
if(fd24 < 0) {
  panic("open failed")
}


fd25 := client.Open("/qksOWSAKgO.txt", client.O_RDWR|client.O_CREATE)
if(fd25 < 0) {
  panic("open failed")
}

//fd state: (22) VO8qjnBUAIIyU8RkhyuN9S

ret = client.Write(fd15, []byte("UVOyZdk6IQVYuXoC8ie7_gOU6YTBRYPy0QMfr9yb3oh127Mh36OCh18RRWQchxPo33CHVC54QjQp0Vt9pQbmUdwFNE_X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) VO8qjnBUAIIyU8RkhyuN9SUVOyZdk6IQVYuXoC8ie7_gOU6YTBRYPy0QMfr9yb3oh127Mh36OCh18RRWQchxPo33CHVC54QjQp0Vt9pQbmUdwFNE_X

ret = client.Close(fd18)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd24, []byte("IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJKYxDrMBH2z1Buyn7oLlSUWtYu7AO0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJKYxDrMBH2z1Buyn7oLlSUWtYu7AO0

fd26 := client.Open("/b09IPCVYla.txt", client.O_RDWR|client.O_CREATE)
if(fd26 < 0) {
  panic("open failed")
}

//fd state: (0) vZ_fdlZxnhaOVoPtgFVfwePtyfb5VHzpXZBRdDTa0jp4i1xjl24KWLDLloaOerXs_s5zyeq_BmLkWnOhn4mq7zVDl

ret = client.Write(fd26, []byte("lq1Q7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) lq1Q7lZxnhaOVoPtgFVfwePtyfb5VHzpXZBRdDTa0jp4i1xjl24KWLDLloaOerXs_s5zyeq_BmLkWnOhn4mq7zVDl
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (18) VoLI6ui5PWAam6QQbZ

ret = client.Write(fd23, []byte("DRh7O1bcZj6NPnlGsqaMcwqZ_ykNi3apzBnfWj3eauGAmEg8qRti0iZCZx8_o5OKhnMNjo2v"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) VoLI6ui5PWAam6QQbZDRh7O1bcZj6NPnlGsqaMcwqZ_ykNi3apzBnfWj3eauGAmEg8qRti0iZCZx8_o5OKhnMNjo2v

buf, ret = client.Read(fd15, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd15)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd25)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd26)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd23, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


buf, ret = client.Read(fd23, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bcZj6NPnlGsqaMcw") {
  panic("wrong data returned")
}


fd27 := client.Open("/JhxzFpd2v3.txt", client.O_RDWR|client.O_CREATE)
if(fd27 < 0) {
  panic("open failed")
}


ret = client.Seek(fd23, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


fd28 := client.Open("/B3ip_bKpPr.txt", client.O_RDWR|client.O_CREATE)
if(fd28 < 0) {
  panic("open failed")
}


fd29 := client.Open("/WFrFE5knfM.txt", client.O_RDWR|client.O_CREATE)
if(fd29 < 0) {
  panic("open failed")
}


fd30 := client.Open("/IXVYf7E6rp.txt", client.O_RDWR|client.O_CREATE)
if(fd30 < 0) {
  panic("open failed")
}


fd31 := client.Open("/cej2zXEjS3.txt", client.O_RDWR|client.O_CREATE)
if(fd31 < 0) {
  panic("open failed")
}


fd32 := client.Open("/JhxzFpd2v3.txt", client.O_RDWR|client.O_CREATE)
if(fd32 < 0) {
  panic("open failed")
}


ret = client.Seek(fd27, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


fd33 := client.Open("/NVHRw1JCF2.txt", client.O_RDWR|client.O_CREATE)
if(fd33 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd23, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "aMcwqZ_ykNi3apzBnfWj3eauG") {
  panic("wrong data returned")
}


ret = client.Close(fd24)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd31, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd23, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


buf, ret = client.Read(fd23, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qZ_ykNi3apzBnfWj3eauGAmEg8qRti0iZCZx8_o5OKhnMNjo") {
  panic("wrong data returned")
}


ret = client.Close(fd29)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd23, 88, client.SEEK_SET)
if(ret != 88) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 88)
  panic("seek failed")
}


buf, ret = client.Read(fd28, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd23, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


ret = client.Close(fd30)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd23, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "j6NPnlGsqaMcwqZ_ykNi3apzBnfWj3eauGAmEg8qRti0iZCZx") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd32, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hPFqeZC_oM7N_tU4TMQ") {
  panic("wrong data returned")
}


ret = client.Close(fd27)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd31, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd31, []byte("kCY3PhBtoJWXjm9QLlBr8sicDs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) kCY3PhBtoJWXjm9QLlBr8sicDs
//fd state: (19) hPFqeZC_oM7N_tU4TMQWHS353Frq_GniK2TtaZhJZi4NtJ6ycLr6j705qVpB5jf6NtsvP5E5FE

ret = client.Write(fd32, []byte("npB9lDZTjeFOgifU0s9h9fLa6URw4u0eKkKivsXGrCO1OA_1PP7FNEpwYXyEw_9cZw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) hPFqeZC_oM7N_tU4TMQnpB9lDZTjeFOgifU0s9h9fLa6URw4u0eKkKivsXGrCO1OA_1PP7FNEpwYXyEw_9cZw

ret = client.Close(fd28)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd23, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Close(fd32)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd33)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd31)
if(ret != 0) {
  panic("close failed")
}

//fd state: (10) VoLI6ui5PWAam6QQbZDRh7O1bcZj6NPnlGsqaMcwqZ_ykNi3apzBnfWj3eauGAmEg8qRti0iZCZx8_o5OKhnMNjo2v

ret = client.Write(fd23, []byte("BjZl_oZ04XVAfSvLbFsCaOzU5LhzywAsGRoj9ii916stUqV3idZMRNJPxKOxN_eiQUGQ6BV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) VoLI6ui5PWBjZl_oZ04XVAfSvLbFsCaOzU5LhzywAsGRoj9ii916stUqV3idZMRNJPxKOxN_eiQUGQ6BVKhnMNjo2v

ret = client.Close(fd23)
if(ret != 0) {
  panic("close failed")
}


fd34 := client.Open("/U0rklyAr5K.txt", client.O_RDWR|client.O_CREATE)
if(fd34 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd34, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd34, []byte("qK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) qK

buf, ret = client.Read(fd34, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd34)
if(ret != 0) {
  panic("close failed")
}


fd35 := client.Open("/fg9OT0wd4Q.txt", client.O_RDWR|client.O_CREATE)
if(fd35 < 0) {
  panic("open failed")
}


ret = client.Seek(fd35, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd35, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd35, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd35, 42)
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

ret = client.Write(fd35, []byte("wkF0Nyca171FaIG3SS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) wkF0Nyca171FaIG3SS

buf, ret = client.Read(fd35, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd35, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd35, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd35, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd35, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (18) wkF0Nyca171FaIG3SS

ret = client.Write(fd35, []byte("UI1LvWB14t0r1f8oVFIn3rOvxCRfpUtyAW4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) wkF0Nyca171FaIG3SSUI1LvWB14t0r1f8oVFIn3rOvxCRfpUtyAW4
//fd state: (53) wkF0Nyca171FaIG3SSUI1LvWB14t0r1f8oVFIn3rOvxCRfpUtyAW4

ret = client.Write(fd35, []byte("L4umW8_bT6nUtxr4QLuxreIxjekqWF1gQykcxP_iz692JMh8_67qe8NOuCyMQbxwv1QCyreK_1T0XkCzhzfDn1W_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (141) wkF0Nyca171FaIG3SSUI1LvWB14t0r1f8oVFIn3rOvxCRfpUtyAW4L4umW8_bT6nUtxr4QLuxreIxjekqWF1gQykcxP_iz692JMh8_67qe8NOuCyMQbxwv1QCyreK_1T0XkCzhzfDn1W_

fd36 := client.Open("/3kZ5zPN0Di.txt", client.O_RDWR|client.O_CREATE)
if(fd36 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd35, 97)
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


buf, ret = client.Read(fd36, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd37 := client.Open("/cej2zXEjS3.txt", client.O_RDWR|client.O_CREATE)
if(fd37 < 0) {
  panic("open failed")
}


fd38 := client.Open("/Se7h93SXZ2.txt", client.O_RDWR|client.O_CREATE)
if(fd38 < 0) {
  panic("open failed")
}

//fd state: (0) kCY3PhBtoJWXjm9QLlBr8sicDs

ret = client.Write(fd37, []byte("EOdU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) EOdUPhBtoJWXjm9QLlBr8sicDs

ret = client.Seek(fd35, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


ret = client.Seek(fd35, 65, client.SEEK_SET)
if(ret != 65) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 65)
  panic("seek failed")
}


ret = client.Close(fd36)
if(ret != 0) {
  panic("close failed")
}


fd39 := client.Open("/9MRCgnwRR5.txt", client.O_RDWR|client.O_CREATE)
if(fd39 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd39, []byte("08JYhq39h3y11QqTs7KCYrJSIyLD_zJoZoQaKqCn1GNwGF0HrWtEYUtq6GfCuJI6LfSFBKPHFMXq4dx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) 08JYhq39h3y11QqTs7KCYrJSIyLD_zJoZoQaKqCn1GNwGF0HrWtEYUtq6GfCuJI6LfSFBKPHFMXq4dx

ret = client.Close(fd38)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd39, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Seek(fd37, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


buf, ret = client.Read(fd39, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oZoQaKqCn1GNwGF0") {
  panic("wrong data returned")
}


ret = client.Close(fd35)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd37)
if(ret != 0) {
  panic("close failed")
}

//fd state: (47) 08JYhq39h3y11QqTs7KCYrJSIyLD_zJoZoQaKqCn1GNwGF0HrWtEYUtq6GfCuJI6LfSFBKPHFMXq4dx

ret = client.Write(fd39, []byte("lnkHJGGfqMxtCw7actkOBPTXC86zaOPe55SO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) 08JYhq39h3y11QqTs7KCYrJSIyLD_zJoZoQaKqCn1GNwGF0lnkHJGGfqMxtCw7actkOBPTXC86zaOPe55SO

buf, ret = client.Read(fd39, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd39, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


buf, ret = client.Read(fd39, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "actkOBPTXC86zaOPe55SO") {
  panic("wrong data returned")
}

//fd state: (83) 08JYhq39h3y11QqTs7KCYrJSIyLD_zJoZoQaKqCn1GNwGF0lnkHJGGfqMxtCw7actkOBPTXC86zaOPe55SO

ret = client.Write(fd39, []byte("YX79B0NdfnZUj57Ewj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) 08JYhq39h3y11QqTs7KCYrJSIyLD_zJoZoQaKqCn1GNwGF0lnkHJGGfqMxtCw7actkOBPTXC86zaOPe55SOYX79B0NdfnZUj57Ewj

ret = client.Seek(fd39, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


ret = client.Close(fd39)
if(ret != 0) {
  panic("close failed")
}


fd40 := client.Open("/qksOWSAKgO.txt", client.O_RDWR|client.O_CREATE)
if(fd40 < 0) {
  panic("open failed")
}


ret = client.Close(fd40)
if(ret != 0) {
  panic("close failed")
}


fd41 := client.Open("/8ldA_h4XX0.txt", client.O_RDWR|client.O_CREATE)
if(fd41 < 0) {
  panic("open failed")
}


ret = client.Close(fd41)
if(ret != 0) {
  panic("close failed")
}


fd42 := client.Open("/DqNYJpnyj0.txt", client.O_RDWR|client.O_CREATE)
if(fd42 < 0) {
  panic("open failed")
}


ret = client.Close(fd42)
if(ret != 0) {
  panic("close failed")
}


fd43 := client.Open("/Se7h93SXZ2.txt", client.O_RDWR|client.O_CREATE)
if(fd43 < 0) {
  panic("open failed")
}


ret = client.Close(fd43)
if(ret != 0) {
  panic("close failed")
}


fd44 := client.Open("/a_pne7JkfX.txt", client.O_RDWR|client.O_CREATE)
if(fd44 < 0) {
  panic("open failed")
}


ret = client.Seek(fd44, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd44)
if(ret != 0) {
  panic("close failed")
}


fd45 := client.Open("/yToNGWmTR3.txt", client.O_RDWR|client.O_CREATE)
if(fd45 < 0) {
  panic("open failed")
}


fd46 := client.Open("/pgOFTKSQaT.txt", client.O_RDWR|client.O_CREATE)
if(fd46 < 0) {
  panic("open failed")
}

//fd state: (0) ju6cvEqolfKV24yQF1KHRBOzP7c7yM_1s

ret = client.Write(fd46, []byte("iOg2CUg2OG0EMMyPbtEshSXjgrWcmzDT9r1LgSgKSD5KftMFSDbSafKNv6t1wMmT9FomfrL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (71) iOg2CUg2OG0EMMyPbtEshSXjgrWcmzDT9r1LgSgKSD5KftMFSDbSafKNv6t1wMmT9FomfrL

fd47 := client.Open("/NVHRw1JCF2.txt", client.O_RDWR|client.O_CREATE)
if(fd47 < 0) {
  panic("open failed")
}


ret = client.Seek(fd45, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd45, []byte("9XUO0DB84eCoNX1yjZPfxCHBboR2EV8OKk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) 9XUO0DB84eCoNX1yjZPfxCHBboR2EV8OKk
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd45)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd46, 91)
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


fd48 := client.Open("/5OR5HN2nWE.txt", client.O_RDWR|client.O_CREATE)
if(fd48 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd46, 54)
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


ret = client.Seek(fd48, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd48, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd48, []byte("qEkJLM6MOHnunELVqpNT1ULt3JGHYij2E50PtnJ7MkTrex3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) qEkJLM6MOHnunELVqpNT1ULt3JGHYij2E50PtnJ7MkTrex3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_

buf, ret = client.Read(fd48, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd48, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd48)
if(ret != 0) {
  panic("close failed")
}


fd49 := client.Open("/Se7h93SXZ2.txt", client.O_RDWR|client.O_CREATE)
if(fd49 < 0) {
  panic("open failed")
}


ret = client.Close(fd49)
if(ret != 0) {
  panic("close failed")
}


fd50 := client.Open("/DqNYJpnyj0.txt", client.O_RDWR|client.O_CREATE)
if(fd50 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd50, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1RuKxeSFOppB1E4nqgWOG") {
  panic("wrong data returned")
}

//fd state: (78) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1RuKxeSFOppB1E4nqgWOGB84Q6GWgNPMRjEDnlxdNvlYhRVshD6jceDTNbuY

ret = client.Write(fd50, []byte("kY1sN7ANfIzWZ3cWSzf8PZmWO8wFFzFT5i5l"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1RuKxeSFOppB1E4nqgWOGkY1sN7ANfIzWZ3cWSzf8PZmWO8wFFzFT5i5lbuY
//fd state: (114) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1RuKxeSFOppB1E4nqgWOGkY1sN7ANfIzWZ3cWSzf8PZmWO8wFFzFT5i5lbuY

ret = client.Write(fd50, []byte("pjlRQ92RhkyIWlR4GE0ht8zNlAODw935dQxCS32hwVSq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (158) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1RuKxeSFOppB1E4nqgWOGkY1sN7ANfIzWZ3cWSzf8PZmWO8wFFzFT5i5lpjlRQ92RhkyIWlR4GE0ht8zNlAODw935dQxCS32hwVSq

ret = client.Close(fd50)
if(ret != 0) {
  panic("close failed")
}


fd51 := client.Open("/DFjf9ipCIz.txt", client.O_RDWR|client.O_CREATE)
if(fd51 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd51, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd51, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd52 := client.Open("/XFfR2z8Yx8.txt", client.O_RDWR|client.O_CREATE)
if(fd52 < 0) {
  panic("open failed")
}


ret = client.Seek(fd51, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd52, []byte("aJMHw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) aJMHw

ret = client.Seek(fd51, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd52, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (2) aJMHw

ret = client.Write(fd52, []byte("HvPTyVZtknF63xCxJonUTuzI2SPZCdCzJMd3qRufXTQw96iP9DfWQT4ezGaz7EE3J"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) aJHvPTyVZtknF63xCxJonUTuzI2SPZCdCzJMd3qRufXTQw96iP9DfWQT4ezGaz7EE3J

ret = client.Close(fd52)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd51)
if(ret != 0) {
  panic("close failed")
}


fd53 := client.Open("/5OvPSr9_7r.txt", client.O_RDWR|client.O_CREATE)
if(fd53 < 0) {
  panic("open failed")
}


ret = client.Close(fd53)
if(ret != 0) {
  panic("close failed")
}


fd54 := client.Open("/qksOWSAKgO.txt", client.O_RDWR|client.O_CREATE)
if(fd54 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd54, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) 

ret = client.Seek(fd54, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd55 := client.Open("/B3ip_bKpPr.txt", client.O_RDWR|client.O_CREATE)
if(fd55 < 0) {
  panic("open failed")
}


ret = client.Close(fd54)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd55, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd55, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd55, []byte("KlyiDU8YaGx7FuXI8g5yptikBB8zbi465x4_8xYJ1TTybzpc3vCiszbWb1Psg4QVjgyPCzWU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) KlyiDU8YaGx7FuXI8g5yptikBB8zbi465x4_8xYJ1TTybzpc3vCiszbWb1Psg4QVjgyPCzWU

fd56 := client.Open("/DFjf9ipCIz.txt", client.O_RDWR|client.O_CREATE)
if(fd56 < 0) {
  panic("open failed")
}

//fd state: (72) KlyiDU8YaGx7FuXI8g5yptikBB8zbi465x4_8xYJ1TTybzpc3vCiszbWb1Psg4QVjgyPCzWU

ret = client.Write(fd55, []byte("Y36osxZA3xgRm6cQnKMJw018cBxv30gHoaz_7qf5gIpVych"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (119) KlyiDU8YaGx7FuXI8g5yptikBB8zbi465x4_8xYJ1TTybzpc3vCiszbWb1Psg4QVjgyPCzWUY36osxZA3xgRm6cQnKMJw018cBxv30gHoaz_7qf5gIpVych

fd57 := client.Open("/JhxzFpd2v3.txt", client.O_RDWR|client.O_CREATE)
if(fd57 < 0) {
  panic("open failed")
}


ret = client.Close(fd57)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd55, 101, client.SEEK_SET)
if(ret != 101) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 101)
  panic("seek failed")
}

//fd state: (101) KlyiDU8YaGx7FuXI8g5yptikBB8zbi465x4_8xYJ1TTybzpc3vCiszbWb1Psg4QVjgyPCzWUY36osxZA3xgRm6cQnKMJw018cBxv30gHoaz_7qf5gIpVych

ret = client.Write(fd55, []byte("yYBx0ex0hACT92F0FLm9FiRcdVGYP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) KlyiDU8YaGx7FuXI8g5yptikBB8zbi465x4_8xYJ1TTybzpc3vCiszbWb1Psg4QVjgyPCzWUY36osxZA3xgRm6cQnKMJw018cBxv3yYBx0ex0hACT92F0FLm9FiRcdVGYP

fd58 := client.Open("/DFjf9ipCIz.txt", client.O_RDWR|client.O_CREATE)
if(fd58 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd58, 39)
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


ret = client.Seek(fd55, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


fd59 := client.Open("/98laAaj4bd.txt", client.O_RDWR|client.O_CREATE)
if(fd59 < 0) {
  panic("open failed")
}


ret = client.Close(fd56)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd59, []byte("nCeoi06fxwEtNgDwxchcRQM4mDwphtJ7NKwx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) nCeoi06fxwEtNgDwxchcRQM4mDwphtJ7NKwx

ret = client.Seek(fd59, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


ret = client.Seek(fd59, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}

//fd state: (36) nCeoi06fxwEtNgDwxchcRQM4mDwphtJ7NKwx

ret = client.Write(fd59, []byte("Z1MJ9YGcbmhw0rUSI_30Lnqyoe7HR8vttj_P916SGf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) nCeoi06fxwEtNgDwxchcRQM4mDwphtJ7NKwxZ1MJ9YGcbmhw0rUSI_30Lnqyoe7HR8vttj_P916SGf
//fd state: (78) nCeoi06fxwEtNgDwxchcRQM4mDwphtJ7NKwxZ1MJ9YGcbmhw0rUSI_30Lnqyoe7HR8vttj_P916SGf

ret = client.Write(fd59, []byte("M8sABmCLH4748pggOjTn0zBpX1VzUXgPzFKMguAsgfN33k6oKcYHop_4LHq8N7wXkybG266"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (149) nCeoi06fxwEtNgDwxchcRQM4mDwphtJ7NKwxZ1MJ9YGcbmhw0rUSI_30Lnqyoe7HR8vttj_P916SGfM8sABmCLH4748pggOjTn0zBpX1VzUXgPzFKMguAsgfN33k6oKcYHop_4LHq8N7wXkybG266

fd60 := client.Open("/hD5fZEysbP.txt", client.O_RDWR|client.O_CREATE)
if(fd60 < 0) {
  panic("open failed")
}


ret = client.Seek(fd59, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}

//fd state: (46) KlyiDU8YaGx7FuXI8g5yptikBB8zbi465x4_8xYJ1TTybzpc3vCiszbWb1Psg4QVjgyPCzWUY36osxZA3xgRm6cQnKMJw018cBxv3yYBx0ex0hACT92F0FLm9FiRcdVGYP

ret = client.Write(fd55, []byte("UdSXvJDjpJVOerF7itzjA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) KlyiDU8YaGx7FuXI8g5yptikBB8zbi465x4_8xYJ1TTybzUdSXvJDjpJVOerF7itzjAPCzWUY36osxZA3xgRm6cQnKMJw018cBxv3yYBx0ex0hACT92F0FLm9FiRcdVGYP

buf, ret = client.Read(fd59, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd61 := client.Open("/LQ9tyeq1jT.txt", client.O_RDWR|client.O_CREATE)
if(fd61 < 0) {
  panic("open failed")
}


fd62 := client.Open("/U0rklyAr5K.txt", client.O_RDWR|client.O_CREATE)
if(fd62 < 0) {
  panic("open failed")
}


ret = client.Seek(fd59, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


fd63 := client.Open("/Uvk9UV6e8V.txt", client.O_RDWR|client.O_CREATE)
if(fd63 < 0) {
  panic("open failed")
}


fd64 := client.Open("/tnVmsNSjJy.txt", client.O_RDWR|client.O_CREATE)
if(fd64 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd55, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PCzWUY36osxZA3xgRm6cQnKMJw018cBxv3yYBx0ex0hAC") {
  panic("wrong data returned")
}


ret = client.Close(fd55)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd60)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd63, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd62, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qK") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd62, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd61, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd62)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd64, []byte("E2Li"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) E2Li

ret = client.Seek(fd61, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (16) nCeoi06fxwEtNgDwxchcRQM4mDwphtJ7NKwxZ1MJ9YGcbmhw0rUSI_30Lnqyoe7HR8vttj_P916SGfM8sABmCLH4748pggOjTn0zBpX1VzUXgPzFKMguAsgfN33k6oKcYHop_4LHq8N7wXkybG266

ret = client.Write(fd59, []byte("HiUr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) nCeoi06fxwEtNgDwHiUrRQM4mDwphtJ7NKwxZ1MJ9YGcbmhw0rUSI_30Lnqyoe7HR8vttj_P916SGfM8sABmCLH4748pggOjTn0zBpX1VzUXgPzFKMguAsgfN33k6oKcYHop_4LHq8N7wXkybG266

buf, ret = client.Read(fd63, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd65 := client.Open("/a_pne7JkfX.txt", client.O_RDWR|client.O_CREATE)
if(fd65 < 0) {
  panic("open failed")
}


ret = client.Close(fd61)
if(ret != 0) {
  panic("close failed")
}


fd66 := client.Open("/JhxzFpd2v3.txt", client.O_RDWR|client.O_CREATE)
if(fd66 < 0) {
  panic("open failed")
}


ret = client.Close(fd63)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd59, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "RQM4mDwphtJ7NKw") {
  panic("wrong data returned")
}

//fd state: (35) nCeoi06fxwEtNgDwHiUrRQM4mDwphtJ7NKwxZ1MJ9YGcbmhw0rUSI_30Lnqyoe7HR8vttj_P916SGfM8sABmCLH4748pggOjTn0zBpX1VzUXgPzFKMguAsgfN33k6oKcYHop_4LHq8N7wXkybG266

ret = client.Write(fd59, []byte("qcazjISHtf7TntDMK9SQeFGUJ44DHq6SZuiPsP_gl5zg5HIs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) nCeoi06fxwEtNgDwHiUrRQM4mDwphtJ7NKwqcazjISHtf7TntDMK9SQeFGUJ44DHq6SZuiPsP_gl5zg5HIsmCLH4748pggOjTn0zBpX1VzUXgPzFKMguAsgfN33k6oKcYHop_4LHq8N7wXkybG266
//fd state: (0) 

ret = client.Write(fd65, []byte("TlTaSiKaPyxBM_FNCUpqAwqfZuv5BMZndMkibocqeazXiDGrm8UtnLwTApUFqkludMHpH2pcuuAqLJz4ag1nRARtQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) TlTaSiKaPyxBM_FNCUpqAwqfZuv5BMZndMkibocqeazXiDGrm8UtnLwTApUFqkludMHpH2pcuuAqLJz4ag1nRARtQ

buf, ret = client.Read(fd65, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd66, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hPFqeZC_oM7N_tU4TMQ") {
  panic("wrong data returned")
}


fd67 := client.Open("/N_tIxiNaQn.txt", client.O_RDWR|client.O_CREATE)
if(fd67 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd65, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (19) hPFqeZC_oM7N_tU4TMQnpB9lDZTjeFOgifU0s9h9fLa6URw4u0eKkKivsXGrCO1OA_1PP7FNEpwYXyEw_9cZw

ret = client.Write(fd66, []byte("aYO2mDcphVFBnsOioUWPgWQQq427JjPY0r2U8TPaP2Yjoy2l2ysLiIaNiCT3EX8uEESdPRLCu6XLEE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) hPFqeZC_oM7N_tU4TMQaYO2mDcphVFBnsOioUWPgWQQq427JjPY0r2U8TPaP2Yjoy2l2ysLiIaNiCT3EX8uEESdPRLCu6XLEE
//fd state: (0) 

ret = client.Write(fd67, []byte("1N_HZXae4g6riiMXZiohWJNDDu66_Mmm6zBYBdZRxnow0P4HV0LGY3qTKKXCIreHFosYTMjSVDMyAb_fnv5C6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) 1N_HZXae4g6riiMXZiohWJNDDu66_Mmm6zBYBdZRxnow0P4HV0LGY3qTKKXCIreHFosYTMjSVDMyAb_fnv5C6

buf, ret = client.Read(fd59, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "mCLH4748pggOjTn0zBpX1VzUXgPzFKMguAsgfN33k6oKcYHop_4LHq8N7wX") {
  panic("wrong data returned")
}

//fd state: (97) hPFqeZC_oM7N_tU4TMQaYO2mDcphVFBnsOioUWPgWQQq427JjPY0r2U8TPaP2Yjoy2l2ysLiIaNiCT3EX8uEESdPRLCu6XLEE

ret = client.Write(fd66, []byte("cDctbwFU0DwGusn4mHqNejskKR0xJEqg1P3VzPLSppWlWWsdUGcAC0jN61mXxqHa3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (162) hPFqeZC_oM7N_tU4TMQaYO2mDcphVFBnsOioUWPgWQQq427JjPY0r2U8TPaP2Yjoy2l2ysLiIaNiCT3EX8uEESdPRLCu6XLEEcDctbwFU0DwGusn4mHqNejskKR0xJEqg1P3VzPLSppWlWWsdUGcAC0jN61mXxqHa3

buf, ret = client.Read(fd64, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd65)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd67)
if(ret != 0) {
  panic("close failed")
}


fd68 := client.Open("/gQ_ImV6qgZ.txt", client.O_RDWR|client.O_CREATE)
if(fd68 < 0) {
  panic("open failed")
}


ret = client.Seek(fd64, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd59, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "kybG266") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd68, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd66)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd68, []byte("BbbdVFZRGyCr_2GEDBUp6Kn6c74EzzmKk4hij01kaq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) BbbdVFZRGyCr_2GEDBUp6Kn6c74EzzmKk4hij01kaq

fd69 := client.Open("/9MRCgnwRR5.txt", client.O_RDWR|client.O_CREATE)
if(fd69 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd59, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd68, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


buf, ret = client.Read(fd68, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "FZRG") {
  panic("wrong data returned")
}


fd70 := client.Open("/fg9OT0wd4Q.txt", client.O_RDWR|client.O_CREATE)
if(fd70 < 0) {
  panic("open failed")
}

//fd state: (0) E2Li

ret = client.Write(fd64, []byte("quffD422E9vksKPRR6IynoZ3j3nJc7CTQuXxHUFc9s0DBVGSALBHyGgwAstibrlUzhbksYaN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) quffD422E9vksKPRR6IynoZ3j3nJc7CTQuXxHUFc9s0DBVGSALBHyGgwAstibrlUzhbksYaN

fd71 := client.Open("/5nEDDggz8U.txt", client.O_RDWR|client.O_CREATE)
if(fd71 < 0) {
  panic("open failed")
}


ret = client.Close(fd64)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd71, []byte("uV9jOsW95BFP1d0k_axYK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) uV9jOsW95BFP1d0k_axYK

fd72 := client.Open("/0n5N6AKRi8.txt", client.O_RDWR|client.O_CREATE)
if(fd72 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd68, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yCr_2GEDBUp6Kn6c") {
  panic("wrong data returned")
}


fd73 := client.Open("/9MRCgnwRR5.txt", client.O_RDWR|client.O_CREATE)
if(fd73 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd73, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "08JYhq39h3y11QqTs7KCYrJSIyLD_zJoZoQaKqC") {
  panic("wrong data returned")
}


fd74 := client.Open("/7nzhcLMCES.txt", client.O_RDWR|client.O_CREATE)
if(fd74 < 0) {
  panic("open failed")
}


ret = client.Close(fd69)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd73)
if(ret != 0) {
  panic("close failed")
}

//fd state: (21) uV9jOsW95BFP1d0k_axYK

ret = client.Write(fd71, []byte("ftzBQ44_JEhTi3Xg_4UqMqnjRKZYxMUOC9lbIlgK7XN1eTpgfr2YzBVU63LBieapKPmVe0GKFYq1zLHg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) uV9jOsW95BFP1d0k_axYKftzBQ44_JEhTi3Xg_4UqMqnjRKZYxMUOC9lbIlgK7XN1eTpgfr2YzBVU63LBieapKPmVe0GKFYq1zLHg
//fd state: (0) wkF0Nyca171FaIG3SSUI1LvWB14t0r1f8oVFIn3rOvxCRfpUtyAW4L4umW8_bT6nUtxr4QLuxreIxjekqWF1gQykcxP_iz692JMh8_67qe8NOuCyMQbxwv1QCyreK_1T0XkCzhzfDn1W_

ret = client.Write(fd70, []byte("hWDv017QOTVjvSrS6Vb2TOHPHSULDO89oGzUoHnfnBUDHYKTvUBoC7dxkduRqXiTwqTOzsKvEO85a97GsAPsFUCou_fN9cX6pN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) hWDv017QOTVjvSrS6Vb2TOHPHSULDO89oGzUoHnfnBUDHYKTvUBoC7dxkduRqXiTwqTOzsKvEO85a97GsAPsFUCou_fN9cX6pNMh8_67qe8NOuCyMQbxwv1QCyreK_1T0XkCzhzfDn1W_

buf, ret = client.Read(fd71, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd75 := client.Open("/JhxzFpd2v3.txt", client.O_RDWR|client.O_CREATE)
if(fd75 < 0) {
  panic("open failed")
}


ret = client.Seek(fd75, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


ret = client.Close(fd70)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd72, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd74, []byte("4jazZ9d"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) 4jazZ9d

fd76 := client.Open("/5OR5HN2nWE.txt", client.O_RDWR|client.O_CREATE)
if(fd76 < 0) {
  panic("open failed")
}


ret = client.Seek(fd68, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


ret = client.Close(fd59)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd68)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd71, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


ret = client.Close(fd72)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd75, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


buf, ret = client.Read(fd74, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) qEkJLM6MOHnunELVqpNT1ULt3JGHYij2E50PtnJ7MkTrex3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_

ret = client.Write(fd76, []byte("cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_

buf, ret = client.Read(fd71, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LBieapKPmVe0GKFYq1zLHg") {
  panic("wrong data returned")
}


ret = client.Seek(fd74, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (3) 4jazZ9d

ret = client.Write(fd74, []byte("aW5mgM2pYs0nDFbscFnCUztfrnFcu5ytJ5hUEQKA5sQSfykrhdSuLd80Z71orDQIYaxACv2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) 4jaaW5mgM2pYs0nDFbscFnCUztfrnFcu5ytJ5hUEQKA5sQSfykrhdSuLd80Z71orDQIYaxACv2

fd77 := client.Open("/3kZ5zPN0Di.txt", client.O_RDWR|client.O_CREATE)
if(fd77 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd75, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7N_tU4TMQaYO2mDcphVFBnsOioUWPgWQQq427JjPY0r2U8TP") {
  panic("wrong data returned")
}


ret = client.Seek(fd71, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


ret = client.Seek(fd74, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd75)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd71, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


ret = client.Seek(fd71, 99, client.SEEK_SET)
if(ret != 99) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 99)
  panic("seek failed")
}


buf, ret = client.Read(fd77, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd78 := client.Open("/PSn4daw9hM.txt", client.O_RDWR|client.O_CREATE)
if(fd78 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd74, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "s0nDFbscFnCUztfrnFcu5ytJ5hUEQKA5sQSfykrhdSuLd80Z71orDQIYaxACv2") {
  panic("wrong data returned")
}


fd79 := client.Open("/fg9OT0wd4Q.txt", client.O_RDWR|client.O_CREATE)
if(fd79 < 0) {
  panic("open failed")
}


ret = client.Seek(fd74, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


ret = client.Close(fd71)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd76, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "3yBRBf8iutTfQjV_BBZKqO") {
  panic("wrong data returned")
}


ret = client.Seek(fd77, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd80 := client.Open("/XFfR2z8Yx8.txt", client.O_RDWR|client.O_CREATE)
if(fd80 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd77, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd79, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hWDv") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd78, []byte("vNHeIG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) vNHeIG

ret = client.Close(fd74)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd76, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


fd81 := client.Open("/gMb60G3hQV.txt", client.O_RDWR|client.O_CREATE)
if(fd81 < 0) {
  panic("open failed")
}


fd82 := client.Open("/qksOWSAKgO.txt", client.O_RDWR|client.O_CREATE)
if(fd82 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd81, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd79, 100, client.SEEK_SET)
if(ret != 100) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 100)
  panic("seek failed")
}


buf, ret = client.Read(fd77, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) aJHvPTyVZtknF63xCxJonUTuzI2SPZCdCzJMd3qRufXTQw96iP9DfWQT4ezGaz7EE3J

ret = client.Write(fd80, []byte("F34DAk2KuCug8XuTQsx6CCyRVjQooGji1QUU0NGioxMQIbowQQupHt4N2xPBo5nX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) F34DAk2KuCug8XuTQsx6CCyRVjQooGji1QUU0NGioxMQIbowQQupHt4N2xPBo5nXE3J

fd83 := client.Open("/4NexkYmnGj.txt", client.O_RDWR|client.O_CREATE)
if(fd83 < 0) {
  panic("open failed")
}


ret = client.Close(fd81)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd83, []byte("8SWfklF5EBz6iAHz_JL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) 8SWfklF5EBz6iAHz_JL

ret = client.Seek(fd82, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd82)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd83, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd79, 73, client.SEEK_SET)
if(ret != 73) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 73)
  panic("seek failed")
}


buf, ret = client.Read(fd80, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "E3J") {
  panic("wrong data returned")
}


ret = client.Close(fd83)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd79, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "O85a97GsAPsFUCou_fN9cX6pNMh8") {
  panic("wrong data returned")
}


fd84 := client.Open("/s7jifbhgZq.txt", client.O_RDWR|client.O_CREATE)
if(fd84 < 0) {
  panic("open failed")
}

//fd state: (6) vNHeIG

ret = client.Write(fd78, []byte("8AE0OXRRuez6GvUd7P58QxF76qelHROmPDoa8pdXX4JVMrq3ZGHiMeMz05ShUf6A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) vNHeIG8AE0OXRRuez6GvUd7P58QxF76qelHROmPDoa8pdXX4JVMrq3ZGHiMeMz05ShUf6A

ret = client.Close(fd76)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd77, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd85 := client.Open("/2pzPsiHvtK.txt", client.O_RDWR|client.O_CREATE)
if(fd85 < 0) {
  panic("open failed")
}


ret = client.Close(fd84)
if(ret != 0) {
  panic("close failed")
}


fd86 := client.Open("/HuKvzU3AWd.txt", client.O_RDWR|client.O_CREATE)
if(fd86 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd78, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd79, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


fd87 := client.Open("/ufMjKnKzWK.txt", client.O_RDWR|client.O_CREATE)
if(fd87 < 0) {
  panic("open failed")
}


ret = client.Close(fd78)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd77)
if(ret != 0) {
  panic("close failed")
}


fd88 := client.Open("/H3B8paSDlB.txt", client.O_RDWR|client.O_CREATE)
if(fd88 < 0) {
  panic("open failed")
}


fd89 := client.Open("/3kZ5zPN0Di.txt", client.O_RDWR|client.O_CREATE)
if(fd89 < 0) {
  panic("open failed")
}


fd90 := client.Open("/f38bZ7Efc2.txt", client.O_RDWR|client.O_CREATE)
if(fd90 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd90, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd87, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd79)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd86, []byte("E_EwwJocqB9RWGMd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (16) E_EwwJocqB9RWGMd

buf, ret = client.Read(fd80, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd85, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd86, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Close(fd90)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd80)
if(ret != 0) {
  panic("close failed")
}

//fd state: (8) E_EwwJocqB9RWGMd

ret = client.Write(fd86, []byte("FAGCQqqWDkasFYseDJafudrUghYmeem_R7qdNolV2Nv0OjoQGg9POekC27mqvc3f7gVaAvgPbIqyv4Iab"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) E_EwwJocFAGCQqqWDkasFYseDJafudrUghYmeem_R7qdNolV2Nv0OjoQGg9POekC27mqvc3f7gVaAvgPbIqyv4Iab
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd87)
if(ret != 0) {
  panic("close failed")
}


fd91 := client.Open("/a_pne7JkfX.txt", client.O_RDWR|client.O_CREATE)
if(fd91 < 0) {
  panic("open failed")
}


ret = client.Seek(fd88, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd89, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd92 := client.Open("/3Xuls7KVgK.txt", client.O_RDWR|client.O_CREATE)
if(fd92 < 0) {
  panic("open failed")
}


ret = client.Seek(fd91, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}

//fd state: (89) E_EwwJocFAGCQqqWDkasFYseDJafudrUghYmeem_R7qdNolV2Nv0OjoQGg9POekC27mqvc3f7gVaAvgPbIqyv4Iab

ret = client.Write(fd86, []byte("94FcQ2b0ZWRMHny6xV4tAsBZU3ULoDaN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (121) E_EwwJocFAGCQqqWDkasFYseDJafudrUghYmeem_R7qdNolV2Nv0OjoQGg9POekC27mqvc3f7gVaAvgPbIqyv4Iab94FcQ2b0ZWRMHny6xV4tAsBZU3ULoDaN

ret = client.Close(fd85)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd89, []byte("HbXLOmMQ9_a9Jsj4Rqi8E8dJQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) HbXLOmMQ9_a9Jsj4Rqi8E8dJQ

ret = client.Close(fd91)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd89)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd86, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (121) E_EwwJocFAGCQqqWDkasFYseDJafudrUghYmeem_R7qdNolV2Nv0OjoQGg9POekC27mqvc3f7gVaAvgPbIqyv4Iab94FcQ2b0ZWRMHny6xV4tAsBZU3ULoDaN

ret = client.Write(fd86, []byte("9IDlgH3D9CD2nloBM8qBSHCj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (145) E_EwwJocFAGCQqqWDkasFYseDJafudrUghYmeem_R7qdNolV2Nv0OjoQGg9POekC27mqvc3f7gVaAvgPbIqyv4Iab94FcQ2b0ZWRMHny6xV4tAsBZU3ULoDaN9IDlgH3D9CD2nloBM8qBSHCj

ret = client.Close(fd86)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd88, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd92, []byte("IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzguBSfl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzguBSfl

ret = client.Close(fd92)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd88, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd88, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd88, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd88, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd88, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd88, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd88, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd88, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd88)
if(ret != 0) {
  panic("close failed")
}


fd93 := client.Open("/80lh0X6OA1.txt", client.O_RDWR|client.O_CREATE)
if(fd93 < 0) {
  panic("open failed")
}


ret = client.Close(fd93)
if(ret != 0) {
  panic("close failed")
}


fd94 := client.Open("/su0syYXcBE.txt", client.O_RDWR|client.O_CREATE)
if(fd94 < 0) {
  panic("open failed")
}


ret = client.Close(fd94)
if(ret != 0) {
  panic("close failed")
}


fd95 := client.Open("/kwPisEKdpD.txt", client.O_RDWR|client.O_CREATE)
if(fd95 < 0) {
  panic("open failed")
}


ret = client.Close(fd95)
if(ret != 0) {
  panic("close failed")
}


fd96 := client.Open("/0vSfmI0PuM.txt", client.O_RDWR|client.O_CREATE)
if(fd96 < 0) {
  panic("open failed")
}


ret = client.Seek(fd96, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd96, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd96, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd96, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd96, []byte("56UEwawz6JUxykZcfsC_9DOGqPpPoYQHeSqvBarJInNKte2GmgxmhKg5jczgKyB3t8G"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) 56UEwawz6JUxykZcfsC_9DOGqPpPoYQHeSqvBarJInNKte2GmgxmhKg5jczgKyB3t8G
//fd state: (67) 56UEwawz6JUxykZcfsC_9DOGqPpPoYQHeSqvBarJInNKte2GmgxmhKg5jczgKyB3t8G

ret = client.Write(fd96, []byte("cZYFvIz4theomeF1wiSa1i_BiudfjCKy07c7ZiW79W2ahE3dJT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) 56UEwawz6JUxykZcfsC_9DOGqPpPoYQHeSqvBarJInNKte2GmgxmhKg5jczgKyB3t8GcZYFvIz4theomeF1wiSa1i_BiudfjCKy07c7ZiW79W2ahE3dJT

ret = client.Seek(fd96, 95, client.SEEK_SET)
if(ret != 95) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 95)
  panic("seek failed")
}


fd97 := client.Open("/o3Jy8jBnOV.txt", client.O_RDWR|client.O_CREATE)
if(fd97 < 0) {
  panic("open failed")
}


fd98 := client.Open("/HA0fbMGRsh.txt", client.O_RDWR|client.O_CREATE)
if(fd98 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd98, []byte("Nub4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) Nub4

fd99 := client.Open("/3kZ5zPN0Di.txt", client.O_RDWR|client.O_CREATE)
if(fd99 < 0) {
  panic("open failed")
}


ret = client.Close(fd97)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd98, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd100 := client.Open("/VznqypOi9t.txt", client.O_RDWR|client.O_CREATE)
if(fd100 < 0) {
  panic("open failed")
}

//fd state: (95) 56UEwawz6JUxykZcfsC_9DOGqPpPoYQHeSqvBarJInNKte2GmgxmhKg5jczgKyB3t8GcZYFvIz4theomeF1wiSa1i_BiudfjCKy07c7ZiW79W2ahE3dJT

ret = client.Write(fd96, []byte("GOqMa2MKRt1nSRZh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) 56UEwawz6JUxykZcfsC_9DOGqPpPoYQHeSqvBarJInNKte2GmgxmhKg5jczgKyB3t8GcZYFvIz4theomeF1wiSa1i_BiudfGOqMa2MKRt1nSRZhhE3dJT

fd101 := client.Open("/fTKEpwYowY.txt", client.O_RDWR|client.O_CREATE)
if(fd101 < 0) {
  panic("open failed")
}


ret = client.Close(fd99)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd98, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd100, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd96, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}


ret = client.Seek(fd100, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (4) Nub4

ret = client.Write(fd98, []byte("zb261bmPM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) Nub4zb261bmPM

ret = client.Close(fd96)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd100, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd102 := client.Open("/f38bZ7Efc2.txt", client.O_RDWR|client.O_CREATE)
if(fd102 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd101, 46)
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


buf, ret = client.Read(fd98, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd98, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd102, []byte("K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) K

ret = client.Seek(fd102, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd100, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd102, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "K") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd100, []byte("19pzCz6ktdSCrSveYuFE84LKWmQswrb4XoEud95X1dyXGj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) 19pzCz6ktdSCrSveYuFE84LKWmQswrb4XoEud95X1dyXGj

ret = client.Close(fd100)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd98)
if(ret != 0) {
  panic("close failed")
}


fd103 := client.Open("/PSn4daw9hM.txt", client.O_RDWR|client.O_CREATE)
if(fd103 < 0) {
  panic("open failed")
}


ret = client.Close(fd102)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd103)
if(ret != 0) {
  panic("close failed")
}


fd104 := client.Open("/0n5N6AKRi8.txt", client.O_RDWR|client.O_CREATE)
if(fd104 < 0) {
  panic("open failed")
}


ret = client.Seek(fd104, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd104)
if(ret != 0) {
  panic("close failed")
}


fd105 := client.Open("/LQ9tyeq1jT.txt", client.O_RDWR|client.O_CREATE)
if(fd105 < 0) {
  panic("open failed")
}


ret = client.Close(fd105)
if(ret != 0) {
  panic("close failed")
}


fd106 := client.Open("/Ch3TfdeKGU.txt", client.O_RDWR|client.O_CREATE)
if(fd106 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd106, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd106, []byte("opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBi

ret = client.Seek(fd106, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


ret = client.Seek(fd106, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


fd107 := client.Open("/9MRCgnwRR5.txt", client.O_RDWR|client.O_CREATE)
if(fd107 < 0) {
  panic("open failed")
}


fd108 := client.Open("/dAkT0T3Pq_.txt", client.O_RDWR|client.O_CREATE)
if(fd108 < 0) {
  panic("open failed")
}

//fd state: (0) 08JYhq39h3y11QqTs7KCYrJSIyLD_zJoZoQaKqCn1GNwGF0lnkHJGGfqMxtCw7actkOBPTXC86zaOPe55SOYX79B0NdfnZUj57Ewj

ret = client.Write(fd107, []byte("hhORrgN5K7sgkTDDeCNGyiLXz8Xtsr98re"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) hhORrgN5K7sgkTDDeCNGyiLXz8Xtsr98reQaKqCn1GNwGF0lnkHJGGfqMxtCw7actkOBPTXC86zaOPe55SOYX79B0NdfnZUj57Ewj

ret = client.Close(fd107)
if(ret != 0) {
  panic("close failed")
}


fd109 := client.Open("/G2EwloCA2u.txt", client.O_RDWR|client.O_CREATE)
if(fd109 < 0) {
  panic("open failed")
}


fd110 := client.Open("/Ch3TfdeKGU.txt", client.O_RDWR|client.O_CREATE)
if(fd110 < 0) {
  panic("open failed")
}

//fd state: (37) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBi

ret = client.Write(fd106, []byte("s3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVBSWzj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVBSWzj

fd111 := client.Open("/XsdIbzj7jo.txt", client.O_RDWR|client.O_CREATE)
if(fd111 < 0) {
  panic("open failed")
}


ret = client.Close(fd109)
if(ret != 0) {
  panic("close failed")
}


fd112 := client.Open("/zjkNSSD_IZ.txt", client.O_RDWR|client.O_CREATE)
if(fd112 < 0) {
  panic("open failed")
}


ret = client.Close(fd111)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd106)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd112)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd110, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYV") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd108, []byte("6KAy19xqY8oLiyqAKXurYg7RRMcbs6RE6EZdYEQGrfzxV40l0Kr4ZYRC74Ps"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) 6KAy19xqY8oLiyqAKXurYg7RRMcbs6RE6EZdYEQGrfzxV40l0Kr4ZYRC74Ps
//fd state: (90) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVBSWzj

ret = client.Write(fd110, []byte("N86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_Ct"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (138) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_Ct
//fd state: (60) 6KAy19xqY8oLiyqAKXurYg7RRMcbs6RE6EZdYEQGrfzxV40l0Kr4ZYRC74Ps

ret = client.Write(fd108, []byte("yhdIHxkcP4S1hmeZKd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) 6KAy19xqY8oLiyqAKXurYg7RRMcbs6RE6EZdYEQGrfzxV40l0Kr4ZYRC74PsyhdIHxkcP4S1hmeZKd

fd113 := client.Open("/MQPxqPotwN.txt", client.O_RDWR|client.O_CREATE)
if(fd113 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd108, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}


buf, ret = client.Read(fd113, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (138) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_Ct

ret = client.Write(fd110, []byte("ngkddGPzHaXXQiR3upF6zZdS0j4voSxTpbrQDv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (176) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_CtngkddGPzHaXXQiR3upF6zZdS0j4voSxTpbrQDv

buf, ret = client.Read(fd113, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd113, []byte("AYnt6MEDhLQBtqbOGFJeK9ew4wpRhOrMeekfSkiZB2WPtnhuvvroe0_iTKvThozwFRNK6uzufscLSrKkuMiQ4X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) AYnt6MEDhLQBtqbOGFJeK9ew4wpRhOrMeekfSkiZB2WPtnhuvvroe0_iTKvThozwFRNK6uzufscLSrKkuMiQ4X
//fd state: (176) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_CtngkddGPzHaXXQiR3upF6zZdS0j4voSxTpbrQDv

ret = client.Write(fd110, []byte("pqKAn0HdNXEvhEPsxK9VEJgjghOi1j3OjF2CEROv2xkb9Fu9yFcX2kYl59HdJrBk_5O1it3WWkb7XcvHP0fMEwW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (263) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_CtngkddGPzHaXXQiR3upF6zZdS0j4voSxTpbrQDvpqKAn0HdNXEvhEPsxK9VEJgjghOi1j3OjF2CEROv2xkb9Fu9yFcX2kYl59HdJrBk_5O1it3WWkb7XcvHP0fMEwW

ret = client.Close(fd110)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd113)
if(ret != 0) {
  panic("close failed")
}


fd114 := client.Open("/G2EwloCA2u.txt", client.O_RDWR|client.O_CREATE)
if(fd114 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd108, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "syhdIHxkcP4S1hmeZKd") {
  panic("wrong data returned")
}


ret = client.Close(fd114)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd108)
if(ret != 0) {
  panic("close failed")
}


fd115 := client.Open("/hQaeVcHR5v.txt", client.O_RDWR|client.O_CREATE)
if(fd115 < 0) {
  panic("open failed")
}


ret = client.Close(fd115)
if(ret != 0) {
  panic("close failed")
}


fd116 := client.Open("/N_tIxiNaQn.txt", client.O_RDWR|client.O_CREATE)
if(fd116 < 0) {
  panic("open failed")
}


fd117 := client.Open("/_9uGqB8bKD.txt", client.O_RDWR|client.O_CREATE)
if(fd117 < 0) {
  panic("open failed")
}


fd118 := client.Open("/pZ1XV6tnhC.txt", client.O_RDWR|client.O_CREATE)
if(fd118 < 0) {
  panic("open failed")
}


ret = client.Seek(fd116, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}

//fd state: (44) 1N_HZXae4g6riiMXZiohWJNDDu66_Mmm6zBYBdZRxnow0P4HV0LGY3qTKKXCIreHFosYTMjSVDMyAb_fnv5C6

ret = client.Write(fd116, []byte("T7rzpjoZtDKBGBQ8vwKnlNK1LGsE43mM4KLSJHqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) 1N_HZXae4g6riiMXZiohWJNDDu66_Mmm6zBYBdZRxnowT7rzpjoZtDKBGBQ8vwKnlNK1LGsE43mM4KLSJHqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3S

ret = client.Close(fd117)
if(ret != 0) {
  panic("close failed")
}


fd119 := client.Open("/3Xuls7KVgK.txt", client.O_RDWR|client.O_CREATE)
if(fd119 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd116, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd118)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd119, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzguBS") {
  panic("wrong data returned")
}


ret = client.Seek(fd119, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Seek(fd119, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}

//fd state: (118) 1N_HZXae4g6riiMXZiohWJNDDu66_Mmm6zBYBdZRxnowT7rzpjoZtDKBGBQ8vwKnlNK1LGsE43mM4KLSJHqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3S

ret = client.Write(fd116, []byte("UpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) 1N_HZXae4g6riiMXZiohWJNDDu66_Mmm6zBYBdZRxnowT7rzpjoZtDKBGBQ8vwKnlNK1LGsE43mM4KLSJHqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

fd120 := client.Open("/a_pne7JkfX.txt", client.O_RDWR|client.O_CREATE)
if(fd120 < 0) {
  panic("open failed")
}


ret = client.Close(fd116)
if(ret != 0) {
  panic("close failed")
}


fd121 := client.Open("/7FGDTeWQh5.txt", client.O_RDWR|client.O_CREATE)
if(fd121 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd120, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TlTaSiKaPyxBM_FNCUpqAwqfZuv5BMZndMkibocqeazXiDGrm8UtnLwTApU") {
  panic("wrong data returned")
}


ret = client.Seek(fd119, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


buf, ret = client.Read(fd121, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (32) IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzguBSfl

ret = client.Write(fd119, []byte("v"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzgvBSfl

buf, ret = client.Read(fd119, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "BSfl") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd119, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd119, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


ret = client.Seek(fd120, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


ret = client.Seek(fd120, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


ret = client.Seek(fd120, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


buf, ret = client.Read(fd120, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nLwTApUFqkludMHpH2pcuuAqLJz4ag1nRARtQ") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd121, []byte("Vp_VJC1oTrE4p3S2GAiVeAhPVh9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) Vp_VJC1oTrE4p3S2GAiVeAhPVh9

ret = client.Seek(fd120, 88, client.SEEK_SET)
if(ret != 88) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 88)
  panic("seek failed")
}


ret = client.Seek(fd121, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Seek(fd120, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Seek(fd119, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Close(fd120)
if(ret != 0) {
  panic("close failed")
}


fd122 := client.Open("/nVFk0Llsip.txt", client.O_RDWR|client.O_CREATE)
if(fd122 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd119, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9L9P8Gw8RHzzj5DWSeYwzgvBSfl") {
  panic("wrong data returned")
}

//fd state: (37) IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzgvBSfl

ret = client.Write(fd119, []byte("qDO5DgxRcAoi3SGu6iDLt86hZEQPTZ5Pak1BJWfo0kt4hcwP8Fj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzgvBSflqDO5DgxRcAoi3SGu6iDLt86hZEQPTZ5Pak1BJWfo0kt4hcwP8Fj

fd123 := client.Open("/phpHm2CqW9.txt", client.O_RDWR|client.O_CREATE)
if(fd123 < 0) {
  panic("open failed")
}


fd124 := client.Open("/NPOgKkI4Qs.txt", client.O_RDWR|client.O_CREATE)
if(fd124 < 0) {
  panic("open failed")
}


ret = client.Close(fd124)
if(ret != 0) {
  panic("close failed")
}


fd125 := client.Open("/HjcABjXgn5.txt", client.O_RDWR|client.O_CREATE)
if(fd125 < 0) {
  panic("open failed")
}


ret = client.Close(fd123)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd122, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd121, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9") {
  panic("wrong data returned")
}


ret = client.Close(fd125)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd121, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd121, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


fd126 := client.Open("/o3Jy8jBnOV.txt", client.O_RDWR|client.O_CREATE)
if(fd126 < 0) {
  panic("open failed")
}


ret = client.Close(fd126)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd119, 84, client.SEEK_SET)
if(ret != 84) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 84)
  panic("seek failed")
}


ret = client.Seek(fd122, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd119)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd122)
if(ret != 0) {
  panic("close failed")
}

//fd state: (5) Vp_VJC1oTrE4p3S2GAiVeAhPVh9

ret = client.Write(fd121, []byte("HJ6NYVb7s7nswJ4WPm73m6JPy9UmeSSIEgFBkWH36l0Z3kyHchtj5Jh47G8mNvu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) Vp_VJHJ6NYVb7s7nswJ4WPm73m6JPy9UmeSSIEgFBkWH36l0Z3kyHchtj5Jh47G8mNvu
//fd state: (68) Vp_VJHJ6NYVb7s7nswJ4WPm73m6JPy9UmeSSIEgFBkWH36l0Z3kyHchtj5Jh47G8mNvu

ret = client.Write(fd121, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) Vp_VJHJ6NYVb7s7nswJ4WPm73m6JPy9UmeSSIEgFBkWH36l0Z3kyHchtj5Jh47G8mNvu

ret = client.Close(fd121)
if(ret != 0) {
  panic("close failed")
}


fd127 := client.Open("/G2EwloCA2u.txt", client.O_RDWR|client.O_CREATE)
if(fd127 < 0) {
  panic("open failed")
}


ret = client.Seek(fd127, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd127, []byte("ugYiWD15NtHBc4g3mvDWr2aXMOuGOyL_TY2VvN0n5ONkKvMMbP6XzO8vtKJtPS4N_jMjpbfqJTCDapCmX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) ugYiWD15NtHBc4g3mvDWr2aXMOuGOyL_TY2VvN0n5ONkKvMMbP6XzO8vtKJtPS4N_jMjpbfqJTCDapCmX

buf, ret = client.Read(fd127, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd127, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd127)
if(ret != 0) {
  panic("close failed")
}


fd128 := client.Open("/jZ1FzHmVGc.txt", client.O_RDWR|client.O_CREATE)
if(fd128 < 0) {
  panic("open failed")
}


fd129 := client.Open("/7UlTBjTyBh.txt", client.O_RDWR|client.O_CREATE)
if(fd129 < 0) {
  panic("open failed")
}


ret = client.Close(fd129)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd128, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd128, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd128, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd128, []byte("2XbURDfrD8pNQkCikR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) 2XbURDfrD8pNQkCikR

ret = client.Close(fd128)
if(ret != 0) {
  panic("close failed")
}


fd130 := client.Open("/ufMjKnKzWK.txt", client.O_RDWR|client.O_CREATE)
if(fd130 < 0) {
  panic("open failed")
}


fd131 := client.Open("/gKjRc6XTZ4.txt", client.O_RDWR|client.O_CREATE)
if(fd131 < 0) {
  panic("open failed")
}


ret = client.Close(fd130)
if(ret != 0) {
  panic("close failed")
}


fd132 := client.Open("/3QNUtjWCV0.txt", client.O_RDWR|client.O_CREATE)
if(fd132 < 0) {
  panic("open failed")
}


ret = client.Close(fd131)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd132, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd132)
if(ret != 0) {
  panic("close failed")
}


fd133 := client.Open("/bmPU4CEsRL.txt", client.O_RDWR|client.O_CREATE)
if(fd133 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd133, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) 

ret = client.Close(fd133)
if(ret != 0) {
  panic("close failed")
}


fd134 := client.Open("/jaPa72Giny.txt", client.O_RDWR|client.O_CREATE)
if(fd134 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd134, []byte("1Dggp9pk9O77dvQtYCsFJ2n05thCac2eS2vZXKb5nZ_mE7pJokWqgjFL8gOrjrIag6khUX47hmW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) 1Dggp9pk9O77dvQtYCsFJ2n05thCac2eS2vZXKb5nZ_mE7pJokWqgjFL8gOrjrIag6khUX47hmW

ret = client.Close(fd134)
if(ret != 0) {
  panic("close failed")
}


fd135 := client.Open("/7Ihi1jG3Qx.txt", client.O_RDWR|client.O_CREATE)
if(fd135 < 0) {
  panic("open failed")
}


ret = client.Seek(fd135, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd135, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd136 := client.Open("/s7jifbhgZq.txt", client.O_RDWR|client.O_CREATE)
if(fd136 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd136, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd137 := client.Open("/6O82q9cXp9.txt", client.O_RDWR|client.O_CREATE)
if(fd137 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd135, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) 

buf, ret = client.Read(fd137, 18)
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


ret = client.Seek(fd135, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd138 := client.Open("/Ch3TfdeKGU.txt", client.O_RDWR|client.O_CREATE)
if(fd138 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd138, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "opcQZzcDBsWO") {
  panic("wrong data returned")
}


ret = client.Seek(fd136, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd135, []byte("QNazqpcx_Dqby"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) QNazqpcx_Dqby
//fd state: (0) 

ret = client.Write(fd136, []byte("jkplCU3muLebBQqxDt2xa3ZPkqB5ZL14SsGUwWz3NUlXc77gUZFWcOnACHnV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) jkplCU3muLebBQqxDt2xa3ZPkqB5ZL14SsGUwWz3NUlXc77gUZFWcOnACHnV

ret = client.Seek(fd138, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}


ret = client.Seek(fd138, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


ret = client.Seek(fd136, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


fd139 := client.Open("/pgOFTKSQaT.txt", client.O_RDWR|client.O_CREATE)
if(fd139 < 0) {
  panic("open failed")
}


fd140 := client.Open("/1hYlxqnLZN.txt", client.O_RDWR|client.O_CREATE)
if(fd140 < 0) {
  panic("open failed")
}

//fd state: (13) QNazqpcx_Dqby

ret = client.Write(fd135, []byte("CCBIjPl4t9iyjKtUvIxD0gbypjNXRV3rXj696pBi8XlSiHG9ehTqgQLVN87"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) QNazqpcx_DqbyCCBIjPl4t9iyjKtUvIxD0gbypjNXRV3rXj696pBi8XlSiHG9ehTqgQLVN87

ret = client.Seek(fd138, 148, client.SEEK_SET)
if(ret != 148) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 148)
  panic("seek failed")
}


ret = client.Seek(fd136, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Seek(fd136, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


ret = client.Seek(fd139, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (24) jkplCU3muLebBQqxDt2xa3ZPkqB5ZL14SsGUwWz3NUlXc77gUZFWcOnACHnV

ret = client.Write(fd136, []byte("V9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5
//fd state: (123) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5

ret = client.Write(fd136, []byte("TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (165) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0

ret = client.Seek(fd140, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (72) QNazqpcx_DqbyCCBIjPl4t9iyjKtUvIxD0gbypjNXRV3rXj696pBi8XlSiHG9ehTqgQLVN87

ret = client.Write(fd135, []byte("2zn0PtF15VfyTBIi0sV7_edEYjZ2UxD2hhKbIHthDyqslMmEt14Xrr7TRUeLy5d9vUF8TzZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (143) QNazqpcx_DqbyCCBIjPl4t9iyjKtUvIxD0gbypjNXRV3rXj696pBi8XlSiHG9ehTqgQLVN872zn0PtF15VfyTBIi0sV7_edEYjZ2UxD2hhKbIHthDyqslMmEt14Xrr7TRUeLy5d9vUF8TzZ

ret = client.Close(fd139)
if(ret != 0) {
  panic("close failed")
}

//fd state: (143) QNazqpcx_DqbyCCBIjPl4t9iyjKtUvIxD0gbypjNXRV3rXj696pBi8XlSiHG9ehTqgQLVN872zn0PtF15VfyTBIi0sV7_edEYjZ2UxD2hhKbIHthDyqslMmEt14Xrr7TRUeLy5d9vUF8TzZ

ret = client.Write(fd135, []byte("cnLTxwNZNES9ZGbU2XVbh3n_goJcIQdRFoyS88GzF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (184) QNazqpcx_DqbyCCBIjPl4t9iyjKtUvIxD0gbypjNXRV3rXj696pBi8XlSiHG9ehTqgQLVN872zn0PtF15VfyTBIi0sV7_edEYjZ2UxD2hhKbIHthDyqslMmEt14Xrr7TRUeLy5d9vUF8TzZcnLTxwNZNES9ZGbU2XVbh3n_goJcIQdRFoyS88GzF

fd141 := client.Open("/7bFIx4fJhh.txt", client.O_RDWR|client.O_CREATE)
if(fd141 < 0) {
  panic("open failed")
}


fd142 := client.Open("/5bNqIKtQYU.txt", client.O_RDWR|client.O_CREATE)
if(fd142 < 0) {
  panic("open failed")
}

//fd state: (184) QNazqpcx_DqbyCCBIjPl4t9iyjKtUvIxD0gbypjNXRV3rXj696pBi8XlSiHG9ehTqgQLVN872zn0PtF15VfyTBIi0sV7_edEYjZ2UxD2hhKbIHthDyqslMmEt14Xrr7TRUeLy5d9vUF8TzZcnLTxwNZNES9ZGbU2XVbh3n_goJcIQdRFoyS88GzF

ret = client.Write(fd135, []byte("AEHlEC75qSYn_J1VDAEbiqcQ1ITy6614GFz6W6aY85xIXpp7lexLO7tWgqWqZqsbS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (249) QNazqpcx_DqbyCCBIjPl4t9iyjKtUvIxD0gbypjNXRV3rXj696pBi8XlSiHG9ehTqgQLVN872zn0PtF15VfyTBIi0sV7_edEYjZ2UxD2hhKbIHthDyqslMmEt14Xrr7TRUeLy5d9vUF8TzZcnLTxwNZNES9ZGbU2XVbh3n_goJcIQdRFoyS88GzFAEHlEC75qSYn_J1VDAEbiqcQ1ITy6614GFz6W6aY85xIXpp7lexLO7tWgqWqZqsbS

fd143 := client.Open("/N_tIxiNaQn.txt", client.O_RDWR|client.O_CREATE)
if(fd143 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd136, 10)
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


buf, ret = client.Read(fd135, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd144 := client.Open("/cej2zXEjS3.txt", client.O_RDWR|client.O_CREATE)
if(fd144 < 0) {
  panic("open failed")
}

//fd state: (0) 1N_HZXae4g6riiMXZiohWJNDDu66_Mmm6zBYBdZRxnowT7rzpjoZtDKBGBQ8vwKnlNK1LGsE43mM4KLSJHqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

ret = client.Write(fd143, []byte("N0kc7vTJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) N0kc7vTJ4g6riiMXZiohWJNDDu66_Mmm6zBYBdZRxnowT7rzpjoZtDKBGBQ8vwKnlNK1LGsE43mM4KLSJHqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

ret = client.Seek(fd135, 242, client.SEEK_SET)
if(ret != 242) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 242)
  panic("seek failed")
}

//fd state: (165) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0

ret = client.Write(fd136, []byte("YPNGbdFXiFg9YFEfYapjqJmSyBlon"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (194) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlon

ret = client.Close(fd135)
if(ret != 0) {
  panic("close failed")
}

//fd state: (194) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlon

ret = client.Write(fd136, []byte("IQyPSJnjvkqLNq6MSW0_YD2l6Pt80sbmXYrB7IAHOpLoDJpEVtB1zfg5Ih6hglMYPMjaJNoaooSnuOINoNVOKOHz4H75jF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (288) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlonIQyPSJnjvkqLNq6MSW0_YD2l6Pt80sbmXYrB7IAHOpLoDJpEVtB1zfg5Ih6hglMYPMjaJNoaooSnuOINoNVOKOHz4H75jF

ret = client.Seek(fd142, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd145 := client.Open("/0n5N6AKRi8.txt", client.O_RDWR|client.O_CREATE)
if(fd145 < 0) {
  panic("open failed")
}

//fd state: (288) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlonIQyPSJnjvkqLNq6MSW0_YD2l6Pt80sbmXYrB7IAHOpLoDJpEVtB1zfg5Ih6hglMYPMjaJNoaooSnuOINoNVOKOHz4H75jF

ret = client.Write(fd136, []byte("sWasBlyzOrGq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (300) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlonIQyPSJnjvkqLNq6MSW0_YD2l6Pt80sbmXYrB7IAHOpLoDJpEVtB1zfg5Ih6hglMYPMjaJNoaooSnuOINoNVOKOHz4H75jFsWasBlyzOrGq

buf, ret = client.Read(fd144, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "EOdUPhBtoJWXjm9QLlBr8sicDs") {
  panic("wrong data returned")
}


fd146 := client.Open("/UWgWPC4Yub.txt", client.O_RDWR|client.O_CREATE)
if(fd146 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd146, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd142, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd138, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "XXQiR3upF6zZdS0j4voSxTpbrQDvpqKAn0HdNXEvhEPsxK9VEJgjghOi1j3OjF2CEROv2xkb9Fu9yFcX2kYl59HdJrB") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd136, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd138, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "k_5O1it3WWkb7XcvHP0fMEwW") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd136, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (26) EOdUPhBtoJWXjm9QLlBr8sicDs

ret = client.Write(fd144, []byte("OtggkvTc35_kkVZCQ65F4DNjghshHw_ta7goiJR0xaWvnI050jm6J52Eqfn1RNOn0NLwY4USzWenc3d17eT8FkefF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) EOdUPhBtoJWXjm9QLlBr8sicDsOtggkvTc35_kkVZCQ65F4DNjghshHw_ta7goiJR0xaWvnI050jm6J52Eqfn1RNOn0NLwY4USzWenc3d17eT8FkefF

buf, ret = client.Read(fd146, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd144)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd143, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}

//fd state: (263) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_CtngkddGPzHaXXQiR3upF6zZdS0j4voSxTpbrQDvpqKAn0HdNXEvhEPsxK9VEJgjghOi1j3OjF2CEROv2xkb9Fu9yFcX2kYl59HdJrBk_5O1it3WWkb7XcvHP0fMEwW

ret = client.Write(fd138, []byte("uokCiB_InMLWragtSRwrCeJ63_GBMrqC65v9gsHBQAMoBmigTbyQ1Oz1mJWOIVA_t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (328) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_CtngkddGPzHaXXQiR3upF6zZdS0j4voSxTpbrQDvpqKAn0HdNXEvhEPsxK9VEJgjghOi1j3OjF2CEROv2xkb9Fu9yFcX2kYl59HdJrBk_5O1it3WWkb7XcvHP0fMEwWuokCiB_InMLWragtSRwrCeJ63_GBMrqC65v9gsHBQAMoBmigTbyQ1Oz1mJWOIVA_t
//fd state: (0) 

ret = client.Write(fd146, []byte("YJ75GOyWf2ax7keJpjXc2X7ZPeCs_b08HKnpcSQMFmWaqeLUx_s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) YJ75GOyWf2ax7keJpjXc2X7ZPeCs_b08HKnpcSQMFmWaqeLUx_s

buf, ret = client.Read(fd142, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd147 := client.Open("/9MRCgnwRR5.txt", client.O_RDWR|client.O_CREATE)
if(fd147 < 0) {
  panic("open failed")
}


fd148 := client.Open("/Zr7CSIZlSr.txt", client.O_RDWR|client.O_CREATE)
if(fd148 < 0) {
  panic("open failed")
}


ret = client.Seek(fd145, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd149 := client.Open("/OSvlBb_1Bb.txt", client.O_RDWR|client.O_CREATE)
if(fd149 < 0) {
  panic("open failed")
}


ret = client.Seek(fd138, 286, client.SEEK_SET)
if(ret != 286) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 286)
  panic("seek failed")
}


ret = client.Close(fd138)
if(ret != 0) {
  panic("close failed")
}


fd150 := client.Open("/vcKGIMt6fS.txt", client.O_RDWR|client.O_CREATE)
if(fd150 < 0) {
  panic("open failed")
}


ret = client.Seek(fd145, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd148, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) hhORrgN5K7sgkTDDeCNGyiLXz8Xtsr98reQaKqCn1GNwGF0lnkHJGGfqMxtCw7actkOBPTXC86zaOPe55SOYX79B0NdfnZUj57Ewj

ret = client.Write(fd147, []byte("RiwgoyosGn7WBKesjkHS1SE6xadBV1dGrzczXPmtkcQjOdzfQMy73m0duCefroN8KQFWxX85lvaH_Uc3ItOEAU9idlTnTvIn4ia"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) RiwgoyosGn7WBKesjkHS1SE6xadBV1dGrzczXPmtkcQjOdzfQMy73m0duCefroN8KQFWxX85lvaH_Uc3ItOEAU9idlTnTvIn4iawj

ret = client.Close(fd146)
if(ret != 0) {
  panic("close failed")
}


fd151 := client.Open("/iool_4YP2J.txt", client.O_RDWR|client.O_CREATE)
if(fd151 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd148, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd150, []byte("J6zoMh5gJuwrMUlIKaReEyr7oD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) J6zoMh5gJuwrMUlIKaReEyr7oD

ret = client.Seek(fd145, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (300) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlonIQyPSJnjvkqLNq6MSW0_YD2l6Pt80sbmXYrB7IAHOpLoDJpEVtB1zfg5Ih6hglMYPMjaJNoaooSnuOINoNVOKOHz4H75jFsWasBlyzOrGq

ret = client.Write(fd136, []byte("5b5hmF7tpBjAukES3ZBLX30Pc6Po9wueNNl4PcMwiFWR8j18hY27NGuvJ6z_5c7LUmP8RYgKxx2GR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (377) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlonIQyPSJnjvkqLNq6MSW0_YD2l6Pt80sbmXYrB7IAHOpLoDJpEVtB1zfg5Ih6hglMYPMjaJNoaooSnuOINoNVOKOHz4H75jFsWasBlyzOrGq5b5hmF7tpBjAukES3ZBLX30Pc6Po9wueNNl4PcMwiFWR8j18hY27NGuvJ6z_5c7LUmP8RYgKxx2GR
//fd state: (0) 

ret = client.Write(fd148, []byte("FeurBNM8PEQkTKAfm3Voc3cYk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) FeurBNM8PEQkTKAfm3Voc3cYk

ret = client.Close(fd148)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd149, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd152 := client.Open("/H3B8paSDlB.txt", client.O_RDWR|client.O_CREATE)
if(fd152 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd142, []byte("BQd_N0FuwM8QFB7YfCTBstNxsUyZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) BQd_N0FuwM8QFB7YfCTBstNxsUyZ

ret = client.Seek(fd150, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


fd153 := client.Open("/3HJvawgO69.txt", client.O_RDWR|client.O_CREATE)
if(fd153 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd143, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rzpjoZtDKBGBQ8vwK") {
  panic("wrong data returned")
}


ret = client.Close(fd149)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd152, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd150, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


ret = client.Seek(fd145, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (21) J6zoMh5gJuwrMUlIKaReEyr7oD

ret = client.Write(fd150, []byte("L"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) J6zoMh5gJuwrMUlIKaReELr7oD

fd154 := client.Open("/LoTeuE2wHL.txt", client.O_RDWR|client.O_CREATE)
if(fd154 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd147, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wj") {
  panic("wrong data returned")
}


ret = client.Seek(fd136, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Seek(fd142, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


fd155 := client.Open("/YlbdOUkjl9.txt", client.O_RDWR|client.O_CREATE)
if(fd155 < 0) {
  panic("open failed")
}


fd156 := client.Open("/su0syYXcBE.txt", client.O_RDWR|client.O_CREATE)
if(fd156 < 0) {
  panic("open failed")
}


fd157 := client.Open("/JOL265Tk2F.txt", client.O_RDWR|client.O_CREATE)
if(fd157 < 0) {
  panic("open failed")
}


fd158 := client.Open("/BJb62m0Xf0.txt", client.O_RDWR|client.O_CREATE)
if(fd158 < 0) {
  panic("open failed")
}


ret = client.Close(fd136)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd152)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd145, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd159 := client.Open("/WwL3oEDSji.txt", client.O_RDWR|client.O_CREATE)
if(fd159 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd155, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd150, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


ret = client.Seek(fd154, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd153)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd145)
if(ret != 0) {
  panic("close failed")
}


fd160 := client.Open("/80lh0X6OA1.txt", client.O_RDWR|client.O_CREATE)
if(fd160 < 0) {
  panic("open failed")
}


ret = client.Close(fd156)
if(ret != 0) {
  panic("close failed")
}


fd161 := client.Open("/q4ub60oZmc.txt", client.O_RDWR|client.O_CREATE)
if(fd161 < 0) {
  panic("open failed")
}


fd162 := client.Open("/FbhHwvwTTv.txt", client.O_RDWR|client.O_CREATE)
if(fd162 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd162, []byte("Rh3AFyMbPWEp86NiXTigFt3jd7Ro6DuUwpKn5W98jKz0p3AAWdV3MdH8J6IAOVxzzCsYXjtA4eGNWXjhovS4Qgs_CZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) Rh3AFyMbPWEp86NiXTigFt3jd7Ro6DuUwpKn5W98jKz0p3AAWdV3MdH8J6IAOVxzzCsYXjtA4eGNWXjhovS4Qgs_CZ
//fd state: (0) 

ret = client.Write(fd151, []byte("un0UI0ExyEVGjC8eqvGu5qAqDxRaPjlJx1uhrXx63uI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) un0UI0ExyEVGjC8eqvGu5qAqDxRaPjlJx1uhrXx63uI

fd163 := client.Open("/2pzPsiHvtK.txt", client.O_RDWR|client.O_CREATE)
if(fd163 < 0) {
  panic("open failed")
}


ret = client.Close(fd147)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd150)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd158, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd164 := client.Open("/gKjRc6XTZ4.txt", client.O_RDWR|client.O_CREATE)
if(fd164 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd155, []byte("66v7avDDW1heGnnccHA15aakzJhoMpAhrZj7RAFG__HXca5vKn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) 66v7avDDW1heGnnccHA15aakzJhoMpAhrZj7RAFG__HXca5vKn

ret = client.Close(fd158)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd160, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd142, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YfCTBstNxsUyZ") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd163, []byte("XnYw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) XnYw

fd165 := client.Open("/MPMXhpOG4e.txt", client.O_RDWR|client.O_CREATE)
if(fd165 < 0) {
  panic("open failed")
}


ret = client.Seek(fd142, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd161, []byte("58YyfutV2MO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsTDVtIse5fp9FLC7S24jUK3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) 58YyfutV2MO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsTDVtIse5fp9FLC7S24jUK3

fd166 := client.Open("/1fD0ZWvax2.txt", client.O_RDWR|client.O_CREATE)
if(fd166 < 0) {
  panic("open failed")
}


ret = client.Close(fd157)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd164, []byte("3VPBh9MV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) 3VPBh9MV
//fd state: (90) Rh3AFyMbPWEp86NiXTigFt3jd7Ro6DuUwpKn5W98jKz0p3AAWdV3MdH8J6IAOVxzzCsYXjtA4eGNWXjhovS4Qgs_CZ

ret = client.Write(fd162, []byte("P5lMKrWFOjb4G7lx5kuTin56Ythr5_miGUUdo4Cx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) Rh3AFyMbPWEp86NiXTigFt3jd7Ro6DuUwpKn5W98jKz0p3AAWdV3MdH8J6IAOVxzzCsYXjtA4eGNWXjhovS4Qgs_CZP5lMKrWFOjb4G7lx5kuTin56Ythr5_miGUUdo4Cx

fd167 := client.Open("/KLDZc9ap3z.txt", client.O_RDWR|client.O_CREATE)
if(fd167 < 0) {
  panic("open failed")
}


ret = client.Close(fd159)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd166, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd143, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Seek(fd166, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd167, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd160, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd163, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (75) 58YyfutV2MO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsTDVtIse5fp9FLC7S24jUK3

ret = client.Write(fd161, []byte("P966HyhdKi4ztI6W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) 58YyfutV2MO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsTDVtIse5fp9FLC7S24jUK3P966HyhdKi4ztI6W

ret = client.Seek(fd166, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd168 := client.Open("/ck3kOvcScG.txt", client.O_RDWR|client.O_CREATE)
if(fd168 < 0) {
  panic("open failed")
}

//fd state: (8) 3VPBh9MV

ret = client.Write(fd164, []byte("_esgeOc1uuawCcUzt4Gb72E1bRup06Hm6WlnprzFKERYXcTMDLm3ZnM9m6XCTjuCK9xys0DtshANs352ETmgLsVi6Cy2NVepyC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (106) 3VPBh9MV_esgeOc1uuawCcUzt4Gb72E1bRup06Hm6WlnprzFKERYXcTMDLm3ZnM9m6XCTjuCK9xys0DtshANs352ETmgLsVi6Cy2NVepyC

ret = client.Seek(fd167, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd142, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yZ") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd165, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd165, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd151)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd160)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd163, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "XnYw") {
  panic("wrong data returned")
}


ret = client.Seek(fd166, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd165, []byte("1B84203NJlhsXiws3iz5uKjyipOo7zt8D7LHX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) 1B84203NJlhsXiws3iz5uKjyipOo7zt8D7LHX

ret = client.Close(fd142)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd161, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd167, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd167, 84)
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


fd169 := client.Open("/BJb62m0Xf0.txt", client.O_RDWR|client.O_CREATE)
if(fd169 < 0) {
  panic("open failed")
}


ret = client.Close(fd161)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd154, []byte("vHad_07sUSPQY9oWG4Z6PgCXyO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) vHad_07sUSPQY9oWG4Z6PgCXyO

ret = client.Seek(fd167, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (5) N0kc7vTJ4g6riiMXZiohWJNDDu66_Mmm6zBYBdZRxnowT7rzpjoZtDKBGBQ8vwKnlNK1LGsE43mM4KLSJHqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

ret = client.Write(fd143, []byte("ZwWSWkEdJAKamsSLMO_pgbQP_d9zqnce8m84P14BTy7umFGnGSmdMquGNscefyWWYlzDtDNe8pHxM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) N0kc7ZwWSWkEdJAKamsSLMO_pgbQP_d9zqnce8m84P14BTy7umFGnGSmdMquGNscefyWWYlzDtDNe8pHxMqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

buf, ret = client.Read(fd165, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd143, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}

//fd state: (4) XnYw

ret = client.Write(fd163, []byte("wwn19frwRjfdt4JAS5telrVM6ElHz8Wmvk3v3rCVL3nYl0QW5PDJXFkkQSxxqj2XsV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) XnYwwwn19frwRjfdt4JAS5telrVM6ElHz8Wmvk3v3rCVL3nYl0QW5PDJXFkkQSxxqj2XsV

buf, ret = client.Read(fd169, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (70) XnYwwwn19frwRjfdt4JAS5telrVM6ElHz8Wmvk3v3rCVL3nYl0QW5PDJXFkkQSxxqj2XsV

ret = client.Write(fd163, []byte("C0mxX89m8yZIZIBnEAYvtg0NTpeR5qKrGpliIfxY9RLnjWs016Uq0z6EkeipmY0_XbcqkRg5ZbOMA8Yvb0Hidkgo4IKnqnMqh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (167) XnYwwwn19frwRjfdt4JAS5telrVM6ElHz8Wmvk3v3rCVL3nYl0QW5PDJXFkkQSxxqj2XsVC0mxX89m8yZIZIBnEAYvtg0NTpeR5qKrGpliIfxY9RLnjWs016Uq0z6EkeipmY0_XbcqkRg5ZbOMA8Yvb0Hidkgo4IKnqnMqh

buf, ret = client.Read(fd167, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd143, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


ret = client.Seek(fd162, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd169, []byte("8bDFp328sAPwsI9RWw1KkDPMvKilRKW5aF55byjwJmxCZ_yEzoHIWlbO7FfsTL27mtR80ZBZcg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) 8bDFp328sAPwsI9RWw1KkDPMvKilRKW5aF55byjwJmxCZ_yEzoHIWlbO7FfsTL27mtR80ZBZcg

ret = client.Close(fd162)
if(ret != 0) {
  panic("close failed")
}


fd170 := client.Open("/L3M3DlnPNb.txt", client.O_RDWR|client.O_CREATE)
if(fd170 < 0) {
  panic("open failed")
}


fd171 := client.Open("/kwPisEKdpD.txt", client.O_RDWR|client.O_CREATE)
if(fd171 < 0) {
  panic("open failed")
}

//fd state: (50) 66v7avDDW1heGnnccHA15aakzJhoMpAhrZj7RAFG__HXca5vKn

ret = client.Write(fd155, []byte("q1w8cwjwvhiRfXZBLo1i8nAS6u8VZGg_KVI961"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) 66v7avDDW1heGnnccHA15aakzJhoMpAhrZj7RAFG__HXca5vKnq1w8cwjwvhiRfXZBLo1i8nAS6u8VZGg_KVI961

ret = client.Seek(fd163, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


fd172 := client.Open("/6O82q9cXp9.txt", client.O_RDWR|client.O_CREATE)
if(fd172 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd172, 52)
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


ret = client.Close(fd170)
if(ret != 0) {
  panic("close failed")
}

//fd state: (106) 3VPBh9MV_esgeOc1uuawCcUzt4Gb72E1bRup06Hm6WlnprzFKERYXcTMDLm3ZnM9m6XCTjuCK9xys0DtshANs352ETmgLsVi6Cy2NVepyC

ret = client.Write(fd164, []byte("f4sRcHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) 3VPBh9MV_esgeOc1uuawCcUzt4Gb72E1bRup06Hm6WlnprzFKERYXcTMDLm3ZnM9m6XCTjuCK9xys0DtshANs352ETmgLsVi6Cy2NVepyCf4sRcHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX
//fd state: (0) 

ret = client.Write(fd168, []byte("MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7r"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7r

fd173 := client.Open("/u1j43GN2O2.txt", client.O_RDWR|client.O_CREATE)
if(fd173 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd143, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WWY") {
  panic("wrong data returned")
}

//fd state: (26) vHad_07sUSPQY9oWG4Z6PgCXyO

ret = client.Write(fd154, []byte("ivTVGGGybzmONdNaqs7JP8weAuoX4mw1tc7p7FJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9D"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) vHad_07sUSPQY9oWG4Z6PgCXyOivTVGGGybzmONdNaqs7JP8weAuoX4mw1tc7p7FJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9D

buf, ret = client.Read(fd167, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd169, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}

//fd state: (32) MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7r

ret = client.Write(fd168, []byte("g7mlGFWHiTouwaph_KvVIYpj_7oytq8A7OHYsjUhjGUIbGEFeTBtEjOMRjaT0GzQC4u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7rg7mlGFWHiTouwaph_KvVIYpj_7oytq8A7OHYsjUhjGUIbGEFeTBtEjOMRjaT0GzQC4u
//fd state: (107) vHad_07sUSPQY9oWG4Z6PgCXyOivTVGGGybzmONdNaqs7JP8weAuoX4mw1tc7p7FJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9D

ret = client.Write(fd154, []byte("U1k2m7DV0Sd44lzORm5fW3VbsTxtO3TER7fGcAILZwZr1azZ0GzAZHbm0i3SOG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (169) vHad_07sUSPQY9oWG4Z6PgCXyOivTVGGGybzmONdNaqs7JP8weAuoX4mw1tc7p7FJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9DU1k2m7DV0Sd44lzORm5fW3VbsTxtO3TER7fGcAILZwZr1azZ0GzAZHbm0i3SOG

fd174 := client.Open("/5OR5HN2nWE.txt", client.O_RDWR|client.O_CREATE)
if(fd174 < 0) {
  panic("open failed")
}

//fd state: (74) XnYwwwn19frwRjfdt4JAS5telrVM6ElHz8Wmvk3v3rCVL3nYl0QW5PDJXFkkQSxxqj2XsVC0mxX89m8yZIZIBnEAYvtg0NTpeR5qKrGpliIfxY9RLnjWs016Uq0z6EkeipmY0_XbcqkRg5ZbOMA8Yvb0Hidkgo4IKnqnMqh

ret = client.Write(fd163, []byte("Ui8vFHG9XFvRtsuFh0SUaUiH4A6LcV5uvJfbFqE6NUv_8Ikd1Xo0a5e"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (129) XnYwwwn19frwRjfdt4JAS5telrVM6ElHz8Wmvk3v3rCVL3nYl0QW5PDJXFkkQSxxqj2XsVC0mxUi8vFHG9XFvRtsuFh0SUaUiH4A6LcV5uvJfbFqE6NUv_8Ikd1Xo0a5epmY0_XbcqkRg5ZbOMA8Yvb0Hidkgo4IKnqnMqh

buf, ret = client.Read(fd173, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd167, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd174, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_") {
  panic("wrong data returned")
}


ret = client.Seek(fd163, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Close(fd163)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd169, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


ret = client.Close(fd143)
if(ret != 0) {
  panic("close failed")
}


fd175 := client.Open("/su0syYXcBE.txt", client.O_RDWR|client.O_CREATE)
if(fd175 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd171, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd175, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd167)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (82) cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_

ret = client.Write(fd174, []byte("JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (155) cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZ

fd176 := client.Open("/3QNUtjWCV0.txt", client.O_RDWR|client.O_CREATE)
if(fd176 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd175, []byte("6qgCJlnFlWwI5JKVCLSKsHH1CTcpRnLOnIXZT82X0nVIVBczBImi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) 6qgCJlnFlWwI5JKVCLSKsHH1CTcpRnLOnIXZT82X0nVIVBczBImi

ret = client.Close(fd154)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd176)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd155)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd164)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd169, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


fd177 := client.Open("/b6UqNkmZMn.txt", client.O_RDWR|client.O_CREATE)
if(fd177 < 0) {
  panic("open failed")
}


ret = client.Seek(fd168, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd173, []byte("0gLCmLMBq1NlfSuSRADchuX1BtLi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) 0gLCmLMBq1NlfSuSRADchuX1BtLi

fd178 := client.Open("/XsdIbzj7jo.txt", client.O_RDWR|client.O_CREATE)
if(fd178 < 0) {
  panic("open failed")
}

//fd state: (155) cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZ

ret = client.Write(fd174, []byte("pepwXJP5Sq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (165) cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5Sq
//fd state: (165) cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5Sq

ret = client.Write(fd174, []byte("heAg0blDhIDsZyXmgpJF07VGfEMha6ECPuveO9v7SJvOeoiC6Ic9Q7NbQeEPYzq_59soWKbDOB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (239) cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5SqheAg0blDhIDsZyXmgpJF07VGfEMha6ECPuveO9v7SJvOeoiC6Ic9Q7NbQeEPYzq_59soWKbDOB

ret = client.Seek(fd173, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd172, []byte("y0ZkzMhkJKlB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) y0ZkzMhkJKlB

fd179 := client.Open("/jzIhJ3ZvR7.txt", client.O_RDWR|client.O_CREATE)
if(fd179 < 0) {
  panic("open failed")
}


fd180 := client.Open("/nr0KUH2sAF.txt", client.O_RDWR|client.O_CREATE)
if(fd180 < 0) {
  panic("open failed")
}


fd181 := client.Open("/ufMjKnKzWK.txt", client.O_RDWR|client.O_CREATE)
if(fd181 < 0) {
  panic("open failed")
}


fd182 := client.Open("/7Ihi1jG3Qx.txt", client.O_RDWR|client.O_CREATE)
if(fd182 < 0) {
  panic("open failed")
}


ret = client.Seek(fd173, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


ret = client.Seek(fd177, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd181)
if(ret != 0) {
  panic("close failed")
}

//fd state: (239) cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5SqheAg0blDhIDsZyXmgpJF07VGfEMha6ECPuveO9v7SJvOeoiC6Ic9Q7NbQeEPYzq_59soWKbDOB

ret = client.Write(fd174, []byte("lI0FXJyR5pGQ77juWJqXqefAoNEVNy9ObQED7MEXq30KgSeTCuBSL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (292) cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5SqheAg0blDhIDsZyXmgpJF07VGfEMha6ECPuveO9v7SJvOeoiC6Ic9Q7NbQeEPYzq_59soWKbDOBlI0FXJyR5pGQ77juWJqXqefAoNEVNy9ObQED7MEXq30KgSeTCuBSL

ret = client.Seek(fd171, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd165)
if(ret != 0) {
  panic("close failed")
}


fd183 := client.Open("/vcKGIMt6fS.txt", client.O_RDWR|client.O_CREATE)
if(fd183 < 0) {
  panic("open failed")
}


ret = client.Close(fd174)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd175, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd184 := client.Open("/iaCC9va3ey.txt", client.O_RDWR|client.O_CREATE)
if(fd184 < 0) {
  panic("open failed")
}


ret = client.Close(fd178)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd169)
if(ret != 0) {
  panic("close failed")
}

//fd state: (52) 6qgCJlnFlWwI5JKVCLSKsHH1CTcpRnLOnIXZT82X0nVIVBczBImi

ret = client.Write(fd175, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) 6qgCJlnFlWwI5JKVCLSKsHH1CTcpRnLOnIXZT82X0nVIVBczBImi

ret = client.Seek(fd168, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}

//fd state: (12) y0ZkzMhkJKlB

ret = client.Write(fd172, []byte("jZ8U2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) y0ZkzMhkJKlBjZ8U2

ret = client.Close(fd175)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd180, []byte("GWQjGcejPWZWZhe8gvzXsZauFC8E5WnH7gC3uQdVtYxyXv5tgOaZ2JW1qGdD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) GWQjGcejPWZWZhe8gvzXsZauFC8E5WnH7gC3uQdVtYxyXv5tgOaZ2JW1qGdD

fd185 := client.Open("/7nzhcLMCES.txt", client.O_RDWR|client.O_CREATE)
if(fd185 < 0) {
  panic("open failed")
}


fd186 := client.Open("/4NexkYmnGj.txt", client.O_RDWR|client.O_CREATE)
if(fd186 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd185, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4jaaW5mgM2pYs0") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd177, []byte("WKWvZ1lw92y1PYCYEDu4aV6EOUcRygq5mB5Jq50FVb0kon1lhM6jlk_CAySw1Mlh8HHq5WM_kjINb0SvC0QaHoS8w7LAd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) WKWvZ1lw92y1PYCYEDu4aV6EOUcRygq5mB5Jq50FVb0kon1lhM6jlk_CAySw1Mlh8HHq5WM_kjINb0SvC0QaHoS8w7LAd

ret = client.Close(fd179)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd173, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "DchuX1BtLi") {
  panic("wrong data returned")
}

//fd state: (17) y0ZkzMhkJKlBjZ8U2

ret = client.Write(fd172, []byte("aLDOFjJLK72WwaF_zoANIMIA02NGrrvKOiuAIMUPwvrEdECZZxVdaUEFor9E1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) y0ZkzMhkJKlBjZ8U2aLDOFjJLK72WwaF_zoANIMIA02NGrrvKOiuAIMUPwvrEdECZZxVdaUEFor9E1

ret = client.Close(fd185)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd168)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd171, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd184, []byte("PLGbtSCiLunsrwMI18vfxWweCyy39DcJljicW2xpnBtcGK9tk9c9cKWKM4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) PLGbtSCiLunsrwMI18vfxWweCyy39DcJljicW2xpnBtcGK9tk9c9cKWKM4

ret = client.Close(fd177)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd172, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd186, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


ret = client.Seek(fd186, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Seek(fd183, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


buf, ret = client.Read(fd172, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd180, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}


buf, ret = client.Read(fd186, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "z_JL") {
  panic("wrong data returned")
}


fd187 := client.Open("/3HJvawgO69.txt", client.O_RDWR|client.O_CREATE)
if(fd187 < 0) {
  panic("open failed")
}

//fd state: (57) GWQjGcejPWZWZhe8gvzXsZauFC8E5WnH7gC3uQdVtYxyXv5tgOaZ2JW1qGdD

ret = client.Write(fd180, []byte("WUSOundME1n4NNIA0ktErIMwtGOp0Y2VlyKcVE1Jt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) GWQjGcejPWZWZhe8gvzXsZauFC8E5WnH7gC3uQdVtYxyXv5tgOaZ2JW1qWUSOundME1n4NNIA0ktErIMwtGOp0Y2VlyKcVE1Jt
//fd state: (0) 

ret = client.Write(fd171, []byte("hgYschvflbSHcx8esE45g_c3QPfFedZLoXALsAmClMQLai4IAlr_JTDrruN5kfJMBPEK4EOHW9nlSThzp1x"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) hgYschvflbSHcx8esE45g_c3QPfFedZLoXALsAmClMQLai4IAlr_JTDrruN5kfJMBPEK4EOHW9nlSThzp1x

ret = client.Seek(fd186, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd183, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


buf, ret = client.Read(fd183, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Lr7oD") {
  panic("wrong data returned")
}


ret = client.Close(fd182)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd186)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd171, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (83) hgYschvflbSHcx8esE45g_c3QPfFedZLoXALsAmClMQLai4IAlr_JTDrruN5kfJMBPEK4EOHW9nlSThzp1x

ret = client.Write(fd171, []byte("a0kMwrDbrsNNjf3RrONe5Oer0M52jFqMbEeUo2sbgF_2LjjrYetUd_biXYsuUI9sy7146"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) hgYschvflbSHcx8esE45g_c3QPfFedZLoXALsAmClMQLai4IAlr_JTDrruN5kfJMBPEK4EOHW9nlSThzp1xa0kMwrDbrsNNjf3RrONe5Oer0M52jFqMbEeUo2sbgF_2LjjrYetUd_biXYsuUI9sy7146
//fd state: (152) hgYschvflbSHcx8esE45g_c3QPfFedZLoXALsAmClMQLai4IAlr_JTDrruN5kfJMBPEK4EOHW9nlSThzp1xa0kMwrDbrsNNjf3RrONe5Oer0M52jFqMbEeUo2sbgF_2LjjrYetUd_biXYsuUI9sy7146

ret = client.Write(fd171, []byte("arw79cBtwNAvN1oQlHw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (171) hgYschvflbSHcx8esE45g_c3QPfFedZLoXALsAmClMQLai4IAlr_JTDrruN5kfJMBPEK4EOHW9nlSThzp1xa0kMwrDbrsNNjf3RrONe5Oer0M52jFqMbEeUo2sbgF_2LjjrYetUd_biXYsuUI9sy7146arw79cBtwNAvN1oQlHw

ret = client.Close(fd173)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd183, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd183, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}

//fd state: (13) J6zoMh5gJuwrMUlIKaReELr7oD

ret = client.Write(fd183, []byte("ICHKJnmai9_UzLxVLCTz3e3YIr6M8KClUBrNU1fLBj8jg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) J6zoMh5gJuwrMICHKJnmai9_UzLxVLCTz3e3YIr6M8KClUBrNU1fLBj8jg

fd188 := client.Open("/u18EbowvjG.txt", client.O_RDWR|client.O_CREATE)
if(fd188 < 0) {
  panic("open failed")
}


ret = client.Close(fd183)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd172, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd171, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd171, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd189 := client.Open("/2g8nvSsXFZ.txt", client.O_RDWR|client.O_CREATE)
if(fd189 < 0) {
  panic("open failed")
}


ret = client.Seek(fd189, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd190 := client.Open("/lC3aPrwuh2.txt", client.O_RDWR|client.O_CREATE)
if(fd190 < 0) {
  panic("open failed")
}


fd191 := client.Open("/yJ9Q0A1km0.txt", client.O_RDWR|client.O_CREATE)
if(fd191 < 0) {
  panic("open failed")
}


ret = client.Seek(fd171, 109, client.SEEK_SET)
if(ret != 109) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 109)
  panic("seek failed")
}


buf, ret = client.Read(fd180, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) VoLI6ui5PWBjZl_oZ04XVAfSvLbFsCaOzU5LhzywAsGRoj9ii916stUqV3idZMRNJPxKOxN_eiQUGQ6BVKhnMNjo2v

ret = client.Write(fd187, []byte("dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobxn6friy_WbUPD9mMu2NsGIx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobxn6friy_WbUPD9mMu2NsGIx

fd192 := client.Open("/phpHm2CqW9.txt", client.O_RDWR|client.O_CREATE)
if(fd192 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd184, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd193 := client.Open("/9MRCgnwRR5.txt", client.O_RDWR|client.O_CREATE)
if(fd193 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd189, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (58) PLGbtSCiLunsrwMI18vfxWweCyy39DcJljicW2xpnBtcGK9tk9c9cKWKM4

ret = client.Write(fd184, []byte("exaN6m6SmQHv7WCpUZ7D0eyR7S5F7mw6cL3OAfddawGgVK4KOfKYccrkmzFK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) PLGbtSCiLunsrwMI18vfxWweCyy39DcJljicW2xpnBtcGK9tk9c9cKWKM4exaN6m6SmQHv7WCpUZ7D0eyR7S5F7mw6cL3OAfddawGgVK4KOfKYccrkmzFK

buf, ret = client.Read(fd171, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "52jFqMbEeUo2s") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd192, []byte("TmSBOqyFw3s1D_3ciG2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) TmSBOqyFw3s1D_3ciG2

ret = client.Seek(fd172, 76, client.SEEK_SET)
if(ret != 76) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 76)
  panic("seek failed")
}


fd194 := client.Open("/sNE434pdqZ.txt", client.O_RDWR|client.O_CREATE)
if(fd194 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd188, 91)
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


ret = client.Close(fd171)
if(ret != 0) {
  panic("close failed")
}


fd195 := client.Open("/IgWgfliKnz.txt", client.O_RDWR|client.O_CREATE)
if(fd195 < 0) {
  panic("open failed")
}


ret = client.Close(fd190)
if(ret != 0) {
  panic("close failed")
}


fd196 := client.Open("/PSn4daw9hM.txt", client.O_RDWR|client.O_CREATE)
if(fd196 < 0) {
  panic("open failed")
}


ret = client.Seek(fd191, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) vNHeIG8AE0OXRRuez6GvUd7P58QxF76qelHROmPDoa8pdXX4JVMrq3ZGHiMeMz05ShUf6A

ret = client.Write(fd196, []byte("vZyCrQkPlbZ_tmIJ24s7xNyaRIDvdteok8bDJZnssxy7jE3mmpKRzyzau"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) vZyCrQkPlbZ_tmIJ24s7xNyaRIDvdteok8bDJZnssxy7jE3mmpKRzyzauiMeMz05ShUf6A
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd191)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd172, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


ret = client.Close(fd187)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd195, []byte("64vkLGSsN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) 64vkLGSsN

ret = client.Close(fd189)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd184)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd188)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd172, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "DOFjJLK72WwaF_zoANIMIA02NGrrvKOiuAIMUPwvrEdE") {
  panic("wrong data returned")
}

//fd state: (98) GWQjGcejPWZWZhe8gvzXsZauFC8E5WnH7gC3uQdVtYxyXv5tgOaZ2JW1qWUSOundME1n4NNIA0ktErIMwtGOp0Y2VlyKcVE1Jt

ret = client.Write(fd180, []byte("HuVHqAGzwaWna6iElmmpJakJtKeTfvqHaX7jj12mTjeypDqXpnLPEI1LXskh3W_xMmE4Fdk9T3PC5AdOat3uEXjewxZAtM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (192) GWQjGcejPWZWZhe8gvzXsZauFC8E5WnH7gC3uQdVtYxyXv5tgOaZ2JW1qWUSOundME1n4NNIA0ktErIMwtGOp0Y2VlyKcVE1JtHuVHqAGzwaWna6iElmmpJakJtKeTfvqHaX7jj12mTjeypDqXpnLPEI1LXskh3W_xMmE4Fdk9T3PC5AdOat3uEXjewxZAtM
//fd state: (57) vZyCrQkPlbZ_tmIJ24s7xNyaRIDvdteok8bDJZnssxy7jE3mmpKRzyzauiMeMz05ShUf6A

ret = client.Write(fd196, []byte("awjv0ptVAvIGEYSxa5B7P4QL6Y1V9hCKfUpXM1Ljo3Bf2K6t8vC_3iFac"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) vZyCrQkPlbZ_tmIJ24s7xNyaRIDvdteok8bDJZnssxy7jE3mmpKRzyzauawjv0ptVAvIGEYSxa5B7P4QL6Y1V9hCKfUpXM1Ljo3Bf2K6t8vC_3iFac

ret = client.Close(fd193)
if(ret != 0) {
  panic("close failed")
}


fd197 := client.Open("/gKjRc6XTZ4.txt", client.O_RDWR|client.O_CREATE)
if(fd197 < 0) {
  panic("open failed")
}

//fd state: (0) 3VPBh9MV_esgeOc1uuawCcUzt4Gb72E1bRup06Hm6WlnprzFKERYXcTMDLm3ZnM9m6XCTjuCK9xys0DtshANs352ETmgLsVi6Cy2NVepyCf4sRcHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX

ret = client.Write(fd197, []byte("12TdoW5oK4Fdre15ATjnGyFMs6lEmWSl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) 12TdoW5oK4Fdre15ATjnGyFMs6lEmWSlbRup06Hm6WlnprzFKERYXcTMDLm3ZnM9m6XCTjuCK9xys0DtshANs352ETmgLsVi6Cy2NVepyCf4sRcHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX

buf, ret = client.Read(fd195, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd196, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd194, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (32) 12TdoW5oK4Fdre15ATjnGyFMs6lEmWSlbRup06Hm6WlnprzFKERYXcTMDLm3ZnM9m6XCTjuCK9xys0DtshANs352ETmgLsVi6Cy2NVepyCf4sRcHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX

ret = client.Write(fd197, []byte("q5ofOFWWZ2UIEOldFSU7M_Wd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) 12TdoW5oK4Fdre15ATjnGyFMs6lEmWSlq5ofOFWWZ2UIEOldFSU7M_WdDLm3ZnM9m6XCTjuCK9xys0DtshANs352ETmgLsVi6Cy2NVepyCf4sRcHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX

buf, ret = client.Read(fd180, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd194, []byte("1efmTy3ecMr_ScncIPf0W8Pl0r8GStympl3iw2vw2u9a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) 1efmTy3ecMr_ScncIPf0W8Pl0r8GStympl3iw2vw2u9a

ret = client.Close(fd194)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd180)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd172)
if(ret != 0) {
  panic("close failed")
}

//fd state: (56) 12TdoW5oK4Fdre15ATjnGyFMs6lEmWSlq5ofOFWWZ2UIEOldFSU7M_WdDLm3ZnM9m6XCTjuCK9xys0DtshANs352ETmgLsVi6Cy2NVepyCf4sRcHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX

ret = client.Write(fd197, []byte("NTFNPH798X0MbOXKl1Cf_g4kbtFVQD0qtQ0t5v92P7f"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) 12TdoW5oK4Fdre15ATjnGyFMs6lEmWSlq5ofOFWWZ2UIEOldFSU7M_WdNTFNPH798X0MbOXKl1Cf_g4kbtFVQD0qtQ0t5v92P7f2NVepyCf4sRcHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX

ret = client.Close(fd196)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd195, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd197, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2NVepyCf4sRcHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX") {
  panic("wrong data returned")
}


ret = client.Close(fd197)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd195, 42)
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


fd198 := client.Open("/yOqQ9bu5R6.txt", client.O_RDWR|client.O_CREATE)
if(fd198 < 0) {
  panic("open failed")
}


ret = client.Close(fd198)
if(ret != 0) {
  panic("close failed")
}


fd199 := client.Open("/MhNRYjnYbb.txt", client.O_RDWR|client.O_CREATE)
if(fd199 < 0) {
  panic("open failed")
}


ret = client.Seek(fd199, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd199, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd199, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd199, []byte("5XmdnaNO1MFJPrC3HiefPzPah5K4SFlK9wx9jUYchOWcM1UmOVq0HtVgLPFB80kAbgjljfFOtohelXf2LlRFPz8EwbnDO3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) 5XmdnaNO1MFJPrC3HiefPzPah5K4SFlK9wx9jUYchOWcM1UmOVq0HtVgLPFB80kAbgjljfFOtohelXf2LlRFPz8EwbnDO3

buf, ret = client.Read(fd199, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd200 := client.Open("/rCKJKWJOO0.txt", client.O_RDWR|client.O_CREATE)
if(fd200 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd200, []byte("oar_tQZ923INkyNkqwOzoat2QB3jEx8HvhyoHrO1piTDMPUx6Q4bcKi25lZAQZKs9kB2vtMtP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) oar_tQZ923INkyNkqwOzoat2QB3jEx8HvhyoHrO1piTDMPUx6Q4bcKi25lZAQZKs9kB2vtMtP

buf, ret = client.Read(fd200, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd199, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


ret = client.Seek(fd200, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


buf, ret = client.Read(fd199, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "V") {
  panic("wrong data returned")
}


fd201 := client.Open("/hD5fZEysbP.txt", client.O_RDWR|client.O_CREATE)
if(fd201 < 0) {
  panic("open failed")
}


fd202 := client.Open("/jUD4VX4_c1.txt", client.O_RDWR|client.O_CREATE)
if(fd202 < 0) {
  panic("open failed")
}


ret = client.Close(fd201)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd202, []byte("gROrUzao3biiAEE9eJ4JfhR9CiHi6ANDV3T"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) gROrUzao3biiAEE9eJ4JfhR9CiHi6ANDV3T

ret = client.Seek(fd202, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


buf, ret = client.Read(fd199, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gLPFB80kAbgjljfFOtohelXf2LlRFPz8EwbnDO3") {
  panic("wrong data returned")
}


fd203 := client.Open("/O5jLTz9ceU.txt", client.O_RDWR|client.O_CREATE)
if(fd203 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd203, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "IFoZ6wJem8tM2m") {
  panic("wrong data returned")
}

//fd state: (94) 5XmdnaNO1MFJPrC3HiefPzPah5K4SFlK9wx9jUYchOWcM1UmOVq0HtVgLPFB80kAbgjljfFOtohelXf2LlRFPz8EwbnDO3

ret = client.Write(fd199, []byte("BsGXY0fMPfcgRDJ9eWxaT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) 5XmdnaNO1MFJPrC3HiefPzPah5K4SFlK9wx9jUYchOWcM1UmOVq0HtVgLPFB80kAbgjljfFOtohelXf2LlRFPz8EwbnDO3BsGXY0fMPfcgRDJ9eWxaT

ret = client.Seek(fd200, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


buf, ret = client.Read(fd200, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "HrO1piTDMPUx6Q4bcKi25lZAQZKs9kB2vtMtP") {
  panic("wrong data returned")
}


ret = client.Seek(fd202, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}

//fd state: (115) 5XmdnaNO1MFJPrC3HiefPzPah5K4SFlK9wx9jUYchOWcM1UmOVq0HtVgLPFB80kAbgjljfFOtohelXf2LlRFPz8EwbnDO3BsGXY0fMPfcgRDJ9eWxaT

ret = client.Write(fd199, []byte("1vb1Fo7S6oxqEjewLHyirRhcNowluq3lsXyIq_wtSWb8iXELrNA0HmAOIxCJ2fHX5g2a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (183) 5XmdnaNO1MFJPrC3HiefPzPah5K4SFlK9wx9jUYchOWcM1UmOVq0HtVgLPFB80kAbgjljfFOtohelXf2LlRFPz8EwbnDO3BsGXY0fMPfcgRDJ9eWxaT1vb1Fo7S6oxqEjewLHyirRhcNowluq3lsXyIq_wtSWb8iXELrNA0HmAOIxCJ2fHX5g2a
//fd state: (15) gROrUzao3biiAEE9eJ4JfhR9CiHi6ANDV3T

ret = client.Write(fd202, []byte("7JZ45sPXmlGSVLppp9S93k1FmopwKyhtP1fE1zOhHVkyirQu8q9KnsjSlDGlbSfpkQEFCSd4HP3ShNDoQn59LrFdmR86bq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (109) gROrUzao3biiAEE7JZ45sPXmlGSVLppp9S93k1FmopwKyhtP1fE1zOhHVkyirQu8q9KnsjSlDGlbSfpkQEFCSd4HP3ShNDoQn59LrFdmR86bq

buf, ret = client.Read(fd200, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd202, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (73) oar_tQZ923INkyNkqwOzoat2QB3jEx8HvhyoHrO1piTDMPUx6Q4bcKi25lZAQZKs9kB2vtMtP

ret = client.Write(fd200, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) oar_tQZ923INkyNkqwOzoat2QB3jEx8HvhyoHrO1piTDMPUx6Q4bcKi25lZAQZKs9kB2vtMtP

fd204 := client.Open("/dqzKim4aOI.txt", client.O_RDWR|client.O_CREATE)
if(fd204 < 0) {
  panic("open failed")
}


fd205 := client.Open("/TuSQnOwhBV.txt", client.O_RDWR|client.O_CREATE)
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


ret = client.Seek(fd200, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


ret = client.Seek(fd203, 69, client.SEEK_SET)
if(ret != 69) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 69)
  panic("seek failed")
}


buf, ret = client.Read(fd202, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd206 := client.Open("/hF6OAOMNoP.txt", client.O_RDWR|client.O_CREATE)
if(fd206 < 0) {
  panic("open failed")
}


ret = client.Close(fd200)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd199, 164, client.SEEK_SET)
if(ret != 164) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 164)
  panic("seek failed")
}


ret = client.Seek(fd206, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd202, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Seek(fd206, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd203, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7AO0") {
  panic("wrong data returned")
}


fd207 := client.Open("/pgOFTKSQaT.txt", client.O_RDWR|client.O_CREATE)
if(fd207 < 0) {
  panic("open failed")
}

//fd state: (26) gROrUzao3biiAEE7JZ45sPXmlGSVLppp9S93k1FmopwKyhtP1fE1zOhHVkyirQu8q9KnsjSlDGlbSfpkQEFCSd4HP3ShNDoQn59LrFdmR86bq

ret = client.Write(fd202, []byte("cH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) gROrUzao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJGGlbSfpkQEFCSd4HP3ShNDoQn59LrFdmR86bq

buf, ret = client.Read(fd207, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "iOg2CUg2O") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd206, []byte("YBa0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) YBa0

buf, ret = client.Read(fd199, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NA0HmA") {
  panic("wrong data returned")
}


fd208 := client.Open("/hCjNEVtEdq.txt", client.O_RDWR|client.O_CREATE)
if(fd208 < 0) {
  panic("open failed")
}


ret = client.Close(fd208)
if(ret != 0) {
  panic("close failed")
}


fd209 := client.Open("/JJw95P1AZF.txt", client.O_RDWR|client.O_CREATE)
if(fd209 < 0) {
  panic("open failed")
}


fd210 := client.Open("/u1j43GN2O2.txt", client.O_RDWR|client.O_CREATE)
if(fd210 < 0) {
  panic("open failed")
}


ret = client.Close(fd199)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd209, []byte("OA7Dw511gPquGZQIjCwsUdug_CdPI81JT9NE3eA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) OA7Dw511gPquGZQIjCwsUdug_CdPI81JT9NE3eA
//fd state: (73) gROrUzao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJGGlbSfpkQEFCSd4HP3ShNDoQn59LrFdmR86bq

ret = client.Write(fd202, []byte("8Hr1Soaw6gy4nWhLw_XUcAKW2u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) gROrUzao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJG8Hr1Soaw6gy4nWhLw_XUcAKW2uLrFdmR86bq

ret = client.Seek(fd207, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


fd211 := client.Open("/b6UqNkmZMn.txt", client.O_RDWR|client.O_CREATE)
if(fd211 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd202, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LrFdmR86bq") {
  panic("wrong data returned")
}

//fd state: (4) YBa0

ret = client.Write(fd206, []byte("qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmSX3cbAML"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmSX3cbAML
//fd state: (0) WKWvZ1lw92y1PYCYEDu4aV6EOUcRygq5mB5Jq50FVb0kon1lhM6jlk_CAySw1Mlh8HHq5WM_kjINb0SvC0QaHoS8w7LAd

ret = client.Write(fd211, []byte("yUYnwURDlPEutGGRUZSgw1FYgmdcxcSXeUiiT_KqFVqnbiA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) yUYnwURDlPEutGGRUZSgw1FYgmdcxcSXeUiiT_KqFVqnbiAlhM6jlk_CAySw1Mlh8HHq5WM_kjINb0SvC0QaHoS8w7LAd

ret = client.Close(fd211)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd210, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Seek(fd209, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (48) iOg2CUg2OG0EMMyPbtEshSXjgrWcmzDT9r1LgSgKSD5KftMFSDbSafKNv6t1wMmT9FomfrL

ret = client.Write(fd207, []byte("8BhrjlgAWAdl10CNCxIT4O4rXXNGYvjjz7AuT_JG2hQaMOgwWlWCky9Jaop1rRzbwoD7ibubhhoTcV_zZATkKAOjzCxcMjz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (143) iOg2CUg2OG0EMMyPbtEshSXjgrWcmzDT9r1LgSgKSD5KftMF8BhrjlgAWAdl10CNCxIT4O4rXXNGYvjjz7AuT_JG2hQaMOgwWlWCky9Jaop1rRzbwoD7ibubhhoTcV_zZATkKAOjzCxcMjz

fd212 := client.Open("/NhWH3hnYlu.txt", client.O_RDWR|client.O_CREATE)
if(fd212 < 0) {
  panic("open failed")
}


fd213 := client.Open("/Y0NpaqExV_.txt", client.O_RDWR|client.O_CREATE)
if(fd213 < 0) {
  panic("open failed")
}


ret = client.Close(fd212)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd209)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd207)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd214 := client.Open("/cej2zXEjS3.txt", client.O_RDWR|client.O_CREATE)
if(fd214 < 0) {
  panic("open failed")
}

//fd state: (2) 0gLCmLMBq1NlfSuSRADchuX1BtLi

ret = client.Write(fd210, []byte("f3F7cyGF9t3PbMgwRbT8XtSRnJk4qIuGEl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) 0gf3F7cyGF9t3PbMgwRbT8XtSRnJk4qIuGEl
//fd state: (59) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmSX3cbAML

ret = client.Write(fd206, []byte("QGcDN0le8SC6AQR8zKkGdwMz6o7b5LzwD_E"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmSX3cbAMLQGcDN0le8SC6AQR8zKkGdwMz6o7b5LzwD_E
//fd state: (94) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmSX3cbAMLQGcDN0le8SC6AQR8zKkGdwMz6o7b5LzwD_E

ret = client.Write(fd206, []byte("y6bSx9OlJuNfMYz0e_B2rEitL4XDXshYwa9EwCMe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmSX3cbAMLQGcDN0le8SC6AQR8zKkGdwMz6o7b5LzwD_Ey6bSx9OlJuNfMYz0e_B2rEitL4XDXshYwa9EwCMe
//fd state: (73) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJKYxDrMBH2z1Buyn7oLlSUWtYu7AO0

ret = client.Write(fd203, []byte("Gc6iHeOB9p87qzcLpiL1kCMiPIJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJKYxDrMBH2z1Buyn7oLlSUWtYu7AO0Gc6iHeOB9p87qzcLpiL1kCMiPIJ

buf, ret = client.Read(fd210, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd206, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) EOdUPhBtoJWXjm9QLlBr8sicDsOtggkvTc35_kkVZCQ65F4DNjghshHw_ta7goiJR0xaWvnI050jm6J52Eqfn1RNOn0NLwY4USzWenc3d17eT8FkefF

ret = client.Write(fd214, []byte("DbxHthGCT3EesMF9EKZ8XMp0FkpDHDYSL8RhgLvkpXpG3tLYl5D_a97gyEmEMSzUef360A5cfVfFpx70DRzuHcTy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) DbxHthGCT3EesMF9EKZ8XMp0FkpDHDYSL8RhgLvkpXpG3tLYl5D_a97gyEmEMSzUef360A5cfVfFpx70DRzuHcTyOn0NLwY4USzWenc3d17eT8FkefF

ret = client.Close(fd203)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd213, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd210, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Close(fd214)
if(ret != 0) {
  panic("close failed")
}


fd215 := client.Open("/gKjRc6XTZ4.txt", client.O_RDWR|client.O_CREATE)
if(fd215 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd213, []byte("7Tza4IwzlxZXvbbw9MYrJud_rHnI7p_tBoUL6L_lxa2PHYRjDOfwvp59rplq8AMfXl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) 7Tza4IwzlxZXvbbw9MYrJud_rHnI7p_tBoUL6L_lxa2PHYRjDOfwvp59rplq8AMfXl

ret = client.Seek(fd213, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


fd216 := client.Open("/MPMXhpOG4e.txt", client.O_RDWR|client.O_CREATE)
if(fd216 < 0) {
  panic("open failed")
}

//fd state: (0) 1B84203NJlhsXiws3iz5uKjyipOo7zt8D7LHX

ret = client.Write(fd216, []byte("v5AY3uvetJDvGFNFq4xx7AyPv0VDCzwLr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) v5AY3uvetJDvGFNFq4xx7AyPv0VDCzwLr7LHX

buf, ret = client.Read(fd202, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd217 := client.Open("/zFmjACN5Xo.txt", client.O_RDWR|client.O_CREATE)
if(fd217 < 0) {
  panic("open failed")
}


ret = client.Seek(fd206, 82, client.SEEK_SET)
if(ret != 82) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 82)
  panic("seek failed")
}

//fd state: (109) gROrUzao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJG8Hr1Soaw6gy4nWhLw_XUcAKW2uLrFdmR86bq

ret = client.Write(fd202, []byte("KTE8S8kEiN86IrFn94PRM32F8RevY2_xmIpmQZAiFxl3WbEcKjUdHXuvKg7UhtARSA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (175) gROrUzao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJG8Hr1Soaw6gy4nWhLw_XUcAKW2uLrFdmR86bqKTE8S8kEiN86IrFn94PRM32F8RevY2_xmIpmQZAiFxl3WbEcKjUdHXuvKg7UhtARSA

ret = client.Close(fd213)
if(ret != 0) {
  panic("close failed")
}


fd218 := client.Open("/FAdcvpgQ2d.txt", client.O_RDWR|client.O_CREATE)
if(fd218 < 0) {
  panic("open failed")
}


ret = client.Seek(fd215, 98, client.SEEK_SET)
if(ret != 98) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 98)
  panic("seek failed")
}


ret = client.Close(fd210)
if(ret != 0) {
  panic("close failed")
}


fd219 := client.Open("/gMb60G3hQV.txt", client.O_RDWR|client.O_CREATE)
if(fd219 < 0) {
  panic("open failed")
}


ret = client.Close(fd206)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd218)
if(ret != 0) {
  panic("close failed")
}

//fd state: (175) gROrUzao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJG8Hr1Soaw6gy4nWhLw_XUcAKW2uLrFdmR86bqKTE8S8kEiN86IrFn94PRM32F8RevY2_xmIpmQZAiFxl3WbEcKjUdHXuvKg7UhtARSA

ret = client.Write(fd202, []byte("QLkRenEXEL4ex0PJ3W_qr1KT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (199) gROrUzao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJG8Hr1Soaw6gy4nWhLw_XUcAKW2uLrFdmR86bqKTE8S8kEiN86IrFn94PRM32F8RevY2_xmIpmQZAiFxl3WbEcKjUdHXuvKg7UhtARSAQLkRenEXEL4ex0PJ3W_qr1KT

buf, ret = client.Read(fd216, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7LHX") {
  panic("wrong data returned")
}


ret = client.Seek(fd202, 162, client.SEEK_SET)
if(ret != 162) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 162)
  panic("seek failed")
}


fd220 := client.Open("/k6tgHqlgOv.txt", client.O_RDWR|client.O_CREATE)
if(fd220 < 0) {
  panic("open failed")
}


ret = client.Close(fd216)
if(ret != 0) {
  panic("close failed")
}


fd221 := client.Open("/YyhcUiTFU4.txt", client.O_RDWR|client.O_CREATE)
if(fd221 < 0) {
  panic("open failed")
}


ret = client.Seek(fd217, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd221)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd217)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd202, 143, client.SEEK_SET)
if(ret != 143) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 143)
  panic("seek failed")
}


fd222 := client.Open("/NrKdemH1gP.txt", client.O_RDWR|client.O_CREATE)
if(fd222 < 0) {
  panic("open failed")
}


fd223 := client.Open("/rFsXYFVDNi.txt", client.O_RDWR|client.O_CREATE)
if(fd223 < 0) {
  panic("open failed")
}


ret = client.Seek(fd220, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd215, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "f2") {
  panic("wrong data returned")
}


ret = client.Close(fd219)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd220)
if(ret != 0) {
  panic("close failed")
}


fd224 := client.Open("/rFsXYFVDNi.txt", client.O_RDWR|client.O_CREATE)
if(fd224 < 0) {
  panic("open failed")
}


ret = client.Seek(fd215, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}

//fd state: (143) gROrUzao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJG8Hr1Soaw6gy4nWhLw_XUcAKW2uLrFdmR86bqKTE8S8kEiN86IrFn94PRM32F8RevY2_xmIpmQZAiFxl3WbEcKjUdHXuvKg7UhtARSAQLkRenEXEL4ex0PJ3W_qr1KT

ret = client.Write(fd202, []byte("nVJo1XAg3EJLWtxsUxNlJ_4W7UMRjpjFfHkTTKCv1e3RW_YJR49"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (194) gROrUzao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJG8Hr1Soaw6gy4nWhLw_XUcAKW2uLrFdmR86bqKTE8S8kEiN86IrFn94PRM32F8RevY2_xmInVJo1XAg3EJLWtxsUxNlJ_4W7UMRjpjFfHkTTKCv1e3RW_YJR49qr1KT

ret = client.Seek(fd223, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd223, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd224, []byte("rksfJ5CALq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) rksfJ5CALq

fd225 := client.Open("/N_tIxiNaQn.txt", client.O_RDWR|client.O_CREATE)
if(fd225 < 0) {
  panic("open failed")
}


ret = client.Close(fd202)
if(ret != 0) {
  panic("close failed")
}


fd226 := client.Open("/QpV8PqvRKB.txt", client.O_RDWR|client.O_CREATE)
if(fd226 < 0) {
  panic("open failed")
}


ret = client.Close(fd222)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd225, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "N0kc7ZwWSWkEd") {
  panic("wrong data returned")
}


fd227 := client.Open("/qfwUZ2XRmv.txt", client.O_RDWR|client.O_CREATE)
if(fd227 < 0) {
  panic("open failed")
}


ret = client.Close(fd227)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd223)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd215)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd225, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JAKamsSLMO_pgbQP_") {
  panic("wrong data returned")
}


fd228 := client.Open("/ck3kOvcScG.txt", client.O_RDWR|client.O_CREATE)
if(fd228 < 0) {
  panic("open failed")
}


fd229 := client.Open("/0bepIOJ5hl.txt", client.O_RDWR|client.O_CREATE)
if(fd229 < 0) {
  panic("open failed")
}


ret = client.Close(fd224)
if(ret != 0) {
  panic("close failed")
}

//fd state: (30) N0kc7ZwWSWkEdJAKamsSLMO_pgbQP_d9zqnce8m84P14BTy7umFGnGSmdMquGNscefyWWYlzDtDNe8pHxMqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

ret = client.Write(fd225, []byte("KOiN2S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) N0kc7ZwWSWkEdJAKamsSLMO_pgbQP_KOiN2Se8m84P14BTy7umFGnGSmdMquGNscefyWWYlzDtDNe8pHxMqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI
//fd state: (0) 

ret = client.Write(fd229, []byte("e7GtGo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) e7GtGo

buf, ret = client.Read(fd226, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd230 := client.Open("/bmPU4CEsRL.txt", client.O_RDWR|client.O_CREATE)
if(fd230 < 0) {
  panic("open failed")
}


fd231 := client.Open("/yJ9Q0A1km0.txt", client.O_RDWR|client.O_CREATE)
if(fd231 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd228, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7rg7mlGFWHiTouwaph_KvVIYpj_7oytq8A7OHY") {
  panic("wrong data returned")
}


ret = client.Seek(fd230, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd232 := client.Open("/9MRCgnwRR5.txt", client.O_RDWR|client.O_CREATE)
if(fd232 < 0) {
  panic("open failed")
}


ret = client.Seek(fd226, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd225)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd228, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}


ret = client.Seek(fd232, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}

//fd state: (6) e7GtGo

ret = client.Write(fd229, []byte("9gJmFzHhRzcGsEGuOwNIvc0RgeSaszv3ZkMhEV4X8cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) e7GtGo9gJmFzHhRzcGsEGuOwNIvc0RgeSaszv3ZkMhEV4X8cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg
//fd state: (0) 

ret = client.Write(fd230, []byte("YI4x3IZsZqhTW8DjuvP4XA0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) YI4x3IZsZqhTW8DjuvP4XA0

ret = client.Seek(fd230, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}

//fd state: (79) e7GtGo9gJmFzHhRzcGsEGuOwNIvc0RgeSaszv3ZkMhEV4X8cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg

ret = client.Write(fd229, []byte("9xXbYIyTnXeA5iyMZRt2cAu6a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (104) e7GtGo9gJmFzHhRzcGsEGuOwNIvc0RgeSaszv3ZkMhEV4X8cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg9xXbYIyTnXeA5iyMZRt2cAu6a

fd233 := client.Open("/qksOWSAKgO.txt", client.O_RDWR|client.O_CREATE)
if(fd233 < 0) {
  panic("open failed")
}


ret = client.Close(fd231)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd226, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (15) YI4x3IZsZqhTW8DjuvP4XA0

ret = client.Write(fd230, []byte("MoFhbYkhta"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) YI4x3IZsZqhTW8DMoFhbYkhta
//fd state: (25) YI4x3IZsZqhTW8DMoFhbYkhta

ret = client.Write(fd230, []byte("vwgIuGvupDyTB6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) YI4x3IZsZqhTW8DMoFhbYkhtavwgIuGvupDyTB6

ret = client.Close(fd226)
if(ret != 0) {
  panic("close failed")
}


fd234 := client.Open("/BJb62m0Xf0.txt", client.O_RDWR|client.O_CREATE)
if(fd234 < 0) {
  panic("open failed")
}


fd235 := client.Open("/gXA10lAm1z.txt", client.O_RDWR|client.O_CREATE)
if(fd235 < 0) {
  panic("open failed")
}


ret = client.Close(fd232)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd233)
if(ret != 0) {
  panic("close failed")
}


fd236 := client.Open("/z7bftk_OK2.txt", client.O_RDWR|client.O_CREATE)
if(fd236 < 0) {
  panic("open failed")
}


ret = client.Close(fd235)
if(ret != 0) {
  panic("close failed")
}

//fd state: (104) e7GtGo9gJmFzHhRzcGsEGuOwNIvc0RgeSaszv3ZkMhEV4X8cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg9xXbYIyTnXeA5iyMZRt2cAu6a

ret = client.Write(fd229, []byte("bhMrV5Ar_zNZLrxFMb8ERTtRP5V1614gAXD26772YV8TDHh8YJE3pOxCxpdqgnklS6QEDjVH8810Oy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (182) e7GtGo9gJmFzHhRzcGsEGuOwNIvc0RgeSaszv3ZkMhEV4X8cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg9xXbYIyTnXeA5iyMZRt2cAu6abhMrV5Ar_zNZLrxFMb8ERTtRP5V1614gAXD26772YV8TDHh8YJE3pOxCxpdqgnklS6QEDjVH8810Oy

ret = client.Close(fd229)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd228, 82, client.SEEK_SET)
if(ret != 82) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 82)
  panic("seek failed")
}


ret = client.Close(fd234)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd230, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd237 := client.Open("/gyu1b9QkCi.txt", client.O_RDWR|client.O_CREATE)
if(fd237 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd236, []byte("jSDzXY82g9nyacFFgIB5qk3FP5QnSORsSSun69"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) jSDzXY82g9nyacFFgIB5qk3FP5QnSORsSSun69
//fd state: (39) YI4x3IZsZqhTW8DMoFhbYkhtavwgIuGvupDyTB6

ret = client.Write(fd230, []byte("wfuhYTlcp6beNFPwOhZjs9CUS6slo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) YI4x3IZsZqhTW8DMoFhbYkhtavwgIuGvupDyTB6wfuhYTlcp6beNFPwOhZjs9CUS6slo

ret = client.Seek(fd236, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd236, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


fd238 := client.Open("/NtMGa3u7Cn.txt", client.O_RDWR|client.O_CREATE)
if(fd238 < 0) {
  panic("open failed")
}


ret = client.Seek(fd230, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}

//fd state: (82) MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7rg7mlGFWHiTouwaph_KvVIYpj_7oytq8A7OHYsjUhjGUIbGEFeTBtEjOMRjaT0GzQC4u

ret = client.Write(fd228, []byte("cNgWg_hNl1OAr7x1i0R72hBYlppsm8js0ZNLN8TX7o"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7rg7mlGFWHiTouwaph_KvVIYpj_7oytq8A7OHYsjUhjGUIbGEFeTcNgWg_hNl1OAr7x1i0R72hBYlppsm8js0ZNLN8TX7o
//fd state: (55) YI4x3IZsZqhTW8DMoFhbYkhtavwgIuGvupDyTB6wfuhYTlcp6beNFPwOhZjs9CUS6slo

ret = client.Write(fd230, []byte("jZU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) YI4x3IZsZqhTW8DMoFhbYkhtavwgIuGvupDyTB6wfuhYTlcp6beNFPwjZUjs9CUS6slo

ret = client.Close(fd236)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd230, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "js9CUS6slo") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd237, []byte("qJCWMeoiiCmCbatMfdBUhrugm0uQx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) qJCWMeoiiCmCbatMfdBUhrugm0uQx

ret = client.Close(fd238)
if(ret != 0) {
  panic("close failed")
}

//fd state: (124) MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7rg7mlGFWHiTouwaph_KvVIYpj_7oytq8A7OHYsjUhjGUIbGEFeTcNgWg_hNl1OAr7x1i0R72hBYlppsm8js0ZNLN8TX7o

ret = client.Write(fd228, []byte("ZtpQjrAQI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7rg7mlGFWHiTouwaph_KvVIYpj_7oytq8A7OHYsjUhjGUIbGEFeTcNgWg_hNl1OAr7x1i0R72hBYlppsm8js0ZNLN8TX7oZtpQjrAQI

buf, ret = client.Read(fd230, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd239 := client.Open("/g_108pF2Kl.txt", client.O_RDWR|client.O_CREATE)
if(fd239 < 0) {
  panic("open failed")
}


fd240 := client.Open("/BGGm6UKLEe.txt", client.O_RDWR|client.O_CREATE)
if(fd240 < 0) {
  panic("open failed")
}


ret = client.Close(fd237)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd228, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd241 := client.Open("/Ep1TY33BtZ.txt", client.O_RDWR|client.O_CREATE)
if(fd241 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd240, []byte("iKVdL3BDwqxKJl6Ryg5eWsTL7AiZgqeFckSgtvx1DIGS6s5v3RP_Sgh2AbJb4Y90iGsSfwTTP4gNfDtv85VrkkB36F"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) iKVdL3BDwqxKJl6Ryg5eWsTL7AiZgqeFckSgtvx1DIGS6s5v3RP_Sgh2AbJb4Y90iGsSfwTTP4gNfDtv85VrkkB36F
//fd state: (90) iKVdL3BDwqxKJl6Ryg5eWsTL7AiZgqeFckSgtvx1DIGS6s5v3RP_Sgh2AbJb4Y90iGsSfwTTP4gNfDtv85VrkkB36F

ret = client.Write(fd240, []byte("quBw9vXB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) iKVdL3BDwqxKJl6Ryg5eWsTL7AiZgqeFckSgtvx1DIGS6s5v3RP_Sgh2AbJb4Y90iGsSfwTTP4gNfDtv85VrkkB36FquBw9vXB

ret = client.Close(fd241)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd228, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd230, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd242 := client.Open("/ck3kOvcScG.txt", client.O_RDWR|client.O_CREATE)
if(fd242 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd230, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd242, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


buf, ret = client.Read(fd242, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "j_7oytq8A7OHYsjUhjGUIbGEFeTcNgWg_hNl1OAr7x1i") {
  panic("wrong data returned")
}


ret = client.Close(fd239)
if(ret != 0) {
  panic("close failed")
}


fd243 := client.Open("/u18EbowvjG.txt", client.O_RDWR|client.O_CREATE)
if(fd243 < 0) {
  panic("open failed")
}


fd244 := client.Open("/nzjL4wl1H6.txt", client.O_RDWR|client.O_CREATE)
if(fd244 < 0) {
  panic("open failed")
}


ret = client.Close(fd240)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd244, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd242, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "0R72hBYlppsm8js0ZNLN8TX7oZtpQjrAQI") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd230, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd230, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Close(fd242)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd243, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd243, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd230, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (133) MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7rg7mlGFWHiTouwaph_KvVIYpj_7oytq8A7OHYsjUhjGUIbGEFeTcNgWg_hNl1OAr7x1i0R72hBYlppsm8js0ZNLN8TX7oZtpQjrAQI

ret = client.Write(fd228, []byte("DUOCcPOQskdU4v3LL3qDrtJpTOQ4qpX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7rg7mlGFWHiTouwaph_KvVIYpj_7oytq8A7OHYsjUhjGUIbGEFeTcNgWg_hNl1OAr7x1i0R72hBYlppsm8js0ZNLN8TX7oZtpQjrAQIDUOCcPOQskdU4v3LL3qDrtJpTOQ4qpX

ret = client.Close(fd230)
if(ret != 0) {
  panic("close failed")
}


fd245 := client.Open("/WNMJT2E_a5.txt", client.O_RDWR|client.O_CREATE)
if(fd245 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd243, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd228, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd246 := client.Open("/tKhMEZbZQS.txt", client.O_RDWR|client.O_CREATE)
if(fd246 < 0) {
  panic("open failed")
}


ret = client.Close(fd246)
if(ret != 0) {
  panic("close failed")
}


fd247 := client.Open("/vQMZf31nQt.txt", client.O_RDWR|client.O_CREATE)
if(fd247 < 0) {
  panic("open failed")
}


fd248 := client.Open("/BGGm6UKLEe.txt", client.O_RDWR|client.O_CREATE)
if(fd248 < 0) {
  panic("open failed")
}


fd249 := client.Open("/WwL3oEDSji.txt", client.O_RDWR|client.O_CREATE)
if(fd249 < 0) {
  panic("open failed")
}


ret = client.Close(fd249)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd243, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd243, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd250 := client.Open("/0TDHoBZ0MP.txt", client.O_RDWR|client.O_CREATE)
if(fd250 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd243, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd248, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "iKVdL3BDwqxKJl6Ryg5eWsTL7AiZgqeFckSgtvx1DIGS6s5v3RP_Sgh2AbJb4Y90iGsSf") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd243, []byte("traAygs_u5kC7IYoROoGQlnwHe0t0qT07VfZvrkhp4Almx45qrPVM0Oi7RMcTyykTouLUy7kx9eded"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) traAygs_u5kC7IYoROoGQlnwHe0t0qT07VfZvrkhp4Almx45qrPVM0Oi7RMcTyykTouLUy7kx9eded

fd251 := client.Open("/BJb62m0Xf0.txt", client.O_RDWR|client.O_CREATE)
if(fd251 < 0) {
  panic("open failed")
}


ret = client.Seek(fd250, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd251, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


buf, ret = client.Read(fd243, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd243, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd244, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd244, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd245, []byte("o0Yi6V584_GDTlmuGd_YrMEYppqzZLehPx4uVmL8NeD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) o0Yi6V584_GDTlmuGd_YrMEYppqzZLehPx4uVmL8NeD

buf, ret = client.Read(fd248, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wTTP4gNfDtv") {
  panic("wrong data returned")
}


ret = client.Seek(fd251, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


ret = client.Close(fd250)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd244, []byte("5YvqDDY1bkYy6NWCBAmmARKLNBejwzBtdJgtjvelC4BOtcofwLVFqEqVxV8raquW0IRBzdGieYJ7XcHugf3mHMXCOsG4OK_u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) 5YvqDDY1bkYy6NWCBAmmARKLNBejwzBtdJgtjvelC4BOtcofwLVFqEqVxV8raquW0IRBzdGieYJ7XcHugf3mHMXCOsG4OK_u

ret = client.Close(fd247)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd228, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


ret = client.Seek(fd228, 109, client.SEEK_SET)
if(ret != 109) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 109)
  panic("seek failed")
}


ret = client.Close(fd248)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd244, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


ret = client.Seek(fd243, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}

//fd state: (33) 5YvqDDY1bkYy6NWCBAmmARKLNBejwzBtdJgtjvelC4BOtcofwLVFqEqVxV8raquW0IRBzdGieYJ7XcHugf3mHMXCOsG4OK_u

ret = client.Write(fd244, []byte("776aCuE0SiJOTYHw7QPX2Ph9xlOUbD9NANZVL6vxGxpxBl3iWuSYU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) 5YvqDDY1bkYy6NWCBAmmARKLNBejwzBtd776aCuE0SiJOTYHw7QPX2Ph9xlOUbD9NANZVL6vxGxpxBl3iWuSYUXCOsG4OK_u

ret = client.Seek(fd228, 124, client.SEEK_SET)
if(ret != 124) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 124)
  panic("seek failed")
}

//fd state: (40) 8bDFp328sAPwsI9RWw1KkDPMvKilRKW5aF55byjwJmxCZ_yEzoHIWlbO7FfsTL27mtR80ZBZcg

ret = client.Write(fd251, []byte("NZ4xtdIEbCtC5gAANpxnJ3tJsXufNvG0blmc48Qtx3FoNQkOK6ZEBVQtJXGGeh3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) 8bDFp328sAPwsI9RWw1KkDPMvKilRKW5aF55byjwNZ4xtdIEbCtC5gAANpxnJ3tJsXufNvG0blmc48Qtx3FoNQkOK6ZEBVQtJXGGeh3

ret = client.Seek(fd245, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Close(fd251)
if(ret != 0) {
  panic("close failed")
}

//fd state: (86) 5YvqDDY1bkYy6NWCBAmmARKLNBejwzBtd776aCuE0SiJOTYHw7QPX2Ph9xlOUbD9NANZVL6vxGxpxBl3iWuSYUXCOsG4OK_u

ret = client.Write(fd244, []byte("VhkdLTySc3Jtz6QUZHrL_WRVUG_8WET"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) 5YvqDDY1bkYy6NWCBAmmARKLNBejwzBtd776aCuE0SiJOTYHw7QPX2Ph9xlOUbD9NANZVL6vxGxpxBl3iWuSYUVhkdLTySc3Jtz6QUZHrL_WRVUG_8WET

fd252 := client.Open("/hD5fZEysbP.txt", client.O_RDWR|client.O_CREATE)
if(fd252 < 0) {
  panic("open failed")
}


fd253 := client.Open("/98laAaj4bd.txt", client.O_RDWR|client.O_CREATE)
if(fd253 < 0) {
  panic("open failed")
}


fd254 := client.Open("/nzjL4wl1H6.txt", client.O_RDWR|client.O_CREATE)
if(fd254 < 0) {
  panic("open failed")
}


fd255 := client.Open("/LoTeuE2wHL.txt", client.O_RDWR|client.O_CREATE)
if(fd255 < 0) {
  panic("open failed")
}


ret = client.Close(fd243)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd252, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd254, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5YvqDDY1bkYy6NWCBAmmARKLNBejwzB") {
  panic("wrong data returned")
}


ret = client.Close(fd253)
if(ret != 0) {
  panic("close failed")
}


fd256 := client.Open("/nYIeV0bwC7.txt", client.O_RDWR|client.O_CREATE)
if(fd256 < 0) {
  panic("open failed")
}


ret = client.Close(fd228)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd256, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd255, 122, client.SEEK_SET)
if(ret != 122) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 122)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd252, []byte("isY8qh3NPha2lFgXejVqao8Pq6eotvwCCiw2GZMhArlAqmU_t_tu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) isY8qh3NPha2lFgXejVqao8Pq6eotvwCCiw2GZMhArlAqmU_t_tu

fd257 := client.Open("/YjFzDv8z8t.txt", client.O_RDWR|client.O_CREATE)
if(fd257 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd257, []byte("6Mpjcg9Bq9pUKvO4ej20AhXmV4Kz5EnBGzjW990zK5ZzXz9SU8k4jqnlo2oRXrhtl5DTfpF0Tlqa4V5OhoXcBUOR4SJSBdu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) 6Mpjcg9Bq9pUKvO4ej20AhXmV4Kz5EnBGzjW990zK5ZzXz9SU8k4jqnlo2oRXrhtl5DTfpF0Tlqa4V5OhoXcBUOR4SJSBdu
//fd state: (2) o0Yi6V584_GDTlmuGd_YrMEYppqzZLehPx4uVmL8NeD

ret = client.Write(fd245, []byte("WYhFJX8YWxvIFCMCGxl60cAUryO8wg4lMOab"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) o0WYhFJX8YWxvIFCMCGxl60cAUryO8wg4lMOabL8NeD

buf, ret = client.Read(fd252, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd258 := client.Open("/H95Lu2se4U.txt", client.O_RDWR|client.O_CREATE)
if(fd258 < 0) {
  panic("open failed")
}


ret = client.Seek(fd254, 115, client.SEEK_SET)
if(ret != 115) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 115)
  panic("seek failed")
}


buf, ret = client.Read(fd252, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd244)
if(ret != 0) {
  panic("close failed")
}

//fd state: (122) vHad_07sUSPQY9oWG4Z6PgCXyOivTVGGGybzmONdNaqs7JP8weAuoX4mw1tc7p7FJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9DU1k2m7DV0Sd44lzORm5fW3VbsTxtO3TER7fGcAILZwZr1azZ0GzAZHbm0i3SOG

ret = client.Write(fd255, []byte("cgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6l"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (218) vHad_07sUSPQY9oWG4Z6PgCXyOivTVGGGybzmONdNaqs7JP8weAuoX4mw1tc7p7FJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6l
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd258, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd245, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}

//fd state: (115) 5YvqDDY1bkYy6NWCBAmmARKLNBejwzBtd776aCuE0SiJOTYHw7QPX2Ph9xlOUbD9NANZVL6vxGxpxBl3iWuSYUVhkdLTySc3Jtz6QUZHrL_WRVUG_8WET

ret = client.Write(fd254, []byte("vC9OhPjKHyZDvrh2TktksdGix_HYVW748ze907L9KJT7VNKntnIJFHUSXKvS1Zr8obX1jDEsRb7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (190) 5YvqDDY1bkYy6NWCBAmmARKLNBejwzBtd776aCuE0SiJOTYHw7QPX2Ph9xlOUbD9NANZVL6vxGxpxBl3iWuSYUVhkdLTySc3Jtz6QUZHrL_WRVUG_8WvC9OhPjKHyZDvrh2TktksdGix_HYVW748ze907L9KJT7VNKntnIJFHUSXKvS1Zr8obX1jDEsRb7

ret = client.Close(fd245)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd255, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd257, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd257, 27)
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


ret = client.Close(fd252)
if(ret != 0) {
  panic("close failed")
}


fd259 := client.Open("/sNE434pdqZ.txt", client.O_RDWR|client.O_CREATE)
if(fd259 < 0) {
  panic("open failed")
}

//fd state: (95) 6Mpjcg9Bq9pUKvO4ej20AhXmV4Kz5EnBGzjW990zK5ZzXz9SU8k4jqnlo2oRXrhtl5DTfpF0Tlqa4V5OhoXcBUOR4SJSBdu

ret = client.Write(fd257, []byte("SdVRh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) 6Mpjcg9Bq9pUKvO4ej20AhXmV4Kz5EnBGzjW990zK5ZzXz9SU8k4jqnlo2oRXrhtl5DTfpF0Tlqa4V5OhoXcBUOR4SJSBduSdVRh

buf, ret = client.Read(fd258, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (218) vHad_07sUSPQY9oWG4Z6PgCXyOivTVGGGybzmONdNaqs7JP8weAuoX4mw1tc7p7FJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6l

ret = client.Write(fd255, []byte("rxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpq9Xy9Vb55endC1v_eRZ6HNjZPejHi0qZHFSlN_dhk7EljE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (302) vHad_07sUSPQY9oWG4Z6PgCXyOivTVGGGybzmONdNaqs7JP8weAuoX4mw1tc7p7FJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpq9Xy9Vb55endC1v_eRZ6HNjZPejHi0qZHFSlN_dhk7EljE

ret = client.Seek(fd259, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd258, []byte("RcVdYxX_cwhAeCpUpmrG5F6pB0Z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) RcVdYxX_cwhAeCpUpmrG5F6pB0Z

ret = client.Close(fd257)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd259)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd258)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd255)
if(ret != 0) {
  panic("close failed")
}


fd260 := client.Open("/LoTeuE2wHL.txt", client.O_RDWR|client.O_CREATE)
if(fd260 < 0) {
  panic("open failed")
}


ret = client.Close(fd254)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd260, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vHad_07sUSPQY9oWG4Z6PgCXyOivTVGGGybzmONdNaqs7JP8weAuoX4mw1tc7p7FJCyFmzcnLHvdBhHZw9E8kqR4zTEWT5") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd260, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "74C2aMJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9D") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd260, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "KFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUd") {
  panic("wrong data returned")
}


ret = client.Seek(fd260, 186, client.SEEK_SET)
if(ret != 186) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 186)
  panic("seek failed")
}


fd261 := client.Open("/tKhMEZbZQS.txt", client.O_RDWR|client.O_CREATE)
if(fd261 < 0) {
  panic("open failed")
}


ret = client.Seek(fd261, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd260, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


buf, ret = client.Read(fd261, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd260, 260, client.SEEK_SET)
if(ret != 260) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 260)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd261, []byte("IroLf7TEO2RpHKU8yzzs4Sfd_r4Y8xhRdEWZWh9Gf7iL0kp5FRBjY0nqTyM_LX0frxuuQ7h0YcDYXY8Ol_a5YxWlu3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) IroLf7TEO2RpHKU8yzzs4Sfd_r4Y8xhRdEWZWh9Gf7iL0kp5FRBjY0nqTyM_LX0frxuuQ7h0YcDYXY8Ol_a5YxWlu3

ret = client.Close(fd260)
if(ret != 0) {
  panic("close failed")
}


fd262 := client.Open("/mdASAirDD_.txt", client.O_RDWR|client.O_CREATE)
if(fd262 < 0) {
  panic("open failed")
}


fd263 := client.Open("/KLDZc9ap3z.txt", client.O_RDWR|client.O_CREATE)
if(fd263 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd262, []byte("WJ_ge0cWSi70VMJMxhZFJvMp2GPNNDxMJWJd3juE5iVvEDtFSouKS7WCwjt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) WJ_ge0cWSi70VMJMxhZFJvMp2GPNNDxMJWJd3juE5iVvEDtFSouKS7WCwjt

fd264 := client.Open("/u1j43GN2O2.txt", client.O_RDWR|client.O_CREATE)
if(fd264 < 0) {
  panic("open failed")
}


ret = client.Seek(fd264, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Close(fd262)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd263, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd265 := client.Open("/8c8dueKjTG.txt", client.O_RDWR|client.O_CREATE)
if(fd265 < 0) {
  panic("open failed")
}


ret = client.Seek(fd263, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd264, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "f3F7cyGF9t3PbMgwRbT8XtSRnJk4qIuGEl") {
  panic("wrong data returned")
}


ret = client.Seek(fd263, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd265)
if(ret != 0) {
  panic("close failed")
}

//fd state: (36) 0gf3F7cyGF9t3PbMgwRbT8XtSRnJk4qIuGEl

ret = client.Write(fd264, []byte("2tLMfdPobsSpmyV1eToSEWKI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) 0gf3F7cyGF9t3PbMgwRbT8XtSRnJk4qIuGEl2tLMfdPobsSpmyV1eToSEWKI
//fd state: (90) IroLf7TEO2RpHKU8yzzs4Sfd_r4Y8xhRdEWZWh9Gf7iL0kp5FRBjY0nqTyM_LX0frxuuQ7h0YcDYXY8Ol_a5YxWlu3

ret = client.Write(fd261, []byte("u8y6TLzexfoqgFyv6XGufHzjw7uC5GtFiSYN7L_Vz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (131) IroLf7TEO2RpHKU8yzzs4Sfd_r4Y8xhRdEWZWh9Gf7iL0kp5FRBjY0nqTyM_LX0frxuuQ7h0YcDYXY8Ol_a5YxWlu3u8y6TLzexfoqgFyv6XGufHzjw7uC5GtFiSYN7L_Vz

ret = client.Close(fd261)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd264, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Seek(fd264, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd263, []byte("20SHsU9J9my5qBwRfDYqyeG6TuyjU8YAnTk_u6lY1N8Dp8XbMqAKUtJIWJ6igBTEoUYLpnF5oVx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) 20SHsU9J9my5qBwRfDYqyeG6TuyjU8YAnTk_u6lY1N8Dp8XbMqAKUtJIWJ6igBTEoUYLpnF5oVx

ret = client.Close(fd263)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd264)
if(ret != 0) {
  panic("close failed")
}


fd266 := client.Open("/F3pbNgbZqI.txt", client.O_RDWR|client.O_CREATE)
if(fd266 < 0) {
  panic("open failed")
}


ret = client.Close(fd266)
if(ret != 0) {
  panic("close failed")
}


fd267 := client.Open("/Nhry0WoS3B.txt", client.O_RDWR|client.O_CREATE)
if(fd267 < 0) {
  panic("open failed")
}


ret = client.Seek(fd267, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd268 := client.Open("/LgDXLTVcW9.txt", client.O_RDWR|client.O_CREATE)
if(fd268 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd267, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd269 := client.Open("/N_tIxiNaQn.txt", client.O_RDWR|client.O_CREATE)
if(fd269 < 0) {
  panic("open failed")
}

//fd state: (0) N0kc7ZwWSWkEdJAKamsSLMO_pgbQP_KOiN2Se8m84P14BTy7umFGnGSmdMquGNscefyWWYlzDtDNe8pHxMqPoDZcXaBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

ret = client.Write(fd269, []byte("rffsCahgymr_Mgi1z1NdDwUxLLkP0IrhtanP9Sam1W31LwXpwMlVmcMXY2fHt8PkqxHJebpPpEHdG157yiCs6V7A7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) rffsCahgymr_Mgi1z1NdDwUxLLkP0IrhtanP9Sam1W31LwXpwMlVmcMXY2fHt8PkqxHJebpPpEHdG157yiCs6V7A7aBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI
//fd state: (0) 

ret = client.Write(fd268, []byte("ODgLvL1lOvNlxhyGun0uHKZawIcCuK1VotUdZUhrrKEWYfLQnDKPTKszs9cC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) ODgLvL1lOvNlxhyGun0uHKZawIcCuK1VotUdZUhrrKEWYfLQnDKPTKszs9cC
//fd state: (0) 

ret = client.Write(fd267, []byte("IMLhv4i7Mt7KBVfo3MZjzO1dLyItZuMGtnT03HqVXI60OylvCuMhl3PS5xoenn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) IMLhv4i7Mt7KBVfo3MZjzO1dLyItZuMGtnT03HqVXI60OylvCuMhl3PS5xoenn

ret = client.Seek(fd268, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}

//fd state: (62) IMLhv4i7Mt7KBVfo3MZjzO1dLyItZuMGtnT03HqVXI60OylvCuMhl3PS5xoenn

ret = client.Write(fd267, []byte("AqZ5A6UvY3nE7EbuluC5MV2GF0PywELri4dyWUOmTbIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF0rJYL4yE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) IMLhv4i7Mt7KBVfo3MZjzO1dLyItZuMGtnT03HqVXI60OylvCuMhl3PS5xoennAqZ5A6UvY3nE7EbuluC5MV2GF0PywELri4dyWUOmTbIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF0rJYL4yE

ret = client.Close(fd269)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd267, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


fd270 := client.Open("/9HIUq0L3Jg.txt", client.O_RDWR|client.O_CREATE)
if(fd270 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd270, []byte("Gsx5686MnwNrzB_a4fV4WqBFf5BZswYvCiGwXu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) Gsx5686MnwNrzB_a4fV4WqBFf5BZswYvCiGwXu

fd271 := client.Open("/g52cddaIMR.txt", client.O_RDWR|client.O_CREATE)
if(fd271 < 0) {
  panic("open failed")
}

//fd state: (5) IMLhv4i7Mt7KBVfo3MZjzO1dLyItZuMGtnT03HqVXI60OylvCuMhl3PS5xoennAqZ5A6UvY3nE7EbuluC5MV2GF0PywELri4dyWUOmTbIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF0rJYL4yE

ret = client.Write(fd267, []byte("mZLPRV8RZ4AE7d08y9KS7BqrHtmmbiwlqh4VpMRT0cjaWx2OwdRlXzhjypQxVIg9URhpIw2lEMjogQzZcAWlaCyV0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) IMLhvmZLPRV8RZ4AE7d08y9KS7BqrHtmmbiwlqh4VpMRT0cjaWx2OwdRlXzhjypQxVIg9URhpIw2lEMjogQzZcAWlaCyV0i4dyWUOmTbIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF0rJYL4yE

buf, ret = client.Read(fd268, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uHKZawIcCuK1VotUdZUhrrKEWYfLQnDKPTKszs9cC") {
  panic("wrong data returned")
}


fd272 := client.Open("/q4ub60oZmc.txt", client.O_RDWR|client.O_CREATE)
if(fd272 < 0) {
  panic("open failed")
}

//fd state: (94) IMLhvmZLPRV8RZ4AE7d08y9KS7BqrHtmmbiwlqh4VpMRT0cjaWx2OwdRlXzhjypQxVIg9URhpIw2lEMjogQzZcAWlaCyV0i4dyWUOmTbIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF0rJYL4yE

ret = client.Write(fd267, []byte("Ny8aKx13H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) IMLhvmZLPRV8RZ4AE7d08y9KS7BqrHtmmbiwlqh4VpMRT0cjaWx2OwdRlXzhjypQxVIg9URhpIw2lEMjogQzZcAWlaCyV0Ny8aKx13HbIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF0rJYL4yE

fd273 := client.Open("/XFfR2z8Yx8.txt", client.O_RDWR|client.O_CREATE)
if(fd273 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd268, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd267, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF0rJYL4yE") {
  panic("wrong data returned")
}


fd274 := client.Open("/5bNqIKtQYU.txt", client.O_RDWR|client.O_CREATE)
if(fd274 < 0) {
  panic("open failed")
}


ret = client.Close(fd270)
if(ret != 0) {
  panic("close failed")
}

//fd state: (152) IMLhvmZLPRV8RZ4AE7d08y9KS7BqrHtmmbiwlqh4VpMRT0cjaWx2OwdRlXzhjypQxVIg9URhpIw2lEMjogQzZcAWlaCyV0Ny8aKx13HbIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF0rJYL4yE

ret = client.Write(fd267, []byte("0yUku2_I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (160) IMLhvmZLPRV8RZ4AE7d08y9KS7BqrHtmmbiwlqh4VpMRT0cjaWx2OwdRlXzhjypQxVIg9URhpIw2lEMjogQzZcAWlaCyV0Ny8aKx13HbIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF0rJYL4yE0yUku2_I

ret = client.Seek(fd274, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


ret = client.Seek(fd268, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


buf, ret = client.Read(fd268, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OvNlxhyGun0uHKZawIcCuK1VotUdZUhrrKEWYfLQnDKPTKszs9cC") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd272, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "58YyfutV") {
  panic("wrong data returned")
}

//fd state: (60) ODgLvL1lOvNlxhyGun0uHKZawIcCuK1VotUdZUhrrKEWYfLQnDKPTKszs9cC

ret = client.Write(fd268, []byte("EJH5mgmF8T8IAXvSGOEkEE7oslN7ldz7W8sipoHHSeZsVFlFWoHPprxSJj05qIuHv4kMnQc6Y2O_lb9L"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (140) ODgLvL1lOvNlxhyGun0uHKZawIcCuK1VotUdZUhrrKEWYfLQnDKPTKszs9cCEJH5mgmF8T8IAXvSGOEkEE7oslN7ldz7W8sipoHHSeZsVFlFWoHPprxSJj05qIuHv4kMnQc6Y2O_lb9L

buf, ret = client.Read(fd271, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd275 := client.Open("/u1j43GN2O2.txt", client.O_RDWR|client.O_CREATE)
if(fd275 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd274, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7YfCTBstNxsUyZ") {
  panic("wrong data returned")
}


fd276 := client.Open("/JJw95P1AZF.txt", client.O_RDWR|client.O_CREATE)
if(fd276 < 0) {
  panic("open failed")
}


ret = client.Seek(fd275, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


buf, ret = client.Read(fd268, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd271, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd267)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd271, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd277 := client.Open("/w3GJXKziAw.txt", client.O_RDWR|client.O_CREATE)
if(fd277 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd271, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd276, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OA7Dw511gPqu") {
  panic("wrong data returned")
}


ret = client.Close(fd277)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd271, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (26) 0gf3F7cyGF9t3PbMgwRbT8XtSRnJk4qIuGEl2tLMfdPobsSpmyV1eToSEWKI

ret = client.Write(fd275, []byte("4lN6DCmVRBisA4PGQSBiLu3VWxLhAzzoXl6yGBbvYvXz_eOqG7o6AzzqQSoxS_kAJOpDd0Cll6hYG1uh6TYRoSj5CbM4a6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) 0gf3F7cyGF9t3PbMgwRbT8XtSR4lN6DCmVRBisA4PGQSBiLu3VWxLhAzzoXl6yGBbvYvXz_eOqG7o6AzzqQSoxS_kAJOpDd0Cll6hYG1uh6TYRoSj5CbM4a6
//fd state: (0) F34DAk2KuCug8XuTQsx6CCyRVjQooGji1QUU0NGioxMQIbowQQupHt4N2xPBo5nXE3J

ret = client.Write(fd273, []byte("7tTVQXHcopQx8K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) 7tTVQXHcopQx8KuTQsx6CCyRVjQooGji1QUU0NGioxMQIbowQQupHt4N2xPBo5nXE3J
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd268, 65, client.SEEK_SET)
if(ret != 65) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 65)
  panic("seek failed")
}


buf, ret = client.Read(fd274, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd278 := client.Open("/2BgRhfO5sI.txt", client.O_RDWR|client.O_CREATE)
if(fd278 < 0) {
  panic("open failed")
}


ret = client.Seek(fd273, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Close(fd272)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd278, 41)
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


ret = client.Seek(fd278, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd278, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd279 := client.Open("/kbPa5OvfPV.txt", client.O_RDWR|client.O_CREATE)
if(fd279 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd271, []byte("XhDfZsalD_4KOYaQiqwN7HcpJxhRkVgyuiMXcVW4c2Z9rcBSulA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) XhDfZsalD_4KOYaQiqwN7HcpJxhRkVgyuiMXcVW4c2Z9rcBSulA

ret = client.Seek(fd279, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd280 := client.Open("/sfmM5ZOnpc.txt", client.O_RDWR|client.O_CREATE)
if(fd280 < 0) {
  panic("open failed")
}

//fd state: (51) XhDfZsalD_4KOYaQiqwN7HcpJxhRkVgyuiMXcVW4c2Z9rcBSulA

ret = client.Write(fd271, []byte("k"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) XhDfZsalD_4KOYaQiqwN7HcpJxhRkVgyuiMXcVW4c2Z9rcBSulAk

ret = client.Seek(fd271, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


ret = client.Close(fd276)
if(ret != 0) {
  panic("close failed")
}

//fd state: (120) 0gf3F7cyGF9t3PbMgwRbT8XtSR4lN6DCmVRBisA4PGQSBiLu3VWxLhAzzoXl6yGBbvYvXz_eOqG7o6AzzqQSoxS_kAJOpDd0Cll6hYG1uh6TYRoSj5CbM4a6

ret = client.Write(fd275, []byte("GcLcHfrTqqtxfHyMct3UmExPTIsnxtwHrma9G86aFkGeMe9pPjg2IglS3E9qDTF4piSMUQzyZqa2l6zGfZsn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (204) 0gf3F7cyGF9t3PbMgwRbT8XtSR4lN6DCmVRBisA4PGQSBiLu3VWxLhAzzoXl6yGBbvYvXz_eOqG7o6AzzqQSoxS_kAJOpDd0Cll6hYG1uh6TYRoSj5CbM4a6GcLcHfrTqqtxfHyMct3UmExPTIsnxtwHrma9G86aFkGeMe9pPjg2IglS3E9qDTF4piSMUQzyZqa2l6zGfZsn
//fd state: (22) XhDfZsalD_4KOYaQiqwN7HcpJxhRkVgyuiMXcVW4c2Z9rcBSulAk

ret = client.Write(fd271, []byte("xx5ZGzt1pUaOmBhVl_8dlRTg8LZRmCm3YEQ4IttUoOCP2692ot"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) XhDfZsalD_4KOYaQiqwN7Hxx5ZGzt1pUaOmBhVl_8dlRTg8LZRmCm3YEQ4IttUoOCP2692ot

buf, ret = client.Read(fd278, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd278, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd281 := client.Open("/ANIOfqGwf0.txt", client.O_RDWR|client.O_CREATE)
if(fd281 < 0) {
  panic("open failed")
}


fd282 := client.Open("/k6tgHqlgOv.txt", client.O_RDWR|client.O_CREATE)
if(fd282 < 0) {
  panic("open failed")
}


fd283 := client.Open("/LQ7H9N02Pi.txt", client.O_RDWR|client.O_CREATE)
if(fd283 < 0) {
  panic("open failed")
}


fd284 := client.Open("/4a6jwdfaVN.txt", client.O_RDWR|client.O_CREATE)
if(fd284 < 0) {
  panic("open failed")
}


ret = client.Close(fd278)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd282, []byte("d2KG6QZ18bnxwE9l2HVd6dydDymtX8wKJeMyOvv8EGB_w441mF0zmUGvX2526aGGUrAPyjORXkHm1VyGqtDCsRVis0g96L22P"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) d2KG6QZ18bnxwE9l2HVd6dydDymtX8wKJeMyOvv8EGB_w441mF0zmUGvX2526aGGUrAPyjORXkHm1VyGqtDCsRVis0g96L22P

ret = client.Close(fd275)
if(ret != 0) {
  panic("close failed")
}

//fd state: (28) BQd_N0FuwM8QFB7YfCTBstNxsUyZ

ret = client.Write(fd274, []byte("MaFw3I099ALEEjh8H_5KTP_ziekMjXcCHQTrYQXT68zgoOv_1I5tlC4F"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) BQd_N0FuwM8QFB7YfCTBstNxsUyZMaFw3I099ALEEjh8H_5KTP_ziekMjXcCHQTrYQXT68zgoOv_1I5tlC4F

buf, ret = client.Read(fd274, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (97) d2KG6QZ18bnxwE9l2HVd6dydDymtX8wKJeMyOvv8EGB_w441mF0zmUGvX2526aGGUrAPyjORXkHm1VyGqtDCsRVis0g96L22P

ret = client.Write(fd282, []byte("H0ZYwCcdZ2uL0V4YnJbVwWdkdPfI7nGsyskPHZAazpkn3zPduW00flwfBUicDniNsME7Prmldt75UIa1gy5KHuGfPQNyni2Ps"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (194) d2KG6QZ18bnxwE9l2HVd6dydDymtX8wKJeMyOvv8EGB_w441mF0zmUGvX2526aGGUrAPyjORXkHm1VyGqtDCsRVis0g96L22PH0ZYwCcdZ2uL0V4YnJbVwWdkdPfI7nGsyskPHZAazpkn3zPduW00flwfBUicDniNsME7Prmldt75UIa1gy5KHuGfPQNyni2Ps

ret = client.Close(fd283)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd279, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (194) d2KG6QZ18bnxwE9l2HVd6dydDymtX8wKJeMyOvv8EGB_w441mF0zmUGvX2526aGGUrAPyjORXkHm1VyGqtDCsRVis0g96L22PH0ZYwCcdZ2uL0V4YnJbVwWdkdPfI7nGsyskPHZAazpkn3zPduW00flwfBUicDniNsME7Prmldt75UIa1gy5KHuGfPQNyni2Ps

ret = client.Write(fd282, []byte("EswQv7PpqQL76egUCM85czMXJ8JW8S7E__FCeGF29DX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (237) d2KG6QZ18bnxwE9l2HVd6dydDymtX8wKJeMyOvv8EGB_w441mF0zmUGvX2526aGGUrAPyjORXkHm1VyGqtDCsRVis0g96L22PH0ZYwCcdZ2uL0V4YnJbVwWdkdPfI7nGsyskPHZAazpkn3zPduW00flwfBUicDniNsME7Prmldt75UIa1gy5KHuGfPQNyni2PsEswQv7PpqQL76egUCM85czMXJ8JW8S7E__FCeGF29DX
//fd state: (0) 

ret = client.Write(fd280, []byte("MjQufC7V8zPluNeCr2vb4qFb_MQnXZCHzCqMDhluSecFp1S9iJojviJV_FgpOvUtxsPpbKMOPdeeTZ1p9YmBh__oKMZQzmT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) MjQufC7V8zPluNeCr2vb4qFb_MQnXZCHzCqMDhluSecFp1S9iJojviJV_FgpOvUtxsPpbKMOPdeeTZ1p9YmBh__oKMZQzmT

buf, ret = client.Read(fd279, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd285 := client.Open("/LoTeuE2wHL.txt", client.O_RDWR|client.O_CREATE)
if(fd285 < 0) {
  panic("open failed")
}


ret = client.Seek(fd284, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd271, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd271, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd281, []byte("2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) 2

ret = client.Seek(fd280, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


fd286 := client.Open("/w3GJXKziAw.txt", client.O_RDWR|client.O_CREATE)
if(fd286 < 0) {
  panic("open failed")
}


ret = client.Close(fd279)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd285, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Close(fd281)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd284, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (62) MjQufC7V8zPluNeCr2vb4qFb_MQnXZCHzCqMDhluSecFp1S9iJojviJV_FgpOvUtxsPpbKMOPdeeTZ1p9YmBh__oKMZQzmT

ret = client.Write(fd280, []byte("o55Rjzf1xEUAOUtOTRbAtdjpgGC5QfjgYKnRWzrN0PTm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (106) MjQufC7V8zPluNeCr2vb4qFb_MQnXZCHzCqMDhluSecFp1S9iJojviJV_FgpOvo55Rjzf1xEUAOUtOTRbAtdjpgGC5QfjgYKnRWzrN0PTm

ret = client.Seek(fd286, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd282, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


ret = client.Seek(fd286, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (13) vHad_07sUSPQY9oWG4Z6PgCXyOivTVGGGybzmONdNaqs7JP8weAuoX4mw1tc7p7FJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpq9Xy9Vb55endC1v_eRZ6HNjZPejHi0qZHFSlN_dhk7EljE

ret = client.Write(fd285, []byte("s3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpq9Xy9Vb55endC1v_eRZ6HNjZPejHi0qZHFSlN_dhk7EljE

ret = client.Close(fd282)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd284)
if(ret != 0) {
  panic("close failed")
}

//fd state: (64) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlJCyFmzcnLHvdBhHZw9E8kqR4zTEWT574C2aMJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpq9Xy9Vb55endC1v_eRZ6HNjZPejHi0qZHFSlN_dhk7EljE

ret = client.Write(fd285, []byte("D4ebxC0HsarSuX28KutOkoOU1jcJi_Ldb8RN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlD4ebxC0HsarSuX28KutOkoOU1jcJi_Ldb8RNJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpq9Xy9Vb55endC1v_eRZ6HNjZPejHi0qZHFSlN_dhk7EljE

buf, ret = client.Read(fd274, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd287 := client.Open("/2g8nvSsXFZ.txt", client.O_RDWR|client.O_CREATE)
if(fd287 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd287, []byte("Arf1bxOFJqakHXAPAcfXmJ0yMQ6dcnEhibz4P6E0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) Arf1bxOFJqakHXAPAcfXmJ0yMQ6dcnEhibz4P6E0
//fd state: (40) Arf1bxOFJqakHXAPAcfXmJ0yMQ6dcnEhibz4P6E0

ret = client.Write(fd287, []byte("xy9PMFa0_oTHnxT4aymvLZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) Arf1bxOFJqakHXAPAcfXmJ0yMQ6dcnEhibz4P6E0xy9PMFa0_oTHnxT4aymvLZ

buf, ret = client.Read(fd274, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd285, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "JfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7w") {
  panic("wrong data returned")
}

//fd state: (84) BQd_N0FuwM8QFB7YfCTBstNxsUyZMaFw3I099ALEEjh8H_5KTP_ziekMjXcCHQTrYQXT68zgoOv_1I5tlC4F

ret = client.Write(fd274, []byte("3U4NTqvg_qwgh0ZZg3yK9ULAgoH5Np54H9NPP4IbCF0KydPYNY2rlD0Z8TTYEHZIrB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) BQd_N0FuwM8QFB7YfCTBstNxsUyZMaFw3I099ALEEjh8H_5KTP_ziekMjXcCHQTrYQXT68zgoOv_1I5tlC4F3U4NTqvg_qwgh0ZZg3yK9ULAgoH5Np54H9NPP4IbCF0KydPYNY2rlD0Z8TTYEHZIrB

ret = client.Seek(fd285, 257, client.SEEK_SET)
if(ret != 257) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 257)
  panic("seek failed")
}


fd288 := client.Open("/su0syYXcBE.txt", client.O_RDWR|client.O_CREATE)
if(fd288 < 0) {
  panic("open failed")
}

//fd state: (150) BQd_N0FuwM8QFB7YfCTBstNxsUyZMaFw3I099ALEEjh8H_5KTP_ziekMjXcCHQTrYQXT68zgoOv_1I5tlC4F3U4NTqvg_qwgh0ZZg3yK9ULAgoH5Np54H9NPP4IbCF0KydPYNY2rlD0Z8TTYEHZIrB

ret = client.Write(fd274, []byte("UhQ8_1bNwd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (160) BQd_N0FuwM8QFB7YfCTBstNxsUyZMaFw3I099ALEEjh8H_5KTP_ziekMjXcCHQTrYQXT68zgoOv_1I5tlC4F3U4NTqvg_qwgh0ZZg3yK9ULAgoH5Np54H9NPP4IbCF0KydPYNY2rlD0Z8TTYEHZIrBUhQ8_1bNwd

ret = client.Close(fd287)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd273, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


fd289 := client.Open("/7bFIx4fJhh.txt", client.O_RDWR|client.O_CREATE)
if(fd289 < 0) {
  panic("open failed")
}


fd290 := client.Open("/Nhry0WoS3B.txt", client.O_RDWR|client.O_CREATE)
if(fd290 < 0) {
  panic("open failed")
}


fd291 := client.Open("/sfmM5ZOnpc.txt", client.O_RDWR|client.O_CREATE)
if(fd291 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd289, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd274)
if(ret != 0) {
  panic("close failed")
}


fd292 := client.Open("/iSviWTXnNc.txt", client.O_RDWR|client.O_CREATE)
if(fd292 < 0) {
  panic("open failed")
}


fd293 := client.Open("/3j6USh_Mur.txt", client.O_RDWR|client.O_CREATE)
if(fd293 < 0) {
  panic("open failed")
}


fd294 := client.Open("/0bepIOJ5hl.txt", client.O_RDWR|client.O_CREATE)
if(fd294 < 0) {
  panic("open failed")
}


ret = client.Close(fd286)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd271)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd294, 142, client.SEEK_SET)
if(ret != 142) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 142)
  panic("seek failed")
}


ret = client.Seek(fd294, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}

//fd state: (106) MjQufC7V8zPluNeCr2vb4qFb_MQnXZCHzCqMDhluSecFp1S9iJojviJV_FgpOvo55Rjzf1xEUAOUtOTRbAtdjpgGC5QfjgYKnRWzrN0PTm

ret = client.Write(fd280, []byte("4X0lPyrBsGWXcxrambBY5fF9pj1Tb4Tj7HTgdKqf2bzyQhbUynQ00IfUu_kk0nU7uuI9zBa4ADc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (181) MjQufC7V8zPluNeCr2vb4qFb_MQnXZCHzCqMDhluSecFp1S9iJojviJV_FgpOvo55Rjzf1xEUAOUtOTRbAtdjpgGC5QfjgYKnRWzrN0PTm4X0lPyrBsGWXcxrambBY5fF9pj1Tb4Tj7HTgdKqf2bzyQhbUynQ00IfUu_kk0nU7uuI9zBa4ADc

buf, ret = client.Read(fd294, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "s3V7W8A8qu5roJQ6_68mnwe8VSx3nTg9xXbYIyTnXeA") {
  panic("wrong data returned")
}


fd295 := client.Open("/9dluZ9EKXT.txt", client.O_RDWR|client.O_CREATE)
if(fd295 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd295, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd295, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) MjQufC7V8zPluNeCr2vb4qFb_MQnXZCHzCqMDhluSecFp1S9iJojviJV_FgpOvo55Rjzf1xEUAOUtOTRbAtdjpgGC5QfjgYKnRWzrN0PTm4X0lPyrBsGWXcxrambBY5fF9pj1Tb4Tj7HTgdKqf2bzyQhbUynQ00IfUu_kk0nU7uuI9zBa4ADc

ret = client.Write(fd291, []byte("Nrg7o"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) Nrg7oC7V8zPluNeCr2vb4qFb_MQnXZCHzCqMDhluSecFp1S9iJojviJV_FgpOvo55Rjzf1xEUAOUtOTRbAtdjpgGC5QfjgYKnRWzrN0PTm4X0lPyrBsGWXcxrambBY5fF9pj1Tb4Tj7HTgdKqf2bzyQhbUynQ00IfUu_kk0nU7uuI9zBa4ADc

ret = client.Seek(fd288, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


buf, ret = client.Read(fd273, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Gji1QUU0NGioxMQIb") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd293, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd295, []byte("Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu
//fd state: (0) 

ret = client.Write(fd293, []byte("2RY7STbQQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) 2RY7STbQQ

buf, ret = client.Read(fd280, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd296 := client.Open("/6O82q9cXp9.txt", client.O_RDWR|client.O_CREATE)
if(fd296 < 0) {
  panic("open failed")
}

//fd state: (257) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlD4ebxC0HsarSuX28KutOkoOU1jcJi_Ldb8RNJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpq9Xy9Vb55endC1v_eRZ6HNjZPejHi0qZHFSlN_dhk7EljE

ret = client.Write(fd285, []byte("DO2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (260) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlD4ebxC0HsarSuX28KutOkoOU1jcJi_Ldb8RNJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpqDO29Vb55endC1v_eRZ6HNjZPejHi0qZHFSlN_dhk7EljE

fd297 := client.Open("/KpCAEyYryL.txt", client.O_RDWR|client.O_CREATE)
if(fd297 < 0) {
  panic("open failed")
}


ret = client.Close(fd297)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd292, []byte("mgcyiF_3Z58FyfIAgS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) mgcyiF_3Z58FyfIAgS
//fd state: (260) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlD4ebxC0HsarSuX28KutOkoOU1jcJi_Ldb8RNJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpqDO29Vb55endC1v_eRZ6HNjZPejHi0qZHFSlN_dhk7EljE

ret = client.Write(fd285, []byte("OhKYGSKIHVR5hWoeN6wL2Pzd_LFtBTEEXboD6kXcJDAtlTXXcYD8cA1TdRtN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (320) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlD4ebxC0HsarSuX28KutOkoOU1jcJi_Ldb8RNJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpqDO2OhKYGSKIHVR5hWoeN6wL2Pzd_LFtBTEEXboD6kXcJDAtlTXXcYD8cA1TdRtN

fd298 := client.Open("/S9l3nzmqvI.txt", client.O_RDWR|client.O_CREATE)
if(fd298 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd289, 77)
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

//fd state: (42) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu

ret = client.Write(fd295, []byte("3Xa8LUU292eZBp84ZnmSvMayWnPIMR08U_9RDJ_pST6KaUNZkB_Pn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu3Xa8LUU292eZBp84ZnmSvMayWnPIMR08U_9RDJ_pST6KaUNZkB_Pn
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd288, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


ret = client.Close(fd294)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd292, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


fd299 := client.Open("/jUD4VX4_c1.txt", client.O_RDWR|client.O_CREATE)
if(fd299 < 0) {
  panic("open failed")
}

//fd state: (18) mgcyiF_3Z58FyfIAgS

ret = client.Write(fd292, []byte("s1DFNfi0BIDOAogT4A9B9avZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) mgcyiF_3Z58FyfIAgSs1DFNfi0BIDOAogT4A9B9avZ
//fd state: (0) y0ZkzMhkJKlBjZ8U2aLDOFjJLK72WwaF_zoANIMIA02NGrrvKOiuAIMUPwvrEdECZZxVdaUEFor9E1

ret = client.Write(fd296, []byte("ba0gK7K2gomflsYkB8MWg_fiSpAOeS7Vimch4MQav8kVMs5jMthbV1yCLdOIM5cLC3sgiYRgtlPZM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) ba0gK7K2gomflsYkB8MWg_fiSpAOeS7Vimch4MQav8kVMs5jMthbV1yCLdOIM5cLC3sgiYRgtlPZM1
//fd state: (22) 6qgCJlnFlWwI5JKVCLSKsHH1CTcpRnLOnIXZT82X0nVIVBczBImi

ret = client.Write(fd288, []byte("8hfaJOZdHq6k1jI5ggB1eipgo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) 6qgCJlnFlWwI5JKVCLSKsH8hfaJOZdHq6k1jI5ggB1eipgozBImi
//fd state: (9) 2RY7STbQQ

ret = client.Write(fd293, []byte("mRN3oE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) 2RY7STbQQmRN3oE
//fd state: (5) Nrg7oC7V8zPluNeCr2vb4qFb_MQnXZCHzCqMDhluSecFp1S9iJojviJV_FgpOvo55Rjzf1xEUAOUtOTRbAtdjpgGC5QfjgYKnRWzrN0PTm4X0lPyrBsGWXcxrambBY5fF9pj1Tb4Tj7HTgdKqf2bzyQhbUynQ00IfUu_kk0nU7uuI9zBa4ADc

ret = client.Write(fd291, []byte("CifGy6Pfsi8vFgU82wVeHNkEpN5jWqCHEphn5yk6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) Nrg7oCifGy6Pfsi8vFgU82wVeHNkEpN5jWqCHEphn5yk61S9iJojviJV_FgpOvo55Rjzf1xEUAOUtOTRbAtdjpgGC5QfjgYKnRWzrN0PTm4X0lPyrBsGWXcxrambBY5fF9pj1Tb4Tj7HTgdKqf2bzyQhbUynQ00IfUu_kk0nU7uuI9zBa4ADc

fd300 := client.Open("/UWgWPC4Yub.txt", client.O_RDWR|client.O_CREATE)
if(fd300 < 0) {
  panic("open failed")
}


ret = client.Close(fd298)
if(ret != 0) {
  panic("close failed")
}

//fd state: (320) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlD4ebxC0HsarSuX28KutOkoOU1jcJi_Ldb8RNJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpqDO2OhKYGSKIHVR5hWoeN6wL2Pzd_LFtBTEEXboD6kXcJDAtlTXXcYD8cA1TdRtN

ret = client.Write(fd285, []byte("shk8YIV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (327) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlD4ebxC0HsarSuX28KutOkoOU1jcJi_Ldb8RNJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpqDO2OhKYGSKIHVR5hWoeN6wL2Pzd_LFtBTEEXboD6kXcJDAtlTXXcYD8cA1TdRtNshk8YIV

buf, ret = client.Read(fd290, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "IMLhvmZLPRV8RZ4AE7d08y9KS7BqrHtmmbiwlqh4VpMRT0cjaWx2OwdRlXz") {
  panic("wrong data returned")
}


ret = client.Close(fd299)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd288)
if(ret != 0) {
  panic("close failed")
}


fd301 := client.Open("/eLMT0M0BqC.txt", client.O_RDWR|client.O_CREATE)
if(fd301 < 0) {
  panic("open failed")
}


ret = client.Seek(fd301, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd280)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd291)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd301)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd293, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Seek(fd273, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


ret = client.Seek(fd292, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


buf, ret = client.Read(fd293, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "RN3oE") {
  panic("wrong data returned")
}


ret = client.Close(fd273)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd300, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


ret = client.Close(fd300)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd295, 71, client.SEEK_SET)
if(ret != 71) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 71)
  panic("seek failed")
}


ret = client.Seek(fd296, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}


fd302 := client.Open("/5bNqIKtQYU.txt", client.O_RDWR|client.O_CREATE)
if(fd302 < 0) {
  panic("open failed")
}


ret = client.Seek(fd302, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (327) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlD4ebxC0HsarSuX28KutOkoOU1jcJi_Ldb8RNJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpqDO2OhKYGSKIHVR5hWoeN6wL2Pzd_LFtBTEEXboD6kXcJDAtlTXXcYD8cA1TdRtNshk8YIV

ret = client.Write(fd285, []byte("pGIsWVWUSfrnHCDpY_8hlT59cNSzfkbFAxo7DEg9fvm0AhZfWx093u9vdtp_Mau"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (390) vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6wIQhFNKZfZyvYjlk6JrIASlD4ebxC0HsarSuX28KutOkoOU1jcJi_Ldb8RNJfVsJ9DU1k2m7DV0Sd44lzcgvm0HPJV9DKFBoKXQlZYYE8Bcobe7Uy6ve9MSnPyMHwaOSze6Z1IbyfdidUdH7wFY5IollSbIHNPdt3iCq3kf18UCo5_y6lrxuV7lDPZcZh0i1k9LxrKMILZZlixbOw6KuGQpqDO2OhKYGSKIHVR5hWoeN6wL2Pzd_LFtBTEEXboD6kXcJDAtlTXXcYD8cA1TdRtNshk8YIVpGIsWVWUSfrnHCDpY_8hlT59cNSzfkbFAxo7DEg9fvm0AhZfWx093u9vdtp_Mau

fd303 := client.Open("/1uvG1cy_pv.txt", client.O_RDWR|client.O_CREATE)
if(fd303 < 0) {
  panic("open failed")
}

//fd state: (44) ba0gK7K2gomflsYkB8MWg_fiSpAOeS7Vimch4MQav8kVMs5jMthbV1yCLdOIM5cLC3sgiYRgtlPZM1

ret = client.Write(fd296, []byte("rJeGxDNb6hyVvMvU1fJr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) ba0gK7K2gomflsYkB8MWg_fiSpAOeS7Vimch4MQav8kVrJeGxDNb6hyVvMvU1fJrC3sgiYRgtlPZM1

buf, ret = client.Read(fd303, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd293)
if(ret != 0) {
  panic("close failed")
}


fd304 := client.Open("/GRTJLmJyYC.txt", client.O_RDWR|client.O_CREATE)
if(fd304 < 0) {
  panic("open failed")
}


ret = client.Close(fd303)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd304, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (16) mgcyiF_3Z58FyfIAgSs1DFNfi0BIDOAogT4A9B9avZ

ret = client.Write(fd292, []byte("f0O4URnFrwbvKR57D9iRhYgs53G6yaWbj261ixhmbDsvGtBW1JG6qh68M5QagXvh1cYXvAxt23eTzuPCN9govVban5ICA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (109) mgcyiF_3Z58FyfIAf0O4URnFrwbvKR57D9iRhYgs53G6yaWbj261ixhmbDsvGtBW1JG6qh68M5QagXvh1cYXvAxt23eTzuPCN9govVban5ICA

ret = client.Seek(fd304, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd305 := client.Open("/zFmjACN5Xo.txt", client.O_RDWR|client.O_CREATE)
if(fd305 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd302, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "BQd_N0FuwM8QFB7YfCTBstNxsUyZMaFw3I099ALE") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd296, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "C3sgiYRgtlPZM1") {
  panic("wrong data returned")
}

//fd state: (40) BQd_N0FuwM8QFB7YfCTBstNxsUyZMaFw3I099ALEEjh8H_5KTP_ziekMjXcCHQTrYQXT68zgoOv_1I5tlC4F3U4NTqvg_qwgh0ZZg3yK9ULAgoH5Np54H9NPP4IbCF0KydPYNY2rlD0Z8TTYEHZIrBUhQ8_1bNwd

ret = client.Write(fd302, []byte("zXOCPosS1SFxYlqrVrvZk4yc4xcXs38KdBhgEKMPIaFGFUbQpVbXA_PdarBBuTjUgVbygEvGrrkTSECWLAOa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) BQd_N0FuwM8QFB7YfCTBstNxsUyZMaFw3I099ALEzXOCPosS1SFxYlqrVrvZk4yc4xcXs38KdBhgEKMPIaFGFUbQpVbXA_PdarBBuTjUgVbygEvGrrkTSECWLAOaCF0KydPYNY2rlD0Z8TTYEHZIrBUhQ8_1bNwd

ret = client.Close(fd304)
if(ret != 0) {
  panic("close failed")
}

//fd state: (124) BQd_N0FuwM8QFB7YfCTBstNxsUyZMaFw3I099ALEzXOCPosS1SFxYlqrVrvZk4yc4xcXs38KdBhgEKMPIaFGFUbQpVbXA_PdarBBuTjUgVbygEvGrrkTSECWLAOaCF0KydPYNY2rlD0Z8TTYEHZIrBUhQ8_1bNwd

ret = client.Write(fd302, []byte("YSLC1vy0Nx6XchHb7yV48LJIxpYGfUpC8ltZMuyEXR2nB4iJyUGtoXn3ehyhP9uw9X2GNo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (194) BQd_N0FuwM8QFB7YfCTBstNxsUyZMaFw3I099ALEzXOCPosS1SFxYlqrVrvZk4yc4xcXs38KdBhgEKMPIaFGFUbQpVbXA_PdarBBuTjUgVbygEvGrrkTSECWLAOaYSLC1vy0Nx6XchHb7yV48LJIxpYGfUpC8ltZMuyEXR2nB4iJyUGtoXn3ehyhP9uw9X2GNo

ret = client.Seek(fd302, 166, client.SEEK_SET)
if(ret != 166) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 166)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd305, []byte("6W4DTiluEoFiWResVd4MhN1OnsaQipn1UxhSNYtTLJYSaEqBmVX9Q6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) 6W4DTiluEoFiWResVd4MhN1OnsaQipn1UxhSNYtTLJYSaEqBmVX9Q6
//fd state: (78) ba0gK7K2gomflsYkB8MWg_fiSpAOeS7Vimch4MQav8kVrJeGxDNb6hyVvMvU1fJrC3sgiYRgtlPZM1

ret = client.Write(fd296, []byte("4nWccmFKgrGX4CHG5cV6wxOivAL0aAZCx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) ba0gK7K2gomflsYkB8MWg_fiSpAOeS7Vimch4MQav8kVrJeGxDNb6hyVvMvU1fJrC3sgiYRgtlPZM14nWccmFKgrGX4CHG5cV6wxOivAL0aAZCx

fd306 := client.Open("/7bFIx4fJhh.txt", client.O_RDWR|client.O_CREATE)
if(fd306 < 0) {
  panic("open failed")
}


fd307 := client.Open("/5OR5HN2nWE.txt", client.O_RDWR|client.O_CREATE)
if(fd307 < 0) {
  panic("open failed")
}


fd308 := client.Open("/Z7sEO5NmmB.txt", client.O_RDWR|client.O_CREATE)
if(fd308 < 0) {
  panic("open failed")
}


ret = client.Seek(fd302, 127, client.SEEK_SET)
if(ret != 127) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 127)
  panic("seek failed")
}


fd309 := client.Open("/6K38SxNJjF.txt", client.O_RDWR|client.O_CREATE)
if(fd309 < 0) {
  panic("open failed")
}


ret = client.Close(fd285)
if(ret != 0) {
  panic("close failed")
}

//fd state: (54) 6W4DTiluEoFiWResVd4MhN1OnsaQipn1UxhSNYtTLJYSaEqBmVX9Q6

ret = client.Write(fd305, []byte("3n9OVdGBZR5cGjPzk7uYD6OVAPTTr178tx2qxHlMP4IulF0fjvEnrBI_YPE_7FQOjiLpJY0Ql1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (128) 6W4DTiluEoFiWResVd4MhN1OnsaQipn1UxhSNYtTLJYSaEqBmVX9Q63n9OVdGBZR5cGjPzk7uYD6OVAPTTr178tx2qxHlMP4IulF0fjvEnrBI_YPE_7FQOjiLpJY0Ql1

buf, ret = client.Read(fd292, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd302)
if(ret != 0) {
  panic("close failed")
}

//fd state: (71) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu3Xa8LUU292eZBp84ZnmSvMayWnPIMR08U_9RDJ_pST6KaUNZkB_Pn

ret = client.Write(fd295, []byte("AC25nm5j9Mp555cYTfWLfJ4EaCL16wmxxiTsTOfWGJ4Td9zvQbTg7xnHHUTJN4YEblSj8qp9zXD0xUDjIqYWP1sJvWA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (162) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu3Xa8LUU292eZBp84ZnmSvMayWnPIMAC25nm5j9Mp555cYTfWLfJ4EaCL16wmxxiTsTOfWGJ4Td9zvQbTg7xnHHUTJN4YEblSj8qp9zXD0xUDjIqYWP1sJvWA

buf, ret = client.Read(fd292, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd310 := client.Open("/kwPisEKdpD.txt", client.O_RDWR|client.O_CREATE)
if(fd310 < 0) {
  panic("open failed")
}


fd311 := client.Open("/PQ99xNzlFm.txt", client.O_RDWR|client.O_CREATE)
if(fd311 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd292, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd308)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd311)
if(ret != 0) {
  panic("close failed")
}


fd312 := client.Open("/5nEDDggz8U.txt", client.O_RDWR|client.O_CREATE)
if(fd312 < 0) {
  panic("open failed")
}

//fd state: (162) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu3Xa8LUU292eZBp84ZnmSvMayWnPIMAC25nm5j9Mp555cYTfWLfJ4EaCL16wmxxiTsTOfWGJ4Td9zvQbTg7xnHHUTJN4YEblSj8qp9zXD0xUDjIqYWP1sJvWA

ret = client.Write(fd295, []byte("sqBwsSLcrAn6rdXsYk99PJzu4AqPS50NvIFKtfdf9YO3Ug59ZuOXqE0kpFMbSXnrQIZwHvF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (233) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu3Xa8LUU292eZBp84ZnmSvMayWnPIMAC25nm5j9Mp555cYTfWLfJ4EaCL16wmxxiTsTOfWGJ4Td9zvQbTg7xnHHUTJN4YEblSj8qp9zXD0xUDjIqYWP1sJvWAsqBwsSLcrAn6rdXsYk99PJzu4AqPS50NvIFKtfdf9YO3Ug59ZuOXqE0kpFMbSXnrQIZwHvF
//fd state: (0) 

ret = client.Write(fd309, []byte("vHT4FWPa17GZJ7LoymjdDosRgf61wgrQf5UYt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) vHT4FWPa17GZJ7LoymjdDosRgf61wgrQf5UYt
//fd state: (0) cBO3dMimpOQG1t4YuYXrF_dVDs2YhLi3WWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5SqheAg0blDhIDsZyXmgpJF07VGfEMha6ECPuveO9v7SJvOeoiC6Ic9Q7NbQeEPYzq_59soWKbDOBlI0FXJyR5pGQ77juWJqXqefAoNEVNy9ObQED7MEXq30KgSeTCuBSL

ret = client.Write(fd307, []byte("LYQZxfFpKRCOEQS3_YxcfogdMTWLG5n1O"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) LYQZxfFpKRCOEQS3_YxcfogdMTWLG5n1OWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5SqheAg0blDhIDsZyXmgpJF07VGfEMha6ECPuveO9v7SJvOeoiC6Ic9Q7NbQeEPYzq_59soWKbDOBlI0FXJyR5pGQ77juWJqXqefAoNEVNy9ObQED7MEXq30KgSeTCuBSL

buf, ret = client.Read(fd292, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd313 := client.Open("/WDl3LmtROa.txt", client.O_RDWR|client.O_CREATE)
if(fd313 < 0) {
  panic("open failed")
}


fd314 := client.Open("/b6UqNkmZMn.txt", client.O_RDWR|client.O_CREATE)
if(fd314 < 0) {
  panic("open failed")
}


ret = client.Seek(fd296, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


ret = client.Close(fd306)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd292, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


buf, ret = client.Read(fd313, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) uV9jOsW95BFP1d0k_axYKftzBQ44_JEhTi3Xg_4UqMqnjRKZYxMUOC9lbIlgK7XN1eTpgfr2YzBVU63LBieapKPmVe0GKFYq1zLHg

ret = client.Write(fd312, []byte("u1iIRuH5MPy9xDVyOvWboeKqI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) u1iIRuH5MPy9xDVyOvWboeKqIQ44_JEhTi3Xg_4UqMqnjRKZYxMUOC9lbIlgK7XN1eTpgfr2YzBVU63LBieapKPmVe0GKFYq1zLHg

ret = client.Close(fd314)
if(ret != 0) {
  panic("close failed")
}

//fd state: (233) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu3Xa8LUU292eZBp84ZnmSvMayWnPIMAC25nm5j9Mp555cYTfWLfJ4EaCL16wmxxiTsTOfWGJ4Td9zvQbTg7xnHHUTJN4YEblSj8qp9zXD0xUDjIqYWP1sJvWAsqBwsSLcrAn6rdXsYk99PJzu4AqPS50NvIFKtfdf9YO3Ug59ZuOXqE0kpFMbSXnrQIZwHvF

ret = client.Write(fd295, []byte("W5OJ1d7gL6VNvlfXuKCTxvHF2ktW2AwrGosgl_pDZ2BSES3ZAVIcC1nuiZ2GASbYZX1K08qYcdkWNK1Znmy0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (317) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu3Xa8LUU292eZBp84ZnmSvMayWnPIMAC25nm5j9Mp555cYTfWLfJ4EaCL16wmxxiTsTOfWGJ4Td9zvQbTg7xnHHUTJN4YEblSj8qp9zXD0xUDjIqYWP1sJvWAsqBwsSLcrAn6rdXsYk99PJzu4AqPS50NvIFKtfdf9YO3Ug59ZuOXqE0kpFMbSXnrQIZwHvFW5OJ1d7gL6VNvlfXuKCTxvHF2ktW2AwrGosgl_pDZ2BSES3ZAVIcC1nuiZ2GASbYZX1K08qYcdkWNK1Znmy0

ret = client.Close(fd313)
if(ret != 0) {
  panic("close failed")
}


fd315 := client.Open("/5FaFxgLgRu.txt", client.O_RDWR|client.O_CREATE)
if(fd315 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd290, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "h") {
  panic("wrong data returned")
}


fd316 := client.Open("/ModFGCjTvL.txt", client.O_RDWR|client.O_CREATE)
if(fd316 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd292, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yfIAf") {
  panic("wrong data returned")
}


ret = client.Close(fd296)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd316, []byte("R5keHUpQVdut1RMooVQdXAKw94Z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) R5keHUpQVdut1RMooVQdXAKw94Z

ret = client.Seek(fd305, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}


ret = client.Seek(fd315, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd295, 248, client.SEEK_SET)
if(ret != 248) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 248)
  panic("seek failed")
}


fd317 := client.Open("/2xFcjyIHJh.txt", client.O_RDWR|client.O_CREATE)
if(fd317 < 0) {
  panic("open failed")
}


ret = client.Seek(fd307, 277, client.SEEK_SET)
if(ret != 277) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 277)
  panic("seek failed")
}


fd318 := client.Open("/TDOT533S0s.txt", client.O_RDWR|client.O_CREATE)
if(fd318 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd290, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jypQxVIg9URhpIw2lEM") {
  panic("wrong data returned")
}


fd319 := client.Open("/ALsGXybpGN.txt", client.O_RDWR|client.O_CREATE)
if(fd319 < 0) {
  panic("open failed")
}


ret = client.Close(fd292)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd315, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd305, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "n1UxhSNYtTLJ") {
  panic("wrong data returned")
}


ret = client.Close(fd295)
if(ret != 0) {
  panic("close failed")
}


fd320 := client.Open("/7bFIx4fJhh.txt", client.O_RDWR|client.O_CREATE)
if(fd320 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (277) LYQZxfFpKRCOEQS3_YxcfogdMTWLG5n1OWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5SqheAg0blDhIDsZyXmgpJF07VGfEMha6ECPuveO9v7SJvOeoiC6Ic9Q7NbQeEPYzq_59soWKbDOBlI0FXJyR5pGQ77juWJqXqefAoNEVNy9ObQED7MEXq30KgSeTCuBSL

ret = client.Write(fd307, []byte("VjcbXt8SithBPuaYLaGnNvsH42ksnRSs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (309) LYQZxfFpKRCOEQS3_YxcfogdMTWLG5n1OWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5SqheAg0blDhIDsZyXmgpJF07VGfEMha6ECPuveO9v7SJvOeoiC6Ic9Q7NbQeEPYzq_59soWKbDOBlI0FXJyR5pGQ77juWJqXqefAoNEVNy9ObQED7MVjcbXt8SithBPuaYLaGnNvsH42ksnRSs

buf, ret = client.Read(fd307, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd290, 144, client.SEEK_SET)
if(ret != 144) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 144)
  panic("seek failed")
}


buf, ret = client.Read(fd305, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YSaEqBmVX9Q63n9OVdGBZR5cGjPzk7uYD6OVAPTTr178tx2qxHlMP4IulF0fjvEn") {
  panic("wrong data returned")
}


fd321 := client.Open("/L68IVOUnGL.txt", client.O_RDWR|client.O_CREATE)
if(fd321 < 0) {
  panic("open failed")
}


ret = client.Seek(fd317, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd317, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd322 := client.Open("/YyhcUiTFU4.txt", client.O_RDWR|client.O_CREATE)
if(fd322 < 0) {
  panic("open failed")
}


ret = client.Seek(fd316, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


ret = client.Close(fd316)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd317, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd322, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd315, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (144) IMLhvmZLPRV8RZ4AE7d08y9KS7BqrHtmmbiwlqh4VpMRT0cjaWx2OwdRlXzhjypQxVIg9URhpIw2lEMjogQzZcAWlaCyV0Ny8aKx13HbIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF0rJYL4yE0yUku2_I

ret = client.Write(fd290, []byte("_a1R44rashdeWGTya4xI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) IMLhvmZLPRV8RZ4AE7d08y9KS7BqrHtmmbiwlqh4VpMRT0cjaWx2OwdRlXzhjypQxVIg9URhpIw2lEMjogQzZcAWlaCyV0Ny8aKx13HbIOnDY7qSkVLs2aHY3YzChgaGpxNaiV5EjMXQYDEF_a1R44rashdeWGTya4xI

ret = client.Close(fd318)
if(ret != 0) {
  panic("close failed")
}

//fd state: (309) LYQZxfFpKRCOEQS3_YxcfogdMTWLG5n1OWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5SqheAg0blDhIDsZyXmgpJF07VGfEMha6ECPuveO9v7SJvOeoiC6Ic9Q7NbQeEPYzq_59soWKbDOBlI0FXJyR5pGQ77juWJqXqefAoNEVNy9ObQED7MVjcbXt8SithBPuaYLaGnNvsH42ksnRSs

ret = client.Write(fd307, []byte("hH6l9Gvqmaxi0bSYrHiiiLEDx5yukpup6Anv1YgtbtQq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (353) LYQZxfFpKRCOEQS3_YxcfogdMTWLG5n1OWagQvWSWg1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDWeDCQXKNjchslB6MZpepwXJP5SqheAg0blDhIDsZyXmgpJF07VGfEMha6ECPuveO9v7SJvOeoiC6Ic9Q7NbQeEPYzq_59soWKbDOBlI0FXJyR5pGQ77juWJqXqefAoNEVNy9ObQED7MVjcbXt8SithBPuaYLaGnNvsH42ksnRSshH6l9Gvqmaxi0bSYrHiiiLEDx5yukpup6Anv1YgtbtQq

fd323 := client.Open("/sZAyeiVvZ5.txt", client.O_RDWR|client.O_CREATE)
if(fd323 < 0) {
  panic("open failed")
}


ret = client.Seek(fd310, 90, client.SEEK_SET)
if(ret != 90) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 90)
  panic("seek failed")
}


fd324 := client.Open("/YyURkaOMEI.txt", client.O_RDWR|client.O_CREATE)
if(fd324 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd307, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd322, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd321, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd310)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd312, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Q44_JEhTi3Xg_4UqMqnjRKZYxMUOC9lbI") {
  panic("wrong data returned")
}


ret = client.Seek(fd305, 92, client.SEEK_SET)
if(ret != 92) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 92)
  panic("seek failed")
}


fd325 := client.Open("/BTpMjqSUbn.txt", client.O_RDWR|client.O_CREATE)
if(fd325 < 0) {
  panic("open failed")
}


ret = client.Close(fd320)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd317, []byte("Lfh04iS8FPFBhIm5PUQUDyNQrc_15R6B_W6E5L0yw1t_Xik9GNj5hqR2yiAp9F6dqreoDz6kU3P2o7glYclI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) Lfh04iS8FPFBhIm5PUQUDyNQrc_15R6B_W6E5L0yw1t_Xik9GNj5hqR2yiAp9F6dqreoDz6kU3P2o7glYclI

ret = client.Seek(fd307, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


fd326 := client.Open("/6HMIdOkz5H.txt", client.O_RDWR|client.O_CREATE)
if(fd326 < 0) {
  panic("open failed")
}

//fd state: (84) Lfh04iS8FPFBhIm5PUQUDyNQrc_15R6B_W6E5L0yw1t_Xik9GNj5hqR2yiAp9F6dqreoDz6kU3P2o7glYclI

ret = client.Write(fd317, []byte("bHPk7H0njwgeTx6BTrhi4URdZmDg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (112) Lfh04iS8FPFBhIm5PUQUDyNQrc_15R6B_W6E5L0yw1t_Xik9GNj5hqR2yiAp9F6dqreoDz6kU3P2o7glYclIbHPk7H0njwgeTx6BTrhi4URdZmDg
//fd state: (0) 

ret = client.Write(fd322, []byte("QJv2H3RdVSn8imaCN0LQ77vtTfiaeYGkxpJNlAbuW00_sEnUSqPTTtgpPBOw855snZLYu4Rl5mO4uKjeN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) QJv2H3RdVSn8imaCN0LQ77vtTfiaeYGkxpJNlAbuW00_sEnUSqPTTtgpPBOw855snZLYu4Rl5mO4uKjeN

buf, ret = client.Read(fd309, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd315, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd327 := client.Open("/q4ub60oZmc.txt", client.O_RDWR|client.O_CREATE)
if(fd327 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd307, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "g1a0u3yBRBf8iutTfQjV_BBZKqO86pq27oSzBr5f_JV_9DGZQC3desDlkxrMUXzzgRbcCaK2_xhz7VyhMwYLUrbqIYpNLVyEDW") {
  panic("wrong data returned")
}


ret = client.Close(fd290)
if(ret != 0) {
  panic("close failed")
}

//fd state: (92) 6W4DTiluEoFiWResVd4MhN1OnsaQipn1UxhSNYtTLJYSaEqBmVX9Q63n9OVdGBZR5cGjPzk7uYD6OVAPTTr178tx2qxHlMP4IulF0fjvEnrBI_YPE_7FQOjiLpJY0Ql1

ret = client.Write(fd305, []byte("PWZ2kg897daKYrfUVrOi9zkMOWbBuNtt8VFaWEMllflBJDB0gXcDsxAtColxmvgJbdLQzndaE09476ExLV2P_eAgUgZVZAVI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (188) 6W4DTiluEoFiWResVd4MhN1OnsaQipn1UxhSNYtTLJYSaEqBmVX9Q63n9OVdGBZR5cGjPzk7uYD6OVAPTTr178tx2qxHPWZ2kg897daKYrfUVrOi9zkMOWbBuNtt8VFaWEMllflBJDB0gXcDsxAtColxmvgJbdLQzndaE09476ExLV2P_eAgUgZVZAVI

ret = client.Seek(fd315, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd307, 206, client.SEEK_SET)
if(ret != 206) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 206)
  panic("seek failed")
}

//fd state: (0) 58YyfutV2MO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsTDVtIse5fp9FLC7S24jUK3P966HyhdKi4ztI6W

ret = client.Write(fd327, []byte("DUtt6FRvZU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) DUtt6FRvZUO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsTDVtIse5fp9FLC7S24jUK3P966HyhdKi4ztI6W

buf, ret = client.Read(fd317, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd315, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd324, []byte("vBYLE_MHtW6vf6rP3WbPgQkCwdusxL4_r6PTkfOitGT_0DU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) vBYLE_MHtW6vf6rP3WbPgQkCwdusxL4_r6PTkfOitGT_0DU
//fd state: (58) u1iIRuH5MPy9xDVyOvWboeKqIQ44_JEhTi3Xg_4UqMqnjRKZYxMUOC9lbIlgK7XN1eTpgfr2YzBVU63LBieapKPmVe0GKFYq1zLHg

ret = client.Write(fd312, []byte("5ZlDYmDRPRGPz6iLIjmXTWoAP08MUGzPx93q0bHjbbOO6nEaDQvDhD9KZV6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) u1iIRuH5MPy9xDVyOvWboeKqIQ44_JEhTi3Xg_4UqMqnjRKZYxMUOC9lbI5ZlDYmDRPRGPz6iLIjmXTWoAP08MUGzPx93q0bHjbbOO6nEaDQvDhD9KZV6

fd328 := client.Open("/o3Jy8jBnOV.txt", client.O_RDWR|client.O_CREATE)
if(fd328 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd323, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd319)
if(ret != 0) {
  panic("close failed")
}


fd329 := client.Open("/hF6OAOMNoP.txt", client.O_RDWR|client.O_CREATE)
if(fd329 < 0) {
  panic("open failed")
}


fd330 := client.Open("/HmduvtDGQA.txt", client.O_RDWR|client.O_CREATE)
if(fd330 < 0) {
  panic("open failed")
}


ret = client.Seek(fd323, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd307)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd325, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd331 := client.Open("/PWtXLMuKMR.txt", client.O_RDWR|client.O_CREATE)
if(fd331 < 0) {
  panic("open failed")
}


fd332 := client.Open("/DRIJ80NCqo.txt", client.O_RDWR|client.O_CREATE)
if(fd332 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd323, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd317, 65, client.SEEK_SET)
if(ret != 65) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 65)
  panic("seek failed")
}


ret = client.Seek(fd325, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd325)
if(ret != 0) {
  panic("close failed")
}


fd333 := client.Open("/IgWgfliKnz.txt", client.O_RDWR|client.O_CREATE)
if(fd333 < 0) {
  panic("open failed")
}


ret = client.Close(fd312)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd330)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd323, 25)
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


buf, ret = client.Read(fd305, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd315)
if(ret != 0) {
  panic("close failed")
}


fd334 := client.Open("/jQdl4Skatz.txt", client.O_RDWR|client.O_CREATE)
if(fd334 < 0) {
  panic("open failed")
}


fd335 := client.Open("/V2hqe83j8X.txt", client.O_RDWR|client.O_CREATE)
if(fd335 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd333, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "64vkLGSsN") {
  panic("wrong data returned")
}


ret = client.Seek(fd327, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Close(fd333)
if(ret != 0) {
  panic("close failed")
}


fd336 := client.Open("/tnVmsNSjJy.txt", client.O_RDWR|client.O_CREATE)
if(fd336 < 0) {
  panic("open failed")
}


ret = client.Seek(fd327, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


ret = client.Seek(fd324, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


fd337 := client.Open("/HjYfMWqq3p.txt", client.O_RDWR|client.O_CREATE)
if(fd337 < 0) {
  panic("open failed")
}


fd338 := client.Open("/LI4HznWdue.txt", client.O_RDWR|client.O_CREATE)
if(fd338 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd322, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd328, []byte("2oDyiI2uj5hQ5SRKaY6bMJApelpQxZCu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) 2oDyiI2uj5hQ5SRKaY6bMJApelpQxZCu

fd339 := client.Open("/Ku8Le4flu7.txt", client.O_RDWR|client.O_CREATE)
if(fd339 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd335, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd336)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd335, []byte("ojMm3KphHvyOUztZl5HW0NHtMAiggvU56l_ri5ALj5YiYWCYLP9jVa9aczM4U5fACE8lYOSTwxZRUGzb1gnqsziLzk2cV5g2q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) ojMm3KphHvyOUztZl5HW0NHtMAiggvU56l_ri5ALj5YiYWCYLP9jVa9aczM4U5fACE8lYOSTwxZRUGzb1gnqsziLzk2cV5g2q

buf, ret = client.Read(fd323, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd321, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (54) DUtt6FRvZUO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsTDVtIse5fp9FLC7S24jUK3P966HyhdKi4ztI6W

ret = client.Write(fd327, []byte("9jQmcAd2B_jOnzjovFZCdLZFTYsg8Ak1erqrchBNgQHs41abtjUye9jEvskWpxQ9NnVSOUyKWMms4e"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) DUtt6FRvZUO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsT9jQmcAd2B_jOnzjovFZCdLZFTYsg8Ak1erqrchBNgQHs41abtjUye9jEvskWpxQ9NnVSOUyKWMms4e

ret = client.Seek(fd332, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd328, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd321, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd339, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd338, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd340 := client.Open("/JcnQY5oHjy.txt", client.O_RDWR|client.O_CREATE)
if(fd340 < 0) {
  panic("open failed")
}


ret = client.Close(fd321)
if(ret != 0) {
  panic("close failed")
}


fd341 := client.Open("/r8EPWv0yyL.txt", client.O_RDWR|client.O_CREATE)
if(fd341 < 0) {
  panic("open failed")
}


ret = client.Seek(fd332, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd322, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd342 := client.Open("/an0N4p7hQP.txt", client.O_RDWR|client.O_CREATE)
if(fd342 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd334, []byte("kyG86np"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) kyG86np

fd343 := client.Open("/eLMT0M0BqC.txt", client.O_RDWR|client.O_CREATE)
if(fd343 < 0) {
  panic("open failed")
}


fd344 := client.Open("/EPIBJZLW_w.txt", client.O_RDWR|client.O_CREATE)
if(fd344 < 0) {
  panic("open failed")
}


ret = client.Seek(fd341, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd332, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd343)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd342, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd344, []byte("5NapMG4B3oXg5jUu6OyqzQwDyD_t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) 5NapMG4B3oXg5jUu6OyqzQwDyD_t
//fd state: (0) 

ret = client.Write(fd332, []byte("kYdSWh8W8rbZLJsXgal"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) kYdSWh8W8rbZLJsXgal

buf, ret = client.Read(fd327, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd305)
if(ret != 0) {
  panic("close failed")
}

//fd state: (7) kyG86np

ret = client.Write(fd334, []byte("YjulZIg2oIDIXkaFtQyI3wQX2pBem1vI3gt29Bw9Q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) kyG86npYjulZIg2oIDIXkaFtQyI3wQX2pBem1vI3gt29Bw9Q
//fd state: (48) kyG86npYjulZIg2oIDIXkaFtQyI3wQX2pBem1vI3gt29Bw9Q

ret = client.Write(fd334, []byte("aiBVZPk7Q5TsPjSNSQADlf1iJc3mRUCoUh5_zkn_z8G2eE0gARXaQf9j88iPjZZsyN5Y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) kyG86npYjulZIg2oIDIXkaFtQyI3wQX2pBem1vI3gt29Bw9QaiBVZPk7Q5TsPjSNSQADlf1iJc3mRUCoUh5_zkn_z8G2eE0gARXaQf9j88iPjZZsyN5Y

fd345 := client.Open("/o3Jy8jBnOV.txt", client.O_RDWR|client.O_CREATE)
if(fd345 < 0) {
  panic("open failed")
}


ret = client.Seek(fd309, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


buf, ret = client.Read(fd329, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5z") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd340, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd329, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "o0vpHHm") {
  panic("wrong data returned")
}


ret = client.Seek(fd345, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}

//fd state: (28) 5NapMG4B3oXg5jUu6OyqzQwDyD_t

ret = client.Write(fd344, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) 5NapMG4B3oXg5jUu6OyqzQwDyD_t

ret = client.Seek(fd327, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


ret = client.Seek(fd338, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd332)
if(ret != 0) {
  panic("close failed")
}


fd346 := client.Open("/bw8gv_bi38.txt", client.O_RDWR|client.O_CREATE)
if(fd346 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd335, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (116) kyG86npYjulZIg2oIDIXkaFtQyI3wQX2pBem1vI3gt29Bw9QaiBVZPk7Q5TsPjSNSQADlf1iJc3mRUCoUh5_zkn_z8G2eE0gARXaQf9j88iPjZZsyN5Y

ret = client.Write(fd334, []byte("pMJYRbFh86g75eW0Jw5uqO_qXr8KJ0bXIV46YyULj6EYDFJwPES_PPyXJtTWWvLxyrTeJpzyX2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (190) kyG86npYjulZIg2oIDIXkaFtQyI3wQX2pBem1vI3gt29Bw9QaiBVZPk7Q5TsPjSNSQADlf1iJc3mRUCoUh5_zkn_z8G2eE0gARXaQf9j88iPjZZsyN5YpMJYRbFh86g75eW0Jw5uqO_qXr8KJ0bXIV46YyULj6EYDFJwPES_PPyXJtTWWvLxyrTeJpzyX2
//fd state: (51) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmSX3cbAMLQGcDN0le8SC6AQR8zKkGdwMz6o7b5LzwD_Ey6bSx9OlJuNfMYz0e_B2rEitL4XDXshYwa9EwCMe

ret = client.Write(fd329, []byte("w4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6pGb4t9b9tXGQwnzrgIBVclz1KMzIpNDhokC7UDRNd2A5z6cYzOgQE9h"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6pGb4t9b9tXGQwnzrgIBVclz1KMzIpNDhokC7UDRNd2A5z6cYzOgQE9h

ret = client.Seek(fd341, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd309, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


fd347 := client.Open("/WNMJT2E_a5.txt", client.O_RDWR|client.O_CREATE)
if(fd347 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd338, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd344)
if(ret != 0) {
  panic("close failed")
}


fd348 := client.Open("/uZNj2PihlI.txt", client.O_RDWR|client.O_CREATE)
if(fd348 < 0) {
  panic("open failed")
}


fd349 := client.Open("/2uQvu09n7Y.txt", client.O_RDWR|client.O_CREATE)
if(fd349 < 0) {
  panic("open failed")
}


ret = client.Close(fd326)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd342, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd350 := client.Open("/KIwtxrfoHo.txt", client.O_RDWR|client.O_CREATE)
if(fd350 < 0) {
  panic("open failed")
}


ret = client.Seek(fd347, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


buf, ret = client.Read(fd337, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd350, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd340)
if(ret != 0) {
  panic("close failed")
}


fd351 := client.Open("/z6HJN7iNRF.txt", client.O_RDWR|client.O_CREATE)
if(fd351 < 0) {
  panic("open failed")
}

//fd state: (190) kyG86npYjulZIg2oIDIXkaFtQyI3wQX2pBem1vI3gt29Bw9QaiBVZPk7Q5TsPjSNSQADlf1iJc3mRUCoUh5_zkn_z8G2eE0gARXaQf9j88iPjZZsyN5YpMJYRbFh86g75eW0Jw5uqO_qXr8KJ0bXIV46YyULj6EYDFJwPES_PPyXJtTWWvLxyrTeJpzyX2

ret = client.Write(fd334, []byte("I1lcJNnY_hR4b75Vg71Sqgu106j"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (217) kyG86npYjulZIg2oIDIXkaFtQyI3wQX2pBem1vI3gt29Bw9QaiBVZPk7Q5TsPjSNSQADlf1iJc3mRUCoUh5_zkn_z8G2eE0gARXaQf9j88iPjZZsyN5YpMJYRbFh86g75eW0Jw5uqO_qXr8KJ0bXIV46YyULj6EYDFJwPES_PPyXJtTWWvLxyrTeJpzyX2I1lcJNnY_hR4b75Vg71Sqgu106j

fd352 := client.Open("/2xFcjyIHJh.txt", client.O_RDWR|client.O_CREATE)
if(fd352 < 0) {
  panic("open failed")
}


ret = client.Seek(fd322, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd339, []byte("gh8ZrcUqP41TT5jUhdQlmP6SfDpCM8GY2TZeTUR4hO8IhS61NFOW1gaYlVv70Hihc9vnly"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) gh8ZrcUqP41TT5jUhdQlmP6SfDpCM8GY2TZeTUR4hO8IhS61NFOW1gaYlVv70Hihc9vnly

buf, ret = client.Read(fd346, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (15) vHT4FWPa17GZJ7LoymjdDosRgf61wgrQf5UYt

ret = client.Write(fd309, []byte("_leuF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) vHT4FWPa17GZJ7L_leuFDosRgf61wgrQf5UYt

ret = client.Seek(fd345, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd351, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd322, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "77vtTfiaeYGkxpJNlAbuW00_sEnUSqPTTtgpPBOw855snZLYu4Rl5mO4uKjeN") {
  panic("wrong data returned")
}


ret = client.Seek(fd350, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) Lfh04iS8FPFBhIm5PUQUDyNQrc_15R6B_W6E5L0yw1t_Xik9GNj5hqR2yiAp9F6dqreoDz6kU3P2o7glYclIbHPk7H0njwgeTx6BTrhi4URdZmDg

ret = client.Write(fd352, []byte("x2onZ9z7QEMYBsx9go6ba1xVywL9PXn2T3hCfHIwPyVx594hkGqbT2mI27J43toIObgB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) x2onZ9z7QEMYBsx9go6ba1xVywL9PXn2T3hCfHIwPyVx594hkGqbT2mI27J43toIObgBDz6kU3P2o7glYclIbHPk7H0njwgeTx6BTrhi4URdZmDg
//fd state: (68) x2onZ9z7QEMYBsx9go6ba1xVywL9PXn2T3hCfHIwPyVx594hkGqbT2mI27J43toIObgBDz6kU3P2o7glYclIbHPk7H0njwgeTx6BTrhi4URdZmDg

ret = client.Write(fd352, []byte("K_P5MPIvYWmaVpbuviksMMDkAyWIhVYnhbJCD94uFfWv3NniteeLxClin_9V9NVR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) x2onZ9z7QEMYBsx9go6ba1xVywL9PXn2T3hCfHIwPyVx594hkGqbT2mI27J43toIObgBK_P5MPIvYWmaVpbuviksMMDkAyWIhVYnhbJCD94uFfWv3NniteeLxClin_9V9NVR
//fd state: (97) ojMm3KphHvyOUztZl5HW0NHtMAiggvU56l_ri5ALj5YiYWCYLP9jVa9aczM4U5fACE8lYOSTwxZRUGzb1gnqsziLzk2cV5g2q

ret = client.Write(fd335, []byte("8Q9mI2t79UWjQ9o7sBqwdlT5SEj8AHaTEimxhi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) ojMm3KphHvyOUztZl5HW0NHtMAiggvU56l_ri5ALj5YiYWCYLP9jVa9aczM4U5fACE8lYOSTwxZRUGzb1gnqsziLzk2cV5g2q8Q9mI2t79UWjQ9o7sBqwdlT5SEj8AHaTEimxhi
//fd state: (132) x2onZ9z7QEMYBsx9go6ba1xVywL9PXn2T3hCfHIwPyVx594hkGqbT2mI27J43toIObgBK_P5MPIvYWmaVpbuviksMMDkAyWIhVYnhbJCD94uFfWv3NniteeLxClin_9V9NVR

ret = client.Write(fd352, []byte("SjQuP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (137) x2onZ9z7QEMYBsx9go6ba1xVywL9PXn2T3hCfHIwPyVx594hkGqbT2mI27J43toIObgBK_P5MPIvYWmaVpbuviksMMDkAyWIhVYnhbJCD94uFfWv3NniteeLxClin_9V9NVRSjQuP

buf, ret = client.Read(fd351, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd334)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd338, 0, client.SEEK_SET)
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

ret = client.Write(fd351, []byte("JU00hh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) JU00hh

ret = client.Seek(fd346, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (32) 2oDyiI2uj5hQ5SRKaY6bMJApelpQxZCu

ret = client.Write(fd328, []byte("eTWPTeuG8kyOOHUMT8ZeUuHpDgzgqetXBvCbMT8wGQGYraFteKtbNt4mJJWD7bzw2dtj9vLGd3Tnxd2mwUO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) 2oDyiI2uj5hQ5SRKaY6bMJApelpQxZCueTWPTeuG8kyOOHUMT8ZeUuHpDgzgqetXBvCbMT8wGQGYraFteKtbNt4mJJWD7bzw2dtj9vLGd3Tnxd2mwUO

buf, ret = client.Read(fd341, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd352)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd341, []byte("1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SKwmBsRs2vebUkwkSWWKmUJg9eqAYqEQd40bLXlFanFLDGI7BzVDJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (90) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SKwmBsRs2vebUkwkSWWKmUJg9eqAYqEQd40bLXlFanFLDGI7BzVDJ

ret = client.Close(fd309)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd328)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd339)
if(ret != 0) {
  panic("close failed")
}


fd353 := client.Open("/HjYfMWqq3p.txt", client.O_RDWR|client.O_CREATE)
if(fd353 < 0) {
  panic("open failed")
}

//fd state: (54) DUtt6FRvZUO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsT9jQmcAd2B_jOnzjovFZCdLZFTYsg8Ak1erqrchBNgQHs41abtjUye9jEvskWpxQ9NnVSOUyKWMms4e

ret = client.Write(fd327, []byte("Tqj698cCTGXVZ2Rl2ZoAC_cxzC5I5kIArFnuzDLUkDN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) DUtt6FRvZUO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsTTqj698cCTGXVZ2Rl2ZoAC_cxzC5I5kIArFnuzDLUkDNs41abtjUye9jEvskWpxQ9NnVSOUyKWMms4e

ret = client.Seek(fd324, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


fd354 := client.Open("/sZAyeiVvZ5.txt", client.O_RDWR|client.O_CREATE)
if(fd354 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd338, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd324, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "sxL4_r6PTkfOitGT_0DU") {
  panic("wrong data returned")
}


ret = client.Close(fd337)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd335, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd347, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


buf, ret = client.Read(fd345, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oDyiI2uj5hQ5SRKaY6bMJApelpQxZCueTWP") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd350, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd327)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd349, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd348)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd353, []byte("WpoGIOgscwlqrMs_MhzQEa1DG9Vn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) WpoGIOgscwlqrMs_MhzQEa1DG9Vn

buf, ret = client.Read(fd338, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd317, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bgBK_P5MPIvYWmaVpbuviksMMDkAyWIhVYnhbJCD94uFfWv3Nnitee") {
  panic("wrong data returned")
}


fd355 := client.Open("/NZOKijET5R.txt", client.O_RDWR|client.O_CREATE)
if(fd355 < 0) {
  panic("open failed")
}


fd356 := client.Open("/ccZLy21HpC.txt", client.O_RDWR|client.O_CREATE)
if(fd356 < 0) {
  panic("open failed")
}


fd357 := client.Open("/N6dWHgfndv.txt", client.O_RDWR|client.O_CREATE)
if(fd357 < 0) {
  panic("open failed")
}


ret = client.Seek(fd353, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}

//fd state: (6) JU00hh

ret = client.Write(fd351, []byte("rKC_tnVdeiqzY1AlaC_9A6GNv3rE44o2Mla8IMZJet9ivw9rRZiExMa5X_Ck4JO43rDF4xbTNFN7Bh8XUREJoDF8fYC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) JU00hhrKC_tnVdeiqzY1AlaC_9A6GNv3rE44o2Mla8IMZJet9ivw9rRZiExMa5X_Ck4JO43rDF4xbTNFN7Bh8XUREJoDF8fYC

ret = client.Seek(fd324, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


fd358 := client.Open("/lzAWuK4aVn.txt", client.O_RDWR|client.O_CREATE)
if(fd358 < 0) {
  panic("open failed")
}

//fd state: (144) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6pGb4t9b9tXGQwnzrgIBVclz1KMzIpNDhokC7UDRNd2A5z6cYzOgQE9h

ret = client.Write(fd329, []byte("u71RxcW_frDMjA0XPhbxq1KYN9Xo3VfJ_kC0kl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (182) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6pGb4t9b9tXGQwnzrgIBVclz1KMzIpNDhokC7UDRNd2A5z6cYzOgQE9hu71RxcW_frDMjA0XPhbxq1KYN9Xo3VfJ_kC0kl

buf, ret = client.Read(fd345, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TeuG8kyOOHUMT8ZeUuHpDgzgqetXBvCbMT8wGQGYraFteKtbNt4mJJWD7bzw2dtj9vLGd3Tnxd2mwUO") {
  panic("wrong data returned")
}


ret = client.Seek(fd341, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd356, []byte("erYkCn44J2JU4H9ZD7OssdAKlyzSyT3Q1UfAesyM9S2xUpw44Nt5Lnpz35c1W_l05E"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) erYkCn44J2JU4H9ZD7OssdAKlyzSyT3Q1UfAesyM9S2xUpw44Nt5Lnpz35c1W_l05E

ret = client.Close(fd338)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd353)
if(ret != 0) {
  panic("close failed")
}

//fd state: (33) vBYLE_MHtW6vf6rP3WbPgQkCwdusxL4_r6PTkfOitGT_0DU

ret = client.Write(fd324, []byte("8oeyNw7XobCr1vIrORDkjH5nH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) vBYLE_MHtW6vf6rP3WbPgQkCwdusxL4_r8oeyNw7XobCr1vIrORDkjH5nH

ret = client.Close(fd358)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd323, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd345, 92, client.SEEK_SET)
if(ret != 92) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 92)
  panic("seek failed")
}


fd359 := client.Open("/ccZLy21HpC.txt", client.O_RDWR|client.O_CREATE)
if(fd359 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd360 := client.Open("/tnVmsNSjJy.txt", client.O_RDWR|client.O_CREATE)
if(fd360 < 0) {
  panic("open failed")
}


fd361 := client.Open("/a_pne7JkfX.txt", client.O_RDWR|client.O_CREATE)
if(fd361 < 0) {
  panic("open failed")
}

//fd state: (119) x2onZ9z7QEMYBsx9go6ba1xVywL9PXn2T3hCfHIwPyVx594hkGqbT2mI27J43toIObgBK_P5MPIvYWmaVpbuviksMMDkAyWIhVYnhbJCD94uFfWv3NniteeLxClin_9V9NVRSjQuP

ret = client.Write(fd317, []byte("76vcV26tfZzZdhoZg7NWuHQPMmVmyXX2Hg6isVOdvdjhf172X5fHInvyUbWoNU5WPtZMiRBTWm71gryVDF3SPPgC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (207) x2onZ9z7QEMYBsx9go6ba1xVywL9PXn2T3hCfHIwPyVx594hkGqbT2mI27J43toIObgBK_P5MPIvYWmaVpbuviksMMDkAyWIhVYnhbJCD94uFfWv3Nnitee76vcV26tfZzZdhoZg7NWuHQPMmVmyXX2Hg6isVOdvdjhf172X5fHInvyUbWoNU5WPtZMiRBTWm71gryVDF3SPPgC

fd362 := client.Open("/phpHm2CqW9.txt", client.O_RDWR|client.O_CREATE)
if(fd362 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd359, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "er") {
  panic("wrong data returned")
}


ret = client.Close(fd346)
if(ret != 0) {
  panic("close failed")
}


fd363 := client.Open("/fGtm1aKOWI.txt", client.O_RDWR|client.O_CREATE)
if(fd363 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd323, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd355, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd360, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "quffD422E9vksKPRR6IynoZ3j3nJc7CTQu") {
  panic("wrong data returned")
}


ret = client.Close(fd342)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd362)
if(ret != 0) {
  panic("close failed")
}


fd364 := client.Open("/WRcTXgVlLq.txt", client.O_RDWR|client.O_CREATE)
if(fd364 < 0) {
  panic("open failed")
}


fd365 := client.Open("/Ecpgj_Bsqe.txt", client.O_RDWR|client.O_CREATE)
if(fd365 < 0) {
  panic("open failed")
}


ret = client.Seek(fd347, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


fd366 := client.Open("/gMb60G3hQV.txt", client.O_RDWR|client.O_CREATE)
if(fd366 < 0) {
  panic("open failed")
}

//fd state: (135) ojMm3KphHvyOUztZl5HW0NHtMAiggvU56l_ri5ALj5YiYWCYLP9jVa9aczM4U5fACE8lYOSTwxZRUGzb1gnqsziLzk2cV5g2q8Q9mI2t79UWjQ9o7sBqwdlT5SEj8AHaTEimxhi

ret = client.Write(fd335, []byte("MZShiDp4sot0iPcp65HeRPxXIDxshG8uJdJ5sC3oDee2o8ggdumc3SEOw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (192) ojMm3KphHvyOUztZl5HW0NHtMAiggvU56l_ri5ALj5YiYWCYLP9jVa9aczM4U5fACE8lYOSTwxZRUGzb1gnqsziLzk2cV5g2q8Q9mI2t79UWjQ9o7sBqwdlT5SEj8AHaTEimxhiMZShiDp4sot0iPcp65HeRPxXIDxshG8uJdJ5sC3oDee2o8ggdumc3SEOw

ret = client.Close(fd322)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd366)
if(ret != 0) {
  panic("close failed")
}


fd367 := client.Open("/xTkLBmiKGg.txt", client.O_RDWR|client.O_CREATE)
if(fd367 < 0) {
  panic("open failed")
}


fd368 := client.Open("/bscOPxCqaM.txt", client.O_RDWR|client.O_CREATE)
if(fd368 < 0) {
  panic("open failed")
}


ret = client.Close(fd323)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd359, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


fd369 := client.Open("/koC3k4dnym.txt", client.O_RDWR|client.O_CREATE)
if(fd369 < 0) {
  panic("open failed")
}


fd370 := client.Open("/PWtXLMuKMR.txt", client.O_RDWR|client.O_CREATE)
if(fd370 < 0) {
  panic("open failed")
}


fd371 := client.Open("/aGmVhCzYpg.txt", client.O_RDWR|client.O_CREATE)
if(fd371 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd335, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd354, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd372 := client.Open("/_gnYfmmFcw.txt", client.O_RDWR|client.O_CREATE)
if(fd372 < 0) {
  panic("open failed")
}


ret = client.Seek(fd317, 149, client.SEEK_SET)
if(ret != 149) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 149)
  panic("seek failed")
}


ret = client.Close(fd361)
if(ret != 0) {
  panic("close failed")
}


fd373 := client.Open("/Ep1TY33BtZ.txt", client.O_RDWR|client.O_CREATE)
if(fd373 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd370, []byte("T89gG2RhDkx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) T89gG2RhDkx

ret = client.Close(fd354)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd363, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd368, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd374 := client.Open("/dOyd_dHSEx.txt", client.O_RDWR|client.O_CREATE)
if(fd374 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd371, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd375 := client.Open("/3HJvawgO69.txt", client.O_RDWR|client.O_CREATE)
if(fd375 < 0) {
  panic("open failed")
}


fd376 := client.Open("/k_7ehxuB5O.txt", client.O_RDWR|client.O_CREATE)
if(fd376 < 0) {
  panic("open failed")
}


fd377 := client.Open("/JJw95P1AZF.txt", client.O_RDWR|client.O_CREATE)
if(fd377 < 0) {
  panic("open failed")
}


fd378 := client.Open("/XRA36JjMkl.txt", client.O_RDWR|client.O_CREATE)
if(fd378 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd363, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd335, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (97) JU00hhrKC_tnVdeiqzY1AlaC_9A6GNv3rE44o2Mla8IMZJet9ivw9rRZiExMa5X_Ck4JO43rDF4xbTNFN7Bh8XUREJoDF8fYC

ret = client.Write(fd351, []byte("dHSmzvdQ2I11QPwLCpLsdPg7iP_wLTedUvKT0oJ2IL9Dyba"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) JU00hhrKC_tnVdeiqzY1AlaC_9A6GNv3rE44o2Mla8IMZJet9ivw9rRZiExMa5X_Ck4JO43rDF4xbTNFN7Bh8XUREJoDF8fYCdHSmzvdQ2I11QPwLCpLsdPg7iP_wLTedUvKT0oJ2IL9Dyba

ret = client.Close(fd363)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd372)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd369, []byte("rKyuK2KVX21x6grr709uSQnw8A8aCvEB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) rKyuK2KVX21x6grr709uSQnw8A8aCvEB

buf, ret = client.Read(fd351, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd379 := client.Open("/CEqlcHC1_z.txt", client.O_RDWR|client.O_CREATE)
if(fd379 < 0) {
  panic("open failed")
}


ret = client.Seek(fd369, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


buf, ret = client.Read(fd377, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OA7Dw") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd374, []byte("sOlIezVCJIhGrXk_WA0bNiQN5rzej8nw2MhL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) sOlIezVCJIhGrXk_WA0bNiQN5rzej8nw2MhL

fd380 := client.Open("/r8EPWv0yyL.txt", client.O_RDWR|client.O_CREATE)
if(fd380 < 0) {
  panic("open failed")
}


ret = client.Close(fd317)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd341, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Jg9eqAYqEQd40bLXlFanFLDGI7BzVDJ") {
  panic("wrong data returned")
}


ret = client.Close(fd356)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd376, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd381 := client.Open("/ex0m66juin.txt", client.O_RDWR|client.O_CREATE)
if(fd381 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd365, []byte("U8kir6pVoYlJZfSVwT7Zh0qzN5n_BNh2yV7TiS8gUmuDRbKDQWG4s6834sGPeF0Ux16KMtzIk8MVxzpTmpU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) U8kir6pVoYlJZfSVwT7Zh0qzN5n_BNh2yV7TiS8gUmuDRbKDQWG4s6834sGPeF0Ux16KMtzIk8MVxzpTmpU

buf, ret = client.Read(fd367, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd365, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd357, []byte("F03x1XC4TTY1tMggRT6o6R2sDAK8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) F03x1XC4TTY1tMggRT6o6R2sDAK8

ret = client.Seek(fd341, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


buf, ret = client.Read(fd380, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1S") {
  panic("wrong data returned")
}


fd382 := client.Open("/CpBTPLAg78.txt", client.O_RDWR|client.O_CREATE)
if(fd382 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd360, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "XxHUFc9s0D") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd382, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd383 := client.Open("/OoNslOuBNK.txt", client.O_RDWR|client.O_CREATE)
if(fd383 < 0) {
  panic("open failed")
}


fd384 := client.Open("/VCgdE3WhbF.txt", client.O_RDWR|client.O_CREATE)
if(fd384 < 0) {
  panic("open failed")
}


ret = client.Close(fd345)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd384, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd385 := client.Open("/6iBimSOOAD.txt", client.O_RDWR|client.O_CREATE)
if(fd385 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd329, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd383, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd375, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobx") {
  panic("wrong data returned")
}


ret = client.Seek(fd385, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd386 := client.Open("/XdRMozFlbD.txt", client.O_RDWR|client.O_CREATE)
if(fd386 < 0) {
  panic("open failed")
}


ret = client.Seek(fd347, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Close(fd377)
if(ret != 0) {
  panic("close failed")
}


fd387 := client.Open("/QpV8PqvRKB.txt", client.O_RDWR|client.O_CREATE)
if(fd387 < 0) {
  panic("open failed")
}


ret = client.Seek(fd376, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (11) T89gG2RhDkx

ret = client.Write(fd370, []byte("9yNYRwwRoMfcrtRNX8Melrfyrq0iVFZKIFGdHR8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) T89gG2RhDkx9yNYRwwRoMfcrtRNX8Melrfyrq0iVFZKIFGdHR8

ret = client.Seek(fd373, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd341, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}

//fd state: (144) JU00hhrKC_tnVdeiqzY1AlaC_9A6GNv3rE44o2Mla8IMZJet9ivw9rRZiExMa5X_Ck4JO43rDF4xbTNFN7Bh8XUREJoDF8fYCdHSmzvdQ2I11QPwLCpLsdPg7iP_wLTedUvKT0oJ2IL9Dyba

ret = client.Write(fd351, []byte("ncdTuDbIikzf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (156) JU00hhrKC_tnVdeiqzY1AlaC_9A6GNv3rE44o2Mla8IMZJet9ivw9rRZiExMa5X_Ck4JO43rDF4xbTNFN7Bh8XUREJoDF8fYCdHSmzvdQ2I11QPwLCpLsdPg7iP_wLTedUvKT0oJ2IL9DybancdTuDbIikzf

ret = client.Close(fd349)
if(ret != 0) {
  panic("close failed")
}

//fd state: (73) dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobxn6friy_WbUPD9mMu2NsGIx

ret = client.Write(fd375, []byte("t3CHFeI1QPc7GzIRJs86RI6Tn3UB9_2bGLQCjQut"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobxt3CHFeI1QPc7GzIRJs86RI6Tn3UB9_2bGLQCjQut
//fd state: (80) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SKwmBsRs2vebUkwkSWWKmUJg9eqAYqEQd40bLXlFanFLDGI7BzVDJ

ret = client.Write(fd341, []byte("3RZ5WvOABuLZXEe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SKwmBsRs2vebUkwkSWWKmUJg9eqAYqEQd40bLXlFanF3RZ5WvOABuLZXEe

ret = client.Seek(fd373, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd369, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "x6grr709uSQnw8A8aCvEB") {
  panic("wrong data returned")
}


fd388 := client.Open("/ccZLy21HpC.txt", client.O_RDWR|client.O_CREATE)
if(fd388 < 0) {
  panic("open failed")
}


fd389 := client.Open("/O5jLTz9ceU.txt", client.O_RDWR|client.O_CREATE)
if(fd389 < 0) {
  panic("open failed")
}


ret = client.Close(fd386)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd371, []byte("NXAYh3NN1TQLQk_71DXInTmlplY0wYdvSKqz7KZMCXawfEOn14Ii5NHiGZaTAtRFnsaBvtS_6LNRL1GU0C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) NXAYh3NN1TQLQk_71DXInTmlplY0wYdvSKqz7KZMCXawfEOn14Ii5NHiGZaTAtRFnsaBvtS_6LNRL1GU0C

buf, ret = client.Read(fd357, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd373, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd329, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd384, []byte("Gz3u5Ivuhtq4Tf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) Gz3u5Ivuhtq4Tf

buf, ret = client.Read(fd374, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd341)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd350, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd383)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd376, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd365, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd371, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd374)
if(ret != 0) {
  panic("close failed")
}


fd390 := client.Open("/HA0fbMGRsh.txt", client.O_RDWR|client.O_CREATE)
if(fd390 < 0) {
  panic("open failed")
}


ret = client.Close(fd385)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd360, 66, client.SEEK_SET)
if(ret != 66) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 66)
  panic("seek failed")
}

//fd state: (14) Gz3u5Ivuhtq4Tf

ret = client.Write(fd384, []byte("PwPVojsU1rHgOBUMuCkOSWuj8C9xbeuNO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) Gz3u5Ivuhtq4TfPwPVojsU1rHgOBUMuCkOSWuj8C9xbeuNO
//fd state: (32) rKyuK2KVX21x6grr709uSQnw8A8aCvEB

ret = client.Write(fd369, []byte("Wl5A57lp3U4JAxn0klXc3YtQPTiSbRY5HsJYbs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) rKyuK2KVX21x6grr709uSQnw8A8aCvEBWl5A57lp3U4JAxn0klXc3YtQPTiSbRY5HsJYbs
//fd state: (38) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SKwmBsRs2vebUkwkSWWKmUJg9eqAYqEQd40bLXlFanF3RZ5WvOABuLZXEe

ret = client.Write(fd380, []byte("BslIa8QaCT7FCQKm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SBslIa8QaCT7FCQKmWWKmUJg9eqAYqEQd40bLXlFanF3RZ5WvOABuLZXEe

ret = client.Close(fd381)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd367, []byte("vCFfXQa06EgdsBDmrxS12aU8odYXDzAWdxNobFLBVvcF_elC46qLTW74M0ZLX82bQX1RbJhld5ui5E77QFUIS2Kym2LMrkdDc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) vCFfXQa06EgdsBDmrxS12aU8odYXDzAWdxNobFLBVvcF_elC46qLTW74M0ZLX82bQX1RbJhld5ui5E77QFUIS2Kym2LMrkdDc

ret = client.Seek(fd390, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd380, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}


ret = client.Seek(fd359, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


fd391 := client.Open("/LoTeuE2wHL.txt", client.O_RDWR|client.O_CREATE)
if(fd391 < 0) {
  panic("open failed")
}


fd392 := client.Open("/zjkNSSD_IZ.txt", client.O_RDWR|client.O_CREATE)
if(fd392 < 0) {
  panic("open failed")
}


ret = client.Seek(fd368, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (113) dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobxt3CHFeI1QPc7GzIRJs86RI6Tn3UB9_2bGLQCjQut

ret = client.Write(fd375, []byte("tj2X8TCntzhCruNWB_2i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobxt3CHFeI1QPc7GzIRJs86RI6Tn3UB9_2bGLQCjQuttj2X8TCntzhCruNWB_2i
//fd state: (0) 

ret = client.Write(fd378, []byte("EqYoao97jeTa8jLQYPg3BnN6u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) EqYoao97jeTa8jLQYPg3BnN6u

ret = client.Seek(fd351, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}


ret = client.Seek(fd347, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


ret = client.Close(fd364)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd384, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


ret = client.Seek(fd368, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd324, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


buf, ret = client.Read(fd390, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ub4zb261bmPM") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd379, []byte("bFMnCMI_uWhIZGagXYuENabII_ca7SzWsXdJEhDGjGfXks4QznG6AkZnna3AR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) bFMnCMI_uWhIZGagXYuENabII_ca7SzWsXdJEhDGjGfXks4QznG6AkZnna3AR

buf, ret = client.Read(fd359, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "35c1W_l05E") {
  panic("wrong data returned")
}


ret = client.Close(fd375)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd355, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd393 := client.Open("/HjcABjXgn5.txt", client.O_RDWR|client.O_CREATE)
if(fd393 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd370, 57)
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


buf, ret = client.Read(fd388, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "erYkCn44J2JU4H9ZD7O") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd365, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd359, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (70) rKyuK2KVX21x6grr709uSQnw8A8aCvEBWl5A57lp3U4JAxn0klXc3YtQPTiSbRY5HsJYbs

ret = client.Write(fd369, []byte("lhKttajr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) rKyuK2KVX21x6grr709uSQnw8A8aCvEBWl5A57lp3U4JAxn0klXc3YtQPTiSbRY5HsJYbslhKttajr
//fd state: (25) EqYoao97jeTa8jLQYPg3BnN6u

ret = client.Write(fd378, []byte("y8R9pFQ4b5a1dUokeB9YALycTfthge2VpmnxlSeA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) EqYoao97jeTa8jLQYPg3BnN6uy8R9pFQ4b5a1dUokeB9YALycTfthge2VpmnxlSeA

fd394 := client.Open("/XdnFXqiD3I.txt", client.O_RDWR|client.O_CREATE)
if(fd394 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd350, []byte("j0V_LTkw2tu1bDOtohk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) j0V_LTkw2tu1bDOtohk

ret = client.Seek(fd351, 145, client.SEEK_SET)
if(ret != 145) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 145)
  panic("seek failed")
}


fd395 := client.Open("/aJEnWn8cux.txt", client.O_RDWR|client.O_CREATE)
if(fd395 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd369, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (66) quffD422E9vksKPRR6IynoZ3j3nJc7CTQuXxHUFc9s0DBVGSALBHyGgwAstibrlUzhbksYaN

ret = client.Write(fd360, []byte("jqNNJZT3TYOWlaSFeR858QFSqlN9a0FfMKiKVCmVTVmaCqVNC1cAiaCk1Hc6sNLIl7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) quffD422E9vksKPRR6IynoZ3j3nJc7CTQuXxHUFc9s0DBVGSALBHyGgwAstibrlUzhjqNNJZT3TYOWlaSFeR858QFSqlN9a0FfMKiKVCmVTVmaCqVNC1cAiaCk1Hc6sNLIl7

ret = client.Close(fd378)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd393, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd335, 98, client.SEEK_SET)
if(ret != 98) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 98)
  panic("seek failed")
}

//fd state: (23) o0WYhFJX8YWxvIFCMCGxl60cAUryO8wg4lMOabL8NeD

ret = client.Write(fd347, []byte("Dfy1QIneeJG88YCNm98WsuuzbOC8ptmY77Nbhr9dpIRbnsh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) o0WYhFJX8YWxvIFCMCGxl60Dfy1QIneeJG88YCNm98WsuuzbOC8ptmY77Nbhr9dpIRbnsh

fd396 := client.Open("/YyURkaOMEI.txt", client.O_RDWR|client.O_CREATE)
if(fd396 < 0) {
  panic("open failed")
}


fd397 := client.Open("/DqNYJpnyj0.txt", client.O_RDWR|client.O_CREATE)
if(fd397 < 0) {
  panic("open failed")
}


fd398 := client.Open("/WNMJT2E_a5.txt", client.O_RDWR|client.O_CREATE)
if(fd398 < 0) {
  panic("open failed")
}


ret = client.Close(fd398)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd355, []byte("28YYVAEc0WcrAmzR8RGqzjBtJDh259KD3KRJimlF4YCODGF2blITNmdoz6mbJieuvE8UjVMR8ssUcRYwJjp9TbIAvIeF_S69MD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) 28YYVAEc0WcrAmzR8RGqzjBtJDh259KD3KRJimlF4YCODGF2blITNmdoz6mbJieuvE8UjVMR8ssUcRYwJjp9TbIAvIeF_S69MD

fd399 := client.Open("/B4ZNEX6Vbw.txt", client.O_RDWR|client.O_CREATE)
if(fd399 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd382, []byte("V2dzZ8G8te7ld6bPiY2193tKSlsrkCb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) V2dzZ8G8te7ld6bPiY2193tKSlsrkCb

ret = client.Close(fd367)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd388)
if(ret != 0) {
  panic("close failed")
}


fd400 := client.Open("/ZXm12GGVei.txt", client.O_RDWR|client.O_CREATE)
if(fd400 < 0) {
  panic("open failed")
}


ret = client.Close(fd396)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd390, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd401 := client.Open("/4NexkYmnGj.txt", client.O_RDWR|client.O_CREATE)
if(fd401 < 0) {
  panic("open failed")
}


fd402 := client.Open("/nYIeV0bwC7.txt", client.O_RDWR|client.O_CREATE)
if(fd402 < 0) {
  panic("open failed")
}


ret = client.Seek(fd387, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd379)
if(ret != 0) {
  panic("close failed")
}


fd403 := client.Open("/rShaVUB4HD.txt", client.O_RDWR|client.O_CREATE)
if(fd403 < 0) {
  panic("open failed")
}


ret = client.Close(fd382)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd368, []byte("1KqKYDG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) 1KqKYDG

fd404 := client.Open("/VcOjpSZndB.txt", client.O_RDWR|client.O_CREATE)
if(fd404 < 0) {
  panic("open failed")
}

//fd state: (28) F03x1XC4TTY1tMggRT6o6R2sDAK8

ret = client.Write(fd357, []byte("Nwkib3k6NQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) F03x1XC4TTY1tMggRT6o6R2sDAK8Nwkib3k6NQ

buf, ret = client.Read(fd371, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd405 := client.Open("/6K38SxNJjF.txt", client.O_RDWR|client.O_CREATE)
if(fd405 < 0) {
  panic("open failed")
}


ret = client.Close(fd365)
if(ret != 0) {
  panic("close failed")
}


fd406 := client.Open("/fGWRLDpC7z.txt", client.O_RDWR|client.O_CREATE)
if(fd406 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd400, []byte("C3v6McwrKLSsD0SoowYMk4kJPaRkYR5lA1zrEtUhkJMoKk7SxBALJVWjzHNPtGRh6gsbYWeCHHHZMNisHkkx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) C3v6McwrKLSsD0SoowYMk4kJPaRkYR5lA1zrEtUhkJMoKk7SxBALJVWjzHNPtGRh6gsbYWeCHHHZMNisHkkx

ret = client.Seek(fd357, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


buf, ret = client.Read(fd406, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd387, []byte("yj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) yj

ret = client.Close(fd395)
if(ret != 0) {
  panic("close failed")
}


fd407 := client.Open("/Uvk9UV6e8V.txt", client.O_RDWR|client.O_CREATE)
if(fd407 < 0) {
  panic("open failed")
}


fd408 := client.Open("/WvrUAZ5kOA.txt", client.O_RDWR|client.O_CREATE)
if(fd408 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd403, []byte("f1TjUZN5f1KmESE_aVG_pByJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) f1TjUZN5f1KmESE_aVG_pByJ
//fd state: (13) Nub4zb261bmPM

ret = client.Write(fd390, []byte("SyNTqo6iUZRU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) Nub4zb261bmPMSyNTqo6iUZRU

ret = client.Close(fd391)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd392, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd392, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd401, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8SWfklF5EBz6iAHz_JL") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd376, 20)
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


ret = client.Close(fd400)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd403, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd350, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Close(fd360)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd407)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd359)
if(ret != 0) {
  panic("close failed")
}


fd409 := client.Open("/a6PEzbIRX2.txt", client.O_RDWR|client.O_CREATE)
if(fd409 < 0) {
  panic("open failed")
}

//fd state: (8) j0V_LTkw2tu1bDOtohk

ret = client.Write(fd350, []byte("JdpZsPkqgpW6B1FugXUFItOvhCT2thnAsVhv0muAOoxVEGWQs5W2I6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) j0V_LTkwJdpZsPkqgpW6B1FugXUFItOvhCT2thnAsVhv0muAOoxVEGWQs5W2I6
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd403, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd389, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "IFo") {
  panic("wrong data returned")
}


fd410 := client.Open("/iool_4YP2J.txt", client.O_RDWR|client.O_CREATE)
if(fd410 < 0) {
  panic("open failed")
}


fd411 := client.Open("/SBGNUumu0z.txt", client.O_RDWR|client.O_CREATE)
if(fd411 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd406, []byte("18DWcnHN73PY1xmBhwMkYuBoRg0TAvqtJ3ObQ4OSMVig0MfVafT1X6lpdI4nq32XrM0ennlYXzz4tJEJCnHzQjyE2XSPFuKH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (96) 18DWcnHN73PY1xmBhwMkYuBoRg0TAvqtJ3ObQ4OSMVig0MfVafT1X6lpdI4nq32XrM0ennlYXzz4tJEJCnHzQjyE2XSPFuKH

fd412 := client.Open("/kwPisEKdpD.txt", client.O_RDWR|client.O_CREATE)
if(fd412 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd399, 25)
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


buf, ret = client.Read(fd390, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd392, []byte("_4drx69HAP8gEJ9J7pdAEjHYHxxIrONED5TCcfGKgtVcoZ2RTLFjLxnAtduf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) _4drx69HAP8gEJ9J7pdAEjHYHxxIrONED5TCcfGKgtVcoZ2RTLFjLxnAtduf

buf, ret = client.Read(fd409, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd380, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WWKmUJg9eqAYqEQd40bLXlFanF3RZ5WvOABuLZXEe") {
  panic("wrong data returned")
}


fd413 := client.Open("/O5jLTz9ceU.txt", client.O_RDWR|client.O_CREATE)
if(fd413 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd392, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd414 := client.Open("/z8PS46PwUh.txt", client.O_RDWR|client.O_CREATE)
if(fd414 < 0) {
  panic("open failed")
}


ret = client.Close(fd393)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd413, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


fd415 := client.Open("/lC3aPrwuh2.txt", client.O_RDWR|client.O_CREATE)
if(fd415 < 0) {
  panic("open failed")
}


ret = client.Close(fd409)
if(ret != 0) {
  panic("close failed")
}


fd416 := client.Open("/IfH1CmBLH1.txt", client.O_RDWR|client.O_CREATE)
if(fd416 < 0) {
  panic("open failed")
}

//fd state: (78) rKyuK2KVX21x6grr709uSQnw8A8aCvEBWl5A57lp3U4JAxn0klXc3YtQPTiSbRY5HsJYbslhKttajr

ret = client.Write(fd369, []byte("j4k6_6ANmi_hRvfsUBDbbnZCZbYOQAyOOgT_yCxluGhLmjJMc8WaS8SXcqjfgNtwSCIuDuc7W4OTHIEnTO_8roKvombp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (170) rKyuK2KVX21x6grr709uSQnw8A8aCvEBWl5A57lp3U4JAxn0klXc3YtQPTiSbRY5HsJYbslhKttajrj4k6_6ANmi_hRvfsUBDbbnZCZbYOQAyOOgT_yCxluGhLmjJMc8WaS8SXcqjfgNtwSCIuDuc7W4OTHIEnTO_8roKvombp

ret = client.Close(fd384)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd404, []byte("0URFdHZtUSPV2b2IsTaDp_KOo9mfGBarPenkxfFlap6TXFP_drUhmMOSflKeHta5i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) 0URFdHZtUSPV2b2IsTaDp_KOo9mfGBarPenkxfFlap6TXFP_drUhmMOSflKeHta5i

fd417 := client.Open("/0TDHoBZ0MP.txt", client.O_RDWR|client.O_CREATE)
if(fd417 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd413, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uyn7oLlSUWtYu") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd408, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd418 := client.Open("/an0N4p7hQP.txt", client.O_RDWR|client.O_CREATE)
if(fd418 < 0) {
  panic("open failed")
}


ret = client.Seek(fd402, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd347, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}


buf, ret = client.Read(fd414, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd419 := client.Open("/jUD4VX4_c1.txt", client.O_RDWR|client.O_CREATE)
if(fd419 < 0) {
  panic("open failed")
}


ret = client.Seek(fd414, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd355)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd414, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd420 := client.Open("/WRcTXgVlLq.txt", client.O_RDWR|client.O_CREATE)
if(fd420 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd408, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd397, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


ret = client.Seek(fd402, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd421 := client.Open("/XsdIbzj7jo.txt", client.O_RDWR|client.O_CREATE)
if(fd421 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd415, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd420, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd416, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcYWO8fYcxG_u293vbLSmk") {
  panic("wrong data returned")
}


ret = client.Close(fd390)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd350, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd416, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC") {
  panic("wrong data returned")
}


ret = client.Seek(fd373, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd417)
if(ret != 0) {
  panic("close failed")
}


fd422 := client.Open("/D2dL3T9Vmw.txt", client.O_RDWR|client.O_CREATE)
if(fd422 < 0) {
  panic("open failed")
}


fd423 := client.Open("/_gnYfmmFcw.txt", client.O_RDWR|client.O_CREATE)
if(fd423 < 0) {
  panic("open failed")
}

//fd state: (19) 8SWfklF5EBz6iAHz_JL

ret = client.Write(fd401, []byte("8fFg1qPdv6MgjkZqhh0O6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) 8SWfklF5EBz6iAHz_JL8fFg1qPdv6MgjkZqhh0O6

ret = client.Close(fd335)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd347, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd399, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd408)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd405)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd406, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd413, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Close(fd351)
if(ret != 0) {
  panic("close failed")
}


fd424 := client.Open("/12b1YLYef1.txt", client.O_RDWR|client.O_CREATE)
if(fd424 < 0) {
  panic("open failed")
}

//fd state: (0) hgYschvflbSHcx8esE45g_c3QPfFedZLoXALsAmClMQLai4IAlr_JTDrruN5kfJMBPEK4EOHW9nlSThzp1xa0kMwrDbrsNNjf3RrONe5Oer0M52jFqMbEeUo2sbgF_2LjjrYetUd_biXYsuUI9sy7146arw79cBtwNAvN1oQlHw

ret = client.Write(fd412, []byte("ti"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) tiYschvflbSHcx8esE45g_c3QPfFedZLoXALsAmClMQLai4IAlr_JTDrruN5kfJMBPEK4EOHW9nlSThzp1xa0kMwrDbrsNNjf3RrONe5Oer0M52jFqMbEeUo2sbgF_2LjjrYetUd_biXYsuUI9sy7146arw79cBtwNAvN1oQlHw

buf, ret = client.Read(fd397, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WOGkY1sN7ANfIz") {
  panic("wrong data returned")
}


ret = client.Close(fd392)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd420, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd425 := client.Open("/7Ihi1jG3Qx.txt", client.O_RDWR|client.O_CREATE)
if(fd425 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd413, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJ") {
  panic("wrong data returned")
}


ret = client.Close(fd347)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd368, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd414, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd376)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd402)
if(ret != 0) {
  panic("close failed")
}

//fd state: (10) F03x1XC4TTY1tMggRT6o6R2sDAK8Nwkib3k6NQ

ret = client.Write(fd357, []byte("8e3yejZCiyyEyCjbeZDQyQCniK9jJmS7ynlcgblJhQX16"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) F03x1XC4TT8e3yejZCiyyEyCjbeZDQyQCniK9jJmS7ynlcgblJhQX16

fd426 := client.Open("/z7bftk_OK2.txt", client.O_RDWR|client.O_CREATE)
if(fd426 < 0) {
  panic("open failed")
}


ret = client.Seek(fd329, 89, client.SEEK_SET)
if(ret != 89) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 89)
  panic("seek failed")
}

//fd state: (40) 8SWfklF5EBz6iAHz_JL8fFg1qPdv6MgjkZqhh0O6

ret = client.Write(fd401, []byte("87Jmk3fcEK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) 8SWfklF5EBz6iAHz_JL8fFg1qPdv6MgjkZqhh0O687Jmk3fcEK
//fd state: (95) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SBslIa8QaCT7FCQKmWWKmUJg9eqAYqEQd40bLXlFanF3RZ5WvOABuLZXEe

ret = client.Write(fd380, []byte("CnaeqMAscLfVt7jkpZRrAD8t7QRDvzqI3YkcBBbBv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SBslIa8QaCT7FCQKmWWKmUJg9eqAYqEQd40bLXlFanF3RZ5WvOABuLZXEeCnaeqMAscLfVt7jkpZRrAD8t7QRDvzqI3YkcBBbBv

fd427 := client.Open("/vcU3VuvRAJ.txt", client.O_RDWR|client.O_CREATE)
if(fd427 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd410, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "un0UI0ExyEVGjC8eqvGu5qAqDxRaPjlJx1uhrXx63uI") {
  panic("wrong data returned")
}

//fd state: (136) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SBslIa8QaCT7FCQKmWWKmUJg9eqAYqEQd40bLXlFanF3RZ5WvOABuLZXEeCnaeqMAscLfVt7jkpZRrAD8t7QRDvzqI3YkcBBbBv

ret = client.Write(fd380, []byte("aHGDmLTTuVQf6w_vCtCSu1WWiYzVrNNR4RitWW9soLhaImN4pMXtkKczHB0t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (196) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SBslIa8QaCT7FCQKmWWKmUJg9eqAYqEQd40bLXlFanF3RZ5WvOABuLZXEeCnaeqMAscLfVt7jkpZRrAD8t7QRDvzqI3YkcBBbBvaHGDmLTTuVQf6w_vCtCSu1WWiYzVrNNR4RitWW9soLhaImN4pMXtkKczHB0t

ret = client.Close(fd397)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd423)
if(ret != 0) {
  panic("close failed")
}


fd428 := client.Open("/6iBimSOOAD.txt", client.O_RDWR|client.O_CREATE)
if(fd428 < 0) {
  panic("open failed")
}


ret = client.Seek(fd421, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd428)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd404, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd429 := client.Open("/MZdRPu9qh_.txt", client.O_RDWR|client.O_CREATE)
if(fd429 < 0) {
  panic("open failed")
}


ret = client.Seek(fd406, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


ret = client.Close(fd429)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd411)
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


fd430 := client.Open("/fGtm1aKOWI.txt", client.O_RDWR|client.O_CREATE)
if(fd430 < 0) {
  panic("open failed")
}


ret = client.Close(fd371)
if(ret != 0) {
  panic("close failed")
}

//fd state: (44) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJKYxDrMBH2z1Buyn7oLlSUWtYu7AO0Gc6iHeOB9p87qzcLpiL1kCMiPIJ

ret = client.Write(fd413, []byte("XchYU0jOu5riE_W_io9ueEeOJmgWXyJ_hxUCW8f9zf9ObyA_q189HWoaUNl0iVn8Zq9xhCD9VNBkUhEIjDmthvyS4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJXchYU0jOu5riE_W_io9ueEeOJmgWXyJ_hxUCW8f9zf9ObyA_q189HWoaUNl0iVn8Zq9xhCD9VNBkUhEIjDmthvyS4
//fd state: (0) 

ret = client.Write(fd394, []byte("ZavIeGKeEdDP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) ZavIeGKeEdDP

fd431 := client.Open("/DqNYJpnyj0.txt", client.O_RDWR|client.O_CREATE)
if(fd431 < 0) {
  panic("open failed")
}


ret = client.Close(fd357)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd415, []byte("g6rGAU7XkOgPUEnxFFSCP1_FHlSqv4UzZS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) g6rGAU7XkOgPUEnxFFSCP1_FHlSqv4UzZS

fd432 := client.Open("/Nhry0WoS3B.txt", client.O_RDWR|client.O_CREATE)
if(fd432 < 0) {
  panic("open failed")
}


ret = client.Close(fd373)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) jSDzXY82g9nyacFFgIB5qk3FP5QnSORsSSun69

ret = client.Write(fd426, []byte("f4YUAg8k"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) f4YUAg8kg9nyacFFgIB5qk3FP5QnSORsSSun69

buf, ret = client.Read(fd431, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1RuKxeSFOppB1E4nqgWOGkY1sN7ANfIzWZ3cWSz") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd419)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd369)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd418, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd406, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "g0MfVafT1X6lpdI4nq32XrM0e") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd418, 78)
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


ret = client.Close(fd420)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd430, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (133) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJXchYU0jOu5riE_W_io9ueEeOJmgWXyJ_hxUCW8f9zf9ObyA_q189HWoaUNl0iVn8Zq9xhCD9VNBkUhEIjDmthvyS4

ret = client.Write(fd413, []byte("X_a2bZBbHEbHFb82Gl89yCUPavhBf65bud9DNgOEqJyAwE_TME0yrAM9H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (190) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJXchYU0jOu5riE_W_io9ueEeOJmgWXyJ_hxUCW8f9zf9ObyA_q189HWoaUNl0iVn8Zq9xhCD9VNBkUhEIjDmthvyS4X_a2bZBbHEbHFb82Gl89yCUPavhBf65bud9DNgOEqJyAwE_TME0yrAM9H
//fd state: (0) 

ret = client.Write(fd430, []byte("60U0cIwZAHe_IJr1kF74cgO1vdsSsjl_2EHgUEyhqoqJVmqmUksGqaorXWW7B1AUd9j"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) 60U0cIwZAHe_IJr1kF74cgO1vdsSsjl_2EHgUEyhqoqJVmqmUksGqaorXWW7B1AUd9j

ret = client.Close(fd422)
if(ret != 0) {
  panic("close failed")
}


fd433 := client.Open("/Ecpgj_Bsqe.txt", client.O_RDWR|client.O_CREATE)
if(fd433 < 0) {
  panic("open failed")
}

//fd state: (0) 1KqKYDG

ret = client.Write(fd368, []byte("uvgxwz4GdXwI7eFtoTrLia7XXDx0E0ZZ_4gejABrg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) uvgxwz4GdXwI7eFtoTrLia7XXDx0E0ZZ_4gejABrg

ret = client.Close(fd425)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd380, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) U8kir6pVoYlJZfSVwT7Zh0qzN5n_BNh2yV7TiS8gUmuDRbKDQWG4s6834sGPeF0Ux16KMtzIk8MVxzpTmpU

ret = client.Write(fd433, []byte("5QiTMVazRoV1vhyVR9McyIh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) 5QiTMVazRoV1vhyVR9McyIhzN5n_BNh2yV7TiS8gUmuDRbKDQWG4s6834sGPeF0Ux16KMtzIk8MVxzpTmpU
//fd state: (41) uvgxwz4GdXwI7eFtoTrLia7XXDx0E0ZZ_4gejABrg

ret = client.Write(fd368, []byte("dcNwD8gLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (106) uvgxwz4GdXwI7eFtoTrLia7XXDx0E0ZZ_4gejABrgdcNwD8gLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG5
//fd state: (68) 18DWcnHN73PY1xmBhwMkYuBoRg0TAvqtJ3ObQ4OSMVig0MfVafT1X6lpdI4nq32XrM0ennlYXzz4tJEJCnHzQjyE2XSPFuKH

ret = client.Write(fd406, []byte("pB05"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) 18DWcnHN73PY1xmBhwMkYuBoRg0TAvqtJ3ObQ4OSMVig0MfVafT1X6lpdI4nq32XrM0epB05Xzz4tJEJCnHzQjyE2XSPFuKH
//fd state: (0) 

ret = client.Write(fd418, []byte("P_TQMFuOjiIByXsrpyaxS5qNQITveHtQW2G3qdnfPneZkkIlUmbSma_8xmlEwEL6l"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) P_TQMFuOjiIByXsrpyaxS5qNQITveHtQW2G3qdnfPneZkkIlUmbSma_8xmlEwEL6l

fd434 := client.Open("/0vSfmI0PuM.txt", client.O_RDWR|client.O_CREATE)
if(fd434 < 0) {
  panic("open failed")
}

//fd state: (106) uvgxwz4GdXwI7eFtoTrLia7XXDx0E0ZZ_4gejABrgdcNwD8gLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG5

ret = client.Write(fd368, []byte("1V23mKFTZbJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (151) uvgxwz4GdXwI7eFtoTrLia7XXDx0E0ZZ_4gejABrgdcNwD8gLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG51V23mKFTZbJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03t

ret = client.Close(fd387)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd414, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd435 := client.Open("/f38bZ7Efc2.txt", client.O_RDWR|client.O_CREATE)
if(fd435 < 0) {
  panic("open failed")
}


ret = client.Seek(fd427, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd406)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd435, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "K") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd433, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "zN5n_BNh2yV7TiS8gUmuDRbKDQWG4s6834sGPeF0Ux16KMtzIk8MVxzpTmpU") {
  panic("wrong data returned")
}

//fd state: (96) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1RuKxeSFOppB1E4nqgWOGkY1sN7ANfIzWZ3cWSzf8PZmWO8wFFzFT5i5lpjlRQ92RhkyIWlR4GE0ht8zNlAODw935dQxCS32hwVSq

ret = client.Write(fd431, []byte("IMMXmYX3I7PMLYyaKAfKQ5YYVZkSDX4XHZ4JXCtAcez0YhQeN7NgBL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) iNruriox_kVjabWuyqdxlG63TEFgvWqSU2Wj2i4T7oPxrWl8iq2wxFhrV1RuKxeSFOppB1E4nqgWOGkY1sN7ANfIzWZ3cWSzIMMXmYX3I7PMLYyaKAfKQ5YYVZkSDX4XHZ4JXCtAcez0YhQeN7NgBLS32hwVSq

buf, ret = client.Read(fd350, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (196) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SBslIa8QaCT7FCQKmWWKmUJg9eqAYqEQd40bLXlFanF3RZ5WvOABuLZXEeCnaeqMAscLfVt7jkpZRrAD8t7QRDvzqI3YkcBBbBvaHGDmLTTuVQf6w_vCtCSu1WWiYzVrNNR4RitWW9soLhaImN4pMXtkKczHB0t

ret = client.Write(fd380, []byte("XqzXQrMPpr8FzdP_rCHkZLm0_nE6mz_hjPJdk5kOoRGMSjMfxTcGDS1ohwpH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (256) 1rfPTChSzbkEbDEGiGnz1RGMYxp3ceFcMJtk1SBslIa8QaCT7FCQKmWWKmUJg9eqAYqEQd40bLXlFanF3RZ5WvOABuLZXEeCnaeqMAscLfVt7jkpZRrAD8t7QRDvzqI3YkcBBbBvaHGDmLTTuVQf6w_vCtCSu1WWiYzVrNNR4RitWW9soLhaImN4pMXtkKczHB0tXqzXQrMPpr8FzdP_rCHkZLm0_nE6mz_hjPJdk5kOoRGMSjMfxTcGDS1ohwpH

fd436 := client.Open("/nVFk0Llsip.txt", client.O_RDWR|client.O_CREATE)
if(fd436 < 0) {
  panic("open failed")
}


ret = client.Seek(fd435, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (43) un0UI0ExyEVGjC8eqvGu5qAqDxRaPjlJx1uhrXx63uI

ret = client.Write(fd410, []byte("Jcod7F1gG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) un0UI0ExyEVGjC8eqvGu5qAqDxRaPjlJx1uhrXx63uIJcod7F1gG

ret = client.Seek(fd434, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


ret = client.Close(fd415)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd432)
if(ret != 0) {
  panic("close failed")
}


fd437 := client.Open("/VznqypOi9t.txt", client.O_RDWR|client.O_CREATE)
if(fd437 < 0) {
  panic("open failed")
}


ret = client.Seek(fd421, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd368, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd436, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd389)
if(ret != 0) {
  panic("close failed")
}

//fd state: (89) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6pGb4t9b9tXGQwnzrgIBVclz1KMzIpNDhokC7UDRNd2A5z6cYzOgQE9hu71RxcW_frDMjA0XPhbxq1KYN9Xo3VfJ_kC0kl

ret = client.Write(fd329, []byte("eV3stMkANkyIMi0zn6dRYYl3FiLVXwxfBM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6eV3stMkANkyIMi0zn6dRYYl3FiLVXwxfBMC7UDRNd2A5z6cYzOgQE9hu71RxcW_frDMjA0XPhbxq1KYN9Xo3VfJ_kC0kl

fd438 := client.Open("/PQ99xNzlFm.txt", client.O_RDWR|client.O_CREATE)
if(fd438 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd424, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd431)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd350, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


fd439 := client.Open("/YuDClGibP1.txt", client.O_RDWR|client.O_CREATE)
if(fd439 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd380, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd424, []byte("k1Tr49rb9cvxxnYnS1eW_mV1Fczxmrw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) k1Tr49rb9cvxxnYnS1eW_mV1Fczxmrw

buf, ret = client.Read(fd424, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd440 := client.Open("/4ztSOL9aUF.txt", client.O_RDWR|client.O_CREATE)
if(fd440 < 0) {
  panic("open failed")
}


fd441 := client.Open("/1uvG1cy_pv.txt", client.O_RDWR|client.O_CREATE)
if(fd441 < 0) {
  panic("open failed")
}


ret = client.Seek(fd426, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}

//fd state: (31) k1Tr49rb9cvxxnYnS1eW_mV1Fczxmrw

ret = client.Write(fd424, []byte("ebOR0FMzfFu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) k1Tr49rb9cvxxnYnS1eW_mV1FczxmrwebOR0FMzfFu

ret = client.Close(fd430)
if(ret != 0) {
  panic("close failed")
}


fd442 := client.Open("/P07sWxGvr6.txt", client.O_RDWR|client.O_CREATE)
if(fd442 < 0) {
  panic("open failed")
}


ret = client.Seek(fd435, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd329, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "C7UDRNd2A5z6cYzOgQE9hu71RxcW_frDMjA0X") {
  panic("wrong data returned")
}


ret = client.Close(fd438)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd416, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_") {
  panic("wrong data returned")
}


fd443 := client.Open("/WFrFE5knfM.txt", client.O_RDWR|client.O_CREATE)
if(fd443 < 0) {
  panic("open failed")
}


ret = client.Close(fd414)
if(ret != 0) {
  panic("close failed")
}


fd444 := client.Open("/hlqnPQas86.txt", client.O_RDWR|client.O_CREATE)
if(fd444 < 0) {
  panic("open failed")
}


fd445 := client.Open("/koC3k4dnym.txt", client.O_RDWR|client.O_CREATE)
if(fd445 < 0) {
  panic("open failed")
}


fd446 := client.Open("/7nzhcLMCES.txt", client.O_RDWR|client.O_CREATE)
if(fd446 < 0) {
  panic("open failed")
}

//fd state: (65) 0URFdHZtUSPV2b2IsTaDp_KOo9mfGBarPenkxfFlap6TXFP_drUhmMOSflKeHta5i

ret = client.Write(fd404, []byte("vGLVOUXKhm9ivMqm1N6BJgP3vh71cjS1f4VfvL5LRTMbaTWrlOrnsF5pQzP9noCFZ6m52zyuXunCFKFBh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (146) 0URFdHZtUSPV2b2IsTaDp_KOo9mfGBarPenkxfFlap6TXFP_drUhmMOSflKeHta5ivGLVOUXKhm9ivMqm1N6BJgP3vh71cjS1f4VfvL5LRTMbaTWrlOrnsF5pQzP9noCFZ6m52zyuXunCFKFBh

ret = client.Close(fd418)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd404)
if(ret != 0) {
  panic("close failed")
}

//fd state: (190) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJXchYU0jOu5riE_W_io9ueEeOJmgWXyJ_hxUCW8f9zf9ObyA_q189HWoaUNl0iVn8Zq9xhCD9VNBkUhEIjDmthvyS4X_a2bZBbHEbHFb82Gl89yCUPavhBf65bud9DNgOEqJyAwE_TME0yrAM9H

ret = client.Write(fd413, []byte("5P_SQ67xmravweA07CZGcLiFLo6c3c"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (220) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJXchYU0jOu5riE_W_io9ueEeOJmgWXyJ_hxUCW8f9zf9ObyA_q189HWoaUNl0iVn8Zq9xhCD9VNBkUhEIjDmthvyS4X_a2bZBbHEbHFb82Gl89yCUPavhBf65bud9DNgOEqJyAwE_TME0yrAM9H5P_SQ67xmravweA07CZGcLiFLo6c3c

buf, ret = client.Read(fd368, 32)
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


ret = client.Close(fd421)
if(ret != 0) {
  panic("close failed")
}

//fd state: (220) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJXchYU0jOu5riE_W_io9ueEeOJmgWXyJ_hxUCW8f9zf9ObyA_q189HWoaUNl0iVn8Zq9xhCD9VNBkUhEIjDmthvyS4X_a2bZBbHEbHFb82Gl89yCUPavhBf65bud9DNgOEqJyAwE_TME0yrAM9H5P_SQ67xmravweA07CZGcLiFLo6c3c

ret = client.Write(fd413, []byte("4nycMI7NJpmDfp34bjNA9JEl4G"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (246) IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJXchYU0jOu5riE_W_io9ueEeOJmgWXyJ_hxUCW8f9zf9ObyA_q189HWoaUNl0iVn8Zq9xhCD9VNBkUhEIjDmthvyS4X_a2bZBbHEbHFb82Gl89yCUPavhBf65bud9DNgOEqJyAwE_TME0yrAM9H5P_SQ67xmravweA07CZGcLiFLo6c3c4nycMI7NJpmDfp34bjNA9JEl4G

ret = client.Close(fd413)
if(ret != 0) {
  panic("close failed")
}

//fd state: (143) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcYWO8fYcxG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_

ret = client.Write(fd416, []byte("DjGBEYoawt98HcK1mh96ePO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (166) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcYWO8fYcxG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePO

buf, ret = client.Read(fd426, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "FgIB5q") {
  panic("wrong data returned")
}


ret = client.Close(fd399)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd444, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd447 := client.Open("/eLMT0M0BqC.txt", client.O_RDWR|client.O_CREATE)
if(fd447 < 0) {
  panic("open failed")
}

//fd state: (0) rKyuK2KVX21x6grr709uSQnw8A8aCvEBWl5A57lp3U4JAxn0klXc3YtQPTiSbRY5HsJYbslhKttajrj4k6_6ANmi_hRvfsUBDbbnZCZbYOQAyOOgT_yCxluGhLmjJMc8WaS8SXcqjfgNtwSCIuDuc7W4OTHIEnTO_8roKvombp

ret = client.Write(fd445, []byte("tdraiIalpd34"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) tdraiIalpd346grr709uSQnw8A8aCvEBWl5A57lp3U4JAxn0klXc3YtQPTiSbRY5HsJYbslhKttajrj4k6_6ANmi_hRvfsUBDbbnZCZbYOQAyOOgT_yCxluGhLmjJMc8WaS8SXcqjfgNtwSCIuDuc7W4OTHIEnTO_8roKvombp
//fd state: (42) k1Tr49rb9cvxxnYnS1eW_mV1FczxmrwebOR0FMzfFu

ret = client.Write(fd424, []byte("7XeRrpemeaiU_1c38l0vtHC0yNaq_icnfmclI28W7UktRqVI1Rit4DK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) k1Tr49rb9cvxxnYnS1eW_mV1FczxmrwebOR0FMzfFu7XeRrpemeaiU_1c38l0vtHC0yNaq_icnfmclI28W7UktRqVI1Rit4DK

fd448 := client.Open("/gMb60G3hQV.txt", client.O_RDWR|client.O_CREATE)
if(fd448 < 0) {
  panic("open failed")
}

//fd state: (38) j0V_LTkwJdpZsPkqgpW6B1FugXUFItOvhCT2thnAsVhv0muAOoxVEGWQs5W2I6

ret = client.Write(fd350, []byte("ZwV0evS4rAddcq54KlYrwpgIvJkgxr1IqLnS34Fqcl1bj4IeTn9_q9clZiJcEXBHnbxwaqBhlzq1Lw5FxyC83Cg5LZZYMUUXuvk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (137) j0V_LTkwJdpZsPkqgpW6B1FugXUFItOvhCT2thZwV0evS4rAddcq54KlYrwpgIvJkgxr1IqLnS34Fqcl1bj4IeTn9_q9clZiJcEXBHnbxwaqBhlzq1Lw5FxyC83Cg5LZZYMUUXuvk

buf, ret = client.Read(fd439, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd447, []byte("k4arljCVCpl8LM98NijSFDr9pIlgzIHWbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WPIs8A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) k4arljCVCpl8LM98NijSFDr9pIlgzIHWbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WPIs8A

ret = client.Close(fd427)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd440, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (21) f4YUAg8kg9nyacFFgIB5qk3FP5QnSORsSSun69

ret = client.Write(fd426, []byte("NE68FaHKpqJesg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) f4YUAg8kg9nyacFFgIB5qNE68FaHKpqJesgn69
//fd state: (166) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcYWO8fYcxG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePO

ret = client.Write(fd416, []byte("gmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (223) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcYWO8fYcxG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7V

buf, ret = client.Read(fd329, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Phbxq1KYN9") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd329, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Xo3VfJ_kC0kl") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd441, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd433, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd439, 96)
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

//fd state: (137) j0V_LTkwJdpZsPkqgpW6B1FugXUFItOvhCT2thZwV0evS4rAddcq54KlYrwpgIvJkgxr1IqLnS34Fqcl1bj4IeTn9_q9clZiJcEXBHnbxwaqBhlzq1Lw5FxyC83Cg5LZZYMUUXuvk

ret = client.Write(fd350, []byte("4tO8jsw2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (145) j0V_LTkwJdpZsPkqgpW6B1FugXUFItOvhCT2thZwV0evS4rAddcq54KlYrwpgIvJkgxr1IqLnS34Fqcl1bj4IeTn9_q9clZiJcEXBHnbxwaqBhlzq1Lw5FxyC83Cg5LZZYMUUXuvk4tO8jsw2

ret = client.Seek(fd444, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (83) k4arljCVCpl8LM98NijSFDr9pIlgzIHWbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WPIs8A

ret = client.Write(fd447, []byte("O2gQuC5LO0EF1jpG0he_W5QG8gElNCJ5O_kxVxhY59XAoIal2VuPAjYA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (139) k4arljCVCpl8LM98NijSFDr9pIlgzIHWbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WPIs8AO2gQuC5LO0EF1jpG0he_W5QG8gElNCJ5O_kxVxhY59XAoIal2VuPAjYA

ret = client.Seek(fd410, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd449 := client.Open("/ijpM9TULvy.txt", client.O_RDWR|client.O_CREATE)
if(fd449 < 0) {
  panic("open failed")
}


ret = client.Close(fd368)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 19pzCz6ktdSCrSveYuFE84LKWmQswrb4XoEud95X1dyXGj

ret = client.Write(fd437, []byte("zQ0UbVvX0tK4owApX3LNHuo2p66OBxUHsUXO9aosiPnGWcZ1n9iFNhF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) zQ0UbVvX0tK4owApX3LNHuo2p66OBxUHsUXO9aosiPnGWcZ1n9iFNhF

ret = client.Seek(fd439, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd394)
if(ret != 0) {
  panic("close failed")
}


fd450 := client.Open("/4XfOwqbLA0.txt", client.O_RDWR|client.O_CREATE)
if(fd450 < 0) {
  panic("open failed")
}


ret = client.Close(fd380)
if(ret != 0) {
  panic("close failed")
}

//fd state: (40) un0UI0ExyEVGjC8eqvGu5qAqDxRaPjlJx1uhrXx63uIJcod7F1gG

ret = client.Write(fd410, []byte("oXDRtusapuOYmkRNCkRx1hnG5xT3Kn_a2EGnev"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) un0UI0ExyEVGjC8eqvGu5qAqDxRaPjlJx1uhrXx6oXDRtusapuOYmkRNCkRx1hnG5xT3Kn_a2EGnev

buf, ret = client.Read(fd439, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (83) 5QiTMVazRoV1vhyVR9McyIhzN5n_BNh2yV7TiS8gUmuDRbKDQWG4s6834sGPeF0Ux16KMtzIk8MVxzpTmpU

ret = client.Write(fd433, []byte("phj5vJfJFcdcwWfz1gyW_74Dab3cQcf9r12WLv0mxt75gOABgZY5LQ2fNb7q9SxlYzezHZZkb4tBEaGE0vPKLV4t5K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (173) 5QiTMVazRoV1vhyVR9McyIhzN5n_BNh2yV7TiS8gUmuDRbKDQWG4s6834sGPeF0Ux16KMtzIk8MVxzpTmpUphj5vJfJFcdcwWfz1gyW_74Dab3cQcf9r12WLv0mxt75gOABgZY5LQ2fNb7q9SxlYzezHZZkb4tBEaGE0vPKLV4t5K

ret = client.Seek(fd450, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd410, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Seek(fd329, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


ret = client.Close(fd424)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd449)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd350)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd446, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


buf, ret = client.Read(fd329, 76)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6eV3stMkANkyIMi0zn6dRYYl3FiLVXwxfB") {
  panic("wrong data returned")
}

//fd state: (223) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcYWO8fYcxG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7V

ret = client.Write(fd416, []byte("BaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (263) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcYWO8fYcxG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv

buf, ret = client.Read(fd439, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd437, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd451 := client.Open("/IfH1CmBLH1.txt", client.O_RDWR|client.O_CREATE)
if(fd451 < 0) {
  panic("open failed")
}

//fd state: (0) xEODE_CL5tzKVusv0EqNljmIHSutuVyakLM0IcYWO8fYcxG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv

ret = client.Write(fd451, []byte("5mGZfIJojFi8gRbmDZgVFHYAwDyp_a142ClO28P1EjobhC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) 5mGZfIJojFi8gRbmDZgVFHYAwDyp_a142ClO28P1EjobhCG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv
//fd state: (139) k4arljCVCpl8LM98NijSFDr9pIlgzIHWbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WPIs8AO2gQuC5LO0EF1jpG0he_W5QG8gElNCJ5O_kxVxhY59XAoIal2VuPAjYA

ret = client.Write(fd447, []byte("RbgubjiyTuJ7yb8qwIUU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (159) k4arljCVCpl8LM98NijSFDr9pIlgzIHWbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WPIs8AO2gQuC5LO0EF1jpG0he_W5QG8gElNCJ5O_kxVxhY59XAoIal2VuPAjYARbgubjiyTuJ7yb8qwIUU

ret = client.Seek(fd437, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}

//fd state: (122) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6eV3stMkANkyIMi0zn6dRYYl3FiLVXwxfBMC7UDRNd2A5z6cYzOgQE9hu71RxcW_frDMjA0XPhbxq1KYN9Xo3VfJ_kC0kl

ret = client.Write(fd329, []byte("UpsSW5WkmfSZntY1XHU7Hmf06EXkXgDVR3utgj8JtJV2NAopjH2D4Qm2iDSklHYldqJ7frYLv1xd33WocC5PkF1A01M"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (213) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6eV3stMkANkyIMi0zn6dRYYl3FiLVXwxfBUpsSW5WkmfSZntY1XHU7Hmf06EXkXgDVR3utgj8JtJV2NAopjH2D4Qm2iDSklHYldqJ7frYLv1xd33WocC5PkF1A01M

ret = client.Seek(fd416, 103, client.SEEK_SET)
if(ret != 103) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 103)
  panic("seek failed")
}


ret = client.Seek(fd437, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


buf, ret = client.Read(fd435, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "K") {
  panic("wrong data returned")
}

//fd state: (12) tdraiIalpd346grr709uSQnw8A8aCvEBWl5A57lp3U4JAxn0klXc3YtQPTiSbRY5HsJYbslhKttajrj4k6_6ANmi_hRvfsUBDbbnZCZbYOQAyOOgT_yCxluGhLmjJMc8WaS8SXcqjfgNtwSCIuDuc7W4OTHIEnTO_8roKvombp

ret = client.Write(fd445, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) tdraiIalpd346grr709uSQnw8A8aCvEBWl5A57lp3U4JAxn0klXc3YtQPTiSbRY5HsJYbslhKttajrj4k6_6ANmi_hRvfsUBDbbnZCZbYOQAyOOgT_yCxluGhLmjJMc8WaS8SXcqjfgNtwSCIuDuc7W4OTHIEnTO_8roKvombp

ret = client.Close(fd450)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd436, []byte("bAPtH9bEtsPwxMV7Buc4AAkfKF3lWlnD24YQ_9SER3Z6L7sgBumqAfxDhFn8EldI2mQuFRD3rL9_EuhSjZ5j0mPHjNoV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) bAPtH9bEtsPwxMV7Buc4AAkfKF3lWlnD24YQ_9SER3Z6L7sgBumqAfxDhFn8EldI2mQuFRD3rL9_EuhSjZ5j0mPHjNoV

ret = client.Close(fd426)
if(ret != 0) {
  panic("close failed")
}

//fd state: (46) 5mGZfIJojFi8gRbmDZgVFHYAwDyp_a142ClO28P1EjobhCG_u293vbLSmklWqvY51SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv

ret = client.Write(fd451, []byte("aqXIL6Vqupzc9HjVX9O"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) 5mGZfIJojFi8gRbmDZgVFHYAwDyp_a142ClO28P1EjobhCaqXIL6Vqupzc9HjVX9OSO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv

buf, ret = client.Read(fd451, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "SO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4K") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd401, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd401)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd441, []byte("eK_NcDFjC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (9) eK_NcDFjC
//fd state: (213) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6eV3stMkANkyIMi0zn6dRYYl3FiLVXwxfBUpsSW5WkmfSZntY1XHU7Hmf06EXkXgDVR3utgj8JtJV2NAopjH2D4Qm2iDSklHYldqJ7frYLv1xd33WocC5PkF1A01M

ret = client.Write(fd329, []byte("jliqmHIjhwOJaaIrZyuJTmZK_hoUpjs662HmYeLJgEnHyOUui7wN6AA5mtL8jtCUomXBxFk_yajV8JiRK0N8_KyUiH3N6WLe76b"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (312) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6eV3stMkANkyIMi0zn6dRYYl3FiLVXwxfBUpsSW5WkmfSZntY1XHU7Hmf06EXkXgDVR3utgj8JtJV2NAopjH2D4Qm2iDSklHYldqJ7frYLv1xd33WocC5PkF1A01MjliqmHIjhwOJaaIrZyuJTmZK_hoUpjs662HmYeLJgEnHyOUui7wN6AA5mtL8jtCUomXBxFk_yajV8JiRK0N8_KyUiH3N6WLe76b

fd452 := client.Open("/nzjL4wl1H6.txt", client.O_RDWR|client.O_CREATE)
if(fd452 < 0) {
  panic("open failed")
}


ret = client.Seek(fd437, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


buf, ret = client.Read(fd416, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1") {
  panic("wrong data returned")
}


ret = client.Seek(fd445, 95, client.SEEK_SET)
if(ret != 95) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 95)
  panic("seek failed")
}


ret = client.Close(fd436)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd439, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd447)
if(ret != 0) {
  panic("close failed")
}

//fd state: (8) zQ0UbVvX0tK4owApX3LNHuo2p66OBxUHsUXO9aosiPnGWcZ1n9iFNhF

ret = client.Write(fd437, []byte("xQaEWwpJXrJ7YBzBZAgyRE4_i81ZL_2rosJe7BfUgvgKEcbElvSdPLf_hqHUI0OeVxUoh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) zQ0UbVvXxQaEWwpJXrJ7YBzBZAgyRE4_i81ZL_2rosJe7BfUgvgKEcbElvSdPLf_hqHUI0OeVxUoh

ret = client.Close(fd446)
if(ret != 0) {
  panic("close failed")
}

//fd state: (107) 5mGZfIJojFi8gRbmDZgVFHYAwDyp_a142ClO28P1EjobhCaqXIL6Vqupzc9HjVX9OSO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KqEoZfbAhT17gFvQkVkbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv

ret = client.Write(fd451, []byte("tlKqwFkRtVlDOv5_lr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (125) 5mGZfIJojFi8gRbmDZgVFHYAwDyp_a142ClO28P1EjobhCaqXIL6Vqupzc9HjVX9OSO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KtlKqwFkRtVlDOv5_lrbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv

ret = client.Seek(fd439, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd451)
if(ret != 0) {
  panic("close failed")
}


fd453 := client.Open("/ck3kOvcScG.txt", client.O_RDWR|client.O_CREATE)
if(fd453 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd444, []byte("fc3_rRNGuSOpsSfxruOVRXJioX5PwzKC3JnAb6lEEvogoGI20YiRC98wbKmW73xtyCOv4Lrrit1c4XFgNHqbUK2A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) fc3_rRNGuSOpsSfxruOVRXJioX5PwzKC3JnAb6lEEvogoGI20YiRC98wbKmW73xtyCOv4Lrrit1c4XFgNHqbUK2A

buf, ret = client.Read(fd444, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd454 := client.Open("/ck3kOvcScG.txt", client.O_RDWR|client.O_CREATE)
if(fd454 < 0) {
  panic("open failed")
}


fd455 := client.Open("/Ch3TfdeKGU.txt", client.O_RDWR|client.O_CREATE)
if(fd455 < 0) {
  panic("open failed")
}

//fd state: (0) opcQZzcDBsWOgvqdwnNYf7HsoCmnG3iqHUxOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_CtngkddGPzHaXXQiR3upF6zZdS0j4voSxTpbrQDvpqKAn0HdNXEvhEPsxK9VEJgjghOi1j3OjF2CEROv2xkb9Fu9yFcX2kYl59HdJrBk_5O1it3WWkb7XcvHP0fMEwWuokCiB_InMLWragtSRwrCeJ63_GBMrqC65v9gsHBQAMoBmigTbyQ1Oz1mJWOIVA_t

ret = client.Write(fd455, []byte("y_97vCaHmX17GTsbbREFgX2NwKRZZIV_AJG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) y_97vCaHmX17GTsbbREFgX2NwKRZZIV_AJGOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_CtngkddGPzHaXXQiR3upF6zZdS0j4voSxTpbrQDvpqKAn0HdNXEvhEPsxK9VEJgjghOi1j3OjF2CEROv2xkb9Fu9yFcX2kYl59HdJrBk_5O1it3WWkb7XcvHP0fMEwWuokCiB_InMLWragtSRwrCeJ63_GBMrqC65v9gsHBQAMoBmigTbyQ1Oz1mJWOIVA_t

ret = client.Close(fd410)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd445)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd452)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd453, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Close(fd440)
if(ret != 0) {
  panic("close failed")
}


fd456 := client.Open("/WmilsB_bL8.txt", client.O_RDWR|client.O_CREATE)
if(fd456 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd437, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (312) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6eV3stMkANkyIMi0zn6dRYYl3FiLVXwxfBUpsSW5WkmfSZntY1XHU7Hmf06EXkXgDVR3utgj8JtJV2NAopjH2D4Qm2iDSklHYldqJ7frYLv1xd33WocC5PkF1A01MjliqmHIjhwOJaaIrZyuJTmZK_hoUpjs662HmYeLJgEnHyOUui7wN6AA5mtL8jtCUomXBxFk_yajV8JiRK0N8_KyUiH3N6WLe76b

ret = client.Write(fd329, []byte("SbGGIATO__"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (322) YBa0qQJHJtY6LuzaEIojRZMpqYypk7QOKHHND_jlMM5zo0vpHHmw4NiDutyfl0dN4OD3pkKGfvNVk70cb0B0tYfT6eV3stMkANkyIMi0zn6dRYYl3FiLVXwxfBUpsSW5WkmfSZntY1XHU7Hmf06EXkXgDVR3utgj8JtJV2NAopjH2D4Qm2iDSklHYldqJ7frYLv1xd33WocC5PkF1A01MjliqmHIjhwOJaaIrZyuJTmZK_hoUpjs662HmYeLJgEnHyOUui7wN6AA5mtL8jtCUomXBxFk_yajV8JiRK0N8_KyUiH3N6WLe76bSbGGIATO__
//fd state: (0) 

ret = client.Write(fd439, []byte("LNnmdEqMzf0x9aD4MheT2j2hWH5iJsrqM7OgqxmSuv8maBs4Nxr05a83A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) LNnmdEqMzf0x9aD4MheT2j2hWH5iJsrqM7OgqxmSuv8maBs4Nxr05a83A

ret = client.Close(fd433)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd437)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd454)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd435, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd448, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd434, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "6JUxykZcfsC_9DOGqPpPoYQH") {
  panic("wrong data returned")
}


ret = client.Seek(fd453, 60, client.SEEK_SET)
if(ret != 60) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 60)
  panic("seek failed")
}


ret = client.Close(fd435)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd453, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "tq8A7OHYsjUhjG") {
  panic("wrong data returned")
}


fd457 := client.Open("/PSn4daw9hM.txt", client.O_RDWR|client.O_CREATE)
if(fd457 < 0) {
  panic("open failed")
}


ret = client.Close(fd453)
if(ret != 0) {
  panic("close failed")
}


fd458 := client.Open("/B3ip_bKpPr.txt", client.O_RDWR|client.O_CREATE)
if(fd458 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd448, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd455, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9") {
  panic("wrong data returned")
}

//fd state: (57) LNnmdEqMzf0x9aD4MheT2j2hWH5iJsrqM7OgqxmSuv8maBs4Nxr05a83A

ret = client.Write(fd439, []byte("vqpyStgln6VN1uRe6Cd3uMEv1IcE9mH9fi1TlZoNDgXK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) LNnmdEqMzf0x9aD4MheT2j2hWH5iJsrqM7OgqxmSuv8maBs4Nxr05a83AvqpyStgln6VN1uRe6Cd3uMEv1IcE9mH9fi1TlZoNDgXK

ret = client.Seek(fd448, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (134) y_97vCaHmX17GTsbbREFgX2NwKRZZIV_AJGOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9V_CtngkddGPzHaXXQiR3upF6zZdS0j4voSxTpbrQDvpqKAn0HdNXEvhEPsxK9VEJgjghOi1j3OjF2CEROv2xkb9Fu9yFcX2kYl59HdJrBk_5O1it3WWkb7XcvHP0fMEwWuokCiB_InMLWragtSRwrCeJ63_GBMrqC65v9gsHBQAMoBmigTbyQ1Oz1mJWOIVA_t

ret = client.Write(fd455, []byte("aGtR1EU9sS1zgBxq5bpYOtA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (157) y_97vCaHmX17GTsbbREFgX2NwKRZZIV_AJGOBs3i0R_ORzMxuYWI3_OG8RCYl7_0XOy42LyatylP3LL_U_DNTktjYVN86o1R1RFRKg9sqOFOmPtZdxtQCTmEsS1M8Yx5Q7Gvy9aGtR1EU9sS1zgBxq5bpYOtA6zZdS0j4voSxTpbrQDvpqKAn0HdNXEvhEPsxK9VEJgjghOi1j3OjF2CEROv2xkb9Fu9yFcX2kYl59HdJrBk_5O1it3WWkb7XcvHP0fMEwWuokCiB_InMLWragtSRwrCeJ63_GBMrqC65v9gsHBQAMoBmigTbyQ1Oz1mJWOIVA_t
//fd state: (32) 56UEwawz6JUxykZcfsC_9DOGqPpPoYQHeSqvBarJInNKte2GmgxmhKg5jczgKyB3t8GcZYFvIz4theomeF1wiSa1i_BiudfGOqMa2MKRt1nSRZhhE3dJT

ret = client.Write(fd434, []byte("s4vHVXBxNsfWAFPMaCmgIXYaiq91sSlo703C9uaelrY5g_OgH3nRGIqRl4xQJh5yy95Ty1zsu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) 56UEwawz6JUxykZcfsC_9DOGqPpPoYQHs4vHVXBxNsfWAFPMaCmgIXYaiq91sSlo703C9uaelrY5g_OgH3nRGIqRl4xQJh5yy95Ty1zsu1nSRZhhE3dJT

buf, ret = client.Read(fd434, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1nSRZhhE3dJT") {
  panic("wrong data returned")
}


ret = client.Close(fd458)
if(ret != 0) {
  panic("close failed")
}


fd459 := client.Open("/Yt5ysR3lnX.txt", client.O_RDWR|client.O_CREATE)
if(fd459 < 0) {
  panic("open failed")
}


ret = client.Close(fd448)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd439, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd459, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd441)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd444, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd459, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd455)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd416, 132, client.SEEK_SET)
if(ret != 132) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 132)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd456, []byte("tKnfEofJTkEXNdBtzfEUIClh3idH0PAsOGaE4ZpTk2nPdmncLHS_KXCFBXLty1v"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) tKnfEofJTkEXNdBtzfEUIClh3idH0PAsOGaE4ZpTk2nPdmncLHS_KXCFBXLty1v

fd460 := client.Open("/1QMCJHImk3.txt", client.O_RDWR|client.O_CREATE)
if(fd460 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd460, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd444, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd329)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd456, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd434, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


buf, ret = client.Read(fd460, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd460, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd439, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd461 := client.Open("/0v9e3KlwIJ.txt", client.O_RDWR|client.O_CREATE)
if(fd461 < 0) {
  panic("open failed")
}


ret = client.Close(fd460)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd416, 148, client.SEEK_SET)
if(ret != 148) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 148)
  panic("seek failed")
}


ret = client.Close(fd456)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd434, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gH3nRGIqRl4xQJh5yy95Ty1zsu1nSRZhhE3dJT") {
  panic("wrong data returned")
}


fd462 := client.Open("/k57M1a0Y2S.txt", client.O_RDWR|client.O_CREATE)
if(fd462 < 0) {
  panic("open failed")
}


ret = client.Close(fd434)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd439)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd416, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Yoawt98HcK1mh96ePOgmRz") {
  panic("wrong data returned")
}

//fd state: (170) 5mGZfIJojFi8gRbmDZgVFHYAwDyp_a142ClO28P1EjobhCaqXIL6Vqupzc9HjVX9OSO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KtlKqwFkRtVlDOv5_lrbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRzGXqpKrZ4ktgI83LUsQp6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv

ret = client.Write(fd416, []byte("3y64XgvS5xa662Z8D1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (188) 5mGZfIJojFi8gRbmDZgVFHYAwDyp_a142ClO28P1EjobhCaqXIL6Vqupzc9HjVX9OSO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KtlKqwFkRtVlDOv5_lrbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRz3y64XgvS5xa662Z8D1p6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv

fd463 := client.Open("/WEwnP8V8uo.txt", client.O_RDWR|client.O_CREATE)
if(fd463 < 0) {
  panic("open failed")
}


ret = client.Close(fd457)
if(ret != 0) {
  panic("close failed")
}


fd464 := client.Open("/J5KDyCcaR4.txt", client.O_RDWR|client.O_CREATE)
if(fd464 < 0) {
  panic("open failed")
}

//fd state: (188) 5mGZfIJojFi8gRbmDZgVFHYAwDyp_a142ClO28P1EjobhCaqXIL6Vqupzc9HjVX9OSO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KtlKqwFkRtVlDOv5_lrbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRz3y64XgvS5xa662Z8D1p6sKp608yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv

ret = client.Write(fd416, []byte("3gJ0WR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (194) 5mGZfIJojFi8gRbmDZgVFHYAwDyp_a142ClO28P1EjobhCaqXIL6Vqupzc9HjVX9OSO3cWE1Q7LRXLzPHf0R4Q6dS8G09HzFC5adyGG8h4KtlKqwFkRtVlDOv5_lrbbYmzCy4RqkeHB3Zm_DjGBEYoawt98HcK1mh96ePOgmRz3y64XgvS5xa662Z8D13gJ0WR08yoZFl1VrqVSUIXQqqz7zNPR_U7VBaNbGNlAbvHgasg1U2FnZ21DnWwsRQ7IPO64b_uv

fd465 := client.Open("/2uQvu09n7Y.txt", client.O_RDWR|client.O_CREATE)
if(fd465 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd464, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd459, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd465, []byte("AQdoNaA61Q2_eyNG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (16) AQdoNaA61Q2_eyNG

fd466 := client.Open("/hlqnPQas86.txt", client.O_RDWR|client.O_CREATE)
if(fd466 < 0) {
  panic("open failed")
}


fd467 := client.Open("/VdQmE9F0jn.txt", client.O_RDWR|client.O_CREATE)
if(fd467 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd466, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "fc3_rRNGuSOpsSfxruOVRXJio") {
  panic("wrong data returned")
}


ret = client.Seek(fd444, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


buf, ret = client.Read(fd466, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "X5PwzKC3JnAb6lEEvogoGI20YiRC98wbKmW73xtyCOv4Lrrit1c4XFgNHq") {
  panic("wrong data returned")
}


ret = client.Close(fd459)
if(ret != 0) {
  panic("close failed")
}


fd468 := client.Open("/JcnQY5oHjy.txt", client.O_RDWR|client.O_CREATE)
if(fd468 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd463, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd469 := client.Open("/_gnYfmmFcw.txt", client.O_RDWR|client.O_CREATE)
if(fd469 < 0) {
  panic("open failed")
}


ret = client.Seek(fd416, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


ret = client.Close(fd462)
if(ret != 0) {
  panic("close failed")
}


fd470 := client.Open("/X41JyBi0m7.txt", client.O_RDWR|client.O_CREATE)
if(fd470 < 0) {
  panic("open failed")
}


ret = client.Close(fd466)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd465, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd470)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd467, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd468, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd465, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


fd471 := client.Open("/Y0NpaqExV_.txt", client.O_RDWR|client.O_CREATE)
if(fd471 < 0) {
  panic("open failed")
}

//fd state: (0) 7Tza4IwzlxZXvbbw9MYrJud_rHnI7p_tBoUL6L_lxa2PHYRjDOfwvp59rplq8AMfXl

ret = client.Write(fd471, []byte("1vAvNO4KXSAK5mSh8KEuyy_flyJJbPOzb0okw9ZeOQ_WfcbZRwunbC4h71FGxc0jSK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) 1vAvNO4KXSAK5mSh8KEuyy_flyJJbPOzb0okw9ZeOQ_WfcbZRwunbC4h71FGxc0jSK
//fd state: (0) 

ret = client.Write(fd469, []byte("VAARa7wAvboH23bRkmnH6SPCUEIyFLYuBmeNAUCyPo6ImCN6uFYjdK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) VAARa7wAvboH23bRkmnH6SPCUEIyFLYuBmeNAUCyPo6ImCN6uFYjdK

fd472 := client.Open("/ZrRGFqOamQ.txt", client.O_RDWR|client.O_CREATE)
if(fd472 < 0) {
  panic("open failed")
}


ret = client.Close(fd467)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd472, []byte("dbV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) dbV

ret = client.Seek(fd416, 111, client.SEEK_SET)
if(ret != 111) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 111)
  panic("seek failed")
}

//fd state: (66) 1vAvNO4KXSAK5mSh8KEuyy_flyJJbPOzb0okw9ZeOQ_WfcbZRwunbC4h71FGxc0jSK

ret = client.Write(fd471, []byte("o76DUISHAiibVEo1prvXIFtYTZAP8ZRgbQeKaov38uSl05Fj4og0yYVC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) 1vAvNO4KXSAK5mSh8KEuyy_flyJJbPOzb0okw9ZeOQ_WfcbZRwunbC4h71FGxc0jSKo76DUISHAiibVEo1prvXIFtYTZAP8ZRgbQeKaov38uSl05Fj4og0yYVC

buf, ret = client.Read(fd461, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd472, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Close(fd461)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd465)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd444, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lEEvogoGI20YiRC98wbKmW73xtyCOv4Lr") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd416, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wFkRtVlDO") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd463, []byte("l61mpyACae2dygDJY5bhph0EERWYE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) l61mpyACae2dygDJY5bhph0EERWYE

fd473 := client.Open("/ex0m66juin.txt", client.O_RDWR|client.O_CREATE)
if(fd473 < 0) {
  panic("open failed")
}


ret = client.Close(fd444)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd416)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd471, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


buf, ret = client.Read(fd469, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd473, []byte("beJ3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) beJ3
//fd state: (0) 

ret = client.Write(fd468, []byte("XY4nIbLsbqY71FPtvyRqGo8qzQrYgZNcwK6Ca71g9jpo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) XY4nIbLsbqY71FPtvyRqGo8qzQrYgZNcwK6Ca71g9jpo

buf, ret = client.Read(fd472, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd474 := client.Open("/GDNyWBC36b.txt", client.O_RDWR|client.O_CREATE)
if(fd474 < 0) {
  panic("open failed")
}


ret = client.Seek(fd468, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd464, []byte("b83r1KqnRaXwe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) b83r1KqnRaXwe

ret = client.Seek(fd463, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}

//fd state: (41) XY4nIbLsbqY71FPtvyRqGo8qzQrYgZNcwK6Ca71g9jpo

ret = client.Write(fd468, []byte("hrlTm0Jc10MHT2DpFiJRQTRmgSGr2l_SXIzMPZ3pPM4gafJmH8lwOjqFICS96D9u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) XY4nIbLsbqY71FPtvyRqGo8qzQrYgZNcwK6Ca71g9hrlTm0Jc10MHT2DpFiJRQTRmgSGr2l_SXIzMPZ3pPM4gafJmH8lwOjqFICS96D9u

ret = client.Seek(fd464, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd468)
if(ret != 0) {
  panic("close failed")
}

//fd state: (54) VAARa7wAvboH23bRkmnH6SPCUEIyFLYuBmeNAUCyPo6ImCN6uFYjdK

ret = client.Write(fd469, []byte("Ku4pgmdjib0Xf07_3d"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) VAARa7wAvboH23bRkmnH6SPCUEIyFLYuBmeNAUCyPo6ImCN6uFYjdKKu4pgmdjib0Xf07_3d

fd475 := client.Open("/V2hqe83j8X.txt", client.O_RDWR|client.O_CREATE)
if(fd475 < 0) {
  panic("open failed")
}

//fd state: (9) l61mpyACae2dygDJY5bhph0EERWYE

ret = client.Write(fd463, []byte("IKvsijlcMQsw0pjPO_kDY_Wp_fD23R_aCeV8H8y4zouEri7411jBLyo4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) l61mpyACaIKvsijlcMQsw0pjPO_kDY_Wp_fD23R_aCeV8H8y4zouEri7411jBLyo4

buf, ret = client.Read(fd463, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd473, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd463, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd469, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (1) beJ3

ret = client.Write(fd473, []byte("oZFjS5IGBM463t_JvEWQiI0HV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (26) boZFjS5IGBM463t_JvEWQiI0HV

fd476 := client.Open("/HA0fbMGRsh.txt", client.O_RDWR|client.O_CREATE)
if(fd476 < 0) {
  panic("open failed")
}


ret = client.Seek(fd475, 172, client.SEEK_SET)
if(ret != 172) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 172)
  panic("seek failed")
}


buf, ret = client.Read(fd474, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (26) boZFjS5IGBM463t_JvEWQiI0HV

ret = client.Write(fd473, []byte("Vu7wxWz6btapnWkPVObOabbsPa7R14BAktYLE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) boZFjS5IGBM463t_JvEWQiI0HVVu7wxWz6btapnWkPVObOabbsPa7R14BAktYLE

ret = client.Close(fd474)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd476, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Nub4zb261bmPMSyNTqo6iUZRU") {
  panic("wrong data returned")
}


ret = client.Close(fd463)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd464)
if(ret != 0) {
  panic("close failed")
}

//fd state: (40) 1vAvNO4KXSAK5mSh8KEuyy_flyJJbPOzb0okw9ZeOQ_WfcbZRwunbC4h71FGxc0jSKo76DUISHAiibVEo1prvXIFtYTZAP8ZRgbQeKaov38uSl05Fj4og0yYVC

ret = client.Write(fd471, []byte("QYIcw9NmyZ7NGuzPPKDNDeB04A7Cjvt9vHh8TW9JWMaqFb0tVUTV9DiduNRfH1DqX2rCtKp5gB7UXugi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) 1vAvNO4KXSAK5mSh8KEuyy_flyJJbPOzb0okw9ZeQYIcw9NmyZ7NGuzPPKDNDeB04A7Cjvt9vHh8TW9JWMaqFb0tVUTV9DiduNRfH1DqX2rCtKp5gB7UXugiVC
//fd state: (172) ojMm3KphHvyOUztZl5HW0NHtMAiggvU56l_ri5ALj5YiYWCYLP9jVa9aczM4U5fACE8lYOSTwxZRUGzb1gnqsziLzk2cV5g2q8Q9mI2t79UWjQ9o7sBqwdlT5SEj8AHaTEimxhiMZShiDp4sot0iPcp65HeRPxXIDxshG8uJdJ5sC3oDee2o8ggdumc3SEOw

ret = client.Write(fd475, []byte("wbO4nR3tcFtvJU7O3themprK5rBQl7nUi0g3xyszukDeP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (217) ojMm3KphHvyOUztZl5HW0NHtMAiggvU56l_ri5ALj5YiYWCYLP9jVa9aczM4U5fACE8lYOSTwxZRUGzb1gnqsziLzk2cV5g2q8Q9mI2t79UWjQ9o7sBqwdlT5SEj8AHaTEimxhiMZShiDp4sot0iPcp65HeRPxXIDxshG8uJdJ5swbO4nR3tcFtvJU7O3themprK5rBQl7nUi0g3xyszukDeP

fd477 := client.Open("/Wn_LlxApk0.txt", client.O_RDWR|client.O_CREATE)
if(fd477 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd476, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd469, 68, client.SEEK_SET)
if(ret != 68) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 68)
  panic("seek failed")
}


fd478 := client.Open("/dpmUqilIPW.txt", client.O_RDWR|client.O_CREATE)
if(fd478 < 0) {
  panic("open failed")
}


ret = client.Close(fd473)
if(ret != 0) {
  panic("close failed")
}

//fd state: (3) dbV

ret = client.Write(fd472, []byte("h3sy3MqbuE2Pvab7InJfn_3EO8QVGpbl93GIqXMGX46aVR_6WeHDaIusUELKP1HsyTxmec"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) dbVh3sy3MqbuE2Pvab7InJfn_3EO8QVGpbl93GIqXMGX46aVR_6WeHDaIusUELKP1HsyTxmec

ret = client.Seek(fd478, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd477, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd475)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd471)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd472, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd479 := client.Open("/qksOWSAKgO.txt", client.O_RDWR|client.O_CREATE)
if(fd479 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd476, 79)
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

//fd state: (0) 

ret = client.Write(fd477, []byte("YOzbUp2AgNu8iOoCyqgl4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) YOzbUp2AgNu8iOoCyqgl4

ret = client.Seek(fd472, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}

//fd state: (18) dbVh3sy3MqbuE2Pvab7InJfn_3EO8QVGpbl93GIqXMGX46aVR_6WeHDaIusUELKP1HsyTxmec

ret = client.Write(fd472, []byte("yTCAGahp3emE78GCgOu0fPd6Uh3VVUigC0yAymgYgEJzhHtit7ce"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) dbVh3sy3MqbuE2PvabyTCAGahp3emE78GCgOu0fPd6Uh3VVUigC0yAymgYgEJzhHtit7cemec
//fd state: (21) YOzbUp2AgNu8iOoCyqgl4

ret = client.Write(fd477, []byte("j13e9I9hcdw6zQPlQYPRBKkjy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjy

buf, ret = client.Read(fd479, 42)
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


ret = client.Close(fd479)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd472, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "mec") {
  panic("wrong data returned")
}

//fd state: (73) dbVh3sy3MqbuE2PvabyTCAGahp3emE78GCgOu0fPd6Uh3VVUigC0yAymgYgEJzhHtit7cemec

ret = client.Write(fd472, []byte("WO2gIF8baD3X56zEyzl9UTjiL5X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) dbVh3sy3MqbuE2PvabyTCAGahp3emE78GCgOu0fPd6Uh3VVUigC0yAymgYgEJzhHtit7cemecWO2gIF8baD3X56zEyzl9UTjiL5X

fd480 := client.Open("/lLf6ieLALU.txt", client.O_RDWR|client.O_CREATE)
if(fd480 < 0) {
  panic("open failed")
}


ret = client.Close(fd472)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd477, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd477, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd480, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd481 := client.Open("/9CD243b5nk.txt", client.O_RDWR|client.O_CREATE)
if(fd481 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd478, []byte("KY3sPdro9Zp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) KY3sPdro9Zp

buf, ret = client.Read(fd477, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd480, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd481, 89)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd478)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd480, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd477, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


buf, ret = client.Read(fd477, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "13e9I9hcdw6") {
  panic("wrong data returned")
}


ret = client.Close(fd480)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd481, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd481)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd477, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "zQPlQYPRBKkjy") {
  panic("wrong data returned")
}

//fd state: (46) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjy

ret = client.Write(fd477, []byte("xL1aaw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw

ret = client.Seek(fd477, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}


ret = client.Seek(fd477, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}


fd482 := client.Open("/xPfTAdq4ti.txt", client.O_RDWR|client.O_CREATE)
if(fd482 < 0) {
  panic("open failed")
}


fd483 := client.Open("/EHYKcgo8eo.txt", client.O_RDWR|client.O_CREATE)
if(fd483 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd483, []byte("ng9qilxe5SOdCWytuT9g6ni6HHPgWxEjkzl4Aj2EhGRq_euNurUKi7Bj_4TkcJNmEJobY1MpfwnfbsOr_5EEx2DB3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) ng9qilxe5SOdCWytuT9g6ni6HHPgWxEjkzl4Aj2EhGRq_euNurUKi7Bj_4TkcJNmEJobY1MpfwnfbsOr_5EEx2DB3

buf, ret = client.Read(fd477, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "aaw") {
  panic("wrong data returned")
}


fd484 := client.Open("/2g8nvSsXFZ.txt", client.O_RDWR|client.O_CREATE)
if(fd484 < 0) {
  panic("open failed")
}


fd485 := client.Open("/Fx14wmZV6Z.txt", client.O_RDWR|client.O_CREATE)
if(fd485 < 0) {
  panic("open failed")
}

//fd state: (52) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw

ret = client.Write(fd477, []byte("1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn9_NIRqSsCo1kZxl_rpWv2Sh8loZFUtlQpyukEue9YtoMytkLq2nzphx3Iih"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn9_NIRqSsCo1kZxl_rpWv2Sh8loZFUtlQpyukEue9YtoMytkLq2nzphx3Iih

ret = client.Close(fd482)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) Arf1bxOFJqakHXAPAcfXmJ0yMQ6dcnEhibz4P6E0xy9PMFa0_oTHnxT4aymvLZ

ret = client.Write(fd484, []byte("w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOJTrqdreNuZLXNgCQ_akukCxfdKIMWKyVRTnvaS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOJTrqdreNuZLXNgCQ_akukCxfdKIMWKyVRTnvaS

ret = client.Close(fd484)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd485, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd485, []byte("43ORMZCWPz2XjGkTjmrHN59C3nH6L95mB8M4ngpYRu_rrG6jrbYmsl1Oh2RFWFouN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) 43ORMZCWPz2XjGkTjmrHN59C3nH6L95mB8M4ngpYRu_rrG6jrbYmsl1Oh2RFWFouN

ret = client.Seek(fd477, 120, client.SEEK_SET)
if(ret != 120) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 120)
  panic("seek failed")
}

//fd state: (120) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn9_NIRqSsCo1kZxl_rpWv2Sh8loZFUtlQpyukEue9YtoMytkLq2nzphx3Iih

ret = client.Write(fd477, []byte("mEh5_fscTpoo4MOAkAHJjuWo4Pb4u25Mx0KS1rZXRT6k5K3mqh6zHh6yiBJi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (180) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn9_NIRqSsCo1kZxl_rpWv2Sh8loZFUmEh5_fscTpoo4MOAkAHJjuWo4Pb4u25Mx0KS1rZXRT6k5K3mqh6zHh6yiBJi
//fd state: (180) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn9_NIRqSsCo1kZxl_rpWv2Sh8loZFUmEh5_fscTpoo4MOAkAHJjuWo4Pb4u25Mx0KS1rZXRT6k5K3mqh6zHh6yiBJi

ret = client.Write(fd477, []byte("GFtU6q2GcIGFVCsxAz2AlI8pSrfGZdTREAz2fblW3IliK3qYqnN9oaXG_mqu_rMqMQbvb0Ow8KOwnWnzpNybhlRJdKT3huEWGb9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (279) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn9_NIRqSsCo1kZxl_rpWv2Sh8loZFUmEh5_fscTpoo4MOAkAHJjuWo4Pb4u25Mx0KS1rZXRT6k5K3mqh6zHh6yiBJiGFtU6q2GcIGFVCsxAz2AlI8pSrfGZdTREAz2fblW3IliK3qYqnN9oaXG_mqu_rMqMQbvb0Ow8KOwnWnzpNybhlRJdKT3huEWGb9

ret = client.Close(fd485)
if(ret != 0) {
  panic("close failed")
}


fd486 := client.Open("/hpYA64tQVQ.txt", client.O_RDWR|client.O_CREATE)
if(fd486 < 0) {
  panic("open failed")
}


fd487 := client.Open("/vcU3VuvRAJ.txt", client.O_RDWR|client.O_CREATE)
if(fd487 < 0) {
  panic("open failed")
}


fd488 := client.Open("/LzEl4xW9O6.txt", client.O_RDWR|client.O_CREATE)
if(fd488 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd487, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd483, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd486, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd488, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd487)
if(ret != 0) {
  panic("close failed")
}


fd489 := client.Open("/Do8eCB4JFH.txt", client.O_RDWR|client.O_CREATE)
if(fd489 < 0) {
  panic("open failed")
}

//fd state: (279) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn9_NIRqSsCo1kZxl_rpWv2Sh8loZFUmEh5_fscTpoo4MOAkAHJjuWo4Pb4u25Mx0KS1rZXRT6k5K3mqh6zHh6yiBJiGFtU6q2GcIGFVCsxAz2AlI8pSrfGZdTREAz2fblW3IliK3qYqnN9oaXG_mqu_rMqMQbvb0Ow8KOwnWnzpNybhlRJdKT3huEWGb9

ret = client.Write(fd477, []byte("BQdPt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (284) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn9_NIRqSsCo1kZxl_rpWv2Sh8loZFUmEh5_fscTpoo4MOAkAHJjuWo4Pb4u25Mx0KS1rZXRT6k5K3mqh6zHh6yiBJiGFtU6q2GcIGFVCsxAz2AlI8pSrfGZdTREAz2fblW3IliK3qYqnN9oaXG_mqu_rMqMQbvb0Ow8KOwnWnzpNybhlRJdKT3huEWGb9BQdPt

fd490 := client.Open("/ZXm12GGVei.txt", client.O_RDWR|client.O_CREATE)
if(fd490 < 0) {
  panic("open failed")
}

//fd state: (89) ng9qilxe5SOdCWytuT9g6ni6HHPgWxEjkzl4Aj2EhGRq_euNurUKi7Bj_4TkcJNmEJobY1MpfwnfbsOr_5EEx2DB3

ret = client.Write(fd483, []byte("jrRHD8Bj9k15byN3yXKRLYuZkymgG35rU_MnV6kIhIqtfUNQ2aU7qR1wg01Wp0oL_RpN_H0GIu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (163) ng9qilxe5SOdCWytuT9g6ni6HHPgWxEjkzl4Aj2EhGRq_euNurUKi7Bj_4TkcJNmEJobY1MpfwnfbsOr_5EEx2DB3jrRHD8Bj9k15byN3yXKRLYuZkymgG35rU_MnV6kIhIqtfUNQ2aU7qR1wg01Wp0oL_RpN_H0GIu

buf, ret = client.Read(fd486, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd491 := client.Open("/8c8dueKjTG.txt", client.O_RDWR|client.O_CREATE)
if(fd491 < 0) {
  panic("open failed")
}


fd492 := client.Open("/PJprT8kWfg.txt", client.O_RDWR|client.O_CREATE)
if(fd492 < 0) {
  panic("open failed")
}


fd493 := client.Open("/f38bZ7Efc2.txt", client.O_RDWR|client.O_CREATE)
if(fd493 < 0) {
  panic("open failed")
}


fd494 := client.Open("/6HMIdOkz5H.txt", client.O_RDWR|client.O_CREATE)
if(fd494 < 0) {
  panic("open failed")
}


fd495 := client.Open("/Ti8vk7aoMw.txt", client.O_RDWR|client.O_CREATE)
if(fd495 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd489, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd483)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd489, []byte("0jRvYQjty4Hk_pECHkOEg21HQAg5NE7I2A7e2zH6_snql6ys2sslshuPA0mMKn5B"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) 0jRvYQjty4Hk_pECHkOEg21HQAg5NE7I2A7e2zH6_snql6ys2sslshuPA0mMKn5B

buf, ret = client.Read(fd477, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd490, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd495, []byte("Hs9CZJ1Aj15UbF5_UWzCpD56svpuGGhBGn49mAL_cYaBmIym4J"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) Hs9CZJ1Aj15UbF5_UWzCpD56svpuGGhBGn49mAL_cYaBmIym4J

ret = client.Seek(fd494, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd490)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd492)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd493, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


fd496 := client.Open("/x40tD7C7XY.txt", client.O_RDWR|client.O_CREATE)
if(fd496 < 0) {
  panic("open failed")
}

//fd state: (284) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn9_NIRqSsCo1kZxl_rpWv2Sh8loZFUmEh5_fscTpoo4MOAkAHJjuWo4Pb4u25Mx0KS1rZXRT6k5K3mqh6zHh6yiBJiGFtU6q2GcIGFVCsxAz2AlI8pSrfGZdTREAz2fblW3IliK3qYqnN9oaXG_mqu_rMqMQbvb0Ow8KOwnWnzpNybhlRJdKT3huEWGb9BQdPt

ret = client.Write(fd477, []byte("NJ8sIzt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (291) YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn9_NIRqSsCo1kZxl_rpWv2Sh8loZFUmEh5_fscTpoo4MOAkAHJjuWo4Pb4u25Mx0KS1rZXRT6k5K3mqh6zHh6yiBJiGFtU6q2GcIGFVCsxAz2AlI8pSrfGZdTREAz2fblW3IliK3qYqnN9oaXG_mqu_rMqMQbvb0Ow8KOwnWnzpNybhlRJdKT3huEWGb9BQdPtNJ8sIzt

buf, ret = client.Read(fd493, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd494, []byte("aQ6A2LxHRWtgYLL653h8gjxIQLFqpvC5x68cEDX2vareJsUW1WP3j9EOjqsYRlhBZCepEeOu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) aQ6A2LxHRWtgYLL653h8gjxIQLFqpvC5x68cEDX2vareJsUW1WP3j9EOjqsYRlhBZCepEeOu

fd497 := client.Open("/ULwj3UEeoU.txt", client.O_RDWR|client.O_CREATE)
if(fd497 < 0) {
  panic("open failed")
}


ret = client.Seek(fd496, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd498 := client.Open("/VCgdE3WhbF.txt", client.O_RDWR|client.O_CREATE)
if(fd498 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd494, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd477)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd495, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


fd499 := client.Open("/qFVLiD0BQj.txt", client.O_RDWR|client.O_CREATE)
if(fd499 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd493, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) Gz3u5Ivuhtq4TfPwPVojsU1rHgOBUMuCkOSWuj8C9xbeuNO

ret = client.Write(fd498, []byte("IYca"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) IYca5Ivuhtq4TfPwPVojsU1rHgOBUMuCkOSWuj8C9xbeuNO
//fd state: (0) 

ret = client.Write(fd497, []byte("ms73HMrSi6mYTSGVtIG4lAaAHa4ALLNGBHlfV0WOZKYQDmWeEEpChVZxsW2Hq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) ms73HMrSi6mYTSGVtIG4lAaAHa4ALLNGBHlfV0WOZKYQDmWeEEpChVZxsW2Hq

buf, ret = client.Read(fd495, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "n49mAL_cYaBmIym4J") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd486, []byte("Cd2achxSHDR1izzhfxMt1UwhWbNL3F5P6XfIaL1ydJgXy1jai2lEpTAPe0vKGo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) Cd2achxSHDR1izzhfxMt1UwhWbNL3F5P6XfIaL1ydJgXy1jai2lEpTAPe0vKGo

ret = client.Seek(fd498, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


fd500 := client.Open("/fTKEpwYowY.txt", client.O_RDWR|client.O_CREATE)
if(fd500 < 0) {
  panic("open failed")
}


ret = client.Seek(fd486, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd491, []byte("oA7rRKmCA7Iw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) oA7rRKmCA7Iw

ret = client.Close(fd497)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd495)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd500, []byte("TNcNBR2ndBLPw8f3Q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) TNcNBR2ndBLPw8f3Q
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd496)
if(ret != 0) {
  panic("close failed")
}

//fd state: (72) aQ6A2LxHRWtgYLL653h8gjxIQLFqpvC5x68cEDX2vareJsUW1WP3j9EOjqsYRlhBZCepEeOu

ret = client.Write(fd494, []byte("5Vbs7f5RVXNAmsjbFDssUvMOn5JKh7OwIKlSvMMtYKQVsEutEbutP_WKfXAcRc39oopmotYHxM2r5s12i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (153) aQ6A2LxHRWtgYLL653h8gjxIQLFqpvC5x68cEDX2vareJsUW1WP3j9EOjqsYRlhBZCepEeOu5Vbs7f5RVXNAmsjbFDssUvMOn5JKh7OwIKlSvMMtYKQVsEutEbutP_WKfXAcRc39oopmotYHxM2r5s12i

ret = client.Close(fd488)
if(ret != 0) {
  panic("close failed")
}

//fd state: (1) K

ret = client.Write(fd493, []byte("OLD9ESj2I2rBu49yWXIXiLk0SnHYMmL9w4InUFneFEKEaiubF3y5TKgyC3FvCzOjVUPlApozPzRN9S3L7Hc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) KOLD9ESj2I2rBu49yWXIXiLk0SnHYMmL9w4InUFneFEKEaiubF3y5TKgyC3FvCzOjVUPlApozPzRN9S3L7Hc

buf, ret = client.Read(fd489, 9)
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


buf, ret = client.Read(fd494, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd489, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


buf, ret = client.Read(fd493, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd500, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Seek(fd499, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (12) oA7rRKmCA7Iw

ret = client.Write(fd491, []byte("XBVIzPFqxKcsGxM62NCT9owA1UVBlqj0dIrKvSXKKAFliFDVLlvJ5Q_y9Mp9O6FKirDbdEXUBnBYbAe15"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) oA7rRKmCA7IwXBVIzPFqxKcsGxM62NCT9owA1UVBlqj0dIrKvSXKKAFliFDVLlvJ5Q_y9Mp9O6FKirDbdEXUBnBYbAe15

buf, ret = client.Read(fd500, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "cNBR2ndBLPw8f3Q") {
  panic("wrong data returned")
}


ret = client.Seek(fd489, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd500)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd493, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (7) 0jRvYQjty4Hk_pECHkOEg21HQAg5NE7I2A7e2zH6_snql6ys2sslshuPA0mMKn5B

ret = client.Write(fd489, []byte("7URR5zsbLHJzOmVteqCW6CppHJm3iRBRnGKzd0RrigjvXEzOuWqrzLjHrd5Uq3mSv_3xh9muk4_1gE7JLGUF9Qtx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) 0jRvYQj7URR5zsbLHJzOmVteqCW6CppHJm3iRBRnGKzd0RrigjvXEzOuWqrzLjHrd5Uq3mSv_3xh9muk4_1gE7JLGUF9Qtx

ret = client.Close(fd486)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd499, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (153) aQ6A2LxHRWtgYLL653h8gjxIQLFqpvC5x68cEDX2vareJsUW1WP3j9EOjqsYRlhBZCepEeOu5Vbs7f5RVXNAmsjbFDssUvMOn5JKh7OwIKlSvMMtYKQVsEutEbutP_WKfXAcRc39oopmotYHxM2r5s12i

ret = client.Write(fd494, []byte("_V4spZIoHLZrwXCw3cDu0QN5UE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (179) aQ6A2LxHRWtgYLL653h8gjxIQLFqpvC5x68cEDX2vareJsUW1WP3j9EOjqsYRlhBZCepEeOu5Vbs7f5RVXNAmsjbFDssUvMOn5JKh7OwIKlSvMMtYKQVsEutEbutP_WKfXAcRc39oopmotYHxM2r5s12i_V4spZIoHLZrwXCw3cDu0QN5UE

fd501 := client.Open("/ez2_D3d8Wg.txt", client.O_RDWR|client.O_CREATE)
if(fd501 < 0) {
  panic("open failed")
}


ret = client.Close(fd494)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd501, []byte("hqmFqvvnvwkoUoOop9PYslqCSIMek_zjvBmU_itMRCkazMgeqCyGwdH_HjF5hw7q1zvLZPKLi9TviC6pOV2S_HVlH7GO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) hqmFqvvnvwkoUoOop9PYslqCSIMek_zjvBmU_itMRCkazMgeqCyGwdH_HjF5hw7q1zvLZPKLi9TviC6pOV2S_HVlH7GO

ret = client.Close(fd491)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd501)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd499)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd493)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd489, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


ret = client.Close(fd489)
if(ret != 0) {
  panic("close failed")
}


fd502 := client.Open("/x40tD7C7XY.txt", client.O_RDWR|client.O_CREATE)
if(fd502 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd502, []byte("X8t3AYJNJ16ytW6qFtl1AF4pUqS9qTDYlLbjgxaH6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) X8t3AYJNJ16ytW6qFtl1AF4pUqS9qTDYlLbjgxaH6

fd503 := client.Open("/EPIBJZLW_w.txt", client.O_RDWR|client.O_CREATE)
if(fd503 < 0) {
  panic("open failed")
}


fd504 := client.Open("/NQGd3iEtki.txt", client.O_RDWR|client.O_CREATE)
if(fd504 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd503, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5NapMG4B3oXg5jUu6OyqzQwDyD_t") {
  panic("wrong data returned")
}


ret = client.Close(fd504)
if(ret != 0) {
  panic("close failed")
}


fd505 := client.Open("/WvrUAZ5kOA.txt", client.O_RDWR|client.O_CREATE)
if(fd505 < 0) {
  panic("open failed")
}


ret = client.Seek(fd505, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd502, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd506 := client.Open("/zjkNSSD_IZ.txt", client.O_RDWR|client.O_CREATE)
if(fd506 < 0) {
  panic("open failed")
}


ret = client.Close(fd503)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) _4drx69HAP8gEJ9J7pdAEjHYHxxIrONED5TCcfGKgtVcoZ2RTLFjLxnAtduf

ret = client.Write(fd506, []byte("Cq2kNiE1sbMFEoTuZJeTlIvlQOcRCdjximM9hLn1RWHTT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) Cq2kNiE1sbMFEoTuZJeTlIvlQOcRCdjximM9hLn1RWHTTZ2RTLFjLxnAtduf

buf, ret = client.Read(fd502, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd505, 63)
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


ret = client.Seek(fd502, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}


ret = client.Seek(fd506, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


buf, ret = client.Read(fd506, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WHTTZ2RTLFjLxnAtduf") {
  panic("wrong data returned")
}

//fd state: (6) X8t3AYJNJ16ytW6qFtl1AF4pUqS9qTDYlLbjgxaH6

ret = client.Write(fd502, []byte("LVsxSzPjfE7HMD4teMOBJD8bNa2DWClM6wufGbUVlLuoPp5bg4H_X2udkW4Ormg4Kv5mn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) X8t3AYLVsxSzPjfE7HMD4teMOBJD8bNa2DWClM6wufGbUVlLuoPp5bg4H_X2udkW4Ormg4Kv5mn

ret = client.Seek(fd502, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


fd507 := client.Open("/ez2_D3d8Wg.txt", client.O_RDWR|client.O_CREATE)
if(fd507 < 0) {
  panic("open failed")
}


ret = client.Close(fd502)
if(ret != 0) {
  panic("close failed")
}

//fd state: (60) Cq2kNiE1sbMFEoTuZJeTlIvlQOcRCdjximM9hLn1RWHTTZ2RTLFjLxnAtduf

ret = client.Write(fd506, []byte("RpBz4RJvsJlB23UcVrLqYcgxLKEZKSYkwQn64N236HDnh760OKcWKqJgUSePNsqfCWx8HlSC01BlDq2fb7nEbD1K5S1JGr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) Cq2kNiE1sbMFEoTuZJeTlIvlQOcRCdjximM9hLn1RWHTTZ2RTLFjLxnAtdufRpBz4RJvsJlB23UcVrLqYcgxLKEZKSYkwQn64N236HDnh760OKcWKqJgUSePNsqfCWx8HlSC01BlDq2fb7nEbD1K5S1JGr

buf, ret = client.Read(fd507, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hqmFqvvnvwkoUoOop9PYslqCSIMek_zjvBmU_itMRCkazMgeqCyGwdH_HjF5hw7q1zvLZPKLi9TviC6pOV2S_HVlH7GO") {
  panic("wrong data returned")
}


fd508 := client.Open("/TzzDnGQvJB.txt", client.O_RDWR|client.O_CREATE)
if(fd508 < 0) {
  panic("open failed")
}

//fd state: (154) Cq2kNiE1sbMFEoTuZJeTlIvlQOcRCdjximM9hLn1RWHTTZ2RTLFjLxnAtdufRpBz4RJvsJlB23UcVrLqYcgxLKEZKSYkwQn64N236HDnh760OKcWKqJgUSePNsqfCWx8HlSC01BlDq2fb7nEbD1K5S1JGr

ret = client.Write(fd506, []byte("7ncPFpqYOr3_ccP4H9Zu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (174) Cq2kNiE1sbMFEoTuZJeTlIvlQOcRCdjximM9hLn1RWHTTZ2RTLFjLxnAtdufRpBz4RJvsJlB23UcVrLqYcgxLKEZKSYkwQn64N236HDnh760OKcWKqJgUSePNsqfCWx8HlSC01BlDq2fb7nEbD1K5S1JGr7ncPFpqYOr3_ccP4H9Zu

ret = client.Close(fd507)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd508)
if(ret != 0) {
  panic("close failed")
}


fd509 := client.Open("/WxPQuN5oUi.txt", client.O_RDWR|client.O_CREATE)
if(fd509 < 0) {
  panic("open failed")
}

//fd state: (174) Cq2kNiE1sbMFEoTuZJeTlIvlQOcRCdjximM9hLn1RWHTTZ2RTLFjLxnAtdufRpBz4RJvsJlB23UcVrLqYcgxLKEZKSYkwQn64N236HDnh760OKcWKqJgUSePNsqfCWx8HlSC01BlDq2fb7nEbD1K5S1JGr7ncPFpqYOr3_ccP4H9Zu

ret = client.Write(fd506, []byte("D2zE6nxqSUsX8kqKTW5odmsIpouKJlpfk2CdDea9d8OFuLv4fGY56EeKFXYBzytFzzJCdCKbONTBiJzk7t4hII"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (260) Cq2kNiE1sbMFEoTuZJeTlIvlQOcRCdjximM9hLn1RWHTTZ2RTLFjLxnAtdufRpBz4RJvsJlB23UcVrLqYcgxLKEZKSYkwQn64N236HDnh760OKcWKqJgUSePNsqfCWx8HlSC01BlDq2fb7nEbD1K5S1JGr7ncPFpqYOr3_ccP4H9ZuD2zE6nxqSUsX8kqKTW5odmsIpouKJlpfk2CdDea9d8OFuLv4fGY56EeKFXYBzytFzzJCdCKbONTBiJzk7t4hII

buf, ret = client.Read(fd509, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd509)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd506, 53, client.SEEK_SET)
if(ret != 53) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 53)
  panic("seek failed")
}


ret = client.Close(fd506)
if(ret != 0) {
  panic("close failed")
}


fd510 := client.Open("/IggtggNFV3.txt", client.O_RDWR|client.O_CREATE)
if(fd510 < 0) {
  panic("open failed")
}


fd511 := client.Open("/dw3wX6Mc3X.txt", client.O_RDWR|client.O_CREATE)
if(fd511 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd511, []byte("ddYTrLzh0Eb7w_2vKsZUT6jtHiBafDV2wFZJk8kMCngmss574AXvqjxz529lqzkRWpg201H63eaL1oaWcdWUJfrrag0Zl55_YK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) ddYTrLzh0Eb7w_2vKsZUT6jtHiBafDV2wFZJk8kMCngmss574AXvqjxz529lqzkRWpg201H63eaL1oaWcdWUJfrrag0Zl55_YK

buf, ret = client.Read(fd510, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd511, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


buf, ret = client.Read(fd511, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "HiBafDV2wFZJk8kMCngmss574AXvqjxz529lq") {
  panic("wrong data returned")
}


ret = client.Close(fd511)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd510, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd510)
if(ret != 0) {
  panic("close failed")
}


fd512 := client.Open("/E34Lq2h34a.txt", client.O_RDWR|client.O_CREATE)
if(fd512 < 0) {
  panic("open failed")
}


ret = client.Seek(fd512, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd512, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd513 := client.Open("/ATxjoIZR9n.txt", client.O_RDWR|client.O_CREATE)
if(fd513 < 0) {
  panic("open failed")
}


ret = client.Close(fd513)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd512, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd512, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd512)
if(ret != 0) {
  panic("close failed")
}


fd514 := client.Open("/bmPU4CEsRL.txt", client.O_RDWR|client.O_CREATE)
if(fd514 < 0) {
  panic("open failed")
}

//fd state: (0) YI4x3IZsZqhTW8DMoFhbYkhtavwgIuGvupDyTB6wfuhYTlcp6beNFPwjZUjs9CUS6slo

ret = client.Write(fd514, []byte("Shb7LbE6vnUuDw7k4Ntu7dDD3_DpaqyOzZc1cogEw78YwvuM1ZAD471"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) Shb7LbE6vnUuDw7k4Ntu7dDD3_DpaqyOzZc1cogEw78YwvuM1ZAD471jZUjs9CUS6slo
//fd state: (55) Shb7LbE6vnUuDw7k4Ntu7dDD3_DpaqyOzZc1cogEw78YwvuM1ZAD471jZUjs9CUS6slo

ret = client.Write(fd514, []byte("1I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) Shb7LbE6vnUuDw7k4Ntu7dDD3_DpaqyOzZc1cogEw78YwvuM1ZAD4711IUjs9CUS6slo

buf, ret = client.Read(fd514, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Ujs9CUS6slo") {
  panic("wrong data returned")
}


ret = client.Close(fd514)
if(ret != 0) {
  panic("close failed")
}


fd515 := client.Open("/PtxNuQ98Qr.txt", client.O_RDWR|client.O_CREATE)
if(fd515 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd515, 59)
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


fd516 := client.Open("/MhNRYjnYbb.txt", client.O_RDWR|client.O_CREATE)
if(fd516 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd516, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5XmdnaNO1MFJPrC3HiefPzPah5K4SFlK9wx9j") {
  panic("wrong data returned")
}


ret = client.Close(fd516)
if(ret != 0) {
  panic("close failed")
}


fd517 := client.Open("/CO9WO8YU1D.txt", client.O_RDWR|client.O_CREATE)
if(fd517 < 0) {
  panic("open failed")
}


ret = client.Seek(fd517, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd517, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd517)
if(ret != 0) {
  panic("close failed")
}


fd518 := client.Open("/qT0qiiJ4v2.txt", client.O_RDWR|client.O_CREATE)
if(fd518 < 0) {
  panic("open failed")
}


fd519 := client.Open("/BNitlfmh8Q.txt", client.O_RDWR|client.O_CREATE)
if(fd519 < 0) {
  panic("open failed")
}


fd520 := client.Open("/QpV8PqvRKB.txt", client.O_RDWR|client.O_CREATE)
if(fd520 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd520, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yj") {
  panic("wrong data returned")
}


ret = client.Seek(fd519, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd518, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd518, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd521 := client.Open("/fFJKS3nhaA.txt", client.O_RDWR|client.O_CREATE)
if(fd521 < 0) {
  panic("open failed")
}


ret = client.Seek(fd519, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (2) yj

ret = client.Write(fd520, []byte("Z0p5mx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) yjZ0p5mx

ret = client.Close(fd520)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd521, []byte("yxC5w1ALc0i6ZplT5CD_GPAJgLqwUcizDHE2fw9eVuroiVk77Zq3F5g2xy4Q3ZYX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) yxC5w1ALc0i6ZplT5CD_GPAJgLqwUcizDHE2fw9eVuroiVk77Zq3F5g2xy4Q3ZYX

fd522 := client.Open("/mPd5U56D9a.txt", client.O_RDWR|client.O_CREATE)
if(fd522 < 0) {
  panic("open failed")
}


fd523 := client.Open("/bCrAMTyNaU.txt", client.O_RDWR|client.O_CREATE)
if(fd523 < 0) {
  panic("open failed")
}


ret = client.Seek(fd518, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd518, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd519, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd519, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd522)
if(ret != 0) {
  panic("close failed")
}


fd524 := client.Open("/mdASAirDD_.txt", client.O_RDWR|client.O_CREATE)
if(fd524 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd518, []byte("PEkd7_j8MAEbc1dYTkV0SHBxkPa3G7ddcTcv1uaawcIEX2QyWropVA1Xzg2nqMc_wGlNhQQpqc9wHyVrt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) PEkd7_j8MAEbc1dYTkV0SHBxkPa3G7ddcTcv1uaawcIEX2QyWropVA1Xzg2nqMc_wGlNhQQpqc9wHyVrt

fd525 := client.Open("/M6lcoqPuI6.txt", client.O_RDWR|client.O_CREATE)
if(fd525 < 0) {
  panic("open failed")
}


ret = client.Close(fd525)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd521, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}

//fd state: (81) PEkd7_j8MAEbc1dYTkV0SHBxkPa3G7ddcTcv1uaawcIEX2QyWropVA1Xzg2nqMc_wGlNhQQpqc9wHyVrt

ret = client.Write(fd518, []byte("oFu3irgZHwfLLW6UGigb3hGclX4ahNZHruoxXfhOER2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) PEkd7_j8MAEbc1dYTkV0SHBxkPa3G7ddcTcv1uaawcIEX2QyWropVA1Xzg2nqMc_wGlNhQQpqc9wHyVrtoFu3irgZHwfLLW6UGigb3hGclX4ahNZHruoxXfhOER2

ret = client.Close(fd518)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd519, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd521, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


ret = client.Close(fd521)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd519, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd524, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}


fd526 := client.Open("/gKjRc6XTZ4.txt", client.O_RDWR|client.O_CREATE)
if(fd526 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd523, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd524, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


buf, ret = client.Read(fd523, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (41) WJ_ge0cWSi70VMJMxhZFJvMp2GPNNDxMJWJd3juE5iVvEDtFSouKS7WCwjt

ret = client.Write(fd524, []byte("tSB3O0xY8PHYvUBksNSSyF6OskkpyZgrKyWqEV9DtbPFOSGd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) WJ_ge0cWSi70VMJMxhZFJvMp2GPNNDxMJWJd3juE5tSB3O0xY8PHYvUBksNSSyF6OskkpyZgrKyWqEV9DtbPFOSGd

buf, ret = client.Read(fd523, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd519, []byte("LgTKjZ9Mx4OedT6Iq9Gn4vZQllthA6f1qV1GyBC63SFoKmXstLQmUcFKgqKOI_U1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) LgTKjZ9Mx4OedT6Iq9Gn4vZQllthA6f1qV1GyBC63SFoKmXstLQmUcFKgqKOI_U1

ret = client.Close(fd523)
if(ret != 0) {
  panic("close failed")
}


fd527 := client.Open("/7bFIx4fJhh.txt", client.O_RDWR|client.O_CREATE)
if(fd527 < 0) {
  panic("open failed")
}


ret = client.Close(fd526)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd524, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd519, 64, client.SEEK_SET)
if(ret != 64) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 64)
  panic("seek failed")
}


fd528 := client.Open("/4BCfB7iO6C.txt", client.O_RDWR|client.O_CREATE)
if(fd528 < 0) {
  panic("open failed")
}


ret = client.Seek(fd524, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}


ret = client.Close(fd528)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd524, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "DtbPFOSGd") {
  panic("wrong data returned")
}


ret = client.Close(fd527)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd519, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}

//fd state: (13) LgTKjZ9Mx4OedT6Iq9Gn4vZQllthA6f1qV1GyBC63SFoKmXstLQmUcFKgqKOI_U1

ret = client.Write(fd519, []byte("y5Ba6C42EufsmgNYi31WJfASu9CpDH4te"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) LgTKjZ9Mx4Oedy5Ba6C42EufsmgNYi31WJfASu9CpDH4teXstLQmUcFKgqKOI_U1
//fd state: (46) LgTKjZ9Mx4Oedy5Ba6C42EufsmgNYi31WJfASu9CpDH4teXstLQmUcFKgqKOI_U1

ret = client.Write(fd519, []byte("qbIXA7emF5TMoP9qN3LH4h_VUqHHGENeiswiaaS0mCtJnYFdGwLnlC5LhOYx9Ie5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) LgTKjZ9Mx4Oedy5Ba6C42EufsmgNYi31WJfASu9CpDH4teqbIXA7emF5TMoP9qN3LH4h_VUqHHGENeiswiaaS0mCtJnYFdGwLnlC5LhOYx9Ie5

ret = client.Seek(fd519, 102, client.SEEK_SET)
if(ret != 102) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 102)
  panic("seek failed")
}


ret = client.Close(fd519)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd524, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd524, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}

//fd state: (9) WJ_ge0cWSi70VMJMxhZFJvMp2GPNNDxMJWJd3juE5tSB3O0xY8PHYvUBksNSSyF6OskkpyZgrKyWqEV9DtbPFOSGd

ret = client.Write(fd524, []byte("oIav26IkzCt17297gdoTkhmw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) WJ_ge0cWSoIav26IkzCt17297gdoTkhmwWJd3juE5tSB3O0xY8PHYvUBksNSSyF6OskkpyZgrKyWqEV9DtbPFOSGd

buf, ret = client.Read(fd524, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WJd3juE5tSB3O0xY8PHYvUBksNS") {
  panic("wrong data returned")
}


fd529 := client.Open("/JJw95P1AZF.txt", client.O_RDWR|client.O_CREATE)
if(fd529 < 0) {
  panic("open failed")
}

//fd state: (60) WJ_ge0cWSoIav26IkzCt17297gdoTkhmwWJd3juE5tSB3O0xY8PHYvUBksNSSyF6OskkpyZgrKyWqEV9DtbPFOSGd

ret = client.Write(fd524, []byte("xScAAC3PcdWnoi4UlDB1s8vKVMo0jQk0dloVqZRKD_BsiBYhJqkPRfyy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) WJ_ge0cWSoIav26IkzCt17297gdoTkhmwWJd3juE5tSB3O0xY8PHYvUBksNSxScAAC3PcdWnoi4UlDB1s8vKVMo0jQk0dloVqZRKD_BsiBYhJqkPRfyy

fd530 := client.Open("/50VBW3gKyR.txt", client.O_RDWR|client.O_CREATE)
if(fd530 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd524, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd529, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "OA7Dw511gPquGZQIjCwsUdug_CdPI81JT9NE3eA") {
  panic("wrong data returned")
}


ret = client.Close(fd524)
if(ret != 0) {
  panic("close failed")
}


fd531 := client.Open("/LgDXLTVcW9.txt", client.O_RDWR|client.O_CREATE)
if(fd531 < 0) {
  panic("open failed")
}


ret = client.Seek(fd529, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}

//fd state: (0) ODgLvL1lOvNlxhyGun0uHKZawIcCuK1VotUdZUhrrKEWYfLQnDKPTKszs9cCEJH5mgmF8T8IAXvSGOEkEE7oslN7ldz7W8sipoHHSeZsVFlFWoHPprxSJj05qIuHv4kMnQc6Y2O_lb9L

ret = client.Write(fd531, []byte("g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8s9LIPzrPxKtTcykmD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8s9LIPzrPxKtTcykmDlN7ldz7W8sipoHHSeZsVFlFWoHPprxSJj05qIuHv4kMnQc6Y2O_lb9L

ret = client.Seek(fd530, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (24) OA7Dw511gPquGZQIjCwsUdug_CdPI81JT9NE3eA

ret = client.Write(fd529, []byte("B2H3stNinfhjaG1auZbaVFmFPXwlknFP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) OA7Dw511gPquGZQIjCwsUdugB2H3stNinfhjaG1auZbaVFmFPXwlknFP

fd532 := client.Open("/V2hqe83j8X.txt", client.O_RDWR|client.O_CREATE)
if(fd532 < 0) {
  panic("open failed")
}

//fd state: (56) OA7Dw511gPquGZQIjCwsUdugB2H3stNinfhjaG1auZbaVFmFPXwlknFP

ret = client.Write(fd529, []byte("LTP9XVRifnzdlcTcimGTjpUp4payaskb1Ps7IUSK3it58J0FEZQG41ivpJwzzAHcVkTuLmgUfCEPa4XC2Y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (138) OA7Dw511gPquGZQIjCwsUdugB2H3stNinfhjaG1auZbaVFmFPXwlknFPLTP9XVRifnzdlcTcimGTjpUp4payaskb1Ps7IUSK3it58J0FEZQG41ivpJwzzAHcVkTuLmgUfCEPa4XC2Y

ret = client.Seek(fd532, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


buf, ret = client.Read(fd532, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lYOSTwxZRUGzb1gnqsziLzk2cV5g2q8Q9mI2t79UWjQ9o7sBqwdlT5SEj8AHaTEimxhiMZShiDp4sot0i") {
  panic("wrong data returned")
}


ret = client.Close(fd529)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd532, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Pcp65HeRPxXIDxshG8uJdJ5swbO4n") {
  panic("wrong data returned")
}


ret = client.Close(fd530)
if(ret != 0) {
  panic("close failed")
}

//fd state: (177) ojMm3KphHvyOUztZl5HW0NHtMAiggvU56l_ri5ALj5YiYWCYLP9jVa9aczM4U5fACE8lYOSTwxZRUGzb1gnqsziLzk2cV5g2q8Q9mI2t79UWjQ9o7sBqwdlT5SEj8AHaTEimxhiMZShiDp4sot0iPcp65HeRPxXIDxshG8uJdJ5swbO4nR3tcFtvJU7O3themprK5rBQl7nUi0g3xyszukDeP

ret = client.Write(fd532, []byte("fGXW1wakTmOq_pGzFFwddunWPAbcHHDV63rgHwCkPgn1nDFOrkv4eKL0blB5Ab3uOGmPVyZOZ_X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (252) ojMm3KphHvyOUztZl5HW0NHtMAiggvU56l_ri5ALj5YiYWCYLP9jVa9aczM4U5fACE8lYOSTwxZRUGzb1gnqsziLzk2cV5g2q8Q9mI2t79UWjQ9o7sBqwdlT5SEj8AHaTEimxhiMZShiDp4sot0iPcp65HeRPxXIDxshG8uJdJ5swbO4nfGXW1wakTmOq_pGzFFwddunWPAbcHHDV63rgHwCkPgn1nDFOrkv4eKL0blB5Ab3uOGmPVyZOZ_X

ret = client.Seek(fd532, 207, client.SEEK_SET)
if(ret != 207) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 207)
  panic("seek failed")
}


fd533 := client.Open("/uIv3bq5PSV.txt", client.O_RDWR|client.O_CREATE)
if(fd533 < 0) {
  panic("open failed")
}


ret = client.Seek(fd533, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd532)
if(ret != 0) {
  panic("close failed")
}

//fd state: (85) g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8s9LIPzrPxKtTcykmDlN7ldz7W8sipoHHSeZsVFlFWoHPprxSJj05qIuHv4kMnQc6Y2O_lb9L

ret = client.Write(fd531, []byte("6mmQ0lFLGCHoH4fRBVMLC3BDBa0V2oEZyyHgK5GxkMsRpgpJlWywdwIUQyarErON8C7Q3jOXwaebhE3YeASIiMmqDzbrPw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (179) g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8s9LIPzrPxKtTcykmD6mmQ0lFLGCHoH4fRBVMLC3BDBa0V2oEZyyHgK5GxkMsRpgpJlWywdwIUQyarErON8C7Q3jOXwaebhE3YeASIiMmqDzbrPw

fd534 := client.Open("/ModFGCjTvL.txt", client.O_RDWR|client.O_CREATE)
if(fd534 < 0) {
  panic("open failed")
}


ret = client.Seek(fd534, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd535 := client.Open("/ALsGXybpGN.txt", client.O_RDWR|client.O_CREATE)
if(fd535 < 0) {
  panic("open failed")
}


ret = client.Close(fd535)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd531, 104, client.SEEK_SET)
if(ret != 104) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 104)
  panic("seek failed")
}


buf, ret = client.Read(fd531, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LC3BDBa0V2oEZyyHgK5GxkMsRpgpJlWywdwIU") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd533, []byte("VwhCGYzEkyAcrqGa3Cajsts0TXpywUzty6PE6ctzYvXYtjAMLfcjxjqwMIp8zd2Dkuw3x"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) VwhCGYzEkyAcrqGa3Cajsts0TXpywUzty6PE6ctzYvXYtjAMLfcjxjqwMIp8zd2Dkuw3x

ret = client.Close(fd534)
if(ret != 0) {
  panic("close failed")
}

//fd state: (141) g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8s9LIPzrPxKtTcykmD6mmQ0lFLGCHoH4fRBVMLC3BDBa0V2oEZyyHgK5GxkMsRpgpJlWywdwIUQyarErON8C7Q3jOXwaebhE3YeASIiMmqDzbrPw

ret = client.Write(fd531, []byte("4_f"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8s9LIPzrPxKtTcykmD6mmQ0lFLGCHoH4fRBVMLC3BDBa0V2oEZyyHgK5GxkMsRpgpJlWywdwIU4_frErON8C7Q3jOXwaebhE3YeASIiMmqDzbrPw

fd536 := client.Open("/bscOPxCqaM.txt", client.O_RDWR|client.O_CREATE)
if(fd536 < 0) {
  panic("open failed")
}


ret = client.Close(fd533)
if(ret != 0) {
  panic("close failed")
}

//fd state: (144) g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8s9LIPzrPxKtTcykmD6mmQ0lFLGCHoH4fRBVMLC3BDBa0V2oEZyyHgK5GxkMsRpgpJlWywdwIU4_frErON8C7Q3jOXwaebhE3YeASIiMmqDzbrPw

ret = client.Write(fd531, []byte("KJ3_VyxZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8s9LIPzrPxKtTcykmD6mmQ0lFLGCHoH4fRBVMLC3BDBa0V2oEZyyHgK5GxkMsRpgpJlWywdwIU4_fKJ3_VyxZQ3jOXwaebhE3YeASIiMmqDzbrPw

ret = client.Seek(fd536, 86, client.SEEK_SET)
if(ret != 86) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 86)
  panic("seek failed")
}


fd537 := client.Open("/fGWRLDpC7z.txt", client.O_RDWR|client.O_CREATE)
if(fd537 < 0) {
  panic("open failed")
}


fd538 := client.Open("/HbRv9wuACc.txt", client.O_RDWR|client.O_CREATE)
if(fd538 < 0) {
  panic("open failed")
}

//fd state: (152) g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8s9LIPzrPxKtTcykmD6mmQ0lFLGCHoH4fRBVMLC3BDBa0V2oEZyyHgK5GxkMsRpgpJlWywdwIU4_fKJ3_VyxZQ3jOXwaebhE3YeASIiMmqDzbrPw

ret = client.Write(fd531, []byte("8nFwSM9vdLlyCyHM9n3IRoRT2hPV8SueBccXbIvJi36BCwgpcMK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (203) g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8s9LIPzrPxKtTcykmD6mmQ0lFLGCHoH4fRBVMLC3BDBa0V2oEZyyHgK5GxkMsRpgpJlWywdwIU4_fKJ3_VyxZ8nFwSM9vdLlyCyHM9n3IRoRT2hPV8SueBccXbIvJi36BCwgpcMK

fd539 := client.Open("/baO6iJtSGv.txt", client.O_RDWR|client.O_CREATE)
if(fd539 < 0) {
  panic("open failed")
}


ret = client.Seek(fd539, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd539, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd539, []byte("tWTVnhyhGaAbBLeKmtialXiKR73iuVOruk0YkuPN5QJjXeG3IZ2rS0v98p8xvaoByhwifVY0KGdxSrFC7i1HfMCJmQJxm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) tWTVnhyhGaAbBLeKmtialXiKR73iuVOruk0YkuPN5QJjXeG3IZ2rS0v98p8xvaoByhwifVY0KGdxSrFC7i1HfMCJmQJxm

buf, ret = client.Read(fd538, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd536, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "doiUFM7NBhH6iXzypGG51V23mKFTZbJ8lSPo2bPH_nt") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd538, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd537, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd531)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd538, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd538, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (2) 18DWcnHN73PY1xmBhwMkYuBoRg0TAvqtJ3ObQ4OSMVig0MfVafT1X6lpdI4nq32XrM0epB05Xzz4tJEJCnHzQjyE2XSPFuKH

ret = client.Write(fd537, []byte("Svz_aLPNAO_i79m"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) 18Svz_aLPNAO_i79mwMkYuBoRg0TAvqtJ3ObQ4OSMVig0MfVafT1X6lpdI4nq32XrM0epB05Xzz4tJEJCnHzQjyE2XSPFuKH
//fd state: (0) 

ret = client.Write(fd538, []byte("KESRQxv3zzWHwYpYZQmE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) KESRQxv3zzWHwYpYZQmE

fd540 := client.Open("/hHnQzi7E8O.txt", client.O_RDWR|client.O_CREATE)
if(fd540 < 0) {
  panic("open failed")
}


ret = client.Seek(fd537, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Seek(fd540, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (93) tWTVnhyhGaAbBLeKmtialXiKR73iuVOruk0YkuPN5QJjXeG3IZ2rS0v98p8xvaoByhwifVY0KGdxSrFC7i1HfMCJmQJxm

ret = client.Write(fd539, []byte("m6LnsA47A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) tWTVnhyhGaAbBLeKmtialXiKR73iuVOruk0YkuPN5QJjXeG3IZ2rS0v98p8xvaoByhwifVY0KGdxSrFC7i1HfMCJmQJxmm6LnsA47A

ret = client.Close(fd537)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd536, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1ODUjb6O09WXjSxKJYc03t") {
  panic("wrong data returned")
}


fd541 := client.Open("/DRIJ80NCqo.txt", client.O_RDWR|client.O_CREATE)
if(fd541 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd540, []byte("YI1WNOtEDR7VqNbS3t6vYBpBkr4fPVFVzOQvnTozxvrVSNdaaRntzq_xjoqLNnZi6mvt1SBWHTl3NW_6H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (81) YI1WNOtEDR7VqNbS3t6vYBpBkr4fPVFVzOQvnTozxvrVSNdaaRntzq_xjoqLNnZi6mvt1SBWHTl3NW_6H

fd542 := client.Open("/Ku8Le4flu7.txt", client.O_RDWR|client.O_CREATE)
if(fd542 < 0) {
  panic("open failed")
}


ret = client.Seek(fd542, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


ret = client.Close(fd538)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) kYdSWh8W8rbZLJsXgal

ret = client.Write(fd541, []byte("xsgC1XaACfGtnqlCCjVxwOohNdqnoAyksxMULAW_sBiZpu9177uXJ3D3Ehw3whDhKcLiea1mfNAnr5l8XiL8ekP6WCFWRQRs3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) xsgC1XaACfGtnqlCCjVxwOohNdqnoAyksxMULAW_sBiZpu9177uXJ3D3Ehw3whDhKcLiea1mfNAnr5l8XiL8ekP6WCFWRQRs3
//fd state: (97) xsgC1XaACfGtnqlCCjVxwOohNdqnoAyksxMULAW_sBiZpu9177uXJ3D3Ehw3whDhKcLiea1mfNAnr5l8XiL8ekP6WCFWRQRs3

ret = client.Write(fd541, []byte("kuaTuOyiA0j170"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) xsgC1XaACfGtnqlCCjVxwOohNdqnoAyksxMULAW_sBiZpu9177uXJ3D3Ehw3whDhKcLiea1mfNAnr5l8XiL8ekP6WCFWRQRs3kuaTuOyiA0j170

buf, ret = client.Read(fd540, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd543 := client.Open("/hpYA64tQVQ.txt", client.O_RDWR|client.O_CREATE)
if(fd543 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd536, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd544 := client.Open("/YmUgNO_SLh.txt", client.O_RDWR|client.O_CREATE)
if(fd544 < 0) {
  panic("open failed")
}


fd545 := client.Open("/tnVmsNSjJy.txt", client.O_RDWR|client.O_CREATE)
if(fd545 < 0) {
  panic("open failed")
}

//fd state: (0) quffD422E9vksKPRR6IynoZ3j3nJc7CTQuXxHUFc9s0DBVGSALBHyGgwAstibrlUzhjqNNJZT3TYOWlaSFeR858QFSqlN9a0FfMKiKVCmVTVmaCqVNC1cAiaCk1Hc6sNLIl7

ret = client.Write(fd545, []byte("ZywEH3A2RWij4v7Z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (16) ZywEH3A2RWij4v7ZR6IynoZ3j3nJc7CTQuXxHUFc9s0DBVGSALBHyGgwAstibrlUzhjqNNJZT3TYOWlaSFeR858QFSqlN9a0FfMKiKVCmVTVmaCqVNC1cAiaCk1Hc6sNLIl7

fd546 := client.Open("/tX3hPwMq2Q.txt", client.O_RDWR|client.O_CREATE)
if(fd546 < 0) {
  panic("open failed")
}

//fd state: (151) uvgxwz4GdXwI7eFtoTrLia7XXDx0E0ZZ_4gejABrgdcNwD8gLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG51V23mKFTZbJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03t

ret = client.Write(fd536, []byte("V8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (246) uvgxwz4GdXwI7eFtoTrLia7XXDx0E0ZZ_4gejABrgdcNwD8gLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG51V23mKFTZbJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

ret = client.Close(fd544)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd539, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


buf, ret = client.Read(fd539, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "98p8xvaoByhwifVY0") {
  panic("wrong data returned")
}


ret = client.Seek(fd545, 120, client.SEEK_SET)
if(ret != 120) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 120)
  panic("seek failed")
}

//fd state: (0) Cd2achxSHDR1izzhfxMt1UwhWbNL3F5P6XfIaL1ydJgXy1jai2lEpTAPe0vKGo

ret = client.Write(fd543, []byte("Y0ztVwJ1o2F7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) Y0ztVwJ1o2F7izzhfxMt1UwhWbNL3F5P6XfIaL1ydJgXy1jai2lEpTAPe0vKGo

ret = client.Seek(fd536, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


fd547 := client.Open("/JhxzFpd2v3.txt", client.O_RDWR|client.O_CREATE)
if(fd547 < 0) {
  panic("open failed")
}


fd548 := client.Open("/ufMjKnKzWK.txt", client.O_RDWR|client.O_CREATE)
if(fd548 < 0) {
  panic("open failed")
}


ret = client.Seek(fd542, 61, client.SEEK_SET)
if(ret != 61) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 61)
  panic("seek failed")
}


ret = client.Close(fd536)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd548, []byte("juvHcsaWGcX9SPSJDXU7YAAbDO18a2za_cxsbIrc73xo5NSCQsGFTcLOh_EgfMeEV1asA_f_WZo1eJxsPOowEbNbblVAQUkdRv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) juvHcsaWGcX9SPSJDXU7YAAbDO18a2za_cxsbIrc73xo5NSCQsGFTcLOh_EgfMeEV1asA_f_WZo1eJxsPOowEbNbblVAQUkdRv

buf, ret = client.Read(fd540, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd542, 65, client.SEEK_SET)
if(ret != 65) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 65)
  panic("seek failed")
}


fd549 := client.Open("/JJw95P1AZF.txt", client.O_RDWR|client.O_CREATE)
if(fd549 < 0) {
  panic("open failed")
}


ret = client.Close(fd545)
if(ret != 0) {
  panic("close failed")
}

//fd state: (65) gh8ZrcUqP41TT5jUhdQlmP6SfDpCM8GY2TZeTUR4hO8IhS61NFOW1gaYlVv70Hihc9vnly

ret = client.Write(fd542, []byte("a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) gh8ZrcUqP41TT5jUhdQlmP6SfDpCM8GY2TZeTUR4hO8IhS61NFOW1gaYlVv70Hihcavnly
//fd state: (12) Y0ztVwJ1o2F7izzhfxMt1UwhWbNL3F5P6XfIaL1ydJgXy1jai2lEpTAPe0vKGo

ret = client.Write(fd543, []byte("uMN7JomDQ38XN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (25) Y0ztVwJ1o2F7uMN7JomDQ38XNbNL3F5P6XfIaL1ydJgXy1jai2lEpTAPe0vKGo
//fd state: (0) OA7Dw511gPquGZQIjCwsUdugB2H3stNinfhjaG1auZbaVFmFPXwlknFPLTP9XVRifnzdlcTcimGTjpUp4payaskb1Ps7IUSK3it58J0FEZQG41ivpJwzzAHcVkTuLmgUfCEPa4XC2Y

ret = client.Write(fd549, []byte("OOpBuHTW5gjVgu0HYJ00pvjmUbgPvbEhHqujak"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) OOpBuHTW5gjVgu0HYJ00pvjmUbgPvbEhHqujak1auZbaVFmFPXwlknFPLTP9XVRifnzdlcTcimGTjpUp4payaskb1Ps7IUSK3it58J0FEZQG41ivpJwzzAHcVkTuLmgUfCEPa4XC2Y

ret = client.Close(fd539)
if(ret != 0) {
  panic("close failed")
}


fd550 := client.Open("/tX3hPwMq2Q.txt", client.O_RDWR|client.O_CREATE)
if(fd550 < 0) {
  panic("open failed")
}


ret = client.Seek(fd541, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Seek(fd541, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


buf, ret = client.Read(fd543, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bNL3F5P6XfIaL1y") {
  panic("wrong data returned")
}


ret = client.Close(fd542)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd546)
if(ret != 0) {
  panic("close failed")
}


fd551 := client.Open("/M6lcoqPuI6.txt", client.O_RDWR|client.O_CREATE)
if(fd551 < 0) {
  panic("open failed")
}


ret = client.Seek(fd543, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}

//fd state: (38) OOpBuHTW5gjVgu0HYJ00pvjmUbgPvbEhHqujak1auZbaVFmFPXwlknFPLTP9XVRifnzdlcTcimGTjpUp4payaskb1Ps7IUSK3it58J0FEZQG41ivpJwzzAHcVkTuLmgUfCEPa4XC2Y

ret = client.Write(fd549, []byte("wIebFJbeYtpJkw73uB0_DfmvGvHzjSYGS9cdzovcKQs0fOjZBYeEBCP4mJaue5Dkxshfuinkn4zzZvqFnoiyysXC9O9G0g0Gzn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) OOpBuHTW5gjVgu0HYJ00pvjmUbgPvbEhHqujakwIebFJbeYtpJkw73uB0_DfmvGvHzjSYGS9cdzovcKQs0fOjZBYeEBCP4mJaue5Dkxshfuinkn4zzZvqFnoiyysXC9O9G0g0Gzn2Y

ret = client.Close(fd541)
if(ret != 0) {
  panic("close failed")
}

//fd state: (81) YI1WNOtEDR7VqNbS3t6vYBpBkr4fPVFVzOQvnTozxvrVSNdaaRntzq_xjoqLNnZi6mvt1SBWHTl3NW_6H

ret = client.Write(fd540, []byte("ik28crYL4ueI2ozYrtba2gHig0WPScLUhHW1JIwvKeEsugi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (128) YI1WNOtEDR7VqNbS3t6vYBpBkr4fPVFVzOQvnTozxvrVSNdaaRntzq_xjoqLNnZi6mvt1SBWHTl3NW_6Hik28crYL4ueI2ozYrtba2gHig0WPScLUhHW1JIwvKeEsugi

buf, ret = client.Read(fd548, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (28) Y0ztVwJ1o2F7uMN7JomDQ38XNbNL3F5P6XfIaL1ydJgXy1jai2lEpTAPe0vKGo

ret = client.Write(fd543, []byte("i6Ba77WBUV6KUxLS7aNpUmu14x8_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) Y0ztVwJ1o2F7uMN7JomDQ38XNbNLi6Ba77WBUV6KUxLS7aNpUmu14x8_e0vKGo

ret = client.Close(fd550)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd547, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "hPFqeZC_oM7N_tU4TMQaYO2mDcphVFBnsOioUWPgWQQq427JjPY0r") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd543, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "e0vKGo") {
  panic("wrong data returned")
}


ret = client.Close(fd548)
if(ret != 0) {
  panic("close failed")
}

//fd state: (128) YI1WNOtEDR7VqNbS3t6vYBpBkr4fPVFVzOQvnTozxvrVSNdaaRntzq_xjoqLNnZi6mvt1SBWHTl3NW_6Hik28crYL4ueI2ozYrtba2gHig0WPScLUhHW1JIwvKeEsugi

ret = client.Write(fd540, []byte("5KfsOi12YG_qqoeJ8TaMSDvoKAyldI8NAv4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (163) YI1WNOtEDR7VqNbS3t6vYBpBkr4fPVFVzOQvnTozxvrVSNdaaRntzq_xjoqLNnZi6mvt1SBWHTl3NW_6Hik28crYL4ueI2ozYrtba2gHig0WPScLUhHW1JIwvKeEsugi5KfsOi12YG_qqoeJ8TaMSDvoKAyldI8NAv4

fd552 := client.Open("/IgWgfliKnz.txt", client.O_RDWR|client.O_CREATE)
if(fd552 < 0) {
  panic("open failed")
}

//fd state: (163) YI1WNOtEDR7VqNbS3t6vYBpBkr4fPVFVzOQvnTozxvrVSNdaaRntzq_xjoqLNnZi6mvt1SBWHTl3NW_6Hik28crYL4ueI2ozYrtba2gHig0WPScLUhHW1JIwvKeEsugi5KfsOi12YG_qqoeJ8TaMSDvoKAyldI8NAv4

ret = client.Write(fd540, []byte("l7NGgr3aYr0AUQZf9sKaCpm_tQHI4xwKoQK_Fh2Km0Eb7QYQZ0SpipRcJ2VZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (223) YI1WNOtEDR7VqNbS3t6vYBpBkr4fPVFVzOQvnTozxvrVSNdaaRntzq_xjoqLNnZi6mvt1SBWHTl3NW_6Hik28crYL4ueI2ozYrtba2gHig0WPScLUhHW1JIwvKeEsugi5KfsOi12YG_qqoeJ8TaMSDvoKAyldI8NAv4l7NGgr3aYr0AUQZf9sKaCpm_tQHI4xwKoQK_Fh2Km0Eb7QYQZ0SpipRcJ2VZ

ret = client.Seek(fd551, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd549)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd552)
if(ret != 0) {
  panic("close failed")
}


fd553 := client.Open("/ySpIGqjdWE.txt", client.O_RDWR|client.O_CREATE)
if(fd553 < 0) {
  panic("open failed")
}


ret = client.Seek(fd543, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


fd554 := client.Open("/NVHRw1JCF2.txt", client.O_RDWR|client.O_CREATE)
if(fd554 < 0) {
  panic("open failed")
}


fd555 := client.Open("/nzteJmsSyV.txt", client.O_RDWR|client.O_CREATE)
if(fd555 < 0) {
  panic("open failed")
}


ret = client.Close(fd551)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd547, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2U8TPaP2Yjoy2l2ysLiIaNiCT3EX8uEESdPRLCu6XLE") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd554, []byte("eopGgOGYGBlbrWzbDDfyb9dkMOUsQuJxjHDbx0GPIhvs5CgT3dNw1g7t4wh2VgZcWruAjIZW3J48xnqEp0jNq3YgQJIaY6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) eopGgOGYGBlbrWzbDDfyb9dkMOUsQuJxjHDbx0GPIhvs5CgT3dNw1g7t4wh2VgZcWruAjIZW3J48xnqEp0jNq3YgQJIaY6

ret = client.Seek(fd543, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd553, []byte("Q8H4YKYVxxu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) Q8H4YKYVxxu

buf, ret = client.Read(fd540, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd547, 92, client.SEEK_SET)
if(ret != 92) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 92)
  panic("seek failed")
}


ret = client.Seek(fd540, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}

//fd state: (11) Q8H4YKYVxxu

ret = client.Write(fd553, []byte("mrkUntI8LfHKRxTCeJl9X0WcLkRak9U_l850PmgDItgmAHF7AEO5zQ7KSUUYfViABJs28Cbo5mgvXptVcxBviQW8y8Y9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) Q8H4YKYVxxumrkUntI8LfHKRxTCeJl9X0WcLkRak9U_l850PmgDItgmAHF7AEO5zQ7KSUUYfViABJs28Cbo5mgvXptVcxBviQW8y8Y9

ret = client.Seek(fd553, 61, client.SEEK_SET)
if(ret != 61) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 61)
  panic("seek failed")
}

//fd state: (94) eopGgOGYGBlbrWzbDDfyb9dkMOUsQuJxjHDbx0GPIhvs5CgT3dNw1g7t4wh2VgZcWruAjIZW3J48xnqEp0jNq3YgQJIaY6

ret = client.Write(fd554, []byte("hbkR21RxbOxpnglyyPULCUSkg4wqQqRPBf1ISO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) eopGgOGYGBlbrWzbDDfyb9dkMOUsQuJxjHDbx0GPIhvs5CgT3dNw1g7t4wh2VgZcWruAjIZW3J48xnqEp0jNq3YgQJIaY6hbkR21RxbOxpnglyyPULCUSkg4wqQqRPBf1ISO

ret = client.Seek(fd553, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (44) Y0ztVwJ1o2F7uMN7JomDQ38XNbNLi6Ba77WBUV6KUxLS7aNpUmu14x8_e0vKGo

ret = client.Write(fd543, []byte("sWOSkOzkxMHG7qAfrjLctxIjkjRF4IBSmQ32B6b8Jcnlj3NR1ngGKnh9Mxp3hjUx91Ehnb11"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) Y0ztVwJ1o2F7uMN7JomDQ38XNbNLi6Ba77WBUV6KUxLSsWOSkOzkxMHG7qAfrjLctxIjkjRF4IBSmQ32B6b8Jcnlj3NR1ngGKnh9Mxp3hjUx91Ehnb11

ret = client.Close(fd540)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd555, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd555, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd556 := client.Open("/gT4qRLxmrv.txt", client.O_RDWR|client.O_CREATE)
if(fd556 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd553, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "H4YKYVxxumrkUntI8LfHKR") {
  panic("wrong data returned")
}


ret = client.Close(fd554)
if(ret != 0) {
  panic("close failed")
}


fd557 := client.Open("/KIwtxrfoHo.txt", client.O_RDWR|client.O_CREATE)
if(fd557 < 0) {
  panic("open failed")
}


ret = client.Seek(fd553, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


ret = client.Seek(fd543, 68, client.SEEK_SET)
if(ret != 68) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 68)
  panic("seek failed")
}


ret = client.Close(fd555)
if(ret != 0) {
  panic("close failed")
}


fd558 := client.Open("/HuKvzU3AWd.txt", client.O_RDWR|client.O_CREATE)
if(fd558 < 0) {
  panic("open failed")
}


fd559 := client.Open("/zWMKnH6tPr.txt", client.O_RDWR|client.O_CREATE)
if(fd559 < 0) {
  panic("open failed")
}

//fd state: (68) Y0ztVwJ1o2F7uMN7JomDQ38XNbNLi6Ba77WBUV6KUxLSsWOSkOzkxMHG7qAfrjLctxIjkjRF4IBSmQ32B6b8Jcnlj3NR1ngGKnh9Mxp3hjUx91Ehnb11

ret = client.Write(fd543, []byte("iTnzcoFd_aoXpzawL_PmEv7yExwj4oq0OEnuOX0esW8fjGSktsU861rWzvANNCQgTk0X6DOOUzmmFrAPAn32Cxz3dW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (158) Y0ztVwJ1o2F7uMN7JomDQ38XNbNLi6Ba77WBUV6KUxLSsWOSkOzkxMHG7qAfrjLctxIjiTnzcoFd_aoXpzawL_PmEv7yExwj4oq0OEnuOX0esW8fjGSktsU861rWzvANNCQgTk0X6DOOUzmmFrAPAn32Cxz3dW

buf, ret = client.Read(fd547, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "6XL") {
  panic("wrong data returned")
}


fd560 := client.Open("/iSviWTXnNc.txt", client.O_RDWR|client.O_CREATE)
if(fd560 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd547, 90, client.SEEK_SET)
if(ret != 90) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 90)
  panic("seek failed")
}


ret = client.Close(fd556)
if(ret != 0) {
  panic("close failed")
}

//fd state: (81) Q8H4YKYVxxumrkUntI8LfHKRxTCeJl9X0WcLkRak9U_l850PmgDItgmAHF7AEO5zQ7KSUUYfViABJs28Cbo5mgvXptVcxBviQW8y8Y9

ret = client.Write(fd553, []byte("7WQNTBUy4Qn0BT1ecfGpd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) Q8H4YKYVxxumrkUntI8LfHKRxTCeJl9X0WcLkRak9U_l850PmgDItgmAHF7AEO5zQ7KSUUYfViABJs28C7WQNTBUy4Qn0BT1ecfGpd9

ret = client.Seek(fd547, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Close(fd559)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd543, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd558)
if(ret != 0) {
  panic("close failed")
}


fd561 := client.Open("/v2DsGdHGbp.txt", client.O_RDWR|client.O_CREATE)
if(fd561 < 0) {
  panic("open failed")
}

//fd state: (0) mgcyiF_3Z58FyfIAf0O4URnFrwbvKR57D9iRhYgs53G6yaWbj261ixhmbDsvGtBW1JG6qh68M5QagXvh1cYXvAxt23eTzuPCN9govVban5ICA

ret = client.Write(fd560, []byte("KTp7tx2XLiFxhi_EWMRzVKj22ZQwtWBSAYoItH1d0V0mGo66Niu_t0roiOGtTN7z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) KTp7tx2XLiFxhi_EWMRzVKj22ZQwtWBSAYoItH1d0V0mGo66Niu_t0roiOGtTN7z1JG6qh68M5QagXvh1cYXvAxt23eTzuPCN9govVban5ICA

ret = client.Close(fd543)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd561, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd562 := client.Open("/MZdRPu9qh_.txt", client.O_RDWR|client.O_CREATE)
if(fd562 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd553, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9") {
  panic("wrong data returned")
}


fd563 := client.Open("/7J0oS48fTq.txt", client.O_RDWR|client.O_CREATE)
if(fd563 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd562, 0)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd563)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd547, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "QQq427JjPY0r2U8TPaP2Yjoy2l2ys") {
  panic("wrong data returned")
}


ret = client.Close(fd561)
if(ret != 0) {
  panic("close failed")
}

//fd state: (64) KTp7tx2XLiFxhi_EWMRzVKj22ZQwtWBSAYoItH1d0V0mGo66Niu_t0roiOGtTN7z1JG6qh68M5QagXvh1cYXvAxt23eTzuPCN9govVban5ICA

ret = client.Write(fd560, []byte("8d0CpVXqjkOY_s20G_bLRP_0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) KTp7tx2XLiFxhi_EWMRzVKj22ZQwtWBSAYoItH1d0V0mGo66Niu_t0roiOGtTN7z8d0CpVXqjkOY_s20G_bLRP_023eTzuPCN9govVban5ICA
//fd state: (0) 

ret = client.Write(fd562, []byte("K2SS3RB5bhsYlbhWrKcsMsRCXKflgBQ00sWt2VPMwiz127ohqw5DHJl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) K2SS3RB5bhsYlbhWrKcsMsRCXKflgBQ00sWt2VPMwiz127ohqw5DHJl

buf, ret = client.Read(fd547, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LiIaNiCT3EX8uEESdPRLCu6XLEEcDctbwFU0DwGusn4mHqNejskKR0") {
  panic("wrong data returned")
}


ret = client.Seek(fd560, 106, client.SEEK_SET)
if(ret != 106) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 106)
  panic("seek failed")
}


buf, ret = client.Read(fd562, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd557, 14)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "j0V_LTkwJdpZsP") {
  panic("wrong data returned")
}


ret = client.Seek(fd560, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


ret = client.Seek(fd562, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd553)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd560, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_t0roiOGtTN7z8d0CpVXqjkOY_s20G_bLRP_023eTzuPCN9govVban5ICA") {
  panic("wrong data returned")
}


ret = client.Close(fd547)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd557, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "kqgpW6B1FugXUFItOvhCT2thZwV0evS4rAddcq54KlYrwpgIvJkgxr1IqLnS34Fqcl1bj4") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd562, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5bhs") {
  panic("wrong data returned")
}


ret = client.Seek(fd560, 109, client.SEEK_SET)
if(ret != 109) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 109)
  panic("seek failed")
}

//fd state: (11) K2SS3RB5bhsYlbhWrKcsMsRCXKflgBQ00sWt2VPMwiz127ohqw5DHJl

ret = client.Write(fd562, []byte("kkWX0LTlB4Bal3pnNuGmd6Nb6eD0cDfJNn8Ly1eNrpdlhFUziRJRsMbHUMlm64k8wviym6sC3VaZnj7lHAPLm3Sdu7rN7OyWD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (108) K2SS3RB5bhskkWX0LTlB4Bal3pnNuGmd6Nb6eD0cDfJNn8Ly1eNrpdlhFUziRJRsMbHUMlm64k8wviym6sC3VaZnj7lHAPLm3Sdu7rN7OyWD

fd564 := client.Open("/Do8eCB4JFH.txt", client.O_RDWR|client.O_CREATE)
if(fd564 < 0) {
  panic("open failed")
}


fd565 := client.Open("/EtjzgRe_ul.txt", client.O_RDWR|client.O_CREATE)
if(fd565 < 0) {
  panic("open failed")
}


fd566 := client.Open("/an0N4p7hQP.txt", client.O_RDWR|client.O_CREATE)
if(fd566 < 0) {
  panic("open failed")
}

//fd state: (109) KTp7tx2XLiFxhi_EWMRzVKj22ZQwtWBSAYoItH1d0V0mGo66Niu_t0roiOGtTN7z8d0CpVXqjkOY_s20G_bLRP_023eTzuPCN9govVban5ICA

ret = client.Write(fd560, []byte("IKTQHtoELlq2YF7f6fNi5XF8i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) KTp7tx2XLiFxhi_EWMRzVKj22ZQwtWBSAYoItH1d0V0mGo66Niu_t0roiOGtTN7z8d0CpVXqjkOY_s20G_bLRP_023eTzuPCN9govVban5ICAIKTQHtoELlq2YF7f6fNi5XF8i

ret = client.Seek(fd557, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}

//fd state: (108) K2SS3RB5bhskkWX0LTlB4Bal3pnNuGmd6Nb6eD0cDfJNn8Ly1eNrpdlhFUziRJRsMbHUMlm64k8wviym6sC3VaZnj7lHAPLm3Sdu7rN7OyWD

ret = client.Write(fd562, []byte("1IEcnnFNAKtt56yHC3DC2B9juzfbEvIdNpa1C5STHmNMMy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (154) K2SS3RB5bhskkWX0LTlB4Bal3pnNuGmd6Nb6eD0cDfJNn8Ly1eNrpdlhFUziRJRsMbHUMlm64k8wviym6sC3VaZnj7lHAPLm3Sdu7rN7OyWD1IEcnnFNAKtt56yHC3DC2B9juzfbEvIdNpa1C5STHmNMMy

ret = client.Seek(fd564, 69, client.SEEK_SET)
if(ret != 69) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 69)
  panic("seek failed")
}


ret = client.Close(fd560)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd562, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}


ret = client.Close(fd566)
if(ret != 0) {
  panic("close failed")
}


fd567 := client.Open("/FPPNgeMK1l.txt", client.O_RDWR|client.O_CREATE)
if(fd567 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd565, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd568 := client.Open("/k57M1a0Y2S.txt", client.O_RDWR|client.O_CREATE)
if(fd568 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd568, []byte("uIWRBUTVmSLW6MJWvcmVbPT4sWhRVdKw7DWUJa2jJCNMrXQ6pLO7uQSsmCiMpNMH8JwgyhgGSbXMRpxDwLpFxm8xmTbtd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) uIWRBUTVmSLW6MJWvcmVbPT4sWhRVdKw7DWUJa2jJCNMrXQ6pLO7uQSsmCiMpNMH8JwgyhgGSbXMRpxDwLpFxm8xmTbtd

fd569 := client.Open("/HuKvzU3AWd.txt", client.O_RDWR|client.O_CREATE)
if(fd569 < 0) {
  panic("open failed")
}


ret = client.Close(fd562)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd568, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


buf, ret = client.Read(fd565, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd570 := client.Open("/3dMK9W1zC2.txt", client.O_RDWR|client.O_CREATE)
if(fd570 < 0) {
  panic("open failed")
}


ret = client.Seek(fd567, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd571 := client.Open("/4miTNfxwtx.txt", client.O_RDWR|client.O_CREATE)
if(fd571 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd564, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "mSv_3xh9muk4_1gE7JLGUF9Qtx") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd571, []byte("yXpMX5fXOU3yYjx_DhdLrXGPBWH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) yXpMX5fXOU3yYjx_DhdLrXGPBWH
//fd state: (27) yXpMX5fXOU3yYjx_DhdLrXGPBWH

ret = client.Write(fd571, []byte("eCSis3D44m_Zc2fSIv__Oy_qn9FVo_SEytU6yuUOz9ToL6tf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) yXpMX5fXOU3yYjx_DhdLrXGPBWHeCSis3D44m_Zc2fSIv__Oy_qn9FVo_SEytU6yuUOz9ToL6tf

ret = client.Seek(fd565, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd564, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


fd572 := client.Open("/4Zu47k6Y33.txt", client.O_RDWR|client.O_CREATE)
if(fd572 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd567, []byte("Ph"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) Ph

ret = client.Seek(fd570, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) E_EwwJocFAGCQqqWDkasFYseDJafudrUghYmeem_R7qdNolV2Nv0OjoQGg9POekC27mqvc3f7gVaAvgPbIqyv4Iab94FcQ2b0ZWRMHny6xV4tAsBZU3ULoDaN9IDlgH3D9CD2nloBM8qBSHCj

ret = client.Write(fd569, []byte("I0dLO0ly_moZBdXEMG1jTUYO9lPlXWP1kybLrUGyeB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) I0dLO0ly_moZBdXEMG1jTUYO9lPlXWP1kybLrUGyeBqdNolV2Nv0OjoQGg9POekC27mqvc3f7gVaAvgPbIqyv4Iab94FcQ2b0ZWRMHny6xV4tAsBZU3ULoDaN9IDlgH3D9CD2nloBM8qBSHCj

ret = client.Seek(fd565, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd573 := client.Open("/HdAlR85KzA.txt", client.O_RDWR|client.O_CREATE)
if(fd573 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd564, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "RBRnGKzd0RrigjvXEzOuWqrzLj") {
  panic("wrong data returned")
}


ret = client.Seek(fd570, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd564, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Hrd5Uq3mSv_3xh9muk4_1gE7JLGUF9Qtx") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd573, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd573, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (42) I0dLO0ly_moZBdXEMG1jTUYO9lPlXWP1kybLrUGyeBqdNolV2Nv0OjoQGg9POekC27mqvc3f7gVaAvgPbIqyv4Iab94FcQ2b0ZWRMHny6xV4tAsBZU3ULoDaN9IDlgH3D9CD2nloBM8qBSHCj

ret = client.Write(fd569, []byte("p3P0ThwG0BGyosN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) I0dLO0ly_moZBdXEMG1jTUYO9lPlXWP1kybLrUGyeBp3P0ThwG0BGyosNg9POekC27mqvc3f7gVaAvgPbIqyv4Iab94FcQ2b0ZWRMHny6xV4tAsBZU3ULoDaN9IDlgH3D9CD2nloBM8qBSHCj

buf, ret = client.Read(fd569, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "g9POekC27mqvc3f7gVaAvgPbIqyv4Iab94FcQ") {
  panic("wrong data returned")
}


ret = client.Close(fd568)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd565, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd572, []byte("xEJxL8yi0WBsxtmISM8nCJIP8wbqTYkFGye6oHsX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) xEJxL8yi0WBsxtmISM8nCJIP8wbqTYkFGye6oHsX

ret = client.Seek(fd570, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd569)
if(ret != 0) {
  panic("close failed")
}


fd574 := client.Open("/wyH_1ywsi5.txt", client.O_RDWR|client.O_CREATE)
if(fd574 < 0) {
  panic("open failed")
}

//fd state: (95) 0jRvYQj7URR5zsbLHJzOmVteqCW6CppHJm3iRBRnGKzd0RrigjvXEzOuWqrzLjHrd5Uq3mSv_3xh9muk4_1gE7JLGUF9Qtx

ret = client.Write(fd564, []byte("Ni5WMn8OZpqn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) 0jRvYQj7URR5zsbLHJzOmVteqCW6CppHJm3iRBRnGKzd0RrigjvXEzOuWqrzLjHrd5Uq3mSv_3xh9muk4_1gE7JLGUF9QtxNi5WMn8OZpqn
//fd state: (0) 

ret = client.Write(fd570, []byte("JFpJuwiTKfjp4Q0qFfqAFCQN84Fo00HKjbaDydHSS9e0be77kB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) JFpJuwiTKfjp4Q0qFfqAFCQN84Fo00HKjbaDydHSS9e0be77kB

ret = client.Close(fd571)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd570)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd567, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd572, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}

//fd state: (107) 0jRvYQj7URR5zsbLHJzOmVteqCW6CppHJm3iRBRnGKzd0RrigjvXEzOuWqrzLjHrd5Uq3mSv_3xh9muk4_1gE7JLGUF9QtxNi5WMn8OZpqn

ret = client.Write(fd564, []byte("qTFyeHuIKUwh3Gip3iWAXgQ4rbnQSks9P"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (140) 0jRvYQj7URR5zsbLHJzOmVteqCW6CppHJm3iRBRnGKzd0RrigjvXEzOuWqrzLjHrd5Uq3mSv_3xh9muk4_1gE7JLGUF9QtxNi5WMn8OZpqnqTFyeHuIKUwh3Gip3iWAXgQ4rbnQSks9P

ret = client.Close(fd572)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd573, []byte("3mKj529b5aflApHDcteqnaMRmmzJxR_6Ku9fGpg3W8nWRjPljJus38dWCtwsnWj9I6aVP8N_hhIxAs_9jv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) 3mKj529b5aflApHDcteqnaMRmmzJxR_6Ku9fGpg3W8nWRjPljJus38dWCtwsnWj9I6aVP8N_hhIxAs_9jv
//fd state: (82) 3mKj529b5aflApHDcteqnaMRmmzJxR_6Ku9fGpg3W8nWRjPljJus38dWCtwsnWj9I6aVP8N_hhIxAs_9jv

ret = client.Write(fd573, []byte("lLb1Xs3chRen1XV1WJcFSqmNVzcAnCN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) 3mKj529b5aflApHDcteqnaMRmmzJxR_6Ku9fGpg3W8nWRjPljJus38dWCtwsnWj9I6aVP8N_hhIxAs_9jvlLb1Xs3chRen1XV1WJcFSqmNVzcAnCN

ret = client.Close(fd574)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd567, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd564)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd557)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd573)
if(ret != 0) {
  panic("close failed")
}


fd575 := client.Open("/iMX4DbJvto.txt", client.O_RDWR|client.O_CREATE)
if(fd575 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd565, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd565, []byte("UJrRcn1G35PjycUKIzPe6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) UJrRcn1G35PjycUKIzPe6

fd576 := client.Open("/ALsGXybpGN.txt", client.O_RDWR|client.O_CREATE)
if(fd576 < 0) {
  panic("open failed")
}


ret = client.Close(fd576)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd565)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) Ph

ret = client.Write(fd567, []byte("JnAWv15UZZUFYU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) JnAWv15UZZUFYU
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd577 := client.Open("/J5KDyCcaR4.txt", client.O_RDWR|client.O_CREATE)
if(fd577 < 0) {
  panic("open failed")
}


ret = client.Close(fd575)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd567, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


fd578 := client.Open("/BTpMjqSUbn.txt", client.O_RDWR|client.O_CREATE)
if(fd578 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd567, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ZZUFYU") {
  panic("wrong data returned")
}


ret = client.Close(fd577)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd567, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd578, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd579 := client.Open("/aHSi5Rp9bT.txt", client.O_RDWR|client.O_CREATE)
if(fd579 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd567, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd579, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd567, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (14) JnAWv15UZZUFYU

ret = client.Write(fd567, []byte("Dj6pJ_a5zcWfBbsHuv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) JnAWv15UZZUFYUDj6pJ_a5zcWfBbsHuv
//fd state: (0) 

ret = client.Write(fd578, []byte("tOM9ePfR8aonmh3Zawrvhis9i_u8qMUn2p7lpZOBQQTELhJZopA8owvCFazYoMNVS_p8jZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) tOM9ePfR8aonmh3Zawrvhis9i_u8qMUn2p7lpZOBQQTELhJZopA8owvCFazYoMNVS_p8jZ

fd580 := client.Open("/KTfoUTGtqF.txt", client.O_RDWR|client.O_CREATE)
if(fd580 < 0) {
  panic("open failed")
}

//fd state: (32) JnAWv15UZZUFYUDj6pJ_a5zcWfBbsHuv

ret = client.Write(fd567, []byte("UlL5HDUf4Urh4VEEHi7CcCXIlA5NZXy9rBEMVXfquSldit_uzMS_o0vSllsY7fi8IFEDy4LzDtxMrv9JPNiWy20wI_QT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) JnAWv15UZZUFYUDj6pJ_a5zcWfBbsHuvUlL5HDUf4Urh4VEEHi7CcCXIlA5NZXy9rBEMVXfquSldit_uzMS_o0vSllsY7fi8IFEDy4LzDtxMrv9JPNiWy20wI_QT

ret = client.Close(fd579)
if(ret != 0) {
  panic("close failed")
}


fd581 := client.Open("/B3ip_bKpPr.txt", client.O_RDWR|client.O_CREATE)
if(fd581 < 0) {
  panic("open failed")
}


fd582 := client.Open("/nVFk0Llsip.txt", client.O_RDWR|client.O_CREATE)
if(fd582 < 0) {
  panic("open failed")
}


ret = client.Seek(fd580, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd583 := client.Open("/SjlUEoWJwN.txt", client.O_RDWR|client.O_CREATE)
if(fd583 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd580, []byte("ygeI8A0VFaRrCWgXpuaDehFfZtlsbA46iTbzD9WTwCHpzIGxpyo2uUbpYWopNAoI03XRoM_JLuZlM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) ygeI8A0VFaRrCWgXpuaDehFfZtlsbA46iTbzD9WTwCHpzIGxpyo2uUbpYWopNAoI03XRoM_JLuZlM

ret = client.Close(fd567)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) bAPtH9bEtsPwxMV7Buc4AAkfKF3lWlnD24YQ_9SER3Z6L7sgBumqAfxDhFn8EldI2mQuFRD3rL9_EuhSjZ5j0mPHjNoV

ret = client.Write(fd582, []byte("oOr9w_cQASB4HiIWo5UjJM2Npbmq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) oOr9w_cQASB4HiIWo5UjJM2NpbmqWlnD24YQ_9SER3Z6L7sgBumqAfxDhFn8EldI2mQuFRD3rL9_EuhSjZ5j0mPHjNoV

ret = client.Close(fd581)
if(ret != 0) {
  panic("close failed")
}

//fd state: (28) oOr9w_cQASB4HiIWo5UjJM2NpbmqWlnD24YQ_9SER3Z6L7sgBumqAfxDhFn8EldI2mQuFRD3rL9_EuhSjZ5j0mPHjNoV

ret = client.Write(fd582, []byte("sFAIZRG2efq83vPhbEKlMLtAIdkRCxG6Z41jvRd_k1Pr59utMX7qVrBT8j"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) oOr9w_cQASB4HiIWo5UjJM2NpbmqsFAIZRG2efq83vPhbEKlMLtAIdkRCxG6Z41jvRd_k1Pr59utMX7qVrBT8jPHjNoV

buf, ret = client.Read(fd582, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PHjN") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd578, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd582)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd580, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd584 := client.Open("/N_tIxiNaQn.txt", client.O_RDWR|client.O_CREATE)
if(fd584 < 0) {
  panic("open failed")
}

//fd state: (70) tOM9ePfR8aonmh3Zawrvhis9i_u8qMUn2p7lpZOBQQTELhJZopA8owvCFazYoMNVS_p8jZ

ret = client.Write(fd578, []byte("vHJKbBDCYMRlDS3xpV8ZBiZY7Om__ysKFjsBZdZ0HlEh1PWsrA3Z1b82DrXKmRIWQDfHSTA0NZOnmLrc8wBGkJlI1X9_wAP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (165) tOM9ePfR8aonmh3Zawrvhis9i_u8qMUn2p7lpZOBQQTELhJZopA8owvCFazYoMNVS_p8jZvHJKbBDCYMRlDS3xpV8ZBiZY7Om__ysKFjsBZdZ0HlEh1PWsrA3Z1b82DrXKmRIWQDfHSTA0NZOnmLrc8wBGkJlI1X9_wAP

ret = client.Seek(fd578, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}

//fd state: (77) ygeI8A0VFaRrCWgXpuaDehFfZtlsbA46iTbzD9WTwCHpzIGxpyo2uUbpYWopNAoI03XRoM_JLuZlM

ret = client.Write(fd580, []byte("AolWvtnbe1gHwWqhTwqbMe9leEMHCYI_qHE4cJ3qHe86c9OJ1YgQIuSHSo1exWObXjOmXCwC4iR2mXmOE6hyDFEcZYTUNwyX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (173) ygeI8A0VFaRrCWgXpuaDehFfZtlsbA46iTbzD9WTwCHpzIGxpyo2uUbpYWopNAoI03XRoM_JLuZlMAolWvtnbe1gHwWqhTwqbMe9leEMHCYI_qHE4cJ3qHe86c9OJ1YgQIuSHSo1exWObXjOmXCwC4iR2mXmOE6hyDFEcZYTUNwyX

ret = client.Close(fd578)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd583, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd585 := client.Open("/ctwAVymDm0.txt", client.O_RDWR|client.O_CREATE)
if(fd585 < 0) {
  panic("open failed")
}


ret = client.Close(fd580)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd585)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) rffsCahgymr_Mgi1z1NdDwUxLLkP0IrhtanP9Sam1W31LwXpwMlVmcMXY2fHt8PkqxHJebpPpEHdG157yiCs6V7A7aBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

ret = client.Write(fd584, []byte("FBJDy8taNpUxqOKLxeIfKLetE2S9NxXIFeAXI7J7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) FBJDy8taNpUxqOKLxeIfKLetE2S9NxXIFeAXI7J71W31LwXpwMlVmcMXY2fHt8PkqxHJebpPpEHdG157yiCs6V7A7aBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

buf, ret = client.Read(fd583, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (40) FBJDy8taNpUxqOKLxeIfKLetE2S9NxXIFeAXI7J71W31LwXpwMlVmcMXY2fHt8PkqxHJebpPpEHdG157yiCs6V7A7aBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

ret = client.Write(fd584, []byte("poJE6De_lg0l1Zyd4BIY2c4IEPMU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) FBJDy8taNpUxqOKLxeIfKLetE2S9NxXIFeAXI7J7poJE6De_lg0l1Zyd4BIY2c4IEPMUebpPpEHdG157yiCs6V7A7aBY2knYQqfqYyibRbe8JfN2eyDY3SUpsrBTKu9WLF9BrN3RyOmUbqVElfH0SI

ret = client.Seek(fd583, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd586 := client.Open("/2j75mwM1GP.txt", client.O_RDWR|client.O_CREATE)
if(fd586 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd584, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ebpPpEHdG157yiCs6V7A7aBY2knYQqfqYyibRbe8JfN") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd586, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd584)
if(ret != 0) {
  panic("close failed")
}


fd587 := client.Open("/kbPa5OvfPV.txt", client.O_RDWR|client.O_CREATE)
if(fd587 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd587, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd583, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd586, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd583)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd586, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd588 := client.Open("/xO3DD4QTKO.txt", client.O_RDWR|client.O_CREATE)
if(fd588 < 0) {
  panic("open failed")
}


ret = client.Close(fd586)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd587)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd588, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd588, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd588, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd588, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd588)
if(ret != 0) {
  panic("close failed")
}


fd589 := client.Open("/XRA36JjMkl.txt", client.O_RDWR|client.O_CREATE)
if(fd589 < 0) {
  panic("open failed")
}


ret = client.Seek(fd589, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}

//fd state: (20) EqYoao97jeTa8jLQYPg3BnN6uy8R9pFQ4b5a1dUokeB9YALycTfthge2VpmnxlSeA

ret = client.Write(fd589, []byte("4Q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) EqYoao97jeTa8jLQYPg34QN6uy8R9pFQ4b5a1dUokeB9YALycTfthge2VpmnxlSeA

fd590 := client.Open("/UpFwxaZwrM.txt", client.O_RDWR|client.O_CREATE)
if(fd590 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd590, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd590, []byte("R1QLnUH8iYE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) R1QLnUH8iYE
//fd state: (11) R1QLnUH8iYE

ret = client.Write(fd590, []byte("0lOVyiRMVFUdQEHat8FgBXTgPEMqJt7nLhMvxSJc4HRRNaZj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) R1QLnUH8iYE0lOVyiRMVFUdQEHat8FgBXTgPEMqJt7nLhMvxSJc4HRRNaZj

ret = client.Seek(fd590, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


ret = client.Close(fd589)
if(ret != 0) {
  panic("close failed")
}


fd591 := client.Open("/tl7XYCDeEa.txt", client.O_RDWR|client.O_CREATE)
if(fd591 < 0) {
  panic("open failed")
}

//fd state: (18) R1QLnUH8iYE0lOVyiRMVFUdQEHat8FgBXTgPEMqJt7nLhMvxSJc4HRRNaZj

ret = client.Write(fd590, []byte("dOKGJPvY6F81vxOGAmvbohNb3XVxv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) R1QLnUH8iYE0lOVyiRdOKGJPvY6F81vxOGAmvbohNb3XVxvxSJc4HRRNaZj

fd592 := client.Open("/ZHmEuATXev.txt", client.O_RDWR|client.O_CREATE)
if(fd592 < 0) {
  panic("open failed")
}

//fd state: (47) R1QLnUH8iYE0lOVyiRdOKGJPvY6F81vxOGAmvbohNb3XVxvxSJc4HRRNaZj

ret = client.Write(fd590, []byte("lchdlTBIXgbdnBZR7lTOAyiHHA4Mm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) R1QLnUH8iYE0lOVyiRdOKGJPvY6F81vxOGAmvbohNb3XVxvlchdlTBIXgbdnBZR7lTOAyiHHA4Mm

buf, ret = client.Read(fd591, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd593 := client.Open("/hJfLk46nNx.txt", client.O_RDWR|client.O_CREATE)
if(fd593 < 0) {
  panic("open failed")
}

//fd state: (76) R1QLnUH8iYE0lOVyiRdOKGJPvY6F81vxOGAmvbohNb3XVxvlchdlTBIXgbdnBZR7lTOAyiHHA4Mm

ret = client.Write(fd590, []byte("4pwqncWq7dS6j6WzHRAxzW13xaDEbSyWrJrg47QBd43B7HpI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (124) R1QLnUH8iYE0lOVyiRdOKGJPvY6F81vxOGAmvbohNb3XVxvlchdlTBIXgbdnBZR7lTOAyiHHA4Mm4pwqncWq7dS6j6WzHRAxzW13xaDEbSyWrJrg47QBd43B7HpI

buf, ret = client.Read(fd592, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd590, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


ret = client.Close(fd590)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd593, []byte("qXJfH9KecWPjOtiS5W_YgNhSnoVSS_yezT29IbOYcw2wflvtRQHgGNyqVctky8Mb7PxDjBNNQhMQRWQhI1ZIWB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) qXJfH9KecWPjOtiS5W_YgNhSnoVSS_yezT29IbOYcw2wflvtRQHgGNyqVctky8Mb7PxDjBNNQhMQRWQhI1ZIWB

ret = client.Close(fd593)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd592, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd592, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd591)
if(ret != 0) {
  panic("close failed")
}


fd594 := client.Open("/8ldA_h4XX0.txt", client.O_RDWR|client.O_CREATE)
if(fd594 < 0) {
  panic("open failed")
}


ret = client.Seek(fd592, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd592, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd594)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd592, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd595 := client.Open("/0bepIOJ5hl.txt", client.O_RDWR|client.O_CREATE)
if(fd595 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd595, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "e7GtGo9gJmFzHhRzcGsEGuOwNIvc0Rg") {
  panic("wrong data returned")
}


fd596 := client.Open("/DRIJ80NCqo.txt", client.O_RDWR|client.O_CREATE)
if(fd596 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd592, []byte("N1qBm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) N1qBm

fd597 := client.Open("/t_LsoQNnzK.txt", client.O_RDWR|client.O_CREATE)
if(fd597 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd592, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd596, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


buf, ret = client.Read(fd595, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "eSaszv3ZkMhEV4X8cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg9xXbYIyTnXeA5iyMZRt2cAu6abhMrV5Ar_zNZLrxFMb8ER") {
  panic("wrong data returned")
}


ret = client.Seek(fd595, 86, client.SEEK_SET)
if(ret != 86) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 86)
  panic("seek failed")
}


fd598 := client.Open("/3SAmTKFcdh.txt", client.O_RDWR|client.O_CREATE)
if(fd598 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (86) e7GtGo9gJmFzHhRzcGsEGuOwNIvc0RgeSaszv3ZkMhEV4X8cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg9xXbYIyTnXeA5iyMZRt2cAu6abhMrV5Ar_zNZLrxFMb8ERTtRP5V1614gAXD26772YV8TDHh8YJE3pOxCxpdqgnklS6QEDjVH8810Oy

ret = client.Write(fd595, []byte("lEsoENN5nmXqn3pnnYjNtpnTVp_xjgUbzK0AaJ9tUJTefwrVi0DJh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (139) e7GtGo9gJmFzHhRzcGsEGuOwNIvc0RgeSaszv3ZkMhEV4X8cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg9xXbYIylEsoENN5nmXqn3pnnYjNtpnTVp_xjgUbzK0AaJ9tUJTefwrVi0DJh26772YV8TDHh8YJE3pOxCxpdqgnklS6QEDjVH8810Oy

ret = client.Seek(fd596, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


buf, ret = client.Read(fd597, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd596)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd592, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd597, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd597, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd592)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd598, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd595, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "26772YV8TDHh8YJE3pOxCxpdqgnklS6QEDjVH8810Oy") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd597, []byte("gu1j7neEQawYN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) gu1j7neEQawYN

buf, ret = client.Read(fd598, 23)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd595, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


ret = client.Seek(fd597, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}

//fd state: (37) e7GtGo9gJmFzHhRzcGsEGuOwNIvc0RgeSaszv3ZkMhEV4X8cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg9xXbYIylEsoENN5nmXqn3pnnYjNtpnTVp_xjgUbzK0AaJ9tUJTefwrVi0DJh26772YV8TDHh8YJE3pOxCxpdqgnklS6QEDjVH8810Oy

ret = client.Write(fd595, []byte("640rmVgmp1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) e7GtGo9gJmFzHhRzcGsEGuOwNIvc0RgeSaszv640rmVgmp1cs3V7W8A8qu5roJQ6_68mnwe8VSx3nTg9xXbYIylEsoENN5nmXqn3pnnYjNtpnTVp_xjgUbzK0AaJ9tUJTefwrVi0DJh26772YV8TDHh8YJE3pOxCxpdqgnklS6QEDjVH8810Oy

ret = client.Seek(fd598, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd599 := client.Open("/koC3k4dnym.txt", client.O_RDWR|client.O_CREATE)
if(fd599 < 0) {
  panic("open failed")
}

//fd state: (13) gu1j7neEQawYN

ret = client.Write(fd597, []byte("naBkN4rBPgC1tH_dI45l7503nE9u5hKpr86oW6Tqhf_Rq8IYlrn2XhGYpU0dyJ6RRkbXYR3PbwN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) gu1j7neEQawYNnaBkN4rBPgC1tH_dI45l7503nE9u5hKpr86oW6Tqhf_Rq8IYlrn2XhGYpU0dyJ6RRkbXYR3PbwN

ret = client.Close(fd595)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd598, []byte("l7cmGwYMTwfXdKksOUVc4a6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMea"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) l7cmGwYMTwfXdKksOUVc4a6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMea

ret = client.Seek(fd597, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}

//fd state: (87) l7cmGwYMTwfXdKksOUVc4a6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMea

ret = client.Write(fd598, []byte("cGtvujY8us5n3eiQhIzWdDWlE1CaKRQCfU0tp3mt6Cc9uJjk6KBOfDS5X03Zl4RlfetEX1TESAONNgKL7wn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (170) l7cmGwYMTwfXdKksOUVc4a6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujY8us5n3eiQhIzWdDWlE1CaKRQCfU0tp3mt6Cc9uJjk6KBOfDS5X03Zl4RlfetEX1TESAONNgKL7wn

buf, ret = client.Read(fd599, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "tdraiIa") {
  panic("wrong data returned")
}


ret = client.Seek(fd598, 6, client.SEEK_SET)
if(ret != 6) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 6)
  panic("seek failed")
}

//fd state: (6) l7cmGwYMTwfXdKksOUVc4a6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujY8us5n3eiQhIzWdDWlE1CaKRQCfU0tp3mt6Cc9uJjk6KBOfDS5X03Zl4RlfetEX1TESAONNgKL7wn

ret = client.Write(fd598, []byte("1kIP2x4CUtWwOQMY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujY8us5n3eiQhIzWdDWlE1CaKRQCfU0tp3mt6Cc9uJjk6KBOfDS5X03Zl4RlfetEX1TESAONNgKL7wn

buf, ret = client.Read(fd598, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujY") {
  panic("wrong data returned")
}


ret = client.Close(fd599)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd597, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "gC1tH_dI45l7503nE9u5hKpr86oW6Tqhf_Rq8IYlrn2XhGYpU0dyJ6RRkbXYR3P") {
  panic("wrong data returned")
}

//fd state: (94) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujY8us5n3eiQhIzWdDWlE1CaKRQCfU0tp3mt6Cc9uJjk6KBOfDS5X03Zl4RlfetEX1TESAONNgKL7wn

ret = client.Write(fd598, []byte("qgZI7hu7kS9s8GjvKh4Dwu6MyVu307wFgCaH4clJFr9kpStgICfQjzP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (149) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujYqgZI7hu7kS9s8GjvKh4Dwu6MyVu307wFgCaH4clJFr9kpStgICfQjzPRlfetEX1TESAONNgKL7wn
//fd state: (85) gu1j7neEQawYNnaBkN4rBPgC1tH_dI45l7503nE9u5hKpr86oW6Tqhf_Rq8IYlrn2XhGYpU0dyJ6RRkbXYR3PbwN

ret = client.Write(fd597, []byte("n6Gtdd5DYtsGayfx257zyi9lLoiW9S94bKx2c2p9GO6f59R2V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) gu1j7neEQawYNnaBkN4rBPgC1tH_dI45l7503nE9u5hKpr86oW6Tqhf_Rq8IYlrn2XhGYpU0dyJ6RRkbXYR3Pn6Gtdd5DYtsGayfx257zyi9lLoiW9S94bKx2c2p9GO6f59R2V

fd600 := client.Open("/ZJ6X2yc2mq.txt", client.O_RDWR|client.O_CREATE)
if(fd600 < 0) {
  panic("open failed")
}


fd601 := client.Open("/K4uEAu5Fwz.txt", client.O_RDWR|client.O_CREATE)
if(fd601 < 0) {
  panic("open failed")
}


ret = client.Close(fd597)
if(ret != 0) {
  panic("close failed")
}


fd602 := client.Open("/JsEWZvUNf1.txt", client.O_RDWR|client.O_CREATE)
if(fd602 < 0) {
  panic("open failed")
}


ret = client.Seek(fd600, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd598, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "R") {
  panic("wrong data returned")
}


ret = client.Seek(fd598, 87, client.SEEK_SET)
if(ret != 87) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 87)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd601, []byte("Cc3vVfiMFscKTleAqBpZx1dP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (24) Cc3vVfiMFscKTleAqBpZx1dP

ret = client.Close(fd600)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd601, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


ret = client.Seek(fd598, 90, client.SEEK_SET)
if(ret != 90) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 90)
  panic("seek failed")
}


ret = client.Seek(fd598, 136, client.SEEK_SET)
if(ret != 136) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 136)
  panic("seek failed")
}


ret = client.Seek(fd598, 83, client.SEEK_SET)
if(ret != 83) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 83)
  panic("seek failed")
}


buf, ret = client.Read(fd598, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oMeacGtvujYqgZI7hu7kS9s8GjvKh4Dwu6MyVu307wFgCaH4clJFr9kp") {
  panic("wrong data returned")
}


ret = client.Close(fd602)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd601, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Close(fd601)
if(ret != 0) {
  panic("close failed")
}

//fd state: (139) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujYqgZI7hu7kS9s8GjvKh4Dwu6MyVu307wFgCaH4clJFr9kpStgICfQjzPRlfetEX1TESAONNgKL7wn

ret = client.Write(fd598, []byte("7uHRzPCUjZOBK0gVjO_Z4Y98s7a0tnvD8HnQGAmb3qxlcQ4gSB4EJV3YwAXsw8MF58fnYQLTn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (212) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujYqgZI7hu7kS9s8GjvKh4Dwu6MyVu307wFgCaH4clJFr9kp7uHRzPCUjZOBK0gVjO_Z4Y98s7a0tnvD8HnQGAmb3qxlcQ4gSB4EJV3YwAXsw8MF58fnYQLTn

ret = client.Seek(fd598, 99, client.SEEK_SET)
if(ret != 99) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 99)
  panic("seek failed")
}


ret = client.Seek(fd598, 111, client.SEEK_SET)
if(ret != 111) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 111)
  panic("seek failed")
}

//fd state: (111) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujYqgZI7hu7kS9s8GjvKh4Dwu6MyVu307wFgCaH4clJFr9kp7uHRzPCUjZOBK0gVjO_Z4Y98s7a0tnvD8HnQGAmb3qxlcQ4gSB4EJV3YwAXsw8MF58fnYQLTn

ret = client.Write(fd598, []byte("jI2wKh8pIjSTv4h_9czwwdnkGDPpZ8ZN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (143) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujYqgZI7hu7kS9s8GjvKjI2wKh8pIjSTv4h_9czwwdnkGDPpZ8ZNzPCUjZOBK0gVjO_Z4Y98s7a0tnvD8HnQGAmb3qxlcQ4gSB4EJV3YwAXsw8MF58fnYQLTn
//fd state: (143) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujYqgZI7hu7kS9s8GjvKjI2wKh8pIjSTv4h_9czwwdnkGDPpZ8ZNzPCUjZOBK0gVjO_Z4Y98s7a0tnvD8HnQGAmb3qxlcQ4gSB4EJV3YwAXsw8MF58fnYQLTn

ret = client.Write(fd598, []byte("5U5ROixd67r9V4sBqyuuTw"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (165) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujYqgZI7hu7kS9s8GjvKjI2wKh8pIjSTv4h_9czwwdnkGDPpZ8ZN5U5ROixd67r9V4sBqyuuTwa0tnvD8HnQGAmb3qxlcQ4gSB4EJV3YwAXsw8MF58fnYQLTn

ret = client.Seek(fd598, 152, client.SEEK_SET)
if(ret != 152) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 152)
  panic("seek failed")
}


buf, ret = client.Read(fd598, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7r9V4sBqyuuTwa0tnvD8HnQGAmb3qxlcQ4gSB4EJV3YwAXsw8MF58fnYQL") {
  panic("wrong data returned")
}


ret = client.Seek(fd598, 87, client.SEEK_SET)
if(ret != 87) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 87)
  panic("seek failed")
}


buf, ret = client.Read(fd598, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "cGtvujYqgZI") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd598, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "7hu7kS9s8GjvKjI2wKh8pIjSTv4h_9czwwdnkGDPpZ8ZN5U5ROixd67r9V4sBqyuuTwa0tnvD8HnQGAmb3qx") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd598, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lcQ4gSB4EJV3YwAXsw8MF58fnYQLTn") {
  panic("wrong data returned")
}


ret = client.Seek(fd598, 106, client.SEEK_SET)
if(ret != 106) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 106)
  panic("seek failed")
}


fd603 := client.Open("/b6UqNkmZMn.txt", client.O_RDWR|client.O_CREATE)
if(fd603 < 0) {
  panic("open failed")
}


fd604 := client.Open("/MhNRYjnYbb.txt", client.O_RDWR|client.O_CREATE)
if(fd604 < 0) {
  panic("open failed")
}

//fd state: (106) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujYqgZI7hu7kS9s8GjvKjI2wKh8pIjSTv4h_9czwwdnkGDPpZ8ZN5U5ROixd67r9V4sBqyuuTwa0tnvD8HnQGAmb3qxlcQ4gSB4EJV3YwAXsw8MF58fnYQLTn

ret = client.Write(fd598, []byte("KMuz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) l7cmGw1kIP2x4CUtWwOQMY6Z5jpXTgaOWu9WYbyoBL98tL0U89WdFVytuIQg90qjLV681oNJDuJZSDCYXVjoMeacGtvujYqgZI7hu7kS9sKMuzKjI2wKh8pIjSTv4h_9czwwdnkGDPpZ8ZN5U5ROixd67r9V4sBqyuuTwa0tnvD8HnQGAmb3qxlcQ4gSB4EJV3YwAXsw8MF58fnYQLTn

ret = client.Close(fd604)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) yUYnwURDlPEutGGRUZSgw1FYgmdcxcSXeUiiT_KqFVqnbiAlhM6jlk_CAySw1Mlh8HHq5WM_kjINb0SvC0QaHoS8w7LAd

ret = client.Write(fd603, []byte("A3CTk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) A3CTkURDlPEutGGRUZSgw1FYgmdcxcSXeUiiT_KqFVqnbiAlhM6jlk_CAySw1Mlh8HHq5WM_kjINb0SvC0QaHoS8w7LAd

ret = client.Seek(fd598, 193, client.SEEK_SET)
if(ret != 193) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 193)
  panic("seek failed")
}


fd605 := client.Open("/5HpqVUKSSg.txt", client.O_RDWR|client.O_CREATE)
if(fd605 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd605, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd605)
if(ret != 0) {
  panic("close failed")
}

//fd state: (5) A3CTkURDlPEutGGRUZSgw1FYgmdcxcSXeUiiT_KqFVqnbiAlhM6jlk_CAySw1Mlh8HHq5WM_kjINb0SvC0QaHoS8w7LAd

ret = client.Write(fd603, []byte("dDwI5hI2Q3fQqD0eRVTj1Yg3rZRSSXb2FB1BIHMLFoyLcD0ZpeAQN4klt0ic1R2myeVrmRU3B4ZT5XvTH5xjP5M2YXX_uSKh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) A3CTkdDwI5hI2Q3fQqD0eRVTj1Yg3rZRSSXb2FB1BIHMLFoyLcD0ZpeAQN4klt0ic1R2myeVrmRU3B4ZT5XvTH5xjP5M2YXX_uSKh

buf, ret = client.Read(fd603, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd598, 165, client.SEEK_SET)
if(ret != 165) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 165)
  panic("seek failed")
}


ret = client.Close(fd598)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd603, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (101) A3CTkdDwI5hI2Q3fQqD0eRVTj1Yg3rZRSSXb2FB1BIHMLFoyLcD0ZpeAQN4klt0ic1R2myeVrmRU3B4ZT5XvTH5xjP5M2YXX_uSKh

ret = client.Write(fd603, []byte("rAM8egc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (108) A3CTkdDwI5hI2Q3fQqD0eRVTj1Yg3rZRSSXb2FB1BIHMLFoyLcD0ZpeAQN4klt0ic1R2myeVrmRU3B4ZT5XvTH5xjP5M2YXX_uSKhrAM8egc

buf, ret = client.Read(fd603, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd603, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


ret = client.Close(fd603)
if(ret != 0) {
  panic("close failed")
}


fd606 := client.Open("/GDNyWBC36b.txt", client.O_RDWR|client.O_CREATE)
if(fd606 < 0) {
  panic("open failed")
}


ret = client.Close(fd606)
if(ret != 0) {
  panic("close failed")
}


fd607 := client.Open("/PJprT8kWfg.txt", client.O_RDWR|client.O_CREATE)
if(fd607 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd607, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd607)
if(ret != 0) {
  panic("close failed")
}


fd608 := client.Open("/LEWZBDjjr3.txt", client.O_RDWR|client.O_CREATE)
if(fd608 < 0) {
  panic("open failed")
}


ret = client.Seek(fd608, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd608, 82)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd608, []byte("WvJ7tJwrxpxsA0cyAiItLd_25b_KDG4r"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) WvJ7tJwrxpxsA0cyAiItLd_25b_KDG4r
//fd state: (32) WvJ7tJwrxpxsA0cyAiItLd_25b_KDG4r

ret = client.Write(fd608, []byte("hNUMNlgLM5CFogWYKgvxKGreHCliYEi4iKgqSfmYZFQEGIpHmztCpJ4JL9QSsbS0x5X07eYw2GknkiEn_BIq19biloW2mReM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (128) WvJ7tJwrxpxsA0cyAiItLd_25b_KDG4rhNUMNlgLM5CFogWYKgvxKGreHCliYEi4iKgqSfmYZFQEGIpHmztCpJ4JL9QSsbS0x5X07eYw2GknkiEn_BIq19biloW2mReM
//fd state: (128) WvJ7tJwrxpxsA0cyAiItLd_25b_KDG4rhNUMNlgLM5CFogWYKgvxKGreHCliYEi4iKgqSfmYZFQEGIpHmztCpJ4JL9QSsbS0x5X07eYw2GknkiEn_BIq19biloW2mReM

ret = client.Write(fd608, []byte("tXcHaoogzWKWXrleUJgwr2QJus4HMx8JVEy8yIX91L1EcrcN0b9u4Ud_s9mlHQmd5ziRYyeBk9a2Xcc1XKEOUhiS9xVP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (220) WvJ7tJwrxpxsA0cyAiItLd_25b_KDG4rhNUMNlgLM5CFogWYKgvxKGreHCliYEi4iKgqSfmYZFQEGIpHmztCpJ4JL9QSsbS0x5X07eYw2GknkiEn_BIq19biloW2mReMtXcHaoogzWKWXrleUJgwr2QJus4HMx8JVEy8yIX91L1EcrcN0b9u4Ud_s9mlHQmd5ziRYyeBk9a2Xcc1XKEOUhiS9xVP

fd609 := client.Open("/ZJ6X2yc2mq.txt", client.O_RDWR|client.O_CREATE)
if(fd609 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd609, []byte("pluQY9kSsqm3yLtm_chpjcbPiIuWdKdWQdAkpQz39vbC6o_C4UKyjqveD_CtpcyFMjQTzPyPSWYf_hGcTYYR61"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) pluQY9kSsqm3yLtm_chpjcbPiIuWdKdWQdAkpQz39vbC6o_C4UKyjqveD_CtpcyFMjQTzPyPSWYf_hGcTYYR61

ret = client.Seek(fd609, 81, client.SEEK_SET)
if(ret != 81) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 81)
  panic("seek failed")
}


fd610 := client.Open("/fu3RTzcgJU.txt", client.O_RDWR|client.O_CREATE)
if(fd610 < 0) {
  panic("open failed")
}


ret = client.Seek(fd608, 119, client.SEEK_SET)
if(ret != 119) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 119)
  panic("seek failed")
}

//fd state: (81) pluQY9kSsqm3yLtm_chpjcbPiIuWdKdWQdAkpQz39vbC6o_C4UKyjqveD_CtpcyFMjQTzPyPSWYf_hGcTYYR61

ret = client.Write(fd609, []byte("6h5hx8uEO0fuoAQsGaihENuLusQAJZElAhp45IKCnqUeFWCDDxjkl7x4lHWB_DPkxdcutk3PI78X280KTaWJXd2VGA4o2w0eEN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (179) pluQY9kSsqm3yLtm_chpjcbPiIuWdKdWQdAkpQz39vbC6o_C4UKyjqveD_CtpcyFMjQTzPyPSWYf_hGcT6h5hx8uEO0fuoAQsGaihENuLusQAJZElAhp45IKCnqUeFWCDDxjkl7x4lHWB_DPkxdcutk3PI78X280KTaWJXd2VGA4o2w0eEN

buf, ret = client.Read(fd610, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd611 := client.Open("/dqzKim4aOI.txt", client.O_RDWR|client.O_CREATE)
if(fd611 < 0) {
  panic("open failed")
}

//fd state: (119) WvJ7tJwrxpxsA0cyAiItLd_25b_KDG4rhNUMNlgLM5CFogWYKgvxKGreHCliYEi4iKgqSfmYZFQEGIpHmztCpJ4JL9QSsbS0x5X07eYw2GknkiEn_BIq19biloW2mReMtXcHaoogzWKWXrleUJgwr2QJus4HMx8JVEy8yIX91L1EcrcN0b9u4Ud_s9mlHQmd5ziRYyeBk9a2Xcc1XKEOUhiS9xVP

ret = client.Write(fd608, []byte("mFfikgRjf9BTO_o5zkjUxKxTHhKe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (147) WvJ7tJwrxpxsA0cyAiItLd_25b_KDG4rhNUMNlgLM5CFogWYKgvxKGreHCliYEi4iKgqSfmYZFQEGIpHmztCpJ4JL9QSsbS0x5X07eYw2GknkiEn_BIq19bmFfikgRjf9BTO_o5zkjUxKxTHhKewr2QJus4HMx8JVEy8yIX91L1EcrcN0b9u4Ud_s9mlHQmd5ziRYyeBk9a2Xcc1XKEOUhiS9xVP

buf, ret = client.Read(fd611, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd608)
if(ret != 0) {
  panic("close failed")
}

//fd state: (179) pluQY9kSsqm3yLtm_chpjcbPiIuWdKdWQdAkpQz39vbC6o_C4UKyjqveD_CtpcyFMjQTzPyPSWYf_hGcT6h5hx8uEO0fuoAQsGaihENuLusQAJZElAhp45IKCnqUeFWCDDxjkl7x4lHWB_DPkxdcutk3PI78X280KTaWJXd2VGA4o2w0eEN

ret = client.Write(fd609, []byte("BysPkKh3H7RONY2Uc8rov_SN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (203) pluQY9kSsqm3yLtm_chpjcbPiIuWdKdWQdAkpQz39vbC6o_C4UKyjqveD_CtpcyFMjQTzPyPSWYf_hGcT6h5hx8uEO0fuoAQsGaihENuLusQAJZElAhp45IKCnqUeFWCDDxjkl7x4lHWB_DPkxdcutk3PI78X280KTaWJXd2VGA4o2w0eENBysPkKh3H7RONY2Uc8rov_SN
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd611, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd611, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd610, []byte("9nmrYxFvpVhGfr2x2p6c1XBQVAx6C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) 9nmrYxFvpVhGfr2x2p6c1XBQVAx6C

ret = client.Seek(fd611, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd609, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd612 := client.Open("/z3BIYvfpSJ.txt", client.O_RDWR|client.O_CREATE)
if(fd612 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd610, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd611)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd609)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd610, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd612, []byte("OgW2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) OgW2

buf, ret = client.Read(fd610, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "x6C") {
  panic("wrong data returned")
}


fd613 := client.Open("/Y9RWxzCpMk.txt", client.O_RDWR|client.O_CREATE)
if(fd613 < 0) {
  panic("open failed")
}


ret = client.Seek(fd610, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}

//fd state: (4) OgW2

ret = client.Write(fd612, []byte("fq8seA65wggcjsiInEWtAleUBjpRXFIUUgNmgD9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) OgW2fq8seA65wggcjsiInEWtAleUBjpRXFIUUgNmgD9

ret = client.Seek(fd613, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd614 := client.Open("/ZraNOL0sam.txt", client.O_RDWR|client.O_CREATE)
if(fd614 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd613, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd615 := client.Open("/axLom8SXqh.txt", client.O_RDWR|client.O_CREATE)
if(fd615 < 0) {
  panic("open failed")
}


ret = client.Seek(fd612, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


fd616 := client.Open("/gyu1b9QkCi.txt", client.O_RDWR|client.O_CREATE)
if(fd616 < 0) {
  panic("open failed")
}


ret = client.Close(fd614)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd613, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd617 := client.Open("/H3B8paSDlB.txt", client.O_RDWR|client.O_CREATE)
if(fd617 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd617, []byte("uXmpE8xxqglfrpsr4eQFF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) uXmpE8xxqglfrpsr4eQFF

buf, ret = client.Read(fd617, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd616, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qJCWMeoiiCmCbatMfdBUhrugm0uQx") {
  panic("wrong data returned")
}

//fd state: (24) OgW2fq8seA65wggcjsiInEWtAleUBjpRXFIUUgNmgD9

ret = client.Write(fd612, []byte("U_2_KqMjY7lBSmZjvEZbi2YRDa0WpzRWx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) OgW2fq8seA65wggcjsiInEWtU_2_KqMjY7lBSmZjvEZbi2YRDa0WpzRWx

ret = client.Seek(fd610, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Seek(fd616, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


buf, ret = client.Read(fd617, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd616)
if(ret != 0) {
  panic("close failed")
}

//fd state: (21) uXmpE8xxqglfrpsr4eQFF

ret = client.Write(fd617, []byte("mbf8gtQfA2jfo1NeaK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (39) uXmpE8xxqglfrpsr4eQFFmbf8gtQfA2jfo1NeaK

buf, ret = client.Read(fd615, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd615, []byte("9zb2pCkJgDPfqRRqvsVYlquBcVv73dnJJ6mw9lp7CpKwhhBGesVUlGsIvFZ_1zAGBq6cAuXJunEcuyPHmZHCfQmbhWOe3oSbxUj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) 9zb2pCkJgDPfqRRqvsVYlquBcVv73dnJJ6mw9lp7CpKwhhBGesVUlGsIvFZ_1zAGBq6cAuXJunEcuyPHmZHCfQmbhWOe3oSbxUj
//fd state: (57) OgW2fq8seA65wggcjsiInEWtU_2_KqMjY7lBSmZjvEZbi2YRDa0WpzRWx

ret = client.Write(fd612, []byte("OJZcPi3PFgjr6roWb8UTXJPaKOacTQf5pP_N"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) OgW2fq8seA65wggcjsiInEWtU_2_KqMjY7lBSmZjvEZbi2YRDa0WpzRWxOJZcPi3PFgjr6roWb8UTXJPaKOacTQf5pP_N

ret = client.Close(fd612)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd610, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Gfr2x2p6c1XBQVAx6C") {
  panic("wrong data returned")
}


ret = client.Seek(fd615, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}


buf, ret = client.Read(fd615, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "XJunEcu") {
  panic("wrong data returned")
}


ret = client.Seek(fd610, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Close(fd615)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd617, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (39) uXmpE8xxqglfrpsr4eQFFmbf8gtQfA2jfo1NeaK

ret = client.Write(fd617, []byte("SO8mhMt5Y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) uXmpE8xxqglfrpsr4eQFFmbf8gtQfA2jfo1NeaKSO8mhMt5Y

fd618 := client.Open("/3HJvawgO69.txt", client.O_RDWR|client.O_CREATE)
if(fd618 < 0) {
  panic("open failed")
}


fd619 := client.Open("/5WEEemNH2J.txt", client.O_RDWR|client.O_CREATE)
if(fd619 < 0) {
  panic("open failed")
}


ret = client.Seek(fd618, 90, client.SEEK_SET)
if(ret != 90) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 90)
  panic("seek failed")
}


ret = client.Close(fd619)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd618, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "s86RI6Tn3UB9_2bGLQCjQuttj2X8TCntzhCruNWB_2i") {
  panic("wrong data returned")
}


fd620 := client.Open("/_9uGqB8bKD.txt", client.O_RDWR|client.O_CREATE)
if(fd620 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd610, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "p6c1XBQVAx6C") {
  panic("wrong data returned")
}


fd621 := client.Open("/mRe0WtXvfw.txt", client.O_RDWR|client.O_CREATE)
if(fd621 < 0) {
  panic("open failed")
}


ret = client.Seek(fd613, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (133) dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobxt3CHFeI1QPc7GzIRJs86RI6Tn3UB9_2bGLQCjQuttj2X8TCntzhCruNWB_2i

ret = client.Write(fd618, []byte("IZhJ0xXFDHDaJniDxyunTKB9h5nGOmsjEA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (167) dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobxt3CHFeI1QPc7GzIRJs86RI6Tn3UB9_2bGLQCjQuttj2X8TCntzhCruNWB_2iIZhJ0xXFDHDaJniDxyunTKB9h5nGOmsjEA

fd622 := client.Open("/xFY5pLXtzh.txt", client.O_RDWR|client.O_CREATE)
if(fd622 < 0) {
  panic("open failed")
}


ret = client.Close(fd620)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd618, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd617)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd613)
if(ret != 0) {
  panic("close failed")
}


fd623 := client.Open("/9dluZ9EKXT.txt", client.O_RDWR|client.O_CREATE)
if(fd623 < 0) {
  panic("open failed")
}

//fd state: (167) dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobxt3CHFeI1QPc7GzIRJs86RI6Tn3UB9_2bGLQCjQuttj2X8TCntzhCruNWB_2iIZhJ0xXFDHDaJniDxyunTKB9h5nGOmsjEA

ret = client.Write(fd618, []byte("Y3bzi14"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (174) dZ5nTo23T4y0QT7m2517jQfzOkTmKiU_DhzFy_VWseE1J6w1FqT6k7fjqmet_xP40ooedUobxt3CHFeI1QPc7GzIRJs86RI6Tn3UB9_2bGLQCjQuttj2X8TCntzhCruNWB_2iIZhJ0xXFDHDaJniDxyunTKB9h5nGOmsjEAY3bzi14

ret = client.Seek(fd618, 58, client.SEEK_SET)
if(ret != 58) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 58)
  panic("seek failed")
}


fd624 := client.Open("/_dAWcMs1ny.txt", client.O_RDWR|client.O_CREATE)
if(fd624 < 0) {
  panic("open failed")
}


ret = client.Close(fd618)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd622, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd622)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd621, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd623, 177, client.SEEK_SET)
if(ret != 177) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 177)
  panic("seek failed")
}


fd625 := client.Open("/MhNRYjnYbb.txt", client.O_RDWR|client.O_CREATE)
if(fd625 < 0) {
  panic("open failed")
}

//fd state: (29) 9nmrYxFvpVhGfr2x2p6c1XBQVAx6C

ret = client.Write(fd610, []byte("Rz6CDDxPq2mRA_lCBYtKe_Tb_JFNLnWBmm29pMO_or9v9dKm2Pgzr2umVGTpifjEjfRVmv4tGDOayWjCdDTb1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (114) 9nmrYxFvpVhGfr2x2p6c1XBQVAx6CRz6CDDxPq2mRA_lCBYtKe_Tb_JFNLnWBmm29pMO_or9v9dKm2Pgzr2umVGTpifjEjfRVmv4tGDOayWjCdDTb1
//fd state: (0) 5XmdnaNO1MFJPrC3HiefPzPah5K4SFlK9wx9jUYchOWcM1UmOVq0HtVgLPFB80kAbgjljfFOtohelXf2LlRFPz8EwbnDO3BsGXY0fMPfcgRDJ9eWxaT1vb1Fo7S6oxqEjewLHyirRhcNowluq3lsXyIq_wtSWb8iXELrNA0HmAOIxCJ2fHX5g2a

ret = client.Write(fd625, []byte("u90RH41Sn94WSUnEzNAFHT5ouL58Xqfd1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) u90RH41Sn94WSUnEzNAFHT5ouL58Xqfd1wx9jUYchOWcM1UmOVq0HtVgLPFB80kAbgjljfFOtohelXf2LlRFPz8EwbnDO3BsGXY0fMPfcgRDJ9eWxaT1vb1Fo7S6oxqEjewLHyirRhcNowluq3lsXyIq_wtSWb8iXELrNA0HmAOIxCJ2fHX5g2a

ret = client.Seek(fd610, 111, client.SEEK_SET)
if(ret != 111) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 111)
  panic("seek failed")
}


buf, ret = client.Read(fd623, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "sYk99PJzu4AqPS50NvIFKtfdf9YO3Ug59ZuOXqE0kpFMbSXnrQIZwHvFW5OJ1d7gL6VNvlfXuKCTxvHF2ktW2AwrGosgl_pDZ") {
  panic("wrong data returned")
}


ret = client.Close(fd625)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd621, []byte("C6vvVWTSx9PlpNtccozeCIN4r3I8QT9jBvpwxLxdGP_OLYs8mLyxgVwciP5Kz0gv0qQybqy_g341OnlclfRIvr_6yMk2gkZ2ki"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) C6vvVWTSx9PlpNtccozeCIN4r3I8QT9jBvpwxLxdGP_OLYs8mLyxgVwciP5Kz0gv0qQybqy_g341OnlclfRIvr_6yMk2gkZ2ki

buf, ret = client.Read(fd624, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (274) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu3Xa8LUU292eZBp84ZnmSvMayWnPIMAC25nm5j9Mp555cYTfWLfJ4EaCL16wmxxiTsTOfWGJ4Td9zvQbTg7xnHHUTJN4YEblSj8qp9zXD0xUDjIqYWP1sJvWAsqBwsSLcrAn6rdXsYk99PJzu4AqPS50NvIFKtfdf9YO3Ug59ZuOXqE0kpFMbSXnrQIZwHvFW5OJ1d7gL6VNvlfXuKCTxvHF2ktW2AwrGosgl_pDZ2BSES3ZAVIcC1nuiZ2GASbYZX1K08qYcdkWNK1Znmy0

ret = client.Write(fd623, []byte("TM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (276) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu3Xa8LUU292eZBp84ZnmSvMayWnPIMAC25nm5j9Mp555cYTfWLfJ4EaCL16wmxxiTsTOfWGJ4Td9zvQbTg7xnHHUTJN4YEblSj8qp9zXD0xUDjIqYWP1sJvWAsqBwsSLcrAn6rdXsYk99PJzu4AqPS50NvIFKtfdf9YO3Ug59ZuOXqE0kpFMbSXnrQIZwHvFW5OJ1d7gL6VNvlfXuKCTxvHF2ktW2AwrGosgl_pDZTMSES3ZAVIcC1nuiZ2GASbYZX1K08qYcdkWNK1Znmy0
//fd state: (111) 9nmrYxFvpVhGfr2x2p6c1XBQVAx6CRz6CDDxPq2mRA_lCBYtKe_Tb_JFNLnWBmm29pMO_or9v9dKm2Pgzr2umVGTpifjEjfRVmv4tGDOayWjCdDTb1

ret = client.Write(fd610, []byte("NVJyPy0xCB9jHc4dffnkmcNyv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) 9nmrYxFvpVhGfr2x2p6c1XBQVAx6CRz6CDDxPq2mRA_lCBYtKe_Tb_JFNLnWBmm29pMO_or9v9dKm2Pgzr2umVGTpifjEjfRVmv4tGDOayWjCdDNVJyPy0xCB9jHc4dffnkmcNyv

ret = client.Close(fd610)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd624, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd621)
if(ret != 0) {
  panic("close failed")
}


fd626 := client.Open("/lddRycIQTm.txt", client.O_RDWR|client.O_CREATE)
if(fd626 < 0) {
  panic("open failed")
}


fd627 := client.Open("/iBDBmf4131.txt", client.O_RDWR|client.O_CREATE)
if(fd627 < 0) {
  panic("open failed")
}


ret = client.Seek(fd624, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd623)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd626, []byte("viedmIij3KVeJgVnY2eUINrfL92xP3j8RfhYqE_Pcp4Z2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) viedmIij3KVeJgVnY2eUINrfL92xP3j8RfhYqE_Pcp4Z2

fd628 := client.Open("/bscOPxCqaM.txt", client.O_RDWR|client.O_CREATE)
if(fd628 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd624, []byte("YXW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) YXW

ret = client.Close(fd624)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd628)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd626)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd627)
if(ret != 0) {
  panic("close failed")
}


fd629 := client.Open("/WEwnP8V8uo.txt", client.O_RDWR|client.O_CREATE)
if(fd629 < 0) {
  panic("open failed")
}


ret = client.Seek(fd629, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}


ret = client.Seek(fd629, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


buf, ret = client.Read(fd629, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "R_aCeV8H8y4zouEri7411jBLyo4") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd629, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (65) l61mpyACaIKvsijlcMQsw0pjPO_kDY_Wp_fD23R_aCeV8H8y4zouEri7411jBLyo4

ret = client.Write(fd629, []byte("dgF9xRp0Da"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (75) l61mpyACaIKvsijlcMQsw0pjPO_kDY_Wp_fD23R_aCeV8H8y4zouEri7411jBLyo4dgF9xRp0Da

ret = client.Close(fd629)
if(ret != 0) {
  panic("close failed")
}


fd630 := client.Open("/hCUNkqYJuK.txt", client.O_RDWR|client.O_CREATE)
if(fd630 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd630)
if(ret != 0) {
  panic("close failed")
}


fd631 := client.Open("/WmilsB_bL8.txt", client.O_RDWR|client.O_CREATE)
if(fd631 < 0) {
  panic("open failed")
}


ret = client.Close(fd631)
if(ret != 0) {
  panic("close failed")
}


fd632 := client.Open("/3dMK9W1zC2.txt", client.O_RDWR|client.O_CREATE)
if(fd632 < 0) {
  panic("open failed")
}


ret = client.Close(fd632)
if(ret != 0) {
  panic("close failed")
}


fd633 := client.Open("/WNMJT2E_a5.txt", client.O_RDWR|client.O_CREATE)
if(fd633 < 0) {
  panic("open failed")
}

//fd state: (0) o0WYhFJX8YWxvIFCMCGxl60Dfy1QIneeJG88YCNm98WsuuzbOC8ptmY77Nbhr9dpIRbnsh

ret = client.Write(fd633, []byte("jGhuTeatzbyKczU0nQRjPYg8ClP6jIgi8buXzjNj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) jGhuTeatzbyKczU0nQRjPYg8ClP6jIgi8buXzjNj98WsuuzbOC8ptmY77Nbhr9dpIRbnsh

ret = client.Seek(fd633, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}

//fd state: (63) jGhuTeatzbyKczU0nQRjPYg8ClP6jIgi8buXzjNj98WsuuzbOC8ptmY77Nbhr9dpIRbnsh

ret = client.Write(fd633, []byte("czu0vaNoygRtjS_AaUffbiGtqmspKeaW457iOKmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (152) jGhuTeatzbyKczU0nQRjPYg8ClP6jIgi8buXzjNj98WsuuzbOC8ptmY77Nbhr9dczu0vaNoygRtjS_AaUffbiGtqmspKeaW457iOKmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyY
//fd state: (152) jGhuTeatzbyKczU0nQRjPYg8ClP6jIgi8buXzjNj98WsuuzbOC8ptmY77Nbhr9dczu0vaNoygRtjS_AaUffbiGtqmspKeaW457iOKmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyY

ret = client.Write(fd633, []byte("kTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (210) jGhuTeatzbyKczU0nQRjPYg8ClP6jIgi8buXzjNj98WsuuzbOC8ptmY77Nbhr9dczu0vaNoygRtjS_AaUffbiGtqmspKeaW457iOKmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyYkTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M

fd634 := client.Open("/c0zFybvjxo.txt", client.O_RDWR|client.O_CREATE)
if(fd634 < 0) {
  panic("open failed")
}


fd635 := client.Open("/gQ_ImV6qgZ.txt", client.O_RDWR|client.O_CREATE)
if(fd635 < 0) {
  panic("open failed")
}


fd636 := client.Open("/8XFLyCw86K.txt", client.O_RDWR|client.O_CREATE)
if(fd636 < 0) {
  panic("open failed")
}

//fd state: (210) jGhuTeatzbyKczU0nQRjPYg8ClP6jIgi8buXzjNj98WsuuzbOC8ptmY77Nbhr9dczu0vaNoygRtjS_AaUffbiGtqmspKeaW457iOKmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyYkTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M

ret = client.Write(fd633, []byte("00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3Xnhl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (266) jGhuTeatzbyKczU0nQRjPYg8ClP6jIgi8buXzjNj98WsuuzbOC8ptmY77Nbhr9dczu0vaNoygRtjS_AaUffbiGtqmspKeaW457iOKmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyYkTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3Xnhl
//fd state: (0) 

ret = client.Write(fd634, []byte("2YEXMaoA5lm8LIBph5CQyaou8JdepqtGFojGME3AjYffc95N2xTP985uqGDriqJWhL57"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) 2YEXMaoA5lm8LIBph5CQyaou8JdepqtGFojGME3AjYffc95N2xTP985uqGDriqJWhL57

ret = client.Seek(fd634, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


ret = client.Seek(fd634, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


buf, ret = client.Read(fd634, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Yffc95N2xTP985uqGDriqJWhL57") {
  panic("wrong data returned")
}


ret = client.Seek(fd634, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Seek(fd634, 39, client.SEEK_SET)
if(ret != 39) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 39)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd636, []byte("IvGvs8F8XnZ4p5C5Ke2WD8Sl_NOnH9dKlaGIr4gChMpyvlCvMH3lTtvZ6eV2y2qf7skCF5cGAP8NANOrWWJPvOqc5aAt_mi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) IvGvs8F8XnZ4p5C5Ke2WD8Sl_NOnH9dKlaGIr4gChMpyvlCvMH3lTtvZ6eV2y2qf7skCF5cGAP8NANOrWWJPvOqc5aAt_mi

fd637 := client.Open("/1hYlxqnLZN.txt", client.O_RDWR|client.O_CREATE)
if(fd637 < 0) {
  panic("open failed")
}


ret = client.Seek(fd634, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


buf, ret = client.Read(fd634, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GFojGME3AjYffc95N2xTP985uqGDriqJWhL57") {
  panic("wrong data returned")
}


ret = client.Close(fd636)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd637, []byte("kNFje5jzWpr4wcLYDGQYe8jJXQmljmLsHCT07Yxt8chFMVnCPeC7uEHr0LZ2Mwf1UZsB32CwWicR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) kNFje5jzWpr4wcLYDGQYe8jJXQmljmLsHCT07Yxt8chFMVnCPeC7uEHr0LZ2Mwf1UZsB32CwWicR

ret = client.Seek(fd634, 18, client.SEEK_SET)
if(ret != 18) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 18)
  panic("seek failed")
}


fd638 := client.Open("/NmitfjLdVG.txt", client.O_RDWR|client.O_CREATE)
if(fd638 < 0) {
  panic("open failed")
}


ret = client.Seek(fd633, 172, client.SEEK_SET)
if(ret != 172) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 172)
  panic("seek failed")
}


buf, ret = client.Read(fd638, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd639 := client.Open("/DFjf9ipCIz.txt", client.O_RDWR|client.O_CREATE)
if(fd639 < 0) {
  panic("open failed")
}


ret = client.Seek(fd637, 73, client.SEEK_SET)
if(ret != 73) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 73)
  panic("seek failed")
}


ret = client.Seek(fd637, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


buf, ret = client.Read(fd638, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd640 := client.Open("/ULwj3UEeoU.txt", client.O_RDWR|client.O_CREATE)
if(fd640 < 0) {
  panic("open failed")
}


ret = client.Seek(fd633, 131, client.SEEK_SET)
if(ret != 131) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 131)
  panic("seek failed")
}


ret = client.Seek(fd633, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


buf, ret = client.Read(fd637, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "QmljmLsHCT07Yxt8chFMVnCPeC7uEHr0LZ2Mwf1UZsB32CwWicR") {
  panic("wrong data returned")
}

//fd state: (0) BbbdVFZRGyCr_2GEDBUp6Kn6c74EzzmKk4hij01kaq

ret = client.Write(fd635, []byte("cVg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (3) cVgdVFZRGyCr_2GEDBUp6Kn6c74EzzmKk4hij01kaq
//fd state: (0) 

ret = client.Write(fd639, []byte("e5WifJPhYBHbnaXpMTW9IBB5mrX28AeLo8kpeFs0hT9VMQV0URuPoJOQGnC1ZyDSksCcYvnRSBoSMy6VUv2Sdw39DYdHRN4q0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) e5WifJPhYBHbnaXpMTW9IBB5mrX28AeLo8kpeFs0hT9VMQV0URuPoJOQGnC1ZyDSksCcYvnRSBoSMy6VUv2Sdw39DYdHRN4q0

ret = client.Close(fd637)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd634, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Close(fd640)
if(ret != 0) {
  panic("close failed")
}


fd641 := client.Open("/3pYSlSk_iq.txt", client.O_RDWR|client.O_CREATE)
if(fd641 < 0) {
  panic("open failed")
}


fd642 := client.Open("/4TdLNMxHgT.txt", client.O_RDWR|client.O_CREATE)
if(fd642 < 0) {
  panic("open failed")
}


fd643 := client.Open("/nG8kZjpIOY.txt", client.O_RDWR|client.O_CREATE)
if(fd643 < 0) {
  panic("open failed")
}


ret = client.Seek(fd642, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd642, []byte("yJFjTlE7bLGBFhZbI8C7RtfoZwI9Vh"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) yJFjTlE7bLGBFhZbI8C7RtfoZwI9Vh

fd644 := client.Open("/aHSi5Rp9bT.txt", client.O_RDWR|client.O_CREATE)
if(fd644 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd644, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (29) jGhuTeatzbyKczU0nQRjPYg8ClP6jIgi8buXzjNj98WsuuzbOC8ptmY77Nbhr9dczu0vaNoygRtjS_AaUffbiGtqmspKeaW457iOKmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyYkTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3Xnhl

ret = client.Write(fd633, []byte("l2DA_SNkf4mQGVqB5QXoQJAiuO7zriLdARzrvVeuo5zt2c9dEONaE0_9ZkescctVzmhoSlZE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (101) jGhuTeatzbyKczU0nQRjPYg8ClP6jl2DA_SNkf4mQGVqB5QXoQJAiuO7zriLdARzrvVeuo5zt2c9dEONaE0_9ZkescctVzmhoSlZEmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyYkTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3Xnhl

fd645 := client.Open("/YyURkaOMEI.txt", client.O_RDWR|client.O_CREATE)
if(fd645 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd643, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd646 := client.Open("/jUD4VX4_c1.txt", client.O_RDWR|client.O_CREATE)
if(fd646 < 0) {
  panic("open failed")
}


ret = client.Seek(fd634, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


ret = client.Seek(fd635, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Close(fd633)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd641, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd644, []byte("Mb6uY5TKsAC9Ocz7wHwLRVINhwgfDAEXY1YWfn3UDqVvNoI3iA_G_wbe5KVrTmQ0BnjDCjjEGh0hS7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) Mb6uY5TKsAC9Ocz7wHwLRVINhwgfDAEXY1YWfn3UDqVvNoI3iA_G_wbe5KVrTmQ0BnjDCjjEGh0hS7

buf, ret = client.Read(fd634, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YEXMaoA5lm8LIBph5CQyaou8JdepqtGFojGME3AjYffc95N2xTP985uqGDriqJWh") {
  panic("wrong data returned")
}


ret = client.Close(fd644)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd641, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd641, []byte("2NE0CD3zDSC3PjzPERIPOEkjgHBqK3QVNyX6nLDgTbT2lpf3ojsVSgLYyjG11feP0ejBvdFyWVU8goclwVv3rFI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) 2NE0CD3zDSC3PjzPERIPOEkjgHBqK3QVNyX6nLDgTbT2lpf3ojsVSgLYyjG11feP0ejBvdFyWVU8goclwVv3rFI

ret = client.Close(fd645)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd635)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd634, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


fd647 := client.Open("/PezyqqHqGJ.txt", client.O_RDWR|client.O_CREATE)
if(fd647 < 0) {
  panic("open failed")
}


fd648 := client.Open("/hFI4pcJrPs.txt", client.O_RDWR|client.O_CREATE)
if(fd648 < 0) {
  panic("open failed")
}


fd649 := client.Open("/Uvk9UV6e8V.txt", client.O_RDWR|client.O_CREATE)
if(fd649 < 0) {
  panic("open failed")
}


ret = client.Close(fd643)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd638, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd642)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd648, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) gROrUzao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJG8Hr1Soaw6gy4nWhLw_XUcAKW2uLrFdmR86bqKTE8S8kEiN86IrFn94PRM32F8RevY2_xmInVJo1XAg3EJLWtxsUxNlJ_4W7UMRjpjFfHkTTKCv1e3RW_YJR49qr1KT

ret = client.Write(fd646, []byte("L5FWKpao"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) L5FWKpao3biiAEE7JZ45sPXmlGcH9kvGcas4Oe0SV0ayartf4MpeshVYnrqCgmRkXbmpPScJG8Hr1Soaw6gy4nWhLw_XUcAKW2uLrFdmR86bqKTE8S8kEiN86IrFn94PRM32F8RevY2_xmInVJo1XAg3EJLWtxsUxNlJ_4W7UMRjpjFfHkTTKCv1e3RW_YJR49qr1KT

fd650 := client.Open("/5OR5HN2nWE.txt", client.O_RDWR|client.O_CREATE)
if(fd650 < 0) {
  panic("open failed")
}


ret = client.Close(fd650)
if(ret != 0) {
  panic("close failed")
}


fd651 := client.Open("/0fk9a1vh_D.txt", client.O_RDWR|client.O_CREATE)
if(fd651 < 0) {
  panic("open failed")
}


ret = client.Close(fd649)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd639, 82, client.SEEK_SET)
if(ret != 82) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 82)
  panic("seek failed")
}


ret = client.Close(fd641)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd646, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd648, []byte("pcBxJdKGjHLXtyYDFqxLpPt_fhEyxaDLeNCCshQ4YkrpESDNRGFnKB5p0wnVK4pbyaLEeoWoN8MOA0fqeg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) pcBxJdKGjHLXtyYDFqxLpPt_fhEyxaDLeNCCshQ4YkrpESDNRGFnKB5p0wnVK4pbyaLEeoWoN8MOA0fqeg

ret = client.Seek(fd634, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Seek(fd651, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (3) 2YEXMaoA5lm8LIBph5CQyaou8JdepqtGFojGME3AjYffc95N2xTP985uqGDriqJWhL57

ret = client.Write(fd634, []byte("8p41AlQ8i3ojS2QBnhzFy0VZhxSnqO1HEfTkBNs"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) 2YE8p41AlQ8i3ojS2QBnhzFy0VZhxSnqO1HEfTkBNsffc95N2xTP985uqGDriqJWhL57

buf, ret = client.Read(fd638, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd651)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd648, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


fd652 := client.Open("/hFI4pcJrPs.txt", client.O_RDWR|client.O_CREATE)
if(fd652 < 0) {
  panic("open failed")
}


fd653 := client.Open("/0fk9a1vh_D.txt", client.O_RDWR|client.O_CREATE)
if(fd653 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd653, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd647)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd639, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


fd654 := client.Open("/8XFLyCw86K.txt", client.O_RDWR|client.O_CREATE)
if(fd654 < 0) {
  panic("open failed")
}


fd655 := client.Open("/XFfR2z8Yx8.txt", client.O_RDWR|client.O_CREATE)
if(fd655 < 0) {
  panic("open failed")
}


ret = client.Seek(fd646, 156, client.SEEK_SET)
if(ret != 156) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 156)
  panic("seek failed")
}


buf, ret = client.Read(fd638, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd652)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd653)
if(ret != 0) {
  panic("close failed")
}


fd656 := client.Open("/hF6OAOMNoP.txt", client.O_RDWR|client.O_CREATE)
if(fd656 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd646)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd634, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


fd657 := client.Open("/8ysb58u2tQ.txt", client.O_RDWR|client.O_CREATE)
if(fd657 < 0) {
  panic("open failed")
}


fd658 := client.Open("/dqzKim4aOI.txt", client.O_RDWR|client.O_CREATE)
if(fd658 < 0) {
  panic("open failed")
}


ret = client.Seek(fd656, 167, client.SEEK_SET)
if(ret != 167) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 167)
  panic("seek failed")
}


buf, ret = client.Read(fd648, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "NCCshQ4YkrpESDNRGFnKB5p0wn") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd658, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd648, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


buf, ret = client.Read(fd648, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "xaDLeNCCshQ4YkrpESDNRG") {
  panic("wrong data returned")
}


ret = client.Seek(fd638, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 7tTVQXHcopQx8KuTQsx6CCyRVjQooGji1QUU0NGioxMQIbowQQupHt4N2xPBo5nXE3J

ret = client.Write(fd655, []byte("S2pIevhg_3VsuH2Wz8OXSbhgoKZyqbi7m6CYiIsQRkxIsmiTqJT0Ii2Wl6ujUvjeMypU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (68) S2pIevhg_3VsuH2Wz8OXSbhgoKZyqbi7m6CYiIsQRkxIsmiTqJT0Ii2Wl6ujUvjeMypU

buf, ret = client.Read(fd656, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AopjH2D4Qm2iDSklHYldqJ7frYLv1xd33WocC5PkF1A01MjliqmHIjhwOJaaIrZyuJTmZ") {
  panic("wrong data returned")
}


ret = client.Close(fd638)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd658, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd659 := client.Open("/qCh1SfXxxY.txt", client.O_RDWR|client.O_CREATE)
if(fd659 < 0) {
  panic("open failed")
}


fd660 := client.Open("/v2DsGdHGbp.txt", client.O_RDWR|client.O_CREATE)
if(fd660 < 0) {
  panic("open failed")
}

//fd state: (16) 2YE8p41AlQ8i3ojS2QBnhzFy0VZhxSnqO1HEfTkBNsffc95N2xTP985uqGDriqJWhL57

ret = client.Write(fd634, []byte("jvaJSUo7qAAmJYsqekejkzYfWUhgkJhmmwyNcqud"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) 2YE8p41AlQ8i3ojSjvaJSUo7qAAmJYsqekejkzYfWUhgkJhmmwyNcqudqGDriqJWhL57

ret = client.Seek(fd639, 86, client.SEEK_SET)
if(ret != 86) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 86)
  panic("seek failed")
}


ret = client.Close(fd655)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd659)
if(ret != 0) {
  panic("close failed")
}

//fd state: (56) 2YE8p41AlQ8i3ojSjvaJSUo7qAAmJYsqekejkzYfWUhgkJhmmwyNcqudqGDriqJWhL57

ret = client.Write(fd634, []byte("yDAcz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) 2YE8p41AlQ8i3ojSjvaJSUo7qAAmJYsqekejkzYfWUhgkJhmmwyNcqudyDAczqJWhL57

ret = client.Close(fd639)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) IvGvs8F8XnZ4p5C5Ke2WD8Sl_NOnH9dKlaGIr4gChMpyvlCvMH3lTtvZ6eV2y2qf7skCF5cGAP8NANOrWWJPvOqc5aAt_mi

ret = client.Write(fd654, []byte("RQWoVCnDTZSkm7K97PKS22au6B5N8Drxb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) RQWoVCnDTZSkm7K97PKS22au6B5N8DrxbaGIr4gChMpyvlCvMH3lTtvZ6eV2y2qf7skCF5cGAP8NANOrWWJPvOqc5aAt_mi

fd661 := client.Open("/L3uDEXPgPs.txt", client.O_RDWR|client.O_CREATE)
if(fd661 < 0) {
  panic("open failed")
}


ret = client.Close(fd654)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd660, 70)
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


ret = client.Close(fd657)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd661, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd660)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd634)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd661, []byte("xuOEnzQOBU2XSbubKw0N0hV0S89rxg7j5PGsbrBxYv3Np_xrQFqvG2P8zvooW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) xuOEnzQOBU2XSbubKw0N0hV0S89rxg7j5PGsbrBxYv3Np_xrQFqvG2P8zvooW

ret = client.Close(fd656)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd661, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}

//fd state: (26) xuOEnzQOBU2XSbubKw0N0hV0S89rxg7j5PGsbrBxYv3Np_xrQFqvG2P8zvooW

ret = client.Write(fd661, []byte("Z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) xuOEnzQOBU2XSbubKw0N0hV0S8Zrxg7j5PGsbrBxYv3Np_xrQFqvG2P8zvooW

ret = client.Seek(fd648, 79, client.SEEK_SET)
if(ret != 79) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 79)
  panic("seek failed")
}


ret = client.Close(fd648)
if(ret != 0) {
  panic("close failed")
}

//fd state: (27) xuOEnzQOBU2XSbubKw0N0hV0S8Zrxg7j5PGsbrBxYv3Np_xrQFqvG2P8zvooW

ret = client.Write(fd661, []byte("4jqFhwGIZij1PaAkIivcwVsD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) xuOEnzQOBU2XSbubKw0N0hV0S8Z4jqFhwGIZij1PaAkIivcwVsDvG2P8zvooW

ret = client.Close(fd661)
if(ret != 0) {
  panic("close failed")
}


fd662 := client.Open("/qSmDSKYTxz.txt", client.O_RDWR|client.O_CREATE)
if(fd662 < 0) {
  panic("open failed")
}


fd663 := client.Open("/E5X4o13iGO.txt", client.O_RDWR|client.O_CREATE)
if(fd663 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd663, []byte("w9lv9WNCyNjb_tHs8Hh__gOhg1tD53k4E8IhKhLN5tqM35gjC0qU5CezzJ3Hb0pzHD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) w9lv9WNCyNjb_tHs8Hh__gOhg1tD53k4E8IhKhLN5tqM35gjC0qU5CezzJ3Hb0pzHD
//fd state: (66) w9lv9WNCyNjb_tHs8Hh__gOhg1tD53k4E8IhKhLN5tqM35gjC0qU5CezzJ3Hb0pzHD

ret = client.Write(fd663, []byte("CEcx9VN0GEQ_NCwybJzFNx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) w9lv9WNCyNjb_tHs8Hh__gOhg1tD53k4E8IhKhLN5tqM35gjC0qU5CezzJ3Hb0pzHDCEcx9VN0GEQ_NCwybJzFNx

fd664 := client.Open("/aW_9PFLHIJ.txt", client.O_RDWR|client.O_CREATE)
if(fd664 < 0) {
  panic("open failed")
}


ret = client.Seek(fd664, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd664, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd663)
if(ret != 0) {
  panic("close failed")
}


fd665 := client.Open("/O5jLTz9ceU.txt", client.O_RDWR|client.O_CREATE)
if(fd665 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd664, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd664, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd665, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "IFoZ6wJem8tM2mgLHoErfmKp5lJAP5AHYK1vTBdWWrAJXchYU0jOu5riE_W_io9ueEeOJmgWXyJ") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd664, []byte("g3X_JITFj4wEmYPF84whorILX9nuNZk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) g3X_JITFj4wEmYPF84whorILX9nuNZk

buf, ret = client.Read(fd665, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_") {
  panic("wrong data returned")
}


ret = client.Seek(fd665, 144, client.SEEK_SET)
if(ret != 144) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 144)
  panic("seek failed")
}


fd666 := client.Open("/lcOs6N5rL8.txt", client.O_RDWR|client.O_CREATE)
if(fd666 < 0) {
  panic("open failed")
}


fd667 := client.Open("/OEztjY7sQ0.txt", client.O_RDWR|client.O_CREATE)
if(fd667 < 0) {
  panic("open failed")
}


ret = client.Close(fd665)
if(ret != 0) {
  panic("close failed")
}


fd668 := client.Open("/0t_D6cCpLL.txt", client.O_RDWR|client.O_CREATE)
if(fd668 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd666, 95)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd667, []byte("hNxaZwV3HjDucCuLtbyUIncnTnR2G3oe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) hNxaZwV3HjDucCuLtbyUIncnTnR2G3oe

ret = client.Close(fd662)
if(ret != 0) {
  panic("close failed")
}

//fd state: (32) hNxaZwV3HjDucCuLtbyUIncnTnR2G3oe

ret = client.Write(fd667, []byte("1cpBOQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) hNxaZwV3HjDucCuLtbyUIncnTnR2G3oe1cpBOQ

buf, ret = client.Read(fd667, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd664, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


ret = client.Seek(fd666, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd668)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd664, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "orILX9nuNZk") {
  panic("wrong data returned")
}


ret = client.Seek(fd664, 9, client.SEEK_SET)
if(ret != 9) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 9)
  panic("seek failed")
}


buf, ret = client.Read(fd667, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd664, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4wEmYPF84whorILX9nuNZk") {
  panic("wrong data returned")
}


ret = client.Seek(fd667, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}


ret = client.Close(fd666)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd667, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}

//fd state: (38) hNxaZwV3HjDucCuLtbyUIncnTnR2G3oe1cpBOQ

ret = client.Write(fd667, []byte("kck6HQg1OyAhF4ocKCNwpO668uRJ66nS04WU29Q972SPJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) hNxaZwV3HjDucCuLtbyUIncnTnR2G3oe1cpBOQkck6HQg1OyAhF4ocKCNwpO668uRJ66nS04WU29Q972SPJ

fd669 := client.Open("/uIv3bq5PSV.txt", client.O_RDWR|client.O_CREATE)
if(fd669 < 0) {
  panic("open failed")
}

//fd state: (0) VwhCGYzEkyAcrqGa3Cajsts0TXpywUzty6PE6ctzYvXYtjAMLfcjxjqwMIp8zd2Dkuw3x

ret = client.Write(fd669, []byte("sjOtsBYxkKA9tJ1qKak1fmgSwGtWgUSB4hTkx2IkRG3_srW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) sjOtsBYxkKA9tJ1qKak1fmgSwGtWgUSB4hTkx2IkRG3_srWMLfcjxjqwMIp8zd2Dkuw3x
//fd state: (83) hNxaZwV3HjDucCuLtbyUIncnTnR2G3oe1cpBOQkck6HQg1OyAhF4ocKCNwpO668uRJ66nS04WU29Q972SPJ

ret = client.Write(fd667, []byte("6u6WizavwtjwxP5b4FDrwzIsMPK2aH10jn7Ub"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) hNxaZwV3HjDucCuLtbyUIncnTnR2G3oe1cpBOQkck6HQg1OyAhF4ocKCNwpO668uRJ66nS04WU29Q972SPJ6u6WizavwtjwxP5b4FDrwzIsMPK2aH10jn7Ub

ret = client.Seek(fd667, 40, client.SEEK_SET)
if(ret != 40) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 40)
  panic("seek failed")
}


ret = client.Close(fd667)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd669, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Seek(fd669, 68, client.SEEK_SET)
if(ret != 68) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 68)
  panic("seek failed")
}


fd670 := client.Open("/ZrRGFqOamQ.txt", client.O_RDWR|client.O_CREATE)
if(fd670 < 0) {
  panic("open failed")
}


ret = client.Seek(fd670, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Seek(fd664, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}

//fd state: (68) sjOtsBYxkKA9tJ1qKak1fmgSwGtWgUSB4hTkx2IkRG3_srWMLfcjxjqwMIp8zd2Dkuw3x

ret = client.Write(fd669, []byte("BDDjGR9n40aNKDM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) sjOtsBYxkKA9tJ1qKak1fmgSwGtWgUSB4hTkx2IkRG3_srWMLfcjxjqwMIp8zd2Dkuw3BDDjGR9n40aNKDM

fd671 := client.Open("/U6jvpoyEz8.txt", client.O_RDWR|client.O_CREATE)
if(fd671 < 0) {
  panic("open failed")
}


ret = client.Close(fd670)
if(ret != 0) {
  panic("close failed")
}


fd672 := client.Open("/M_yRoj1Az0.txt", client.O_RDWR|client.O_CREATE)
if(fd672 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd672, []byte("Qxm17gsGy0I3stdo0XUrbTPLonD9wNsIp5_CRyciNzh1TXirckKt_gfSD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) Qxm17gsGy0I3stdo0XUrbTPLonD9wNsIp5_CRyciNzh1TXirckKt_gfSD

buf, ret = client.Read(fd669, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd669)
if(ret != 0) {
  panic("close failed")
}


fd673 := client.Open("/sZAyeiVvZ5.txt", client.O_RDWR|client.O_CREATE)
if(fd673 < 0) {
  panic("open failed")
}


ret = client.Close(fd673)
if(ret != 0) {
  panic("close failed")
}


fd674 := client.Open("/pgOFTKSQaT.txt", client.O_RDWR|client.O_CREATE)
if(fd674 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd671, []byte("Hh0erg8Hj8Y_vRuCzydSB39XicLpTKRyBY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) Hh0erg8Hj8Y_vRuCzydSB39XicLpTKRyBY

ret = client.Seek(fd671, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


fd675 := client.Open("/s7jifbhgZq.txt", client.O_RDWR|client.O_CREATE)
if(fd675 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd675, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jkplCU3muLebBQqxDt2xa3ZPV9") {
  panic("wrong data returned")
}


ret = client.Seek(fd674, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Close(fd674)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd664, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "mYPF84whorILX9nuNZk") {
  panic("wrong data returned")
}


ret = client.Seek(fd671, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


fd676 := client.Open("/JBxzKGqlST.txt", client.O_RDWR|client.O_CREATE)
if(fd676 < 0) {
  panic("open failed")
}


fd677 := client.Open("/VCgdE3WhbF.txt", client.O_RDWR|client.O_CREATE)
if(fd677 < 0) {
  panic("open failed")
}

//fd state: (57) Qxm17gsGy0I3stdo0XUrbTPLonD9wNsIp5_CRyciNzh1TXirckKt_gfSD

ret = client.Write(fd672, []byte("f3rBg_a6Yh87GC9gMp1_vFMK7gkMG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) Qxm17gsGy0I3stdo0XUrbTPLonD9wNsIp5_CRyciNzh1TXirckKt_gfSDf3rBg_a6Yh87GC9gMp1_vFMK7gkMG

fd678 := client.Open("/tX3hPwMq2Q.txt", client.O_RDWR|client.O_CREATE)
if(fd678 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd664, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd676, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd677, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


fd679 := client.Open("/elTreHWyzI.txt", client.O_RDWR|client.O_CREATE)
if(fd679 < 0) {
  panic("open failed")
}

//fd state: (8) Hh0erg8Hj8Y_vRuCzydSB39XicLpTKRyBY

ret = client.Write(fd671, []byte("37nZyqpy0I26arLOkF1eXYAnOd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) Hh0erg8H37nZyqpy0I26arLOkF1eXYAnOd

fd680 := client.Open("/CO9WO8YU1D.txt", client.O_RDWR|client.O_CREATE)
if(fd680 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd680, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) 

buf, ret = client.Read(fd676, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (26) jkplCU3muLebBQqxDt2xa3ZPV9WIWCUXPIqNJva9uKUplrlaMujplTTBEgCCNytjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlonIQyPSJnjvkqLNq6MSW0_YD2l6Pt80sbmXYrB7IAHOpLoDJpEVtB1zfg5Ih6hglMYPMjaJNoaooSnuOINoNVOKOHz4H75jFsWasBlyzOrGq5b5hmF7tpBjAukES3ZBLX30Pc6Po9wueNNl4PcMwiFWR8j18hY27NGuvJ6z_5c7LUmP8RYgKxx2GR

ret = client.Write(fd675, []byte("RSZsMujMaLZw8mzxPLZzfrjGKwQJYhCbBN6K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) jkplCU3muLebBQqxDt2xa3ZPV9RSZsMujMaLZw8mzxPLZzfrjGKwQJYhCbBN6KtjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlonIQyPSJnjvkqLNq6MSW0_YD2l6Pt80sbmXYrB7IAHOpLoDJpEVtB1zfg5Ih6hglMYPMjaJNoaooSnuOINoNVOKOHz4H75jFsWasBlyzOrGq5b5hmF7tpBjAukES3ZBLX30Pc6Po9wueNNl4PcMwiFWR8j18hY27NGuvJ6z_5c7LUmP8RYgKxx2GR

ret = client.Close(fd676)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd671, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd675, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


fd681 := client.Open("/CpBTPLAg78.txt", client.O_RDWR|client.O_CREATE)
if(fd681 < 0) {
  panic("open failed")
}


ret = client.Seek(fd681, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


ret = client.Seek(fd664, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}


fd682 := client.Open("/q4q9fzYtS_.txt", client.O_RDWR|client.O_CREATE)
if(fd682 < 0) {
  panic("open failed")
}


ret = client.Close(fd681)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd682)
if(ret != 0) {
  panic("close failed")
}

//fd state: (45) jkplCU3muLebBQqxDt2xa3ZPV9RSZsMujMaLZw8mzxPLZzfrjGKwQJYhCbBN6KtjiyggDkPAAwa_IsJfuUs_9Ue17WkmMeuNTxE8uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlonIQyPSJnjvkqLNq6MSW0_YD2l6Pt80sbmXYrB7IAHOpLoDJpEVtB1zfg5Ih6hglMYPMjaJNoaooSnuOINoNVOKOHz4H75jFsWasBlyzOrGq5b5hmF7tpBjAukES3ZBLX30Pc6Po9wueNNl4PcMwiFWR8j18hY27NGuvJ6z_5c7LUmP8RYgKxx2GR

ret = client.Write(fd675, []byte("l3e7zzwUHuE03x4kSh5VRUPZMToHKwPm7eW2BbjEhM5bSL6CJWqWDOi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) jkplCU3muLebBQqxDt2xa3ZPV9RSZsMujMaLZw8mzxPLZl3e7zzwUHuE03x4kSh5VRUPZMToHKwPm7eW2BbjEhM5bSL6CJWqWDOiuZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YPNGbdFXiFg9YFEfYapjqJmSyBlonIQyPSJnjvkqLNq6MSW0_YD2l6Pt80sbmXYrB7IAHOpLoDJpEVtB1zfg5Ih6hglMYPMjaJNoaooSnuOINoNVOKOHz4H75jFsWasBlyzOrGq5b5hmF7tpBjAukES3ZBLX30Pc6Po9wueNNl4PcMwiFWR8j18hY27NGuvJ6z_5c7LUmP8RYgKxx2GR
//fd state: (26) IYca5Ivuhtq4TfPwPVojsU1rHgOBUMuCkOSWuj8C9xbeuNO

ret = client.Write(fd677, []byte("hOEUoI8b_F0GUTS3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (42) IYca5Ivuhtq4TfPwPVojsU1rHghOEUoI8b_F0GUTS3beuNO

buf, ret = client.Read(fd675, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uZlpoMrrmeMWYrv0mlOQRw5TwP2ucCPex1_Wx8dV0cqGZiAoumHu1171uUs9wWzv0YP") {
  panic("wrong data returned")
}


fd683 := client.Open("/aHSi5Rp9bT.txt", client.O_RDWR|client.O_CREATE)
if(fd683 < 0) {
  panic("open failed")
}


ret = client.Close(fd683)
if(ret != 0) {
  panic("close failed")
}


fd684 := client.Open("/m80E198Few.txt", client.O_RDWR|client.O_CREATE)
if(fd684 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd684, []byte("rYRmgqf7CHjsZWt4jqaMVOdCCdrO1z93kYKeryDpLvVhHJ0QzLZsdETLineQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) rYRmgqf7CHjsZWt4jqaMVOdCCdrO1z93kYKeryDpLvVhHJ0QzLZsdETLineQ

ret = client.Seek(fd671, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


ret = client.Close(fd680)
if(ret != 0) {
  panic("close failed")
}

//fd state: (86) Qxm17gsGy0I3stdo0XUrbTPLonD9wNsIp5_CRyciNzh1TXirckKt_gfSDf3rBg_a6Yh87GC9gMp1_vFMK7gkMG

ret = client.Write(fd672, []byte("HrWAsp7ifB5y1x3ZVm2DiuNTJUuG4NyJjxAhTM2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (125) Qxm17gsGy0I3stdo0XUrbTPLonD9wNsIp5_CRyciNzh1TXirckKt_gfSDf3rBg_a6Yh87GC9gMp1_vFMK7gkMGHrWAsp7ifB5y1x3ZVm2DiuNTJUuG4NyJjxAhTM2

ret = client.Close(fd677)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd679, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd678, []byte("f0sJGEGBvzBzkWry6GqnLdkz0I2MQ3yEbrheP_mUTaY257gOdqzXBf186KRAe9pCHujeNq0V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) f0sJGEGBvzBzkWry6GqnLdkz0I2MQ3yEbrheP_mUTaY257gOdqzXBf186KRAe9pCHujeNq0V
//fd state: (72) f0sJGEGBvzBzkWry6GqnLdkz0I2MQ3yEbrheP_mUTaY257gOdqzXBf186KRAe9pCHujeNq0V

ret = client.Write(fd678, []byte("ePlkqwj6GOlz3ZjvNpfS37vNO2jKZFHQNDagXEOvfL5x3wPa5H6WFD_olQuNYg9RboMSIbrpwwjUgtaE_vu4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (156) f0sJGEGBvzBzkWry6GqnLdkz0I2MQ3yEbrheP_mUTaY257gOdqzXBf186KRAe9pCHujeNq0VePlkqwj6GOlz3ZjvNpfS37vNO2jKZFHQNDagXEOvfL5x3wPa5H6WFD_olQuNYg9RboMSIbrpwwjUgtaE_vu4

fd685 := client.Open("/LQ7H9N02Pi.txt", client.O_RDWR|client.O_CREATE)
if(fd685 < 0) {
  panic("open failed")
}


fd686 := client.Open("/3xxC_0CJ4b.txt", client.O_RDWR|client.O_CREATE)
if(fd686 < 0) {
  panic("open failed")
}


ret = client.Close(fd678)
if(ret != 0) {
  panic("close failed")
}


fd687 := client.Open("/bc6Z6rSkIG.txt", client.O_RDWR|client.O_CREATE)
if(fd687 < 0) {
  panic("open failed")
}

//fd state: (24) g3X_JITFj4wEmYPF84whorILX9nuNZk

ret = client.Write(fd664, []byte("wU5vbT8v5NIHYf4kspODIYfLwX3FTqntX9X3osG8nHnhiEEwVbjzSEmSz0_TO9qUN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) g3X_JITFj4wEmYPF84whorILwU5vbT8v5NIHYf4kspODIYfLwX3FTqntX9X3osG8nHnhiEEwVbjzSEmSz0_TO9qUN

ret = client.Seek(fd687, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd685, 1)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd688 := client.Open("/VznqypOi9t.txt", client.O_RDWR|client.O_CREATE)
if(fd688 < 0) {
  panic("open failed")
}

//fd state: (125) Qxm17gsGy0I3stdo0XUrbTPLonD9wNsIp5_CRyciNzh1TXirckKt_gfSDf3rBg_a6Yh87GC9gMp1_vFMK7gkMGHrWAsp7ifB5y1x3ZVm2DiuNTJUuG4NyJjxAhTM2

ret = client.Write(fd672, []byte("sw9PalsUrVu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) Qxm17gsGy0I3stdo0XUrbTPLonD9wNsIp5_CRyciNzh1TXirckKt_gfSDf3rBg_a6Yh87GC9gMp1_vFMK7gkMGHrWAsp7ifB5y1x3ZVm2DiuNTJUuG4NyJjxAhTM2sw9PalsUrVu

ret = client.Close(fd671)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd675)
if(ret != 0) {
  panic("close failed")
}


fd689 := client.Open("/2qm1UAFxsb.txt", client.O_RDWR|client.O_CREATE)
if(fd689 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd684, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd689, []byte("kPbSKY3eUWqUIdnMNcfq0LTxwohOMKkvmABKSvCBVIWcPXIQbjSA967zvJLAPVHZa8h"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (67) kPbSKY3eUWqUIdnMNcfq0LTxwohOMKkvmABKSvCBVIWcPXIQbjSA967zvJLAPVHZa8h

fd690 := client.Open("/dqzKim4aOI.txt", client.O_RDWR|client.O_CREATE)
if(fd690 < 0) {
  panic("open failed")
}


ret = client.Seek(fd684, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}

//fd state: (67) kPbSKY3eUWqUIdnMNcfq0LTxwohOMKkvmABKSvCBVIWcPXIQbjSA967zvJLAPVHZa8h

ret = client.Write(fd689, []byte("Ac9ZQnvhH1J3Bn2PjXnj3L0RS1wZ7aJcC3qmB9Lg0v_izkdJk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) kPbSKY3eUWqUIdnMNcfq0LTxwohOMKkvmABKSvCBVIWcPXIQbjSA967zvJLAPVHZa8hAc9ZQnvhH1J3Bn2PjXnj3L0RS1wZ7aJcC3qmB9Lg0v_izkdJk

ret = client.Seek(fd679, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd672, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) zQ0UbVvXxQaEWwpJXrJ7YBzBZAgyRE4_i81ZL_2rosJe7BfUgvgKEcbElvSdPLf_hqHUI0OeVxUoh

ret = client.Write(fd688, []byte("qM2YRCpsaQbhAWq4moksG58gzjeIJVIMOzhDYYZPMTAkAB417F1H94eAzl9eZhaHxcdM3S36Ev4S0K"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) qM2YRCpsaQbhAWq4moksG58gzjeIJVIMOzhDYYZPMTAkAB417F1H94eAzl9eZhaHxcdM3S36Ev4S0K

ret = client.Seek(fd684, 52, client.SEEK_SET)
if(ret != 52) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 52)
  panic("seek failed")
}


buf, ret = client.Read(fd685, 51)
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


buf, ret = client.Read(fd664, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd684, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


fd691 := client.Open("/TzzDnGQvJB.txt", client.O_RDWR|client.O_CREATE)
if(fd691 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd686, 64)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd686)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd664, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd687, []byte("VGRXiTiRDs_zTiYcGpVgddsxGX4VUhUKNiVMpop4PKpJPcpXK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) VGRXiTiRDs_zTiYcGpVgddsxGX4VUhUKNiVMpop4PKpJPcpXK

ret = client.Seek(fd664, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


ret = client.Seek(fd672, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}

//fd state: (49) VGRXiTiRDs_zTiYcGpVgddsxGX4VUhUKNiVMpop4PKpJPcpXK

ret = client.Write(fd687, []byte("ZWqPCsSzOiwsjtzb7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) VGRXiTiRDs_zTiYcGpVgddsxGX4VUhUKNiVMpop4PKpJPcpXKZWqPCsSzOiwsjtzb7

fd692 := client.Open("/baO6iJtSGv.txt", client.O_RDWR|client.O_CREATE)
if(fd692 < 0) {
  panic("open failed")
}


fd693 := client.Open("/WvrUAZ5kOA.txt", client.O_RDWR|client.O_CREATE)
if(fd693 < 0) {
  panic("open failed")
}

//fd state: (25) g3X_JITFj4wEmYPF84whorILwU5vbT8v5NIHYf4kspODIYfLwX3FTqntX9X3osG8nHnhiEEwVbjzSEmSz0_TO9qUN

ret = client.Write(fd664, []byte("3yN2YBfDG6FQtmKHepFffah9fB1XymaklmVxJvkbShZACyDm6pC5FHlew_ZIEokrUoljAYnqcJCCZ3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) g3X_JITFj4wEmYPF84whorILw3yN2YBfDG6FQtmKHepFffah9fB1XymaklmVxJvkbShZACyDm6pC5FHlew_ZIEokrUoljAYnqcJCCZ3
//fd state: (0) 

ret = client.Write(fd691, []byte("dwY9G54cN6X_7tjm7CXjdpGl9K0O3GZMejUUvg6i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) dwY9G54cN6X_7tjm7CXjdpGl9K0O3GZMejUUvg6i

ret = client.Close(fd684)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd690, []byte("fYc_e9nynKAI_Qbs30iGKpLRVjURJ0mFUnMEPw1GMc4uBRomcFY7eT_K4xncjvzM8uiefx1rwJCDuKsTWnG0N"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) fYc_e9nynKAI_Qbs30iGKpLRVjURJ0mFUnMEPw1GMc4uBRomcFY7eT_K4xncjvzM8uiefx1rwJCDuKsTWnG0N

ret = client.Close(fd690)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd691)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd664, 95, client.SEEK_SET)
if(ret != 95) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 95)
  panic("seek failed")
}


fd694 := client.Open("/7VYdgfQyUf.txt", client.O_RDWR|client.O_CREATE)
if(fd694 < 0) {
  panic("open failed")
}


fd695 := client.Open("/N6wbOW8u8N.txt", client.O_RDWR|client.O_CREATE)
if(fd695 < 0) {
  panic("open failed")
}


fd696 := client.Open("/e3GmYXSUks.txt", client.O_RDWR|client.O_CREATE)
if(fd696 < 0) {
  panic("open failed")
}


ret = client.Close(fd695)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd692, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Close(fd672)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd687, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}


buf, ret = client.Read(fd688, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd685, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd697 := client.Open("/Zr7CSIZlSr.txt", client.O_RDWR|client.O_CREATE)
if(fd697 < 0) {
  panic("open failed")
}


ret = client.Seek(fd692, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}


ret = client.Seek(fd664, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


fd698 := client.Open("/3kZ5zPN0Di.txt", client.O_RDWR|client.O_CREATE)
if(fd698 < 0) {
  panic("open failed")
}

//fd state: (78) qM2YRCpsaQbhAWq4moksG58gzjeIJVIMOzhDYYZPMTAkAB417F1H94eAzl9eZhaHxcdM3S36Ev4S0K

ret = client.Write(fd688, []byte("3Ye_r3FT5TvdpxRbPaapZQ88Ybr5MYkmqOYv4g9Zd1LhbKbrPgpnPPvJ0ZfFyDcOI2QuYmol7h_nK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (155) qM2YRCpsaQbhAWq4moksG58gzjeIJVIMOzhDYYZPMTAkAB417F1H94eAzl9eZhaHxcdM3S36Ev4S0K3Ye_r3FT5TvdpxRbPaapZQ88Ybr5MYkmqOYv4g9Zd1LhbKbrPgpnPPvJ0ZfFyDcOI2QuYmol7h_nK

ret = client.Close(fd692)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd664)
if(ret != 0) {
  panic("close failed")
}


fd699 := client.Open("/PWtXLMuKMR.txt", client.O_RDWR|client.O_CREATE)
if(fd699 < 0) {
  panic("open failed")
}


ret = client.Seek(fd694, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd696, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd700 := client.Open("/6zJp0s_4j6.txt", client.O_RDWR|client.O_CREATE)
if(fd700 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd687, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GRXiTiRDs_zTiYcGpVgddsxGX4VUhUKNiVMpop4PKpJPcpXKZWqPCsSzOiwsjtzb7") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd696, []byte("3xsZJrE7Y0cY2Wzr8JT_vq2qY7a7kSyvM9qjN5AyuwflnFgifWVUYkU3mejpzBZYlAKRuOMLOALH_6UMzp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) 3xsZJrE7Y0cY2Wzr8JT_vq2qY7a7kSyvM9qjN5AyuwflnFgifWVUYkU3mejpzBZYlAKRuOMLOALH_6UMzp

ret = client.Seek(fd698, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd698)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd696, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}


buf, ret = client.Read(fd687, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) T89gG2RhDkx9yNYRwwRoMfcrtRNX8Melrfyrq0iVFZKIFGdHR8

ret = client.Write(fd699, []byte("jbpUaUrWA1FheViOVusRp_HHvcrHpJr7EZhEcsZwoGvG0VsgeR1aPoSc0NlmeqECgSCCgLkH3HbuJS7w"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) jbpUaUrWA1FheViOVusRp_HHvcrHpJr7EZhEcsZwoGvG0VsgeR1aPoSc0NlmeqECgSCCgLkH3HbuJS7w

buf, ret = client.Read(fd689, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd700)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd688, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd701 := client.Open("/gcbBJzu1Aq.txt", client.O_RDWR|client.O_CREATE)
if(fd701 < 0) {
  panic("open failed")
}


fd702 := client.Open("/gcbBJzu1Aq.txt", client.O_RDWR|client.O_CREATE)
if(fd702 < 0) {
  panic("open failed")
}


ret = client.Close(fd702)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd688, 2)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) FeurBNM8PEQkTKAfm3Voc3cYk

ret = client.Write(fd697, []byte("UVIyOBgDDHaSFooc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (16) UVIyOBgDDHaSFoocm3Voc3cYk

fd703 := client.Open("/ghuzLxuO9g.txt", client.O_RDWR|client.O_CREATE)
if(fd703 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd687, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd687)
if(ret != 0) {
  panic("close failed")
}


fd704 := client.Open("/L68IVOUnGL.txt", client.O_RDWR|client.O_CREATE)
if(fd704 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd685, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd694, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd704, []byte("LIs2ETL8TFfgpyXc9rU4JfgBCo2nm2k"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) LIs2ETL8TFfgpyXc9rU4JfgBCo2nm2k

buf, ret = client.Read(fd694, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd699)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd693)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd701, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (31) LIs2ETL8TFfgpyXc9rU4JfgBCo2nm2k

ret = client.Write(fd704, []byte("C44is1XzE2n2gqcbkU4vipsVSP6TXjlywUaJRvJGmFGzqQGj2egQGl6uAa2tv6nEUKwzG0hFzFO77y05Y9ssWjOdBgW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) LIs2ETL8TFfgpyXc9rU4JfgBCo2nm2kC44is1XzE2n2gqcbkU4vipsVSP6TXjlywUaJRvJGmFGzqQGj2egQGl6uAa2tv6nEUKwzG0hFzFO77y05Y9ssWjOdBgW

ret = client.Seek(fd701, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd705 := client.Open("/0t_D6cCpLL.txt", client.O_RDWR|client.O_CREATE)
if(fd705 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd685, 27)
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


buf, ret = client.Read(fd694, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (116) kPbSKY3eUWqUIdnMNcfq0LTxwohOMKkvmABKSvCBVIWcPXIQbjSA967zvJLAPVHZa8hAc9ZQnvhH1J3Bn2PjXnj3L0RS1wZ7aJcC3qmB9Lg0v_izkdJk

ret = client.Write(fd689, []byte("lMgSyJXQvaDKuJ1QRNMvcs1HOy4oc6kfpUrPxujOBO1i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (160) kPbSKY3eUWqUIdnMNcfq0LTxwohOMKkvmABKSvCBVIWcPXIQbjSA967zvJLAPVHZa8hAc9ZQnvhH1J3Bn2PjXnj3L0RS1wZ7aJcC3qmB9Lg0v_izkdJklMgSyJXQvaDKuJ1QRNMvcs1HOy4oc6kfpUrPxujOBO1i

fd706 := client.Open("/Ti8vk7aoMw.txt", client.O_RDWR|client.O_CREATE)
if(fd706 < 0) {
  panic("open failed")
}


ret = client.Close(fd697)
if(ret != 0) {
  panic("close failed")
}


fd707 := client.Open("/BMECz1_7Wb.txt", client.O_RDWR|client.O_CREATE)
if(fd707 < 0) {
  panic("open failed")
}


fd708 := client.Open("/kKAM8FZ2Mx.txt", client.O_RDWR|client.O_CREATE)
if(fd708 < 0) {
  panic("open failed")
}


ret = client.Close(fd696)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd707)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd706)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd689, 120, client.SEEK_SET)
if(ret != 120) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 120)
  panic("seek failed")
}


buf, ret = client.Read(fd689, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yJXQvaDKuJ1QRNMvcs1HOy4oc6kfpUrPxujOBO1i") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd688, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd708)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd689, 142, client.SEEK_SET)
if(ret != 142) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 142)
  panic("seek failed")
}


fd709 := client.Open("/d8bf98ibB6.txt", client.O_RDWR|client.O_CREATE)
if(fd709 < 0) {
  panic("open failed")
}


ret = client.Seek(fd703, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd710 := client.Open("/dw3wX6Mc3X.txt", client.O_RDWR|client.O_CREATE)
if(fd710 < 0) {
  panic("open failed")
}


ret = client.Seek(fd685, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd710)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd703)
if(ret != 0) {
  panic("close failed")
}


fd711 := client.Open("/JyCJqcy6n8.txt", client.O_RDWR|client.O_CREATE)
if(fd711 < 0) {
  panic("open failed")
}


ret = client.Close(fd685)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd688)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd689, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4oc6kfpUrPxujOBO1i") {
  panic("wrong data returned")
}


ret = client.Close(fd701)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd709, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (122) LIs2ETL8TFfgpyXc9rU4JfgBCo2nm2kC44is1XzE2n2gqcbkU4vipsVSP6TXjlywUaJRvJGmFGzqQGj2egQGl6uAa2tv6nEUKwzG0hFzFO77y05Y9ssWjOdBgW

ret = client.Write(fd704, []byte("7nXrMYXNCfCZ9FMgMp_jz35ppYe_Xu_sqlYknURIsyqGABi8gOMefMFo3TkBMpfCvAttRSZiRjTxzyAAL4GKwuso04h4Vfe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (217) LIs2ETL8TFfgpyXc9rU4JfgBCo2nm2kC44is1XzE2n2gqcbkU4vipsVSP6TXjlywUaJRvJGmFGzqQGj2egQGl6uAa2tv6nEUKwzG0hFzFO77y05Y9ssWjOdBgW7nXrMYXNCfCZ9FMgMp_jz35ppYe_Xu_sqlYknURIsyqGABi8gOMefMFo3TkBMpfCvAttRSZiRjTxzyAAL4GKwuso04h4Vfe

ret = client.Seek(fd704, 206, client.SEEK_SET)
if(ret != 206) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 206)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd709, []byte("_HHX_JCctTXH78qY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (16) _HHX_JCctTXH78qY

fd712 := client.Open("/BvTYAUZthR.txt", client.O_RDWR|client.O_CREATE)
if(fd712 < 0) {
  panic("open failed")
}

//fd state: (16) _HHX_JCctTXH78qY

ret = client.Write(fd709, []byte("QoVhp9eryskMrFDVAq0x5R2TNeM6426LJCOaXzpIZB_3Xj5jEgps3LYEOm0NOSiuySTNNvFpLrbEVoRmnMk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) _HHX_JCctTXH78qYQoVhp9eryskMrFDVAq0x5R2TNeM6426LJCOaXzpIZB_3Xj5jEgps3LYEOm0NOSiuySTNNvFpLrbEVoRmnMk

ret = client.Close(fd694)
if(ret != 0) {
  panic("close failed")
}

//fd state: (160) kPbSKY3eUWqUIdnMNcfq0LTxwohOMKkvmABKSvCBVIWcPXIQbjSA967zvJLAPVHZa8hAc9ZQnvhH1J3Bn2PjXnj3L0RS1wZ7aJcC3qmB9Lg0v_izkdJklMgSyJXQvaDKuJ1QRNMvcs1HOy4oc6kfpUrPxujOBO1i

ret = client.Write(fd689, []byte("rvYYqO6W00bYinhcbPQiKarTK6jf6dAQmcpGAPTgp93Yyrke7krP8R2TcA1t"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (220) kPbSKY3eUWqUIdnMNcfq0LTxwohOMKkvmABKSvCBVIWcPXIQbjSA967zvJLAPVHZa8hAc9ZQnvhH1J3Bn2PjXnj3L0RS1wZ7aJcC3qmB9Lg0v_izkdJklMgSyJXQvaDKuJ1QRNMvcs1HOy4oc6kfpUrPxujOBO1irvYYqO6W00bYinhcbPQiKarTK6jf6dAQmcpGAPTgp93Yyrke7krP8R2TcA1t

fd713 := client.Open("/6iBimSOOAD.txt", client.O_RDWR|client.O_CREATE)
if(fd713 < 0) {
  panic("open failed")
}

//fd state: (220) kPbSKY3eUWqUIdnMNcfq0LTxwohOMKkvmABKSvCBVIWcPXIQbjSA967zvJLAPVHZa8hAc9ZQnvhH1J3Bn2PjXnj3L0RS1wZ7aJcC3qmB9Lg0v_izkdJklMgSyJXQvaDKuJ1QRNMvcs1HOy4oc6kfpUrPxujOBO1irvYYqO6W00bYinhcbPQiKarTK6jf6dAQmcpGAPTgp93Yyrke7krP8R2TcA1t

ret = client.Write(fd689, []byte("8znWTUd_SdcjneK6fYtQCWx9NQCiWB8plDNiiyUUvZdbw2UmrrUX1bBioGUmQUGhhANxWq5sDxRlrhMf1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (301) kPbSKY3eUWqUIdnMNcfq0LTxwohOMKkvmABKSvCBVIWcPXIQbjSA967zvJLAPVHZa8hAc9ZQnvhH1J3Bn2PjXnj3L0RS1wZ7aJcC3qmB9Lg0v_izkdJklMgSyJXQvaDKuJ1QRNMvcs1HOy4oc6kfpUrPxujOBO1irvYYqO6W00bYinhcbPQiKarTK6jf6dAQmcpGAPTgp93Yyrke7krP8R2TcA1t8znWTUd_SdcjneK6fYtQCWx9NQCiWB8plDNiiyUUvZdbw2UmrrUX1bBioGUmQUGhhANxWq5sDxRlrhMf1
//fd state: (0) 

ret = client.Write(fd711, []byte("sKiTJp4lmRe8L0zpkMZ9FrAfDNSkHhq5bEfmBlH_D0nSdWCKkr4sQmhSDQzkGi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) sKiTJp4lmRe8L0zpkMZ9FrAfDNSkHhq5bEfmBlH_D0nSdWCKkr4sQmhSDQzkGi

ret = client.Seek(fd713, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd689)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd712, []byte("trW9AWBuURfZ1TEOfMxeOOl6Xxlv4XUGse"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (34) trW9AWBuURfZ1TEOfMxeOOl6Xxlv4XUGse

fd714 := client.Open("/FPPNgeMK1l.txt", client.O_RDWR|client.O_CREATE)
if(fd714 < 0) {
  panic("open failed")
}


fd715 := client.Open("/GO3lDUGzDz.txt", client.O_RDWR|client.O_CREATE)
if(fd715 < 0) {
  panic("open failed")
}


fd716 := client.Open("/WxPQuN5oUi.txt", client.O_RDWR|client.O_CREATE)
if(fd716 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd704, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wuso04h4Vfe") {
  panic("wrong data returned")
}


ret = client.Close(fd711)
if(ret != 0) {
  panic("close failed")
}


fd717 := client.Open("/ukzm_5Sa9e.txt", client.O_RDWR|client.O_CREATE)
if(fd717 < 0) {
  panic("open failed")
}


fd718 := client.Open("/kKAM8FZ2Mx.txt", client.O_RDWR|client.O_CREATE)
if(fd718 < 0) {
  panic("open failed")
}


fd719 := client.Open("/Wn_LlxApk0.txt", client.O_RDWR|client.O_CREATE)
if(fd719 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd718, []byte("7fkhcM9H"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (8) 7fkhcM9H

ret = client.Seek(fd704, 102, client.SEEK_SET)
if(ret != 102) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 102)
  panic("seek failed")
}


ret = client.Close(fd719)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd717)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd709)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd713, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd720 := client.Open("/FZZ8y6anbP.txt", client.O_RDWR|client.O_CREATE)
if(fd720 < 0) {
  panic("open failed")
}


ret = client.Close(fd715)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd714)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd720)
if(ret != 0) {
  panic("close failed")
}

//fd state: (8) 7fkhcM9H

ret = client.Write(fd718, []byte("wADws7DLpR2_q2oharOA6DAaEvssVYVC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) 7fkhcM9HwADws7DLpR2_q2oharOA6DAaEvssVYVC
//fd state: (34) trW9AWBuURfZ1TEOfMxeOOl6Xxlv4XUGse

ret = client.Write(fd712, []byte("xhA0kQI9AI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) trW9AWBuURfZ1TEOfMxeOOl6Xxlv4XUGsexhA0kQI9AI
//fd state: (44) trW9AWBuURfZ1TEOfMxeOOl6Xxlv4XUGsexhA0kQI9AI

ret = client.Write(fd712, []byte("aNH4Up4_CuGoJWZlUYn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) trW9AWBuURfZ1TEOfMxeOOl6Xxlv4XUGsexhA0kQI9AIaNH4Up4_CuGoJWZlUYn

ret = client.Seek(fd718, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


fd721 := client.Open("/N6wbOW8u8N.txt", client.O_RDWR|client.O_CREATE)
if(fd721 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd721, []byte("pwQC5FFMIs2NShg9HhAK69RYOZs3GNmlYc3oZUXyB4Hte1ivXrjQa"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) pwQC5FFMIs2NShg9HhAK69RYOZs3GNmlYc3oZUXyB4Hte1ivXrjQa

fd722 := client.Open("/Yt5ysR3lnX.txt", client.O_RDWR|client.O_CREATE)
if(fd722 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd713, []byte("yr04WD68hSZ33E3h9ktsmxp8RGNWpJWRt_nf3cAxWSGN5L51h99MRje46s1UV87j91cX3ar1GjNWIVovtK16_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) yr04WD68hSZ33E3h9ktsmxp8RGNWpJWRt_nf3cAxWSGN5L51h99MRje46s1UV87j91cX3ar1GjNWIVovtK16_

fd723 := client.Open("/V2oSgaa_77.txt", client.O_RDWR|client.O_CREATE)
if(fd723 < 0) {
  panic("open failed")
}


ret = client.Seek(fd722, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd721, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd724 := client.Open("/lEKHoBp1sO.txt", client.O_RDWR|client.O_CREATE)
if(fd724 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (102) LIs2ETL8TFfgpyXc9rU4JfgBCo2nm2kC44is1XzE2n2gqcbkU4vipsVSP6TXjlywUaJRvJGmFGzqQGj2egQGl6uAa2tv6nEUKwzG0hFzFO77y05Y9ssWjOdBgW7nXrMYXNCfCZ9FMgMp_jz35ppYe_Xu_sqlYknURIsyqGABi8gOMefMFo3TkBMpfCvAttRSZiRjTxzyAAL4GKwuso04h4Vfe

ret = client.Write(fd704, []byte("L39PLouIkOMutZls__0iNkoxZE3aWriHF0I"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (137) LIs2ETL8TFfgpyXc9rU4JfgBCo2nm2kC44is1XzE2n2gqcbkU4vipsVSP6TXjlywUaJRvJGmFGzqQGj2egQGl6uAa2tv6nEUKwzG0hL39PLouIkOMutZls__0iNkoxZE3aWriHF0IgMp_jz35ppYe_Xu_sqlYknURIsyqGABi8gOMefMFo3TkBMpfCvAttRSZiRjTxzyAAL4GKwuso04h4Vfe

fd725 := client.Open("/cCJFYITFef.txt", client.O_RDWR|client.O_CREATE)
if(fd725 < 0) {
  panic("open failed")
}


ret = client.Seek(fd722, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd718, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Seek(fd723, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd712, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Close(fd704)
if(ret != 0) {
  panic("close failed")
}

//fd state: (53) pwQC5FFMIs2NShg9HhAK69RYOZs3GNmlYc3oZUXyB4Hte1ivXrjQa

ret = client.Write(fd721, []byte("qiy8VD68iNlG7YsdCcBp1mEUeWolE_lrQ8M6HYBmwCqwgxG_W6w0yQuZ_3vPzjBnt8ceMuOe2OOtF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) pwQC5FFMIs2NShg9HhAK69RYOZs3GNmlYc3oZUXyB4Hte1ivXrjQaqiy8VD68iNlG7YsdCcBp1mEUeWolE_lrQ8M6HYBmwCqwgxG_W6w0yQuZ_3vPzjBnt8ceMuOe2OOtF

ret = client.Close(fd712)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd723, []byte("xCcAXavwS5pE0yD23H6Xfr4s3xBzXYN4X8qNUdO_DvFP1X4xdtOkAgKCd4rUUaIDvjtFR6YBlc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (74) xCcAXavwS5pE0yD23H6Xfr4s3xBzXYN4X8qNUdO_DvFP1X4xdtOkAgKCd4rUUaIDvjtFR6YBlc

ret = client.Close(fd718)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd724, []byte("XRuOrX7baFpUVtaIZ2LwFf2j6WuQoWdyXphcQM_WmSn8ggGS3Qx4_o30Chtth"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) XRuOrX7baFpUVtaIZ2LwFf2j6WuQoWdyXphcQM_WmSn8ggGS3Qx4_o30Chtth

buf, ret = client.Read(fd722, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd724, 67)
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


buf, ret = client.Read(fd725, 87)
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


ret = client.Seek(fd725, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (61) XRuOrX7baFpUVtaIZ2LwFf2j6WuQoWdyXphcQM_WmSn8ggGS3Qx4_o30Chtth

ret = client.Write(fd724, []byte("oxkUnJjxkT1JXeJd1HbF16lA5IPLjN2p6VvwFMeLFWVSURfsb70adSdQlp_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) XRuOrX7baFpUVtaIZ2LwFf2j6WuQoWdyXphcQM_WmSn8ggGS3Qx4_o30ChtthoxkUnJjxkT1JXeJd1HbF16lA5IPLjN2p6VvwFMeLFWVSURfsb70adSdQlp_

ret = client.Seek(fd723, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


ret = client.Seek(fd716, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd721)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd724, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


fd726 := client.Open("/mCw4x9xWLU.txt", client.O_RDWR|client.O_CREATE)
if(fd726 < 0) {
  panic("open failed")
}


ret = client.Close(fd726)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd716, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (25) xCcAXavwS5pE0yD23H6Xfr4s3xBzXYN4X8qNUdO_DvFP1X4xdtOkAgKCd4rUUaIDvjtFR6YBlc

ret = client.Write(fd723, []byte("fZkLX9cNl48jiYmRmgKzyAJuqHZu_2yujBSbir0E9lla6CE5nHIifp6ScNiwE2LIbpcCGjMY48uA1XNFbPqLsswqiRgD5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) xCcAXavwS5pE0yD23H6Xfr4s3fZkLX9cNl48jiYmRmgKzyAJuqHZu_2yujBSbir0E9lla6CE5nHIifp6ScNiwE2LIbpcCGjMY48uA1XNFbPqLsswqiRgD5
//fd state: (16) XRuOrX7baFpUVtaIZ2LwFf2j6WuQoWdyXphcQM_WmSn8ggGS3Qx4_o30ChtthoxkUnJjxkT1JXeJd1HbF16lA5IPLjN2p6VvwFMeLFWVSURfsb70adSdQlp_

ret = client.Write(fd724, []byte("hoM6RODxfgsFTJFk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) XRuOrX7baFpUVtaIhoM6RODxfgsFTJFkXphcQM_WmSn8ggGS3Qx4_o30ChtthoxkUnJjxkT1JXeJd1HbF16lA5IPLjN2p6VvwFMeLFWVSURfsb70adSdQlp_

ret = client.Seek(fd724, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


ret = client.Seek(fd724, 30, client.SEEK_SET)
if(ret != 30) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 30)
  panic("seek failed")
}


fd727 := client.Open("/L68IVOUnGL.txt", client.O_RDWR|client.O_CREATE)
if(fd727 < 0) {
  panic("open failed")
}


ret = client.Close(fd724)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd725, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd728 := client.Open("/hof9DnQvTf.txt", client.O_RDWR|client.O_CREATE)
if(fd728 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd725, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd728)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd727)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd716, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd723)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd716, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd725, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd716, []byte("saPlCq53CyUOcJHmhQEM8Bs77EcECn9fn3WhBqO3sEeZUuS6_R9vhaSg0gxdLc8Hj_801qPr21brHaMFt5dJYMAGGP0cyzmaoPU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) saPlCq53CyUOcJHmhQEM8Bs77EcECn9fn3WhBqO3sEeZUuS6_R9vhaSg0gxdLc8Hj_801qPr21brHaMFt5dJYMAGGP0cyzmaoPU

ret = client.Seek(fd725, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd729 := client.Open("/lEoNAXEQIT.txt", client.O_RDWR|client.O_CREATE)
if(fd729 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd729, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd725)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd716)
if(ret != 0) {
  panic("close failed")
}


fd730 := client.Open("/6iBimSOOAD.txt", client.O_RDWR|client.O_CREATE)
if(fd730 < 0) {
  panic("open failed")
}


fd731 := client.Open("/ccZLy21HpC.txt", client.O_RDWR|client.O_CREATE)
if(fd731 < 0) {
  panic("open failed")
}

//fd state: (0) yr04WD68hSZ33E3h9ktsmxp8RGNWpJWRt_nf3cAxWSGN5L51h99MRje46s1UV87j91cX3ar1GjNWIVovtK16_

ret = client.Write(fd730, []byte("UJA7UvqJLrZKorg1C_1wr8pV0GAkRZbYtfjz1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) UJA7UvqJLrZKorg1C_1wr8pV0GAkRZbYtfjz1cAxWSGN5L51h99MRje46s1UV87j91cX3ar1GjNWIVovtK16_

fd732 := client.Open("/zeQqJNOdYy.txt", client.O_RDWR|client.O_CREATE)
if(fd732 < 0) {
  panic("open failed")
}


fd733 := client.Open("/phbm4OEGKb.txt", client.O_RDWR|client.O_CREATE)
if(fd733 < 0) {
  panic("open failed")
}


ret = client.Close(fd732)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd729)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd731, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}


buf, ret = client.Read(fd731, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "4J2JU4H9ZD7OssdAKlyzSyT3Q1UfAesyM9S2") {
  panic("wrong data returned")
}


ret = client.Close(fd730)
if(ret != 0) {
  panic("close failed")
}

//fd state: (43) erYkCn44J2JU4H9ZD7OssdAKlyzSyT3Q1UfAesyM9S2xUpw44Nt5Lnpz35c1W_l05E

ret = client.Write(fd731, []byte("PAc8hRRxH7deOET43xXnQtGpAPaNiKMzo0JPuYSQbytap3GEoRyOnIa23QLXAOsu_mDZWtmTyKwIQnRhOW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (125) erYkCn44J2JU4H9ZD7OssdAKlyzSyT3Q1UfAesyM9S2PAc8hRRxH7deOET43xXnQtGpAPaNiKMzo0JPuYSQbytap3GEoRyOnIa23QLXAOsu_mDZWtmTyKwIQnRhOW

ret = client.Close(fd733)
if(ret != 0) {
  panic("close failed")
}


fd734 := client.Open("/qQGpJOi0Ag.txt", client.O_RDWR|client.O_CREATE)
if(fd734 < 0) {
  panic("open failed")
}


ret = client.Seek(fd731, 119, client.SEEK_SET)
if(ret != 119) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 119)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd734, []byte("fcCcF9Gnp4Rk5qqmIYwIiJ3qOlu4DCoOwvDJR1t6gx1gXHOlgVhZq7ncm7teCWWBqAI3xiP2KC8tRR8wuu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) fcCcF9Gnp4Rk5qqmIYwIiJ3qOlu4DCoOwvDJR1t6gx1gXHOlgVhZq7ncm7teCWWBqAI3xiP2KC8tRR8wuu

buf, ret = client.Read(fd731, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "QnRhOW") {
  panic("wrong data returned")
}

//fd state: (125) erYkCn44J2JU4H9ZD7OssdAKlyzSyT3Q1UfAesyM9S2PAc8hRRxH7deOET43xXnQtGpAPaNiKMzo0JPuYSQbytap3GEoRyOnIa23QLXAOsu_mDZWtmTyKwIQnRhOW

ret = client.Write(fd731, []byte("7KgR3hJgji2O72huCDftJFErh378d3QReFQUBzrx2c7gKg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (171) erYkCn44J2JU4H9ZD7OssdAKlyzSyT3Q1UfAesyM9S2PAc8hRRxH7deOET43xXnQtGpAPaNiKMzo0JPuYSQbytap3GEoRyOnIa23QLXAOsu_mDZWtmTyKwIQnRhOW7KgR3hJgji2O72huCDftJFErh378d3QReFQUBzrx2c7gKg

ret = client.Seek(fd731, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd735 := client.Open("/2g8nvSsXFZ.txt", client.O_RDWR|client.O_CREATE)
if(fd735 < 0) {
  panic("open failed")
}


ret = client.Close(fd734)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd735, 71)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOJTrqdreNuZLXNgCQ_akukCxfdKIMWKyVR") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd735, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TnvaS") {
  panic("wrong data returned")
}


ret = client.Close(fd731)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd735, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


ret = client.Seek(fd735, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}


ret = client.Seek(fd735, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}

//fd state: (38) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOJTrqdreNuZLXNgCQ_akukCxfdKIMWKyVRTnvaS

ret = client.Write(fd735, []byte("IkF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOIkFqdreNuZLXNgCQ_akukCxfdKIMWKyVRTnvaS
//fd state: (41) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOIkFqdreNuZLXNgCQ_akukCxfdKIMWKyVRTnvaS

ret = client.Write(fd735, []byte("J7QAFKXKZIeMm_V3wPQgvEEXKPYaNLWz5mIOB6J6l1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOIkFJ7QAFKXKZIeMm_V3wPQgvEEXKPYaNLWz5mIOB6J6l1
//fd state: (83) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOIkFJ7QAFKXKZIeMm_V3wPQgvEEXKPYaNLWz5mIOB6J6l1

ret = client.Write(fd735, []byte("PA7x_CKgZdfqC4fYfbIAU18wafCzu0eUGV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOIkFJ7QAFKXKZIeMm_V3wPQgvEEXKPYaNLWz5mIOB6J6l1PA7x_CKgZdfqC4fYfbIAU18wafCzu0eUGV
//fd state: (117) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOIkFJ7QAFKXKZIeMm_V3wPQgvEEXKPYaNLWz5mIOB6J6l1PA7x_CKgZdfqC4fYfbIAU18wafCzu0eUGV

ret = client.Write(fd735, []byte("kBafBroa86iCoW8pIse9ZSrF9NhK5GhrERCFASY9JbmJDOwixNfNwPrecg9ezxiQ18icpyhVzOl9Oa7gP3TqWuS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (204) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOIkFJ7QAFKXKZIeMm_V3wPQgvEEXKPYaNLWz5mIOB6J6l1PA7x_CKgZdfqC4fYfbIAU18wafCzu0eUGVkBafBroa86iCoW8pIse9ZSrF9NhK5GhrERCFASY9JbmJDOwixNfNwPrecg9ezxiQ18icpyhVzOl9Oa7gP3TqWuS
//fd state: (204) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOIkFJ7QAFKXKZIeMm_V3wPQgvEEXKPYaNLWz5mIOB6J6l1PA7x_CKgZdfqC4fYfbIAU18wafCzu0eUGVkBafBroa86iCoW8pIse9ZSrF9NhK5GhrERCFASY9JbmJDOwixNfNwPrecg9ezxiQ18icpyhVzOl9Oa7gP3TqWuS

ret = client.Write(fd735, []byte("v_697SIDrGdAgqIT57d0yclmUeE91C"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (234) w41O_9WZzABoMPDuGVNn3e7RgeBWZQST1oHXWOIkFJ7QAFKXKZIeMm_V3wPQgvEEXKPYaNLWz5mIOB6J6l1PA7x_CKgZdfqC4fYfbIAU18wafCzu0eUGVkBafBroa86iCoW8pIse9ZSrF9NhK5GhrERCFASY9JbmJDOwixNfNwPrecg9ezxiQ18icpyhVzOl9Oa7gP3TqWuSv_697SIDrGdAgqIT57d0yclmUeE91C

buf, ret = client.Read(fd735, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd735, 75, client.SEEK_SET)
if(ret != 75) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 75)
  panic("seek failed")
}


fd736 := client.Open("/GO3lDUGzDz.txt", client.O_RDWR|client.O_CREATE)
if(fd736 < 0) {
  panic("open failed")
}


ret = client.Seek(fd735, 143, client.SEEK_SET)
if(ret != 143) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 143)
  panic("seek failed")
}


ret = client.Close(fd735)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd736, []byte("V7HI8TRadG_urYXps539QTAEnrFloAoKPwFJjWwAXTp_3DNMY2nHGfz74QHj74GzMSsO5EIZ8TmZ3e_Iit22EI0yN_yudaDN9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) V7HI8TRadG_urYXps539QTAEnrFloAoKPwFJjWwAXTp_3DNMY2nHGfz74QHj74GzMSsO5EIZ8TmZ3e_Iit22EI0yN_yudaDN9

ret = client.Seek(fd736, 94, client.SEEK_SET)
if(ret != 94) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 94)
  panic("seek failed")
}


ret = client.Seek(fd736, 70, client.SEEK_SET)
if(ret != 70) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 70)
  panic("seek failed")
}


ret = client.Seek(fd736, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


fd737 := client.Open("/WZrIgODs1e.txt", client.O_RDWR|client.O_CREATE)
if(fd737 < 0) {
  panic("open failed")
}


ret = client.Close(fd737)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd736, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TAEnrFloAoKPwFJjWwAXTp_3DNMY2nH") {
  panic("wrong data returned")
}


fd738 := client.Open("/nPCl6unWmR.txt", client.O_RDWR|client.O_CREATE)
if(fd738 < 0) {
  panic("open failed")
}

//fd state: (52) V7HI8TRadG_urYXps539QTAEnrFloAoKPwFJjWwAXTp_3DNMY2nHGfz74QHj74GzMSsO5EIZ8TmZ3e_Iit22EI0yN_yudaDN9

ret = client.Write(fd736, []byte("ONzlExiD0xje61sR4n_252MG3HTNCHM1ulmv0eEwuRpalyi3YeNgTyCQ7ZzxJvy105JfnnQe1uPg_vRFk94"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) V7HI8TRadG_urYXps539QTAEnrFloAoKPwFJjWwAXTp_3DNMY2nHONzlExiD0xje61sR4n_252MG3HTNCHM1ulmv0eEwuRpalyi3YeNgTyCQ7ZzxJvy105JfnnQe1uPg_vRFk94

ret = client.Seek(fd736, 83, client.SEEK_SET)
if(ret != 83) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 83)
  panic("seek failed")
}


ret = client.Seek(fd738, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd739 := client.Open("/tnVmsNSjJy.txt", client.O_RDWR|client.O_CREATE)
if(fd739 < 0) {
  panic("open failed")
}


ret = client.Seek(fd739, 78, client.SEEK_SET)
if(ret != 78) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 78)
  panic("seek failed")
}


ret = client.Seek(fd739, 100, client.SEEK_SET)
if(ret != 100) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 100)
  panic("seek failed")
}

//fd state: (100) ZywEH3A2RWij4v7ZR6IynoZ3j3nJc7CTQuXxHUFc9s0DBVGSALBHyGgwAstibrlUzhjqNNJZT3TYOWlaSFeR858QFSqlN9a0FfMKiKVCmVTVmaCqVNC1cAiaCk1Hc6sNLIl7

ret = client.Write(fd739, []byte("jrPJPn8DE1e5r06oU4vokAfEHX3BYnvEOYbmTbtz4nXPDd8V0ovXGJfqRlmiOCFst7wh6WFUbK9PduZKopD2u_imma62Mq2V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (196) ZywEH3A2RWij4v7ZR6IynoZ3j3nJc7CTQuXxHUFc9s0DBVGSALBHyGgwAstibrlUzhjqNNJZT3TYOWlaSFeR858QFSqlN9a0FfMKjrPJPn8DE1e5r06oU4vokAfEHX3BYnvEOYbmTbtz4nXPDd8V0ovXGJfqRlmiOCFst7wh6WFUbK9PduZKopD2u_imma62Mq2V

ret = client.Seek(fd739, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


ret = client.Close(fd739)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd738, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd740 := client.Open("/Zs6nTuJ6Ur.txt", client.O_RDWR|client.O_CREATE)
if(fd740 < 0) {
  panic("open failed")
}


ret = client.Seek(fd740, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (83) V7HI8TRadG_urYXps539QTAEnrFloAoKPwFJjWwAXTp_3DNMY2nHONzlExiD0xje61sR4n_252MG3HTNCHM1ulmv0eEwuRpalyi3YeNgTyCQ7ZzxJvy105JfnnQe1uPg_vRFk94

ret = client.Write(fd736, []byte("nBROlyM4WOkDCJXGGFBISL_G9kheMQI29eF3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (119) V7HI8TRadG_urYXps539QTAEnrFloAoKPwFJjWwAXTp_3DNMY2nHONzlExiD0xje61sR4n_252MG3HTNCHMnBROlyM4WOkDCJXGGFBISL_G9kheMQI29eF3fnnQe1uPg_vRFk94

ret = client.Seek(fd736, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd738, []byte("pLhUkrK335"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) pLhUkrK335

ret = client.Close(fd740)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd738)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd736, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}


ret = client.Seek(fd736, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Close(fd736)
if(ret != 0) {
  panic("close failed")
}


fd741 := client.Open("/3Xuls7KVgK.txt", client.O_RDWR|client.O_CREATE)
if(fd741 < 0) {
  panic("open failed")
}

//fd state: (0) IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzgvBSflqDO5DgxRcAoi3SGu6iDLt86hZEQPTZ5Pak1BJWfo0kt4hcwP8Fj

ret = client.Write(fd741, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzgvBSflqDO5DgxRcAoi3SGu6iDLt86hZEQPTZ5Pak1BJWfo0kt4hcwP8Fj

ret = client.Seek(fd741, 29, client.SEEK_SET)
if(ret != 29) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 29)
  panic("seek failed")
}


fd742 := client.Open("/MQPxqPotwN.txt", client.O_RDWR|client.O_CREATE)
if(fd742 < 0) {
  panic("open failed")
}


ret = client.Seek(fd742, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


ret = client.Seek(fd742, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


ret = client.Seek(fd742, 55, client.SEEK_SET)
if(ret != 55) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 55)
  panic("seek failed")
}


buf, ret = client.Read(fd741, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "wzgvBSflqDO5DgxRcAoi3SGu6iDLt") {
  panic("wrong data returned")
}

//fd state: (58) IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzgvBSflqDO5DgxRcAoi3SGu6iDLt86hZEQPTZ5Pak1BJWfo0kt4hcwP8Fj

ret = client.Write(fd741, []byte("XMm9WtQnaI4MzXqi7yRG46Yq2NcbSm8wVU4IGZ8yfLKCbzphZIfM8JKeGzeerpHaglUxQewr5w0GhOcm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (138) IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzgvBSflqDO5DgxRcAoi3SGu6iDLtXMm9WtQnaI4MzXqi7yRG46Yq2NcbSm8wVU4IGZ8yfLKCbzphZIfM8JKeGzeerpHaglUxQewr5w0GhOcm

ret = client.Seek(fd742, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


buf, ret = client.Read(fd742, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TKvT") {
  panic("wrong data returned")
}


ret = client.Close(fd742)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd741)
if(ret != 0) {
  panic("close failed")
}


fd743 := client.Open("/g52cddaIMR.txt", client.O_RDWR|client.O_CREATE)
if(fd743 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd743, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "XhDfZsalD_4KOYaQiqwN7Hxx5ZGzt1pUaOmBhVl_8dlRTg8LZRmCm3YEQ4IttUoOCP2") {
  panic("wrong data returned")
}


ret = client.Close(fd743)
if(ret != 0) {
  panic("close failed")
}


fd744 := client.Open("/5HpqVUKSSg.txt", client.O_RDWR|client.O_CREATE)
if(fd744 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd744, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd744, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd744, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd745 := client.Open("/qksOWSAKgO.txt", client.O_RDWR|client.O_CREATE)
if(fd745 < 0) {
  panic("open failed")
}


ret = client.Close(fd744)
if(ret != 0) {
  panic("close failed")
}


fd746 := client.Open("/7Ihi1jG3Qx.txt", client.O_RDWR|client.O_CREATE)
if(fd746 < 0) {
  panic("open failed")
}


ret = client.Seek(fd745, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd745, 21)
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


ret = client.Seek(fd746, 187, client.SEEK_SET)
if(ret != 187) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 187)
  panic("seek failed")
}


ret = client.Seek(fd746, 196, client.SEEK_SET)
if(ret != 196) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 196)
  panic("seek failed")
}

//fd state: (196) QNazqpcx_DqbyCCBIjPl4t9iyjKtUvIxD0gbypjNXRV3rXj696pBi8XlSiHG9ehTqgQLVN872zn0PtF15VfyTBIi0sV7_edEYjZ2UxD2hhKbIHthDyqslMmEt14Xrr7TRUeLy5d9vUF8TzZcnLTxwNZNES9ZGbU2XVbh3n_goJcIQdRFoyS88GzFAEHlEC75qSYn_J1VDAEbiqcQ1ITy6614GFz6W6aY85xIXpp7lexLO7tWgqWqZqsbS

ret = client.Write(fd746, []byte("lt0thOAV0BE0Nw7CNGdw6Yb4JNusG4bSRl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (230) QNazqpcx_DqbyCCBIjPl4t9iyjKtUvIxD0gbypjNXRV3rXj696pBi8XlSiHG9ehTqgQLVN872zn0PtF15VfyTBIi0sV7_edEYjZ2UxD2hhKbIHthDyqslMmEt14Xrr7TRUeLy5d9vUF8TzZcnLTxwNZNES9ZGbU2XVbh3n_goJcIQdRFoyS88GzFAEHlEC75qSYnlt0thOAV0BE0Nw7CNGdw6Yb4JNusG4bSRlp7lexLO7tWgqWqZqsbS

ret = client.Close(fd746)
if(ret != 0) {
  panic("close failed")
}


fd747 := client.Open("/p148qPA_oh.txt", client.O_RDWR|client.O_CREATE)
if(fd747 < 0) {
  panic("open failed")
}


ret = client.Close(fd747)
if(ret != 0) {
  panic("close failed")
}


fd748 := client.Open("/2BgRhfO5sI.txt", client.O_RDWR|client.O_CREATE)
if(fd748 < 0) {
  panic("open failed")
}


ret = client.Close(fd748)
if(ret != 0) {
  panic("close failed")
}


fd749 := client.Open("/WNMJT2E_a5.txt", client.O_RDWR|client.O_CREATE)
if(fd749 < 0) {
  panic("open failed")
}


fd750 := client.Open("/sP6wfz2wrs.txt", client.O_RDWR|client.O_CREATE)
if(fd750 < 0) {
  panic("open failed")
}


fd751 := client.Open("/pSvqpvAdra.txt", client.O_RDWR|client.O_CREATE)
if(fd751 < 0) {
  panic("open failed")
}


fd752 := client.Open("/FZZ8y6anbP.txt", client.O_RDWR|client.O_CREATE)
if(fd752 < 0) {
  panic("open failed")
}


fd753 := client.Open("/laYd5cF8cy.txt", client.O_RDWR|client.O_CREATE)
if(fd753 < 0) {
  panic("open failed")
}


fd754 := client.Open("/N07ScZRqX2.txt", client.O_RDWR|client.O_CREATE)
if(fd754 < 0) {
  panic("open failed")
}


fd755 := client.Open("/j6CCRdMK5y.txt", client.O_RDWR|client.O_CREATE)
if(fd755 < 0) {
  panic("open failed")
}


ret = client.Close(fd751)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd752, []byte("mPHAWHjoGFH6JrC9kVSeUNEH5X2Ln3P_wKkeMWJUE0usmiPyvLuUfUlyLc8EZk82mxOsSOp5_t2AcK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) mPHAWHjoGFH6JrC9kVSeUNEH5X2Ln3P_wKkeMWJUE0usmiPyvLuUfUlyLc8EZk82mxOsSOp5_t2AcK
//fd state: (0) 

ret = client.Write(fd753, []byte("oRbtEUvZEEeD0XcHhHwf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) oRbtEUvZEEeD0XcHhHwf
//fd state: (78) mPHAWHjoGFH6JrC9kVSeUNEH5X2Ln3P_wKkeMWJUE0usmiPyvLuUfUlyLc8EZk82mxOsSOp5_t2AcK

ret = client.Write(fd752, []byte("oPFlrPATLLqWhL7aryb3WuTINjsLtGRb4i7r9n6gz6IXcQhL87As"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (130) mPHAWHjoGFH6JrC9kVSeUNEH5X2Ln3P_wKkeMWJUE0usmiPyvLuUfUlyLc8EZk82mxOsSOp5_t2AcKoPFlrPATLLqWhL7aryb3WuTINjsLtGRb4i7r9n6gz6IXcQhL87As
//fd state: (130) mPHAWHjoGFH6JrC9kVSeUNEH5X2Ln3P_wKkeMWJUE0usmiPyvLuUfUlyLc8EZk82mxOsSOp5_t2AcKoPFlrPATLLqWhL7aryb3WuTINjsLtGRb4i7r9n6gz6IXcQhL87As

ret = client.Write(fd752, []byte("PwJCc8S2E3MipvkQQO7Oy3mNRm2MizYsPT70VKlUHo6XRrd_vCy6_5Z8EzZJIhWm7_8I2X2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (201) mPHAWHjoGFH6JrC9kVSeUNEH5X2Ln3P_wKkeMWJUE0usmiPyvLuUfUlyLc8EZk82mxOsSOp5_t2AcKoPFlrPATLLqWhL7aryb3WuTINjsLtGRb4i7r9n6gz6IXcQhL87AsPwJCc8S2E3MipvkQQO7Oy3mNRm2MizYsPT70VKlUHo6XRrd_vCy6_5Z8EzZJIhWm7_8I2X2

buf, ret = client.Read(fd754, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd749, 5)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "jGhuT") {
  panic("wrong data returned")
}


ret = client.Seek(fd754, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd754, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd753, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


fd756 := client.Open("/H95Lu2se4U.txt", client.O_RDWR|client.O_CREATE)
if(fd756 < 0) {
  panic("open failed")
}


ret = client.Close(fd754)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd755, []byte("uQcGj0BblK_1ZPvfD0m9abQNxxSRtx6OGh9ZOMQyv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (41) uQcGj0BblK_1ZPvfD0m9abQNxxSRtx6OGh9ZOMQyv

buf, ret = client.Read(fd756, 92)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "RcVdYxX_cwhAeCpUpmrG5F6pB0Z") {
  panic("wrong data returned")
}

//fd state: (5) jGhuTeatzbyKczU0nQRjPYg8ClP6jl2DA_SNkf4mQGVqB5QXoQJAiuO7zriLdARzrvVeuo5zt2c9dEONaE0_9ZkescctVzmhoSlZEmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyYkTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3Xnhl

ret = client.Write(fd749, []byte("JWN8bSOLIzKLYMquz5zJF7juenwmeowhKh2J1052SZOfhTAiGFIaKUn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) jGhuTJWN8bSOLIzKLYMquz5zJF7juenwmeowhKh2J1052SZOfhTAiGFIaKUndARzrvVeuo5zt2c9dEONaE0_9ZkescctVzmhoSlZEmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyYkTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3Xnhl

buf, ret = client.Read(fd750, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd750, []byte("qY4gGJQfcZZRvICUGnOgBhEwZtdS6A_r"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) qY4gGJQfcZZRvICUGnOgBhEwZtdS6A_r

buf, ret = client.Read(fd756, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd757 := client.Open("/NhWH3hnYlu.txt", client.O_RDWR|client.O_CREATE)
if(fd757 < 0) {
  panic("open failed")
}

//fd state: (32) qY4gGJQfcZZRvICUGnOgBhEwZtdS6A_r

ret = client.Write(fd750, []byte("nUggXGsFpDi05733QsoWfaelKGzYxKO7259GbDCTbQ4DC9bHcvxj6u7mbjaQtq4n8VXqzc6yulbJM64mu6h"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) qY4gGJQfcZZRvICUGnOgBhEwZtdS6A_rnUggXGsFpDi05733QsoWfaelKGzYxKO7259GbDCTbQ4DC9bHcvxj6u7mbjaQtq4n8VXqzc6yulbJM64mu6h

ret = client.Close(fd757)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd756)
if(ret != 0) {
  panic("close failed")
}


fd758 := client.Open("/HvvyN1FrRe.txt", client.O_RDWR|client.O_CREATE)
if(fd758 < 0) {
  panic("open failed")
}


ret = client.Close(fd750)
if(ret != 0) {
  panic("close failed")
}

//fd state: (41) uQcGj0BblK_1ZPvfD0m9abQNxxSRtx6OGh9ZOMQyv

ret = client.Write(fd755, []byte("ucGwm7gigFgjHN3ONFklMqriXpVyYvY0z8k3OwqD4fAAuo08TsPHAcfyhS2Ph"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (102) uQcGj0BblK_1ZPvfD0m9abQNxxSRtx6OGh9ZOMQyvucGwm7gigFgjHN3ONFklMqriXpVyYvY0z8k3OwqD4fAAuo08TsPHAcfyhS2Ph

fd759 := client.Open("/yNSeHLcutx.txt", client.O_RDWR|client.O_CREATE)
if(fd759 < 0) {
  panic("open failed")
}


fd760 := client.Open("/qCwNyLAYQa.txt", client.O_RDWR|client.O_CREATE)
if(fd760 < 0) {
  panic("open failed")
}


fd761 := client.Open("/5HpqVUKSSg.txt", client.O_RDWR|client.O_CREATE)
if(fd761 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd760, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd758, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd762 := client.Open("/cVKOQbpRpD.txt", client.O_RDWR|client.O_CREATE)
if(fd762 < 0) {
  panic("open failed")
}


ret = client.Seek(fd758, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd763 := client.Open("/H6rFCmF66g.txt", client.O_RDWR|client.O_CREATE)
if(fd763 < 0) {
  panic("open failed")
}


fd764 := client.Open("/LgDXLTVcW9.txt", client.O_RDWR|client.O_CREATE)
if(fd764 < 0) {
  panic("open failed")
}


ret = client.Seek(fd749, 184, client.SEEK_SET)
if(ret != 184) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 184)
  panic("seek failed")
}


ret = client.Close(fd752)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd761)
if(ret != 0) {
  panic("close failed")
}


fd765 := client.Open("/dAkT0T3Pq_.txt", client.O_RDWR|client.O_CREATE)
if(fd765 < 0) {
  panic("open failed")
}

//fd state: (0) 6KAy19xqY8oLiyqAKXurYg7RRMcbs6RE6EZdYEQGrfzxV40l0Kr4ZYRC74PsyhdIHxkcP4S1hmeZKd

ret = client.Write(fd765, []byte("6GYX6r4CSbQwL9DE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (16) 6GYX6r4CSbQwL9DEKXurYg7RRMcbs6RE6EZdYEQGrfzxV40l0Kr4ZYRC74PsyhdIHxkcP4S1hmeZKd

fd766 := client.Open("/cVKOQbpRpD.txt", client.O_RDWR|client.O_CREATE)
if(fd766 < 0) {
  panic("open failed")
}


ret = client.Seek(fd766, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd749, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8nGpxQljJuWSv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3Xn") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd766, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd758, []byte("rGiSWR8HelcjXQExYGvBdhiH5g0V4ujkYldf8Oy70ccNln0qisMnHLQCn6npw_OqmX_9Exnd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) rGiSWR8HelcjXQExYGvBdhiH5g0V4ujkYldf8Oy70ccNln0qisMnHLQCn6npw_OqmX_9Exnd

ret = client.Close(fd760)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd762, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd765)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd762, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd755, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


ret = client.Seek(fd766, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd767 := client.Open("/DPMWRgC7B4.txt", client.O_RDWR|client.O_CREATE)
if(fd767 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd762, []byte("DHySLeoovBWGcKsz0tc6wc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) DHySLeoovBWGcKsz0tc6wc

fd768 := client.Open("/nPCl6unWmR.txt", client.O_RDWR|client.O_CREATE)
if(fd768 < 0) {
  panic("open failed")
}


ret = client.Close(fd753)
if(ret != 0) {
  panic("close failed")
}

//fd state: (264) jGhuTJWN8bSOLIzKLYMquz5zJF7juenwmeowhKh2J1052SZOfhTAiGFIaKUndARzrvVeuo5zt2c9dEONaE0_9ZkescctVzmhoSlZEmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyYkTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3Xnhl

ret = client.Write(fd749, []byte("lIP061t5GV4wpOD_IZZibK_kRiGlXN_xi"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (297) jGhuTJWN8bSOLIzKLYMquz5zJF7juenwmeowhKh2J1052SZOfhTAiGFIaKUndARzrvVeuo5zt2c9dEONaE0_9ZkescctVzmhoSlZEmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyYkTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3XnlIP061t5GV4wpOD_IZZibK_kRiGlXN_xi

buf, ret = client.Read(fd764, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "g8fgNROr2pdn1THyrFPrGvu6VAE050zlS7O03A9aKXA_Kpj18KGoUUHONi_EXWxL2qn8") {
  panic("wrong data returned")
}


ret = client.Close(fd758)
if(ret != 0) {
  panic("close failed")
}


fd769 := client.Open("/uIv3bq5PSV.txt", client.O_RDWR|client.O_CREATE)
if(fd769 < 0) {
  panic("open failed")
}


fd770 := client.Open("/Yt5ysR3lnX.txt", client.O_RDWR|client.O_CREATE)
if(fd770 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd768, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pLhUkrK335") {
  panic("wrong data returned")
}


ret = client.Seek(fd749, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


ret = client.Close(fd766)
if(ret != 0) {
  panic("close failed")
}


fd771 := client.Open("/gIZcMTy6C7.txt", client.O_RDWR|client.O_CREATE)
if(fd771 < 0) {
  panic("open failed")
}


fd772 := client.Open("/LKNpLlv1eM.txt", client.O_RDWR|client.O_CREATE)
if(fd772 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd769, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "sjOtsBYxkKA9tJ1qKak1fmgSwGtWgUSB4hTkx2IkRG3_srW") {
  panic("wrong data returned")
}


ret = client.Close(fd755)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd769, 47, client.SEEK_SET)
if(ret != 47) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 47)
  panic("seek failed")
}


ret = client.Close(fd764)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd772, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd762, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd762)
if(ret != 0) {
  panic("close failed")
}


fd773 := client.Open("/XQFqEQurVL.txt", client.O_RDWR|client.O_CREATE)
if(fd773 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd773, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd774 := client.Open("/az63yNz5FL.txt", client.O_RDWR|client.O_CREATE)
if(fd774 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd767, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd774, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd775 := client.Open("/qtH7OOL_RQ.txt", client.O_RDWR|client.O_CREATE)
if(fd775 < 0) {
  panic("open failed")
}


ret = client.Seek(fd772, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd768)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd771)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd749, 104, client.SEEK_SET)
if(ret != 104) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 104)
  panic("seek failed")
}


ret = client.Seek(fd759, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd774)
if(ret != 0) {
  panic("close failed")
}


fd776 := client.Open("/1hYlxqnLZN.txt", client.O_RDWR|client.O_CREATE)
if(fd776 < 0) {
  panic("open failed")
}


ret = client.Seek(fd773, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd759, 8)
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


buf, ret = client.Read(fd775, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd767, []byte("10_inySFR66YwxBifUVEE1Nl54ueAluEkS476UVbjnfCyK2j_4_fg2cZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) 10_inySFR66YwxBifUVEE1Nl54ueAluEkS476UVbjnfCyK2j_4_fg2cZ
//fd state: (0) kNFje5jzWpr4wcLYDGQYe8jJXQmljmLsHCT07Yxt8chFMVnCPeC7uEHr0LZ2Mwf1UZsB32CwWicR

ret = client.Write(fd776, []byte("fWKiDuMl7ndE5mTlJpX0rmbqm3CDLCN2u3Xhk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) fWKiDuMl7ndE5mTlJpX0rmbqm3CDLCN2u3XhkYxt8chFMVnCPeC7uEHr0LZ2Mwf1UZsB32CwWicR
//fd state: (0) 

ret = client.Write(fd759, []byte("lT7XKr_yQC6HCWPtbdNiw7WWY9zlrnPllVNckGz5q5wh_QGfbbWbermPSJkQrNMqAR8Zx2TSW2FcB2TQyplrbHr3HGMKJwW2d"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) lT7XKr_yQC6HCWPtbdNiw7WWY9zlrnPllVNckGz5q5wh_QGfbbWbermPSJkQrNMqAR8Zx2TSW2FcB2TQyplrbHr3HGMKJwW2d

ret = client.Close(fd776)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd770, []byte("8T291IinDr75hy4_R"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (17) 8T291IinDr75hy4_R
//fd state: (104) jGhuTJWN8bSOLIzKLYMquz5zJF7juenwmeowhKh2J1052SZOfhTAiGFIaKUndARzrvVeuo5zt2c9dEONaE0_9ZkescctVzmhoSlZEmWq21Ne9zoCBFV4XqKP8HJs2fjlukmMFBfLpZoaadAqfc9beHyYkTbZJR5NY53L3Yy4huWfOR9N1A4QUptR8nGpxQljJuWSv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3XnlIP061t5GV4wpOD_IZZibK_kRiGlXN_xi

ret = client.Write(fd749, []byte("IgKTn1bWiIHTkkusYcmzN04h06Bs0HAUGwLhRiEWdO4Ttlrz_Ujy0VZ2ikW8LP7g4KEql7R__xTmoMfCrg29ytLE4cyv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (196) jGhuTJWN8bSOLIzKLYMquz5zJF7juenwmeowhKh2J1052SZOfhTAiGFIaKUndARzrvVeuo5zt2c9dEONaE0_9ZkescctVzmhoSlZEmWqIgKTn1bWiIHTkkusYcmzN04h06Bs0HAUGwLhRiEWdO4Ttlrz_Ujy0VZ2ikW8LP7g4KEql7R__xTmoMfCrg29ytLE4cyvv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3XnlIP061t5GV4wpOD_IZZibK_kRiGlXN_xi

buf, ret = client.Read(fd749, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "v1R1757BGOaD2M00xNp2fT7Q7Tv") {
  panic("wrong data returned")
}

//fd state: (17) 8T291IinDr75hy4_R

ret = client.Write(fd770, []byte("7PJ5_8Z0Z9VkEtALvbYeU7mo2dvYm_Vd5d1FHkS7P_ONzJspX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) 8T291IinDr75hy4_R7PJ5_8Z0Z9VkEtALvbYeU7mo2dvYm_Vd5d1FHkS7P_ONzJspX

ret = client.Seek(fd763, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (223) jGhuTJWN8bSOLIzKLYMquz5zJF7juenwmeowhKh2J1052SZOfhTAiGFIaKUndARzrvVeuo5zt2c9dEONaE0_9ZkescctVzmhoSlZEmWqIgKTn1bWiIHTkkusYcmzN04h06Bs0HAUGwLhRiEWdO4Ttlrz_Ujy0VZ2ikW8LP7g4KEql7R__xTmoMfCrg29ytLE4cyvv1R1757BGOaD2M00xNp2fT7Q7TvacbtHlTN3IMYW9z0CrM35fAOUBO3eSuGUXUJ9o3XnlIP061t5GV4wpOD_IZZibK_kRiGlXN_xi

ret = client.Write(fd749, []byte("Fa8FWInmyzkZw3GvBJM3Fc3ZEudM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (251) jGhuTJWN8bSOLIzKLYMquz5zJF7juenwmeowhKh2J1052SZOfhTAiGFIaKUndARzrvVeuo5zt2c9dEONaE0_9ZkescctVzmhoSlZEmWqIgKTn1bWiIHTkkusYcmzN04h06Bs0HAUGwLhRiEWdO4Ttlrz_Ujy0VZ2ikW8LP7g4KEql7R__xTmoMfCrg29ytLE4cyvv1R1757BGOaD2M00xNp2fT7Q7TvFa8FWInmyzkZw3GvBJM3Fc3ZEudMeSuGUXUJ9o3XnlIP061t5GV4wpOD_IZZibK_kRiGlXN_xi
//fd state: (47) sjOtsBYxkKA9tJ1qKak1fmgSwGtWgUSB4hTkx2IkRG3_srWMLfcjxjqwMIp8zd2Dkuw3BDDjGR9n40aNKDM

ret = client.Write(fd769, []byte("m7czrTfHNowkl9AY7g3zZVS4DZoz3yfqu4djaXz_JAjNcrQVFU55H2XRyUT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (106) sjOtsBYxkKA9tJ1qKak1fmgSwGtWgUSB4hTkx2IkRG3_srWm7czrTfHNowkl9AY7g3zZVS4DZoz3yfqu4djaXz_JAjNcrQVFU55H2XRyUT

buf, ret = client.Read(fd773, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd749, 138, client.SEEK_SET)
if(ret != 138) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 138)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd775, []byte("OYfsz4cl_EBAYg6VfcCSoGlHnhkQ4EMuJClP01jMw4VaSV60fYTlFq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) OYfsz4cl_EBAYg6VfcCSoGlHnhkQ4EMuJClP01jMw4VaSV60fYTlFq

fd777 := client.Open("/YqBc9d_V9a.txt", client.O_RDWR|client.O_CREATE)
if(fd777 < 0) {
  panic("open failed")
}


ret = client.Close(fd767)
if(ret != 0) {
  panic("close failed")
}


fd778 := client.Open("/3SAmTKFcdh.txt", client.O_RDWR|client.O_CREATE)
if(fd778 < 0) {
  panic("open failed")
}


ret = client.Seek(fd770, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


ret = client.Seek(fd769, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


ret = client.Close(fd749)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd778)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd763, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd759, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


ret = client.Seek(fd770, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}

//fd state: (17) sjOtsBYxkKA9tJ1qKak1fmgSwGtWgUSB4hTkx2IkRG3_srWm7czrTfHNowkl9AY7g3zZVS4DZoz3yfqu4djaXz_JAjNcrQVFU55H2XRyUT

ret = client.Write(fd769, []byte("z5sBAaSlaZG0WuTAkCueluwOIyx4iGb1LYQVg051oNp4vOFjbrLtdxvzf8QGvTj7lIFm5Pj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) sjOtsBYxkKA9tJ1qKz5sBAaSlaZG0WuTAkCueluwOIyx4iGb1LYQVg051oNp4vOFjbrLtdxvzf8QGvTj7lIFm5PjAjNcrQVFU55H2XRyUT
//fd state: (0) 

ret = client.Write(fd777, []byte("Zr_K0QNVZ0U53KZ_XQLPfabcli_SFnCg_AtYn_lw6rbiVnsFh8noTcxVFTTVA8pT2tTdvewCcTbAnZG2xXmiilaNBnD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) Zr_K0QNVZ0U53KZ_XQLPfabcli_SFnCg_AtYn_lw6rbiVnsFh8noTcxVFTTVA8pT2tTdvewCcTbAnZG2xXmiilaNBnD

fd779 := client.Open("/LGqspx2N6R.txt", client.O_RDWR|client.O_CREATE)
if(fd779 < 0) {
  panic("open failed")
}


ret = client.Seek(fd779, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd769, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AjNcrQVFU55H2XRyUT") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd770, 36)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YeU7mo2dvYm_Vd5d1FHkS7P_ONzJspX") {
  panic("wrong data returned")
}


ret = client.Seek(fd763, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd769, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd763, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd777, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (66) 8T291IinDr75hy4_R7PJ5_8Z0Z9VkEtALvbYeU7mo2dvYm_Vd5d1FHkS7P_ONzJspX

ret = client.Write(fd770, []byte("AMGWBj20WgrAqlp7s1gH6QhD9uoiXo_8yfcXmbJXcw6veTm9LzjvcKX5nryXvLvHX3a"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (133) 8T291IinDr75hy4_R7PJ5_8Z0Z9VkEtALvbYeU7mo2dvYm_Vd5d1FHkS7P_ONzJspXAMGWBj20WgrAqlp7s1gH6QhD9uoiXo_8yfcXmbJXcw6veTm9LzjvcKX5nryXvLvHX3a

ret = client.Seek(fd775, 8, client.SEEK_SET)
if(ret != 8) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 8)
  panic("seek failed")
}


buf, ret = client.Read(fd769, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd759, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


ret = client.Seek(fd773, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd770, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd769, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd780 := client.Open("/ez2_D3d8Wg.txt", client.O_RDWR|client.O_CREATE)
if(fd780 < 0) {
  panic("open failed")
}


ret = client.Close(fd769)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd779, 56)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd781 := client.Open("/fTKEpwYowY.txt", client.O_RDWR|client.O_CREATE)
if(fd781 < 0) {
  panic("open failed")
}

//fd state: (91) Zr_K0QNVZ0U53KZ_XQLPfabcli_SFnCg_AtYn_lw6rbiVnsFh8noTcxVFTTVA8pT2tTdvewCcTbAnZG2xXmiilaNBnD

ret = client.Write(fd777, []byte("kbpzq9ZNiov5cr8Hh2lW1pK3U6wM4dQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) Zr_K0QNVZ0U53KZ_XQLPfabcli_SFnCg_AtYn_lw6rbiVnsFh8noTcxVFTTVA8pT2tTdvewCcTbAnZG2xXmiilaNBnDkbpzq9ZNiov5cr8Hh2lW1pK3U6wM4dQ

fd782 := client.Open("/MIUhU1fa3X.txt", client.O_RDWR|client.O_CREATE)
if(fd782 < 0) {
  panic("open failed")
}


fd783 := client.Open("/uS4PfPrncQ.txt", client.O_RDWR|client.O_CREATE)
if(fd783 < 0) {
  panic("open failed")
}


fd784 := client.Open("/eckusOytLJ.txt", client.O_RDWR|client.O_CREATE)
if(fd784 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd783, []byte("WUORDdFhX7aGgUN2zIsira4i2V7MF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (29) WUORDdFhX7aGgUN2zIsira4i2V7MF
//fd state: (0) 

ret = client.Write(fd763, []byte("uO7tWjKJmT7LVzJf9ezREnUSuyu363jqHLkzvJxuqa8JlZ0h1a5HOpL7b83b6AuEhVBbZsau97S5TsC0b9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) uO7tWjKJmT7LVzJf9ezREnUSuyu363jqHLkzvJxuqa8JlZ0h1a5HOpL7b83b6AuEhVBbZsau97S5TsC0b9

ret = client.Close(fd763)
if(ret != 0) {
  panic("close failed")
}


fd785 := client.Open("/gKjRc6XTZ4.txt", client.O_RDWR|client.O_CREATE)
if(fd785 < 0) {
  panic("open failed")
}


fd786 := client.Open("/FZZ8y6anbP.txt", client.O_RDWR|client.O_CREATE)
if(fd786 < 0) {
  panic("open failed")
}


ret = client.Close(fd786)
if(ret != 0) {
  panic("close failed")
}

//fd state: (133) 8T291IinDr75hy4_R7PJ5_8Z0Z9VkEtALvbYeU7mo2dvYm_Vd5d1FHkS7P_ONzJspXAMGWBj20WgrAqlp7s1gH6QhD9uoiXo_8yfcXmbJXcw6veTm9LzjvcKX5nryXvLvHX3a

ret = client.Write(fd770, []byte("4Dhg9py6PMPcQCTYnprnLcOhhMgBSL5GvCmVmTHVc2on53N3_kesb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (186) 8T291IinDr75hy4_R7PJ5_8Z0Z9VkEtALvbYeU7mo2dvYm_Vd5d1FHkS7P_ONzJspXAMGWBj20WgrAqlp7s1gH6QhD9uoiXo_8yfcXmbJXcw6veTm9LzjvcKX5nryXvLvHX3a4Dhg9py6PMPcQCTYnprnLcOhhMgBSL5GvCmVmTHVc2on53N3_kesb

buf, ret = client.Read(fd759, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "SJkQrNMqAR8Zx") {
  panic("wrong data returned")
}

//fd state: (29) WUORDdFhX7aGgUN2zIsira4i2V7MF

ret = client.Write(fd783, []byte("gHqxsIZKaa5LmZ9WUrmu5RADZN9jb5urL2DGG6A3_FWD9JNRkXsN4FZryw2C3Uy7YBdoZS_ySQfT0cTkTzoCHUmYw23X"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (121) WUORDdFhX7aGgUN2zIsira4i2V7MFgHqxsIZKaa5LmZ9WUrmu5RADZN9jb5urL2DGG6A3_FWD9JNRkXsN4FZryw2C3Uy7YBdoZS_ySQfT0cTkTzoCHUmYw23X

ret = client.Close(fd783)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd777, 12, client.SEEK_SET)
if(ret != 12) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 12)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (8) OYfsz4cl_EBAYg6VfcCSoGlHnhkQ4EMuJClP01jMw4VaSV60fYTlFq

ret = client.Write(fd775, []byte("9PYc4Zj2L7434hObeaKkrtGxPcG4g7Vxpc1v5JCol4cX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) OYfsz4cl9PYc4Zj2L7434hObeaKkrtGxPcG4g7Vxpc1v5JCol4cXFq

ret = client.Seek(fd782, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd787 := client.Open("/5bYZ4qlxNy.txt", client.O_RDWR|client.O_CREATE)
if(fd787 < 0) {
  panic("open failed")
}


ret = client.Seek(fd770, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd782, []byte("eazdovy4uzn7fzn9NLxn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) eazdovy4uzn7fzn9NLxn

fd788 := client.Open("/XdRMozFlbD.txt", client.O_RDWR|client.O_CREATE)
if(fd788 < 0) {
  panic("open failed")
}


ret = client.Seek(fd775, 50, client.SEEK_SET)
if(ret != 50) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 50)
  panic("seek failed")
}


fd789 := client.Open("/lWg_UGb3WU.txt", client.O_RDWR|client.O_CREATE)
if(fd789 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd789, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd785, 65, client.SEEK_SET)
if(ret != 65) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 65)
  panic("seek failed")
}


fd790 := client.Open("/ex0m66juin.txt", client.O_RDWR|client.O_CREATE)
if(fd790 < 0) {
  panic("open failed")
}


ret = client.Close(fd784)
if(ret != 0) {
  panic("close failed")
}


fd791 := client.Open("/eshGbVAHWt.txt", client.O_RDWR|client.O_CREATE)
if(fd791 < 0) {
  panic("open failed")
}


fd792 := client.Open("/fiHZii4BMm.txt", client.O_RDWR|client.O_CREATE)
if(fd792 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd779, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd790)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd785, 142, client.SEEK_SET)
if(ret != 142) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 142)
  panic("seek failed")
}


fd793 := client.Open("/VCgdE3WhbF.txt", client.O_RDWR|client.O_CREATE)
if(fd793 < 0) {
  panic("open failed")
}


fd794 := client.Open("/8rhje1DUVt.txt", client.O_RDWR|client.O_CREATE)
if(fd794 < 0) {
  panic("open failed")
}


fd795 := client.Open("/LoTeuE2wHL.txt", client.O_RDWR|client.O_CREATE)
if(fd795 < 0) {
  panic("open failed")
}

//fd state: (0) IYca5Ivuhtq4TfPwPVojsU1rHghOEUoI8b_F0GUTS3beuNO

ret = client.Write(fd793, []byte("tUDG8D"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (6) tUDG8Dvuhtq4TfPwPVojsU1rHghOEUoI8b_F0GUTS3beuNO

ret = client.Close(fd792)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd789, []byte("FTQESMaEGavTx9Dgv7QZzw85qSf2Ri3wxgRr72nmsUwLFEoiq_mdMneqZswav2brP167vsEFmBTMdoN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) FTQESMaEGavTx9Dgv7QZzw85qSf2Ri3wxgRr72nmsUwLFEoiq_mdMneqZswav2brP167vsEFmBTMdoN

buf, ret = client.Read(fd773, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd795, 322, client.SEEK_SET)
if(ret != 322) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 322)
  panic("seek failed")
}


ret = client.Close(fd793)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd773, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (12) Zr_K0QNVZ0U53KZ_XQLPfabcli_SFnCg_AtYn_lw6rbiVnsFh8noTcxVFTTVA8pT2tTdvewCcTbAnZG2xXmiilaNBnDkbpzq9ZNiov5cr8Hh2lW1pK3U6wM4dQ

ret = client.Write(fd777, []byte("NouWQQMbjQniqHytyUfaJf7Adnr5E3mpvzzmtlHSbNW4XK8DRfxCbre3_tXSIuonJ3ygejaVjzERfeGsNkq5baOLtLvDXadx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (108) Zr_K0QNVZ0U5NouWQQMbjQniqHytyUfaJf7Adnr5E3mpvzzmtlHSbNW4XK8DRfxCbre3_tXSIuonJ3ygejaVjzERfeGsNkq5baOLtLvDXadx2lW1pK3U6wM4dQ

buf, ret = client.Read(fd775, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "cXFq") {
  panic("wrong data returned")
}


ret = client.Seek(fd787, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd796 := client.Open("/YyhcUiTFU4.txt", client.O_RDWR|client.O_CREATE)
if(fd796 < 0) {
  panic("open failed")
}


ret = client.Seek(fd796, 77, client.SEEK_SET)
if(ret != 77) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 77)
  panic("seek failed")
}


fd797 := client.Open("/bscOPxCqaM.txt", client.O_RDWR|client.O_CREATE)
if(fd797 < 0) {
  panic("open failed")
}

//fd state: (20) eazdovy4uzn7fzn9NLxn

ret = client.Write(fd782, []byte("YZun21x0Ahg6g3XBcodgCEmFbaOKDmd30ZY5MN7a6Q9AY9J62DuZTWcwLfBjg__VCzz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) eazdovy4uzn7fzn9NLxnYZun21x0Ahg6g3XBcodgCEmFbaOKDmd30ZY5MN7a6Q9AY9J62DuZTWcwLfBjg__VCzz

buf, ret = client.Read(fd775, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd798 := client.Open("/Se7h93SXZ2.txt", client.O_RDWR|client.O_CREATE)
if(fd798 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd782, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd788)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd798, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd782)
if(ret != 0) {
  panic("close failed")
}

//fd state: (77) QJv2H3RdVSn8imaCN0LQ77vtTfiaeYGkxpJNlAbuW00_sEnUSqPTTtgpPBOw855snZLYu4Rl5mO4uKjeN

ret = client.Write(fd796, []byte("5ECilv9s_jP93cFd2v0HwbiemgaeruhJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (109) QJv2H3RdVSn8imaCN0LQ77vtTfiaeYGkxpJNlAbuW00_sEnUSqPTTtgpPBOw855snZLYu4Rl5mO4u5ECilv9s_jP93cFd2v0HwbiemgaeruhJ
//fd state: (54) OYfsz4cl9PYc4Zj2L7434hObeaKkrtGxPcG4g7Vxpc1v5JCol4cXFq

ret = client.Write(fd775, []byte("aKRPoSXNY425ZUgneRJ28JvSjca_QSEyBazCQlvpI4uaqxV6G49w7O_A_FzXye9avKLw6wjkOu41yr"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) OYfsz4cl9PYc4Zj2L7434hObeaKkrtGxPcG4g7Vxpc1v5JCol4cXFqaKRPoSXNY425ZUgneRJ28JvSjca_QSEyBazCQlvpI4uaqxV6G49w7O_A_FzXye9avKLw6wjkOu41yr

fd799 := client.Open("/6s3j3b8f8n.txt", client.O_RDWR|client.O_CREATE)
if(fd799 < 0) {
  panic("open failed")
}


ret = client.Close(fd789)
if(ret != 0) {
  panic("close failed")
}


fd800 := client.Open("/v2DsGdHGbp.txt", client.O_RDWR|client.O_CREATE)
if(fd800 < 0) {
  panic("open failed")
}


ret = client.Close(fd795)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd781, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


ret = client.Seek(fd796, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


fd801 := client.Open("/PtxNuQ98Qr.txt", client.O_RDWR|client.O_CREATE)
if(fd801 < 0) {
  panic("open failed")
}


fd802 := client.Open("/Q_2vsRUS7G.txt", client.O_RDWR|client.O_CREATE)
if(fd802 < 0) {
  panic("open failed")
}

//fd state: (0) uvgxwz4GdXwI7eFtoTrLia7XXDx0E0ZZ_4gejABrgdcNwD8gLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG51V23mKFTZbJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

ret = client.Write(fd797, []byte("65ocKPNQhErm3i9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (15) 65ocKPNQhErm3i9toTrLia7XXDx0E0ZZ_4gejABrgdcNwD8gLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG51V23mKFTZbJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

ret = client.Close(fd801)
if(ret != 0) {
  panic("close failed")
}

//fd state: (4) 8T291IinDr75hy4_R7PJ5_8Z0Z9VkEtALvbYeU7mo2dvYm_Vd5d1FHkS7P_ONzJspXAMGWBj20WgrAqlp7s1gH6QhD9uoiXo_8yfcXmbJXcw6veTm9LzjvcKX5nryXvLvHX3a4Dhg9py6PMPcQCTYnprnLcOhhMgBSL5GvCmVmTHVc2on53N3_kesb

ret = client.Write(fd770, []byte("1_jIaWt1rCJLvqEoCJsMc_KDMgGjJJMiG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) 8T291_jIaWt1rCJLvqEoCJsMc_KDMgGjJJMiGU7mo2dvYm_Vd5d1FHkS7P_ONzJspXAMGWBj20WgrAqlp7s1gH6QhD9uoiXo_8yfcXmbJXcw6veTm9LzjvcKX5nryXvLvHX3a4Dhg9py6PMPcQCTYnprnLcOhhMgBSL5GvCmVmTHVc2on53N3_kesb
//fd state: (69) lT7XKr_yQC6HCWPtbdNiw7WWY9zlrnPllVNckGz5q5wh_QGfbbWbermPSJkQrNMqAR8Zx2TSW2FcB2TQyplrbHr3HGMKJwW2d

ret = client.Write(fd759, []byte("TG8GiMDoZVb9ja8dtoy9vnMCZEQR_kuwcwBLw3ofRAeu55qDF0hyyX2g6crAxk25S2uBvUajSd1q8KHHRwfIWOt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (156) lT7XKr_yQC6HCWPtbdNiw7WWY9zlrnPllVNckGz5q5wh_QGfbbWbermPSJkQrNMqAR8ZxTG8GiMDoZVb9ja8dtoy9vnMCZEQR_kuwcwBLw3ofRAeu55qDF0hyyX2g6crAxk25S2uBvUajSd1q8KHHRwfIWOt

ret = client.Close(fd787)
if(ret != 0) {
  panic("close failed")
}


fd803 := client.Open("/gMb60G3hQV.txt", client.O_RDWR|client.O_CREATE)
if(fd803 < 0) {
  panic("open failed")
}


fd804 := client.Open("/tNDiw25xFC.txt", client.O_RDWR|client.O_CREATE)
if(fd804 < 0) {
  panic("open failed")
}


ret = client.Close(fd777)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd773, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd805 := client.Open("/RQKmSwTj6V.txt", client.O_RDWR|client.O_CREATE)
if(fd805 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd794, []byte("fQOswQIzKfLZUv6VS6xjEWVzzmj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) fQOswQIzKfLZUv6VS6xjEWVzzmj

buf, ret = client.Read(fd781, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "f3Q") {
  panic("wrong data returned")
}


ret = client.Close(fd799)
if(ret != 0) {
  panic("close failed")
}


fd806 := client.Open("/CHEKYDs2R7.txt", client.O_RDWR|client.O_CREATE)
if(fd806 < 0) {
  panic("open failed")
}

//fd state: (156) lT7XKr_yQC6HCWPtbdNiw7WWY9zlrnPllVNckGz5q5wh_QGfbbWbermPSJkQrNMqAR8ZxTG8GiMDoZVb9ja8dtoy9vnMCZEQR_kuwcwBLw3ofRAeu55qDF0hyyX2g6crAxk25S2uBvUajSd1q8KHHRwfIWOt

ret = client.Write(fd759, []byte("jF_XTJnZAS2kjwMlxk365XloWokJddmQILW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (191) lT7XKr_yQC6HCWPtbdNiw7WWY9zlrnPllVNckGz5q5wh_QGfbbWbermPSJkQrNMqAR8ZxTG8GiMDoZVb9ja8dtoy9vnMCZEQR_kuwcwBLw3ofRAeu55qDF0hyyX2g6crAxk25S2uBvUajSd1q8KHHRwfIWOtjF_XTJnZAS2kjwMlxk365XloWokJddmQILW

ret = client.Seek(fd781, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Seek(fd803, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd807 := client.Open("/OoNslOuBNK.txt", client.O_RDWR|client.O_CREATE)
if(fd807 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd759, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd785, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}


fd808 := client.Open("/Ue36Yd0iZk.txt", client.O_RDWR|client.O_CREATE)
if(fd808 < 0) {
  panic("open failed")
}


ret = client.Close(fd798)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd780)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd807)
if(ret != 0) {
  panic("close failed")
}


fd809 := client.Open("/ghuzLxuO9g.txt", client.O_RDWR|client.O_CREATE)
if(fd809 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd809, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (2) TNcNBR2ndBLPw8f3Q

ret = client.Write(fd781, []byte("8k20gOg0pFELmoJ6K8KPttkGb_G_T0HyoderTcAbV5azPDzc_OBw5WyLPUmuzg1wM8wSTAHHeiX3X88i73v7I5mCEyf4ISKb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (98) TN8k20gOg0pFELmoJ6K8KPttkGb_G_T0HyoderTcAbV5azPDzc_OBw5WyLPUmuzg1wM8wSTAHHeiX3X88i73v7I5mCEyf4ISKb

buf, ret = client.Read(fd794, 18)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd810 := client.Open("/bscOPxCqaM.txt", client.O_RDWR|client.O_CREATE)
if(fd810 < 0) {
  panic("open failed")
}

//fd state: (44) 12TdoW5oK4Fdre15ATjnGyFMs6lEmWSlq5ofOFWWZ2UIEOldFSU7M_WdNTFNPH798X0MbOXKl1Cf_g4kbtFVQD0qtQ0t5v92P7f2NVepyCf4sRcHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX

ret = client.Write(fd785, []byte("ME7F8XCrC1sRZPl23PlXInHAJH8SDDG0bemSEAXUvrsF9qzhzR54J2WNEBMcu0A7xiq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (111) 12TdoW5oK4Fdre15ATjnGyFMs6lEmWSlq5ofOFWWZ2UIME7F8XCrC1sRZPl23PlXInHAJH8SDDG0bemSEAXUvrsF9qzhzR54J2WNEBMcu0A7xiqHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX

fd811 := client.Open("/Do8eCB4JFH.txt", client.O_RDWR|client.O_CREATE)
if(fd811 < 0) {
  panic("open failed")
}


fd812 := client.Open("/_gnYfmmFcw.txt", client.O_RDWR|client.O_CREATE)
if(fd812 < 0) {
  panic("open failed")
}


fd813 := client.Open("/tNDiw25xFC.txt", client.O_RDWR|client.O_CREATE)
if(fd813 < 0) {
  panic("open failed")
}


ret = client.Seek(fd802, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd810, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "65ocKPNQhErm3i9toTrLia7XXDx0E0ZZ_4gejABr") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd791, 50)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd814 := client.Open("/12bHyglhYH.txt", client.O_RDWR|client.O_CREATE)
if(fd814 < 0) {
  panic("open failed")
}


fd815 := client.Open("/KLDZc9ap3z.txt", client.O_RDWR|client.O_CREATE)
if(fd815 < 0) {
  panic("open failed")
}

//fd state: (15) 65ocKPNQhErm3i9toTrLia7XXDx0E0ZZ_4gejABrgdcNwD8gLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG51V23mKFTZbJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

ret = client.Write(fd797, []byte("_D7cWaefooTa2yn72d8tkQDGn6M3lyKiD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) 65ocKPNQhErm3i9_D7cWaefooTa2yn72d8tkQDGn6M3lyKiDLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG51V23mKFTZbJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

ret = client.Close(fd800)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd813)
if(ret != 0) {
  panic("close failed")
}


fd816 := client.Open("/2pzPsiHvtK.txt", client.O_RDWR|client.O_CREATE)
if(fd816 < 0) {
  panic("open failed")
}


ret = client.Seek(fd796, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd773, []byte("AEFmGU2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (7) AEFmGU2

buf, ret = client.Read(fd796, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "77vtTfiaeYGkxpJNlAbuW00_sEnUSqPTTtgpPBOw855snZLYu4Rl5mO4u5ECilv") {
  panic("wrong data returned")
}


ret = client.Seek(fd781, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}

//fd state: (27) fQOswQIzKfLZUv6VS6xjEWVzzmj

ret = client.Write(fd794, []byte("B563bTFgOW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (37) fQOswQIzKfLZUv6VS6xjEWVzzmjB563bTFgOW

ret = client.Seek(fd791, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd816, 105, client.SEEK_SET)
if(ret != 105) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 105)
  panic("seek failed")
}


ret = client.Seek(fd794, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd817 := client.Open("/ydVPWBuhMs.txt", client.O_RDWR|client.O_CREATE)
if(fd817 < 0) {
  panic("open failed")
}

//fd state: (111) 12TdoW5oK4Fdre15ATjnGyFMs6lEmWSlq5ofOFWWZ2UIME7F8XCrC1sRZPl23PlXInHAJH8SDDG0bemSEAXUvrsF9qzhzR54J2WNEBMcu0A7xiqHTCYu9HWbmVZHFNQe_2O8NYhUa1gr6DIX

ret = client.Write(fd785, []byte("b8CfVzJKAPaqWTsd7Dz3qxbJFjBnKZjS4WH6VFv1M2LBrYv04QwyrbV0cjCD1zxfsRySDuvkmkQMOBMO5VGfrt0q2xkGzroT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (207) 12TdoW5oK4Fdre15ATjnGyFMs6lEmWSlq5ofOFWWZ2UIME7F8XCrC1sRZPl23PlXInHAJH8SDDG0bemSEAXUvrsF9qzhzR54J2WNEBMcu0A7xiqb8CfVzJKAPaqWTsd7Dz3qxbJFjBnKZjS4WH6VFv1M2LBrYv04QwyrbV0cjCD1zxfsRySDuvkmkQMOBMO5VGfrt0q2xkGzroT
//fd state: (83) QJv2H3RdVSn8imaCN0LQ77vtTfiaeYGkxpJNlAbuW00_sEnUSqPTTtgpPBOw855snZLYu4Rl5mO4u5ECilv9s_jP93cFd2v0HwbiemgaeruhJ

ret = client.Write(fd796, []byte("HgEzWpDT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) QJv2H3RdVSn8imaCN0LQ77vtTfiaeYGkxpJNlAbuW00_sEnUSqPTTtgpPBOw855snZLYu4Rl5mO4u5ECilvHgEzWpDTFd2v0HwbiemgaeruhJ

ret = client.Close(fd806)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd802, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd812, 72, client.SEEK_SET)
if(ret != 72) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 72)
  panic("seek failed")
}


ret = client.Seek(fd805, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd818 := client.Open("/jioZXwkhDi.txt", client.O_RDWR|client.O_CREATE)
if(fd818 < 0) {
  panic("open failed")
}

//fd state: (48) 65ocKPNQhErm3i9_D7cWaefooTa2yn72d8tkQDGn6M3lyKiDLrC76ZwK72ESiPdlKgT1aujlNJXrG0_j2a8607doiUFM7NBhH6iXzypGG51V23mKFTZbJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

ret = client.Write(fd797, []byte("reQhvIpahKMV7tFa1USvE5cNfjXKSRgEGFj2mBLGgveJJG1WVyvo1vU7h4y5WP_nhp_A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (116) 65ocKPNQhErm3i9_D7cWaefooTa2yn72d8tkQDGn6M3lyKiDreQhvIpahKMV7tFa1USvE5cNfjXKSRgEGFj2mBLGgveJJG1WVyvo1vU7h4y5WP_nhp_AJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

fd819 := client.Open("/XQupscYfNb.txt", client.O_RDWR|client.O_CREATE)
if(fd819 < 0) {
  panic("open failed")
}


fd820 := client.Open("/NnsTDHvSVb.txt", client.O_RDWR|client.O_CREATE)
if(fd820 < 0) {
  panic("open failed")
}


ret = client.Close(fd818)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd804, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd814, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd819, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd821 := client.Open("/TKrqbB__4d.txt", client.O_RDWR|client.O_CREATE)
if(fd821 < 0) {
  panic("open failed")
}


ret = client.Close(fd816)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd809, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd812)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd794)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd810)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd770, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "U7mo2dvYm_Vd5d1FHkS") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd821, []byte("ab3tbvnEjQ4ONOUxs3PDzAaL8cvc7QjKpvCB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) ab3tbvnEjQ4ONOUxs3PDzAaL8cvc7QjKpvCB

fd822 := client.Open("/_9uGqB8bKD.txt", client.O_RDWR|client.O_CREATE)
if(fd822 < 0) {
  panic("open failed")
}


fd823 := client.Open("/ck3kOvcScG.txt", client.O_RDWR|client.O_CREATE)
if(fd823 < 0) {
  panic("open failed")
}

//fd state: (56) 8T291_jIaWt1rCJLvqEoCJsMc_KDMgGjJJMiGU7mo2dvYm_Vd5d1FHkS7P_ONzJspXAMGWBj20WgrAqlp7s1gH6QhD9uoiXo_8yfcXmbJXcw6veTm9LzjvcKX5nryXvLvHX3a4Dhg9py6PMPcQCTYnprnLcOhhMgBSL5GvCmVmTHVc2on53N3_kesb

ret = client.Write(fd770, []byte("xTXOqWrMTxygkpRGk8usaAz7rDjyjOD5ACLAbm71klXrLgxZ6LWJzpDVuW7"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) 8T291_jIaWt1rCJLvqEoCJsMc_KDMgGjJJMiGU7mo2dvYm_Vd5d1FHkSxTXOqWrMTxygkpRGk8usaAz7rDjyjOD5ACLAbm71klXrLgxZ6LWJzpDVuW7zjvcKX5nryXvLvHX3a4Dhg9py6PMPcQCTYnprnLcOhhMgBSL5GvCmVmTHVc2on53N3_kesb

fd824 := client.Open("/kfZ_4vRm9s.txt", client.O_RDWR|client.O_CREATE)
if(fd824 < 0) {
  panic("open failed")
}


ret = client.Close(fd817)
if(ret != 0) {
  panic("close failed")
}

//fd state: (4) TN8k20gOg0pFELmoJ6K8KPttkGb_G_T0HyoderTcAbV5azPDzc_OBw5WyLPUmuzg1wM8wSTAHHeiX3X88i73v7I5mCEyf4ISKb

ret = client.Write(fd781, []byte("VvnxmZlrz9LZIx0qchpsVK2vGk8Q_jtNYOHPPl4PvD7Kfwz0l8C3HB8UBL3iua857oa_Lj9yqJljlBUWI85p"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) TN8kVvnxmZlrz9LZIx0qchpsVK2vGk8Q_jtNYOHPPl4PvD7Kfwz0l8C3HB8UBL3iua857oa_Lj9yqJljlBUWI85pmCEyf4ISKb

fd825 := client.Open("/WxPQuN5oUi.txt", client.O_RDWR|client.O_CREATE)
if(fd825 < 0) {
  panic("open failed")
}


ret = client.Seek(fd770, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd809, []byte("o4NLujFJmcfM_gx9YigRiVk4r5dvrd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) o4NLujFJmcfM_gx9YigRiVk4r5dvrd

ret = client.Seek(fd805, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd821, 7)
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


ret = client.Seek(fd808, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd823)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd819, 87)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd826 := client.Open("/EyEjU2mTvD.txt", client.O_RDWR|client.O_CREATE)
if(fd826 < 0) {
  panic("open failed")
}


fd827 := client.Open("/D2dL3T9Vmw.txt", client.O_RDWR|client.O_CREATE)
if(fd827 < 0) {
  panic("open failed")
}


ret = client.Seek(fd803, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd785, 184, client.SEEK_SET)
if(ret != 184) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 184)
  panic("seek failed")
}


buf, ret = client.Read(fd808, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (88) TN8kVvnxmZlrz9LZIx0qchpsVK2vGk8Q_jtNYOHPPl4PvD7Kfwz0l8C3HB8UBL3iua857oa_Lj9yqJljlBUWI85pmCEyf4ISKb

ret = client.Write(fd781, []byte("0zZFOmkKOImptSUrIPW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) TN8kVvnxmZlrz9LZIx0qchpsVK2vGk8Q_jtNYOHPPl4PvD7Kfwz0l8C3HB8UBL3iua857oa_Lj9yqJljlBUWI85p0zZFOmkKOImptSUrIPW

fd828 := client.Open("/lfyrzZTJQG.txt", client.O_RDWR|client.O_CREATE)
if(fd828 < 0) {
  panic("open failed")
}


ret = client.Close(fd791)
if(ret != 0) {
  panic("close failed")
}


fd829 := client.Open("/qQGpJOi0Ag.txt", client.O_RDWR|client.O_CREATE)
if(fd829 < 0) {
  panic("open failed")
}


ret = client.Close(fd815)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd825)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd828)
if(ret != 0) {
  panic("close failed")
}

//fd state: (30) o4NLujFJmcfM_gx9YigRiVk4r5dvrd

ret = client.Write(fd809, []byte("PpYVmXUCZpEktcOrnv0V5swWqRJHxWK1YHGDZp3gXQTpWRtCyt"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) o4NLujFJmcfM_gx9YigRiVk4r5dvrdPpYVmXUCZpEktcOrnv0V5swWqRJHxWK1YHGDZp3gXQTpWRtCyt

fd830 := client.Open("/su0syYXcBE.txt", client.O_RDWR|client.O_CREATE)
if(fd830 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd779, 96)
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


buf, ret = client.Read(fd809, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 6qgCJlnFlWwI5JKVCLSKsH8hfaJOZdHq6k1jI5ggB1eipgozBImi

ret = client.Write(fd830, []byte("G9ROa74a0anrO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) G9ROa74a0anrOJKVCLSKsH8hfaJOZdHq6k1jI5ggB1eipgozBImi
//fd state: (0) 

ret = client.Write(fd803, []byte("70fK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) 70fK
//fd state: (36) ab3tbvnEjQ4ONOUxs3PDzAaL8cvc7QjKpvCB

ret = client.Write(fd821, []byte("S42vmKe7DdWPk3EiUUubWB0GOEv0BN4zA_nveMo_6Lko9v9yOlIxDqyzY3pJpbxRo1_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) ab3tbvnEjQ4ONOUxs3PDzAaL8cvc7QjKpvCBS42vmKe7DdWPk3EiUUubWB0GOEv0BN4zA_nveMo_6Lko9v9yOlIxDqyzY3pJpbxRo1_

ret = client.Close(fd829)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd803, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


ret = client.Seek(fd819, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd809, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


fd831 := client.Open("/NPOgKkI4Qs.txt", client.O_RDWR|client.O_CREATE)
if(fd831 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd822, []byte("nXatej4H2wF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) nXatej4H2wF

fd832 := client.Open("/ck3kOvcScG.txt", client.O_RDWR|client.O_CREATE)
if(fd832 < 0) {
  panic("open failed")
}

//fd state: (103) ab3tbvnEjQ4ONOUxs3PDzAaL8cvc7QjKpvCBS42vmKe7DdWPk3EiUUubWB0GOEv0BN4zA_nveMo_6Lko9v9yOlIxDqyzY3pJpbxRo1_

ret = client.Write(fd821, []byte("6CvoDOq14bRukM1eow4fyllUKouP4s2MKxIna"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (140) ab3tbvnEjQ4ONOUxs3PDzAaL8cvc7QjKpvCBS42vmKe7DdWPk3EiUUubWB0GOEv0BN4zA_nveMo_6Lko9v9yOlIxDqyzY3pJpbxRo1_6CvoDOq14bRukM1eow4fyllUKouP4s2MKxIna
//fd state: (0) 

ret = client.Write(fd808, []byte("DL8hwxw3j2vtubdUOg3a_kIxqDomJK0RbFq4hrL_KK064FdB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) DL8hwxw3j2vtubdUOg3a_kIxqDomJK0RbFq4hrL_KK064FdB

buf, ret = client.Read(fd770, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2dv") {
  panic("wrong data returned")
}


ret = client.Close(fd819)
if(ret != 0) {
  panic("close failed")
}


fd833 := client.Open("/6G9beguEkr.txt", client.O_RDWR|client.O_CREATE)
if(fd833 < 0) {
  panic("open failed")
}


ret = client.Seek(fd802, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd826, []byte("zSF1r"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (5) zSF1r

buf, ret = client.Read(fd781, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd775)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) MWfJD4J_HA9lvBpS0dKaJs9VAsmiQy7rg7mlGFWHiTouwaph_KvVIYpj_7oytq8A7OHYsjUhjGUIbGEFeTcNgWg_hNl1OAr7x1i0R72hBYlppsm8js0ZNLN8TX7oZtpQjrAQIDUOCcPOQskdU4v3LL3qDrtJpTOQ4qpX

ret = client.Write(fd832, []byte("5jZUzTgGBRd05AtxM2Qm73lMg0MMK45e7nx7MNobZOLLXCLHxHeuRiQmnfUyG6mDRN30B8QGxCmi99XG3cCHrD82tP_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) 5jZUzTgGBRd05AtxM2Qm73lMg0MMK45e7nx7MNobZOLLXCLHxHeuRiQmnfUyG6mDRN30B8QGxCmi99XG3cCHrD82tP_1OAr7x1i0R72hBYlppsm8js0ZNLN8TX7oZtpQjrAQIDUOCcPOQskdU4v3LL3qDrtJpTOQ4qpX

fd834 := client.Open("/j_YKWOwzTu.txt", client.O_RDWR|client.O_CREATE)
if(fd834 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd802, []byte("XbYREHH_P6VoK3zSLC9c3SEMsSZ3NAk157GBRUOjfH3frmdwSGyK3CzbDdUTTvir_neqJWcnKzWSWmyK9EQXFrHjWmb8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) XbYREHH_P6VoK3zSLC9c3SEMsSZ3NAk157GBRUOjfH3frmdwSGyK3CzbDdUTTvir_neqJWcnKzWSWmyK9EQXFrHjWmb8
//fd state: (91) 5jZUzTgGBRd05AtxM2Qm73lMg0MMK45e7nx7MNobZOLLXCLHxHeuRiQmnfUyG6mDRN30B8QGxCmi99XG3cCHrD82tP_1OAr7x1i0R72hBYlppsm8js0ZNLN8TX7oZtpQjrAQIDUOCcPOQskdU4v3LL3qDrtJpTOQ4qpX

ret = client.Write(fd832, []byte("e5tJoKh87rZ1ADJ_NDnEgL"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) 5jZUzTgGBRd05AtxM2Qm73lMg0MMK45e7nx7MNobZOLLXCLHxHeuRiQmnfUyG6mDRN30B8QGxCmi99XG3cCHrD82tP_e5tJoKh87rZ1ADJ_NDnEgLs0ZNLN8TX7oZtpQjrAQIDUOCcPOQskdU4v3LL3qDrtJpTOQ4qpX

fd835 := client.Open("/JYA1gHZys5.txt", client.O_RDWR|client.O_CREATE)
if(fd835 < 0) {
  panic("open failed")
}


ret = client.Close(fd820)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd827, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd836 := client.Open("/j6CCRdMK5y.txt", client.O_RDWR|client.O_CREATE)
if(fd836 < 0) {
  panic("open failed")
}


fd837 := client.Open("/AX1EuCJwHs.txt", client.O_RDWR|client.O_CREATE)
if(fd837 < 0) {
  panic("open failed")
}


ret = client.Seek(fd834, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd811, 24, client.SEEK_SET)
if(ret != 24) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 24)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd834, []byte("2UXzfwyUHy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (10) 2UXzfwyUHy

buf, ret = client.Read(fd835, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd835, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd838 := client.Open("/4NexkYmnGj.txt", client.O_RDWR|client.O_CREATE)
if(fd838 < 0) {
  panic("open failed")
}


fd839 := client.Open("/XM7E0K2UJu.txt", client.O_RDWR|client.O_CREATE)
if(fd839 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd831, 38)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (107) TN8kVvnxmZlrz9LZIx0qchpsVK2vGk8Q_jtNYOHPPl4PvD7Kfwz0l8C3HB8UBL3iua857oa_Lj9yqJljlBUWI85p0zZFOmkKOImptSUrIPW

ret = client.Write(fd781, []byte("MUEO0xrycDuInc_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (122) TN8kVvnxmZlrz9LZIx0qchpsVK2vGk8Q_jtNYOHPPl4PvD7Kfwz0l8C3HB8UBL3iua857oa_Lj9yqJljlBUWI85p0zZFOmkKOImptSUrIPWMUEO0xrycDuInc_

buf, ret = client.Read(fd836, 9)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uQcGj0Bbl") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd834, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd770, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Ym_Vd5") {
  panic("wrong data returned")
}


ret = client.Close(fd826)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd797, 12)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "J8lSPo2bPH_n") {
  panic("wrong data returned")
}


ret = client.Close(fd796)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd797, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "t1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5Uk") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd831, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd837)
if(ret != 0) {
  panic("close failed")
}


fd840 := client.Open("/OIVNzc67re.txt", client.O_RDWR|client.O_CREATE)
if(fd840 < 0) {
  panic("open failed")
}


fd841 := client.Open("/zK0Bie8jKe.txt", client.O_RDWR|client.O_CREATE)
if(fd841 < 0) {
  panic("open failed")
}


ret = client.Close(fd824)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd770, 5, client.SEEK_SET)
if(ret != 5) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 5)
  panic("seek failed")
}


ret = client.Close(fd803)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd781, 114, client.SEEK_SET)
if(ret != 114) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 114)
  panic("seek failed")
}


buf, ret = client.Read(fd831, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd770, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_jIaWt1rCJLvqEoCJsMc_") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd802, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd822)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd835, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (24) 0jRvYQj7URR5zsbLHJzOmVteqCW6CppHJm3iRBRnGKzd0RrigjvXEzOuWqrzLjHrd5Uq3mSv_3xh9muk4_1gE7JLGUF9QtxNi5WMn8OZpqnqTFyeHuIKUwh3Gip3iWAXgQ4rbnQSks9P

ret = client.Write(fd811, []byte("xz0rIKbMjCfAlacKh7NTv2MtksEsOUESRUk6Yhcv_V7DrFHhE_t1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) 0jRvYQj7URR5zsbLHJzOmVtexz0rIKbMjCfAlacKh7NTv2MtksEsOUESRUk6Yhcv_V7DrFHhE_t19muk4_1gE7JLGUF9QtxNi5WMn8OZpqnqTFyeHuIKUwh3Gip3iWAXgQ4rbnQSks9P

ret = client.Close(fd770)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd811, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


ret = client.Seek(fd781, 71, client.SEEK_SET)
if(ret != 71) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 71)
  panic("seek failed")
}


buf, ret = client.Read(fd838, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8SWfklF5EBz6iAHz_JL8fFg1qPdv6MgjkZqhh0O687Jmk3fcEK") {
  panic("wrong data returned")
}

//fd state: (7) AEFmGU2

ret = client.Write(fd773, []byte("OnjZ7gZZW8QtStMJt6MwlU1ahSqGV_6K3yKZ3dwxOhIAZKLd6GO7m3D5YYScx6yEHaljwa8RzIbT20a06czdu67g"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) AEFmGU2OnjZ7gZZW8QtStMJt6MwlU1ahSqGV_6K3yKZ3dwxOhIAZKLd6GO7m3D5YYScx6yEHaljwa8RzIbT20a06czdu67g
//fd state: (92) XbYREHH_P6VoK3zSLC9c3SEMsSZ3NAk157GBRUOjfH3frmdwSGyK3CzbDdUTTvir_neqJWcnKzWSWmyK9EQXFrHjWmb8

ret = client.Write(fd802, []byte("223DM6tb4GZesrVgNdIlU1y51ANAE8p"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) XbYREHH_P6VoK3zSLC9c3SEMsSZ3NAk157GBRUOjfH3frmdwSGyK3CzbDdUTTvir_neqJWcnKzWSWmyK9EQXFrHjWmb8223DM6tb4GZesrVgNdIlU1y51ANAE8p

buf, ret = client.Read(fd781, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "_Lj9yqJljlBUWI85p0zZFOmkKOImptSUrIPWMUEO0xrycDuInc_") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd840, []byte("hL7_gyQ9_zQZdQQOW08Cm4l2Jql3zopmHrZc2Cv21Qd22pyNkTOtJAd1ic"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) hL7_gyQ9_zQZdQQOW08Cm4l2Jql3zopmHrZc2Cv21Qd22pyNkTOtJAd1ic

buf, ret = client.Read(fd802, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd804)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd831, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd802, 110, client.SEEK_SET)
if(ret != 110) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 110)
  panic("seek failed")
}


ret = client.Seek(fd785, 164, client.SEEK_SET)
if(ret != 164) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 164)
  panic("seek failed")
}


fd842 := client.Open("/zWMKnH6tPr.txt", client.O_RDWR|client.O_CREATE)
if(fd842 < 0) {
  panic("open failed")
}


ret = client.Close(fd841)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd830)
if(ret != 0) {
  panic("close failed")
}


fd843 := client.Open("/LQ9tyeq1jT.txt", client.O_RDWR|client.O_CREATE)
if(fd843 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd805, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd844 := client.Open("/LoTeuE2wHL.txt", client.O_RDWR|client.O_CREATE)
if(fd844 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd779, []byte("3PfWCpfYadSTTPFmFB3AQIba91ehZCV4uGF_AEUtgH4IDe03TJpKNH5wZ1SnTKvMBIk2144"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (71) 3PfWCpfYadSTTPFmFB3AQIba91ehZCV4uGF_AEUtgH4IDe03TJpKNH5wZ1SnTKvMBIk2144

buf, ret = client.Read(fd839, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd844, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vHad_07sUSPQYs3WDfHxRuTjNFxvKBDMpIjMpR5l6") {
  panic("wrong data returned")
}


ret = client.Close(fd843)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd785, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "bV0cjCD1zxf") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd808, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd779)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd808)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd781, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd845 := client.Open("/g1InCUZb90.txt", client.O_RDWR|client.O_CREATE)
if(fd845 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd821, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd773)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd821)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd831, []byte("O5V03IGsQuqA5332QKk2R"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) O5V03IGsQuqA5332QKk2R

buf, ret = client.Read(fd839, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd781, 26, client.SEEK_SET)
if(ret != 26) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 26)
  panic("seek failed")
}


fd846 := client.Open("/9YeGzW1_xA.txt", client.O_RDWR|client.O_CREATE)
if(fd846 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd827, []byte("PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVFavfRWPp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVFavfRWPp

ret = client.Seek(fd835, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (46) 0jRvYQj7URR5zsbLHJzOmVtexz0rIKbMjCfAlacKh7NTv2MtksEsOUESRUk6Yhcv_V7DrFHhE_t19muk4_1gE7JLGUF9QtxNi5WMn8OZpqnqTFyeHuIKUwh3Gip3iWAXgQ4rbnQSks9P

ret = client.Write(fd811, []byte("TdoT9mz6mfO55uY0o3B"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) 0jRvYQj7URR5zsbLHJzOmVtexz0rIKbMjCfAlacKh7NTv2TdoT9mz6mfO55uY0o3BV7DrFHhE_t19muk4_1gE7JLGUF9QtxNi5WMn8OZpqnqTFyeHuIKUwh3Gip3iWAXgQ4rbnQSks9P

fd847 := client.Open("/L3M3DlnPNb.txt", client.O_RDWR|client.O_CREATE)
if(fd847 < 0) {
  panic("open failed")
}


ret = client.Close(fd811)
if(ret != 0) {
  panic("close failed")
}

//fd state: (113) 5jZUzTgGBRd05AtxM2Qm73lMg0MMK45e7nx7MNobZOLLXCLHxHeuRiQmnfUyG6mDRN30B8QGxCmi99XG3cCHrD82tP_e5tJoKh87rZ1ADJ_NDnEgLs0ZNLN8TX7oZtpQjrAQIDUOCcPOQskdU4v3LL3qDrtJpTOQ4qpX

ret = client.Write(fd832, []byte("dXyHkZuTcvZIGWfDJu7fcG501nYpfQXgtgVwaum7XAjE_4wm6DA37_tjqFe4RV5cYlpgljiMGrpUiN_VpLxk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (197) 5jZUzTgGBRd05AtxM2Qm73lMg0MMK45e7nx7MNobZOLLXCLHxHeuRiQmnfUyG6mDRN30B8QGxCmi99XG3cCHrD82tP_e5tJoKh87rZ1ADJ_NDnEgLdXyHkZuTcvZIGWfDJu7fcG501nYpfQXgtgVwaum7XAjE_4wm6DA37_tjqFe4RV5cYlpgljiMGrpUiN_VpLxk

ret = client.Seek(fd797, 145, client.SEEK_SET)
if(ret != 145) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 145)
  panic("seek failed")
}


ret = client.Seek(fd797, 69, client.SEEK_SET)
if(ret != 69) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 69)
  panic("seek failed")
}


fd848 := client.Open("/CpBTPLAg78.txt", client.O_RDWR|client.O_CREATE)
if(fd848 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd781, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2vGk8Q") {
  panic("wrong data returned")
}


ret = client.Close(fd847)
if(ret != 0) {
  panic("close failed")
}

//fd state: (69) 65ocKPNQhErm3i9_D7cWaefooTa2yn72d8tkQDGn6M3lyKiDreQhvIpahKMV7tFa1USvE5cNfjXKSRgEGFj2mBLGgveJJG1WVyvo1vU7h4y5WP_nhp_AJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

ret = client.Write(fd797, []byte("LrS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) 65ocKPNQhErm3i9_D7cWaefooTa2yn72d8tkQDGn6M3lyKiDreQhvIpahKMV7tFa1USvELrSfjXKSRgEGFj2mBLGgveJJG1WVyvo1vU7h4y5WP_nhp_AJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY
//fd state: (21) O5V03IGsQuqA5332QKk2R

ret = client.Write(fd831, []byte("L4OAXX6N_6uuU83mdVXR7_g6dM_qlpZuT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) O5V03IGsQuqA5332QKk2RL4OAXX6N_6uuU83mdVXR7_g6dM_qlpZuT
//fd state: (0) 

ret = client.Write(fd833, []byte("IntWhIPvD_LgLEkEbf0U"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) IntWhIPvD_LgLEkEbf0U

buf, ret = client.Read(fd832, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (110) XbYREHH_P6VoK3zSLC9c3SEMsSZ3NAk157GBRUOjfH3frmdwSGyK3CzbDdUTTvir_neqJWcnKzWSWmyK9EQXFrHjWmb8223DM6tb4GZesrVgNdIlU1y51ANAE8p

ret = client.Write(fd802, []byte("f1juU0YJT8qlcY4MHuVMFi5Il8xt5GLGlaLutps5bdY7ZHyY4BWR6Y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) XbYREHH_P6VoK3zSLC9c3SEMsSZ3NAk157GBRUOjfH3frmdwSGyK3CzbDdUTTvir_neqJWcnKzWSWmyK9EQXFrHjWmb8223DM6tb4GZesrVgNdf1juU0YJT8qlcY4MHuVMFi5Il8xt5GLGlaLutps5bdY7ZHyY4BWR6Y

fd849 := client.Open("/bCrAMTyNaU.txt", client.O_RDWR|client.O_CREATE)
if(fd849 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd846, []byte("hegnpOrEEOMWLW7aTmuuuD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) hegnpOrEEOMWLW7aTmuuuD

buf, ret = client.Read(fd809, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "9YigRiVk4r5dvrdPpYVmXUCZpEktcOrnv0V5swWqRJHxWK1YHGDZp3gXQTpWRtCyt") {
  panic("wrong data returned")
}


fd850 := client.Open("/qQGpJOi0Ag.txt", client.O_RDWR|client.O_CREATE)
if(fd850 < 0) {
  panic("open failed")
}

//fd state: (9) uQcGj0BblK_1ZPvfD0m9abQNxxSRtx6OGh9ZOMQyvucGwm7gigFgjHN3ONFklMqriXpVyYvY0z8k3OwqD4fAAuo08TsPHAcfyhS2Ph

ret = client.Write(fd836, []byte("hh9DjwqUPXXIokuxysUBVJry6qDWJmed2FRtDhLDwoNLSShwFxOEBcQZ1SFRESJzdk6qDzTBj5GqJr1scYE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) uQcGj0Bblhh9DjwqUPXXIokuxysUBVJry6qDWJmed2FRtDhLDwoNLSShwFxOEBcQZ1SFRESJzdk6qDzTBj5GqJr1scYEHAcfyhS2Ph

ret = client.Seek(fd844, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}

//fd state: (54) O5V03IGsQuqA5332QKk2RL4OAXX6N_6uuU83mdVXR7_g6dM_qlpZuT

ret = client.Write(fd831, []byte("bTBboyUbcYHb62vmOMzkfM3LjrTvRq8iKvOuTV4UW_8s4T2EQepcIHFGTjSpxxGnZsIQy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (123) O5V03IGsQuqA5332QKk2RL4OAXX6N_6uuU83mdVXR7_g6dM_qlpZuTbTBboyUbcYHb62vmOMzkfM3LjrTvRq8iKvOuTV4UW_8s4T2EQepcIHFGTjSpxxGnZsIQy
//fd state: (50) 8SWfklF5EBz6iAHz_JL8fFg1qPdv6MgjkZqhh0O687Jmk3fcEK

ret = client.Write(fd838, []byte("lGgovDOlomLtuvxHCXMHpm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) 8SWfklF5EBz6iAHz_JL8fFg1qPdv6MgjkZqhh0O687Jmk3fcEKlGgovDOlomLtuvxHCXMHpm

fd851 := client.Open("/1x4Po35g3G.txt", client.O_RDWR|client.O_CREATE)
if(fd851 < 0) {
  panic("open failed")
}


ret = client.Close(fd835)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd805, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd844, 177, client.SEEK_SET)
if(ret != 177) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 177)
  panic("seek failed")
}


ret = client.Seek(fd831, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}

//fd state: (72) 65ocKPNQhErm3i9_D7cWaefooTa2yn72d8tkQDGn6M3lyKiDreQhvIpahKMV7tFa1USvELrSfjXKSRgEGFj2mBLGgveJJG1WVyvo1vU7h4y5WP_nhp_AJ8lSPo2bPH_nt1ODUjb6O09WXjSxKJYc03tV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

ret = client.Write(fd797, []byte("qmFO97N8KJ1MrpYHuKFdphPCfMy2AjOzmpI9NSBrWRqQVh3slkjCjuny_DCHS44H9JSJX6dRP62f6m"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) 65ocKPNQhErm3i9_D7cWaefooTa2yn72d8tkQDGn6M3lyKiDreQhvIpahKMV7tFa1USvELrSqmFO97N8KJ1MrpYHuKFdphPCfMy2AjOzmpI9NSBrWRqQVh3slkjCjuny_DCHS44H9JSJX6dRP62f6mtV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

fd852 := client.Open("/9dluZ9EKXT.txt", client.O_RDWR|client.O_CREATE)
if(fd852 < 0) {
  panic("open failed")
}

//fd state: (80) o4NLujFJmcfM_gx9YigRiVk4r5dvrdPpYVmXUCZpEktcOrnv0V5swWqRJHxWK1YHGDZp3gXQTpWRtCyt

ret = client.Write(fd809, []byte("_rOn86WJ8hydAkMS0mfwHZ4o50G4K0Ii58MuxJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (118) o4NLujFJmcfM_gx9YigRiVk4r5dvrdPpYVmXUCZpEktcOrnv0V5swWqRJHxWK1YHGDZp3gXQTpWRtCyt_rOn86WJ8hydAkMS0mfwHZ4o50G4K0Ii58MuxJ

ret = client.Seek(fd832, 129, client.SEEK_SET)
if(ret != 129) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 129)
  panic("seek failed")
}


fd853 := client.Open("/KuBxwFzoD6.txt", client.O_RDWR|client.O_CREATE)
if(fd853 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd842, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd854 := client.Open("/a6PEzbIRX2.txt", client.O_RDWR|client.O_CREATE)
if(fd854 < 0) {
  panic("open failed")
}

//fd state: (0) Xlv3T0SCSl4qyWfXC5OSppB4wzkmP2P1E2hgmfBocu3Xa8LUU292eZBp84ZnmSvMayWnPIMAC25nm5j9Mp555cYTfWLfJ4EaCL16wmxxiTsTOfWGJ4Td9zvQbTg7xnHHUTJN4YEblSj8qp9zXD0xUDjIqYWP1sJvWAsqBwsSLcrAn6rdXsYk99PJzu4AqPS50NvIFKtfdf9YO3Ug59ZuOXqE0kpFMbSXnrQIZwHvFW5OJ1d7gL6VNvlfXuKCTxvHF2ktW2AwrGosgl_pDZTMSES3ZAVIcC1nuiZ2GASbYZX1K08qYcdkWNK1Znmy0

ret = client.Write(fd852, []byte("4DRHOxHDXVLrmod0MHCxwDM1m00pCvhWSf51rWIY7CV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) 4DRHOxHDXVLrmod0MHCxwDM1m00pCvhWSf51rWIY7CVXa8LUU292eZBp84ZnmSvMayWnPIMAC25nm5j9Mp555cYTfWLfJ4EaCL16wmxxiTsTOfWGJ4Td9zvQbTg7xnHHUTJN4YEblSj8qp9zXD0xUDjIqYWP1sJvWAsqBwsSLcrAn6rdXsYk99PJzu4AqPS50NvIFKtfdf9YO3Ug59ZuOXqE0kpFMbSXnrQIZwHvFW5OJ1d7gL6VNvlfXuKCTxvHF2ktW2AwrGosgl_pDZTMSES3ZAVIcC1nuiZ2GASbYZX1K08qYcdkWNK1Znmy0

buf, ret = client.Read(fd845, 65)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd839, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd855 := client.Open("/TDOT533S0s.txt", client.O_RDWR|client.O_CREATE)
if(fd855 < 0) {
  panic("open failed")
}


ret = client.Seek(fd850, 25, client.SEEK_SET)
if(ret != 25) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 25)
  panic("seek failed")
}


fd856 := client.Open("/OeXkXPuXV6.txt", client.O_RDWR|client.O_CREATE)
if(fd856 < 0) {
  panic("open failed")
}


fd857 := client.Open("/e2AgzPM538.txt", client.O_RDWR|client.O_CREATE)
if(fd857 < 0) {
  panic("open failed")
}


ret = client.Close(fd849)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd838, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) V2dzZ8G8te7ld6bPiY2193tKSlsrkCb

ret = client.Write(fd848, []byte("5915fhJg7V33TuoaorAwiIxIicDu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) 5915fhJg7V33TuoaorAwiIxIicDukCb

ret = client.Close(fd848)
if(ret != 0) {
  panic("close failed")
}


fd858 := client.Open("/fu3RTzcgJU.txt", client.O_RDWR|client.O_CREATE)
if(fd858 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd805, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (10) 2UXzfwyUHy

ret = client.Write(fd834, []byte("0Ef2K4TXqX2kcul2sa33D_mLK9SaZX4bFRGWy7F6yxKXah2G0bFos5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) 2UXzfwyUHy0Ef2K4TXqX2kcul2sa33D_mLK9SaZX4bFRGWy7F6yxKXah2G0bFos5

ret = client.Seek(fd832, 151, client.SEEK_SET)
if(ret != 151) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 151)
  panic("seek failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd785)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd852, 204, client.SEEK_SET)
if(ret != 204) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 204)
  panic("seek failed")
}


buf, ret = client.Read(fd856, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd851, []byte("9mRErHP3CsvWao9tvqx5qUuhQSFWEL4Fhq2QRaIAOBlGULf7uAHLFkgje35Eoc2Wsw3e6b6nV0BcRdLbnb1gp5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) 9mRErHP3CsvWao9tvqx5qUuhQSFWEL4Fhq2QRaIAOBlGULf7uAHLFkgje35Eoc2Wsw3e6b6nV0BcRdLbnb1gp5

ret = client.Close(fd833)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd854, []byte("awLNIzTV8Mmpqn4PlFitjVUHptjFZfpa6ukSulojktNGKQNKwLfGs5FudhPh6kZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) awLNIzTV8Mmpqn4PlFitjVUHptjFZfpa6ukSulojktNGKQNKwLfGs5FudhPh6kZ

ret = client.Close(fd854)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd846, 1, client.SEEK_SET)
if(ret != 1) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 1)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd855, []byte("5qtNU8CBaURHjj28Ry72B4ZSlZ5XH_MaodsSBrReGq80kAUBh5bugU8OXABcAwSD5Qjg9N3_o_OT23bhC9odNKt0G3aPE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (93) 5qtNU8CBaURHjj28Ry72B4ZSlZ5XH_MaodsSBrReGq80kAUBh5bugU8OXABcAwSD5Qjg9N3_o_OT23bhC9odNKt0G3aPE

buf, ret = client.Read(fd839, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd832, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "m7XAjE_4wm6DA37_tjqFe4RV5cYlpgljiMGrpUiN_VpLxk") {
  panic("wrong data returned")
}


ret = client.Seek(fd781, 121, client.SEEK_SET)
if(ret != 121) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 121)
  panic("seek failed")
}


fd859 := client.Open("/EcKZwuR0vW.txt", client.O_RDWR|client.O_CREATE)
if(fd859 < 0) {
  panic("open failed")
}


ret = client.Seek(fd846, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd859)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd846, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WLW7aTmuuuD") {
  panic("wrong data returned")
}


fd860 := client.Open("/imHM3T_ti2.txt", client.O_RDWR|client.O_CREATE)
if(fd860 < 0) {
  panic("open failed")
}


fd861 := client.Open("/J7dw2n8Y5L.txt", client.O_RDWR|client.O_CREATE)
if(fd861 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd861, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd845, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd862 := client.Open("/JyCJqcy6n8.txt", client.O_RDWR|client.O_CREATE)
if(fd862 < 0) {
  panic("open failed")
}


ret = client.Close(fd781)
if(ret != 0) {
  panic("close failed")
}


fd863 := client.Open("/MF_bH4w72p.txt", client.O_RDWR|client.O_CREATE)
if(fd863 < 0) {
  panic("open failed")
}


ret = client.Seek(fd827, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


buf, ret = client.Read(fd846, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd844, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "fdidUdH7wFY5IollSbIHNPdt3iCq3k") {
  panic("wrong data returned")
}


ret = client.Close(fd802)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd846, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd864 := client.Open("/QjQHErqwDF.txt", client.O_RDWR|client.O_CREATE)
if(fd864 < 0) {
  panic("open failed")
}


ret = client.Close(fd832)
if(ret != 0) {
  panic("close failed")
}


fd865 := client.Open("/cCJFYITFef.txt", client.O_RDWR|client.O_CREATE)
if(fd865 < 0) {
  panic("open failed")
}


ret = client.Seek(fd834, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


ret = client.Close(fd852)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd856, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd809, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd846, 59)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd850, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lu4DCoOwvDJR1t6gx1gXHOlgVhZq7ncm7teCWWBqAI3xiP2KC8tRR") {
  panic("wrong data returned")
}


ret = client.Close(fd846)
if(ret != 0) {
  panic("close failed")
}


fd866 := client.Open("/k57M1a0Y2S.txt", client.O_RDWR|client.O_CREATE)
if(fd866 < 0) {
  panic("open failed")
}


ret = client.Close(fd864)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) uIWRBUTVmSLW6MJWvcmVbPT4sWhRVdKw7DWUJa2jJCNMrXQ6pLO7uQSsmCiMpNMH8JwgyhgGSbXMRpxDwLpFxm8xmTbtd

ret = client.Write(fd866, []byte("As8R"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) As8RBUTVmSLW6MJWvcmVbPT4sWhRVdKw7DWUJa2jJCNMrXQ6pLO7uQSsmCiMpNMH8JwgyhgGSbXMRpxDwLpFxm8xmTbtd

ret = client.Close(fd858)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd834)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) sKiTJp4lmRe8L0zpkMZ9FrAfDNSkHhq5bEfmBlH_D0nSdWCKkr4sQmhSDQzkGi

ret = client.Write(fd862, []byte("reW3brQ5WtJ7h2U4TpG12WqkOd_p9q6FUy023DBOQcBgKIGIGoMxFLcD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) reW3brQ5WtJ7h2U4TpG12WqkOd_p9q6FUy023DBOQcBgKIGIGoMxFLcDDQzkGi

fd867 := client.Open("/HjYfMWqq3p.txt", client.O_RDWR|client.O_CREATE)
if(fd867 < 0) {
  panic("open failed")
}

//fd state: (41) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVFavfRWPp

ret = client.Write(fd827, []byte("sKMqDZIiTpMAnBpw32LOHqMKw5YcXmLra2siBHAupNg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVsKMqDZIiTpMAnBpw32LOHqMKw5YcXmLra2siBHAupNg

fd868 := client.Open("/E34Lq2h34a.txt", client.O_RDWR|client.O_CREATE)
if(fd868 < 0) {
  panic("open failed")
}


ret = client.Seek(fd838, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}


ret = client.Close(fd838)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd840, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


ret = client.Seek(fd857, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd869 := client.Open("/F_DiFXQHjM.txt", client.O_RDWR|client.O_CREATE)
if(fd869 < 0) {
  panic("open failed")
}


ret = client.Close(fd851)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd853, 20)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd870 := client.Open("/elTreHWyzI.txt", client.O_RDWR|client.O_CREATE)
if(fd870 < 0) {
  panic("open failed")
}


fd871 := client.Open("/uADXMtzZGV.txt", client.O_RDWR|client.O_CREATE)
if(fd871 < 0) {
  panic("open failed")
}

//fd state: (150) 65ocKPNQhErm3i9_D7cWaefooTa2yn72d8tkQDGn6M3lyKiDreQhvIpahKMV7tFa1USvELrSqmFO97N8KJ1MrpYHuKFdphPCfMy2AjOzmpI9NSBrWRqQVh3slkjCjuny_DCHS44H9JSJX6dRP62f6mtV8e_q1vxukpnaZr5UkIzWe4h9BOTlDZGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY

ret = client.Write(fd797, []byte("RYayh2H2qtIDX_oN6u3poP71h6IEViKA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (182) 65ocKPNQhErm3i9_D7cWaefooTa2yn72d8tkQDGn6M3lyKiDreQhvIpahKMV7tFa1USvELrSqmFO97N8KJ1MrpYHuKFdphPCfMy2AjOzmpI9NSBrWRqQVh3slkjCjuny_DCHS44H9JSJX6dRP62f6mRYayh2H2qtIDX_oN6u3poP71h6IEViKAGRtlVcxk6ZdAfRHvSUIP2xb7l8NfeH0G5XUYoz7hIRGf5fRmaIOUETZiJbjNSngY
//fd state: (0) 

ret = client.Write(fd857, []byte("8srmQGmytQHtj0HMxZL4rzxHyZBn6i36f6IOp9YJIabPhT0Rat5SOUyv2jgtFg_TQ8mXTmbbnsfAWeC6ZFp87kau3DAe7UR0_wq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (99) 8srmQGmytQHtj0HMxZL4rzxHyZBn6i36f6IOp9YJIabPhT0Rat5SOUyv2jgtFg_TQ8mXTmbbnsfAWeC6ZFp87kau3DAe7UR0_wq

buf, ret = client.Read(fd809, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd842, 43)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd872 := client.Open("/9jmuO2RAmg.txt", client.O_RDWR|client.O_CREATE)
if(fd872 < 0) {
  panic("open failed")
}


ret = client.Close(fd856)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd871, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd867, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}


ret = client.Seek(fd844, 180, client.SEEK_SET)
if(ret != 180) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 180)
  panic("seek failed")
}


ret = client.Close(fd845)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd868)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd850, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


fd873 := client.Open("/FvsiKZeF_q.txt", client.O_RDWR|client.O_CREATE)
if(fd873 < 0) {
  panic("open failed")
}

//fd state: (99) 8srmQGmytQHtj0HMxZL4rzxHyZBn6i36f6IOp9YJIabPhT0Rat5SOUyv2jgtFg_TQ8mXTmbbnsfAWeC6ZFp87kau3DAe7UR0_wq

ret = client.Write(fd857, []byte("Esfc17vecwuvR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (112) 8srmQGmytQHtj0HMxZL4rzxHyZBn6i36f6IOp9YJIabPhT0Rat5SOUyv2jgtFg_TQ8mXTmbbnsfAWeC6ZFp87kau3DAe7UR0_wqEsfc17vecwuvR

ret = client.Close(fd844)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd839, []byte("Hmi0zD0FeScNuK8JxztGf86sFNckZR5R_ElFBQ_xDrnQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) Hmi0zD0FeScNuK8JxztGf86sFNckZR5R_ElFBQ_xDrnQ
//fd state: (0) 

ret = client.Write(fd853, []byte("jrHbUiF4bKS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) jrHbUiF4bKS

ret = client.Close(fd797)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd842)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd855)
if(ret != 0) {
  panic("close failed")
}

//fd state: (51) hL7_gyQ9_zQZdQQOW08Cm4l2Jql3zopmHrZc2Cv21Qd22pyNkTOtJAd1ic

ret = client.Write(fd840, []byte("rCRwiunof2DWWK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) hL7_gyQ9_zQZdQQOW08Cm4l2Jql3zopmHrZc2Cv21Qd22pyNkTOrCRwiunof2DWWK

ret = client.Seek(fd873, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd836)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd831, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "TvRq8iKv") {
  panic("wrong data returned")
}


fd874 := client.Open("/QjQHErqwDF.txt", client.O_RDWR|client.O_CREATE)
if(fd874 < 0) {
  panic("open failed")
}

//fd state: (11) jrHbUiF4bKS

ret = client.Write(fd853, []byte("nTSWVC2hlpHB_BDEP_bzvRGciDGo7CONtM_Swb61BJx4srYHrEh9BxzhQUvT3JHUGQ94PASYck7I2EzXo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) jrHbUiF4bKSnTSWVC2hlpHB_BDEP_bzvRGciDGo7CONtM_Swb61BJx4srYHrEh9BxzhQUvT3JHUGQ94PASYck7I2EzXo

buf, ret = client.Read(fd863, 13)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd872, 8)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd875 := client.Open("/gzkpsBF3YU.txt", client.O_RDWR|client.O_CREATE)
if(fd875 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd860, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd876 := client.Open("/P4ToA7VMb5.txt", client.O_RDWR|client.O_CREATE)
if(fd876 < 0) {
  panic("open failed")
}


ret = client.Seek(fd862, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}


buf, ret = client.Read(fd865, 49)
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


ret = client.Close(fd872)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd805, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd871, []byte("5xTa9f8999RYvy4oc3kHA_cDAfyxxnjiovl2eSW8Dilo9vMn0SlbAIoPK8DVfanMGZLPdPjWP6kpdTxZJ0us"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) 5xTa9f8999RYvy4oc3kHA_cDAfyxxnjiovl2eSW8Dilo9vMn0SlbAIoPK8DVfanMGZLPdPjWP6kpdTxZJ0us

ret = client.Close(fd874)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd839, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd877 := client.Open("/42C0UQibkO.txt", client.O_RDWR|client.O_CREATE)
if(fd877 < 0) {
  panic("open failed")
}


ret = client.Close(fd876)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd861, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd865, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (19) reW3brQ5WtJ7h2U4TpG12WqkOd_p9q6FUy023DBOQcBgKIGIGoMxFLcDDQzkGi

ret = client.Write(fd862, []byte("Izrsjj7Tzz4JPfo4aLEkJnUevEUvytI1S_5rxMs7WMUiXpDQz1Sl17CrBWo3t6SqC6u8U"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (88) reW3brQ5WtJ7h2U4TpGIzrsjj7Tzz4JPfo4aLEkJnUevEUvytI1S_5rxMs7WMUiXpDQz1Sl17CrBWo3t6SqC6u8U
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
buf, ret = client.Read(fd857, 29)
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


fd878 := client.Open("/fDoQce4Ark.txt", client.O_RDWR|client.O_CREATE)
if(fd878 < 0) {
  panic("open failed")
}


fd879 := client.Open("/12b1YLYef1.txt", client.O_RDWR|client.O_CREATE)
if(fd879 < 0) {
  panic("open failed")
}


ret = client.Close(fd805)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd869, []byte("evqXE9vigV50"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (12) evqXE9vigV50

buf, ret = client.Read(fd863, 24)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd877, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd875)
if(ret != 0) {
  panic("close failed")
}


fd880 := client.Open("/ll75qwO0GR.txt", client.O_RDWR|client.O_CREATE)
if(fd880 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd863, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd879, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "k1Tr49rb9cvxxnYnS1eW_mV1FczxmrwebOR0FMzfF") {
  panic("wrong data returned")
}


ret = client.Close(fd860)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd867)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd866, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "BUTVmSLW6MJWvcmVbPT4sWhRVdKw7DWUJa2jJCNMrXQ6pLO7uQSsmCiMpNMH8JwgyhgGSbXMRpxDwLpFxm8xmTbt") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd850, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "J3qOlu4DCoOwvDJR1t6gx1gXHOlgVhZq7ncm7teCWWBqAI3xiP2KC8t") {
  panic("wrong data returned")
}


ret = client.Seek(fd878, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd881 := client.Open("/EoL2bkUCkD.txt", client.O_RDWR|client.O_CREATE)
if(fd881 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd869, 97)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd840)
if(ret != 0) {
  panic("close failed")
}


fd882 := client.Open("/_joua8qqGG.txt", client.O_RDWR|client.O_CREATE)
if(fd882 < 0) {
  panic("open failed")
}


fd883 := client.Open("/FZZ8y6anbP.txt", client.O_RDWR|client.O_CREATE)
if(fd883 < 0) {
  panic("open failed")
}


ret = client.Seek(fd831, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Close(fd809)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd850)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd879)
if(ret != 0) {
  panic("close failed")
}


fd884 := client.Open("/KIwtxrfoHo.txt", client.O_RDWR|client.O_CREATE)
if(fd884 < 0) {
  panic("open failed")
}


fd885 := client.Open("/2qm1UAFxsb.txt", client.O_RDWR|client.O_CREATE)
if(fd885 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd880, 28)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd885)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd878, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (112) 8srmQGmytQHtj0HMxZL4rzxHyZBn6i36f6IOp9YJIabPhT0Rat5SOUyv2jgtFg_TQ8mXTmbbnsfAWeC6ZFp87kau3DAe7UR0_wqEsfc17vecwuvR

ret = client.Write(fd857, []byte("moeHwWZm4jDdQZVyrUJs8YVq59xHqAqhls3OEWPO2iBL05e_W9gm9VOIVhPqs8o_kE1ykvyVpb8HUGdQBpdH9nan"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (200) 8srmQGmytQHtj0HMxZL4rzxHyZBn6i36f6IOp9YJIabPhT0Rat5SOUyv2jgtFg_TQ8mXTmbbnsfAWeC6ZFp87kau3DAe7UR0_wqEsfc17vecwuvRmoeHwWZm4jDdQZVyrUJs8YVq59xHqAqhls3OEWPO2iBL05e_W9gm9VOIVhPqs8o_kE1ykvyVpb8HUGdQBpdH9nan

fd886 := client.Open("/vg_zIyook5.txt", client.O_RDWR|client.O_CREATE)
if(fd886 < 0) {
  panic("open failed")
}


fd887 := client.Open("/b09IPCVYla.txt", client.O_RDWR|client.O_CREATE)
if(fd887 < 0) {
  panic("open failed")
}


ret = client.Seek(fd882, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd888 := client.Open("/VdQmE9F0jn.txt", client.O_RDWR|client.O_CREATE)
if(fd888 < 0) {
  panic("open failed")
}


fd889 := client.Open("/kfZ_4vRm9s.txt", client.O_RDWR|client.O_CREATE)
if(fd889 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd880, []byte("tQUWGC1GeaA8gGOZs9kHD6groTP78FB8pKa38HAECZR0F"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) tQUWGC1GeaA8gGOZs9kHD6groTP78FB8pKa38HAECZR0F

fd890 := client.Open("/3889mJmyIM.txt", client.O_RDWR|client.O_CREATE)
if(fd890 < 0) {
  panic("open failed")
}


fd891 := client.Open("/7lLjkB_6Sg.txt", client.O_RDWR|client.O_CREATE)
if(fd891 < 0) {
  panic("open failed")
}


ret = client.Seek(fd878, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd887, 84)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lq1Q7lZxnhaOVoPtgFVfwePtyfb5VHzpXZBRdDTa0jp4i1xjl24KWLDLloaOerXs_s5zyeq_BmLkWnOhn4mq") {
  panic("wrong data returned")
}


ret = client.Seek(fd866, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


ret = client.Close(fd884)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd889)
if(ret != 0) {
  panic("close failed")
}

//fd state: (92) jrHbUiF4bKSnTSWVC2hlpHB_BDEP_bzvRGciDGo7CONtM_Swb61BJx4srYHrEh9BxzhQUvT3JHUGQ94PASYck7I2EzXo

ret = client.Write(fd853, []byte("p7EKIcS6_NB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) jrHbUiF4bKSnTSWVC2hlpHB_BDEP_bzvRGciDGo7CONtM_Swb61BJx4srYHrEh9BxzhQUvT3JHUGQ94PASYck7I2EzXop7EKIcS6_NB
//fd state: (88) reW3brQ5WtJ7h2U4TpGIzrsjj7Tzz4JPfo4aLEkJnUevEUvytI1S_5rxMs7WMUiXpDQz1Sl17CrBWo3t6SqC6u8U

ret = client.Write(fd862, []byte("KUWSp5Edy3qrKOBeM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (105) reW3brQ5WtJ7h2U4TpGIzrsjj7Tzz4JPfo4aLEkJnUevEUvytI1S_5rxMs7WMUiXpDQz1Sl17CrBWo3t6SqC6u8UKUWSp5Edy3qrKOBeM

ret = client.Close(fd890)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd866, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "2jJCNMrXQ6pLO7uQSsmCiMpNMH8JwgyhgGSbXMRpxDwLpFxm8") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd877, []byte("Be38cBfU34B0d87QFTFPZXaKyP2VeFyDKXIAjhglh2QRxXx0qXe9W"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) Be38cBfU34B0d87QFTFPZXaKyP2VeFyDKXIAjhglh2QRxXx0qXe9W
//fd state: (12) evqXE9vigV50

ret = client.Write(fd869, []byte("ZTOPqfk1IJsVRVxeNcydSdwWipgYTDGHCwHWmz8vXeyQLnd9QpClkK2Mp0fvPmAAZhzhanFRKs3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (87) evqXE9vigV50ZTOPqfk1IJsVRVxeNcydSdwWipgYTDGHCwHWmz8vXeyQLnd9QpClkK2Mp0fvPmAAZhzhanFRKs3

buf, ret = client.Read(fd871, 62)
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


ret = client.Seek(fd886, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (84) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVsKMqDZIiTpMAnBpw32LOHqMKw5YcXmLra2siBHAupNg

ret = client.Write(fd827, []byte("qLbCDlkxBHg6zD3C07J1ldKX6RnTzo6DUhWKkY1xZNEduHbx3gZmfmGgI5R5bXCAvWeOAvWMtdoySZfkzi6PPQYILoXrM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (177) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVsKMqDZIiTpMAnBpw32LOHqMKw5YcXmLra2siBHAupNgqLbCDlkxBHg6zD3C07J1ldKX6RnTzo6DUhWKkY1xZNEduHbx3gZmfmGgI5R5bXCAvWeOAvWMtdoySZfkzi6PPQYILoXrM

ret = client.Seek(fd871, 56, client.SEEK_SET)
if(ret != 56) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 56)
  panic("seek failed")
}


ret = client.Seek(fd886, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd839)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd880, 44, client.SEEK_SET)
if(ret != 44) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 44)
  panic("seek failed")
}


ret = client.Seek(fd857, 122, client.SEEK_SET)
if(ret != 122) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 122)
  panic("seek failed")
}


ret = client.Close(fd883)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd878)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd871, 34, client.SEEK_SET)
if(ret != 34) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 34)
  panic("seek failed")
}


fd892 := client.Open("/Wn_LlxApk0.txt", client.O_RDWR|client.O_CREATE)
if(fd892 < 0) {
  panic("open failed")
}


ret = client.Seek(fd881, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd893 := client.Open("/WK5l2pN7oo.txt", client.O_RDWR|client.O_CREATE)
if(fd893 < 0) {
  panic("open failed")
}


ret = client.Close(fd887)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd857, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}


ret = client.Seek(fd886, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd893)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd870, []byte("mJmGLspdAZeidGxS1UA6M1JZScUp7ZlO3_NOELMCsrOcVjO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) mJmGLspdAZeidGxS1UA6M1JZScUp7ZlO3_NOELMCsrOcVjO

ret = client.Close(fd880)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd871, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "l2eSW8Dilo9vMn0SlbAIoPK8DVfanMGZLPdPjWP6kpdTxZJ0us") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd882, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd892, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "YOzbUp2AgNu8iOoCyqgl4j13e9I9hcdw6zQPlQYPRBKkjyxL1aaw1hH9r55v0oUE9SwMQgyIzz55Eh_ylnJr8ErgwQn") {
  panic("wrong data returned")
}


ret = client.Seek(fd870, 38, client.SEEK_SET)
if(ret != 38) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 38)
  panic("seek failed")
}


buf, ret = client.Read(fd882, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd888, []byte("XAq1p9zxhdnmePpYubPepWpl1TD0SLE_2RGTrJlJUte9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (44) XAq1p9zxhdnmePpYubPepWpl1TD0SLE_2RGTrJlJUte9
//fd state: (0) 

ret = client.Write(fd863, []byte("kvSI9jkdaytubpmLiMWakHbDoD5E58Oy3kvwoINgnut3nYo44"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) kvSI9jkdaytubpmLiMWakHbDoD5E58Oy3kvwoINgnut3nYo44

fd894 := client.Open("/N_tIxiNaQn.txt", client.O_RDWR|client.O_CREATE)
if(fd894 < 0) {
  panic("open failed")
}


ret = client.Close(fd866)
if(ret != 0) {
  panic("close failed")
}


fd895 := client.Open("/4miTNfxwtx.txt", client.O_RDWR|client.O_CREATE)
if(fd895 < 0) {
  panic("open failed")
}


fd896 := client.Open("/TuxJ_iXijk.txt", client.O_RDWR|client.O_CREATE)
if(fd896 < 0) {
  panic("open failed")
}


ret = client.Seek(fd862, 60, client.SEEK_SET)
if(ret != 60) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 60)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd886, []byte("8pUcclXsFSCiMnmnijwNN"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (21) 8pUcclXsFSCiMnmnijwNN

fd897 := client.Open("/wSeTMDhPF3.txt", client.O_RDWR|client.O_CREATE)
if(fd897 < 0) {
  panic("open failed")
}


fd898 := client.Open("/nUgmRLHSa0.txt", client.O_RDWR|client.O_CREATE)
if(fd898 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd897, []byte("GyzF5SQrqxR9pRM6eqY0exITf1ZKos1HLymvbSiHcQL_nfDHv_wfkyeKqQUx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) GyzF5SQrqxR9pRM6eqY0exITf1ZKos1HLymvbSiHcQL_nfDHv_wfkyeKqQUx

ret = client.Close(fd862)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd827, 105, client.SEEK_SET)
if(ret != 105) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 105)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd881, []byte("B2QKsmrPdgWUcsyA0qUuMz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (22) B2QKsmrPdgWUcsyA0qUuMz

ret = client.Close(fd897)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd831)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd870, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "MCsrOcVjO") {
  panic("wrong data returned")
}


fd899 := client.Open("/58hhHlmzPu.txt", client.O_RDWR|client.O_CREATE)
if(fd899 < 0) {
  panic("open failed")
}


fd900 := client.Open("/c70JhCNiXm.txt", client.O_RDWR|client.O_CREATE)
if(fd900 < 0) {
  panic("open failed")
}


ret = client.Close(fd894)
if(ret != 0) {
  panic("close failed")
}


fd901 := client.Open("/HmduvtDGQA.txt", client.O_RDWR|client.O_CREATE)
if(fd901 < 0) {
  panic("open failed")
}


fd902 := client.Open("/5OvPSr9_7r.txt", client.O_RDWR|client.O_CREATE)
if(fd902 < 0) {
  panic("open failed")
}


ret = client.Close(fd892)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (87) evqXE9vigV50ZTOPqfk1IJsVRVxeNcydSdwWipgYTDGHCwHWmz8vXeyQLnd9QpClkK2Mp0fvPmAAZhzhanFRKs3

ret = client.Write(fd869, []byte("zdpQqr4dL7gt7qUxfIOjqUbCMlMM8VpUu1JhR1xCHwneio6fv4qb7iXSOZo1S76AW_GKRuPr3qff_FXK23b2cULYIh1_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (179) evqXE9vigV50ZTOPqfk1IJsVRVxeNcydSdwWipgYTDGHCwHWmz8vXeyQLnd9QpClkK2Mp0fvPmAAZhzhanFRKs3zdpQqr4dL7gt7qUxfIOjqUbCMlMM8VpUu1JhR1xCHwneio6fv4qb7iXSOZo1S76AW_GKRuPr3qff_FXK23b2cULYIh1_

fd903 := client.Open("/p821DSywl7.txt", client.O_RDWR|client.O_CREATE)
if(fd903 < 0) {
  panic("open failed")
}


ret = client.Close(fd886)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd903, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd904 := client.Open("/7BiGgDUEJp.txt", client.O_RDWR|client.O_CREATE)
if(fd904 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd904, 54)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd903, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd905 := client.Open("/Ue36Yd0iZk.txt", client.O_RDWR|client.O_CREATE)
if(fd905 < 0) {
  panic("open failed")
}


ret = client.Seek(fd891, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd902, []byte("lxlQgnvZQcRGE_xFYAz3e5Jvv5XYk1VEqlHdiJChAJTZcTj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) lxlQgnvZQcRGE_xFYAz3e5Jvv5XYk1VEqlHdiJChAJTZcTj
//fd state: (0) DL8hwxw3j2vtubdUOg3a_kIxqDomJK0RbFq4hrL_KK064FdB

ret = client.Write(fd905, []byte("rbIOtGGVpXEcP4s0ke75g_JxOOaHqY5D7rSqSiu7R0eCzuovn_cqktnRaa7os7qzKj"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) rbIOtGGVpXEcP4s0ke75g_JxOOaHqY5D7rSqSiu7R0eCzuovn_cqktnRaa7os7qzKj

ret = client.Close(fd853)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd863, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd898)
if(ret != 0) {
  panic("close failed")
}

//fd state: (105) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVsKMqDZIiTpMAnBpw32LOHqMKw5YcXmLra2siBHAupNgqLbCDlkxBHg6zD3C07J1ldKX6RnTzo6DUhWKkY1xZNEduHbx3gZmfmGgI5R5bXCAvWeOAvWMtdoySZfkzi6PPQYILoXrM

ret = client.Write(fd827, []byte("Sr8mUhUtqlvEugXc6zj5ledwi_2c5bwMp9h0IKaEYqQoRG8JjFvZRu12RuflF3Dzg8V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (172) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVsKMqDZIiTpMAnBpw32LOHqMKw5YcXmLra2siBHAupNgqLbCDlkxBHg6zD3C07J1lSr8mUhUtqlvEugXc6zj5ledwi_2c5bwMp9h0IKaEYqQoRG8JjFvZRu12RuflF3Dzg8VLoXrM

buf, ret = client.Read(fd891, 61)
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


ret = client.Close(fd903)
if(ret != 0) {
  panic("close failed")
}


fd906 := client.Open("/MF_bH4w72p.txt", client.O_RDWR|client.O_CREATE)
if(fd906 < 0) {
  panic("open failed")
}


ret = client.Close(fd901)
if(ret != 0) {
  panic("close failed")
}


fd907 := client.Open("/i1lq0NaG_I.txt", client.O_RDWR|client.O_CREATE)
if(fd907 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd904, []byte("Kn4BJR6_os8yXcJmWBqEvlCoY3wQmg6d7Ztsqj_XQphpHC2RK9s13hJRGCrYGfNf97qZqu0xyaFFiHnsNtijxM7an"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) Kn4BJR6_os8yXcJmWBqEvlCoY3wQmg6d7Ztsqj_XQphpHC2RK9s13hJRGCrYGfNf97qZqu0xyaFFiHnsNtijxM7an

fd908 := client.Open("/G2EwloCA2u.txt", client.O_RDWR|client.O_CREATE)
if(fd908 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd863, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (47) mJmGLspdAZeidGxS1UA6M1JZScUp7ZlO3_NOELMCsrOcVjO

ret = client.Write(fd870, []byte("vpKuL33HmbdnKQhbaVx4FK28xq9popf048gGNMySh2NfS7thr7LjbIlP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (103) mJmGLspdAZeidGxS1UA6M1JZScUp7ZlO3_NOELMCsrOcVjOvpKuL33HmbdnKQhbaVx4FK28xq9popf048gGNMySh2NfS7thr7LjbIlP
//fd state: (49) kvSI9jkdaytubpmLiMWakHbDoD5E58Oy3kvwoINgnut3nYo44

ret = client.Write(fd863, []byte("18VpdIx8GiiMnJOfKaZWNGSUmZXIMrbXiqzImqKAPDuvmlHC2Gcqjyi9aji3Mly2OUOPEV0UUw5tf1qpaJDAQhf4R_HkzrMxyay"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (148) kvSI9jkdaytubpmLiMWakHbDoD5E58Oy3kvwoINgnut3nYo4418VpdIx8GiiMnJOfKaZWNGSUmZXIMrbXiqzImqKAPDuvmlHC2Gcqjyi9aji3Mly2OUOPEV0UUw5tf1qpaJDAQhf4R_HkzrMxyay

buf, ret = client.Read(fd900, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (172) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVsKMqDZIiTpMAnBpw32LOHqMKw5YcXmLra2siBHAupNgqLbCDlkxBHg6zD3C07J1lSr8mUhUtqlvEugXc6zj5ledwi_2c5bwMp9h0IKaEYqQoRG8JjFvZRu12RuflF3Dzg8VLoXrM

ret = client.Write(fd827, []byte("Sl6hZ2s2Pqco5NcnGqfllC0ZP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (197) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVsKMqDZIiTpMAnBpw32LOHqMKw5YcXmLra2siBHAupNgqLbCDlkxBHg6zD3C07J1lSr8mUhUtqlvEugXc6zj5ledwi_2c5bwMp9h0IKaEYqQoRG8JjFvZRu12RuflF3Dzg8VSl6hZ2s2Pqco5NcnGqfllC0ZP

fd909 := client.Open("/MPMXhpOG4e.txt", client.O_RDWR|client.O_CREATE)
if(fd909 < 0) {
  panic("open failed")
}

//fd state: (0) ugYiWD15NtHBc4g3mvDWr2aXMOuGOyL_TY2VvN0n5ONkKvMMbP6XzO8vtKJtPS4N_jMjpbfqJTCDapCmX

ret = client.Write(fd908, []byte("5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcdu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (49) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduP6XzO8vtKJtPS4N_jMjpbfqJTCDapCmX

buf, ret = client.Read(fd899, 7)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd857, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}

//fd state: (89) Kn4BJR6_os8yXcJmWBqEvlCoY3wQmg6d7Ztsqj_XQphpHC2RK9s13hJRGCrYGfNf97qZqu0xyaFFiHnsNtijxM7an

ret = client.Write(fd904, []byte("w6ozLJXvRXr59OnBUVYRMia85zlemSR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) Kn4BJR6_os8yXcJmWBqEvlCoY3wQmg6d7Ztsqj_XQphpHC2RK9s13hJRGCrYGfNf97qZqu0xyaFFiHnsNtijxM7anw6ozLJXvRXr59OnBUVYRMia85zlemSR

ret = client.Close(fd904)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd870)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd882, []byte("9tQvMtJoyJ62lo"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (14) 9tQvMtJoyJ62lo

fd910 := client.Open("/RQKmSwTj6V.txt", client.O_RDWR|client.O_CREATE)
if(fd910 < 0) {
  panic("open failed")
}


fd911 := client.Open("/ObQwqCflA9.txt", client.O_RDWR|client.O_CREATE)
if(fd911 < 0) {
  panic("open failed")
}

//fd state: (0) kvSI9jkdaytubpmLiMWakHbDoD5E58Oy3kvwoINgnut3nYo4418VpdIx8GiiMnJOfKaZWNGSUmZXIMrbXiqzImqKAPDuvmlHC2Gcqjyi9aji3Mly2OUOPEV0UUw5tf1qpaJDAQhf4R_HkzrMxyay

ret = client.Write(fd906, []byte("bQ4SElzbXUD7JxDzu4fS0zX006ng5KR8IO4ZwDOOOhYUtx2s_eUC0x0N34dxpQZY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) bQ4SElzbXUD7JxDzu4fS0zX006ng5KR8IO4ZwDOOOhYUtx2s_eUC0x0N34dxpQZYfKaZWNGSUmZXIMrbXiqzImqKAPDuvmlHC2Gcqjyi9aji3Mly2OUOPEV0UUw5tf1qpaJDAQhf4R_HkzrMxyay

ret = client.Seek(fd865, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd906)
if(ret != 0) {
  panic("close failed")
}


fd912 := client.Open("/iSviWTXnNc.txt", client.O_RDWR|client.O_CREATE)
if(fd912 < 0) {
  panic("open failed")
}

//fd state: (49) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduP6XzO8vtKJtPS4N_jMjpbfqJTCDapCmX

ret = client.Write(fd908, []byte("CefufyMh63jhejAMMs_WsmEXF5PQ0vOGMfutut0VLIuHRKk5TxwhJDQy8RnbgBb3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (113) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduCefufyMh63jhejAMMs_WsmEXF5PQ0vOGMfutut0VLIuHRKk5TxwhJDQy8RnbgBb3

ret = client.Close(fd902)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd910, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd900)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd911)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd863, 67, client.SEEK_SET)
if(ret != 67) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 67)
  panic("seek failed")
}


fd913 := client.Open("/zK0Bie8jKe.txt", client.O_RDWR|client.O_CREATE)
if(fd913 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd865, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd907, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd914 := client.Open("/FPPNgeMK1l.txt", client.O_RDWR|client.O_CREATE)
if(fd914 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd905, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (113) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduCefufyMh63jhejAMMs_WsmEXF5PQ0vOGMfutut0VLIuHRKk5TxwhJDQy8RnbgBb3

ret = client.Write(fd908, []byte("hDFonL9zPdCZB46ma7NO6MCoOgkI6VHAhmnB9uYlsWcigbMVtEtcNkWOr6lTqbm3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (177) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduCefufyMh63jhejAMMs_WsmEXF5PQ0vOGMfutut0VLIuHRKk5TxwhJDQy8RnbgBb3hDFonL9zPdCZB46ma7NO6MCoOgkI6VHAhmnB9uYlsWcigbMVtEtcNkWOr6lTqbm3

buf, ret = client.Read(fd913, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd915 := client.Open("/jQdl4Skatz.txt", client.O_RDWR|client.O_CREATE)
if(fd915 < 0) {
  panic("open failed")
}


ret = client.Seek(fd915, 43, client.SEEK_SET)
if(ret != 43) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 43)
  panic("seek failed")
}


fd916 := client.Open("/syBqqWYO2I.txt", client.O_RDWR|client.O_CREATE)
if(fd916 < 0) {
  panic("open failed")
}


fd917 := client.Open("/EjZosYJW8N.txt", client.O_RDWR|client.O_CREATE)
if(fd917 < 0) {
  panic("open failed")
}


ret = client.Seek(fd877, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}

//fd state: (0) v5AY3uvetJDvGFNFq4xx7AyPv0VDCzwLr7LHX

ret = client.Write(fd909, []byte("XHuQNL0HtI0Se5uA7V61T0_XyME6a4OaPqMWGtLscR4kElJReoHn8B2H8"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (57) XHuQNL0HtI0Se5uA7V61T0_XyME6a4OaPqMWGtLscR4kElJReoHn8B2H8

ret = client.Seek(fd891, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd910, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd913, []byte("Ww8zLyrDVSm53DdatvaOciEv7TYJXE22imlO5ZI7hrT9PmWA9zRi3v9Fu3LmAvihh7UC9S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) Ww8zLyrDVSm53DdatvaOciEv7TYJXE22imlO5ZI7hrT9PmWA9zRi3v9Fu3LmAvihh7UC9S

ret = client.Seek(fd871, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}

//fd state: (177) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduCefufyMh63jhejAMMs_WsmEXF5PQ0vOGMfutut0VLIuHRKk5TxwhJDQy8RnbgBb3hDFonL9zPdCZB46ma7NO6MCoOgkI6VHAhmnB9uYlsWcigbMVtEtcNkWOr6lTqbm3

ret = client.Write(fd908, []byte("R91df5eMpxVbYiWkCUtKwSDrvR_VMLeTmWb4wzlZIKIx"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (221) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduCefufyMh63jhejAMMs_WsmEXF5PQ0vOGMfutut0VLIuHRKk5TxwhJDQy8RnbgBb3hDFonL9zPdCZB46ma7NO6MCoOgkI6VHAhmnB9uYlsWcigbMVtEtcNkWOr6lTqbm3R91df5eMpxVbYiWkCUtKwSDrvR_VMLeTmWb4wzlZIKIx

ret = client.Seek(fd896, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd917, []byte("S9tn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (4) S9tn

ret = client.Close(fd881)
if(ret != 0) {
  panic("close failed")
}


fd918 := client.Open("/uS4PfPrncQ.txt", client.O_RDWR|client.O_CREATE)
if(fd918 < 0) {
  panic("open failed")
}


ret = client.Seek(fd865, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd919 := client.Open("/5bNqIKtQYU.txt", client.O_RDWR|client.O_CREATE)
if(fd919 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd908, 29)
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


buf, ret = client.Read(fd908, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd920 := client.Open("/gnQflFAoey.txt", client.O_RDWR|client.O_CREATE)
if(fd920 < 0) {
  panic("open failed")
}


fd921 := client.Open("/a1Vqa8McUR.txt", client.O_RDWR|client.O_CREATE)
if(fd921 < 0) {
  panic("open failed")
}


ret = client.Close(fd912)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd917, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd891, []byte("535v2k5K9rLUKWEqnonh48PfQfJ8LZnK_b1tcYRb2o4uM_gskZ2HhkZE8c5JO5IWXTn3w99sKL24fO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) 535v2k5K9rLUKWEqnonh48PfQfJ8LZnK_b1tcYRb2o4uM_gskZ2HhkZE8c5JO5IWXTn3w99sKL24fO

fd922 := client.Open("/h9gQ2G1psB.txt", client.O_RDWR|client.O_CREATE)
if(fd922 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd913, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd905)
if(ret != 0) {
  panic("close failed")
}


fd923 := client.Open("/LJuG61n_0Q.txt", client.O_RDWR|client.O_CREATE)
if(fd923 < 0) {
  panic("open failed")
}


fd924 := client.Open("/5rI6XWi8iP.txt", client.O_RDWR|client.O_CREATE)
if(fd924 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd882, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) JnAWv15UZZUFYUDj6pJ_a5zcWfBbsHuvUlL5HDUf4Urh4VEEHi7CcCXIlA5NZXy9rBEMVXfquSldit_uzMS_o0vSllsY7fi8IFEDy4LzDtxMrv9JPNiWy20wI_QT

ret = client.Write(fd914, []byte("GWSuJLfYYQiWQ3DIKP77hyyJXjLNDqiqWzN1W5YT7OAztpNqGV2PoR49"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (56) GWSuJLfYYQiWQ3DIKP77hyyJXjLNDqiqWzN1W5YT7OAztpNqGV2PoR49lA5NZXy9rBEMVXfquSldit_uzMS_o0vSllsY7fi8IFEDy4LzDtxMrv9JPNiWy20wI_QT

buf, ret = client.Read(fd907, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd908, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd888, 39)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd925 := client.Open("/z8TQLK8GRZ.txt", client.O_RDWR|client.O_CREATE)
if(fd925 < 0) {
  panic("open failed")
}


ret = client.Close(fd914)
if(ret != 0) {
  panic("close failed")
}


fd926 := client.Open("/zgdT4mG5bm.txt", client.O_RDWR|client.O_CREATE)
if(fd926 < 0) {
  panic("open failed")
}


fd927 := client.Open("/8XfedTKzMs.txt", client.O_RDWR|client.O_CREATE)
if(fd927 < 0) {
  panic("open failed")
}

//fd state: (57) XHuQNL0HtI0Se5uA7V61T0_XyME6a4OaPqMWGtLscR4kElJReoHn8B2H8

ret = client.Write(fd909, []byte("EYUzeSnSzLZLxrRcGvyrkoEElap46omt2USSEGsfU5Wzwc4XmW95Ma_ue_Gxd6XytIrIFu58p8Mj2VuvQ9x7TQO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) XHuQNL0HtI0Se5uA7V61T0_XyME6a4OaPqMWGtLscR4kElJReoHn8B2H8EYUzeSnSzLZLxrRcGvyrkoEElap46omt2USSEGsfU5Wzwc4XmW95Ma_ue_Gxd6XytIrIFu58p8Mj2VuvQ9x7TQO

fd928 := client.Open("/XzHLmOJsV2.txt", client.O_RDWR|client.O_CREATE)
if(fd928 < 0) {
  panic("open failed")
}


ret = client.Seek(fd920, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd928, []byte("ck78CUniA51uAXRLevdke4LSLW6TvkUvulQmSiqJzd1P53ct7o_k3pgJJy04vQRnh7DQvWFwWPjV8dgcoKd"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) ck78CUniA51uAXRLevdke4LSLW6TvkUvulQmSiqJzd1P53ct7o_k3pgJJy04vQRnh7DQvWFwWPjV8dgcoKd

ret = client.Close(fd909)
if(ret != 0) {
  panic("close failed")
}


fd929 := client.Open("/eLMT0M0BqC.txt", client.O_RDWR|client.O_CREATE)
if(fd929 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now//fd state: (67) bQ4SElzbXUD7JxDzu4fS0zX006ng5KR8IO4ZwDOOOhYUtx2s_eUC0x0N34dxpQZYfKaZWNGSUmZXIMrbXiqzImqKAPDuvmlHC2Gcqjyi9aji3Mly2OUOPEV0UUw5tf1qpaJDAQhf4R_HkzrMxyay

ret = client.Write(fd863, []byte("wPQ8_L5CA4d8Wzo0uqZ2frTMn9mjuQWOIpKrcBWOWooB6x6PLmAFha679KQ9Slpel58Jx1MfeZ3CVg"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (145) bQ4SElzbXUD7JxDzu4fS0zX006ng5KR8IO4ZwDOOOhYUtx2s_eUC0x0N34dxpQZYfKawPQ8_L5CA4d8Wzo0uqZ2frTMn9mjuQWOIpKrcBWOWooB6x6PLmAFha679KQ9Slpel58Jx1MfeZ3CVgyay
//fd state: (20) 5xTa9f8999RYvy4oc3kHA_cDAfyxxnjiovl2eSW8Dilo9vMn0SlbAIoPK8DVfanMGZLPdPjWP6kpdTxZJ0us

ret = client.Write(fd871, []byte("0FLzIR_kayzGxlmRCy3GFJawxl1E6wQAVvMUd1C9iMJyXRjX0tiDWvvDelW_5ZfQqyoxfXElcz"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) 5xTa9f8999RYvy4oc3kH0FLzIR_kayzGxlmRCy3GFJawxl1E6wQAVvMUd1C9iMJyXRjX0tiDWvvDelW_5ZfQqyoxfXElcz

ret = client.Close(fd917)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd927)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd908, 217, client.SEEK_SET)
if(ret != 217) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 217)
  panic("seek failed")
}


buf, ret = client.Read(fd922, 88)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd919)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd865, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (78) 535v2k5K9rLUKWEqnonh48PfQfJ8LZnK_b1tcYRb2o4uM_gskZ2HhkZE8c5JO5IWXTn3w99sKL24fO

ret = client.Write(fd891, []byte("2fnbCvhLsMJZx4wosY6tU61ibGaxkLivSO_bGMDCKDdx22tiWZf2nqkRZv17tIVFgD"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) 535v2k5K9rLUKWEqnonh48PfQfJ8LZnK_b1tcYRb2o4uM_gskZ2HhkZE8c5JO5IWXTn3w99sKL24fO2fnbCvhLsMJZx4wosY6tU61ibGaxkLivSO_bGMDCKDdx22tiWZf2nqkRZv17tIVFgD
//fd state: (0) 

ret = client.Write(fd865, []byte("t87EI8vprHTuXlNCFMCmmn1GqOSqxpKbFLCoIsni50NcRd8AHTqRP6gRkspmbMctk3P1MR1u_bEHfeOuWyARV0BLH679"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) t87EI8vprHTuXlNCFMCmmn1GqOSqxpKbFLCoIsni50NcRd8AHTqRP6gRkspmbMctk3P1MR1u_bEHfeOuWyARV0BLH679

fd930 := client.Open("/0jtYObXPre.txt", client.O_RDWR|client.O_CREATE)
if(fd930 < 0) {
  panic("open failed")
}


fd931 := client.Open("/kwPisEKdpD.txt", client.O_RDWR|client.O_CREATE)
if(fd931 < 0) {
  panic("open failed")
}


ret = client.Seek(fd910, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd921, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd913)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) k4arljCVCpl8LM98NijSFDr9pIlgzIHWbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WPIs8AO2gQuC5LO0EF1jpG0he_W5QG8gElNCJ5O_kxVxhY59XAoIal2VuPAjYARbgubjiyTuJ7yb8qwIUU

ret = client.Write(fd929, []byte("w_icxKWZ9sVtHdec32IR1oQ96LH9PfM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (31) w_icxKWZ9sVtHdec32IR1oQ96LH9PfMWbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WPIs8AO2gQuC5LO0EF1jpG0he_W5QG8gElNCJ5O_kxVxhY59XAoIal2VuPAjYARbgubjiyTuJ7yb8qwIUU

buf, ret = client.Read(fd918, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WUORDdFhX7aGgUN2zIsira4i2V7MFgHqxsIZKaa5LmZ9WUrmu5RAD") {
  panic("wrong data returned")
}


fd932 := client.Open("/6HMIdOkz5H.txt", client.O_RDWR|client.O_CREATE)
if(fd932 < 0) {
  panic("open failed")
}


ret = client.Seek(fd908, 164, client.SEEK_SET)
if(ret != 164) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 164)
  panic("seek failed")
}


fd933 := client.Open("/oGnzWegge0.txt", client.O_RDWR|client.O_CREATE)
if(fd933 < 0) {
  panic("open failed")
}

//fd state: (179) evqXE9vigV50ZTOPqfk1IJsVRVxeNcydSdwWipgYTDGHCwHWmz8vXeyQLnd9QpClkK2Mp0fvPmAAZhzhanFRKs3zdpQqr4dL7gt7qUxfIOjqUbCMlMM8VpUu1JhR1xCHwneio6fv4qb7iXSOZo1S76AW_GKRuPr3qff_FXK23b2cULYIh1_

ret = client.Write(fd869, []byte("GLyBqBWC7H68e0hMBdx_zXFNUvr84RrONG70VSOAwLTWgmclNpdTn6ciQEcJ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (239) evqXE9vigV50ZTOPqfk1IJsVRVxeNcydSdwWipgYTDGHCwHWmz8vXeyQLnd9QpClkK2Mp0fvPmAAZhzhanFRKs3zdpQqr4dL7gt7qUxfIOjqUbCMlMM8VpUu1JhR1xCHwneio6fv4qb7iXSOZo1S76AW_GKRuPr3qff_FXK23b2cULYIh1_GLyBqBWC7H68e0hMBdx_zXFNUvr84RrONG70VSOAwLTWgmclNpdTn6ciQEcJ

ret = client.Close(fd869)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd907, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd933, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (53) WUORDdFhX7aGgUN2zIsira4i2V7MFgHqxsIZKaa5LmZ9WUrmu5RADZN9jb5urL2DGG6A3_FWD9JNRkXsN4FZryw2C3Uy7YBdoZS_ySQfT0cTkTzoCHUmYw23X

ret = client.Write(fd918, []byte("IAdvvpM_kK_xkA9PPvTYkjDm8sdqA6nbvQa2SfeTufV53vd810Zqp3Y6TMISNaAGb274RFCPT1KlRr35R4UzNmKIMFf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (144) WUORDdFhX7aGgUN2zIsira4i2V7MFgHqxsIZKaa5LmZ9WUrmu5RADIAdvvpM_kK_xkA9PPvTYkjDm8sdqA6nbvQa2SfeTufV53vd810Zqp3Y6TMISNaAGb274RFCPT1KlRr35R4UzNmKIMFf

ret = client.Close(fd933)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd932)
if(ret != 0) {
  panic("close failed")
}

//fd state: (94) 5xTa9f8999RYvy4oc3kH0FLzIR_kayzGxlmRCy3GFJawxl1E6wQAVvMUd1C9iMJyXRjX0tiDWvvDelW_5ZfQqyoxfXElcz

ret = client.Write(fd871, []byte("lpEyht3bkcwEI19GYw1yV8SKIjBs0H3bcYcj3qFt0Y5y8Dr2g6wk9KYUQSs5p0KVCgtC8hAcdWo3BSFTB1hBQpJScK_gQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (187) 5xTa9f8999RYvy4oc3kH0FLzIR_kayzGxlmRCy3GFJawxl1E6wQAVvMUd1C9iMJyXRjX0tiDWvvDelW_5ZfQqyoxfXElczlpEyht3bkcwEI19GYw1yV8SKIjBs0H3bcYcj3qFt0Y5y8Dr2g6wk9KYUQSs5p0KVCgtC8hAcdWo3BSFTB1hBQpJScK_gQ

buf, ret = client.Read(fd877, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "aKyP2VeFyDKXIAjhglh2QRxXx0qXe9W") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd920, []byte("PDfgYjKpVWH5q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) PDfgYjKpVWH5q

ret = client.Seek(fd922, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd934 := client.Open("/_KHs6d6YBx.txt", client.O_RDWR|client.O_CREATE)
if(fd934 < 0) {
  panic("open failed")
}


ret = client.Seek(fd871, 85, client.SEEK_SET)
if(ret != 85) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 85)
  panic("seek failed")
}


buf, ret = client.Read(fd896, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd907, []byte("w2wiHMIwkojuN_ViyXwEpJX80QYE"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (28) w2wiHMIwkojuN_ViyXwEpJX80QYE
//fd state: (0) 

ret = client.Write(fd923, []byte("z9UKp12xhCJfTkoM200QXyIRfTcZW4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (30) z9UKp12xhCJfTkoM200QXyIRfTcZW4

ret = client.Close(fd923)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd916, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd877)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd925, 68)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd935 := client.Open("/3QNUtjWCV0.txt", client.O_RDWR|client.O_CREATE)
if(fd935 < 0) {
  panic("open failed")
}


ret = client.Close(fd926)
if(ret != 0) {
  panic("close failed")
}

//fd state: (197) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVsKMqDZIiTpMAnBpw32LOHqMKw5YcXmLra2siBHAupNgqLbCDlkxBHg6zD3C07J1lSr8mUhUtqlvEugXc6zj5ledwi_2c5bwMp9h0IKaEYqQoRG8JjFvZRu12RuflF3Dzg8VSl6hZ2s2Pqco5NcnGqfllC0ZP

ret = client.Write(fd827, []byte("q3_1RFmVUpBFvkVclCi6tUzoeW7oyiSrM_ZXMWKgtwKzVVbLryJWjsyPoU1avMU_5dN5iTGhnKLnJNUNpHqqOpuNtD3z"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (289) PY0kvtzValqfqt_thy5TceXVYNBKM6kqQJwQYQmwVsKMqDZIiTpMAnBpw32LOHqMKw5YcXmLra2siBHAupNgqLbCDlkxBHg6zD3C07J1lSr8mUhUtqlvEugXc6zj5ledwi_2c5bwMp9h0IKaEYqQoRG8JjFvZRu12RuflF3Dzg8VSl6hZ2s2Pqco5NcnGqfllC0ZPq3_1RFmVUpBFvkVclCi6tUzoeW7oyiSrM_ZXMWKgtwKzVVbLryJWjsyPoU1avMU_5dN5iTGhnKLnJNUNpHqqOpuNtD3z

buf, ret = client.Read(fd929, 48)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "WbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WP") {
  panic("wrong data returned")
}


ret = client.Close(fd928)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd899)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd896, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd910, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd891)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd925, []byte("UHWxEtP_gpj5bUH7vO1IjnSNvtEAgJZpfo7qGQydjrWqcEeRkD14FzjhwU7bmhaR1XFY51NfbCm7ZzTMz7z9y3qFjrgb3i"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (94) UHWxEtP_gpj5bUH7vO1IjnSNvtEAgJZpfo7qGQydjrWqcEeRkD14FzjhwU7bmhaR1XFY51NfbCm7ZzTMz7z9y3qFjrgb3i

fd936 := client.Open("/pTzqUWCEtV.txt", client.O_RDWR|client.O_CREATE)
if(fd936 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd910, []byte("LzC9iFR6wItiTm_eI2D1ELO_MUf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (27) LzC9iFR6wItiTm_eI2D1ELO_MUf

fd937 := client.Open("/igl0L0GvBm.txt", client.O_RDWR|client.O_CREATE)
if(fd937 < 0) {
  panic("open failed")
}


ret = client.Close(fd882)
if(ret != 0) {
  panic("close failed")
}

//fd state: (92) t87EI8vprHTuXlNCFMCmmn1GqOSqxpKbFLCoIsni50NcRd8AHTqRP6gRkspmbMctk3P1MR1u_bEHfeOuWyARV0BLH679

ret = client.Write(fd865, []byte("89lFeByAGrK7tcH0dmxESgUkMVPTZZSSrMvQ370HDp_1u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (137) t87EI8vprHTuXlNCFMCmmn1GqOSqxpKbFLCoIsni50NcRd8AHTqRP6gRkspmbMctk3P1MR1u_bEHfeOuWyARV0BLH67989lFeByAGrK7tcH0dmxESgUkMVPTZZSSrMvQ370HDp_1u
//fd state: (0) 

ret = client.Write(fd934, []byte(""))
if(ret != 0) {
  panic("write failed")
}

//fd state: (0) 

fd938 := client.Open("/DIihw8XIVK.txt", client.O_RDWR|client.O_CREATE)
if(fd938 < 0) {
  panic("open failed")
}


fd939 := client.Open("/MBbzK_9s28.txt", client.O_RDWR|client.O_CREATE)
if(fd939 < 0) {
  panic("open failed")
}


ret = client.Close(fd857)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd918, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd937, []byte("sHw1MuENrHD3ahRhy9ESJHR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (23) sHw1MuENrHD3ahRhy9ESJHR

ret = client.Close(fd934)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd924, []byte("tIL_DA6QZEr7oEO57pg5TAq78qT4RVKK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) tIL_DA6QZEr7oEO57pg5TAq78qT4RVKK

ret = client.Close(fd938)
if(ret != 0) {
  panic("close failed")
}


fd940 := client.Open("/ZZSoebNSVw.txt", client.O_RDWR|client.O_CREATE)
if(fd940 < 0) {
  panic("open failed")
}


ret = client.Close(fd907)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd871, 27)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "yoxfXElczlpEyht3bkcwEI19GYw") {
  panic("wrong data returned")
}


ret = client.Close(fd865)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd924, 16, client.SEEK_SET)
if(ret != 16) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 16)
  panic("seek failed")
}


ret = client.Seek(fd940, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd936, []byte("xYQ0yC_3sRbh2dp2n48utynDw3bFtTELXuos1syR8XpUhDXQ6xAoEcat18KrpGIC1w2dm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) xYQ0yC_3sRbh2dp2n48utynDw3bFtTELXuos1syR8XpUhDXQ6xAoEcat18KrpGIC1w2dm

ret = client.Seek(fd827, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


ret = client.Close(fd921)
if(ret != 0) {
  panic("close failed")
}


fd941 := client.Open("/jaPa72Giny.txt", client.O_RDWR|client.O_CREATE)
if(fd941 < 0) {
  panic("open failed")
}


fd942 := client.Open("/fBJDl1Lbkh.txt", client.O_RDWR|client.O_CREATE)
if(fd942 < 0) {
  panic("open failed")
}


ret = client.Close(fd827)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd896, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd871, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "1yV8SKIjBs0H3bcYcj3qFt0Y5y8Dr2g6wk9KYUQSs5p0KVCgtC8hAcdWo3BSFTB1hBQ") {
  panic("wrong data returned")
}


ret = client.Close(fd935)
if(ret != 0) {
  panic("close failed")
}


fd943 := client.Open("/aXUl_uJx3o.txt", client.O_RDWR|client.O_CREATE)
if(fd943 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd916, []byte("LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PG"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PG

buf, ret = client.Read(fd918, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "RDdFhX7aGgUN2zIsira4i2V7M") {
  panic("wrong data returned")
}

//fd state: (44) XAq1p9zxhdnmePpYubPepWpl1TD0SLE_2RGTrJlJUte9

ret = client.Write(fd888, []byte("RSDzTNwYu5gM0AOHZacJP15iZi8OV7Y7Fj5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) XAq1p9zxhdnmePpYubPepWpl1TD0SLE_2RGTrJlJUte9RSDzTNwYu5gM0AOHZacJP15iZi8OV7Y7Fj5

fd944 := client.Open("/_joua8qqGG.txt", client.O_RDWR|client.O_CREATE)
if(fd944 < 0) {
  panic("open failed")
}


fd945 := client.Open("/4TdLNMxHgT.txt", client.O_RDWR|client.O_CREATE)
if(fd945 < 0) {
  panic("open failed")
}


ret = client.Close(fd863)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd931, 150, client.SEEK_SET)
if(ret != 150) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 150)
  panic("seek failed")
}


ret = client.Seek(fd945, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd939, []byte("rRiEIozF5g89hR9ysYDDiNML0zxYKNVk3BZppTCBqKURJqOsav"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (50) rRiEIozF5g89hR9ysYDDiNML0zxYKNVk3BZppTCBqKURJqOsav

ret = client.Seek(fd930, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd918, 0)
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


ret = client.Close(fd922)
if(ret != 0) {
  panic("close failed")
}

//fd state: (28) yJFjTlE7bLGBFhZbI8C7RtfoZwI9Vh

ret = client.Write(fd945, []byte("Uw98L8SVthlAAXjQUhlOUzgOyOzpnzRhQ8Ull6mST6Xrbje64s9PPoeTaSAFcPb"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (91) yJFjTlE7bLGBFhZbI8C7RtfoZwI9Uw98L8SVthlAAXjQUhlOUzgOyOzpnzRhQ8Ull6mST6Xrbje64s9PPoeTaSAFcPb

fd946 := client.Open("/bRTrk37Phi.txt", client.O_RDWR|client.O_CREATE)
if(fd946 < 0) {
  panic("open failed")
}


ret = client.Close(fd871)
if(ret != 0) {
  panic("close failed")
}


fd947 := client.Open("/NPBGZkAtqp.txt", client.O_RDWR|client.O_CREATE)
if(fd947 < 0) {
  panic("open failed")
}

//fd state: (48) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PG

ret = client.Write(fd916, []byte("NaDirOIX0L8lkbJIIY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (66) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PGNaDirOIX0L8lkbJIIY

ret = client.Seek(fd947, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd948 := client.Open("/WmilsB_bL8.txt", client.O_RDWR|client.O_CREATE)
if(fd948 < 0) {
  panic("open failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Close(fd948)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd939, 49, client.SEEK_SET)
if(ret != 49) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 49)
  panic("seek failed")
}

//fd state: (150) tiYschvflbSHcx8esE45g_c3QPfFedZLoXALsAmClMQLai4IAlr_JTDrruN5kfJMBPEK4EOHW9nlSThzp1xa0kMwrDbrsNNjf3RrONe5Oer0M52jFqMbEeUo2sbgF_2LjjrYetUd_biXYsuUI9sy7146arw79cBtwNAvN1oQlHw

ret = client.Write(fd931, []byte("U0vFQSnAkkaSmHQOv26aLpheyEh_N8PAu47_vfDDYIi8BQScDudq3a6WGq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (208) tiYschvflbSHcx8esE45g_c3QPfFedZLoXALsAmClMQLai4IAlr_JTDrruN5kfJMBPEK4EOHW9nlSThzp1xa0kMwrDbrsNNjf3RrONe5Oer0M52jFqMbEeUo2sbgF_2LjjrYetUd_biXYsuUI9sy71U0vFQSnAkkaSmHQOv26aLpheyEh_N8PAu47_vfDDYIi8BQScDudq3a6WGq

fd949 := client.Open("/Aglota8V8i.txt", client.O_RDWR|client.O_CREATE)
if(fd949 < 0) {
  panic("open failed")
}


fd950 := client.Open("/m80E198Few.txt", client.O_RDWR|client.O_CREATE)
if(fd950 < 0) {
  panic("open failed")
}


ret = client.Close(fd949)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd896)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd920)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd930, []byte("1rVl2ooX1bbvNNegJzVq7NC7I9ZIETE_G10TDsazvVKS3mM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) 1rVl2ooX1bbvNNegJzVq7NC7I9ZIETE_G10TDsazvVKS3mM

buf, ret = client.Read(fd925, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd947, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd936, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


buf, ret = client.Read(fd918, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "FgHqxsIZKaa5LmZ9WUrmu5RADIAdvv") {
  panic("wrong data returned")
}

//fd state: (79) XAq1p9zxhdnmePpYubPepWpl1TD0SLE_2RGTrJlJUte9RSDzTNwYu5gM0AOHZacJP15iZi8OV7Y7Fj5

ret = client.Write(fd888, []byte("AcJLZKkM5FeQiRsps5nfBIVq1OYLrW6DVbN3124ze1aVpKqoEK1t5"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (132) XAq1p9zxhdnmePpYubPepWpl1TD0SLE_2RGTrJlJUte9RSDzTNwYu5gM0AOHZacJP15iZi8OV7Y7Fj5AcJLZKkM5FeQiRsps5nfBIVq1OYLrW6DVbN3124ze1aVpKqoEK1t5

fd951 := client.Open("/hHqhVtDSIg.txt", client.O_RDWR|client.O_CREATE)
if(fd951 < 0) {
  panic("open failed")
}

//fd state: (0) rYRmgqf7CHjsZWt4jqaMVOdCCdrO1z93kYKeryDpLvVhHJ0QzLZsdETLineQ

ret = client.Write(fd950, []byte("5VCiVgUfqKCbJizAhRS7t7TuKh_PnjB5KLqH_mkFPqV05OX3yY_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) 5VCiVgUfqKCbJizAhRS7t7TuKh_PnjB5KLqH_mkFPqV05OX3yY_sdETLineQ

fd952 := client.Open("/PW_XUw6X1R.txt", client.O_RDWR|client.O_CREATE)
if(fd952 < 0) {
  panic("open failed")
}


ret = client.Close(fd940)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd950, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Close(fd945)
if(ret != 0) {
  panic("close failed")
}

//fd state: (66) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PGNaDirOIX0L8lkbJIIY

ret = client.Write(fd916, []byte("8T9uy0DG4Jq7160YxqjfTs0TvcUcOiv4zYpsWx12Af6iB8sVeVkPFf42THlsw9R8qetZVcmEmnBHzUm_gWvTV9jgTHwh1Upuc"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (163) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PGNaDirOIX0L8lkbJIIY8T9uy0DG4Jq7160YxqjfTs0TvcUcOiv4zYpsWx12Af6iB8sVeVkPFf42THlsw9R8qetZVcmEmnBHzUm_gWvTV9jgTHwh1Upuc

ret = client.Close(fd941)
if(ret != 0) {
  panic("close failed")
}


fd953 := client.Open("/jullFWfs9H.txt", client.O_RDWR|client.O_CREATE)
if(fd953 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd929, 58)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Is8AO2gQuC5LO0EF1jpG0he_W5QG8gElNCJ5O_kxVxhY59XAoIal2VuPAj") {
  panic("wrong data returned")
}


ret = client.Close(fd953)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd918, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "pM_kK_xkA9PPvTYkjDm8sdqA6nbvQa2SfeTufV53vd810Zqp3Y6TMISNaAGb274RFCPT1KlRr35R4U") {
  panic("wrong data returned")
}


ret = client.Seek(fd950, 54, client.SEEK_SET)
if(ret != 54) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 54)
  panic("seek failed")
}

//fd state: (16) tIL_DA6QZEr7oEO57pg5TAq78qT4RVKK

ret = client.Write(fd924, []byte("w_mVN4zPC9PnhhRzvBl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (35) tIL_DA6QZEr7oEO5w_mVN4zPC9PnhhRzvBl
//fd state: (27) LzC9iFR6wItiTm_eI2D1ELO_MUf

ret = client.Write(fd910, []byte("IB4JhTRaVrVAyd7qI58MDuyam2D0o3JZk14kII2cbTLyyiPY7_Jrn3_NotkqKJAhC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (92) LzC9iFR6wItiTm_eI2D1ELO_MUfIB4JhTRaVrVAyd7qI58MDuyam2D0o3JZk14kII2cbTLyyiPY7_Jrn3_NotkqKJAhC

ret = client.Close(fd931)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd937, 4, client.SEEK_SET)
if(ret != 4) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 4)
  panic("seek failed")
}


fd954 := client.Open("/sINhP7j7hc.txt", client.O_RDWR|client.O_CREATE)
if(fd954 < 0) {
  panic("open failed")
}


ret = client.Close(fd942)
if(ret != 0) {
  panic("close failed")
}

//fd state: (35) tIL_DA6QZEr7oEO5w_mVN4zPC9PnhhRzvBl

ret = client.Write(fd924, []byte("0YPIDToAYOfWmqogAg9psFz2hJECP6UTmXtFe_L0soRcoiKdrsqzPcWRozocrvemUeC7bNrhGiSPJPlSh6M4wYa6cNtE_InLSR2"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (134) tIL_DA6QZEr7oEO5w_mVN4zPC9PnhhRzvBl0YPIDToAYOfWmqogAg9psFz2hJECP6UTmXtFe_L0soRcoiKdrsqzPcWRozocrvemUeC7bNrhGiSPJPlSh6M4wYa6cNtE_InLSR2
//fd state: (4) sHw1MuENrHD3ahRhy9ESJHR

ret = client.Write(fd937, []byte("VACC3ObrUGlr673yz6acDJrEErjJWJWQWqEEqkd3juN3_b42irp7cu6mu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) sHw1VACC3ObrUGlr673yz6acDJrEErjJWJWQWqEEqkd3juN3_b42irp7cu6mu

ret = client.Seek(fd888, 21, client.SEEK_SET)
if(ret != 21) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 21)
  panic("seek failed")
}


buf, ret = client.Read(fd930, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd916, 116, client.SEEK_SET)
if(ret != 116) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 116)
  panic("seek failed")
}


ret = client.Close(fd918)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd952, []byte("2stqOktIjc03B3JrWkbG7OpqeJ4HeX6gAENC6McNdyiDsOiQuJLrKGecFN9ZFiTggdDb0kLsF3R418uB7GcF_"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) 2stqOktIjc03B3JrWkbG7OpqeJ4HeX6gAENC6McNdyiDsOiQuJLrKGecFN9ZFiTggdDb0kLsF3R418uB7GcF_

ret = client.Seek(fd925, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


buf, ret = client.Read(fd947, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd929, 95, client.SEEK_SET)
if(ret != 95) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 95)
  panic("seek failed")
}


ret = client.Seek(fd950, 27, client.SEEK_SET)
if(ret != 27) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 27)
  panic("seek failed")
}


ret = client.Close(fd925)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd946, []byte("MMlKUte_f5QatcnPocQ1YXQ1YWkPxd1FUOJk2FYdubv4rE6vq6zw0gKxTZm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (59) MMlKUte_f5QatcnPocQ1YXQ1YWkPxd1FUOJk2FYdubv4rE6vq6zw0gKxTZm

ret = client.Close(fd954)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd910, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd946, 30)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd947)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd937, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd955 := client.Open("/XdnFXqiD3I.txt", client.O_RDWR|client.O_CREATE)
if(fd955 < 0) {
  panic("open failed")
}


ret = client.Seek(fd924, 14, client.SEEK_SET)
if(ret != 14) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 14)
  panic("seek failed")
}


fd956 := client.Open("/us2x7vtui0.txt", client.O_RDWR|client.O_CREATE)
if(fd956 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd936, 21)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "XQ6xAoEcat18KrpGIC1w2") {
  panic("wrong data returned")
}


ret = client.Seek(fd930, 45, client.SEEK_SET)
if(ret != 45) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 45)
  panic("seek failed")
}


ret = client.Close(fd937)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd924)
if(ret != 0) {
  panic("close failed")
}


fd957 := client.Open("/PkNTTq3tWC.txt", client.O_RDWR|client.O_CREATE)
if(fd957 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd910, 96)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd916, 71, client.SEEK_SET)
if(ret != 71) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 71)
  panic("seek failed")
}


buf, ret = client.Read(fd888, 55)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "Wpl1TD0SLE_2RGTrJlJUte9RSDzTNwYu5gM0AOHZacJP15iZi8OV7Y7") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd950, 32)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "PnjB5KLqH_mkFPqV05OX3yY_sdETLine") {
  panic("wrong data returned")
}


ret = client.Close(fd952)
if(ret != 0) {
  panic("close failed")
}

//fd state: (164) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduCefufyMh63jhejAMMs_WsmEXF5PQ0vOGMfutut0VLIuHRKk5TxwhJDQy8RnbgBb3hDFonL9zPdCZB46ma7NO6MCoOgkI6VHAhmnB9uYlsWcigbMVtEtcNkWOr6lTqbm3R91df5eMpxVbYiWkCUtKwSDrvR_VMLeTmWb4wzlZIKIx

ret = client.Write(fd908, []byte("K9v8A2HRRNGNs3XrJoZzhjZiWrs17"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (193) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduCefufyMh63jhejAMMs_WsmEXF5PQ0vOGMfutut0VLIuHRKk5TxwhJDQy8RnbgBb3hDFonL9zPdCZB46ma7NO6MCoOgkI6VHAhmnB9uYlsWcigbMVtEtK9v8A2HRRNGNs3XrJoZzhjZiWrs17CUtKwSDrvR_VMLeTmWb4wzlZIKIx

fd958 := client.Open("/A3BV1OD5Bq.txt", client.O_RDWR|client.O_CREATE)
if(fd958 < 0) {
  panic("open failed")
}

//fd state: (45) 1rVl2ooX1bbvNNegJzVq7NC7I9ZIETE_G10TDsazvVKS3mM

ret = client.Write(fd930, []byte("sqarSallsKB1ghMxxJ1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (64) 1rVl2ooX1bbvNNegJzVq7NC7I9ZIETE_G10TDsazvVKS3sqarSallsKB1ghMxxJ1
//fd state: (59) MMlKUte_f5QatcnPocQ1YXQ1YWkPxd1FUOJk2FYdubv4rE6vq6zw0gKxTZm

ret = client.Write(fd946, []byte("cwv9GFzrJs90rxyYaVOP4VsVZQSEFEDla4nbYapDnhBUnYxRPCcn4JZqbA"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (117) MMlKUte_f5QatcnPocQ1YXQ1YWkPxd1FUOJk2FYdubv4rE6vq6zw0gKxTZmcwv9GFzrJs90rxyYaVOP4VsVZQSEFEDla4nbYapDnhBUnYxRPCcn4JZqbA

ret = client.Close(fd946)
if(ret != 0) {
  panic("close failed")
}


fd959 := client.Open("/a1Vqa8McUR.txt", client.O_RDWR|client.O_CREATE)
if(fd959 < 0) {
  panic("open failed")
}


fd960 := client.Open("/ModFGCjTvL.txt", client.O_RDWR|client.O_CREATE)
if(fd960 < 0) {
  panic("open failed")
}


ret = client.Seek(fd956, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (95) w_icxKWZ9sVtHdec32IR1oQ96LH9PfMWbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WPIs8AO2gQuC5LO0EF1jpG0he_W5QG8gElNCJ5O_kxVxhY59XAoIal2VuPAjYARbgubjiyTuJ7yb8qwIUU

ret = client.Write(fd929, []byte("ZFz0S"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) w_icxKWZ9sVtHdec32IR1oQ96LH9PfMWbLygiz7G0t7aSCKHskID3FduwB3c0xdxxhvCpxAhAe1b0WPIs8AO2gQuC5LO0EFZFz0She_W5QG8gElNCJ5O_kxVxhY59XAoIal2VuPAjYARbgubjiyTuJ7yb8qwIUU

ret = client.Close(fd910)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd888)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd955, 7, client.SEEK_SET)
if(ret != 7) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 7)
  panic("seek failed")
}

//fd state: (49) rRiEIozF5g89hR9ysYDDiNML0zxYKNVk3BZppTCBqKURJqOsav

ret = client.Write(fd939, []byte("Bt9ko3c4ZuvdrRRsRHQ7DTyYy9CWyJy9WXt1lZRXIwLJ7E"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) rRiEIozF5g89hR9ysYDDiNML0zxYKNVk3BZppTCBqKURJqOsaBt9ko3c4ZuvdrRRsRHQ7DTyYy9CWyJy9WXt1lZRXIwLJ7E

ret = client.Seek(fd908, 178, client.SEEK_SET)
if(ret != 178) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 178)
  panic("seek failed")
}


buf, ret = client.Read(fd955, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "eEdDP") {
  panic("wrong data returned")
}

//fd state: (59) 5VCiVgUfqKCbJizAhRS7t7TuKh_PnjB5KLqH_mkFPqV05OX3yY_sdETLineQ

ret = client.Write(fd950, []byte("1V0rQH7I7advExPLabITU5LwwqaW7gnZEczdSJLzGKgWNXVyb_etOLTzPqPxlOtzQRokOAK3HKSDUowGn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (140) 5VCiVgUfqKCbJizAhRS7t7TuKh_PnjB5KLqH_mkFPqV05OX3yY_sdETLine1V0rQH7I7advExPLabITU5LwwqaW7gnZEczdSJLzGKgWNXVyb_etOLTzPqPxlOtzQRokOAK3HKSDUowGn

buf, ret = client.Read(fd950, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd916, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "0DG4Jq7160YxqjfTs0TvcUcOiv4zYpsWx12Af6iB8sVeVkPFf42THlsw9R8qetZVcm") {
  panic("wrong data returned")
}


ret = client.Close(fd939)
if(ret != 0) {
  panic("close failed")
}


fd961 := client.Open("/sGEi4tWRrI.txt", client.O_RDWR|client.O_CREATE)
if(fd961 < 0) {
  panic("open failed")
}


ret = client.Close(fd955)
if(ret != 0) {
  panic("close failed")
}


fd962 := client.Open("/1fD0ZWvax2.txt", client.O_RDWR|client.O_CREATE)
if(fd962 < 0) {
  panic("open failed")
}

//fd state: (67) xYQ0yC_3sRbh2dp2n48utynDw3bFtTELXuos1syR8XpUhDXQ6xAoEcat18KrpGIC1w2dm

ret = client.Write(fd936, []byte("jz8ier_ghvrmZozD1khPg_Xxusa_qTjeW3Pti3g29AA1r5z2KXEBPEpADalqvfJgfS71hVwTnGWCMcq35yDpavK9r"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (156) xYQ0yC_3sRbh2dp2n48utynDw3bFtTELXuos1syR8XpUhDXQ6xAoEcat18KrpGIC1w2jz8ier_ghvrmZozD1khPg_Xxusa_qTjeW3Pti3g29AA1r5z2KXEBPEpADalqvfJgfS71hVwTnGWCMcq35yDpavK9r
//fd state: (0) 

ret = client.Write(fd951, []byte("c"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (1) c

fd963 := client.Open("/jioZXwkhDi.txt", client.O_RDWR|client.O_CREATE)
if(fd963 < 0) {
  panic("open failed")
}


fd964 := client.Open("/qfwUZ2XRmv.txt", client.O_RDWR|client.O_CREATE)
if(fd964 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd956, []byte("kPsqQKeEh6YIVB3UxV6gNpYiUdD3zc6G789Z42YpiIuL5CYaZJ7EHM2dA5RqdeUNvnQohwExeaGBSO_N1Ry"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (83) kPsqQKeEh6YIVB3UxV6gNpYiUdD3zc6G789Z42YpiIuL5CYaZJ7EHM2dA5RqdeUNvnQohwExeaGBSO_N1Ry

buf, ret = client.Read(fd950, 69)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd962, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd936, 83, client.SEEK_SET)
if(ret != 83) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 83)
  panic("seek failed")
}


ret = client.Seek(fd943, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd961, 29)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd965 := client.Open("/ghuzLxuO9g.txt", client.O_RDWR|client.O_CREATE)
if(fd965 < 0) {
  panic("open failed")
}


ret = client.Close(fd943)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd951)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd936)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd908, 80, client.SEEK_SET)
if(ret != 80) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 80)
  panic("seek failed")
}


fd966 := client.Open("/3Xuls7KVgK.txt", client.O_RDWR|client.O_CREATE)
if(fd966 < 0) {
  panic("open failed")
}


ret = client.Seek(fd950, 73, client.SEEK_SET)
if(ret != 73) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 73)
  panic("seek failed")
}


ret = client.Close(fd958)
if(ret != 0) {
  panic("close failed")
}


fd967 := client.Open("/EtryQom2t1.txt", client.O_RDWR|client.O_CREATE)
if(fd967 < 0) {
  panic("open failed")
}


fd968 := client.Open("/i2P0PZS57Y.txt", client.O_RDWR|client.O_CREATE)
if(fd968 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd962, []byte("F3zxoEX8la8uUBPXXNnvLTuVi8MF3itgR_2Y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (36) F3zxoEX8la8uUBPXXNnvLTuVi8MF3itgR_2Y

buf, ret = client.Read(fd908, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "GMfutut0VLIuHRKk5TxwhJDQy8RnbgBb3hDFonL9zPdCZB46ma7NO") {
  panic("wrong data returned")
}


ret = client.Seek(fd968, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) R5keHUpQVdut1RMooVQdXAKw94Z

ret = client.Write(fd960, []byte("qLytBiaHoMoTg2N_2veyrHcHr4We6YWUPWz4fdTcjnA4AXLW5zokPIeYQkjeTqxnK"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) qLytBiaHoMoTg2N_2veyrHcHr4We6YWUPWz4fdTcjnA4AXLW5zokPIeYQkjeTqxnK

fd969 := client.Open("/P07sWxGvr6.txt", client.O_RDWR|client.O_CREATE)
if(fd969 < 0) {
  panic("open failed")
}


ret = client.Close(fd964)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd929)
if(ret != 0) {
  panic("close failed")
}


fd970 := client.Open("/v2DsGdHGbp.txt", client.O_RDWR|client.O_CREATE)
if(fd970 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd966, 91)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "IRolh6rlrs9L9P8Gw8RHzzj5DWSeYwzgvBSflqDO5DgxRcAoi3SGu6iDLtXMm9WtQnaI4MzXqi7yRG46Yq2NcbSm8wV") {
  panic("wrong data returned")
}


ret = client.Close(fd950)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd967, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd970, []byte("xJ21MfQZLujOTNcIgbSd_Lknn1nWBK7LqcpE5_3vuvcwZoWDNA7eMSE9bpnF"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (60) xJ21MfQZLujOTNcIgbSd_Lknn1nWBK7LqcpE5_3vuvcwZoWDNA7eMSE9bpnF
//fd state: (0) 

ret = client.Write(fd957, []byte("zR5pvhU0cAIs3_0d0j4atL5EwBKCCY4zA_DfotqmtTNGuiGv9HkohIB29UzbA6kIlcILbWCLfg_D"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (76) zR5pvhU0cAIs3_0d0j4atL5EwBKCCY4zA_DfotqmtTNGuiGv9HkohIB29UzbA6kIlcILbWCLfg_D

buf, ret = client.Read(fd969, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd971 := client.Open("/VfSkX9551W.txt", client.O_RDWR|client.O_CREATE)
if(fd971 < 0) {
  panic("open failed")
}


ret = client.Seek(fd963, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd963, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd969, []byte("JTrzYjM4YNFGUaOW_iV86TrQXYJoIpMKbM6OW5S2NXcQjDnP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (48) JTrzYjM4YNFGUaOW_iV86TrQXYJoIpMKbM6OW5S2NXcQjDnP

ret = client.Seek(fd960, 51, client.SEEK_SET)
if(ret != 51) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 51)
  panic("seek failed")
}


ret = client.Close(fd967)
if(ret != 0) {
  panic("close failed")
}


fd972 := client.Open("/Vj7T1Yr3JX.txt", client.O_RDWR|client.O_CREATE)
if(fd972 < 0) {
  panic("open failed")
}


ret = client.Close(fd957)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd966)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd956, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd916, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "EmnBHzUm_gWvTV9jgTHwh1Upuc") {
  panic("wrong data returned")
}


fd973 := client.Open("/_KHs6d6YBx.txt", client.O_RDWR|client.O_CREATE)
if(fd973 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd956, 66)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd968, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd916, 45)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd973)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd956, 28, client.SEEK_SET)
if(ret != 28) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 28)
  panic("seek failed")
}

//fd state: (36) F3zxoEX8la8uUBPXXNnvLTuVi8MF3itgR_2Y

ret = client.Write(fd962, []byte("VYcKm6mxs1tFfNuqhfpV8Ua5ZcAPIZEzssdDq8aA2MqZ4T7Uq61Iglawtr9vaFNvynczuJl2uv"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (110) F3zxoEX8la8uUBPXXNnvLTuVi8MF3itgR_2YVYcKm6mxs1tFfNuqhfpV8Ua5ZcAPIZEzssdDq8aA2MqZ4T7Uq61Iglawtr9vaFNvynczuJl2uv

fd974 := client.Open("/tuVdxQFOfF.txt", client.O_RDWR|client.O_CREATE)
if(fd974 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd968, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd962, 100, client.SEEK_SET)
if(ret != 100) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 100)
  panic("seek failed")
}


ret = client.Seek(fd963, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd975 := client.Open("/WDmAlke9N0.txt", client.O_RDWR|client.O_CREATE)
if(fd975 < 0) {
  panic("open failed")
}


ret = client.Seek(fd972, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd963, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd976 := client.Open("/MF_bH4w72p.txt", client.O_RDWR|client.O_CREATE)
if(fd976 < 0) {
  panic("open failed")
}


ret = client.Seek(fd930, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Close(fd971)
if(ret != 0) {
  panic("close failed")
}

//fd state: (28) kPsqQKeEh6YIVB3UxV6gNpYiUdD3zc6G789Z42YpiIuL5CYaZJ7EHM2dA5RqdeUNvnQohwExeaGBSO_N1Ry

ret = client.Write(fd956, []byte("LLSvz3CeEgcYULW1meY8FCDO5qLNgJtJW5wxT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (65) kPsqQKeEh6YIVB3UxV6gNpYiUdD3LLSvz3CeEgcYULW1meY8FCDO5qLNgJtJW5wxTnQohwExeaGBSO_N1Ry

ret = client.Seek(fd969, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Seek(fd930, 59, client.SEEK_SET)
if(ret != 59) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 59)
  panic("seek failed")
}

//fd state: (0) bQ4SElzbXUD7JxDzu4fS0zX006ng5KR8IO4ZwDOOOhYUtx2s_eUC0x0N34dxpQZYfKawPQ8_L5CA4d8Wzo0uqZ2frTMn9mjuQWOIpKrcBWOWooB6x6PLmAFha679KQ9Slpel58Jx1MfeZ3CVgyay

ret = client.Write(fd976, []byte("B3VCID_sYC4Ebyo84OKQKdPTOkydoJfdkOlGO7yx7V7EtQj3m0tijaCX66eB0hzPCz3HI7bjIhckeM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (78) B3VCID_sYC4Ebyo84OKQKdPTOkydoJfdkOlGO7yx7V7EtQj3m0tijaCX66eB0hzPCz3HI7bjIhckeM8Wzo0uqZ2frTMn9mjuQWOIpKrcBWOWooB6x6PLmAFha679KQ9Slpel58Jx1MfeZ3CVgyay

fd977 := client.Open("/k8VXHxxYvV.txt", client.O_RDWR|client.O_CREATE)
if(fd977 < 0) {
  panic("open failed")
}


fd978 := client.Open("/cHSpTAv2xS.txt", client.O_RDWR|client.O_CREATE)
if(fd978 < 0) {
  panic("open failed")
}


fd979 := client.Open("/GxEOkRdSJm.txt", client.O_RDWR|client.O_CREATE)
if(fd979 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd972, 44)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd908, 106, client.SEEK_SET)
if(ret != 106) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 106)
  panic("seek failed")
}


fd980 := client.Open("/zMkKhqfeQu.txt", client.O_RDWR|client.O_CREATE)
if(fd980 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd959, []byte("48q85lrnrDKNAal6B2q"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (19) 48q85lrnrDKNAal6B2q

fd981 := client.Open("/CCVft7HP0n.txt", client.O_RDWR|client.O_CREATE)
if(fd981 < 0) {
  panic("open failed")
}

//fd state: (41) JTrzYjM4YNFGUaOW_iV86TrQXYJoIpMKbM6OW5S2NXcQjDnP

ret = client.Write(fd969, []byte("pZ0br"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) JTrzYjM4YNFGUaOW_iV86TrQXYJoIpMKbM6OW5S2NpZ0brnP

ret = client.Seek(fd979, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd962)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd976)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd977, 62)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd972, []byte("VlmGL1xbotRxZJ3UpHKNDnTgeCbZCuzyjnjOJoh85_CWDtCJQvEoIbhsG3L2kBJts3YsvowgaKwHzhdQ9IiAP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (85) VlmGL1xbotRxZJ3UpHKNDnTgeCbZCuzyjnjOJoh85_CWDtCJQvEoIbhsG3L2kBJts3YsvowgaKwHzhdQ9IiAP

buf, ret = client.Read(fd970, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd982 := client.Open("/rCKJKWJOO0.txt", client.O_RDWR|client.O_CREATE)
if(fd982 < 0) {
  panic("open failed")
}


ret = client.Close(fd956)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd974, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd983 := client.Open("/P4ToA7VMb5.txt", client.O_RDWR|client.O_CREATE)
if(fd983 < 0) {
  panic("open failed")
}


ret = client.Seek(fd974, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd930)
if(ret != 0) {
  panic("close failed")
}

//fd state: (60) xJ21MfQZLujOTNcIgbSd_Lknn1nWBK7LqcpE5_3vuvcwZoWDNA7eMSE9bpnF

ret = client.Write(fd970, []byte("Yi46tGYYdQM03f3PQ0JkXpc4N0EFO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (89) xJ21MfQZLujOTNcIgbSd_Lknn1nWBK7LqcpE5_3vuvcwZoWDNA7eMSE9bpnFYi46tGYYdQM03f3PQ0JkXpc4N0EFO

buf, ret = client.Read(fd983, 99)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd981, []byte("AKHjhivSe6Iu3"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (13) AKHjhivSe6Iu3

ret = client.Close(fd974)
if(ret != 0) {
  panic("close failed")
}


fd984 := client.Open("/Fx14wmZV6Z.txt", client.O_RDWR|client.O_CREATE)
if(fd984 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd983, []byte("f2Nbs4T2rpE8CXnzcRf_Fs7j4eXfhks1Ot6_LLaCK9X3rGsMdajUx5GTeN9QrOXI2YX_YYh5u"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) f2Nbs4T2rpE8CXnzcRf_Fs7j4eXfhks1Ot6_LLaCK9X3rGsMdajUx5GTeN9QrOXI2YX_YYh5u

buf, ret = client.Read(fd984, 11)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "43ORMZCWPz2") {
  panic("wrong data returned")
}

//fd state: (73) f2Nbs4T2rpE8CXnzcRf_Fs7j4eXfhks1Ot6_LLaCK9X3rGsMdajUx5GTeN9QrOXI2YX_YYh5u

ret = client.Write(fd983, []byte("QT5UAT5mndoD_5BOLnyon7R0fAt1SvDUqUeMNQrUrxvB7gm9lBJCdDEfMo3xv7f"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (136) f2Nbs4T2rpE8CXnzcRf_Fs7j4eXfhks1Ot6_LLaCK9X3rGsMdajUx5GTeN9QrOXI2YX_YYh5uQT5UAT5mndoD_5BOLnyon7R0fAt1SvDUqUeMNQrUrxvB7gm9lBJCdDEfMo3xv7f

ret = client.Seek(fd975, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd981)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd972, 60)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd979, []byte("eCWUNku92jv3I5MQVhMh2cj6s4wnvedGbxyxOMjnmAvYe"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (45) eCWUNku92jv3I5MQVhMh2cj6s4wnvedGbxyxOMjnmAvYe

fd985 := client.Open("/z3BIYvfpSJ.txt", client.O_RDWR|client.O_CREATE)
if(fd985 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd963, []byte("WQMS1wcrR178fzA4Jt8QKLyCaBPlgG4A"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (32) WQMS1wcrR178fzA4Jt8QKLyCaBPlgG4A

ret = client.Close(fd960)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd975)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd984, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


ret = client.Close(fd959)
if(ret != 0) {
  panic("close failed")
}

//fd state: (32) WQMS1wcrR178fzA4Jt8QKLyCaBPlgG4A

ret = client.Write(fd963, []byte("EXef5hHFESv2TmMONd1jGysyH5RsYjsIShqnC_s8nXTNeYNU"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) WQMS1wcrR178fzA4Jt8QKLyCaBPlgG4AEXef5hHFESv2TmMONd1jGysyH5RsYjsIShqnC_s8nXTNeYNU

buf, ret = client.Read(fd965, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "o4NLujFJmcfM_gx9YigRiVk4r5dvrdPpYVmXUCZpEktcOrnv0V5swWqRJHxWK1YHGDZp3gXQTpWRtCyt_rO") {
  panic("wrong data returned")
}


ret = client.Seek(fd968, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd979, 26)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd986 := client.Open("/oGnzWegge0.txt", client.O_RDWR|client.O_CREATE)
if(fd986 < 0) {
  panic("open failed")
}


fd987 := client.Open("/6ryJr0VJ2o.txt", client.O_RDWR|client.O_CREATE)
if(fd987 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd986, []byte("4V2rrHZogWcsxe5BoiA7xnu6XEui0AEFCWo5Pvn9VR8hvKR0ttm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (51) 4V2rrHZogWcsxe5BoiA7xnu6XEui0AEFCWo5Pvn9VR8hvKR0ttm
//fd state: (0) 

ret = client.Write(fd977, []byte("DLpM_cdkC7GGJdAH3s3GqVDnlYE7UMSpY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (33) DLpM_cdkC7GGJdAH3s3GqVDnlYE7UMSpY

buf, ret = client.Read(fd916, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd968, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd988 := client.Open("/5OvPSr9_7r.txt", client.O_RDWR|client.O_CREATE)
if(fd988 < 0) {
  panic("open failed")
}


ret = client.Seek(fd968, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd987, 31)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (163) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PGNaDirOIX0L8lkbJIIY8T9uy0DG4Jq7160YxqjfTs0TvcUcOiv4zYpsWx12Af6iB8sVeVkPFf42THlsw9R8qetZVcmEmnBHzUm_gWvTV9jgTHwh1Upuc

ret = client.Write(fd916, []byte("kNc2m7MLQVgISBibVhpqi6JFBQk0WDi0TMjcqFJ2pogTeWYA7D4xJBtCwlUD1XidlwQhZoqHq6gpQ7vLdo26nQYvSe_s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (255) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PGNaDirOIX0L8lkbJIIY8T9uy0DG4Jq7160YxqjfTs0TvcUcOiv4zYpsWx12Af6iB8sVeVkPFf42THlsw9R8qetZVcmEmnBHzUm_gWvTV9jgTHwh1UpuckNc2m7MLQVgISBibVhpqi6JFBQk0WDi0TMjcqFJ2pogTeWYA7D4xJBtCwlUD1XidlwQhZoqHq6gpQ7vLdo26nQYvSe_s

ret = client.Seek(fd970, 36, client.SEEK_SET)
if(ret != 36) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 36)
  panic("seek failed")
}


ret = client.Seek(fd965, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd989 := client.Open("/FZZ8y6anbP.txt", client.O_RDWR|client.O_CREATE)
if(fd989 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd988, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "lxlQgnvZQcRGE_xFYAz3e5Jvv5XYk1VEqlHdiJChAJTZcTj") {
  panic("wrong data returned")
}


ret = client.Seek(fd963, 31, client.SEEK_SET)
if(ret != 31) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 31)
  panic("seek failed")
}


ret = client.Close(fd989)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd984, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}


ret = client.Seek(fd968, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd987, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd908, 100, client.SEEK_SET)
if(ret != 100) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 100)
  panic("seek failed")
}


fd990 := client.Open("/dPntpWiOAH.txt", client.O_RDWR|client.O_CREATE)
if(fd990 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd986, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd969, 22)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "nP") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd961, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd991 := client.Open("/wpxBs2zOJa.txt", client.O_RDWR|client.O_CREATE)
if(fd991 < 0) {
  panic("open failed")
}


fd992 := client.Open("/fTKEpwYowY.txt", client.O_RDWR|client.O_CREATE)
if(fd992 < 0) {
  panic("open failed")
}


ret = client.Seek(fd990, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd993 := client.Open("/KuBxwFzoD6.txt", client.O_RDWR|client.O_CREATE)
if(fd993 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd982, 53)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oar_tQZ923INkyNkqwOzoat2QB3jEx8HvhyoHrO1piTDMPUx6Q4bc") {
  panic("wrong data returned")
}


ret = client.Seek(fd970, 20, client.SEEK_SET)
if(ret != 20) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 20)
  panic("seek failed")
}


fd994 := client.Open("/ok6_v9o6CM.txt", client.O_RDWR|client.O_CREATE)
if(fd994 < 0) {
  panic("open failed")
}


ret = client.Seek(fd990, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd995 := client.Open("/gwWCYYUhIS.txt", client.O_RDWR|client.O_CREATE)
if(fd995 < 0) {
  panic("open failed")
}

//fd state: (100) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduCefufyMh63jhejAMMs_WsmEXF5PQ0vOGMfutut0VLIuHRKk5TxwhJDQy8RnbgBb3hDFonL9zPdCZB46ma7NO6MCoOgkI6VHAhmnB9uYlsWcigbMVtEtK9v8A2HRRNGNs3XrJoZzhjZiWrs17CUtKwSDrvR_VMLeTmWb4wzlZIKIx

ret = client.Write(fd908, []byte("VHe11Kq"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (107) 5FeCaVzju_TlX1MXQ7glL9N5x_JZCEcXNoAXwouWjiz5hVcduCefufyMh63jhejAMMs_WsmEXF5PQ0vOGMfutut0VLIuHRKk5TxwVHe11KqnbgBb3hDFonL9zPdCZB46ma7NO6MCoOgkI6VHAhmnB9uYlsWcigbMVtEtK9v8A2HRRNGNs3XrJoZzhjZiWrs17CUtKwSDrvR_VMLeTmWb4wzlZIKIx

buf, ret = client.Read(fd961, 93)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd996 := client.Open("/ATxjoIZR9n.txt", client.O_RDWR|client.O_CREATE)
if(fd996 < 0) {
  panic("open failed")
}


ret = client.Seek(fd993, 35, client.SEEK_SET)
if(ret != 35) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 35)
  panic("seek failed")
}


ret = client.Seek(fd985, 22, client.SEEK_SET)
if(ret != 22) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 22)
  panic("seek failed")
}


fd997 := client.Open("/sioso0n3f8.txt", client.O_RDWR|client.O_CREATE)
if(fd997 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd997, []byte("5Y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (2) 5Y

ret = client.Seek(fd970, 33, client.SEEK_SET)
if(ret != 33) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 33)
  panic("seek failed")
}


buf, ret = client.Read(fd963, 85)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "AEXef5hHFESv2TmMONd1jGysyH5RsYjsIShqnC_s8nXTNeYNU") {
  panic("wrong data returned")
}


fd998 := client.Open("/adn_5gM8I7.txt", client.O_RDWR|client.O_CREATE)
if(fd998 < 0) {
  panic("open failed")
}


fd999 := client.Open("/MhNRYjnYbb.txt", client.O_RDWR|client.O_CREATE)
if(fd999 < 0) {
  panic("open failed")
}


ret = client.Seek(fd985, 76, client.SEEK_SET)
if(ret != 76) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 76)
  panic("seek failed")
}

//fd state: (33) xJ21MfQZLujOTNcIgbSd_Lknn1nWBK7LqcpE5_3vuvcwZoWDNA7eMSE9bpnFYi46tGYYdQM03f3PQ0JkXpc4N0EFO

ret = client.Write(fd970, []byte("v309arK_kuApjggpmo_z1YAlYonelu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (63) xJ21MfQZLujOTNcIgbSd_Lknn1nWBK7Lqv309arK_kuApjggpmo_z1YAlYonelu6tGYYdQM03f3PQ0JkXpc4N0EFO
//fd state: (0) 

ret = client.Write(fd980, []byte("CzTACK7MDCRuyJSCzSSihxXKnBXfOkf1jQr5acHmHYP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) CzTACK7MDCRuyJSCzSSihxXKnBXfOkf1jQr5acHmHYP

ret = client.Seek(fd983, 132, client.SEEK_SET)
if(ret != 132) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 132)
  panic("seek failed")
}


ret = client.Seek(fd990, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd990, []byte("DFK8ImPtDhWky7YM26ji4F56nBLrd77riK7ewgsRFYvPi4cJkeHFuW"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (54) DFK8ImPtDhWky7YM26ji4F56nBLrd77riK7ewgsRFYvPi4cJkeHFuW

ret = client.Seek(fd979, 42, client.SEEK_SET)
if(ret != 42) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 42)
  panic("seek failed")
}


fd1000 := client.Open("/px3F6TMyCR.txt", client.O_RDWR|client.O_CREATE)
if(fd1000 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd965, 46)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "LujFJmcfM_gx9YigRiVk4r5dvrdPpYVmXUCZpEktcOrnv0") {
  panic("wrong data returned")
}


ret = client.Close(fd972)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd984, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rbYmsl1Oh2RFWFouN") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd994, []byte("GGxVbsK87FcxhKmK5QsljrayXeeZXtyQjRE8Vs2EQ1y"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (43) GGxVbsK87FcxhKmK5QsljrayXeeZXtyQjRE8Vs2EQ1y
//fd state: (65) 43ORMZCWPz2XjGkTjmrHN59C3nH6L95mB8M4ngpYRu_rrG6jrbYmsl1Oh2RFWFouN

ret = client.Write(fd984, []byte("9JjkX5BAvs1WNkHgDLthvBo542XfQf"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (95) 43ORMZCWPz2XjGkTjmrHN59C3nH6L95mB8M4ngpYRu_rrG6jrbYmsl1Oh2RFWFouN9JjkX5BAvs1WNkHgDLthvBo542XfQf

ret = client.Seek(fd992, 93, client.SEEK_SET)
if(ret != 93) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 93)
  panic("seek failed")
}

//fd state: (48) JTrzYjM4YNFGUaOW_iV86TrQXYJoIpMKbM6OW5S2NpZ0brnP

ret = client.Write(fd969, []byte("9GW2GyZCRAhKr3Kjp7AzWbpZy7hY9n89oVnb94b6vjhvBmyfXzoMl6MqIa6VZlKk"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (112) JTrzYjM4YNFGUaOW_iV86TrQXYJoIpMKbM6OW5S2NpZ0brnP9GW2GyZCRAhKr3Kjp7AzWbpZy7hY9n89oVnb94b6vjhvBmyfXzoMl6MqIa6VZlKk

ret = client.Close(fd970)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd999)
if(ret != 0) {
  panic("close failed")
}

//fd state: (54) DFK8ImPtDhWky7YM26ji4F56nBLrd77riK7ewgsRFYvPi4cJkeHFuW

ret = client.Write(fd990, []byte("hEWwwGfBx6ibrJ5tUHec57PmG1yVXMRd5tSWqL5zvCKPK0MAAyoZPosjA9Pgn"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (115) DFK8ImPtDhWky7YM26ji4F56nBLrd77riK7ewgsRFYvPi4cJkeHFuWhEWwwGfBx6ibrJ5tUHec57PmG1yVXMRd5tSWqL5zvCKPK0MAAyoZPosjA9Pgn

ret = client.Seek(fd992, 32, client.SEEK_SET)
if(ret != 32) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 32)
  panic("seek failed")
}


buf, ret = client.Read(fd988, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1001 := client.Open("/jnES8PlKT_.txt", client.O_RDWR|client.O_CREATE)
if(fd1001 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd978, 90)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd998, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd985, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}

//fd state: (255) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PGNaDirOIX0L8lkbJIIY8T9uy0DG4Jq7160YxqjfTs0TvcUcOiv4zYpsWx12Af6iB8sVeVkPFf42THlsw9R8qetZVcmEmnBHzUm_gWvTV9jgTHwh1UpuckNc2m7MLQVgISBibVhpqi6JFBQk0WDi0TMjcqFJ2pogTeWYA7D4xJBtCwlUD1XidlwQhZoqHq6gpQ7vLdo26nQYvSe_s

ret = client.Write(fd916, []byte("5yaI5xps33zhk2qzh0H5ZbePdkIh0yw1M4EHH1y5_e9d74xDE9F0qcwoUGGjwUCN16jmP"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (324) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PGNaDirOIX0L8lkbJIIY8T9uy0DG4Jq7160YxqjfTs0TvcUcOiv4zYpsWx12Af6iB8sVeVkPFf42THlsw9R8qetZVcmEmnBHzUm_gWvTV9jgTHwh1UpuckNc2m7MLQVgISBibVhpqi6JFBQk0WDi0TMjcqFJ2pogTeWYA7D4xJBtCwlUD1XidlwQhZoqHq6gpQ7vLdo26nQYvSe_s5yaI5xps33zhk2qzh0H5ZbePdkIh0yw1M4EHH1y5_e9d74xDE9F0qcwoUGGjwUCN16jmP

fd1002 := client.Open("/WNMJT2E_a5.txt", client.O_RDWR|client.O_CREATE)
if(fd1002 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1001, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (32) TN8kVvnxmZlrz9LZIx0qchpsVK2vGk8Q_jtNYOHPPl4PvD7Kfwz0l8C3HB8UBL3iua857oa_Lj9yqJljlBUWI85p0zZFOmkKOImptSUrIPWMUEO0xrycDuInc_

ret = client.Write(fd992, []byte("WWrBSe6LxSInmj2iaxUNHHYie1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (58) TN8kVvnxmZlrz9LZIx0qchpsVK2vGk8QWWrBSe6LxSInmj2iaxUNHHYie18UBL3iua857oa_Lj9yqJljlBUWI85p0zZFOmkKOImptSUrIPWMUEO0xrycDuInc_

buf, ret = client.Read(fd988, 40)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (43) GGxVbsK87FcxhKmK5QsljrayXeeZXtyQjRE8Vs2EQ1y

ret = client.Write(fd994, []byte("A_YcDiLdv0uoOUcdalwJtZEMazdIkgyN5tv6"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) GGxVbsK87FcxhKmK5QsljrayXeeZXtyQjRE8Vs2EQ1yA_YcDiLdv0uoOUcdalwJtZEMazdIkgyN5tv6

buf, ret = client.Read(fd987, 35)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1003 := client.Open("/nAZdvLNkiZ.txt", client.O_RDWR|client.O_CREATE)
if(fd1003 < 0) {
  panic("open failed")
}


ret = client.Seek(fd969, 85, client.SEEK_SET)
if(ret != 85) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 85)
  panic("seek failed")
}


fd1004 := client.Open("/1QMCJHImk3.txt", client.O_RDWR|client.O_CREATE)
if(fd1004 < 0) {
  panic("open failed")
}


ret = client.Close(fd1002)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd980, 83)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1005 := client.Open("/B6ZnymC1_b.txt", client.O_RDWR|client.O_CREATE)
if(fd1005 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd990, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd1003, []byte("zNFAKa8G1PXUHVmN52RaRwKhHKpOktYJ2t2e6BSCTztRgO0h1E3lu1d8KA0FbXm5S3CPy0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (70) zNFAKa8G1PXUHVmN52RaRwKhHKpOktYJ2t2e6BSCTztRgO0h1E3lu1d8KA0FbXm5S3CPy0
//fd state: (70) zNFAKa8G1PXUHVmN52RaRwKhHKpOktYJ2t2e6BSCTztRgO0h1E3lu1d8KA0FbXm5S3CPy0

ret = client.Write(fd1003, []byte("x3U7ar5Es3E5gGqFiD3UetLOEFajvC"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (100) zNFAKa8G1PXUHVmN52RaRwKhHKpOktYJ2t2e6BSCTztRgO0h1E3lu1d8KA0FbXm5S3CPy0x3U7ar5Es3E5gGqFiD3UetLOEFajvC
//fd state: (0) 

ret = client.Write(fd968, []byte("WT3FGiHMaB8DCP_W0oaSXLnlO7GrrjBMrkpvQUtwVFlktII4kxy11"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (53) WT3FGiHMaB8DCP_W0oaSXLnlO7GrrjBMrkpvQUtwVFlktII4kxy11

ret = client.Seek(fd990, 97, client.SEEK_SET)
if(ret != 97) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 97)
  panic("seek failed")
}


buf, ret = client.Read(fd984, 17)
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


ret = client.Close(fd982)
if(ret != 0) {
  panic("close failed")
}


fd1006 := client.Open("/0jtYObXPre.txt", client.O_RDWR|client.O_CREATE)
if(fd1006 < 0) {
  panic("open failed")
}


ret = client.Seek(fd963, 23, client.SEEK_SET)
if(ret != 23) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 23)
  panic("seek failed")
}


fd1007 := client.Open("/k8VXHxxYvV.txt", client.O_RDWR|client.O_CREATE)
if(fd1007 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd984, 15)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd994, 67)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd1005, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd1007, 13, client.SEEK_SET)
if(ret != 13) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 13)
  panic("seek failed")
}


ret = client.Close(fd993)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd984, 10, client.SEEK_SET)
if(ret != 10) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 10)
  panic("seek failed")
}


ret = client.Close(fd963)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd1000)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd977)
if(ret != 0) {
  panic("close failed")
}


fd1008 := client.Open("/6tacWVf56J.txt", client.O_RDWR|client.O_CREATE)
if(fd1008 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd994, 98)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd998, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd985)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
fd1009 := client.Open("/LI4HznWdue.txt", client.O_RDWR|client.O_CREATE)
if(fd1009 < 0) {
  panic("open failed")
}


ret = client.Seek(fd997, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd908, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


buf, ret = client.Read(fd980, 79)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd986)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd978, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1010 := client.Open("/qLk_5cS8Ya.txt", client.O_RDWR|client.O_CREATE)
if(fd1010 < 0) {
  panic("open failed")
}

//fd state: (100) zNFAKa8G1PXUHVmN52RaRwKhHKpOktYJ2t2e6BSCTztRgO0h1E3lu1d8KA0FbXm5S3CPy0x3U7ar5Es3E5gGqFiD3UetLOEFajvC

ret = client.Write(fd1003, []byte("0pHFPjzMHt4x_Xgvk3qYmxe6rcfALrZ0mtEo6ooECjcBkOPZE19ZuT4uSD8BZwYa27jUWhdwXth6n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (177) zNFAKa8G1PXUHVmN52RaRwKhHKpOktYJ2t2e6BSCTztRgO0h1E3lu1d8KA0FbXm5S3CPy0x3U7ar5Es3E5gGqFiD3UetLOEFajvC0pHFPjzMHt4x_Xgvk3qYmxe6rcfALrZ0mtEo6ooECjcBkOPZE19ZuT4uSD8BZwYa27jUWhdwXth6n

ret = client.Seek(fd996, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1011 := client.Open("/nr0KUH2sAF.txt", client.O_RDWR|client.O_CREATE)
if(fd1011 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd997, 51)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "5Y") {
  panic("wrong data returned")
}


ret = client.Seek(fd1007, 19, client.SEEK_SET)
if(ret != 19) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 19)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd991, []byte("8T5VBxvDQPYuSKFClvMB7A7c3rnbP55wsch3tWtO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (40) 8T5VBxvDQPYuSKFClvMB7A7c3rnbP55wsch3tWtO

ret = client.Close(fd978)
if(ret != 0) {
  panic("close failed")
}


fd1012 := client.Open("/cs8WHp2mJ6.txt", client.O_RDWR|client.O_CREATE)
if(fd1012 < 0) {
  panic("open failed")
}


ret = client.Seek(fd984, 69, client.SEEK_SET)
if(ret != 69) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 69)
  panic("seek failed")
}


buf, ret = client.Read(fd996, 91)
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


buf, ret = client.Read(fd916, 41)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd994)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd980, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1003, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1013 := client.Open("/1FB_ZwzNZh.txt", client.O_RDWR|client.O_CREATE)
if(fd1013 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd916, 77)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd997, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}


ret = client.Seek(fd996, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (132) f2Nbs4T2rpE8CXnzcRf_Fs7j4eXfhks1Ot6_LLaCK9X3rGsMdajUx5GTeN9QrOXI2YX_YYh5uQT5UAT5mndoD_5BOLnyon7R0fAt1SvDUqUeMNQrUrxvB7gm9lBJCdDEfMo3xv7f

ret = client.Write(fd983, []byte("yDp"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (135) f2Nbs4T2rpE8CXnzcRf_Fs7j4eXfhks1Ot6_LLaCK9X3rGsMdajUx5GTeN9QrOXI2YX_YYh5uQT5UAT5mndoD_5BOLnyon7R0fAt1SvDUqUeMNQrUrxvB7gm9lBJCdDEfMo3yDpf
//fd state: (0) 

ret = client.Write(fd998, []byte("8t1b5wRm3rJQ8iLPGhFhSGj4vJBNOqFzD71_D4w3AAWqCwQ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (47) 8t1b5wRm3rJQ8iLPGhFhSGj4vJBNOqFzD71_D4w3AAWqCwQ
//fd state: (43) CzTACK7MDCRuyJSCzSSihxXKnBXfOkf1jQr5acHmHYP

ret = client.Write(fd980, []byte("8LkWkcLnjmCtE7sFlaiozL0kIQPQdIALVu"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (77) CzTACK7MDCRuyJSCzSSihxXKnBXfOkf1jQr5acHmHYP8LkWkcLnjmCtE7sFlaiozL0kIQPQdIALVu
//fd state: (177) zNFAKa8G1PXUHVmN52RaRwKhHKpOktYJ2t2e6BSCTztRgO0h1E3lu1d8KA0FbXm5S3CPy0x3U7ar5Es3E5gGqFiD3UetLOEFajvC0pHFPjzMHt4x_Xgvk3qYmxe6rcfALrZ0mtEo6ooECjcBkOPZE19ZuT4uSD8BZwYa27jUWhdwXth6n

ret = client.Write(fd1003, []byte("YHZ1bsixY2egXh3kd_2J6qErzLa44_NR6HtDQqAwuUDfKLWt1TyQbA6qa9PkkPd8wmWAKZXScH8CuecCJdQR"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (261) zNFAKa8G1PXUHVmN52RaRwKhHKpOktYJ2t2e6BSCTztRgO0h1E3lu1d8KA0FbXm5S3CPy0x3U7ar5Es3E5gGqFiD3UetLOEFajvC0pHFPjzMHt4x_Xgvk3qYmxe6rcfALrZ0mtEo6ooECjcBkOPZE19ZuT4uSD8BZwYa27jUWhdwXth6nYHZ1bsixY2egXh3kd_2J6qErzLa44_NR6HtDQqAwuUDfKLWt1TyQbA6qa9PkkPd8wmWAKZXScH8CuecCJdQR
//fd state: (0) 

ret = client.Write(fd1013, []byte("KLuHt9tazZ9lxP7qGAZ7nNqg8Xw9C8b6aa3SLiz1nVj2POr5iBKpAv_PFAqqIOq_78Sontm7n"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) KLuHt9tazZ9lxP7qGAZ7nNqg8Xw9C8b6aa3SLiz1nVj2POr5iBKpAv_PFAqqIOq_78Sontm7n

ret = client.Seek(fd990, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


ret = client.Seek(fd984, 63, client.SEEK_SET)
if(ret != 63) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 63)
  panic("seek failed")
}


fd1014 := client.Open("/pJyYoW8Q2N.txt", client.O_RDWR|client.O_CREATE)
if(fd1014 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1008, 31)
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


fd1015 := client.Open("/lygaNFnozr.txt", client.O_RDWR|client.O_CREATE)
if(fd1015 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd1005, []byte("AfVft0lk6CP9Vff_fW1sQNOIHHxaD1njlhu7pnB8NT0odqxdyqJjbOJVCWSBg4"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (62) AfVft0lk6CP9Vff_fW1sQNOIHHxaD1njlhu7pnB8NT0odqxdyqJjbOJVCWSBg4

fd1016 := client.Open("/RllLaNygFM.txt", client.O_RDWR|client.O_CREATE)
if(fd1016 < 0) {
  panic("open failed")
}


ret = client.Seek(fd961, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1017 := client.Open("/zL35j1zplT.txt", client.O_RDWR|client.O_CREATE)
if(fd1017 < 0) {
  panic("open failed")
}


ret = client.Seek(fd983, 17, client.SEEK_SET)
if(ret != 17) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 17)
  panic("seek failed")
}


fd1018 := client.Open("/AE89pFv5BA.txt", client.O_RDWR|client.O_CREATE)
if(fd1018 < 0) {
  panic("open failed")
}


fd1019 := client.Open("/UMnmYcjsNF.txt", client.O_RDWR|client.O_CREATE)
if(fd1019 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1015, 47)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1020 := client.Open("/EBLUvzsp0B.txt", client.O_RDWR|client.O_CREATE)
if(fd1020 < 0) {
  panic("open failed")
}


ret = client.Close(fd1006)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1012, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd979, 10)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vYe") {
  panic("wrong data returned")
}

//fd state: (49) o4NLujFJmcfM_gx9YigRiVk4r5dvrdPpYVmXUCZpEktcOrnv0V5swWqRJHxWK1YHGDZp3gXQTpWRtCyt_rOn86WJ8hydAkMS0mfwHZ4o50G4K0Ii58MuxJ

ret = client.Write(fd965, []byte("6pPkarHoHlCXlX9kesggDoIF6DlmedCPcelsn3JVIgWIj6INFnOen26JJ_e9"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (109) o4NLujFJmcfM_gx9YigRiVk4r5dvrdPpYVmXUCZpEktcOrnv06pPkarHoHlCXlX9kesggDoIF6DlmedCPcelsn3JVIgWIj6INFnOen26JJ_e90Ii58MuxJ

ret = client.Close(fd996)
if(ret != 0) {
  panic("close failed")
}


fd1021 := client.Open("/r0lIDh29XE.txt", client.O_RDWR|client.O_CREATE)
if(fd1021 < 0) {
  panic("open failed")
}


ret = client.Close(fd1008)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1004, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd984)
if(ret != 0) {
  panic("close failed")
}


fd1022 := client.Open("/_xUVgVqU6r.txt", client.O_RDWR|client.O_CREATE)
if(fd1022 < 0) {
  panic("open failed")
}


fd1023 := client.Open("/lVuo_WIGlh.txt", client.O_RDWR|client.O_CREATE)
if(fd1023 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1015, 86)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1024 := client.Open("/2BgRhfO5sI.txt", client.O_RDWR|client.O_CREATE)
if(fd1024 < 0) {
  panic("open failed")
}


fd1025 := client.Open("/_CP4mokCjO.txt", client.O_RDWR|client.O_CREATE)
if(fd1025 < 0) {
  panic("open failed")
}


ret = client.Close(fd1012)
if(ret != 0) {
  panic("close failed")
}

//fd state: (45) eCWUNku92jv3I5MQVhMh2cj6s4wnvedGbxyxOMjnmAvYe

ret = client.Write(fd979, []byte("dlrgNFtnwTqdMKqc8aEbKMbbXo63CJbGQ85"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (80) eCWUNku92jv3I5MQVhMh2cj6s4wnvedGbxyxOMjnmAvYedlrgNFtnwTqdMKqc8aEbKMbbXo63CJbGQ85

ret = client.Close(fd1020)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1005, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd980)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd992, 19)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "8UBL3iua857oa_Lj9yq") {
  panic("wrong data returned")
}


ret = client.Seek(fd1019, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd1024, 74)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (47) lxlQgnvZQcRGE_xFYAz3e5Jvv5XYk1VEqlHdiJChAJTZcTj

ret = client.Write(fd988, []byte("PnqnzOPbNSAp6GDIQe1ulxcPGr5C96CfdJbTU4s"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (86) lxlQgnvZQcRGE_xFYAz3e5Jvv5XYk1VEqlHdiJChAJTZcTjPnqnzOPbNSAp6GDIQe1ulxcPGr5C96CfdJbTU4s

ret = client.Seek(fd1022, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd990, 3, client.SEEK_SET)
if(ret != 3) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 3)
  panic("seek failed")
}


fd1026 := client.Open("/WDmAlke9N0.txt", client.O_RDWR|client.O_CREATE)
if(fd1026 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1015, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1027 := client.Open("/9MRCgnwRR5.txt", client.O_RDWR|client.O_CREATE)
if(fd1027 < 0) {
  panic("open failed")
}


ret = client.Seek(fd990, 66, client.SEEK_SET)
if(ret != 66) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 66)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd1026, []byte("N9P0G8PuNw9vpRzI2TVkKpX3ncRnvjyzKG850lb6mKtozKWidMJ328jxkRp3LK9CiuyQUzSquHf10ST6VY"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (82) N9P0G8PuNw9vpRzI2TVkKpX3ncRnvjyzKG850lb6mKtozKWidMJ328jxkRp3LK9CiuyQUzSquHf10ST6VY

fd1028 := client.Open("/atxR_KBzyK.txt", client.O_RDWR|client.O_CREATE)
if(fd1028 < 0) {
  panic("open failed")
}


ret = client.Close(fd983)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd987, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}

//fd state: (40) 8T5VBxvDQPYuSKFClvMB7A7c3rnbP55wsch3tWtO

ret = client.Write(fd991, []byte("D_ivrll6mAaY9mRQ7sQWQJrZ62toEnA0oRITB8M7bRRmKamtLPVNRYR9Ue7FgzbHiF7P6Y3xRRzIdC6G"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (120) 8T5VBxvDQPYuSKFClvMB7A7c3rnbP55wsch3tWtOD_ivrll6mAaY9mRQ7sQWQJrZ62toEnA0oRITB8M7bRRmKamtLPVNRYR9Ue7FgzbHiF7P6Y3xRRzIdC6G

buf, ret = client.Read(fd997, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd991, 11, client.SEEK_SET)
if(ret != 11) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 11)
  panic("seek failed")
}

//fd state: (0) 

ret = client.Write(fd1017, []byte("kIyIuNKmT2RIaPKoYUgBdXIFwrwSXDPHc5apULz6d0H8DC3W3MWrxThASqSnRDBqnrikmF9IS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (73) kIyIuNKmT2RIaPKoYUgBdXIFwrwSXDPHc5apULz6d0H8DC3W3MWrxThASqSnRDBqnrikmF9IS

fd1029 := client.Open("/fTKEpwYowY.txt", client.O_RDWR|client.O_CREATE)
if(fd1029 < 0) {
  panic("open failed")
}


fd1030 := client.Open("/tpL6efUjLY.txt", client.O_RDWR|client.O_CREATE)
if(fd1030 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd961, 17)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Close(fd1026)
if(ret != 0) {
  panic("close failed")
}


fd1031 := client.Open("/3pYSlSk_iq.txt", client.O_RDWR|client.O_CREATE)
if(fd1031 < 0) {
  panic("open failed")
}


fd1032 := client.Open("/KyuyMsniun.txt", client.O_RDWR|client.O_CREATE)
if(fd1032 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1013, 62, client.SEEK_SET)
if(ret != 62) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 62)
  panic("seek failed")
}


ret = client.Close(fd998)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd965, 49)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "0Ii58MuxJ") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1019, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd988, 48, client.SEEK_SET)
if(ret != 48) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 48)
  panic("seek failed")
}

//fd state: (0) 2NE0CD3zDSC3PjzPERIPOEkjgHBqK3QVNyX6nLDgTbT2lpf3ojsVSgLYyjG11feP0ejBvdFyWVU8goclwVv3rFI

ret = client.Write(fd1031, []byte("2ZDWSotFFNm"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (11) 2ZDWSotFFNm3PjzPERIPOEkjgHBqK3QVNyX6nLDgTbT2lpf3ojsVSgLYyjG11feP0ejBvdFyWVU8goclwVv3rFI

buf, ret = client.Read(fd1022, 61)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1033 := client.Open("/Xl1CCdrndP.txt", client.O_RDWR|client.O_CREATE)
if(fd1033 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1023, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd1013, 57, client.SEEK_SET)
if(ret != 57) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 57)
  panic("seek failed")
}

//fd state: (77) TN8kVvnxmZlrz9LZIx0qchpsVK2vGk8QWWrBSe6LxSInmj2iaxUNHHYie18UBL3iua857oa_Lj9yqJljlBUWI85p0zZFOmkKOImptSUrIPWMUEO0xrycDuInc_

ret = client.Write(fd992, []byte("GMghP7Z5Q8TgeR6oZOEZFW6_VUhpOqhTbbacDc4pxkDBlcG079CRNHh1OKT4jIYficmTzDWEPi9x_4qhGIAO7KKXIQTpuG0EAS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (175) TN8kVvnxmZlrz9LZIx0qchpsVK2vGk8QWWrBSe6LxSInmj2iaxUNHHYie18UBL3iua857oa_Lj9yqGMghP7Z5Q8TgeR6oZOEZFW6_VUhpOqhTbbacDc4pxkDBlcG079CRNHh1OKT4jIYficmTzDWEPi9x_4qhGIAO7KKXIQTpuG0EAS
now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
ret = client.Seek(fd1021, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd990, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "rJ5tUHec57PmG1yVXMRd5tSWqL5zvCKPK0MAAyoZPosjA9Pgn") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd1014, []byte("E3LuEiCLMf7u3YtmCv8almnLM9pk3HpBcHaYvTVFiSsFud"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (46) E3LuEiCLMf7u3YtmCv8almnLM9pk3HpBcHaYvTVFiSsFud

ret = client.Close(fd1033)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1032, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd1028, []byte("jbXe6nwxusYmLQjg3x6eAOUFrw5VmAUDManBwZMXivJfz9vt8WUT"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (52) jbXe6nwxusYmLQjg3x6eAOUFrw5VmAUDManBwZMXivJfz9vt8WUT

fd1034 := client.Open("/q4ub60oZmc.txt", client.O_RDWR|client.O_CREATE)
if(fd1034 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1011, 74, client.SEEK_SET)
if(ret != 74) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 74)
  panic("seek failed")
}


ret = client.Seek(fd1013, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1035 := client.Open("/yEg4zmRGmr.txt", client.O_RDWR|client.O_CREATE)
if(fd1035 < 0) {
  panic("open failed")
}

//fd state: (118) o4NLujFJmcfM_gx9YigRiVk4r5dvrdPpYVmXUCZpEktcOrnv06pPkarHoHlCXlX9kesggDoIF6DlmedCPcelsn3JVIgWIj6INFnOen26JJ_e90Ii58MuxJ

ret = client.Write(fd965, []byte("dMo0Nky_CYQvD8ia4O9Nw8aPM8tiBWya"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (150) o4NLujFJmcfM_gx9YigRiVk4r5dvrdPpYVmXUCZpEktcOrnv06pPkarHoHlCXlX9kesggDoIF6DlmedCPcelsn3JVIgWIj6INFnOen26JJ_e90Ii58MuxJdMo0Nky_CYQvD8ia4O9Nw8aPM8tiBWya

buf, ret = client.Read(fd991, 75)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "uSKFClvMB7A7c3rnbP55wsch3tWtOD_ivrll6mAaY9mRQ7sQWQJrZ62toEnA0oRITB8M7bRRmKa") {
  panic("wrong data returned")
}


ret = client.Close(fd995)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1004, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1036 := client.Open("/VyvFnecLmb.txt", client.O_RDWR|client.O_CREATE)
if(fd1036 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1021, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1037 := client.Open("/zG9ey33jcs.txt", client.O_RDWR|client.O_CREATE)
if(fd1037 < 0) {
  panic("open failed")
}


fd1038 := client.Open("/bfcWynsjAY.txt", client.O_RDWR|client.O_CREATE)
if(fd1038 < 0) {
  panic("open failed")
}


fd1039 := client.Open("/WZrIgODs1e.txt", client.O_RDWR|client.O_CREATE)
if(fd1039 < 0) {
  panic("open failed")
}


ret = client.Close(fd1015)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd1004)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd1022)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1023, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd1035)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd988)
if(ret != 0) {
  panic("close failed")
}


fd1040 := client.Open("/gmKKZnvKrX.txt", client.O_RDWR|client.O_CREATE)
if(fd1040 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1021, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd965, 63)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


ret = client.Seek(fd1029, 15, client.SEEK_SET)
if(ret != 15) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 15)
  panic("seek failed")
}


fd1041 := client.Open("/tswNiHc8WS.txt", client.O_RDWR|client.O_CREATE)
if(fd1041 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1039, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


buf, ret = client.Read(fd1019, 3)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1042 := client.Open("/Fx14wmZV6Z.txt", client.O_RDWR|client.O_CREATE)
if(fd1042 < 0) {
  panic("open failed")
}


ret = client.Seek(fd1032, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd1031, 41, client.SEEK_SET)
if(ret != 41) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 41)
  panic("seek failed")
}


buf, ret = client.Read(fd1025, 42)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1029, 25)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "ZIx0qchpsVK2vGk8QWWrBSe6L") {
  panic("wrong data returned")
}


ret = client.Seek(fd1013, 37, client.SEEK_SET)
if(ret != 37) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 37)
  panic("seek failed")
}

//fd state: (73) kIyIuNKmT2RIaPKoYUgBdXIFwrwSXDPHc5apULz6d0H8DC3W3MWrxThASqSnRDBqnrikmF9IS

ret = client.Write(fd1017, []byte("7TOSZBVshgUPWLaFPNwI56tF_ZiT4_Nhi67c7JWrege3Su031N0d4W7QaXy5QAyn9CjvdjC98RAv278QueF5OB"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (159) kIyIuNKmT2RIaPKoYUgBdXIFwrwSXDPHc5apULz6d0H8DC3W3MWrxThASqSnRDBqnrikmF9IS7TOSZBVshgUPWLaFPNwI56tF_ZiT4_Nhi67c7JWrege3Su031N0d4W7QaXy5QAyn9CjvdjC98RAv278QueF5OB
//fd state: (159) kIyIuNKmT2RIaPKoYUgBdXIFwrwSXDPHc5apULz6d0H8DC3W3MWrxThASqSnRDBqnrikmF9IS7TOSZBVshgUPWLaFPNwI56tF_ZiT4_Nhi67c7JWrege3Su031N0d4W7QaXy5QAyn9CjvdjC98RAv278QueF5OB

ret = client.Write(fd1017, []byte("DrUga"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (164) kIyIuNKmT2RIaPKoYUgBdXIFwrwSXDPHc5apULz6d0H8DC3W3MWrxThASqSnRDBqnrikmF9IS7TOSZBVshgUPWLaFPNwI56tF_ZiT4_Nhi67c7JWrege3Su031N0d4W7QaXy5QAyn9CjvdjC98RAv278QueF5OBDrUga

ret = client.Close(fd979)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd961, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1043 := client.Open("/zgdT4mG5bm.txt", client.O_RDWR|client.O_CREATE)
if(fd1043 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1040, 4)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1024, 73)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd1038, []byte("Sl6mBLWKAgg5YSk6Jvigj5IKbzp9sB9FUb4bIyaF2dRACboBAHX0uOEVjAA4e3JDzDSxwTY9l6KRjw8wXe6V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (84) Sl6mBLWKAgg5YSk6Jvigj5IKbzp9sB9FUb4bIyaF2dRACboBAHX0uOEVjAA4e3JDzDSxwTY9l6KRjw8wXe6V

ret = client.Close(fd1031)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1018, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Close(fd1013)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd1025)
if(ret != 0) {
  panic("close failed")
}


fd1044 := client.Open("/1hYlxqnLZN.txt", client.O_RDWR|client.O_CREATE)
if(fd1044 < 0) {
  panic("open failed")
}

//fd state: (86) 8T5VBxvDQPYuSKFClvMB7A7c3rnbP55wsch3tWtOD_ivrll6mAaY9mRQ7sQWQJrZ62toEnA0oRITB8M7bRRmKamtLPVNRYR9Ue7FgzbHiF7P6Y3xRRzIdC6G

ret = client.Write(fd991, []byte("ChT9wGalUv_TcF0OB9TeVZzdIVod2XY4XKe3j5HBfXsHx9Q6LeGKrArfg2Y3FZYpbLzGf3eBUECWQUHB4ouyxoUz7lM"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (177) 8T5VBxvDQPYuSKFClvMB7A7c3rnbP55wsch3tWtOD_ivrll6mAaY9mRQ7sQWQJrZ62toEnA0oRITB8M7bRRmKaChT9wGalUv_TcF0OB9TeVZzdIVod2XY4XKe3j5HBfXsHx9Q6LeGKrArfg2Y3FZYpbLzGf3eBUECWQUHB4ouyxoUz7lM
//fd state: (177) 8T5VBxvDQPYuSKFClvMB7A7c3rnbP55wsch3tWtOD_ivrll6mAaY9mRQ7sQWQJrZ62toEnA0oRITB8M7bRRmKaChT9wGalUv_TcF0OB9TeVZzdIVod2XY4XKe3j5HBfXsHx9Q6LeGKrArfg2Y3FZYpbLzGf3eBUECWQUHB4ouyxoUz7lM

ret = client.Write(fd991, []byte("3A6qTdiyI1m590P09G9LeYP2PCWKHuGwo7GoavbU5Ff_4rd50hjqBIWUJpuFTX"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (239) 8T5VBxvDQPYuSKFClvMB7A7c3rnbP55wsch3tWtOD_ivrll6mAaY9mRQ7sQWQJrZ62toEnA0oRITB8M7bRRmKaChT9wGalUv_TcF0OB9TeVZzdIVod2XY4XKe3j5HBfXsHx9Q6LeGKrArfg2Y3FZYpbLzGf3eBUECWQUHB4ouyxoUz7lM3A6qTdiyI1m590P09G9LeYP2PCWKHuGwo7GoavbU5Ff_4rd50hjqBIWUJpuFTX

buf, ret = client.Read(fd1021, 72)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) RiwgoyosGn7WBKesjkHS1SE6xadBV1dGrzczXPmtkcQjOdzfQMy73m0duCefroN8KQFWxX85lvaH_Uc3ItOEAU9idlTnTvIn4iawj

ret = client.Write(fd1027, []byte("WQY7FLZmWdvsuapXFZMMmvroW20AMbq68L09gnXOmogLu2ns7zpkzC4hQ0sVy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (61) WQY7FLZmWdvsuapXFZMMmvroW20AMbq68L09gnXOmogLu2ns7zpkzC4hQ0sVyoN8KQFWxX85lvaH_Uc3ItOEAU9idlTnTvIn4iawj

buf, ret = client.Read(fd997, 81)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (164) kIyIuNKmT2RIaPKoYUgBdXIFwrwSXDPHc5apULz6d0H8DC3W3MWrxThASqSnRDBqnrikmF9IS7TOSZBVshgUPWLaFPNwI56tF_ZiT4_Nhi67c7JWrege3Su031N0d4W7QaXy5QAyn9CjvdjC98RAv278QueF5OBDrUga

ret = client.Write(fd1017, []byte("BFLQfyUHA9L4txBnAz5gsdDHaDLIQX_Um7Cz5eLMNJMD8soDnJNqAbcVajAXh5THBcf5rGoqITwige8OxG6y0tEZPy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (254) kIyIuNKmT2RIaPKoYUgBdXIFwrwSXDPHc5apULz6d0H8DC3W3MWrxThASqSnRDBqnrikmF9IS7TOSZBVshgUPWLaFPNwI56tF_ZiT4_Nhi67c7JWrege3Su031N0d4W7QaXy5QAyn9CjvdjC98RAv278QueF5OBDrUgaBFLQfyUHA9L4txBnAz5gsdDHaDLIQX_Um7Cz5eLMNJMD8soDnJNqAbcVajAXh5THBcf5rGoqITwige8OxG6y0tEZPy

buf, ret = client.Read(fd1040, 34)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (19) DLpM_cdkC7GGJdAH3s3GqVDnlYE7UMSpY

ret = client.Write(fd1007, []byte("L"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) DLpM_cdkC7GGJdAH3s3LqVDnlYE7UMSpY
//fd state: (0) DUtt6FRvZUO8nZ7vqcqOJf_k869A8Vtkt4kiFGWOKifssUqTadhYsTTqj698cCTGXVZ2Rl2ZoAC_cxzC5I5kIArFnuzDLUkDNs41abtjUye9jEvskWpxQ9NnVSOUyKWMms4e

ret = client.Write(fd1034, []byte("_4rt30VzYNqdiPbeYiDRGZGer2DIWmRWOSyGL2ZQKdL5hloXgpLmMIqpTVROcMr4yPIgroqzmcUvNyV"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (79) _4rt30VzYNqdiPbeYiDRGZGer2DIWmRWOSyGL2ZQKdL5hloXgpLmMIqpTVROcMr4yPIgroqzmcUvNyVC5I5kIArFnuzDLUkDNs41abtjUye9jEvskWpxQ9NnVSOUyKWMms4e

ret = client.Close(fd997)
if(ret != 0) {
  panic("close failed")
}


ret = client.Seek(fd1021, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


fd1045 := client.Open("/bscOPxCqaM.txt", client.O_RDWR|client.O_CREATE)
if(fd1045 < 0) {
  panic("open failed")
}


ret = client.Close(fd1039)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd1024, []byte("k7bOlEEPVuzNfn7OCDEzAhh8swFoW9nuxXUZDrGDb216xjivaQhN12V"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (55) k7bOlEEPVuzNfn7OCDEzAhh8swFoW9nuxXUZDrGDb216xjivaQhN12V

buf, ret = client.Read(fd1007, 57)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "qVDnlYE7UMSpY") {
  panic("wrong data returned")
}


ret = client.Seek(fd1019, 0, client.SEEK_SET)
if(ret != 0) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 0)
  panic("seek failed")
}


ret = client.Seek(fd1028, 46, client.SEEK_SET)
if(ret != 46) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 46)
  panic("seek failed")
}


buf, ret = client.Read(fd1042, 33)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "43ORMZCWPz2XjGkTjmrHN59C3nH6L95mB") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1027, 78)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "oN8KQFWxX85lvaH_Uc3ItOEAU9idlTnTvIn4iawj") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd1041, []byte("cF3KJ6we9vprLCXbFPtiCzNhP5rBi9wgCG4Y4L"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (38) cF3KJ6we9vprLCXbFPtiCzNhP5rBi9wgCG4Y4L

fd1046 := client.Open("/6yBMQiluJ2.txt", client.O_RDWR|client.O_CREATE)
if(fd1046 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1028, 52)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "vt8WUT") {
  panic("wrong data returned")
}


ret = client.Close(fd1030)
if(ret != 0) {
  panic("close failed")
}


ret = client.Close(fd968)
if(ret != 0) {
  panic("close failed")
}

//fd state: (0) 

ret = client.Write(fd1009, []byte("Kl8eSCIOn1VMsNLQe4V3go97A5IockdFVBeedvtSIOnGMM3sdqNmQaXVAHdHDzvEVzGeI"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (69) Kl8eSCIOn1VMsNLQe4V3go97A5IockdFVBeedvtSIOnGMM3sdqNmQaXVAHdHDzvEVzGeI

fd1047 := client.Open("/PCTC9S86Zb.txt", client.O_RDWR|client.O_CREATE)
if(fd1047 < 0) {
  panic("open failed")
}

//fd state: (0) 

ret = client.Write(fd1032, []byte("nqWZ1mhpou_v4FFdkBi8cmGte3LWeeudvypVk6UXkgF4rDr6LdKop3Z0iE13YNG2KtCiBuBLFYcU1vS97sCiHSmxKQHgYRREZ"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (97) nqWZ1mhpou_v4FFdkBi8cmGte3LWeeudvypVk6UXkgF4rDr6LdKop3Z0iE13YNG2KtCiBuBLFYcU1vS97sCiHSmxKQHgYRREZ

ret = client.Seek(fd1005, 2, client.SEEK_SET)
if(ret != 2) {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, 2)
  panic("seek failed")
}

//fd state: (38) cF3KJ6we9vprLCXbFPtiCzNhP5rBi9wgCG4Y4L

ret = client.Write(fd1041, []byte("U937rIJTucqB_2Xlxt6DlE8goGQqH1l0p0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (72) cF3KJ6we9vprLCXbFPtiCzNhP5rBi9wgCG4Y4LU937rIJTucqB_2Xlxt6DlE8goGQqH1l0p0

ret = client.Close(fd1014)
if(ret != 0) {
  panic("close failed")
}

//fd state: (79) _4rt30VzYNqdiPbeYiDRGZGer2DIWmRWOSyGL2ZQKdL5hloXgpLmMIqpTVROcMr4yPIgroqzmcUvNyVC5I5kIArFnuzDLUkDNs41abtjUye9jEvskWpxQ9NnVSOUyKWMms4e

ret = client.Write(fd1034, []byte("LdVn02kRerQKvnBJWWt4Y1S0i_seqThlko2DDD2llCn95JlRd0n0"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (131) _4rt30VzYNqdiPbeYiDRGZGer2DIWmRWOSyGL2ZQKdL5hloXgpLmMIqpTVROcMr4yPIgroqzmcUvNyVLdVn02kRerQKvnBJWWt4Y1S0i_seqThlko2DDD2llCn95JlRd0n0e
//fd state: (0) 

ret = client.Write(fd987, []byte("3zF4LaM1j5M8g7yizaAS"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (20) 3zF4LaM1j5M8g7yizaAS

buf, ret = client.Read(fd1047, 94)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (0) 

ret = client.Write(fd1010, []byte("CuyRg584hjs9ZDZ4s1"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (18) CuyRg584hjs9ZDZ4s1

ret = client.Close(fd1046)
if(ret != 0) {
  panic("close failed")
}


buf, ret = client.Read(fd1009, 37)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (115) DFK8ImPtDhWky7YM26ji4F56nBLrd77riK7ewgsRFYvPi4cJkeHFuWhEWwwGfBx6ibrJ5tUHec57PmG1yVXMRd5tSWqL5zvCKPK0MAAyoZPosjA9Pgn

ret = client.Write(fd990, []byte("Rf5Bt9SiguAZxMkvcVZTcIb_EEza217Ps8s21KjqoT_sx0CE9JVpZDDLcvYJ0BvVwUj55LLZO5Zl"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (191) DFK8ImPtDhWky7YM26ji4F56nBLrd77riK7ewgsRFYvPi4cJkeHFuWhEWwwGfBx6ibrJ5tUHec57PmG1yVXMRd5tSWqL5zvCKPK0MAAyoZPosjA9PgnRf5Bt9SiguAZxMkvcVZTcIb_EEza217Ps8s21KjqoT_sx0CE9JVpZDDLcvYJ0BvVwUj55LLZO5Zl

buf, ret = client.Read(fd1036, 80)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


buf, ret = client.Read(fd1047, 16)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}


fd1048 := client.Open("/3Xuls7KVgK.txt", client.O_RDWR|client.O_CREATE)
if(fd1048 < 0) {
  panic("open failed")
}


ret = client.Close(fd1047)
if(ret != 0) {
  panic("close failed")
}

//fd state: (131) _4rt30VzYNqdiPbeYiDRGZGer2DIWmRWOSyGL2ZQKdL5hloXgpLmMIqpTVROcMr4yPIgroqzmcUvNyVLdVn02kRerQKvnBJWWt4Y1S0i_seqThlko2DDD2llCn95JlRd0n0e

ret = client.Write(fd1034, []byte("km_s2ln1ZXsbPpSV69OK9VdOJOPy"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (159) _4rt30VzYNqdiPbeYiDRGZGer2DIWmRWOSyGL2ZQKdL5hloXgpLmMIqpTVROcMr4yPIgroqzmcUvNyVLdVn02kRerQKvnBJWWt4Y1S0i_seqThlko2DDD2llCn95JlRd0n0km_s2ln1ZXsbPpSV69OK9VdOJOPy

ret = client.Close(fd1041)
if(ret != 0) {
  panic("close failed")
}


fd1049 := client.Open("/atxR_KBzyK.txt", client.O_RDWR|client.O_CREATE)
if(fd1049 < 0) {
  panic("open failed")
}


buf, ret = client.Read(fd1028, 70)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (324) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PGNaDirOIX0L8lkbJIIY8T9uy0DG4Jq7160YxqjfTs0TvcUcOiv4zYpsWx12Af6iB8sVeVkPFf42THlsw9R8qetZVcmEmnBHzUm_gWvTV9jgTHwh1UpuckNc2m7MLQVgISBibVhpqi6JFBQk0WDi0TMjcqFJ2pogTeWYA7D4xJBtCwlUD1XidlwQhZoqHq6gpQ7vLdo26nQYvSe_s5yaI5xps33zhk2qzh0H5ZbePdkIh0yw1M4EHH1y5_e9d74xDE9F0qcwoUGGjwUCN16jmP

ret = client.Write(fd916, []byte("VqAQLOWPP1kMjBhm9S7L0kkXRX2grI9MwHTe9oppiOE2LO59DV9u_4eK9lrO"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (384) LS8LRt3Yw0Z2wltkBiK9UvWEq_a7oKzSyGNbPDjDO6zbg8PGNaDirOIX0L8lkbJIIY8T9uy0DG4Jq7160YxqjfTs0TvcUcOiv4zYpsWx12Af6iB8sVeVkPFf42THlsw9R8qetZVcmEmnBHzUm_gWvTV9jgTHwh1UpuckNc2m7MLQVgISBibVhpqi6JFBQk0WDi0TMjcqFJ2pogTeWYA7D4xJBtCwlUD1XidlwQhZoqHq6gpQ7vLdo26nQYvSe_s5yaI5xps33zhk2qzh0H5ZbePdkIh0yw1M4EHH1y5_e9d74xDE9F0qcwoUGGjwUCN16jmPVqAQLOWPP1kMjBhm9S7L0kkXRX2grI9MwHTe9oppiOE2LO59DV9u_4eK9lrO

buf, ret = client.Read(fd1028, 6)
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "") {
  panic("wrong data returned")
}

//fd state: (191) DFK8ImPtDhWky7YM26ji4F56nBLrd77riK7ewgsRFYvPi4cJkeHFuWhEWwwGfBx6ibrJ5tUHec57PmG1yVXMRd5tSWqL5zvCKPK0MAAyoZPosjA9PgnRf5Bt9SiguAZxMkvcVZTcIb_EEza217Ps8s21KjqoT_sx0CE9JVpZDDLcvYJ0BvVwUj55LLZO5Zl

ret = client.Write(fd990, []byte("bJAxTY0hJ9GsvL5sBlyLdEm12RBurqNStUnjxBFLtKRYB5QbPcERADxlPP7XuCK4KAH"))
if(ret != 0) {
  panic("write failed")
}

//fd state: (258) DFK8ImPtDhWky7YM26ji4F56nBLrd77riK7ewgsRFYvPi4cJkeHFuWhEWwwGfBx6ibrJ5tUHec57PmG1yVXMRd5tSWqL5zvCKPK0MAAyoZPosjA9PgnRf5Bt9SiguAZxMkvcVZTcIb_EEza217Ps8s21KjqoT_sx0CE9JVpZDDLcvYJ0BvVwUj55LLZO5ZlbJAxTY0hJ9GsvL5sBlyLdEm12RBurqNStUnjxBFLtKRYB5QbPcERADxlPP7XuCK4KAH

ret = client.Close(fd1009)
if(ret != 0) {
  panic("close failed")
}

now = time.Nanoseconds()
fmt.Printf("Elapsed = %d\nn", now - last)
last = now
	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}
