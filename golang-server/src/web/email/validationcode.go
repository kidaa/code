/**
 * Created by Michael on 2015/7/28.
 */
package email
import (
	"image"
	"image/png"
	"io"
	"math/rand"
	"time"
	"strconv"
	"io/ioutil"
	"freetype"
	"image/draw"
)
type ValidationCode struct {
}

//生成一个新的验证码
func (this *ValidationCode) NewCdoe(len int) string {
	r := rand.New(rand.NewSource(int64(time.Now().Second())))
	var strCode string
	for i := 0; i < len; i++ {
		n := r.Intn(10)
		strCode += strconv.Itoa(n)
	}
	return strCode
}

func (this *ValidationCode) DrawToImg(strCode string, w io.Writer) error {
	arrFontFile := []string{"03.ttf"}
	r := rand.New(rand.NewSource(int64(time.Now().Second())))
	fIndex := r.Intn(len(arrFontFile))
	strFontFile := "./fonts/" + arrFontFile[fIndex]
	fontBytes, err := ioutil.ReadFile(strFontFile)
	if err != nil {
		return err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	c := freetype.NewContext()
	var fontSize float64
	fontSize = 20
	c.SetDPI(120)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	//计算字符串的宽度，对于高度，还有此问题，懂的可以改改
	width, startY := c.MeasureString(strCode)
	heigth := c.FUnitToPixelRU(font.UnitsPerEm())
	width += 10
	heigth += 10
	fg, bg := image.Black, image.Transparent			// 背景透明
	rgba := image.NewRGBA(image.Rect(0, 0, width, heigth))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	//	this.disturbBitmap(rgba)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	pt := freetype.Pt(5, 5+startY)
	_, err = c.DrawString(strCode, pt)
	if err != nil {
		return err
	}
	err = png.Encode(w, rgba)
	return nil
}

////绘制干扰背景
//func (this *ValidationCode) disturbBitmap(img *image.RGBA) {
//	r := rand.New(rand.NewSource(int64(time.Now().Second())))
//	for i := 0; i < img.Rect.Max.X; i++ {
//		for j := 0; j < img.Rect.Max.Y; j++ {
//			n := r.Intn(100)
//			if n < 40 {
//				c := color.NRGBA{uint8(r.Intn(150)), uint8(r.Intn(150)), uint8(r.Intn(150)), uint8(r.Intn(100))}
//				img.Set(i, j, c)
//			}
//		}
//	}
//}
