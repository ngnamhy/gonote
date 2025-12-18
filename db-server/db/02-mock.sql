-- ===============================
-- Mock data for Reddit-style schema
-- ===============================

-- ===== USERS =====
INSERT INTO users (id, username, email, password) VALUES
(gen_random_uuid(), 'alice', 'alice@example.com', 'hashed_password_1'),
(gen_random_uuid(), 'bob',   'bob@example.com',   'hashed_password_2'),
(gen_random_uuid(), 'carol', 'carol@example.com', 'hashed_password_3');

-- ===== POSTS =====
INSERT INTO posts (id, user_id, title, body) VALUES
(
    gen_random_uuid(),
    (SELECT id FROM users WHERE username = 'alice'),
    'Go + Postgres + Docker',
    'How do you structure a clean backend with Go and Postgres?'
),
(
    gen_random_uuid(),
    (SELECT id FROM users WHERE username = 'bob'),
    'Learning backend seriously',
    'I want to build something like Reddit or Notion. Any advice?'
);

-- ===== TAGS =====
INSERT INTO tags (id, name) VALUES
(gen_random_uuid(), 'golang'),
(gen_random_uuid(), 'backend'),
(gen_random_uuid(), 'postgres'),
(gen_random_uuid(), 'docker');

-- ===== POST_TAGS =====
INSERT INTO post_tags (post_id, tag_id)
SELECT p.id, t.id
FROM posts p, tags t
WHERE p.title = 'Go + Postgres + Docker'
  AND t.name IN ('golang', 'backend', 'postgres', 'docker');

-- ===== COMMENTS (level 1) =====
INSERT INTO comments (id, post_id, user_id, body) VALUES
(
    gen_random_uuid(),
    (SELECT id FROM posts WHERE title = 'Go + Postgres + Docker'),
    (SELECT id FROM users WHERE username = 'bob'),
    'You should start with a clear schema and avoid AutoMigrate in production.'
),
(
    gen_random_uuid(),
    (SELECT id FROM posts WHERE title = 'Go + Postgres + Docker'),
    (SELECT id FROM users WHERE username = 'carol'),
    'Docker + SQL migrations will save you a lot of pain later.'
);

-- ===== COMMENTS (reply) =====
INSERT INTO comments (post_id, user_id, parent_id, body) VALUES
(
    (SELECT id FROM posts WHERE title = 'Go + Postgres + Docker'),
    (SELECT id FROM users WHERE username = 'alice'),
    (SELECT id FROM comments WHERE body LIKE 'You should start with%'),
    'Totally agree, that is what I am trying to do now.'
);

-- ===== VOTES =====
-- Upvote post
INSERT INTO votes (user_id, target_type, target_id, value) VALUES
(
    (SELECT id FROM users WHERE username = 'bob'),
    'post',
    (SELECT id FROM posts WHERE title = 'Go + Postgres + Docker'),
    1
),
(
    (SELECT id FROM users WHERE username = 'carol'),
    'post',
    (SELECT id FROM posts WHERE title = 'Go + Postgres + Docker'),
    1
);

-- Downvote a comment
INSERT INTO votes (user_id, target_type, target_id, value) VALUES
(
    (SELECT id FROM users WHERE username = 'alice'),
    'comment',
    (SELECT id FROM comments WHERE body LIKE 'Docker + SQL migrations%'),
    -1
);

-- ===== UPDATE SCORE (simulate app logic) =====
UPDATE posts
SET score = (
    SELECT COALESCE(SUM(value), 0)
    FROM votes
    WHERE target_type = 'post'
      AND target_id = posts.id
);

UPDATE comments
SET score = (
    SELECT COALESCE(SUM(value), 0)
    FROM votes
    WHERE target_type = 'comment'
      AND target_id = comments.id
);
