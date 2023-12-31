package svg

import (
	"fmt"
	"io"
	"time"

	cf "github.com/brunoquindeler/go-with-tests/math/clockface"
)

const secondHandLenght = 90
const minuteHandLenght = 80
const hourHandLenght = 50
const clockCentreX = 150
const clockCentreY = 150

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

// Write writes an SVG representation of an analogue clock, showing the time t, to the writer w
func Write(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := makeHand(cf.SecondHandPoint(t), secondHandLenght)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(cf.MinuteHandPoint(t), minuteHandLenght)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:7px;"/>`, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(cf.HourHandPoint(t), hourHandLenght)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:7px;"/>`, p.X, p.Y)
}

func makeHand(p cf.Point, lenght float64) cf.Point {
	p = cf.Point{X: p.X * lenght, Y: p.Y * lenght}                // scale
	p = cf.Point{X: p.X, Y: -p.Y}                                 // flip
	return cf.Point{X: p.X + clockCentreX, Y: p.Y + clockCentreY} // translate
}
