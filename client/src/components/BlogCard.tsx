interface BlogCardType {
    // avatar:
    author: string;
    title: string;
    content: string;
}

export default function BlogCard(props: BlogCardType) {
    return (
        <section>
            <header>
                <h1>{props.title}</h1>
                <p>{props.author}</p>
            </header>
            <article>{props.content}</article>
        </section>
    );
}