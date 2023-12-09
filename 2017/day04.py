inp = """aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa"""

result = 0
for line in inp.split("\n"):
    words = line.split()
    n = len(words)
    if n == len(set(words)):
        result += 1
print(result)

inp = """abcde fghij
abcde xyz ecdab
a ab abc abd abf abj
iiii oiii ooii oooi oooo
oiii ioii iioi iiio"""

result = 0
for line in inp.split("\n"):
    words = line.split()
    n = len(words)
    if n == len(set([str(sorted(w)) for w in words])):
        result += 1
print(result)
