const std = @import("std");

fn priority(c: u8) u8 {
    return switch (c) {
        'a'...'z' => c - 'a' + 1,
        'A'...'Z' => c - 'A' + 27,
        else => unreachable,
    };
}

fn part1(input: []u8) u32 {
    var output: u32 = 0;
    var it = std.mem.splitScalar(u8, input, '\n');

    while (it.next()) |line| {
        const len = line.len;
        const left = line[0 .. len / 2];
        const right = line[len / 2 ..];
        var seen: [53]u1 = [_]u1{0} ** 53;

        for (left) |c| seen[priority(c)] = 1;
        for (right) |c| {
            if (seen[priority(c)] == 1) {
                output += priority(c);
                break;
            }
        }
    }
    return output;
}

fn part2(input: []u8) u32 {
    var it = std.mem.splitScalar(u8, input, '\n');
    var seen: [53]u2 = [_]u2{0} ** 53;

    var curr_group: u32 = 0;
    var output: u32 = 0;

    while (it.next()) |line| {
        for (line) |c| {
            if (seen[priority(c)] == curr_group) {
                const new_val = seen[priority(c)] + 1;
                if (new_val == 3) {
                    output += priority(c);
                    break;
                }
                seen[priority(c)] = new_val;
            } else if (curr_group == 0) {
                seen[priority(c)] = 1;
            }
        }
        curr_group += 1;
        if (curr_group == 3) {
            seen = [_]u2{0} ** 53;
            curr_group = 0;
        }
    }
    return output;
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

    std.debug.print("Part 1: {d}\n", .{part1(input)});
    std.debug.print("Part 2: {d}\n", .{part2(input)});
}
