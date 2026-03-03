const std = @import("std");

fn findMarker(input: []const u8, length: u4) u32 {
    var mask: u26 = 0;
    for (0..length - 1) |i| {
        mask ^= @as(u26, 1) << @intCast(input[i] - 'a');
    }
    for (0..input.len - length) |i| {
        const first = input[i];
        const last = input[i + length - 1];
        mask ^= @as(u26, 1) << @intCast(last - 'a');
        if (@popCount(mask) == length) {
            return @intCast(i + length);
        }
        mask ^= @as(u26, 1) << @intCast(first - 'a');
    }
    return 0;
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

    std.debug.print("Part 1: {d}\n", .{findMarker(input[0 .. input.len - 1], 4)});
    std.debug.print("Part 2: {d}\n", .{findMarker(input[0 .. input.len - 1], 14)});
}
