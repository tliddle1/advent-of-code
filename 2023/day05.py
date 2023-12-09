def convert(src,m):
    dst = []
    for x in m.split('\n')[1:]:
        mappings = x.split()
        destination_start = int(mappings[0])
        source_start = int(mappings[1])
        range_length = int(mappings[2])
        for i in src:
            if source_start <= i < source_start+range_length:
                dst.append((destination_start+i-source_start))
                src.remove(i)
    return dst + src

def solve(lines):
    seeds, seeds_soil, soil_fert, fert_water, water_light, light_temp, temp_humid, humid_loc = lines
    seeds = seeds.split()[1:]
    seeds = [int(x) for x in seeds]
    result = []
    for s in seeds:
        soils = convert([s],seeds_soil)
        ferts = convert(soils,soil_fert)
        waters = convert(ferts,fert_water)
        lights = convert(waters,water_light)
        temps = convert(lights,light_temp)
        humids = convert(temps,temp_humid)
        locs = convert(humids,humid_loc)
        result.append(locs[0])
    return min(result)


inp = """seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4"""

lines = inp.split("\n\n")
print(solve(lines))

def convert(src,m):
    result = []
    for i in range(0,len(src),2):
        s = src[i]
        s_length = src[i+1]
        converted = False
        for x in m.split("\n")[1:]:
            x = x.split()
            dst_start = int(x[0])
            src_start = int(x[1])
            length = int(x[2])
            if src_start <= s < src_start + length:
                change = (dst_start-src_start)
                result.append(s+change)
                result.append(s_length)
                converted = True
                break
        if not converted:
            result.append(s)
            result.append(s_length)
    return result



def seed_loc(seeds, seeds_soil, soil_fert, fert_water, water_light, light_temp, temp_humid, humid_loc):
    new_seeds = getNewRanges(seeds,seeds_soil)
    soils = convert(new_seeds, seeds_soil)

    new_soils = getNewRanges(soils,soil_fert)
    ferts = convert(new_soils, soil_fert)

    new_ferts = getNewRanges(ferts,fert_water)
    waters = convert(new_ferts, fert_water)

    new_waters = getNewRanges(waters,water_light)
    lights = convert(new_waters, water_light)

    new_lights = getNewRanges(lights,light_temp)
    temps = convert(new_lights, light_temp)

    new_temps = getNewRanges(temps,temp_humid)
    humids = convert(new_temps, temp_humid)

    new_humids = getNewRanges(humids,humid_loc)
    locs = convert(new_humids, humid_loc)

    return locs

def getNewRanges(ranges, src_dst_mappings):
    new_ranges = []
    for i in range(0,len(ranges),2):
        range_start = ranges[i]
        range_length = ranges[i+1]
        converted = False
        for x in src_dst_mappings.split('\n')[1:]:
            mappings = x.split()
            dst_start = int(mappings[0])
            src_start = int(mappings[1])
            length = int(mappings[2])
            if range_start + range_length <= src_start or src_start + length <= range_start:
                pass
            elif src_start <= range_start and range_start + range_length <= src_start + length:
                new_ranges.append(range_start)
                new_ranges.append(range_length)
                converted = True
                break
            elif range_start < src_start and range_start + range_length > src_start and range_start + range_length <= src_start + length:
                new_ranges.append(src_start)
                new_ranges.append(range_start + range_length - src_start)
                # Get other ranges
                for n in getNewRanges([range_start, range_start + range_length - src_start],src_dst_mappings):
                    new_ranges.append(n)
                converted = True
                break
            elif src_start <= range_start and range_start < src_start + length and range_start + range_length > src_start + length:
                new_ranges.append(range_start)
                new_ranges.append(src_start+length-range_start)
                # Get other ranges
                for n in getNewRanges([src_start+length, range_length - (src_start+length-range_start)],src_dst_mappings):
                    new_ranges.append(n)
                converted = True
                break
            elif range_start < src_start and range_start + range_length > src_start + length:
                new_ranges.append(src_start)
                new_ranges.append(length)
                for n in getNewRanges([range_start,src_start-range_start,src_start+length, range_start+range_length-(src_start+length)], src_dst_mappings):
                    new_ranges.append(n)
                converted = True
                break
        if not converted:
            new_ranges.append(range_start)
            new_ranges.append(range_length)
    return new_ranges


def solve(lines):
    # seeds, seeds_soil = lines
    seeds, seeds_soil, soil_fert, fert_water, water_light, light_temp, temp_humid, humid_loc = lines
    seeds = seeds.split()[1:]
    seeds = [int(x) for x in seeds]
    result = []
    locs = seed_loc(seeds, seeds_soil, soil_fert, fert_water, water_light, light_temp, temp_humid, humid_loc)
    min_loc = float('inf')
    for i in range(0,len(locs),2):
        min_loc = min(locs[i], min_loc)
    return min_loc

lines = inp.split("\n\n")
print(solve(lines))
