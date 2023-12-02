from functools import reduce
import operator
from aocd import data, submit

cubes_max = {'red': 12, 'green': 13, 'blue': 14}
test_date = """Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
"""

lines = data.splitlines() # type: ignore
plausible_games = []
power_set_games = []

def check_game(game) -> tuple[bool, int]:
    turns = game.split("; ")
    plausible = True
    minimal_cubes_game = {'red': 0, 'green': 0, 'blue': 0}
    for turn in turns:
        stones = turn.split(", ")
        for stone in stones:
            amount, color = stone.split(" ")
            if cubes_max[color] < int(amount):
                plausible = False
            if minimal_cubes_game[color] < int(amount):
                minimal_cubes_game[color] = int(amount)
    return plausible, reduce(operator.mul, minimal_cubes_game.values(), 1)

for line in lines:
    game_text, game = line.split(": ")
    (is_plausible, minimal_sets_power) = check_game(game)
    if is_plausible:
        _, game_nr = game_text.split(" ")
        plausible_games.append(int(game_nr))
    power_set_games.append(minimal_sets_power)

submit(sum(plausible_games), part='a') # type: ignore
submit(sum(power_set_games), part='b') # type: ignore
