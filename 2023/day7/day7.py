import functools
from aocd import data, submit

test_data = """32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
"""

card_order_a = {"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
card_order_b = {"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 1, "Q": 12, "K": 13, "A": 14}
part = 'a'

def compare_card_strength(card_a: str, card_b: str) -> int:
    if part == 'a':
        if card_order_a[card_a] > card_order_a[card_b]:
            return 1
        if card_order_a[card_a] < card_order_a[card_b]:
            return -1
    else:
        if card_order_b[card_a] > card_order_b[card_b]:
            return 1
        if card_order_b[card_a] < card_order_b[card_b]:
            return -1
    return 0

def compare_cards_strength(hand_a: str, hand_b: str) -> int:
    for i in range(len(hand_a[0])):
        if compare_card_strength(hand_a[0][i], hand_b[0][i]) == 1:
            return 1
        if compare_card_strength(hand_a[0][i], hand_b[0][i]) == -1:
            return -1
    return 0

def get_hand_type_part_a(hand):
    card_sets = []
    for character in hand:
        if hand.count(character) > 1:
            card_sets.append(hand.count(character))

    hand_type = "Hc"
    if 5 in card_sets:
        hand_type = "5k"
    elif 4 in card_sets:
        hand_type = "4k"
    elif 3 in card_sets and 2 in card_sets:
        hand_type = "Fh"
    elif 3 in card_sets:
        hand_type = "3k"
    elif card_sets.count(2) > 2:
        hand_type = "2p"
    elif 2 in card_sets:
        hand_type = "1p"
    return hand_type

def get_hand_type_part_b(hand_type, amount_of_jacks):  # noqa: C901
    if amount_of_jacks == 5 or amount_of_jacks == 4:
        hand_type = "5k"
    elif amount_of_jacks == 3:
        hand_type = "5k" if hand_type == "1p" else "4k"
    elif amount_of_jacks == 2:
        if hand_type == "3k":
            hand_type = "5k"
        elif hand_type == "1p" or hand_type == "Fh":
            hand_type = "4k"
        else:
            hand_type = "3k"
    elif amount_of_jacks == 1:
        if hand_type == "4k":
            hand_type = "5k"
        elif hand_type == "3k":
            hand_type = "4k"
        elif hand_type == "2p":
            hand_type = "Fh"
        elif hand_type == "1p":
            hand_type = "3k"
        else:
            hand_type = "1p"
    return hand_type

def get_hand_type(hand: str, part) -> str:
    if part == 'a':
        return get_hand_type_part_a(hand)

    amount_of_jacks = hand.count('J')
    hand = hand.replace('J', '')
    hand_type = get_hand_type_part_a(hand)
    if amount_of_jacks > 0:
        return get_hand_type_part_b(hand_type, amount_of_jacks)
    return hand_type


camel_cards = data.splitlines() # type: ignore
for part in ['a', 'b']:
    ordered_by_kind = {"5k": [], "4k": [], "Fh": [], "3k": [], "2p": [], "1p": [], "Hc": []}
    for camel_card in camel_cards:
        (card, strength) = camel_card.split(" ")
        ordered_by_kind[get_hand_type(str(card), part)].append((str(card), int(strength)))

    ranked_hands = []
    for kind in reversed(ordered_by_kind):
        ranked_hands += sorted(ordered_by_kind[kind], key=functools.cmp_to_key(compare_cards_strength))

    total_score = 0
    for rank, value in enumerate(ranked_hands):
        total_score += (rank + 1) * value[1]
    submit(total_score, part=part) # type: ignore



