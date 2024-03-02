-- +goose Up
CREATE TABLE languages (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT UNIQUE NOT NULL,
  logo_slug TEXT UNIQUE NOT NULL
);

INSERT INTO languages (name, logo_slug)
VALUES ('JavaScript', 'javascript'),
       ('Python',     'python'),
       ('Ruby',       'ruby'),
       ('Java',       'java'),
       ('PHP',        'php'),
       ('C#',         'csharp'),
       ('Go',         'go'),
       ('TypeScript', 'typescript'),
       ('C++',        'cplusplus'),
       ('Perl',       'perl'),
       ('Haskell',    'haskell'),
       ('Julia',      'julia'),
       ('OCaml',      'ocaml'),
       ('C',          'c'),
       ('Rust',       'rust'),
       ('Elixir',     'elixir'),
       ('Lua',        'lua'),
       ('Clojure',    'clojure'),
       ('Kotlin',     'kotlin'),
       ('Swift',      'swift'),
       ('R',          'r');


-- +goose Down
DROP TABLE languages;
