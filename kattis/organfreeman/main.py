#!/usr/bin/env python
import math

def find_x(fx: int):
    total = 0
    while fx > 0:
        total += math.factorial(fx % 10)
        fx //= 10
        if fx == 0:
            return total
    return None


print(f"{find_x(3)=}")
print(f"{find_x(20)=}")
print(f"{find_x(3758)=}")
print(f"{find_x(45483)=}")
