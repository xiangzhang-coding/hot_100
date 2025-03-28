# median of two sorted arrays

## 方法一

把两个数组排序，然后根据 $m + n$ 的奇偶找出中位数。总的时间复杂度为 O( n log( n ) ),空间复杂度为 O(n)。

## 方法二

把两个数组归并排序 $k$ 次，其中 $k = \frac{m+n}{2}$。时间复杂度比方法一好一点为 O(n)，空间复杂度为 O(n)

## 方法三

首先的重点是找出中位数的技巧，若两个数组长度分别为 $m$和$n$，则他们的中位数的位置为第$\frac{(m+n+1)}{2}$和$\frac{(m+n+2)}{2}$ 个数。
其次使用了分治的思想，每次去掉数组中的一部分数（若要求第k个数，先去除k/2个数）直至剩余找k=1时的情况。
再次就是注意边界情况
