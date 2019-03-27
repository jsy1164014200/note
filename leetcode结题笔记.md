# leetcode结题笔记

## 方法

1. 递归



## 2. 链表两数相加

my code

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    incrementCount := 0
    for true {
        if l1 == nil && l2 == nil && incrementCount ==0 {
            break
        }
        
        a := 0 
        if l1 != nil {
            a = l1.Val
            l1 = l1.Next
        }
        b := 0
        if l2 != nil {
            b = l2.Val
            l2 = l2.Next
        }
        
        x := &ListNode{(a+b+incrementCount) % 10, nil}
        addListNode(result, x)
       
        
        if a + b + incrementCount >= 10 {
            incrementCount = 1
        } else {
            incrementCount = 0
        }
        
    }    
    return result.Next
}


func addListNode(l *ListNode,t *ListNode) {
    if l == nil {
        l = t
        return 
    }
    x := l
    for true {
        if x.Next == nil {
            x.Next = t
            break
        }
        x = x.Next
    }
}
```

your code

```c++
class Solution {
public:
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
        ListNode *res = new ListNode(-1);
        ListNode *cur = res;
        int carry = 0;
        while (l1 || l2) {
            int n1 = l1 ? l1->val : 0;
            int n2 = l2 ? l2->val : 0;
            int sum = n1 + n2 + carry;
            carry = sum / 10;
            cur->next = new ListNode(sum % 10);
            cur = cur->next;
            if (l1) l1 = l1->next;
            if (l2) l2 = l2->next;
        }
        if (carry) cur->next = new ListNode(1);
        return res->next;
    }
};
```

## *3. 无重复的最长子串

滑动窗口是数组/字符串问题中常用的抽象概念。 窗口通常是在数组/字符串中由开始和结束索引定义的一系列元素的集合，即 [i, j)[i,j)（左闭，右开）。而滑动窗口是可以将两个边界向某一方向“滑动”的窗口。例如，我们将 [i, j)[i,j) 向右滑动 11 个元素，则它将变为 [i+1, j+1)[i+1,j+1)（左闭，右开）。

```go
type Slide struct {
	Begin int
	End   int
}

func lengthOfLongestSubstring(s string) int {
	slide := &Slide{0, 0}
	for slide.End < len(s) {
		for i := slide.End; i < len(s); i++ {
			if isUnique(s[slide.Begin:slide.End]) && !strings.Contains(s[slide.Begin:slide.End], s[i:i+1]) {
				slide.End++
			} else {
				slide.Begin++
				slide.End++
				break
			}
		}
	}
	return slide.End - slide.Begin
}

func isUnique(s string) bool {

	for index, _ := range s {
		if strings.Count(s, s[index:index+1]) > 1 {
			return false
		}
	}
	return true
}
```



## 4. 寻找两个有序数组中的中位数



```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	c := len(nums1) + len(nums2)
	arr := make([]int, c)

	i := 0
	j := 0
	k := 0
	for i < c {
		if j >= len(nums1) {
			arr[i] = nums2[k]
			i++
			k++
			continue
		}
		if k >= len(nums2) {
			arr[i] = nums1[j]
			i++
			j++
			continue
		}
		if nums1[j] <= nums2[k] {
			arr[i] = nums1[j]
			i++
			j++
		} else {
			arr[i] = nums2[k]
			i++
			k++
		}
	}

	if c%2 == 0 {
		return (float64(arr[c/2]) + float64(arr[(c-2)/2])) / 2
	} else {
		return float64(arr[(c-1)/2])
	}

}
```



## *5. 最长回文子串

1. 最长公共子串

我们可以看到，当 SS 的其他部分中存在非回文子串的反向副本时，最长公共子串法就会失败。为了纠正这一点，每当我们找到最长的公共子串的候选项时，都需要检查子串的索引是否与反向子串的原始索引相同。如果相同，那么我们尝试更新目前为止找到的最长回文子串；如果不是，我们就跳过这个候选项并继续寻找下一个候选。

2. 中心扩展算法

   2n - 1 个中心

```go
func longestPalindrome(s string) string {
    if s == "" {
		return ""
	}
	start := 0        // s开始的索引
	end := len(s) - 1 // s结束的索引
	startIndex := 0
	endIndex := 0
	resultStart := 0 // 	记录 最长回文子串的 开始索引
	resultEnd := 0   //  记录 最长回文子串的 结束索引
	for index, _ := range s {
		resultStart = index - 1
		resultEnd = index + 1
		for resultStart >= start && resultEnd <= end {
			if s[resultStart] == s[resultEnd] {
				resultStart--
				resultEnd++
			} else {
				break
			}
		}
		if resultEnd-resultStart-2 > endIndex-startIndex {
			endIndex = resultEnd - 1
			startIndex = resultStart + 1
		}

		resultStart = index
		resultEnd = index + 1
		for resultStart >= start && resultEnd <= end {
			if s[resultStart] == s[resultEnd] {
				resultStart--
				resultEnd++
			} else {
				break
			}
		}
		if resultEnd-resultStart-2 > endIndex-startIndex {
			endIndex = resultEnd - 1
			startIndex = resultStart + 1
		}
	}

	return s[startIndex : endIndex+1]
}
```



## 6. Z字形变换

找 规律

```go
func convert(s string, numRows int) string {
	totalCount := len(s)
	result := ""
	index := 0
	for i := 0; i < numRows; i++ {
		// 最后一个是 2*n - 2
		// 从 i 开始 到
		index = i
		if i == 0 || i == numRows-1 {
			for index < totalCount {
				result = result + s[index:index+1]
				index += 2*numRows - 2
			}
		} else {
			for index < totalCount {
				result = result + s[index:index+1]
				index += 2*numRows - 2 - 2*i
				if index >= totalCount {
					break
				}
				result = result + s[index:index+1]
				index += 2 * i
			}
		}

	}

	return result
}
```

## 7. 整数反转

```go
func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}

	end := 0
	result := []int{}
	temp := 1
	for true {
		if x/temp%10 == 0 && x/temp == 0 {
			break
		}
		result = append(result, x/temp%10)
		temp *= 10
	}

	length := len(result)
	for index, value := range result {
		for i := 0; i < length-1-index; i++ {
			value *= 10
		}
		end += value
	}
	if end > 2147483647 {
		return 0
	}
	return end
}
```



## 8. 

```go

```

