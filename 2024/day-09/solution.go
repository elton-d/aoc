package main

import (
	"fmt"
	"strconv"

	"github.com/elton-d/aoc/util"
)

func main() {
	input := util.GetInputStr("https://adventofcode.com/2024/day/9/input")

	fmt.Printf("Part 1: %d\n", checksum(compact(expand(input))))
	fmt.Printf("Part 2: %d\n", checksum(compact2(expand(input))))
}

type ListNode struct {
	idx    int
	blocks int
	next   *ListNode
}

func expand(input string) []string {
	var out []string
	var id int

	for i, c := range input {
		times, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			for range times {
				out = append(out, fmt.Sprintf("%d", id))
			}
			id++
		} else {
			for range times {
				out = append(out, ".")
			}
		}
	}
	return out
}

func compact(input []string) []string {
	var l, r int
	output := make([]string, len(input))
	copy(output, input)
	r = len(input) - 1

	for l < r {
		if output[l] != "." {
			l++
			continue
		}
		if output[r] == "." {
			r--
			continue
		}
		output[l] = output[r]
		output[r] = "."
		l++
		r--
	}
	return output
}

func compact2(input []string) []string {
	var head, tail *ListNode

	var k int
	for k = 0; k < len(input); {
		if input[k] == "." {
			start := k
			end := k
			for end = k; end < len(input); end++ {
				if input[end] != "." {
					end--
					break
				}
			}
			newNode := &ListNode{idx: start, blocks: end - start + 1}
			if head == nil {
				head = newNode
				tail = newNode
			} else {
				tail.next = newNode
				tail = newNode
			}
			k = end + 1
		} else {
			k++
		}
	}
	output := make([]string, len(input))
	copy(output, input)

	var srcEnd, srcStart int

	for r := len(input) - 1; r >= 0; {
		if input[r] != "." {
			srcEnd = r
			srcStart = r
			fileID := input[r]

			for i := r; i >= 0; i-- {
				if input[i] == fileID {
					srcStart = i
				} else {
					break
				}
			}

			fileSize := srcEnd - srcStart + 1
			prev := &ListNode{}
			for node := head; node != nil; node = node.next {
				if node.idx >= srcStart {
					break
				}
				if node.blocks >= fileSize {
					dst := node.idx
					for j := srcStart; j <= srcEnd; j++ {
						output[dst] = output[j]
						output[j] = "."
						dst++
					}
					node.blocks -= fileSize
					node.idx += fileSize
					if node.blocks == 0 {
						if prev.next == node {
							prev.next = node.next
						} else {
							head = node.next
						}
					}
					break
				}
				prev = node
			}
			r = srcStart - 1
		} else {
			r--
		}
	}

	return output
}

func checksum(input []string) int {
	var sum int
	for i, id := range input {
		if id == "." {
			continue
		}

		n, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		sum += i * n
	}
	return sum
}
