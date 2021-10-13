
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"github.com/pipe4/lang/pipe4/gogll/parser/symbols"
)

type Label int

const(
	GoGLL0R0 Label = iota
	GoGLL0R1
	Import0R0
	Import0R1
	Import0R2
	ImportBlock0R0
	ImportBlock0R1
	ImportBlock0R2
	ImportBlock0R3
	ImportBlock0R4
	ImportName0R0
	ImportName0R1
	ImportPath0R0
	ImportPath0R1
	ImportStatements0R0
	ImportStatements0R1
	ImportStatements1R0
	ImportStatements1R1
	ImportStatements1R2
)

type Slot struct {
	NT      symbols.NT
	Alt     int
	Pos     int
	Symbols symbols.Symbols
	Label 	Label
}

type Index struct {
	NT      symbols.NT
	Alt     int
	Pos     int
}

func GetAlternates(nt symbols.NT) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt symbols.NT, alt, pos int) Label {
	l, exist := slotIndex[Index{nt,alt,pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() symbols.NT {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() symbols.Symbols {
	return l.Slot().Symbols
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

var slots = map[Label]*Slot{ 
	GoGLL0R0: {
		symbols.NT_GoGLL, 0, 0, 
		symbols.Symbols{  
			symbols.NT_ImportBlock,
		}, 
		GoGLL0R0, 
	},
	GoGLL0R1: {
		symbols.NT_GoGLL, 0, 1, 
		symbols.Symbols{  
			symbols.NT_ImportBlock,
		}, 
		GoGLL0R1, 
	},
	Import0R0: {
		symbols.NT_Import, 0, 0, 
		symbols.Symbols{  
			symbols.NT_ImportName, 
			symbols.NT_ImportPath,
		}, 
		Import0R0, 
	},
	Import0R1: {
		symbols.NT_Import, 0, 1, 
		symbols.Symbols{  
			symbols.NT_ImportName, 
			symbols.NT_ImportPath,
		}, 
		Import0R1, 
	},
	Import0R2: {
		symbols.NT_Import, 0, 2, 
		symbols.Symbols{  
			symbols.NT_ImportName, 
			symbols.NT_ImportPath,
		}, 
		Import0R2, 
	},
	ImportBlock0R0: {
		symbols.NT_ImportBlock, 0, 0, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_0, 
			symbols.NT_ImportStatements, 
			symbols.T_1,
		}, 
		ImportBlock0R0, 
	},
	ImportBlock0R1: {
		symbols.NT_ImportBlock, 0, 1, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_0, 
			symbols.NT_ImportStatements, 
			symbols.T_1,
		}, 
		ImportBlock0R1, 
	},
	ImportBlock0R2: {
		symbols.NT_ImportBlock, 0, 2, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_0, 
			symbols.NT_ImportStatements, 
			symbols.T_1,
		}, 
		ImportBlock0R2, 
	},
	ImportBlock0R3: {
		symbols.NT_ImportBlock, 0, 3, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_0, 
			symbols.NT_ImportStatements, 
			symbols.T_1,
		}, 
		ImportBlock0R3, 
	},
	ImportBlock0R4: {
		symbols.NT_ImportBlock, 0, 4, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_0, 
			symbols.NT_ImportStatements, 
			symbols.T_1,
		}, 
		ImportBlock0R4, 
	},
	ImportName0R0: {
		symbols.NT_ImportName, 0, 0, 
		symbols.Symbols{  
			symbols.T_3,
		}, 
		ImportName0R0, 
	},
	ImportName0R1: {
		symbols.NT_ImportName, 0, 1, 
		symbols.Symbols{  
			symbols.T_3,
		}, 
		ImportName0R1, 
	},
	ImportPath0R0: {
		symbols.NT_ImportPath, 0, 0, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		ImportPath0R0, 
	},
	ImportPath0R1: {
		symbols.NT_ImportPath, 0, 1, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		ImportPath0R1, 
	},
	ImportStatements0R0: {
		symbols.NT_ImportStatements, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Import,
		}, 
		ImportStatements0R0, 
	},
	ImportStatements0R1: {
		symbols.NT_ImportStatements, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Import,
		}, 
		ImportStatements0R1, 
	},
	ImportStatements1R0: {
		symbols.NT_ImportStatements, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Import, 
			symbols.NT_ImportStatements,
		}, 
		ImportStatements1R0, 
	},
	ImportStatements1R1: {
		symbols.NT_ImportStatements, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Import, 
			symbols.NT_ImportStatements,
		}, 
		ImportStatements1R1, 
	},
	ImportStatements1R2: {
		symbols.NT_ImportStatements, 1, 2, 
		symbols.Symbols{  
			symbols.NT_Import, 
			symbols.NT_ImportStatements,
		}, 
		ImportStatements1R2, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_GoGLL,0,0 }: GoGLL0R0,
	Index{ symbols.NT_GoGLL,0,1 }: GoGLL0R1,
	Index{ symbols.NT_Import,0,0 }: Import0R0,
	Index{ symbols.NT_Import,0,1 }: Import0R1,
	Index{ symbols.NT_Import,0,2 }: Import0R2,
	Index{ symbols.NT_ImportBlock,0,0 }: ImportBlock0R0,
	Index{ symbols.NT_ImportBlock,0,1 }: ImportBlock0R1,
	Index{ symbols.NT_ImportBlock,0,2 }: ImportBlock0R2,
	Index{ symbols.NT_ImportBlock,0,3 }: ImportBlock0R3,
	Index{ symbols.NT_ImportBlock,0,4 }: ImportBlock0R4,
	Index{ symbols.NT_ImportName,0,0 }: ImportName0R0,
	Index{ symbols.NT_ImportName,0,1 }: ImportName0R1,
	Index{ symbols.NT_ImportPath,0,0 }: ImportPath0R0,
	Index{ symbols.NT_ImportPath,0,1 }: ImportPath0R1,
	Index{ symbols.NT_ImportStatements,0,0 }: ImportStatements0R0,
	Index{ symbols.NT_ImportStatements,0,1 }: ImportStatements0R1,
	Index{ symbols.NT_ImportStatements,1,0 }: ImportStatements1R0,
	Index{ symbols.NT_ImportStatements,1,1 }: ImportStatements1R1,
	Index{ symbols.NT_ImportStatements,1,2 }: ImportStatements1R2,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_GoGLL:[]Label{ GoGLL0R0 },
	symbols.NT_ImportBlock:[]Label{ ImportBlock0R0 },
	symbols.NT_ImportStatements:[]Label{ ImportStatements0R0,ImportStatements1R0 },
	symbols.NT_Import:[]Label{ Import0R0 },
	symbols.NT_ImportName:[]Label{ ImportName0R0 },
	symbols.NT_ImportPath:[]Label{ ImportPath0R0 },
}

