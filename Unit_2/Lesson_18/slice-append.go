// slice-appeng.go
package main

import "fmt"

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Heumea", "Makemake", "Eris"}
	dump("dwarfs1", dwarfs1)
	dwarfs2 := append(dwarfs1, "Orcus")
	dump("dwarfs2", dwarfs2)
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs3", dwarfs3)
	dwarfs3[1] = "Pluto!"
	/* dwarfs3 and dwarfs2 are changed
	 * but dwarfs1 is not because it points
	 * at another array */
	dump("dwarfs3 changed", dwarfs3)
	dump("dwarfs2 changed", dwarfs2)
	dump("dwarfs1 changed", dwarfs1)
}
