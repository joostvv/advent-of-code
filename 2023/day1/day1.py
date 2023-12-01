import sys
from operator import itemgetter
from aocd import data, submit

MAX_INT = sys.maxsize * 2 + 1
translations = {
    "one": '1',
    "two": '2',
    "three": '3',
    "four": '4',
    "five": '5',
    "six": '6',
    "seven": '7',
    "eight": '8',
    "nine": '9'
}

def do_second_part(translations, targets, word):
    min_target = min([(word.find(target) if word.find(target) > -1 else MAX_INT, target) for target in targets], key=itemgetter(0))[1]
    max_target = max([(word.rfind(target), target) for target in targets], key=itemgetter(0))[1]
    return translations.get(min_target, min_target) + translations.get(max_target, max_target)

def do_first_part(translated_word):
    digits = [int(digit) for digit in filter(str.isdigit, translated_word)]
    return digits[0] * 10 + (digits[0] if len(digits) == 1 else + digits[-1])
            
lines = data.splitlines() # type: ignore
result_a = sum([do_first_part(word) for word in lines])
targets = list(translations.keys()) + list(translations.values())
result_b = sum([do_first_part(do_second_part(translations, targets, word)) for word in lines])

submit(result_a, part='a') # type: ignore
submit(result_b, part='b') # type: ignore
