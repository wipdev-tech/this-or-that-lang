-- +goose Up
CREATE TABLE languages (
  id INT PRIMARY KEY,
  name TEXT NOT NULL,
  logo_slug TEXT NOT NULL
);

INSERT INTO languages (id, name, logo_slug)
VALUES (1,  'JavaScript', 'javascript'),
       (2,  'Python',     'python'),
       (3,  'Ruby',       'ruby'),
       (4,  'Java',       'java'),
       (5,  'PHP',        'php'),
       (6,  'C#',         'csharp'),
       (7,  'Go',         'go'),
       (8,  'TypeScript', 'typescript'),
       (9,  'C++',        'cplusplus'),
       (10, 'Perl',       'perl');


-- +goose Down
DROP TABLE languages;
