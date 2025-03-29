package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand/v2"
	"os"
)

// type user struct {
// 	id       int
// 	name     string
// 	username string
// 	password string
// 	isLogin  bool
// 	text     rune
// 	number1  int32
// 	number12 int16
// 	number   int8
// 	number3  uintptr
// 	number14 uint8
// }

// func main() {

// userLogin(user{
// 	name:     "Henry ",
// 	username: "Hcore",
// 	password: "tony@2015164218",
// })

// }

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {

	// s, sep := "", ""
	// for _, arg := range os.Args[0:] {
	// 	fmt.Println(os.Args)
	// 	s += sep + arg
	// 	sep = " "
	// }

	// fmt.Println(s)
	// fmt.Println(os.Args[1:])
	lissajous(os.Stdout)

	// counts := make(map[string]int)

	// files := os.Args[1:]

	// if len(files) == 0 {
	// 	countLines(os.Stdin, counts)
	// } else {
	// 	for _, arg := range files {
	// 		f, err := os.Open(arg)
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	// 			continue
	// 		}
	// 		countLines(f, counts)
	// 		f.Close()
	// 	}
	// }

	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for range nframes {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}

// }

// username := "Hcore"
// password := "tony@2015164218"

// counts := make(map[int16] string  ,int, );
// counts[]

// if value.name == "" {
// 	fmt.Println("no user name input ")
// 	return false
// }
// if value.password == "" {
// 	fmt.Println("no password input")
// 	return false

// }
// if value.username != username {
// 	fmt.Println("incorrect username")
// 	return false
// }
// if value.password != password {
// 	fmt.Println("wrong password")
// }
// fmt.Println("sucesful login")
// fmt.Println("welcome", value.name)

// return true
