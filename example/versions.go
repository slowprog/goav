package main

import (
	"log"

	"github.com/slowprog/goav/avcodec"
	"github.com/slowprog/goav/avdevice"
	"github.com/slowprog/goav/avfilter"
	"github.com/slowprog/goav/avformat"
	"github.com/slowprog/goav/avutil"
	"github.com/slowprog/goav/swresample"
	"github.com/slowprog/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
