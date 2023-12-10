from functools import reduce
from aocd import data, submit

test_data = """0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
"""

next_numbers_in_sequence = []
previous_numbers_in_sequence = []
def calculate_deltas(sequence):
    return [sequence[i+1] - sequence[i] for i in range(len(sequence) - 1)]

sequences = data.splitlines() # type: ignore
length_sequences = len(list(map(int, sequences[0].split(" "))))
for sequence in sequences:
    deltas = [0] + list(map(int, sequence.split(" "))) + [0]
    while True:
        deltas = calculate_deltas(deltas)
        if len(deltas) == 2:
            break
    next_numbers_in_sequence.append(deltas[1] * -1)
    previous_numbers_in_sequence.append(deltas[0])

result_b = sum(previous_numbers_in_sequence)
if length_sequences % 2 != 1:
    result_b *= -1

submit(sum(next_numbers_in_sequence), part='a') # type: ignore
submit(result_b, part='b') # type: ignore
