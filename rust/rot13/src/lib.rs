pub fn rotate_file(input: String, output: String) {
    rotate_impl(input, output, &StandardFileIO {});
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

impl FileIO for StandardFileIO {
     fn read(&self, filename: String) -> String {
         String::default()
     }

     fn write(&self, filename: String, contents: String) {

     }
}

#[cfg(test)]
mod tests {
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
        fn read(&self, filename: String) -> String {
            self.read_contents.clone()
        }

        fn write(&self, filename: String, contents: String) {
            *self.written_filename.borrow_mut() = filename;
            *self.written_contents.borrow_mut() = contents;
        }
    }

    #[test]
    fn rotate_writes_contents_of_input_with_rot13_applied() {
        let mock_fileio = MockFileIO::new(String::from("Testing, 1, 2, 3"));
        rotate_impl(String::from("in.txt"), String::from("out.txt"), &mock_fileio);

        assert_eq!(String::from("out.txt"), *mock_fileio.written_filename.borrow());
        assert_eq!(String::from("Grfgvat, 1, 2, 3"), *mock_fileio.written_contents.borrow());
    }
}
