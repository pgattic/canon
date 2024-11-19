#![allow(non_snake_case)]

use dioxus::desktop::WindowBuilder;
use dirs::home_dir;
use std::path::PathBuf;
use dioxus::prelude::*;
//use dioxus_logger::tracing::{info, Level};
use libcanon::reference::Reference;
use libcanon::*;
use dioxus::desktop::tao::dpi::PhysicalPosition;

fn main() {
    let cfg = dioxus::desktop::Config::new()
        .with_custom_head(r#"
<style>
  body {
    margin: 0;
    overflow-x: hidden;
    /*padding: 0;*/
  }

  @media (prefers-color-scheme: dark) {
    body {
      background-color: #222222;
      color: #dddddd;
    }
  }

  * {
    -webkit-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }
</style>"#.to_string())
        .with_window(
            WindowBuilder::new()
                .with_title("Canon")
                //.with_decorations(false)
                .with_position(PhysicalPosition::new(0, 0)));

    LaunchBuilder::desktop().with_cfg(cfg).launch(Home);
}

#[component]
fn ScriptureView(query: String, show_numbers: bool) -> Element {
    let canon_path: PathBuf = home_dir().unwrap().join(".canon").join("texts");
    //let mut selected_text = use_signal(|| String::from(""));


    // Parse the reference
    let reference = Reference::from_str(&query).unwrap();
    let result = citation::cite(&canon_path, &reference);
    match result {
        Ok(citation) => {
            rsx! {
                //p { "Selected text: {selected_text}" }
                //h1 { "{citation.book_name}" }
                for ch in citation.chapters.iter() {
                    if ch.entire_chapter {
                        h2 {
                            r#style: "
                                text-align: center;
                                font-weight: normal;
                            ",
                            "CHAPTER {ch.path.file_name().unwrap().to_str().unwrap()}"
                        }
                    }
                    div {
                        //onselect: move |e| {
                        //    selected_text.set(e.)
                        //}
                        for v in &ch.verses {
                            p {
                                //style: "text-align: justify;",
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

#[component]
fn Home() -> Element {
    let mut query = use_signal(|| String::from("1ne3"));
    let mut show_numbers = use_signal(|| true);

    rsx! {
        nav {
            r#style: "
                position: sticky;
                top: 0;
                width: 100%;
                padding: 8px;
                background: #444444;
            ",
            input {
                //r#style: "
                //    background: #f00;
                //",
                r#type: "text",
                value: "{query}",
                oninput: move |e| {query.set(e.value());},
            }
            button {
                onclick: move |_| {show_numbers.toggle();},
                "Show/hide numbers"
            }
        }
        div {
            r#style: "
                margin: 0 auto;
                max-width: 800px;
                padding: 0 32px 16px;
            ",
            ScriptureView { query: query, show_numbers: show_numbers() }
        }
    }
}

