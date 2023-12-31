inp = """Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"""

def calc(line):
    cards = line.split(" | ")
    winning = set(cards[0].split(": ")[1].split())
    mine = set(cards[1].split())

    return 2**(len(winning.intersection(mine))) // 2
def solve(lines):
    result = 0
    for line in lines:
        result += calc(line)
    return result

lines = inp.split("\n")
print(solve(lines))

# One liner
print(sum([(2**(len(set(l.split(" | ")[0].split(": ")[1].split()).intersection(set(l.split(" | ")[1].split()))))//2) for l in inp.split("\n")]))

def calc(line):
    cards = line.split(" | ")
    winning = set(cards[0].split(": ")[1].split())
    mine = set(cards[1].split())

    return len(winning.intersection(mine))
def solve(lines):

    result = 0
    i = 1
    d = {}
    for line in lines:
        score = calc(line)

        for x in range(i+1,i+score+1):

            if x in d.keys():
                if i in d.keys():
                    d[x] += d[i]+1
                else:
                    d[x] += 1
            else:
                if i in d.keys():
                    d[x] = d[i]+1
                else:
                    d[x] = 1
        i += 1
    return sum(d.values())+len(lines)

inp = """Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"""

lines = inp.split("\n")
print(solve(lines))
