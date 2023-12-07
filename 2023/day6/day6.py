from functools import reduce
import operator
from aocd import data, submit

test_data = """Time:      7  15   30
Distance:  9  40  200
"""

def calculate_options_naive(time, distance):
    count = 0
    for seconds in range(time):
        if (time - seconds) * seconds > distance:
            count += 1
    return count

data = data.splitlines() # type: ignore
_, times = data[0].split(": ")
_, distances = data[1].split(": ")

# Assignment 1
time = int("".join(list(filter(None, times.split(" ")))))
distance = int("".join(list(filter(None, distances.split(" ")))))
winning_possibility = calculate_options_naive(time, distance)

# Assignment 2
times = list(map(int, filter(None, times.split(" "))))
distances = list(map(int, filter(None, distances.split(" "))))
winning_possibilities = []
for time, distance in zip(times, distances):  # noqa: B905
    winning_possibilities.append(calculate_options_naive(time, distance))

submit(reduce(operator.mul, winning_possibilities, 1), part='a') # type: ignore
submit(winning_possibility, part='b') # type: ignore
