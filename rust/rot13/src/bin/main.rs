extern crate rot13;
use std::env;

fn main() {
	let args: Vec<String> = env::args().collect();
	if args.len() < 3 {
		panic!("Too few arguments");
	}
	let input_file = args[1].as_str();
	let output_file = args[2].as_str();
	let is_production = env::var("DEV").is_err();
	rot13::rotate_file(input_file, output_file, rot13::Config::new(is_production));
}
