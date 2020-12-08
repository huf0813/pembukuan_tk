package sqlite

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

type ConnSqlite struct {
}

type ConnSqliteInterface interface {
	SqliteConnInit() *sql.DB
	AutoDropDB() error
}

func (sqliteConn *ConnSqlite) SqliteConnInit() *sql.DB {
	result, err := sql.Open("sqlite3", "./db/sqlite/pembukuan_db")
	if err != nil {
		panic(err)
	}
	return result
}

func (sqliteConn *ConnSqlite) AutoDropDB() error {
	conn := sqliteConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return errors.New("connection failed to db")
	}

	// customers table
	if _, err := conn.Exec("drop table if exists customers"); err != nil {
		return err
	} else {
		if _, err := conn.Exec("CREATE TABLE IF NOT EXISTS `customers` (`id` INTEGER NOT NULL, `name` TEXT NOT NULL,`phone` TEXT NOT NULL UNIQUE,`email` TEXT UNIQUE,`address` TEXT NOT NULL,`created_at` TEXT NOT NULL,`updated_at` TEXT NOT NULL,`deleted_at` TEXT, PRIMARY KEY(`id`));"); err != nil {
			return err
		}
	}

	// user_types table
	if _, err := conn.Exec("drop table if exists user_types"); err != nil {
		return err
	} else {
		if _, err := conn.Exec("CREATE TABLE IF NOT EXISTS `user_types` (`id` INTEGER NOT NULL, `name` TEXT NOT NULL, `created_at` TEXT NOT NULL, `updated_at` TEXT NOT NULL, `deleted_at` TEXT, PRIMARY KEY(`id`));"); err != nil {
			return err
		}
	}

	// invoices table
	if _, err := conn.Exec("drop table if exists invoices"); err != nil {
		return err
	} else {
		if _, err := conn.Exec("CREATE TABLE IF NOT EXISTS `invoices` (`id` INTEGER NOT NULL,`customer_id` INTEGER NOT NULL,`user_id` INTEGER NOT NULL,`created_at` TEXT NOT NULL,`updated_at` TEXT NOT NULL,`deleted_at` TEXT,FOREIGN KEY(`user_id`) REFERENCES `users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,FOREIGN KEY(`customer_id`) REFERENCES `customers`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,PRIMARY KEY(`id`));"); err != nil {
			return err
		}
	}

	// products table
	if _, err := conn.Exec("drop table if exists products"); err != nil {
		return err
	} else {
		if _, err := conn.Exec("CREATE TABLE IF NOT EXISTS `products` (`id` INTEGER NOT NULL,`name` TEXT NOT NULL,`price` TEXT NOT NULL,`created_at` TEXT NOT NULL,`updated_at` TEXT NOT NULL,`deleted_at` TEXT,PRIMARY KEY(`id`));"); err != nil {
			return err
		}
	}

	// product_decreases table
	if _, err := conn.Exec("drop table if exists product_decreases"); err != nil {
		return err
	} else {
		if _, err := conn.Exec("CREATE TABLE IF NOT EXISTS `product_decreases` (`id` INTEGER NOT NULL,`product_id` INTEGER NOT NULL,`quantity` INTEGER NOT NULL,`invoice_id` INTEGER NOT NULL,FOREIGN KEY(`product_id`) REFERENCES `products`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,FOREIGN KEY(`invoice_id`) REFERENCES `invoices`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,PRIMARY KEY(`id`));"); err != nil {
			return err
		}
	}

	// product_increases table
	if _, err := conn.Exec("drop table if exists product_increases"); err != nil {
		return err
	} else {
		if _, err := conn.Exec("CREATE TABLE IF NOT EXISTS `product_increases` (`id` INTEGER NOT NULL,`product_id` INTEGER NOT NULL,`quantity` INTEGER NOT NULL,`user_id` INTEGER NOT NULL,`created_at` TEXT NOT NULL,`updated_at` TEXT NOT NULL,`deleted_at` TEXT,FOREIGN KEY(`user_id`) REFERENCES `users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,FOREIGN KEY(`product_id`) REFERENCES `products`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,PRIMARY KEY(`id`));"); err != nil {
			return err
		}
	}

	// users table
	if _, err := conn.Exec("drop table if exists users"); err != nil {
		return err
	} else {
		if _, err := conn.Exec("CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER NOT NULL,`user_type_id` INTEGER NOT NULL,`username` TEXT NOT NULL UNIQUE,`password` TEXT NOT NULL,`created_at` TEXT NOT NULL,`updated_at`TEXT NOT NULL,`deleted_at` TEXT,FOREIGN KEY(`user_type_id`) REFERENCES `user_types`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,PRIMARY KEY(`id`));"); err != nil {
			return err
		} else {

		}
	}

	// relationship
	if _, err := conn.Exec("CREATE INDEX IF NOT EXISTS `users_fkIdx_68` ON `users` (`user_type_id`);CREATE INDEX IF NOT EXISTS `product_increases_fkIdx_93` ON `product_increases` (`user_id`);CREATE INDEX IF NOT EXISTS `product_increases_fkIdx_88` ON `product_increases` (`product_id`);CREATE INDEX IF NOT EXISTS `product_decreases_fkIdx_95` ON `product_decreases` (`product_id`);CREATE INDEX IF NOT EXISTS `product_decreases_fkIdx_109` ON `product_decreases` (`invoice_id`);CREATE INDEX IF NOT EXISTS `invoices_fkIdx_96` ON `invoices` (`user_id`);CREATE INDEX IF NOT EXISTS `invoices_fkIdx_81` ON `invoices` (`customer_id`);"); err != nil {
		return err
	}
	return nil
}
