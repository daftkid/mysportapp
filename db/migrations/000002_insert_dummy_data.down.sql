DELETE FROM targets WHERE
        comment = 'test target for user 1';

DELETE FROM users WHERE
        name = 'test-user-1';

DELETE FROM avatars WHERE
    alt = 'alt text' AND path = '';

DELETE FROM credentials WHERE
    email = 'test@test.com' AND hpassword = 'testhashedpassword';
