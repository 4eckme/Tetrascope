//Tetrascope sample
//R=1024px, scale=100%

package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "os"
    "strings"
    "strconv"
)

func ConvertToColor(s string) (r string) {

    r = s
    if len(s)<6 {
        r = s+strings.Repeat("0", 6-len(s))
    } else if len(s)>6 {
        r = string(s[0:6])
    }
    return r
}

func ParseHexColorFast(s string) (c color.RGBA) {

    c.A = 0xff

    if s[0] != '#' {
        return c
    }

    hexToByte := func(b byte) byte {
        switch {
        case b >= '0' && b <= '9':
            return b - '0'
        case b >= 'a' && b <= 'f':
            return b - 'a' + 10
        case b >= 'A' && b <= 'F':
            return b - 'A' + 10
        }
        return 0
    }

    c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
    c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
    c.B = hexToByte(s[5])<<4 + hexToByte(s[6])

    return
}

func main() {
    
    R:=1024
    D:=R*2+1;    
    new_png_file := "tetrascope.png"
    
    myimage := image.NewRGBA(image.Rect(0, 0, D, D))
    bgcolor := color.RGBA{0, 0, 0, 0}
    draw.Draw(myimage, myimage.Bounds(), &image.Uniform{bgcolor}, image.ZP, draw.Src)

    for x:=0; x<=R; x++ {
        for y:=0; y<=R; y++ {
            x1 := R-x
            x2 := R+x
            y1 := R-y
            y2 := R+y
            if x*x+y*y <= R*R {
                var c int64;                
                c = int64(x*x+y*y);
                col := "#"+ConvertToColor(strconv.FormatInt(c, 16))
                color := ParseHexColorFast(col)
                pixel1 := image.Rect(x1, y1, x1+1, y1+1)
                pixel2 := image.Rect(x1, y2, x1+1, y2+1)
                pixel3 := image.Rect(x2, y1, x2+1, y1+1)
                pixel4 := image.Rect(x2, y2, x2+1, y2+1)
                draw.Draw(myimage, pixel1, &image.Uniform{color}, image.ZP, draw.Src)
                draw.Draw(myimage, pixel2, &image.Uniform{color}, image.ZP, draw.Src)
                draw.Draw(myimage, pixel3, &image.Uniform{color}, image.ZP, draw.Src)
                draw.Draw(myimage, pixel4, &image.Uniform{color}, image.ZP, draw.Src)
            }
        }   
    }

    myfile, err := os.Create(new_png_file)
    if err != nil {
        panic(err)
    }
    png.Encode(myfile, myimage)
}
