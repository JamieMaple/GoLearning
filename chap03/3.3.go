// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	red           = "#f00"
	blue          = "#00f"
)

var sin30, cos30 = math.Tan(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", getPic)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func getPic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, isPeak := corner(i+1, j+1)
			style := blue
			if isPeak {
				style = red
			}
			fmt.Fprintf(w, "<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				style, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprint(w, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	isPeak := false
	if z >= 0 {
		isPeak = true
	} else {
		isPeak = false
	}
	return sx, sy, isPeak
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	if r != 0 {
		return math.Sin(r) / r
	} else {
		return 0
	}
}
