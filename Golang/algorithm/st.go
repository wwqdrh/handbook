package algorithm

import "math/bits"

func sumSubarrayMins(arr []int) int {
    st := NewST(arr)

    mod := int(1e9)+7
    res := 0
    n := len(arr)
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            res = (res + st.Query(i, j)) % mod
        }
    }
    return res
}

type ST [][]int

func NewST(nums []int) ST {
    numsLen := len(nums)
    numsLenBit := bits.Len(uint(numsLen))
    st := make(ST, numsLen)
    for i := range st {
        st[i] = make([]int, numsLenBit+1)
        st[i][0] = nums[i]
    }
    for j := 1; j << 1 <= numsLen; j++ {
        for i := 0; i + (j << 1) <= numsLen; i++ {
            st[i][j] = st.Op(st[i][j-1], st[i+(1<<(j-1))][j-1])
        }
    }
    return st
}

func (s ST) Query(l, r int) int {
    if l == r {
        return s[l][0]
    }
    k := bits.Len(uint(r-l)) - 1
    return s.Op(s[l][k], s[r-(1<<k)][k])
}

func (s ST) Op(a, b int) int {
    if a <= b {
        return a
    }
    return b
}