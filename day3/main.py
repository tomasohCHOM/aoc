import collections

f = open("input.dat", "r")
lines = [x[:-1] for x in f.readlines()]
f.close()

lines = [[c for c in x] for x in lines]


def check(r, i, j):
    for k in range(i - 1, j + 2):
        if k < 0 or k >= len(lines[0]):
            continue
        for s in range(r - 1, r + 2):
            if (
                s >= 0
                and s < len(lines)
                and lines[s][k] != "."
                and not lines[s][k].isdigit()
            ):
                return True
    return False


def part1():
    res = 0
    for r, line in enumerate(lines):
        i = 0
        n = len(line)
        while i < n:
            if line[i].isdigit():
                j = i
                curr = line[i]
                while j + 1 < n and line[j + 1].isdigit():
                    curr += line[j + 1]
                    j += 1
                digits = int(curr)
                if check(r, i, j):
                    res += digits
                if i != j:
                    i = j + 1
                else:
                    i += 1
            else:
                i += 1
    return res


def check2(r, i, j, locations, digits):
    for k in range(i - 1, j + 2):
        if k < 0 or k >= len(lines[0]):
            continue
        for s in range(r - 1, r + 2):
            if s >= 0 and s < len(lines) and lines[s][k] == "*":
                locations[(s, k)].append(digits)
    return False


def part2():
    res = 0
    locations = collections.defaultdict(list)
    for r, line in enumerate(lines):
        i = 0
        n = len(line)
        while i < n:
            if line[i].isdigit():
                j = i
                curr = line[i]
                while j + 1 < n and line[j + 1].isdigit():
                    curr += line[j + 1]
                    j += 1
                digits = int(curr)
                check2(r, i, j, locations, digits)
                if i != j:
                    i = j + 1
                else:
                    i += 1
            else:
                i += 1
    for v in locations.values():
        if len(v) == 2:
            res += v[0] * v[1]

    return res


if __name__ == "__main__":
    print(part1())
    print(part2())
