package text

import (
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// LinesByWidth は文字列を face で描写した時に width に収まる文字列に分割したリストを取得する
func LinesByWidth(face font.Face, s string, width int) []string {
	s = strings.Trim(s, "\n")
	if len(s) == 0 {
		return []string{""}
	}

	w := fixed.I(width)
	ngAdv, _ := face.GlyphAdvance('\ufffd')

	// 改行コードは無条件で改行させるので事前に行単位に分割する
	linesBr := strings.Split(s, "\n")
	lines := make([]string, 0, len(linesBr))

	for _, line := range linesBr {
		var start int
		var x fixed.Int26_6
		var prevR rune
		// font.MeasureString のロジックがベース
		// ただし横幅が上限超過した時点で折り返したいのでMeasureStringは使えない
		for i, r := range line {
			if prevR > 0 {
				// 行頭でなければ文字間のマージンを追加
				x += face.Kern(prevR, r)
			}
			prevR = r
			adv, ok := face.GlyphAdvance(r)
			if !ok {
				// フォントにコードポイントが含まれない場合はNG文字の幅を想定する
				adv = ngAdv
			}
			if x+adv > w {
				lines = append(lines, line[start:i])
				start = i
				x = adv
				continue
			}
			x += adv
		}
		lines = append(lines, line[start:])
	}

	return lines
}
