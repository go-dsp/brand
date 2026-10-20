// Command mark writes the go-dsp mark, svg/color/go-dsp.svg:
//
//	cd mark && go run . > ../svg/color/go-dsp.svg
//
// go-dsp is signal processing, and the mark is a signal: a tone of three
// periods under a Hann window, the burst a filter or a spectral estimator is
// handed. The trace is a formula, y = A sin(2 pi p t) sin(pi t) for t from 0 to
// 1, sampled at evenly spaced points and joined by a polyline whose round joins
// hide the sampling. Nothing is placed by hand: change periods, amplitude or
// samples and the mark redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#F08AB0" // the gradient's light stop
	dark  = "#B3245F" // its dark stop

	stroke = 13.0 // the family's line width

	left, right = 36.0, 220.0 // the glyph's horizontal extent
	amplitude   = 58.0        // peak of the window, from the axis
	periods     = 3.0         // periods of the tone
	samples     = 96          // segments of the polyline
)

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func main() {
	pts := make([]string, 0, samples+1)
	for i := 0; i <= samples; i++ {
		t := float64(i) / samples
		x := left + (right-left)*t
		y := 128 - amplitude*math.Sin(2*math.Pi*periods*t)*math.Sin(math.Pi*t)
		pts = append(pts, num(x)+","+num(y))
	}
	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-dsp">
  <defs><linearGradient id="dspgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#dspgrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    <polyline points="%s"/>
  </g>
</svg>
`, light, dark, num(stroke), strings.Join(pts, " "))
}
