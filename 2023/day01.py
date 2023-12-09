
# Part 1

inp = """1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet"""

def calculateCalibrationValue(line):
    result = 0
    for i in range(len(line)):
        c = line[i]
        try:
            result += int(c)*10
            break
        except:
            pass
    for i in range(len(line)-1,-1,-1):
        c = line[i]
        try:
            result += int(c)
            break
        except:
            pass
    return result

def solve(lines):
    result = 0
    for line in lines:
        result += calculateCalibrationValue(line)
    return result

lines = inp.split("\n")
print(solve(lines))

# Part 2

inp = """two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen"""

digits = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
    "zero": 0,
}

def caluclateCalibrationValue(n):
    result = 0

    letters = ""
    for i in range(len(n)):
        c = n[i]
        try:
            result += int(c)*10
            break
        except:
            letters += c
            found = False
            for d in digits:
                if d in letters:
                    result += digits[d]*10
                    found = True
                    break
            if found:
                break

    letters = ""
    for i in range(len(n)-1,-1,-1):
        c = n[i]
        try:
            result += int(c)
            break
        except:
            letters += c
            found = False
            for d in digits:
                if d in letters[::-1]:
                    result += digits[d]
                    found = True
                    break
            if found:
                break

    return result

def solve(lines):
    result = 0
    for l in lines:
        result += caluclateCalibrationValue(l)
    return result

lines = inp.split("\n")
print(solve(lines))
