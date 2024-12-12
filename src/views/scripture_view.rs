use libcanon::reference::Reference;
use libcanon::citation;
use dioxus::prelude::*;
use crate::constants::canon_home;

#[component]
pub fn ScriptureView(query: String, show_numbers: bool) -> Element {

    // Parse the reference
    let reference = Reference::from_str(&query).unwrap();
    let result = citation::cite(&canon_home(), &reference);
    match result {
        Ok(citation) => {
            rsx! {
                for ch in citation.chapters.iter() {
                    if ch.entire_chapter {
                        h2 {
                            text_align: "center",
                            font_weight: "normal",
                            "CHAPTER {ch.path.file_name().unwrap().to_str().unwrap()}"
                        }
                    }
                    div {
                        for v in &ch.verses {
                            p {
                                if show_numbers {
                                    b {"{v.verse} "} // Verse number
                                }
                                span { // Verse content
                                    r#style: "
                                        -webkit-user-select: text;
                                        -ms-user-select: text;
                                        user-select: text;
                                    ",
                                    "{v.content}"
                                }
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

