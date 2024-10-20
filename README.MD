# Benchmarking Arithmetic Operations in Go

This benchmark measures the performance of basic arithmetic operations (addition, subtraction, multiplication, and division) for both integers and floating-point numbers in Go. The goal was to determine whether there is any significant performance difference between integer (`int`) and floating-point (`float64`) operations.

## Results

The benchmark was run on an AMD Ryzen 7 1700 Eight-Core Processor with the following results:

| Operation     | Type       | Time per Operation (ns) |
|---------------|------------|-------------------------|
| Addition      | `int`      | 0.2913 ns/op            |
| Subtraction   | `int`      | 0.2742 ns/op            |
| Multiplication| `int`      | 0.2705 ns/op            |
| Division      | `int`      | 0.2725 ns/op            |
| Addition      | `float64`  | 0.2682 ns/op            |
| Subtraction   | `float64`  | 0.2686 ns/op            |
| Multiplication| `float64`  | 0.2682 ns/op            |
| Division      | `float64`  | 0.2754 ns/op            |

## Key Insights

### 1. **Negligible Difference Between Integers and Floats**

From the benchmark, we can observe that both integer and floating-point operations take almost the same amount of time. The time per operation is around **0.27 to 0.29 nanoseconds**, regardless of whether we're working with `int` or `float64`.

### 2. **Floats Are Not Slower**

Historically, floating-point operations were more expensive than integer operations, especially in older processors or systems without dedicated floating-point hardware. However, in modern CPUs, like the AMD Ryzen 7 1700 used in this benchmark, the performance gap between floating-point and integer arithmetic has essentially disappeared.

### 3. **No Need to Optimize by Rewriting Floats as Integers**

In today’s era, with modern compilers and CPUs, there’s no significant advantage in rewriting floating-point code to use integers purely for performance reasons. Optimizing code in this way might have been important decades ago, but with modern hardware, the cost of floating-point arithmetic is almost identical to that of integer arithmetic.

### 4. **Focus on Code Clarity and Precision**

Since floating-point operations are no longer a bottleneck, it makes more sense to choose the appropriate data type based on the problem domain. For applications requiring precision (e.g., financial calculations), using integers (e.g., representing currency in cents) can still be a good choice. However, in cases where real numbers or fractional values are needed, there’s no performance-driven reason to avoid `float64` in favor of `int`.

## Conclusion

This benchmark demonstrates that in modern computing, there is no longer a significant performance difference between integer and floating-point arithmetic. Therefore, decisions on data types should be based on the specific requirements of the application (e.g., precision or value range) rather than performance concerns.
