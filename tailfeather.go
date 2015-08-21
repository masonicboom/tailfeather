// Tailfeather prints fields of delimited input in colors based on their values.
//
// The intended purpose is to aid visual inspection of log files, by coloring
// repeated values consistently, making patterns and anomalies easier to spot.

// Overall Strategy: for each field, track the last N unique values observed,
// and assign a different color to each one. When a new value is observed,
// assign it the least-recently assigned color.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var inputDelimiter string
var outputDelimiter string

func init() {
	flag.StringVar(&inputDelimiter, "input-delimiter", " ", "Input field delimiter.")
	flag.StringVar(&outputDelimiter, "output-delimiter", "\t", "Output field delimiter.")
}

// These are the colors available for use.
var colors = []ct.Color{
	ct.White,
	ct.Magenta,
	ct.Cyan,
	ct.Blue,
	ct.Green,
	ct.Yellow,
	ct.Red,
}

// MaxValues is limited to the number of distinct colors that can be used.
const maxValues = 7

// Each Field is effectively a FIFO queue of color indices.
type Field struct {
	ValueIndex         map[string]int    // Map from value to its index; used to determine color of a value.
	Values             [maxValues]string // Map from index to its value; used to evict old values from the ValueIndex.
	NextIndex          int               // The next index to evict.
	PrevDisplayedColor ct.Color          // The previously-displayed color; used to avoid repetition.
}

func NewField() *Field {
	return &Field{
		ValueIndex:         make(map[string]int),
		Values:             [maxValues]string{},
		NextIndex:          0,
		PrevDisplayedColor: colors[len(colors)-1], // This is just chosen to not be the first color.
	}
}

func main() {
	flag.Parse()

	// Reset console color if interrupted.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		ct.ResetColor()
		fmt.Println()
		os.Exit(0)
	}()

	// The fields data structure tracks coloration info for each field.
	var fields map[int]*Field
	var numFields = 0

	// Process input lines one-by-one, analyzing field format and coloring output based on the fields data structure.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()
		lineFields := strings.Split(l, inputDelimiter)

		// If a new line format is encountered, re-initialize coloration state.
		if len(lineFields) != numFields {
			n := len(lineFields)

			fields = make(map[int]*Field)
			for i := 0; i < n; i++ {
				fields[i] = NewField()
			}

			numFields = n
		}

		// Output each field in the appropriate color, separated by outputDelimiter.
		for fi, v := range lineFields {
			f := fields[fi]

			// Determine the color index of the current field value,
			// if it has been seen before.
			vi, ok := f.ValueIndex[v]
			c := colors[vi]

			// When a novel value is encountered, choose its color
			// by putting it in the FIFO queue for this field.
			if !ok {
				vi = f.NextIndex
				c = colors[vi]

				// Avoid repeating the last color used.
				if c == f.PrevDisplayedColor {
					vi = (f.NextIndex + 1) % maxValues
					c = colors[vi]
				}

				// Evict previous value, if necessary.
				prevValue := f.Values[vi]
				delete(f.ValueIndex, prevValue)

				// Put the new value into the FIFO queue.
				f.Values[vi] = v
				f.ValueIndex[v] = vi
				f.NextIndex = (vi + 1) % maxValues
			}

			// Print the value in the appropriate color.
			ct.ChangeColor(c, false, ct.None, false)
			fmt.Print(v)

			// Record the color that was used, to avoid repetition.
			// If a new color is cycled in on the next row, it is
			// possible to repeat the same color that was shown on
			// this row, if that color is "up to bat". Recording
			// the current color lets us handle that case.
			f.PrevDisplayedColor = c

			// Follow the value with the outputDelimiter, if there are more fields.
			if fi < numFields-1 {
				fmt.Print(outputDelimiter)
			}
		}
		fmt.Println()
	}

	ct.ResetColor()
}
