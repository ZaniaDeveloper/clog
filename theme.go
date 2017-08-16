package plog

import (
	"strings"
	"github.com/mgutz/ansi"
	"reflect"
)

const (
	THEME_STRING_SEPARATORS = ","
	THEME_VALUE_SEPARATOR   = "="

	LEVEL_THEME = "TRAC,DBUG=grey,INFO=cyan+h,WARN=yellow+h,ERRO=red,FATL=white+h:red"
	STACK_THEME = "stak=red,func,file,line=red+h"
	LOG_THEME = "tim=blue+h,pre=yellow,msg=magenta+h,key=green+h,val=green,oth=default"
	DEFAULT_THEME = LOG_THEME + "," + STACK_THEME + "," + LEVEL_THEME
)

var defaultTheme = parseTheme(DEFAULT_THEME)

type Theme struct {
	LogKey       string `theme:"key"`
	LogValue     string `theme:"val"`
	LogTimestamp string `theme:"tim"`
	LogPrefix    string `theme:"pre"`
	LogMessage   string `theme:"msg"`
	LogOther     string `theme:"oth"`

	StackFunction string `theme:"func"`
	StackFilename string `theme:"file"`
	StackLine     string `theme:"line"`
	StackOther    string `theme:"stak"`

	LevelTrace string `theme:"TRAC"`
	LevelDebug string `theme:"DBUG"`
	LevelInfo  string `theme:"INFO"`
	LevelWarn  string `theme:"WARN"`
	LevelError string `theme:"ERRO"`
	LevelFatal string `theme:"FATL"`
}

func adaptColor(color string) string {
	if strings.HasPrefix(color, "grey") || strings.HasPrefix(color, "gray") {
		color = "black+h" + color[5:]
	} else {
		switch color[0] {
		case "-":
			color = "default"
			if len(color) > 1 {
				color += color[1:]
			}

		case "+":
			color = "default" + color
		}
	}

	return color
}

func parseTheme(theme string) *Theme {
	pairs := strings.Split(theme, THEME_STRING_SEPARATORS)
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
