package display

import (
	"sort"
	"strconv"
	"strings"
)

func New() *Display {
	return &Display{}
}

type Display struct {
	mapping map[int]string
}

/*
0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
*/

//revive:disable:cognitive-complexity
//revive:disable:cyclomatic

func (d *Display) Calibrate(inChunks []string) {
	mapping := make(map[int]string) // this ister claibraation
	// Do the easy ones
	for _, chunk := range inChunks {
		switch len(chunk) {
		case 2: // Figure 1
			mapping[1] = chunk
		case 3: // Figure 7
			mapping[7] = chunk
		case 4: // Figure 4
			mapping[4] = chunk
		case 7: // figure 8
			mapping[8] = chunk
		}
	}
	// Now fill in the gaps
	for _, chunk := range inChunks {
		switch len(chunk) {
		case 5: // 2, 3, 5
			// contains 7?
			if overlap(chunk, mapping[7]) {
				mapping[3] = chunk
			} else {
				if overlapCount(chunk, mapping[4]) == 3 {
					mapping[5] = chunk
				} else {
					mapping[2] = chunk
				}
			}
		case 6: // 0, 6, 9
			// contains 4?
			if overlap(chunk, mapping[4]) {
				mapping[9] = chunk
			} else {
				if overlap(chunk, mapping[7]) {
					mapping[0] = chunk
				} else {
					mapping[6] = chunk
				}
			}
		}
	}
	d.mapping = mapping
}

//revive:enable:cognitive-complexity
//revive:enable:cyclomatic

func (d *Display) Decode(encoded []string) int {
	numerals := []int{}
	for _, code := range encoded {
		for numeral, val := range d.mapping {
			if overlapExact(val, code) {
				numerals = append(numerals, numeral)
			}
		}
	}
	return toNumber(numerals)
}

func toNumber(input []int) int {
	numeralsString := []string{}
	for _, val := range input {
		str := strconv.Itoa(val)
		numeralsString = append(numeralsString, str)
	}
	ret, err := strconv.Atoi(strings.Join(numeralsString, ""))
	if err != nil {
		panic("not expected")
	}
	return ret
}

func overlapCount(chunk, subChunk string) int {
	chunkParts := strings.Split(chunk, "")
	subChunkParts := strings.Split(subChunk, "")
	overlapCount := 0
	for _, sub := range subChunkParts {
		for _, chunk := range chunkParts {
			if sub == chunk {
				overlapCount++
			}
		}
	}
	return overlapCount
}

func overlapExact(chunk, subChunk string) bool {
	chunkParts := strings.Split(chunk, "")
	subChunkParts := strings.Split(subChunk, "")
	sort.Strings(chunkParts)
	sort.Strings(subChunkParts)
	return strings.Join(chunkParts, "") == strings.Join(subChunkParts, "")
}

func overlap(chunk, subChunk string) bool {
	chunkParts := strings.Split(chunk, "")
	subChunkParts := strings.Split(subChunk, "")
	for _, sub := range subChunkParts {
		match := false
		for _, chunk := range chunkParts {
			if sub == chunk {
				match = true
			}
		}
		if !match {
			return false
		}
	}
	return true
}
