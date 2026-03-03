const std = @import("std");

fn part1(input: std.ArrayList([4]u8)) u32 {
    var output: u32 = 0;
    for (input.items) |ranges| {
        if ((ranges[2] >= ranges[0] and ranges[1] >= ranges[3]) or
            (ranges[0] >= ranges[2] and ranges[3] >= ranges[1]))
        {
            output += 1;
        }
    }
    return output;
}

fn part2(input: std.ArrayList([4]u8)) u32 {
    var output: u32 = 0;
    for (input.items) |ranges| {
        if (ranges[0] <= ranges[2] and ranges[1] >= ranges[2]) {
            output += 1;
        } else if (ranges[2] <= ranges[0] and ranges[3] >= ranges[0]) {
            output += 1;
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

    var it = std.mem.splitScalar(u8, input, '\n');
    var parsed_input: std.ArrayList([4]u8) = .empty;
    defer parsed_input.deinit(allocator);

    while (it.next()) |line| {
        if (line.len == 0) break;
        var lineIt = std.mem.splitScalar(u8, line, ',');
        const first_range = try parseRangeNumbers(lineIt.next().?);
        const second_range = try parseRangeNumbers(lineIt.next().?);
        std.debug.print("{any}, {any}\n", .{ first_range, second_range });

        try parsed_input.append(allocator, [4]u8{
            first_range[0],
            first_range[1],
            second_range[0],
            second_range[1],
        });
    }

    std.debug.print("Part 1: {d}\n", .{part1(parsed_input)});
    std.debug.print("Part 2: {d}\n", .{part2(parsed_input)});
}

fn parseRangeNumbers(range: []const u8) ![2]u8 {
    var it = std.mem.splitScalar(u8, range, '-');
    const start = try std.fmt.parseInt(u8, it.next().?, 10);
    const end = try std.fmt.parseInt(u8, it.next().?, 10);

    return [2]u8{ start, end };
}
