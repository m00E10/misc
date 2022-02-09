////////////////////////////////////////////////////////////////////////////////
// So this will be a kind of pseudo bit pack. It will not be ready for
// practical application, this is more just an exercise in understanding how
// bitpacking actually works.
//
// values to bitpack:
// 00000000 00000000 01111101 11101110   16 bits
// 00000000 00000000 00000000 10100001    8 bits
// 00000010 10100011 00010010 00001110   25 bits
//
// Bitpacking strategy:
// 1-bit header followed by a known number of bits representing a value
// With this strategy we are adding 1 bit to every value, so we must save more
// than 2 bits on half of our values to get any compression
//
// Todo:
// X Pseudocode
// Implementation in Golang, linearly
// Revise by seperating things into functions
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

func main() {
  // The three pseudo 32bit values we will be packing
  bitvals := [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                      0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0},
                     {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                      0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
                     {0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1,
                      0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1}}

  fmt.Println("Unpacked bitvalues:")
  for i := 0 ; i < 3 ; i++ {
  	for j := 0 ; j < 32 ; j++ {
    	fmt.Print(bitvals[i][j])
    }
    fmt.Println()
  }
  fmt.Println()
  
  // Pack all values that have 16 leading 0s, place all others alongside
  var packed []int
  for i := 0 ; i < 3 ; i++ {
    consecutive_0s := 0
    for j := 0 ; j < 31 ; j++ {
      if bitvals[i][j] == 0 && bitvals[i][j+1] == 0 {
        consecutive_0s++
      }
    }
    consecutive_0s++ // makes up for the first undocumented 0
    fmt.Println("Consecutive Zeros:", consecutive_0s)

    if consecutive_0s > 15 {
      packed = append(packed, 1)
      for j := 16 ; j < 32 ; j++ {
        packed = append(packed, bitvals[i][j])
      }
    } else {
      packed = append(packed, 0)
      for j := 0 ; j < 32 ; j++ {
        packed = append(packed, bitvals[i][j])
      }
    }

  }
  fmt.Println()
  fmt.Println("Packed bitvalues:")
	for _, val := range packed {
	  fmt.Print(val)
  }
  fmt.Println()

  var depacked []int
  fmt.Println(); fmt.Println("Packed length", len(packed)); fmt.Println();
  for i := 0 ; i < len(packed) ; {
    fmt.Println("Value of i", i)
    if packed[i] == 1 {
      depacked = append(depacked, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
      i++
      for j := i ; j < i + 15 ; j++ {
        depacked = append(depacked, packed[j])
      }
      fmt.Println(depacked)
    } else {
      i++
      for j := i ; j < i + 32; j++ {
        depacked = append(depacked, packed[j])
      }
      fmt.Println(depacked)
    }
  }
  fmt.Println()
  fmt.Println("Depacked bitvalues:")
  for _, val := range depacked {
    fmt.Print(val)
  }
  fmt.Println()

}
// DECODER
// for (int j=0; j<length(packed)) {
//   if (packed[j] == 1) {
//      echo "0000000000000000" >> unpacked
//      for (int i=j+1; i<j+16; i++) {
//        packed[i] >> unpacked
//      }
//   }
//   if(packed[j] == 0) {
//      for (int i=j+1; i<j+32; i++) {
//        packed[i] >> unpacked
//      }
//   }
// }
// print("Unpacked values:\n");
// print(unpacked);
