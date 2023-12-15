def HASH(s, cur=0):
    if len(s) == 0:
        return cur
    return HASH(s[1:],((cur + (ord(s[0])))*17) % 256)

def focusingPower(box, pos, n):
    return (int(box)+1)*pos*int(n)

inp = """rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"""

hashmap = {}
for i in range(256):
    hashmap[i] = {}
count = 0
for x in inp.split(","):
    if '-' in x:
        label = x[:-1]
        box = HASH(label)
        if label in hashmap[box]:
            del hashmap[box][label]
    else:
        label, n = x.split('=')
        box = HASH(label)
        if label in hashmap[box]:
            pos, _ = hashmap[box][label]
            hashmap[box][label] = (pos, n)
        else:
            hashmap[box][label] = (count, n)
    count += 1

result = 0
for box,lens_dict in hashmap.items():
    pos = 1
    for lens in sorted(lens_dict.values()):
        result += focusingPower(box, pos, lens[1])
        pos += 1

print(sum([HASH(x) for x in inp.split(",")]))
print(result)
