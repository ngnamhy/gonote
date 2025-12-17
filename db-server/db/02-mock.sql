-- =========================
-- USERS
-- =========================
INSERT INTO users (id, username, email, password) VALUES
('11111111-1111-1111-1111-111111111111', 'alice', 'alice@example.com', 'hashed_pw_1'),
('22222222-2222-2222-2222-222222222222', 'bob', 'bob@example.com', 'hashed_pw_2'),
('33333333-3333-3333-3333-333333333333', 'charlie', 'charlie@example.com', 'hashed_pw_3'),
('44444444-4444-4444-4444-444444444444', 'david', 'david@example.com', 'hashed_pw_4');

-- =========================
-- TAGS
-- =========================
INSERT INTO tags (id, name) VALUES
('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'golang'),
('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'postgresql'),
('cccccccc-cccc-cccc-cccc-cccccccccccc', 'backend'),
('dddddddd-dddd-dddd-dddd-dddddddddddd', 'api'),
('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'gin');

-- =========================
-- POSTS
-- =========================
INSERT INTO posts (id, user_id, title, body, view_count) VALUES
('p1111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111',
 'How to use Gin framework?', 'I am new to Gin. How do I create a REST API?', 120),

('p2222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222',
 'PostgreSQL foreign key question', 'How does ON DELETE CASCADE work?', 85),

('p3333333-3333-3333-3333-333333333333', '33333333-3333-3333-3333-333333333333',
 'Best practices for REST API', 'What are the best REST API practices?', 200);

-- =========================
-- POST TAGS
-- =========================
INSERT INTO post_tags (post_id, tag_id) VALUES
('p1111111-1111-1111-1111-111111111111', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'),
('p1111111-1111-1111-1111-111111111111', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee'),
('p1111111-1111-1111-1111-111111111111', 'dddddddd-dddd-dddd-dddd-dddddddddddd'),
('p2222222-2222-2222-2222-222222222222', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'),
('p2222222-2222-2222-2222-222222222222', 'cccccccc-cccc-cccc-cccc-cccccccccccc'),
('p3333333-3333-3333-3333-333333333333', 'dddddddd-dddd-dddd-dddd-dddddddddddd'),
('p3333333-3333-3333-3333-333333333333', 'cccccccc-cccc-cccc-cccc-cccccccccccc');

-- =========================
-- ANSWERS
-- =========================
INSERT INTO answers (id, post_id, user_id, body) VALUES
('a1111111-1111-1111-1111-111111111111',
 'p1111111-1111-1111-1111-111111111111',
 '22222222-2222-2222-2222-222222222222',
 'You should start by creating a gin.Engine and defining routes.'),

('a2222222-2222-2222-2222-222222222222',
 'p1111111-1111-1111-1111-111111111111',
 '33333333-3333-3333-3333-333333333333',
 'Gin is very fast and easy to use for REST APIs.');

-- =========================
-- COMMENTS
-- =========================
INSERT INTO comments (user_id, target_type, target_id, body) VALUES
('22222222-2222-2222-2222-222222222222', 'post',
 'p1111111-1111-1111-1111-111111111111',
 'Good question, I was wondering the same.'),

('11111111-1111-1111-1111-111111111111', 'answer',
 'a1111111-1111-1111-1111-111111111111',
 'This answer helped me a lot, thanks!');
