use std::collections::HashMap;

pub fn compare_two_strings(first: &str, second: &str) -> f64 {
    let first = remove_spaces_fast(first);
    let second = remove_spaces_fast(second);
    if first == second {
        return 1 as f64;
    };

    if first.chars().count() < 2 || second.chars().count() < 2 {
        return 0 as f64;
    };

    let mut map = HashMap::new();
    let first_chars: Vec<char> = first.chars().collect();
    for window in first_chars.windows(2) {
        let bigram = String::from_iter(window);
        *map.entry(bigram).or_insert(0) += 1;
    }

    let mut intersection_size = 0;
    let second_chars: Vec<char> = second.chars().collect();
    for window in second_chars.windows(2) {
        let bigram = String::from_iter(window);
        if let Some(count) = map.get_mut(&bigram) {
            if *count > 0 {
                *count -= 1;
                intersection_size += 1;
            }
        }
    }
    (2.0 * intersection_size as f64) / (first.chars().count() + second.chars().count() - 2) as f64
}

#[derive(Debug)]
pub struct Rating {
    pub target: String,
    pub rating: f64,
}

#[derive(Debug)]
pub struct MatchResult {
    pub ratings: Vec<Rating>,
    pub best_match: Rating,
    pub best_match_index: usize,
}

pub fn find_best_match(
    main_string: &str,
    target_strings: &[String],
) -> Result<MatchResult, &'static str> {
    if !are_args_valid(main_string, target_strings) {
        return Err("Bad arguments: First argument should be a string, second should be an array of strings");
    }

    let mut ratings = Vec::new();
    let mut best_match_index = 0;

    for (i, current_target_string) in target_strings.iter().enumerate() {
        let current_rating = compare_two_strings(main_string, current_target_string);
        ratings.push(Rating {
            target: current_target_string.clone(),
            rating: current_rating,
        });

        if current_rating > ratings[best_match_index].rating {
            best_match_index = i;
        }
    }

    let best_match = Rating {
        target: ratings[best_match_index].target.clone(),
        rating: ratings[best_match_index].rating,
    };

    Ok(MatchResult {
        ratings,
        best_match,
        best_match_index,
    })
}

fn are_args_valid(main_string: &str, target_strings: &[String]) -> bool {
    !target_strings.is_empty()
}

fn remove_spaces_fast(s: &str) -> String {
    let mut result = String::with_capacity(s.len());
    s.chars()
        .filter(|c| !c.is_whitespace())
        .for_each(|c| result.push(c));
    result
}
