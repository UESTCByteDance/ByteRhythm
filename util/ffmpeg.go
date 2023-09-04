package util

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strings"
	"unsafe"

	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"github.com/giorgisio/goav/swscale"
)

// VideoGetNetImgCount 网络拉流抽帧 图片存储本地
func VideoGetNetImgCount(frameNum int, url string) string {
	//视频流地址
	// url="rtmp://192.168.20.221:30200/live/2dfp52anvad7g"
	//初始化网络库
	filename := url
	avformat.AvformatNetworkInit()
	// Open video file
	pFormatContext := avformat.AvformatAllocContext()

	if avformat.AvformatOpenInput(&pFormatContext, filename, nil, nil) != 0 {
		fmt.Printf("Unable to open file %s\n", filename)
		return ""
	}

	// Retrieve stream information
	if pFormatContext.AvformatFindStreamInfo(nil) < 0 {
		fmt.Println("Couldn't find stream information")
		return ""
	}

	// Dump information about file onto standard error
	//pFormatContext.AvDumpFormat(0, filename, 0)

	// Find the first video stream
	for i := 0; i < int(pFormatContext.NbStreams()); i++ {
		switch pFormatContext.Streams()[i].CodecParameters().AvCodecGetType() {
		case avformat.AVMEDIA_TYPE_VIDEO:

			// Get a pointer to the codec context for the video stream
			pCodecCtxOrig := pFormatContext.Streams()[i].Codec()
			// Find the decoder for the video stream
			pCodec := avcodec.AvcodecFindDecoder(avcodec.CodecId(pCodecCtxOrig.GetCodecId()))
			if pCodec == nil {
				fmt.Println("Unsupported codec!")
				return ""
			}
			// Copy context
			pCodecCtx := pCodec.AvcodecAllocContext3()
			if pCodecCtx.AvcodecCopyContext((*avcodec.Context)(unsafe.Pointer(pCodecCtxOrig))) != 0 {
				fmt.Println("Couldn't copy codec context")
				return ""
			}

			// Open codec
			if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
				fmt.Println("Could not open codec")
				return ""
			}

			// Allocate video frame
			pFrame := avutil.AvFrameAlloc()

			// Allocate an AVFrame structure
			pFrameRGB := avutil.AvFrameAlloc()
			if pFrameRGB == nil {
				fmt.Println("Unable to allocate RGB Frame")
				return ""
			}

			// Determine required buffer size and allocate buffer
			numBytes := uintptr(avcodec.AvpictureGetSize(avcodec.AV_PIX_FMT_RGB24, pCodecCtx.Width(),
				pCodecCtx.Height()))
			buffer := avutil.AvMalloc(numBytes)

			// Assign appropriate parts of buffer to image planes in pFrameRGB
			// Note that pFrameRGB is an AVFrame, but AVFrame is a superset
			// of AVPicture
			avp := (*avcodec.Picture)(unsafe.Pointer(pFrameRGB))
			avp.AvpictureFill((*uint8)(buffer), avcodec.AV_PIX_FMT_RGB24,
				pCodecCtx.Width(), pCodecCtx.Height())

			// initialize SWS context for software scaling
			swsCtx := swscale.SwsGetcontext(
				pCodecCtx.Width(),
				pCodecCtx.Height(),
				(swscale.PixelFormat)(pCodecCtx.PixFmt()),
				pCodecCtx.Width(),
				pCodecCtx.Height(),
				avcodec.AV_PIX_FMT_RGB24,
				avcodec.SWS_BILINEAR,
				nil,
				nil,
				nil,
			)

			// Read frames and save first five frames to disk
			frameNumber := 1
			packet := avcodec.AvPacketAlloc()
			for pFormatContext.AvReadFrame(packet) >= 0 {
				// Is this a packet from the video stream?
				if packet.StreamIndex() == i {
					// Decode video frame
					response := pCodecCtx.AvcodecSendPacket(packet)
					if response < 0 {
						fmt.Printf("Error while sending a packet to the decoder: %s\n", avutil.ErrorFromCode(response))
					}
					for response >= 0 {
						response = pCodecCtx.AvcodecReceiveFrame((*avcodec.Frame)(unsafe.Pointer(pFrame)))
						if response == avutil.AvErrorEAGAIN || response == avutil.AvErrorEOF {
							break
						} else if response < 0 {
							if response == -11 {
								// fmt.Println("response error = - 11 ReadEagain") //ReadEagain错误
								break
							}
							fmt.Printf("Error while receiving a frame from the decoder: %s\n", avutil.ErrorFromCode(response))
							return ""
						}

						if frameNumber <= frameNum {
							// Convert the image from its native format to RGB
							swscale.SwsScale2(swsCtx, avutil.Data(pFrame),
								avutil.Linesize(pFrame), 0, pCodecCtx.Height(),
								avutil.Data(pFrameRGB), avutil.Linesize(pFrameRGB))

							// Save the frame to disk
							urlStrings := strings.Split(url, ".")
							urlstrings2 := strings.Split(urlStrings[len(urlStrings)-2], "/")
							movieUrl := urlstrings2[1]
							return SaveFrameJpg(pFrameRGB, pCodecCtx.Width(), pCodecCtx.Height(), movieUrl)
						}
						//frameNumber++
					}
				}

				// Free the packet that was allocated by av_read_frame
				packet.AvFreePacket()
			}

			// Free the RGB image
			avutil.AvFree(buffer)
			avutil.AvFrameFree(pFrameRGB)

			// Free the YUV frame
			avutil.AvFrameFree(pFrame)

			// Close the codecs
			pCodecCtx.AvcodecClose()
			(*avcodec.Context)(unsafe.Pointer(pCodecCtxOrig)).AvcodecClose()

			// Close the video file
			pFormatContext.AvformatCloseInput()

			// Stop after saving frames of first video straem
			break
		case avformat.AVMEDIA_TYPE_DATA:
			i = 2
			// Get a pointer to the codec context for the video stream
			pCodecCtxOrig := pFormatContext.Streams()[i].Codec()
			// Find the decoder for the video stream
			pCodec := avcodec.AvcodecFindDecoder(avcodec.CodecId(pCodecCtxOrig.GetCodecId()))
			if pCodec == nil {
				fmt.Println("Unsupported codec!")
				return ""
			}
			// Copy context
			pCodecCtx := pCodec.AvcodecAllocContext3()
			if pCodecCtx.AvcodecCopyContext((*avcodec.Context)(unsafe.Pointer(pCodecCtxOrig))) != 0 {
				fmt.Println("Couldn't copy codec context")
				return ""
			}

			// Open codec
			if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
				fmt.Println("Could not open codec")
				return ""
			}

			// Allocate video frame
			pFrame := avutil.AvFrameAlloc()

			// Allocate an AVFrame structure
			pFrameRGB := avutil.AvFrameAlloc()
			if pFrameRGB == nil {
				fmt.Println("Unable to allocate RGB Frame")
				return ""
			}

			// Determine required buffer size and allocate buffer
			numBytes := uintptr(avcodec.AvpictureGetSize(avcodec.AV_PIX_FMT_RGB24, pCodecCtx.Width(),
				pCodecCtx.Height()))
			buffer := avutil.AvMalloc(numBytes)

			// Assign appropriate parts of buffer to image planes in pFrameRGB
			// Note that pFrameRGB is an AVFrame, but AVFrame is a superset
			// of AVPicture
			avp := (*avcodec.Picture)(unsafe.Pointer(pFrameRGB))
			avp.AvpictureFill((*uint8)(buffer), avcodec.AV_PIX_FMT_RGB24, pCodecCtx.Width(), pCodecCtx.Height())

			// initialize SWS context for software scaling
			swsCtx := swscale.SwsGetcontext(
				pCodecCtx.Width(),
				pCodecCtx.Height(),
				(swscale.PixelFormat)(pCodecCtx.PixFmt()),
				pCodecCtx.Width(),
				pCodecCtx.Height(),
				avcodec.AV_PIX_FMT_RGB24,
				avcodec.SWS_BILINEAR,
				nil,
				nil,
				nil,
			)

			// Read frames and save first five frames to disk
			frameNumber := 1
			packet := avcodec.AvPacketAlloc()
			for pFormatContext.AvReadFrame(packet) >= 0 {
				// Is this a packet from the video stream?
				if packet.StreamIndex() == i {
					// Decode video frame
					response := pCodecCtx.AvcodecSendPacket(packet)
					if response < 0 {
						fmt.Printf("Error while sending a packet to the decoder: %s\n", avutil.ErrorFromCode(response))
					}
					for response >= 0 {
						response = pCodecCtx.AvcodecReceiveFrame((*avcodec.Frame)(unsafe.Pointer(pFrame)))
						if response == avutil.AvErrorEAGAIN || response == avutil.AvErrorEOF {
							break
						} else if response < 0 {
							if response == -11 {
								// fmt.Println("response error = - 11 ReadEagain") //ReadEagain错误
								break
							}
							fmt.Printf("Error while receiving a frame from the decoder: %s\n", avutil.ErrorFromCode(response))
							return ""
						}

						if frameNumber <= frameNum {
							// Convert the image from its native format to RGB
							swscale.SwsScale2(swsCtx, avutil.Data(pFrame),
								avutil.Linesize(pFrame), 0, pCodecCtx.Height(),
								avutil.Data(pFrameRGB), avutil.Linesize(pFrameRGB))

							// Save the frame to disk
							//fmt.Printf("Writing frame %d\n", frameNumber)
							urlStrings := strings.Split(url, ".")
							urlstrings2 := strings.Split(urlStrings[len(urlStrings)-2], "/")
							movieUrl := urlstrings2[1]
							return SaveFrameJpg(pFrameRGB, pCodecCtx.Width(), pCodecCtx.Height(), movieUrl)
						}
						frameNumber++
					}
				}

				// Free the packet that was allocated by av_read_frame
				packet.AvFreePacket()
			}

			// Free the RGB image
			avutil.AvFree(buffer)
			avutil.AvFrameFree(pFrameRGB)

			// Free the YUV frame
			avutil.AvFrameFree(pFrame)

			// Close the codecs
			pCodecCtx.AvcodecClose()
			(*avcodec.Context)(unsafe.Pointer(pCodecCtxOrig)).AvcodecClose()

			// Close the video file
			pFormatContext.AvformatCloseInput()

			// Stop after saving frames of first video straem
			break
		default:
			fmt.Println("Didn't find a video stream")
			return ""
		}
	}
	return ""
}

// SaveFrameJpg SaveFrame writes a single frame to disk as a JPG file
func SaveFrameJpg(frame *avutil.Frame, width, height int, movieUrl string) string {
	// Open file
	fileName := fmt.Sprintf("./app/video/tmp/" + movieUrl + "_cover.jpg")
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error Reading")
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		data0 := avutil.Data(frame)[0]
		// startPos := uintptr(unsafe.Pointer(data0)) + uintptr(y)*uintptr(avutil.Linesize(frame)[0])
		//fmt.Println("startPos: ", startPos)
		xxx := 0
		for x := 0; x < width; x++ {
			var pixel = make([]byte, 3)
			for i := 0; i < 3; i++ {
				element := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data0)) + uintptr(y)*uintptr(avutil.Linesize(frame)[0]) + uintptr(xxx)))
				pixel[i] = element
				// fmt.Println(element)
				xxx++
			}
			img.SetRGBA(x, y, color.RGBA{pixel[0], pixel[1], pixel[2], 0xff})
		}
	}
	//img image.Image
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		fmt.Println("jpeg.Encode err: ", err)
		fmt.Println("Error jpeg.Encode")
		return ""
	}
	return fileName
}
