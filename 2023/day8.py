import math

def solvePart1(lines):
    d = {}
    instructions = lines[0]
    k = 'AAA'
    for line in lines[1].split("\n"):
        key = line.split()[0]
        left = line.split()[2][1:4]
        right = line.split()[3][:3]
        d[key] = (left,right)
    i = 0
    n = len(instructions)
    while k != 'ZZZ':
        cur = instructions[i%n]
        v = d[k]
        if cur == 'L':
            k = v[0]
        else:
            k = v[1]
        i += 1
    return i

inp = """LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)"""

lines = inp.split("\n\n")
print(solvePart1(lines))

def solvePart2(lines):
    d = {}
    instructions = lines[0]
    keys = []
    for line in lines[1].split("\n"):
        key = line.split()[0]
        left = line.split()[2][1:4]
        right = line.split()[3][:3]
        d[key] = (left,right)
        if key[-1] == "A":
            keys.append(key)
    
    n = len(instructions)
    results = []
    for k in keys:
        i = 0
        while k[-1] != 'Z':
            cur = instructions[i%n]
            v = d[k]
            if cur == 'L':
                k = v[0]
            else:
                k = v[1]
            i += 1
        results.append(i)
    return math.lcm(*results)

inp = """LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)"""

lines = inp.split("\n\n")
print(solvePart2(lines))
