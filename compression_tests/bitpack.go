////////////////////////////////////////////////////////////////////////////////
// This is a pseudo bit packing program. It's terribly unoptimized, filled to
// the brim with needless conditionals, and can't even be used in any practical
// manner. It's just an excercise for myself in understanding how bitpacking
// works, where it's limitations are, and to generally get more experience
// programming in go. I will eventually revisit this and make it suck less.
//
// Bitpacking strategy:
// 1-bit header followed by a known number of bits representing a value
// With this strategy we are adding 1 bit to every value, so we must save more
// than 2 bits on half of our values to get any compression
////////////////////////////////////////////////////////////////////////////////
package main


import (
  "os"
  "fmt"
  "strconv"
  "encoding/hex"
)


// Take hex values from user, validate, then convert to slice stored bit lengths
func hexGetter() [][]int {
  var input string
  var bits []string

  fmt.Println("Input hex strings of 8 length. Ex \"0000FDEE\"")
  fmt.Println("When done, input \"done\"")
  for {
    fmt.Scanln(&input)
    if input == "done" {
      break
    }
    if len(input) != 8 {
      fmt.Println("Hex length must be 8, try again")
    } else { 
      _, err := hex.DecodeString(input)
      if err != nil { fmt.Println("Like, not a valid hex scoob"); os.Exit(0) }

      fromhex, err := strconv.ParseInt(input, 16, 64)
      if err != nil { fmt.Println("Something happened heeheehoo"); os.Exit(0) }
      fromint := fmt.Sprintf("%032b", fromhex) // convert to 32 bit binary value
      bits = append(bits, fromint)
    }
  }

  // Make and initialize bits_final slice based off user input
  bits_final := make([][]int, len(bits))
  for i := 0; i < len(bits); i++ {
    bits_final[i] = make([]int, len(bits[0]))
  }

  // Convert and add individual bits to the 2d slice bits_final
  i := 0
  for _, row := range bits {
    j := 0
    for _, bit := range row {
      value, _ := strconv.Atoi(string(bit))
      bits_final[i][j] = value
      j++
    }
    i++
  }

  fmt.Println("Unpacked bitvalues:")
  for _, row := range bits_final {
    for _, bit := range row {
      fmt.Print(bit)
    }
  }
  fmt.Println()

  global_unpack = len(bits_final)*len(bits_final[0])
  return bits_final
}


// Encoder: pack all values that have 16 leading 0s, leave all others as-is
func bitPacker(packerSource [][]int) []int {
  var packed []int

  // Count the amount of consecutive 0s in a particular bitlength
  for i := 0; i < len(packerSource); i++ {
    consecutive_0s := 0
    for j := 0 ; j < 31 ; j++ {
      if packerSource[i][j] == 0 && packerSource[i][j+1] == 0 {
        consecutive_0s++
      } else { break }
    }
    consecutive_0s++ // makes up for the first undocumented 0

    // Header value "1" assigned to bit lengths with 16 consecutive 0s, else "0"
    if consecutive_0s > 15 {
      packed = append(packed, 1)
      for j := 16 ; j < 32 ; j++ {
        packed = append(packed, packerSource[i][j])
      }
    } else {
      packed = append(packed, 0)
      for j := 0 ; j < 32 ; j++ {
        packed = append(packed, packerSource[i][j])
      }
    }
  }

  fmt.Println(); fmt.Println("Packed bitvalues:")
	for _, val := range packed {
	  fmt.Print(val)
  }
  fmt.Println(); fmt.Println()

  global_packed = len(packed)
  return packed
}


// Decoder: read the header value, take action, append to depacked slice
func bitDepacker(unpackerSource []int) []int {
  var depacked []int

  // Track unpacked location with i, read header values and adjust location 
  for i := 0 ; i < len(unpackerSource) ; {
    if unpackerSource[i] == 1 {
      depacked = append(depacked, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
      for j := i + 1 ; j < i + 17 ; j++ {
        depacked = append(depacked, unpackerSource[j])
      }
      i = i + 17
    } else {
      for j := i + 1 ; j < i + 33; j++ {
        depacked = append(depacked, unpackerSource[j])
      }
      i = i + 33
    }
    if i == len(unpackerSource) { break }
  }

  fmt.Println("Depacked bitvalues:")
  for _, val := range depacked {
    fmt.Print(val)
  }
  fmt.Println()

  return depacked
}


var global_unpack int
var global_packed int

func main() {
  bitvals := hexGetter()
  packed := bitPacker(bitvals)
  bitDepacker(packed)

  // Should make this its own function but global variables work
  fmt.Println("\nUnpacked bit size:", global_unpack)
  fmt.Println("Packed bit size:", global_packed)
  fmt.Println("You saved", global_unpack - global_packed, "bits")
}
