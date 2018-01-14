use std::fs::File;
use std::io::prelude::*;

pub struct Config {
    is_production: bool
}

impl Config {
    pub fn debug() -> Config {
        Config {
            is_production: false
        }
    }

    pub fn production() -> Config {
        Config {
            is_production: true
        }
    }

    pub fn new(is_production: bool) -> Config {
        Config {
            is_production
        }
    }
}

pub fn rotate_file(input: &str, output: &str, config: Config) {
    rotate_impl(input, output, config, &StandardFileIO::new());
}

static TEST_PATH: &'static str = "test/";

fn rotate_impl<T>(input: &str, output: &str, config: Config, file_io: &T) where T: FileIO {
    let input = decide_path(input, &config);
    let output = decide_path(output, &config);
    let rotated_contents: String = file_io.read(input.as_str()).chars().map(rotate).collect();
    file_io.write(output.as_str(), rotated_contents);
}

fn decide_path(path: &str, config: &Config) -> String {
    if config.is_production {
        String::from(path)
    } else {
        format!("{}{}", TEST_PATH, path)
    }
}

fn rotate(input: char) -> char {
    if input.is_lowercase() {
        rotate_char(input, 'a')
    } else if input.is_uppercase() {
        rotate_char(input, 'A')
    } else {
        input
    }
}

fn rotate_char(input: char, base_char: char) -> char {
    let base_val = base_char as u8;
    let input_val = input as u8;
    let rotated_val = (input_val - base_val + 13) % 26;
    char::from(base_val + rotated_val)
}

trait FileIO {
    fn read(&self, filename: &str) -> String;
    fn write(&self, filename: &str, contents: String);
}

struct StandardFileIO;

impl StandardFileIO {
    fn new() -> StandardFileIO {
        StandardFileIO {}
    }
}

impl FileIO for StandardFileIO {
     fn read(&self, filename: &str) -> String {
         let mut file = File::open(filename).expect("file not found");
         let mut contents = String::default();
         file.read_to_string(&mut contents).expect("could not read file");
         contents
     }

     fn write(&self, filename: &str, contents: String) {
         let mut file = File::create(filename).expect("could not create file");
         file.write_all(contents.as_bytes()).expect("failed to write file contents");
     }
}

#[cfg(test)]
mod rotate_tests {
    use super::*;
    use std::cell::RefCell;

    struct MockFileIO {
        read_filename: RefCell<String>,
        read_contents: String,
        written_contents: RefCell<String>,
        written_filename: RefCell<String>,
    }

    impl MockFileIO {
        fn new(read_contents: &str) -> MockFileIO {
            MockFileIO {
                read_filename: RefCell::default(),
                read_contents: String::from(read_contents),
                written_contents: RefCell::default(),
                written_filename: RefCell::default(),
            }
        }

        fn verify_read_filename(&self, expected: &str) {
            assert_eq!(String::from(expected), *self.read_filename.borrow());
        }

        fn verify_written_filename(&self, expected: &str) {
            assert_eq!(String::from(expected), *self.written_filename.borrow());
        }

        fn verify_written_contents(&self, expected: &str) {
            assert_eq!(String::from(expected), *self.written_contents.borrow());
        }
    }

    impl FileIO for MockFileIO {
        fn read(&self, filename: &str) -> String {
            *self.read_filename.borrow_mut() = String::from(filename);
            self.read_contents.clone()
        }

        fn write(&self, filename: &str, contents: String) {
            *self.written_filename.borrow_mut() = String::from(filename);
            *self.written_contents.borrow_mut() = contents;
        }
    }

    #[test]
    fn writes_contents_of_input_with_rot13_applied() {
        let mock_fileio = MockFileIO::new("Testing, 1, 2, 3");

        rotate_impl("in.txt", "out.txt", Config::production(), &mock_fileio);

        mock_fileio.verify_read_filename("in.txt");
        mock_fileio.verify_written_filename("out.txt");
        mock_fileio.verify_written_contents("Grfgvat, 1, 2, 3");
    }

    #[test]
    fn reads_from_test_folder_when_in_debug_mode() {
        let mock_fileio = MockFileIO::new("Testing, 1, 2, 3");

        rotate_impl("in.txt", "out.txt", Config::debug(), &mock_fileio);

        mock_fileio.verify_read_filename("test/in.txt");
    }

    #[test]
    fn writes_to_test_folder_when_in_debug_mode() {
        let mock_fileio = MockFileIO::new("Testing, 1, 2, 3");

        rotate_impl("in.txt", "out.txt", Config::debug(), &mock_fileio);

        mock_fileio.verify_written_filename("test/out.txt");
    }
}

#[cfg(test)]
mod fileio_tests {
    use super::*;

    #[test]
    fn read_returns_contents_of_file() {
        let fileio = StandardFileIO::new();
        let contents = fileio.read("in.txt");
        assert_eq!(String::from("The dog barks at midnight."), contents);
    }

    #[test]
    fn write_writes_contents_to_file() {
        let filename = "out.txt";
        std::fs::remove_file(&filename).ok();

        let fileio = StandardFileIO::new();
        let contents = String::from("The dog barks at midnight.");
        fileio.write(filename, contents.clone());
        assert_eq!(contents, fileio.read(filename));

        std::fs::remove_file(filename).ok();
    }
}
