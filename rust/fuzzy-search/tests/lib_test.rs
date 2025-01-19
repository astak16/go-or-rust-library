use fuzzy_search;

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_fuzzy_search() {
        let testData = vec![
            ("car", "cartwheel", true),
            ("cwhl", "cartwheel", true),
            ("cwheel", "cartwheel", true),
            ("cartwheel", "cartwheel", true),
            ("cwheeel", "cartwheel", false),
            ("lw", "cartwheel", false),
            ("语言", "php语言", true),
            ("hp语", "php语言", true),
            ("Py开发", "Python开发者", true),
            ("Py 开发", "Python开发者", false),
            ("爪哇进阶", "爪哇开发进阶", true),
            ("格式工具", "非常简单的格式化工具", true),
            ("正则", "学习正则表达式怎么学习", true),
            ("学习正则", "正则表达式怎么学习", false),
        ];

        for (needly, haystack, expected) in testData {
            assert_eq!(fuzzy_search::fuzzy_search(needly, haystack), expected);
        }
    }
}
