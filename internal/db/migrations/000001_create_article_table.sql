create table public.articles
(
    id         serial PRIMARY KEY,
    title      varchar(64) not null ,
    author     varchar(64) not null ,
    content    text not null ,
    created_at TIMESTAMP WITH TIME ZONE not null DEFAULT now()
);

comment on table public.articles is '投稿';
comment on column public.articles.id is '投稿ID';
comment on column public.articles.title is '投稿タイトル';
comment on column public.articles.author is '投稿作者';
comment on column public.articles.content is '投稿文';
comment on column public.articles.created_at is '作成日時';

create index articles_draft_index
    on public.articles (id, title, author, created_at);