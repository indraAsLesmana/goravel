CREATE TABLE attachments (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  created_at DATETIME(3) NOT NULL,
  updated_at DATETIME(3) NOT NULL,
  blogpost_id BIGINT(20) UNSIGNED NOT NULL,
  title VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL,
  deleted_at datetime(3),
  PRIMARY KEY (id),
  KEY idx_attachment_created_at (created_at),
  KEY idx_attachment_updated_at (updated_at),
  KEY fk_blogpost_id (blogpost_id),
  CONSTRAINT fk_blogpost
    FOREIGN KEY (blogpost_id)
    REFERENCES blogposts(id)
    ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
