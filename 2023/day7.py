inp = """32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483"""

def getLargestCardGroup(hand):
    num_j = 0
    d = {}
    for c in hand:
        if c in d.keys():
            d[c] += 1
        else:
            d[c] = 1
    max_v = 0
    for v in d.values():
        if v > max_v:
            max_v = v
    return max_v

def getTieBreakerScore(hand):
    d = {"A":"14", "K":"13", "Q":"12", "J":"11", "T":"10", "9":"09", "8":"08", "7":"07", "6":"06", "5":"05", "4":"04", "3":"03", "2":"02"}
    divisor = 100000
    result = ""
    for c in hand:
        result += d[c]
    return float(result) / divisor

def getType(hand):
    multiplier = 100000 # Used to differentiate hands without floating point error being a problem
    priority = 0
    cards = set()
    for c in hand:
        cards.add(c)
    if len(cards) <= 1:
        # Five of a kind
        priority = 6
    elif len(cards) == 2:
        if getLargestCardGroup(hand) == 4:
            # Four of a kind
            priority = 5
        else:
            # Full house
            priority = 4
    elif len(cards) == 3:
        if getLargestCardGroup(hand) == 3:
            # Three of a kind
            priority = 3
        else:
            # Two pair
            priority = 2
    elif len(cards) == 4:
        # One pair
        priority = 1
    return priority * multiplier

def scoreHand(hand):
    return getType(hand)+getTieBreakerScore(hand)

def solvePart1(lines):
    hands = []
    for line in lines:
        hand, bid = line.split()
        bid = int(bid)
        score = scoreHand(hand)
        hands.append((score,bid))
    ranked = sorted(hands)[::-1]
    result = 0
    rank = len(ranked)
    for hand in ranked:
        _, bid = hand
        result += bid*rank
        rank -= 1
    return result

lines = inp.split("\n")
print(solvePart1(lines))


def getLargestCardGroup(hand):
    num_j = 0
    d = {}
    for c in hand:
        if c == 'J':
            num_j += 1
        else:
            if c in d.keys():
                d[c] += 1
            else:
                d[c] = 1
    max_v = 0
    for v in d.values():
        if v > max_v:
            max_v = v
    return max_v + num_j

def getTieBreakerScore(hand):
    # J cards are now the weakest individual cards
    d = {"A":"14", "K":"13", "Q":"12", "J":"01", "T":"10", "9":"09", "8":"08", "7":"07", "6":"06", "5":"05", "4":"04", "3":"03", "2":"02"}
    divisor = 100000
    result = ""
    for c in hand:
        result += d[c]
    return float(result) / divisor

def getType(hand):
    multiplier = 100000 # Used to differentiate hands without floating point error being a problem
    priority = 0
    cards = set()
    for c in hand:
        cards.add(c)
    # Remove J cards because they would be changed to another card in the hand
    if 'J' in hand:
        cards.remove('J')
    if len(cards) <= 1:
        # Five of a kind
        priority = 6
    elif len(cards) == 2:
        if getLargestCardGroup(hand) == 4:
            # Four of a kind
            priority = 5
        else:
            # Full house
            priority = 4
    elif len(cards) == 3:
        if getLargestCardGroup(hand) == 3:
            # Three of a kind
            priority = 3
        else:
            # Two pair
            priority = 2
    elif len(cards) == 4:
        # One pair
        priority = 1
    return priority * multiplier

def scoreHand(hand):
    return getType(hand)+getTieBreakerScore(hand)

def solvePart2(lines):
    hands = []
    for line in lines:
        hand, bid = line.split()
        bid = int(bid)
        score = scoreHand(hand)
        hands.append((score,bid))
    ranked = sorted(hands)[::-1]
    result = 0
    rank = len(ranked)
    for hand in ranked:
        _, bid = hand
        result += bid*rank
        rank -= 1
    return result

lines = inp.split("\n")
print(solvePart2(lines))
