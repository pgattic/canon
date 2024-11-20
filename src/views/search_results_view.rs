use dirs::home_dir;
use std::path::PathBuf;
use dioxus::prelude::*;

#[component]
pub fn SearchResultsView(query: String) -> Element {
    let canon_path: PathBuf = home_dir().unwrap().join(".canon").join("texts");
    //let mut selected_text = use_signal(|| String::from(""));


    // Parse the reference
    let search_results = if query.len() > 0 {libcanon::search::search(&canon_path, &query)} else { Err("Search something fun!")};
    match search_results {
        Ok(matches) => {
            //matches[0].re
            rsx! {
                //p { "Selected text: {selected_text}" }
                //h1 { "{citation.book_name}" }
                for m in matches {
                    div {
                        //onselect: move |e| {
                        //    selected_text.set(e.)
                        //}
                        p {
                            //style: "text-align: justify;",
                            b {
                                "{m.reference} "
                            }
                            span { // Verse content
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

