package main

import (
  "fmt"
  "image/jpeg"
  "io/ioutil"
  //"log"
  "image"
  "bufio"
  "image/draw"
  "image/png"
  //"image/color"
  "github.com/golang/freetype/truetype"
  "github.com/golang/freetype"
  "os"
  "golang.org/x/image/font"
)

var (
  fontFile = "InputMono-Regular.ttf"
  imageFilename = "sample.jpg"

  dpi = 72.0
  size = 125.0
  hinting = "full"
  text="kirhgoff"
)

func main() {
  fmt.Println("Starting ASCII converter")
  fmt.Println("========================")

  var img = loadImage(imageFilename)
  fmt.Printf("Loaded file bounds: %#v\n", img.Bounds())

  fontBytes, err := ioutil.ReadFile(fontFile)
  if err != nil {
    panic(err.Error())
  }

  parsedFont, err := truetype.Parse(fontBytes)
  if err != nil {
    panic(err.Error())
  }

  // Freetype context
  //fg, bg := image.White, image.Black
  rgba := image.NewRGBA(image.Rect(0, 0, 1200, 200))
  draw.Draw(rgba, rgba.Bounds(), image.Black, image.ZP, draw.Src)
  context := freetype.NewContext()
  context.SetDPI(dpi)
  context.SetFont(parsedFont)
  context.SetFontSize(size)
  context.SetClip(rgba.Bounds())
  context.SetDst(rgba)
  context.SetSrc(image.White)
  switch hinting {
  default:
    context.SetHinting(font.HintingNone)
  case "full":
    context.SetHinting(font.HintingFull)
  }

  options := truetype.Options{}
  options.Size = size
  face := truetype.NewFace(parsedFont, &options)

  // Calculate the widths and print to image
  for i, letter := range (text) {
    advancedWidth, ok := face.GlyphAdvance(rune(letter))
    if ok != true {
      panic(err.Error())
    }
    letterWidth := int(float64(advancedWidth) / 64)
    fmt.Printf("%+v\n", letterWidth)

    pt := freetype.Pt(i * 250 + (125 - letterWidth / 2), 128)
    context.DrawString(string(letter), pt)
    fmt.Printf("%+v\n", advancedWidth)
  }

  saveImage("out.png", rgba)
}

//--------------------------------------------------
// Saves image to disk
func saveImage(filename string, rgba * image.RGBA) {
  outFile, err := os.Create(filename)
  if err != nil {
    panic(err.Error())
  }
  defer outFile.Close()

  bf := bufio.NewWriter(outFile)
  err = png.Encode(bf, rgba)
  if err != nil {
    panic(err.Error())
  }
  err = bf.Flush()
  if err != nil {
    panic(err.Error())
  }
  fmt.Println("Wrote out.png OK.")
}

//-------------------------------------------
// Loads image from disk
func loadImage(filename string) image.Image {
  fmt.Printf("loading image %s\n", filename)
  infile, err := os.Open(filename)
  if err != nil {
    panic(err.Error())
  }
  defer infile.Close()

  img, err := jpeg.Decode(infile)
  if err != nil {
    panic(err.Error())
  }
  return img
}
