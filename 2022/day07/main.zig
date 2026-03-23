const std = @import("std");

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
    var stack: std.ArrayList([]const u8) = .empty;
    defer stack.deinit(allocator);

    var sizes: std.StringHashMap(u64) = .init(allocator);
    defer {
        var it_cleanup = sizes.iterator();
        while (it_cleanup.next()) |entry| {
            allocator.free(entry.key_ptr.*);
        }
        sizes.deinit();
    }

    while (it.next()) |line| {
        if (line.len == 0) break;

        if (std.mem.startsWith(u8, line, "$ cd")) {
            const idx = std.mem.lastIndexOfScalar(u8, line, ' ') orelse unreachable;
            const target = line[idx + 1 ..];
            if (std.mem.eql(u8, target, "/")) {
                stack.clearRetainingCapacity();
                try stack.append(allocator, "/");
            } else if (std.mem.eql(u8, target, "")) {
                _ = stack.pop();
            } else {
                try stack.append(allocator, target);
            }
        } else if (std.ascii.isDigit(line[0])) {
            var parts = std.mem.splitScalar(u8, line, ' ');
            const size_str = parts.next().?;
            const file_size = try std.fmt.parseInt(u64, size_str, 10);

            var path_buf: std.ArrayList(u8) = .empty;
            defer path_buf.deinit(allocator);

            for (stack.items, 0..) |dir, i| {
                if (i == 0) {
                    try path_buf.appendSlice(allocator, "/");
                } else {
                    if (!std.mem.eql(u8, path_buf.items, "/")) {
                        try path_buf.append(allocator, '/');
                    }
                    try path_buf.appendSlice(allocator, dir);
                }
                var entry = try sizes.getOrPut(path_buf.items);

                if (!entry.found_existing) {
                    entry.key_ptr.* = try allocator.dupe(u8, path_buf.items);
                    entry.value_ptr.* = 0;
                }

                entry.value_ptr.* += file_size;
            }
        }

        std.debug.print("{s}\n", .{line});
    }

    // std.debug.print("Part 1: {d}\n", .{part1});
    // std.debug.print("Part 2: {d}\n", .{part2});
}
