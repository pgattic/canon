use dirs::home_dir;
use std::path::PathBuf;
use dioxus::prelude::*;

#[component]
pub fn SearchResultsView(query: String) -> Element {
    let canon_path: PathBuf = home_dir().unwrap().join(".canon").join("texts");

    // Time the search function
    let render_start = std::time::Instant::now();
    let search_results = if query.len() > 0 {
        libcanon::search::search(&canon_path, &query)
    } else {
        Err("Search something fun!")
    };
    let duration = render_start.elapsed();

    match search_results {
        Ok(matches) => {
            rsx! {
                p { "{matches.len()} results found in {duration.as_secs_f32()} seconds" }
                for m in matches {
                    div {
                        p {
                            b {
                                "{m.reference} "
                            }
                            span {
                                r#style: "
                                    -webkit-user-select: text;
                                    -ms-user-select: text;
                                    user-select: text;
                                ",
                                "{m.text}"
                            }
                        }
                    }
                }
            }
        }
        Err(problem) => {
            rsx! {
                p { "Error: {problem}" }
            }
        }
    }
}

