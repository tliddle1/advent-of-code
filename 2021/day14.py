inp = """NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C"""

def iterate(amts,pairs,d):
    new_pairs = {}
    for pair, num in pairs.items():
        # Get new letter
        new_letter = d[pair]
        
        # Make new pairs
        p1 = pair[0] + new_letter
        p2 = new_letter + pair[1]
        
        # Add new pairs
        for p in [p1,p2]:
            if p in new_pairs:
                new_pairs[p] += num
            else:
                new_pairs[p] = num

        # Add new letter freqs
        if new_letter in amts:
            amts[new_letter] += num
        else:
            amts[new_letter] = num
    return amts, new_pairs


def solve(lines, iterations):
    start = lines[0]
    
    # Get key value pairs
    d = {}
    for line in lines[2:]:
        k,v = line.split(" -> ")
        d[k] = v

    # Get Amounts
    amts = {}
    for c in start:
        if c in amts:
            amts[c] += 1
        else:
            amts[c] = 1
    
    # Get Pairs
    pairs = {}
    for i in range(len(start)-1):
        pair = start[i] + start[i+1]
        if pair in pairs:   
            pairs[pair] += 1
        else:
            pairs[pair] = 1
    
    # Solve
    for i in range(iterations):
        amts, pairs = iterate(amts,pairs,d)
    return max(amts.values())-min(amts.values())

print(solve(inp.split("\n"), 10))
print(solve(inp.split("\n"), 40))
