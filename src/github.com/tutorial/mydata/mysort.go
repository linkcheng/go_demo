package mydata

import "fmt"

// 升序

// 冒泡排序
// 冒泡排序：一次冒泡过程，只能把一个元素放到正确位置。
// 实现：相邻的两个元素比较一下，如果 s [i]>s [i+1]，替换位置。
// 最多需要冒泡 n 次；
func Bubble(s []int) {
	length := len(s)

	if length < 2 {
		return 
	}

	for i:=0; i<length; i++ {
		for j:=0; j<length-1; j++ {
			// 前一个 > 后一个，则交换
			if s[j] > s[j+1] {
                s[j], s[j+1] = s[j+1], s[j]
            }
		}
	}
}

// 插入排序
// 就是把数据分成两部分，左边部分为排序好的，右边部分为未排序的。
// 每次从右边取出一个元素 s[j], 和左边排序好的元素依次比较。
// 然后合适的地方插入。
func Insert(s []int) {
	length := len(s)

	if length < 2 {
		return 
	}

	for i:=1; i<length; i++ {
		j := i-1
		val := s[i]
		for ; j>=0; j-- {
			// 待排序值 < 当前有序值
			if val < s[j] {
				// 有序向右移动
				s[j+1] = s[j]
			} else {
				break
			}
		}
		// 放入0位置 j=-1，
		// 或者放入当前有序值后一个位置， val>=s[j]
		s[j+1] = val
	}
}

// 归并排序
// 如果要排序一个数组，我们先把数组从中间分成前后两部分， 然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了。
// 复杂度是 O(nlogn)，非原地排序，比较耗空间。
func MergeSort(s []int) []int {
	length := len(s)
	if length < 2 {
		return s
	}

	return mergeSort(s)
}


func mergeSort(s []int) []int {
	length := len(s)
	if length < 2 {
        return s //最后切割只剩下一个元素
    }

	mid := length / 2
	left := mergeSort(s[:mid])
	right := mergeSort(s[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	lLen := len(left)
	rLen := len(right)
	res := make([]int, 0, lLen+rLen)

	lIdx, rIdx := 0, 0
	for lIdx < lLen && rIdx < rLen {
		if left[lIdx] < right[rIdx] {
			res = append(res, left[lIdx])
			lIdx ++
		} else {
			res = append(res, right[rIdx])
			rIdx ++
		}
	}

	if lIdx < lLen {
		res = append(res, left[lIdx:]...)
	}
	if rIdx < rLen {
		res = append(res, right[rIdx:]...)
	}

	return res
}


// 快排
// 快速排序是原地排序，就是在 s 数组中不断交换位置。
// 也是分治思想，选中一个标尺，比如第一个元素。然后把其他元素依次和他比较，把小于它的元素放左边，大于它的元素放右边
// 递归处理左边部分，右边部分
func QuickSort(s []int) {
	length := len(s)
	if length < 2 {
		return
	}
	quickSort(s)
}

func quickSort(s []int) {
	length := len(s)
	if length < 2 {
		return
	}

	value := s[0]
	// 左边区域 < value, 右边区域 > value
	left, right := 0, length-1
	
	for left < right {
		cur := left+1
		if s[cur] < value {
			s[left], s[cur] = s[cur], s[left]
			left ++
		} else if s[cur] > value {
			s[right], s[cur] = s[cur], s[right]
			right --
		} else {
			left ++
		}
	}
	// fmt.Printf("left=%d, right=%d\n", left, right)

	quickSort(s[:left])
	quickSort(s[left+1:])
}

// 二分法通过对折的方式查找一个数据，条件是必须是一个有序的数组。
// 数组是顺序，可以随机获取一个值 O(log n).
func BinarySearch(target int, s []int) int {
	length := len(s)
	if length == 0 {
		return -1
	}
	left, right := 0, length-1
	for left <= right {
		mid := (left+right) / 2
		if target < s[mid] {
			right = mid - 1
		} else if target > s[mid] {
			left = mid + 1
		} else {
			return s[mid]
		}
	}

	return -1
}

// 堆排序, 递增
func HeapSort(s []int) {
	length := len(s)
	if length < 2 {
		return
	}

	// 构建大根堆
	fmt.Printf("===============%v==============\n", s)
	for i:=0; i<length; i++ {
		heapInsert(s, i)
	}
	fmt.Printf("###############%v##############\n", s)
	
	// 一次把堆顶位置的值与堆中最后一个值交换，
	// 通过 heapify 继续保持堆结构
	for r:=length-1; r>0; r-- {
		s[0], s[r] = s[r], s[0]
		heapify(s, 0, r)		
	}
}

// 大根堆
func heapInsert(s []int, i int) {
	// parent 位置
	p := (i-1) >> 1
	// fmt.Printf("p=%d\n", p)
	for p >=0 && s[i] > s[p] {
		// fmt.Printf("s=%d, p=%d, s[i]=%d, s[p]=%d\n", i, p, s[i], s[p])
		s[p], s[i] = s[i], s[p]
		i = p
		p = (i-1) >> 1
	}
}

// 保持以 root 为根节点的大根堆在 size 范围内依然是大根堆结构
func heapify(s []int, root, size int) {
	left := (root << 1) + 1
	for left < size {
		largest := left
		right := left + 1
		if right < size && s[right] > s[left] {
			largest = right 
		} 

		if s[root] > s[largest] {
			break
		}
		
		s[largest], s[root] = s[root], s[largest]
		root = largest
		left = (root << 1) + 1
	}
}
