package main


import (
	"log"
    "fmt"
    "time"
    "math/rand"
	"github.com/hajimehoshi/ebiten/v2"
	
)

const (
	
    screenWidth  = 320
	screenHeight = 240
)

var pass int = 0
type World struct {
	area   []bool
	width  int
	height int
}
///////// Список правил //////////////  
func Rules(x int, w *World) {
 
  switch (x) {
      case 1:    
        GameLife( w )
      
  }
    
}
/////// Правила игры "Жизнь" ////////////////
  func GameLife( w *World ) {
        
    width := w.width
	height := w.height
	next := make([]bool, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pop := neighbourCount(w.area, width, height, x, y)
			switch {
  			  case (pop == 2 || pop == 3 ) && w.area[y*width+x]:
				next[y*width+x] = true

			  case ( pop < 2 || pop > 3 ):
				next[y*width+x] = false

			  case pop == 3 :
				next[y*width+x] = true
			}
		}
	}
	w.area = next
	
  }
////////////// Вычисление соседей //////////////  
  func neighbourCount(a []bool, width, height, x, y int) int {
	c := 0
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if i == 0 && j == 0 {
				continue
			}
			x2 := x + i
			y2 := y + j
			if x2 < 0 || y2 < 0 || width <= x2 || height <= y2 {
				continue
			}
			if a[y2*width+x2] {
				c++
			}
		}
	}
	return c
}
////////////// Цвет пикселя //////////////
func (w *World) Draw(pix []byte) {
	for i, v := range w.area {
		if v {
            // Белый цвет //
 			pix[4*i] = 0xff     // R
  			pix[4*i+1] = 0xff   // G
  			pix[4*i+2] = 0xff   // B
  			pix[4*i+3] = 0xff   // A
  			
		} else {
			pix[4*i] = 0
 			pix[4*i+1] = 0
 			pix[4*i+2] = 0
 			pix[4*i+3] = 0
		}
		
		
	}
}
//// ( 1 ) Инициализируем мир /////////  
func NewWorld(width, height int, maxInitLiveCells int) *World {
	w := &World{
		area:   make([]bool, width*height),
		width:  width,
		height: height,
	}
	w.init(maxInitLiveCells)
	return w
}  
  
//// ( 2 ) Инициализируем мир /////////    
func (w *World) init(maxLiveCells int) {
	for i := 0; i < maxLiveCells; i++ {
        x := rand.Intn(w.width)
		y := rand.Intn(w.height)
		w.area[ y*w.width + x ] = true
	}
}  

////////////// Обновление мира ///////////////
func (w *World) Update() {
	
   Rules( 1, w )
}

               ////////////// Game /////////////  
type Game struct {
	
    world  *World
	pixels []byte
}

func (g *Game) Update() error {
	
    g.world.Update()
    pass ++ 
    fmt.Println("pass: ",pass)
    return nil
}

func (g *Game) Draw( screen *ebiten.Image ) {
	if g.pixels == nil {
		g.pixels = make([]byte, screenWidth*screenHeight*4) // длина пикселя
	}
	
	g.world.Draw(g.pixels)
	screen.ReplacePixels(g.pixels) 
}

func (g *Game) Layout(outsideWidth, outsideHeight int) ( int, int ) {
	
    return screenWidth, screenHeight
}

func main() {
	
    rand.Seed(time.Now().UnixNano())
    
    g := &Game {
		world: NewWorld(screenWidth, screenHeight, int((screenWidth*screenHeight)/10)),
	}
    
    ebiten.SetWindowSize( screenWidth*2, screenHeight*2 )
	ebiten.SetWindowTitle( "Hello, World!" )
	
    if err := ebiten.RunGame( g ); err != nil {
		log.Fatal( err )
	}
}
