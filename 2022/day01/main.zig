const std = @import("std");
const parseInt = std.fmt.parseInt;

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

    var max_sum: u32 = 0;
    var second_max: u32 = 0;
    var third_max: u32 = 0;
    var curr_sum: u32 = 0;

    while (try reader.takeDelimiter('\n')) |line| {
        if (std.mem.eql(u8, line, "")) {
            if (curr_sum > max_sum) {
                third_max = second_max;
                second_max = max_sum;
                max_sum = curr_sum;
            } else if (curr_sum > second_max) {
                third_max = second_max;
                second_max = curr_sum;
            } else if (curr_sum > third_max) {
                third_max = curr_sum;
            }

            curr_sum = 0;
            continue;
        }

        const num = try parseInt(u32, line, 10);
        curr_sum += num;
    }

    std.debug.print("Part 1: {d}\n", .{max_sum});
    std.debug.print("Part 2: {d}\n", .{max_sum + second_max + third_max});
}
