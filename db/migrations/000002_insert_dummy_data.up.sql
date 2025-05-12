INSERT INTO avatars (path, alt) VALUES (
    '', 'alt text'
);

INSERT INTO credentials (email, hpassword) VALUES (
    'test@test.com', 'testhashedpassword'
);

INSERT INTO users (name, bio, avatar_id, creds_id) VALUES (
    'test-user-1', 'test user 1 bio', 1, 1
 );

INSERT INTO targets (numofdays, numofreps, comment, user_id) VALUES (
    30, 1000, 'test target for user 1', 1
);