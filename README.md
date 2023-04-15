# Kaiseki ("analysis")

So you have a serial log from u-boot you need parsed
into binary am I right? Well look no further, because your
dreams have come true -- in rust as well! No.. No, not really.

That being said, this *does* do the trick so clean up your output,
and get ready to parse some data!

`
usage: kaiseki <input> <output>
`

## Compiling

It's a rust project managed by Cargo -- what more could you ask for??

Install it like this:
`
git clone https://github.com/crux161/kaiseki.git
cd kaiseki
cargo b --release
cargo install --path .		# goes to /home/$USER/.cargo/bin most likely
`

Maybe I'll write a Makefile or some setup script to do this, but for
now don't be lazy! It's called copy and paste you heathens! :P


## Rust Support via Kaiseki

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
