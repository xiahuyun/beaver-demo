CREATE TABLE notificationEvent (
    id bigint NOT NULL AUTO_INCREMENT,
    eventId varchar(255) NOT NULL DEFAULT '' COMMENT 'The event id',
    eventType varchar(255) NOT NULL DEFAULT '' COMMENT 'The event type',
    category varchar(255) NOT NULL DEFAULT '' COMMENT 'The event category',
    version int NOT NULL DEFAULT 0 COMMENT 'The event version',
    fromUserId varchar(255) DEFAULT '' COMMENT 'The from user id',
    targetId varchar(255) DEFAULT '' COMMENT 'The target user id',
    targetType varchar(255) DEFAULT '' COMMENT 'The target user type',
    payload json COMMENT 'The event payload',
    priority int NOT NULL DEFAULT 5 COMMENT 'The event priority',
    status int NOT NULL DEFAULT 1 COMMENT 'The event status',
    dedupHash varchar(255) COMMENT 'The event dedup hash',
    create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE event_id_index (eventId),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'notification event table';

CREATE TABLE notificationInbox (
    id bigint NOT NULL AUTO_INCREMENT,
    userId varchar(64) NOT NULL COMMENT '收件人',
    eventId VARCHAR(64) NOT NULL,
    eventType VARCHAR(32) NOT NULL,
    category VARCHAR(32) NOT NULL,
    version BIGINT NOT NULL DEFAULT 0,
    isRead TINYINT(1) NOT NULL DEFAULT 0,
    readAt TIMESTAMP NULL,
    status TINYINT NOT NULL DEFAULT 1,
    isDeleted TINYINT(1) NOT NULL DEFAULT 0,
    silent TINYINT(1) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    INDEX idx_inbox_user_event (userId, eventId),
    INDEX idx_event_type (eventType),
    INDEX idx_category (category),
    INDEX idx_version (version),
    INDEX idx_is_read (isRead),
    INDEX idx_status (status),
    INDEX idx_is_deleted (isDeleted),
    INDEX idx_deleted_at (deleted_at),

    PRIMARY KEY (id)
) ENGINE=InnoDB COLLATE utf8mb4_unicode_ci COMMENT='通知收件箱表';
