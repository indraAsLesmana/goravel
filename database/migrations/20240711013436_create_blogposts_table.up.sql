CREATE TABLE blogposts (
  id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  title varchar(255) NOT NULL,
  url varchar(255),
  deleted_at datetime(3),
  PRIMARY KEY (id),
  KEY idx_blogpost_created_at (created_at),
  KEY idx_blogpost_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;