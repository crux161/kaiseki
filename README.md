# firmwaretools

A set of scripts and tools for various firmware analysis tasks. 

## Rust Support via Kaiseki (the `parse-uboot-dump.py`-rs)

The original python script in this project has been ported to rust.
At least, I did my best trying to understand and recreate the logic.
There is no help if something goes wrong in the parsing, it will simply
complain and error out enigmatically.. This.. needs to be fixed.

That being said, not everyone has python installed everywhere they need
to run this, and being rust you can compile a version for your platform
very easily.

Despite being written in rust -- and timing a few seconds ahead of it's
python port -- a lot could probably be done to make this thing faster.

But then again, you just let picocom dump data overnight-- what's
another 15 seconds waiting to parse ~200MB?

## Naming

This contribution of mine is named Kaiseki after the Japanese
word 「解析」meaning "analysis"... I thought it was cute anyways. :P
