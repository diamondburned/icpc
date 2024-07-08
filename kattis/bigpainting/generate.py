#!/usr/bin/env python3
import random

painting = (500, 500)
masterpiece = (2000, 2000)

print(f"{painting[0]} {painting[1]} {masterpiece[0]} {masterpiece[1]}")

for _ in range(0, painting[1]):
    for _ in range(0, painting[0]):
        print("o", end="")
    print()

for _ in range(0, masterpiece[1]-1):
    for _ in range(0, masterpiece[0]):
        print("o", end="")
    print()
for _ in range(0, masterpiece[0]):
    print("x", end="")
print()
