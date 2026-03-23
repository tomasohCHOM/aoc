import argparse

if __name__ == "__main__":
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument(
        "--sample",
        "-s",
        action="store_true",
        help="Use sample.txt (manually filled in)",
    )
    args = arg_parser.parse_args()
    input_file = "input.txt"
    if args.sample:
        print("USING SAMPLE INPUT FILE")
        input_file = "sample.txt"

    input = [line for line in open(input_file).read().strip().split("\n")]

    stack, sizes = [], {}
    for line in input:
        if line.startswith("$ cd "):
            target = line.split()[-1]
            if target == "/":
                stack = ["/"]
            elif target == "..":
                stack.pop()
            else:
                stack.append(target)
        elif line[0].isdigit():
            size = int(line.split()[0])
            for i in range(1, len(stack) + 1):
                key = "/".join(stack[:i])
                sizes[key] = sizes.get(key, 0) + size

    part1 = sum(size for size in sizes.values() if size <= 100_000)

    total_space = 70_000_000
    needed_space = 30_000_000
    free_space = total_space - sizes["/"]
    delete_min = needed_space - free_space

    part2 = min(size for size in sizes.values() if size >= delete_min)

    print("Part 1:", part1)
    print("Part 2:", part2)
