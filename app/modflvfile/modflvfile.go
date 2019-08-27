package main

import (
	"flag"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/nezha/pkg/errors"
	"github.com/q191201771/nezha/pkg/log"
	"io"
	"os"
)

// 修改flv文件的一些信息（比如某些tag的时间戳）后另存文件
//
// Usage:
// ./bin/modflvfile -i /tmp/in.flv -o /tmp/out.flv

var countA int
var countV int

func hookTag(tag *httpflv.Tag) {
	log.Infof("%+v", tag.Header)
	if tag.Header.T == httpflv.TagTypeAudio {

		//if countA < 3 {
		//	httpflv.ModTagTimestamp(tag, 16777205)
		//}
		//countA++
	}
	if tag.Header.T == httpflv.TagTypeVideo {
		//if countV < 3 {
		//	httpflv.ModTagTimestamp(tag, 16777205)
		//}
		//countV++
		if tag.IsAVCKeySeqHeader() {
			log.Info("key seq header.")
		}
		if tag.IsAVCKeyNalu() {
			log.Info("key nalu.")
		}
	}
}

func main() {
	var err error
	inFileName, outFileName := parseFlag()

	var ffr httpflv.FlvFileReader
	err = ffr.Open(inFileName)
	errors.PanicIfErrorOccur(err)
	defer ffr.Dispose()
	log.Infof("open input flv file succ.")

	var ffw httpflv.FlvFileWriter
	err = ffw.Open(outFileName)
	errors.PanicIfErrorOccur(err)
	defer ffw.Dispose()
	log.Infof("open output flv file succ.")

	flvHeader, err := ffr.ReadFlvHeader()
	errors.PanicIfErrorOccur(err)

	err = ffw.WriteRaw(flvHeader)
	errors.PanicIfErrorOccur(err)

	//for i:=0; i < 10; i++{
	for {
		tag, err := ffr.ReadTag()
		if err == io.EOF {
			log.Infof("EOF.")
			break
		}
		errors.PanicIfErrorOccur(err)

		//log.Infof("> hook. %+v", tag)
		hookTag(tag)
		//log.Infof("< hook. %+v", tag)
		err = ffw.WriteRaw(tag.Raw)
		errors.PanicIfErrorOccur(err)
	}
}

func parseFlag() (string, string) {
	i := flag.String("i", "", "specify input flv file")
	o := flag.String("o", "", "specify ouput flv file")
	flag.Parse()
	if *i == "" || *o == "" {
		flag.Usage()
		os.Exit(1)
	}
	return *i, *o
}
