f = open("input.dat", "r")
lines = [x[:-1] for x in f.readlines()]
f.close()


def part1():
    output = 0
    max_colors = {"red": 12, "green": 13, "blue": 14}

    for i, line in enumerate(lines):
        game_id = i + 1
        is_possible = True

        i = line.index(":")
        digits = 0
        while i < len(line):
            color = ""
            if line[i].isdigit():
                digits = line[i]
                while i + 1 < len(line) and line[i + 1].isdigit():
                    digits += line[i + 1]
                    i += 1
                digits = int(digits)
            if line[i].isalpha():
                color = line[i]
                while i + 1 < len(line) and line[i + 1].isalpha():
                    color += line[i + 1]
                    i += 1
            if digits and color and digits > max_colors[color]:
                is_possible = False
                break
            i += 1

        if is_possible:
            output += game_id
    return output


def part2():
    output = 0
    for i, line in enumerate(lines):
        max_colors = {"red": 0, "green": 0, "blue": 0}

        i = line.index(":")
        digits = 0
        while i < len(line):
            color = ""
            if line[i].isdigit():
                digits = line[i]
                while i + 1 < len(line) and line[i + 1].isdigit():
                    digits += line[i + 1]
                    i += 1
                digits = int(digits)
            if line[i].isalpha():
                color = line[i]
                while i + 1 < len(line) and line[i + 1].isalpha():
                    color += line[i + 1]
                    i += 1
            if digits and color and digits > max_colors[color]:
                max_colors[color] = digits
            i += 1

        curr = 1
        for v in max_colors.values():
            curr *= v
        output += curr
    return output


if __name__ == "__main__":
    print(part1())
    print(part2())
