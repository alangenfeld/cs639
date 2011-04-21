package proxy  

import (
	"./master"
	"fmt"
	"../include/sfs"
	"os"
	"rand"
	"time"
	//"flag"
	//"strconv"
)

var  m = new(master.Master);

var delay int64;
type Proxy int;
var delay_prob int64;

func ( p *Proxy) DeleteFile(args *sfs.DeleteArgs, ret *sfs.DeleteReturn) os.Error {
	//time.Sleep(delay)
	Delay()
	fmt.Println("Inside proxy DeleteFile")
	return m.DeleteFile(args,ret)
}

func( p *Proxy) ReadOpen(args *sfs.OpenArgs, info *sfs.OpenReturn) os.Error {
        //time.Sleep(delay)
	Delay()
	fmt.Println("Inside proxy ReadOpen")     
	return m.ReadOpen(args,info)
}


func(p *Proxy) ReadDir(args *sfs.ReadDirArgs, ret *sfs.ReadDirReturn) os.Error {
	//time.Sleep(delay)
	Delay()
	fmt.Println("Inside proxy ReadDir")
	return m.ReadDir(args,ret)
}


func(p *Proxy) MapChunkToFile(args *sfs.MapChunkToFileArgs, ret *sfs.MapChunkToFileReturn) os.Error {
        //time.Sleep(delay)   
	Delay()
	fmt.Println("Inside proxy MapChunkToFile")
	return m.MapChunkToFile(args,ret)
}

func (p *Proxy) GetNewChunk(args *sfs.GetNewChunkArgs, ret *sfs.GetNewChunkReturn) os.Error {
        //time.Sleep(delay)
        Delay()
	fmt.Println("Inside proxy GetNewChunk")
	return m.GetNewChunk(args,ret) 
}

func (p *Proxy) ReportWrite(args *sfs.ReportWriteArgs, ret *sfs.ReportWriteReturn) os.Error {
	//time.Sleep(delay)
        Delay()
	fmt.Println("Inside proxy ReportWrite")
	return m.ReportWrite(args,ret)
}

func (p *Proxy) RemoveFile(args *sfs.RemoveArgs, result *sfs.RemoveReturn) os.Error {
	//time.Sleep(delay)
        Delay()
	fmt.Println("Inside proxy RemoveFile")
	return m.RemoveFile(args,result)
}
func (p *Proxy) BirthChunk(args *sfs.ChunkBirthArgs, info *sfs.ChunkBirthReturn) os.Error {
	Delay()
        fmt.Println("Inside proxy BirthChunk")
	return m.BirthChunk(args,info)
}

func (p *Proxy) BeatHeart(args *sfs.HeartbeatArgs, info *sfs.HeartbeatReturn) os.Error {
	Delay()
        fmt.Println("Inside proxy BeatHeart")
	return m.BeatHeart(args,info)
}
func (p *Proxy) MakeDir(args *sfs.MakeDirArgs, ret *sfs.MakeDirReturn) os.Error {
	//time.Sleep(delay)
	Delay()
	fmt.Println("Inside makedir")
	return m.MakeDir(args,ret)
}


func Init(del int64,del_prob int64){
	fmt.Println("Inside proxy initializer")
	m = new(master.Master);	
	delay = del;
	delay_prob = del_prob;
	//if err != nil{
	//	fmt.Println(err)
	//}

}

func Delay(){
	m := rand.Int63n(100)
	if m < delay_prob {
		fmt.Println("Delaying ...")
		time.Sleep(delay);
	}
}


