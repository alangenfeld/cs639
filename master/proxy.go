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


func ( p *Proxy) DeleteFile(args *sfs.DeleteArgs, ret *sfs.DeleteReturn) os.Error {
	time.Sleep(delay)
	fmt.Println("Inside proxy DeleteFile")
	return m.DeleteFile(args,ret)
}

func( p *Proxy) ReadOpen(args *sfs.OpenArgs, info *sfs.OpenReturn) os.Error {
        time.Sleep(delay)
	fmt.Println("Inside proxy ReadOpen")     
	return m.ReadOpen(args,info)
}


func(p *Proxy) ReadDir(args *sfs.ReadDirArgs, ret *sfs.ReadDirReturn) os.Error {
	time.Sleep(delay)
	fmt.Println("Inside proxy ReadDir")
	return m.ReadDir(args,ret)
}


func(p *Proxy) MapChunkToFile(args *sfs.MapChunkToFileArgs, ret *sfs.MapChunkToFileReturn) os.Error {
        time.Sleep(delay)   
	fmt.Println("Inside proxy MapChunkToFile")
	return m.MapChunkToFile(args,ret)
}

func (p *Proxy) GetNewChunk(args *sfs.GetNewChunkArgs, ret *sfs.GetNewChunkReturn) os.Error {
        time.Sleep(delay)
        fmt.Println("Inside proxy GetNewChunk")
	return m.GetNewChunk(args,ret) 
}

func (p *Proxy) ReportWrite(args *sfs.ReportWriteArgs, ret *sfs.ReportWriteReturn) os.Error {
	time.Sleep(delay)
        fmt.Println("Inside proxy ReportWrite")
	return m.ReportWrite(args,ret)
}

func (p *Proxy) RemoveFile(args *sfs.RemoveArgs, result *sfs.RemoveReturn) os.Error {
	time.Sleep(delay)
        fmt.Println("Inside proxy RemoveFile")
	return m.RemoveFile(args,result)
}
func (p *Proxy) BirthChunk(args *sfs.ChunkBirthArgs, info *sfs.ChunkBirthReturn) os.Error {
	        time.Sleep(delay)
        fmt.Println("Inside proxy BirthChunk")
	return m.BirthChunk(args,info)
}

func (p *Proxy) BeatHeart(args *sfs.HeartbeatArgs, info *sfs.HeartbeatReturn) os.Error {
	if shouldDrop() == true {
	time.Sleep(delay)
	}
        fmt.Println("Inside proxy BeatHeart")
	return m.BeatHeart(args,info)
}
func (p *Proxy) MakeDir(args *sfs.MakeDirArgs, ret *sfs.MakeDirReturn) os.Error {
	time.Sleep(delay)
	fmt.Println("Inside makedir")
	return m.MakeDir(args,ret)
}


func init(){
	fmt.Println("Inside proxy initializer")
	m = new(master.Master);
	//flag.Parse();
	
	delay =  00000000;
	//if err != nil{
	//	fmt.Println(err)
	//}

}

func shouldDrop() bool {
	m := rand.Int31n(3)
	if m > 1 {
		return true;
	}
	return false;
}


