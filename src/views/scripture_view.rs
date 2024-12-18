use libcanon::citation::Citation;
use dioxus::prelude::*;

#[component]
pub fn ScriptureView(content: Citation, show_numbers: bool) -> Element {
    rsx! {
        div {
            line_height: "1.5",
            for ch in content.chapters {
                div {
                    if show_numbers {
                        h2 {
                            text_align: "center",
                            font_weight: "normal",
                            "CHAPTER {ch.path.file_name().unwrap().to_str().unwrap()}"
                        }
                    }
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
}

