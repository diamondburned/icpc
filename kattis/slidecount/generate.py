#!/usr/bin/env python3
import random

n = 100_000
c = 1_000_000

print(f"{n} {c}")

l2 = ""
for _ in range(0, n):
    l2 += f"{random.randint(0, c//1000000)} "
print(l2)
