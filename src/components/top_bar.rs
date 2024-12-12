use dioxus::prelude::*;

#[derive(PartialEq, Clone, Props)]
pub struct TopBarProps {
    bar: Element,
    content: Element
}

#[component]
pub fn TopBar(props: TopBarProps) -> Element {
    rsx! {
        div {
            margin: "0",
            display: "grid",
            grid_template_rows: "auto 1fr",
            min_height: "0",
            width: "100vw",
            height: "100%",
            color: "light-dark(black, white)",
            background_color: "light-dark(#ffffff, #222222)",
            nav {
                padding: "8px",
                box_shadow: "0 2px 5px rgba(0,0,0,0.2)",
                background_color: "light-dark(#eeeeee, #333333)",
                {props.bar}
            }
            div {
                margin: "0 auto",
                max_width: "800px",
                padding: "0 32px 16px",
                overflow: "scroll",
                {props.content}
            }
        }
    }
}

