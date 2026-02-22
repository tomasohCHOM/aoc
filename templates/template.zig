const std = @import("std");

fn part1(_: anytype) u32 {
    return 0;
}

fn part2(_: anytype) u32 {
    return 0;
}

pub fn main(init: std.process.Init) !void {
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

    // Use any parsing method here
    while (try reader.takeDelimiter('\n')) |_| {
        // Placeholder
    }

    std.debug.print("Part 1: {d}\n", .{part1(0)});
    std.debug.print("Part 2: {d}\n", .{part2(0)});
}
