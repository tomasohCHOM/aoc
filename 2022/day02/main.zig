const std = @import("std");

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

    var part1: u32 = 0;
    var part2: u32 = 0;

    while (try reader.takeDelimiter('\n')) |line| {
        const them = line[0] - 'A';
        const us = line[2] - 'X';
        // Part 1
        part1 += us + 1;
        if (us == them) {
            part1 += 3;
        } else if (@mod(us + 3 - them, 3) == 1) {
            part1 += 6;
        }
        // Part 2
        if (us == 0) {
            part2 += @mod(them + 2, 3) + 1;
        } else if (us == 1) {
            part2 += 3 + them + 1;
        } else {
            part2 += 6 + @mod(them + 1, 3) + 1;
        }
    }

    std.debug.print("Part 1: {d}\n", .{part1});
    std.debug.print("Part 2: {d}\n", .{part2});
}
