-- Active: 1765765503303@@127.0.0.1@3306@beaverdemo_usercenter
/*
CREATE DATABASE IF NOT EXISTS beaverdemo_usercenter;
USE beaverdemo_usercenter;

CREATE TABLE IF NOT EXISTS user;

/*
* delete user for debug
*/
DELETE FROM user WHERE email = 'xiahuyun043@126.com';

CREATE TABLE user (
    id bigint NOT NULL AUTO_INCREMENT,
    name varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The user name',
    nickname varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The user nickname',
    password varchar(255) NOT NULL DEFAULT '' COMMENT 'The user password',
    phone varchar(255) NOT NULL DEFAULT '' COMMENT 'The mobile phone number',
    email varchar(255) NOT NULL DEFAULT '' COMMENT 'The user email',
    gender char(10) NOT NULL DEFAULT 'unknown' COMMENT 'gender,male|female|unknown',
    create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    avatar varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    register_type varchar(255) NOT NULL DEFAULT '' COMMENT 'The user register type, email|phone',
    UNIQUE email_index (email),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'user table';

/*
* goctl model mysql ddl --src=./app/usercenter/model/usercenter.sql --dir ./app/usercenter/model/ --strict -c
*/
