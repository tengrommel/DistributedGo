package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
)

func main() {
	xys, err := readData("ch16/data.txt")
	if err != nil{
		log.Fatalf("could not read data.txt: %v", err)
	}
	err = plotData("out.png", xys)
	if err != nil{
		log.Fatalf("could not plot data: %v", err)
	}
}
func plotData(path string, xys []xy) error {
	f, err := os.Create(path)
	if err != nil{
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p, err := plot.New()
	if err != nil{
		return fmt.Errorf("could not create plot: %v", err)
	}

	pxys := make(plotter.XYs, len(xys))
	for i, xy :=range xys{
		pxys[i].X = xy.x
		pxys[i].Y = xy.y
	}

	s, err := plotter.NewScatter(pxys)
	if err != nil{
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R:255, A:255}
	p.Add(s)

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil{
		return fmt.Errorf("could not create writer: %v", err)
	}


	_, err = wt.WriteTo(f)
	if err != nil{
		return fmt.Errorf("could not write to %s: %v",path, err)
	}
	if err := f.Close(); err != nil{
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}

type xy struct {
	x, y float64
}

func readData(path string) ([]xy, error) {
	f, err := os.Open(path)
	if err != nil{
		return nil, err
	}

	defer f.Close()
	var xys []xy
	s := bufio.NewScanner(f)
	for s.Scan(){
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil{
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		xys = append(xys, xy{x,y})
	}
	if err := s.Err(); err != nil{
		return nil, fmt.Errorf("could not scan: %v", )
	}
	return xys, nil
}
