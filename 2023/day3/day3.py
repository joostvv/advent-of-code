from aocd import data, submit

test_data = """467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
"""

board = data.splitlines() # type: ignore
part_numbers = []

def is_special_symbol(character):
    return not character.isdigit() and character != '.'

def is_gear_symbol(character):
    return character == '*'


def check_around_for_digits(xmax, ymax, board, x, y):  # noqa: PLR0911
    # Top
    if y > 0 and board[y-1][x]:
        return True
    # Top right
    if y > 0 and x < xmax and board[y-1][x+1].isdigit():
        return True
    # Top left
    if y > 0 and x > 0 and board[y-1][x-1].isdigit():
        return True
    # Left
    if x > 0 and is_special_symbol(board[y][x-1]):
        return True
    # Right
    if x < xmax and is_special_symbol(board[y][x+1]):
        return True
    # Bottom
    if y < ymax and is_special_symbol(board[y+1][x]):
        return True
    # Bottom right
    if y < ymax and x < xmax and is_special_symbol(board[y+1][x+1]):
        return True
    # Bottom left
    if y < ymax and x > 0 and is_special_symbol(board[y+1][x-1]):
        return True

    return False


ymax = len(board) - 1
for y, row in enumerate(board):
    print(row)
    xmax = len(row) - 1
    digit_word = ""
    is_machine_part = False
    for x, character in enumerate(row):
        if character.isdigit():
            digit_word += character
            if not is_machine_part and check_around_for_digits(xmax, ymax, board, x, y):
                is_machine_part = True
            if x == xmax and is_machine_part:
                part_numbers.append(int(digit_word))
                is_machine_part = False
        elif digit_word:
            if is_machine_part:
                part_numbers.append(int(digit_word))
                is_machine_part = False
            digit_word = ""
        
submit(sum(part_numbers), part='a') # type: ignore
