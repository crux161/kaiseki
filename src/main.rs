use std::env;
use std::fs::File;
use std::io::Read;
use std::io::Write;
use regex::Regex;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args: Vec<String> = env::args().collect();
    let input = &args[1];
    let output = &args[2];

    let mut i = File::open(input)?;
    let mut o = File::create(output)?;

    let mut buf = String::new();
    i.read_to_string(&mut buf)?;

    let r = Regex::new(r"^[0-9a-f]{8}:")?;

    for line in buf.lines() {
        let line = line.trim();
        if r.is_match(line) {
            let line = line.split(':').collect::<Vec<&str>>(); // a little confusing but stick with
            if line.len() == 2 {                               // me here a little longer...
                let line = line[1].replace(" ", "").chars().take(32).collect::<String>();
                let data = hex::decode(line)?;
                o.write_all(&data)?;
            }
        }
    }

    Ok(())      // ok? OK!
}
