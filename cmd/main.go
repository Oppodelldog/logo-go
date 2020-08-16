package main

import (
	"fmt"
	"github.com/Oppodelldog/logo-go/arts"
	"github.com/Oppodelldog/logo-go/turtle"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
)

const width = 1920
const height = 1080

type ArtWork struct {
	drawFrames func()
	frameRate  int
	codec      Codec
	makeMovie  bool
}

type Codec struct {
	video     string
	pixel     string
	container string
}

var apgnCodec = Codec{
	video:     "apng",
	pixel:     "rgb24",
	container: ".apng",
}
var gifCodec = Codec{
	video:     "gif",
	pixel:     "rgb24",
	container: ".gif",
}
var mp4Codec = Codec{
	video:     "libx264",
	pixel:     "yuv420p",
	container: ".mp4",
}
var mp4Codec2 = Codec{
	video:     "h264",
	pixel:     "yuv420p",
	container: ".mp4",
}
var mp4Codec3 = Codec{
	video:     "mjpeg",
	pixel:     "yuv420p",
	container: ".mp4",
}

var artWorks = map[string]ArtWork{
	"ChelsyWagner": {
		makeMovie: true,
		frameRate: 20,
		codec:     mp4Codec,
		drawFrames: func() {
			for i := 0; i < 360; i++ {
				dest, gc := initImage()
				arts.ChelsyWagner(float64(i))(turtle.New(gc), width, height)
				finishImage(gc, dest, i)
			}
		}},
	"DieselMalone": {
		makeMovie: true,
		frameRate: 30,
		codec:     mp4Codec3,
		drawFrames: func() {
			for i := 1; i < 360; i++ {
				dest, gc := initImage()
				arts.DieselMalone(float64(i))(turtle.New(gc), width, height)
				finishImage(gc, dest, i)
			}
		},
	},
	"DanicaVelez": {
		makeMovie: true,
		frameRate: 30,
		codec:     mp4Codec3,
		drawFrames: func() {
			var frame int
			for i := float64(0); i < 360; i += 0.2 {
				dest, gc := initImage()
				arts.DanicaVelez(float64(i))(turtle.New(gc), width, height)
				finishImage(gc, dest, frame)
				frame++
			}
		},
	},
	"UrsulaConnolly": {
		makeMovie: true,
		frameRate: 30,
		codec:     mp4Codec3,
		drawFrames: func() {
			var frame int
			for i := float64(0); i < 30; i += 1 {
				fmt.Print(".")
				dest, gc := initImage()
				arts.UrsulaConnolly(float64(i))(turtle.New(gc), width, height)
				finishImage(gc, dest, frame)
				frame++
			}
		},
	},
	"MalaikaKelly": {
		makeMovie: true,
		frameRate: 30,
		codec:     mp4Codec3,
		drawFrames: func() {
			var frame int
			for i := float64(0); i < 30; i += 0.1 {
				if int(i)%10==0{
					fmt.Print("X")
				}else{
					fmt.Print(".")
				}

				dest, gc := initImage()
				arts.MalaikaKelly(float64(i))(turtle.New(gc), width, height)
				finishImage(gc, dest, frame)
				frame++
			}
		},
	},
}

func main() {
	doThe("MalaikaKelly")
}

func doThe(name string) {
	artWork := artWorks[name]

	fmt.Println("draw frames")
	artWork.drawFrames()
	if artWorks[name].makeMovie {
		fmt.Println("make movie")
		makeMovie(name, artWork)
	}
}

func makeMovie(outname string, artWork ArtWork) {
	cmd := exec.Command("ffmpeg", "-y",
		"-framerate", strconv.Itoa(artWork.frameRate),
		"-r", strconv.Itoa(artWork.frameRate),
		"-i", "%04d.png", "-c:v",
		artWork.codec.video,
		"-pix_fmt", artWork.codec.pixel,
		"../"+outname+artWork.codec.container)
	cmd.Dir = "out/img"
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("ERROR creating video")
		fmt.Println(string(output))
	}

	filepath.Walk("out/img", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) == ".png" {
			os.Remove(path)
		}

		return nil
	})
}

func finishImage(gc *draw2dimg.GraphicContext, dest *image.RGBA, frame int) {
	gc.Stroke()
	err := draw2dimg.SaveToPngFile(path.Join("out", "img", fmt.Sprintf("%04d.png", frame)), dest)
	if err != nil {
		panic(err)
	}
}

func initImage() (*image.RGBA, *draw2dimg.GraphicContext) {
	dest := image.NewRGBA(image.Rect(0, 0, width, height))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x0, 0x0, 0x0, 0xff})
	gc.SetStrokeColor(color.RGBA{0x0, 0xff, 0x0, 0xff})
	gc.SetLineWidth(1)

	gc.Clear()
	gc.BeginPath()
	return dest, gc
}
