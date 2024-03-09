-- +goose Up
CREATE TABLE votes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    lang1_id INTEGER NOT NULL,
    lang2_id INTEGER NOT NULL,
    choice INTEGER,

    constraint fk_lang1_id FOREIGN KEY(lang1_id) REFERENCES languages(id),
    constraint fk_lang2_id FOREIGN KEY(lang2_id) REFERENCES languages(id),
    constraint unique_lang_pair UNIQUE ( lang1_id, lang2_id ),
    constraint lang_alpha_order CHECK ( lang1_id < lang2_id )
);

-- +goose Down
DROP TABLE votes;
