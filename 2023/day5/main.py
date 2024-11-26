f = open("input.dat", "r")
lines = [x[:-1] for x in f.readlines()]
f.close()


def part1():
    seeds = [int(x) for x in lines[0].split(": ")[1].split()]

    ans = float("inf")

    # [79, 14, 55, 13]
    for i in range(len(seeds)):
        done = False

        num = seeds[i]
        line_num = 2
        # ALL THE LINES
        while line_num < len(lines):
            if len(lines[line_num]) == 0:
                line_num += 1
                continue
            if lines[line_num][0].isalpha():
                line_num += 1
                done = False
                continue

            if done:
                line_num += 1
                continue

            destination, src, rng = [int(x) for x in lines[line_num].split()]

            if num >= src and num <= (src + rng):
                num = num + (destination - src)
                done = True
            line_num += 1

        if num < ans:
            ans = num

    return ans


# For part 2
transitions = [[], [], [], [], [], [], []]


def solve(l, r, lvl):
    if lvl == len(transitions):
        return l
    scores = []
    for destination, src, rng in transitions[lvl]:
        left = max(l, src)
        right = min(r, src + rng)
        if left > right:
            continue
        scores.append(
            solve(left - src + destination, right - src + destination, lvl + 1)
        )
    return min(scores) if scores else float("inf")


def part2():
    seeds = [int(x) for x in lines[0].split(": ")[1].split()]

    ans = float("inf")
    i = -2

    for line in lines:
        if len(line) == 0:
            continue
        if line[0].isalpha():
            i += 1
            continue
        transitions[i].append([int(x) for x in line.split()])

    for i in range(0, len(seeds), 2):
        num = solve(seeds[i], seeds[i] + seeds[i + 1], 0)
        if num < ans:
            ans = num

    return ans


if __name__ == "__main__":
    print(part1())
    print(part2())
