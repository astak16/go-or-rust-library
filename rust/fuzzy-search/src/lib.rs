pub fn fuzzy_search(needle: &str, haystack: &str) -> bool {
    let hlen = haystack.len();
    let nlen = needle.len();

    // 长度检查
    if nlen > hlen {
        return false;
    }
    if nlen == hlen {
        return needle == haystack;
    }

    let needle_chars: Vec<char> = needle.chars().collect();
    let haystack_chars: Vec<char> = haystack.chars().collect();

    let mut j = 0;
    'outer: for &nch in needle_chars.iter() {
        while j < haystack_chars.len() {
            if nch == haystack_chars[j] {
                j += 1;
                continue 'outer;
            }
            j += 1;
        }
        return false;
    }
    true
}
