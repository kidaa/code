//对freetype的扩展加入一个MeasureString 来获取字符串所点的高度，宽度
package freetype

import (
	"freetype/truetype"
)

//从 DrawString 复制修改，用来计算字符串所点的宽度
func (c *Context) MeasureString(s string) (width, maskY int) {
	if c.font == nil {
		panic("freetype: DrawText called with a nil font")
	}

	prev, hasPrev := truetype.Index(0), false
	arrRune := []rune(s)
	p := Pt(0, 0+c.FUnitToPixelRU(c.font.UnitsPerEm()))
	for _, r := range arrRune {
		index := c.font.Index(r)
		if hasPrev {
			p.X += c.FUnitToFix32(int(c.font.Kerning(prev, index)))
		}
		mask, _, _ := c.glyph(index, p)
		b := mask.Bounds()
		if b.Max.Y > maskY {
			maskY = b.Max.Y
		}

		if hasPrev {
			u := int(c.font.Kerning(prev, index))
			width += int(u)
			p.X += c.FUnitToFix32(u)
		}
		u := int(c.font.HMetric(index).AdvanceWidth)
		width += u
		p.X += c.FUnitToFix32(u)
		prev, hasPrev = index, true
	}
	width = c.FUnitToPixelRU(width)
	return
}
