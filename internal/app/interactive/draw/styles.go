package draw

import "github.com/gdamore/tcell/v2"

var searchHighlightColor = tcell.Color172
var mainColor = tcell.Color126
var bgColor = tcell.Color233
var textColor = tcell.ColorWhite
var inputColor = tcell.Color235
var strongTextColor = tcell.Color141
var weakTextColor = tcell.Color139
var errorColor = tcell.Color160
var statusColor = tcell.Color106
var activeColor = tcell.Color189

var DefaultStyle = tcell.StyleDefault.Background(bgColor)
var inputStyle = DefaultStyle.Background(inputColor).Foreground(textColor)
var buttonStyle = DefaultStyle.Background(mainColor).Foreground(textColor)
var activeTabStyle = DefaultStyle.Foreground(bgColor).Background(activeColor)
var activeStyle = DefaultStyle.Foreground(mainColor).Background(activeColor)
var textStyle = DefaultStyle.Foreground(textColor)
var weakTextStyle = DefaultStyle.Foreground(weakTextColor)
