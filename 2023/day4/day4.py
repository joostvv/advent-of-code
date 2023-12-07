from aocd import data, submit

test_data = """Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
"""
cards_points_total = 0
scratchcards = data.splitlines()
cards_amount = [1] * len(scratchcards)

for card in scratchcards:
    card_text, numbers = card.split(": ")
    winning_numbers, found_numbers = numbers.split(" | ")
    winning_numbers = list(map(int, filter(None, winning_numbers.split(" "))))
    found_numbers = list(map(int, filter(None, found_numbers.split(" "))))
    card_points_total = 1
    card_nr = card_counter = int(card_text[4:].strip())
    
    for nr in found_numbers:
        if nr in winning_numbers:
            card_points_total *= 2
            card_counter += 1
            cards_amount[card_counter - 1] += cards_amount[card_nr - 1]
    cards_points_total += card_points_total // 2

submit(cards_points_total, part='a') # type: ignore
submit(sum(cards_amount), part='b') # type: ignore
