# 前后缀和 and 最值

## 前缀和/后缀和、前缀/后缀最值

设数组为 $a[0], a[1], \ldots, a[n-1]$。

---

### 1. 前缀和

$$\text{prefix}[i] = a[0] + a[1] + \cdots + a[i]$$

$$\text{prefix}[i] = \text{prefix}[i-1] + a[i]$$

**查询区间和** $[l, r]$：

$$\sum_{k=l}^{r} a[k] = \text{prefix}[r] - \text{prefix}[l-1]$$

（约定 $\text{prefix}[-1] = 0$）

**构建：**

```
prefix[0] = a[0]
for i := 1 to n-1:
    prefix[i] = prefix[i-1] + a[i]
```

---

### 2. 后缀和

$$\text{suffix}[i] = a[i] + a[i+1] + \cdots + a[n-1]$$

$$\text{suffix}[i] = \text{suffix}[i+1] + a[i] $$

**查询区间和** $[l, r]$：

$$\sum_{k=l}^{r} a[k] = \text{suffix}[l] - \text{suffix}[r+1]$$

（约定 $\text{suffix}[n] = 0$）

**构建：**

```
suffix[n-1] = a[n-1]
for i := n-2 downto 0:
    suffix[i] = suffix[i+1] + a[i]
```

---

### 3. 前缀最值

$$\text{prefixMax}[i] = \max(a[0], a[1], \ldots, a[i])$$

$$\text{prefixMax}[i] = \max(\text{prefixMax}[i-1], a[i])$$

**用途**：$O(1)$ 查询 $a[0..i]$ 中的最大值。

**构建：**

```
prefixMax[0] = a[0]
for i := 1 to n-1:
    prefixMax[i] = max(prefixMax[i-1], a[i])
```

**典型场景**：买卖股票——遍历卖出日时，用前缀最小值获取之前的最低买入价。

---

### 4. 后缀最值

$$\text{suffixMax}[i] = \max(a[i], a[i+1], \ldots, a[n-1])$$

$$\text{suffixMax}[i] = \max(\text{suffixMax}[i+1], a[i])$$

**用途**：$O(1)$ 查询 $a[i..n-1]$ 中的最大值。

**构建：**

```
suffixMax[n-1] = a[n-1]
for i := n-2 downto 0:
    suffixMax[i] = max(suffixMax[i+1], a[i])
```

**典型场景**：会议室身高（atcoder abc463_c），接雨水问题中的右侧最大高度。

---

### 对比总结

|          | 构建方向 | 构建复杂度 | 单次查询 | 典型用途             |
| -------- | -------- | ---------- | -------- | -------------------- |
| 前缀和   | 左→右    | $O(N)$     | $O(1)$   | 区间求和             |
| 后缀和   | 右→左    | $O(N)$     | $O(1)$   | 区间求和（反向视角） |
| 前缀最值 | 左→右    | $O(N)$     | $O(1)$   | 左侧历史最值         |
| 后缀最值 | 右→左    | $O(N)$     | $O(1)$   | 右侧未来最值         |

---

### 组合技巧

会**同时用前缀和后缀**，比如：

- **接雨水**：`min(prefixMax[i], suffixMax[i]) - a[i]`
- **除自身以外数组的积**：`prefixProd[i-1] * suffixProd[i+1]`
- **分割数组**：枚举分割点，左边用前缀信息，右边用后缀信息

核心思想都是一样的：**一次线性扫描预处理，换取每次查询 $O(1)$**。

## 应用场景/题目

快速求解子数组和（如LeetCode 560、303、304）。

结合哈希表统计特定和的子数组数量。

扩展到二维矩阵的区间和计算（二维前缀和）。

### 前缀/后缀最值

- AtCoder
  - abc463_c

## 代码模板实现

[suffixPrefixArray.go](./suffixPrefixArray/suffixPrefixArray.go)