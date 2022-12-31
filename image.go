package fltk

/*
#include <stdlib.h>
#include "cfltk/cfl_image.h"
*/
import "C"
import (
	"errors"
	goimage "image"
	"unsafe"
)

type image struct {
	iPtr *C.Fl_Image
}

type Image interface {
	getImage() *image
}

var ErrImageDestroyed = errors.New("image is destroyed")

func (i *image) getImage() *image { return i }

func (i *image) ptr() *C.Fl_Image {
	if i.iPtr == nil {
		panic(ErrImageDestroyed)
	}
	return i.iPtr
}

func initImage(i Image, p unsafe.Pointer) {
	i.getImage().iPtr = (*C.Fl_Image)(p)
}

func (i *image) Destroy() {
	C.Fl_Image_delete(i.ptr())
	i.iPtr = nil
}

func (i *image) Draw(x, y, w, h int) {
	C.Fl_Image_draw(i.ptr(), C.int(x), C.int(y), C.int(w), C.int(h))
}

func (i *image) W() int {
	return int(C.Fl_Image_width(i.ptr()))
}

func (i *image) H() int {
	return int(C.Fl_Image_height(i.ptr()))
}

func (i *image) Count() int {
	return int(C.Fl_Image_count(i.ptr()))
}

func (i *image) Scale(width int, height int, proportional bool, can_expand bool) {
	prop := 0
	expand := 0
	if proportional {
		prop = 1
	}
	if can_expand {
		expand = 1
	}
	C.Fl_Image_scale(i.ptr(), C.int(width), C.int(height), C.int(prop), C.int(expand))
}

func (i *image) fail() int {
	return int(C.Fl_Image_fail(i.ptr()))
}

func (i *image) DataW() int {
	return int(C.Fl_Image_data_w(i.ptr()))
}

func (i *image) DataH() int {
	return int(C.Fl_Image_data_h(i.ptr()))
}

func (i *image) D() int {
	return int(C.Fl_Image_d(i.ptr()))
}

func (i *image) Ld() int {
	return int(C.Fl_Image_ld(i.ptr()))
}

func (i *image) Inactive() {
	C.Fl_Image_inactive(i.ptr())
}

func image_error(val int) error {
	if val == -1 {
		return errors.New("no Image was found")
	} else if val == -2 {
		return errors.New("file access error")
	} else if val == -3 {
		return errors.New("image format error")
	} else {
		return nil
	}
}

type SvgImage struct {
	image
}

func NewSvgImageLoad(path string) (*SvgImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &SvgImage{}
	initImage(img, unsafe.Pointer(C.Fl_SVG_Image_new(fileStr)))
	return img, image_error(img.fail())
}

func NewSvgImageFromString(str string) (*SvgImage, error) {
	imagestr := C.CString(str)
	defer C.free(unsafe.Pointer(imagestr))
	img := &SvgImage{}
	initImage(img, unsafe.Pointer(C.Fl_SVG_Image_from(imagestr)))
	return img, image_error(img.fail())
}

type PngImage struct {
	image
}

func NewPngImageLoad(path string) (*PngImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &PngImage{}
	initImage(img, unsafe.Pointer(C.Fl_PNG_Image_new(fileStr)))
	return img, image_error(img.fail())
}

func NewPngImageFromData(data []uint8) (*PngImage, error) {
	buf := (*C.uchar)(unsafe.Pointer(&data[0]))

	img := &PngImage{}

	initImage(img, unsafe.Pointer(C.Fl_PNG_Image_from(buf, C.int(len(data)))))
	return img, image_error(img.fail())
}

type JpegImage struct {
	image
}

func NewJpegImageLoad(path string) (*JpegImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &JpegImage{}
	initImage(img, unsafe.Pointer(C.Fl_JPEG_Image_new(fileStr)))
	return img, image_error(img.fail())
}

func NewJpegImageFromData(data []uint8) (*JpegImage, error) {
	buf := (*C.uchar)(unsafe.Pointer(&data[0]))

	img := &JpegImage{}

	initImage(img, unsafe.Pointer(C.Fl_JPEG_Image_from(buf)))
	return img, image_error(img.fail())
}

type BmpImage struct {
	image
}

func NewBmpImageLoad(path string) (*BmpImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &BmpImage{}
	initImage(img, unsafe.Pointer(C.Fl_BMP_Image_new(fileStr)))
	return img, image_error(img.fail())
}

func NewBmpImageFromData(data []uint8) (*BmpImage, error) {
	buf := (*C.uchar)(unsafe.Pointer(&data[0]))

	img := &BmpImage{}

	initImage(img, unsafe.Pointer(C.Fl_BMP_Image_from(buf, C.long(len(data)))))
	return img, image_error(img.fail())
}

type RgbImage struct {
	image
	data []uint8
}

func NewRgbImage(data []uint8, w, h, depth int) (*RgbImage, error) {
	img := &RgbImage{}
	img.data = append(make([]uint8, 0, len(data)), data...)

	buf := (*C.uchar)(unsafe.Pointer(&img.data[0]))

	initImage(img, unsafe.Pointer(C.Fl_RGB_Image_new(buf, C.int(w), C.int(h), C.int(depth), C.int(0))))
	return img, image_error(img.fail())
}

func NewRgbImageFromImage(image goimage.Image) (*RgbImage, error) {
	rgbImage := &RgbImage{}
	var w, h, stride, depth int
	if grayImage, ok := image.(*goimage.Gray); ok {
		rgbImage.data = append(make([]uint8, 0, len(grayImage.Pix)), grayImage.Pix...)
		w, h, stride = grayImage.Rect.Dx(), grayImage.Rect.Dy(), grayImage.Stride
		depth = 1
	} else if rgbaImage, ok := image.(*goimage.RGBA); ok {
		rgbImage.data = append(make([]uint8, 0, len(rgbaImage.Pix)), rgbaImage.Pix...)
		w, h, stride = rgbaImage.Rect.Dx(), rgbaImage.Rect.Dy(), rgbaImage.Stride
		depth = 4
	} else {
		w, h, stride = image.Bounds().Dx(), image.Bounds().Dy(), 0
		depth = 4
		rgbImage.data = make([]uint8, 0, w*h*4)
		for dy := 0; dy < h; dy++ {
			y := image.Bounds().Min.Y + dy
			for dx := 0; dx < w; dx++ {
				x := image.Bounds().Min.X + dx
				r, g, b, a := image.At(x, y).RGBA()
				r8, g8, b8, a8 := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)
				rgbImage.data = append(rgbImage.data, r8, g8, b8, a8)
			}
		}
	}
	buf := (*C.uchar)(unsafe.Pointer(&rgbImage.data[0]))
	initImage(rgbImage, unsafe.Pointer(C.Fl_RGB_Image_from_data(buf, C.int(w), C.int(h), C.int(depth), C.int(stride))))
	return rgbImage, image_error(rgbImage.fail())
}

type SharedImage struct {
	image
}

var shared_init bool

func register_images() {
	C.Fl_register_images()
}

func NewSharedImageLoad(path string) (*SharedImage, error) {
	if !shared_init {
		register_images()
		shared_init = true
	}
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &SharedImage{}
	initImage(img, unsafe.Pointer(C.Fl_Shared_Image_get(fileStr, C.int(0), C.int(0))))
	if img.iPtr == nil {
		return nil, errors.New("Shared Image initialization error")
	}
	return img, image_error(img.fail())
}
