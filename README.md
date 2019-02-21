# FibonacciJourney
## 主题
要传达的东西
1. 动态规划
2. 缓存算法
3. 算法时间空间复杂度
4. 快速幂
5. 如何代码优化
6. 如何追求卓越
7. 如何掌握学习方法
## 提纲
### A.由第一个问题开头
爬楼的问题引入主题---斐波那契数列的实现
### B.第一种解法：纯递归
分析递归法的优缺点，并引入缓存法
* 优点：代码容易实现
* 缺点：代码空间效率极低
### C.第二种解法：缓存法
分析缓存法优缺点，并引入迭代法
* 优点：比纯递归快
* 缺点：数据空间略大
### D.第三种解法：迭代法
实现迭代法
* 演示迭代法效率上的提升
* 根据上条的空间复杂度引入时间复杂度及分析方法
### E.第二个问题：还能更快吗？
引入两个方式
* 矩阵乘法实现的斐波那契算法
* 快速幂
### F.讲解快速幂算法
讲解并实现快速幂，并分析快速幂的时间复杂度
### G.第四个解法：矩阵+快速幂
根据前两条的思路实现矩阵+快速幂的算法
* 分析效率和复杂度
* 引入斐波那契数列的公式算法
### H.第五个解法：公式法
根据公式法实现斐波那契
* 两者都是O(c*lgN)，比较两者的常数c
* √5提前算完
* 解释为啥buildin的方法没有自己写得快
### I.总结
看似是一场炫技之旅，其实我们真正要讲的是什么
1. 分析第一个问题所引入的动态规划思想
2. 算法的时空模型
3. 过早的优化是万恶之源
4. 优化代码的各种方法
    * 每次计算数据->缓存计算好的数据
    * 递归->迭代
    * 引用数学技巧（看看得了）
    * 别迷信系统库，具体问题具体分析
### J.哲学思考
我们应该如何精益求精，如何简历学习方法
### K.end
Thanks all！

## 各个算法结果
### A.正确性
```
=== RUN   TestFibRecursion
1.4930352e+07
--- PASS: TestFibRecursion (0.06s)
=== RUN   TestFibCache
7.0330367711422765e+208
--- PASS: TestFibCache (0.00s)
=== RUN   TestFibIterator
7.0330367711422765e+208
--- PASS: TestFibIterator (0.00s)
=== RUN   TestFibMatrix
7.033036771142278e+208
--- PASS: TestFibMatrix (0.00s)
=== RUN   TestFibFormulaBuildIn
7.033036771142261e+208
--- PASS: TestFibFormulaBuildIn (0.00s)
=== RUN   TestFibFormula
7.033036771142261e+208
--- PASS: TestFibFormula (0.00s)
PASS
```

### B.效率
```
goos: darwin
goarch: amd64
pkg: github.com/Ninlgde/FibonacciJourney/go
BenchmarkFibRecursion-12         	      30	  50489203 ns/op
BenchmarkFibCache-12             	  200000	      7859 ns/op
BenchmarkFibIterator-12          	 1000000	      1503 ns/op
BenchmarkFibMatrix-12            	 5000000	       339 ns/op
BenchmarkFibFormulaBuildIn-12    	20000000	        69.9 ns/op
BenchmarkFibFormula-12           	100000000	        17.9 ns/op
PASS
```