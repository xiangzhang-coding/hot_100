# median of two sorted arrays

## 方法一

把两个数组排序，然后根据 $m\plus n$ 的奇偶找出中位数。总的时间复杂度为 O(nlog(n)),空间复杂度为 O(n)。

## 方法二

把两个数组归并排序 $k$ 次，其中 $k = \frac{m+n}{2}$。时间复杂度比方法一好一点为 O(n)，空间复杂度为 O(n)

## 方法三
