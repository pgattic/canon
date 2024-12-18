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
                padding: "0 16px",
                box_shadow: "0 2px 5px rgba(0,0,0,0.2)",
                background_color: "light-dark(#eeeeee, #333333)",
                height: "58px",
                display: "flex",
                gap: "8px",
                align_items: "center",
                {props.bar}
            }
            div {
                margin: "0 auto",
                max_width: "800px",
                padding: "8px 32px",
                overflow: "scroll",
                {props.content}
            }
        }
    }
}

