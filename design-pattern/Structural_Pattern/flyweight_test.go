/*
@Time : 2022/10/29 20:34
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : flyweight_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Structural_Pattern

import (
	"fmt"
	"testing"
)

type ImageFlyweight struct {
	data string
}

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

var imageFactory *ImageFlyweightFactory

func GetImageFlyweightFactory () *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}
	return image
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	return &ImageFlyweight{
		data: fmt.Sprintf("image data %s", filename),
	}
}

func (i *ImageFlyweight)Data() string {
	return i.data
}

type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{image}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.data)
}

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")
	viewer1.Display()
	viewer2.Display()
	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		fmt.Println("viewer1.ImageFlyweight != viewer2.ImageFlyweight")
	} else {
		fmt.Println("viewer1.ImageFlyweight == viewer2.ImageFlyweight")
	}
}