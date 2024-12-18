use crate::canon_home;
use dioxus::prelude::*;
use crate::views::ScriptureView;
use crate::components::TopBar;

#[component]
pub fn Search() -> Element {
    let mut search_input = use_signal(|| String::new());
    let mut search_duration = use_signal(|| 0.0);
    let mut search_results: Signal<Result<Vec<(libcanon::reference::Reference, libcanon::citation::Citation)>, &str>> = use_signal(|| Err("Search is case-sensitive"));

    rsx! {
        TopBar {
            bar: rsx! {
                input {
                    r#type: "text",
                    padding: "4px",
                    oninput: move |e| {search_input.set(e.value());},
                }
                button {
                    onclick: move |_| {
                        let render_start = std::time::Instant::now();
                        if search_input().len() > 0 {
                            search_results.set(libcanon::search::search(&canon_home(), &search_input()))
                        } else {
                            search_results.set(Err("Search is case-sensitive"))
                        };
                        let duration = render_start.elapsed();
                        search_duration.set(duration.as_secs_f32());
                    },
                    "Search"
                }
            },
            content: match search_results() {
                Ok(matches) => {
                    rsx! {
                        p { "{matches.len()} results found in {search_duration} seconds" }
                        for (reference, citation) in matches {
                            p {"{reference}"},
                            ScriptureView { content: citation, show_numbers: true }
                        }
                    }
                }
                Err(problem) => {
                    rsx! {
                        p { "{problem}" }
                    }
                }
            }
        }
    }
}

