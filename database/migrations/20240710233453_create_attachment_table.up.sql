CREATE TABLE attachments (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    created_at datetime(3) NOT NULL,
    updated_at datetime(3) NOT NULL,
    blogpost_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    KEY idx_attachment_created_at (created_at),
    KEY idx_attachment_updated_at (updated_at),
    FOREIGN KEY (blogpost_id) REFERENCES blogposts(id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;