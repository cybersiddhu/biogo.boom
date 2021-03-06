package main

import (
	"fmt"
	"github.com/cybersiddhu/biogo.boom"
)

func main() {
	boom.Verbosity(0)
	sf, err := boom.OpenSAM("test.sam", "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s %d\n%s", sf.RefNames(), sf.RefLengths(), sf.Text())
	for {
		r, _, err := sf.Read()
		if err != nil {
			break
		}
		//fmt.Printf("%d %s %d-%d (%d) %v %d %d %s %v Fl:%v %q\n", r.RefID(), r.Name(), r.Start(), r.End(), r.Len(), r.Cigar(), r.Strand(), r.Score(), r.Seq(), r.Quality(), r.Flags(), r.Tags())
		fmt.Printf("%d %s %d\n",r.RefID(),r.Name(),r.Score())
	}
}
