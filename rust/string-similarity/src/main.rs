use std::collections::HashMap;
fn main() {
    let mut map2 = HashMap::new();
    let key = 'a';

    // 1. 使用 entry
    let entry = map2.entry(key);

    // 2. 使用 or_insert 设置默认值
    let count = entry.or_insert(0);
}
