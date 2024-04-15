pub mod utils {
    pub fn hamming_distance(x: &[u8], y: &[u8]) -> usize {
        if x.len() != y.len() {
            panic!("Unequal sizes!!!")
        }

        let mut acc: usize = 0;

        for i in 0..x.len() {
            acc += (x[i] ^ y[i]).count_ones() as usize;
        }

        acc
    }

    pub fn repeating_key_xor(s: &str, key: &str) -> Vec<u8> {
        let bytes_s = s.as_bytes();
        let bytes_key = key.as_bytes();

        (0..s.len())
            .map(|idx| bytes_s[idx] ^ bytes_key[idx % key.len()])
            .collect()
    }

    pub fn score_word(s: &str) -> usize {
        let mut acc: usize = 0;

        let content: &str = "uldrhsnioate";

        for c in s.chars() {
            if content.contains(c) {
                let index_element = content.chars().position(|x| x == c).unwrap();
                acc += index_element;
            }
        }

        acc
    }

    pub fn hex_to_bytes(s: &str) -> Vec<u8> {
        (0..s.len())
            .step_by(2)
            .map(|idx| u8::from_str_radix(&s[idx..idx + 2], 16).expect("Invalid hex"))
            .collect()
    }

    pub fn bytes_to_hex(x: &[u8]) -> String {
        (0..x.len()).map(|idx| format!("{:02x}", x[idx])).collect()
    }

    pub fn xor(x: &[u8], y: &[u8]) -> Option<Vec<u8>> {
        if x.len() != y.len() {
            None
        } else {
            Some((0..x.len()).map(|idx| x[idx] ^ y[idx]).collect())
        }
    }

    pub fn crack_xor_key(b: &[u8]) -> u8 {
        let mut max = 0usize;
        let mut best_char = 0u8;

        for i in 0u8..255 {
            let y = vec![i; b.len()];
            if let Some(c) = xor(b, &y) {
                if let Ok(s) = std::str::from_utf8(&c) {
                    let score = score_word(s);

                    if score > max {
                        max = score;
                        best_char = i;
                    }
                }
            }
        }

        best_char
    }

    pub fn crypt_xor(input: &str, key: &str) -> Vec<u8> {
        let mut result = Vec::new();
        for (i, a) in input.bytes().enumerate() {
            let key_char = key.chars().nth(i % key.len()).unwrap();
            result.push(a ^ key_char as u8);
        }

        result
    }

    pub fn crack_single_xor_bytes(b: Vec<u8>) -> (String, usize) {
        let mut max_word = String::new();
        let mut max: usize = 0;

        for i in 0..255 {
            let y = vec![i; b.len()];

            let c = xor(&b, &y).unwrap();

            let s = match std::str::from_utf8(&c) {
                Ok(v) => v.to_owned(),
                Err(_) => String::from(""),
            };

            let score = score_word(&s);

            if score > max {
                max_word = s.clone();
                max = score;
            }
        }

        (max_word, max)
    }

    pub fn crack_single_xor(s: &str) -> (String, usize) {
        let b = hex_to_bytes(s);

        let mut max_word = String::new();
        let mut max: usize = 0;

        for i in 0..255 {
            let y = vec![i; b.len()];

            let c = xor(&b, &y).unwrap();

            let s = match std::str::from_utf8(&c) {
                Ok(v) => v.to_owned(),
                Err(_) => String::from(""),
            };

            let score = score_word(&s);

            if score > max {
                max_word = s.clone();
                max = score;
            }
        }

        (max_word, max)
    }
}
