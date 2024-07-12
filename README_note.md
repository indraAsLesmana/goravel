$ go run . artisan migrate:rollback
Migration rollback success


go run . artisan migrate
Migration success


go run . artisan migrate:fresh --seed --seeder=BlogpostSeeder
2024/07/09 23:42:04 /go/pkg/mod/github.com/goravel/framework@v1.14.1/database/gorm/query.go:107
[33.587ms] [rows:1] INSERT INTO `blogposts` (`created_at`,`updated_at`,`title`,`attach_file`,`url`,`deleted_at`) VALUES ('2024-07-09 23:42:04.848','2024-07-09 23:42:04.848','Millioner 2024','','',NULL) RETURNING `id`
Database seeding completed successfully.
Migration fresh success


go run . artisan migrate:fresh
Migration fresh success


go run . artisan make:migration create_attachments_table


// rollback when migrating 2 table
go run . artisan migrate:rollback --step=2