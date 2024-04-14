use base64::{engine::general_purpose, Engine};
use openssl::symm::{decrypt, Cipher};

use std::fs;
use std::str;

fn main() {
    // let s = fs::read_to_string("6.txt")
    //     .unwrap()
    //     .split('\n')
    //     .collect::<Vec<_>>()
    //     .join("");

    // let b = general_purpose::STANDARD.decode(s).unwrap();
    // let mut distances = vec![];

    // for keysize in 2..40 {
    //     let first = b.chunks(keysize).nth(0).unwrap();
    //     let second = b.chunks(keysize).nth(1).unwrap();
    //     let third = b.chunks(keysize).nth(2).unwrap();
    //     let fourth = b.chunks(keysize).nth(3).unwrap();

    //     let distance1 = hamming_distance(first, second) as f32 / keysize as f32;
    //     let distance2 = hamming_distance(third, fourth) as f32 / keysize as f32;

    //     distances.push((keysize, (distance1 + distance2) / 2.0));
    // }

    // distances.sort_by(|a, b| a.1.partial_cmp(&b.1).unwrap());

    // let mut min_score = std::f32::MAX;
    // let mut min_key = String::new();
    // let mut min_result = vec![];
}

fn hamming_distance(x: &[u8], y: &[u8]) -> usize {
    if x.len() != y.len() {
        panic!("Unequal sizes!!!")
    }

    let mut acc: usize = 0;

    for i in 0..x.len() {
        acc += (x[i] ^ y[i]).count_ones() as usize;
    }

    acc
}

fn repeating_key_xor(s: &str, key: &str) -> Vec<u8> {
    let bytes_s = s.as_bytes();
    let bytes_key = key.as_bytes();

    (0..s.len())
        .map(|idx| bytes_s[idx] ^ bytes_key[idx % key.len()])
        .collect()
}

