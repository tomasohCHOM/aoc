const std = @import("std");

fn part1(allocator: std.mem.Allocator, drawing: []const u8, moves: []const u8) !u32 {
    const stacks = try parseStacks(allocator, drawing);
    defer freeStacks(allocator, stacks);

    var instructions = try parseInstructions(allocator, moves);
    defer instructions.deinit(allocator);

    std.debug.print("{any}\n{any}\n", .{ stacks, instructions });

    const output = 0;
    return output;
}

fn part2(allocator: std.mem.Allocator, drawing: []const u8, moves: []const u8) !u32 {
    const stacks = try parseStacks(allocator, drawing);
    defer freeStacks(allocator, stacks);

    var instructions = try parseInstructions(allocator, moves);
    defer instructions.deinit(allocator);

    const output = 0;
    return output;
}

fn parseStacks(allocator: std.mem.Allocator, drawing: []const u8) ![]std.ArrayList(u8) {
    var lines_it = std.mem.splitScalar(u8, drawing, '\n');
    var lines: std.ArrayList([]const u8) = .empty;
    defer lines.deinit(allocator);
    while (lines_it.next()) |line| {
        try lines.append(allocator, line);
    }

    const number_lines = lines.items[lines.items.len - 1];
    const stack_count = (number_lines.len + 1) / 4;
    const stacks = try allocator.alloc(std.ArrayList(u8), stack_count);
    for (stacks) |*s| {
        s.* = .empty;
    }

    var row: usize = lines.items.len - 1;
    while (row > 0) {
        row -= 1;
        const line = lines.items[row];
        for (stacks, 0..) |*stack, i| {
            const col = 1 + i * 4;
            if (col < line.len) {
                const c = line[col];
                if (c != ' ') {
                    try stack.append(allocator, c);
                }
            }
        }
    }

    return stacks;
}

const Move = struct {
    count: u32,
    from: u32,
    to: u32,
};

fn parseInstructions(allocator: std.mem.Allocator, move_instructions: []const u8) !std.ArrayList(Move) {
    var moves: std.ArrayList(Move) = .empty;
    var it = std.mem.splitScalar(u8, move_instructions, '\n');
    while (it.next()) |line| {
        if (line.len == 0) continue;

        var tok = std.mem.tokenizeScalar(u8, line, ' ');

        _ = tok.next(); // "move"
        const count_str = tok.next().?;
        _ = tok.next(); // "from"
        const from_str = tok.next().?;
        _ = tok.next(); // "to"
        const to_str = tok.next().?;

        const move = Move{
            .count = try std.fmt.parseInt(u32, count_str, 10),
            .from = try std.fmt.parseInt(u32, from_str, 10) - 1,
            .to = try std.fmt.parseInt(u32, to_str, 10) - 1,
        };

        try moves.append(allocator, move);
    }

    return moves;
}

fn freeStacks(allocator: std.mem.Allocator, stacks: []std.ArrayList(u8)) void {
    for (stacks) |*s| {
        s.deinit(allocator);
    }
    allocator.free(stacks);
}

pub fn main(init: std.process.Init) !void {
    var gpa: std.heap.DebugAllocator(.{}) = .init;
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();

    const io = init.io;
    var args = init.minimal.args.iterate();
    _ = args.skip();

    var use_sample = false;
    while (args.next()) |arg| {
        if (std.mem.eql(u8, arg, "--sample") or std.mem.eql(u8, arg, "-s")) {
            std.debug.print("USING SAMPLE INPUT FILE\n", .{});
            use_sample = true;
            break;
        }
    }

    const input_file = if (use_sample) "sample.txt" else "input.txt";
    const cwd = std.Io.Dir.cwd();
    var file = try cwd.openFile(io, input_file, .{ .mode = .read_only });
    defer file.close(io);

    var read_buf: [1024]u8 = undefined;
    var file_reader = file.reader(io, &read_buf);
    const reader = &file_reader.interface;

    const stat = try file.stat(io);
    const size = stat.size;

    const input = try allocator.alloc(u8, size);
    defer allocator.free(input);
    _ = try reader.readSliceAll(input);

    var it = std.mem.splitSequence(u8, input, "\n\n");

    const drawing = it.next().?;
    const moves = it.next().?;

    std.debug.print("Part 1: {d}\n", .{try part1(allocator, drawing, moves)});
    std.debug.print("Part 2: {d}\n", .{try part2(allocator, drawing, moves)});
}
