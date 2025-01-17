use string_similarity;

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_compare_two_strings() {
        let test_cases = vec![
            ("french", "quebec", 0.0),
            ("france", "france", 1.0),
            ("fRaNce", "france", 0.2),
            ("healed", "sealed", 0.8),
            (
                "web applications",
                "applications of the web",
                0.7878787878787878,
            ),
            (
                "this will have a typo somewhere",
                "this will huve a typo somewhere",
                0.92,
            ),
            (
                "Olive-green table for sale, in extremely good condition.",
                "For sale: table in very good  condition, olive green in colour.",
                0.6060606060606061,
            ),
            (
                "Olive-green table for sale, in extremely good condition.",
                "For sale: green Subaru Impreza, 210,000 miles",
                0.2558139534883721,
            ),
            (
                "Olive-green table for sale, in extremely good condition.",
                "Wanted: mountain bike with at least 21 gears.",
                0.1411764705882353,
            ),
            (
                "this has one extra word",
                "this has one word",
                0.7741935483870968,
            ),
            ("a", "a", 1.0),
            ("a", "b", 0.0),
            ("", "", 1.0),
            ("a", "", 0.0),
            ("", "a", 0.0),
            ("apple event", "apple    event", 1.0),
            ("iphone", "iphone x", 0.9090909090909091),
        ];

        for (first, second, expected) in test_cases {
            let result = string_similarity::compare_two_strings(first, second);
            assert!(
                (result - expected).abs() < 1e-10,
                "Testing '{}' vs '{}': expected {}, got {}",
                first,
                second,
                expected,
                result
            );
        }
    }

    const BAD_ARGS_ERROR_MSG: &str =
        "Bad arguments: First argument should be a string, second should be an array of strings";

    #[test]
    // "one", ["two", "three"]
    fn accepts_string_and_array_returns_object() {
        let target_strings = vec![String::from("two"), String::from("three")];
        let output = string_similarity::find_best_match("one", &target_strings);
        assert!(output.is_ok());
    }

    #[test]
    // "hello", []
    fn throws_error_if_second_arg_empty_array() {
        let empty_strings: Vec<String> = vec![];
        let result = string_similarity::find_best_match("hello", &empty_strings);
        assert!(result.is_err());
        assert_eq!(result.unwrap_err(), BAD_ARGS_ERROR_MSG);
    }

    #[test]
    // "hello", ["something"，""]
    fn if_second_arg_contains_empty_string_returns_object() {
        let target_strings = vec![String::from("something"), String::from("")];
        let result = string_similarity::find_best_match("hello", &target_strings);
        assert!(result.is_ok());
    }

    #[test]
    fn test_find_best_match_rating() {
        let query = "healed";
        let target_strings = vec![
            String::from("mailed"),
            String::from("edward"),
            String::from("sealed"),
            String::from("theatre"),
        ];
        let matches = string_similarity::find_best_match(query, &target_strings).unwrap();
        // 验证返回的最佳匹配
        assert_eq!(matches.best_match_index, 2); // "sealed" 应该是最佳匹配
        assert_eq!(matches.best_match.target, "sealed");
        assert!((matches.best_match.rating - 0.8).abs() < 1e-10);

        let expected_ratings = vec![
            string_similarity::Rating {
                target: String::from("mailed"),
                rating: 0.4,
            },
            string_similarity::Rating {
                target: String::from("edward"),
                rating: 0.2,
            },
            string_similarity::Rating {
                target: String::from("sealed"),
                rating: 0.8,
            },
            string_similarity::Rating {
                target: String::from("theatre"),
                rating: 0.36363636363636365,
            },
        ];
        assert_eq!(matches.ratings.len(), expected_ratings.len());
        for (actual, expected) in matches.ratings.iter().zip(expected_ratings.iter()) {
            assert_eq!(actual.target, expected.target);
            assert!(
                (actual.rating - expected.rating).abs() < 1e-10,
                "For target '{}': expected rating {}, got {}",
                actual.target,
                expected.rating,
                actual.rating
            );
        }
    }
}