fn score_word(s: &str) -> usize {
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

fn hex_to_bytes(s: &str) -> Vec<u8> {
    (0..s.len())
        .step_by(2)
        .map(|idx| u8::from_str_radix(&s[idx..idx + 2], 16).expect("Invalid hex"))
        .collect()
}

fn bytes_to_hex(x: &[u8]) -> String {
    (0..x.len()).map(|idx| format!("{:02x}", x[idx])).collect()
}

fn xor(x: &[u8], y: &[u8]) -> Option<Vec<u8>> {
    if x.len() != y.len() {
        None
    } else {
        Some((0..x.len()).map(|idx| x[idx] ^ y[idx]).collect())
    }
}

fn crack_single_xor(s: &str) -> (String, usize) {
    let b = hex_to_bytes(s);

    let mut max_word = String::new();
    let mut max: usize = 0;

    for i in 0..128 {
        let y = vec![i; b.len()];

        let c = xor(&b, &y).unwrap();

        let s = match str::from_utf8(&c) {
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

#[cfg(test)]
mod tests {
    use super::*;
    use base64::prelude::*;
    use openssl::symm::Cipher;

    #[test]
    fn s1c1() {
        let hex = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d";

        let b = hex_to_bytes(hex);

        let mut result = String::new();

        BASE64_STANDARD.encode_string(b, &mut result);

        assert_eq!(
            result,
            "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
        )
    }

    #[test]
    fn s1c2() {
        let x = "1c0111001f010100061a024b53535009181c";
        let y = "686974207468652062756c6c277320657965";

        let bytes = match xor(&hex_to_bytes(x), &hex_to_bytes(y)) {
            Some(c) => c,
            None => panic!("Unable to XOR, sry, bai!!!"),
        };

        assert_eq!(bytes_to_hex(&bytes), "746865206b696420646f6e277420706c6179")
    }

    #[test]
    fn s1c3() {
        let x = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736";

        let b = hex_to_bytes(x);

        let mut max_word = String::new();
        let mut max: usize = 0;

        for i in 0..128 {
            let y: Vec<u8> = vec![i; b.len()];

            let c = xor(&b, &y).unwrap();

            let s = match str::from_utf8(&c) {
                Ok(v) => v.to_owned(),
                Err(e) => panic!("Invalid UTF-8 sequence: {}", e),
            };

            let score = score_word(&s);

            if score > max {
                max_word = s.clone();
                max = score;
            }
        }

        assert_eq!("Cooking MC's like a pound of bacon", max_word);
    }

    #[test]
    fn s1c4() {
        let s = fs::read_to_string("4.txt").unwrap();

        let mut max: usize = 0;
        let mut max_string: String = String::from("");

        for line in s.lines() {
            let (cracked, cracked_num) = crack_single_xor(line);

            if cracked_num > max {
                max = cracked_num;
                max_string = cracked;
            }
        }

        assert_eq!(
            "Now that the party is jumping",
            max_string.strip_suffix('\n').unwrap()
        );
    }

    #[test]
    fn s1c5() {
        let s = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal";
        assert_eq!("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f",bytes_to_hex(&repeating_key_xor(s, "ICE")));
    }

    #[test]
    fn hamming_test() {
        assert_eq!(
            37,
            hamming_distance("this is a test".as_bytes(), "wokka wokka!!!".as_bytes())
        );
    }

    #[test]
    fn s1c6() {}

    #[test]
    fn s1c7() {
        let key = "YELLOW SUBMARINE".as_bytes();

        let s = fs::read_to_string("7.txt")
            .unwrap()
            .split('\n')
            .collect::<Vec<_>>()
            .join("");

        let b = general_purpose::STANDARD.decode(s).unwrap();

        let cipher = Cipher::aes_128_ecb();

        let ciphertext = decrypt(cipher, key, None, &b).unwrap();

        let plaintext = str::from_utf8(&ciphertext).unwrap();

        assert_eq!(plaintext, "I'm back and I'm ringin' the bell \nA rockin' on the mike while the fly girls yell \nIn ecstasy in the back of me \nWell that's my DJ Deshay cuttin' all them Z's \nHittin' hard and the girlies goin' crazy \nVanilla's on the mike, man I'm not lazy. \n\nI'm lettin' my drug kick in \nIt controls my mouth and I begin \nTo just let it flow, let my concepts go \nMy posse's to the side yellin', Go Vanilla Go! \n\nSmooth 'cause that's the way I will be \nAnd if you don't give a damn, then \nWhy you starin' at me \nSo get off 'cause I control the stage \nThere's no dissin' allowed \nI'm in my own phase \nThe girlies sa y they love me and that is ok \nAnd I can dance better than any kid n' play \n\nStage 2 -- Yea the one ya' wanna listen to \nIt's off my head so let the beat play through \nSo I can funk it up and make it sound good \n1-2-3 Yo -- Knock on some wood \nFor good luck, I like my rhymes atrocious \nSupercalafragilisticexpialidocious \nI'm an effect and that you can bet \nI can take a fly girl and make her wet. \n\nI'm like Samson -- Samson to Delilah \nThere's no denyin', You can try to hang \nBut you'll keep tryin' to get my style \nOver and over, practice makes perfect \nBut not if you're a loafer. \n\nYou'll get nowhere, no place, no time, no girls \nSoon -- Oh my God, homebody, you probably eat \nSpaghetti with a spoon! Come on and say it! \n\nVIP. Vanilla Ice yep, yep, I'm comin' hard like a rhino \nIntoxicating so you stagger like a wino \nSo punks stop trying and girl stop cryin' \nVanilla Ice is sellin' and you people are buyin' \n'Cause why the freaks are jockin' like Crazy Glue \nMovin' and groovin' trying to sing along \nAll through the ghetto groovin' this here song \nNow you're amazed by the VIP posse. \n\nSteppin' so hard like a German Nazi \nStartled by the bases hittin' ground \nThere's no trippin' on mine, I'm just gettin' down \nSparkamatic, I'm hangin' tight like a fanatic \nYou trapped me once and I thought that \nYou might have it \nSo step down and lend me your ear \n'89 in my time! You, '90 is my year. \n\nYou're weakenin' fast, YO! and I can tell it \nYour body's gettin' hot, so, so I can smell it \nSo don't be mad and don't be sad \n'Cause the lyrics belong to ICE, You can call me Dad \nYou're pitchin' a fit, so step back and endure \nLet the witch doctor, Ice, do the dance to cure \nSo come up close and don't be square \nYou wanna battle me -- Anytime, anywhere \n\nYou thought that I was weak, Boy, you're dead wrong \nSo come on, everybody and sing this song \n\nSay -- Play that funky music Say, go white boy, go white boy go \nplay that funky music Go white boy, go white boy, go \nLay down and boogie and play that funky music till you die. \n\nPlay that funky music Come on, Come on, let me hear \nPlay that funky music white boy you say it, say it \nPlay that funky music A little louder now \nPlay that funky music, white boy Come on, Come on, Come on \nPlay that funky music \n")
    }

    #[test]
    fn s1c8() {
        let data = std::fs::read_to_string("8.txt")
            .unwrap()
            .lines()
            .map(|lines| hex::decode(lines).unwrap())
            .collect::<Vec<_>>();

        // Find the line with the lowest average hamming distance between chunks.
        let mut min_score = std::f32::MAX;
        let mut best_line = 0;

        for (idx, line) in data.iter().enumerate() {
            let mut score = 0f32;

            for (j, chunk_1) in line.chunks(16).enumerate() {
                for (i, chunk_2) in line.chunks(16).enumerate() {
                    if i == j {
                        continue;
                    }

                    score += hamming_distance(chunk_1, chunk_2) as f32 / 16.0;
                }
            }

            if score < min_score {
                min_score = score;
                best_line = idx;
            }
        }

        assert_eq!(best_line, 132);
    }
}
