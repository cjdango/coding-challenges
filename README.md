This repository contains all of the coding challenges I have tried. Doing these challenges helps me sharpen my data structure and algorithm skills. Solutions may or may not be mine.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Counting minimum steps](#counting-minimum-steps)
    - [Available steps](#available-steps)
    - [Examples](#examples)
    - [Solution &#035;1 (brute force)](#solution-1-brute-force)
    - [Solution &#035;2 (memoization)](#solution-2-memoization)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


<!-- START Counting minimum steps -->
# Counting minimum steps
Given a positive integer n, return the minimum number of steps to minimize `n` to 1

### Available steps

1. Decrement n by 1 `n-1`
2. If n is divisible by 2, then divide n by 2 `n/2`
3. If n is divisible by 3, then divide n by 3 `n/3`

### Examples

- 10 ⇒ 3 steps (10 ⇒ 9 ⇒ 3 ⇒ 1)
- 15 ⇒ 4 steps (15 ⇒ 5 ⇒ 4 ⇒ 2 ⇒ 1)
- 6 ⇒ 2 steps (6 ⇒ 2 ⇒ 1) or (6 ⇒ 3 ⇒ 1)

### Solution #1 (brute force)

Let's take the example above, `n=6` . 

There are numbers of paths we can take to minimize `n` to 1 and we can't possibly know how many steps does the fastest path takes without going through all of them, so we'll do a brute force approach.

We can create a [recursive function](https://www.geeksforgeeks.org/recursive-functions/) to facilitate the steps we can take.

We can think about this as a tree that can have three branches each node. Each branch represent each steps, `n-1` , `n/2` (if n is divisible by 2), and `n/3` (if n is divisible by 3). To further understand this solution, let's visualize it.

![recursive tree](/minsteps/counting-min-steps-top-down-tree.png)

As you can see from the image above, there are 3  fastest combinations of steps we can take to get 6 minimized to 1. All of the 3 combinations requires two steps, so `2` is the answer we are looking for, `2` is the count of minimum steps to minimize 6 to 1.

**Fastest Combinations:**

- `(6/3) / 2 = 1`
    - Step count 1 = 6/3
    - Step count 2 = /2

- `(6/3) - 1 = 1`
    - Step count 1 = 6/3
    - Step count 2 = -1
- `(6/2) / 3 = 1`
    - Step count 1 = 6/3
    - Step count 2 = /3

But how do we actually count the steps? Remember that we are doing a recursive function approach, so the return values of each function call are the actual step counts from the bottom to top (`n=1` to `n=6`).

Our recursive function looks like this `fn(n)` , if `n` is `1` our function returns `0`, since there are zero steps required to minimize `1` to `1`. Else it will further process the call, calling itself applying `step 1`, and optionally for `step 2 and 3`. For `step 2 and 3` we will compare which step is the minimum between the result of previous step and the result of applying the current step to the function `fn`, then add `1` to the processes` result, naturally counting the steps from bottom to top call stack:

```jsx
fn(n):
  if n === 1
    return 0
  
  // Step 1
  result = fn(n - 1)
  
  // Step 2
  if n is divisible by 2:
    result = min(result, fn(n/2))

  // Step 3
  if n is divisible by 3:
    result = min(result, fn(n/3))

  return result + 1
```

You can see my golang implementation of this solution [here](https://github.com/cjdango/coding-challenges/blob/main/minsteps/count.go).

### Solution #2 (memoization)

The previous solution solves the problem. It can get the minimum number of steps to minimize `n` to 1. The problem with the brute force approach is that it is very inefficient. We are repeating ourselves multiple times and it can take a very long time if `n` is a very big number. You can try the solution above yourself and try to input `1000` as `n`, like `fn(1000)`.

A lot of nodes are repeating it's calculations. What we can do is to still use the previous solution, then implement the concept of [memoization](https://en.wikipedia.org/wiki/Memoization). If you're familiar with react, this concept is what powers the `memo` react hook.

Memoization is a form of caching and memoized functions becomes optimized for speed in exchange for higher use of memory space. 

In our example above, calculations for `fn(3)` and `fn(2)` are being repeated multiple times, in other words we are calling `fn(2)` and `fn(3)` repeatedly but they always return the same results. 

Our function is [idempotent](https://en.wikipedia.org/wiki/Idempotence) , meaning we can execute it several times without changing the final result, given that we give it the same arguments.

e.g. `fn(6)` always returns `2`, `fn(3)` always returns `1`

So as you can imagine, we can just save the result of a function call in the memory then grab that result and use it as the return value the next time we encounter the same function call instead of calculating again and again. This way we can save a lot of time. 

e.g. <br/>
![table](/minsteps/table.png)

To memoized our function, we need some kind of memory storage to store the results of each unique nodes. We can utilize a list and it's indices since we are just dealing with positive integers. 

First we need to change our function signature from `fn(int) int` to `fn(int, int[]) int`

```jsx
fn(n, memo[]):
```

**We also need to make sure that the length of our `memo` is `n+1`. That's because we are going to use `n` as our index. And we also need to fill our `memo` with zeros.** Languages like Java and Golang initializes int arrays with zeros as their indices' empty values.

Then we need to change and add a few lines of code in our previous solution.

```jsx
fn(n int, memo int[]):
  if n === 1
    return 0

  // if n is already calculated and
  // stored its result in memo[n]
  if memo[n] != 0 {
    return memo[n] // grab and return result from memo
  }
  
  // Step 1
  result = fn(n - 1, memo) // pass memo
  
  // Step 2
  if n is divisible by 2:
    result = min(result, fn(n/2, memo)) // pass memo

  // Step 3
  if n is divisible by 3:
    result = min(result, fn(n/3, memo)) // pass memo

  memo[n] = result + 1 // save result in memo

  return memo[n] // return result
```

You can see my golang implementation of this solution [here](https://github.com/cjdango/coding-challenges/blob/main/minsteps/countmemo.go).
<!-- END Counting minimum steps -->
