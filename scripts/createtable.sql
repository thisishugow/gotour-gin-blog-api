create table blog_tag (
    id serial NOT NULL, 
    name varchar(100) default '',
    -- Common columns &start
    created_on numeric default 0,
    created_by varchar(100) default '',
    modified_on numeric default 0,
    modified_by varchar(100) default '',
    delete_on numeric default 0,
    is_del numeric default 0, 
    -- Common columns *end
    state numeric default 1, -- 0, disabled; 1, enabled.
    primary key(id) 
);

create table blog_article (
    id serial not null,
    title varchar(100) default '',
    desc varchar(100) default '',
    cover_image_url default '', 
    content text default '',
    -- Common columns &start
    created_on numeric default 0,
    created_by varchar(100) default '',
    modified_on numeric default 0,
    modified_by varchar(100) default '',
    delete_on numeric default 0,
    is_del numeric default 0, 
    -- Common columns *end
    state numeric default 1,
    primary key(id) 
);

create table blog_article_tag (
    id serial not null,
    article_id numeric not null default 0,
    tag_id numeric not null default 0,
    -- Common columns &start
    created_on numeric default 0,
    created_by varchar(100) default '',
    modified_on numeric default 0,
    modified_by varchar(100) default '',
    delete_on numeric default 0,
    is_del numeric default 0, 
    -- Common columns *end
    state numeric default 1,
    primary key(id) 
);