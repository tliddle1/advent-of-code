inp = """Time:      7  15   30
Distance:  9  40  200"""

import math

lines = inp.split("\n")
time = lines[0].split()[1:]
distance = lines[1].split()[1:]
result = 1
for i in range(len(time)):
    t = float(time[i])
    d = float(distance[i])
    # calculate the discriminant
    det = (t**2) - (4*d)
    sol1 = (t-math.sqrt(det))/(2)
    sol2 = (t+math.sqrt(det))/(2)
    if sol1.is_integer():
        sol1 = sol1+1
    else:
        sol1 = math.ceil(sol1)
    if sol2.is_integer():
        sol2 = sol2-1
    else:
        sol2 = math.floor(sol2)
    
    result *= int(sol2-sol1+1)
print(result)

inp = """Time:      7  15   30
Distance:  9  40  200"""

import math

lines = inp.split("\n")
time = lines[0].split()[1:]
distance = lines[1].split()[1:]
result = 1
t = ""
d = ""
for i in range(len(time)):
    t += time[i]
    d += distance[i]
t = float(t)
d = float(d)
# calculate the discriminant
det = (t**2) - (4*d)
sol1 = (t-math.sqrt(det))/(2)
sol2 = (t+math.sqrt(det))/(2)
if sol1.is_integer():
    sol1 = sol1+1
else:
    sol1 = math.ceil(sol1)
if sol2.is_integer():
    sol2 = sol2-1
else:
    sol2 = math.floor(sol2)
print(sol2-sol1+1)
