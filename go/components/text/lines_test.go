package text

import (
	"testing"

	"github.com/meian/hansei-board/go/core/font"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinesByWidth(t *testing.T) {
	face, err := font.LoadFace(font.Propotional, 20)
	require.NoError(t, err)
	type args struct {
		s string
		w int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "横幅が狭いと改行される",
			args: args{
				s: "あいうえおかきくけこさしすせそ",
				w: 100,
			},
			want: []string{"あいうえお", "かきくけこさ", "しすせそ"},
		},
		{
			name: "横幅が広いと改行されない",
			args: args{
				s: "あいうえおかきくけこさしすせそ",
				w: 500,
			},
			want: []string{"あいうえおかきくけこさしすせそ"},
		},
		{
			name: "空文字は空で返る",
			args: args{
				s: "",
				w: 500,
			},
			want: []string{""},
		},
		{
			name: "改行のみ",
			args: args{
				s: "\n",
				w: 100,
			},
			want: []string{""},
		},
		{
			name: "改行箇所は任意に改行される",
			args: args{
				s: "あいうえお\nかきくけこさしすせそ",
				w: 85,
			},
			want: []string{"あいうえ", "お", "かきくけこ", "さしすせ", "そ"},
		},
		{
			name: "先頭に改行がある",
			args: args{
				s: "\nあいうえおかきくけこさしすせそ",
				w: 100,
			},
			want: []string{"あいうえお", "かきくけこさ", "しすせそ"},
		},
		{
			name: "末尾に改行がある",
			args: args{
				s: "あいうえおかきくけこさしすせそ\n",
				w: 100,
			},
			want: []string{"あいうえお", "かきくけこさ", "しすせそ"},
		},
		{
			name: "横幅位置に改行がある",
			args: args{
				s: "あいうえおかきくけこさ\nしすせそ",
				w: 100,
			},
			want: []string{"あいうえお", "かきくけこさ", "しすせそ"},
		},
		{
			name: "横幅位置に複数改行がある",
			args: args{
				s: "あいうえおかきくけこさ\n\nしすせそ",
				w: 100,
			},
			want: []string{"あいうえお", "かきくけこさ", "", "しすせそ"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := LinesByWidth(face, tt.args.s, tt.args.w)
			t.Logf("lines=%#v", lines)
			assert.Equal(t, tt.want, lines)
		})
	}
}
