const std = @import("std");

fn priority(c: u8) u8 {
    return switch (c) {
        'a'...'z' => c - 'a' + 1,
        'A'...'Z' => c - 'A' + 27,
        else => unreachable,
    };
}

fn part1(reader: *std.Io.Reader, allocator: std.mem.Allocator) !u32 {
    var output: u32 = 0;
    while (try reader.takeDelimiter('\n')) |line| {
        const len = line.len;
        const left = line[0 .. len / 2];
        const right = line[len / 2 ..];
        var map: std.AutoHashMap(u8, void) = .init(allocator);

        for (left) |c| try map.put(c, {});
        for (right) |c| {
            if (map.contains(c)) {
                output += priority(c);
                break;
            }
        }
        map.deinit();
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

    std.debug.print("Part 1: {d}\n", .{try part1(reader, allocator)});
    // std.debug.print("Part 2: {d}\n", .{part2(0)});
}
