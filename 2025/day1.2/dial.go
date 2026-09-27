package main

import (
	//"fmt"
	"iter"
	"strconv"
)

type dial struct {
	pos *node
	seq iter.Seq[string]
}

type node struct {
	val   int
	right *node
	left  *node
}

func newDial(start, max int, seq iter.Seq[string]) *dial {
	return &dial{
		pos: buildNodes(start, max),
		seq: seq,
	}
}

func (d *dial) runSeqWatchZero() int {
	count := 0
	for dialMove := range d.seq {
		if len(dialMove) < 1 {
			continue
		}
		dir := dialMove[0]
		ticks := 0
		var e error
		if len(dialMove) > 1 {
			//fmt.Printf("dialMove=%q, sliced=%q\n", dialMove, dialMove[1:])
			ticks, e = strconv.Atoi(string(dialMove[1:]))
			check(e)
		}
		count += d.tickWatchZeroCount(rune(dir), ticks)
	}
	return count
}

func (d *dial) tickWatchZeroCount(direction rune, ticks int) int {
	count := 0
	for range ticks {
		if direction == 'R' {
			d.pos = d.pos.right
		} else {
			d.pos = d.pos.left
		}

		if d.pos.val == 0 {
			count++
		}
	}
	return count
}

func buildNodes(start, max int) *node {
	zeroNode := newNode(0)
	var startNode *node
	prev := zeroNode
	for i := 1; i <= max; i++ {
		curr := newNode(i)
		prev.right = curr
		curr.left = prev

		if i == start {
			startNode = curr
		}

		prev = curr
	}
	// prev is at 99 after loop
	prev.right = zeroNode
	zeroNode.left = prev
	return startNode
}

func newNode(val int) *node {
	return &node{
		val:   val,
		left:  nil,
		right: nil,
	}
}
