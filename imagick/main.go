package main

import "gopkg.in/gographics/imagick.v2/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	source := imagick.NewMagickWand()
	lgtm := imagick.NewMagickWand()
	result := imagick.NewMagickWand()

	source.ReadImage("source.gif")
	lgtm.ReadImage("lgtm.gif")

	sourceWidth := source.GetImageWidth()
	sourceHeight := source.GetImageHeight()
	lgtm.ScaleImage(sourceWidth, sourceHeight)

	coalescedImages := source.CoalesceImages()
	source.Destroy()

	for i := 1; i < int(coalescedImages.GetNumberImages()); i++ {
		coalescedImages.SetIteratorIndex(i)
		tmpImage := coalescedImages.GetImage()
		tmpImage.CompositeImage(lgtm, imagick.COMPOSITE_OP_OVER, 0, 0)
		result.AddImage(tmpImage)
		tmpImage.Destroy()
	}

	lgtm.Destroy()
	coalescedImages.Destroy()

	result.WriteImages("result.gif", true)
	result.Destroy()
}
