CREATE TABLE friend (
    id bigint NOT NULL AUTO_INCREMENT,
    friendId varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The friend id',
    friendshipId varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The friendship id',
    sendUserId varchar(255) NOT NULL DEFAULT '' COMMENT 'The send user id',
    revUserId varchar(255) NOT NULL DEFAULT '' COMMENT 'The receive user id',
    sendUserNotice varchar(255) NOT NULL DEFAULT '' COMMENT 'The send user notice',
    revUserNotice varchar(255) NOT NULL DEFAULT '' COMMENT 'The receive user notice',
    create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    source varchar(255) NOT NULL DEFAULT '' COMMENT 'The friend source',
    isDeleted varchar(255) NOT NULL DEFAULT '' COMMENT 'The friend is deleted',
    UNIQUE friend_id_index (friendId),
    UNIQUE friendship_id_index (friendshipId),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'friend table';

CREATE TABLE friendverify (
    id bigint NOT NULL AUTO_INCREMENT,
    verifyId varchar(255) NOT NULL DEFAULT '' COMMENT 'The verify id',
    sendUserId varchar(255) NOT NULL DEFAULT '' COMMENT 'The send user id',
    revUserId varchar(255) NOT NULL DEFAULT '' COMMENT 'The receive user id',
    sendStatus varchar(255) NOT NULL DEFAULT '' COMMENT 'The send user status, 发起方状态 0:未处理 1:已通过 2:已拒绝 3: 忽略 4:删除',
    revStatus varchar(255) NOT NULL DEFAULT '' COMMENT 'The receive user status, 接收方状态 0:未处理 1:已通过 2:已拒绝 3: 忽略 4:删除',
    createat timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updateat timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    source varchar(255) NOT NULL DEFAULT '' COMMENT 'The friend source',
    message varchar(255) NOT NULL DEFAULT '' COMMENT 'The friend verify message',
    version int NOT NULL DEFAULT 0 COMMENT 'The friend verify version',
    UNIQUE verify_id_index (verifyId),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'friend verify table';
