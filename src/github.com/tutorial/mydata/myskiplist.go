package mydata

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
    maxLevel int     = 6
    p        float32 = 0.5
    Esp      float64 = 0.0000001
)

func RandomLevel() int {
    level := 1
    rand.Seed(time.Now().UnixNano())
    for rand.Float32() < p && level < maxLevel {
        level++
    }
    return level
}

func IsEqual(a, b float64) bool {
    return math.Abs(a-b) < Esp
}

// Element is an Element of a skiplist.
type Element struct {
    Score   float64
    Value   interface{}
    next []*Element
}

func NewElement(score float64, value interface{}, level int) *Element {
    return &Element{
        Score:   score,
        Value:   value,
        next: make([]*Element, level),
    }
}

// Next returns first element after e.
func (e *Element) Next() *Element {
    if e != nil {
        return e.next[0]
    }
    return nil
}

// SkipList represents a skiplist.
// The zero value from SkipList is an empty skiplist ready to use.
type SkipList struct {
    Header *Element // header is a dummy element
    Len    int      // 长度，不包括头节点, 
    Level  int      // 高度，暂时只增不减
}

// New returns a new empty SkipList.
func NewSkipList() *SkipList {
    return &SkipList{Header: &Element{}}
}

func (s *SkipList) Front() *Element {
    return s.Header.next[0]
}

// Search the skiplist to findout element with the given score.
// Returns (*Element, true) if the given score present, otherwise returns (nil, false).
func (s *SkipList) Search(score float64) (*Element, bool) {
    cur := s.Header
    // 从上到下
    for i:=s.Level-1; i>=0; i-- {
        // 从左到右
        for cur.next[i] != nil && cur.next[i].Score < score {
            cur = cur.next[i]
        }
    }
    if cur.next == nil {
        return nil, false
    }

    cur = cur.next[0]
    if cur != nil && IsEqual(cur.Score, score) {
        return cur, true
    }
    return nil, false
}

// Insert (score, value) pair to the skiplist and returns pointer of element.
func (s *SkipList) Insert(score float64, value interface{}) *Element {
    ele, exist := s.Search(score)
    if exist {
        ele.Value = value
        return ele
    }

    start := len(s.Header.next)
    level := RandomLevel()

    fmt.Printf("Insert: score=%.1f, value=%v, level=%d\n", score, value,level)
    ele = NewElement(score, value, level)
    for s.Level < level {
        s.Header.next = append(s.Header.next, ele)
        s.Level ++
    }

    if level < start {
        start = level
    }

    cur := s.Header
    for i:=start-1; i>=0; i-- {
        for cur.next[i] != nil && cur.next[i].Score < score {
            cur = cur.next[i]
        }
        ele.next[i] = cur.next[i]
        cur.next[i] = ele
    }
    
    s.Len++
    return ele
}

// Delete remove and return element with given score, return nil if element not present
func (s *SkipList) Delete(score float64) *Element {
    ele, exist := s.Search(score)
    if !exist {
        return nil
    }

    cur := s.Header
    for i:=s.Level-1; i>=0; i-- {
        for cur.next[i] != nil && cur.next[i].Score < score {
            cur = cur.next[i]
        }
        if cur.next[i] != nil && IsEqual(cur.next[i].Score, score) {
            cur.next[i] = ele.next[i]
        }
    }

    s.Len--
    return ele
}

func (s *SkipList) Scan() {
    // 记录节点顺序, 包含头节点
    index := make([]*Element, 0, s.Len+1)
    cur := s.Header

    fmt.Printf("level: %d =>", 1)
    for cur.next[0] != nil {
        fmt.Printf("[s=%.1f, v=%v, h=%d] -> ", cur.Score, cur.Value, len(cur.next))
        index = append(index, cur)
        cur = cur.next[0]
    }
    fmt.Printf("[s=%.1f, v=%v, h=%d]。\n", cur.Score, cur.Value, len(cur.next))
    index = append(index, cur)
    
    for i:=1; i<s.Level; i++ {
        cur := s.Header
        if cur.next[i] == nil {
            continue
        }
        fmt.Printf("level: %d =>", i+1)
        
        j:=0;
        for ; cur.next[i] != nil; j++ {
            e := index[j]
            if IsEqual(e.Score, cur.Score) {
                fmt.Printf("[s=%.1f, v=%v, h=%d] -> ", cur.Score, cur.Value, len(cur.next))
            } else {
                fmt.Printf("[===============] -> ")
                continue
            }
            cur = cur.next[i]
        }
        // 最后一个节点处理
        for ; j<s.Len+1; j++ {
            e := index[j]
            if IsEqual(e.Score, cur.Score) {
                fmt.Printf("[s=%.1f, v=%v, h=%d]。\n", cur.Score, cur.Value, len(cur.next))
                break
            } else {
                fmt.Printf("[***************] -> ")
            }
        }
    }

}

