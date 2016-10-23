package main

import "gopkg.in/gographics/imagick.v2/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	source := imagick.NewMagickWand()
	lgtm := imagick.NewMagickWand()

	source.ReadImage("source.gif")
	lgtm.ReadImage("lgtm.gif")

	source.AddImage(lgtm)
	lgtm.Destroy()
	lgtm = source.CoalesceImages()

	source.Destroy()
	source = imagick.NewMagickWand()

	for i := 1; i < int(lgtm.GetNumberImages()); i++ {
		lgtm.SetIteratorIndex(i)
		tmpImage := lgtm.GetImage()
		source.AddImage(tmpImage)
		tmpImage.Destroy()
	}

	source.ResetIterator()
	lgtm.Destroy()

	lgtm = source.CompareImageLayers(imagick.IMAGE_LAYER_COMPARE_ANY)
	lgtm.SetOption("loop", "0")

	lgtm.WriteImages("result.gif", true)
}
