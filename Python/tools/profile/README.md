性能分析相关工具的使用

- timeit: timeit只输出被测试代码的总运行时间，单位为秒，没有详细的统计。
- profile: 纯Python实现的性能测试模块，接口和cProfile一样。
    - ncall: 函数运行次数
    - tottime: 函数的总的运行时间，减去函数中调用子函数的运行时间
    - cumtime: 函数及其所有子函数调整的运行时间，也就是函数开始调用到结束的时间。
- cProfile: c语言实现的性能测试模块，接口和profile一样。
- line_profiler: line_profiler可以统计每行代码的执行次数和执行时间等，时间单位为微妙。
- memory_profiler: 统计每行代码啊占用的内存大小
