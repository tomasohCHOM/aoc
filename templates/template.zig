const std = @import("std");

fn part1(_: anytype) u32 {
    return 0;
}

fn part2(_: anytype) u32 {
    return 0;
}

pub fn main(init: std.process.Init) !void {
    const io = init.io;
    const cwd = std.Io.Dir.cwd();

    var file = try cwd.openFile(io, "input.txt", .{ .mode = .read_only });
    defer file.close(io);

    var read_buf: [1024]u8 = undefined;
    var file_reader = file.reader(io, &read_buf);
    const reader = &file_reader.interface;

    // Use any parsing method here
    while (try reader.takeDelimiter('\n')) |line| {
        std.debug.print("{s}\n", .{line}); // Placeholder
    }

    std.debug.print("Part 1: {d}\n", .{part1(0)});
    std.debug.print("Part 2: {d}\n", .{part2(0)});
}
