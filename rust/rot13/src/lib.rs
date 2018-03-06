mod fileio;
pub mod config;
use fileio::*;
use config::Config;

pub fn rotate_file(input: &str, output: &str, config: Config) {
    rotate_impl(input, output, config, &StandardFileIO::new());
}

static TEST_PATH: &'static str = "test/";

fn rotate_impl<T>(input: &str, output: &str, config: Config, file_io: &T) where T: FileIO {
    let input = decide_path(input, &config);
    let output = decide_path(output, &config);
    let rotated_contents: FileContents = file_io.read(input.as_str()).map(rotate_string);
    //file_io.write(output.as_str(), rotated_contents);
}

fn decide_path(path: &str, config: &Config) -> String {
    if config.is_production {
        String::from(path)
    } else {
        format!("{}{}", TEST_PATH, path)
    }
}

fn rotate_string(input: String) -> String {
    input.chars().map(rotate).collect::<String>()
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

#[cfg(test)]
mod tests {
    use super::*;
    use std::cell::RefCell;

    #[test]
    fn writes_contents_of_input_with_rot13_applied() {
        let mock_fileio = MockFileIO::new("Testing, 1, 2, 3");

        rotate_impl("in.txt", "out.txt", config::production(), &mock_fileio);

        mock_fileio.verify_read_filename("in.txt");
        mock_fileio.verify_written_filename("out.txt");
        mock_fileio.verify_written_contents("Grfgvat, 1, 2, 3");
    }

    #[test]
    fn reads_from_test_folder_when_in_debug_mode() {
        let mock_fileio = MockFileIO::new("Testing, 1, 2, 3");

        rotate_impl("in.txt", "out.txt", config::debug(), &mock_fileio);

        mock_fileio.verify_read_filename("test/in.txt");
    }

    #[test]
    fn writes_to_test_folder_when_in_debug_mode() {
        let mock_fileio = MockFileIO::new("Testing, 1, 2, 3");

        rotate_impl("in.txt", "out.txt", config::debug(), &mock_fileio);

        mock_fileio.verify_written_filename("test/out.txt");
    }

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
        fn read(&self, filename: &str) -> Box<FileContents> {
            *self.read_filename.borrow_mut() = String::from(filename);
            Box::new(MockFileContents::new(self.read_contents.as_str()))
        }

        fn write(&self, filename: &str, contents: Box<FileContents>) {
            *self.written_filename.borrow_mut() = String::from(filename);
            *self.written_contents.borrow_mut() = contents.collect::<String>();
        }
    }
}
