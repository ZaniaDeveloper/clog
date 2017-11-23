package clog

import (
	"github.com/mgutz/ansi"
	"reflect"
	"strings"
)

const (
	THEME_STRING_SEPARATOR = ","
	THEME_VALUE_SEPARATOR  = "="

	_DEF_LEVEL_THEME = "TRAC,DBUG=grey,INFO,MESG=cyan+h,WARN=yellow+h,ERRO=red,FATL=white+h:red"
	_DEF_STACK_THEME = "stak=red,func,file,line=red+h"
	_DEF_LOG_THEME   = "tim=blue+h,pre=yellow,msg=magenta+h,key=green+h,val=green,oth=default"
	DEFAULT_THEME    = _DEF_LOG_THEME + "," + _DEF_STACK_THEME + "," + _DEF_LEVEL_THEME
)

var defaultTheme = parseTheme(DEFAULT_THEME)

type Theme struct {
	LogKey       string `theme:"key" json:"key,omitempty" yaml:"key,omitempty"`
	LogValue     string `theme:"val" json:"val,omitempty" yaml:"val,omitempty"`
	LogTimestamp string `theme:"tim" json:"tim,omitempty" yaml:"tim,omitempty"`
	LogPrefix    string `theme:"pre" json:"pre,omitempty" yaml:"pre,omitempty"`
	LogMessage   string `theme:"msg" json:"msg,omitempty" yaml:"msg,omitempty"`
	LogOther     string `theme:"oth" json:"oth,omitempty" yaml:"oth,omitempty"`

	StackFunction string `theme:"func" json:"func,omitempty" yaml:"func,omitempty"`
	StackFilename string `theme:"file" json:"file,omitempty" yaml:"file,omitempty"`
	StackLine     string `theme:"line" json:"line,omitempty" yaml:"line,omitempty"`
	StackOther    string `theme:"stak" json:"stak,omitempty" yaml:"stak,omitempty"`

	LevelTrace   string `theme:"TRAC" json:"TRAC,omitempty" yaml:"TRAC,omitempty"`
	LevelDebug   string `theme:"DBUG" json:"DBUG,omitempty" yaml:"DBUG,omitempty"`
	LevelMessage string `theme:"MESG" json:"MESG,omitempty" yaml:"MESG,omitempty"`
	LevelInfo    string `theme:"INFO" json:"INFO,omitempty" yaml:"INFO,omitempty"`
	LevelWarn    string `theme:"WARN" json:"WARN,omitempty" yaml:"WARN,omitempty"`
	LevelError   string `theme:"ERRO" json:"ERRO,omitempty" yaml:"ERRO,omitempty"`
	LevelFatal   string `theme:"FATL" json:"FATL,omitempty" yaml:"FATL,omitempty"`
}

func adaptColor(color string) string {
	if strings.HasPrefix(color, "grey") || strings.HasPrefix(color, "gray") {
		color = "black+h" + color[5:]
	} else {
		switch color[0] {
		case '-':
			color = "default"
			if len(color) > 1 {
				color += color[1:]
			}

		case '+':
			color = "default" + color
		}
	}

	return color
}

func parseTheme(theme string) *Theme {
	pairs := strings.Split(theme, THEME_STRING_SEPARATOR)
	if len(pairs) == 0 {
		return nil
	}

	var keys []string

	styles := map[string]string{}
	for _, pair := range pairs {
		if pair == "" {
			continue
		}

		parts := strings.Split(pair, THEME_VALUE_SEPARATOR)
		keys = append(keys, parts[0])

		if len(parts) > 1 {
			style := parts[1]
			for _, key := range keys {
				styles[key] = style
			}
		}
	}

	var color = func(key string) string {
		style := styles[key]
		switch style {
		case "~":
			style = "reset"

		case "grey", "gray", "grey+", "gray+":
			style = "black+h"

		case "":
			return ""

		default:
			colors := strings.Split(style, ":")

			for i, color := range colors {
				colors[i] = adaptColor(color)
			}

			style = strings.Join(colors, ":")
		}

		return ansi.ColorCode(style)
	}

	res := new(Theme)
	allStyle, ok := styles["*"]
	if !ok {
		allStyle, ok = styles["ALL"]
	}

	val := reflect.ValueOf(res).Elem()

	if ok {
		col := reflect.ValueOf(ansi.ColorCode(allStyle))
		for i := 0; i < val.NumField(); i++ {
			val.Field(i).Set(col)
		}
	} else if defaultTheme != nil {
		*res = *defaultTheme
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		col := color(typ.Field(i).Tag.Get("theme"))
		if col != "" {
			val.Field(i).Set(reflect.ValueOf(col))
		}
	}

	return res
}
