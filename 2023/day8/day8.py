from aocd import data, submit
import math

test_data_1 = """RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
"""

test_data_2 = """LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
"""

test_data_3 = """LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
"""

input_data = data.splitlines() # type: ignore
instructions = input_data[0]
amount_of_instructions = len(instructions)
nodes = {}
for node in input_data[2:]:
    node_name, connected_nodes = node.split(" = ")
    connected_nodes = connected_nodes[1:-1].split(", ")
    nodes[node_name] = (connected_nodes[0], connected_nodes[1])

def get_left_item(nodes, current_node):
    return nodes[current_node][0]

def get_right_item(nodes, current_node):
    return nodes[current_node][1]

def do_part_a(instructions, amount_of_instructions, nodes):
    current_node = "AAA"
    total_steps = 0
    current_instruction = 0
    while current_node != "ZZZ":
        next_step = instructions[current_instruction]
        if next_step == 'L':
            current_node = get_left_item(nodes, current_node)
        else:
            current_node = get_right_item(nodes, current_node)
        current_instruction = (current_instruction + 1) % amount_of_instructions
        total_steps += 1
    return total_steps

submit(do_part_a(instructions, amount_of_instructions, nodes), part='a') # type: ignore

current_nodes = list(filter(lambda node: node.endswith("A"), nodes))
total_steps = 0
current_instruction = 0
loop_values = []

while len(current_nodes) != 0:
    next_step = instructions[current_instruction]
    for i, current_node in enumerate(current_nodes):
        if next_step == 'L':
            current_nodes[i] = get_left_item(nodes, current_node)
        else:
            current_nodes[i] = get_right_item(nodes, current_node)

    total_steps += 1
    for current_node in current_nodes:
        if current_node.endswith("Z"):
            loop_values.append(total_steps)
            current_nodes.remove(current_node)
    current_instruction = (current_instruction + 1) % amount_of_instructions

# Inspiration by reddit
submit(math.lcm(*loop_values), part='b') # type: ignore
