use std::fs::File;
use std::io::prelude::*;

pub fn rotate_file(input: String, output: String) {
    rotate_impl(input, output, &StandardFileIO::new());
}

fn rotate_impl<T>(input: String, output: String, file_io: &T) where T: FileIO {
    let rotated_contents: String = file_io.read(input).chars().map(rotate).collect();
    file_io.write(output, rotated_contents);
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
    fn read(&self, filename: String) -> String;
    fn write(&self, filename: String, contents: String);
}

struct StandardFileIO;

impl StandardFileIO {
    fn new() -> StandardFileIO {
        StandardFileIO {}
    }
}

impl FileIO for StandardFileIO {
     fn read(&self, filename: String) -> String {
         let mut file = File::open(filename).expect("file not found");
         let mut contents = String::default();
         file.read_to_string(&mut contents).expect("could not read file");
         contents
     }

     fn write(&self, filename: String, contents: String) {
         let mut file = File::create(filename).expect("could not create file");
         file.write_all(contents.as_bytes()).expect("failed to write file contents");
     }
}

#[cfg(test)]
mod rotate_tests {
    use super::*;
    use std::cell::RefCell;

    struct MockFileIO {
        read_contents: String,
        written_contents: RefCell<String>,
        written_filename: RefCell<String>,
    }

    impl MockFileIO {
        fn new(read_contents: String) -> MockFileIO {
            MockFileIO {
                read_contents,
                written_contents: RefCell::default(),
                written_filename: RefCell::default(),
            }
        }
    }

    impl FileIO for MockFileIO {
        fn read(&self, _filename: String) -> String {
            self.read_contents.clone()
        }

        fn write(&self, filename: String, contents: String) {
            *self.written_filename.borrow_mut() = filename;
            *self.written_contents.borrow_mut() = contents;
        }
    }

    #[test]
    fn writes_contents_of_input_with_rot13_applied() {
        let mock_fileio = MockFileIO::new(String::from("Testing, 1, 2, 3"));
        rotate_impl(String::from("in.txt"), String::from("out.txt"), &mock_fileio);

        assert_eq!(String::from("out.txt"), *mock_fileio.written_filename.borrow());
        assert_eq!(String::from("Grfgvat, 1, 2, 3"), *mock_fileio.written_contents.borrow());
    }
}

#[cfg(test)]
mod fileio_tests {
    use super::*;

    #[test]
    fn read_returns_contents_of_file() {
        let fileio = StandardFileIO::new();
        let contents = fileio.read(String::from("in.txt"));
        assert_eq!(String::from("The dog barks at midnight."), contents);
    }

    #[test]
    fn write_writes_contents_to_file() {
        let filename = String::from("out.txt");
        std::fs::remove_file(&filename).ok();

        let fileio = StandardFileIO::new();
        let contents = String::from("The dog barks at midnight.");
        fileio.write(filename.clone(), contents.clone());
        assert_eq!(contents, fileio.read(filename.clone()));

        std::fs::remove_file(&filename).ok();
    }
}
